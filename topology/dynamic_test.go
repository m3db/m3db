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
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/m3db/m3cluster/client"
	"github.com/m3db/m3cluster/services"
	"github.com/m3db/m3cluster/shard"
	"github.com/m3db/m3x/retry"
	"github.com/m3db/m3x/watch"
	"github.com/stretchr/testify/assert"
)

func testSetup(ctrl *gomock.Controller) (DynamicOptions, xwatch.Watch) {
	opts := NewDynamicOptions()
	opts = opts.RetryOptions(xretry.NewOptions().InitialBackoff(time.Millisecond))

	watch := newTestWatch(time.Millisecond, time.Millisecond, 10, 10)
	mockCSServices := services.NewMockServices(ctrl)
	mockCSServices.EXPECT().WatchInstances(opts.GetService(), opts.GetQueryOptions()).Return(watch, nil)

	mockCSClient := client.NewMockClient(ctrl)
	mockCSClient.EXPECT().Services().Return(mockCSServices)
	opts = opts.ConfigServiceClient(mockCSClient)
	return opts, watch
}

func TestInitTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts, _ := testSetup(ctrl)
	topo, err := newDynamicTopology(opts.InitTimeout(10 * time.Millisecond))
	assert.Equal(t, errInitTimeOut, err)
	assert.Nil(t, topo)
}

func TestInitNoTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts, watch := testSetup(ctrl)
	go watch.(*testWatch).run()
	topo, err := newDynamicTopology(opts)

	assert.NoError(t, err)
	assert.NotNil(t, topo)
	topo.Close()
	// safe to close again
	topo.Close()
}

func TestBackoffPoll(t *testing.T) {
	opts := NewDynamicOptions().RetryOptions(xretry.NewOptions().InitialBackoff(time.Millisecond))
	w := newTestWatch(time.Millisecond, time.Millisecond, 10, 10)
	close(w.(*testWatch).ch)
	input := newServiceTopologyInput(w, opts.GetHashGen(), opts.GetRetryOptions())
	data, err := input.Poll()
	assert.Equal(t, xwatch.ErrSourceClosed, err)
	assert.Nil(t, data)

	w = newTestWatch(time.Millisecond, time.Millisecond, 0, 10)
	go w.(*testWatch).run()
	input = newServiceTopologyInput(w, opts.GetHashGen(), opts.GetRetryOptions())
	data, err = input.Poll()
	assert.Error(t, err)
	assert.Nil(t, data)

	w = newTestWatch(time.Millisecond, time.Millisecond, 10, 10)
	go w.(*testWatch).run()
	input = newServiceTopologyInput(w, opts.GetHashGen(), opts.GetRetryOptions())
	data, err = input.Poll()
	assert.NoError(t, err)
	assert.Equal(t, 2, data.(Map).Replicas())
	assert.Equal(t, 3, data.(Map).HostsLen())
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts, watch := testSetup(ctrl)
	go watch.(*testWatch).run()
	topo, err := newDynamicTopology(opts)
	assert.NoError(t, err)

	m := topo.Get()
	assert.Equal(t, 2, m.Replicas())
}

func TestGetAndSubscribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts, watch := testSetup(ctrl)
	go watch.(*testWatch).run()
	topo, err := newDynamicTopology(opts)
	assert.NoError(t, err)

	m, w, err := topo.GetAndSubscribe()
	assert.Equal(t, 2, m.Replicas())
	assert.Equal(t, 2, w.Get().(Map).Replicas())

	for _ = range w.C() {
		assert.Equal(t, 2, w.Get().(Map).Replicas())
	}
}

func TestGetUniqueShardsAndReplicas(t *testing.T) {
	shards, replica := uniqueShardsAndReplicas(goodCase)
	assert.Equal(t, 3, len(shards))
	assert.Equal(t, 2, replica)
}

type testWatch struct {
	sync.RWMutex

	data                  interface{}
	firstDelay, nextDelay time.Duration
	errAfter, closeAfter  int
	currentCalled         int
	ch                    chan struct{}
	closed                bool
}

func newTestWatch(firstDelay, nextDelay time.Duration, errAfter, closeAfter int) xwatch.Watch {
	w := testWatch{firstDelay: firstDelay, nextDelay: nextDelay, errAfter: errAfter, closeAfter: closeAfter}
	w.ch = make(chan struct{})
	return &w
}

func (w *testWatch) run() {
	time.Sleep(w.firstDelay)
	w.update()
	for w.currentCalled < w.closeAfter {
		time.Sleep(w.nextDelay)
		w.update()
	}
	close(w.ch)
}

func (w *testWatch) update() {
	w.Lock()
	defer w.Unlock()
	if w.currentCalled < w.errAfter {
		w.data = goodCase
	} else {
		w.data = nil
	}
	w.ch <- struct{}{}
	w.currentCalled++
}

func (w *testWatch) Close() {
}

func (w *testWatch) Get() interface{} {
	w.RLock()
	defer w.RUnlock()
	return w.data
}

func (w *testWatch) C() <-chan struct{} {
	return w.ch
}

var (
	i1 = services.NewServiceInstance().SetShards(shard.NewShards(
		[]shard.Shard{
			shard.NewShard(0),
			shard.NewShard(1),
		})).SetID("h1").SetEndpoint("h1:9000")

	i2 = services.NewServiceInstance().SetShards(shard.NewShards(
		[]shard.Shard{
			shard.NewShard(1),
			shard.NewShard(2),
		})).SetID("h2").SetEndpoint("h2:9000")

	i3 = services.NewServiceInstance().SetShards(shard.NewShards(
		[]shard.Shard{
			shard.NewShard(2),
			shard.NewShard(0),
		})).SetID("h3").SetEndpoint("h3:9000")

	goodCase = []services.ServiceInstance{i1, i2, i3}
)
