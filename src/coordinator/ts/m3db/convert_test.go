// Copyright (c) 2018 Uber Technologies, Inc.
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

package m3db

import (
	"io"
	"testing"
	"time"

	"github.com/m3db/m3db/src/coordinator/storage"
	"github.com/m3db/m3db/src/coordinator/test"
	"github.com/m3db/m3db/src/dbnode/encoding"
	"github.com/m3db/m3db/src/dbnode/encoding/m3tsz"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	seriesID        = test.SeriesID
	seriesNamespace = test.SeriesNamespace

	testTags = test.TestTags

	blockSize = test.BlockSize

	start       = test.Start
	seriesStart = test.SeriesStart
	middle      = test.Middle
	end         = test.End

	// required for iterator pool
	testIterAlloc = func(r io.Reader) encoding.ReaderIterator {
		return m3tsz.NewReaderIterator(r, m3tsz.DefaultIntOptimizationEnabled, encoding.NewOptions())
	}
)

func TestConversion(t *testing.T) {
	iter, err := test.BuildTestSeriesIterator()
	require.NoError(t, err)
	iterators := encoding.NewSeriesIterators([]encoding.SeriesIterator{iter}, nil)

	blocks, err := ConvertM3DBSeriesIterators(iterators, testIterAlloc)
	require.NoError(t, err)

	for _, block := range blocks {
		assert.Equal(t, seriesID, block.ID.String())
		assert.Equal(t, seriesNamespace, block.Namespace.String())

		convertedTags, err := storage.FromIdentTagIteratorToTags(block.Tags)
		require.NoError(t, err)
		assert.Equal(t, testTags["foo"], convertedTags["foo"])
		assert.Equal(t, testTags["baz"], convertedTags["baz"])

		blockOneSeriesIterator := block.Blocks[0].SeriesIterator
		blockTwoSeriesIterator := block.Blocks[1].SeriesIterator

		assert.Equal(t, start.Add(2*time.Minute), blockOneSeriesIterator.Start())
		assert.Equal(t, middle, blockOneSeriesIterator.End())

		for i := 3; blockOneSeriesIterator.Next(); i++ {
			dp, _, _ := blockOneSeriesIterator.Current()
			assert.Equal(t, float64(i), dp.Value)
		}

		assert.Equal(t, middle, blockTwoSeriesIterator.Start())
		assert.Equal(t, end, blockTwoSeriesIterator.End())

		for i := 101; blockTwoSeriesIterator.Next(); i++ {
			dp, _, _ := blockTwoSeriesIterator.Current()
			assert.Equal(t, float64(i), dp.Value)
		}
	}
}
