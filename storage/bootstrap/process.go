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

package bootstrap

import (
	"errors"
	"time"

	"github.com/m3db/m3x/time"
)

var (
	errUnfulfilledRanges = errors.New("bootstrap finished with unfulfilled ranges")
)

// bootstrapProcess represents the bootstrapping process.
type bootstrapProcess struct {
	opts         Options
	bootstrapper Bootstrapper
}

// NewBootstrapProcess creates a new bootstrap process.
func NewBootstrapProcess(
	opts Options,
	bootstrapper Bootstrapper,
) Bootstrap {
	return &bootstrapProcess{
		opts:         opts,
		bootstrapper: bootstrapper,
	}
}

// initTimeRanges initializes the target time ranges.
// NB(xichen): bootstrapping is now a two-step process: we bootstrap the data between
// [writeStart - retentionPeriod, writeStart - bufferPast) in the first step, and the
// data between [writeStart - bufferPast, writeStart + bufferFuture) in the second step.
// The reason why this is necessary is because writers can write data in the past or in
// the future, and as a result, future writes between [writeStart - bufferFuture, writeStart)
// are lost, which means the data buffered between [writeStart, writeStart + bufferFuture)
// cannot be trusted and therefore need to be bootstrapped as well, which means in the worst
// case we need to wait till writeStart + bufferPast + bufferFuture before bootstrap can complete.
func (b *bootstrapProcess) initTimeRanges(writeStart time.Time) xtime.Ranges {
	start := writeStart.Add(-b.opts.GetRetentionOptions().GetRetentionPeriod())
	midPoint := writeStart.Add(-b.opts.GetRetentionOptions().GetBufferPast())
	end := writeStart.Add(b.opts.GetRetentionOptions().GetBufferFuture())

	return xtime.NewRanges().
		AddRange(xtime.Range{Start: start, End: midPoint}).
		AddRange(xtime.Range{Start: midPoint, End: end})
}

// bootstrap returns the bootstrapped results for a given shard time ranges.
func (b *bootstrapProcess) bootstrap(timeRanges xtime.Ranges, shards []uint32) (Result, error) {
	result := NewResult()

	it := timeRanges.Iter()
	for it.Next() {
		shardsTimeRanges := make(map[uint32]xtime.Ranges, len(shards))

		r := xtime.NewRanges()
		r.AddRange(it.Value())
		for _, s := range shards {
			shardsTimeRanges[s] = r
		}

		res, err := b.bootstrapper.Bootstrap(shardsTimeRanges)
		if err != nil {
			return nil, err
		}
		result.AddResult(res)
	}

	return result, nil
}

// Run initiates the bootstrap process, where writeStart is when we start
// accepting incoming writes.
func (b *bootstrapProcess) Run(writeStart time.Time, shards []uint32) (Result, error) {

	// Initialize the target range we'd like to bootstrap
	unfulfilledRanges := b.initTimeRanges(writeStart)

	return b.bootstrap(unfulfilledRanges, shards)
}
