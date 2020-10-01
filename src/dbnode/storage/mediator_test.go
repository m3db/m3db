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

package storage

import (
	"testing"
	"time"

	xclock "github.com/m3db/m3/src/x/clock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
)

func TestDatabaseMediatorOpenClose(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := DefaultTestOptions().SetRepairEnabled(false)
	now := time.Now()
	opts = opts.
		SetBootstrapProcessProvider(nil).
		SetClockOptions(opts.ClockOptions().SetNowFn(func() time.Time {
			return now
		})).
		SetInstrumentOptions(opts.InstrumentOptions().SetReportInterval(time.Millisecond))

	db := NewMockdatabase(ctrl)
	db.EXPECT().Options().Return(opts).AnyTimes()
	db.EXPECT().OwnedNamespaces().Return(nil, nil).AnyTimes()
	db.EXPECT().BootstrapState().Return(DatabaseBootstrapState{}).AnyTimes()
	db.EXPECT().IsBootstrappedAndDurable().Return(true).AnyTimes()
	m, err := newMediator(db, nil, opts)
	require.NoError(t, err)

	var executed, reported atomic.Bool

	backgroundProcess := NewMockBackgroundProcess(ctrl)
	backgroundProcess.EXPECT().Report().Do(func() {
		reported.Store(true)
	}).AnyTimes()
	gomock.InOrder(
		backgroundProcess.EXPECT().Run().Do(func() {
			executed.Store(true)
		}),
		backgroundProcess.EXPECT().Stop(),
	)

	m.RegisterBackgroundProcess(backgroundProcess)

	require.Equal(t, errMediatorNotOpen, m.Close())

	require.NoError(t, m.Open())
	require.Equal(t, errMediatorAlreadyOpen, m.Open())

	xclock.WaitUntil(func() bool {
		return executed.Load() && reported.Load()
	}, time.Second)
	require.True(t, executed.Load(), "failed to execute")
	require.True(t, reported.Load(), "failed to report")

	require.NoError(t, m.Close())
	require.Equal(t, errMediatorAlreadyClosed, m.Close())
}

func TestDatabaseMediatorDisableFileOpsAndWait(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := DefaultTestOptions().SetRepairEnabled(false)
	now := time.Now()
	opts = opts.
		SetBootstrapProcessProvider(nil).
		SetClockOptions(opts.ClockOptions().SetNowFn(func() time.Time {
			return now
		}))

	db := NewMockdatabase(ctrl)
	db.EXPECT().Options().Return(opts).AnyTimes()
	med, err := newMediator(db, nil, opts)
	require.NoError(t, err)

	m := med.(*mediator)
	fsm := NewMockdatabaseFileSystemManager(ctrl)
	m.databaseFileSystemManager = fsm
	var slept []time.Duration
	m.sleepFn = func(d time.Duration) { slept = append(slept, d) }

	gomock.InOrder(
		fsm.EXPECT().Disable().Return(fileOpInProgress),
		fsm.EXPECT().Status().Return(fileOpInProgress),
		fsm.EXPECT().Status().Return(fileOpInProgress),
		fsm.EXPECT().Status().Return(fileOpNotStarted),
	)

	m.DisableFileOpsAndWait()
	require.Equal(t, 3, len(slept))
}
