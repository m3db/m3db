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
// Source: github.com/m3db/m3ninx/search/types.go

// Package search is a generated GoMock package.
package search

import (
	"reflect"

	"github.com/m3db/m3ninx/doc"
	"github.com/m3db/m3ninx/generated/proto/querypb"
	"github.com/m3db/m3ninx/index"
	"github.com/m3db/m3ninx/postings"

	"github.com/golang/mock/gomock"
)

// MockExecutor is a mock of Executor interface
type MockExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorMockRecorder
}

// MockExecutorMockRecorder is the mock recorder for MockExecutor
type MockExecutorMockRecorder struct {
	mock *MockExecutor
}

// NewMockExecutor creates a new mock instance
func NewMockExecutor(ctrl *gomock.Controller) *MockExecutor {
	mock := &MockExecutor{ctrl: ctrl}
	mock.recorder = &MockExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecutor) EXPECT() *MockExecutorMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockExecutor) Execute(q Query) (doc.Iterator, error) {
	ret := m.ctrl.Call(m, "Execute", q)
	ret0, _ := ret[0].(doc.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockExecutorMockRecorder) Execute(q interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockExecutor)(nil).Execute), q)
}

// Close mocks base method
func (m *MockExecutor) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockExecutorMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockExecutor)(nil).Close))
}

// MockQuery is a mock of Query interface
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// String mocks base method
func (m *MockQuery) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockQueryMockRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockQuery)(nil).String))
}

// Searcher mocks base method
func (m *MockQuery) Searcher(rs index.Readers) (Searcher, error) {
	ret := m.ctrl.Call(m, "Searcher", rs)
	ret0, _ := ret[0].(Searcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Searcher indicates an expected call of Searcher
func (mr *MockQueryMockRecorder) Searcher(rs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Searcher", reflect.TypeOf((*MockQuery)(nil).Searcher), rs)
}

// Equal mocks base method
func (m *MockQuery) Equal(q Query) bool {
	ret := m.ctrl.Call(m, "Equal", q)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockQueryMockRecorder) Equal(q interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockQuery)(nil).Equal), q)
}

// Marshal mocks base method
func (m *MockQuery) Marshal() ([]byte, error) {
	ret := m.ctrl.Call(m, "Marshal")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal
func (mr *MockQueryMockRecorder) Marshal() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockQuery)(nil).Marshal))
}

// Unmarshal mocks base method
func (m *MockQuery) Unmarshal(data []byte) error {
	ret := m.ctrl.Call(m, "Unmarshal", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal
func (mr *MockQueryMockRecorder) Unmarshal(data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockQuery)(nil).Unmarshal), data)
}

// ToProto mocks base method
func (m *MockQuery) ToProto() *querypb.Query {
	ret := m.ctrl.Call(m, "ToProto")
	ret0, _ := ret[0].(*querypb.Query)
	return ret0
}

// ToProto indicates an expected call of ToProto
func (mr *MockQueryMockRecorder) ToProto() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToProto", reflect.TypeOf((*MockQuery)(nil).ToProto))
}

// MockSearcher is a mock of Searcher interface
type MockSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockSearcherMockRecorder
}

// MockSearcherMockRecorder is the mock recorder for MockSearcher
type MockSearcherMockRecorder struct {
	mock *MockSearcher
}

// NewMockSearcher creates a new mock instance
func NewMockSearcher(ctrl *gomock.Controller) *MockSearcher {
	mock := &MockSearcher{ctrl: ctrl}
	mock.recorder = &MockSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearcher) EXPECT() *MockSearcherMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockSearcher) Next() bool {
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockSearcherMockRecorder) Next() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockSearcher)(nil).Next))
}

// Current mocks base method
func (m *MockSearcher) Current() postings.List {
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(postings.List)
	return ret0
}

// Current indicates an expected call of Current
func (mr *MockSearcherMockRecorder) Current() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockSearcher)(nil).Current))
}

// Err mocks base method
func (m *MockSearcher) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockSearcherMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSearcher)(nil).Err))
}

// NumReaders mocks base method
func (m *MockSearcher) NumReaders() int {
	ret := m.ctrl.Call(m, "NumReaders")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumReaders indicates an expected call of NumReaders
func (mr *MockSearcherMockRecorder) NumReaders() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumReaders", reflect.TypeOf((*MockSearcher)(nil).NumReaders))
}
