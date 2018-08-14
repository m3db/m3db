// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/digest (interfaces: ReaderWithDigest)

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

// Package digest is a generated GoMock package.
package digest

import (
	"hash"
	"io"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockReaderWithDigest is a mock of ReaderWithDigest interface
type MockReaderWithDigest struct {
	ctrl     *gomock.Controller
	recorder *MockReaderWithDigestMockRecorder
}

// MockReaderWithDigestMockRecorder is the mock recorder for MockReaderWithDigest
type MockReaderWithDigestMockRecorder struct {
	mock *MockReaderWithDigest
}

// NewMockReaderWithDigest creates a new mock instance
func NewMockReaderWithDigest(ctrl *gomock.Controller) *MockReaderWithDigest {
	mock := &MockReaderWithDigest{ctrl: ctrl}
	mock.recorder = &MockReaderWithDigestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReaderWithDigest) EXPECT() *MockReaderWithDigestMockRecorder {
	return m.recorder
}

// Digest mocks base method
func (m *MockReaderWithDigest) Digest() hash.Hash32 {
	ret := m.ctrl.Call(m, "Digest")
	ret0, _ := ret[0].(hash.Hash32)
	return ret0
}

// Digest indicates an expected call of Digest
func (mr *MockReaderWithDigestMockRecorder) Digest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Digest", reflect.TypeOf((*MockReaderWithDigest)(nil).Digest))
}

// Read mocks base method
func (m *MockReaderWithDigest) Read(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReaderWithDigestMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReaderWithDigest)(nil).Read), arg0)
}

// Reset mocks base method
func (m *MockReaderWithDigest) Reset(arg0 io.Reader) {
	m.ctrl.Call(m, "Reset", arg0)
}

// Reset indicates an expected call of Reset
func (mr *MockReaderWithDigestMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockReaderWithDigest)(nil).Reset), arg0)
}

// Validate mocks base method
func (m *MockReaderWithDigest) Validate(arg0 uint32) error {
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockReaderWithDigestMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockReaderWithDigest)(nil).Validate), arg0)
}
