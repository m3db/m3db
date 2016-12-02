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

package sharding

import (
	"errors"
	"math"

	"github.com/m3db/m3cluster/shard"
	"github.com/m3db/m3db/ts"

	"github.com/spaolacci/murmur3"
)

var (
	// ErrDuplicateShards returned when shard set is empty
	ErrDuplicateShards = errors.New("duplicate shards")
)

type shardSet struct {
	shards []shard.Shard
	ids    []uint32
	fn     HashFn
}

// NewShardSet creates a new sharding scheme with a set of shards
func NewShardSet(shards []shard.Shard, fn HashFn) (ShardSet, error) {
	if err := validateShards(shards); err != nil {
		return nil, err
	}
	ids := make([]uint32, len(shards))
	for i := range ids {
		ids[i] = shards[i].ID()
	}
	return &shardSet{
		shards: shards,
		ids:    ids,
		fn:     fn,
	}, nil
}

func (s *shardSet) Lookup(identifier ts.ID) uint32 {
	return s.fn(identifier)
}

func (s *shardSet) All() []shard.Shard {
	return s.shards[:]
}

func (s *shardSet) AllIDs() []uint32 {
	return s.ids[:]
}

func (s *shardSet) Min() uint32 {
	min := uint32(math.MaxUint32)
	for _, shard := range s.ids {
		if shard < min {
			min = shard
		}
	}
	return min
}

func (s *shardSet) Max() uint32 {
	max := uint32(0)
	for _, shard := range s.ids {
		if shard > max {
			max = shard
		}
	}
	return max
}

func (s *shardSet) HashFn() HashFn {
	return s.fn
}

// NewShards returns a new slice of shards with a specified state
func NewShards(ids []uint32, state shard.State) []shard.Shard {
	shards := make([]shard.Shard, len(ids))
	for i, id := range ids {
		shards[i] = shard.NewShard(id).SetState(state)
	}
	return shards
}

// IDs returns a new slice of shard IDs for a set of shards
func IDs(shards []shard.Shard) []uint32 {
	ids := make([]uint32, len(shards))
	for i := range ids {
		ids[i] = shards[i].ID()
	}
	return ids
}

func validateShards(shards []shard.Shard) error {
	uniqueShards := make(map[uint32]struct{}, len(shards))
	for _, s := range shards {
		if _, exist := uniqueShards[s.ID()]; exist {
			return ErrDuplicateShards
		}
		uniqueShards[s.ID()] = struct{}{}
	}
	return nil
}

// DefaultHashGen generates a HashFn based on murmur32
func DefaultHashGen(length int) HashFn {
	return func(id ts.ID) uint32 {
		return murmur3.Sum32(id.Data()) % uint32(length)
	}
}
