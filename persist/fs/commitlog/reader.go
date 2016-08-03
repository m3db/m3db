// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package commitlog

import (
	"bufio"
	"encoding/binary"
	"errors"
	"hash/adler32"
	"io"
	"os"
	"time"

	"github.com/m3db/m3db/generated/proto/schema"
	"github.com/m3db/m3db/interfaces/m3db"
	"github.com/m3db/m3x/time"

	"github.com/golang/protobuf/proto"
)

var (
	errCommitLogReaderAlreadyOpen               = errors.New("commit log reader already open")
	errCommitLogReaderChunkSizeChecksumMismatch = errors.New("commit log reader encountered chunk size checksum mismatch")
	errCommitLogReaderChunkDataChecksumMismatch = errors.New("commit log reader encountered chunk data checksum mismatch")
	errCommitLogReaderChunkUnexpectedEnd        = errors.New("commit log reader encountered message in chunk with unexpected end")
	errCommitLogReaderMissingLogMetadata        = errors.New("commit log reader encountered message missing metadata")
)

type commitLogReader interface {
	io.Closer

	// Open opens the commit log for reading
	Open(filePath string) (time.Time, time.Duration, int, error)

	// Read returns the next key and data pair or error, will return io.EOF at end of volume
	Read() (m3db.CommitLogSeries, m3db.Datapoint, xtime.Unit, []byte, error)
}

type reader struct {
	opts           m3db.DatabaseOptions
	fd             *os.File
	buffer         *bufio.Reader
	chunkRemaining uint32
	sizeBuffer     []byte
	info           schema.CommitLogInfo
	log            schema.CommitLog
	metadata       schema.CommitLogMetadata
	metadataLookup map[uint64]*readerSeries
}

type readerSeries struct {
	idx   uint64
	id    string
	shard uint32
}

func (s readerSeries) UniqueIndex() uint64 {
	return s.idx
}

func (s readerSeries) ID() string {
	return s.id
}

func (s readerSeries) Shard() uint32 {
	return s.shard
}

func newCommitLogReader(opts m3db.DatabaseOptions) commitLogReader {
	return &reader{
		opts:           opts,
		buffer:         bufio.NewReaderSize(nil, bufferWriteSize),
		sizeBuffer:     make([]byte, binary.MaxVarintLen64),
		metadataLookup: make(map[uint64]*readerSeries),
	}
}

func (r *reader) Open(filePath string) (time.Time, time.Duration, int, error) {
	if r.fd != nil {
		return timeZero, 0, 0, errCommitLogReaderAlreadyOpen
	}

	var err error
	r.fd, err = os.Open(filePath)
	if err != nil {
		return timeZero, 0, 0, err
	}

	r.buffer.Reset(r.fd)
	r.info = schema.CommitLogInfo{}
	if err := r.read(&r.info); err != nil {
		r.Close()
		return timeZero, 0, 0, err
	}

	start := time.Unix(0, r.info.Start)
	duration := time.Duration(r.info.Duration)
	index := int(r.info.Index)
	return start, duration, index, nil
}

func (r *reader) Read() (
	series m3db.CommitLogSeries,
	datapoint m3db.Datapoint,
	unit xtime.Unit,
	annotation []byte,
	resultErr error,
) {
	if err := r.read(&r.log); err != nil {
		resultErr = err
		return
	}

	if len(r.log.Metadata) != 0 {
		if err := proto.Unmarshal(r.log.Metadata, &r.metadata); err != nil {
			resultErr = err
			return
		}
		r.metadataLookup[r.log.Idx] = &readerSeries{
			idx:   r.log.Idx,
			id:    r.metadata.Id,
			shard: r.metadata.Shard,
		}
	}

	metadata, ok := r.metadataLookup[r.log.Idx]
	if !ok {
		resultErr = errCommitLogReaderMissingLogMetadata
		return
	}

	series = metadata
	datapoint = m3db.Datapoint{
		Timestamp: time.Unix(0, r.log.Timestamp),
		Value:     r.log.Value,
	}
	unit = xtime.Unit(byte(r.log.Unit))
	annotation = r.log.Annotation
	return
}

func (r *reader) read(message proto.Message) error {
	if r.chunkRemaining == 0 {
		// Peek instead of read as we take reference to the underlying buffer
		// rather than have to copy from the buffer to our own buffer
		header, err := r.buffer.Peek(chunkReserveHeaderLen)
		if err != nil {
			return err
		}

		size := endianness.Uint32(header[:4])
		checksumSize := endianness.Uint32(header[4:8])
		checksumData := endianness.Uint32(header[8:12])

		// Verify size checksum
		if adler32.Checksum(header[:4]) != checksumSize {
			return errCommitLogReaderChunkSizeChecksumMismatch
		}

		// Discard the peeked header
		if _, err := r.buffer.Discard(chunkReserveHeaderLen); err != nil {
			return err
		}

		// Verify data checksum
		data, err := r.buffer.Peek(int(size))
		if err != nil {
			return err
		}

		if adler32.Checksum(data) != checksumData {
			return errCommitLogReaderChunkSizeChecksumMismatch
		}

		r.chunkRemaining = size
	}

	// Read size of message
	size, err := binary.ReadUvarint(r.buffer)
	if err != nil {
		return err
	}

	// Track consumed
	consumed := binary.PutUvarint(r.sizeBuffer, size)
	r.chunkRemaining -= uint32(consumed)

	// Read message
	if r.chunkRemaining < uint32(size) {
		return errCommitLogReaderChunkUnexpectedEnd
	}

	data, err := r.buffer.Peek(int(size))
	if err != nil {
		return err
	}
	if err := proto.Unmarshal(data, message); err != nil {
		return err
	}

	// Discard the peeked data
	if _, err := r.buffer.Discard(int(size)); err != nil {
		return err
	}

	// Track consumed
	r.chunkRemaining -= uint32(size)
	return nil
}

func (r *reader) Close() error {
	if r.fd == nil {
		return nil
	}

	if err := r.fd.Close(); err != nil {
		return err
	}

	r.fd = nil
	r.chunkRemaining = 0
	r.metadataLookup = make(map[uint64]*readerSeries)
	return nil
}
