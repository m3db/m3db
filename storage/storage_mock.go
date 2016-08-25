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
// Source: github.com/m3db/m3db/storage/types.go

package storage

import (
	gomock "github.com/golang/mock/gomock"
	clock "github.com/m3db/m3db/clock"
	context "github.com/m3db/m3db/context"
	encoding "github.com/m3db/m3db/encoding"
	instrument "github.com/m3db/m3db/instrument"
	persist "github.com/m3db/m3db/persist"
	commitlog "github.com/m3db/m3db/persist/fs/commitlog"
	pool "github.com/m3db/m3db/pool"
	retention "github.com/m3db/m3db/retention"
	block "github.com/m3db/m3db/storage/block"
	bootstrap "github.com/m3db/m3db/storage/bootstrap"
	io "github.com/m3db/m3db/x/io"
	time "github.com/m3db/m3x/time"
	time0 "time"
)

// Mock of FetchBlockResult interface
type MockFetchBlockResult struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlockResultRecorder
}

// Recorder for MockFetchBlockResult (not exported)
type _MockFetchBlockResultRecorder struct {
	mock *MockFetchBlockResult
}

func NewMockFetchBlockResult(ctrl *gomock.Controller) *MockFetchBlockResult {
	mock := &MockFetchBlockResult{ctrl: ctrl}
	mock.recorder = &_MockFetchBlockResultRecorder{mock}
	return mock
}

func (_m *MockFetchBlockResult) EXPECT() *_MockFetchBlockResultRecorder {
	return _m.recorder
}

func (_m *MockFetchBlockResult) Start() time0.Time {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockFetchBlockResultRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockFetchBlockResult) Readers() []io.SegmentReader {
	ret := _m.ctrl.Call(_m, "Readers")
	ret0, _ := ret[0].([]io.SegmentReader)
	return ret0
}

func (_mr *_MockFetchBlockResultRecorder) Readers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Readers")
}

func (_m *MockFetchBlockResult) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFetchBlockResultRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

// Mock of FetchBlocksMetadataResult interface
type MockFetchBlocksMetadataResult struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlocksMetadataResultRecorder
}

// Recorder for MockFetchBlocksMetadataResult (not exported)
type _MockFetchBlocksMetadataResultRecorder struct {
	mock *MockFetchBlocksMetadataResult
}

func NewMockFetchBlocksMetadataResult(ctrl *gomock.Controller) *MockFetchBlocksMetadataResult {
	mock := &MockFetchBlocksMetadataResult{ctrl: ctrl}
	mock.recorder = &_MockFetchBlocksMetadataResultRecorder{mock}
	return mock
}

func (_m *MockFetchBlocksMetadataResult) EXPECT() *_MockFetchBlocksMetadataResultRecorder {
	return _m.recorder
}

func (_m *MockFetchBlocksMetadataResult) ID() string {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockFetchBlocksMetadataResultRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockFetchBlocksMetadataResult) Blocks() []FetchBlockMetadataResult {
	ret := _m.ctrl.Call(_m, "Blocks")
	ret0, _ := ret[0].([]FetchBlockMetadataResult)
	return ret0
}

func (_mr *_MockFetchBlocksMetadataResultRecorder) Blocks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Blocks")
}

// Mock of FetchBlockMetadataResult interface
type MockFetchBlockMetadataResult struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlockMetadataResultRecorder
}

// Recorder for MockFetchBlockMetadataResult (not exported)
type _MockFetchBlockMetadataResultRecorder struct {
	mock *MockFetchBlockMetadataResult
}

func NewMockFetchBlockMetadataResult(ctrl *gomock.Controller) *MockFetchBlockMetadataResult {
	mock := &MockFetchBlockMetadataResult{ctrl: ctrl}
	mock.recorder = &_MockFetchBlockMetadataResultRecorder{mock}
	return mock
}

func (_m *MockFetchBlockMetadataResult) EXPECT() *_MockFetchBlockMetadataResultRecorder {
	return _m.recorder
}

func (_m *MockFetchBlockMetadataResult) Start() time0.Time {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockFetchBlockMetadataResultRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockFetchBlockMetadataResult) Size() *int64 {
	ret := _m.ctrl.Call(_m, "Size")
	ret0, _ := ret[0].(*int64)
	return ret0
}

func (_mr *_MockFetchBlockMetadataResultRecorder) Size() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Size")
}

func (_m *MockFetchBlockMetadataResult) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFetchBlockMetadataResultRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

// Mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseRecorder
}

// Recorder for MockDatabase (not exported)
type _MockDatabaseRecorder struct {
	mock *MockDatabase
}

func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &_MockDatabaseRecorder{mock}
	return mock
}

func (_m *MockDatabase) EXPECT() *_MockDatabaseRecorder {
	return _m.recorder
}

func (_m *MockDatabase) Options() Options {
	ret := _m.ctrl.Call(_m, "Options")
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Options() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Options")
}

func (_m *MockDatabase) Open() error {
	ret := _m.ctrl.Call(_m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Open() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open")
}

func (_m *MockDatabase) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockDatabase) Write(ctx context.Context, id string, timestamp time0.Time, value float64, unit time.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, id, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Write(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockDatabase) ReadEncoded(ctx context.Context, id string, start time0.Time, end time0.Time) ([][]io.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, id, start, end)
	ret0, _ := ret[0].([][]io.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) ReadEncoded(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2, arg3)
}

func (_m *MockDatabase) FetchBlocks(ctx context.Context, shard uint32, id string, starts []time0.Time) ([]FetchBlockResult, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, shard, id, starts)
	ret0, _ := ret[0].([]FetchBlockResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) FetchBlocks(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1, arg2, arg3)
}

func (_m *MockDatabase) FetchBlocksMetadata(ctx context.Context, shardID uint32, limit int64, pageToken int64, includeSizes bool) ([]FetchBlocksMetadataResult, *int64, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, shardID, limit, pageToken, includeSizes)
	ret0, _ := ret[0].([]FetchBlocksMetadataResult)
	ret1, _ := ret[1].(*int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockDatabaseRecorder) FetchBlocksMetadata(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockDatabase) Bootstrap() error {
	ret := _m.ctrl.Call(_m, "Bootstrap")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Bootstrap() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap")
}

func (_m *MockDatabase) IsBootstrapped() bool {
	ret := _m.ctrl.Call(_m, "IsBootstrapped")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDatabaseRecorder) IsBootstrapped() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsBootstrapped")
}

// Mock of databaseShard interface
type MockdatabaseShard struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseShardRecorder
}

// Recorder for MockdatabaseShard (not exported)
type _MockdatabaseShardRecorder struct {
	mock *MockdatabaseShard
}

func NewMockdatabaseShard(ctrl *gomock.Controller) *MockdatabaseShard {
	mock := &MockdatabaseShard{ctrl: ctrl}
	mock.recorder = &_MockdatabaseShardRecorder{mock}
	return mock
}

func (_m *MockdatabaseShard) EXPECT() *_MockdatabaseShardRecorder {
	return _m.recorder
}

func (_m *MockdatabaseShard) ID() uint32 {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(uint32)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockdatabaseShard) Tick() {
	_m.ctrl.Call(_m, "Tick")
}

func (_mr *_MockdatabaseShardRecorder) Tick() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tick")
}

func (_m *MockdatabaseShard) Write(ctx context.Context, id string, timestamp time0.Time, value float64, unit time.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, id, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) Write(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockdatabaseShard) ReadEncoded(ctx context.Context, id string, start time0.Time, end time0.Time) ([][]io.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, id, start, end)
	ret0, _ := ret[0].([][]io.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdatabaseShardRecorder) ReadEncoded(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2, arg3)
}

func (_m *MockdatabaseShard) FetchBlocks(ctx context.Context, id string, starts []time0.Time) []FetchBlockResult {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, id, starts)
	ret0, _ := ret[0].([]FetchBlockResult)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) FetchBlocks(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1, arg2)
}

func (_m *MockdatabaseShard) FetchBlocksMetadata(ctx context.Context, limit int64, pageToken int64, includeSizes bool) ([]FetchBlocksMetadataResult, *int64) {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, limit, pageToken, includeSizes)
	ret0, _ := ret[0].([]FetchBlocksMetadataResult)
	ret1, _ := ret[1].(*int64)
	return ret0, ret1
}

func (_mr *_MockdatabaseShardRecorder) FetchBlocksMetadata(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1, arg2, arg3)
}

func (_m *MockdatabaseShard) Bootstrap(bs bootstrap.Bootstrap, writeStart time0.Time, cutover time0.Time) error {
	ret := _m.ctrl.Call(_m, "Bootstrap", bs, writeStart, cutover)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) Bootstrap(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap", arg0, arg1, arg2)
}

func (_m *MockdatabaseShard) Flush(ctx context.Context, blockStart time0.Time, pm persist.Manager) error {
	ret := _m.ctrl.Call(_m, "Flush", ctx, blockStart, pm)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) Flush(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0, arg1, arg2)
}

// Mock of databaseSeries interface
type MockdatabaseSeries struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseSeriesRecorder
}

// Recorder for MockdatabaseSeries (not exported)
type _MockdatabaseSeriesRecorder struct {
	mock *MockdatabaseSeries
}

func NewMockdatabaseSeries(ctrl *gomock.Controller) *MockdatabaseSeries {
	mock := &MockdatabaseSeries{ctrl: ctrl}
	mock.recorder = &_MockdatabaseSeriesRecorder{mock}
	return mock
}

func (_m *MockdatabaseSeries) EXPECT() *_MockdatabaseSeriesRecorder {
	return _m.recorder
}

func (_m *MockdatabaseSeries) ID() string {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockdatabaseSeriesRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockdatabaseSeries) Tick() error {
	ret := _m.ctrl.Call(_m, "Tick")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseSeriesRecorder) Tick() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tick")
}

func (_m *MockdatabaseSeries) Write(ctx context.Context, timestamp time0.Time, value float64, unit time.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseSeriesRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockdatabaseSeries) ReadEncoded(ctx context.Context, start time0.Time, end time0.Time) ([][]io.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, start, end)
	ret0, _ := ret[0].([][]io.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdatabaseSeriesRecorder) ReadEncoded(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2)
}

func (_m *MockdatabaseSeries) FetchBlocks(ctx context.Context, starts []time0.Time) []FetchBlockResult {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, starts)
	ret0, _ := ret[0].([]FetchBlockResult)
	return ret0
}

func (_mr *_MockdatabaseSeriesRecorder) FetchBlocks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1)
}

func (_m *MockdatabaseSeries) FetchBlocksMetadata(ctx context.Context, includeSizes bool) FetchBlocksMetadataResult {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, includeSizes)
	ret0, _ := ret[0].(FetchBlocksMetadataResult)
	return ret0
}

func (_mr *_MockdatabaseSeriesRecorder) FetchBlocksMetadata(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1)
}

func (_m *MockdatabaseSeries) Empty() bool {
	ret := _m.ctrl.Call(_m, "Empty")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseSeriesRecorder) Empty() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Empty")
}

func (_m *MockdatabaseSeries) Bootstrap(rs block.DatabaseSeriesBlocks, cutover time0.Time) error {
	ret := _m.ctrl.Call(_m, "Bootstrap", rs, cutover)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseSeriesRecorder) Bootstrap(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap", arg0, arg1)
}

func (_m *MockdatabaseSeries) Flush(ctx context.Context, blockStart time0.Time, persistFn persist.Fn) error {
	ret := _m.ctrl.Call(_m, "Flush", ctx, blockStart, persistFn)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseSeriesRecorder) Flush(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0, arg1, arg2)
}

// Mock of databaseBuffer interface
type MockdatabaseBuffer struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseBufferRecorder
}

// Recorder for MockdatabaseBuffer (not exported)
type _MockdatabaseBufferRecorder struct {
	mock *MockdatabaseBuffer
}

func NewMockdatabaseBuffer(ctrl *gomock.Controller) *MockdatabaseBuffer {
	mock := &MockdatabaseBuffer{ctrl: ctrl}
	mock.recorder = &_MockdatabaseBufferRecorder{mock}
	return mock
}

func (_m *MockdatabaseBuffer) EXPECT() *_MockdatabaseBufferRecorder {
	return _m.recorder
}

func (_m *MockdatabaseBuffer) Write(ctx context.Context, timestamp time0.Time, value float64, unit time.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockdatabaseBuffer) ReadEncoded(ctx context.Context, start time0.Time, end time0.Time) [][]io.SegmentReader {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, start, end)
	ret0, _ := ret[0].([][]io.SegmentReader)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) ReadEncoded(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2)
}

func (_m *MockdatabaseBuffer) FetchBlocks(ctx context.Context, starts []time0.Time) []FetchBlockResult {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, starts)
	ret0, _ := ret[0].([]FetchBlockResult)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) FetchBlocks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1)
}

func (_m *MockdatabaseBuffer) FetchBlocksMetadata(ctx context.Context, includeSizes bool) []FetchBlockMetadataResult {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, includeSizes)
	ret0, _ := ret[0].([]FetchBlockMetadataResult)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) FetchBlocksMetadata(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1)
}

func (_m *MockdatabaseBuffer) Empty() bool {
	ret := _m.ctrl.Call(_m, "Empty")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) Empty() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Empty")
}

func (_m *MockdatabaseBuffer) NeedsDrain() bool {
	ret := _m.ctrl.Call(_m, "NeedsDrain")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) NeedsDrain() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NeedsDrain")
}

func (_m *MockdatabaseBuffer) DrainAndReset(forced bool) {
	_m.ctrl.Call(_m, "DrainAndReset", forced)
}

func (_mr *_MockdatabaseBufferRecorder) DrainAndReset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DrainAndReset", arg0)
}

// Mock of databaseBootstrapManager interface
type MockdatabaseBootstrapManager struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseBootstrapManagerRecorder
}

// Recorder for MockdatabaseBootstrapManager (not exported)
type _MockdatabaseBootstrapManagerRecorder struct {
	mock *MockdatabaseBootstrapManager
}

func NewMockdatabaseBootstrapManager(ctrl *gomock.Controller) *MockdatabaseBootstrapManager {
	mock := &MockdatabaseBootstrapManager{ctrl: ctrl}
	mock.recorder = &_MockdatabaseBootstrapManagerRecorder{mock}
	return mock
}

func (_m *MockdatabaseBootstrapManager) EXPECT() *_MockdatabaseBootstrapManagerRecorder {
	return _m.recorder
}

func (_m *MockdatabaseBootstrapManager) IsBootstrapped() bool {
	ret := _m.ctrl.Call(_m, "IsBootstrapped")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseBootstrapManagerRecorder) IsBootstrapped() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsBootstrapped")
}

func (_m *MockdatabaseBootstrapManager) Bootstrap() error {
	ret := _m.ctrl.Call(_m, "Bootstrap")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseBootstrapManagerRecorder) Bootstrap() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap")
}

// Mock of databaseFlushManager interface
type MockdatabaseFlushManager struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseFlushManagerRecorder
}

// Recorder for MockdatabaseFlushManager (not exported)
type _MockdatabaseFlushManagerRecorder struct {
	mock *MockdatabaseFlushManager
}

func NewMockdatabaseFlushManager(ctrl *gomock.Controller) *MockdatabaseFlushManager {
	mock := &MockdatabaseFlushManager{ctrl: ctrl}
	mock.recorder = &_MockdatabaseFlushManagerRecorder{mock}
	return mock
}

func (_m *MockdatabaseFlushManager) EXPECT() *_MockdatabaseFlushManagerRecorder {
	return _m.recorder
}

<<<<<<< 189c89dfd70b8996bcfd511498f84566c69edb78
func (_m *MockdatabaseFlushManager) HasFlushed(t time0.Time) bool {
	ret := _m.ctrl.Call(_m, "HasFlushed", t)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseFlushManagerRecorder) HasFlushed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasFlushed", arg0)
}

func (_m *MockdatabaseFlushManager) FlushTimeStart(t time0.Time) time0.Time {
	ret := _m.ctrl.Call(_m, "FlushTimeStart", t)
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockdatabaseFlushManagerRecorder) FlushTimeStart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushTimeStart", arg0)
}

func (_m *MockdatabaseFlushManager) FlushTimeEnd(t time0.Time) time0.Time {
	ret := _m.ctrl.Call(_m, "FlushTimeEnd", t)
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockdatabaseFlushManagerRecorder) FlushTimeEnd(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushTimeEnd", arg0)
}

func (_m *MockdatabaseFlushManager) Flush(t time0.Time) error {
	ret := _m.ctrl.Call(_m, "Flush", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseFlushManagerRecorder) Flush(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0)
}

// Mock of databaseCleanupManager interface
type MockdatabaseCleanupManager struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseCleanupManagerRecorder
}

// Recorder for MockdatabaseCleanupManager (not exported)
type _MockdatabaseCleanupManagerRecorder struct {
	mock *MockdatabaseCleanupManager
}

func NewMockdatabaseCleanupManager(ctrl *gomock.Controller) *MockdatabaseCleanupManager {
	mock := &MockdatabaseCleanupManager{ctrl: ctrl}
	mock.recorder = &_MockdatabaseCleanupManagerRecorder{mock}
	return mock
}

func (_m *MockdatabaseCleanupManager) EXPECT() *_MockdatabaseCleanupManagerRecorder {
	return _m.recorder
}

func (_m *MockdatabaseCleanupManager) Cleanup(t time0.Time) error {
	ret := _m.ctrl.Call(_m, "Cleanup", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseCleanupManagerRecorder) Cleanup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Cleanup", arg0)
}

// Mock of databaseFileSystemManager interface
type MockdatabaseFileSystemManager struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseFileSystemManagerRecorder
}

// Recorder for MockdatabaseFileSystemManager (not exported)
type _MockdatabaseFileSystemManagerRecorder struct {
	mock *MockdatabaseFileSystemManager
}

func NewMockdatabaseFileSystemManager(ctrl *gomock.Controller) *MockdatabaseFileSystemManager {
	mock := &MockdatabaseFileSystemManager{ctrl: ctrl}
	mock.recorder = &_MockdatabaseFileSystemManagerRecorder{mock}
	return mock
}

func (_m *MockdatabaseFileSystemManager) EXPECT() *_MockdatabaseFileSystemManagerRecorder {
	return _m.recorder
}

func (_m *MockdatabaseFileSystemManager) HasFlushed(t time0.Time) bool {
	ret := _m.ctrl.Call(_m, "HasFlushed", t)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) HasFlushed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasFlushed", arg0)
}

func (_m *MockdatabaseFileSystemManager) FlushTimeStart(t time0.Time) time0.Time {
	ret := _m.ctrl.Call(_m, "FlushTimeStart", t)
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) FlushTimeStart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushTimeStart", arg0)
}

func (_m *MockdatabaseFileSystemManager) FlushTimeEnd(t time0.Time) time0.Time {
	ret := _m.ctrl.Call(_m, "FlushTimeEnd", t)
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) FlushTimeEnd(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushTimeEnd", arg0)
}

func (_m *MockdatabaseFileSystemManager) Flush(t time0.Time) error {
	ret := _m.ctrl.Call(_m, "Flush", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) Flush(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0)
}

func (_m *MockdatabaseFileSystemManager) Cleanup(t time0.Time) error {
	ret := _m.ctrl.Call(_m, "Cleanup", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) Cleanup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Cleanup", arg0)
}

func (_m *MockdatabaseFileSystemManager) ShouldRun(t time0.Time) bool {
	ret := _m.ctrl.Call(_m, "ShouldRun", t)
=======
func (_m *MockdatabaseFlushManager) NeedsFlush(t time0.Time) bool {
	ret := _m.ctrl.Call(_m, "NeedsFlush", t)
>>>>>>> Strip GOPATH from source path in auto-generated mocks
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) ShouldRun(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShouldRun", arg0)
}

<<<<<<< 189c89dfd70b8996bcfd511498f84566c69edb78
func (_m *MockdatabaseFileSystemManager) Run(t time0.Time, async bool) {
	_m.ctrl.Call(_m, "Run", t, async)
=======
func (_m *MockdatabaseFlushManager) Flush(t time0.Time, async bool) {
	_m.ctrl.Call(_m, "Flush", t, async)
>>>>>>> Strip GOPATH from source path in auto-generated mocks
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0, arg1)
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

func (_m *MockOptions) ClockOptions(value clock.Options) Options {
	ret := _m.ctrl.Call(_m, "ClockOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ClockOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClockOptions", arg0)
}

func (_m *MockOptions) GetClockOptions() clock.Options {
	ret := _m.ctrl.Call(_m, "GetClockOptions")
	ret0, _ := ret[0].(clock.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetClockOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetClockOptions")
}

func (_m *MockOptions) InstrumentOptions(value instrument.Options) Options {
	ret := _m.ctrl.Call(_m, "InstrumentOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) InstrumentOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstrumentOptions", arg0)
}

func (_m *MockOptions) GetInstrumentOptions() instrument.Options {
	ret := _m.ctrl.Call(_m, "GetInstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetInstrumentOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetInstrumentOptions")
}

func (_m *MockOptions) RetentionOptions(value retention.Options) Options {
	ret := _m.ctrl.Call(_m, "RetentionOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) RetentionOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetentionOptions", arg0)
}

func (_m *MockOptions) GetRetentionOptions() retention.Options {
	ret := _m.ctrl.Call(_m, "GetRetentionOptions")
	ret0, _ := ret[0].(retention.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetRetentionOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRetentionOptions")
}

func (_m *MockOptions) DatabaseBlockOptions(value block.Options) Options {
	ret := _m.ctrl.Call(_m, "DatabaseBlockOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseBlockOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseBlockOptions", arg0)
}

func (_m *MockOptions) GetDatabaseBlockOptions() block.Options {
	ret := _m.ctrl.Call(_m, "GetDatabaseBlockOptions")
	ret0, _ := ret[0].(block.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetDatabaseBlockOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDatabaseBlockOptions")
}

func (_m *MockOptions) CommitLogOptions(value commitlog.Options) Options {
	ret := _m.ctrl.Call(_m, "CommitLogOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) CommitLogOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CommitLogOptions", arg0)
}

func (_m *MockOptions) GetCommitLogOptions() commitlog.Options {
	ret := _m.ctrl.Call(_m, "GetCommitLogOptions")
	ret0, _ := ret[0].(commitlog.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetCommitLogOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCommitLogOptions")
}

func (_m *MockOptions) EncodingM3TSZPooled() Options {
	ret := _m.ctrl.Call(_m, "EncodingM3TSZPooled")
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) EncodingM3TSZPooled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EncodingM3TSZPooled")
}

func (_m *MockOptions) EncodingM3TSZ() Options {
	ret := _m.ctrl.Call(_m, "EncodingM3TSZ")
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) EncodingM3TSZ() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EncodingM3TSZ")
}

func (_m *MockOptions) NewEncoderFn(value encoding.NewEncoderFn) Options {
	ret := _m.ctrl.Call(_m, "NewEncoderFn", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) NewEncoderFn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewEncoderFn", arg0)
}

func (_m *MockOptions) GetNewEncoderFn() encoding.NewEncoderFn {
	ret := _m.ctrl.Call(_m, "GetNewEncoderFn")
	ret0, _ := ret[0].(encoding.NewEncoderFn)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetNewEncoderFn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNewEncoderFn")
}

func (_m *MockOptions) NewDecoderFn(value encoding.NewDecoderFn) Options {
	ret := _m.ctrl.Call(_m, "NewDecoderFn", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) NewDecoderFn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewDecoderFn", arg0)
}

func (_m *MockOptions) GetNewDecoderFn() encoding.NewDecoderFn {
	ret := _m.ctrl.Call(_m, "GetNewDecoderFn")
	ret0, _ := ret[0].(encoding.NewDecoderFn)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetNewDecoderFn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNewDecoderFn")
}

func (_m *MockOptions) NewBootstrapFn(value NewBootstrapFn) Options {
	ret := _m.ctrl.Call(_m, "NewBootstrapFn", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) NewBootstrapFn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewBootstrapFn", arg0)
}

func (_m *MockOptions) GetNewBootstrapFn() NewBootstrapFn {
	ret := _m.ctrl.Call(_m, "GetNewBootstrapFn")
	ret0, _ := ret[0].(NewBootstrapFn)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetNewBootstrapFn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNewBootstrapFn")
}

func (_m *MockOptions) NewPersistManagerFn(value NewPersistManagerFn) Options {
	ret := _m.ctrl.Call(_m, "NewPersistManagerFn", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) NewPersistManagerFn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewPersistManagerFn", arg0)
}

func (_m *MockOptions) GetNewPersistManagerFn() NewPersistManagerFn {
	ret := _m.ctrl.Call(_m, "GetNewPersistManagerFn")
	ret0, _ := ret[0].(NewPersistManagerFn)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetNewPersistManagerFn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNewPersistManagerFn")
}

func (_m *MockOptions) MaxFlushRetries(value int) Options {
	ret := _m.ctrl.Call(_m, "MaxFlushRetries", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) MaxFlushRetries(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MaxFlushRetries", arg0)
}

func (_m *MockOptions) GetMaxFlushRetries() int {
	ret := _m.ctrl.Call(_m, "GetMaxFlushRetries")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetMaxFlushRetries() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMaxFlushRetries")
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
