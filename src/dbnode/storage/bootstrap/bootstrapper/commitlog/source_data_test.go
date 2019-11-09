// Copyright (c) 2018 Uber Technologies, Inc.
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

package commitlog

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"testing"
	"time"

	"github.com/m3db/m3/src/dbnode/digest"
	"github.com/m3db/m3/src/dbnode/encoding"
	"github.com/m3db/m3/src/dbnode/namespace"
	"github.com/m3db/m3/src/dbnode/persist"
	"github.com/m3db/m3/src/dbnode/persist/fs"
	"github.com/m3db/m3/src/dbnode/persist/fs/commitlog"
	"github.com/m3db/m3/src/dbnode/storage/block"
	"github.com/m3db/m3/src/dbnode/storage/bootstrap"
	"github.com/m3db/m3/src/dbnode/storage/bootstrap/result"
	"github.com/m3db/m3/src/dbnode/storage/series"
	"github.com/m3db/m3/src/dbnode/topology"
	"github.com/m3db/m3/src/dbnode/ts"
	"github.com/m3db/m3/src/x/checked"
	"github.com/m3db/m3/src/x/context"
	"github.com/m3db/m3/src/x/ident"
	"github.com/m3db/m3/src/x/pool"
	xtime "github.com/m3db/m3/src/x/time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	testNamespaceID    = ident.StringID("commitlog_test_ns")
	testDefaultRunOpts = bootstrap.NewRunOptions().
				SetInitialTopologyState(&topology.StateSnapshot{})
	minCommitLogRetention = 10 * time.Minute
)

func testNsMetadata(t *testing.T) namespace.Metadata {
	md, err := namespace.NewMetadata(testNamespaceID, namespace.NewOptions())
	require.NoError(t, err)
	return md
}

func TestAvailableEmptyRangeError(t *testing.T) {
	var (
		opts     = testDefaultOpts
		src      = newCommitLogSource(opts, fs.Inspection{})
		res, err = src.AvailableData(testNsMetadata(t), result.ShardTimeRanges{}, testDefaultRunOpts)
	)
	require.NoError(t, err)
	require.True(t, result.ShardTimeRanges{}.Equal(res))
}

func TestReadEmpty(t *testing.T) {
	opts := testDefaultOpts

	src := newCommitLogSource(opts, fs.Inspection{})
	md := testNsMetadata(t)
	target := result.ShardTimeRanges{}
	tester := bootstrap.BuildNamespacesTester(t, testDefaultRunOpts, target, md)
	defer tester.Finish()

	tester.TestReadWith(src)
	tester.TestUnfulfilledForNamespaceIsEmpty(md)
}

func TestReadErrorOnNewIteratorError(t *testing.T) {
	opts := testDefaultOpts
	src := newCommitLogSource(opts, fs.Inspection{}).(*commitLogSource)

	src.newIteratorFn = func(
		_ commitlog.IteratorOpts,
	) (commitlog.Iterator, []commitlog.ErrorWithPath, error) {
		return nil, nil, fmt.Errorf("an error")
	}

	ranges := xtime.Ranges{}
	ranges = ranges.AddRange(xtime.Range{
		Start: time.Now(),
		End:   time.Now().Add(time.Hour),
	})

	md := testNsMetadata(t)
	target := result.ShardTimeRanges{0: ranges}
	tester := bootstrap.BuildNamespacesTester(t, testDefaultRunOpts, target, md)
	defer tester.Finish()

	res, err := src.Read(tester.Namespaces)
	require.Error(t, err)
	require.Nil(t, res.Results)
}

func TestReadOrderedValues(t *testing.T) {
	opts := testDefaultOpts
	md := testNsMetadata(t)
	testReadOrderedValues(t, opts, md, nil)
}

func testReadOrderedValues(t *testing.T, opts Options, md namespace.Metadata, setAnn setAnnotation) {
	nsCtx := namespace.NewContextFrom(md)

	src := newCommitLogSource(opts, fs.Inspection{}).(*commitLogSource)

	blockSize := md.Options().RetentionOptions().BlockSize()
	now := time.Now()
	start := now.Truncate(blockSize).Add(-blockSize)
	end := now.Truncate(blockSize)

	// Request a little after the start of data, because always reading full blocks
	// it should return the entire block beginning from "start"
	require.True(t, blockSize >= minCommitLogRetention)
	ranges := xtime.Ranges{}
	ranges = ranges.AddRange(xtime.Range{
		Start: start,
		End:   end,
	})

	foo := ts.Series{Namespace: nsCtx.ID, Shard: 0, ID: ident.StringID("foo")}
	bar := ts.Series{Namespace: nsCtx.ID, Shard: 1, ID: ident.StringID("bar")}
	baz := ts.Series{Namespace: nsCtx.ID, Shard: 2, ID: ident.StringID("baz")}

	values := []testValue{
		{foo, start, 1.0, xtime.Second, nil},
		{foo, start.Add(1 * time.Minute), 2.0, xtime.Second, nil},
		{bar, start.Add(2 * time.Minute), 1.0, xtime.Second, nil},
		{bar, start.Add(3 * time.Minute), 2.0, xtime.Second, nil},
		// "baz" is in shard 2 and should not be returned
		{baz, start.Add(4 * time.Minute), 1.0, xtime.Second, nil},
	}
	if setAnn != nil {
		values = setAnn(values)
	}

	src.newIteratorFn = func(
		_ commitlog.IteratorOpts,
	) (commitlog.Iterator, []commitlog.ErrorWithPath, error) {
		return newTestCommitLogIterator(values, nil), nil, nil
	}

	targetRanges := result.ShardTimeRanges{0: ranges, 1: ranges}
	tester := bootstrap.BuildNamespacesTester(t, testDefaultRunOpts, targetRanges, md)
	defer tester.Finish()

	tester.TestReadWith(src)
	tester.TestUnfulfilledForNamespaceIsEmpty(md)

	read := tester.DumpWritesForNamespace(md)
	require.Equal(t, 2, len(read))
	enforceValuesAreCorrect(t, values[:4], read)
}

func TestReadUnorderedValues(t *testing.T) {
	opts := testDefaultOpts
	md := testNsMetadata(t)
	testReadUnorderedValues(t, opts, md, nil)
}

func testReadUnorderedValues(t *testing.T, opts Options, md namespace.Metadata, setAnn setAnnotation) {
	nsCtx := namespace.NewContextFrom(md)
	src := newCommitLogSource(opts, fs.Inspection{}).(*commitLogSource)

	blockSize := md.Options().RetentionOptions().BlockSize()
	now := time.Now()
	start := now.Truncate(blockSize).Add(-blockSize)
	end := now.Truncate(blockSize)

	// Request a little after the start of data, because always reading full blocks
	// it should return the entire block beginning from "start"
	require.True(t, blockSize >= minCommitLogRetention)
	ranges := xtime.Ranges{}
	ranges = ranges.AddRange(xtime.Range{
		Start: start,
		End:   end,
	})

	foo := ts.Series{Namespace: nsCtx.ID, Shard: 0, ID: ident.StringID("foo")}

	values := []testValue{
		{foo, start.Add(10 * time.Minute), 1.0, xtime.Second, nil},
		{foo, start.Add(1 * time.Minute), 2.0, xtime.Second, nil},
		{foo, start.Add(2 * time.Minute), 3.0, xtime.Second, nil},
		{foo, start.Add(3 * time.Minute), 4.0, xtime.Second, nil},
		{foo, start, 5.0, xtime.Second, nil},
	}
	if setAnn != nil {
		values = setAnn(values)
	}

	src.newIteratorFn = func(
		_ commitlog.IteratorOpts,
	) (commitlog.Iterator, []commitlog.ErrorWithPath, error) {
		return newTestCommitLogIterator(values, nil), nil, nil
	}

	targetRanges := result.ShardTimeRanges{0: ranges, 1: ranges}
	tester := bootstrap.BuildNamespacesTester(t, testDefaultRunOpts, targetRanges, md)
	defer tester.Finish()

	tester.TestReadWith(src)
	tester.TestUnfulfilledForNamespaceIsEmpty(md)

	read := tester.DumpWritesForNamespace(md)
	require.Equal(t, 1, len(read))
	enforceValuesAreCorrect(t, values, read)
}

// TestReadHandlesDifferentSeriesWithIdenticalUniqueIndex was added as a
// regression test to make sure that the commit log bootstrapper does not make
// any assumptions about series having a unique index because that only holds
// for the duration that an M3DB node is on, but commit log files can span
// multiple M3DB processes which means that unique indexes could be re-used
// for multiple different series.
func TestReadHandlesDifferentSeriesWithIdenticalUniqueIndex(t *testing.T) {
	opts := testDefaultOpts
	md := testNsMetadata(t)

	nsCtx := namespace.NewContextFrom(md)
	src := newCommitLogSource(opts, fs.Inspection{}).(*commitLogSource)

	blockSize := md.Options().RetentionOptions().BlockSize()
	now := time.Now()
	start := now.Truncate(blockSize).Add(-blockSize)
	end := now.Truncate(blockSize)

	require.True(t, blockSize >= minCommitLogRetention)
	ranges := xtime.Ranges{}
	ranges = ranges.AddRange(xtime.Range{
		Start: start,
		End:   end,
	})

	// All series need to be in the same shard to exercise the regression.
	foo := ts.Series{
		Namespace: nsCtx.ID, Shard: 0, ID: ident.StringID("foo"), UniqueIndex: 0}
	bar := ts.Series{
		Namespace: nsCtx.ID, Shard: 0, ID: ident.StringID("bar"), UniqueIndex: 0}

	values := []testValue{
		{foo, start, 1.0, xtime.Second, nil},
		{bar, start, 2.0, xtime.Second, nil},
	}

	src.newIteratorFn = func(
		_ commitlog.IteratorOpts,
	) (commitlog.Iterator, []commitlog.ErrorWithPath, error) {
		return newTestCommitLogIterator(values, nil), nil, nil
	}

	targetRanges := result.ShardTimeRanges{0: ranges, 1: ranges}
	tester := bootstrap.BuildNamespacesTester(t, testDefaultRunOpts, targetRanges, md)
	defer tester.Finish()

	tester.TestReadWith(src)
	tester.TestUnfulfilledForNamespaceIsEmpty(md)

	read := tester.DumpWritesForNamespace(md)
	require.Equal(t, 2, len(read))
	enforceValuesAreCorrect(t, values, read)
}

func TestReadTrimsToRanges(t *testing.T) {
	opts := testDefaultOpts
	md := testNsMetadata(t)
	testReadTrimsToRanges(t, opts, md, nil)
}

func testReadTrimsToRanges(t *testing.T, opts Options, md namespace.Metadata, setAnn setAnnotation) {
	nsCtx := namespace.NewContextFrom(md)
	src := newCommitLogSource(opts, fs.Inspection{}).(*commitLogSource)

	blockSize := md.Options().RetentionOptions().BlockSize()
	now := time.Now()
	start := now.Truncate(blockSize).Add(-blockSize)
	end := now.Truncate(blockSize)

	// Request a little after the start of data, because always reading full blocks
	// it should return the entire block beginning from "start"
	require.True(t, blockSize >= minCommitLogRetention)
	ranges := xtime.Ranges{}
	ranges = ranges.AddRange(xtime.Range{
		Start: start,
		End:   end,
	})

	foo := ts.Series{Namespace: nsCtx.ID, Shard: 0, ID: ident.StringID("foo")}

	values := []testValue{
		{foo, start.Add(-1 * time.Minute), 1.0, xtime.Nanosecond, nil},
		{foo, start, 2.0, xtime.Nanosecond, nil},
		{foo, start.Add(1 * time.Minute), 3.0, xtime.Nanosecond, nil},
		{foo, end.Truncate(blockSize).Add(blockSize).Add(time.Nanosecond), 4.0, xtime.Nanosecond, nil},
	}
	if setAnn != nil {
		values = setAnn(values)
	}

	src.newIteratorFn = func(
		_ commitlog.IteratorOpts,
	) (commitlog.Iterator, []commitlog.ErrorWithPath, error) {
		return newTestCommitLogIterator(values, nil), nil, nil
	}

	targetRanges := result.ShardTimeRanges{0: ranges, 1: ranges}
	tester := bootstrap.BuildNamespacesTester(t, testDefaultRunOpts, targetRanges, md)
	defer tester.Finish()

	tester.TestReadWith(src)
	tester.TestUnfulfilledForNamespaceIsEmpty(md)

	read := tester.DumpWritesForNamespace(md)
	require.Equal(t, 1, len(read))
	enforceValuesAreCorrect(t, values[1:3], read)
}

func TestItMergesSnapshotsAndCommitLogs(t *testing.T) {
	opts := testDefaultOpts
	md := testNsMetadata(t)

	testItMergesSnapshotsAndCommitLogs(t, opts, md, nil)
}

func testItMergesSnapshotsAndCommitLogs(t *testing.T, opts Options,
	md namespace.Metadata, setAnn setAnnotation) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		nsCtx     = namespace.NewContextFrom(md)
		src       = newCommitLogSource(opts, fs.Inspection{}).(*commitLogSource)
		blockSize = md.Options().RetentionOptions().BlockSize()
		now       = time.Now()
		start     = now.Truncate(blockSize).Add(-blockSize)
		end       = now.Truncate(blockSize)
		ranges    = xtime.Ranges{}

		foo             = ts.Series{Namespace: nsCtx.ID, Shard: 0, ID: ident.StringID("foo")}
		commitLogValues = []testValue{
			{foo, start.Add(2 * time.Minute), 1.0, xtime.Nanosecond, nil},
			{foo, start.Add(3 * time.Minute), 2.0, xtime.Nanosecond, nil},
			{foo, start.Add(4 * time.Minute), 3.0, xtime.Nanosecond, nil},

			// Should not be present
			{foo, end.Truncate(blockSize).Add(blockSize).Add(time.Nanosecond), 4.0, xtime.Nanosecond, nil},
		}
	)
	if setAnn != nil {
		commitLogValues = setAnn(commitLogValues)
	}

	// Request a little after the start of data, because always reading full blocks it
	// should return the entire block beginning from "start".
	require.True(t, blockSize >= minCommitLogRetention)

	ranges = ranges.AddRange(xtime.Range{
		Start: start,
		End:   end,
	})

	src.newIteratorFn = func(
		_ commitlog.IteratorOpts,
	) (commitlog.Iterator, []commitlog.ErrorWithPath, error) {
		return newTestCommitLogIterator(commitLogValues, nil), nil, nil
	}

	src.snapshotFilesFn = func(
		filePathPrefix string,
		namespace ident.ID,
		shard uint32,
	) (fs.FileSetFilesSlice, error) {
		return fs.FileSetFilesSlice{
			fs.FileSetFile{
				ID: fs.FileSetFileIdentifier{
					Namespace:   namespace,
					BlockStart:  start,
					Shard:       shard,
					VolumeIndex: 0,
				},
				// Make sure path passes the "is snapshot" check in SnapshotTimeAndID method.
				AbsoluteFilepaths:               []string{"snapshots/checkpoint"},
				CachedHasCompleteCheckpointFile: fs.EvalTrue,
				CachedSnapshotTime:              start.Add(time.Minute),
			},
		}, nil
	}

	mockReader := fs.NewMockDataFileSetReader(ctrl)
	mockReader.EXPECT().Open(fs.ReaderOpenOptionsMatcher{
		ID: fs.FileSetFileIdentifier{
			Namespace:   nsCtx.ID,
			BlockStart:  start,
			Shard:       0,
			VolumeIndex: 0,
		},
		FileSetType: persist.FileSetSnapshotType,
	}).Return(nil).AnyTimes()
	mockReader.EXPECT().Entries().Return(1).AnyTimes()
	mockReader.EXPECT().Close().Return(nil).AnyTimes()

	snapshotValues := []testValue{
		{foo, start.Add(1 * time.Minute), 1.0, xtime.Nanosecond, nil},
	}
	if setAnn != nil {
		snapshotValues = setAnn(snapshotValues)
	}

	encoderPool := opts.ResultOptions().DatabaseBlockOptions().EncoderPool()
	encoder := encoderPool.Get()
	encoder.Reset(snapshotValues[0].t, 10, nsCtx.Schema)
	for _, value := range snapshotValues {
		dp := ts.Datapoint{
			Timestamp: value.t,
			Value:     value.v,
		}
		encoder.Encode(dp, value.u, value.a)
	}

	ctx := context.NewContext()
	defer ctx.Close()

	reader, ok := encoder.Stream(ctx)
	require.True(t, ok)

	seg, err := reader.Segment()
	require.NoError(t, err)

	bytes := make([]byte, seg.Len())
	_, err = reader.Read(bytes)
	require.NoError(t, err)

	mockReader.EXPECT().Read().Return(
		foo.ID,
		ident.EmptyTagIterator,
		checked.NewBytes(bytes, nil),
		digest.Checksum(bytes),
		nil,
	)
	mockReader.EXPECT().Read().Return(nil, nil, nil, uint32(0), io.EOF)

	src.newReaderFn = func(
		bytesPool pool.CheckedBytesPool,
		opts fs.Options,
	) (fs.DataFileSetReader, error) {
		return mockReader, nil
	}

	targetRanges := result.ShardTimeRanges{0: ranges}
	tester := bootstrap.BuildNamespacesTesterWithReaderIteratorPool(
		t,
		testDefaultRunOpts,
		targetRanges,
		opts.ResultOptions().DatabaseBlockOptions().MultiReaderIteratorPool(),
		md,
	)

	defer tester.Finish()
	tester.TestReadWith(src)
	tester.TestUnfulfilledForNamespaceIsEmpty(md)

	read := tester.DumpWritesForNamespace(md)
	require.Equal(t, 1, len(read))
	enforceValuesAreCorrect(t, commitLogValues[0:3], read)

	// NB: this case is a little tricky in that it's combining writes that come
	// through both the `.Write()` and `.LoadBlock()` methods, which get read
	// separately in the bootstrap test utility. Thus it's necessary to combine
	// the results from both paths here prior to comparison.
	read = tester.DumpValuesForNamespace(md)
	enforceValuesAreCorrect(t, snapshotValues, read)
}

type setAnnotation func([]testValue) []testValue
type annotationEqual func([]byte, []byte) bool

type testValue struct {
	s ts.Series
	t time.Time
	v float64
	u xtime.Unit
	a ts.Annotation
}

type seriesShardResultBlock struct {
	encoder encoding.Encoder
}

type seriesShardResult struct {
	blocks map[xtime.UnixNano]*seriesShardResultBlock
	result block.DatabaseSeriesBlocks
}

func equals(ac series.DecodedTestValue, ex testValue) bool {
	return ac.Timestamp.Equal(ex.t) &&
		math.Abs(ac.Value-ex.v) < 0.000001 &&
		ac.Unit == ex.u &&
		bytes.Equal(ac.Annotation, ex.a)
}

func enforceValuesAreCorrect(
	t *testing.T,
	values []testValue,
	actual bootstrap.DecodedBlockMap,
) {
	require.NoError(t, verifyValuesAreCorrect(values, actual))
}

func verifyValuesAreCorrect(
	values []testValue,
	actual bootstrap.DecodedBlockMap,
) error {
	if actual == nil || len(actual) == 0 {
		if 0 != len(values) {
			return fmt.Errorf("expected %v, got nil", values)
		}
	}

	count := 0
	for _, ex := range values {
		id := ex.s.ID.String()
		acList := actual[id]
		found := false
		for _, ac := range acList {
			if equals(ac, ex) {
				count++
				found = true
			}
		}

		if !found {
			return fmt.Errorf("could not find %s", id)
		}
	}

	// Ensure there are no extra values in the actual result.
	actualCount := 0
	for _, v := range actual {
		actualCount += len(v)
	}

	if actualCount != count {
		return fmt.Errorf("expected %d values, got %d values", count, actualCount)
	}

	return nil
}

type testCommitLogIterator struct {
	values []testValue
	idx    int
	err    error
	closed bool
}

type testValuesByTime []testValue

func (v testValuesByTime) Len() int      { return len(v) }
func (v testValuesByTime) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v testValuesByTime) Less(i, j int) bool {
	return v[i].t.Before(v[j].t)
}

func newTestCommitLogIterator(values []testValue, err error) *testCommitLogIterator {
	return &testCommitLogIterator{values: values, idx: -1, err: err}
}

func (i *testCommitLogIterator) Next() bool {
	i.idx++
	return i.idx < len(i.values)
}

func (i *testCommitLogIterator) Current() commitlog.LogEntry {
	idx := i.idx
	if idx == -1 {
		idx = 0
	}
	v := i.values[idx]
	return commitlog.LogEntry{
		Series:     v.s,
		Datapoint:  ts.Datapoint{Timestamp: v.t, Value: v.v},
		Unit:       v.u,
		Annotation: v.a,
		Metadata: commitlog.LogEntryMetadata{
			FileReadID:        uint64(idx) + 1,
			SeriesUniqueIndex: v.s.UniqueIndex,
		},
	}
}

func (i *testCommitLogIterator) Err() error {
	return i.err
}

func (i *testCommitLogIterator) Close() {
	i.closed = true
}
