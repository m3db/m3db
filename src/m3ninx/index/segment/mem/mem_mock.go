// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/m3ninx/index/segment/mem (interfaces: ReadableSegment)

// Copyright (c) 2020 Uber Technologies, Inc.
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

// Package mem is a generated GoMock package.
package mem

import (
	"reflect"
	"regexp"

	"github.com/m3db/m3/src/m3ninx/doc"
	"github.com/m3db/m3/src/m3ninx/index/segment"
	"github.com/m3db/m3/src/m3ninx/postings"

	"github.com/golang/mock/gomock"
)

// MockReadableSegment is a mock of ReadableSegment interface
type MockReadableSegment struct {
	ctrl     *gomock.Controller
	recorder *MockReadableSegmentMockRecorder
}

// MockReadableSegmentMockRecorder is the mock recorder for MockReadableSegment
type MockReadableSegmentMockRecorder struct {
	mock *MockReadableSegment
}

// NewMockReadableSegment creates a new mock instance
func NewMockReadableSegment(ctrl *gomock.Controller) *MockReadableSegment {
	mock := &MockReadableSegment{ctrl: ctrl}
	mock.recorder = &MockReadableSegmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadableSegment) EXPECT() *MockReadableSegmentMockRecorder {
	return m.recorder
}

// ContainsField mocks base method
func (m *MockReadableSegment) ContainsField(arg0 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsField", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainsField indicates an expected call of ContainsField
func (mr *MockReadableSegmentMockRecorder) ContainsField(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsField", reflect.TypeOf((*MockReadableSegment)(nil).ContainsField), arg0)
}

// Fields mocks base method
func (m *MockReadableSegment) Fields() (segment.FieldsIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].(segment.FieldsIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fields indicates an expected call of Fields
func (mr *MockReadableSegmentMockRecorder) Fields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MockReadableSegment)(nil).Fields))
}

// Terms mocks base method
func (m *MockReadableSegment) Terms(arg0 []byte) (segment.TermsIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Terms", arg0)
	ret0, _ := ret[0].(segment.TermsIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Terms indicates an expected call of Terms
func (mr *MockReadableSegmentMockRecorder) Terms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terms", reflect.TypeOf((*MockReadableSegment)(nil).Terms), arg0)
}

// getDoc mocks base method
func (m *MockReadableSegment) getDoc(arg0 postings.ID) (doc.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getDoc", arg0)
	ret0, _ := ret[0].(doc.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getDoc indicates an expected call of getDoc
func (mr *MockReadableSegmentMockRecorder) getDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getDoc", reflect.TypeOf((*MockReadableSegment)(nil).getDoc), arg0)
}

// matchRegexp mocks base method
func (m *MockReadableSegment) matchRegexp(arg0 []byte, arg1 *regexp.Regexp) (postings.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "matchRegexp", arg0, arg1)
	ret0, _ := ret[0].(postings.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// matchRegexp indicates an expected call of matchRegexp
func (mr *MockReadableSegmentMockRecorder) matchRegexp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "matchRegexp", reflect.TypeOf((*MockReadableSegment)(nil).matchRegexp), arg0, arg1)
}

// matchTerm mocks base method
func (m *MockReadableSegment) matchTerm(arg0, arg1 []byte) (postings.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "matchTerm", arg0, arg1)
	ret0, _ := ret[0].(postings.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// matchTerm indicates an expected call of matchTerm
func (mr *MockReadableSegmentMockRecorder) matchTerm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "matchTerm", reflect.TypeOf((*MockReadableSegment)(nil).matchTerm), arg0, arg1)
}
