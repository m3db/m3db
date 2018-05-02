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

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3db/x/xio (interfaces: BlockReader,SegmentReader)

package xio

import (
	"reflect"
	"time"

	"github.com/m3db/m3db/ts"

	"github.com/golang/mock/gomock"
)

// MockBlockReader is a mock of BlockReader interface
type MockBlockReader struct {
	ctrl     *gomock.Controller
	recorder *MockBlockReaderMockRecorder
}

// MockBlockReaderMockRecorder is the mock recorder for MockBlockReader
type MockBlockReaderMockRecorder struct {
	mock *MockBlockReader
}

// NewMockBlockReader creates a new mock instance
func NewMockBlockReader(ctrl *gomock.Controller) *MockBlockReader {
	mock := &MockBlockReader{ctrl: ctrl}
	mock.recorder = &MockBlockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockBlockReader) EXPECT() *MockBlockReaderMockRecorder {
	return _m.recorder
}

// Clone mocks base method
func (_m *MockBlockReader) Clone() (SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "Clone")
	ret0, _ := ret[0].(SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Clone indicates an expected call of Clone
func (_mr *MockBlockReaderMockRecorder) Clone() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Clone", reflect.TypeOf((*MockBlockReader)(nil).Clone))
}

// End mocks base method
func (_m *MockBlockReader) End() time.Time {
	ret := _m.ctrl.Call(_m, "End")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// End indicates an expected call of End
func (_mr *MockBlockReaderMockRecorder) End() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "End", reflect.TypeOf((*MockBlockReader)(nil).End))
}

// Finalize mocks base method
func (_m *MockBlockReader) Finalize() {
	_m.ctrl.Call(_m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (_mr *MockBlockReaderMockRecorder) Finalize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Finalize", reflect.TypeOf((*MockBlockReader)(nil).Finalize))
}

// Read mocks base method
func (_m *MockBlockReader) Read(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Read", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (_mr *MockBlockReaderMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Read", reflect.TypeOf((*MockBlockReader)(nil).Read), arg0)
}

// Reset mocks base method
func (_m *MockBlockReader) Reset(_param0 ts.Segment) {
	_m.ctrl.Call(_m, "Reset", _param0)
}

// Reset indicates an expected call of Reset
func (_mr *MockBlockReaderMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Reset", reflect.TypeOf((*MockBlockReader)(nil).Reset), arg0)
}

// ResetWindowed mocks base method
func (_m *MockBlockReader) ResetWindowed(_param0 ts.Segment, _param1 time.Time, _param2 time.Time) {
	_m.ctrl.Call(_m, "ResetWindowed", _param0, _param1, _param2)
}

// ResetWindowed indicates an expected call of ResetWindowed
func (_mr *MockBlockReaderMockRecorder) ResetWindowed(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ResetWindowed", reflect.TypeOf((*MockBlockReader)(nil).ResetWindowed), arg0, arg1, arg2)
}

// Segment mocks base method
func (_m *MockBlockReader) Segment() (ts.Segment, error) {
	ret := _m.ctrl.Call(_m, "Segment")
	ret0, _ := ret[0].(ts.Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Segment indicates an expected call of Segment
func (_mr *MockBlockReaderMockRecorder) Segment() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Segment", reflect.TypeOf((*MockBlockReader)(nil).Segment))
}

// SegmentReader mocks base method
func (_m *MockBlockReader) SegmentReader() (SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "SegmentReader")
	ret0, _ := ret[0].(SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SegmentReader indicates an expected call of SegmentReader
func (_mr *MockBlockReaderMockRecorder) SegmentReader() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SegmentReader", reflect.TypeOf((*MockBlockReader)(nil).SegmentReader))
}

// Start mocks base method
func (_m *MockBlockReader) Start() time.Time {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Start indicates an expected call of Start
func (_mr *MockBlockReaderMockRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Start", reflect.TypeOf((*MockBlockReader)(nil).Start))
}

// MockSegmentReader is a mock of SegmentReader interface
type MockSegmentReader struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentReaderMockRecorder
}

// MockSegmentReaderMockRecorder is the mock recorder for MockSegmentReader
type MockSegmentReaderMockRecorder struct {
	mock *MockSegmentReader
}

// NewMockSegmentReader creates a new mock instance
func NewMockSegmentReader(ctrl *gomock.Controller) *MockSegmentReader {
	mock := &MockSegmentReader{ctrl: ctrl}
	mock.recorder = &MockSegmentReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockSegmentReader) EXPECT() *MockSegmentReaderMockRecorder {
	return _m.recorder
}

// Clone mocks base method
func (_m *MockSegmentReader) Clone() (SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "Clone")
	ret0, _ := ret[0].(SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Clone indicates an expected call of Clone
func (_mr *MockSegmentReaderMockRecorder) Clone() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Clone", reflect.TypeOf((*MockSegmentReader)(nil).Clone))
}

// Finalize mocks base method
func (_m *MockSegmentReader) Finalize() {
	_m.ctrl.Call(_m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (_mr *MockSegmentReaderMockRecorder) Finalize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Finalize", reflect.TypeOf((*MockSegmentReader)(nil).Finalize))
}

// Read mocks base method
func (_m *MockSegmentReader) Read(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Read", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (_mr *MockSegmentReaderMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Read", reflect.TypeOf((*MockSegmentReader)(nil).Read), arg0)
}

// Reset mocks base method
func (_m *MockSegmentReader) Reset(_param0 ts.Segment) {
	_m.ctrl.Call(_m, "Reset", _param0)
}

// Reset indicates an expected call of Reset
func (_mr *MockSegmentReaderMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Reset", reflect.TypeOf((*MockSegmentReader)(nil).Reset), arg0)
}

// Segment mocks base method
func (_m *MockSegmentReader) Segment() (ts.Segment, error) {
	ret := _m.ctrl.Call(_m, "Segment")
	ret0, _ := ret[0].(ts.Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Segment indicates an expected call of Segment
func (_mr *MockSegmentReaderMockRecorder) Segment() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Segment", reflect.TypeOf((*MockSegmentReader)(nil).Segment))
}
