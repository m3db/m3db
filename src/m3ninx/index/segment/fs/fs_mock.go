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
// Source: github.com/m3db/m3ninx/index/segment/fs (interfaces: Writer,Segment)

// Package fs is a generated GoMock package.
package fs

import (
	"io"
	"reflect"
	"regexp"

	"github.com/m3db/m3ninx/doc"
	"github.com/m3db/m3ninx/index"
	"github.com/m3db/m3ninx/index/segment"
	"github.com/m3db/m3ninx/postings"

	"github.com/golang/mock/gomock"
)

// MockWriter is a mock of Writer interface
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// MajorVersion mocks base method
func (m *MockWriter) MajorVersion() int {
	ret := m.ctrl.Call(m, "MajorVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// MajorVersion indicates an expected call of MajorVersion
func (mr *MockWriterMockRecorder) MajorVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MajorVersion", reflect.TypeOf((*MockWriter)(nil).MajorVersion))
}

// Metadata mocks base method
func (m *MockWriter) Metadata() []byte {
	ret := m.ctrl.Call(m, "Metadata")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Metadata indicates an expected call of Metadata
func (mr *MockWriterMockRecorder) Metadata() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockWriter)(nil).Metadata))
}

// MinorVersion mocks base method
func (m *MockWriter) MinorVersion() int {
	ret := m.ctrl.Call(m, "MinorVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// MinorVersion indicates an expected call of MinorVersion
func (mr *MockWriterMockRecorder) MinorVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinorVersion", reflect.TypeOf((*MockWriter)(nil).MinorVersion))
}

// Reset mocks base method
func (m *MockWriter) Reset(arg0 segment.MutableSegment) error {
	ret := m.ctrl.Call(m, "Reset", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset
func (mr *MockWriterMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockWriter)(nil).Reset), arg0)
}

// WriteDocumentsData mocks base method
func (m *MockWriter) WriteDocumentsData(arg0 io.Writer) error {
	ret := m.ctrl.Call(m, "WriteDocumentsData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteDocumentsData indicates an expected call of WriteDocumentsData
func (mr *MockWriterMockRecorder) WriteDocumentsData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteDocumentsData", reflect.TypeOf((*MockWriter)(nil).WriteDocumentsData), arg0)
}

// WriteDocumentsIndex mocks base method
func (m *MockWriter) WriteDocumentsIndex(arg0 io.Writer) error {
	ret := m.ctrl.Call(m, "WriteDocumentsIndex", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteDocumentsIndex indicates an expected call of WriteDocumentsIndex
func (mr *MockWriterMockRecorder) WriteDocumentsIndex(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteDocumentsIndex", reflect.TypeOf((*MockWriter)(nil).WriteDocumentsIndex), arg0)
}

// WriteFSTFields mocks base method
func (m *MockWriter) WriteFSTFields(arg0 io.Writer) error {
	ret := m.ctrl.Call(m, "WriteFSTFields", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFSTFields indicates an expected call of WriteFSTFields
func (mr *MockWriterMockRecorder) WriteFSTFields(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFSTFields", reflect.TypeOf((*MockWriter)(nil).WriteFSTFields), arg0)
}

// WriteFSTTerms mocks base method
func (m *MockWriter) WriteFSTTerms(arg0 io.Writer) error {
	ret := m.ctrl.Call(m, "WriteFSTTerms", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFSTTerms indicates an expected call of WriteFSTTerms
func (mr *MockWriterMockRecorder) WriteFSTTerms(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFSTTerms", reflect.TypeOf((*MockWriter)(nil).WriteFSTTerms), arg0)
}

// WritePostingsOffsets mocks base method
func (m *MockWriter) WritePostingsOffsets(arg0 io.Writer) error {
	ret := m.ctrl.Call(m, "WritePostingsOffsets", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WritePostingsOffsets indicates an expected call of WritePostingsOffsets
func (mr *MockWriterMockRecorder) WritePostingsOffsets(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePostingsOffsets", reflect.TypeOf((*MockWriter)(nil).WritePostingsOffsets), arg0)
}

// MockSegment is a mock of Segment interface
type MockSegment struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentMockRecorder
}

// MockSegmentMockRecorder is the mock recorder for MockSegment
type MockSegmentMockRecorder struct {
	mock *MockSegment
}

// NewMockSegment creates a new mock instance
func NewMockSegment(ctrl *gomock.Controller) *MockSegment {
	mock := &MockSegment{ctrl: ctrl}
	mock.recorder = &MockSegmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSegment) EXPECT() *MockSegmentMockRecorder {
	return m.recorder
}

// AllDocs mocks base method
func (m *MockSegment) AllDocs() (index.IDDocIterator, error) {
	ret := m.ctrl.Call(m, "AllDocs")
	ret0, _ := ret[0].(index.IDDocIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllDocs indicates an expected call of AllDocs
func (mr *MockSegmentMockRecorder) AllDocs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllDocs", reflect.TypeOf((*MockSegment)(nil).AllDocs))
}

// Close mocks base method
func (m *MockSegment) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockSegmentMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSegment)(nil).Close))
}

// ContainsID mocks base method
func (m *MockSegment) ContainsID(arg0 []byte) (bool, error) {
	ret := m.ctrl.Call(m, "ContainsID", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainsID indicates an expected call of ContainsID
func (mr *MockSegmentMockRecorder) ContainsID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsID", reflect.TypeOf((*MockSegment)(nil).ContainsID), arg0)
}

// Doc mocks base method
func (m *MockSegment) Doc(arg0 postings.ID) (doc.Document, error) {
	ret := m.ctrl.Call(m, "Doc", arg0)
	ret0, _ := ret[0].(doc.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Doc indicates an expected call of Doc
func (mr *MockSegmentMockRecorder) Doc(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Doc", reflect.TypeOf((*MockSegment)(nil).Doc), arg0)
}

// Docs mocks base method
func (m *MockSegment) Docs(arg0 postings.List) (doc.Iterator, error) {
	ret := m.ctrl.Call(m, "Docs", arg0)
	ret0, _ := ret[0].(doc.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Docs indicates an expected call of Docs
func (mr *MockSegmentMockRecorder) Docs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Docs", reflect.TypeOf((*MockSegment)(nil).Docs), arg0)
}

// Fields mocks base method
func (m *MockSegment) Fields() ([][]byte, error) {
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fields indicates an expected call of Fields
func (mr *MockSegmentMockRecorder) Fields() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MockSegment)(nil).Fields))
}

// MatchAll mocks base method
func (m *MockSegment) MatchAll() (postings.MutableList, error) {
	ret := m.ctrl.Call(m, "MatchAll")
	ret0, _ := ret[0].(postings.MutableList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchAll indicates an expected call of MatchAll
func (mr *MockSegmentMockRecorder) MatchAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchAll", reflect.TypeOf((*MockSegment)(nil).MatchAll))
}

// MatchRegexp mocks base method
func (m *MockSegment) MatchRegexp(arg0, arg1 []byte, arg2 *regexp.Regexp) (postings.List, error) {
	ret := m.ctrl.Call(m, "MatchRegexp", arg0, arg1, arg2)
	ret0, _ := ret[0].(postings.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchRegexp indicates an expected call of MatchRegexp
func (mr *MockSegmentMockRecorder) MatchRegexp(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchRegexp", reflect.TypeOf((*MockSegment)(nil).MatchRegexp), arg0, arg1, arg2)
}

// MatchTerm mocks base method
func (m *MockSegment) MatchTerm(arg0, arg1 []byte) (postings.List, error) {
	ret := m.ctrl.Call(m, "MatchTerm", arg0, arg1)
	ret0, _ := ret[0].(postings.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchTerm indicates an expected call of MatchTerm
func (mr *MockSegmentMockRecorder) MatchTerm(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchTerm", reflect.TypeOf((*MockSegment)(nil).MatchTerm), arg0, arg1)
}

// Reader mocks base method
func (m *MockSegment) Reader() (index.Reader, error) {
	ret := m.ctrl.Call(m, "Reader")
	ret0, _ := ret[0].(index.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reader indicates an expected call of Reader
func (mr *MockSegmentMockRecorder) Reader() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockSegment)(nil).Reader))
}

// Size mocks base method
func (m *MockSegment) Size() int64 {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockSegmentMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockSegment)(nil).Size))
}

// Terms mocks base method
func (m *MockSegment) Terms(arg0 []byte) ([][]byte, error) {
	ret := m.ctrl.Call(m, "Terms", arg0)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Terms indicates an expected call of Terms
func (mr *MockSegmentMockRecorder) Terms(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terms", reflect.TypeOf((*MockSegment)(nil).Terms), arg0)
}
