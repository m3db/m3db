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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: /Users/r/go/src/github.com/m3db/m3db/client/types.go

package client

import (
	gomock "github.com/golang/mock/gomock"
	clock "github.com/m3db/m3db/clock"
	encoding "github.com/m3db/m3db/encoding"
	rpc "github.com/m3db/m3db/generated/thrift/rpc"
	instrument "github.com/m3db/m3db/instrument"
	pool "github.com/m3db/m3db/pool"
	topology "github.com/m3db/m3db/topology"
	time0 "github.com/m3db/m3x/time"
	tchannel_go "github.com/uber/tchannel-go"
	time "time"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) NewSession() (Session, error) {
	ret := _m.ctrl.Call(_m, "NewSession")
	ret0, _ := ret[0].(Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) NewSession() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewSession")
}

// Mock of Session interface
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *_MockSessionRecorder
}

// Recorder for MockSession (not exported)
type _MockSessionRecorder struct {
	mock *MockSession
}

func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &_MockSessionRecorder{mock}
	return mock
}

func (_m *MockSession) EXPECT() *_MockSessionRecorder {
	return _m.recorder
}

func (_m *MockSession) Write(id string, t time.Time, value float64, unit time0.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", id, t, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSessionRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockSession) Fetch(id string, startInclusive time.Time, endExclusive time.Time) (encoding.SeriesIterator, error) {
	ret := _m.ctrl.Call(_m, "Fetch", id, startInclusive, endExclusive)
	ret0, _ := ret[0].(encoding.SeriesIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSessionRecorder) Fetch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fetch", arg0, arg1, arg2)
}

func (_m *MockSession) FetchAll(ids []string, startInclusive time.Time, endExclusive time.Time) (encoding.SeriesIterators, error) {
	ret := _m.ctrl.Call(_m, "FetchAll", ids, startInclusive, endExclusive)
	ret0, _ := ret[0].(encoding.SeriesIterators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSessionRecorder) FetchAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchAll", arg0, arg1, arg2)
}

func (_m *MockSession) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSessionRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of clientSession interface
type MockclientSession struct {
	ctrl     *gomock.Controller
	recorder *_MockclientSessionRecorder
}

// Recorder for MockclientSession (not exported)
type _MockclientSessionRecorder struct {
	mock *MockclientSession
}

func NewMockclientSession(ctrl *gomock.Controller) *MockclientSession {
	mock := &MockclientSession{ctrl: ctrl}
	mock.recorder = &_MockclientSessionRecorder{mock}
	return mock
}

func (_m *MockclientSession) EXPECT() *_MockclientSessionRecorder {
	return _m.recorder
}

func (_m *MockclientSession) Write(id string, t time.Time, value float64, unit time0.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", id, t, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockclientSessionRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockclientSession) Fetch(id string, startInclusive time.Time, endExclusive time.Time) (encoding.SeriesIterator, error) {
	ret := _m.ctrl.Call(_m, "Fetch", id, startInclusive, endExclusive)
	ret0, _ := ret[0].(encoding.SeriesIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockclientSessionRecorder) Fetch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fetch", arg0, arg1, arg2)
}

func (_m *MockclientSession) FetchAll(ids []string, startInclusive time.Time, endExclusive time.Time) (encoding.SeriesIterators, error) {
	ret := _m.ctrl.Call(_m, "FetchAll", ids, startInclusive, endExclusive)
	ret0, _ := ret[0].(encoding.SeriesIterators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockclientSessionRecorder) FetchAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchAll", arg0, arg1, arg2)
}

func (_m *MockclientSession) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockclientSessionRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockclientSession) Open() error {
	ret := _m.ctrl.Call(_m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockclientSessionRecorder) Open() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open")
}

// Mock of hostQueue interface
type MockhostQueue struct {
	ctrl     *gomock.Controller
	recorder *_MockhostQueueRecorder
}

// Recorder for MockhostQueue (not exported)
type _MockhostQueueRecorder struct {
	mock *MockhostQueue
}

func NewMockhostQueue(ctrl *gomock.Controller) *MockhostQueue {
	mock := &MockhostQueue{ctrl: ctrl}
	mock.recorder = &_MockhostQueueRecorder{mock}
	return mock
}

func (_m *MockhostQueue) EXPECT() *_MockhostQueueRecorder {
	return _m.recorder
}

func (_m *MockhostQueue) Open() {
	_m.ctrl.Call(_m, "Open")
}

func (_mr *_MockhostQueueRecorder) Open() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open")
}

func (_m *MockhostQueue) Len() int {
	ret := _m.ctrl.Call(_m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockhostQueueRecorder) Len() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Len")
}

func (_m *MockhostQueue) Enqueue(op op) error {
	ret := _m.ctrl.Call(_m, "Enqueue", op)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockhostQueueRecorder) Enqueue(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Enqueue", arg0)
}

func (_m *MockhostQueue) GetConnectionCount() int {
	ret := _m.ctrl.Call(_m, "GetConnectionCount")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockhostQueueRecorder) GetConnectionCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetConnectionCount")
}

func (_m *MockhostQueue) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockhostQueueRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of connectionPool interface
type MockconnectionPool struct {
	ctrl     *gomock.Controller
	recorder *_MockconnectionPoolRecorder
}

// Recorder for MockconnectionPool (not exported)
type _MockconnectionPoolRecorder struct {
	mock *MockconnectionPool
}

func NewMockconnectionPool(ctrl *gomock.Controller) *MockconnectionPool {
	mock := &MockconnectionPool{ctrl: ctrl}
	mock.recorder = &_MockconnectionPoolRecorder{mock}
	return mock
}

func (_m *MockconnectionPool) EXPECT() *_MockconnectionPoolRecorder {
	return _m.recorder
}

func (_m *MockconnectionPool) Open() {
	_m.ctrl.Call(_m, "Open")
}

func (_mr *_MockconnectionPoolRecorder) Open() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open")
}

func (_m *MockconnectionPool) GetConnectionCount() int {
	ret := _m.ctrl.Call(_m, "GetConnectionCount")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockconnectionPoolRecorder) GetConnectionCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetConnectionCount")
}

func (_m *MockconnectionPool) NextClient() (rpc.TChanNode, error) {
	ret := _m.ctrl.Call(_m, "NextClient")
	ret0, _ := ret[0].(rpc.TChanNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockconnectionPoolRecorder) NextClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NextClient")
}

func (_m *MockconnectionPool) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockconnectionPoolRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of op interface
type Mockop struct {
	ctrl     *gomock.Controller
	recorder *_MockopRecorder
}

// Recorder for Mockop (not exported)
type _MockopRecorder struct {
	mock *Mockop
}

func NewMockop(ctrl *gomock.Controller) *Mockop {
	mock := &Mockop{ctrl: ctrl}
	mock.recorder = &_MockopRecorder{mock}
	return mock
}

func (_m *Mockop) EXPECT() *_MockopRecorder {
	return _m.recorder
}

func (_m *Mockop) Size() int {
	ret := _m.ctrl.Call(_m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockopRecorder) Size() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Size")
}

func (_m *Mockop) GetCompletionFn() completionFn {
	ret := _m.ctrl.Call(_m, "GetCompletionFn")
	ret0, _ := ret[0].(completionFn)
	return ret0
}

func (_mr *_MockopRecorder) GetCompletionFn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCompletionFn")
}

// Mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *_MockOptionsRecorder
}

// Recorder for MockOptions (not exported)
type _MockOptionsRecorder struct {
	mock *MockOptions
}

func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &_MockOptionsRecorder{mock}
	return mock
}

func (_m *MockOptions) EXPECT() *_MockOptionsRecorder {
	return _m.recorder
}

func (_m *MockOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOptionsRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Validate")
}

func (_m *MockOptions) ClockOptions(value clock.Options) Options {
	ret := _m.ctrl.Call(_m, "ClockOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ClockOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClockOptions", arg0)
}

func (_m *MockOptions) GetClockOptions() clock.Options {
	ret := _m.ctrl.Call(_m, "GetClockOptions")
	ret0, _ := ret[0].(clock.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetClockOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetClockOptions")
}

func (_m *MockOptions) InstrumentOptions(value instrument.Options) Options {
	ret := _m.ctrl.Call(_m, "InstrumentOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) InstrumentOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstrumentOptions", arg0)
}

func (_m *MockOptions) GetInstrumentOptions() instrument.Options {
	ret := _m.ctrl.Call(_m, "GetInstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetInstrumentOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetInstrumentOptions")
}

func (_m *MockOptions) EncodingTsz() Options {
	ret := _m.ctrl.Call(_m, "EncodingTsz")
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) EncodingTsz() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EncodingTsz")
}

func (_m *MockOptions) TopologyType(value topology.TopologyType) Options {
	ret := _m.ctrl.Call(_m, "TopologyType", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) TopologyType(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TopologyType", arg0)
}

func (_m *MockOptions) GetTopologyType() topology.TopologyType {
	ret := _m.ctrl.Call(_m, "GetTopologyType")
	ret0, _ := ret[0].(topology.TopologyType)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetTopologyType() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTopologyType")
}

func (_m *MockOptions) ConsistencyLevel(value topology.ConsistencyLevel) Options {
	ret := _m.ctrl.Call(_m, "ConsistencyLevel", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ConsistencyLevel(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConsistencyLevel", arg0)
}

func (_m *MockOptions) GetConsistencyLevel() topology.ConsistencyLevel {
	ret := _m.ctrl.Call(_m, "GetConsistencyLevel")
	ret0, _ := ret[0].(topology.ConsistencyLevel)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetConsistencyLevel() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetConsistencyLevel")
}

func (_m *MockOptions) ChannelOptions(value *tchannel_go.ChannelOptions) Options {
	ret := _m.ctrl.Call(_m, "ChannelOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ChannelOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ChannelOptions", arg0)
}

func (_m *MockOptions) GetChannelOptions() *tchannel_go.ChannelOptions {
	ret := _m.ctrl.Call(_m, "GetChannelOptions")
	ret0, _ := ret[0].(*tchannel_go.ChannelOptions)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetChannelOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetChannelOptions")
}

func (_m *MockOptions) MaxConnectionCount(value int) Options {
	ret := _m.ctrl.Call(_m, "MaxConnectionCount", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) MaxConnectionCount(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MaxConnectionCount", arg0)
}

func (_m *MockOptions) GetMaxConnectionCount() int {
	ret := _m.ctrl.Call(_m, "GetMaxConnectionCount")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetMaxConnectionCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMaxConnectionCount")
}

func (_m *MockOptions) MinConnectionCount(value int) Options {
	ret := _m.ctrl.Call(_m, "MinConnectionCount", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) MinConnectionCount(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MinConnectionCount", arg0)
}

func (_m *MockOptions) GetMinConnectionCount() int {
	ret := _m.ctrl.Call(_m, "GetMinConnectionCount")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetMinConnectionCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMinConnectionCount")
}

func (_m *MockOptions) HostConnectTimeout(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "HostConnectTimeout", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) HostConnectTimeout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HostConnectTimeout", arg0)
}

func (_m *MockOptions) GetHostConnectTimeout() time.Duration {
	ret := _m.ctrl.Call(_m, "GetHostConnectTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetHostConnectTimeout() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetHostConnectTimeout")
}

func (_m *MockOptions) ClusterConnectTimeout(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "ClusterConnectTimeout", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ClusterConnectTimeout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClusterConnectTimeout", arg0)
}

func (_m *MockOptions) GetClusterConnectTimeout() time.Duration {
	ret := _m.ctrl.Call(_m, "GetClusterConnectTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetClusterConnectTimeout() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetClusterConnectTimeout")
}

func (_m *MockOptions) ClusterConnectConsistencyLevel(value topology.ConsistencyLevel) Options {
	ret := _m.ctrl.Call(_m, "ClusterConnectConsistencyLevel", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ClusterConnectConsistencyLevel(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClusterConnectConsistencyLevel", arg0)
}

func (_m *MockOptions) GetClusterConnectConsistencyLevel() topology.ConsistencyLevel {
	ret := _m.ctrl.Call(_m, "GetClusterConnectConsistencyLevel")
	ret0, _ := ret[0].(topology.ConsistencyLevel)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetClusterConnectConsistencyLevel() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetClusterConnectConsistencyLevel")
}

func (_m *MockOptions) WriteRequestTimeout(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "WriteRequestTimeout", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) WriteRequestTimeout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteRequestTimeout", arg0)
}

func (_m *MockOptions) GetWriteRequestTimeout() time.Duration {
	ret := _m.ctrl.Call(_m, "GetWriteRequestTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetWriteRequestTimeout() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWriteRequestTimeout")
}

func (_m *MockOptions) FetchRequestTimeout(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "FetchRequestTimeout", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) FetchRequestTimeout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchRequestTimeout", arg0)
}

func (_m *MockOptions) GetFetchRequestTimeout() time.Duration {
	ret := _m.ctrl.Call(_m, "GetFetchRequestTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetFetchRequestTimeout() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFetchRequestTimeout")
}

func (_m *MockOptions) BackgroundConnectInterval(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "BackgroundConnectInterval", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) BackgroundConnectInterval(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BackgroundConnectInterval", arg0)
}

func (_m *MockOptions) GetBackgroundConnectInterval() time.Duration {
	ret := _m.ctrl.Call(_m, "GetBackgroundConnectInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetBackgroundConnectInterval() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBackgroundConnectInterval")
}

func (_m *MockOptions) BackgroundConnectStutter(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "BackgroundConnectStutter", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) BackgroundConnectStutter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BackgroundConnectStutter", arg0)
}

func (_m *MockOptions) GetBackgroundConnectStutter() time.Duration {
	ret := _m.ctrl.Call(_m, "GetBackgroundConnectStutter")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetBackgroundConnectStutter() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBackgroundConnectStutter")
}

func (_m *MockOptions) BackgroundHealthCheckInterval(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "BackgroundHealthCheckInterval", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) BackgroundHealthCheckInterval(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BackgroundHealthCheckInterval", arg0)
}

func (_m *MockOptions) GetBackgroundHealthCheckInterval() time.Duration {
	ret := _m.ctrl.Call(_m, "GetBackgroundHealthCheckInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetBackgroundHealthCheckInterval() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBackgroundHealthCheckInterval")
}

func (_m *MockOptions) BackgroundHealthCheckStutter(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "BackgroundHealthCheckStutter", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) BackgroundHealthCheckStutter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BackgroundHealthCheckStutter", arg0)
}

func (_m *MockOptions) GetBackgroundHealthCheckStutter() time.Duration {
	ret := _m.ctrl.Call(_m, "GetBackgroundHealthCheckStutter")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetBackgroundHealthCheckStutter() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBackgroundHealthCheckStutter")
}

func (_m *MockOptions) WriteOpPoolSize(value int) Options {
	ret := _m.ctrl.Call(_m, "WriteOpPoolSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) WriteOpPoolSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteOpPoolSize", arg0)
}

func (_m *MockOptions) GetWriteOpPoolSize() int {
	ret := _m.ctrl.Call(_m, "GetWriteOpPoolSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetWriteOpPoolSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWriteOpPoolSize")
}

func (_m *MockOptions) FetchBatchOpPoolSize(value int) Options {
	ret := _m.ctrl.Call(_m, "FetchBatchOpPoolSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) FetchBatchOpPoolSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBatchOpPoolSize", arg0)
}

func (_m *MockOptions) GetFetchBatchOpPoolSize() int {
	ret := _m.ctrl.Call(_m, "GetFetchBatchOpPoolSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetFetchBatchOpPoolSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFetchBatchOpPoolSize")
}

func (_m *MockOptions) WriteBatchSize(value int) Options {
	ret := _m.ctrl.Call(_m, "WriteBatchSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) WriteBatchSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteBatchSize", arg0)
}

func (_m *MockOptions) GetWriteBatchSize() int {
	ret := _m.ctrl.Call(_m, "GetWriteBatchSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetWriteBatchSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWriteBatchSize")
}

func (_m *MockOptions) FetchBatchSize(value int) Options {
	ret := _m.ctrl.Call(_m, "FetchBatchSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) FetchBatchSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBatchSize", arg0)
}

func (_m *MockOptions) GetFetchBatchSize() int {
	ret := _m.ctrl.Call(_m, "GetFetchBatchSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetFetchBatchSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFetchBatchSize")
}

func (_m *MockOptions) HostQueueOpsFlushSize(value int) Options {
	ret := _m.ctrl.Call(_m, "HostQueueOpsFlushSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) HostQueueOpsFlushSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HostQueueOpsFlushSize", arg0)
}

func (_m *MockOptions) GetHostQueueOpsFlushSize() int {
	ret := _m.ctrl.Call(_m, "GetHostQueueOpsFlushSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetHostQueueOpsFlushSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetHostQueueOpsFlushSize")
}

func (_m *MockOptions) HostQueueOpsFlushInterval(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "HostQueueOpsFlushInterval", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) HostQueueOpsFlushInterval(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HostQueueOpsFlushInterval", arg0)
}

func (_m *MockOptions) GetHostQueueOpsFlushInterval() time.Duration {
	ret := _m.ctrl.Call(_m, "GetHostQueueOpsFlushInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetHostQueueOpsFlushInterval() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetHostQueueOpsFlushInterval")
}

func (_m *MockOptions) HostQueueOpsArrayPoolSize(value int) Options {
	ret := _m.ctrl.Call(_m, "HostQueueOpsArrayPoolSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) HostQueueOpsArrayPoolSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HostQueueOpsArrayPoolSize", arg0)
}

func (_m *MockOptions) GetHostQueueOpsArrayPoolSize() int {
	ret := _m.ctrl.Call(_m, "GetHostQueueOpsArrayPoolSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetHostQueueOpsArrayPoolSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetHostQueueOpsArrayPoolSize")
}

func (_m *MockOptions) SeriesIteratorPoolSize(value int) Options {
	ret := _m.ctrl.Call(_m, "SeriesIteratorPoolSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SeriesIteratorPoolSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SeriesIteratorPoolSize", arg0)
}

func (_m *MockOptions) GetSeriesIteratorPoolSize() int {
	ret := _m.ctrl.Call(_m, "GetSeriesIteratorPoolSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetSeriesIteratorPoolSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSeriesIteratorPoolSize")
}

func (_m *MockOptions) SeriesIteratorArrayPoolBuckets(value []pool.PoolBucket) Options {
	ret := _m.ctrl.Call(_m, "SeriesIteratorArrayPoolBuckets", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SeriesIteratorArrayPoolBuckets(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SeriesIteratorArrayPoolBuckets", arg0)
}

func (_m *MockOptions) GetSeriesIteratorArrayPoolBuckets() []pool.PoolBucket {
	ret := _m.ctrl.Call(_m, "GetSeriesIteratorArrayPoolBuckets")
	ret0, _ := ret[0].([]pool.PoolBucket)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetSeriesIteratorArrayPoolBuckets() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSeriesIteratorArrayPoolBuckets")
}

func (_m *MockOptions) ReaderIteratorAllocate(value encoding.ReaderIteratorAllocate) Options {
	ret := _m.ctrl.Call(_m, "ReaderIteratorAllocate", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ReaderIteratorAllocate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReaderIteratorAllocate", arg0)
}

func (_m *MockOptions) GetReaderIteratorAllocate() encoding.ReaderIteratorAllocate {
	ret := _m.ctrl.Call(_m, "GetReaderIteratorAllocate")
	ret0, _ := ret[0].(encoding.ReaderIteratorAllocate)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetReaderIteratorAllocate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetReaderIteratorAllocate")
}
