// Copyright (c) 2017 Uber Technologies, Inc.
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
	"sync"
	"time"

	schema "github.com/m3db/m3aggregator/generated/proto/flush"
	"github.com/m3db/m3x/clock"
	"github.com/m3db/m3x/log"
	xsync "github.com/m3db/m3x/sync"
	"github.com/m3db/m3x/watch"

	"github.com/uber-go/tally"
)

type standardFlusherMetrics struct {
	shardNotFound        tally.Counter
	resolutionNotFound   tally.Counter
	flushWindowsNotEnded tally.Counter
}

func newStandardFlusherMetrics(scope tally.Scope) standardFlusherMetrics {
	return standardFlusherMetrics{
		shardNotFound:        scope.Counter("shard-not-found"),
		resolutionNotFound:   scope.Counter("resolution-not-found"),
		flushWindowsNotEnded: scope.Counter("flush-windows-not-ended"),
	}
}

type forwardedFlusherMetrics struct {
	shardNotFound             tally.Counter
	resolutionNotFound        tally.Counter
	nilForwardedTimes         tally.Counter
	numForwardedTimesNotFound tally.Counter
	flushWindowsNotEnded      tally.Counter
}

func newForwardedFlusherMetrics(scope tally.Scope) forwardedFlusherMetrics {
	return forwardedFlusherMetrics{
		shardNotFound:             scope.Counter("shard-not-found"),
		resolutionNotFound:        scope.Counter("resolution-not-found"),
		nilForwardedTimes:         scope.Counter("nil-forwarded-times"),
		numForwardedTimesNotFound: scope.Counter("num-forwarded-times-not-found"),
		flushWindowsNotEnded:      scope.Counter("flush-windows-not-ended"),
	}
}

type followerFlushManagerMetrics struct {
	watchCreateErrors tally.Counter
	kvUpdateFlush     tally.Counter
	forcedFlush       tally.Counter
	notCampaigning    tally.Counter
	standard          standardFlusherMetrics
	forwarded         forwardedFlusherMetrics
}

func newFollowerFlushManagerMetrics(scope tally.Scope) followerFlushManagerMetrics {
	standardScope := scope.Tagged(map[string]string{"flusher-type": "standard"})
	forwardedScope := scope.Tagged(map[string]string{"flusher-type": "forwarded"})
	return followerFlushManagerMetrics{
		watchCreateErrors: scope.Counter("watch-create-errors"),
		kvUpdateFlush:     scope.Counter("kv-update-flush"),
		forcedFlush:       scope.Counter("forced-flush"),
		notCampaigning:    scope.Counter("not-campaigning"),
		standard:          newStandardFlusherMetrics(standardScope),
		forwarded:         newForwardedFlusherMetrics(forwardedScope),
	}
}

type followerFlushManager struct {
	sync.RWMutex
	sync.WaitGroup

	nowFn                 clock.NowFn
	checkEvery            time.Duration
	workers               xsync.WorkerPool
	placementManager      PlacementManager
	electionManager       ElectionManager
	flushTimesManager     FlushTimesManager
	maxBufferSize         time.Duration
	forcedFlushWindowSize time.Duration
	logger                log.Logger
	scope                 tally.Scope

	doneCh          <-chan struct{}
	received        *schema.ShardSetFlushTimes
	processed       *schema.ShardSetFlushTimes
	flushTimesState flushTimesState
	flushMode       followerFlushMode
	lastFlushed     time.Time
	openedAt        time.Time
	flushTask       *followerFlushTask
	sleepFn         sleepFn
	metrics         followerFlushManagerMetrics
}

func newFollowerFlushManager(
	doneCh <-chan struct{},
	opts FlushManagerOptions,
) roleBasedFlushManager {
	nowFn := opts.ClockOptions().NowFn()
	instrumentOpts := opts.InstrumentOptions()
	scope := instrumentOpts.MetricsScope()
	mgr := &followerFlushManager{
		nowFn:                 nowFn,
		checkEvery:            opts.CheckEvery(),
		workers:               opts.WorkerPool(),
		placementManager:      opts.PlacementManager(),
		electionManager:       opts.ElectionManager(),
		flushTimesManager:     opts.FlushTimesManager(),
		maxBufferSize:         opts.MaxBufferSize(),
		forcedFlushWindowSize: opts.ForcedFlushWindowSize(),
		logger:                instrumentOpts.Logger(),
		scope:                 scope,
		doneCh:                doneCh,
		flushTimesState:       flushTimesUninitialized,
		flushMode:             unknownFollowerFlush,
		lastFlushed:           nowFn(),
		sleepFn:               time.Sleep,
		metrics:               newFollowerFlushManagerMetrics(scope),
	}
	mgr.flushTask = &followerFlushTask{mgr: mgr}
	return mgr
}

func (mgr *followerFlushManager) Open() {
	mgr.Lock()
	defer mgr.Unlock()

	mgr.openedAt = mgr.nowFn()
	mgr.Add(1)
	go mgr.watchFlushTimes()
}

// NB(xichen): no actions needed for initializing the follower flush manager.
func (mgr *followerFlushManager) Init([]*flushBucket) {}

func (mgr *followerFlushManager) Prepare(buckets []*flushBucket) (flushTask, time.Duration) {
	// NB(xichen): a flush is triggered in the following scenarios:
	// * The flush times persisted in kv have been updated since last flush, or
	// * Sufficient time (a.k.a. maxBufferSize) has elapsed since last flush.
	mgr.Lock()
	defer mgr.Unlock()

	var (
		now                = mgr.nowFn()
		needsFlush         bool
		flushersByInterval []flushersGroup
	)
	if mgr.flushTimesState == flushTimesUpdated {
		mgr.flushTimesState = flushTimesProcessed
		mgr.processed = mgr.received
		mgr.flushMode = kvUpdateFollowerFlush
		flushersByInterval = mgr.flushersFromKVUpdateWithLock(buckets)
		needsFlush = true
		mgr.metrics.kvUpdateFlush.Inc(1)
	} else {
		durationSinceLastFlush := now.Sub(mgr.lastFlushed)
		// If the follower has accumulated enough data since last flush without receiving a KV
		// update, we enter the forced flush mode.
		if mgr.flushMode != forcedFollowerFlush && durationSinceLastFlush >= mgr.maxBufferSize {
			mgr.flushMode = forcedFollowerFlush
		}
		// Each time the follower flushes the olded data accumulated for a given window size.
		if mgr.flushMode == forcedFollowerFlush && durationSinceLastFlush >= mgr.forcedFlushWindowSize {
			flushBeforeNanos := now.Add(-mgr.maxBufferSize).UnixNano()
			flushersByInterval = mgr.flushersFromForcedFlush(buckets, flushBeforeNanos)
			needsFlush = true
			mgr.metrics.forcedFlush.Inc(1)
		}
	}

	if !needsFlush {
		return nil, mgr.checkEvery
	}
	mgr.lastFlushed = now
	mgr.flushTask.flushersByInterval = flushersByInterval
	return mgr.flushTask, 0
}

// NB(xichen): the follower flush manager flushes data based on the flush times
// stored in kv and does not need to take extra actions when a new bucket is added.
func (mgr *followerFlushManager) OnBucketAdded(int, *flushBucket) {}

// The follower flush manager may only lead if and only if all the following conditions
// are met:
// * The instance is campaigning.
// * All the aggregation windows since the flush manager is opened have ended.
func (mgr *followerFlushManager) CanLead() bool {
	mgr.RLock()
	defer mgr.RUnlock()

	if !mgr.electionManager.IsCampaigning() {
		mgr.metrics.notCampaigning.Inc(1)
		return false
	}
	if mgr.processed == nil {
		return false
	}

	for _, shardFlushTimes := range mgr.processed.ByShard {
		// If the shard is tombstoned, there is no need to examine its flush times.
		if shardFlushTimes.Tombstoned {
			continue
		}

		// Check that for standard metrics, all the open windows containing the process
		// start time are closed, meaning the standard metrics that didn't make to the
		// process have been flushed successfully downstream.
		for windowNanos, lastFlushedNanos := range shardFlushTimes.StandardByResolution {
			windowSize := time.Duration(windowNanos)
			windowEndAt := mgr.openedAt.Truncate(windowSize)
			if windowEndAt.Before(mgr.openedAt) {
				windowEndAt = windowEndAt.Add(windowSize)
			}
			if lastFlushedNanos < windowEndAt.UnixNano() {
				mgr.metrics.standard.flushWindowsNotEnded.Inc(1)
				return false
			}
		}

		// Check that the forwarded metrics have been flushed past the process start
		// time, meaning the forwarded metrics that didn't make to the process have been
		// flushed successfully downstream.
		for windowNanos, fbr := range shardFlushTimes.ForwardedByResolution {
			if fbr == nil {
				mgr.metrics.forwarded.nilForwardedTimes.Inc(1)
				return false
			}
			// Since the timestamps of the forwarded metrics are aligned to the resolution
			// boundaries, we simply need to make sure that all forwarded metrics in or before
			// the window containing the process start time have been flushed to assert that
			// the process can safely take over leadership.
			windowSize := time.Duration(windowNanos)
			waitTillFlushedTime := mgr.openedAt.Truncate(windowSize)
			if waitTillFlushedTime.Equal(mgr.openedAt) {
				waitTillFlushedTime = waitTillFlushedTime.Add(-windowSize)
			}
			for _, lastFlushedNanos := range fbr.ByNumForwardedTimes {
				if lastFlushedNanos <= waitTillFlushedTime.UnixNano() {
					mgr.metrics.forwarded.flushWindowsNotEnded.Inc(1)
					return false
				}
			}
		}
	}

	return true
}

func (mgr *followerFlushManager) Close() { mgr.Wait() }

func (mgr *followerFlushManager) flushersFromKVUpdateWithLock(buckets []*flushBucket) []flushersGroup {
	flushersByInterval := make([]flushersGroup, len(buckets))
	for i, bucket := range buckets {
		bucketID := bucket.bucketID
		flushersByInterval[i].interval = bucket.interval
		flushersByInterval[i].duration = bucket.duration
		switch bucketID.listType {
		case standardMetricListType:
			flushersByInterval[i].flushers = mgr.standardFlushersFromKVUpdateWithLock(bucketID.standard, bucket.flushers)
		case forwardedMetricListType:
			flushersByInterval[i].flushers = mgr.forwardedFlushersFromKVUpdateWithLock(bucketID.forwarded, bucket.flushers)
		default:
			panic("should never get here")
		}
	}
	return flushersByInterval
}

func (mgr *followerFlushManager) standardFlushersFromKVUpdateWithLock(
	listID standardMetricListID,
	flushers []flushingMetricList,
) []flusherWithTime {
	var (
		resolution       = listID.resolution
		flushersWithTime = make([]flusherWithTime, 0, defaultInitialFlushCapacity)
	)
	for _, flusher := range flushers {
		shard := flusher.Shard()
		shardFlushTimes, exists := mgr.received.ByShard[shard]
		if !exists {
			mgr.metrics.standard.shardNotFound.Inc(1)
			mgr.logger.WithFields(
				log.NewField("shard", shard),
			).Warn("shard not found in standard flusher flush times")
			continue
		}
		lastFlushedAtNanos, exists := shardFlushTimes.StandardByResolution[int64(resolution)]
		if !exists {
			mgr.metrics.standard.resolutionNotFound.Inc(1)
			mgr.logger.WithFields(
				log.NewField("shard", shard),
				log.NewField("resolution", resolution.String()),
			).Warn("resolution not found in standard flusher flush times")
			continue
		}
		newFlushTarget := flusherWithTime{
			flusher:          flusher,
			flushBeforeNanos: lastFlushedAtNanos,
		}
		flushersWithTime = append(flushersWithTime, newFlushTarget)
	}
	return flushersWithTime
}

func (mgr *followerFlushManager) forwardedFlushersFromKVUpdateWithLock(
	listID forwardedMetricListID,
	flushers []flushingMetricList,
) []flusherWithTime {
	var (
		resolution        = listID.resolution
		numForwardedTimes = listID.numForwardedTimes
		flushersWithTime  = make([]flusherWithTime, 0, defaultInitialFlushCapacity)
	)
	for _, flusher := range flushers {
		shard := flusher.Shard()
		shardFlushTimes, exists := mgr.received.ByShard[shard]
		if !exists {
			mgr.metrics.forwarded.shardNotFound.Inc(1)
			mgr.logger.WithFields(
				log.NewField("shard", shard),
			).Warn("shard not found in forwarded flusher flush times")
			continue
		}
		flushTimesForResolution, exists := shardFlushTimes.ForwardedByResolution[int64(resolution)]
		if !exists {
			mgr.metrics.forwarded.resolutionNotFound.Inc(1)
			mgr.logger.WithFields(
				log.NewField("shard", shard),
				log.NewField("resolution", resolution.String()),
			).Warn("resolution not found in forwarded flusher flush times")
			continue
		}
		if flushTimesForResolution == nil {
			mgr.metrics.forwarded.nilForwardedTimes.Inc(1)
			mgr.logger.WithFields(
				log.NewField("shard", shard),
				log.NewField("resolution", resolution.String()),
			).Warn("nil forwarded flusher flush times")
			continue
		}
		lastFlushedAtNanos, exists := flushTimesForResolution.ByNumForwardedTimes[int32(numForwardedTimes)]
		if !exists {
			mgr.metrics.forwarded.numForwardedTimesNotFound.Inc(1)
			mgr.logger.WithFields(
				log.NewField("shard", shard),
				log.NewField("resolution", resolution.String()),
				log.NewField("numForwardedTimes", numForwardedTimes),
			).Warn("numForwardedTimes not found in forwarded flush times")
			continue
		}
		newFlushTarget := flusherWithTime{
			flusher:          flusher,
			flushBeforeNanos: lastFlushedAtNanos,
		}
		flushersWithTime = append(flushersWithTime, newFlushTarget)
	}
	return flushersWithTime
}

func (mgr *followerFlushManager) flushersFromForcedFlush(
	buckets []*flushBucket,
	flushBeforeNanos int64,
) []flushersGroup {
	flushersByInterval := make([]flushersGroup, len(buckets))
	for i, bucket := range buckets {
		flushersByInterval[i].interval = bucket.interval
		flushersByInterval[i].duration = bucket.duration
		flushersByInterval[i].flushers = make([]flusherWithTime, 0, defaultInitialFlushCapacity)
		for _, flusher := range bucket.flushers {
			newFlushTarget := flusherWithTime{
				flusher:          flusher,
				flushBeforeNanos: flushBeforeNanos,
			}
			flushersByInterval[i].flushers = append(flushersByInterval[i].flushers, newFlushTarget)
		}
	}
	return flushersByInterval
}

func (mgr *followerFlushManager) watchFlushTimes() {
	defer mgr.Done()

	var (
		throttlePeriod  = time.Second
		flushTimesWatch watch.Watch
		err             error
	)

	for {
		if flushTimesWatch == nil {
			flushTimesWatch, err = mgr.flushTimesManager.Watch()
			if err != nil {
				mgr.metrics.watchCreateErrors.Inc(1)
				mgr.sleepFn(throttlePeriod)
				continue
			}
		}

		select {
		case <-flushTimesWatch.C():
			mgr.Lock()
			mgr.received = flushTimesWatch.Get().(*schema.ShardSetFlushTimes)
			mgr.flushTimesState = flushTimesUpdated
			mgr.Unlock()
		case <-mgr.doneCh:
			return
		}
	}
}

type followerFlushTask struct {
	mgr                *followerFlushManager
	flushersByInterval []flushersGroup
}

func (t *followerFlushTask) Run() {
	var (
		mgr       = t.mgr
		wgWorkers sync.WaitGroup
	)
	for _, group := range t.flushersByInterval {
		start := mgr.nowFn()
		for _, flusherWithTime := range group.flushers {
			flusherWithTime := flusherWithTime
			wgWorkers.Add(1)
			mgr.workers.Go(func() {
				flusherWithTime.flusher.DiscardBefore(flusherWithTime.flushBeforeNanos)
				wgWorkers.Done()
			})
		}
		wgWorkers.Wait()
		group.duration.Record(mgr.nowFn().Sub(start))
	}
}

type flushTimesState int

const (
	flushTimesUninitialized flushTimesState = iota
	flushTimesUpdated
	flushTimesProcessed
)

type followerFlushMode int

const (
	unknownFollowerFlush followerFlushMode = iota
	kvUpdateFollowerFlush
	forcedFollowerFlush
)

type flusherWithTime struct {
	flusher          flushingMetricList
	flushBeforeNanos int64
}

type flushersGroup struct {
	interval time.Duration
	duration tally.Timer
	flushers []flusherWithTime
}
