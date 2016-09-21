package storage

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"

	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/context"
	"github.com/m3db/m3db/instrument"
	"github.com/m3db/m3db/persist"
	"github.com/m3db/m3db/persist/fs/commitlog"
	"github.com/m3db/m3db/pool"
	"github.com/m3db/m3db/sharding"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/storage/bootstrap"
	"github.com/m3db/m3db/storage/namespace"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3db/x/io"
	"github.com/m3db/m3x/errors"
	"github.com/m3db/m3x/log"
	"github.com/m3db/m3x/time"

	"github.com/uber-go/tally"
)

func commitLogWriteNoOp(
	series commitlog.Series,
	datapoint ts.Datapoint,
	unit xtime.Unit,
	annotation ts.Annotation,
) error {
	return nil
}

type dbNamespace struct {
	sync.RWMutex

	name     string
	shardSet sharding.ShardSet
	nopts    namespace.Options
	sopts    Options
	nowFn    clock.NowFn
	log      xlog.Logger
	bs       bootstrapState

	// Contains an entry to all shards for fast shard lookup, an
	// entry will be nil when this shard does not belong to current database
	shards []databaseShard

	increasingIndex  increasingIndex
	writeCommitLogFn writeCommitLogFn

	metrics databaseNamespaceMetrics
}

type databaseNamespaceMetrics struct {
	bootstrap   instrument.MethodMetrics
	flush       instrument.MethodMetrics
	unfulfilled tally.Counter
}

func newDatabaseNamespaceMetrics(scope tally.Scope, samplingRate float64) databaseNamespaceMetrics {
	return databaseNamespaceMetrics{
		bootstrap:   instrument.NewMethodMetrics(scope, "bootstrap", samplingRate),
		flush:       instrument.NewMethodMetrics(scope, "flush", samplingRate),
		unfulfilled: scope.Counter("bootstrap.unfulfilled"),
	}
}

func newDatabaseNamespace(
	metadata namespace.Metadata,
	shardSet sharding.ShardSet,
	increasingIndex increasingIndex,
	writeCommitLogFn writeCommitLogFn,
	sopts Options,
) databaseNamespace {
	name := metadata.Name()
	nopts := metadata.Options()
	fn := writeCommitLogFn
	if !nopts.WritesToCommitLog() {
		fn = commitLogWriteNoOp
	}

	iops := sopts.InstrumentOptions()
	scope := iops.MetricsScope().SubScope("database").Tagged(map[string]string{"namespace": name})

	n := &dbNamespace{
		name:             name,
		shardSet:         shardSet,
		nopts:            nopts,
		sopts:            sopts,
		nowFn:            sopts.ClockOptions().NowFn(),
		log:              sopts.InstrumentOptions().Logger(),
		increasingIndex:  increasingIndex,
		writeCommitLogFn: fn,
		metrics:          newDatabaseNamespaceMetrics(scope, iops.MetricsSamplingRate()),
	}

	n.initShards()

	return n
}

func (n *dbNamespace) Name() string {
	return n.name
}

func (n *dbNamespace) Tick() {
	shards := n.getOwnedShards()
	if len(shards) == 0 {
		return
	}

	// Tick through the shards sequentially to avoid parallel data flushes
	for _, shard := range shards {
		shard.Tick()
	}
}

func (n *dbNamespace) Write(
	ctx context.Context,
	id string,
	timestamp time.Time,
	value float64,
	unit xtime.Unit,
	annotation []byte,
) error {
	shardID := n.shardSet.Shard(id)
	shard, err := n.shardAt(shardID)
	if err != nil {
		return err
	}
	return shard.Write(ctx, id, timestamp, value, unit, annotation)
}

func (n *dbNamespace) ReadEncoded(
	ctx context.Context,
	id string,
	start, end time.Time,
) ([][]xio.SegmentReader, error) {
	shardID := n.shardSet.Shard(id)
	shard, err := n.shardAt(shardID)
	if err != nil {
		return nil, err
	}
	return shard.ReadEncoded(ctx, id, start, end)
}

func (n *dbNamespace) FetchBlocks(
	ctx context.Context,
	shardID uint32,
	id string,
	starts []time.Time,
) ([]block.FetchBlockResult, error) {
	shard, err := n.shardAt(shardID)
	if err != nil {
		return nil, xerrors.NewInvalidParamsError(err)
	}
	return shard.FetchBlocks(ctx, id, starts), nil
}

func (n *dbNamespace) FetchBlocksMetadata(
	ctx context.Context,
	shardID uint32,
	limit int64,
	pageToken int64,
	includeSizes bool,
	includeChecksums bool,
) ([]block.FetchBlocksMetadataResult, *int64, error) {
	shard, err := n.shardAt(shardID)
	if err != nil {
		return nil, nil, xerrors.NewInvalidParamsError(err)
	}
	res, nextPageToken := shard.FetchBlocksMetadata(ctx, limit, pageToken, includeSizes, includeChecksums)
	return res, nextPageToken, nil
}

func (n *dbNamespace) Bootstrap(
	bs bootstrap.Bootstrap,
	targetRanges xtime.Ranges,
	writeStart time.Time,
	cutover time.Time,
) error {
	callStart := n.nowFn()

	n.Lock()
	if n.bs == bootstrapped {
		n.Unlock()
		n.metrics.bootstrap.ReportSuccess(n.nowFn().Sub(callStart))
		return nil
	}
	if n.bs == bootstrapping {
		n.Unlock()
		n.metrics.bootstrap.ReportError(n.nowFn().Sub(callStart))
		return errNamespaceIsBootstrapping
	}
	n.bs = bootstrapping
	n.Unlock()

	defer func() {
		n.Lock()
		n.bs = bootstrapped
		n.Unlock()
	}()

	if !n.nopts.NeedsBootstrap() {
		n.metrics.bootstrap.ReportSuccess(n.nowFn().Sub(callStart))
		return nil
	}

	shards := n.getOwnedShards()
	shardIDs := make([]uint32, len(shards))
	for i, shard := range shards {
		shardIDs[i] = shard.ID()
	}

	result, err := bs.Run(targetRanges, n.name, shardIDs)
	if err != nil {
		n.log.Errorf("bootstrap for namespace %s aborted due to error: %v", n.name, err)
		n.metrics.bootstrap.ReportError(n.nowFn().Sub(callStart))
		return err
	}
	n.metrics.bootstrap.Success.Inc(1)

	// Bootstrap shards using at least half the CPUs available
	workers := pool.NewWorkerPool(int(math.Ceil(float64(runtime.NumCPU()) / 2)))
	workers.Init()

	var (
		multiErr = xerrors.NewMultiError()
		results  = result.ShardResults()
		mutex    sync.Mutex
		wg       sync.WaitGroup
	)
	for _, shard := range shards {
		shard := shard
		wg.Add(1)
		workers.Go(func() {
			var bootstrapped map[string]block.DatabaseSeriesBlocks
			if result, ok := results[shard.ID()]; ok {
				bootstrapped = result.AllSeries()
			}

			err := shard.Bootstrap(bootstrapped, writeStart, cutover)

			mutex.Lock()
			multiErr = multiErr.Add(err)
			mutex.Unlock()

			wg.Done()
		})
	}

	wg.Wait()

	// Counter, tag this with namespace
	unfulfilled := int64(len(result.Unfulfilled()))
	n.metrics.unfulfilled.Inc(unfulfilled)
	if unfulfilled > 0 {
		str := result.Unfulfilled().SummaryString()
		n.log.Errorf("bootstrap for namespace %s completed with unfulfilled ranges: %s", n.name, str)
	}

	err = multiErr.FinalError()
	n.metrics.bootstrap.ReportSuccessOrError(err, n.nowFn().Sub(callStart))
	return err
}

func (n *dbNamespace) Flush(
	ctx context.Context,
	blockStart time.Time,
	pm persist.Manager,
) error {
	callStart := n.nowFn()

	n.RLock()
	if n.bs != bootstrapped {
		n.RUnlock()
		n.metrics.flush.ReportError(n.nowFn().Sub(callStart))
		return errNamespaceNotBootstrapped
	}
	n.RUnlock()

	if !n.nopts.NeedsFlush() {
		n.metrics.flush.ReportSuccess(n.nowFn().Sub(callStart))
		return nil
	}

	multiErr := xerrors.NewMultiError()
	shards := n.getOwnedShards()
	for _, shard := range shards {
		// NB(xichen): we still want to proceed if a shard fails to flush its data.
		// Probably want to emit a counter here, but for now just log it.
		if err := shard.Flush(ctx, n.name, blockStart, pm); err != nil {
			detailedErr := fmt.Errorf("shard %d failed to flush data: %v", shard.ID(), err)
			multiErr = multiErr.Add(detailedErr)
		}
	}

	// if nil, succeeded

	if res := multiErr.FinalError(); res != nil {
		n.metrics.flush.ReportError(n.nowFn().Sub(callStart))
		return res
	}
	n.metrics.flush.ReportSuccess(n.nowFn().Sub(callStart))
	return nil
}

func (n *dbNamespace) CleanupFileset(earliestToRetain time.Time) error {
	if !n.nopts.NeedsFilesetCleanup() {
		return nil
	}

	multiErr := xerrors.NewMultiError()
	shards := n.getOwnedShards()
	for _, shard := range shards {
		if err := shard.CleanupFileset(n.name, earliestToRetain); err != nil {
			multiErr = multiErr.Add(err)
		}
	}

	return multiErr.FinalError()
}

func (n *dbNamespace) Truncate() (int64, error) {
	var totalNumSeries int64

	n.RLock()
	shards := n.shardSet.Shards()
	for _, shard := range shards {
		totalNumSeries += n.shards[shard].NumSeries()
	}
	n.RUnlock()

	// For now we are simply dropping all the objects (e.g., shards, series, blocks etc) owned by the
	// namespace, which means the memory will be reclaimed the next time GC kicks in and returns the
	// reclaimed memory to the OS. In the future, we might investigate whether it's worth returning
	// the pooled objects to the pools if the pool is low and needs replenishing.
	n.initShards()

	// NB(xichen): possibly also clean up disk files and force a GC here to reclaim memory immediately
	return totalNumSeries, nil
}

func (n *dbNamespace) Repair(repairer databaseShardRepairer) error {
	var (
		numShardsRepaired     int
		numTotalSeries        int64
		numTotalBlocks        int64
		numSizeDiffSeries     int64
		numSizeDiffBlocks     int64
		numChecksumDiffSeries int64
		numChecksumDiffBlocks int64
	)

	multiErr := xerrors.NewMultiError()
	shards := n.getOwnedShards()
	throttle := repairer.Options().RepairShardThrottle()

	for _, shard := range shards {
		metadataRes, err := shard.Repair(n.name, repairer)
		if err != nil {
			multiErr = multiErr.Add(err)
		} else {
			numShardsRepaired++
			numTotalSeries += metadataRes.NumSeries
			numTotalBlocks += metadataRes.NumBlocks
			numSizeDiffSeries += metadataRes.SizeDifferences.NumSeries()
			numSizeDiffBlocks += metadataRes.SizeDifferences.NumBlocks()
			numChecksumDiffSeries += metadataRes.ChecksumDifferences.NumSeries()
			numChecksumDiffBlocks += metadataRes.ChecksumDifferences.NumBlocks()
		}
		if throttle > 0 {
			time.Sleep(throttle)
		}
	}

	n.log.WithFields(
		xlog.NewLogField("namespace", n.name),
		xlog.NewLogField("numTotalShards", len(shards)),
		xlog.NewLogField("numShardsRepaired", numShardsRepaired),
		xlog.NewLogField("numTotalSeries", numTotalSeries),
		xlog.NewLogField("numTotalBlocks", numTotalBlocks),
		xlog.NewLogField("numSizeDiffSeries", numSizeDiffSeries),
		xlog.NewLogField("numSizeDiffBlocks", numSizeDiffBlocks),
		xlog.NewLogField("numChecksumDiffSeries", numChecksumDiffSeries),
		xlog.NewLogField("numChecksumDiffBlocks", numChecksumDiffBlocks),
	).Infof("repair result")

	return multiErr.FinalError()
}

func (n *dbNamespace) getOwnedShards() []databaseShard {
	n.RLock()
	shards := n.shardSet.Shards()
	databaseShards := make([]databaseShard, len(shards))
	for i, shard := range shards {
		databaseShards[i] = n.shards[shard]
	}
	n.RUnlock()
	return databaseShards
}

func (n *dbNamespace) shardAt(shardID uint32) (databaseShard, error) {
	n.RLock()
	shard := n.shards[shardID]
	n.RUnlock()
	if shard == nil {
		return nil, fmt.Errorf("not responsible for shard %d", shardID)
	}
	return shard, nil
}

func (n *dbNamespace) initShards() {
	shards := n.shardSet.Shards()
	dbShards := make([]databaseShard, n.shardSet.Max()+1)
	for _, shard := range shards {
		dbShards[shard] = newDatabaseShard(shard, n.increasingIndex, n.writeCommitLogFn, n.nopts.NeedsBootstrap(), n.sopts)
	}
	n.Lock()
	n.shards = dbShards
	n.Unlock()
}
