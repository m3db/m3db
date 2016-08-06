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
// Source: github.com/m3db/m3db/interfaces/m3db (interfaces: ShardResult,Bootstrap,Bootstrapper,Source)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	m3db "github.com/m3db/m3db/interfaces/m3db"
	time0 "github.com/m3db/m3x/time"
	time "time"
)

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

func (_m *MockShardResult) AddBlock(_param0 string, _param1 m3db.DatabaseBlock) {
	_m.ctrl.Call(_m, "AddBlock", _param0, _param1)
}

func (_mr *_MockShardResultRecorder) AddBlock(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddBlock", arg0, arg1)
}

func (_m *MockShardResult) AddResult(_param0 m3db.ShardResult) {
	_m.ctrl.Call(_m, "AddResult", _param0)
}

func (_mr *_MockShardResultRecorder) AddResult(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddResult", arg0)
}

func (_m *MockShardResult) AddSeries(_param0 string, _param1 m3db.DatabaseSeriesBlocks) {
	_m.ctrl.Call(_m, "AddSeries", _param0, _param1)
}

func (_mr *_MockShardResultRecorder) AddSeries(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddSeries", arg0, arg1)
}

func (_m *MockShardResult) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockShardResultRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockShardResult) GetAllSeries() map[string]m3db.DatabaseSeriesBlocks {
	ret := _m.ctrl.Call(_m, "GetAllSeries")
	ret0, _ := ret[0].(map[string]m3db.DatabaseSeriesBlocks)
	return ret0
}

func (_mr *_MockShardResultRecorder) GetAllSeries() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllSeries")
}

func (_m *MockShardResult) IsEmpty() bool {
	ret := _m.ctrl.Call(_m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockShardResultRecorder) IsEmpty() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsEmpty")
}

func (_m *MockShardResult) RemoveSeries(_param0 string) {
	_m.ctrl.Call(_m, "RemoveSeries", _param0)
}

func (_mr *_MockShardResultRecorder) RemoveSeries(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveSeries", arg0)
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

func (_m *MockBootstrap) Run(_param0 time.Time, _param1 uint32) (m3db.ShardResult, error) {
	ret := _m.ctrl.Call(_m, "Run", _param0, _param1)
	ret0, _ := ret[0].(m3db.ShardResult)
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

func (_m *MockBootstrapper) Bootstrap(_param0 uint32, _param1 time0.Ranges) (m3db.ShardResult, time0.Ranges) {
	ret := _m.ctrl.Call(_m, "Bootstrap", _param0, _param1)
	ret0, _ := ret[0].(m3db.ShardResult)
	ret1, _ := ret[1].(time0.Ranges)
	return ret0, ret1
}

func (_mr *_MockBootstrapperRecorder) Bootstrap(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap", arg0, arg1)
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

func (_m *MockSource) GetAvailability(_param0 uint32, _param1 time0.Ranges) time0.Ranges {
	ret := _m.ctrl.Call(_m, "GetAvailability", _param0, _param1)
	ret0, _ := ret[0].(time0.Ranges)
	return ret0
}

func (_mr *_MockSourceRecorder) GetAvailability(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAvailability", arg0, arg1)
}

func (_m *MockSource) ReadData(_param0 uint32, _param1 time0.Ranges) (m3db.ShardResult, time0.Ranges) {
	ret := _m.ctrl.Call(_m, "ReadData", _param0, _param1)
	ret0, _ := ret[0].(m3db.ShardResult)
	ret1, _ := ret[1].(time0.Ranges)
	return ret0, ret1
}

func (_mr *_MockSourceRecorder) ReadData(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadData", arg0, arg1)
}
