// +build integration

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
	"os"
	"testing"
	"time"

	"github.com/m3db/m3db/bootstrap"
	"github.com/m3db/m3db/bootstrap/bootstrapper"
	bfs "github.com/m3db/m3db/bootstrap/bootstrapper/fs"
	"github.com/m3db/m3db/interfaces/m3db"
	"github.com/m3db/m3db/services/m3dbnode/server"
	xtime "github.com/m3db/m3db/x/time"

	"github.com/stretchr/testify/require"
)

func writeToDisk(
	writer m3db.FileSetWriter,
	shardingScheme m3db.ShardScheme,
	encoder m3db.Encoder,
	start time.Time,
	dm dataMap,
) error {
	idsPerShard := make(map[uint32][]string)
	for id := range dm {
		shard := shardingScheme.Shard(id)
		idsPerShard[shard] = append(idsPerShard[shard], id)
	}
	segmentHolder := make([][]byte, 2)
	for shard, ids := range idsPerShard {
		if err := writer.Open(shard, start); err != nil {
			return err
		}
		for _, id := range ids {
			encoder.Reset(start, 0)
			for _, dp := range dm[id] {
				if err := encoder.Encode(dp, xtime.Second, nil); err != nil {
					return err
				}
			}
			segment := encoder.Stream().Segment()
			segmentHolder[0] = segment.Head
			segmentHolder[1] = segment.Tail
			if err := writer.WriteAll(id, segmentHolder); err != nil {
				return err
			}
		}
		if err := writer.Close(); err != nil {
			return err
		}
	}
	return nil
}

func TestFilesystemBootstrap(t *testing.T) {
	// Test setup
	opts, now, err := setup()
	require.NoError(t, err)
	blockSize := opts.GetBlockSize()
	filePathPrefix := opts.GetFilePathPrefix()
	opts = opts.RetentionPeriod(2 * time.Hour).NewBootstrapFn(func() m3db.Bootstrap {
		noOpAll := bootstrapper.NewNoOpAllBootstrapper()
		bs := bfs.NewFileSystemBootstrapper(filePathPrefix, opts, noOpAll)
		bsOpts := bootstrap.NewOptions()
		return bootstrap.NewBootstrapProcess(bsOpts, opts, bs)
	})
	defer os.RemoveAll(filePathPrefix)

	writerFn := opts.GetNewFileSetWriterFn()
	writer := writerFn(blockSize, filePathPrefix)
	shardingScheme, err := server.ShardingScheme()
	require.NoError(t, err)
	encoder := opts.GetEncoderPool().Get()
	dataMaps := make(map[time.Time]dataMap)

	// Write test data
	inputData := []struct {
		metricNames []string
		numPoints   int
		start       time.Time
	}{
		{[]string{"foo", "bar"}, 100, (*now).Add(-blockSize)},
		{[]string{"foo", "baz"}, 50, *now},
	}
	for _, input := range inputData {
		testData := generateTestData(input.metricNames, input.numPoints, input.start)
		dataMaps[input.start] = testData
		require.NoError(t, writeToDisk(writer, shardingScheme, encoder, input.start, testData))
	}

	// Start the server with filesystem bootstrapper
	log := opts.GetLogger()
	log.Debug("filesystem bootstrap test")
	doneCh := make(chan struct{})
	require.NoError(t, startServer(opts, doneCh))
	log.Debug("server is now up")

	// Stop the server
	defer func() {
		require.NoError(t, stopServer(doneCh))
		log.Debug("server is now down")
	}()

	// Verify in-memory data match what we expect
	verifyDataMaps(t, blockSize, dataMaps)
}
