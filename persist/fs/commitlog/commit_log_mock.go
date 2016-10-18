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
// Source: github.com/m3db/m3db/persist/fs/commitlog/types.go

package commitlog

import (
	gomock "github.com/golang/mock/gomock"
	clock "github.com/m3db/m3db/clock"
	context "github.com/m3db/m3db/context"
	instrument "github.com/m3db/m3db/instrument"
	fs "github.com/m3db/m3db/persist/fs"
	retention "github.com/m3db/m3db/retention"
	ts "github.com/m3db/m3db/ts"
	time0 "github.com/m3db/m3x/time"
	time "time"
)

// Mock of CommitLog interface
type MockCommitLog struct {
	ctrl     *gomock.Controller
	recorder *_MockCommitLogRecorder
}

// Recorder for MockCommitLog (not exported)
type _MockCommitLogRecorder struct {
	mock *MockCommitLog
}

func NewMockCommitLog(ctrl *gomock.Controller) *MockCommitLog {
	mock := &MockCommitLog{ctrl: ctrl}
	mock.recorder = &_MockCommitLogRecorder{mock}
	return mock
}

func (_m *MockCommitLog) EXPECT() *_MockCommitLogRecorder {
	return _m.recorder
}

func (_m *MockCommitLog) Open() error {
	ret := _m.ctrl.Call(_m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCommitLogRecorder) Open() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open")
}

func (_m *MockCommitLog) Write(ctx context.Context, series Series, datapoint ts.Datapoint, unit time0.Unit, annotation ts.Annotation) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, series, datapoint, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCommitLogRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockCommitLog) WriteBehind(ctx context.Context, series Series, datapoint ts.Datapoint, unit time0.Unit, annotation ts.Annotation) error {
	ret := _m.ctrl.Call(_m, "WriteBehind", ctx, series, datapoint, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCommitLogRecorder) WriteBehind(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteBehind", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockCommitLog) Iter() (Iterator, error) {
	ret := _m.ctrl.Call(_m, "Iter")
	ret0, _ := ret[0].(Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCommitLogRecorder) Iter() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Iter")
}

func (_m *MockCommitLog) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCommitLogRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of Iterator interface
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *_MockIteratorRecorder
}

// Recorder for MockIterator (not exported)
type _MockIteratorRecorder struct {
	mock *MockIterator
}

func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &_MockIteratorRecorder{mock}
	return mock
}

func (_m *MockIterator) EXPECT() *_MockIteratorRecorder {
	return _m.recorder
}

func (_m *MockIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockIteratorRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

func (_m *MockIterator) Current() (Series, ts.Datapoint, time0.Unit, ts.Annotation) {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(Series)
	ret1, _ := ret[1].(ts.Datapoint)
	ret2, _ := ret[2].(time0.Unit)
	ret3, _ := ret[3].(ts.Annotation)
	return ret0, ret1, ret2, ret3
}

func (_mr *_MockIteratorRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

func (_m *MockIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIteratorRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockIterator) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockIteratorRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
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

func (_m *MockOptions) SetClockOptions(value clock.Options) Options {
	ret := _m.ctrl.Call(_m, "SetClockOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetClockOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetClockOptions", arg0)
}

func (_m *MockOptions) ClockOptions() clock.Options {
	ret := _m.ctrl.Call(_m, "ClockOptions")
	ret0, _ := ret[0].(clock.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ClockOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClockOptions")
}

func (_m *MockOptions) SetInstrumentOptions(value instrument.Options) Options {
	ret := _m.ctrl.Call(_m, "SetInstrumentOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetInstrumentOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetInstrumentOptions", arg0)
}

func (_m *MockOptions) InstrumentOptions() instrument.Options {
	ret := _m.ctrl.Call(_m, "InstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) InstrumentOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstrumentOptions")
}

func (_m *MockOptions) SetRetentionOptions(value retention.Options) Options {
	ret := _m.ctrl.Call(_m, "SetRetentionOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetRetentionOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRetentionOptions", arg0)
}

func (_m *MockOptions) RetentionOptions() retention.Options {
	ret := _m.ctrl.Call(_m, "RetentionOptions")
	ret0, _ := ret[0].(retention.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) RetentionOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetentionOptions")
}

func (_m *MockOptions) SetFilesystemOptions(value fs.Options) Options {
	ret := _m.ctrl.Call(_m, "SetFilesystemOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetFilesystemOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetFilesystemOptions", arg0)
}

func (_m *MockOptions) FilesystemOptions() fs.Options {
	ret := _m.ctrl.Call(_m, "FilesystemOptions")
	ret0, _ := ret[0].(fs.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) FilesystemOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilesystemOptions")
}

func (_m *MockOptions) SetFlushSize(value int) Options {
	ret := _m.ctrl.Call(_m, "SetFlushSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetFlushSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetFlushSize", arg0)
}

func (_m *MockOptions) FlushSize() int {
	ret := _m.ctrl.Call(_m, "FlushSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) FlushSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushSize")
}

func (_m *MockOptions) SetStrategy(value Strategy) Options {
	ret := _m.ctrl.Call(_m, "SetStrategy", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetStrategy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetStrategy", arg0)
}

func (_m *MockOptions) Strategy() Strategy {
	ret := _m.ctrl.Call(_m, "Strategy")
	ret0, _ := ret[0].(Strategy)
	return ret0
}

func (_mr *_MockOptionsRecorder) Strategy() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Strategy")
}

func (_m *MockOptions) SetFlushInterval(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetFlushInterval", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetFlushInterval(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetFlushInterval", arg0)
}

func (_m *MockOptions) FlushInterval() time.Duration {
	ret := _m.ctrl.Call(_m, "FlushInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) FlushInterval() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushInterval")
}

func (_m *MockOptions) SetBacklogQueueSize(value int) Options {
	ret := _m.ctrl.Call(_m, "SetBacklogQueueSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetBacklogQueueSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetBacklogQueueSize", arg0)
}

func (_m *MockOptions) BacklogQueueSize() int {
	ret := _m.ctrl.Call(_m, "BacklogQueueSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) BacklogQueueSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BacklogQueueSize")
}
