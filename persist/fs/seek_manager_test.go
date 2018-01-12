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

package fs

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	xtime "github.com/m3db/m3x/time"
	"github.com/stretchr/testify/require"
)

func TestSeekerManagerCacheShardIndices(t *testing.T) {
	shards := []uint32{2, 5, 9, 478, 1023}
	m := NewSeekerManager(nil, NewOptions(), NewBlockRetrieverOptions().FetchConcurrency()).(*seekerManager)
	var byTimes []*seekersByTime
	m.openAnyUnopenSeekersFn = func(byTime *seekersByTime) error {
		byTimes = append(byTimes, byTime)
		return nil
	}

	require.NoError(t, m.CacheShardIndices(shards))

	// Assert captured byTime objects match expectations
	require.Equal(t, len(shards), len(byTimes))
	for idx, shard := range shards {
		byTimes[idx].shard = shard
	}

	// Assert seeksByShardIdx match expectations
	shardSet := make(map[uint32]struct{}, len(shards))
	for _, shard := range shards {
		shardSet[shard] = struct{}{}
	}
	for shard, byTime := range m.seekersByShardIdx {
		_, exists := shardSet[uint32(shard)]
		if !exists {
			require.False(t, byTime.accessed)
		} else {
			require.True(t, byTime.accessed)
			require.Equal(t, uint32(shard), byTime.shard)
		}
	}
}

// TestSeekerManagerBorrowOpenSeekersLazy tests that the Borrow() method will
// open seekers lazily if they're not already open.
func TestSeekerManagerBorrowOpenSeekersLazy(t *testing.T) {
	ctrl := gomock.NewController(t)

	shards := []uint32{2, 5, 9, 478, 1023}
	m := NewSeekerManager(nil, NewOptions(), NewBlockRetrieverOptions().FetchConcurrency()).(*seekerManager)
	m.newOpenSeekerFn = func(
		shard uint32,
		blockStart time.Time,
	) (FileSetSeeker, error) {
		mock := NewMockFileSetSeeker(ctrl)
		mock.EXPECT().Open(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mock.EXPECT().Clone().Return(mock, nil)
		mock.EXPECT().ConcurrentIDBloomFilter().Return(nil)
		mock.EXPECT().ConcurrentIDBloomFilter().Return(nil)
		return mock, nil
	}

	metadata := testNs1Metadata(t)
	require.NoError(t, m.Open(metadata))
	for _, shard := range shards {
		_, err := m.Borrow(shard, time.Time{})
		require.NoError(t, err)
		byTime := m.seekersByTime(shard)
		seekers := byTime.seekers[xtime.ToUnixNano(time.Time{})]
		require.Equal(t, NewBlockRetrieverOptions().FetchConcurrency(), len(seekers.seekers))
	}
}

// TestSeekerManagerOpenCloseLoop tests the openCloseLoop of the SeekerManager
// by making sure that it makes the right decisions with regards to cleaning
// up resources based on their state.
func TestSeekerManagerOpenCloseLoop(t *testing.T) {
	// Prevent the test from running too slowly
	defer runWithSeekManagerCloseInteral(time.Millisecond)()

	ctrl := gomock.NewController(t)

	shards := []uint32{2, 5, 9, 478, 1023}
	m := NewSeekerManager(nil, NewOptions(), NewBlockRetrieverOptions().FetchConcurrency()).(*seekerManager)
	clockOpts := m.opts.ClockOptions()
	now := clockOpts.NowFn()()
	startNano := xtime.ToUnixNano(now)

	// Initialize some seekers for a time period
	m.openAnyUnopenSeekersFn = func(byTime *seekersByTime) error {
		byTime.Lock()

		// Don't overwrite if called again
		if len(byTime.seekers) != 0 {
			byTime.Unlock()
			return nil
		}
		mock := NewMockFileSetSeeker(ctrl)
		mock.EXPECT().Close().Return(nil)
		mocks := []borrowableSeeker{}
		mocks = append(mocks, borrowableSeeker{seeker: mock})
		byTime.seekers[startNano] = seekersAndBloom{
			seekers:     mocks,
			bloomFilter: nil,
		}
		byTime.Unlock()
		return nil
	}

	// Force all the seekers to be opened
	require.NoError(t, m.CacheShardIndices(shards))

	// Create a signaling mechanism so that we're notified everytime the
	// openCloseLoop ticks
	doneCh := make(chan struct{})
	m.openCloseLoopCallback = func() {
		doneCh <- struct{}{}
	}

	// // Modify the clock on the SeekerManager such that all the seekers we
	// // created are now out of retention
	metadata := testNs1Metadata(t)

	seekers := []FileSetSeeker{}

	require.NoError(t, m.Open(metadata))
	// Steps is a series of steps for the test. It is guaranteed that at least
	// one (not exactly one!) tick of the openCloseLoop will occur between every step.
	steps := []struct {
		title string
		step  func()
	}{
		{
			title: "NOOP, allow at least one tick before the next step",
			step:  func() {},
		},
		{
			title: "Make sure it didn't clean up the seekers which are still in retention",
			step: func() {
				m.RLock()
				for _, shard := range shards {
					require.Equal(t, 1, len(m.seekersByTime(shard).seekers[startNano].seekers))
				}
				m.RUnlock()
			},
		},
		{
			title: "Borrow a seeker from each shard and start the openCloseLoop",
			step: func() {
				for _, shard := range shards {
					seeker, err := m.Borrow(shard, now)
					require.NoError(t, err)
					require.NotNil(t, seeker)
					seekers = append(seekers, seeker)
				}
			},
		},
		{
			title: "Modify the clock on the seekerManager such that all of the seekers are now out of retention",
			step: func() {
				newNowFn := func() time.Time {
					return now.Add(10 * metadata.Options().RetentionOptions().RetentionPeriod())
				}
				clockOpts = clockOpts.SetNowFn(newNowFn)
				m.Lock()
				m.opts = m.opts.SetClockOptions(clockOpts)
				m.Unlock()
			},
		},
		{
			title: "Make sure the seeker manager cant be closed while seekers are borrowed",
			step: func() {
				require.Equal(t, errCantCloseSeekerManagerWhileSeekersAreBorrowed, m.Close())
			},
		},
		{
			title: "Make sure that none of the seekers were cleaned up during the openCloseLoop tick (because they're still borrowed)",
			step: func() {
				m.RLock()
				for _, shard := range shards {
					require.Equal(t, 1, len(m.seekersByTime(shard).seekers[startNano].seekers))
				}
				m.RUnlock()
			},
		},
		{
			title: "Return the borrowed seekers",
			step: func() {
				for i, seeker := range seekers {
					require.NoError(t, m.Return(shards[i], now, seeker))
				}
			},
		},
		{
			title: "Make sure that the returned seekers were cleaned up during the openCloseLoop tick",
			step: func() {
				m.RLock()
				for _, shard := range shards {
					_, ok := m.seekersByTime(shard).seekers[startNano]
					require.False(t, ok)
				}
				m.RUnlock()
			},
		},
		{
			title: "Close the seeker manager",
			step: func() {
				require.NoError(t, m.Close())
			},
		},
	}

	for _, step := range steps {
		step.step()
		<-doneCh
	}
}

func runWithSeekManagerCloseInteral(interval time.Duration) func() {
	old := seekManagerCloseInterval
	seekManagerCloseInterval = time.Millisecond
	return func() {
		seekManagerCloseInterval = old
	}
}
