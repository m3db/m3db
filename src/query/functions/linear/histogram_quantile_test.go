// Copyright (c) 2019 Uber Technologies, Inc.
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

package linear

import (
	"math"
	"testing"
	"time"

	"github.com/m3db/m3/src/query/block"
	"github.com/m3db/m3/src/query/executor/transform"
	"github.com/m3db/m3/src/query/models"
	"github.com/m3db/m3/src/query/parser"
	"github.com/m3db/m3/src/query/test"
	"github.com/m3db/m3/src/query/test/executor"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGatherSeriesToBuckets(t *testing.T) {
	name := []byte("name")
	bucket := []byte("bucket")
	tagOpts := models.NewTagOptions().
		SetIDSchemeType(models.TypeQuoted).
		SetMetricName(name).
		SetBucketName(bucket)

	tags := models.NewTags(3, tagOpts).SetName([]byte("foo")).AddTag(models.Tag{
		Name:  []byte("bar"),
		Value: []byte("baz"),
	})

	noBucketMeta := block.SeriesMeta{Tags: tags}
	invalidBucketMeta := block.SeriesMeta{Tags: tags.Clone().SetBucket([]byte("string"))}
	validMeta := block.SeriesMeta{Tags: tags.Clone().SetBucket([]byte("0.1"))}
	validMeta2 := block.SeriesMeta{Tags: tags.Clone().SetBucket([]byte("0.1"))}
	validMeta3 := block.SeriesMeta{Tags: tags.Clone().SetBucket([]byte("10"))}
	infMeta := block.SeriesMeta{Tags: tags.Clone().SetBucket([]byte("Inf"))}
	validMetaMoreTags := block.SeriesMeta{Tags: tags.Clone().SetBucket([]byte("0.1")).AddTag(models.Tag{
		Name:  []byte("qux"),
		Value: []byte("qar"),
	})}

	metas := []block.SeriesMeta{
		validMeta, noBucketMeta, invalidBucketMeta, validMeta2, validMetaMoreTags, validMeta3, infMeta,
	}

	actual := gatherSeriesToBuckets(metas)
	expected := bucketedSeries{
		`{bar="baz"}`: indexedBuckets{
			buckets: []indexedBucket{
				{upperBound: 0.1, idx: 0},
				{upperBound: 0.1, idx: 3},
				{upperBound: 10, idx: 5},
				{upperBound: math.Inf(1), idx: 6},
			},
			tags: models.NewTags(1, tagOpts).AddTag(models.Tag{
				Name:  []byte("bar"),
				Value: []byte("baz"),
			}),
		},
		`{bar="baz",qux="qar"}`: indexedBuckets{
			buckets: []indexedBucket{
				{upperBound: 0.1, idx: 4},
			},
			tags: models.NewTags(1, tagOpts).AddTag(models.Tag{
				Name:  []byte("bar"),
				Value: []byte("baz"),
			}).AddTag(models.Tag{
				Name:  []byte("qux"),
				Value: []byte("qar"),
			}),
		},
	}

	assert.Equal(t, expected, actual)
}

func TestSanitizeBuckets(t *testing.T) {
	bucketed := bucketedSeries{
		`{bar="baz"}`: indexedBuckets{
			buckets: []indexedBucket{
				{upperBound: 10, idx: 5},
				{upperBound: math.Inf(1), idx: 6},
				{upperBound: 1, idx: 0},
				{upperBound: 2, idx: 3},
			},
		},
		`{with="neginf"}`: indexedBuckets{
			buckets: []indexedBucket{
				{upperBound: 10, idx: 5},
				{upperBound: math.Inf(-1), idx: 6},
				{upperBound: 1, idx: 0},
				{upperBound: 2, idx: 3},
			},
		},
		`{no="infinity"}`: indexedBuckets{
			buckets: []indexedBucket{
				{upperBound: 0.1, idx: 4},
				{upperBound: 0.2, idx: 14},
				{upperBound: 0.3, idx: 114},
			},
		},
		`{just="infinity"}`: indexedBuckets{
			buckets: []indexedBucket{
				{upperBound: math.Inf(1), idx: 4},
			},
		},
		`{just="neg-infinity"}`: indexedBuckets{
			buckets: []indexedBucket{
				{upperBound: math.Inf(-1), idx: 4},
			},
		},
	}

	actual := bucketedSeries{
		`{bar="baz"}`: indexedBuckets{
			buckets: []indexedBucket{
				{upperBound: 1, idx: 0},
				{upperBound: 2, idx: 3},
				{upperBound: 10, idx: 5},
				{upperBound: math.Inf(1), idx: 6},
			},
		},
	}

	sanitizeBuckets(bucketed)
	assert.Equal(t, actual, bucketed)
}

func TestBucketQuantile(t *testing.T) {
	// single bucket returns nan
	actual := bucketQuantile(0.5, []bucketValue{{upperBound: 1, value: 1}})
	assert.True(t, math.IsNaN(actual))

	// bucket with no infinity returns nan
	actual = bucketQuantile(0.5, []bucketValue{
		{upperBound: 1, value: 1},
		{upperBound: 2, value: 2},
	})
	assert.True(t, math.IsNaN(actual))

	// bucket with negative infinity bound returns nan
	actual = bucketQuantile(0.5, []bucketValue{
		{upperBound: 1, value: 1},
		{upperBound: 2, value: 2},
		{upperBound: math.Inf(-1), value: 22},
	})
	assert.True(t, math.IsNaN(actual))

	actual = bucketQuantile(0.5, []bucketValue{
		{upperBound: 1, value: 1},
		{upperBound: math.Inf(1), value: 22},
	})
	assert.Equal(t, float64(1), actual)

	// NB: tested against Prom
	buckets := []bucketValue{
		{upperBound: 1, value: 1},
		{upperBound: 2, value: 2},
		{upperBound: 5, value: 5},
		{upperBound: 10, value: 10},
		{upperBound: 20, value: 15},
		{upperBound: math.Inf(1), value: 16},
	}

	actual = bucketQuantile(0, buckets)
	assert.InDelta(t, float64(0), actual, 0.0001)

	actual = bucketQuantile(0.15, buckets)
	assert.InDelta(t, 2.4, actual, 0.0001)

	actual = bucketQuantile(0.2, buckets)
	assert.InDelta(t, float64(3.2), actual, 0.0001)

	actual = bucketQuantile(0.5, buckets)
	assert.InDelta(t, float64(8), actual, 0.0001)

	actual = bucketQuantile(0.8, buckets)
	assert.InDelta(t, float64(15.6), actual, 0.0001)

	actual = bucketQuantile(1, buckets)
	assert.InDelta(t, float64(20), actual, 0.0001)
}

func TestNewOp(t *testing.T) {
	args := make([]interface{}, 0, 1)
	_, err := NewHistogramQuantileOp(args, HistogramQuantileType)
	assert.Error(t, err)

	args = append(args, "invalid")
	_, err = NewHistogramQuantileOp(args, HistogramQuantileType)
	assert.Error(t, err)

	args[0] = 2.0
	_, err = NewHistogramQuantileOp(args, ClampMaxType)
	assert.Error(t, err)

	op, err := NewHistogramQuantileOp(args, HistogramQuantileType)
	assert.NoError(t, err)

	assert.Equal(t, HistogramQuantileType, op.OpType())
	assert.Equal(t, "type: histogram_quantile", op.String())
}

func testQuantileFunctionWithQ(t *testing.T, q float64) [][]float64 {
	args := make([]interface{}, 0, 1)
	args = append(args, q)
	op, err := NewHistogramQuantileOp(args, HistogramQuantileType)
	require.NoError(t, err)

	name := []byte("name")
	bucket := []byte("bucket")
	tagOpts := models.NewTagOptions().
		SetIDSchemeType(models.TypeQuoted).
		SetMetricName(name).
		SetBucketName(bucket)

	tags := models.NewTags(3, tagOpts).SetName([]byte("foo")).AddTag(models.Tag{
		Name:  []byte("bar"),
		Value: []byte("baz"),
	})

	seriesMetas := []block.SeriesMeta{
		{Tags: tags.Clone().SetBucket([]byte("1"))},
		{Tags: tags.Clone().SetBucket([]byte("2"))},
		{Tags: tags.Clone().SetBucket([]byte("5"))},
		{Tags: tags.Clone().SetBucket([]byte("10"))},
		{Tags: tags.Clone().SetBucket([]byte("20"))},
		{Tags: tags.Clone().SetBucket([]byte("Inf"))},
		// this series should not be part of the output, since it has no bucket tag.
		{Tags: tags.Clone()},
	}

	v := [][]float64{
		{1, 1, 11, 12, 1},
		{2, 2, 12, 13, 2},
		{5, 5, 15, 40, 5},
		{10, 10, 20, 50, 10},
		{15, 15, 25, 70, 15},
		{16, 19, 26, 71, 1},
	}

	bounds := models.Bounds{
		Start:    time.Now(),
		Duration: time.Minute * 5,
		StepSize: time.Minute,
	}

	bl := test.NewBlockFromValuesWithSeriesMeta(bounds, seriesMetas, v)
	c, sink := executor.NewControllerWithSink(parser.NodeID(1))
	node := op.(histogramQuantileOp).Node(c, transform.Options{})
	err = node.Process(parser.NodeID(0), bl)
	require.NoError(t, err)

	return sink.Values
}

var (
	inf  = math.Inf(+1)
	ninf = math.Inf(-1)
)

func TestQuantileFunctionForInvalidQValues(t *testing.T) {
	actual := testQuantileFunctionWithQ(t, -1)
	assert.Equal(t, [][]float64{{ninf, ninf, ninf, ninf, ninf}}, actual)
	actual = testQuantileFunctionWithQ(t, 1.1)
	assert.Equal(t, [][]float64{{inf, inf, inf, inf, inf}}, actual)

	actual = testQuantileFunctionWithQ(t, 0.8)
	test.EqualsWithNansWithDelta(t, [][]float64{{15.6, 20, 11.6, 13.4, 0.8}}, actual, 0.00001)
}
