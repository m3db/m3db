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

package integration

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/m3db/m3db/client"
	"github.com/m3db/m3db/encoding"
	"github.com/m3db/m3db/instrument"
	"github.com/m3db/m3db/persist/fs"
	"github.com/m3db/m3db/retention"
	"github.com/m3db/m3db/sharding"
	"github.com/m3db/m3db/storage"
	"github.com/m3db/m3db/storage/bootstrap"
	"github.com/m3db/m3db/storage/bootstrap/bootstrapper"
	bfs "github.com/m3db/m3db/storage/bootstrap/bootstrapper/fs"
	"github.com/m3db/m3db/storage/bootstrap/bootstrapper/peers"
	"github.com/m3db/m3db/topology"
	"github.com/m3db/m3x/log"
	"github.com/m3db/m3x/time"

	"github.com/stretchr/testify/require"
)

const (
	multiAddrPortStart = 9000
	multiAddrPortEach  = 4
)

type conditionFn func() bool

func waitUntil(fn conditionFn, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if fn() {
			return true
		}
		time.Sleep(time.Second)
	}
	return false
}

func newTestSleepFn() (func(d time.Duration), func(overrideFn func(d time.Duration))) {
	var mutex sync.RWMutex
	sleepFn := time.Sleep
	return func(d time.Duration) {
			mutex.RLock()
			defer mutex.RUnlock()
			sleepFn(d)
		}, func(fn func(d time.Duration)) {
			mutex.Lock()
			defer mutex.Unlock()
			sleepFn = fn
		}
}

func newMultiAddrTestOptions(opts testOptions, instance int) testOptions {
	bind := "0.0.0.0"
	start := multiAddrPortStart + (instance * multiAddrPortEach)
	return opts.
		SetTChannelNodeAddr(fmt.Sprintf("%s:%d", bind, start)).
		SetTChannelClusterAddr(fmt.Sprintf("%s:%d", bind, start+1)).
		SetHTTPNodeAddr(fmt.Sprintf("%s:%d", bind, start+2)).
		SetHTTPClusterAddr(fmt.Sprintf("%s:%d", bind, start+3))
}

func newMultiAddrAdminSession(
	t *testing.T,
	instrumentOpts instrument.Options,
	shardSet sharding.ShardSet,
	replicas int,
	instance int,
) client.AdminSession {
	var (
		start         = multiAddrPortStart
		hostShardSets []topology.HostShardSet
		origin        topology.Host
	)
	for i := 0; i < replicas; i++ {
		id := fmt.Sprintf("localhost%d", i)
		nodeAddr := fmt.Sprintf("127.0.0.1:%d", start+(i*multiAddrPortEach))
		host := topology.NewHost(id, nodeAddr)
		if i == instance {
			origin = host
		}
		hostShardSet := topology.NewHostShardSet(host, shardSet)
		hostShardSets = append(hostShardSets, hostShardSet)
	}

	staticOptions := topology.NewStaticOptions().
		SetShardSet(shardSet).
		SetReplicas(replicas).
		SetHostShardSets(hostShardSets)

	opts := client.NewAdminOptions().
		SetOrigin(origin).
		SetInstrumentOptions(instrumentOpts).
		SetTopologyInitializer(topology.NewStaticInitializer(staticOptions)).
		SetClusterConnectConsistencyLevel(client.ConnectConsistencyLevelAny).
		SetClusterConnectTimeout(time.Second)

	adminClient, err := client.NewAdminClient(opts.(client.AdminOptions))
	require.NoError(t, err)

	session, err := adminClient.NewAdminSession()
	require.NoError(t, err)

	return session
}

func newBootstrappableTestSetup(
	t *testing.T,
	opts testOptions,
	retentionOpts retention.Options,
	newBootstrapFn storage.NewBootstrapFn,
) *testSetup {
	setup, err := newTestSetup(opts)
	require.NoError(t, err)

	setup.storageOpts = setup.storageOpts.
		SetRetentionOptions(retentionOpts).
		SetNewBootstrapFn(newBootstrapFn)
	return setup
}

type bootstrappableTestSetupOptions struct {
	disablePeersBootstrapper bool
	sleepFn                  peers.SleepFn
}

type closeFn func()

func newDefaultBootstrappableTestSetups(
	t *testing.T,
	opts testOptions,
	retentionOpts retention.Options,
	setupOpts []bootstrappableTestSetupOptions,
) (testSetups, closeFn) {
	var (
		replicas        = len(setupOpts)
		setups          []*testSetup
		cleanupFns      []func()
		cleanupFnsMutex sync.RWMutex
	)
	appendCleanupFn := func(fn func()) {
		cleanupFnsMutex.Lock()
		defer cleanupFnsMutex.Unlock()
		cleanupFns = append(cleanupFns, fn)
	}
	for i := 0; i < replicas; i++ {
		var (
			instance              = i
			usingPeersBoostrapper = !setupOpts[i].disablePeersBootstrapper
			sleepFn               = setupOpts[i].sleepFn
			instanceOpts          = newMultiAddrTestOptions(opts, instance)
			setup                 *testSetup
		)
		setup = newBootstrappableTestSetup(t, instanceOpts, retentionOpts, func() bootstrap.Bootstrap {
			instrumentOpts := setup.storageOpts.InstrumentOptions()

			bsOpts := bootstrap.NewOptions().
				SetClockOptions(setup.storageOpts.ClockOptions()).
				SetRetentionOptions(setup.storageOpts.RetentionOptions()).
				SetInstrumentOptions(instrumentOpts)

			noOpAll := bootstrapper.NewNoOpAllBootstrapper()

			var session client.AdminSession
			if usingPeersBoostrapper {
				session = newMultiAddrAdminSession(t, instrumentOpts, setup.shardSet, replicas, instance)
				appendCleanupFn(func() {
					session.Close()
				})
			}

			peersOpts := peers.NewOptions().
				SetBootstrapOptions(bsOpts).
				SetAdminSession(session)
			if sleepFn != nil {
				peersOpts = peersOpts.SetSleepFn(sleepFn)
			}
			peersBootstrapper := peers.NewPeersBootstrapper(peersOpts, noOpAll)

			fsOpts := setup.storageOpts.CommitLogOptions().FilesystemOptions()
			filePathPrefix := fsOpts.FilePathPrefix()

			bfsOpts := bfs.NewOptions().
				SetBootstrapOptions(bsOpts).
				SetFilesystemOptions(fsOpts)

			var fsBootstrapper bootstrap.Bootstrapper
			if usingPeersBoostrapper {
				fsBootstrapper = bfs.NewFileSystemBootstrapper(filePathPrefix, bfsOpts, peersBootstrapper)
			} else {
				fsBootstrapper = bfs.NewFileSystemBootstrapper(filePathPrefix, bfsOpts, noOpAll)
			}

			return bootstrap.NewBootstrapProcess(bsOpts, fsBootstrapper)
		})
		logger := setup.storageOpts.InstrumentOptions().Logger()
		logger = logger.WithFields(xlog.NewLogField("instance", instance))
		iopts := setup.storageOpts.InstrumentOptions().SetLogger(logger)
		setup.storageOpts = setup.storageOpts.SetInstrumentOptions(iopts)

		setups = append(setups, setup)
		appendCleanupFn(func() {
			setup.close()
		})
	}

	return setups, func() {
		cleanupFnsMutex.RLock()
		defer cleanupFnsMutex.RUnlock()
		for _, fn := range cleanupFns {
			fn()
		}
	}
}

type testData struct {
	ids       []string
	numPoints int
	start     time.Time
}

func writeTestDataToDisk(
	namespace string,
	testSetup *testSetup,
	seriesMaps map[time.Time]seriesList,
) error {
	storageOpts := testSetup.storageOpts
	fsOpts := storageOpts.CommitLogOptions().FilesystemOptions()
	writerBufferSize := fsOpts.WriterBufferSize()
	blockSize := storageOpts.RetentionOptions().BlockSize()
	filePathPrefix := fsOpts.FilePathPrefix()
	newFileMode := fsOpts.NewFileMode()
	newDirectoryMode := fsOpts.NewDirectoryMode()
	writer := fs.NewWriter(blockSize, filePathPrefix, writerBufferSize, newFileMode, newDirectoryMode)
	encoder := storageOpts.EncoderPool().Get()

	for start, data := range seriesMaps {
		err := writeToDisk(writer, testSetup.shardSet, encoder, start, namespace, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeToDisk(
	writer fs.FileSetWriter,
	shardSet sharding.ShardSet,
	encoder encoding.Encoder,
	start time.Time,
	namespace string,
	seriesList seriesList,
) error {
	seriesPerShard := make(map[uint32][]series)
	for _, s := range seriesList {
		shard := shardSet.Shard(s.id)
		seriesPerShard[shard] = append(seriesPerShard[shard], s)
	}
	segmentHolder := make([][]byte, 2)
	for shard, seriesList := range seriesPerShard {
		if err := writer.Open(namespace, shard, start); err != nil {
			return err
		}
		for _, series := range seriesList {
			encoder.Reset(start, 0)
			for _, dp := range series.data {
				if err := encoder.Encode(dp, xtime.Second, nil); err != nil {
					return err
				}
			}
			segment := encoder.Stream().Segment()
			segmentHolder[0] = segment.Head
			segmentHolder[1] = segment.Tail
			if err := writer.WriteAll(series.id, segmentHolder); err != nil {
				return err
			}
		}
		if err := writer.Close(); err != nil {
			return err
		}
	}

	return nil
}
