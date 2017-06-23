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

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/uber-go/tally"

	"github.com/m3db/m3db/storage/namespace"
)

func newMultipleFlushManagerNeedsFlush(t *testing.T, ctrl *gomock.Controller) (
	*flushManager,
	*MockdatabaseNamespace,
	*MockdatabaseNamespace,
) {
	options := namespace.NewOptions()
	namespace := NewMockdatabaseNamespace(ctrl)
	namespace.EXPECT().Options().Return(options).AnyTimes()
	otherNamespace := NewMockdatabaseNamespace(ctrl)
	otherNamespace.EXPECT().Options().Return(options).AnyTimes()

	db := newMockDatabase()
	db.namespaces = map[string]databaseNamespace{
		defaultTestNamespaceID.String(): namespace,
		"someString":                    otherNamespace,
	}

	fm := newFlushManager(db, tally.NoopScope).(*flushManager)

	return fm, namespace, otherNamespace
}

func TestFlushManagerNeedsFlushSingle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fm, ns1, ns2 := newMultipleFlushManagerNeedsFlush(t, ctrl)
	now := time.Now()

	// short circuit second ns for this test case
	ns2.EXPECT().NeedsFlush(now, now).Return(false).AnyTimes()

	ns1.EXPECT().NeedsFlush(now, now).Return(true)
	require.True(t, fm.NeedsFlush(now, now))

	ns1.EXPECT().NeedsFlush(now, now).Return(false)
	require.False(t, fm.NeedsFlush(now, now))
}

func TestFlushManagerNeedsFlushMultipleAllFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fm, ns1, ns2 := newMultipleFlushManagerNeedsFlush(t, ctrl)
	now := time.Now()

	ns2.EXPECT().NeedsFlush(now, now).Return(false)
	ns1.EXPECT().NeedsFlush(now, now).Return(false)
	require.False(t, fm.NeedsFlush(now, now))
}

func TestFlushManagerNeedsFlushMultipleAllTrue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fm, ns1, ns2 := newMultipleFlushManagerNeedsFlush(t, ctrl)
	now := time.Now()

	ns2.EXPECT().NeedsFlush(now, now).Return(true).AnyTimes()
	ns1.EXPECT().NeedsFlush(now, now).Return(true).AnyTimes()
	require.True(t, fm.NeedsFlush(now, now))
}

func TestFlushManagerNeedsFlushMultipleSingleTrueA(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fm, ns1, ns2 := newMultipleFlushManagerNeedsFlush(t, ctrl)
	now := time.Now()

	ns1.EXPECT().NeedsFlush(now, now).Return(false).AnyTimes()
	ns2.EXPECT().NeedsFlush(now, now).Return(true)
	require.True(t, fm.NeedsFlush(now, now))
}

func TestFlushManagerNeedsFlushMultipleSingleTrueb(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fm, ns1, ns2 := newMultipleFlushManagerNeedsFlush(t, ctrl)
	now := time.Now()

	ns1.EXPECT().NeedsFlush(now, now).Return(true)
	ns2.EXPECT().NeedsFlush(now, now).Return(false).AnyTimes()
	require.True(t, fm.NeedsFlush(now, now))
}

func TestFlushManagerFlushTimeStart(t *testing.T) {
	inputs := []struct {
		ts       time.Time
		expected time.Time
	}{
		{time.Unix(86400*2, 0), time.Unix(0, 0)},
		{time.Unix(86400*2+7200, 0), time.Unix(7200, 0)},
		{time.Unix(86400*2+10800, 0), time.Unix(7200, 0)},
	}

	database := newMockDatabase()
	fm := newFlushManager(database, tally.NoopScope).(*flushManager)
	for _, input := range inputs {
		start, _ := fm.flushRange(defaultTestRetentionOptions, input.ts)
		require.Equal(t, input.expected, start)
	}
}

func TestFlushManagerFlushTimeEnd(t *testing.T) {
	inputs := []struct {
		ts       time.Time
		expected time.Time
	}{
		{time.Unix(7800, 0), time.Unix(0, 0)},
		{time.Unix(8000, 0), time.Unix(0, 0)},
		{time.Unix(15200, 0), time.Unix(7200, 0)},
	}
	database := newMockDatabase()
	fm := newFlushManager(database, tally.NoopScope).(*flushManager)
	for _, input := range inputs {
		_, end := fm.flushRange(defaultTestRetentionOptions, input.ts)
		require.Equal(t, input.expected, end)
	}
}

func TestFlushManagerNamespaceFlushTimes(t *testing.T) {
	// TODO(prateek): write this test
}
