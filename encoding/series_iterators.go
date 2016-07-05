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

package encoding

import "github.com/m3db/m3db/interfaces/m3db"

type seriesIterators struct {
	iters  []m3db.SeriesIterator
	closed bool
	pool   m3db.MutableSeriesIteratorsPool
}

// NewSeriesIterators creates a new series iterators collection
func NewSeriesIterators(
	iters []m3db.SeriesIterator,
	pool m3db.MutableSeriesIteratorsPool,
) m3db.MutableSeriesIterators {
	it := &seriesIterators{iters: iters}
	it.Reset(0)
	return it
}

func (iters *seriesIterators) Iters() []m3db.SeriesIterator {
	return iters.iters
}

func (iters *seriesIterators) Close() {
	if iters.closed {
		return
	}
	iters.closed = true
	for _, iter := range iters.iters {
		iter.Close()
	}
	if iters.pool != nil {
		iters.pool.Put(iters)
	}
}

func (iters *seriesIterators) Len() int {
	return len(iters.iters)
}

func (iters *seriesIterators) Cap() int {
	return cap(iters.iters)
}

func (iters *seriesIterators) SetAt(idx int, iter m3db.SeriesIterator) {
	iters.iters[idx] = iter
}

func (iters *seriesIterators) Reset(size int) {
	iters.iters = iters.iters[:size]
}
