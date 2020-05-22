// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/storage/series/buffer.go

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

// Package series is a generated GoMock package.
package series

import (
	"reflect"
	"time"

	"github.com/m3db/m3/src/dbnode/namespace"
	"github.com/m3db/m3/src/dbnode/persist"
	"github.com/m3db/m3/src/dbnode/storage/block"
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
func (m *MockdatabaseBuffer) Write(ctx context.Context, timestamp time.Time, value float64, unit time0.Unit, annotation []byte, wOpts WriteOptions) (bool, WriteType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", ctx, timestamp, value, unit, annotation, wOpts)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(WriteType)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Write indicates an expected call of Write
func (mr *MockdatabaseBufferMockRecorder) Write(ctx, timestamp, value, unit, annotation, wOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockdatabaseBuffer)(nil).Write), ctx, timestamp, value, unit, annotation, wOpts)
}

// Snapshot mocks base method
func (m *MockdatabaseBuffer) Snapshot(ctx context.Context, blockStart time.Time, id ident.ID, tags ident.Tags, persistFn persist.DataFn, nsCtx namespace.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", ctx, blockStart, id, tags, persistFn, nsCtx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Snapshot indicates an expected call of Snapshot
func (mr *MockdatabaseBufferMockRecorder) Snapshot(ctx, blockStart, id, tags, persistFn, nsCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockdatabaseBuffer)(nil).Snapshot), ctx, blockStart, id, tags, persistFn, nsCtx)
}

// WarmFlush mocks base method
func (m *MockdatabaseBuffer) WarmFlush(ctx context.Context, blockStart time.Time, id ident.ID, tags ident.Tags, persistFn persist.DataFn, nsCtx namespace.Context) (FlushOutcome, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WarmFlush", ctx, blockStart, id, tags, persistFn, nsCtx)
	ret0, _ := ret[0].(FlushOutcome)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WarmFlush indicates an expected call of WarmFlush
func (mr *MockdatabaseBufferMockRecorder) WarmFlush(ctx, blockStart, id, tags, persistFn, nsCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarmFlush", reflect.TypeOf((*MockdatabaseBuffer)(nil).WarmFlush), ctx, blockStart, id, tags, persistFn, nsCtx)
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

// FetchBlocksForColdFlush mocks base method
func (m *MockdatabaseBuffer) FetchBlocksForColdFlush(ctx context.Context, start time.Time, version int, nsCtx namespace.Context) (block.FetchBlockResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBlocksForColdFlush", ctx, start, version, nsCtx)
	ret0, _ := ret[0].(block.FetchBlockResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBlocksForColdFlush indicates an expected call of FetchBlocksForColdFlush
func (mr *MockdatabaseBufferMockRecorder) FetchBlocksForColdFlush(ctx, start, version, nsCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlocksForColdFlush", reflect.TypeOf((*MockdatabaseBuffer)(nil).FetchBlocksForColdFlush), ctx, start, version, nsCtx)
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

// ColdFlushBlockStarts mocks base method
func (m *MockdatabaseBuffer) ColdFlushBlockStarts(blockStates map[time0.UnixNano]BlockState) OptimizedTimes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ColdFlushBlockStarts", blockStates)
	ret0, _ := ret[0].(OptimizedTimes)
	return ret0
}

// ColdFlushBlockStarts indicates an expected call of ColdFlushBlockStarts
func (mr *MockdatabaseBufferMockRecorder) ColdFlushBlockStarts(blockStates interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ColdFlushBlockStarts", reflect.TypeOf((*MockdatabaseBuffer)(nil).ColdFlushBlockStarts), blockStates)
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
func (m *MockdatabaseBuffer) Tick(versions ShardBlockStateSnapshot, nsCtx namespace.Context) bufferTickResult {
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

// Load mocks base method
func (m *MockdatabaseBuffer) Load(bl block.DatabaseBlock, writeType WriteType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Load", bl, writeType)
}

// Load indicates an expected call of Load
func (mr *MockdatabaseBufferMockRecorder) Load(bl, writeType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockdatabaseBuffer)(nil).Load), bl, writeType)
}

// Reset mocks base method
func (m *MockdatabaseBuffer) Reset(opts databaseBufferResetOptions) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", opts)
}

// Reset indicates an expected call of Reset
func (mr *MockdatabaseBufferMockRecorder) Reset(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockdatabaseBuffer)(nil).Reset), opts)
}
