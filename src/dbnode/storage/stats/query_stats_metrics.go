// Copyright (c) 2020 Uber Technologies, Inc.
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

package stats

import (
	"sync"
	"time"

	"github.com/m3db/m3/src/x/instrument"
	"github.com/uber-go/tally"
)

// Tracker implementation that emits query stats as metrics.
type queryStatsMetricsTracker struct {
	sync.Mutex
	recentDocs tally.Gauge
}

// DefaultQueryStatsTrackerForMetrics provides a tracker
// implementation that emits query stats as metrics.
func DefaultQueryStatsTrackerForMetrics(opts instrument.Options) QueryStatsTracker {
	scope := opts.
		MetricsScope().
		SubScope("query-stats")
	return &queryStatsMetricsTracker{
		recentDocs: scope.Gauge("recentDocs"),
	}
}

func (t *queryStatsMetricsTracker) TrackDocs(recentDocs int) error {
	t.Lock()
	t.recentDocs.Update(float64(recentDocs))
	t.Unlock()
	return nil
}

func (t *queryStatsMetricsTracker) Lookback() time.Duration {
	return time.Second
}
