// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/ts/writes/types.go

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

// Package writes is a generated GoMock package.
package writes

import (
	"reflect"
	"time"

	"github.com/m3db/m3/src/dbnode/ts"
	"github.com/m3db/m3/src/x/checked"
	"github.com/m3db/m3/src/x/ident"
	time0 "github.com/m3db/m3/src/x/time"

	"github.com/golang/mock/gomock"
)

// MockWriteBatch is a mock of WriteBatch interface
type MockWriteBatch struct {
	ctrl     *gomock.Controller
	recorder *MockWriteBatchMockRecorder
}

// MockWriteBatchMockRecorder is the mock recorder for MockWriteBatch
type MockWriteBatchMockRecorder struct {
	mock *MockWriteBatch
}

// NewMockWriteBatch creates a new mock instance
func NewMockWriteBatch(ctrl *gomock.Controller) *MockWriteBatch {
	mock := &MockWriteBatch{ctrl: ctrl}
	mock.recorder = &MockWriteBatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriteBatch) EXPECT() *MockWriteBatchMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockWriteBatch) Add(originalIndex int, id ident.ID, timestamp time.Time, value float64, unit time0.Unit, annotation checked.Bytes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", originalIndex, id, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockWriteBatchMockRecorder) Add(originalIndex, id, timestamp, value, unit, annotation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockWriteBatch)(nil).Add), originalIndex, id, timestamp, value, unit, annotation)
}

// AddTagged mocks base method
func (m *MockWriteBatch) AddTagged(originalIndex int, id ident.ID, tags ident.TagIterator, encodedTags checked.Bytes, timestamp time.Time, value float64, unit time0.Unit, annotation checked.Bytes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTagged", originalIndex, id, tags, encodedTags, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTagged indicates an expected call of AddTagged
func (mr *MockWriteBatchMockRecorder) AddTagged(originalIndex, id, tags, encodedTags, timestamp, value, unit, annotation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTagged", reflect.TypeOf((*MockWriteBatch)(nil).AddTagged), originalIndex, id, tags, encodedTags, timestamp, value, unit, annotation)
}

// SetFinalizeEncodedTagsFn mocks base method
func (m *MockWriteBatch) SetFinalizeEncodedTagsFn(f FinalizeEncodedTagsFn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizeEncodedTagsFn", f)
}

// SetFinalizeEncodedTagsFn indicates an expected call of SetFinalizeEncodedTagsFn
func (mr *MockWriteBatchMockRecorder) SetFinalizeEncodedTagsFn(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizeEncodedTagsFn", reflect.TypeOf((*MockWriteBatch)(nil).SetFinalizeEncodedTagsFn), f)
}

// SetFinalizeAnnotationFn mocks base method
func (m *MockWriteBatch) SetFinalizeAnnotationFn(f FinalizeAnnotationFn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizeAnnotationFn", f)
}

// SetFinalizeAnnotationFn indicates an expected call of SetFinalizeAnnotationFn
func (mr *MockWriteBatchMockRecorder) SetFinalizeAnnotationFn(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizeAnnotationFn", reflect.TypeOf((*MockWriteBatch)(nil).SetFinalizeAnnotationFn), f)
}

// Iter mocks base method
func (m *MockWriteBatch) Iter() []BatchWrite {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].([]BatchWrite)
	return ret0
}

// Iter indicates an expected call of Iter
func (mr *MockWriteBatchMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockWriteBatch)(nil).Iter))
}

// SetPendingIndex mocks base method
func (m *MockWriteBatch) SetPendingIndex(idx int, pending PendingIndexInsert) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPendingIndex", idx, pending)
}

// SetPendingIndex indicates an expected call of SetPendingIndex
func (mr *MockWriteBatchMockRecorder) SetPendingIndex(idx, pending interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPendingIndex", reflect.TypeOf((*MockWriteBatch)(nil).SetPendingIndex), idx, pending)
}

// PendingIndex mocks base method
func (m *MockWriteBatch) PendingIndex() []PendingIndexInsert {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingIndex")
	ret0, _ := ret[0].([]PendingIndexInsert)
	return ret0
}

// PendingIndex indicates an expected call of PendingIndex
func (mr *MockWriteBatchMockRecorder) PendingIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingIndex", reflect.TypeOf((*MockWriteBatch)(nil).PendingIndex))
}

// SetError mocks base method
func (m *MockWriteBatch) SetError(idx int, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetError", idx, err)
}

// SetError indicates an expected call of SetError
func (mr *MockWriteBatchMockRecorder) SetError(idx, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetError", reflect.TypeOf((*MockWriteBatch)(nil).SetError), idx, err)
}

// SetSeries mocks base method
func (m *MockWriteBatch) SetSeries(idx int, series ts.Series) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSeries", idx, series)
}

// SetSeries indicates an expected call of SetSeries
func (mr *MockWriteBatchMockRecorder) SetSeries(idx, series interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSeries", reflect.TypeOf((*MockWriteBatch)(nil).SetSeries), idx, series)
}

// SetSkipWrite mocks base method
func (m *MockWriteBatch) SetSkipWrite(idx int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSkipWrite", idx)
}

// SetSkipWrite indicates an expected call of SetSkipWrite
func (mr *MockWriteBatchMockRecorder) SetSkipWrite(idx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSkipWrite", reflect.TypeOf((*MockWriteBatch)(nil).SetSkipWrite), idx)
}

// Reset mocks base method
func (m *MockWriteBatch) Reset(batchSize int, ns ident.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", batchSize, ns)
}

// Reset indicates an expected call of Reset
func (mr *MockWriteBatchMockRecorder) Reset(batchSize, ns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockWriteBatch)(nil).Reset), batchSize, ns)
}

// Finalize mocks base method
func (m *MockWriteBatch) Finalize() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (mr *MockWriteBatchMockRecorder) Finalize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockWriteBatch)(nil).Finalize))
}

// cap mocks base method
func (m *MockWriteBatch) cap() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cap")
	ret0, _ := ret[0].(int)
	return ret0
}

// cap indicates an expected call of cap
func (mr *MockWriteBatchMockRecorder) cap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cap", reflect.TypeOf((*MockWriteBatch)(nil).cap))
}

// MockBatchWriter is a mock of BatchWriter interface
type MockBatchWriter struct {
	ctrl     *gomock.Controller
	recorder *MockBatchWriterMockRecorder
}

// MockBatchWriterMockRecorder is the mock recorder for MockBatchWriter
type MockBatchWriterMockRecorder struct {
	mock *MockBatchWriter
}

// NewMockBatchWriter creates a new mock instance
func NewMockBatchWriter(ctrl *gomock.Controller) *MockBatchWriter {
	mock := &MockBatchWriter{ctrl: ctrl}
	mock.recorder = &MockBatchWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBatchWriter) EXPECT() *MockBatchWriterMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockBatchWriter) Add(originalIndex int, id ident.ID, timestamp time.Time, value float64, unit time0.Unit, annotation checked.Bytes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", originalIndex, id, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockBatchWriterMockRecorder) Add(originalIndex, id, timestamp, value, unit, annotation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockBatchWriter)(nil).Add), originalIndex, id, timestamp, value, unit, annotation)
}

// AddTagged mocks base method
func (m *MockBatchWriter) AddTagged(originalIndex int, id ident.ID, tags ident.TagIterator, encodedTags checked.Bytes, timestamp time.Time, value float64, unit time0.Unit, annotation checked.Bytes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTagged", originalIndex, id, tags, encodedTags, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTagged indicates an expected call of AddTagged
func (mr *MockBatchWriterMockRecorder) AddTagged(originalIndex, id, tags, encodedTags, timestamp, value, unit, annotation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTagged", reflect.TypeOf((*MockBatchWriter)(nil).AddTagged), originalIndex, id, tags, encodedTags, timestamp, value, unit, annotation)
}

// SetFinalizeEncodedTagsFn mocks base method
func (m *MockBatchWriter) SetFinalizeEncodedTagsFn(f FinalizeEncodedTagsFn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizeEncodedTagsFn", f)
}

// SetFinalizeEncodedTagsFn indicates an expected call of SetFinalizeEncodedTagsFn
func (mr *MockBatchWriterMockRecorder) SetFinalizeEncodedTagsFn(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizeEncodedTagsFn", reflect.TypeOf((*MockBatchWriter)(nil).SetFinalizeEncodedTagsFn), f)
}

// SetFinalizeAnnotationFn mocks base method
func (m *MockBatchWriter) SetFinalizeAnnotationFn(f FinalizeAnnotationFn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizeAnnotationFn", f)
}

// SetFinalizeAnnotationFn indicates an expected call of SetFinalizeAnnotationFn
func (mr *MockBatchWriterMockRecorder) SetFinalizeAnnotationFn(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizeAnnotationFn", reflect.TypeOf((*MockBatchWriter)(nil).SetFinalizeAnnotationFn), f)
}
