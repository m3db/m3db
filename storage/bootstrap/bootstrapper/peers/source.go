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

package peers

import (
	"fmt"
	"sync"
	"time"

	"github.com/m3db/m3db/client"
	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/persist"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/storage/bootstrap"
	"github.com/m3db/m3db/storage/bootstrap/result"
	"github.com/m3db/m3db/storage/namespace"
	"github.com/m3db/m3db/storage/series"
	"github.com/m3db/m3x/context"
	xlog "github.com/m3db/m3x/log"
	xsync "github.com/m3db/m3x/sync"
	xtime "github.com/m3db/m3x/time"
)

type peersSource struct {
	opts  Options
	log   xlog.Logger
	nowFn clock.NowFn
}

type incrementalFlush struct {
	nsMetadata        namespace.Metadata
	shard             uint32
	shardRetrieverMgr block.DatabaseShardBlockRetrieverManager
	shardResult       result.ShardResult
	timeRange         xtime.Range
}

func newPeersSource(opts Options) bootstrap.Source {
	return &peersSource{
		opts:  opts,
		log:   opts.ResultOptions().InstrumentOptions().Logger(),
		nowFn: opts.ResultOptions().ClockOptions().NowFn(),
	}
}

func (s *peersSource) Can(strategy bootstrap.Strategy) bool {
	switch strategy {
	case bootstrap.BootstrapSequential:
		return true
	}
	return false
}

func (s *peersSource) Available(
	nsMetadata namespace.Metadata,
	shardsTimeRanges result.ShardTimeRanges,
) result.ShardTimeRanges {
	// Peers should be able to fulfill all data
	return shardsTimeRanges
}

func (s *peersSource) Read(
	nsMetadata namespace.Metadata,
	shardsTimeRanges result.ShardTimeRanges,
	opts bootstrap.RunOptions,
) (result.BootstrapResult, error) {
	if shardsTimeRanges.IsEmpty() {
		return nil, nil
	}

	var (
		namespace         = nsMetadata.ID()
		blockRetriever    block.DatabaseBlockRetriever
		shardRetrieverMgr block.DatabaseShardBlockRetrieverManager
		persistFlush      persist.Flush
		incremental       = false
		seriesCachePolicy = s.opts.ResultOptions().SeriesCachePolicy()
	)
	if opts.Incremental() && seriesCachePolicy != series.CacheAll {
		retrieverMgr := s.opts.DatabaseBlockRetrieverManager()
		persistManager := s.opts.PersistManager()

		// Neither of these should ever happen
		if seriesCachePolicy != series.CacheAll && retrieverMgr == nil {
			s.log.Fatal("tried to perform incremental flush without retriever manager")
		}
		if seriesCachePolicy != series.CacheAll && persistManager == nil {
			s.log.Fatal("tried to perform incremental flush without persist manager")
		}

		s.log.WithFields(
			xlog.NewField("namespace", namespace.String()),
		).Infof("peers bootstrapper resolving block retriever")

		r, err := retrieverMgr.Retriever(nsMetadata)
		if err != nil {
			return nil, err
		}

		persist, err := persistManager.StartPersist()
		if err != nil {
			return nil, err
		}

		defer persist.Done()

		incremental = true
		blockRetriever = r
		shardRetrieverMgr = block.NewDatabaseShardBlockRetrieverManager(r)
		persistFlush = persist
	}

	result := result.NewBootstrapResult()
	session, err := s.opts.AdminClient().DefaultAdminSession()
	if err != nil {
		s.log.Errorf("peers bootstrapper cannot get default admin session: %v", err)
		result.SetUnfulfilled(shardsTimeRanges)
		return nil, err
	}

	var (
		resultLock              sync.Mutex
		wg                      sync.WaitGroup
		incrementalWorkerDoneCh = make(chan struct{})
		incrementalMaxQueue     = s.opts.IncrementalPersistMaxQueueSize()
		incrementalQueue        = make(chan incrementalFlush, incrementalMaxQueue)
		bopts                   = s.opts.ResultOptions()
		count                   = len(shardsTimeRanges)
		concurrency             = s.opts.DefaultShardConcurrency()
		blockSize               = nsMetadata.Options().RetentionOptions().BlockSize()
	)
	if incremental {
		concurrency = s.opts.IncrementalShardConcurrency()
	}

	s.log.WithFields(
		xlog.NewField("shards", count),
		xlog.NewField("concurrency", concurrency),
		xlog.NewField("incremental", incremental),
	).Infof("peers bootstrapper bootstrapping shards for ranges")
	if incremental {
		go s.startIncrementalQueueWorkerLoop(
			incrementalWorkerDoneCh, incrementalQueue, persistFlush, result, &resultLock)
	}

	workers := xsync.NewWorkerPool(concurrency)
	workers.Init()
	for shard, ranges := range shardsTimeRanges {
		shard, ranges := shard, ranges
		wg.Add(1)
		workers.Go(func() {
			defer wg.Done()
			s.fetchBootstrapBlocksFromPeers(shard, ranges, nsMetadata,
				session, bopts, result, &resultLock, incremental, incrementalQueue, shardRetrieverMgr, blockSize)
		})
	}

	wg.Wait()
	close(incrementalQueue)
	if incremental {
		// Wait for the incrementalQueueWorker to finish incrementally flushing everything
		<-incrementalWorkerDoneCh
	}

	if incremental {
		// Now cache the incremental results
		err := s.cacheShardIndices(shardsTimeRanges, blockRetriever)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// startIncrementalQueueWorkerLoop is meant to be run in its own goroutine, and it creates a worker that
// loops through the incrementalQueue and performs an incrementalFlush for each entry, ensuring that
// no more than one incremental flush is ever happening at once. Once the incrementalQueue channel
// is closed, and the worker has completed flushing all the remaining entries, it will close the
// provided doneCh so that callers can block until everything has been successfully flushed.
func (s *peersSource) startIncrementalQueueWorkerLoop(
	doneCh chan struct{},
	incrementalQueue chan incrementalFlush,
	persistFlush persist.Flush,
	bootstrapResult result.BootstrapResult,
	lock *sync.Mutex,
) {
	// If performing an incremental bootstrap then flush one
	// at a time as shard results are gathered
	for flush := range incrementalQueue {
		err := s.incrementalFlush(persistFlush, flush.nsMetadata, flush.shard,
			flush.shardRetrieverMgr, flush.shardResult, flush.timeRange)
		if err == nil {
			// Safe to add to the shared bootstrap result now
			lock.Lock()
			bootstrapResult.Add(flush.shard, flush.shardResult, xtime.Ranges{})
			lock.Unlock()
			continue
		}

		// Remove results and make unfulfilled if an error occurred
		s.log.WithFields(
			xlog.NewField("error", err.Error()),
		).Errorf("peers bootstrapper incremental flush encountered error")

		// Make unfulfilled
		lock.Lock()
		bootstrapResult.Add(flush.shard, nil, xtime.Ranges{}.AddRange(flush.timeRange))
		lock.Unlock()
	}
	close(doneCh)
}

// fetchBootstrapBlocksFromPeers loops through all the provided ranges for a given shard and
// fetches all the bootstrap blocks from the appropriate peers.
// 		Non-incremental case: Immediately add the results to the bootstrap result
// 		Incremental case: Don't add the results yet, but push an incrementalFlush into the
// 						  incremental queue. The incrementalQueue worker will eventually
// 						  add the results once its performed the incremental flush.
func (s *peersSource) fetchBootstrapBlocksFromPeers(
	shard uint32,
	ranges xtime.Ranges,
	nsMetadata namespace.Metadata,
	session client.AdminSession,
	bopts result.Options,
	bootstrapResult result.BootstrapResult,
	lock *sync.Mutex,
	incremental bool,
	incrementalQueue chan incrementalFlush,
	shardRetrieverMgr block.DatabaseShardBlockRetrieverManager,
	blockSize time.Duration,
) {
	it := ranges.Iter()
	for it.Next() {
		currRange := it.Value()

		for blockStart := currRange.Start; blockStart.Before(currRange.End); blockStart = blockStart.Add(blockSize) {
			version := s.opts.FetchBlocksMetadataEndpointVersion()
			blockEnd := blockStart.Add(blockSize)
			shardResult, err := session.FetchBootstrapBlocksFromPeers(
				nsMetadata, shard, blockStart, blockEnd, bopts, version)

			s.logFetchBootstrapBlocksFromPeersOutcome(shard, shardResult, err)

			if err != nil {
				// Do not add result at all to the bootstrap result
				lock.Lock()
				bootstrapResult.Add(shard, nil, xtime.Ranges{}.AddRange(currRange))
				lock.Unlock()
				continue
			}

			if incremental {
				incrementalQueue <- incrementalFlush{
					nsMetadata:        nsMetadata,
					shard:             shard,
					shardRetrieverMgr: shardRetrieverMgr,
					shardResult:       shardResult,
					timeRange:         xtime.Range{Start: blockStart, End: blockEnd},
				}
				continue
			}

			// If not waiting to incremental flush, add straight away to bootstrap result
			lock.Lock()
			bootstrapResult.Add(shard, shardResult, xtime.Ranges{})
			lock.Unlock()
		}
	}
}

func (s *peersSource) logFetchBootstrapBlocksFromPeersOutcome(
	shard uint32,
	shardResult result.ShardResult,
	err error,
) {
	if err == nil {
		shardBlockSeriesCounter := map[xtime.UnixNano]int64{}
		for _, entry := range shardResult.AllSeries().Iter() {
			series := entry.Value()
			for blockStart := range series.Blocks.AllBlocks() {
				shardBlockSeriesCounter[blockStart]++
			}
		}

		for block, numSeries := range shardBlockSeriesCounter {
			s.log.WithFields(
				xlog.NewField("shard", shard),
				xlog.NewField("numSeries", numSeries),
				xlog.NewField("block", block),
			).Info("peer bootstrapped shard")
		}
	} else {
		s.log.WithFields(
			xlog.NewField("shard", shard),
			xlog.NewField("error", err.Error()),
		).Error("error fetching bootstrap blocks from peers")
	}
}

// incrementalFlush is used to incrementally flush peer-bootstrapped shards
// to disk as they finish so that we're not (necessarily) holding everything
// in memory at once.
// incrementalFlush starts by looping through every block in a timerange for
// a given shard, and then subsequently looping through every series in that
// shard/block and flushing it to disk. Depending on the series caching policy,
// the series will either be held in memory, or removed from memory once
// flushing has completed.
// Once everything has been flushed to disk then depending on the series
// caching policy the function is either done, or in the case of the
// CacheAllMetadata policy we loop through every series and make every block
// retrievable (so that we can retrieve data for the blocks that we're caching
// the metadata for).
// In addition, if the caching policy is not CacheAll or CacheAllMetadata, then
// at the end we remove all the series objects from the shard result as well
// (since all their corresponding blocks have been removed anyways) to prevent
// a huge memory spike caused by adding lots of unused series to the Shard
// object and then immediately evicting them in the next tick.
func (s *peersSource) incrementalFlush(
	flush persist.Flush,
	nsMetadata namespace.Metadata,
	shard uint32,
	shardRetrieverMgr block.DatabaseShardBlockRetrieverManager,
	shardResult result.ShardResult,
	tr xtime.Range,
) error {
	var (
		ropts             = nsMetadata.Options().RetentionOptions()
		blockSize         = ropts.BlockSize()
		shardRetriever    = shardRetrieverMgr.ShardRetriever(shard)
		tmpCtx            = context.NewContext()
		seriesCachePolicy = s.opts.ResultOptions().SeriesCachePolicy()
	)
	if seriesCachePolicy == series.CacheAllMetadata && shardRetriever == nil {
		return fmt.Errorf("shard retriever missing for shard: %d", shard)
	}

	for start := tr.Start; start.Before(tr.End); start = start.Add(blockSize) {
		prepareOpts := persist.PrepareOptions{
			NamespaceMetadata: nsMetadata,
			Shard:             shard,
			BlockStart:        start,
		}
		prepared, err := flush.Prepare(prepareOpts)
		if err != nil {
			return err
		}

		var blockErr error
		for _, entry := range shardResult.AllSeries().Iter() {
			s := entry.Value()
			bl, ok := s.Blocks.BlockAt(start)
			if !ok {
				continue
			}

			tmpCtx.Reset()
			stream, err := bl.Stream(tmpCtx)
			if err != nil {
				tmpCtx.BlockingClose()
				blockErr = err // Need to call prepared.Close, avoid return
				break
			}

			segment, err := stream.Segment()
			if err != nil {
				tmpCtx.BlockingClose()
				blockErr = err // Need to call prepared.Close, avoid return
				break
			}

			err = prepared.Persist(s.ID, segment, bl.Checksum())
			tmpCtx.BlockingClose()
			if err != nil {
				blockErr = err // Need to call prepared.Close, avoid return
				break
			}

			switch seriesCachePolicy {
			case series.CacheAll:
				// Leave the blocks in the shard result, we need to return all blocks
				// so we can cache in memory
			case series.CacheAllMetadata:
				// NB(r): We can now make the flushed blocks retrievable, note that we
				// explicitly perform another loop here and lookup the block again
				// to avoid a large expensive allocation to hold onto the blocks
				// that we just flushed that would have to be pooled.
				// We are explicitly trading CPU time here for lower GC pressure.
				metadata := block.RetrievableBlockMetadata{
					ID:       s.ID,
					Length:   bl.Len(),
					Checksum: bl.Checksum(),
				}
				bl.ResetRetrievable(start, blockSize, shardRetriever, metadata)
			default:
				// Not caching the series or metadata in memory so finalize the block,
				// better to do this as we loop through to make blocks return to the
				// pool earlier than at the end of this flush cycle
				s.Blocks.RemoveBlockAt(start)
				bl.Close()
			}
		}

		// Always close before attempting to check if block error occurred,
		// avoid using a defer here as this needs to be done for each inner loop
		err = prepared.Close()
		if blockErr != nil {
			// A block error is more interesting to bubble up than a close error
			err = blockErr
		}

		if err != nil {
			return err
		}
	}

	if seriesCachePolicy != series.CacheAll && seriesCachePolicy != series.CacheAllMetadata {
		// TODO: We need this right now because nodes with older versions of M3DB will return an extra
		// block when requesting bootstrapped blocks. Once all the clusters have been upgraded we can
		// remove this code.
		for _, entry := range shardResult.AllSeries().Iter() {
			s := entry.Value()
			bl, ok := s.Blocks.BlockAt(tr.End)
			if !ok {
				continue
			}
			s.Blocks.RemoveBlockAt(tr.End)
			bl.Close()
		}

		// If we're not going to keep all of the data, or at least all of the metadata in memory
		// then we don't want to keep these series in the shard result. If we leave them in, then
		// they will all get loaded into the shard object, and then immediately evicted on the next
		// tick which causes unnecessary memory pressure.
		numSeriesTriedToRemoveWithRemainingBlocks := 0
		for _, entry := range shardResult.AllSeries().Iter() {
			series := entry.Value()
			numBlocksRemaining := len(series.Blocks.AllBlocks())
			// Should never happen since we removed all the block in the previous loop and fetching
			// bootstrap blocks should always be exclusive on the end side.
			if numBlocksRemaining > 0 {
				numSeriesTriedToRemoveWithRemainingBlocks++
				continue
			}

			shardResult.RemoveSeries(series.ID)
			series.Blocks.Close()
			// Safe to finalize these IDs because the prepared object was the only other thing
			// using them, and it has been closed.
			series.ID.Finalize()
		}
		if numSeriesTriedToRemoveWithRemainingBlocks > 0 {
			s.log.WithFields(
				xlog.NewField("start", tr.Start.Unix()),
				xlog.NewField("end", tr.End.Unix()),
				xlog.NewField("numTimes", numSeriesTriedToRemoveWithRemainingBlocks),
			).Error("error tried to remove series that still has blocks")
		}
	}

	return nil
}

func (s *peersSource) cacheShardIndices(
	shardsTimeRanges result.ShardTimeRanges,
	blockRetriever block.DatabaseBlockRetriever,
) error {
	// Now cache the incremental results
	shards := make([]uint32, 0, len(shardsTimeRanges))
	for shard := range shardsTimeRanges {
		shards = append(shards, shard)
	}

	return blockRetriever.CacheShardIndices(shards)
}
