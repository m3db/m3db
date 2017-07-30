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

package commitlog_test

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"
	"time"

	"github.com/m3db/m3db/context"
	"github.com/m3db/m3db/persist/fs/commitlog"
	"github.com/m3db/m3db/ts"
	xtime "github.com/m3db/m3x/time"
	"github.com/stretchr/testify/require"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/commands"
	"github.com/leanovate/gopter/gen"
)

var (
	defaultCommitLogSeed     int64  = 12345678
	defaultShardNum          uint32 = 1234
	defaultTestNamespace            = ts.StringID("testNamespace")
	defaultTestFlushInterval        = time.Millisecond
)

func TestCommitLogReadWrite(t *testing.T) {
	baseTestPath, err := ioutil.TempDir("", "commit-log-test-base-dir")
	require.NoError(t, err)
	defer os.RemoveAll(baseTestPath)

	opts := commitlog.NewOptions()
	fsOpts := opts.FilesystemOptions().SetFilePathPrefix(baseTestPath)
	opts = opts.SetFilesystemOptions(fsOpts).SetFlushInterval(time.Millisecond)

	cl := commitlog.NewCommitLog(opts)
	require.NoError(t, cl.Open())

	params := gopter.DefaultGenParameters()
	writeResult := gen.SliceOfN(100, genWrite())(params)
	writesInterface, ok := writeResult.Retrieve()
	require.True(t, ok)
	writes, ok := writesInterface.([]generatedWrite)

	ctx := context.NewContext()
	for _, w := range writes {
		require.NoError(t, cl.WriteBehind(ctx, w.series, w.datapoint, w.unit, w.annotation))
	}
	ctx.Close()
	require.NoError(t, cl.Close())

	i := 0
	iter, err := commitlog.NewIterator(opts)
	require.NoError(t, err)
	defer iter.Close()
	for ; iter.Next(); i++ {
		series, datapoint, _, _ := iter.Current()
		write := writes[i]
		require.Equal(t, write.series.ID.String(), series.ID.String())
		require.Equal(t, write.datapoint.Value, datapoint.Value)
		require.True(t, write.datapoint.Timestamp.Equal(datapoint.Timestamp))
	}
	require.Equal(t, len(writes), i)
}

func TestCommitLogRWOnce(t *testing.T) {
	baseTestPath, err := ioutil.TempDir("", "commit-log-test-base-dir")
	require.NoError(t, err)
	defer os.RemoveAll(baseTestPath)

	opts := commitlog.NewOptions()
	fsOpts := opts.FilesystemOptions().SetFilePathPrefix(baseTestPath)
	opts = opts.SetFilesystemOptions(fsOpts).SetFlushInterval(time.Millisecond)

	params := gopter.DefaultGenParameters()
	writeResult := genWrite()(params)
	writesInterface, ok := writeResult.Retrieve()
	require.True(t, ok)
	w, ok := writesInterface.(generatedWrite)
	require.True(t, ok)

	cl := commitlog.NewCommitLog(opts)
	require.NoError(t, cl.Open())

	ctx := context.NewContext()
	require.NoError(t, cl.Write(ctx, w.series, w.datapoint, w.unit, w.annotation))
	ctx.Close()
	require.NoError(t, cl.Close())

	iter, err := commitlog.NewIterator(opts)
	require.NoError(t, err)
	defer iter.Close()
	i := 0
	for ; iter.Next(); i++ {
		series, datapoint, _, _ := iter.Current()
		require.Equal(t, w.series.ID.String(), series.ID.String())
		require.Equal(t, w.datapoint.Value, datapoint.Value)
		require.True(t, w.datapoint.Timestamp.Equal(datapoint.Timestamp))
	}
	require.Equal(t, 1, i)
}

func TestCommitLogPropTest(t *testing.T) {
	basePath, err := ioutil.TempDir("", "commit-log-tests")
	require.NoError(t, err)
	defer os.RemoveAll(basePath)

	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	properties := gopter.NewProperties(parameters)

	comms := clCommandFunctor(basePath, t)
	properties.Property("CommitLog System", commands.Prop(comms))
	properties.TestingRun(t)
}

// clCommandFunctor is a var which implements the command.Commands interface,
// i.e. is responsible for creating/destroying the system under test and generating
// commands and initial states (clState)
func clCommandFunctor(basePath string, t *testing.T) *commands.ProtoCommands {
	return &commands.ProtoCommands{
		NewSystemUnderTestFunc: func(initialState commands.State) commands.SystemUnderTest {
			return initialState
		},
		InitialStateGen: genState(basePath, t).WithLabel("state"),
		InitialPreConditionFunc: func(state commands.State) bool {
			if state == nil {
				return false
			}
			_, ok := state.(*clState)
			return ok
		},
		GenCommandFunc: func(state commands.State) gopter.Gen {
			return gen.OneGenOf(genOpenCommand, genCloseCommand, genWriteBehindCommand)
		},
	}
}

// operations on *clState

var genOpenCommand = gen.Const(&commands.ProtoCommand{
	Name: "Open",
	PreConditionFunc: func(state commands.State) bool {
		return !state.(*clState).open
	},
	RunFunc: func(q commands.SystemUnderTest) commands.Result {
		s := q.(*clState)
		s.cLog = commitlog.NewCommitLog(s.opts)
		return s.cLog.Open()
	},
	NextStateFunc: func(state commands.State) commands.State {
		s := state.(*clState)
		s.open = true
		return s
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		if result == nil {
			return &gopter.PropResult{Status: gopter.PropTrue}
		}
		return &gopter.PropResult{
			Status: gopter.PropFalse,
			Error:  result.(error),
		}
	},
})

var genCloseCommand = gen.Const(&commands.ProtoCommand{
	Name: "Close",
	PreConditionFunc: func(state commands.State) bool {
		return state.(*clState).open
	},
	RunFunc: func(q commands.SystemUnderTest) commands.Result {
		s := q.(*clState)
		return s.cLog.Close()
	},
	NextStateFunc: func(state commands.State) commands.State {
		s := state.(*clState)
		s.open = false
		s.cLog = nil
		return s
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		if result != nil {
			return &gopter.PropResult{
				Status: gopter.PropError,
				Error:  result.(error),
			}
		}
		s := state.(*clState)
		err := s.writesArePresent(s.pendingWrites...)
		if err != nil {
			return &gopter.PropResult{
				Status: gopter.PropError,
				Error:  err.(error),
			}
		}
		return &gopter.PropResult{Status: gopter.PropTrue}
	},
})

var genWriteBehindCommand = genWrite().
	Map(func(w generatedWrite) commands.Command {
		return &commands.ProtoCommand{
			Name: "WriteBehind",
			PreConditionFunc: func(state commands.State) bool {
				return state.(*clState).open
			},
			RunFunc: func(q commands.SystemUnderTest) commands.Result {
				s := q.(*clState)
				ctx := context.NewContext()
				defer ctx.Close()
				return s.cLog.WriteBehind(ctx, w.series, w.datapoint, w.unit, w.annotation)
			},
			NextStateFunc: func(state commands.State) commands.State {
				s := state.(*clState)
				s.pendingWrites = append(s.pendingWrites, w)
				return s
			},
			PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
				if result == nil {
					return &gopter.PropResult{Status: gopter.PropTrue}
				}
				return &gopter.PropResult{
					Status: gopter.PropFalse,
					Error:  result.(error),
				}
			},
		}
	})

// clState holds the expected state (i.e. its the commands.State), and we use it as the SystemUnderTest
type clState struct {
	basePath      string
	opts          commitlog.Options
	open          bool
	cLog          commitlog.CommitLog
	pendingWrites []generatedWrite
}

// generator for commit log write
func genState(basePath string, t *testing.T) gopter.Gen {
	return gen.Identifier().
		MapResult(func(r *gopter.GenResult) *gopter.GenResult {
			iface, ok := r.Retrieve()
			if !ok {
				return gopter.NewEmptyResult(reflect.PtrTo(reflect.TypeOf(clState{})))
			}
			p, ok := iface.(string)
			if !ok {
				return gopter.NewEmptyResult(reflect.PtrTo(reflect.TypeOf(clState{})))
			}
			initPath := path.Join(basePath, p)
			result := newInitState(initPath, t)
			return gopter.NewGenResult(result, gopter.NoShrinker)
		})
}

func newInitState(dir string, t *testing.T) *clState {
	opts := commitlog.NewOptions().SetFlushInterval(defaultTestFlushInterval)
	fsOpts := opts.FilesystemOptions().SetFilePathPrefix(dir)
	opts = opts.SetFilesystemOptions(fsOpts)
	return &clState{
		basePath: dir,
		opts:     opts,
	}
}

func (s *clState) writesArePresent(writes ...generatedWrite) error {
	writesOnDisk := make(map[ts.Hash]map[time.Time]float64)
	iter, err := commitlog.NewIterator(s.opts)
	if err != nil {
		return err
	}

	defer iter.Close()
	for iter.Next() {
		series, datapoint, _, _ := iter.Current()
		idHash := series.ID.Hash()
		seriesMap, ok := writesOnDisk[idHash]
		if !ok {
			seriesMap = make(map[time.Time]float64)
			writesOnDisk[idHash] = seriesMap
		}
		seriesMap[datapoint.Timestamp] = datapoint.Value
	}
	if err := iter.Err(); err != nil {
		return err
	}

	missingErr := fmt.Errorf("writesOnDisk: %+v, writes: %+v", writesOnDisk, writes)
	for _, w := range writes {
		idHash := w.series.ID.Hash()
		seriesMap, ok := writesOnDisk[idHash]
		if !ok {
			return missingErr
		}
		val, ok := seriesMap[w.datapoint.Timestamp]
		if !ok {
			return missingErr
		}
		if val != w.datapoint.Value {
			return missingErr
		}
	}
	return nil
}

type generatedWrite struct {
	series     commitlog.Series
	datapoint  ts.Datapoint
	unit       xtime.Unit
	annotation ts.Annotation
}

func (w generatedWrite) String() string {
	return fmt.Sprintf("ID = %v, Datapoint = %+v", w.series.ID.String(), w.datapoint)
}

// generator for commit log write
func genWrite() gopter.Gen {
	return gopter.CombineGens(
		gen.Identifier(),
		gen.TimeRange(time.Now(), 15*time.Minute),
		gen.Float64(),
	).Map(func(val []interface{}) generatedWrite {
		id := val[0].(string)
		t := val[1].(time.Time)
		v := val[2].(float64)

		return generatedWrite{
			series: commitlog.Series{
				ID:          ts.StringID(id),
				Namespace:   defaultTestNamespace,
				Shard:       defaultShardNum,
				UniqueIndex: uniqueID(id),
			},
			datapoint: ts.Datapoint{
				Timestamp: t,
				Value:     v,
			},
			unit: xtime.Nanosecond,
		}
	})
}

func uniqueID(s string) uint64 {
	return uint64(crc32.Checksum([]byte(s), crc32.IEEETable))
}
