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
	result "github.com/m3db/m3db/storage/bootstrap/result"
	ts "github.com/m3db/m3db/ts"
	time "github.com/m3db/m3x/time"
)

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

func (_m *MockBootstrap) Run(targetRanges time.Ranges, namespace ts.ID, shards []uint32) (result.BootstrapResult, error) {
	ret := _m.ctrl.Call(_m, "Run", targetRanges, namespace, shards)
	ret0, _ := ret[0].(result.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBootstrapRecorder) Run(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0, arg1, arg2)
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

func (_m *MockBootstrapper) String() string {
	ret := _m.ctrl.Call(_m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockBootstrapperRecorder) String() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "String")
}

func (_m *MockBootstrapper) Can(strategy Strategy) bool {
	ret := _m.ctrl.Call(_m, "Can", strategy)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockBootstrapperRecorder) Can(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Can", arg0)
}

func (_m *MockBootstrapper) Bootstrap(namespace ts.ID, shardsTimeRanges result.ShardTimeRanges) (result.BootstrapResult, error) {
	ret := _m.ctrl.Call(_m, "Bootstrap", namespace, shardsTimeRanges)
	ret0, _ := ret[0].(result.BootstrapResult)
	ret1, _ := ret[1].(error)
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

func (_m *MockSource) Can(strategy Strategy) bool {
	ret := _m.ctrl.Call(_m, "Can", strategy)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockSourceRecorder) Can(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Can", arg0)
}

func (_m *MockSource) Available(namespace ts.ID, shardsTimeRanges result.ShardTimeRanges) result.ShardTimeRanges {
	ret := _m.ctrl.Call(_m, "Available", namespace, shardsTimeRanges)
	ret0, _ := ret[0].(result.ShardTimeRanges)
	return ret0
}

func (_mr *_MockSourceRecorder) Available(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Available", arg0, arg1)
}

func (_m *MockSource) Read(namespace ts.ID, shardsTimeRanges result.ShardTimeRanges) (result.BootstrapResult, error) {
	ret := _m.ctrl.Call(_m, "Read", namespace, shardsTimeRanges)
	ret0, _ := ret[0].(result.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSourceRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0, arg1)
}
