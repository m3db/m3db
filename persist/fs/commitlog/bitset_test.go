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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitSetSetValue(t *testing.T) {
	bs := newBitset().(*set)
	var values []uint64

	// Setting a value smaller than the bitset length doesn't
	// trigger reallocations
	oldLen := bs.Len()
	values = append(values, uint64(oldLen-1))
	bs.set(values[len(values)-1])
	require.Equal(t, oldLen, bs.Len())
	for _, v := range values {
		require.True(t, bs.has(v))
	}

	// Setting a value bigger than the bitset length,
	// which triggers an reallocation, and verify the capacity
	// has grown and all the existing data are kept
	values = append(values, uint64(oldLen+1))
	bs.set(values[len(values)-1])
	require.Equal(t, 2*oldLen, bs.Len())
	for _, v := range values {
		require.True(t, bs.has(v))
	}

	// Setting a value bigger than 2 times the bitset length
	// will trigger an reallocation and set the length of
	// the new underlying bitset to value + 1
	newVal := bs.Len()*2 + 10
	values = append(values, uint64(newVal))
	bs.set(values[len(values)-1])
	require.Equal(t, newVal+1, bs.Len())
	for _, v := range values {
		require.True(t, bs.has(v))
	}
}
