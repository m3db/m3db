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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3db/persist/types.go

package persist

import (
	time "time"

	namespace "github.com/m3db/m3db/storage/namespace"

	gomock "github.com/golang/mock/gomock"
)

// Mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *_MockManagerRecorder
}

// Recorder for MockManager (not exported)
type _MockManagerRecorder struct {
	mock *MockManager
}

func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &_MockManagerRecorder{mock}
	return mock
}

func (_m *MockManager) EXPECT() *_MockManagerRecorder {
	return _m.recorder
}

func (_m *MockManager) StartFlush() (Flush, error) {
	ret := _m.ctrl.Call(_m, "StartFlush")
	ret0, _ := ret[0].(Flush)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) StartFlush() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartFlush")
}

// Mock of Flush interface
type MockFlush struct {
	ctrl     *gomock.Controller
	recorder *_MockFlushRecorder
}

// Recorder for MockFlush (not exported)
type _MockFlushRecorder struct {
	mock *MockFlush
}

func NewMockFlush(ctrl *gomock.Controller) *MockFlush {
	mock := &MockFlush{ctrl: ctrl}
	mock.recorder = &_MockFlushRecorder{mock}
	return mock
}

func (_m *MockFlush) EXPECT() *_MockFlushRecorder {
	return _m.recorder
}

func (_m *MockFlush) Prepare(ns namespace.Metadata, shard uint32, blockStart time.Time) (PreparedPersist, error) {
	ret := _m.ctrl.Call(_m, "Prepare", ns, shard, blockStart)
	ret0, _ := ret[0].(PreparedPersist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFlushRecorder) Prepare(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prepare", arg0, arg1, arg2)
}

func (_m *MockFlush) Done() error {
	ret := _m.ctrl.Call(_m, "Done")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFlushRecorder) Done() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Done")
}
