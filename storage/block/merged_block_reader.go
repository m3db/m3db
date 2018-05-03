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

package block

import (
	"fmt"
	"sync"
	"time"

	"github.com/m3db/m3db/encoding"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3db/x/xio"
)

type dbMergedBlockReader struct {
	sync.RWMutex
	opts       Options
	blockStart time.Time
	blockEnd   time.Time
	streams    [2]xio.SegmentReader
	readers    [2]xio.SegmentReader
	merged     xio.Block
	encoder    encoding.Encoder
	err        error
}

func newDatabaseMergedBlockReader(
	blockStart time.Time,
	blockEnd time.Time,
	streamA, streamB xio.SegmentReader,
	opts Options,
) xio.Block {
	r := &dbMergedBlockReader{
		opts:       opts,
		blockStart: blockStart,
		blockEnd:   blockEnd,
	}
	r.streams[0] = streamA
	r.streams[1] = streamB
	r.readers[0] = streamA
	r.readers[1] = streamB
	return xio.Block{
		SegmentReader: r,
		Start:         blockStart,
		End:           blockEnd,
	}
}

func (r *dbMergedBlockReader) mergedReader() (xio.Block, error) {
	r.RLock()
	if !r.merged.IsEmpty() || r.err != nil {
		r.RUnlock()
		return r.merged, r.err
	}
	r.RUnlock()

	r.Lock()
	defer r.Unlock()

	if !r.merged.IsEmpty() || r.err != nil {
		return r.merged, r.err
	}

	multiIter := r.opts.MultiReaderIteratorPool().Get()
	multiIter.Reset(r.readers[:], r.blockStart, r.blockEnd)
	defer multiIter.Close()

	r.encoder = r.opts.EncoderPool().Get()
	r.encoder.Reset(r.blockStart, r.opts.DatabaseBlockAllocSize())

	for multiIter.Next() {
		dp, unit, annotation := multiIter.Current()
		err := r.encoder.Encode(dp, unit, annotation)
		if err != nil {
			r.encoder.Close()
			r.err = err
			return xio.EmptyBlock, err
		}
	}
	if err := multiIter.Err(); err != nil {
		r.encoder.Close()
		r.err = err
		return xio.EmptyBlock, err
	}

	// Release references to the existing streams
	for i := range r.streams {
		r.streams[i].Finalize()
		r.streams[i] = nil
	}
	for i := range r.readers {
		r.readers[i] = nil
	}

	r.merged = xio.Block{
		SegmentReader: r.encoder.Stream(),
		Start:         r.blockStart,
		End:           r.blockEnd,
	}

	return r.merged, nil
}

func (r *dbMergedBlockReader) Clone() (xio.SegmentReader, error) {
	s0, err := r.streams[0].Clone()
	if err != nil {
		return nil, err
	}
	s1, err := r.streams[1].Clone()
	if err != nil {
		return nil, err
	}
	return newDatabaseMergedBlockReader(
		r.blockStart,
		r.blockEnd,
		s0,
		s1,
		r.opts,
	), nil
}

func (r *dbMergedBlockReader) Start() time.Time {
	return r.blockStart
}

func (r *dbMergedBlockReader) End() time.Time {
	return r.blockEnd
}

func (r *dbMergedBlockReader) Read(b []byte) (int, error) {
	reader, err := r.mergedReader()
	if err != nil {
		return 0, err
	}
	return reader.Read(b)
}

func (r *dbMergedBlockReader) Segment() (ts.Segment, error) {
	reader, err := r.mergedReader()
	if err != nil {
		return ts.Segment{}, err
	}
	return reader.Segment()
}

func (r *dbMergedBlockReader) SegmentReader() (xio.SegmentReader, error) {
	reader, err := r.mergedReader()
	if err != nil {
		return nil, err
	}
	return reader.SegmentReader, nil
}

func (r *dbMergedBlockReader) Reset(_ ts.Segment) {
	panic(fmt.Errorf("merged block reader not available for re-use"))
}

func (r *dbMergedBlockReader) ResetWindowed(_ ts.Segment, _, _ time.Time) {
	panic(fmt.Errorf("merged block reader not available for re-use"))
}

func (r *dbMergedBlockReader) Finalize() {
	r.Lock()

	r.blockStart = time.Time{}

	for i := range r.streams {
		if r.streams[i] != nil {
			r.streams[i].Finalize()
			r.streams[i] = nil
		}
	}
	for i := range r.readers {
		if r.readers[i] != nil {
			r.readers[i] = nil
		}
	}

	if !r.merged.IsEmpty() {
		r.merged.Finalize()
	}
	r.merged = xio.EmptyBlock

	if r.encoder != nil {
		r.encoder.Close()
	}
	r.encoder = nil

	r.err = nil

	r.Unlock()
}
