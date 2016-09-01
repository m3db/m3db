// Copyright (c) 2016 Uber Technologies, Inc.
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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3db/storage/block/types.go

package block

import (
	gomock "github.com/golang/mock/gomock"
	context "github.com/m3db/m3db/context"
	encoding "github.com/m3db/m3db/encoding"
	pool "github.com/m3db/m3db/pool"
	ts "github.com/m3db/m3db/ts"
	io "github.com/m3db/m3db/x/io"
	time0 "github.com/m3db/m3x/time"
	time "time"
)

// Mock of DatabaseBlock interface
type MockDatabaseBlock struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseBlockRecorder
}

// Recorder for MockDatabaseBlock (not exported)
type _MockDatabaseBlockRecorder struct {
	mock *MockDatabaseBlock
}

func NewMockDatabaseBlock(ctrl *gomock.Controller) *MockDatabaseBlock {
	mock := &MockDatabaseBlock{ctrl: ctrl}
	mock.recorder = &_MockDatabaseBlockRecorder{mock}
	return mock
}

func (_m *MockDatabaseBlock) EXPECT() *_MockDatabaseBlockRecorder {
	return _m.recorder
}

func (_m *MockDatabaseBlock) StartTime() time.Time {
	ret := _m.ctrl.Call(_m, "StartTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockDatabaseBlockRecorder) StartTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartTime")
}

func (_m *MockDatabaseBlock) IsSealed() bool {
	ret := _m.ctrl.Call(_m, "IsSealed")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDatabaseBlockRecorder) IsSealed() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsSealed")
}

func (_m *MockDatabaseBlock) Write(timestamp time.Time, value float64, unit time0.Unit, annotation ts.Annotation) error {
	ret := _m.ctrl.Call(_m, "Write", timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseBlockRecorder) Write(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3)
}

func (_m *MockDatabaseBlock) Stream(blocker context.Context) (io.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "Stream", blocker)
	ret0, _ := ret[0].(io.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseBlockRecorder) Stream(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stream", arg0)
}

func (_m *MockDatabaseBlock) Reset(startTime time.Time, encoder encoding.Encoder) {
	_m.ctrl.Call(_m, "Reset", startTime, encoder)
}

func (_mr *_MockDatabaseBlockRecorder) Reset(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0, arg1)
}

func (_m *MockDatabaseBlock) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockDatabaseBlockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockDatabaseBlock) Seal() {
	_m.ctrl.Call(_m, "Seal")
}

func (_mr *_MockDatabaseBlockRecorder) Seal() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Seal")
}

// Mock of DatabaseSeriesBlocks interface
type MockDatabaseSeriesBlocks struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseSeriesBlocksRecorder
}

// Recorder for MockDatabaseSeriesBlocks (not exported)
type _MockDatabaseSeriesBlocksRecorder struct {
	mock *MockDatabaseSeriesBlocks
}

func NewMockDatabaseSeriesBlocks(ctrl *gomock.Controller) *MockDatabaseSeriesBlocks {
	mock := &MockDatabaseSeriesBlocks{ctrl: ctrl}
	mock.recorder = &_MockDatabaseSeriesBlocksRecorder{mock}
	return mock
}

func (_m *MockDatabaseSeriesBlocks) EXPECT() *_MockDatabaseSeriesBlocksRecorder {
	return _m.recorder
}

func (_m *MockDatabaseSeriesBlocks) Len() int {
	ret := _m.ctrl.Call(_m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) Len() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Len")
}

func (_m *MockDatabaseSeriesBlocks) AddBlock(block DatabaseBlock) {
	_m.ctrl.Call(_m, "AddBlock", block)
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) AddBlock(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddBlock", arg0)
}

func (_m *MockDatabaseSeriesBlocks) AddSeries(other DatabaseSeriesBlocks) {
	_m.ctrl.Call(_m, "AddSeries", other)
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) AddSeries(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddSeries", arg0)
}

func (_m *MockDatabaseSeriesBlocks) GetMinTime() time.Time {
	ret := _m.ctrl.Call(_m, "GetMinTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) GetMinTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMinTime")
}

func (_m *MockDatabaseSeriesBlocks) GetMaxTime() time.Time {
	ret := _m.ctrl.Call(_m, "GetMaxTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) GetMaxTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMaxTime")
}

func (_m *MockDatabaseSeriesBlocks) GetBlockAt(t time.Time) (DatabaseBlock, bool) {
	ret := _m.ctrl.Call(_m, "GetBlockAt", t)
	ret0, _ := ret[0].(DatabaseBlock)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) GetBlockAt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBlockAt", arg0)
}

func (_m *MockDatabaseSeriesBlocks) GetBlockOrAdd(t time.Time) DatabaseBlock {
	ret := _m.ctrl.Call(_m, "GetBlockOrAdd", t)
	ret0, _ := ret[0].(DatabaseBlock)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) GetBlockOrAdd(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBlockOrAdd", arg0)
}

func (_m *MockDatabaseSeriesBlocks) GetAllBlocks() map[time.Time]DatabaseBlock {
	ret := _m.ctrl.Call(_m, "GetAllBlocks")
	ret0, _ := ret[0].(map[time.Time]DatabaseBlock)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) GetAllBlocks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllBlocks")
}

func (_m *MockDatabaseSeriesBlocks) RemoveBlockAt(t time.Time) {
	_m.ctrl.Call(_m, "RemoveBlockAt", t)
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) RemoveBlockAt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveBlockAt", arg0)
}

func (_m *MockDatabaseSeriesBlocks) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of DatabaseBlockPool interface
type MockDatabaseBlockPool struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseBlockPoolRecorder
}

// Recorder for MockDatabaseBlockPool (not exported)
type _MockDatabaseBlockPoolRecorder struct {
	mock *MockDatabaseBlockPool
}

func NewMockDatabaseBlockPool(ctrl *gomock.Controller) *MockDatabaseBlockPool {
	mock := &MockDatabaseBlockPool{ctrl: ctrl}
	mock.recorder = &_MockDatabaseBlockPoolRecorder{mock}
	return mock
}

func (_m *MockDatabaseBlockPool) EXPECT() *_MockDatabaseBlockPoolRecorder {
	return _m.recorder
}

func (_m *MockDatabaseBlockPool) Init(alloc DatabaseBlockAllocate) {
	_m.ctrl.Call(_m, "Init", alloc)
}

func (_mr *_MockDatabaseBlockPoolRecorder) Init(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0)
}

func (_m *MockDatabaseBlockPool) Get() DatabaseBlock {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(DatabaseBlock)
	return ret0
}

func (_mr *_MockDatabaseBlockPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockDatabaseBlockPool) Put(block DatabaseBlock) {
	_m.ctrl.Call(_m, "Put", block)
}

func (_mr *_MockDatabaseBlockPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *_MockOptionsRecorder
}

// Recorder for MockOptions (not exported)
type _MockOptionsRecorder struct {
	mock *MockOptions
}

func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &_MockOptionsRecorder{mock}
	return mock
}

func (_m *MockOptions) EXPECT() *_MockOptionsRecorder {
	return _m.recorder
}

func (_m *MockOptions) DatabaseBlockAllocSize(value int) Options {
	ret := _m.ctrl.Call(_m, "DatabaseBlockAllocSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseBlockAllocSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseBlockAllocSize", arg0)
}

func (_m *MockOptions) GetDatabaseBlockAllocSize() int {
	ret := _m.ctrl.Call(_m, "GetDatabaseBlockAllocSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetDatabaseBlockAllocSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDatabaseBlockAllocSize")
}

func (_m *MockOptions) DatabaseBlockPool(value DatabaseBlockPool) Options {
	ret := _m.ctrl.Call(_m, "DatabaseBlockPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseBlockPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseBlockPool", arg0)
}

func (_m *MockOptions) GetDatabaseBlockPool() DatabaseBlockPool {
	ret := _m.ctrl.Call(_m, "GetDatabaseBlockPool")
	ret0, _ := ret[0].(DatabaseBlockPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetDatabaseBlockPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDatabaseBlockPool")
}

func (_m *MockOptions) ContextPool(value context.Pool) Options {
	ret := _m.ctrl.Call(_m, "ContextPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ContextPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContextPool", arg0)
}

func (_m *MockOptions) GetContextPool() context.Pool {
	ret := _m.ctrl.Call(_m, "GetContextPool")
	ret0, _ := ret[0].(context.Pool)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetContextPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetContextPool")
}

func (_m *MockOptions) EncoderPool(value encoding.EncoderPool) Options {
	ret := _m.ctrl.Call(_m, "EncoderPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) EncoderPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EncoderPool", arg0)
}

func (_m *MockOptions) GetEncoderPool() encoding.EncoderPool {
	ret := _m.ctrl.Call(_m, "GetEncoderPool")
	ret0, _ := ret[0].(encoding.EncoderPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetEncoderPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEncoderPool")
}

func (_m *MockOptions) ReaderIteratorPool(value encoding.ReaderIteratorPool) Options {
	ret := _m.ctrl.Call(_m, "ReaderIteratorPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ReaderIteratorPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReaderIteratorPool", arg0)
}

func (_m *MockOptions) GetReaderIteratorPool() encoding.ReaderIteratorPool {
	ret := _m.ctrl.Call(_m, "GetReaderIteratorPool")
	ret0, _ := ret[0].(encoding.ReaderIteratorPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetReaderIteratorPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetReaderIteratorPool")
}

func (_m *MockOptions) MultiReaderIteratorPool(value encoding.MultiReaderIteratorPool) Options {
	ret := _m.ctrl.Call(_m, "MultiReaderIteratorPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) MultiReaderIteratorPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MultiReaderIteratorPool", arg0)
}

func (_m *MockOptions) GetMultiReaderIteratorPool() encoding.MultiReaderIteratorPool {
	ret := _m.ctrl.Call(_m, "GetMultiReaderIteratorPool")
	ret0, _ := ret[0].(encoding.MultiReaderIteratorPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetMultiReaderIteratorPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMultiReaderIteratorPool")
}

func (_m *MockOptions) SegmentReaderPool(value io.SegmentReaderPool) Options {
	ret := _m.ctrl.Call(_m, "SegmentReaderPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SegmentReaderPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SegmentReaderPool", arg0)
}

func (_m *MockOptions) GetSegmentReaderPool() io.SegmentReaderPool {
	ret := _m.ctrl.Call(_m, "GetSegmentReaderPool")
	ret0, _ := ret[0].(io.SegmentReaderPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetSegmentReaderPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSegmentReaderPool")
}

func (_m *MockOptions) BytesPool(value pool.BytesPool) Options {
	ret := _m.ctrl.Call(_m, "BytesPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) BytesPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BytesPool", arg0)
}

func (_m *MockOptions) GetBytesPool() pool.BytesPool {
	ret := _m.ctrl.Call(_m, "GetBytesPool")
	ret0, _ := ret[0].(pool.BytesPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetBytesPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBytesPool")
}
