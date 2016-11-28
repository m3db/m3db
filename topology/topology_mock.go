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
// Source: github.com/m3db/m3db/topology/types.go

package topology

import (
	gomock "github.com/golang/mock/gomock"
	client "github.com/m3db/m3cluster/client"
	services "github.com/m3db/m3cluster/services"
	instrument "github.com/m3db/m3db/instrument"
	sharding "github.com/m3db/m3db/sharding"
	ts "github.com/m3db/m3db/ts"
	time "time"
)

// Mock of Host interface
type MockHost struct {
	ctrl     *gomock.Controller
	recorder *_MockHostRecorder
}

// Recorder for MockHost (not exported)
type _MockHostRecorder struct {
	mock *MockHost
}

func NewMockHost(ctrl *gomock.Controller) *MockHost {
	mock := &MockHost{ctrl: ctrl}
	mock.recorder = &_MockHostRecorder{mock}
	return mock
}

func (_m *MockHost) EXPECT() *_MockHostRecorder {
	return _m.recorder
}

func (_m *MockHost) ID() string {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockHostRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockHost) Address() string {
	ret := _m.ctrl.Call(_m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockHostRecorder) Address() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Address")
}

func (_m *MockHost) String() string {
	ret := _m.ctrl.Call(_m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockHostRecorder) String() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "String")
}

// Mock of HostShardSet interface
type MockHostShardSet struct {
	ctrl     *gomock.Controller
	recorder *_MockHostShardSetRecorder
}

// Recorder for MockHostShardSet (not exported)
type _MockHostShardSetRecorder struct {
	mock *MockHostShardSet
}

func NewMockHostShardSet(ctrl *gomock.Controller) *MockHostShardSet {
	mock := &MockHostShardSet{ctrl: ctrl}
	mock.recorder = &_MockHostShardSetRecorder{mock}
	return mock
}

func (_m *MockHostShardSet) EXPECT() *_MockHostShardSetRecorder {
	return _m.recorder
}

func (_m *MockHostShardSet) Host() Host {
	ret := _m.ctrl.Call(_m, "Host")
	ret0, _ := ret[0].(Host)
	return ret0
}

func (_mr *_MockHostShardSetRecorder) Host() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Host")
}

func (_m *MockHostShardSet) ShardSet() sharding.ShardSet {
	ret := _m.ctrl.Call(_m, "ShardSet")
	ret0, _ := ret[0].(sharding.ShardSet)
	return ret0
}

func (_mr *_MockHostShardSetRecorder) ShardSet() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShardSet")
}

// Mock of Initializer interface
type MockInitializer struct {
	ctrl     *gomock.Controller
	recorder *_MockInitializerRecorder
}

// Recorder for MockInitializer (not exported)
type _MockInitializerRecorder struct {
	mock *MockInitializer
}

func NewMockInitializer(ctrl *gomock.Controller) *MockInitializer {
	mock := &MockInitializer{ctrl: ctrl}
	mock.recorder = &_MockInitializerRecorder{mock}
	return mock
}

func (_m *MockInitializer) EXPECT() *_MockInitializerRecorder {
	return _m.recorder
}

func (_m *MockInitializer) Init() (Topology, error) {
	ret := _m.ctrl.Call(_m, "Init")
	ret0, _ := ret[0].(Topology)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInitializerRecorder) Init() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init")
}

// Mock of Topology interface
type MockTopology struct {
	ctrl     *gomock.Controller
	recorder *_MockTopologyRecorder
}

// Recorder for MockTopology (not exported)
type _MockTopologyRecorder struct {
	mock *MockTopology
}

func NewMockTopology(ctrl *gomock.Controller) *MockTopology {
	mock := &MockTopology{ctrl: ctrl}
	mock.recorder = &_MockTopologyRecorder{mock}
	return mock
}

func (_m *MockTopology) EXPECT() *_MockTopologyRecorder {
	return _m.recorder
}

func (_m *MockTopology) Get() Map {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Map)
	return ret0
}

func (_mr *_MockTopologyRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockTopology) Watch() (MapWatch, error) {
	ret := _m.ctrl.Call(_m, "Watch")
	ret0, _ := ret[0].(MapWatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTopologyRecorder) Watch() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Watch")
}

func (_m *MockTopology) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockTopologyRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of MapWatch interface
type MockMapWatch struct {
	ctrl     *gomock.Controller
	recorder *_MockMapWatchRecorder
}

// Recorder for MockMapWatch (not exported)
type _MockMapWatchRecorder struct {
	mock *MockMapWatch
}

func NewMockMapWatch(ctrl *gomock.Controller) *MockMapWatch {
	mock := &MockMapWatch{ctrl: ctrl}
	mock.recorder = &_MockMapWatchRecorder{mock}
	return mock
}

func (_m *MockMapWatch) EXPECT() *_MockMapWatchRecorder {
	return _m.recorder
}

func (_m *MockMapWatch) C() <-chan struct{} {
	ret := _m.ctrl.Call(_m, "C")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

func (_mr *_MockMapWatchRecorder) C() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "C")
}

func (_m *MockMapWatch) Get() Map {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Map)
	return ret0
}

func (_mr *_MockMapWatchRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockMapWatch) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockMapWatchRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of Map interface
type MockMap struct {
	ctrl     *gomock.Controller
	recorder *_MockMapRecorder
}

// Recorder for MockMap (not exported)
type _MockMapRecorder struct {
	mock *MockMap
}

func NewMockMap(ctrl *gomock.Controller) *MockMap {
	mock := &MockMap{ctrl: ctrl}
	mock.recorder = &_MockMapRecorder{mock}
	return mock
}

func (_m *MockMap) EXPECT() *_MockMapRecorder {
	return _m.recorder
}

func (_m *MockMap) Hosts() []Host {
	ret := _m.ctrl.Call(_m, "Hosts")
	ret0, _ := ret[0].([]Host)
	return ret0
}

func (_mr *_MockMapRecorder) Hosts() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Hosts")
}

func (_m *MockMap) HostShardSets() []HostShardSet {
	ret := _m.ctrl.Call(_m, "HostShardSets")
	ret0, _ := ret[0].([]HostShardSet)
	return ret0
}

func (_mr *_MockMapRecorder) HostShardSets() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HostShardSets")
}

func (_m *MockMap) LookupHostShardSet(hostID string) (HostShardSet, bool) {
	ret := _m.ctrl.Call(_m, "LookupHostShardSet", hostID)
	ret0, _ := ret[0].(HostShardSet)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockMapRecorder) LookupHostShardSet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LookupHostShardSet", arg0)
}

func (_m *MockMap) HostsLen() int {
	ret := _m.ctrl.Call(_m, "HostsLen")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockMapRecorder) HostsLen() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HostsLen")
}

func (_m *MockMap) ShardSet() sharding.ShardSet {
	ret := _m.ctrl.Call(_m, "ShardSet")
	ret0, _ := ret[0].(sharding.ShardSet)
	return ret0
}

func (_mr *_MockMapRecorder) ShardSet() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShardSet")
}

func (_m *MockMap) Route(id ts.ID) (uint32, []Host, error) {
	ret := _m.ctrl.Call(_m, "Route", id)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].([]Host)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockMapRecorder) Route(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Route", arg0)
}

func (_m *MockMap) RouteForEach(id ts.ID, forEachFn RouteForEachFn) error {
	ret := _m.ctrl.Call(_m, "RouteForEach", id, forEachFn)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMapRecorder) RouteForEach(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteForEach", arg0, arg1)
}

func (_m *MockMap) RouteShard(shard uint32) ([]Host, error) {
	ret := _m.ctrl.Call(_m, "RouteShard", shard)
	ret0, _ := ret[0].([]Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMapRecorder) RouteShard(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteShard", arg0)
}

func (_m *MockMap) RouteShardForEach(shard uint32, forEachFn RouteForEachFn) error {
	ret := _m.ctrl.Call(_m, "RouteShardForEach", shard, forEachFn)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMapRecorder) RouteShardForEach(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteShardForEach", arg0, arg1)
}

func (_m *MockMap) Replicas() int {
	ret := _m.ctrl.Call(_m, "Replicas")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockMapRecorder) Replicas() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Replicas")
}

func (_m *MockMap) MajorityReplicas() int {
	ret := _m.ctrl.Call(_m, "MajorityReplicas")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockMapRecorder) MajorityReplicas() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MajorityReplicas")
}

// Mock of StaticOptions interface
type MockStaticOptions struct {
	ctrl     *gomock.Controller
	recorder *_MockStaticOptionsRecorder
}

// Recorder for MockStaticOptions (not exported)
type _MockStaticOptionsRecorder struct {
	mock *MockStaticOptions
}

func NewMockStaticOptions(ctrl *gomock.Controller) *MockStaticOptions {
	mock := &MockStaticOptions{ctrl: ctrl}
	mock.recorder = &_MockStaticOptionsRecorder{mock}
	return mock
}

func (_m *MockStaticOptions) EXPECT() *_MockStaticOptionsRecorder {
	return _m.recorder
}

func (_m *MockStaticOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStaticOptionsRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Validate")
}

func (_m *MockStaticOptions) SetShardSet(value sharding.ShardSet) StaticOptions {
	ret := _m.ctrl.Call(_m, "SetShardSet", value)
	ret0, _ := ret[0].(StaticOptions)
	return ret0
}

func (_mr *_MockStaticOptionsRecorder) SetShardSet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetShardSet", arg0)
}

func (_m *MockStaticOptions) ShardSet() sharding.ShardSet {
	ret := _m.ctrl.Call(_m, "ShardSet")
	ret0, _ := ret[0].(sharding.ShardSet)
	return ret0
}

func (_mr *_MockStaticOptionsRecorder) ShardSet() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShardSet")
}

func (_m *MockStaticOptions) SetReplicas(value int) StaticOptions {
	ret := _m.ctrl.Call(_m, "SetReplicas", value)
	ret0, _ := ret[0].(StaticOptions)
	return ret0
}

func (_mr *_MockStaticOptionsRecorder) SetReplicas(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetReplicas", arg0)
}

func (_m *MockStaticOptions) Replicas() int {
	ret := _m.ctrl.Call(_m, "Replicas")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockStaticOptionsRecorder) Replicas() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Replicas")
}

func (_m *MockStaticOptions) SetHostShardSets(value []HostShardSet) StaticOptions {
	ret := _m.ctrl.Call(_m, "SetHostShardSets", value)
	ret0, _ := ret[0].(StaticOptions)
	return ret0
}

func (_mr *_MockStaticOptionsRecorder) SetHostShardSets(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetHostShardSets", arg0)
}

func (_m *MockStaticOptions) HostShardSets() []HostShardSet {
	ret := _m.ctrl.Call(_m, "HostShardSets")
	ret0, _ := ret[0].([]HostShardSet)
	return ret0
}

func (_mr *_MockStaticOptionsRecorder) HostShardSets() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HostShardSets")
}

// Mock of DynamicOptions interface
type MockDynamicOptions struct {
	ctrl     *gomock.Controller
	recorder *_MockDynamicOptionsRecorder
}

// Recorder for MockDynamicOptions (not exported)
type _MockDynamicOptionsRecorder struct {
	mock *MockDynamicOptions
}

func NewMockDynamicOptions(ctrl *gomock.Controller) *MockDynamicOptions {
	mock := &MockDynamicOptions{ctrl: ctrl}
	mock.recorder = &_MockDynamicOptionsRecorder{mock}
	return mock
}

func (_m *MockDynamicOptions) EXPECT() *_MockDynamicOptionsRecorder {
	return _m.recorder
}

func (_m *MockDynamicOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Validate")
}

func (_m *MockDynamicOptions) SetConfigServiceClient(c client.Client) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetConfigServiceClient", c)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) SetConfigServiceClient(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetConfigServiceClient", arg0)
}

func (_m *MockDynamicOptions) ConfigServiceClient() client.Client {
	ret := _m.ctrl.Call(_m, "ConfigServiceClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) ConfigServiceClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigServiceClient")
}

func (_m *MockDynamicOptions) SetService(s string) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetService", s)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) SetService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetService", arg0)
}

func (_m *MockDynamicOptions) Service() string {
	ret := _m.ctrl.Call(_m, "Service")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) Service() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Service")
}

func (_m *MockDynamicOptions) SetQueryOptions(value services.QueryOptions) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetQueryOptions", value)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) SetQueryOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetQueryOptions", arg0)
}

func (_m *MockDynamicOptions) QueryOptions() services.QueryOptions {
	ret := _m.ctrl.Call(_m, "QueryOptions")
	ret0, _ := ret[0].(services.QueryOptions)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) QueryOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryOptions")
}

func (_m *MockDynamicOptions) SetInstrumentOptions(value instrument.Options) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetInstrumentOptions", value)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) SetInstrumentOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetInstrumentOptions", arg0)
}

func (_m *MockDynamicOptions) InstrumentOptions() instrument.Options {
	ret := _m.ctrl.Call(_m, "InstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) InstrumentOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstrumentOptions")
}

func (_m *MockDynamicOptions) SetInitTimeout(value time.Duration) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetInitTimeout", value)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) SetInitTimeout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetInitTimeout", arg0)
}

func (_m *MockDynamicOptions) InitTimeout() time.Duration {
	ret := _m.ctrl.Call(_m, "InitTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) InitTimeout() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InitTimeout")
}

func (_m *MockDynamicOptions) SetHashGen(h sharding.HashGen) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetHashGen", h)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) SetHashGen(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetHashGen", arg0)
}

func (_m *MockDynamicOptions) HashGen() sharding.HashGen {
	ret := _m.ctrl.Call(_m, "HashGen")
	ret0, _ := ret[0].(sharding.HashGen)
	return ret0
}

func (_mr *_MockDynamicOptionsRecorder) HashGen() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HashGen")
}
