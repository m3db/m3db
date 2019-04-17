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

package server

// This file contains reporters and setup for our query/cost.ChainedEnforcer
// instances.
import (
	"sync"

	"github.com/m3db/m3/src/cmd/services/m3query/config"
	qcost "github.com/m3db/m3/src/query/cost"
	"github.com/m3db/m3/src/x/cost"
	"github.com/m3db/m3/src/x/instrument"

	"github.com/uber-go/tally"
)

const (
	queriesOverLimitMetric  = "over_datapoints_limit"
	datapointsMetric        = "datapoints"
	datapointsCounterMetric = "datapoints_counter"
	maxDatapointsHistMetric = "max_datapoints_hist"
)

// newConfiguredChainedEnforcer returns a ChainedEnforcer with 3 configured
// levels: global, per-query, per-block. Global and per-query both have limits
// on them (as configured by cfg.Limits); per-block is purely for accounting
// purposes.
// Our enforcers report at least these stats:
//   cost.global.datapoints: gauge; the number of datapoints currently in use
//                           by this instance.
//
//   cost.global.datapoints_counter: counter; counter representation of the
//                                   number of datapoints in use by this instance
//
//   cost.{per_query,global}.over_datapoints_limit: counter; how many queries are over the
//    datapoint limit
//
//   cost.per_query.max_datapoints_hist: histogram; represents the
//     distribution of the maximum datapoints used at any point in each query.
func newConfiguredChainedEnforcer(cfg *config.Configuration, instrumentOptions instrument.Options) (qcost.ChainedEnforcer, error) {
	costScope := instrumentOptions.MetricsScope().SubScope("cost")
	costIops := instrumentOptions.SetMetricsScope(costScope)
	limitMgr := cost.NewStaticLimitManager(cfg.Limits.Global.AsLimitManagerOptions().SetInstrumentOptions(costIops))
	tracker := cost.NewTracker()

	globalEnforcer := cost.NewEnforcer(limitMgr, tracker,
		cost.NewEnforcerOptions().SetReporter(
			newGlobalReporter(costScope.SubScope("global")),
		).SetCostExceededMessage("limits.global.maxFetchedDatapoints exceeded"),
	)

	queryEnforcerOpts := cost.NewEnforcerOptions().SetCostExceededMessage("limits.perQuery.maxFetchedDatapoints exceeded").
		SetReporter(newPerQueryReporter(costScope.
			SubScope("per_query")))

	queryEnforcer := cost.NewEnforcer(
		cost.NewStaticLimitManager(cfg.Limits.PerQuery.AsLimitManagerOptions()),
		cost.NewTracker(),
		queryEnforcerOpts)

	blockEnforcer := cost.NewEnforcer(
		cost.NewStaticLimitManager(cost.NewLimitManagerOptions().SetDefaultLimit(cost.Limit{Enabled: false})),
		cost.NewTracker(),
		nil,
	)

	return qcost.NewChainedEnforcer(qcost.GlobalLevel, []cost.Enforcer{
		globalEnforcer,
		queryEnforcer,
		blockEnforcer,
	})
}

// globalReporter records ChainedEnforcer statistics for the global enforcer.
type globalReporter struct {
	datapoints        tally.Gauge
	datapointsCounter tally.Counter
	overLimit         overLimitReporter
}

// assert we implement the interface
var _ cost.EnforcerReporter = (*globalReporter)(nil)

func newGlobalReporter(s tally.Scope) *globalReporter {
	return &globalReporter{
		datapoints:        s.Gauge(datapointsMetric),
		datapointsCounter: s.Counter(datapointsCounterMetric),
		overLimit:         newOverLimitReporter(s),
	}
}

func (gr *globalReporter) ReportCurrent(c cost.Cost) {
	gr.datapoints.Update(float64(c))
}

// ReportCost for global reporters sends the new incoming cost to a counter.
// Since counters can only be incremented, it ignores negative values.
func (gr *globalReporter) ReportCost(c cost.Cost) {
	if c > 0 {
		gr.datapointsCounter.Inc(int64(c))
	}
}

// ReportOverLimit delegates to gr.overLimit
func (gr *globalReporter) ReportOverLimit(enabled bool) {
	gr.overLimit.ReportOverLimit(enabled)
}

// perQueryReporter records ChainedEnforcer statistics on a per query level.
type perQueryReporter struct {
	mu            *sync.Mutex
	maxDatapoints cost.Cost
	queryHisto    tally.Histogram
	overLimit     overLimitReporter
}

// assert we implement the interface
var _ qcost.ChainedReporter = (*perQueryReporter)(nil)

func newPerQueryReporter(scope tally.Scope) *perQueryReporter {
	return &perQueryReporter{
		mu:            &sync.Mutex{},
		maxDatapoints: 0,
		queryHisto: scope.Histogram(maxDatapointsHistMetric,
			tally.MustMakeExponentialValueBuckets(10.0, 10.0, 6)),
		overLimit: newOverLimitReporter(scope),
	}
}

// ReportCost is a noop for perQueryReporter because it's noisy to report
// the current cost for every query (hard to meaningfully divide out).
// Instead, we report the max datapoints at the end of the query--see on
// release.
func (perQueryReporter) ReportCost(c cost.Cost) {}

// ReportCurrent is a noop for perQueryReporter--see ReportCost for
// explanation.
func (perQueryReporter) ReportCurrent(c cost.Cost) {}

// ReportOverLimit reports when a query is over its per query limit.
func (pr *perQueryReporter) ReportOverLimit(enabled bool) {
	pr.overLimit.ReportOverLimit(enabled)
}

// OnChildClose takes the max of the current cost for this query and the
// previously recorded cost. We do this OnChildRelease instead of on
// ReportCurrent to avoid locking every time we add to the Enforcer.
func (pr *perQueryReporter) OnChildClose(curCost cost.Cost) {
	pr.mu.Lock()
	if curCost > pr.maxDatapoints {
		pr.maxDatapoints = curCost
	}
	pr.mu.Unlock()
}

// OnClose records the maximum cost seen by this reporter.
func (pr *perQueryReporter) OnClose(curCost cost.Cost) {
	pr.mu.Lock()
	pr.queryHisto.RecordValue(float64(pr.maxDatapoints))
	pr.mu.Unlock()
}

// overLimitReporter factors out reporting over limit cases for both global
// and per query enforcer reporters.
type overLimitReporter struct {
	queriesOverLimitDisabled tally.Counter
	queriesOverLimitEnabled  tally.Counter
}

func newOverLimitReporter(scope tally.Scope) overLimitReporter {
	return overLimitReporter{
		queriesOverLimitDisabled: scope.Tagged(map[string]string{
			"enabled": "false",
		}).Counter(queriesOverLimitMetric),

		queriesOverLimitEnabled: scope.Tagged(map[string]string{
			"enabled": "true",
		}).Counter(queriesOverLimitMetric),
	}
}

// ReportOverLimit increments <prefix>.over_limit, tagged by
// "enabled".
func (olr overLimitReporter) ReportOverLimit(enabled bool) {
	if enabled {
		olr.queriesOverLimitEnabled.Inc(1)
	} else {
		olr.queriesOverLimitDisabled.Inc(1)
	}
}
