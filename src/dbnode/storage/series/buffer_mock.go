// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/storage/series/buffer.go

// Copyright (c) 2019 Uber Technologies, Inc.
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

// Package series is a generated GoMock package.
package series

import (
	"reflect"
	"time"

	"github.com/m3db/m3/src/dbnode/persist"
	"github.com/m3db/m3/src/dbnode/storage/block"
	"github.com/m3db/m3/src/dbnode/storage/namespace"
	"github.com/m3db/m3/src/dbnode/x/xio"
	"github.com/m3db/m3/src/x/context"
	"github.com/m3db/m3/src/x/ident"
	time0 "github.com/m3db/m3/src/x/time"

	"github.com/golang/mock/gomock"
)

// MockdatabaseBuffer is a mock of databaseBuffer interface
type MockdatabaseBuffer struct {
	ctrl     *gomock.Controller
	recorder *MockdatabaseBufferMockRecorder
}

// MockdatabaseBufferMockRecorder is the mock recorder for MockdatabaseBuffer
type MockdatabaseBufferMockRecorder struct {
	mock *MockdatabaseBuffer
}

// NewMockdatabaseBuffer creates a new mock instance
func NewMockdatabaseBuffer(ctrl *gomock.Controller) *MockdatabaseBuffer {
	mock := &MockdatabaseBuffer{ctrl: ctrl}
	mock.recorder = &MockdatabaseBufferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockdatabaseBuffer) EXPECT() *MockdatabaseBufferMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockdatabaseBuffer) Write(ctx context.Context, timestamp time.Time, value float64, unit time0.Unit, annotation []byte, wOpts WriteOptions) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", ctx, timestamp, value, unit, annotation, wOpts)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockdatabaseBufferMockRecorder) Write(ctx, timestamp, value, unit, annotation, wOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockdatabaseBuffer)(nil).Write), ctx, timestamp, value, unit, annotation, wOpts)
}

// Snapshot mocks base method
func (m *MockdatabaseBuffer) Snapshot(ctx context.Context, blockStart time.Time, nsCtx namespace.Context) (xio.SegmentReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", ctx, blockStart, nsCtx)
	ret0, _ := ret[0].(xio.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot
func (mr *MockdatabaseBufferMockRecorder) Snapshot(ctx, blockStart, nsCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockdatabaseBuffer)(nil).Snapshot), ctx, blockStart, nsCtx)
}

// Flush mocks base method
func (m *MockdatabaseBuffer) Flush(ctx context.Context, blockStart time.Time, id ident.ID, tags ident.Tags, persistFn persist.DataFn, version int, nsCtx namespace.Context) (FlushOutcome, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", ctx, blockStart, id, tags, persistFn, version, nsCtx)
	ret0, _ := ret[0].(FlushOutcome)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Flush indicates an expected call of Flush
func (mr *MockdatabaseBufferMockRecorder) Flush(ctx, blockStart, id, tags, persistFn, version, nsCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockdatabaseBuffer)(nil).Flush), ctx, blockStart, id, tags, persistFn, version, nsCtx)
}

// ReadEncoded mocks base method
func (m *MockdatabaseBuffer) ReadEncoded(ctx context.Context, start, end time.Time, nsCtx namespace.Context) ([][]xio.BlockReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadEncoded", ctx, start, end, nsCtx)
	ret0, _ := ret[0].([][]xio.BlockReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadEncoded indicates an expected call of ReadEncoded
func (mr *MockdatabaseBufferMockRecorder) ReadEncoded(ctx, start, end, nsCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadEncoded", reflect.TypeOf((*MockdatabaseBuffer)(nil).ReadEncoded), ctx, start, end, nsCtx)
}

// FetchBlocks mocks base method
func (m *MockdatabaseBuffer) FetchBlocks(ctx context.Context, starts []time.Time, nsCtx namespace.Context) []block.FetchBlockResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBlocks", ctx, starts, nsCtx)
	ret0, _ := ret[0].([]block.FetchBlockResult)
	return ret0
}

// FetchBlocks indicates an expected call of FetchBlocks
func (mr *MockdatabaseBufferMockRecorder) FetchBlocks(ctx, starts, nsCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlocks", reflect.TypeOf((*MockdatabaseBuffer)(nil).FetchBlocks), ctx, starts, nsCtx)
}

// FetchBlocksMetadata mocks base method
func (m *MockdatabaseBuffer) FetchBlocksMetadata(ctx context.Context, start, end time.Time, opts FetchBlocksMetadataOptions) (block.FetchBlockMetadataResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBlocksMetadata", ctx, start, end, opts)
	ret0, _ := ret[0].(block.FetchBlockMetadataResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBlocksMetadata indicates an expected call of FetchBlocksMetadata
func (mr *MockdatabaseBufferMockRecorder) FetchBlocksMetadata(ctx, start, end, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlocksMetadata", reflect.TypeOf((*MockdatabaseBuffer)(nil).FetchBlocksMetadata), ctx, start, end, opts)
}

// IsEmpty mocks base method
func (m *MockdatabaseBuffer) IsEmpty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmpty indicates an expected call of IsEmpty
func (mr *MockdatabaseBufferMockRecorder) IsEmpty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmpty", reflect.TypeOf((*MockdatabaseBuffer)(nil).IsEmpty))
}

// Stats mocks base method
func (m *MockdatabaseBuffer) Stats() bufferStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(bufferStats)
	return ret0
}

// Stats indicates an expected call of Stats
func (mr *MockdatabaseBufferMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockdatabaseBuffer)(nil).Stats))
}

// Tick mocks base method
func (m *MockdatabaseBuffer) Tick(versions map[time0.UnixNano]BlockState, nsCtx namespace.Context) bufferTickResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tick", versions, nsCtx)
	ret0, _ := ret[0].(bufferTickResult)
	return ret0
}

// Tick indicates an expected call of Tick
func (mr *MockdatabaseBufferMockRecorder) Tick(versions, nsCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tick", reflect.TypeOf((*MockdatabaseBuffer)(nil).Tick), versions, nsCtx)
}

// Bootstrap mocks base method
func (m *MockdatabaseBuffer) Bootstrap(bl block.DatabaseBlock) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Bootstrap", bl)
}

// Bootstrap indicates an expected call of Bootstrap
func (mr *MockdatabaseBufferMockRecorder) Bootstrap(bl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockdatabaseBuffer)(nil).Bootstrap), bl)
}

// Reset mocks base method
func (m *MockdatabaseBuffer) Reset(opts Options) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", opts)
}

// Reset indicates an expected call of Reset
func (mr *MockdatabaseBufferMockRecorder) Reset(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockdatabaseBuffer)(nil).Reset), opts)
}
