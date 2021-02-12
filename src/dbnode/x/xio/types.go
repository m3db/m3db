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

package xio

import (
	"io"
	"time"

	"github.com/m3db/m3/src/dbnode/ts"
	"github.com/m3db/m3/src/x/context"
	"github.com/m3db/m3/src/x/pool"
	xresource "github.com/m3db/m3/src/x/resource"
)

// BlockReader represents a block reader backed by a
// SegmentReader with start time and block size.
type BlockReader struct {
	SegmentReader
	Start     time.Time
	BlockSize time.Duration
}

// ReadSegment reads the Segment, blocking until the segment is available or the deadline expires.
func (b BlockReader) ReadSegment(ctx context.Context) (ts.Segment, error) {
	done := make(chan struct{}, 1)
	var (
		segment ts.Segment
		err error
	)
	go func() {
		segment, err = b.Segment()
		done <- struct{}{}
	}()
	select {
	case <-ctx.GoContext().Done():
		return ts.Segment{}, ctx.GoContext().Err()
	case <-done:
		return segment, err
	}
}

// EmptyBlockReader represents the default block reader.
var EmptyBlockReader = BlockReader{}

// SegmentReader implements the io reader interface backed by a segment.
type SegmentReader interface {
	io.Reader
	xresource.Finalizer

	// Segment gets the segment read by this reader.
	Segment() (ts.Segment, error)

	// Reset resets the reader to read a new segment.
	Reset(segment ts.Segment)

	// Clone returns a clone of a copy of the underlying data reset,
	// with an optional byte pool to reduce allocations.
	Clone(pool pool.CheckedBytesPool) (SegmentReader, error)
}

// SegmentReaderPool provides a pool for segment readers.
type SegmentReaderPool interface {
	// Init will initialize the pool.
	Init()

	// Get provides a segment reader from the pool.
	Get() SegmentReader

	// Put returns a segment reader to the pool.
	Put(sr SegmentReader)
}

// ReaderSliceOfSlicesIterator is an iterator that iterates
// through an array of reader arrays.
type ReaderSliceOfSlicesIterator interface {
	// Next moves to the next item.
	Next() bool

	// CurrentReaders returns the current length, start time, and block size.
	CurrentReaders() (length int, start time.Time, blockSize time.Duration)

	// CurrentReaderAt returns the current reader in the slice
	// of readers at an index.
	CurrentReaderAt(idx int) BlockReader

	// Close closes the iterator.
	Close()

	// Size gives the size of bytes in this iterator.
	Size() (int, error)

	// RewindToIndex returns the iterator to a specific index.
	// This operation is invalid if any of the block readers have been read.
	RewindToIndex(idx int)

	// Index returns the iterator's current index.
	Index() int
}

// ReaderSliceOfSlicesFromBlockReadersIterator is an iterator
// that iterates through an array of reader arrays.
type ReaderSliceOfSlicesFromBlockReadersIterator interface {
	ReaderSliceOfSlicesIterator

	// Reset resets the iterator with a new array of block readers arrays.
	Reset(blocks [][]BlockReader)
}
