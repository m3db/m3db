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
	"sync"

	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/context"
	"github.com/m3db/m3db/persist"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/storage/bootstrap"
	"github.com/m3db/m3db/storage/bootstrap/result"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3x/log"
	"github.com/m3db/m3x/sync"
	"github.com/m3db/m3x/time"
)

type peersSource struct {
	opts  Options
	log   xlog.Logger
	nowFn clock.NowFn
}

type incrementalFlush struct {
	namespace         ts.ID
	persistManager    persist.Manager
	shard             uint32
	shardRetrieverMgr block.DatabaseShardBlockRetrieverManager
	shardResult       result.ShardResult
	timeRange         xtime.Range
}

type incrementalFlushedBlock struct {
	id    ts.ID
	block block.DatabaseBlock
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
	namespace ts.ID,
	shardsTimeRanges result.ShardTimeRanges,
) result.ShardTimeRanges {
	// Peers should be able to fulfill all data
	return shardsTimeRanges
}

func (s *peersSource) Read(
	namespace ts.ID,
	shardsTimeRanges result.ShardTimeRanges,
	opts bootstrap.RunOptions,
) (result.BootstrapResult, error) {
	if shardsTimeRanges.IsEmpty() {
		return nil, nil
	}

	var (
		blockRetriever    block.DatabaseBlockRetriever
		shardRetrieverMgr block.DatabaseShardBlockRetrieverManager
		persistManager    persist.Manager
		incremental       = false
	)
	if opts.Incremental() {
		retrieverMgr := s.opts.DatabaseBlockRetrieverManager()
		newPersistMgrFn := s.opts.NewPersistManagerFn()
		if retrieverMgr != nil && newPersistMgrFn != nil {
			s.log.WithFields(
				xlog.NewLogField("namespace", namespace.String()),
			).Infof("peers bootstrapper resolving block retriever")

			r, err := retrieverMgr.Retriever(namespace)
			if err != nil {
				return nil, err
			}

			incremental = true
			blockRetriever = r
			shardRetrieverMgr = block.NewDatabaseShardBlockRetrieverManager(r)
			persistManager = newPersistMgrFn()
		} else {
			s.log.WithFields(
				xlog.NewLogField("namespace", namespace.String()),
				xlog.NewLogField("noRetrieverMgr", retrieverMgr == nil),
				xlog.NewLogField("noNewPersistMgrFn", newPersistMgrFn == nil),
			).Infof("peers bootstrapper skipping incremental run")
		}
	}

	result := result.NewBootstrapResult()
	session, err := s.opts.AdminClient().DefaultAdminSession()
	if err != nil {
		s.log.Errorf("peers bootstrapper cannot get default admin session: %v", err)
		result.SetUnfulfilled(shardsTimeRanges)
		return result, nil
	}

	var (
		lock                sync.Mutex
		wg                  sync.WaitGroup
		incrementalWg       sync.WaitGroup
		incrementalMaxQueue = s.opts.IncrementalBootstrapPersistMaxQueue()
		incrementalQueue    = make(chan incrementalFlush, incrementalMaxQueue)
		bopts               = s.opts.ResultOptions()
		count               = len(shardsTimeRanges)
		concurrency         = s.opts.DefaultBootstrapShardConcurrency()
	)
	if incremental {
		concurrency = s.opts.IncrementalBootstrapShardConcurrency()
	}

	s.log.WithFields(
		xlog.NewLogField("shards", count),
		xlog.NewLogField("concurrency", concurrency),
		xlog.NewLogField("incremental", incremental),
	).Infof("peers bootstrapper bootstrapping shards for ranges")
	if incremental {
		// If performing an incremental bootstrap then flush one
		// at a time as shard results are gathered
		incrementalWg.Add(1)
		go func() {
			defer incrementalWg.Done()

			for flush := range incrementalQueue {
				err := s.incrementalFlush(flush.namespace, flush.persistManager,
					flush.shard, flush.shardRetrieverMgr, flush.shardResult, flush.timeRange)
				if err != nil {
					s.log.WithFields(
						xlog.NewLogField("error", err.Error()),
					).Infof("peers bootstrapper incremental flush encountered error")
				}
			}
		}()
	}

	workers := xsync.NewWorkerPool(concurrency)
	workers.Init()
	for shard, ranges := range shardsTimeRanges {
		shard, ranges := shard, ranges
		wg.Add(1)
		workers.Go(func() {
			defer wg.Done()

			it := ranges.Iter()
			for it.Next() {
				currRange := it.Value()

				shardResult, err := session.FetchBootstrapBlocksFromPeers(namespace,
					shard, currRange.Start, currRange.End, bopts)

				if err == nil && incremental {
					incrementalQueue <- incrementalFlush{
						namespace:         namespace,
						persistManager:    persistManager,
						shard:             shard,
						shardRetrieverMgr: shardRetrieverMgr,
						shardResult:       shardResult,
						timeRange:         currRange,
					}
				}

				lock.Lock()
				if err == nil {
					result.Add(shard, shardResult, nil)
				} else {
					result.Add(shard, nil, xtime.NewRanges().AddRange(currRange))
				}
				lock.Unlock()
			}
		})
	}

	wg.Wait()

	close(incrementalQueue)
	incrementalWg.Wait()

	if incremental {
		// Now cache the incremental results
		shards := make([]uint32, 0, len(shardsTimeRanges))
		for shard := range shardsTimeRanges {
			shards = append(shards, shard)
		}

		if err = blockRetriever.CacheShardIndices(shards); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (s *peersSource) incrementalFlush(
	namespace ts.ID,
	persistManager persist.Manager,
	shard uint32,
	shardRetrieverMgr block.DatabaseShardBlockRetrieverManager,
	shardResult result.ShardResult,
	tr xtime.Range,
) error {
	var (
		numSeries      = len(shardResult.AllSeries())
		flushedBlocks  = make([]incrementalFlushedBlock, 0, numSeries)
		ropts          = s.opts.ResultOptions().RetentionOptions()
		blockSize      = ropts.BlockSize()
		shardRetriever = shardRetrieverMgr.ShardRetriever(shard)
		tmpCtx         = context.NewContext()
	)
	for start := tr.Start; start.Before(tr.End); start = start.Add(blockSize) {
		prepared, err := persistManager.Prepare(namespace, shard, start)
		if err != nil {
			return err
		}

		flushedBlocks = flushedBlocks[:0]

		for _, series := range shardResult.AllSeries() {
			bl, ok := series.Blocks.BlockAt(start)
			if !ok {
				continue
			}

			tmpCtx.Reset()
			stream, err := bl.Stream(tmpCtx)
			if err != nil {
				return err
			}

			segment, err := stream.Segment()
			if err != nil {
				return err
			}

			err = prepared.Persist(series.ID, segment, bl.Checksum())
			tmpCtx.BlockingClose()
			if err != nil {
				return err
			}

			entry := incrementalFlushedBlock{
				id:    series.ID,
				block: bl,
			}
			flushedBlocks = append(flushedBlocks, entry)
		}

		if err := prepared.Close(); err != nil {
			return err
		}

		// We can now make flushed blocks retrievable
		for _, entry := range flushedBlocks {
			metadata := block.RetrievableBlockMetadata{
				ID:       entry.id,
				Length:   entry.block.Len(),
				Checksum: entry.block.Checksum(),
			}
			entry.block.ResetRetrievable(start, shardRetriever, metadata)
		}
	}

	return nil
}
