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

package topology

import (
	"github.com/m3db/m3db/sharding"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3x/watch"
)

type staticMap struct {
	shardSet            sharding.ShardSet
	hostShardSets       []HostShardSet
	hostShardSetsByID   map[string]HostShardSet
	orderedHosts        []Host
	hostsByShard        [][]Host
	orderedHostsByShard [][]orderedHost
	replicas            int
	majority            int
}

func newStaticMap(opts StaticOptions) Map {
	totalShards := len(opts.ShardSet().Shards())
	hostShardSets := opts.HostShardSets()
	topoMap := staticMap{
		shardSet:            opts.ShardSet(),
		hostShardSets:       hostShardSets,
		hostShardSetsByID:   make(map[string]HostShardSet),
		orderedHosts:        make([]Host, 0, len(hostShardSets)),
		hostsByShard:        make([][]Host, totalShards),
		orderedHostsByShard: make([][]orderedHost, totalShards),
		replicas:            opts.Replicas(),
		majority:            majority(opts.Replicas()),
	}

	for idx, hostShardSet := range hostShardSets {
		host := hostShardSet.Host()
		topoMap.hostShardSetsByID[host.ID()] = hostShardSet
		topoMap.orderedHosts = append(topoMap.orderedHosts, host)
		for _, shard := range hostShardSet.ShardSet().Shards() {
			topoMap.hostsByShard[shard] = append(topoMap.hostsByShard[shard], host)
			topoMap.orderedHostsByShard[shard] = append(topoMap.orderedHostsByShard[shard], orderedHost{
				idx:  idx,
				host: host,
			})
		}
	}

	return &topoMap
}

type orderedHost struct {
	idx  int
	host Host
}

func (t *staticMap) Hosts() []Host {
	return t.orderedHosts
}

func (t *staticMap) HostShardSets() []HostShardSet {
	return t.hostShardSets
}

func (t *staticMap) LookupHostShardSet(id string) (HostShardSet, bool) {
	value, ok := t.hostShardSetsByID[id]
	return value, ok
}

func (t *staticMap) HostsLen() int {
	return len(t.orderedHosts)
}

func (t *staticMap) ShardSet() sharding.ShardSet {
	return t.shardSet
}

func (t *staticMap) Route(id ts.ID) (uint32, []Host, error) {
	shard := t.shardSet.Shard(id)
	if int(shard) >= len(t.hostsByShard) {
		return shard, nil, errUnownedShard
	}
	return shard, t.hostsByShard[shard], nil
}

func (t *staticMap) RouteForEach(id ts.ID, forEachFn RouteForEachFn) error {
	return t.RouteShardForEach(t.shardSet.Shard(id), forEachFn)
}

func (t *staticMap) RouteShard(shard uint32) ([]Host, error) {
	if int(shard) >= len(t.hostsByShard) {
		return nil, errUnownedShard
	}
	return t.hostsByShard[shard], nil
}

func (t *staticMap) RouteShardForEach(shard uint32, forEachFn RouteForEachFn) error {
	if int(shard) >= len(t.orderedHostsByShard) {
		return errUnownedShard
	}
	orderedHosts := t.orderedHostsByShard[shard]
	for i := range orderedHosts {
		forEachFn(orderedHosts[i].idx, orderedHosts[i].host)
	}
	return nil
}

func (t *staticMap) Replicas() int {
	return t.replicas
}

func (t *staticMap) MajorityReplicas() int {
	return t.majority
}

type mapWatch struct {
	xwatch.Watch
}

func newMapWatch(w xwatch.Watch) MapWatch {
	return &mapWatch{w}
}

func (w *mapWatch) C() <-chan struct{} {
	return w.Watch.C()
}

func (w *mapWatch) Get() Map {
	value := w.Watch.Get()
	if value == nil {
		return nil
	}
	return value.(Map)
}

func (w *mapWatch) Close() {
	w.Watch.Close()
}
