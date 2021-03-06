// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/ctl/service/r2/store (interfaces: Store)

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

// Package store is a generated GoMock package.
package store

import (
	"reflect"

	"github.com/m3db/m3/src/metrics/rules/view"
	"github.com/m3db/m3/src/metrics/rules/view/changes"

	"github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStore) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStore)(nil).Close))
}

// CreateMappingRule mocks base method.
func (m *MockStore) CreateMappingRule(arg0 string, arg1 view.MappingRule, arg2 UpdateOptions) (view.MappingRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMappingRule", arg0, arg1, arg2)
	ret0, _ := ret[0].(view.MappingRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMappingRule indicates an expected call of CreateMappingRule.
func (mr *MockStoreMockRecorder) CreateMappingRule(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMappingRule", reflect.TypeOf((*MockStore)(nil).CreateMappingRule), arg0, arg1, arg2)
}

// CreateNamespace mocks base method.
func (m *MockStore) CreateNamespace(arg0 string, arg1 UpdateOptions) (view.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespace", arg0, arg1)
	ret0, _ := ret[0].(view.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNamespace indicates an expected call of CreateNamespace.
func (mr *MockStoreMockRecorder) CreateNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespace", reflect.TypeOf((*MockStore)(nil).CreateNamespace), arg0, arg1)
}

// CreateRollupRule mocks base method.
func (m *MockStore) CreateRollupRule(arg0 string, arg1 view.RollupRule, arg2 UpdateOptions) (view.RollupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRollupRule", arg0, arg1, arg2)
	ret0, _ := ret[0].(view.RollupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRollupRule indicates an expected call of CreateRollupRule.
func (mr *MockStoreMockRecorder) CreateRollupRule(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRollupRule", reflect.TypeOf((*MockStore)(nil).CreateRollupRule), arg0, arg1, arg2)
}

// DeleteMappingRule mocks base method.
func (m *MockStore) DeleteMappingRule(arg0, arg1 string, arg2 UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMappingRule", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMappingRule indicates an expected call of DeleteMappingRule.
func (mr *MockStoreMockRecorder) DeleteMappingRule(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMappingRule", reflect.TypeOf((*MockStore)(nil).DeleteMappingRule), arg0, arg1, arg2)
}

// DeleteNamespace mocks base method.
func (m *MockStore) DeleteNamespace(arg0 string, arg1 UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNamespace indicates an expected call of DeleteNamespace.
func (mr *MockStoreMockRecorder) DeleteNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockStore)(nil).DeleteNamespace), arg0, arg1)
}

// DeleteRollupRule mocks base method.
func (m *MockStore) DeleteRollupRule(arg0, arg1 string, arg2 UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRollupRule", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRollupRule indicates an expected call of DeleteRollupRule.
func (mr *MockStoreMockRecorder) DeleteRollupRule(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRollupRule", reflect.TypeOf((*MockStore)(nil).DeleteRollupRule), arg0, arg1, arg2)
}

// FetchMappingRule mocks base method.
func (m *MockStore) FetchMappingRule(arg0, arg1 string) (view.MappingRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMappingRule", arg0, arg1)
	ret0, _ := ret[0].(view.MappingRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMappingRule indicates an expected call of FetchMappingRule.
func (mr *MockStoreMockRecorder) FetchMappingRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMappingRule", reflect.TypeOf((*MockStore)(nil).FetchMappingRule), arg0, arg1)
}

// FetchMappingRuleHistory mocks base method.
func (m *MockStore) FetchMappingRuleHistory(arg0, arg1 string) ([]view.MappingRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMappingRuleHistory", arg0, arg1)
	ret0, _ := ret[0].([]view.MappingRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMappingRuleHistory indicates an expected call of FetchMappingRuleHistory.
func (mr *MockStoreMockRecorder) FetchMappingRuleHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMappingRuleHistory", reflect.TypeOf((*MockStore)(nil).FetchMappingRuleHistory), arg0, arg1)
}

// FetchNamespaces mocks base method.
func (m *MockStore) FetchNamespaces() (view.Namespaces, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchNamespaces")
	ret0, _ := ret[0].(view.Namespaces)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchNamespaces indicates an expected call of FetchNamespaces.
func (mr *MockStoreMockRecorder) FetchNamespaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchNamespaces", reflect.TypeOf((*MockStore)(nil).FetchNamespaces))
}

// FetchRollupRule mocks base method.
func (m *MockStore) FetchRollupRule(arg0, arg1 string) (view.RollupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRollupRule", arg0, arg1)
	ret0, _ := ret[0].(view.RollupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRollupRule indicates an expected call of FetchRollupRule.
func (mr *MockStoreMockRecorder) FetchRollupRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRollupRule", reflect.TypeOf((*MockStore)(nil).FetchRollupRule), arg0, arg1)
}

// FetchRollupRuleHistory mocks base method.
func (m *MockStore) FetchRollupRuleHistory(arg0, arg1 string) ([]view.RollupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRollupRuleHistory", arg0, arg1)
	ret0, _ := ret[0].([]view.RollupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRollupRuleHistory indicates an expected call of FetchRollupRuleHistory.
func (mr *MockStoreMockRecorder) FetchRollupRuleHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRollupRuleHistory", reflect.TypeOf((*MockStore)(nil).FetchRollupRuleHistory), arg0, arg1)
}

// FetchRuleSetSnapshot mocks base method.
func (m *MockStore) FetchRuleSetSnapshot(arg0 string) (view.RuleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRuleSetSnapshot", arg0)
	ret0, _ := ret[0].(view.RuleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRuleSetSnapshot indicates an expected call of FetchRuleSetSnapshot.
func (mr *MockStoreMockRecorder) FetchRuleSetSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRuleSetSnapshot", reflect.TypeOf((*MockStore)(nil).FetchRuleSetSnapshot), arg0)
}

// UpdateMappingRule mocks base method.
func (m *MockStore) UpdateMappingRule(arg0, arg1 string, arg2 view.MappingRule, arg3 UpdateOptions) (view.MappingRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMappingRule", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(view.MappingRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMappingRule indicates an expected call of UpdateMappingRule.
func (mr *MockStoreMockRecorder) UpdateMappingRule(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMappingRule", reflect.TypeOf((*MockStore)(nil).UpdateMappingRule), arg0, arg1, arg2, arg3)
}

// UpdateRollupRule mocks base method.
func (m *MockStore) UpdateRollupRule(arg0, arg1 string, arg2 view.RollupRule, arg3 UpdateOptions) (view.RollupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRollupRule", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(view.RollupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRollupRule indicates an expected call of UpdateRollupRule.
func (mr *MockStoreMockRecorder) UpdateRollupRule(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRollupRule", reflect.TypeOf((*MockStore)(nil).UpdateRollupRule), arg0, arg1, arg2, arg3)
}

// UpdateRuleSet mocks base method.
func (m *MockStore) UpdateRuleSet(arg0 changes.RuleSetChanges, arg1 int, arg2 UpdateOptions) (view.RuleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRuleSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(view.RuleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRuleSet indicates an expected call of UpdateRuleSet.
func (mr *MockStoreMockRecorder) UpdateRuleSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRuleSet", reflect.TypeOf((*MockStore)(nil).UpdateRuleSet), arg0, arg1, arg2)
}

// ValidateRuleSet mocks base method.
func (m *MockStore) ValidateRuleSet(arg0 view.RuleSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRuleSet", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateRuleSet indicates an expected call of ValidateRuleSet.
func (mr *MockStoreMockRecorder) ValidateRuleSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRuleSet", reflect.TypeOf((*MockStore)(nil).ValidateRuleSet), arg0)
}
