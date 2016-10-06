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
	repair "github.com/m3db/m3db/storage/repair"
	series "github.com/m3db/m3db/storage/series"
	io "github.com/m3db/m3db/x/io"
	time0 "github.com/m3db/m3x/time"
	time "time"
)

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

func (_m *MockDatabase) Write(ctx context.Context, namespace string, id string, timestamp time.Time, value float64, unit time0.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, namespace, id, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Write(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

func (_m *MockDatabase) ReadEncoded(ctx context.Context, namespace string, id string, start time.Time, end time.Time) ([][]io.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, namespace, id, start, end)
	ret0, _ := ret[0].([][]io.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) ReadEncoded(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockDatabase) FetchBlocks(ctx context.Context, namespace string, shard uint32, id string, starts []time.Time) ([]block.FetchBlockResult, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, namespace, shard, id, starts)
	ret0, _ := ret[0].([]block.FetchBlockResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) FetchBlocks(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockDatabase) FetchBlocksMetadata(ctx context.Context, namespace string, shard uint32, limit int64, pageToken int64, includeSizes bool, includeChecksums bool) ([]block.FetchBlocksMetadataResult, *int64, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, namespace, shard, limit, pageToken, includeSizes, includeChecksums)
	ret0, _ := ret[0].([]block.FetchBlocksMetadataResult)
	ret1, _ := ret[1].(*int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockDatabaseRecorder) FetchBlocksMetadata(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
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

func (_m *MockDatabase) Repair() error {
	ret := _m.ctrl.Call(_m, "Repair")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Repair() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Repair")
}

func (_m *MockDatabase) Truncate(namespace string) (int64, error) {
	ret := _m.ctrl.Call(_m, "Truncate", namespace)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) Truncate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Truncate", arg0)
}

// Mock of databaseNamespace interface
type MockdatabaseNamespace struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseNamespaceRecorder
}

// Recorder for MockdatabaseNamespace (not exported)
type _MockdatabaseNamespaceRecorder struct {
	mock *MockdatabaseNamespace
}

func NewMockdatabaseNamespace(ctrl *gomock.Controller) *MockdatabaseNamespace {
	mock := &MockdatabaseNamespace{ctrl: ctrl}
	mock.recorder = &_MockdatabaseNamespaceRecorder{mock}
	return mock
}

func (_m *MockdatabaseNamespace) EXPECT() *_MockdatabaseNamespaceRecorder {
	return _m.recorder
}

func (_m *MockdatabaseNamespace) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockdatabaseNamespaceRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockdatabaseNamespace) NumSeries() int64 {
	ret := _m.ctrl.Call(_m, "NumSeries")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockdatabaseNamespaceRecorder) NumSeries() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NumSeries")
}

func (_m *MockdatabaseNamespace) Tick(softDeadline time.Duration) {
	_m.ctrl.Call(_m, "Tick", softDeadline)
}

func (_mr *_MockdatabaseNamespaceRecorder) Tick(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tick", arg0)
}

func (_m *MockdatabaseNamespace) Write(ctx context.Context, id string, timestamp time.Time, value float64, unit time0.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, id, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseNamespaceRecorder) Write(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockdatabaseNamespace) ReadEncoded(ctx context.Context, id string, start time.Time, end time.Time) ([][]io.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, id, start, end)
	ret0, _ := ret[0].([][]io.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdatabaseNamespaceRecorder) ReadEncoded(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2, arg3)
}

func (_m *MockdatabaseNamespace) FetchBlocks(ctx context.Context, shardID uint32, id string, starts []time.Time) ([]block.FetchBlockResult, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, shardID, id, starts)
	ret0, _ := ret[0].([]block.FetchBlockResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdatabaseNamespaceRecorder) FetchBlocks(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1, arg2, arg3)
}

func (_m *MockdatabaseNamespace) FetchBlocksMetadata(ctx context.Context, shardID uint32, limit int64, pageToken int64, includeSizes bool, includeChecksums bool) ([]block.FetchBlocksMetadataResult, *int64, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, shardID, limit, pageToken, includeSizes, includeChecksums)
	ret0, _ := ret[0].([]block.FetchBlocksMetadataResult)
	ret1, _ := ret[1].(*int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockdatabaseNamespaceRecorder) FetchBlocksMetadata(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockdatabaseNamespace) Bootstrap(bs bootstrap.Bootstrap, targetRanges time0.Ranges, writeStart time.Time) error {
	ret := _m.ctrl.Call(_m, "Bootstrap", bs, targetRanges, writeStart)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseNamespaceRecorder) Bootstrap(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap", arg0, arg1, arg2)
}

func (_m *MockdatabaseNamespace) Flush(blockStart time.Time, pm persist.Manager) error {
	ret := _m.ctrl.Call(_m, "Flush", blockStart, pm)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseNamespaceRecorder) Flush(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0, arg1)
}

func (_m *MockdatabaseNamespace) CleanupFileset(earliestToRetain time.Time) error {
	ret := _m.ctrl.Call(_m, "CleanupFileset", earliestToRetain)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseNamespaceRecorder) CleanupFileset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CleanupFileset", arg0)
}

func (_m *MockdatabaseNamespace) Truncate() (int64, error) {
	ret := _m.ctrl.Call(_m, "Truncate")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdatabaseNamespaceRecorder) Truncate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Truncate")
}

func (_m *MockdatabaseNamespace) Repair(repairer databaseShardRepairer) error {
	ret := _m.ctrl.Call(_m, "Repair", repairer)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseNamespaceRecorder) Repair(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Repair", arg0)
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

func (_m *MockdatabaseShard) NumSeries() int64 {
	ret := _m.ctrl.Call(_m, "NumSeries")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) NumSeries() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NumSeries")
}

func (_m *MockdatabaseShard) Tick(softDeadline time.Duration) {
	_m.ctrl.Call(_m, "Tick", softDeadline)
}

func (_mr *_MockdatabaseShardRecorder) Tick(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tick", arg0)
}

func (_m *MockdatabaseShard) Write(ctx context.Context, id string, timestamp time.Time, value float64, unit time0.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, id, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) Write(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockdatabaseShard) ReadEncoded(ctx context.Context, id string, start time.Time, end time.Time) ([][]io.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, id, start, end)
	ret0, _ := ret[0].([][]io.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdatabaseShardRecorder) ReadEncoded(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2, arg3)
}

func (_m *MockdatabaseShard) FetchBlocks(ctx context.Context, id string, starts []time.Time) []block.FetchBlockResult {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, id, starts)
	ret0, _ := ret[0].([]block.FetchBlockResult)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) FetchBlocks(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1, arg2)
}

func (_m *MockdatabaseShard) FetchBlocksMetadata(ctx context.Context, limit int64, pageToken int64, includeSizes bool, includeChecksums bool) ([]block.FetchBlocksMetadataResult, *int64) {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, limit, pageToken, includeSizes, includeChecksums)
	ret0, _ := ret[0].([]block.FetchBlocksMetadataResult)
	ret1, _ := ret[1].(*int64)
	return ret0, ret1
}

func (_mr *_MockdatabaseShardRecorder) FetchBlocksMetadata(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockdatabaseShard) Bootstrap(bootstrappedSeries map[string]block.DatabaseSeriesBlocks, writeStart time.Time) error {
	ret := _m.ctrl.Call(_m, "Bootstrap", bootstrappedSeries, writeStart)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) Bootstrap(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap", arg0, arg1)
}

func (_m *MockdatabaseShard) Flush(namespace string, blockStart time.Time, pm persist.Manager) error {
	ret := _m.ctrl.Call(_m, "Flush", namespace, blockStart, pm)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) Flush(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0, arg1, arg2)
}

func (_m *MockdatabaseShard) CleanupFileset(namespace string, earliestToRetain time.Time) error {
	ret := _m.ctrl.Call(_m, "CleanupFileset", namespace, earliestToRetain)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) CleanupFileset(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CleanupFileset", arg0, arg1)
}

func (_m *MockdatabaseShard) Repair(namespace string, repairer databaseShardRepairer) (repair.MetadataComparisonResult, error) {
	ret := _m.ctrl.Call(_m, "Repair", namespace, repairer)
	ret0, _ := ret[0].(repair.MetadataComparisonResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdatabaseShardRecorder) Repair(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Repair", arg0, arg1)
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

func (_m *MockdatabaseFlushManager) HasFlushed(t time.Time) bool {
	ret := _m.ctrl.Call(_m, "HasFlushed", t)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseFlushManagerRecorder) HasFlushed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasFlushed", arg0)
}

func (_m *MockdatabaseFlushManager) FlushTimeStart(t time.Time) time.Time {
	ret := _m.ctrl.Call(_m, "FlushTimeStart", t)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockdatabaseFlushManagerRecorder) FlushTimeStart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushTimeStart", arg0)
}

func (_m *MockdatabaseFlushManager) FlushTimeEnd(t time.Time) time.Time {
	ret := _m.ctrl.Call(_m, "FlushTimeEnd", t)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockdatabaseFlushManagerRecorder) FlushTimeEnd(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushTimeEnd", arg0)
}

func (_m *MockdatabaseFlushManager) Flush(t time.Time) error {
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

func (_m *MockdatabaseCleanupManager) Cleanup(t time.Time) error {
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

func (_m *MockdatabaseFileSystemManager) HasFlushed(t time.Time) bool {
	ret := _m.ctrl.Call(_m, "HasFlushed", t)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) HasFlushed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasFlushed", arg0)
}

func (_m *MockdatabaseFileSystemManager) FlushTimeStart(t time.Time) time.Time {
	ret := _m.ctrl.Call(_m, "FlushTimeStart", t)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) FlushTimeStart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushTimeStart", arg0)
}

func (_m *MockdatabaseFileSystemManager) FlushTimeEnd(t time.Time) time.Time {
	ret := _m.ctrl.Call(_m, "FlushTimeEnd", t)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) FlushTimeEnd(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushTimeEnd", arg0)
}

func (_m *MockdatabaseFileSystemManager) Flush(t time.Time) error {
	ret := _m.ctrl.Call(_m, "Flush", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) Flush(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0)
}

func (_m *MockdatabaseFileSystemManager) Cleanup(t time.Time) error {
	ret := _m.ctrl.Call(_m, "Cleanup", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) Cleanup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Cleanup", arg0)
}

func (_m *MockdatabaseFileSystemManager) ShouldRun(t time.Time) bool {
	ret := _m.ctrl.Call(_m, "ShouldRun", t)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) ShouldRun(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShouldRun", arg0)
}

func (_m *MockdatabaseFileSystemManager) Run(t time.Time, async bool) {
	_m.ctrl.Call(_m, "Run", t, async)
}

func (_mr *_MockdatabaseFileSystemManagerRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0, arg1)
}

// Mock of databaseShardRepairer interface
type MockdatabaseShardRepairer struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseShardRepairerRecorder
}

// Recorder for MockdatabaseShardRepairer (not exported)
type _MockdatabaseShardRepairerRecorder struct {
	mock *MockdatabaseShardRepairer
}

func NewMockdatabaseShardRepairer(ctrl *gomock.Controller) *MockdatabaseShardRepairer {
	mock := &MockdatabaseShardRepairer{ctrl: ctrl}
	mock.recorder = &_MockdatabaseShardRepairerRecorder{mock}
	return mock
}

func (_m *MockdatabaseShardRepairer) EXPECT() *_MockdatabaseShardRepairerRecorder {
	return _m.recorder
}

func (_m *MockdatabaseShardRepairer) Options() repair.Options {
	ret := _m.ctrl.Call(_m, "Options")
	ret0, _ := ret[0].(repair.Options)
	return ret0
}

func (_mr *_MockdatabaseShardRepairerRecorder) Options() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Options")
}

func (_m *MockdatabaseShardRepairer) Repair(namespace string, shard databaseShard) (repair.MetadataComparisonResult, error) {
	ret := _m.ctrl.Call(_m, "Repair", namespace, shard)
	ret0, _ := ret[0].(repair.MetadataComparisonResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdatabaseShardRepairerRecorder) Repair(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Repair", arg0, arg1)
}

// Mock of databaseRepairer interface
type MockdatabaseRepairer struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseRepairerRecorder
}

// Recorder for MockdatabaseRepairer (not exported)
type _MockdatabaseRepairerRecorder struct {
	mock *MockdatabaseRepairer
}

func NewMockdatabaseRepairer(ctrl *gomock.Controller) *MockdatabaseRepairer {
	mock := &MockdatabaseRepairer{ctrl: ctrl}
	mock.recorder = &_MockdatabaseRepairerRecorder{mock}
	return mock
}

func (_m *MockdatabaseRepairer) EXPECT() *_MockdatabaseRepairerRecorder {
	return _m.recorder
}

func (_m *MockdatabaseRepairer) Start() {
	_m.ctrl.Call(_m, "Start")
}

func (_mr *_MockdatabaseRepairerRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockdatabaseRepairer) Stop() {
	_m.ctrl.Call(_m, "Stop")
}

func (_mr *_MockdatabaseRepairerRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

func (_m *MockdatabaseRepairer) Repair() error {
	ret := _m.ctrl.Call(_m, "Repair")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseRepairerRecorder) Repair() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Repair")
}

func (_m *MockdatabaseRepairer) IsRepairing() bool {
	ret := _m.ctrl.Call(_m, "IsRepairing")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseRepairerRecorder) IsRepairing() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsRepairing")
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

func (_m *MockOptions) SetClockOptions(value clock.Options) Options {
	ret := _m.ctrl.Call(_m, "SetClockOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetClockOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetClockOptions", arg0)
}

func (_m *MockOptions) ClockOptions() clock.Options {
	ret := _m.ctrl.Call(_m, "ClockOptions")
	ret0, _ := ret[0].(clock.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ClockOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClockOptions")
}

func (_m *MockOptions) SetInstrumentOptions(value instrument.Options) Options {
	ret := _m.ctrl.Call(_m, "SetInstrumentOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetInstrumentOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetInstrumentOptions", arg0)
}

func (_m *MockOptions) InstrumentOptions() instrument.Options {
	ret := _m.ctrl.Call(_m, "InstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) InstrumentOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstrumentOptions")
}

func (_m *MockOptions) SetRetentionOptions(value retention.Options) Options {
	ret := _m.ctrl.Call(_m, "SetRetentionOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetRetentionOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRetentionOptions", arg0)
}

func (_m *MockOptions) RetentionOptions() retention.Options {
	ret := _m.ctrl.Call(_m, "RetentionOptions")
	ret0, _ := ret[0].(retention.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) RetentionOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetentionOptions")
}

func (_m *MockOptions) SetDatabaseBlockOptions(value block.Options) Options {
	ret := _m.ctrl.Call(_m, "SetDatabaseBlockOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetDatabaseBlockOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDatabaseBlockOptions", arg0)
}

func (_m *MockOptions) DatabaseBlockOptions() block.Options {
	ret := _m.ctrl.Call(_m, "DatabaseBlockOptions")
	ret0, _ := ret[0].(block.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseBlockOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseBlockOptions")
}

func (_m *MockOptions) SetCommitLogOptions(value commitlog.Options) Options {
	ret := _m.ctrl.Call(_m, "SetCommitLogOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetCommitLogOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetCommitLogOptions", arg0)
}

func (_m *MockOptions) CommitLogOptions() commitlog.Options {
	ret := _m.ctrl.Call(_m, "CommitLogOptions")
	ret0, _ := ret[0].(commitlog.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) CommitLogOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CommitLogOptions")
}

func (_m *MockOptions) SetRepairOptions(value repair.Options) Options {
	ret := _m.ctrl.Call(_m, "SetRepairOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetRepairOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRepairOptions", arg0)
}

func (_m *MockOptions) RepairOptions() repair.Options {
	ret := _m.ctrl.Call(_m, "RepairOptions")
	ret0, _ := ret[0].(repair.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) RepairOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RepairOptions")
}

func (_m *MockOptions) SetEncodingM3TSZPooled() Options {
	ret := _m.ctrl.Call(_m, "SetEncodingM3TSZPooled")
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetEncodingM3TSZPooled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetEncodingM3TSZPooled")
}

func (_m *MockOptions) SetEncodingM3TSZ() Options {
	ret := _m.ctrl.Call(_m, "SetEncodingM3TSZ")
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetEncodingM3TSZ() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetEncodingM3TSZ")
}

func (_m *MockOptions) SetNewEncoderFn(value encoding.NewEncoderFn) Options {
	ret := _m.ctrl.Call(_m, "SetNewEncoderFn", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetNewEncoderFn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetNewEncoderFn", arg0)
}

func (_m *MockOptions) NewEncoderFn() encoding.NewEncoderFn {
	ret := _m.ctrl.Call(_m, "NewEncoderFn")
	ret0, _ := ret[0].(encoding.NewEncoderFn)
	return ret0
}

func (_mr *_MockOptionsRecorder) NewEncoderFn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewEncoderFn")
}

func (_m *MockOptions) SetNewDecoderFn(value encoding.NewDecoderFn) Options {
	ret := _m.ctrl.Call(_m, "SetNewDecoderFn", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetNewDecoderFn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetNewDecoderFn", arg0)
}

func (_m *MockOptions) NewDecoderFn() encoding.NewDecoderFn {
	ret := _m.ctrl.Call(_m, "NewDecoderFn")
	ret0, _ := ret[0].(encoding.NewDecoderFn)
	return ret0
}

func (_mr *_MockOptionsRecorder) NewDecoderFn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewDecoderFn")
}

func (_m *MockOptions) SetNewBootstrapFn(value NewBootstrapFn) Options {
	ret := _m.ctrl.Call(_m, "SetNewBootstrapFn", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetNewBootstrapFn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetNewBootstrapFn", arg0)
}

func (_m *MockOptions) NewBootstrapFn() NewBootstrapFn {
	ret := _m.ctrl.Call(_m, "NewBootstrapFn")
	ret0, _ := ret[0].(NewBootstrapFn)
	return ret0
}

func (_mr *_MockOptionsRecorder) NewBootstrapFn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewBootstrapFn")
}

func (_m *MockOptions) SetNewPersistManagerFn(value NewPersistManagerFn) Options {
	ret := _m.ctrl.Call(_m, "SetNewPersistManagerFn", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetNewPersistManagerFn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetNewPersistManagerFn", arg0)
}

func (_m *MockOptions) NewPersistManagerFn() NewPersistManagerFn {
	ret := _m.ctrl.Call(_m, "NewPersistManagerFn")
	ret0, _ := ret[0].(NewPersistManagerFn)
	return ret0
}

func (_mr *_MockOptionsRecorder) NewPersistManagerFn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewPersistManagerFn")
}

func (_m *MockOptions) SetMaxFlushRetries(value int) Options {
	ret := _m.ctrl.Call(_m, "SetMaxFlushRetries", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetMaxFlushRetries(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetMaxFlushRetries", arg0)
}

func (_m *MockOptions) MaxFlushRetries() int {
	ret := _m.ctrl.Call(_m, "MaxFlushRetries")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) MaxFlushRetries() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MaxFlushRetries")
}

func (_m *MockOptions) SetContextPool(value context.Pool) Options {
	ret := _m.ctrl.Call(_m, "SetContextPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetContextPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetContextPool", arg0)
}

func (_m *MockOptions) ContextPool() context.Pool {
	ret := _m.ctrl.Call(_m, "ContextPool")
	ret0, _ := ret[0].(context.Pool)
	return ret0
}

func (_mr *_MockOptionsRecorder) ContextPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContextPool")
}

func (_m *MockOptions) SetDatabaseSeriesPool(value series.DatabaseSeriesPool) Options {
	ret := _m.ctrl.Call(_m, "SetDatabaseSeriesPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetDatabaseSeriesPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDatabaseSeriesPool", arg0)
}

func (_m *MockOptions) DatabaseSeriesPool() series.DatabaseSeriesPool {
	ret := _m.ctrl.Call(_m, "DatabaseSeriesPool")
	ret0, _ := ret[0].(series.DatabaseSeriesPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseSeriesPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseSeriesPool")
}

func (_m *MockOptions) SetBytesPool(value pool.BytesPool) Options {
	ret := _m.ctrl.Call(_m, "SetBytesPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetBytesPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetBytesPool", arg0)
}

func (_m *MockOptions) BytesPool() pool.BytesPool {
	ret := _m.ctrl.Call(_m, "BytesPool")
	ret0, _ := ret[0].(pool.BytesPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) BytesPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BytesPool")
}

func (_m *MockOptions) SetEncoderPool(value encoding.EncoderPool) Options {
	ret := _m.ctrl.Call(_m, "SetEncoderPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetEncoderPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetEncoderPool", arg0)
}

func (_m *MockOptions) EncoderPool() encoding.EncoderPool {
	ret := _m.ctrl.Call(_m, "EncoderPool")
	ret0, _ := ret[0].(encoding.EncoderPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) EncoderPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EncoderPool")
}

func (_m *MockOptions) SetSegmentReaderPool(value io.SegmentReaderPool) Options {
	ret := _m.ctrl.Call(_m, "SetSegmentReaderPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetSegmentReaderPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSegmentReaderPool", arg0)
}

func (_m *MockOptions) SegmentReaderPool() io.SegmentReaderPool {
	ret := _m.ctrl.Call(_m, "SegmentReaderPool")
	ret0, _ := ret[0].(io.SegmentReaderPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) SegmentReaderPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SegmentReaderPool")
}

func (_m *MockOptions) SetReaderIteratorPool(value encoding.ReaderIteratorPool) Options {
	ret := _m.ctrl.Call(_m, "SetReaderIteratorPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetReaderIteratorPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetReaderIteratorPool", arg0)
}

func (_m *MockOptions) ReaderIteratorPool() encoding.ReaderIteratorPool {
	ret := _m.ctrl.Call(_m, "ReaderIteratorPool")
	ret0, _ := ret[0].(encoding.ReaderIteratorPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) ReaderIteratorPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReaderIteratorPool")
}

func (_m *MockOptions) SetMultiReaderIteratorPool(value encoding.MultiReaderIteratorPool) Options {
	ret := _m.ctrl.Call(_m, "SetMultiReaderIteratorPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetMultiReaderIteratorPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetMultiReaderIteratorPool", arg0)
}

func (_m *MockOptions) MultiReaderIteratorPool() encoding.MultiReaderIteratorPool {
	ret := _m.ctrl.Call(_m, "MultiReaderIteratorPool")
	ret0, _ := ret[0].(encoding.MultiReaderIteratorPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) MultiReaderIteratorPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MultiReaderIteratorPool")
}
