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
	"time"

	"github.com/m3db/m3db/ts"
)

// CloneBlock returns a clone of the block with the underlying data reset
func (b BlockReader) CloneBlock() (BlockReader, error) {
	sr, err := b.SegmentReader.Clone()
	if err != nil {
		return EmptyBlockReader, err
	}
	return BlockReader{
		SegmentReader: sr,
		Start:         b.Start,
		End:           b.End,
	}, nil
}

// IsEmpty returns true for the empty block
func (b BlockReader) IsEmpty() bool {
	return b.Start.Equal(timeZero) && b.End.Equal(timeZero) && b.SegmentReader == nil
}

// IsNotEmpty returns false for the empty block
func (b BlockReader) IsNotEmpty() bool {
	return !b.IsEmpty()
}

// ResetWindowed resets the underlying reader window, as well as start and end times for the block
func (b *BlockReader) ResetWindowed(segment ts.Segment, start, end time.Time) {
	b.Reset(segment)
	b.Start = start
	b.End = end
}
