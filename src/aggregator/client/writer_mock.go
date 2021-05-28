// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/aggregator/client/writer.go

// Copyright (c) 2021 Uber Technologies, Inc.
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

// Package client is a generated GoMock package.
package client

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockinstanceWriter is a mock of instanceWriter interface.
type MockinstanceWriter struct {
	ctrl     *gomock.Controller
	recorder *MockinstanceWriterMockRecorder
}

// MockinstanceWriterMockRecorder is the mock recorder for MockinstanceWriter.
type MockinstanceWriterMockRecorder struct {
	mock *MockinstanceWriter
}

// NewMockinstanceWriter creates a new mock instance.
func NewMockinstanceWriter(ctrl *gomock.Controller) *MockinstanceWriter {
	mock := &MockinstanceWriter{ctrl: ctrl}
	mock.recorder = &MockinstanceWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockinstanceWriter) EXPECT() *MockinstanceWriterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockinstanceWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockinstanceWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockinstanceWriter)(nil).Close))
}

// Flush mocks base method.
func (m *MockinstanceWriter) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockinstanceWriterMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockinstanceWriter)(nil).Flush))
}

// QueueSize mocks base method.
func (m *MockinstanceWriter) QueueSize() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// QueueSize indicates an expected call of QueueSize.
func (mr *MockinstanceWriterMockRecorder) QueueSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueSize", reflect.TypeOf((*MockinstanceWriter)(nil).QueueSize))
}

// Write mocks base method.
func (m *MockinstanceWriter) Write(shard uint32, payload payloadUnion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", shard, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockinstanceWriterMockRecorder) Write(shard, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockinstanceWriter)(nil).Write), shard, payload)
}
