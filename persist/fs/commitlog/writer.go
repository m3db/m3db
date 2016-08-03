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
	"github.com/m3db/m3db/persist/fs"
	"github.com/m3db/m3x/time"

	"github.com/golang/protobuf/proto"
)

const (
	// The length to reserve for a chunk header:
	// - size uint32
	// - checksumSize uint32
	// - checksumData uint32
	chunkReserveHeaderLen = 4 + 4 + 4

	// TODO: make configurable by DatabaseOptions once rebased with options from master
	bufferWriteSize = 65536
)

var (
	errCommitLogWriterAlreadyOpen                = errors.New("commit log writer already open")
	errCommitLogWriterFlushWithoutReservedLength = errors.New("commit log writer flushed without header reserve")

	endianness = binary.LittleEndian
)

type commitLogWriter interface {
	io.Closer

	// Open opens the commit log for writing data
	Open(start time.Time, duration time.Duration) error

	// IsOpen returns whether open or not
	IsOpen() bool

	// Write will write an entry in the commit log for a given series
	Write(series m3db.CommitLogSeries, datapoint m3db.Datapoint, unit xtime.Unit, annotation []byte) error

	// Flush will flush the contents to the disk, useful when first testing if first commit log is writable
	Flush() error
}

type flushFn func(err error)

type writer struct {
	filePathPrefix     string
	newFileMode        os.FileMode
	newDirectoryMode   os.FileMode
	nowFn              m3db.NowFn
	bitset             bitset
	start              time.Time
	duration           time.Duration
	chunkWriter        chunkWriter
	chunkReserveHeader []byte
	buffer             *bufio.Writer
	info               schema.CommitLogInfo
	log                schema.CommitLog
	metadata           schema.CommitLogMetadata
	sizeBuffer         []byte
	infoBuffer         *proto.Buffer
	metadataBuffer     *proto.Buffer
	logBuffer          *proto.Buffer
}

func newCommitLogWriter(
	flushFn flushFn,
	opts m3db.DatabaseOptions,
) commitLogWriter {
	return &writer{
		filePathPrefix:     opts.GetFilePathPrefix(),
		newFileMode:        opts.GetFileWriterOptions().GetNewFileMode(),
		newDirectoryMode:   opts.GetFileWriterOptions().GetNewDirectoryMode(),
		nowFn:              opts.GetNowFn(),
		chunkWriter:        chunkWriter{flushFn: flushFn},
		chunkReserveHeader: make([]byte, chunkReserveHeaderLen),
		buffer:             bufio.NewWriterSize(nil, bufferWriteSize),
		bitset:             newBitset(),
		sizeBuffer:         make([]byte, binary.MaxVarintLen64),
		infoBuffer:         proto.NewBuffer(nil),
		metadataBuffer:     proto.NewBuffer(nil),
		logBuffer:          proto.NewBuffer(nil),
	}
}

func (w *writer) Open(start time.Time, duration time.Duration) error {
	if w.IsOpen() {
		return errCommitLogWriterAlreadyOpen
	}

	commitLogsDir := fs.CommitLogsDirPath(w.filePathPrefix)
	if err := os.MkdirAll(commitLogsDir, w.newDirectoryMode); err != nil {
		return err
	}

	filePath, index := fs.NextCommitLogsFile(w.filePathPrefix, start)
	w.info = schema.CommitLogInfo{
		Start:    start.UnixNano(),
		Duration: int64(duration),
		Index:    int64(index),
	}
	w.infoBuffer.Reset()
	if err := w.infoBuffer.Marshal(&w.info); err != nil {
		return err
	}

	fd, err := fs.OpenWritable(filePath, w.newFileMode)
	if err != nil {
		return err
	}

	w.chunkWriter.fd = fd
	w.buffer.Reset(&w.chunkWriter)
	if err := w.write(w.infoBuffer.Bytes()); err != nil {
		w.Close()
		return err
	}

	w.start = start
	w.duration = duration
	return nil
}

func (w *writer) IsOpen() bool {
	return w.chunkWriter.fd != nil
}

func (w *writer) Write(series m3db.CommitLogSeries, datapoint m3db.Datapoint, unit xtime.Unit, annotation []byte) error {
	w.log = schema.CommitLog{}
	w.log.Created = w.nowFn().UnixNano()
	w.log.Idx = series.UniqueIndex()

	seen := w.bitset.has(w.log.Idx)
	if !seen {
		// If "idx" hasn't been written to commit log
		// yet we need to include series metadata
		w.metadata = schema.CommitLogMetadata{}
		w.metadata.Id = series.ID()
		w.metadata.Shard = series.Shard()

		w.metadataBuffer.Reset()
		if err := w.metadataBuffer.Marshal(&w.metadata); err != nil {
			return err
		}
		w.log.Metadata = w.metadataBuffer.Bytes()
	}

	w.log.Timestamp = datapoint.Timestamp.UnixNano()
	w.log.Value = datapoint.Value
	w.log.Unit = uint32(unit)
	w.log.Annotation = annotation

	w.logBuffer.Reset()
	if err := w.logBuffer.Marshal(&w.log); err != nil {
		return err
	}
	if err := w.write(w.logBuffer.Bytes()); err != nil {
		return err
	}

	if !seen {
		// Record we have seen this series
		w.bitset.set(w.log.Idx)
	}
	return nil
}

func (w *writer) Flush() error {
	return w.buffer.Flush()
}

func (w *writer) Close() error {
	if !w.IsOpen() {
		return nil
	}

	if err := w.chunkWriter.fd.Close(); err != nil {
		return err
	}

	w.chunkWriter.fd = nil
	w.chunkWriter.flushFn = nil
	w.bitset.clearAll()
	w.start = timeZero
	w.duration = 0
	return nil
}

func (w *writer) write(data []byte) error {
	if w.buffer.Buffered() == 0 {
		// Reserve bytes for checksums and size prepend
		if _, err := w.buffer.Write(w.chunkReserveHeader); err != nil {
			return err
		}
	}

	dataLen := len(data)
	sizeLen := binary.PutUvarint(w.sizeBuffer, uint64(dataLen))
	if dataLen+sizeLen > w.buffer.Available() {
		// Ensure to never write across write boundary
		if err := w.buffer.Flush(); err != nil {
			return err
		}
		// Reserve bytes for checksums and size prepend
		if _, err := w.buffer.Write(w.chunkReserveHeader); err != nil {
			return err
		}
	}
	if _, err := w.buffer.Write(w.sizeBuffer[:sizeLen]); err != nil {
		return err
	}
	_, err := w.buffer.Write(data)
	return err
}

type chunkWriter struct {
	fd      *os.File
	flushFn flushFn
}

func (w *chunkWriter) Write(p []byte) (int, error) {
	rawLen := len(p)
	if rawLen <= chunkReserveHeaderLen {
		return 0, errCommitLogWriterFlushWithoutReservedLength
	}

	size := rawLen - chunkReserveHeaderLen

	// Write size
	slice := p[:4]
	endianness.PutUint32(slice, uint32(size))

	// Calculate checksums
	checksumSize := adler32.Checksum(slice)
	checksumData := adler32.Checksum(p[12:])

	// Write checksums
	slice = p[4:8]
	endianness.PutUint32(slice, uint32(checksumSize))
	slice = p[8:12]
	endianness.PutUint32(slice, uint32(checksumData))

	// Write to file descriptor
	n, err := w.fd.Write(p)

	// Fire flush callback
	w.flushFn(err)
	return n, err
}
