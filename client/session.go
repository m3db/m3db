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

package client

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/m3db/m3cluster/shard"
	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/digest"
	"github.com/m3db/m3db/encoding"
	"github.com/m3db/m3db/generated/thrift/rpc"
	"github.com/m3db/m3db/network/server/tchannelthrift/convert"
	"github.com/m3db/m3db/serialize"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/storage/bootstrap/result"
	"github.com/m3db/m3db/storage/index"
	"github.com/m3db/m3db/storage/namespace"
	"github.com/m3db/m3db/topology"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3db/x/xio"
	"github.com/m3db/m3x/checked"
	"github.com/m3db/m3x/context"
	xerrors "github.com/m3db/m3x/errors"
	"github.com/m3db/m3x/ident"
	xlog "github.com/m3db/m3x/log"
	"github.com/m3db/m3x/pool"
	xretry "github.com/m3db/m3x/retry"
	xsync "github.com/m3db/m3x/sync"
	xtime "github.com/m3db/m3x/time"

	"github.com/uber-go/tally"
	"github.com/uber/tchannel-go/thrift"
)

const (
	clusterConnectWaitInterval           = 10 * time.Millisecond
	blocksMetadataInitialCapacity        = 64
	blocksMetadataChannelInitialCapacity = 4096
	gaugeReportInterval                  = 500 * time.Millisecond
	blockMetadataChBufSize               = 4096
)

type resultTypeEnum string

const (
	resultTypeMetadata  resultTypeEnum = "metadata"
	resultTypeBootstrap                = "bootstrap"
	resultTypeRaw                      = "raw"
)

// FetchBlocksMetadataEndpointVersion represents an endpoint version for the
// fetch blocks metadata endpoint
type FetchBlocksMetadataEndpointVersion int

const (
	// FetchBlocksMetadataEndpointDefault represents to use the default fetch blocks metadata endpoint
	FetchBlocksMetadataEndpointDefault FetchBlocksMetadataEndpointVersion = iota
	// FetchBlocksMetadataEndpointV1 represents v1 of the fetch blocks metadata endpoint
	FetchBlocksMetadataEndpointV1
	// FetchBlocksMetadataEndpointV2 represents v2 of the fetch blocks metadata endpoint
	FetchBlocksMetadataEndpointV2 = FetchBlocksMetadataEndpointDefault
)

var (
	validFetchBlocksMetadataEndpoints = []FetchBlocksMetadataEndpointVersion{
		FetchBlocksMetadataEndpointV1,
		FetchBlocksMetadataEndpointV2,
	}
	errFetchBlocksMetadataEndpointVersionUnspecified = errors.New(
		"fetch blocks metadata endpoint version unspecified")
	errUnknownWriteAttemptType = errors.New(
		"unknown write attempt type specified, internal error")
	errNotImplemented = errors.New("not implemented")
)

var (
	// ErrClusterConnectTimeout is raised when connecting to the cluster and
	// ensuring at least each partition has an up node with a connection to it
	ErrClusterConnectTimeout = errors.New("timed out establishing min connections to cluster")
	// errSessionStateNotInitial is raised when trying to open a session and
	// its not in the initial clean state
	errSessionStateNotInitial = errors.New("session not in initial state")
	// errSessionStateNotOpen is raised when operations are requested when the
	// session is not in the open state
	errSessionStateNotOpen = errors.New("session not in open state")
	// errSessionBadBlockResultFromPeer is raised when there is a bad block
	// return from a peer when fetching blocks from peers
	errSessionBadBlockResultFromPeer = errors.New("session fetched bad block result from peer")
	// errSessionInvalidConnectClusterConnectConsistencyLevel is raised when
	// the connect consistency level specified is not recognized
	errSessionInvalidConnectClusterConnectConsistencyLevel = errors.New("session has invalid connect consistency level specified")
	// errSessionHasNoHostQueueForHost is raised when host queue requested for a missing host
	errSessionHasNoHostQueueForHost = errors.New("session has no host queue for host")
	// errInvalidFetchBlocksMetadataVersion is raised when an invalid fetch blocks
	// metadata endpoint version is provided
	errInvalidFetchBlocksMetadataVersion = errors.New("invalid fetch blocks metadata endpoint version")
	// errUnableToEncodeTags is raised when the server is unable to encode provided tags
	// to be sent over the wire.
	errUnableToEncodeTags = errors.New("unable to include tags")
)

type session struct {
	sync.RWMutex

	opts                             Options
	scope                            tally.Scope
	nowFn                            clock.NowFn
	log                              xlog.Logger
	writeLevel                       topology.ConsistencyLevel
	readLevel                        ReadConsistencyLevel
	newHostQueueFn                   newHostQueueFn
	topo                             topology.Topology
	topoMap                          topology.Map
	topoWatch                        topology.MapWatch
	replicas                         int32
	majority                         int32
	queues                           []hostQueue
	queuesByHostID                   map[string]hostQueue
	state                            state
	writeRetrier                     xretry.Retrier
	fetchRetrier                     xretry.Retrier
	streamBlocksRetrier              xretry.Retrier
	contextPool                      context.Pool
	idPool                           ident.Pool
	writeOperationPool               *writeOperationPool
	writeTaggedOperationPool         *writeTaggedOperationPool
	fetchBatchOpPool                 *fetchBatchOpPool
	fetchBatchOpArrayArrayPool       *fetchBatchOpArrayArrayPool
	iteratorArrayPool                encoding.IteratorArrayPool
	tagEncoderPool                   serialize.TagEncoderPool
	readerSliceOfSlicesIteratorPool  *readerSliceOfSlicesIteratorPool
	multiReaderIteratorPool          encoding.MultiReaderIteratorPool
	seriesIteratorPool               encoding.SeriesIteratorPool
	seriesIteratorsPool              encoding.MutableSeriesIteratorsPool
	writeAttemptPool                 *writeAttemptPool
	writeStatePool                   *writeStatePool
	fetchAttemptPool                 *fetchAttemptPool
	fetchBatchSize                   int
	newPeerBlocksQueueFn             newPeerBlocksQueueFn
	reattemptStreamBlocksFromPeersFn reattemptStreamBlocksFromPeersFn
	pickBestPeerFn                   pickBestPeerFn
	origin                           topology.Host
	streamBlocksMaxBlockRetries      int
	streamBlocksWorkers              xsync.WorkerPool
	streamBlocksBatchSize            int
	streamBlocksMetadataBatchTimeout time.Duration
	streamBlocksBatchTimeout         time.Duration
	bootstrapLevel                   ReadConsistencyLevel
	metrics                          sessionMetrics
}

type shardMetricsKey struct {
	shardID    uint32
	resultType resultTypeEnum
}

type sessionMetrics struct {
	sync.RWMutex
	writeSuccess               tally.Counter
	writeErrors                tally.Counter
	writeNodesRespondingErrors []tally.Counter
	fetchSuccess               tally.Counter
	fetchErrors                tally.Counter
	fetchNodesRespondingErrors []tally.Counter
	topologyUpdatedSuccess     tally.Counter
	topologyUpdatedError       tally.Counter
	streamFromPeersMetrics     map[shardMetricsKey]streamFromPeersMetrics
}

func newSessionMetrics(scope tally.Scope) sessionMetrics {
	return sessionMetrics{
		writeSuccess:           scope.Counter("write.success"),
		writeErrors:            scope.Counter("write.errors"),
		fetchSuccess:           scope.Counter("fetch.success"),
		fetchErrors:            scope.Counter("fetch.errors"),
		topologyUpdatedSuccess: scope.Counter("topology.updated-success"),
		topologyUpdatedError:   scope.Counter("topology.updated-error"),
		streamFromPeersMetrics: make(map[shardMetricsKey]streamFromPeersMetrics),
	}
}

type streamFromPeersMetrics struct {
	fetchBlocksFromPeers       tally.Gauge
	metadataFetches            tally.Gauge
	metadataFetchBatchCall     tally.Counter
	metadataFetchBatchSuccess  tally.Counter
	metadataFetchBatchError    tally.Counter
	metadataFetchBatchBlockErr tally.Counter
	metadataReceived           tally.Counter
	fetchBlockSuccess          tally.Counter
	fetchBlockError            tally.Counter
	fetchBlockFullRetry        tally.Counter
	fetchBlockFinalError       tally.Counter
	fetchBlockRetriesReqError  tally.Counter
	fetchBlockRetriesRespError tally.Counter
	blocksEnqueueChannel       tally.Gauge
}

type newHostQueueFn func(
	host topology.Host,
	writeBatchRawRequestPool writeBatchRawRequestPool,
	writeBatchRawRequestElementArrayPool writeBatchRawRequestElementArrayPool,
	writeTaggedBatchRawRequestPool writeTaggedBatchRawRequestPool,
	writeTaggedBatchRawRequestElementArrayPool writeTaggedBatchRawRequestElementArrayPool,
	opts Options,
) hostQueue

func newSession(opts Options) (clientSession, error) {
	topo, err := opts.TopologyInitializer().Init()
	if err != nil {
		return nil, err
	}

	scope := opts.InstrumentOptions().MetricsScope()

	s := &session{
		opts:                 opts,
		scope:                scope,
		nowFn:                opts.ClockOptions().NowFn(),
		log:                  opts.InstrumentOptions().Logger(),
		writeLevel:           opts.WriteConsistencyLevel(),
		readLevel:            opts.ReadConsistencyLevel(),
		newHostQueueFn:       newHostQueue,
		queuesByHostID:       make(map[string]hostQueue),
		topo:                 topo,
		fetchBatchSize:       opts.FetchBatchSize(),
		newPeerBlocksQueueFn: newPeerBlocksQueue,
		writeRetrier:         opts.WriteRetrier(),
		fetchRetrier:         opts.FetchRetrier(),
		contextPool:          opts.ContextPool(),
		idPool:               opts.IdentifierPool(),
		metrics:              newSessionMetrics(scope),
	}
	s.reattemptStreamBlocksFromPeersFn = s.streamBlocksReattemptFromPeers
	s.pickBestPeerFn = s.streamBlocksPickBestPeer
	writeAttemptPoolOpts := pool.NewObjectPoolOptions().
		SetSize(opts.WriteOpPoolSize()).
		SetInstrumentOptions(opts.InstrumentOptions().SetMetricsScope(
			scope.SubScope("write-attempt-pool"),
		))
	s.writeAttemptPool = newWriteAttemptPool(s, writeAttemptPoolOpts)
	s.writeAttemptPool.Init()

	fetchAttemptPoolOpts := pool.NewObjectPoolOptions().
		SetSize(opts.FetchBatchOpPoolSize()).
		SetInstrumentOptions(opts.InstrumentOptions().SetMetricsScope(
			scope.SubScope("fetch-attempt-pool"),
		))
	s.fetchAttemptPool = newFetchAttemptPool(s, fetchAttemptPoolOpts)
	s.fetchAttemptPool.Init()

	tagEncoderPoolOpts := pool.NewObjectPoolOptions().
		SetSize(opts.TagEncoderPoolSize()).
		SetInstrumentOptions(opts.InstrumentOptions().SetMetricsScope(
			scope.SubScope("tag-encoder-pool"),
		))
	s.tagEncoderPool = serialize.NewTagEncoderPool(opts.TagEncoderOptions(), tagEncoderPoolOpts)
	s.tagEncoderPool.Init()

	if opts, ok := opts.(AdminOptions); ok {
		s.origin = opts.Origin()
		s.streamBlocksMaxBlockRetries = opts.FetchSeriesBlocksMaxBlockRetries()
		s.streamBlocksWorkers = xsync.NewWorkerPool(opts.FetchSeriesBlocksBatchConcurrency())
		s.streamBlocksWorkers.Init()
		s.streamBlocksBatchSize = opts.FetchSeriesBlocksBatchSize()
		s.streamBlocksMetadataBatchTimeout = opts.FetchSeriesBlocksMetadataBatchTimeout()
		s.streamBlocksBatchTimeout = opts.FetchSeriesBlocksBatchTimeout()
		s.streamBlocksRetrier = opts.StreamBlocksRetrier()
		s.bootstrapLevel = opts.BootstrapConsistencyLevel()
	}

	return s, nil
}

func (s *session) ShardID(id ident.ID) (uint32, error) {
	s.RLock()
	if s.state != stateOpen {
		s.RUnlock()
		return 0, errSessionStateNotOpen
	}
	value := s.topoMap.ShardSet().Lookup(id)
	s.RUnlock()
	return value, nil
}

// newPeerMetadataStreamingProgressMetrics returns a struct with an embedded
// list of fields that can be used to emit metrics about the current state of
// the peer metadata streaming process
func (s *session) newPeerMetadataStreamingProgressMetrics(
	shard uint32,
	resultType resultTypeEnum,
) *streamFromPeersMetrics {
	mKey := shardMetricsKey{shardID: shard, resultType: resultType}
	s.metrics.RLock()
	m, ok := s.metrics.streamFromPeersMetrics[mKey]
	s.metrics.RUnlock()

	if ok {
		return &m
	}

	scope := s.opts.InstrumentOptions().MetricsScope()

	s.metrics.Lock()
	m, ok = s.metrics.streamFromPeersMetrics[mKey]
	if ok {
		s.metrics.Unlock()
		return &m
	}
	scope = scope.SubScope("stream-from-peers").Tagged(map[string]string{
		"shard":      fmt.Sprintf("%d", shard),
		"resultType": string(resultType),
	})
	m = streamFromPeersMetrics{
		fetchBlocksFromPeers:       scope.Gauge("fetch-blocks-inprogress"),
		metadataFetches:            scope.Gauge("fetch-metadata-peers-inprogress"),
		metadataFetchBatchCall:     scope.Counter("fetch-metadata-peers-batch-call"),
		metadataFetchBatchSuccess:  scope.Counter("fetch-metadata-peers-batch-success"),
		metadataFetchBatchError:    scope.Counter("fetch-metadata-peers-batch-error"),
		metadataFetchBatchBlockErr: scope.Counter("fetch-metadata-peers-batch-block-err"),
		metadataReceived:           scope.Counter("fetch-metadata-peers-received"),
		fetchBlockSuccess:          scope.Counter("fetch-block-success"),
		fetchBlockError:            scope.Counter("fetch-block-error"),
		fetchBlockFinalError:       scope.Counter("fetch-block-final-error"),
		fetchBlockRetriesReqError: scope.Tagged(map[string]string{
			"reason": "request-error",
		}).Counter("fetch-block-retries"),
		fetchBlockRetriesRespError: scope.Tagged(map[string]string{
			"reason": "response-error",
		}).Counter("fetch-block-retries"),
		blocksEnqueueChannel: scope.Gauge("fetch-blocks-enqueue-channel-length"),
	}
	s.metrics.streamFromPeersMetrics[mKey] = m
	s.metrics.Unlock()
	return &m
}

func (s *session) incWriteMetrics(consistencyResultErr error, respErrs int32) {
	if idx := s.nodesRespondingErrorsMetricIndex(respErrs); idx >= 0 {
		s.metrics.writeNodesRespondingErrors[idx].Inc(1)
	}
	if consistencyResultErr == nil {
		s.metrics.writeSuccess.Inc(1)
	} else {
		s.metrics.writeErrors.Inc(1)
	}
}

func (s *session) incFetchMetrics(consistencyResultErr error, respErrs int32) {
	if idx := s.nodesRespondingErrorsMetricIndex(respErrs); idx >= 0 {
		s.metrics.fetchNodesRespondingErrors[idx].Inc(1)
	}
	if consistencyResultErr == nil {
		s.metrics.fetchSuccess.Inc(1)
	} else {
		s.metrics.fetchErrors.Inc(1)
	}
}

func (s *session) nodesRespondingErrorsMetricIndex(respErrs int32) int32 {
	idx := respErrs - 1
	replicas := int32(s.Replicas())
	if respErrs > replicas {
		// Cap to the max replicas, we might get more errors
		// when a node is initializing a shard causing replicas + 1
		// nodes to respond to operations
		idx = replicas - 1
	}
	return idx
}

func (s *session) Open() error {
	s.Lock()
	if s.state != stateNotOpen {
		s.Unlock()
		return errSessionStateNotInitial
	}

	watch, err := s.topo.Watch()
	if err != nil {
		s.Unlock()
		return err
	}

	// Wait for the topology to be available
	<-watch.C()

	topoMap := watch.Get()

	queues, replicas, majority, err := s.hostQueues(topoMap, nil)
	if err != nil {
		s.Unlock()
		return err
	}
	s.setTopologyWithLock(topoMap, queues, replicas, majority)
	s.topoWatch = watch

	// NB(r): Alloc pools that can take some time in Open, expectation
	// is already that Open will take some time
	writeOperationPoolOpts := pool.NewObjectPoolOptions().
		SetSize(s.opts.WriteOpPoolSize()).
		SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
			s.scope.SubScope("write-op-pool"),
		))
	s.writeOperationPool = newWriteOperationPool(writeOperationPoolOpts)
	s.writeOperationPool.Init()

	writeTaggedOperationPoolOpts := pool.NewObjectPoolOptions().
		SetSize(s.opts.WriteTaggedOpPoolSize()).
		SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
			s.scope.SubScope("write-op-tagged-pool"),
		))
	s.writeTaggedOperationPool = newWriteTaggedOpPool(writeTaggedOperationPoolOpts)
	s.writeTaggedOperationPool.Init()

	writeStatePoolSize := s.opts.WriteOpPoolSize()
	if s.opts.WriteTaggedOpPoolSize() > writeStatePoolSize {
		writeStatePoolSize = s.opts.WriteTaggedOpPoolSize()
	}
	writeStatePoolOpts := pool.NewObjectPoolOptions().
		SetSize(writeStatePoolSize).
		SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
			s.scope.SubScope("write-state-pool"),
		))
	s.writeStatePool = newWriteStatePool(s.writeLevel, s.tagEncoderPool, writeStatePoolOpts)
	s.writeStatePool.Init()

	fetchBatchOpPoolOpts := pool.NewObjectPoolOptions().
		SetSize(s.opts.FetchBatchOpPoolSize()).
		SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
			s.scope.SubScope("fetch-batch-op-pool"),
		))
	s.fetchBatchOpPool = newFetchBatchOpPool(fetchBatchOpPoolOpts, s.fetchBatchSize)
	s.fetchBatchOpPool.Init()

	seriesIteratorPoolOpts := pool.NewObjectPoolOptions().
		SetSize(s.opts.SeriesIteratorPoolSize()).
		SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
			s.scope.SubScope("series-iterator-pool"),
		))
	s.seriesIteratorPool = encoding.NewSeriesIteratorPool(seriesIteratorPoolOpts)
	s.seriesIteratorPool.Init()
	s.seriesIteratorsPool = encoding.NewMutableSeriesIteratorsPool(s.opts.SeriesIteratorArrayPoolBuckets())
	s.seriesIteratorsPool.Init()
	s.state = stateOpen
	s.Unlock()

	go func() {
		for range watch.C() {
			s.log.Info("received update for topology")
			topoMap := watch.Get()

			s.RLock()
			existingQueues := s.queues
			s.RUnlock()

			queues, replicas, majority, err := s.hostQueues(topoMap, existingQueues)
			if err != nil {
				s.log.Errorf("could not update topology map: %v", err)
				s.metrics.topologyUpdatedError.Inc(1)
				continue
			}
			s.Lock()
			s.setTopologyWithLock(topoMap, queues, replicas, majority)
			s.Unlock()
			s.metrics.topologyUpdatedSuccess.Inc(1)
		}
	}()

	return nil
}

func (s *session) BorrowConnection(hostID string, fn withConnectionFn) error {
	s.RLock()
	unlocked := false
	queue, ok := s.queuesByHostID[hostID]
	if !ok {
		s.RUnlock()
		return errSessionHasNoHostQueueForHost
	}
	err := queue.BorrowConnection(func(c rpc.TChanNode) {
		// Unlock early on success
		s.RUnlock()
		unlocked = true

		// Execute function with borrowed connection
		fn(c)
	})
	if !unlocked {
		s.RUnlock()
	}
	return err
}

func (s *session) hostQueues(
	topoMap topology.Map,
	existing []hostQueue,
) ([]hostQueue, int, int, error) {
	// NB(r): we leave existing writes in the host queues to finish
	// as they are already enroute to their destination. This is an edge case
	// that might result in leaving nodes counting towards quorum, but fixing it
	// would result in additional chatter.

	start := s.nowFn()

	existingByHostID := make(map[string]hostQueue, len(existing))
	for _, queue := range existing {
		existingByHostID[queue.Host().ID()] = queue
	}

	hosts := topoMap.Hosts()
	queues := make([]hostQueue, 0, len(hosts))
	newQueues := make([]hostQueue, 0, len(hosts))
	for _, host := range hosts {
		if existingQueue, ok := existingByHostID[host.ID()]; ok {
			queues = append(queues, existingQueue)
			continue
		}
		newQueue := s.newHostQueue(host, topoMap)
		queues = append(queues, newQueue)
		newQueues = append(newQueues, newQueue)
	}

	shards := topoMap.ShardSet().AllIDs()
	minConnectionCount := s.opts.MinConnectionCount()
	replicas := topoMap.Replicas()
	majority := topoMap.MajorityReplicas()

	firstConnectConsistencyLevel := s.opts.ClusterConnectConsistencyLevel()
	if firstConnectConsistencyLevel == ConnectConsistencyLevelNone {
		// Return immediately if no connect consistency required
		return queues, replicas, majority, nil
	}

	connectConsistencyLevel := firstConnectConsistencyLevel
	if connectConsistencyLevel == ConnectConsistencyLevelAny {
		// If level any specified, first attempt all then proceed lowering requirement
		connectConsistencyLevel = ConnectConsistencyLevelAll
	}

	// Abort if we do not connect
	connected := false
	defer func() {
		if !connected {
			for _, queue := range newQueues {
				queue.Close()
			}
		}
	}()

	for {
		if now := s.nowFn(); now.Sub(start) >= s.opts.ClusterConnectTimeout() {
			switch firstConnectConsistencyLevel {
			case ConnectConsistencyLevelAny:
				// If connecting with connect any strategy then keep
				// trying but lower consistency requirement
				start = now
				connectConsistencyLevel--
				if connectConsistencyLevel == ConnectConsistencyLevelNone {
					// Already tried to resolve all consistency requirements, just
					// return successfully at this point
					err := fmt.Errorf("timed out connecting, returning success")
					s.log.Warnf("cluster connect with consistency any: %v", err)
					connected = true
					return queues, replicas, majority, nil
				}
			default:
				// Timed out connecting to a specific consistency requirement
				return nil, 0, 0, ErrClusterConnectTimeout
			}
		}
		// Be optimistic
		clusterAvailable := true
		for _, shard := range shards {
			shardReplicasAvailable := 0
			routeErr := topoMap.RouteShardForEach(shard, func(idx int, _ topology.Host) {
				if queues[idx].ConnectionCount() >= minConnectionCount {
					shardReplicasAvailable++
				}
			})
			if routeErr != nil {
				return nil, 0, 0, routeErr
			}
			var clusterAvailableForShard bool
			switch connectConsistencyLevel {
			case ConnectConsistencyLevelAll:
				clusterAvailableForShard = shardReplicasAvailable == replicas
			case ConnectConsistencyLevelMajority:
				clusterAvailableForShard = shardReplicasAvailable >= majority
			case ConnectConsistencyLevelOne:
				clusterAvailableForShard = shardReplicasAvailable > 0
			default:
				return nil, 0, 0, errSessionInvalidConnectClusterConnectConsistencyLevel
			}
			if !clusterAvailableForShard {
				clusterAvailable = false
				break
			}
		}
		if clusterAvailable { // All done
			break
		}
		time.Sleep(clusterConnectWaitInterval)
	}

	connected = true
	return queues, replicas, majority, nil
}

func (s *session) setTopologyWithLock(topoMap topology.Map, queues []hostQueue, replicas, majority int) {
	prevQueues := s.queues

	newQueuesByHostID := make(map[string]hostQueue, len(queues))
	for _, queue := range queues {
		newQueuesByHostID[queue.Host().ID()] = queue
	}

	s.queues = queues
	s.queuesByHostID = newQueuesByHostID

	s.topoMap = topoMap

	atomic.StoreInt32(&s.replicas, int32(replicas))
	atomic.StoreInt32(&s.majority, int32(majority))

	// NB(r): Always recreate the fetch batch op array array pool as it must be
	// the exact length of the queues as we index directly into the return array in
	// in fetch calls
	poolOpts := pool.NewObjectPoolOptions().
		SetSize(s.opts.FetchBatchOpPoolSize()).
		SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
			s.scope.SubScope("fetch-batch-op-array-array-pool"),
		))
	s.fetchBatchOpArrayArrayPool = newFetchBatchOpArrayArrayPool(
		poolOpts,
		len(queues),
		s.opts.FetchBatchOpPoolSize()/len(queues))
	s.fetchBatchOpArrayArrayPool.Init()

	if s.iteratorArrayPool == nil {
		s.iteratorArrayPool = encoding.NewIteratorArrayPool([]pool.Bucket{
			pool.Bucket{
				Capacity: replicas,
				Count:    s.opts.SeriesIteratorPoolSize(),
			},
		})
		s.iteratorArrayPool.Init()
	}
	if s.readerSliceOfSlicesIteratorPool == nil {
		size := replicas * s.opts.SeriesIteratorPoolSize()
		poolOpts := pool.NewObjectPoolOptions().
			SetSize(size).
			SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
				s.scope.SubScope("reader-slice-of-slices-iterator-pool"),
			))
		s.readerSliceOfSlicesIteratorPool = newReaderSliceOfSlicesIteratorPool(poolOpts)
		s.readerSliceOfSlicesIteratorPool.Init()
	}
	if s.multiReaderIteratorPool == nil {
		size := replicas * s.opts.SeriesIteratorPoolSize()
		poolOpts := pool.NewObjectPoolOptions().
			SetSize(size).
			SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
				s.scope.SubScope("multi-reader-iterator-pool"),
			))
		s.multiReaderIteratorPool = encoding.NewMultiReaderIteratorPool(poolOpts)
		s.multiReaderIteratorPool.Init(s.opts.ReaderIteratorAllocate())
	}
	if replicas > len(s.metrics.writeNodesRespondingErrors) {
		curr := len(s.metrics.writeNodesRespondingErrors)
		for i := curr; i < replicas; i++ {
			tags := map[string]string{"nodes": fmt.Sprintf("%d", i+1)}
			counter := s.scope.Tagged(tags).Counter("write.nodes-responding-error")
			s.metrics.writeNodesRespondingErrors =
				append(s.metrics.writeNodesRespondingErrors, counter)
		}
	}
	if replicas > len(s.metrics.fetchNodesRespondingErrors) {
		curr := len(s.metrics.fetchNodesRespondingErrors)
		for i := curr; i < replicas; i++ {
			tags := map[string]string{"nodes": fmt.Sprintf("%d", i+1)}
			counter := s.scope.Tagged(tags).Counter("fetch.nodes-responding-error")
			s.metrics.fetchNodesRespondingErrors =
				append(s.metrics.fetchNodesRespondingErrors, counter)
		}
	}

	// Asynchronously close the set of host queues no longer in use
	go func() {
		for _, queue := range prevQueues {
			newQueue, ok := newQueuesByHostID[queue.Host().ID()]
			if !ok || newQueue != queue {
				queue.Close()
			}
		}
	}()

	s.log.Infof("successfully updated topology to %d hosts", topoMap.HostsLen())
}

func (s *session) newHostQueue(host topology.Host, topoMap topology.Map) hostQueue {
	// NB(r): Due to hosts being replicas we have:
	// = replica * numWrites
	// = total writes to all hosts
	// We need to pool:
	// = replica * (numWrites / writeBatchSize)
	// = number of batch request structs to pool
	// For purposes of simplifying the options for pooling the write op pool size
	// represents the number of ops to pool not including replication, this is due
	// to the fact that the ops are shared between the different host queue replicas.
	writeOpPoolSize := s.opts.WriteOpPoolSize()
	if s.opts.WriteTaggedOpPoolSize() > writeOpPoolSize {
		writeOpPoolSize = s.opts.WriteTaggedOpPoolSize()
	}
	totalBatches := topoMap.Replicas() *
		int(math.Ceil(float64(writeOpPoolSize)/float64(s.opts.WriteBatchSize())))
	hostBatches := int(math.Ceil(float64(totalBatches) / float64(topoMap.HostsLen())))

	writeBatchRequestPoolOpts := pool.NewObjectPoolOptions().
		SetSize(hostBatches).
		SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
			s.scope.SubScope("write-batch-request-pool"),
		))
	writeBatchRequestPool := newWriteBatchRawRequestPool(writeBatchRequestPoolOpts)
	writeBatchRequestPool.Init()

	writeTaggedBatchRequestPoolOpts := pool.NewObjectPoolOptions().
		SetSize(hostBatches).
		SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
			s.scope.SubScope("write-tagged-batch-request-pool"),
		))
	writeTaggedBatchRequestPool := newWriteTaggedBatchRawRequestPool(writeTaggedBatchRequestPoolOpts)
	writeTaggedBatchRequestPool.Init()

	writeBatchRawRequestElementArrayPoolOpts := pool.NewObjectPoolOptions().
		SetSize(hostBatches).
		SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
			s.scope.SubScope("id-datapoint-array-pool"),
		))
	writeBatchRawRequestElementArrayPool := newWriteBatchRawRequestElementArrayPool(
		writeBatchRawRequestElementArrayPoolOpts, s.opts.WriteBatchSize())
	writeBatchRawRequestElementArrayPool.Init()

	writeTaggedBatchRawRequestElementArrayPoolOpts := pool.NewObjectPoolOptions().
		SetSize(hostBatches).
		SetInstrumentOptions(s.opts.InstrumentOptions().SetMetricsScope(
			s.scope.SubScope("id-tagged-datapoint-array-pool"),
		))
	writeTaggedBatchRawRequestElementArrayPool := newWriteTaggedBatchRawRequestElementArrayPool(
		writeTaggedBatchRawRequestElementArrayPoolOpts, s.opts.WriteBatchSize())
	writeTaggedBatchRawRequestElementArrayPool.Init()

	hostQueue := s.newHostQueueFn(host,
		writeBatchRequestPool, writeBatchRawRequestElementArrayPool,
		writeTaggedBatchRequestPool, writeTaggedBatchRawRequestElementArrayPool,
		s.opts)
	hostQueue.Open()
	return hostQueue
}

func (s *session) Write(
	namespace, id ident.ID,
	t time.Time,
	value float64,
	unit xtime.Unit,
	annotation []byte,
) error {
	w := s.writeAttemptPool.Get()
	w.args.attemptType = untaggedWriteAttemptType
	w.args.namespace, w.args.id = namespace, id
	w.args.tags = ident.EmptyTagIterator
	w.args.t, w.args.value, w.args.unit, w.args.annotation =
		t, value, unit, annotation
	err := s.writeRetrier.Attempt(w.attemptFn)
	s.writeAttemptPool.Put(w)
	return err
}

func (s *session) WriteTagged(
	namespace, id ident.ID,
	tags ident.TagIterator,
	t time.Time,
	value float64,
	unit xtime.Unit,
	annotation []byte,
) error {
	w := s.writeAttemptPool.Get()
	w.args.attemptType = taggedWriteAttemptType
	w.args.namespace, w.args.id, w.args.tags = namespace, id, tags
	w.args.t, w.args.value, w.args.unit, w.args.annotation =
		t, value, unit, annotation
	err := s.writeRetrier.Attempt(w.attemptFn)
	s.writeAttemptPool.Put(w)
	return err
}

func (s *session) writeAttempt(
	wType writeAttemptType,
	namespace, id ident.ID,
	inputTags ident.TagIterator,
	t time.Time,
	value float64,
	unit xtime.Unit,
	annotation []byte,
) error {
	timeType, timeTypeErr := convert.ToTimeType(unit)
	if timeTypeErr != nil {
		return timeTypeErr
	}

	timestamp, timestampErr := convert.ToValue(t, timeType)
	if timestampErr != nil {
		return timestampErr
	}

	if s.RLock(); s.state != stateOpen {
		s.RUnlock()
		return errSessionStateNotOpen
	}

	state, majority, enqueued, err := s.writeAttemptWithRLock(
		wType, namespace, id, inputTags, timestamp, value, timeType, annotation)
	s.RUnlock()

	if err != nil {
		return err
	}

	// it's safe to Wait() here, as we still own the lock on state, after it's
	// returned from writeAttemptWithRLock.
	state.Wait()

	err = s.writeConsistencyResult(s.writeLevel, majority, enqueued,
		enqueued-state.pending, int32(len(state.errors)), state.errors)

	s.incWriteMetrics(err, int32(len(state.errors)))

	// must Unlock before decRef'ing, as the latter releases the writeState back into a
	// pool if ref count == 0.
	state.Unlock()
	state.decRef()

	return err
}

// NB(prateek): the returned writeState, if valid, still holds the lock. Its ownership
// is transferred to the calling function, and is expected to manage the lifecycle of
// of the object (including releasing the lock/decRef'ing it).
func (s *session) writeAttemptWithRLock(
	wType writeAttemptType,
	namespace, id ident.ID,
	inputTags ident.TagIterator,
	timestamp int64,
	value float64,
	timeType rpc.TimeType,
	annotation []byte,
) (*writeState, int32, int32, error) {
	var (
		majority = atomic.LoadInt32(&s.majority)
		enqueued int32
	)

	// NB(prateek): We retain an individual copy of the namespace, ID per
	// writeState, as each writeState tracks the lifecycle of it's resources in
	// use in the various queues. Tracking per writeAttempt isn't sufficient as
	// we may enqueue multiple writeStates concurrently depending on retries
	// and consistency level checks.
	nsID := s.idPool.Clone(namespace)
	tsID := s.idPool.Clone(id)
	var tagEncoder serialize.TagEncoder
	if wType == taggedWriteAttemptType {
		tagEncoder = s.tagEncoderPool.Get()
		defer tagEncoder.Finalize()
		if err := tagEncoder.Encode(inputTags); err != nil {
			return nil, 0, 0, err
		}
	}

	var op writeOp
	switch wType {
	case untaggedWriteAttemptType:
		wop := s.writeOperationPool.Get()
		wop.namespace = nsID
		wop.shardID = s.topoMap.ShardSet().Lookup(tsID)
		wop.request.ID = tsID.Data().Get()
		wop.request.Datapoint.Value = value
		wop.request.Datapoint.Timestamp = timestamp
		wop.request.Datapoint.TimestampTimeType = timeType
		wop.request.Datapoint.Annotation = annotation
		op = wop
	case taggedWriteAttemptType:
		wop := s.writeTaggedOperationPool.Get()
		wop.namespace = nsID
		wop.shardID = s.topoMap.ShardSet().Lookup(tsID)
		wop.request.ID = tsID.Data().Get()
		encodedTagBytes, ok := tagEncoder.Data()
		if !ok {
			return nil, 0, 0, errUnableToEncodeTags
		}
		wop.request.EncodedTags = encodedTagBytes.Get()
		wop.request.Datapoint.Value = value
		wop.request.Datapoint.Timestamp = timestamp
		wop.request.Datapoint.TimestampTimeType = timeType
		wop.request.Datapoint.Annotation = annotation
		op = wop
	default:
		// should never happen
		return nil, 0, 0, errUnknownWriteAttemptType
	}

	state := s.writeStatePool.Get()
	state.topoMap = s.topoMap
	state.incRef()

	// todo@bl: Can we combine the writeOpPool and the writeStatePool?
	state.op, state.majority = op, majority
	state.nsID, state.tsID, state.tagEncoder = nsID, tsID, tagEncoder
	op.SetCompletionFn(state.completionFn)

	if err := s.topoMap.RouteForEach(tsID, func(idx int, host topology.Host) {
		// Count pending write requests before we enqueue the completion fns,
		// which rely on the count when executing
		state.pending++
		state.queues = append(state.queues, s.queues[idx])
	}); err != nil {
		state.decRef()
		return nil, 0, 0, err
	}

	state.Lock()
	for i := range state.queues {
		state.incRef()
		if err := state.queues[i].Enqueue(state.op); err != nil {
			state.Unlock()
			state.decRef()

			// NB(r): if this happens we have a bug, once we are in the read
			// lock the current queues should never be closed
			s.opts.InstrumentOptions().Logger().Errorf("failed to enqueue write: %v", err)
			return nil, 0, 0, err
		}
		enqueued++
	}

	// NB(prateek): the current go-routine still holds a lock on the
	// returned writeState object.
	return state, majority, enqueued, nil
}

func (s *session) Fetch(
	namespace ident.ID,
	id ident.ID,
	startInclusive, endExclusive time.Time,
) (encoding.SeriesIterator, error) {
	tsIDs := ident.NewIDsIterator(id)
	results, err := s.FetchIDs(namespace, tsIDs, startInclusive, endExclusive)
	if err != nil {
		return nil, err
	}
	mutableResults := results.(encoding.MutableSeriesIterators)
	iters := mutableResults.Iters()
	iter := iters[0]
	// Reset to zero so that when we close this results set the iter doesn't get closed
	mutableResults.Reset(0)
	mutableResults.Close()
	return iter, nil
}

func (s *session) FetchIDs(
	namespace ident.ID,
	ids ident.Iterator,
	startInclusive, endExclusive time.Time,
) (encoding.SeriesIterators, error) {
	f := s.fetchAttemptPool.Get()
	f.args.namespace, f.args.ids = namespace, ids
	f.args.start, f.args.end = startInclusive, endExclusive
	err := s.fetchRetrier.Attempt(f.attemptFn)
	result := f.result
	s.fetchAttemptPool.Put(f)
	return result, err
}

func (s *session) FetchTagged(
	q index.Query, opts index.QueryOptions,
) (encoding.SeriesIterators, bool, error) {
	return nil, false, errNotImplemented
}

func (s *session) FetchTaggedIDs(
	q index.Query, opts index.QueryOptions,
) (index.QueryResults, error) {
	return index.QueryResults{}, errNotImplemented
}

func (s *session) fetchIDsAttempt(
	inputNamespace ident.ID,
	inputIDs ident.Iterator,
	startInclusive, endExclusive time.Time,
) (encoding.SeriesIterators, error) {
	var (
		wg                     sync.WaitGroup
		allPending             int32
		routeErr               error
		enqueueErr             error
		resultErrLock          sync.RWMutex
		resultErr              error
		resultErrs             int32
		majority               int32
		consistencyLevel       ReadConsistencyLevel
		fetchBatchOpsByHostIdx [][]*fetchBatchOp
		success                = false
	)

	// NB(prateek): need to make a copy of inputNamespace and inputIDs to control
	// their life-cycle within this function.
	namespace := s.idPool.Clone(inputNamespace)
	// First, we duplicate the iterator (only the struct referencing the underlying slice,
	// not the slice itself). Need this to be able to iterate the original iterator
	// multiple times in case of retries.
	ids := inputIDs.Duplicate()

	rangeStart, tsErr := convert.ToValue(startInclusive, rpc.TimeType_UNIX_NANOSECONDS)
	if tsErr != nil {
		return nil, tsErr
	}

	rangeEnd, tsErr := convert.ToValue(endExclusive, rpc.TimeType_UNIX_NANOSECONDS)
	if tsErr != nil {
		return nil, tsErr
	}

	s.RLock()
	if s.state != stateOpen {
		s.RUnlock()
		return nil, errSessionStateNotOpen
	}

	iters := s.seriesIteratorsPool.Get(ids.Remaining())
	iters.Reset(ids.Remaining())

	defer func() {
		// NB(r): Ensure we cover all edge cases and close the iters in any case
		// of an error being returned
		if !success {
			iters.Close()
		}
	}()

	// NB(r): We must take and return pooled items in the session read lock for the
	// pools that change during a topology update.
	// This is due to when a queue is re-initialized it enqueues a fixed number
	// of entries into the backing channel for the pool and will forever stall
	// on the last few puts if any unexpected entries find their way there
	// while it is filling.
	fetchBatchOpsByHostIdx = s.fetchBatchOpArrayArrayPool.Get()

	consistencyLevel = s.readLevel
	majority = atomic.LoadInt32(&s.majority)

	// NB(prateek): namespaceAccessors tracks the number of pending accessors for nsID.
	// It is set to incremented by `replica` for each requested ID during fetch enqueuing,
	// and once by initial request, and is decremented for each replica retrieved, inside
	// completionFn, and once by the allCompletionFn. So know we can Finalize `namespace`
	// once it's value reaches 0.
	namespaceAccessors := int32(0)

	for idx := 0; ids.Next(); idx++ {
		var (
			idx  = idx // capture loop variable
			tsID = s.idPool.Clone(ids.Current())

			wgIsDone int32
			// NB(xichen): resultsAccessors and idAccessors get initialized to number of replicas + 1
			// before enqueuing (incremented when iterating over the replicas for this ID), and gets
			// decremented for each replica as well as inside the allCompletionFn so we know when
			// resultsAccessors is 0, results are no longer accessed and it's safe to return results
			// to the pool.
			resultsAccessors int32 = 1
			idAccessors      int32 = 1
			resultsLock      sync.RWMutex
			results          []encoding.Iterator
			enqueued         int32
			pending          int32
			success          int32
			errors           []error
			errs             int32
		)

		// increment namespaceAccesors by 1 to indicate it still needs to be handled by the
		// allCompletionFn for tsID.
		atomic.AddInt32(&namespaceAccessors, 1)

		wg.Add(1)
		allCompletionFn := func() {
			var reportErrors []error
			errsLen := atomic.LoadInt32(&errs)
			if errsLen > 0 {
				resultErrLock.RLock()
				reportErrors = errors[:]
				resultErrLock.RUnlock()
			}
			responded := enqueued - atomic.LoadInt32(&pending)
			err := s.readConsistencyResult(consistencyLevel, majority, enqueued,
				responded, errsLen, reportErrors)
			s.incFetchMetrics(err, errsLen)
			if err != nil {
				resultErrLock.Lock()
				if resultErr == nil {
					resultErr = err
				}
				resultErrs++
				resultErrLock.Unlock()
			} else {
				resultsLock.RLock()
				successIters := results[:success]
				resultsLock.RUnlock()
				iter := s.seriesIteratorPool.Get()
				// NB(prateek): we need to allocate a copy of ident.ID to allow the seriesIterator
				// to have control over the lifecycle of ID. We cannot allow seriesIterator
				// to control the lifecycle of the original ident.ID, as it might still be in use
				// due to a pending request in queue.
				seriesID := s.idPool.Clone(tsID)
				namespaceID := s.idPool.Clone(namespace)
				iter.Reset(seriesID, namespaceID, startInclusive, endExclusive, successIters)
				iters.SetAt(idx, iter)
			}
			if atomic.AddInt32(&resultsAccessors, -1) == 0 {
				s.iteratorArrayPool.Put(results)
			}
			if atomic.AddInt32(&idAccessors, -1) == 0 {
				tsID.Finalize()
			}
			if atomic.AddInt32(&namespaceAccessors, -1) == 0 {
				namespace.Finalize()
			}
			wg.Done()
		}
		completionFn := func(result interface{}, err error) {
			var snapshotSuccess int32
			if err != nil {
				atomic.AddInt32(&errs, 1)
				// NB(r): reuse the error lock here as we do not want to create
				// a whole lot of locks for every single ID fetched due to size
				// of mutex being non-trivial and likely to cause more stack growth
				// or GC pressure if ends up on heap which is likely due to naive
				// escape analysis.
				resultErrLock.Lock()
				errors = append(errors, err)
				resultErrLock.Unlock()
			} else {
				slicesIter := s.readerSliceOfSlicesIteratorPool.Get()
				slicesIter.Reset(result.([]*rpc.Segments))
				multiIter := s.multiReaderIteratorPool.Get()
				multiIter.ResetSliceOfSlices(slicesIter)
				// Results is pre-allocated after creating fetch ops for this ID below
				resultsLock.Lock()
				results[success] = multiIter
				success++
				snapshotSuccess = success
				resultsLock.Unlock()
			}
			// NB(xichen): decrementing pending and checking remaining against zero must
			// come after incrementing success, otherwise we might end up passing results[:success]
			// to iter.Reset down below before setting the iterator in the results array,
			// which would cause a nil pointer exception.
			remaining := atomic.AddInt32(&pending, -1)
			doneAll := remaining == 0
			switch s.readLevel {
			case ReadConsistencyLevelOne:
				complete := snapshotSuccess > 0 || doneAll
				if complete && atomic.CompareAndSwapInt32(&wgIsDone, 0, 1) {
					allCompletionFn()
				}
			case ReadConsistencyLevelMajority, ReadConsistencyLevelUnstrictMajority:
				complete := snapshotSuccess >= majority || doneAll
				if complete && atomic.CompareAndSwapInt32(&wgIsDone, 0, 1) {
					allCompletionFn()
				}
			case ReadConsistencyLevelAll:
				if doneAll && atomic.CompareAndSwapInt32(&wgIsDone, 0, 1) {
					allCompletionFn()
				}
			}

			if atomic.AddInt32(&resultsAccessors, -1) == 0 {
				s.iteratorArrayPool.Put(results)
			}
			if atomic.AddInt32(&idAccessors, -1) == 0 {
				tsID.Finalize()
			}
			if atomic.AddInt32(&namespaceAccessors, -1) == 0 {
				namespace.Finalize()
			}
		}

		if err := s.topoMap.RouteForEach(tsID, func(hostIdx int, host topology.Host) {
			// Inc safely as this for each is sequential
			enqueued++
			pending++
			allPending++
			resultsAccessors++
			namespaceAccessors++
			idAccessors++

			ops := fetchBatchOpsByHostIdx[hostIdx]

			var f *fetchBatchOp
			if len(ops) > 0 {
				// Find the last and potentially current fetch op for this host
				f = ops[len(ops)-1]
			}
			if f == nil || f.Size() >= s.fetchBatchSize {
				// If no current fetch op or existing one is at batch capacity add one
				// NB(r): Note that we defer to the host queue to take ownership
				// of these ops and for returning the ops to the pool when done as
				// they know when their use is complete.
				f = s.fetchBatchOpPool.Get()
				f.IncRef()
				fetchBatchOpsByHostIdx[hostIdx] = append(fetchBatchOpsByHostIdx[hostIdx], f)
				f.request.RangeStart = rangeStart
				f.request.RangeEnd = rangeEnd
				f.request.RangeTimeType = rpc.TimeType_UNIX_NANOSECONDS
			}

			// Append IDWithNamespace to this request
			f.append(namespace.Data().Get(), tsID.Data().Get(), completionFn)
		}); err != nil {
			routeErr = err
			break
		}

		// Once we've enqueued we know how many to expect so retrieve and set length
		results = s.iteratorArrayPool.Get(int(enqueued))
		results = results[:enqueued]
	}

	if routeErr != nil {
		s.RUnlock()
		return nil, routeErr
	}

	// Enqueue fetch ops
	for idx := range fetchBatchOpsByHostIdx {
		for _, f := range fetchBatchOpsByHostIdx[idx] {
			// Passing ownership of the op itself to the host queue
			f.DecRef()
			if err := s.queues[idx].Enqueue(f); err != nil && enqueueErr == nil {
				enqueueErr = err
				break
			}
		}
		if enqueueErr != nil {
			break
		}
	}
	s.fetchBatchOpArrayArrayPool.Put(fetchBatchOpsByHostIdx)
	s.RUnlock()

	if enqueueErr != nil {
		s.log.Errorf("failed to enqueue fetch: %v", enqueueErr)
		return nil, enqueueErr
	}

	wg.Wait()

	resultErrLock.RLock()
	retErr := resultErr
	resultErrLock.RUnlock()
	if retErr != nil {
		return nil, retErr
	}
	success = true
	return iters, nil
}

func (s *session) writeConsistencyResult(
	level topology.ConsistencyLevel,
	majority, enqueued, responded, resultErrs int32,
	errs []error,
) error {
	// Check consistency level satisfied
	success := enqueued - resultErrs
	if !s.writeConsistencyAchieved(level, int(majority), int(enqueued), int(success)) {
		return newConsistencyResultError(level, int(enqueued), int(responded), errs)
	}
	return nil
}

func (s *session) writeConsistencyAchieved(level topology.ConsistencyLevel, majority, enqueued, success int) bool {
	switch level {
	case topology.ConsistencyLevelAll:
		if success == enqueued { // Meets all
			return true
		}
	case topology.ConsistencyLevelMajority:
		if success >= majority { // Meets majority
			return true
		}
	case topology.ConsistencyLevelOne:
		if success > 0 { // Meets one
			return true
		}
	}
	return false
}

func (s *session) readConsistencyResult(
	level ReadConsistencyLevel,
	majority, enqueued, responded, resultErrs int32,
	errs []error,
) error {
	// Check consistency level satisfied
	success := enqueued - resultErrs
	if !s.readConsistencyAchieved(level, int(majority), int(enqueued), int(success)) {
		return newConsistencyResultError(level, int(enqueued), int(responded), errs)
	}
	return nil
}

func (s *session) readConsistencyAchieved(level ReadConsistencyLevel, majority, enqueued, success int) bool {
	switch level {
	case ReadConsistencyLevelAll:
		if success == enqueued { // Meets all
			return true
		}
	case ReadConsistencyLevelMajority:
		if success >= majority { // Meets majority
			return true
		}
	case ReadConsistencyLevelOne, ReadConsistencyLevelUnstrictMajority:
		if success > 0 { // Meets one
			return true
		}
	}
	return false
}

func (s *session) Close() error {
	s.Lock()
	if s.state != stateOpen {
		s.Unlock()
		return errSessionStateNotOpen
	}
	s.state = stateClosed
	s.Unlock()

	for _, q := range s.queues {
		q.Close()
	}

	s.topoWatch.Close()
	s.topo.Close()
	return nil
}

func (s *session) Origin() topology.Host {
	return s.origin
}

func (s *session) Replicas() int {
	return int(atomic.LoadInt32(&s.replicas))
}

func (s *session) Truncate(namespace ident.ID) (int64, error) {
	var (
		wg            sync.WaitGroup
		enqueueErr    xerrors.MultiError
		resultErrLock sync.Mutex
		resultErr     xerrors.MultiError
		truncated     int64
	)

	t := &truncateOp{}
	t.request.NameSpace = namespace.Data().Get()
	t.completionFn = func(result interface{}, err error) {
		if err != nil {
			resultErrLock.Lock()
			resultErr = resultErr.Add(err)
			resultErrLock.Unlock()
		} else {
			res := result.(*rpc.TruncateResult_)
			atomic.AddInt64(&truncated, res.NumSeries)
		}
		wg.Done()
	}

	s.RLock()
	for idx := range s.queues {
		wg.Add(1)
		if err := s.queues[idx].Enqueue(t); err != nil {
			wg.Done()
			enqueueErr = enqueueErr.Add(err)
		}
	}
	s.RUnlock()

	if err := enqueueErr.FinalError(); err != nil {
		s.log.Errorf("failed to enqueue request: %v", err)
		return 0, err
	}

	// Wait for namespace to be truncated on all replicas
	wg.Wait()

	return truncated, resultErr.FinalError()
}

type peers struct {
	peers            []peer
	selfExcluded     bool
	selfHostShardSet topology.HostShardSet
}

func (s *session) peersForShard(shard uint32) (peers, error) {
	var (
		lookupErr error
		result    = peers{peers: make([]peer, 0, s.topoMap.Replicas())}
	)

	s.RLock()
	err := s.topoMap.RouteShardForEach(shard, func(idx int, host topology.Host) {
		if s.origin != nil && s.origin.ID() == host.ID() {
			// Don't include the origin host
			result.selfExcluded = true
			// Include the origin host shard set for help determining quorum
			hostShardSet, ok := s.topoMap.LookupHostShardSet(host.ID())
			if !ok {
				lookupErr = fmt.Errorf("could not find shard set for host ID: %s", host.ID())
			}
			result.selfHostShardSet = hostShardSet
			return
		}
		result.peers = append(result.peers, newPeer(s, host))
	})
	s.RUnlock()
	if resultErr := xerrors.FirstError(err, lookupErr); resultErr != nil {
		return peers{}, resultErr
	}
	return result, nil
}

func (s *session) FetchBlocksMetadataFromPeers(
	namespace ident.ID,
	shard uint32,
	start, end time.Time,
	consistencyLevel ReadConsistencyLevel,
	resultOpts result.Options,
	version FetchBlocksMetadataEndpointVersion,
) (PeerBlockMetadataIter, error) {
	peers, err := s.peersForShard(shard)
	if err != nil {
		return nil, err
	}

	var (
		metadataCh = make(chan receivedBlockMetadata,
			blocksMetadataChannelInitialCapacity)
		errCh = make(chan error, 1)
		meta  = resultTypeMetadata
		m     = s.newPeerMetadataStreamingProgressMetrics(shard, meta)
	)
	go func() {
		errCh <- s.streamBlocksMetadataFromPeers(namespace, shard,
			peers, start, end, consistencyLevel, metadataCh, resultOpts, m, version)
		close(metadataCh)
		close(errCh)
	}()

	return newMetadataIter(metadataCh, errCh), nil
}

// FetchBootstrapBlocksFromPeers will fetch the specified blocks from peers for
// bootstrapping purposes. Refer to peer_bootstrapping.md for more details.
func (s *session) FetchBootstrapBlocksFromPeers(
	nsMetadata namespace.Metadata,
	shard uint32,
	start, end time.Time,
	opts result.Options,
	version FetchBlocksMetadataEndpointVersion,
) (result.ShardResult, error) {
	if !IsValidFetchBlocksMetadataEndpoint(version) {
		return nil, errInvalidFetchBlocksMetadataVersion
	}

	var (
		result           = newBulkBlocksResult(s.opts, opts)
		doneCh           = make(chan struct{})
		progress         = s.newPeerMetadataStreamingProgressMetrics(shard, resultTypeBootstrap)
		consistencyLevel = s.bootstrapLevel
	)

	// Determine which peers own the specified shard
	peers, err := s.peersForShard(shard)
	if err != nil {
		return nil, err
	}

	// Emit a gauge indicating whether we're done or not
	go func() {
		for {
			select {
			case <-doneCh:
				progress.fetchBlocksFromPeers.Update(0)
				return
			default:
				progress.fetchBlocksFromPeers.Update(1)
				time.Sleep(gaugeReportInterval)
			}
		}
	}()
	defer close(doneCh)

	// Begin pulling metadata, if one or multiple peers fail no error will
	// be returned from this routine as long as one peer succeeds completely
	metadataCh := make(chan receivedBlockMetadata, blockMetadataChBufSize)
	// Spin up a background goroutine which will begin streaming metadata from
	// all the peers and pushing them into the metadatach
	errCh := make(chan error, 1)
	go func() {
		errCh <- s.streamBlocksMetadataFromPeers(nsMetadata.ID(), shard,
			peers, start, end, consistencyLevel, metadataCh, opts, progress, version)
		close(metadataCh)
	}()

	// Begin consuming metadata and making requests. This will block until all
	// data has been streamed (or failed to stream). Note that this function does
	// not return an error and if anything goes wrong here we won't report it to
	// the caller, but metrics and logs are emitted internally. Also note that the
	// streamAndGroupCollectedBlocksMetadata function is injected.
	s.streamBlocksFromPeers(nsMetadata, shard, peers.peers, metadataCh,
		opts, result, progress, s.streamAndGroupCollectedBlocksMetadata)

	// Check if an error occurred during the metadata streaming
	if err = <-errCh; err != nil {
		return nil, err
	}

	return result.result, nil
}

func (s *session) FetchBlocksFromPeers(
	nsMetadata namespace.Metadata,
	shard uint32,
	metadatas []block.ReplicaMetadata,
	opts result.Options,
) (PeerBlocksIter, error) {

	var (
		logger   = opts.InstrumentOptions().Logger()
		complete = int64(0)
		doneCh   = make(chan error, 1)
		outputCh = make(chan peerBlocksDatapoint, 4096)
		result   = newStreamBlocksResult(s.opts, opts, outputCh)
		onDone   = func(err error) {
			atomic.StoreInt64(&complete, 1)
			select {
			case doneCh <- err:
			default:
			}
		}
		progress = s.newPeerMetadataStreamingProgressMetrics(shard, resultTypeRaw)
	)

	peers, err := s.peersForShard(shard)
	if err != nil {
		return nil, err
	}
	peersByHost := make(map[string]peer, len(peers.peers))
	for _, peer := range peers.peers {
		peersByHost[peer.Host().ID()] = peer
	}

	go func() {
		for atomic.LoadInt64(&complete) == 0 {
			progress.fetchBlocksFromPeers.Update(1)
			time.Sleep(gaugeReportInterval)
		}
		progress.fetchBlocksFromPeers.Update(0)
	}()

	metadataCh := make(chan receivedBlockMetadata, blockMetadataChBufSize)
	go func() {
		for _, rb := range metadatas {
			peer, ok := peersByHost[rb.Host.ID()]
			if !ok {
				logger.WithFields(
					xlog.NewField("peer", rb.Host.String()),
					xlog.NewField("id", rb.ID.String()),
					xlog.NewField("start", rb.Start.String()),
				).Warnf("replica requested from unknown peer, skipping")
				continue
			}
			metadataCh <- receivedBlockMetadata{
				id:   rb.ID,
				peer: peer,
				block: blockMetadata{
					start:    rb.Start,
					size:     rb.Size,
					checksum: rb.Checksum,
					lastRead: rb.LastRead,
				},
			}
		}
		close(metadataCh)
	}()

	// Begin consuming metadata and making requests
	go func() {
		s.streamBlocksFromPeers(nsMetadata, shard, peers.peers, metadataCh,
			opts, result, progress, s.passThroughBlocksMetadata)
		close(outputCh)
		onDone(nil)
	}()

	pbi := newPeerBlocksIter(outputCh, doneCh)
	return pbi, nil
}

func (s *session) streamBlocksMetadataFromPeers(
	namespace ident.ID,
	shardID uint32,
	peers peers,
	start, end time.Time,
	consistencyLevel ReadConsistencyLevel,
	metadataCh chan<- receivedBlockMetadata,
	resultOpts result.Options,
	progress *streamFromPeersMetrics,
	version FetchBlocksMetadataEndpointVersion,
) error {
	var (
		wg       sync.WaitGroup
		errLock  sync.RWMutex
		errors   = make(map[int]error)
		setError = func(idx int, err error) {
			errLock.Lock()
			errors[idx] = err
			errLock.Unlock()
		}
		getErrors = func() []error {
			var result []error
			errLock.RLock()
			for _, err := range errors {
				if err == nil {
					continue
				}
				result = append(result, err)
			}
			errLock.RUnlock()
			return result
		}
		abortError    error
		setAbortError = func(err error) {
			errLock.Lock()
			abortError = err
			errLock.Unlock()
		}
		getAbortError = func() error {
			errLock.RLock()
			result := abortError
			errLock.RUnlock()
			return result
		}
		pending   = int64(len(peers.peers))
		majority  = atomic.LoadInt32(&s.majority)
		enqueued  = int32(len(peers.peers))
		responded int32
		success   int32
	)
	if peers.selfExcluded {
		state, err := peers.selfHostShardSet.ShardSet().LookupStateByID(shardID)
		if err == nil && state == shard.Available {
			// If we excluded ourselves from fetching, we basically treat ourselves
			// as a successful peer response since we can bootstrap from ourselves
			// just fine
			enqueued++
			success++
		}
	}

	progress.metadataFetches.Update(float64(pending))
	for idx, peer := range peers.peers {
		idx := idx
		peer := peer

		wg.Add(1)
		go func() {
			defer func() {
				// Success or error counts towards a response
				atomic.AddInt32(&responded, 1)

				// Decrement pending
				progress.metadataFetches.Update(float64(atomic.AddInt64(&pending, -1)))

				// Mark done
				wg.Done()
			}()

			var (
				firstAttempt  = true
				currPageToken pageToken
			)
			condition := func() bool {
				if firstAttempt {
					// Always attempt at least once
					firstAttempt = false
					return true
				}
				level := consistencyLevel
				majority := int(majority)
				enqueued := int(enqueued)
				success := int(atomic.LoadInt32(&success))
				return !s.readConsistencyAchieved(level, majority, enqueued, success) &&
					getAbortError() == nil
			}
			for condition() {
				var err error
				switch version {
				case FetchBlocksMetadataEndpointV1:
					currPageToken, err = s.streamBlocksMetadataFromPeer(namespace, shardID,
						peer, start, end, currPageToken, metadataCh, progress)
				case FetchBlocksMetadataEndpointV2:
					currPageToken, err = s.streamBlocksMetadataFromPeerV2(namespace, shardID,
						peer, start, end, currPageToken, metadataCh, resultOpts, progress)
				default:
					// Should never happen - we validate the version before this function is
					// ever called
					err = xerrors.NewNonRetryableError(errInvalidFetchBlocksMetadataVersion)
				}

				// Set error or success if err is nil
				setError(idx, err)

				// Check exit criteria
				if err != nil && xerrors.IsNonRetryableError(err) {
					setAbortError(err)
					return // Cannot recover from this error, so we break from the loop
				}
				if err == nil {
					atomic.AddInt32(&success, 1)
					return
				}
			}
		}()
	}

	wg.Wait()

	if err := getAbortError(); err != nil {
		return err
	}

	errs := getErrors()
	return s.readConsistencyResult(consistencyLevel, majority, enqueued,
		atomic.LoadInt32(&responded), int32(len(errs)), errs)
}

// pageToken is just an opaque type that needs to be downcasted to expected
// page token type, this makes it easy to use the page token across the two
// versions
type pageToken interface{}

// TODO(rartoul): Delete this once we delete the V1 code path
func (s *session) streamBlocksMetadataFromPeer(
	namespace ident.ID,
	shard uint32,
	peer peer,
	start, end time.Time,
	startPageToken pageToken,
	ch chan<- receivedBlockMetadata,
	progress *streamFromPeersMetrics,
) (pageToken, error) {
	var pageToken *int64
	if startPageToken != nil {
		var ok bool
		pageToken, ok = startPageToken.(*int64)
		if !ok {
			err := fmt.Errorf("unexpected start page token type: %s",
				reflect.TypeOf(startPageToken).Elem().String())
			return nil, xerrors.NewNonRetryableError(err)
		}
	}

	var (
		optionIncludeSizes     = true
		optionIncludeChecksums = true
		optionIncludeLastRead  = true
		moreResults            = true

		// Only used for logs
		peerStr              = peer.Host().ID()
		metadataCountByBlock = map[xtime.UnixNano]int64{}
	)

	// Only used for logs
	defer func() {
		for block, numMetadata := range metadataCountByBlock {
			s.log.WithFields(
				xlog.NewField("shard", shard),
				xlog.NewField("peer", peerStr),
				xlog.NewField("numMetadata", numMetadata),
				xlog.NewField("block", block),
			).Debug("finished streaming blocks metadata from peer")
		}
	}()

	// Declare before loop to avoid redeclaring each iteration
	attemptFn := func(client rpc.TChanNode) error {
		tctx, _ := thrift.NewContext(s.streamBlocksMetadataBatchTimeout)
		req := rpc.NewFetchBlocksMetadataRawRequest()
		req.NameSpace = namespace.Data().Get()
		req.Shard = int32(shard)
		req.RangeStart = start.UnixNano()
		req.RangeEnd = end.UnixNano()
		req.Limit = int64(s.streamBlocksBatchSize)
		req.PageToken = pageToken
		req.IncludeSizes = &optionIncludeSizes
		req.IncludeChecksums = &optionIncludeChecksums
		req.IncludeLastRead = &optionIncludeLastRead

		progress.metadataFetchBatchCall.Inc(1)
		result, err := client.FetchBlocksMetadataRaw(tctx, req)
		if err != nil {
			progress.metadataFetchBatchError.Inc(1)
			return err
		}

		progress.metadataFetchBatchSuccess.Inc(1)
		progress.metadataReceived.Inc(int64(len(result.Elements)))

		if result.NextPageToken != nil {
			// Create space on the heap for the page token and take it's
			// address to avoid having to keep the entire result around just
			// for the page token
			resultPageToken := *result.NextPageToken
			pageToken = &resultPageToken
		} else {
			// No further results
			moreResults = false
		}

		for _, elem := range result.Elements {
			blockID := ident.BinaryID(checked.NewBytes(elem.ID, nil))
			for _, b := range elem.Blocks {
				blockStart := time.Unix(0, b.Start)

				// Error occurred retrieving block metadata, use default values
				if b.Err != nil {
					progress.metadataFetchBatchBlockErr.Inc(1)
					s.log.WithFields(
						xlog.NewField("shard", shard),
						xlog.NewField("peer", peerStr),
						xlog.NewField("block", blockStart),
						xlog.NewField("error", err),
					).Error("error occurred retrieving block metadata")
					// Enqueue with a zerod checksum which cause a fanout fetch
					ch <- receivedBlockMetadata{
						peer: peer,
						id:   blockID,
						block: blockMetadata{
							start: blockStart,
						},
					}
					continue
				}

				var size int64
				if b.Size != nil {
					size = *b.Size
				}

				var pChecksum *uint32
				if b.Checksum != nil {
					value := uint32(*b.Checksum)
					pChecksum = &value
				}

				var lastRead time.Time
				if b.LastRead != nil {
					value, err := convert.ToTime(*b.LastRead, b.LastReadTimeType)
					if err == nil {
						lastRead = value
					}
				}

				ch <- receivedBlockMetadata{
					peer: peer,
					id:   blockID,
					block: blockMetadata{
						start:    blockStart,
						size:     size,
						checksum: pChecksum,
						lastRead: lastRead,
					},
				}

				// Only used for logs
				metadataCountByBlock[xtime.ToUnixNano(blockStart)]++
			}
		}

		return nil
	}

	// NB(r): split the following methods up so they don't allocate
	// a closure per fetch blocks call
	var attemptErr error
	checkedAttemptFn := func(client rpc.TChanNode) {
		attemptErr = attemptFn(client)
	}

	fetchFn := func() error {
		borrowErr := peer.BorrowConnection(checkedAttemptFn)
		return xerrors.FirstError(borrowErr, attemptErr)
	}

	for moreResults {
		if err := s.streamBlocksRetrier.Attempt(fetchFn); err != nil {
			return pageToken, err
		}
	}
	return nil, nil
}

// streamBlocksMetadataFromPeerV2 has several heap allocated anonymous
// function, however, they're only allocated once per peer/shard combination
// for the entire peer bootstrapping process so performance is acceptable
func (s *session) streamBlocksMetadataFromPeerV2(
	namespace ident.ID,
	shard uint32,
	peer peer,
	start, end time.Time,
	startPageToken pageToken,
	metadataCh chan<- receivedBlockMetadata,
	resultOpts result.Options,
	progress *streamFromPeersMetrics,
) (pageToken, error) {
	var pageToken []byte
	if startPageToken != nil {
		var ok bool
		pageToken, ok = startPageToken.([]byte)
		if !ok {
			err := fmt.Errorf("unexpected start page token type: %s",
				reflect.TypeOf(startPageToken).Elem().String())
			return nil, xerrors.NewNonRetryableError(err)
		}
	}

	var (
		optionIncludeSizes     = true
		optionIncludeChecksums = true
		optionIncludeLastRead  = true
		moreResults            = true
		idPool                 = s.idPool
		bytesPool              = resultOpts.DatabaseBlockOptions().BytesPool()

		// Only used for logs
		peerStr              = peer.Host().ID()
		metadataCountByBlock = map[xtime.UnixNano]int64{}
	)

	// Only used for logs
	defer func() {
		for block, numMetadata := range metadataCountByBlock {
			s.log.WithFields(
				xlog.NewField("shard", shard),
				xlog.NewField("peer", peerStr),
				xlog.NewField("numMetadata", numMetadata),
				xlog.NewField("block", block),
			).Debug("finished streaming blocks metadata from peer")
		}
	}()

	// Declare before loop to avoid redeclaring each iteration
	attemptFn := func(client rpc.TChanNode) error {
		tctx, _ := thrift.NewContext(s.streamBlocksMetadataBatchTimeout)
		req := rpc.NewFetchBlocksMetadataRawV2Request()
		req.NameSpace = namespace.Data().Get()
		req.Shard = int32(shard)
		req.RangeStart = start.UnixNano()
		req.RangeEnd = end.UnixNano()
		req.Limit = int64(s.streamBlocksBatchSize)
		req.PageToken = pageToken
		req.IncludeSizes = &optionIncludeSizes
		req.IncludeChecksums = &optionIncludeChecksums
		req.IncludeLastRead = &optionIncludeLastRead

		progress.metadataFetchBatchCall.Inc(1)
		result, err := client.FetchBlocksMetadataRawV2(tctx, req)
		if err != nil {
			progress.metadataFetchBatchError.Inc(1)
			return err
		}

		progress.metadataFetchBatchSuccess.Inc(1)
		progress.metadataReceived.Inc(int64(len(result.Elements)))

		if result.NextPageToken != nil {
			// Reset pageToken + copy new pageToken into previously allocated memory,
			// extending as necessary
			pageToken = append(pageToken[:0], result.NextPageToken...)
		} else {
			// No further results
			moreResults = false
		}

		for _, elem := range result.Elements {
			blockStart := time.Unix(0, elem.Start)

			data := bytesPool.Get(len(elem.ID))
			data.IncRef()
			data.AppendAll(elem.ID)
			data.DecRef()

			clonedID := idPool.BinaryID(data)

			// Error occurred retrieving block metadata, use default values
			if elem.Err != nil {
				progress.metadataFetchBatchBlockErr.Inc(1)
				s.log.WithFields(
					xlog.NewField("shard", shard),
					xlog.NewField("peer", peerStr),
					xlog.NewField("block", blockStart),
					xlog.NewField("error", err),
				).Error("error occurred retrieving block metadata")
				// Enqueue with a zerod checksum which cause a fanout fetch
				metadataCh <- receivedBlockMetadata{
					peer: peer,
					id:   clonedID,
					block: blockMetadata{
						start: blockStart,
					},
				}
				continue
			}

			var size int64
			if elem.Size != nil {
				size = *elem.Size
			}

			var pChecksum *uint32
			if elem.Checksum != nil {
				value := uint32(*elem.Checksum)
				pChecksum = &value
			}

			var lastRead time.Time
			if elem.LastRead != nil {
				value, err := convert.ToTime(*elem.LastRead, elem.LastReadTimeType)
				if err == nil {
					lastRead = value
				}
			}

			metadataCh <- receivedBlockMetadata{
				peer: peer,
				id:   clonedID,
				block: blockMetadata{
					start:    blockStart,
					size:     size,
					checksum: pChecksum,
					lastRead: lastRead,
				},
			}
			// Only used for logs
			metadataCountByBlock[xtime.ToUnixNano(blockStart)]++
		}
		return nil
	}

	var attemptErr error
	checkedAttemptFn := func(client rpc.TChanNode) {
		attemptErr = attemptFn(client)
	}

	fetchFn := func() error {
		borrowErr := peer.BorrowConnection(checkedAttemptFn)
		return xerrors.FirstError(borrowErr, attemptErr)
	}

	for moreResults {
		if err := s.streamBlocksRetrier.Attempt(fetchFn); err != nil {
			return pageToken, err
		}
	}
	return nil, nil
}

func (s *session) streamBlocksFromPeers(
	nsMetadata namespace.Metadata,
	shard uint32,
	peers []peer,
	metadataCh <-chan receivedBlockMetadata,
	opts result.Options,
	result blocksResult,
	progress *streamFromPeersMetrics,
	streamMetadataFn streamBlocksMetadataFn,
) {
	var (
		enqueueCh           = newEnqueueChannel(progress)
		peerBlocksBatchSize = s.streamBlocksBatchSize
	)

	// Consume the incoming metadata and enqueue to the ready channel
	// Spin up background goroutine to consume
	go func() {
		streamMetadataFn(len(peers), metadataCh, enqueueCh)
		// Begin assessing the queue and how much is processed, once queue
		// is entirely processed then we can close the enqueue channel
		enqueueCh.closeOnAllProcessed()
	}()

	// Fetch blocks from peers as results become ready
	peerQueues := make(peerBlocksQueues, 0, len(peers))
	for _, peer := range peers {
		peer := peer
		size := peerBlocksBatchSize
		workers := s.streamBlocksWorkers
		drainEvery := 100 * time.Millisecond
		queue := s.newPeerBlocksQueueFn(peer, size, drainEvery, workers,
			func(batch []receivedBlockMetadata) {
				s.streamBlocksBatchFromPeer(nsMetadata, shard, peer, batch, opts,
					result, enqueueCh, s.streamBlocksRetrier, progress)
			})
		peerQueues = append(peerQueues, queue)
	}

	var (
		selected             []receivedBlockMetadata
		pooled               selectPeersFromPerPeerBlockMetadatasPooledResources
		onQueueItemProcessed = func() {
			enqueueCh.trackProcessed(1)
		}
	)
	for perPeerBlocksMetadata := range enqueueCh.get() {
		// Filter and select which blocks to retrieve from which peers
		selected, pooled = s.selectPeersFromPerPeerBlockMetadatas(
			perPeerBlocksMetadata, peerQueues, enqueueCh,
			pooled, progress)

		if len(selected) == 0 {
			onQueueItemProcessed()
			continue
		}

		if len(selected) == 1 {
			queue := peerQueues.findQueue(selected[0].peer)
			queue.enqueue(selected[0], onQueueItemProcessed)
			continue
		}

		// Need to fan out, only track this as processed once all peer
		// queues have completed their fetches
		peerQueueEnqueues := uint32(len(selected))
		completed := uint32(0)
		onDone := func() {
			// Mark completion of work from the enqueue channel when all queues drained
			if atomic.AddUint32(&completed, 1) != peerQueueEnqueues {
				return
			}
			enqueueCh.trackProcessed(1)
		}

		for _, receivedBlockMetadata := range selected {
			queue := peerQueues.findQueue(receivedBlockMetadata.peer)
			queue.enqueue(receivedBlockMetadata, onDone)
		}
	}

	// Close all queues
	peerQueues.closeAll()
}

type streamBlocksMetadataFn func(
	peersLen int,
	ch <-chan receivedBlockMetadata,
	enqueueCh enqueueChannel,
)

func (s *session) passThroughBlocksMetadata(
	peersLen int,
	ch <-chan receivedBlockMetadata,
	enqueueCh enqueueChannel,
) {
	// Receive off of metadata channel
	for {
		m, ok := <-ch
		if !ok {
			break
		}
		res := []receivedBlockMetadata{m}
		enqueueCh.enqueue(res)
	}
}

func (s *session) streamAndGroupCollectedBlocksMetadata(
	peersLen int,
	metadataCh <-chan receivedBlockMetadata,
	enqueueCh enqueueChannel,
) {
	metadata := make(map[hashAndBlockStart]receivedBlocks)

	for {
		m, ok := <-metadataCh
		if !ok {
			break
		}

		key := hashAndBlockStart{
			hash:       m.id.Hash(),
			blockStart: m.block.start.UnixNano(),
		}
		received, ok := metadata[key]
		if !ok {
			received = receivedBlocks{
				results: make([]receivedBlockMetadata, 0, peersLen),
			}
		}

		// The entry has already been enqueued which means the metadata we just
		// received is a duplicate. Discard it and move on.
		if received.enqueued {
			s.emitDuplicateMetadataLogV2(received, m)
			continue
		}

		// Determine if the incoming metadata is a duplicate by checking if we've
		// already received metadata from this peer.
		existingIndex := -1
		for i, existingMetadata := range received.results {
			if existingMetadata.peer.Host().ID() == m.peer.Host().ID() {
				existingIndex = i
				break
			}
		}

		if existingIndex != -1 {
			// If it is a duplicate, then overwrite it (always keep the most recent
			// duplicate)
			received.results[existingIndex] = m
		} else {
			// Otherwise it's not a duplicate, so its safe to append.
			received.results = append(received.results, m)
		}

		// Since we always perform an overwrite instead of an append for duplicates
		// from the same peer, once len(received.results == peersLen) then we know
		// that we've received at least one metadata from every peer and its safe
		// to enqueue the entry.
		if len(received.results) == peersLen {
			enqueueCh.enqueue(received.results)
			received.enqueued = true
		}

		// Ensure tracking enqueued by setting modified result back to map
		metadata[key] = received
	}

	// Enqueue all unenqueued received metadata. Note that these entries will have
	// metadata from only a subset of their peers.
	for _, received := range metadata {
		if received.enqueued {
			continue
		}
		enqueueCh.enqueue(received.results)
	}
}

// TODO(rartoul): Delete this when we delete the V1 code path
func (s *session) emitDuplicateMetadataLog(
	received receivedBlocks,
	metadata receivedBlockMetadata,
) {
	fields := make([]xlog.Field, 0, len(received.results)+1)
	fields = append(fields, xlog.NewField(
		"incomingMetadata",
		fmt.Sprintf("ID: %s, peer: %s", metadata.id.String(), metadata.peer.Host().String()),
	))
	for i, result := range received.results {
		fields = append(fields, xlog.NewField(
			fmt.Sprintf("existingMetadata_%d", i),
			fmt.Sprintf("ID: %s, peer: %s", result.id.String(), result.peer.Host().String()),
		))
	}
	s.log.WithFields(fields...).Warnf(
		"received metadata, but peer metadata has already been submitted")
}

// emitDuplicateMetadataLogV2 emits a log with the details of the duplicate metadata
// event. Note that we're unable to log the blocks themselves because they're contained
// in a slice that is not safe for concurrent access (I.E logging them here would be
// racey because other code could be modifying the slice)
func (s *session) emitDuplicateMetadataLogV2(
	received receivedBlocks,
	metadata receivedBlockMetadata,
) {
	fields := make([]xlog.Field, 0, len(received.results)+1)
	fields = append(fields, xlog.NewField(
		"incomingMetadata",
		fmt.Sprintf(
			"ID: %s, peer: %s",
			metadata.id.String(),
			metadata.peer.Host().String(),
		),
	))
	for i, result := range received.results {
		fields = append(fields, xlog.NewField(
			fmt.Sprintf("existingMetadata_%d", i),
			fmt.Sprintf(
				"ID: %s, peer: %s",
				result.id.String(),
				result.peer.Host().String(),
			),
		))
	}
	// Debug-level because this is a common enough occurrence that logging it by
	// default would be noisy
	s.log.WithFields(fields...).Debugf(
		"received metadata, but peer metadata has already been submitted")
}

type pickBestPeerFn func(
	perPeerBlockMetadata []receivedBlockMetadata,
	peerQueues peerBlocksQueues,
	resources pickBestPeerPooledResources,
) (index int, pooled pickBestPeerPooledResources)

type pickBestPeerPooledResources struct {
	ranking []receivedBlockMetadataQueue
}

func (s *session) streamBlocksPickBestPeer(
	perPeerBlockMetadata []receivedBlockMetadata,
	peerQueues peerBlocksQueues,
	pooled pickBestPeerPooledResources,
) (int, pickBestPeerPooledResources) {
	// Order by least attempts then by least outstanding blocks being fetched
	pooled.ranking = pooled.ranking[:0]
	for i := range perPeerBlockMetadata {
		elem := receivedBlockMetadataQueue{
			blockMetadata: perPeerBlockMetadata[i],
			queue:         peerQueues.findQueue(perPeerBlockMetadata[i].peer),
		}
		pooled.ranking = append(pooled.ranking, elem)
	}
	elems := receivedBlockMetadataQueuesByAttemptsAscOutstandingAsc(pooled.ranking)
	sort.Stable(elems)

	// Return index of the best peer
	var (
		bestPeer = pooled.ranking[0].queue.peer
		idx      int
	)
	for i := range perPeerBlockMetadata {
		if bestPeer == perPeerBlockMetadata[i].peer {
			idx = i
			break
		}
	}
	return idx, pooled
}

type selectPeersFromPerPeerBlockMetadatasPooledResources struct {
	currEligible                []receivedBlockMetadata
	result                      []receivedBlockMetadata
	pickBestPeerPooledResources pickBestPeerPooledResources
}

func (s *session) selectPeersFromPerPeerBlockMetadatas(
	perPeerBlocksMetadata []receivedBlockMetadata,
	peerQueues peerBlocksQueues,
	reEnqueueCh enqueueChannel,
	pooled selectPeersFromPerPeerBlockMetadatasPooledResources,
	m *streamFromPeersMetrics,
) ([]receivedBlockMetadata, selectPeersFromPerPeerBlockMetadatasPooledResources) {
	// Get references to pooled array
	pooled.currEligible = pooled.currEligible[:0]
	for _, metadata := range perPeerBlocksMetadata {
		pooled.currEligible = append(pooled.currEligible, metadata)
	}

	currEligible := pooled.currEligible[:]

	// Sort the per peer metadatas by peer ID for consistent results
	sort.Sort(peerBlockMetadataByID(currEligible))


	// Only select from peers not already attempted
	currID := currEligible[0].id
	currBlock := currEligible[0].block
	for i := len(currEligible) - 1; i >= 0; i-- {
		if currEligible[i].block.reattempt.attempt == 0 {
			// Not attempted yet
			continue
		}

		// Check if eligible
		n := s.streamBlocksMaxBlockRetries
		if currEligible[i].block.reattempt.peerAttempts(currEligible[i].peer) >= n {
			// Swap current entry to tail
			receivedBlockMetadatas(currEligible).swap(i, len(currEligible)-1)
			// Trim newly last entry
			currEligible = currEligible[:len(currEligible)-1]
			continue
		}
	}

	if len(currEligible) == 0 {
		// No current eligible peers to select from
		if currBlock.reattempt.failAllowed != nil && atomic.AddInt32(currBlock.reattempt.failAllowed, -1) > 0 {
			// Some peers may still return results so we don't report
			// error here, just skip considering this block
			return nil, pooled
		}

		// Retry all of them again
		errMsg := "all retries failed for streaming blocks from peers"
		// err := fmt.Errorf(errMsg+": attempts=%d", currBlock.reattempt.attempt)
		// reattemptReason := retriesExhaustedErrReason
		// reattemptType := fullRetryReattemptType
		// retryBlock := currBlock
		// numPeers := int32(len(retryBlock.reattempt.fetchedPeersMetadata))
		// retryBlock.reattempt.failAllowed = &numPeers
		// s.reattemptStreamBlocksFromPeersFn([]blockMetadata{retryBlock}, reEnqueueCh,
		// 	err, reattemptReason, reattemptType, m)

		m.fetchBlockFinalError.Inc(1)
		s.log.WithFields(
			xlog.NewField("id", currID.String()),
			xlog.NewField("start", currBlock.start.String()),
			xlog.NewField("attempted", currBlock.reattempt.attempt),
			xlog.NewField("attemptErrs", xerrors.Errors(currBlock.reattempt.errs).Error()),
		).Error(errMsg)

		return nil, pooled
	}

	var (
		singlePeer         = len(currEligible) == 1
		sameNonNilChecksum = true
		curChecksum        *uint32
	)
	for i := range currEligible {
		// If any peer has a nil checksum, this might be the most recent block
		// and therefore not sealed so we want to merge from all peers
		if currEligible[i].block.checksum == nil {
			sameNonNilChecksum = false
			break
		}
		if curChecksum == nil {
			curChecksum = currEligible[i].block.checksum
		} else if *curChecksum != *currEligible[i].block.checksum {
			sameNonNilChecksum = false
			break
		}
	}

	// Prepare the reattempt peers metadata so we can retry from any of the peers on failure
	peersMetadata := make([]blockMetadataReattemptPeerMetadata, 0, len(currEligible))
	for i := range currEligible {
		metadata := blockMetadataReattemptPeerMetadata{
			peer:     currEligible[i].peer,
			start:    currEligible[i].block.start,
			size:     currEligible[i].block.size,
			checksum: currEligible[i].block.checksum,
		}
		peersMetadata = append(peersMetadata, metadata)
	}

	// If all the peers have the same non-nil checksum, we pick the peer with the
	// fewest attempts and fewest outstanding requests
	if singlePeer || sameNonNilChecksum {
		var idx int
		if singlePeer {
			idx = 0
		} else {
			pooledResources := pooled.pickBestPeerPooledResources
			idx, pooledResources = s.pickBestPeerFn(currEligible, peerQueues,
				pooledResources)
			pooled.pickBestPeerPooledResources = pooledResources
		}

		// Set the reattempt metadata
		selected := currEligible[idx]
		selected.block.reattempt.id = currID
		selected.block.reattempt.attempt++
		selected.block.reattempt.attempted =
			append(selected.block.reattempt.attempted, selected.peer)
		selected.block.reattempt.failAllowed = nil
		selected.block.reattempt.retryPeersMetadata = peersMetadata
		selected.block.reattempt.fetchedPeersMetadata = peersMetadata

		// Return just the single peer we selected
		currEligible = currEligible[:1]
		currEligible[0] = selected
	} else {
		numPeers := int32(len(currEligible))
		for i := range currEligible {
			// Set the reattempt metadata
			// NB(xichen): each block will only be retried on the same peer because we
			// already fan out the request to all peers. This means we merge data on
			// a best-effort basis and only fail if we failed to read data from all peers.
			peer := currEligible[i].peer
			block := currEligible[i].block
			currEligible[i].block.reattempt.id = currID
			currEligible[i].block.reattempt.attempt++
			currEligible[i].block.reattempt.attempted =
				append(currEligible[i].block.reattempt.attempted, peer)
			currEligible[i].block.reattempt.failAllowed = &numPeers
			currEligible[i].block.reattempt.retryPeersMetadata = []blockMetadataReattemptPeerMetadata{
				{
					peer:     peer,
					start:    block.start,
					size:     block.size,
					checksum: block.checksum,
				},
			}
			currEligible[i].block.reattempt.fetchedPeersMetadata = peersMetadata
		}
	}

	return currEligible, pooled
}

func (s *session) streamBlocksBatchFromPeer(
	namespaceMetadata namespace.Metadata,
	shard uint32,
	peer peer,
	batch []receivedBlockMetadata,
	opts result.Options,
	blocksResult blocksResult,
	enqueueCh enqueueChannel,
	retrier xretry.Retrier,
	m *streamFromPeersMetrics,
) {
	// Prepare request
	var (
		req          = rpc.NewFetchBlocksRawRequest()
		result       *rpc.FetchBlocksRawResult_
		reqBlocksLen uint

		nowFn              = opts.ClockOptions().NowFn()
		ropts              = namespaceMetadata.Options().RetentionOptions()
		blockSize          = ropts.BlockSize()
		retention          = ropts.RetentionPeriod()
		earliestBlockStart = nowFn().Add(-retention).Truncate(blockSize)
	)
	req.NameSpace = namespaceMetadata.ID().Data().Get()
	req.Shard = int32(shard)
	req.Elements = make([]*rpc.FetchBlocksRawRequestElement, 0, len(batch))
	for i := range batch {
		blockStart := batch[i].block.start
		if blockStart.Before(earliestBlockStart) {
			continue // Fell out of retention while we were streaming blocks
		}
		req.Elements = append(req.Elements, &rpc.FetchBlocksRawRequestElement{
			ID:     batch[i].id.Data().Get(),
			Starts: []int64{blockStart.UnixNano()},
		})
		reqBlocksLen++
	}
	if reqBlocksLen == 0 {
		// All blocks fell out of retention while streaming
		return
	}

	// Attempt request
	if err := retrier.Attempt(func() error {
		var attemptErr error
		borrowErr := peer.BorrowConnection(func(client rpc.TChanNode) {
			tctx, _ := thrift.NewContext(s.streamBlocksBatchTimeout)
			result, attemptErr = client.FetchBlocksRaw(tctx, req)
		})
		err := xerrors.FirstError(borrowErr, attemptErr)
		// Do not retry if cannot borrow the connection or
		// if the connection pool has no connections
		switch err {
		case errSessionHasNoHostQueueForHost,
			errConnectionPoolHasNoConnections:
			err = xerrors.NewNonRetryableError(err)
		}
		return err
	}); err != nil {
		blocksErr := fmt.Errorf(
			"stream blocks request error: error=%s, peer=%s",
			err.Error(), peer.Host().String(),
		)
		s.reattemptStreamBlocksFromPeersFn(batch, enqueueCh, blocksErr,
			reqErrReason, nextRetryReattemptType, m)
		m.fetchBlockError.Inc(int64(reqBlocksLen))
		s.log.Debugf(blocksErr.Error())
		return
	}

	// Parse and act on result
	tooManyIDsLogged := false
	for i := range result.Elements {
		if i >= len(batch) {
			m.fetchBlockError.Inc(int64(len(req.Elements[i].Starts)))
			m.fetchBlockFinalError.Inc(int64(len(req.Elements[i].Starts)))
			if !tooManyIDsLogged {
				tooManyIDsLogged = true
				s.log.WithFields(
					xlog.NewField("peer", peer.Host().String()),
				).Errorf("stream blocks more IDs than expected")
			}
			continue
		}

		id := batch[i].id
		if !bytes.Equal(id.Data().Get(), result.Elements[i].ID) {
			blocksErr := fmt.Errorf(
				"stream blocks mismatched ID: expectedID=%s, actualID=%s, indexID=%d, peer=%s",
				batch[i].id.String(), id.String(), i, peer.Host().String(),
			)
			b := []receivedBlockMetadata{batch[i]}
			s.reattemptStreamBlocksFromPeersFn(b, enqueueCh, blocksErr,
				respErrReason, nextRetryReattemptType, m)
			m.fetchBlockError.Inc(int64(len(req.Elements[i].Starts)))
			s.log.Debugf(blocksErr.Error())
			continue
		}

		if len(result.Elements[i].Blocks) == 0 {
			// If fell out of retention during request this is healthy, otherwise
			// missing blocks will be repaired during an active repair
			continue
		}

		tooManyBlocksLogged := false
		for j := range result.Elements[i].Blocks {
			if j >= len(req.Elements[i].Starts) {
				m.fetchBlockError.Inc(int64(len(req.Elements[i].Starts)))
				m.fetchBlockFinalError.Inc(int64(len(req.Elements[i].Starts)))
				if !tooManyBlocksLogged {
					tooManyBlocksLogged = true
					s.log.WithFields(
						xlog.NewField("id", id.String()),
						xlog.NewField("expectedStarts", newTimesByUnixNanos(req.Elements[i].Starts)),
						xlog.NewField("actualStarts", newTimesByRPCBlocks(result.Elements[i].Blocks)),
						xlog.NewField("peer", peer.Host().String()),
					).Errorf("stream blocks returned more blocks than expected")
				}
				continue
			}

			block := result.Elements[i].Blocks[j]

			// Verify and if verify succeeds add the block from the peer
			err := s.verifyFetchedBlock(block)
			if err == nil {
				err = blocksResult.addBlockFromPeer(id, peer.Host(), block)
			}

			if err != nil {
				failed := []receivedBlockMetadata{batch[i]}
				blocksErr := fmt.Errorf(
					"stream blocks bad block: id=%s, start=%d, error=%s, indexID=%d, indexBlock=%d, peer=%s",
					id.String(), block.Start, err.Error(), i, j, peer.Host().String(),
				)
				s.reattemptStreamBlocksFromPeersFn(failed, enqueueCh, blocksErr,
					respErrReason, nextRetryReattemptType, m)
				m.fetchBlockError.Inc(1)
				s.log.Debugf(blocksErr.Error())
				continue
			}

			m.fetchBlockSuccess.Inc(1)
		}
	}
}

func (s *session) verifyFetchedBlock(block *rpc.Block) error {
	if block.Err != nil {
		return fmt.Errorf("block error from peer: %s %s", block.Err.Type.String(), block.Err.Message)
	}
	if block.Segments == nil {
		return fmt.Errorf("block segments is bad: segments is nil")
	}
	if block.Segments.Merged == nil && len(block.Segments.Unmerged) == 0 {
		return fmt.Errorf("block segments is bad: merged and unmerged not set")
	}

	if checksum := block.Checksum; checksum != nil {
		var (
			d        = digest.NewDigest()
			expected = uint32(*checksum)
		)
		if merged := block.Segments.Merged; merged != nil {
			d = d.Update(merged.Head).Update(merged.Tail)
		} else {
			for _, s := range block.Segments.Unmerged {
				d = d.Update(s.Head).Update(s.Tail)
			}
		}
		if actual := d.Sum32(); actual != expected {
			return fmt.Errorf("block checksum is bad: expected=%d, actual=%d", expected, actual)
		}
	}

	return nil
}

type reason int

const (
	reqErrReason reason = iota
	respErrReason
	retriesExhaustedErrReason
)

type reattemptType int

const (
	nextRetryReattemptType reattemptType = iota
	fullRetryReattemptType
)

type reattemptStreamBlocksFromPeersFn func(
	[]receivedBlockMetadata,
	enqueueChannel,
	error,
	reason,
	reattemptType,
	*streamFromPeersMetrics,
)

func (s *session) streamBlocksReattemptFromPeers(
	blocks []receivedBlockMetadata,
	enqueueCh enqueueChannel,
	attemptErr error,
	reason reason,
	reattemptType reattemptType,
	m *streamFromPeersMetrics,
) {
	switch reason {
	case reqErrReason:
		m.fetchBlockRetriesReqError.Inc(int64(len(blocks)))
	case respErrReason:
		m.fetchBlockRetriesRespError.Inc(int64(len(blocks)))
	}

	// Must do this asynchronously or else could get into a deadlock scenario
	// where cannot enqueue into the reattempt channel because no more work is
	// getting done because new attempts are blocked on existing attempts completing
	// and existing attempts are trying to enqueue into a full reattempt channel
	enqueue := enqueueCh.enqueueDelayed(len(blocks))
	go s.streamBlocksReattemptFromPeersEnqueue(blocks, attemptErr, reattemptType, enqueue)
}

func (s *session) streamBlocksReattemptFromPeersEnqueue(
	blocks []receivedBlockMetadata,
	attemptErr error,
	reattemptType reattemptType,
	enqueueFn func([]receivedBlockMetadata),
) {
	for i := range blocks {
		var reattemptPeersMetadata []blockMetadataReattemptPeerMetadata
		switch reattemptType {
		case nextRetryReattemptType:
			reattemptPeersMetadata = blocks[i].block.reattempt.retryPeersMetadata
		case fullRetryReattemptType:
			reattemptPeersMetadata = blocks[i].block.reattempt.fetchedPeersMetadata
		}
		if len(reattemptPeersMetadata) == 0 {
			continue
		}

		// Reconstruct peers metadata for reattempt
		reattemptBlocksMetadata := make([]receivedBlockMetadata, len(reattemptPeersMetadata))
		for j := range reattemptPeersMetadata {
			reattempt := blocks[i].block.reattempt

			// Copy the errors for every peer so they don't shard the same error
			// slice and therefore are not subject to race conditions when the
			// error slice is modified
			reattemptErrs := make([]error, len(reattempt.errs)+1)
			n := copy(reattemptErrs, reattempt.errs)
			reattemptErrs[n] = attemptErr
			reattempt.errs = reattemptErrs

			reattemptBlocksMetadata[j] = receivedBlockMetadata{
				peer: reattemptPeersMetadata[j].peer,
				id:   reattempt.id,
				block: blockMetadata{
					start:     reattemptPeersMetadata[j].start,
					size:      reattemptPeersMetadata[j].size,
					checksum:  reattemptPeersMetadata[j].checksum,
					reattempt: reattempt,
				},
			}
		}
		// Re-enqueue the block to be fetched
		enqueueFn(reattemptBlocksMetadata)
	}
}

type blocksResult interface {
	addBlockFromPeer(id ident.ID, peer topology.Host, block *rpc.Block) error
}

type baseBlocksResult struct {
	blockOpts               block.Options
	blockAllocSize          int
	contextPool             context.Pool
	encoderPool             encoding.EncoderPool
	multiReaderIteratorPool encoding.MultiReaderIteratorPool
}

func newBaseBlocksResult(
	opts Options,
	resultOpts result.Options,
) baseBlocksResult {
	blockOpts := resultOpts.DatabaseBlockOptions()
	return baseBlocksResult{
		blockOpts:               blockOpts,
		blockAllocSize:          blockOpts.DatabaseBlockAllocSize(),
		contextPool:             opts.ContextPool(),
		encoderPool:             blockOpts.EncoderPool(),
		multiReaderIteratorPool: blockOpts.MultiReaderIteratorPool(),
	}
}

func (b *baseBlocksResult) segmentForBlock(seg *rpc.Segment) ts.Segment {
	var (
		bytesPool  = b.blockOpts.BytesPool()
		head, tail checked.Bytes
	)
	if len(seg.Head) > 0 {
		head = bytesPool.Get(len(seg.Head))
		head.IncRef()
		head.AppendAll(seg.Head)
		head.DecRef()
	}
	if len(seg.Tail) > 0 {
		tail = bytesPool.Get(len(seg.Tail))
		tail.IncRef()
		tail.AppendAll(seg.Tail)
		tail.DecRef()
	}
	return ts.NewSegment(head, tail, ts.FinalizeHead&ts.FinalizeTail)
}

func (b *baseBlocksResult) mergeReaders(start time.Time, readers []io.Reader) (encoding.Encoder, error) {
	iter := b.multiReaderIteratorPool.Get()
	iter.Reset(readers)
	defer iter.Close()

	encoder := b.encoderPool.Get()
	encoder.Reset(start, b.blockAllocSize)

	for iter.Next() {
		dp, unit, annotation := iter.Current()
		if err := encoder.Encode(dp, unit, annotation); err != nil {
			encoder.Close()
			return nil, err
		}
	}
	if err := iter.Err(); err != nil {
		encoder.Close()
		return nil, err
	}

	return encoder, nil
}

func (b *baseBlocksResult) newDatabaseBlock(block *rpc.Block) (block.DatabaseBlock, error) {
	var (
		start    = time.Unix(0, block.Start)
		segments = block.Segments
		result   = b.blockOpts.DatabaseBlockPool().Get()
	)

	if segments == nil {
		result.Close() // return block to pool
		return nil, errSessionBadBlockResultFromPeer
	}

	switch {
	case segments.Merged != nil:
		// Unmerged, can insert directly into a single block
		result.Reset(start, b.segmentForBlock(segments.Merged))

	case segments.Unmerged != nil:
		// Must merge to provide a single block
		segmentReaderPool := b.blockOpts.SegmentReaderPool()
		readers := make([]io.Reader, len(segments.Unmerged))
		for i := range segments.Unmerged {
			segmentReader := segmentReaderPool.Get()
			segmentReader.Reset(b.segmentForBlock(segments.Unmerged[i]))
			readers[i] = segmentReader
		}

		encoder, err := b.mergeReaders(start, readers)
		for _, reader := range readers {
			// Close each reader
			segmentReader := reader.(xio.SegmentReader)
			segmentReader.Finalize()
		}

		if err != nil {
			// mergeReaders(...) already calls encoder.Close() upon error
			result.Close() // return block to pool
			return nil, err
		}

		// Set the block data
		result.Reset(start, encoder.Discard())

	default:
		result.Close() // return block to pool
		return nil, errSessionBadBlockResultFromPeer
	}

	return result, nil
}

type streamBlocksResult struct {
	baseBlocksResult
	outputCh chan<- peerBlocksDatapoint
}

func newStreamBlocksResult(
	opts Options,
	resultOpts result.Options,
	outputCh chan<- peerBlocksDatapoint,
) *streamBlocksResult {
	return &streamBlocksResult{
		baseBlocksResult: newBaseBlocksResult(opts, resultOpts),
		outputCh:         outputCh,
	}
}

type peerBlocksDatapoint struct {
	id    ident.ID
	peer  topology.Host
	block block.DatabaseBlock
}

func (s *streamBlocksResult) addBlockFromPeer(id ident.ID, peer topology.Host, block *rpc.Block) error {
	result, err := s.newDatabaseBlock(block)
	if err != nil {
		return err
	}
	s.outputCh <- peerBlocksDatapoint{
		id:    id,
		peer:  peer,
		block: result,
	}
	return nil
}

type peerBlocksIter struct {
	inputCh <-chan peerBlocksDatapoint
	errCh   <-chan error
	current peerBlocksDatapoint
	err     error
	done    bool
}

func newPeerBlocksIter(
	inputC <-chan peerBlocksDatapoint,
	errC <-chan error,
) *peerBlocksIter {
	return &peerBlocksIter{
		inputCh: inputC,
		errCh:   errC,
	}
}

func (it *peerBlocksIter) Current() (topology.Host, ident.ID, block.DatabaseBlock) {
	return it.current.peer, it.current.id, it.current.block
}

func (it *peerBlocksIter) Err() error {
	return it.err
}

func (it *peerBlocksIter) Next() bool {
	if it.done || it.err != nil {
		return false
	}
	m, more := <-it.inputCh

	if !more {
		it.err = <-it.errCh
		it.done = true
		return false
	}

	it.current = m
	return true
}

type bulkBlocksResult struct {
	sync.RWMutex
	baseBlocksResult
	result result.ShardResult
}

func newBulkBlocksResult(
	opts Options,
	resultOpts result.Options,
) *bulkBlocksResult {
	return &bulkBlocksResult{
		baseBlocksResult: newBaseBlocksResult(opts, resultOpts),
		result:           result.NewShardResult(4096, resultOpts),
	}
}

func (r *bulkBlocksResult) addBlockFromPeer(id ident.ID, peer topology.Host, block *rpc.Block) error {
	start := time.Unix(0, block.Start)
	result, err := r.newDatabaseBlock(block)
	if err != nil {
		return err
	}

	for {
		r.Lock()
		currBlock, exists := r.result.BlockAt(id, start)
		if !exists {
			r.result.AddBlock(id, result)
			r.Unlock()
			break
		}

		// Remove the existing block from the result so it doesn't get
		// merged again
		r.result.RemoveBlockAt(id, start)
		r.Unlock()

		// If we've already received data for this block, merge them
		// with the new block if possible
		tmpCtx := r.contextPool.Get()
		currReader, err := currBlock.Stream(tmpCtx)
		if err != nil {
			return err
		}

		// If there are no data in the current block, there is no
		// need to merge
		if currReader == nil {
			continue
		}

		resultReader, err := result.Stream(tmpCtx)
		if err != nil {
			return err
		}
		if resultReader == nil {
			return nil
		}

		readers := []io.Reader{currReader, resultReader}
		encoder, err := r.mergeReaders(start, readers)

		if err != nil {
			return err
		}

		result.Close()

		result = r.blockOpts.DatabaseBlockPool().Get()
		result.Reset(start, encoder.Discard())

		tmpCtx.Close()
	}

	return nil
}

type enqueueCh struct {
	enqueued         uint64
	processed        uint64
	peersMetadataCh  chan []receivedBlockMetadata
	closed           int64
	enqueueDelayedFn func(peersMetadata []receivedBlockMetadata)
	metrics          *streamFromPeersMetrics
}

const enqueueChannelDefaultLen = 32768

func newEnqueueChannel(m *streamFromPeersMetrics) enqueueChannel {
	c := &enqueueCh{
		peersMetadataCh: make(chan []receivedBlockMetadata, enqueueChannelDefaultLen),
		closed:          0,
		metrics:         m,
	}
	// Allocate the enqueue delayed fn just once
	c.enqueueDelayedFn = func(peersMetadata []receivedBlockMetadata) {
		c.peersMetadataCh <- peersMetadata
	}
	go func() {
		for atomic.LoadInt64(&c.closed) == 0 {
			m.blocksEnqueueChannel.Update(float64(len(c.peersMetadataCh)))
			time.Sleep(gaugeReportInterval)
		}
	}()
	return c
}

func (c *enqueueCh) enqueue(peersMetadata []receivedBlockMetadata) {
	atomic.AddUint64(&c.enqueued, 1)
	c.peersMetadataCh <- peersMetadata
}

func (c *enqueueCh) enqueueDelayed(numToEnqueue int) func([]receivedBlockMetadata) {
	atomic.AddUint64(&c.enqueued, uint64(numToEnqueue))
	return c.enqueueDelayedFn
}

func (c *enqueueCh) get() <-chan []receivedBlockMetadata {
	return c.peersMetadataCh
}

func (c *enqueueCh) trackProcessed(amount int) {
	atomic.AddUint64(&c.processed, uint64(amount))
}

func (c *enqueueCh) unprocessedLen() int {
	return int(atomic.LoadUint64(&c.enqueued) - atomic.LoadUint64(&c.processed))
}

func (c *enqueueCh) closeOnAllProcessed() {
	defer func() {
		atomic.StoreInt64(&c.closed, 1)
	}()
	for {
		if c.unprocessedLen() == 0 {
			// Will only ever be zero after all is processed if called
			// after enqueueing the desired set of entries as long as
			// the guarantee that reattempts are enqueued before the
			// failed attempt is marked as processed is upheld
			close(c.peersMetadataCh)
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

type receivedBlocks struct {
	enqueued bool
	results  []receivedBlockMetadata
}

type processFn func(batch []receivedBlockMetadata)

// peerBlocksQueue is a per peer queue of blocks to be retrieved from a peer
type peerBlocksQueue struct {
	sync.RWMutex
	closed       bool
	peer         peer
	queue        []receivedBlockMetadata
	doneFns      []func()
	assigned     uint64
	completed    uint64
	maxQueueSize int
	workers      xsync.WorkerPool
	processFn    processFn
}

type newPeerBlocksQueueFn func(
	peer peer,
	maxQueueSize int,
	interval time.Duration,
	workers xsync.WorkerPool,
	processFn processFn,
) *peerBlocksQueue

func newPeerBlocksQueue(
	peer peer,
	maxQueueSize int,
	interval time.Duration,
	workers xsync.WorkerPool,
	processFn processFn,
) *peerBlocksQueue {
	q := &peerBlocksQueue{
		peer:         peer,
		maxQueueSize: maxQueueSize,
		workers:      workers,
		processFn:    processFn,
	}
	if interval > 0 {
		go q.drainEvery(interval)
	}
	return q
}

func (q *peerBlocksQueue) drainEvery(interval time.Duration) {
	for {
		q.Lock()
		if q.closed {
			q.Unlock()
			return
		}
		q.drainWithLock()
		q.Unlock()
		time.Sleep(interval)
	}
}

func (q *peerBlocksQueue) close() {
	q.Lock()
	defer q.Unlock()
	q.closed = true
}

func (q *peerBlocksQueue) trackAssigned(amount int) {
	atomic.AddUint64(&q.assigned, uint64(amount))
}

func (q *peerBlocksQueue) trackCompleted(amount int) {
	atomic.AddUint64(&q.completed, uint64(amount))
}

func (q *peerBlocksQueue) enqueue(bl receivedBlockMetadata, doneFn func()) {
	q.Lock()

	if len(q.queue) == 0 && cap(q.queue) < q.maxQueueSize {
		// Lazy initialize queue
		q.queue = make([]receivedBlockMetadata, 0, q.maxQueueSize)
	}
	if len(q.doneFns) == 0 && cap(q.doneFns) < q.maxQueueSize {
		// Lazy initialize doneFns
		q.doneFns = make([]func(), 0, q.maxQueueSize)
	}
	q.queue = append(q.queue, bl)
	if doneFn != nil {
		q.doneFns = append(q.doneFns, doneFn)
	}
	q.trackAssigned(1)

	// Determine if should drain immediately
	if len(q.queue) < q.maxQueueSize {
		// Require more to fill up block
		q.Unlock()
		return
	}
	q.drainWithLock()

	q.Unlock()
}

func (q *peerBlocksQueue) drain() {
	q.Lock()
	q.drainWithLock()
	q.Unlock()
}

func (q *peerBlocksQueue) drainWithLock() {
	if len(q.queue) == 0 {
		// None to drain
		return
	}
	enqueued := q.queue
	doneFns := q.doneFns
	q.queue = nil
	q.doneFns = nil
	q.workers.Go(func() {
		q.processFn(enqueued)
		// Call done callbacks
		for i := range doneFns {
			doneFns[i]()
		}
		// Track completed blocks
		q.trackCompleted(len(enqueued))
	})
}

type peerBlocksQueues []*peerBlocksQueue

func (qs peerBlocksQueues) findQueue(peer peer) *peerBlocksQueue {
	for _, q := range qs {
		if q.peer == peer {
			return q
		}
	}
	return nil
}

func (qs peerBlocksQueues) closeAll() {
	for _, q := range qs {
		q.close()
	}
}

type receivedBlockMetadata struct {
	peer  peer
	id    ident.ID
	block blockMetadata
}

type receivedBlockMetadatas []receivedBlockMetadata

func (arr receivedBlockMetadatas) swap(i, j int) { arr[i], arr[j] = arr[j], arr[i] }

type peerBlockMetadataByID []receivedBlockMetadata

func (arr peerBlockMetadataByID) Len() int      { return len(arr) }
func (arr peerBlockMetadataByID) Swap(i, j int) { arr[i], arr[j] = arr[j], arr[i] }
func (arr peerBlockMetadataByID) Less(i, j int) bool {
	return strings.Compare(arr[i].peer.Host().ID(), arr[j].peer.Host().ID()) < 0
}

type receivedBlockMetadataQueue struct {
	blockMetadata receivedBlockMetadata
	queue         *peerBlocksQueue
}

type receivedBlockMetadataQueuesByAttemptsAscOutstandingAsc []receivedBlockMetadataQueue

func (arr receivedBlockMetadataQueuesByAttemptsAscOutstandingAsc) Len() int {
	return len(arr)
}
func (arr receivedBlockMetadataQueuesByAttemptsAscOutstandingAsc) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr receivedBlockMetadataQueuesByAttemptsAscOutstandingAsc) Less(i, j int) bool {
	peerI := arr[i].queue.peer
	peerJ := arr[j].queue.peer
	attemptsI := arr[i].blockMetadata.block.reattempt.peerAttempts(peerI)
	attemptsJ := arr[j].blockMetadata.block.reattempt.peerAttempts(peerJ)
	if attemptsI != attemptsJ {
		return attemptsI < attemptsJ
	}

	outstandingI :=
		atomic.LoadUint64(&arr[i].queue.assigned) -
			atomic.LoadUint64(&arr[i].queue.completed)
	outstandingJ :=
		atomic.LoadUint64(&arr[j].queue.assigned) -
			atomic.LoadUint64(&arr[j].queue.completed)
	return outstandingI < outstandingJ
}

type blockMetadata struct {
	start     time.Time
	size      int64
	checksum  *uint32
	lastRead  time.Time
	reattempt blockMetadataReattempt
}

type blockMetadataReattempt struct {
	attempt              int
	failAllowed          *int32
	id                   ident.ID
	attempted            []peer
	errs                 []error
	retryPeersMetadata   []blockMetadataReattemptPeerMetadata
	fetchedPeersMetadata []blockMetadataReattemptPeerMetadata
}

type blockMetadataReattemptPeerMetadata struct {
	peer     peer
	start    time.Time
	size     int64
	checksum *uint32
}

func (b blockMetadataReattempt) peerAttempts(p peer) int {
	r := 0
	for i := range b.attempted {
		if b.attempted[i] == p {
			r++
		}
	}
	return r
}

func newTimesByUnixNanos(values []int64) []time.Time {
	result := make([]time.Time, len(values))
	for i := range values {
		result[i] = time.Unix(0, values[i])
	}
	return result
}

func newTimesByRPCBlocks(values []*rpc.Block) []time.Time {
	result := make([]time.Time, len(values))
	for i := range values {
		result[i] = time.Unix(0, values[i].Start)
	}
	return result
}

type metadataIter struct {
	inputCh  <-chan receivedBlockMetadata
	errCh    <-chan error
	host     topology.Host
	metadata block.Metadata
	done     bool
	err      error
}

func newMetadataIter(
	inputCh <-chan receivedBlockMetadata,
	errCh <-chan error,
) PeerBlockMetadataIter {
	return &metadataIter{
		inputCh: inputCh,
		errCh:   errCh,
	}
}

func (it *metadataIter) Next() bool {
	if it.done || it.err != nil {
		return false
	}
	m, more := <-it.inputCh
	if !more {
		it.err = <-it.errCh
		it.done = true
		return false
	}
	it.host = m.peer.Host()
	it.metadata = block.NewMetadata(m.id, m.block.start,
		m.block.size, m.block.checksum, m.block.lastRead)
	return true
}

func (it *metadataIter) Current() (topology.Host, block.Metadata) {
	return it.host, it.metadata
}

func (it *metadataIter) Err() error {
	return it.err
}

type hashAndBlockStart struct {
	hash       ident.Hash
	blockStart int64
}

// IsValidFetchBlocksMetadataEndpoint returns a bool indicating whether the
// specified endpointVersion is valid
func IsValidFetchBlocksMetadataEndpoint(endpointVersion FetchBlocksMetadataEndpointVersion) bool {
	for _, version := range validFetchBlocksMetadataEndpoints {
		if version == endpointVersion {
			return true
		}
	}
	return false
}

// UnmarshalYAML unmarshals an FetchBlocksMetadataEndpointVersion into a valid type from string.
func (v *FetchBlocksMetadataEndpointVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if str == "" {
		return errFetchBlocksMetadataEndpointVersionUnspecified
	}
	for _, valid := range validFetchBlocksMetadataEndpoints {
		if str == valid.String() {
			*v = valid
			return nil
		}
	}
	return fmt.Errorf("invalid FetchBlocksMetadataEndpointVersion '%s' valid types are: %v",
		str, validFetchBlocksMetadataEndpoints)
}

func (v FetchBlocksMetadataEndpointVersion) String() string {
	switch v {
	case FetchBlocksMetadataEndpointV1:
		return "v1"
	case FetchBlocksMetadataEndpointV2:
		return "v2"
	}
	return unknown
}
