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

package storage

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/m3db/m3db/client"
	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/context"
	"github.com/m3db/m3db/retention"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/storage/repair"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3x/errors"
	"github.com/m3db/m3x/log"
	"github.com/m3db/m3x/time"

	"github.com/m3db/m3db/topology"
	"github.com/uber-go/tally"
)

var (
	errNoRepairOptions  = errors.New("no repair options")
	errRepairInProgress = errors.New("repair already in progress")
)

type recordFn func(namespace ts.ID, shard databaseShard, diffRes repair.MetadataComparisonResult)
type repairShardFn func(namespace ts.ID, shard databaseShard, diffRes repair.MetadataComparisonResult) error

type shardRepairer struct {
	opts      Options
	rpopts    repair.Options
	rtopts    retention.Options
	client    client.AdminClient
	recordFn  recordFn
	repairFn  repairShardFn
	logger    xlog.Logger
	scope     tally.Scope
	nowFn     clock.NowFn
	blockSize time.Duration
}

func newShardRepairer(opts Options, rpopts repair.Options) (databaseShardRepairer, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	iopts := opts.InstrumentOptions()
	scope := iopts.MetricsScope().SubScope("database.repair").Tagged(map[string]string{"host": hostname})
	rtopts := opts.RetentionOptions()

	r := shardRepairer{
		opts:      opts,
		rpopts:    rpopts,
		rtopts:    rtopts,
		client:    rpopts.AdminClient(),
		logger:    iopts.Logger(),
		scope:     scope,
		nowFn:     opts.ClockOptions().NowFn(),
		blockSize: rtopts.BlockSize(),
	}
	r.recordFn = r.recordDifferences
	r.repairFn = r.repairDifferences

	return r, nil
}

func (r shardRepairer) Options() repair.Options {
	return r.rpopts
}

func (r shardRepairer) Repair(
	ctx context.Context,
	namespace ts.ID,
	tr xtime.Range,
	shard databaseShard,
) (repair.MetadataComparisonResult, error) {
	session, err := r.client.DefaultAdminSession()
	if err != nil {
		return repair.MetadataComparisonResult{}, err
	}

	var (
		start    = tr.Start
		end      = tr.End
		origin   = session.Origin()
		replicas = session.Replicas()
	)

	metadata := repair.NewReplicaMetadataComparer(replicas, r.rpopts)
	ctx.RegisterFinalizer(metadata)

	// Add local metadata
	localMetadata, _ := shard.FetchBlocksMetadata(ctx, start, end, math.MaxInt64, 0, true, true)
	ctx.RegisterFinalizer(context.FinalizerFn(localMetadata.Close))

	localIter := block.NewFilteredBlocksMetadataIter(localMetadata)
	metadata.AddLocalMetadata(origin, localIter)

	// Add peer metadata
	peerIter, err := session.FetchBlocksMetadataFromPeers(namespace, shard.ID(), start, end)
	if err != nil {
		return repair.MetadataComparisonResult{}, err
	}
	if err := metadata.AddPeerMetadata(peerIter); err != nil {
		return repair.MetadataComparisonResult{}, err
	}

	metadataRes := metadata.Compare()

	r.recordFn(namespace, shard, metadataRes)
	// TODO(prateek): flipr config for this(!)
	if err = r.repairFn(namespace, shard, metadataRes); err != nil {
		return repair.MetadataComparisonResult{}, err
	}

	// TODO(prateek):
	// - change the return type to include a RepairResult construct
	// - trace up the chain here, make sure we mark state to re-attempt any pending repairs
	return metadataRes, nil
}

type hostSet map[topology.Host]repair.HostBlockMetadata

type blockID struct {
	id    ts.ID
	start time.Time
}

func newHostSet() *hostSet {
	hs := make(hostSet, 3)
	return &hs
}

func (h *hostSet) insert(hbm repair.HostBlockMetadata) {
	(*h)[hbm.Host] = hbm
}

func (h *hostSet) insertOneFromSet(oh *hostSet) {
	for _, hbm := range *oh {
		h.insert(hbm)
		return
	}
}

func (h *hostSet) contains(host topology.Host) bool {
	_, ok := (*h)[host]
	return ok
}

func (h *hostSet) remove(host topology.Host) {
	delete(*h, host)
}

func (h *hostSet) empty() bool {
	return len(*h) == 0
}

// TODO(prateek): tests for constructRequiredRepairBlocks
func (r shardRepairer) constructRequiredRepairBlocks(
	namespace ts.ID,
	shard databaseShard,
	diffRes repair.MetadataComparisonResult,
	originHost topology.Host,
) ([]block.ReplicaMetadata, map[blockID]*hostSet) {
	// TODO(prateek): pooling object creation in this method
	replicaState := make(map[blockID]*hostSet, diffRes.NumBlocks)

	// add size differences
	for _, rsm := range diffRes.SizeDifferences.Series() {
		for start, blk := range rsm.Metadata.Blocks() {
			blkID := blockID{rsm.ID, start}
			// find all unique sizes seen for this block
			uniqueValues := make(map[int64]*hostSet, 3)
			for _, hBlk := range blk.Metadata() {
				sz := hBlk.Size
				if _, ok := uniqueValues[sz]; !ok {
					uniqueValues[sz] = newHostSet()
				}
				uniqueValues[sz].insert(hBlk)
			}
			for _, hs := range uniqueValues {
				if hs.contains(originHost) {
					// we already have originHost data available locally, no need to fetch it
					continue
				}
				// insert value into replicaState
				if _, ok := replicaState[blkID]; !ok {
					replicaState[blkID] = newHostSet()
				}
				// only care about a single value from the unique set
				replicaState[blkID].insertOneFromSet(hs)
			}
		}
	}

	// add checksum differences
	for _, rsm := range diffRes.ChecksumDifferences.Series() {
		for start, blk := range rsm.Metadata.Blocks() {
			blkID := blockID{rsm.ID, start}
			// find all unique checksums seen for this block
			uniqueValues := make(map[uint32]*hostSet, 3)
			for _, hBlk := range blk.Metadata() {
				cs := hBlk.Checksum
				if cs == nil {
					continue // checksum unavailable, don't include in diff
				}
				if _, ok := uniqueValues[*cs]; !ok {
					uniqueValues[*cs] = newHostSet()
				}
				uniqueValues[*cs].insert(hBlk)
			}
			for _, hs := range uniqueValues {
				if hs.contains(originHost) {
					// we already have originHost data available locally, no need to fetch it
					continue
				}
				// insert value into replicaState
				if _, ok := replicaState[blkID]; !ok {
					replicaState[blkID] = newHostSet()
				}
				// only care about a single value from the unique set
				replicaState[blkID].insertOneFromSet(hs)
			}
		}
	}

	repairBlocks := make([]block.ReplicaMetadata, 0, diffRes.NumBlocks)
	for blkID, hs := range replicaState {
		for _, h := range *hs {
			repairBlocks = append(repairBlocks, block.ReplicaMetadata{
				Start:    blkID.start,
				ID:       blkID.id,
				Peer:     h.Host,
				Checksum: h.Checksum,
				Size:     h.Size,
			})
		}
	}
	return repairBlocks, replicaState
}

func (r shardRepairer) repairDifferences(
	namespace ts.ID,
	shard databaseShard,
	diffRes repair.MetadataComparisonResult,
) error {
	var (
		logger       = r.opts.InstrumentOptions().Logger()
		session, err = r.client.DefaultAdminSession()
		multiErr     xerrors.MultiError
	)
	if err != nil {
		return err
	}

	reqBlocks, blockState := r.constructRequiredRepairBlocks(namespace, shard, diffRes, session.Origin())
	blocksIter, err := session.FetchRepairBlocksFromPeers(namespace, shard.ID(), reqBlocks, r.rpopts.ResultOptions())
	if err != nil {
		return err
	}

	for blocksIter.Next() {
		host, id, blk := blocksIter.Current()
		// TODO(prateek): does Close() need to be called on blk
		// i.e. figure out ownership - is it reset by iterator or not

		blkID := blockID{id, blk.StartTime()}
		if hs, ok := blockState[blkID]; !ok || !hs.contains(host) {
			// should never happen
			logger.WithFields(
				xlog.NewLogField("id", id.String()),
				xlog.NewLogField("host", host.String()),
				xlog.NewLogField("blockID", blkID),
			).Warnf("received un-requested block, session.FetchRepairBlockFromPeers violated contract, skipping.")
			continue
		}

		blockState[blkID].remove(host)
		// TODO(prateek): current implementation of UpdateSeries marks shards flushed internally,
		// this is something we should amortize to be done at the end of iterating through this function
		// concern being, we want to minimize the number of flushes we can potentially induce
		markFlushStateDirty := true
		if err := shard.UpdateSeries(id, blk, markFlushStateDirty); err != nil {
			multiErr.Add(fmt.Errorf(
				"unable to update series [ id = %v, block_start = %v, err = %v ]", id.String(), blkID.start, err))
			// TODO(prateek): increment error count in return object
			// TODO(prateek): publish metrics for this too
			continue
		}
		//	track number of "repaired" blocks, report metric
	}
	// for any cached, and not consolidated, track in errors metric
	// or, if blocksIter.Err() is :(

	// TODO(prateek): figure out how to transfer shardResult -> local shards
	// and then write those files on disk

	// for writes:
	// 	- don't always write, factor in a minimum number of blocks repaired,
	//    and number of blocks still requiring repair
	// 	- write a new version of the file for the timestamp, keep last 'n' versions,
	//    do NOT delete old version before writing a new version

	// TODO(prateek): change the return type to include a RepairResult construct
	return multiErr.FinalError()
}

func (r shardRepairer) recordDifferences(
	namespace ts.ID,
	shard databaseShard,
	diffRes repair.MetadataComparisonResult,
) {
	var (
		shardScope = r.scope.Tagged(map[string]string{
			"namespace": namespace.String(),
			"shard":     strconv.Itoa(int(shard.ID())),
		})
		totalScope        = shardScope.Tagged(map[string]string{"resultType": "total"})
		sizeDiffScope     = shardScope.Tagged(map[string]string{"resultType": "sizeDiff"})
		checksumDiffScope = shardScope.Tagged(map[string]string{"resultType": "checksumDiff"})
	)

	// Record total number of series and total number of blocks
	totalScope.Counter("series").Inc(diffRes.NumSeries)
	totalScope.Counter("blocks").Inc(diffRes.NumBlocks)

	// Record size differences
	sizeDiffScope.Counter("series").Inc(diffRes.SizeDifferences.NumSeries())
	sizeDiffScope.Counter("blocks").Inc(diffRes.SizeDifferences.NumBlocks())

	// Record checksum differences
	checksumDiffScope.Counter("series").Inc(diffRes.ChecksumDifferences.NumSeries())
	checksumDiffScope.Counter("blocks").Inc(diffRes.ChecksumDifferences.NumBlocks())
}

type repairFn func() error

type sleepFn func(d time.Duration)

type repairStatus int

const (
	repairNotStarted repairStatus = iota
	repairSuccess
	repairFailed
)

type repairState struct {
	Status      repairStatus
	NumFailures int
}

type dbRepairer struct {
	sync.Mutex

	database      database
	ropts         repair.Options
	rtopts        retention.Options
	shardRepairer databaseShardRepairer
	repairStates  map[time.Time]repairState

	repairFn            repairFn
	sleepFn             sleepFn
	nowFn               clock.NowFn
	logger              xlog.Logger
	repairInterval      time.Duration
	repairTimeOffset    time.Duration
	repairTimeJitter    time.Duration
	repairCheckInterval time.Duration
	repairMaxRetries    int
	closed              bool
	running             int32
}

func newDatabaseRepairer(database database) (databaseRepairer, error) {
	opts := database.Options()
	nowFn := opts.ClockOptions().NowFn()
	ropts := opts.RepairOptions()
	if ropts == nil {
		return nil, errNoRepairOptions
	}
	if err := ropts.Validate(); err != nil {
		return nil, err
	}

	shardRepairer, err := newShardRepairer(opts, ropts)
	if err != nil {
		return nil, err
	}

	var jitter time.Duration
	if repairJitter := ropts.RepairTimeJitter(); repairJitter > 0 {
		src := rand.NewSource(nowFn().UnixNano())
		jitter = time.Duration(float64(repairJitter) * (float64(src.Int63()) / float64(math.MaxInt64)))
	}

	r := &dbRepairer{
		database:            database,
		ropts:               ropts,
		rtopts:              opts.RetentionOptions(),
		shardRepairer:       shardRepairer,
		repairStates:        make(map[time.Time]repairState),
		sleepFn:             time.Sleep,
		nowFn:               nowFn,
		logger:              opts.InstrumentOptions().Logger(),
		repairInterval:      ropts.RepairInterval(),
		repairTimeOffset:    ropts.RepairTimeOffset(),
		repairTimeJitter:    jitter,
		repairCheckInterval: ropts.RepairCheckInterval(),
		repairMaxRetries:    ropts.RepairMaxRetries(),
	}
	r.repairFn = r.Repair

	return r, nil
}

func (r *dbRepairer) run() {
	var curIntervalStart time.Time

	for {
		r.Lock()
		closed := r.closed
		r.Unlock()

		if closed {
			break
		}

		r.sleepFn(r.repairCheckInterval)

		now := r.nowFn()
		intervalStart := now.Truncate(r.repairInterval)

		// If we haven't reached the offset yet, skip
		target := intervalStart.Add(r.repairTimeOffset + r.repairTimeJitter)
		if now.Before(target) {
			continue
		}

		// If we are in the same interval, we must have already repaired, skip
		if intervalStart == curIntervalStart {
			continue
		}

		curIntervalStart = intervalStart
		if err := r.repairFn(); err != nil {
			r.logger.Errorf("error repairing database: %v", err)
		}
	}
}

func (r *dbRepairer) repairTimeRanges() xtime.Ranges {
	var (
		now       = r.nowFn()
		blockSize = r.rtopts.BlockSize()
		start     = now.Add(-r.rtopts.RetentionPeriod()).Truncate(blockSize)
		end       = now.Add(-r.rtopts.BufferPast()).Truncate(blockSize)
	)

	targetRanges := xtime.NewRanges().AddRange(xtime.Range{Start: start, End: end})
	for t := range r.repairStates {
		if !r.needsRepair(t) {
			targetRanges = targetRanges.RemoveRange(xtime.Range{Start: t, End: t.Add(blockSize)})
		}
	}

	return targetRanges
}

func (r *dbRepairer) needsRepair(t time.Time) bool {
	repairState, exists := r.repairStates[t]
	if !exists {
		return true
	}
	return repairState.Status == repairFailed && repairState.NumFailures < r.repairMaxRetries
}

func (r *dbRepairer) Start() {
	if r.repairInterval <= 0 {
		return
	}

	go r.run()
}

func (r *dbRepairer) Stop() {
	r.Lock()
	r.closed = true
	r.Unlock()
}

func (r *dbRepairer) Repair() error {
	// Don't attempt a repair if the database is not bootstrapped yet
	if !r.database.IsBootstrapped() {
		return nil
	}

	if !atomic.CompareAndSwapInt32(&r.running, 0, 1) {
		return errRepairInProgress
	}

	defer func() {
		atomic.StoreInt32(&r.running, 0)
	}()

	multiErr := xerrors.NewMultiError()
	blockSize := r.rtopts.BlockSize()
	iter := r.repairTimeRanges().Iter()
	for iter.Next() {
		tr := iter.Value()
		err := r.repairWithTimeRange(tr)
		for t := tr.Start; t.Before(tr.End); t = t.Add(blockSize) {
			repairState := r.repairStates[t]
			if err == nil {
				repairState.Status = repairSuccess
			} else {
				repairState.Status = repairFailed
				repairState.NumFailures++
			}
			r.repairStates[t] = repairState
		}
		multiErr = multiErr.Add(err)
	}

	return multiErr.FinalError()
}

func (r *dbRepairer) repairWithTimeRange(tr xtime.Range) error {
	multiErr := xerrors.NewMultiError()
	namespaces := r.database.getOwnedNamespaces()
	for _, n := range namespaces {
		if err := n.Repair(r.shardRepairer, tr); err != nil {
			detailedErr := fmt.Errorf("namespace %s failed to repair time range %v: %v", n.ID().String(), tr, err)
			multiErr = multiErr.Add(detailedErr)
		}
	}
	return multiErr.FinalError()
}

func (r *dbRepairer) IsRepairing() bool {
	return atomic.LoadInt32(&r.running) == 1
}
