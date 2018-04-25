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
	"github.com/m3db/m3db/ts"
)

// nolint: deadcode
type nullSegmentReader struct{}

func (r nullSegmentReader) Read([]byte) (n int, err error) { return 0, nil }
func (r nullSegmentReader) Segment() (ts.Segment, error)   { return ts.Segment{}, nil }
func (r nullSegmentReader) Reset(ts.Segment)               {}
func (r nullSegmentReader) Finalize()                      {}
func (r nullSegmentReader) Clone() (Reader, error)         { return r, nil }

// nolint: deadcode
type nullSegment struct{}

func (r nullSegment) Read([]byte) (n int, err error) { return 0, nil }
func (r nullSegment) Clone() (Reader, error)         { return r, nil }
