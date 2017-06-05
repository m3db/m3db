package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/m3db/m3db/clock"
	mocksharding "github.com/m3db/m3db/sharding/mocks"
	"github.com/m3db/m3db/ts"

	"github.com/golang/mock/gomock"
	xtime "github.com/m3db/m3x/time"
	"github.com/stretchr/testify/require"
)

func newTestOptions(dir string) Options {
	return NewOptions().
		SetWriteEmptyShards(false).
		SetFilePathPrefix(dir).
		SetRetentionPeriod(4 * time.Hour).
		SetBlockSize(time.Hour)

}
func TestSingleConf(t *testing.T) {
	dir, err := ioutil.TempDir("", "test-single-conf")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	opts := newTestOptions(dir)
	writer := NewWriter(opts)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeShard := mocksharding.NewMockShardSet(ctrl)
	fakeShard.EXPECT().AllIDs().Return([]uint32{123}).AnyTimes()
	fakeShard.EXPECT().Lookup(gomock.Any()).Return(uint32(123)).AnyTimes()

	now := opts.ClockOptions().NowFn()().Truncate(opts.BlockSize())
	seriesMaps := make(SeriesBlocksByStart)
	seriesMaps[now] = Block(BlockConfig{
		IDs:       []string{"one-series-id"},
		NumPoints: 10,
		Start:     now,
	})
	require.NoError(t, writer.Write(ts.StringID("testmetrics"), fakeShard, seriesMaps))

	te := newFileInfoExtractor()
	require.NoError(t, filepath.Walk(dir, te.visit))
	shards := te.sortedShards()
	require.Equal(t, 1, len(shards))
	require.Equal(t, uint32(123), shards[0])
	require.Equal(t, 1, len(te.times))
	require.Equal(t, now, te.sortedTimes()[0])
}

func TestMultipleConf(t *testing.T) {
	dir, err := ioutil.TempDir("", "test-multiple-conf")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	opts := newTestOptions(dir)
	writer := NewWriter(opts)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeShard := mocksharding.NewMockShardSet(ctrl)
	fakeShard.EXPECT().AllIDs().Return([]uint32{123}).AnyTimes()
	fakeShard.EXPECT().Lookup(gomock.Any()).Return(uint32(123)).AnyTimes()

	now := opts.ClockOptions().NowFn()().Truncate(opts.BlockSize())
	past := now.Add(-opts.BlockSize())
	seriesMaps := make(SeriesBlocksByStart)
	seriesMaps[now] = Block(BlockConfig{
		IDs:       []string{"one-series-id"},
		NumPoints: 10,
		Start:     now,
	})
	seriesMaps[past] = Block(BlockConfig{
		IDs:       []string{"two-series-id"},
		NumPoints: 10,
		Start:     past,
	})
	require.NoError(t, writer.Write(ts.StringID("testmetrics"), fakeShard, seriesMaps))

	te := newFileInfoExtractor()
	require.NoError(t, filepath.Walk(dir, te.visit))
	shards := te.sortedShards()
	require.Equal(t, 1, len(shards))
	require.Equal(t, uint32(123), shards[0])
	require.Equal(t, 2, len(te.times))
	times := te.sortedTimes()
	require.Equal(t, past, times[0])
	require.Equal(t, now, times[1])
}
func TestDelayedConf(t *testing.T) {
	dir, err := ioutil.TempDir("", "test-multiple-conf")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	delay := func(d time.Duration) clock.NowFn {
		return func() time.Time {
			return time.Now().Add(-1 * d)
		}
	}
	co := clock.NewOptions().SetNowFn(delay(24 * time.Hour))
	opts := newTestOptions(dir).SetClockOptions(co)
	writer := NewWriter(opts)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeShard := mocksharding.NewMockShardSet(ctrl)
	fakeShard.EXPECT().AllIDs().Return([]uint32{123}).AnyTimes()
	fakeShard.EXPECT().Lookup(gomock.Any()).Return(uint32(123)).AnyTimes()

	now := opts.ClockOptions().NowFn()().Truncate(opts.BlockSize())
	seriesMaps := make(SeriesBlocksByStart)
	seriesMaps[now] = Block(BlockConfig{
		IDs:       []string{"one-series-id"},
		NumPoints: 10,
		Start:     now,
	})
	require.NoError(t, writer.Write(ts.StringID("testmetrics"), fakeShard, seriesMaps))

	te := newFileInfoExtractor()
	require.NoError(t, filepath.Walk(dir, te.visit))
	shards := te.sortedShards()
	require.Equal(t, 1, len(shards))
	require.Equal(t, uint32(123), shards[0])
	require.Equal(t, 1, len(te.times))
	times := te.sortedTimes()
	require.Equal(t, now, times[0])
}

type fileInfoExtractor struct {
	shards map[uint32]struct{}
	times  map[int64]struct{}
}

func newFileInfoExtractor() *fileInfoExtractor {
	return &fileInfoExtractor{
		shards: make(map[uint32]struct{}),
		times:  make(map[int64]struct{}),
	}
}

func (t *fileInfoExtractor) sortedShards() []uint32 {
	shards := make([]uint32, 0, len(t.shards))
	for i := range t.shards {
		shards = append(shards, i)
	}
	sort.Sort(uint32arr(shards))
	return shards
}

func (t *fileInfoExtractor) sortedTimes() []time.Time {
	times := make([]int64, 0, len(t.times))
	for i := range t.times {
		times = append(times, i)
	}
	sort.Sort(int64arr(times))

	timets := make([]time.Time, 0, len(t.times))
	for _, ts := range times {
		timets = append(timets, xtime.FromNanoseconds(ts))
	}
	return timets
}

func (t *fileInfoExtractor) visit(fPath string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}
	shardDir := path.Base(path.Dir(fPath))
	shardNum, err := strconv.ParseUint(shardDir, 10, 32)
	if err != nil {
		return err
	}
	t.shards[uint32(shardNum)] = struct{}{}

	name := f.Name()
	first := strings.Index(name, "-")
	if first == -1 {
		return fmt.Errorf("unable to find '-' in %v", name)
	}
	last := strings.LastIndex(name, "-")
	if last == -1 {
		return fmt.Errorf("unable to find '-' in %v", name)
	}
	if first == last {
		return fmt.Errorf("found only single '-' in %v", name)
	}
	num, parseErr := strconv.ParseInt(name[first+1:last], 10, 64)
	if parseErr != nil {
		return err
	}
	t.times[num] = struct{}{}
	return nil
}

type uint32arr []uint32

func (a uint32arr) Len() int           { return len(a) }
func (a uint32arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a uint32arr) Less(i, j int) bool { return a[i] < a[j] }

type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }
