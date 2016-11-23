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

// +build linux

package pool

import (
	"fmt"
	"syscall"
)

const (
	mmapHugePageSize = 2 << 20
)

func mmap(n int) []byte {
	r, err := syscall.Mmap(-1, 0, n, mmapReadWriteAccess, mmapMemory)
	if err != nil {
		return nil, err
	}

	if n < mmapHugePageSize {
		return r, nil
	} else if err := syscall.Madvise(r, syscall.MADV_HUGEPAGE); err != nil {
		fmt.Printf("unable to turn THP on for %p: %v\n", &r[0], err)
	}

	return r, nil
}
