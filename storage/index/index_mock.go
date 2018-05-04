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
// Source: github.com/m3db/m3db/storage/index (interfaces: Results)

// Package index is a generated GoMock package.
package index

import (
	"reflect"

	"github.com/m3db/m3ninx/doc"
	"github.com/m3db/m3x/ident"

	"github.com/golang/mock/gomock"
)

// MockResults is a mock of Results interface
type MockResults struct {
	ctrl     *gomock.Controller
	recorder *MockResultsMockRecorder
}

// MockResultsMockRecorder is the mock recorder for MockResults
type MockResultsMockRecorder struct {
	mock *MockResults
}

// NewMockResults creates a new mock instance
func NewMockResults(ctrl *gomock.Controller) *MockResults {
	mock := &MockResults{ctrl: ctrl}
	mock.recorder = &MockResultsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResults) EXPECT() *MockResultsMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockResults) Add(arg0 doc.Document) (bool, error) {
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockResultsMockRecorder) Add(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockResults)(nil).Add), arg0)
}

// Finalize mocks base method
func (m *MockResults) Finalize() {
	m.ctrl.Call(m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (mr *MockResultsMockRecorder) Finalize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockResults)(nil).Finalize))
}

// Map mocks base method
func (m *MockResults) Map() *ResultsMap {
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(*ResultsMap)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockResultsMockRecorder) Map() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockResults)(nil).Map))
}

// Namespace mocks base method
func (m *MockResults) Namespace() ident.ID {
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(ident.ID)
	return ret0
}

// Namespace indicates an expected call of Namespace
func (mr *MockResultsMockRecorder) Namespace() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockResults)(nil).Namespace))
}

// Reset mocks base method
func (m *MockResults) Reset(arg0 ident.ID) {
	m.ctrl.Call(m, "Reset", arg0)
}

// Reset indicates an expected call of Reset
func (mr *MockResultsMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockResults)(nil).Reset), arg0)
}
