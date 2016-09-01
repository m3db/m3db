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
// Source: github.com/m3db/m3db/storage/bootstrap/types.go

package bootstrap

import (
	gomock "github.com/golang/mock/gomock"
	instrument "github.com/m3db/m3db/instrument"
	retention "github.com/m3db/m3db/retention"
	block "github.com/m3db/m3db/storage/block"
	time "github.com/m3db/m3x/time"
	time0 "time"
)

// Mock of Result interface
type MockResult struct {
	ctrl     *gomock.Controller
	recorder *_MockResultRecorder
}

// Recorder for MockResult (not exported)
type _MockResultRecorder struct {
	mock *MockResult
}

func NewMockResult(ctrl *gomock.Controller) *MockResult {
	mock := &MockResult{ctrl: ctrl}
	mock.recorder = &_MockResultRecorder{mock}
	return mock
}

func (_m *MockResult) EXPECT() *_MockResultRecorder {
	return _m.recorder
}

func (_m *MockResult) ShardResults() ShardResults {
	ret := _m.ctrl.Call(_m, "ShardResults")
	ret0, _ := ret[0].(ShardResults)
	return ret0
}

func (_mr *_MockResultRecorder) ShardResults() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShardResults")
}

func (_m *MockResult) Unfulfilled() ShardTimeRanges {
	ret := _m.ctrl.Call(_m, "Unfulfilled")
	ret0, _ := ret[0].(ShardTimeRanges)
	return ret0
}

func (_mr *_MockResultRecorder) Unfulfilled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unfulfilled")
}

func (_m *MockResult) AddShardResult(shard uint32, result ShardResult, unfulfilled time.Ranges) {
	_m.ctrl.Call(_m, "AddShardResult", shard, result, unfulfilled)
}

func (_mr *_MockResultRecorder) AddShardResult(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddShardResult", arg0, arg1, arg2)
}

func (_m *MockResult) SetUnfulfilled(unfulfilled ShardTimeRanges) {
	_m.ctrl.Call(_m, "SetUnfulfilled", unfulfilled)
}

func (_mr *_MockResultRecorder) SetUnfulfilled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetUnfulfilled", arg0)
}

func (_m *MockResult) AddResult(other Result) {
	_m.ctrl.Call(_m, "AddResult", other)
}

func (_mr *_MockResultRecorder) AddResult(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddResult", arg0)
}

// Mock of ShardResult interface
type MockShardResult struct {
	ctrl     *gomock.Controller
	recorder *_MockShardResultRecorder
}

// Recorder for MockShardResult (not exported)
type _MockShardResultRecorder struct {
	mock *MockShardResult
}

func NewMockShardResult(ctrl *gomock.Controller) *MockShardResult {
	mock := &MockShardResult{ctrl: ctrl}
	mock.recorder = &_MockShardResultRecorder{mock}
	return mock
}

func (_m *MockShardResult) EXPECT() *_MockShardResultRecorder {
	return _m.recorder
}

func (_m *MockShardResult) IsEmpty() bool {
	ret := _m.ctrl.Call(_m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockShardResultRecorder) IsEmpty() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsEmpty")
}

func (_m *MockShardResult) AllSeries() map[string]block.DatabaseSeriesBlocks {
	ret := _m.ctrl.Call(_m, "AllSeries")
	ret0, _ := ret[0].(map[string]block.DatabaseSeriesBlocks)
	return ret0
}

func (_mr *_MockShardResultRecorder) AllSeries() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AllSeries")
}

func (_m *MockShardResult) AddBlock(id string, block block.DatabaseBlock) {
	_m.ctrl.Call(_m, "AddBlock", id, block)
}

func (_mr *_MockShardResultRecorder) AddBlock(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddBlock", arg0, arg1)
}

func (_m *MockShardResult) AddSeries(id string, rawSeries block.DatabaseSeriesBlocks) {
	_m.ctrl.Call(_m, "AddSeries", id, rawSeries)
}

func (_mr *_MockShardResultRecorder) AddSeries(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddSeries", arg0, arg1)
}

func (_m *MockShardResult) AddResult(other ShardResult) {
	_m.ctrl.Call(_m, "AddResult", other)
}

func (_mr *_MockShardResultRecorder) AddResult(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddResult", arg0)
}

func (_m *MockShardResult) RemoveSeries(id string) {
	_m.ctrl.Call(_m, "RemoveSeries", id)
}

func (_mr *_MockShardResultRecorder) RemoveSeries(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveSeries", arg0)
}

func (_m *MockShardResult) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockShardResultRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of Bootstrap interface
type MockBootstrap struct {
	ctrl     *gomock.Controller
	recorder *_MockBootstrapRecorder
}

// Recorder for MockBootstrap (not exported)
type _MockBootstrapRecorder struct {
	mock *MockBootstrap
}

func NewMockBootstrap(ctrl *gomock.Controller) *MockBootstrap {
	mock := &MockBootstrap{ctrl: ctrl}
	mock.recorder = &_MockBootstrapRecorder{mock}
	return mock
}

func (_m *MockBootstrap) EXPECT() *_MockBootstrapRecorder {
	return _m.recorder
}

func (_m *MockBootstrap) Run(writeStart time0.Time, shards []uint32) (Result, error) {
	ret := _m.ctrl.Call(_m, "Run", writeStart, shards)
	ret0, _ := ret[0].(Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBootstrapRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0, arg1)
}

// Mock of Bootstrapper interface
type MockBootstrapper struct {
	ctrl     *gomock.Controller
	recorder *_MockBootstrapperRecorder
}

// Recorder for MockBootstrapper (not exported)
type _MockBootstrapperRecorder struct {
	mock *MockBootstrapper
}

func NewMockBootstrapper(ctrl *gomock.Controller) *MockBootstrapper {
	mock := &MockBootstrapper{ctrl: ctrl}
	mock.recorder = &_MockBootstrapperRecorder{mock}
	return mock
}

func (_m *MockBootstrapper) EXPECT() *_MockBootstrapperRecorder {
	return _m.recorder
}

func (_m *MockBootstrapper) Can(strategy Strategy) bool {
	ret := _m.ctrl.Call(_m, "Can", strategy)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockBootstrapperRecorder) Can(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Can", arg0)
}

func (_m *MockBootstrapper) Bootstrap(shardsTimeRanges ShardTimeRanges) (Result, error) {
	ret := _m.ctrl.Call(_m, "Bootstrap", shardsTimeRanges)
	ret0, _ := ret[0].(Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBootstrapperRecorder) Bootstrap(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap", arg0)
}

// Mock of Source interface
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *_MockSourceRecorder
}

// Recorder for MockSource (not exported)
type _MockSourceRecorder struct {
	mock *MockSource
}

func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &_MockSourceRecorder{mock}
	return mock
}

func (_m *MockSource) EXPECT() *_MockSourceRecorder {
	return _m.recorder
}

func (_m *MockSource) Can(strategy Strategy) bool {
	ret := _m.ctrl.Call(_m, "Can", strategy)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockSourceRecorder) Can(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Can", arg0)
}

func (_m *MockSource) Available(shardsTimeRanges ShardTimeRanges) ShardTimeRanges {
	ret := _m.ctrl.Call(_m, "Available", shardsTimeRanges)
	ret0, _ := ret[0].(ShardTimeRanges)
	return ret0
}

func (_mr *_MockSourceRecorder) Available(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Available", arg0)
}

func (_m *MockSource) Read(shardsTimeRanges ShardTimeRanges) (Result, error) {
	ret := _m.ctrl.Call(_m, "Read", shardsTimeRanges)
	ret0, _ := ret[0].(Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSourceRecorder) Read(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0)
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

func (_m *MockOptions) RetentionOptions(value retention.Options) Options {
	ret := _m.ctrl.Call(_m, "RetentionOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) RetentionOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetentionOptions", arg0)
}

func (_m *MockOptions) GetRetentionOptions() retention.Options {
	ret := _m.ctrl.Call(_m, "GetRetentionOptions")
	ret0, _ := ret[0].(retention.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetRetentionOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRetentionOptions")
}

func (_m *MockOptions) DatabaseBlockOptions(value block.Options) Options {
	ret := _m.ctrl.Call(_m, "DatabaseBlockOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseBlockOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseBlockOptions", arg0)
}

func (_m *MockOptions) GetDatabaseBlockOptions() block.Options {
	ret := _m.ctrl.Call(_m, "GetDatabaseBlockOptions")
	ret0, _ := ret[0].(block.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetDatabaseBlockOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDatabaseBlockOptions")
}
