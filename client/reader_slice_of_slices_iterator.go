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

package client

import (
	"time"

	"github.com/m3db/m3db/generated/thrift/rpc"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3db/x/xio"
	"github.com/m3db/m3x/checked"
	xtime "github.com/m3db/m3x/time"
)

var timeZero = time.Time{}

type readerSliceOfSlicesIterator struct {
	segments     []*rpc.Segments
	blockReaders []xio.BlockReader
	idx          int
	closed       bool
	pool         *readerSliceOfSlicesIteratorPool
}

func newReaderSliceOfSlicesIterator(
	segments []*rpc.Segments,
	pool *readerSliceOfSlicesIteratorPool,
) *readerSliceOfSlicesIterator {
	it := &readerSliceOfSlicesIterator{pool: pool}
	it.Reset(segments)
	return it
}

func (it *readerSliceOfSlicesIterator) Next() bool {
	if !(it.idx+1 < len(it.segments)) {
		return false
	}
	it.idx++

	// Extend segment readers if not enough available
	currLen := it.CurrentLen()
	if len(it.blockReaders) < currLen {
		diff := currLen - len(it.blockReaders)
		for i := 0; i < diff; i++ {
			seg := ts.NewSegment(nil, nil, ts.FinalizeNone)
			sr := xio.NewSegmentReader(seg)
			br := xio.NewBlockReader(sr, it.CurrentStart(), it.CurrentEnd())
			it.blockReaders = append(it.blockReaders, br)
		}
	}

	// Set the segment readers to reader from current segment pieces
	segment := it.segments[it.idx]
	if segment.Merged != nil {
		it.resetReader(it.blockReaders[0], segment.Merged)
	} else {
		for i := 0; i < currLen; i++ {
			it.resetReader(it.blockReaders[i], segment.Unmerged[i])
		}
	}

	return true
}

func (it *readerSliceOfSlicesIterator) resetReader(
	r xio.BlockReader,
	seg *rpc.Segment,
) {
	rseg, err := r.Segment()
	if err != nil {
		r.ResetWindowed(ts.Segment{}, it.CurrentStart(), it.CurrentEnd())
		return
	}

	var (
		head = rseg.Head
		tail = rseg.Tail
	)
	if head == nil {
		head = checked.NewBytes(seg.Head, nil)
		head.IncRef()
	} else {
		head.Reset(seg.Head)
	}
	if tail == nil {
		tail = checked.NewBytes(seg.Tail, nil)
		tail.IncRef()
	} else {
		tail.Reset(seg.Tail)
	}
	r.ResetWindowed(ts.NewSegment(head, tail, ts.FinalizeNone), it.CurrentStart(), it.CurrentEnd())
}

func (it *readerSliceOfSlicesIterator) CurrentLen() int {
	if it.segments[it.idx].Merged != nil {
		return 1
	}
	return len(it.segments[it.idx].Unmerged)
}

func timeConvert(ticks *int64) time.Time {
	if ticks == nil {
		return timeZero
	}
	return xtime.FromNormalizedTime(*ticks, time.Nanosecond)
}

func (it *readerSliceOfSlicesIterator) CurrentStart() time.Time {
	segments := it.segments[it.idx]
	if segments.Merged != nil {
		return timeConvert(segments.Merged.StartTime)
	}
	if len(segments.Unmerged) == 0 {
		return timeZero
	}
	return timeConvert(segments.Unmerged[0].StartTime)
}

func (it *readerSliceOfSlicesIterator) CurrentEnd() time.Time {
	segments := it.segments[it.idx]
	if segments.Merged != nil {
		return timeConvert(segments.Merged.EndTime)
	}
	return timeConvert(segments.Unmerged[0].EndTime)
}

func (it *readerSliceOfSlicesIterator) CurrentAt(idx int) xio.Reader {
	if idx >= it.CurrentLen() {
		return nil
	}
	return it.blockReaders[idx]
}

func (it *readerSliceOfSlicesIterator) Close() {
	if it.closed {
		return
	}
	it.closed = true
	// Release any refs to segments
	it.segments = nil
	// Release any refs to segment byte slices
	for i := range it.blockReaders {
		seg, err := it.blockReaders[i].Segment()
		if err != nil {
			continue
		}
		if seg.Head != nil {
			seg.Head.Reset(nil)
		}
		if seg.Tail != nil {
			seg.Tail.Reset(nil)
		}
	}
	if pool := it.pool; pool != nil {
		pool.Put(it)
	}
}

func (it *readerSliceOfSlicesIterator) Reset(segments []*rpc.Segments) {
	it.segments = segments
	it.idx = -1
	it.closed = false
}
