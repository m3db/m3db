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

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package aggregator

import (
	"fmt"
	"math"
	"sync"
	"time"

	maggregation "github.com/m3db/m3/src/metrics/aggregation"
	"github.com/m3db/m3/src/metrics/metric/id"
	"github.com/m3db/m3/src/metrics/metric/unaggregated"
	"github.com/m3db/m3/src/metrics/pipeline/applied"
	"github.com/m3db/m3/src/metrics/policy"
	"github.com/m3db/m3/src/metrics/transformation"

	"github.com/willf/bitset"
)

type lockedCounterAggregation struct {
	sync.Mutex

	closed      bool
	sourcesSeen *bitset.BitSet
	aggregation counterAggregation
}

type timedCounter struct {
	startAtNanos int64 // start time of an aggregation window
	lockedAgg    *lockedCounterAggregation
}

func (ta *timedCounter) Reset() {
	ta.startAtNanos = 0
	ta.lockedAgg = nil
}

// CounterElem is an element storing time-bucketed aggregations.
type CounterElem struct {
	elemBase
	counterElemBase

	values              []timedCounter             // metric aggregations sorted by time in ascending order
	toConsume           []timedCounter             // small buffer to avoid memory allocations during consumption
	lastConsumedAtNanos int64                      // last consumed at in Unix nanoseconds
	lastConsumedValues  []transformation.Datapoint // last consumed values
}

// NewCounterElem creates a new element for the given metric type.
func NewCounterElem(
	id id.RawID,
	sp policy.StoragePolicy,
	aggTypes maggregation.Types,
	pipeline applied.Pipeline,
	numForwardedTimes int,
	idPrefixSuffixType IDPrefixSuffixType,
	opts Options,
) (*CounterElem, error) {
	e := &CounterElem{
		elemBase: newElemBase(opts),
		values:   make([]timedCounter, 0, defaultNumAggregations), // in most cases values will have two entries
	}
	if err := e.ResetSetData(id, sp, aggTypes, pipeline, numForwardedTimes, idPrefixSuffixType); err != nil {
		return nil, err
	}
	return e, nil
}

// MustNewCounterElem creates a new element, or panics if the input is invalid.
func MustNewCounterElem(
	id id.RawID,
	sp policy.StoragePolicy,
	aggTypes maggregation.Types,
	pipeline applied.Pipeline,
	numForwardedTimes int,
	idPrefixSuffixType IDPrefixSuffixType,
	opts Options,
) *CounterElem {
	elem, err := NewCounterElem(id, sp, aggTypes, pipeline, numForwardedTimes, idPrefixSuffixType, opts)
	if err != nil {
		panic(fmt.Errorf("unable to create element: %v", err))
	}
	return elem
}

// ResetSetData resets the element and sets data.
func (e *CounterElem) ResetSetData(
	id id.RawID,
	sp policy.StoragePolicy,
	aggTypes maggregation.Types,
	pipeline applied.Pipeline,
	numForwardedTimes int,
	idPrefixSuffixType IDPrefixSuffixType,
) error {
	useDefaultAggregation := aggTypes.IsDefault()
	if useDefaultAggregation {
		aggTypes = e.DefaultAggregationTypes(e.aggTypesOpts)
	}
	if err := e.elemBase.resetSetData(id, sp, aggTypes, useDefaultAggregation, pipeline, numForwardedTimes, idPrefixSuffixType); err != nil {
		return err
	}
	if err := e.counterElemBase.ResetSetData(e.aggTypesOpts, aggTypes, useDefaultAggregation); err != nil {
		return err
	}
	// If the pipeline contains derivative transformations, we need to store past
	// values in order to compute the derivatives.
	if !e.parsedPipeline.HasDerivativeTransform {
		return nil
	}
	numAggTypes := len(e.aggTypes)
	if cap(e.lastConsumedValues) < numAggTypes {
		e.lastConsumedValues = make([]transformation.Datapoint, numAggTypes)
	}
	e.lastConsumedValues = e.lastConsumedValues[:numAggTypes]
	for i := 0; i < len(e.lastConsumedValues); i++ {
		e.lastConsumedValues[i] = transformation.Datapoint{Value: nan}
	}
	return nil
}

// AddUnion adds a metric value union at a given timestamp.
func (e *CounterElem) AddUnion(timestamp time.Time, mu unaggregated.MetricUnion) error {
	alignedStart := timestamp.Truncate(e.sp.Resolution().Window).UnixNano()
	lockedAgg, err := e.findOrCreate(alignedStart, createAggregationOptions{})
	if err != nil {
		return err
	}
	lockedAgg.Lock()
	if lockedAgg.closed {
		lockedAgg.Unlock()
		return errAggregationClosed
	}
	lockedAgg.aggregation.AddUnion(timestamp, mu)
	lockedAgg.Unlock()
	return nil
}

// AddValue adds a metric value at a given timestamp.
func (e *CounterElem) AddValue(timestamp time.Time, value float64) error {
	alignedStart := timestamp.Truncate(e.sp.Resolution().Window).UnixNano()
	lockedAgg, err := e.findOrCreate(alignedStart, createAggregationOptions{})
	if err != nil {
		return err
	}
	lockedAgg.Lock()
	if lockedAgg.closed {
		lockedAgg.Unlock()
		return errAggregationClosed
	}
	lockedAgg.aggregation.Add(timestamp, value)
	lockedAgg.Unlock()
	return nil
}

// AddUnique adds a metric value from a given source at a given timestamp.
// If previous values from the same source have already been added to the
// same aggregation, the incoming value is discarded.
func (e *CounterElem) AddUnique(timestamp time.Time, values []float64, sourceID uint32) error {
	alignedStart := timestamp.Truncate(e.sp.Resolution().Window).UnixNano()
	lockedAgg, err := e.findOrCreate(alignedStart, createAggregationOptions{initSourceSet: true})
	if err != nil {
		return err
	}
	lockedAgg.Lock()
	if lockedAgg.closed {
		lockedAgg.Unlock()
		return errAggregationClosed
	}
	source := uint(sourceID)
	if lockedAgg.sourcesSeen.Test(source) {
		lockedAgg.Unlock()
		return errDuplicateForwardingSource
	}
	lockedAgg.sourcesSeen.Set(source)
	for _, v := range values {
		lockedAgg.aggregation.Add(timestamp, v)
	}
	lockedAgg.Unlock()
	return nil
}

// Consume consumes values before a given time and removes them from the element
// after they are consumed, returning whether the element can be collected after
// the consumption is completed.
// NB: Consume is not thread-safe and must be called within a single goroutine
// to avoid race conditions.
func (e *CounterElem) Consume(
	targetNanos int64,
	isEarlierThanFn isEarlierThanFn,
	timestampNanosFn timestampNanosFn,
	flushLocalFn flushLocalMetricFn,
	flushForwardedFn flushForwardedMetricFn,
	onForwardedFlushedFn onForwardingElemFlushedFn,
) bool {
	resolution := e.sp.Resolution().Window
	e.Lock()
	if e.closed {
		e.Unlock()
		return false
	}
	idx := 0
	for range e.values {
		// Bail as soon as the timestamp is no later than the target time.
		if !isEarlierThanFn(e.values[idx].startAtNanos, resolution, targetNanos) {
			break
		}
		idx++
	}
	e.toConsume = e.toConsume[:0]
	if idx > 0 {
		// Shift remaining values to the left and shrink the values slice.
		e.toConsume = append(e.toConsume, e.values[:idx]...)
		n := copy(e.values[0:], e.values[idx:])
		// Clear out the invalid items to avoid holding references to objects
		// for reduced GC overhead..
		for i := n; i < len(e.values); i++ {
			e.values[i].Reset()
		}
		e.values = e.values[:n]
	}
	canCollect := len(e.values) == 0 && e.tombstoned
	e.Unlock()

	// Process the aggregations that are ready for consumption.
	for i := range e.toConsume {
		timeNanos := timestampNanosFn(e.toConsume[i].startAtNanos, resolution)
		e.toConsume[i].lockedAgg.Lock()
		e.processValueWithAggregationLock(timeNanos, e.toConsume[i].lockedAgg, flushLocalFn, flushForwardedFn)
		// Closes the aggregation object after it's processed.
		e.toConsume[i].lockedAgg.closed = true
		e.toConsume[i].lockedAgg.aggregation.Close()
		if e.toConsume[i].lockedAgg.sourcesSeen != nil {
			e.cachedSourceSetsLock.Lock()
			// This is to make sure there aren't too many cached source sets taking up
			// too much space.
			if len(e.cachedSourceSets) < e.opts.MaxNumCachedSourceSets() {
				e.cachedSourceSets = append(e.cachedSourceSets, e.toConsume[i].lockedAgg.sourcesSeen)
			}
			e.cachedSourceSetsLock.Unlock()
			e.toConsume[i].lockedAgg.sourcesSeen = nil
		}
		e.toConsume[i].lockedAgg.Unlock()
		e.toConsume[i].Reset()
	}

	if e.parsedPipeline.HasRollup {
		forwardedAggregationKey, _ := e.ForwardedAggregationKey()
		onForwardedFlushedFn(e.onForwardedAggregationWrittenFn, forwardedAggregationKey)
	}

	return canCollect
}

// Close closes the element.
func (e *CounterElem) Close() {
	e.Lock()
	if e.closed {
		e.Unlock()
		return
	}
	e.closed = true
	e.id = nil
	e.parsedPipeline = parsedPipeline{}
	e.writeForwardedMetricFn = nil
	e.onForwardedAggregationWrittenFn = nil
	for idx := range e.cachedSourceSets {
		e.cachedSourceSets[idx] = nil
	}
	e.cachedSourceSets = nil
	for idx := range e.values {
		// Close the underlying aggregation objects.
		e.values[idx].lockedAgg.sourcesSeen = nil
		e.values[idx].lockedAgg.aggregation.Close()
		e.values[idx].Reset()
	}
	e.values = e.values[:0]
	e.toConsume = e.toConsume[:0]
	e.lastConsumedValues = e.lastConsumedValues[:0]
	e.counterElemBase.Close()
	aggTypesPool := e.aggTypesOpts.TypesPool()
	pool := e.ElemPool(e.opts)
	e.Unlock()

	if !e.useDefaultAggregation {
		aggTypesPool.Put(e.aggTypes)
	}
	pool.Put(e)
}

// findOrCreate finds the aggregation for a given time, or creates one
// if it doesn't exist.
func (e *CounterElem) findOrCreate(
	alignedStart int64,
	createOpts createAggregationOptions,
) (*lockedCounterAggregation, error) {
	e.RLock()
	if e.closed {
		e.RUnlock()
		return nil, errElemClosed
	}
	idx, found := e.indexOfWithLock(alignedStart)
	if found {
		agg := e.values[idx].lockedAgg
		e.RUnlock()
		return agg, nil
	}
	e.RUnlock()

	e.Lock()
	if e.closed {
		e.Unlock()
		return nil, errElemClosed
	}
	idx, found = e.indexOfWithLock(alignedStart)
	if found {
		agg := e.values[idx].lockedAgg
		e.Unlock()
		return agg, nil
	}

	// If not found, create a new aggregation.
	numValues := len(e.values)
	e.values = append(e.values, timedCounter{})
	copy(e.values[idx+1:numValues+1], e.values[idx:numValues])

	var sourcesSeen *bitset.BitSet
	if createOpts.initSourceSet {
		e.cachedSourceSetsLock.Lock()
		if numCachedSourceSets := len(e.cachedSourceSets); numCachedSourceSets > 0 {
			sourcesSeen = e.cachedSourceSets[numCachedSourceSets-1]
			e.cachedSourceSets[numCachedSourceSets-1] = nil
			e.cachedSourceSets = e.cachedSourceSets[:numCachedSourceSets-1]
			sourcesSeen.ClearAll()
		} else {
			sourcesSeen = bitset.New(defaultNumSources)
		}
		e.cachedSourceSetsLock.Unlock()
	}
	e.values[idx] = timedCounter{
		startAtNanos: alignedStart,
		lockedAgg: &lockedCounterAggregation{
			sourcesSeen: sourcesSeen,
			aggregation: e.NewAggregation(e.opts, e.aggOpts),
		},
	}
	agg := e.values[idx].lockedAgg
	e.Unlock()
	return agg, nil
}

// indexOfWithLock finds the smallest element index whose timestamp
// is no smaller than the start time passed in, and true if it's an
// exact match, false otherwise.
func (e *CounterElem) indexOfWithLock(alignedStart int64) (int, bool) {
	numValues := len(e.values)
	// Optimize for the common case.
	if numValues > 0 && e.values[numValues-1].startAtNanos == alignedStart {
		return numValues - 1, true
	}
	// Binary search for the unusual case. We intentionally do not
	// use the sort.Search() function because it requires passing
	// in a closure.
	left, right := 0, numValues
	for left < right {
		mid := left + (right-left)/2 // avoid overflow
		if e.values[mid].startAtNanos < alignedStart {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// If the current timestamp is equal to or larger than the target time,
	// return the index as is.
	if left < numValues && e.values[left].startAtNanos == alignedStart {
		return left, true
	}
	return left, false
}

func (e *CounterElem) processValueWithAggregationLock(
	timeNanos int64,
	lockedAgg *lockedCounterAggregation,
	flushLocalFn flushLocalMetricFn,
	flushForwardedFn flushForwardedMetricFn,
) {
	var (
		transformations  = e.parsedPipeline.Transformations
		discardNaNValues = e.opts.DiscardNaNAggregatedValues()
	)
	for aggTypeIdx, aggType := range e.aggTypes {
		toFlush := make([]transformation.Datapoint, 0)
		value := lockedAgg.aggregation.ValueOf(aggType)
		for _, transformOp := range transformations {

			unaryOp, isUnaryOp := transformOp.UnaryTransform()
			binaryOp, isBinaryOp := transformOp.BinaryTransform()
			unaryMultiOp, isUnaryMultiOp := transformOp.UnaryMultiOutputTransform()
			switch {
			case isUnaryOp:
				curr := transformation.Datapoint{
					TimeNanos: timeNanos,
					Value:     value,
				}

				res := unaryOp.Evaluate(curr)

				value = res.Value

			case isBinaryOp:
				lastTimeNanos := e.lastConsumedAtNanos
				prev := transformation.Datapoint{
					TimeNanos: lastTimeNanos,
					Value:     e.lastConsumedValues[aggTypeIdx].Value,
				}

				currTimeNanos := timeNanos
				curr := transformation.Datapoint{
					TimeNanos: currTimeNanos,
					Value:     value,
				}

				res := binaryOp.Evaluate(prev, curr)

				// NB: we only need to record the value needed for derivative transformations.
				// We currently only support first-order derivative transformations so we only
				// need to keep one value. In the future if we need to support higher-order
				// derivative transformations, we need to store an array of values here.
				if !math.IsNaN(curr.Value) {
					e.lastConsumedValues[aggTypeIdx] = curr
				}

				value = res.Value
			case isUnaryMultiOp:
				curr := transformation.Datapoint{
					TimeNanos: timeNanos,
					Value:     value,
				}
				res, others := unaryMultiOp.Evaluate(curr)
				for _, o := range others {
					toFlush = append(toFlush, o)
				}
				value = res.Value
			}
		}

		if discardNaNValues && math.IsNaN(value) {
			continue
		}

		if !e.parsedPipeline.HasRollup {
			toFlush = append(toFlush, transformation.Datapoint{
				TimeNanos: timeNanos,
				Value:     value,
			})
			for _, point := range toFlush {
				switch e.idPrefixSuffixType {
				case NoPrefixNoSuffix:
					flushLocalFn(nil, e.id, nil, point.TimeNanos, point.Value, e.sp)
				case WithPrefixWithSuffix:
					flushLocalFn(e.FullPrefix(e.opts), e.id, e.TypeStringFor(e.aggTypesOpts, aggType),
						point.TimeNanos, point.Value, e.sp)
				}
			}
		} else {
			forwardedAggregationKey, _ := e.ForwardedAggregationKey()
			flushForwardedFn(e.writeForwardedMetricFn, forwardedAggregationKey, timeNanos, value)
		}
	}
	e.lastConsumedAtNanos = timeNanos
}
