// +build integration
//
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
	"sync/atomic"
	"testing"
	"time"

	"github.com/m3db/m3cluster/services"
	"github.com/m3db/m3cluster/shard"
	"github.com/m3db/m3db/client"
	"github.com/m3db/m3db/storage/index"
	"github.com/m3db/m3db/topology"
	"github.com/m3db/m3ninx/idx"
	xclock "github.com/m3db/m3x/clock"
	"github.com/m3db/m3x/ident"
	xtime "github.com/m3db/m3x/time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndexMultipleNodeHighConcurrency(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}
	var (
		concurrency = 10
		writeEach   = 100
		numTags     = 10
	)

	levels := []topology.ReadConsistencyLevel{
		topology.ReadConsistencyLevelOne,
		topology.ReadConsistencyLevelUnstrictMajority,
		topology.ReadConsistencyLevelMajority,
		topology.ReadConsistencyLevelAll,
	}
	for _, lvl := range levels {
		t.Run(
			fmt.Sprintf("running test for %v", lvl),
			func(t *testing.T) {
				numShards := defaultNumShards
				minShard := uint32(0)
				maxShard := uint32(numShards - 1)

				// nodes = m3db nodes
				nodes, closeFn, clientopts := makeMultiNodeSetup(t, numShards, true, true, []services.ServiceInstance{
					node(t, 0, newClusterShardsRange(minShard, maxShard, shard.Available)),
					node(t, 1, newClusterShardsRange(minShard, maxShard, shard.Available)),
					node(t, 2, newClusterShardsRange(minShard, maxShard, shard.Available)),
				})

				defer closeFn()
				log := nodes[0].storageOpts.InstrumentOptions().Logger()
				// Start the nodes
				for _, n := range nodes {
					require.NoError(t, n.startServer())
				}

				c, err := client.NewClient(clientopts.SetReadConsistencyLevel(lvl))
				require.NoError(t, err)
				session, err := c.NewSession()
				require.NoError(t, err)
				defer session.Close()

				var (
					insertWg       sync.WaitGroup
					numTotalErrors uint32
				)
				now := nodes[0].db.Options().ClockOptions().NowFn()()
				start := time.Now()
				log.Info("starting data write")

				for i := 0; i < concurrency; i++ {
					insertWg.Add(1)
					idx := i
					go func() {
						numErrors := uint32(0)
						for j := 0; j < writeEach; j++ {
							id, tags := genIDTags(idx, j, numTags)
							err := session.WriteTagged(testNamespaces[0], id, tags, now, float64(1.0), xtime.Second, nil)
							if err != nil {
								numErrors++
							}
						}
						atomic.AddUint32(&numTotalErrors, numErrors)
						insertWg.Done()
					}()
				}

				insertWg.Wait()
				require.Zero(t, numTotalErrors)
				log.Infof("test data written in %v", time.Since(start))
				log.Infof("waiting to see if data is indexed")

				var (
					indexTimeout = 10 * time.Second
					fetchWg      sync.WaitGroup
				)
				for i := 0; i < concurrency; i++ {
					fetchWg.Add(1)
					idx := i
					go func() {
						id, tags := genIDTags(idx, writeEach-1, numTags)
						indexed := xclock.WaitUntil(func() bool {
							found := isIndexed(t, session, testNamespaces[0], id, tags)
							return found
						}, indexTimeout)
						assert.True(t, indexed, "timed out waiting for index retrieval")
						fetchWg.Done()
					}()
				}
				fetchWg.Wait()
				log.Infof("data is indexed in %v", time.Since(start))
			})
	}
}

func isIndexed(t *testing.T, s client.Session, ns ident.ID, id ident.ID, tags ident.TagIterator) bool {
	q := newQuery(t, tags)
	iter, _, err := s.FetchTaggedIDs(ns, index.Query{q}, index.QueryOptions{
		StartInclusive: time.Now(),
		EndExclusive:   time.Now(),
		Limit:          10})
	if err != nil {
		return false
	}
	if !iter.Next() {
		return false
	}
	cuNs, cuID, cuTag := iter.Current()
	if ns.String() != cuNs.String() {
		return false
	}
	if id.String() != cuID.String() {
		return false
	}
	return ident.NewTagIterMatcher(tags).Matches(cuTag)
}

func newQuery(t *testing.T, tags ident.TagIterator) idx.Query {
	tags = tags.Duplicate()
	filters := make([]idx.Query, 0, tags.Remaining())
	for tags.Next() {
		tag := tags.Current()
		tq := idx.NewTermQuery(tag.Name.Bytes(), tag.Value.Bytes())
		filters = append(filters, tq)
	}
	return idx.NewConjunctionQuery(filters...)
}

func genIDTags(i int, j int, numTags int) (ident.ID, ident.TagIterator) {
	id := fmt.Sprintf("foo.%d.%d", i, j)
	tags := make([]ident.Tag, 0, numTags)
	for i := 0; i < numTags; i++ {
		tags = append(tags, ident.StringTag(
			fmt.Sprintf("%s.tagname.%d", id, i),
			fmt.Sprintf("%s.tagvalue.%d", id, i),
		))
	}
	tags = append(tags,
		ident.StringTag("commoni", fmt.Sprintf("%d", i)),
		ident.StringTag("shared", "shared"))
	return ident.StringID(id), ident.NewTagSliceIterator(tags)
}
