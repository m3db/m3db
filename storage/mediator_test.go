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

	"github.com/m3db/m3db/storage/bootstrap"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDatabaseMediatorOpenClose(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := testDatabaseOptions().SetRepairEnabled(false)
	now := time.Now()
	opts = opts.SetNewBootstrapFn(func() bootstrap.Bootstrap {
		return nil
	}).SetClockOptions(opts.ClockOptions().SetNowFn(func() time.Time {
		return now
	}))

	db := &mockDatabase{opts: opts}
	med, err := newMediator(db, opts)
	require.NoError(t, err)

	m := med.(*mediator)
	tm := NewMockdatabaseTickManager(ctrl)
	fsm := NewMockdatabaseFileSystemManager(ctrl)
	m.databaseTickManager = tm
	m.databaseFileSystemManager = fsm

	deadline := opts.RetentionOptions().BufferDrain()
	tm.EXPECT().Tick(deadline, false).Return(true).AnyTimes()
	fsm.EXPECT().Run(now, true, false).AnyTimes()

	require.Equal(t, errMediatorNotOpen, m.Close())

	require.NoError(t, m.Open())
	require.Equal(t, errMediatorAlreadyOpen, m.Open())

	require.NoError(t, m.Close())
	require.Equal(t, errMediatorAlreadyClosed, m.Close())
}

func TestDatabaseBootstrapWithFileOpInProgess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := testDatabaseOptions().SetRepairEnabled(false)
	now := time.Now()
	opts = opts.SetNewBootstrapFn(func() bootstrap.Bootstrap {
		return nil
	}).SetClockOptions(opts.ClockOptions().SetNowFn(func() time.Time {
		return now
	}))

	db := &mockDatabase{opts: opts}
	med, err := newMediator(db, opts)
	require.NoError(t, err)

	m := med.(*mediator)
	bsm := NewMockdatabaseBootstrapManager(ctrl)
	tm := NewMockdatabaseTickManager(ctrl)
	fsm := NewMockdatabaseFileSystemManager(ctrl)
	m.databaseBootstrapManager = bsm
	m.databaseTickManager = tm
	m.databaseFileSystemManager = fsm
	var slept []time.Duration
	m.sleepFn = func(d time.Duration) { slept = append(slept, d) }

	fsm.EXPECT().Disable().Return(true)
	fsm.EXPECT().Enable().AnyTimes()
	gomock.InOrder(
		fsm.EXPECT().IsRunning().Return(true),
		fsm.EXPECT().IsRunning().Return(true),
		fsm.EXPECT().IsRunning().Return(false),
		fsm.EXPECT().Run(now, false, true),
	)
	bsm.EXPECT().Bootstrap().Return(nil)
	tm.EXPECT().Tick(time.Duration(0), true).Return(true)

	require.Nil(t, m.Bootstrap())
	require.Equal(t, 3, len(slept))
}
