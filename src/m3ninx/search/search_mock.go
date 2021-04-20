// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/m3ninx/search/types.go

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

// Package search is a generated GoMock package.
package search

import (
	"reflect"

	"github.com/m3db/m3/src/m3ninx/doc"
	"github.com/m3db/m3/src/m3ninx/generated/proto/querypb"
	"github.com/m3db/m3/src/m3ninx/index"
	"github.com/m3db/m3/src/m3ninx/postings"
	"github.com/m3db/m3/src/x/context"

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
func (m *MockExecutor) Execute(ctx context.Context, q Query) (doc.QueryDocIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, q)
	ret0, _ := ret[0].(doc.QueryDocIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockExecutorMockRecorder) Execute(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockExecutor)(nil).Execute), ctx, q)
}

// Close mocks base method
func (m *MockExecutor) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockExecutorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
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
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockQueryMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockQuery)(nil).String))
}

// Searcher mocks base method
func (m *MockQuery) Searcher() (Searcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Searcher")
	ret0, _ := ret[0].(Searcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Searcher indicates an expected call of Searcher
func (mr *MockQueryMockRecorder) Searcher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Searcher", reflect.TypeOf((*MockQuery)(nil).Searcher))
}

// Equal mocks base method
func (m *MockQuery) Equal(q Query) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", q)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockQueryMockRecorder) Equal(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockQuery)(nil).Equal), q)
}

// ToProto mocks base method
func (m *MockQuery) ToProto() *querypb.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToProto")
	ret0, _ := ret[0].(*querypb.Query)
	return ret0
}

// ToProto indicates an expected call of ToProto
func (mr *MockQueryMockRecorder) ToProto() *gomock.Call {
	mr.mock.ctrl.T.Helper()
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

// Search mocks base method
func (m *MockSearcher) Search(arg0 index.Reader) (postings.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].(postings.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockSearcherMockRecorder) Search(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockSearcher)(nil).Search), arg0)
}

// MockReadThroughSegmentSearcher is a mock of ReadThroughSegmentSearcher interface
type MockReadThroughSegmentSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockReadThroughSegmentSearcherMockRecorder
}

// MockReadThroughSegmentSearcherMockRecorder is the mock recorder for MockReadThroughSegmentSearcher
type MockReadThroughSegmentSearcherMockRecorder struct {
	mock *MockReadThroughSegmentSearcher
}

// NewMockReadThroughSegmentSearcher creates a new mock instance
func NewMockReadThroughSegmentSearcher(ctrl *gomock.Controller) *MockReadThroughSegmentSearcher {
	mock := &MockReadThroughSegmentSearcher{ctrl: ctrl}
	mock.recorder = &MockReadThroughSegmentSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadThroughSegmentSearcher) EXPECT() *MockReadThroughSegmentSearcherMockRecorder {
	return m.recorder
}

// Search mocks base method
func (m *MockReadThroughSegmentSearcher) Search(query Query, searcher Searcher) (postings.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query, searcher)
	ret0, _ := ret[0].(postings.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockReadThroughSegmentSearcherMockRecorder) Search(query, searcher interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockReadThroughSegmentSearcher)(nil).Search), query, searcher)
}
