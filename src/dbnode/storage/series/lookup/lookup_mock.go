// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/storage/series/lookup (interfaces: IndexWriter)

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

// Package lookup is a generated GoMock package.
package lookup

import (
	"reflect"
	"time"

	"github.com/m3db/m3/src/dbnode/ts/writes"
	time0 "github.com/m3db/m3/src/x/time"

	"github.com/golang/mock/gomock"
)

// MockIndexWriter is a mock of IndexWriter interface.
type MockIndexWriter struct {
	ctrl     *gomock.Controller
	recorder *MockIndexWriterMockRecorder
}

// MockIndexWriterMockRecorder is the mock recorder for MockIndexWriter.
type MockIndexWriterMockRecorder struct {
	mock *MockIndexWriter
}

// NewMockIndexWriter creates a new mock instance.
func NewMockIndexWriter(ctrl *gomock.Controller) *MockIndexWriter {
	mock := &MockIndexWriter{ctrl: ctrl}
	mock.recorder = &MockIndexWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexWriter) EXPECT() *MockIndexWriterMockRecorder {
	return m.recorder
}

// BlockStartForWriteTime mocks base method.
func (m *MockIndexWriter) BlockStartForWriteTime(arg0 time.Time) time0.UnixNano {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockStartForWriteTime", arg0)
	ret0, _ := ret[0].(time0.UnixNano)
	return ret0
}

// BlockStartForWriteTime indicates an expected call of BlockStartForWriteTime.
func (mr *MockIndexWriterMockRecorder) BlockStartForWriteTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockStartForWriteTime", reflect.TypeOf((*MockIndexWriter)(nil).BlockStartForWriteTime), arg0)
}

// WritePending mocks base method.
func (m *MockIndexWriter) WritePending(arg0 []writes.PendingIndexInsert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WritePending", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WritePending indicates an expected call of WritePending.
func (mr *MockIndexWriterMockRecorder) WritePending(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePending", reflect.TypeOf((*MockIndexWriter)(nil).WritePending), arg0)
}
