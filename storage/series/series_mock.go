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
// Source: github.com/m3db/m3db/storage/series (interfaces: DatabaseSeries,QueryableBlockRetriever)

package series

import (
	"reflect"
	"time"

	"github.com/m3db/m3db/persist"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3db/x/xio"
	"github.com/m3db/m3x/context"
	"github.com/m3db/m3x/ident"
	time0 "github.com/m3db/m3x/time"

	"github.com/golang/mock/gomock"
)

// MockDatabaseSeries is a mock of DatabaseSeries interface
type MockDatabaseSeries struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseSeriesMockRecorder
}

// MockDatabaseSeriesMockRecorder is the mock recorder for MockDatabaseSeries
type MockDatabaseSeriesMockRecorder struct {
	mock *MockDatabaseSeries
}

// NewMockDatabaseSeries creates a new mock instance
func NewMockDatabaseSeries(ctrl *gomock.Controller) *MockDatabaseSeries {
	mock := &MockDatabaseSeries{ctrl: ctrl}
	mock.recorder = &MockDatabaseSeriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDatabaseSeries) EXPECT() *MockDatabaseSeriesMockRecorder {
	return _m.recorder
}

// Bootstrap mocks base method
func (_m *MockDatabaseSeries) Bootstrap(_param0 block.DatabaseSeriesBlocks) error {
	ret := _m.ctrl.Call(_m, "Bootstrap", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bootstrap indicates an expected call of Bootstrap
func (_mr *MockDatabaseSeriesMockRecorder) Bootstrap(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Bootstrap", reflect.TypeOf((*MockDatabaseSeries)(nil).Bootstrap), arg0)
}

// Close mocks base method
func (_m *MockDatabaseSeries) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockDatabaseSeriesMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockDatabaseSeries)(nil).Close))
}

// FetchBlocks mocks base method
func (_m *MockDatabaseSeries) FetchBlocks(_param0 context.Context, _param1 []time.Time) ([]block.FetchBlockResult, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocks", _param0, _param1)
	ret0, _ := ret[0].([]block.FetchBlockResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBlocks indicates an expected call of FetchBlocks
func (_mr *MockDatabaseSeriesMockRecorder) FetchBlocks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FetchBlocks", reflect.TypeOf((*MockDatabaseSeries)(nil).FetchBlocks), arg0, arg1)
}

// FetchBlocksMetadata mocks base method
func (_m *MockDatabaseSeries) FetchBlocksMetadata(_param0 context.Context, _param1 time.Time, _param2 time.Time, _param3 FetchBlocksMetadataOptions) block.FetchBlocksMetadataResult {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(block.FetchBlocksMetadataResult)
	return ret0
}

// FetchBlocksMetadata indicates an expected call of FetchBlocksMetadata
func (_mr *MockDatabaseSeriesMockRecorder) FetchBlocksMetadata(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FetchBlocksMetadata", reflect.TypeOf((*MockDatabaseSeries)(nil).FetchBlocksMetadata), arg0, arg1, arg2, arg3)
}

// Flush mocks base method
func (_m *MockDatabaseSeries) Flush(_param0 context.Context, _param1 time.Time, _param2 persist.Fn) error {
	ret := _m.ctrl.Call(_m, "Flush", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (_mr *MockDatabaseSeriesMockRecorder) Flush(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Flush", reflect.TypeOf((*MockDatabaseSeries)(nil).Flush), arg0, arg1, arg2)
}

// ID mocks base method
func (_m *MockDatabaseSeries) ID() ident.ID {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(ident.ID)
	return ret0
}

// ID indicates an expected call of ID
func (_mr *MockDatabaseSeriesMockRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ID", reflect.TypeOf((*MockDatabaseSeries)(nil).ID))
}

// IsBootstrapped mocks base method
func (_m *MockDatabaseSeries) IsBootstrapped() bool {
	ret := _m.ctrl.Call(_m, "IsBootstrapped")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBootstrapped indicates an expected call of IsBootstrapped
func (_mr *MockDatabaseSeriesMockRecorder) IsBootstrapped() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "IsBootstrapped", reflect.TypeOf((*MockDatabaseSeries)(nil).IsBootstrapped))
}

// IsEmpty mocks base method
func (_m *MockDatabaseSeries) IsEmpty() bool {
	ret := _m.ctrl.Call(_m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmpty indicates an expected call of IsEmpty
func (_mr *MockDatabaseSeriesMockRecorder) IsEmpty() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "IsEmpty", reflect.TypeOf((*MockDatabaseSeries)(nil).IsEmpty))
}

// NumActiveBlocks mocks base method
func (_m *MockDatabaseSeries) NumActiveBlocks() int {
	ret := _m.ctrl.Call(_m, "NumActiveBlocks")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumActiveBlocks indicates an expected call of NumActiveBlocks
func (_mr *MockDatabaseSeriesMockRecorder) NumActiveBlocks() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "NumActiveBlocks", reflect.TypeOf((*MockDatabaseSeries)(nil).NumActiveBlocks))
}

// OnEvictedFromWiredList mocks base method
func (_m *MockDatabaseSeries) OnEvictedFromWiredList(_param0 ident.ID, _param1 time.Time) {
	_m.ctrl.Call(_m, "OnEvictedFromWiredList", _param0, _param1)
}

// OnEvictedFromWiredList indicates an expected call of OnEvictedFromWiredList
func (_mr *MockDatabaseSeriesMockRecorder) OnEvictedFromWiredList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "OnEvictedFromWiredList", reflect.TypeOf((*MockDatabaseSeries)(nil).OnEvictedFromWiredList), arg0, arg1)
}

// OnRetrieveBlock mocks base method
func (_m *MockDatabaseSeries) OnRetrieveBlock(_param0 ident.ID, _param1 time.Time, _param2 ts.Segment) {
	_m.ctrl.Call(_m, "OnRetrieveBlock", _param0, _param1, _param2)
}

// OnRetrieveBlock indicates an expected call of OnRetrieveBlock
func (_mr *MockDatabaseSeriesMockRecorder) OnRetrieveBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "OnRetrieveBlock", reflect.TypeOf((*MockDatabaseSeries)(nil).OnRetrieveBlock), arg0, arg1, arg2)
}

// ReadEncoded mocks base method
func (_m *MockDatabaseSeries) ReadEncoded(_param0 context.Context, _param1 time.Time, _param2 time.Time) ([][]xio.Block, error) {
	ret := _m.ctrl.Call(_m, "ReadEncoded", _param0, _param1, _param2)
	ret0, _ := ret[0].([][]xio.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadEncoded indicates an expected call of ReadEncoded
func (_mr *MockDatabaseSeriesMockRecorder) ReadEncoded(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadEncoded", reflect.TypeOf((*MockDatabaseSeries)(nil).ReadEncoded), arg0, arg1, arg2)
}

// Reset mocks base method
func (_m *MockDatabaseSeries) Reset(_param0 ident.ID, _param1 ident.Tags, _param2 QueryableBlockRetriever, _param3 block.OnRetrieveBlock, _param4 block.OnEvictedFromWiredList, _param5 Options) {
	_m.ctrl.Call(_m, "Reset", _param0, _param1, _param2, _param3, _param4, _param5)
}

// Reset indicates an expected call of Reset
func (_mr *MockDatabaseSeriesMockRecorder) Reset(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Reset", reflect.TypeOf((*MockDatabaseSeries)(nil).Reset), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Snapshot mocks base method
func (_m *MockDatabaseSeries) Snapshot(_param0 context.Context, _param1 time.Time, _param2 persist.Fn) error {
	ret := _m.ctrl.Call(_m, "Snapshot", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Snapshot indicates an expected call of Snapshot
func (_mr *MockDatabaseSeriesMockRecorder) Snapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Snapshot", reflect.TypeOf((*MockDatabaseSeries)(nil).Snapshot), arg0, arg1, arg2)
}

// Tags mocks base method
func (_m *MockDatabaseSeries) Tags() ident.Tags {
	ret := _m.ctrl.Call(_m, "Tags")
	ret0, _ := ret[0].(ident.Tags)
	return ret0
}

// Tags indicates an expected call of Tags
func (_mr *MockDatabaseSeriesMockRecorder) Tags() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Tags", reflect.TypeOf((*MockDatabaseSeries)(nil).Tags))
}

// Tick mocks base method
func (_m *MockDatabaseSeries) Tick() (TickResult, error) {
	ret := _m.ctrl.Call(_m, "Tick")
	ret0, _ := ret[0].(TickResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tick indicates an expected call of Tick
func (_mr *MockDatabaseSeriesMockRecorder) Tick() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Tick", reflect.TypeOf((*MockDatabaseSeries)(nil).Tick))
}

// Write mocks base method
func (_m *MockDatabaseSeries) Write(_param0 context.Context, _param1 time.Time, _param2 float64, _param3 time0.Unit, _param4 []byte) error {
	ret := _m.ctrl.Call(_m, "Write", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (_mr *MockDatabaseSeriesMockRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Write", reflect.TypeOf((*MockDatabaseSeries)(nil).Write), arg0, arg1, arg2, arg3, arg4)
}

// MockQueryableBlockRetriever is a mock of QueryableBlockRetriever interface
type MockQueryableBlockRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockQueryableBlockRetrieverMockRecorder
}

// MockQueryableBlockRetrieverMockRecorder is the mock recorder for MockQueryableBlockRetriever
type MockQueryableBlockRetrieverMockRecorder struct {
	mock *MockQueryableBlockRetriever
}

// NewMockQueryableBlockRetriever creates a new mock instance
func NewMockQueryableBlockRetriever(ctrl *gomock.Controller) *MockQueryableBlockRetriever {
	mock := &MockQueryableBlockRetriever{ctrl: ctrl}
	mock.recorder = &MockQueryableBlockRetrieverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockQueryableBlockRetriever) EXPECT() *MockQueryableBlockRetrieverMockRecorder {
	return _m.recorder
}

// IsBlockRetrievable mocks base method
func (_m *MockQueryableBlockRetriever) IsBlockRetrievable(_param0 time.Time) bool {
	ret := _m.ctrl.Call(_m, "IsBlockRetrievable", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBlockRetrievable indicates an expected call of IsBlockRetrievable
func (_mr *MockQueryableBlockRetrieverMockRecorder) IsBlockRetrievable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "IsBlockRetrievable", reflect.TypeOf((*MockQueryableBlockRetriever)(nil).IsBlockRetrievable), arg0)
}

// Stream mocks base method
func (_m *MockQueryableBlockRetriever) Stream(_param0 context.Context, _param1 ident.ID, _param2 time.Time, _param3 time.Time, _param4 block.OnRetrieveBlock) (xio.Block, error) {
	ret := _m.ctrl.Call(_m, "Stream", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(xio.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stream indicates an expected call of Stream
func (_mr *MockQueryableBlockRetrieverMockRecorder) Stream(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Stream", reflect.TypeOf((*MockQueryableBlockRetriever)(nil).Stream), arg0, arg1, arg2, arg3, arg4)
}
