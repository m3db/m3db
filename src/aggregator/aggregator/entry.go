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
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/m3db/m3aggregator/bitset"
	"github.com/m3db/m3aggregator/rate"
	"github.com/m3db/m3aggregator/runtime"
	"github.com/m3db/m3metrics/aggregation"
	"github.com/m3db/m3metrics/metadata"
	"github.com/m3db/m3metrics/metric"
	"github.com/m3db/m3metrics/metric/aggregated"
	metricid "github.com/m3db/m3metrics/metric/id"
	"github.com/m3db/m3metrics/metric/unaggregated"
	"github.com/m3db/m3metrics/op/applied"
	"github.com/m3db/m3metrics/policy"
	xerrors "github.com/m3db/m3x/errors"

	"github.com/uber-go/tally"
)

const (
	// initialAggregationCapacity is the initial number of slots
	// allocated for aggregation metadata.
	initialAggregationCapacity = 2
)

var (
	errEntryClosed                 = errors.New("entry is closed")
	errWriteValueRateLimitExceeded = errors.New("write value rate limit is exceeded")
	errEmptyMetadatas              = errors.New("empty metadata list")
	errNoApplicableMetadata        = errors.New("no applicable metadata")
	errNoPipelinesInMetadata       = errors.New("no pipelines in metadata")
	errArrivedTooLate              = errors.New("arrived too late")
)

type rateLimitEntryMetrics struct {
	valueRateLimitExceeded tally.Counter
	droppedValues          tally.Counter
}

func newRateLimitEntryMetrics(scope tally.Scope) rateLimitEntryMetrics {
	return rateLimitEntryMetrics{
		valueRateLimitExceeded: scope.Counter("value-rate-limit-exceeded"),
		droppedValues:          scope.Counter("dropped-values"),
	}
}

type untimedEntryMetrics struct {
	rateLimit               rateLimitEntryMetrics
	emptyMetadatas          tally.Counter
	noApplicableMetadata    tally.Counter
	noPipelinesInMetadata   tally.Counter
	emptyPipeline           tally.Counter
	noAggregationInPipeline tally.Counter
	staleMetadata           tally.Counter
	tombstonedMetadata      tally.Counter
	metadatasUpdates        tally.Counter
}

func newUntimedEntryMetrics(scope tally.Scope) untimedEntryMetrics {
	return untimedEntryMetrics{
		rateLimit:               newRateLimitEntryMetrics(scope),
		emptyMetadatas:          scope.Counter("empty-metadatas"),
		noApplicableMetadata:    scope.Counter("no-applicable-metadata"),
		noPipelinesInMetadata:   scope.Counter("no-pipelines-in-metadata"),
		emptyPipeline:           scope.Counter("empty-pipeline"),
		noAggregationInPipeline: scope.Counter("no-aggregation-in-pipeline"),
		staleMetadata:           scope.Counter("stale-metadata"),
		tombstonedMetadata:      scope.Counter("tombstoned-metadata"),
		metadatasUpdates:        scope.Counter("metadatas-updates"),
	}
}

type forwardedEntryMetrics struct {
	rateLimit        rateLimitEntryMetrics
	arrivedTooLate   tally.Counter
	duplicateSources tally.Counter
	metadataUpdates  tally.Counter
}

func newForwardedEntryMetrics(scope tally.Scope) forwardedEntryMetrics {
	return forwardedEntryMetrics{
		rateLimit:        newRateLimitEntryMetrics(scope),
		arrivedTooLate:   scope.Counter("arrived-too-late"),
		duplicateSources: scope.Counter("duplicate-sources"),
		metadataUpdates:  scope.Counter("metadata-updates"),
	}
}

type entryMetrics struct {
	untimed   untimedEntryMetrics
	forwarded forwardedEntryMetrics
}

func newEntryMetrics(scope tally.Scope) entryMetrics {
	untimedEntryScope := scope.Tagged(map[string]string{"entry-type": "untimed"})
	forwardedEntryScope := scope.Tagged(map[string]string{"entry-type": "forwarded"})
	return entryMetrics{
		untimed:   newUntimedEntryMetrics(untimedEntryScope),
		forwarded: newForwardedEntryMetrics(forwardedEntryScope),
	}
}

// Entry keeps track of a metric's aggregations alongside the aggregation
// metadatas including storage policies, aggregation types, and remaining pipeline
// steps if any.
type Entry struct {
	sync.RWMutex

	closed              bool
	opts                Options
	rateLimiter         *rate.Limiter
	hasDefaultMetadatas bool
	cutoverNanos        int64
	lists               *metricLists
	numWriters          int32
	lastAccessNanos     int64
	aggregations        aggregationValues
	metrics             entryMetrics
	// The entry keeps a decompressor to reuse the bitset in it, so we can
	// save some heap allocations.
	decompressor aggregation.IDDecompressor
}

// NewEntry creates a new entry.
func NewEntry(lists *metricLists, runtimeOpts runtime.Options, opts Options) *Entry {
	scope := opts.InstrumentOptions().MetricsScope().SubScope("entry")
	e := &Entry{
		aggregations: make(aggregationValues, 0, initialAggregationCapacity),
		metrics:      newEntryMetrics(scope),
		decompressor: aggregation.NewPooledIDDecompressor(opts.AggregationTypesOptions().TypesPool()),
	}
	e.ResetSetData(lists, runtimeOpts, opts)
	return e
}

// IncWriter increases the writer count.
func (e *Entry) IncWriter() { atomic.AddInt32(&e.numWriters, 1) }

// DecWriter decreases the writer count.
func (e *Entry) DecWriter() { atomic.AddInt32(&e.numWriters, -1) }

// ResetSetData resets the entry and sets initial data.
// NB(xichen): we need to reset the options here to use the correct
// time lock contained in the options.
func (e *Entry) ResetSetData(lists *metricLists, runtimeOpts runtime.Options, opts Options) {
	e.Lock()
	e.closed = false
	e.opts = opts
	e.resetRateLimiterWithLock(runtimeOpts)
	e.hasDefaultMetadatas = false
	e.cutoverNanos = uninitializedCutoverNanos
	e.lists = lists
	e.numWriters = 0
	e.recordLastAccessed(e.opts.ClockOptions().NowFn()())
	e.Unlock()
}

// SetRuntimeOptions updates the parameters of the rate limiter.
func (e *Entry) SetRuntimeOptions(opts runtime.Options) {
	e.Lock()
	if e.closed {
		e.Unlock()
		return
	}
	e.resetRateLimiterWithLock(opts)
	e.Unlock()
}

// AddUntimed adds an untimed metric along with its metadatas.
func (e *Entry) AddUntimed(
	metricUnion unaggregated.MetricUnion,
	metadatas metadata.StagedMetadatas,
) error {
	switch metricUnion.Type {
	case metric.TimerType:
		var err error
		if err = e.applyValueRateLimit(
			int64(len(metricUnion.BatchTimerVal)),
			e.metrics.untimed.rateLimit,
		); err == nil {
			err = e.writeBatchTimerWithMetadatas(metricUnion, metadatas)
		}
		if metricUnion.BatchTimerVal != nil && metricUnion.TimerValPool != nil {
			metricUnion.TimerValPool.Put(metricUnion.BatchTimerVal)
		}
		return err
	default:
		// For counters and gauges, there is a single value in the metric union.
		if err := e.applyValueRateLimit(1, e.metrics.untimed.rateLimit); err != nil {
			return err
		}
		return e.addUntimed(metricUnion, metadatas)
	}
}

// AddForwarded adds a forwarded metric alongside its metadata.
func (e *Entry) AddForwarded(
	metric aggregated.Metric,
	metadata metadata.ForwardMetadata,
) error {
	if err := e.applyValueRateLimit(1, e.metrics.untimed.rateLimit); err != nil {
		return err
	}
	return e.addForwarded(metric, metadata)
}

// ShouldExpire returns whether the entry should expire.
func (e *Entry) ShouldExpire(now time.Time) bool {
	e.RLock()
	if e.closed {
		e.RUnlock()
		return false
	}
	e.RUnlock()

	return e.shouldExpire(now)
}

// TryExpire attempts to expire the entry, returning true
// if the entry is expired, and false otherwise.
func (e *Entry) TryExpire(now time.Time) bool {
	e.Lock()
	if e.closed {
		e.Unlock()
		return false
	}
	if !e.shouldExpire(now) {
		e.Unlock()
		return false
	}
	e.closed = true
	// Empty out the aggregation elements so they don't hold references
	// to other objects after being put back to pool to reduce GC overhead.
	for i := range e.aggregations {
		e.aggregations[i].elem.Value.(metricElem).MarkAsTombstoned()
		e.aggregations[i] = aggregationValue{}
	}
	e.aggregations = e.aggregations[:0]
	e.lists = nil
	pool := e.opts.EntryPool()
	e.Unlock()

	pool.Put(e)
	return true
}

func (e *Entry) writeBatchTimerWithMetadatas(
	metric unaggregated.MetricUnion,
	metadatas metadata.StagedMetadatas,
) error {
	// If there is no limit on the maximum batch size per write, write
	// all timers at once.
	maxTimerBatchSizePerWrite := e.opts.MaxTimerBatchSizePerWrite()
	if maxTimerBatchSizePerWrite == 0 {
		return e.addUntimed(metric, metadatas)
	}

	// Otherwise, honor maximum timer batch size.
	var (
		timerValues    = metric.BatchTimerVal
		numTimerValues = len(timerValues)
		start, end     int
	)
	for start = 0; start < numTimerValues; start = end {
		end = start + maxTimerBatchSizePerWrite
		if end > numTimerValues {
			end = numTimerValues
		}
		splitTimer := metric
		splitTimer.BatchTimerVal = timerValues[start:end]
		if err := e.addUntimed(splitTimer, metadatas); err != nil {
			return err
		}
	}
	return nil
}

func (e *Entry) addUntimed(
	metric unaggregated.MetricUnion,
	metadatas metadata.StagedMetadatas,
) error {
	timeLock := e.opts.TimeLock()
	timeLock.RLock()

	// NB(xichen): it is important that we determine the current time
	// within the time lock. This ensures time ordering by wrapping
	// actions that need to happen before a given time within a read lock,
	// so it is guaranteed that actions before when a write lock is acquired
	// must have all completed. This is used to ensure we never write metrics
	// for times that have already been flushed.
	currTime := e.opts.ClockOptions().NowFn()()
	e.recordLastAccessed(currTime)

	e.RLock()
	if e.closed {
		e.RUnlock()
		timeLock.RUnlock()
		return errEntryClosed
	}

	// Fast exit path for the common case where the metric has default metadatas for aggregation.
	hasDefaultMetadatas := metadatas.IsDefault()
	if e.hasDefaultMetadatas && hasDefaultMetadatas {
		err := e.addUntimedWithLock(currTime, metric)
		e.RUnlock()
		timeLock.RUnlock()
		return err
	}

	sm, err := e.activeStagedMetadataWithLock(currTime, metadatas)
	if err != nil {
		e.RUnlock()
		timeLock.RUnlock()
		return err
	}

	// If the metadata indicates the (rollup) metric has been tombstoned, the metric is
	// not ingested for aggregation. However, we do not update the policies asssociated
	// with this entry and mark it tombstoned because there may be a different raw metric
	// generating this same (rollup) metric that is actively emitting, meaning this entry
	// may still be very much alive.
	if sm.Tombstoned {
		e.RUnlock()
		timeLock.RUnlock()
		e.metrics.untimed.tombstonedMetadata.Inc(1)
		return nil
	}

	// It is expected that there is at least one pipeline in the metadata.
	if len(sm.Pipelines) == 0 {
		e.RUnlock()
		timeLock.RUnlock()
		e.metrics.untimed.noPipelinesInMetadata.Inc(1)
		return errNoPipelinesInMetadata
	}

	if !e.shouldUpdateStagedMetadatasWithLock(sm) {
		err = e.addUntimedWithLock(currTime, metric)
		e.RUnlock()
		timeLock.RUnlock()
		return err
	}
	e.RUnlock()

	e.Lock()
	if e.closed {
		e.Unlock()
		timeLock.RUnlock()
		return errEntryClosed
	}

	if e.shouldUpdateStagedMetadatasWithLock(sm) {
		if err = e.updateStagedMetadatasWithLock(metric, hasDefaultMetadatas, sm); err != nil {
			// NB(xichen): if an error occurred during policy update, the policies
			// will remain as they are, i.e., there are no half-updated policies.
			e.Unlock()
			timeLock.RUnlock()
			return err
		}
	}

	err = e.addUntimedWithLock(currTime, metric)
	e.Unlock()
	timeLock.RUnlock()

	return err
}

// NB(xichen): we assume the metadatas are sorted by their cutover times
// in ascending order.
func (e *Entry) activeStagedMetadataWithLock(
	t time.Time,
	metadatas metadata.StagedMetadatas,
) (metadata.StagedMetadata, error) {
	// If we have no metadata to apply, simply bail.
	if len(metadatas) == 0 {
		e.metrics.untimed.emptyMetadatas.Inc(1)
		return metadata.DefaultStagedMetadata, errEmptyMetadatas
	}
	timeNanos := t.UnixNano()
	for idx := len(metadatas) - 1; idx >= 0; idx-- {
		if metadatas[idx].CutoverNanos <= timeNanos {
			return metadatas[idx], nil
		}
	}
	e.metrics.untimed.noApplicableMetadata.Inc(1)
	return metadata.DefaultStagedMetadata, errNoApplicableMetadata
}

// NB: The metadata passed in is guaranteed to have cut over based on the current time.
func (e *Entry) shouldUpdateStagedMetadatasWithLock(sm metadata.StagedMetadata) bool {
	// If this is a stale metadata, we don't update the existing metadata.
	if e.cutoverNanos > sm.CutoverNanos {
		e.metrics.untimed.staleMetadata.Inc(1)
		return false
	}

	// If this is a newer metadata, we always update.
	if e.cutoverNanos < sm.CutoverNanos {
		return true
	}

	// Iterate over the list of pipelines and check whether we have metadata changes.
	// NB: If the incoming metadata have the same set of aggregation keys as the cached
	// set but also have duplicates, there is no need to update metadatas as long as
	// the cached set has all aggregation keys in the incoming metadata and vice versa.
	bs := bitset.New(uint(len(e.aggregations)))
	for _, pipeline := range sm.Pipelines {
		storagePolicies := e.storagePolicies(pipeline.StoragePolicies)
		for _, storagePolicy := range storagePolicies {
			key := aggregationKey{
				aggregationID: pipeline.AggregationID,
				storagePolicy: storagePolicy,
				pipeline:      pipeline.Pipeline,
			}
			idx := e.aggregations.index(key)
			if idx < 0 {
				return true
			}
			bs.Set(uint(idx))
		}
	}
	return !bs.All(uint(len(e.aggregations)))
}

func (e *Entry) storagePolicies(policies []policy.StoragePolicy) []policy.StoragePolicy {
	if !policy.IsDefaultStoragePolicies(policies) {
		return policies
	}
	return e.opts.DefaultStoragePolicies()
}

func (e *Entry) maybeCopyIDWithLock(id metricid.RawID) metricid.RawID {
	// If there are existing elements for this id, try reusing
	// the id from the elements because those are owned by us.
	if len(e.aggregations) > 0 {
		return e.aggregations[0].elem.Value.(metricElem).ID()
	}

	// Otherwise it is necessary to make a copy because it's not owned by us.
	elemID := make(metricid.RawID, len(id))
	copy(elemID, id)
	return elemID
}

// addAggregationKey adds a new aggregation key to the list of new aggregations.
func (e *Entry) addNewAggregationKey(
	metricType metric.Type,
	metricID metricid.RawID,
	key aggregationKey,
	listID metricListID,
	newAggregations aggregationValues,
) (aggregationValues, error) {
	// Remove duplicate aggregation pipelines.
	if newAggregations.contains(key) {
		return newAggregations, nil
	}
	if idx := e.aggregations.index(key); idx >= 0 {
		newAggregations = append(newAggregations, e.aggregations[idx])
		return newAggregations, nil
	}
	aggTypes, err := e.decompressor.Decompress(key.aggregationID)
	if err != nil {
		return nil, err
	}
	var newElem metricElem
	switch metricType {
	case metric.CounterType:
		newElem = e.opts.CounterElemPool().Get()
	case metric.TimerType:
		newElem = e.opts.TimerElemPool().Get()
	case metric.GaugeType:
		newElem = e.opts.GaugeElemPool().Get()
	default:
		return nil, errInvalidMetricType
	}
	// NB: The pipeline may not be owned by us and as such we need to make a copy here.
	key.pipeline = key.pipeline.Clone()
	if err = newElem.ResetSetData(metricID, key.storagePolicy, aggTypes, key.pipeline, key.numForwardedTimes); err != nil {
		return nil, err
	}
	list, err := e.lists.FindOrCreate(listID)
	if err != nil {
		return nil, err
	}
	newListElem, err := list.PushBack(newElem)
	if err != nil {
		return nil, err
	}
	newAggregations = append(newAggregations, aggregationValue{key: key, elem: newListElem})
	return newAggregations, nil
}

func (e *Entry) removeOldAggregations(newAggregations aggregationValues) {
	for _, val := range e.aggregations {
		if !newAggregations.contains(val.key) {
			val.elem.Value.(metricElem).MarkAsTombstoned()
		}
	}
}

func (e *Entry) updateStagedMetadatasWithLock(
	metric unaggregated.MetricUnion,
	hasDefaultMetadatas bool,
	sm metadata.StagedMetadata,
) error {
	var (
		elemID          = e.maybeCopyIDWithLock(metric.ID)
		newAggregations = make(aggregationValues, 0, initialAggregationCapacity)
	)

	// Update the metadatas.
	for _, pipeline := range sm.Pipelines {
		storagePolicies := e.storagePolicies(pipeline.StoragePolicies)
		for _, storagePolicy := range storagePolicies {
			key := aggregationKey{
				aggregationID: pipeline.AggregationID,
				storagePolicy: storagePolicy,
				pipeline:      pipeline.Pipeline,
			}
			listID := standardMetricListID{
				resolution: storagePolicy.Resolution().Window,
			}.toMetricListID()
			var err error
			newAggregations, err = e.addNewAggregationKey(metric.Type, elemID, key, listID, newAggregations)
			if err != nil {
				return err
			}
		}
	}

	// Mark the outdated elements as tombstoned.
	e.removeOldAggregations(newAggregations)

	// Replace the existing aggregations with new aggregations.
	e.aggregations = newAggregations
	e.hasDefaultMetadatas = hasDefaultMetadatas
	e.cutoverNanos = sm.CutoverNanos
	e.metrics.untimed.metadatasUpdates.Inc(1)

	return nil
}

func (e *Entry) addUntimedWithLock(timestamp time.Time, mu unaggregated.MetricUnion) error {
	multiErr := xerrors.NewMultiError()
	for _, val := range e.aggregations {
		if err := val.elem.Value.(metricElem).AddUnion(timestamp, mu); err != nil {
			multiErr = multiErr.Add(err)
		}
	}
	return multiErr.FinalError()
}

func (e *Entry) addForwarded(
	metric aggregated.Metric,
	metadata metadata.ForwardMetadata,
) error {
	timeLock := e.opts.TimeLock()
	timeLock.RLock()

	// NB(xichen): it is important that we determine the current time
	// within the time lock. This ensures time ordering by wrapping
	// actions that need to happen before a given time within a read lock,
	// so it is guaranteed that actions before when a write lock is acquired
	// must have all completed. This is used to ensure we never write metrics
	// for times that have already been flushed.
	currTime := e.opts.ClockOptions().NowFn()()
	e.recordLastAccessed(currTime)

	e.RLock()
	if e.closed {
		e.RUnlock()
		timeLock.RUnlock()
		return errEntryClosed
	}

	// Reject datapoints that arrive too late.
	if err := e.checkLateness(
		currTime.UnixNano(),
		metric.TimeNanos,
		metadata.StoragePolicy.Resolution().Window,
		metadata.NumForwardedTimes,
	); err != nil {
		e.RUnlock()
		timeLock.RUnlock()
		return err
	}

	// Check if we should update metadata, and add metric if not.
	if !e.shouldUpdateForwardMetadataWithLock(metadata) {
		err := e.addForwardedWithLock(metric, metadata.SourceID)
		e.RUnlock()
		timeLock.RUnlock()
		return err
	}
	e.RUnlock()

	e.Lock()
	if e.closed {
		e.Unlock()
		timeLock.RUnlock()
		return errEntryClosed
	}

	// Update metatadata, and add metric.
	if e.shouldUpdateForwardMetadataWithLock(metadata) {
		if err := e.updateForwardMetadataWithLock(metric, metadata); err != nil {
			e.Unlock()
			timeLock.RUnlock()
			return err
		}
	}

	err := e.addForwardedWithLock(metric, metadata.SourceID)
	e.Unlock()
	timeLock.RUnlock()

	return err
}

func (e *Entry) checkLateness(
	currNanos, metricTimeNanos int64,
	resolution time.Duration,
	numForwardedTimes int,
) error {
	maxAllowedForwardingDelayFn := e.opts.MaxAllowedForwardingDelayFn()
	maxLatenessAllowed := maxAllowedForwardingDelayFn(resolution, numForwardedTimes)
	if currNanos-metricTimeNanos <= maxLatenessAllowed.Nanoseconds() {
		return nil
	}
	e.metrics.forwarded.arrivedTooLate.Inc(1)
	return errArrivedTooLate
}

// NB: For forwarded metrics, an entry must be associated with one and only one metadata.
func (e *Entry) shouldUpdateForwardMetadataWithLock(metadata metadata.ForwardMetadata) bool {
	key := aggregationKey{
		aggregationID:     metadata.AggregationID,
		storagePolicy:     metadata.StoragePolicy,
		pipeline:          metadata.Pipeline,
		numForwardedTimes: metadata.NumForwardedTimes,
	}
	return e.aggregations.index(key) < 0
}

func (e *Entry) updateForwardMetadataWithLock(
	metric aggregated.Metric,
	metadata metadata.ForwardMetadata,
) error {
	var (
		elemID          = e.maybeCopyIDWithLock(metric.ID)
		newAggregations = make(aggregationValues, 0, 1)
		err             error
	)

	// Update the forward metadata.
	key := aggregationKey{
		aggregationID:     metadata.AggregationID,
		storagePolicy:     metadata.StoragePolicy,
		pipeline:          metadata.Pipeline,
		numForwardedTimes: metadata.NumForwardedTimes,
	}
	listID := forwardedMetricListID{
		resolution:        metadata.StoragePolicy.Resolution().Window,
		numForwardedTimes: metadata.NumForwardedTimes,
	}.toMetricListID()
	newAggregations, err = e.addNewAggregationKey(metric.Type, elemID, key, listID, newAggregations)
	if err != nil {
		return err
	}

	// Mark the outdated elements as tombstoned.
	e.removeOldAggregations(newAggregations)

	e.aggregations = newAggregations
	e.metrics.forwarded.metadataUpdates.Inc(1)
	return nil
}

func (e *Entry) addForwardedWithLock(metric aggregated.Metric, sourceID []byte) error {
	var (
		timestamp = time.Unix(0, metric.TimeNanos)
		multiErr  = xerrors.NewMultiError()
	)
	for _, val := range e.aggregations {
		err := val.elem.Value.(metricElem).AddUnique(timestamp, metric.Value, sourceID)
		if err == errDuplicateForwardingSource {
			// Duplicate forwarding sources may occur during a leader re-election and is not
			// considered an external facing error. Hence, we record it and move on.
			e.metrics.forwarded.duplicateSources.Inc(1)
			continue
		}
		if err != nil {
			multiErr = multiErr.Add(err)
		}
	}
	return multiErr.FinalError()
}

func (e *Entry) writerCount() int        { return int(atomic.LoadInt32(&e.numWriters)) }
func (e *Entry) lastAccessed() time.Time { return time.Unix(0, atomic.LoadInt64(&e.lastAccessNanos)) }

func (e *Entry) recordLastAccessed(currTime time.Time) {
	atomic.StoreInt64(&e.lastAccessNanos, currTime.UnixNano())
}

func (e *Entry) shouldExpire(now time.Time) bool {
	// Only expire the entry if there are no active writers
	// and it has reached its ttl since last accessed.
	return e.writerCount() == 0 && now.After(e.lastAccessed().Add(e.opts.EntryTTL()))
}

func (e *Entry) resetRateLimiterWithLock(runtimeOpts runtime.Options) {
	newLimit := runtimeOpts.WriteValuesPerMetricLimitPerSecond()
	if newLimit <= 0 {
		e.rateLimiter = nil
		return
	}
	if e.rateLimiter == nil {
		nowFn := e.opts.ClockOptions().NowFn()
		e.rateLimiter = rate.NewLimiter(newLimit, nowFn)
		return
	}
	e.rateLimiter.Reset(newLimit)
}

func (e *Entry) applyValueRateLimit(numValues int64, m rateLimitEntryMetrics) error {
	e.RLock()
	rateLimiter := e.rateLimiter
	e.RUnlock()
	if rateLimiter == nil {
		return nil
	}
	if rateLimiter.IsAllowed(numValues) {
		return nil
	}
	m.valueRateLimitExceeded.Inc(1)
	m.droppedValues.Inc(numValues)
	return errWriteValueRateLimitExceeded
}

type aggregationKey struct {
	aggregationID     aggregation.ID
	storagePolicy     policy.StoragePolicy
	pipeline          applied.Pipeline
	numForwardedTimes int
}

func (k aggregationKey) Equal(other aggregationKey) bool {
	return k.aggregationID == other.aggregationID &&
		k.storagePolicy == other.storagePolicy &&
		k.pipeline.Equal(other.pipeline) &&
		k.numForwardedTimes == other.numForwardedTimes
}

type aggregationValue struct {
	key  aggregationKey
	elem *list.Element
}

// TODO(xichen): benchmark the performance of using a single slice
// versus a map with a partial key versus a map with a hash of full key.
type aggregationValues []aggregationValue

func (vals aggregationValues) index(k aggregationKey) int {
	for i, val := range vals {
		if val.key.Equal(k) {
			return i
		}
	}
	return -1
}

func (vals aggregationValues) contains(k aggregationKey) bool {
	return vals.index(k) != -1
}
