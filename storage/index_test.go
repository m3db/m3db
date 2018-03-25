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

package storage

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/storage/index"
	"github.com/m3db/m3ninx/doc"
	"github.com/m3db/m3ninx/index/segment"
	"github.com/m3db/m3x/context"
	"github.com/m3db/m3x/ident"

	"github.com/fortytw2/leaktest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/uber-go/tally"
)

func testDatabaseIndexOptions() index.Options {
	return index.NewOptions()
}

func newTestDatabaseIndex(t *testing.T, ctrl *gomock.Controller) (databaseIndex, *MockdatabaseIndexInsertQueue) {
	q := NewMockdatabaseIndexInsertQueue(ctrl)
	newFn := func(fn dbIndexInsertBatchFn, nowFn clock.NowFn, s tally.Scope) databaseIndexInsertQueue {
		return q
	}
	q.EXPECT().Start().Return(nil)
	idx, err := newDatabaseIndex(newFn, testDatabaseIndexOptions())
	assert.NoError(t, err)
	return idx, q
}

func TestDatabaseIndexHappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	q := NewMockdatabaseIndexInsertQueue(ctrl)
	newFn := func(fn dbIndexInsertBatchFn, nowFn clock.NowFn, s tally.Scope) databaseIndexInsertQueue {
		return q
	}
	q.EXPECT().Start().Return(nil)

	idx, err := newDatabaseIndex(newFn, testDatabaseIndexOptions())
	assert.NoError(t, err)
	assert.NotNil(t, idx)

	q.EXPECT().Stop().Return(nil)
	assert.NoError(t, idx.Close())
}

func TestDatabaseIndexStartErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	q := NewMockdatabaseIndexInsertQueue(ctrl)
	newFn := func(fn dbIndexInsertBatchFn, nowFn clock.NowFn, s tally.Scope) databaseIndexInsertQueue {
		return q
	}
	q.EXPECT().Start().Return(fmt.Errorf("random err"))

	idx, err := newDatabaseIndex(newFn, testDatabaseIndexOptions())
	assert.Error(t, err)
	assert.Nil(t, idx)
}

func TestDatabaseIndexStopErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	q := NewMockdatabaseIndexInsertQueue(ctrl)
	newFn := func(fn dbIndexInsertBatchFn, nowFn clock.NowFn, s tally.Scope) databaseIndexInsertQueue {
		return q
	}
	q.EXPECT().Start().Return(nil)

	idx, err := newDatabaseIndex(newFn, testDatabaseIndexOptions())
	assert.NoError(t, err)
	assert.NotNil(t, idx)

	q.EXPECT().Stop().Return(fmt.Errorf("random err"))
	assert.Error(t, idx.Close())
}

func TestDatabaseIndexInvalidDocConversion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbIdx, _ := newTestDatabaseIndex(t, ctrl)
	idx, ok := dbIdx.(*dbIndex)
	assert.True(t, ok)

	id := ident.StringID("foo")
	ns := ident.StringID("bar")
	tags := ident.Tags{
		ident.StringTag(string(index.ReservedFieldNameNamespace), "value"),
	}

	_, err := idx.doc(ns, id, tags)
	assert.Error(t, err)
}

func TestDatabaseIndexInvalidDocWrite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbIdx, _ := newTestDatabaseIndex(t, ctrl)
	idx, ok := dbIdx.(*dbIndex)
	assert.True(t, ok)

	id := ident.StringID("foo")
	ns := ident.StringID("bar")
	tags := ident.Tags{
		ident.StringTag(string(index.ReservedFieldNameNamespace), "value"),
	}

	lifecycle := &testLifecycleHooks{}
	assert.Error(t, idx.Write(ns, id, tags, lifecycle))

	// ensure lifecycle is finalized despite failure
	lifecycle.Lock()
	defer lifecycle.Unlock()
	assert.True(t, lifecycle.finalized)
}

func TestDatabaseIndexWriteAfterClose(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbIdx, q := newTestDatabaseIndex(t, ctrl)
	idx, ok := dbIdx.(*dbIndex)
	assert.True(t, ok)

	id := ident.StringID("foo")
	ns := ident.StringID("bar")
	tags := ident.Tags{
		ident.StringTag("name", "value"),
	}

	q.EXPECT().Stop().Return(nil)
	assert.NoError(t, idx.Close())

	lifecycle := &testLifecycleHooks{}
	assert.Error(t, idx.Write(ns, id, tags, lifecycle))

	// ensure lifecycle is finalized despite failure
	lifecycle.Lock()
	defer lifecycle.Unlock()
	assert.True(t, lifecycle.finalized)
}

func TestDatabaseIndexWriteQueueError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbIdx, q := newTestDatabaseIndex(t, ctrl)
	idx, ok := dbIdx.(*dbIndex)
	assert.True(t, ok)

	id := ident.StringID("foo")
	ns := ident.StringID("bar")
	tags := ident.Tags{
		ident.StringTag("name", "value"),
	}

	lifecycle := &testLifecycleHooks{}
	q.EXPECT().
		Insert(gomock.Any(), lifecycle).
		Return(nil, fmt.Errorf("random err"))
	assert.Error(t, idx.Write(ns, id, tags, lifecycle))

	// ensure lifecycle is finalized despite failure
	lifecycle.Lock()
	defer lifecycle.Unlock()
	assert.True(t, lifecycle.finalized)
}

func TestDatabaseIndexDocConversion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbIdx, _ := newTestDatabaseIndex(t, ctrl)
	idx, ok := dbIdx.(*dbIndex)
	assert.True(t, ok)

	id := ident.StringID("foo")
	ns := ident.StringID("bar")
	tags := ident.Tags{
		ident.StringTag("name", "value"),
	}

	d, err := idx.doc(ns, id, tags)
	assert.NoError(t, err)
	assert.Equal(t, "foo", string(d.ID))
	assert.Len(t, d.Fields, 2)
	assert.Equal(t, index.ReservedFieldNameNamespace, d.Fields[0].Name)
	assert.Equal(t, "bar", string(d.Fields[0].Value))
	assert.Equal(t, "name", string(d.Fields[1].Name))
	assert.Equal(t, "value", string(d.Fields[1].Value))
}

func TestDatabaseIndexInsertQueueInteraction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbIdx, q := newTestDatabaseIndex(t, ctrl)
	idx, ok := dbIdx.(*dbIndex)
	assert.True(t, ok)

	var (
		id   = ident.StringID("foo")
		ns   = ident.StringID("bar")
		tags = ident.Tags{
			ident.StringTag("name", "value"),
		}
	)

	d, err := idx.doc(ns, id, tags)
	assert.NoError(t, err)

	var wg sync.WaitGroup
	lifecycleFns := &testLifecycleHooks{}
	q.EXPECT().Insert(docMatcher{d}, gomock.Any()).Return(&wg, nil)
	assert.NoError(t, idx.Write(ns, id, tags, lifecycleFns))
}

func TestDatabaseIndexInsertQuery(t *testing.T) {
	defer leaktest.CheckTimeout(t, 2*time.Second)()

	newFn := func(fn dbIndexInsertBatchFn, nowFn clock.NowFn, s tally.Scope) databaseIndexInsertQueue {
		q := newDatabaseIndexInsertQueue(fn, nowFn, s)
		q.(*dbIndexInsertQueue).indexBatchBackoff = 10 * time.Millisecond
		return q
	}
	idx, err := newDatabaseIndex(newFn, testDatabaseIndexOptions())
	assert.NoError(t, err)
	defer idx.Close()

	var (
		id   = ident.StringID("foo")
		ns   = ident.StringID("bar")
		tags = ident.Tags{
			ident.StringTag("name", "value"),
		}
		ctx          = context.NewContext()
		lifecycleFns = &testLifecycleHooks{}
	)
	// make insert mode sync for tests
	idx.(*dbIndex).insertMode = dbIndexInsertSync
	assert.NoError(t, idx.Write(ns, id, tags, lifecycleFns))

	res, err := idx.Query(ctx, index.Query{
		segment.Query{
			Conjunction: segment.AndConjunction,
			Filters: []segment.Filter{
				segment.Filter{
					FieldName:        []byte("name"),
					FieldValueFilter: []byte("val.*"),
					Regexp:           true,
				},
			},
		},
	}, index.QueryOptions{})
	assert.NoError(t, err)

	assert.True(t, res.Exhaustive)
	iter := res.Iterator
	assert.True(t, iter.Next())

	cNs, cID, cTags := iter.Current()
	assert.Equal(t, "foo", cID.String())
	assert.Equal(t, "bar", cNs.String())
	assert.Len(t, cTags, 1)
	assert.Equal(t, "name", cTags[0].Name.String())
	assert.Equal(t, "value", cTags[0].Value.String())
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

type docMatcher struct{ d doc.Document }

func (dm docMatcher) Matches(x interface{}) bool {
	other, ok := x.(doc.Document)
	if !ok {
		return false
	}
	if !bytes.Equal(dm.d.ID, other.ID) {
		return false
	}
	if len(dm.d.Fields) != len(other.Fields) {
		return false
	}
	for i := range dm.d.Fields {
		if !bytes.Equal(dm.d.Fields[i].Name, other.Fields[i].Name) {
			return false
		}
		if !bytes.Equal(dm.d.Fields[i].Value, other.Fields[i].Value) {
			return false
		}
	}
	return true
}

func (dm docMatcher) String() string {
	return fmt.Sprintf("doc is %+v", dm.d)
}

type testLifecycleHooks struct {
	sync.Mutex
	writeTime time.Time
	finalized bool
}

func (t *testLifecycleHooks) OnIndexSuccess(ts time.Time) {
	t.Lock()
	t.writeTime = ts
	t.Unlock()
}

func (t *testLifecycleHooks) OnIndexFinalize() {
	t.Lock()
	if t.finalized {
		// fine to do as it's only used during tests
		panic("already finalized")
	}
	t.finalized = true
	t.Unlock()
}
