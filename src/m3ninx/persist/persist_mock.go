// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/m3ninx/persist/types.go

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

// Package persist is a generated GoMock package.
package persist

import (
	"io"
	"reflect"

	"github.com/m3db/m3/src/m3ninx/index/segment"

	"github.com/golang/mock/gomock"
)

// MockIndexFileSetWriter is a mock of IndexFileSetWriter interface
type MockIndexFileSetWriter struct {
	ctrl     *gomock.Controller
	recorder *MockIndexFileSetWriterMockRecorder
}

// MockIndexFileSetWriterMockRecorder is the mock recorder for MockIndexFileSetWriter
type MockIndexFileSetWriterMockRecorder struct {
	mock *MockIndexFileSetWriter
}

// NewMockIndexFileSetWriter creates a new mock instance
func NewMockIndexFileSetWriter(ctrl *gomock.Controller) *MockIndexFileSetWriter {
	mock := &MockIndexFileSetWriter{ctrl: ctrl}
	mock.recorder = &MockIndexFileSetWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexFileSetWriter) EXPECT() *MockIndexFileSetWriterMockRecorder {
	return m.recorder
}

// WriteSegmentFileSet mocks base method
func (m *MockIndexFileSetWriter) WriteSegmentFileSet(segmentFileSet IndexSegmentFileSetWriter) error {
	ret := m.ctrl.Call(m, "WriteSegmentFileSet", segmentFileSet)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteSegmentFileSet indicates an expected call of WriteSegmentFileSet
func (mr *MockIndexFileSetWriterMockRecorder) WriteSegmentFileSet(segmentFileSet interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteSegmentFileSet", reflect.TypeOf((*MockIndexFileSetWriter)(nil).WriteSegmentFileSet), segmentFileSet)
}

// MockIndexSegmentFileSetWriter is a mock of IndexSegmentFileSetWriter interface
type MockIndexSegmentFileSetWriter struct {
	ctrl     *gomock.Controller
	recorder *MockIndexSegmentFileSetWriterMockRecorder
}

// MockIndexSegmentFileSetWriterMockRecorder is the mock recorder for MockIndexSegmentFileSetWriter
type MockIndexSegmentFileSetWriterMockRecorder struct {
	mock *MockIndexSegmentFileSetWriter
}

// NewMockIndexSegmentFileSetWriter creates a new mock instance
func NewMockIndexSegmentFileSetWriter(ctrl *gomock.Controller) *MockIndexSegmentFileSetWriter {
	mock := &MockIndexSegmentFileSetWriter{ctrl: ctrl}
	mock.recorder = &MockIndexSegmentFileSetWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexSegmentFileSetWriter) EXPECT() *MockIndexSegmentFileSetWriterMockRecorder {
	return m.recorder
}

// SegmentType mocks base method
func (m *MockIndexSegmentFileSetWriter) SegmentType() IndexSegmentType {
	ret := m.ctrl.Call(m, "SegmentType")
	ret0, _ := ret[0].(IndexSegmentType)
	return ret0
}

// SegmentType indicates an expected call of SegmentType
func (mr *MockIndexSegmentFileSetWriterMockRecorder) SegmentType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SegmentType", reflect.TypeOf((*MockIndexSegmentFileSetWriter)(nil).SegmentType))
}

// MajorVersion mocks base method
func (m *MockIndexSegmentFileSetWriter) MajorVersion() int {
	ret := m.ctrl.Call(m, "MajorVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// MajorVersion indicates an expected call of MajorVersion
func (mr *MockIndexSegmentFileSetWriterMockRecorder) MajorVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MajorVersion", reflect.TypeOf((*MockIndexSegmentFileSetWriter)(nil).MajorVersion))
}

// MinorVersion mocks base method
func (m *MockIndexSegmentFileSetWriter) MinorVersion() int {
	ret := m.ctrl.Call(m, "MinorVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// MinorVersion indicates an expected call of MinorVersion
func (mr *MockIndexSegmentFileSetWriterMockRecorder) MinorVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinorVersion", reflect.TypeOf((*MockIndexSegmentFileSetWriter)(nil).MinorVersion))
}

// SegmentMetadata mocks base method
func (m *MockIndexSegmentFileSetWriter) SegmentMetadata() []byte {
	ret := m.ctrl.Call(m, "SegmentMetadata")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// SegmentMetadata indicates an expected call of SegmentMetadata
func (mr *MockIndexSegmentFileSetWriterMockRecorder) SegmentMetadata() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SegmentMetadata", reflect.TypeOf((*MockIndexSegmentFileSetWriter)(nil).SegmentMetadata))
}

// Files mocks base method
func (m *MockIndexSegmentFileSetWriter) Files() []IndexSegmentFileType {
	ret := m.ctrl.Call(m, "Files")
	ret0, _ := ret[0].([]IndexSegmentFileType)
	return ret0
}

// Files indicates an expected call of Files
func (mr *MockIndexSegmentFileSetWriterMockRecorder) Files() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Files", reflect.TypeOf((*MockIndexSegmentFileSetWriter)(nil).Files))
}

// WriteFile mocks base method
func (m *MockIndexSegmentFileSetWriter) WriteFile(fileType IndexSegmentFileType, writer io.Writer) error {
	ret := m.ctrl.Call(m, "WriteFile", fileType, writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile
func (mr *MockIndexSegmentFileSetWriterMockRecorder) WriteFile(fileType, writer interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockIndexSegmentFileSetWriter)(nil).WriteFile), fileType, writer)
}

// MockMutableSegmentFileSetWriter is a mock of MutableSegmentFileSetWriter interface
type MockMutableSegmentFileSetWriter struct {
	ctrl     *gomock.Controller
	recorder *MockMutableSegmentFileSetWriterMockRecorder
}

// MockMutableSegmentFileSetWriterMockRecorder is the mock recorder for MockMutableSegmentFileSetWriter
type MockMutableSegmentFileSetWriterMockRecorder struct {
	mock *MockMutableSegmentFileSetWriter
}

// NewMockMutableSegmentFileSetWriter creates a new mock instance
func NewMockMutableSegmentFileSetWriter(ctrl *gomock.Controller) *MockMutableSegmentFileSetWriter {
	mock := &MockMutableSegmentFileSetWriter{ctrl: ctrl}
	mock.recorder = &MockMutableSegmentFileSetWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMutableSegmentFileSetWriter) EXPECT() *MockMutableSegmentFileSetWriterMockRecorder {
	return m.recorder
}

// SegmentType mocks base method
func (m *MockMutableSegmentFileSetWriter) SegmentType() IndexSegmentType {
	ret := m.ctrl.Call(m, "SegmentType")
	ret0, _ := ret[0].(IndexSegmentType)
	return ret0
}

// SegmentType indicates an expected call of SegmentType
func (mr *MockMutableSegmentFileSetWriterMockRecorder) SegmentType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SegmentType", reflect.TypeOf((*MockMutableSegmentFileSetWriter)(nil).SegmentType))
}

// MajorVersion mocks base method
func (m *MockMutableSegmentFileSetWriter) MajorVersion() int {
	ret := m.ctrl.Call(m, "MajorVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// MajorVersion indicates an expected call of MajorVersion
func (mr *MockMutableSegmentFileSetWriterMockRecorder) MajorVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MajorVersion", reflect.TypeOf((*MockMutableSegmentFileSetWriter)(nil).MajorVersion))
}

// MinorVersion mocks base method
func (m *MockMutableSegmentFileSetWriter) MinorVersion() int {
	ret := m.ctrl.Call(m, "MinorVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// MinorVersion indicates an expected call of MinorVersion
func (mr *MockMutableSegmentFileSetWriterMockRecorder) MinorVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinorVersion", reflect.TypeOf((*MockMutableSegmentFileSetWriter)(nil).MinorVersion))
}

// SegmentMetadata mocks base method
func (m *MockMutableSegmentFileSetWriter) SegmentMetadata() []byte {
	ret := m.ctrl.Call(m, "SegmentMetadata")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// SegmentMetadata indicates an expected call of SegmentMetadata
func (mr *MockMutableSegmentFileSetWriterMockRecorder) SegmentMetadata() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SegmentMetadata", reflect.TypeOf((*MockMutableSegmentFileSetWriter)(nil).SegmentMetadata))
}

// Files mocks base method
func (m *MockMutableSegmentFileSetWriter) Files() []IndexSegmentFileType {
	ret := m.ctrl.Call(m, "Files")
	ret0, _ := ret[0].([]IndexSegmentFileType)
	return ret0
}

// Files indicates an expected call of Files
func (mr *MockMutableSegmentFileSetWriterMockRecorder) Files() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Files", reflect.TypeOf((*MockMutableSegmentFileSetWriter)(nil).Files))
}

// WriteFile mocks base method
func (m *MockMutableSegmentFileSetWriter) WriteFile(fileType IndexSegmentFileType, writer io.Writer) error {
	ret := m.ctrl.Call(m, "WriteFile", fileType, writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile
func (mr *MockMutableSegmentFileSetWriterMockRecorder) WriteFile(fileType, writer interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockMutableSegmentFileSetWriter)(nil).WriteFile), fileType, writer)
}

// Reset mocks base method
func (m *MockMutableSegmentFileSetWriter) Reset(arg0 segment.MutableSegment) error {
	ret := m.ctrl.Call(m, "Reset", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset
func (mr *MockMutableSegmentFileSetWriterMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockMutableSegmentFileSetWriter)(nil).Reset), arg0)
}

// MockIndexFileSetReader is a mock of IndexFileSetReader interface
type MockIndexFileSetReader struct {
	ctrl     *gomock.Controller
	recorder *MockIndexFileSetReaderMockRecorder
}

// MockIndexFileSetReaderMockRecorder is the mock recorder for MockIndexFileSetReader
type MockIndexFileSetReaderMockRecorder struct {
	mock *MockIndexFileSetReader
}

// NewMockIndexFileSetReader creates a new mock instance
func NewMockIndexFileSetReader(ctrl *gomock.Controller) *MockIndexFileSetReader {
	mock := &MockIndexFileSetReader{ctrl: ctrl}
	mock.recorder = &MockIndexFileSetReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexFileSetReader) EXPECT() *MockIndexFileSetReaderMockRecorder {
	return m.recorder
}

// SegmentFileSets mocks base method
func (m *MockIndexFileSetReader) SegmentFileSets() int {
	ret := m.ctrl.Call(m, "SegmentFileSets")
	ret0, _ := ret[0].(int)
	return ret0
}

// SegmentFileSets indicates an expected call of SegmentFileSets
func (mr *MockIndexFileSetReaderMockRecorder) SegmentFileSets() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SegmentFileSets", reflect.TypeOf((*MockIndexFileSetReader)(nil).SegmentFileSets))
}

// ReadSegmentFileSet mocks base method
func (m *MockIndexFileSetReader) ReadSegmentFileSet() (IndexSegmentFileSet, error) {
	ret := m.ctrl.Call(m, "ReadSegmentFileSet")
	ret0, _ := ret[0].(IndexSegmentFileSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadSegmentFileSet indicates an expected call of ReadSegmentFileSet
func (mr *MockIndexFileSetReaderMockRecorder) ReadSegmentFileSet() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSegmentFileSet", reflect.TypeOf((*MockIndexFileSetReader)(nil).ReadSegmentFileSet))
}

// MockIndexSegmentFileSet is a mock of IndexSegmentFileSet interface
type MockIndexSegmentFileSet struct {
	ctrl     *gomock.Controller
	recorder *MockIndexSegmentFileSetMockRecorder
}

// MockIndexSegmentFileSetMockRecorder is the mock recorder for MockIndexSegmentFileSet
type MockIndexSegmentFileSetMockRecorder struct {
	mock *MockIndexSegmentFileSet
}

// NewMockIndexSegmentFileSet creates a new mock instance
func NewMockIndexSegmentFileSet(ctrl *gomock.Controller) *MockIndexSegmentFileSet {
	mock := &MockIndexSegmentFileSet{ctrl: ctrl}
	mock.recorder = &MockIndexSegmentFileSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexSegmentFileSet) EXPECT() *MockIndexSegmentFileSetMockRecorder {
	return m.recorder
}

// SegmentType mocks base method
func (m *MockIndexSegmentFileSet) SegmentType() IndexSegmentType {
	ret := m.ctrl.Call(m, "SegmentType")
	ret0, _ := ret[0].(IndexSegmentType)
	return ret0
}

// SegmentType indicates an expected call of SegmentType
func (mr *MockIndexSegmentFileSetMockRecorder) SegmentType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SegmentType", reflect.TypeOf((*MockIndexSegmentFileSet)(nil).SegmentType))
}

// MajorVersion mocks base method
func (m *MockIndexSegmentFileSet) MajorVersion() int {
	ret := m.ctrl.Call(m, "MajorVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// MajorVersion indicates an expected call of MajorVersion
func (mr *MockIndexSegmentFileSetMockRecorder) MajorVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MajorVersion", reflect.TypeOf((*MockIndexSegmentFileSet)(nil).MajorVersion))
}

// MinorVersion mocks base method
func (m *MockIndexSegmentFileSet) MinorVersion() int {
	ret := m.ctrl.Call(m, "MinorVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// MinorVersion indicates an expected call of MinorVersion
func (mr *MockIndexSegmentFileSetMockRecorder) MinorVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinorVersion", reflect.TypeOf((*MockIndexSegmentFileSet)(nil).MinorVersion))
}

// SegmentMetadata mocks base method
func (m *MockIndexSegmentFileSet) SegmentMetadata() []byte {
	ret := m.ctrl.Call(m, "SegmentMetadata")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// SegmentMetadata indicates an expected call of SegmentMetadata
func (mr *MockIndexSegmentFileSetMockRecorder) SegmentMetadata() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SegmentMetadata", reflect.TypeOf((*MockIndexSegmentFileSet)(nil).SegmentMetadata))
}

// Files mocks base method
func (m *MockIndexSegmentFileSet) Files() []IndexSegmentFile {
	ret := m.ctrl.Call(m, "Files")
	ret0, _ := ret[0].([]IndexSegmentFile)
	return ret0
}

// Files indicates an expected call of Files
func (mr *MockIndexSegmentFileSetMockRecorder) Files() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Files", reflect.TypeOf((*MockIndexSegmentFileSet)(nil).Files))
}

// MockIndexSegmentFile is a mock of IndexSegmentFile interface
type MockIndexSegmentFile struct {
	ctrl     *gomock.Controller
	recorder *MockIndexSegmentFileMockRecorder
}

// MockIndexSegmentFileMockRecorder is the mock recorder for MockIndexSegmentFile
type MockIndexSegmentFileMockRecorder struct {
	mock *MockIndexSegmentFile
}

// NewMockIndexSegmentFile creates a new mock instance
func NewMockIndexSegmentFile(ctrl *gomock.Controller) *MockIndexSegmentFile {
	mock := &MockIndexSegmentFile{ctrl: ctrl}
	mock.recorder = &MockIndexSegmentFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexSegmentFile) EXPECT() *MockIndexSegmentFileMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockIndexSegmentFile) Read(p []byte) (int, error) {
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockIndexSegmentFileMockRecorder) Read(p interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockIndexSegmentFile)(nil).Read), p)
}

// Close mocks base method
func (m *MockIndexSegmentFile) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIndexSegmentFileMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIndexSegmentFile)(nil).Close))
}

// SegmentFileType mocks base method
func (m *MockIndexSegmentFile) SegmentFileType() IndexSegmentFileType {
	ret := m.ctrl.Call(m, "SegmentFileType")
	ret0, _ := ret[0].(IndexSegmentFileType)
	return ret0
}

// SegmentFileType indicates an expected call of SegmentFileType
func (mr *MockIndexSegmentFileMockRecorder) SegmentFileType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SegmentFileType", reflect.TypeOf((*MockIndexSegmentFile)(nil).SegmentFileType))
}

// Bytes mocks base method
func (m *MockIndexSegmentFile) Bytes() ([]byte, error) {
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bytes indicates an expected call of Bytes
func (mr *MockIndexSegmentFileMockRecorder) Bytes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockIndexSegmentFile)(nil).Bytes))
}
