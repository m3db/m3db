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

package aggregator

import (
	"container/list"
	"math"
	"sync"
	"time"

	"github.com/m3db/m3aggregator/id"
	"github.com/m3db/m3metrics/metric"
	"github.com/m3db/m3metrics/metric/unaggregated"
	"github.com/m3db/m3metrics/policy"
	"github.com/m3db/m3x/clock"

	"github.com/uber-go/tally"
)

var (
	emptyHashedEntry hashedEntry
)

type hashedEntry struct {
	idHash id.Hash
	entry  *Entry
}

type metricMapMetrics struct {
	numEntries tally.Gauge
}

func newMetricMapMetrics(scope tally.Scope) metricMapMetrics {
	return metricMapMetrics{
		numEntries: scope.Gauge("num-entries"),
	}
}

// NB(xichen): use a type-specific list for hashedEntry if the conversion
// overhead between interface{} and hashedEntry becomes a problem.
type metricMap struct {
	sync.RWMutex

	opts         Options
	nowFn        clock.NowFn
	entryPool    EntryPool
	batchPercent float64

	metricLists *metricLists
	entries     map[id.Hash]*list.Element
	entryList   *list.List
	waitForFn   waitForFn
	metrics     metricMapMetrics
}

func newMetricMap(opts Options) *metricMap {
	scope := opts.InstrumentOptions().MetricsScope().SubScope("map")
	metricLists := newMetricLists(opts)

	return &metricMap{
		opts:         opts,
		nowFn:        opts.ClockOptions().NowFn(),
		entryPool:    opts.EntryPool(),
		batchPercent: opts.EntryCheckBatchPercent(),
		metricLists:  metricLists,
		entries:      make(map[id.Hash]*list.Element),
		entryList:    list.New(),
		waitForFn:    time.After,
		metrics:      newMetricMapMetrics(scope),
	}
}

func (m *metricMap) AddMetricWithPolicies(
	mu unaggregated.MetricUnion,
	policies policy.VersionedPolicies,
) error {
	e := m.findOrCreate(mu.ID)
	err := e.AddMetricWithPolicies(mu, policies)
	e.DecWriter()
	return err
}

func (m *metricMap) DeleteExpired(target time.Duration) int64 {
	now := m.nowFn()

	// Determine batch size.
	m.RLock()
	elemsLen := m.entryList.Len()
	if elemsLen == 0 {
		// If the list is empty, nothing to do.
		m.RUnlock()
		return 0
	}
	batchSize := int(math.Max(1.0, math.Ceil(m.batchPercent*float64(elemsLen))))
	numBatches := int(math.Ceil(float64(elemsLen) / float64(batchSize)))
	targetPerBatch := target / time.Duration(numBatches)
	currElem := m.entryList.Front()
	m.RUnlock()

	// NB(xichen): if this runs frequently enough, stash the expired buffer
	// in the map object and reuse the buffer.
	var (
		numExpired int64
		expired    = make([]hashedEntry, 0, batchSize)
		batchIdx   int
	)
	for currElem != nil {
		m.RLock()
		for numChecked := 0; numChecked < batchSize && currElem != nil; numChecked++ {
			nextElem := currElem.Next()
			hashedEntry := currElem.Value.(hashedEntry)
			if hashedEntry.entry.ShouldExpire(now) {
				expired = append(expired, hashedEntry)
			}
			currElem = nextElem
		}
		m.RUnlock()

		// Actually purging the expired entries.
		if len(expired) >= batchSize {
			numExpired += m.purgeExpired(now, expired)
			for i := range expired {
				expired[i] = emptyHashedEntry
			}
			expired = expired[:0]
		}

		batchIdx++
		targetTime := now.Add(time.Duration(batchIdx) * targetPerBatch)
		currTime := m.nowFn()
		if currTime.Before(targetTime) {
			m.waitForFn(targetTime.Sub(currTime))
		}
	}

	// Purge if there are remaining expired entries.
	numExpired += m.purgeExpired(now, expired)
	for i := range expired {
		expired[i] = emptyHashedEntry
	}
	return numExpired
}

func (m *metricMap) Close() {
	m.metricLists.Close()
}

func (m *metricMap) findOrCreate(mid metric.ID) *Entry {
	idHash := id.HashFn(mid)
	m.RLock()
	if entry, found := m.lookupEntryWithLock(idHash); found {
		// NB(xichen): it is important to increase number of writers
		// within a lock so we can account for active writers
		// when deleting expired entries.
		entry.IncWriter()
		m.RUnlock()
		return entry
	}
	m.RUnlock()

	m.Lock()
	entry, found := m.lookupEntryWithLock(idHash)
	if !found {
		entry = m.entryPool.Get()
		entry.ResetSetData(m.metricLists)
		m.entries[idHash] = m.entryList.PushBack(hashedEntry{
			idHash: idHash,
			entry:  entry,
		})
	}
	entry.IncWriter()
	m.Unlock()

	return entry
}

func (m *metricMap) lookupEntryWithLock(hash id.Hash) (*Entry, bool) {
	elem, exists := m.entries[hash]
	if !exists {
		return nil, false
	}
	return elem.Value.(hashedEntry).entry, true
}

func (m *metricMap) purgeExpired(now time.Time, entries []hashedEntry) int64 {
	if len(entries) == 0 {
		return 0
	}
	var numExpired int64
	m.Lock()
	for i := range entries {
		if entries[i].entry.TryExpire(now) {
			elem := m.entries[entries[i].idHash]
			delete(m.entries, entries[i].idHash)
			elem.Value = nil
			m.entryList.Remove(elem)
			numExpired++
		}
	}
	m.Unlock()
	return numExpired
}

func (m *metricMap) reportMetrics() {
	m.RLock()
	numEntries := m.entryList.Len()
	m.RUnlock()

	m.metrics.numEntries.Update(float64(numEntries))
	m.metricLists.reportMetrics()
}
