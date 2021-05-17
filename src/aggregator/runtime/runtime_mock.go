// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/aggregator/runtime (interfaces: OptionsWatcher)

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

// Package runtime is a generated GoMock package.
package runtime

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockOptionsWatcher is a mock of OptionsWatcher interface.
type MockOptionsWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsWatcherMockRecorder
}

// MockOptionsWatcherMockRecorder is the mock recorder for MockOptionsWatcher.
type MockOptionsWatcherMockRecorder struct {
	mock *MockOptionsWatcher
}

// NewMockOptionsWatcher creates a new mock instance.
func NewMockOptionsWatcher(ctrl *gomock.Controller) *MockOptionsWatcher {
	mock := &MockOptionsWatcher{ctrl: ctrl}
	mock.recorder = &MockOptionsWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOptionsWatcher) EXPECT() *MockOptionsWatcherMockRecorder {
	return m.recorder
}

// SetRuntimeOptions mocks base method.
func (m *MockOptionsWatcher) SetRuntimeOptions(arg0 Options) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRuntimeOptions", arg0)
}

// SetRuntimeOptions indicates an expected call of SetRuntimeOptions.
func (mr *MockOptionsWatcherMockRecorder) SetRuntimeOptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRuntimeOptions", reflect.TypeOf((*MockOptionsWatcher)(nil).SetRuntimeOptions), arg0)
}
