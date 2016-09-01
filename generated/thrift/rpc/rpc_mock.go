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
// Source: github.com/m3db/m3db/generated/thrift/rpc/tchan-rpc.go

package rpc

import (
	gomock "github.com/golang/mock/gomock"
	thrift "github.com/uber/tchannel-go/thrift"
)

// Mock of TChanCluster interface
type MockTChanCluster struct {
	ctrl     *gomock.Controller
	recorder *_MockTChanClusterRecorder
}

// Recorder for MockTChanCluster (not exported)
type _MockTChanClusterRecorder struct {
	mock *MockTChanCluster
}

func NewMockTChanCluster(ctrl *gomock.Controller) *MockTChanCluster {
	mock := &MockTChanCluster{ctrl: ctrl}
	mock.recorder = &_MockTChanClusterRecorder{mock}
	return mock
}

func (_m *MockTChanCluster) EXPECT() *_MockTChanClusterRecorder {
	return _m.recorder
}

func (_m *MockTChanCluster) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	ret := _m.ctrl.Call(_m, "Fetch", ctx, req)
	ret0, _ := ret[0].(*FetchResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanClusterRecorder) Fetch(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fetch", arg0, arg1)
}

func (_m *MockTChanCluster) Health(ctx thrift.Context) (*HealthResult_, error) {
	ret := _m.ctrl.Call(_m, "Health", ctx)
	ret0, _ := ret[0].(*HealthResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanClusterRecorder) Health(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Health", arg0)
}

func (_m *MockTChanCluster) Write(ctx thrift.Context, req *WriteRequest) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTChanClusterRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1)
}

// Mock of TChanNode interface
type MockTChanNode struct {
	ctrl     *gomock.Controller
	recorder *_MockTChanNodeRecorder
}

// Recorder for MockTChanNode (not exported)
type _MockTChanNodeRecorder struct {
	mock *MockTChanNode
}

func NewMockTChanNode(ctrl *gomock.Controller) *MockTChanNode {
	mock := &MockTChanNode{ctrl: ctrl}
	mock.recorder = &_MockTChanNodeRecorder{mock}
	return mock
}

func (_m *MockTChanNode) EXPECT() *_MockTChanNodeRecorder {
	return _m.recorder
}

func (_m *MockTChanNode) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	ret := _m.ctrl.Call(_m, "Fetch", ctx, req)
	ret0, _ := ret[0].(*FetchResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) Fetch(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fetch", arg0, arg1)
}

func (_m *MockTChanNode) FetchBlocks(ctx thrift.Context, req *FetchBlocksRequest) (*FetchBlocksResult_, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, req)
	ret0, _ := ret[0].(*FetchBlocksResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) FetchBlocks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1)
}

func (_m *MockTChanNode) FetchBlocksMetadata(ctx thrift.Context, req *FetchBlocksMetadataRequest) (*FetchBlocksMetadataResult_, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, req)
	ret0, _ := ret[0].(*FetchBlocksMetadataResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) FetchBlocksMetadata(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1)
}

func (_m *MockTChanNode) FetchRawBatch(ctx thrift.Context, req *FetchRawBatchRequest) (*FetchRawBatchResult_, error) {
	ret := _m.ctrl.Call(_m, "FetchRawBatch", ctx, req)
	ret0, _ := ret[0].(*FetchRawBatchResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) FetchRawBatch(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchRawBatch", arg0, arg1)
}

func (_m *MockTChanNode) Health(ctx thrift.Context) (*HealthResult_, error) {
	ret := _m.ctrl.Call(_m, "Health", ctx)
	ret0, _ := ret[0].(*HealthResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) Health(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Health", arg0)
}

func (_m *MockTChanNode) TruncateNamespace(ctx thrift.Context, req *TruncateNamespaceRequest) error {
	ret := _m.ctrl.Call(_m, "TruncateNamespace", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTChanNodeRecorder) TruncateNamespace(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TruncateNamespace", arg0, arg1)
}

func (_m *MockTChanNode) Write(ctx thrift.Context, req *WriteRequest) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTChanNodeRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1)
}

func (_m *MockTChanNode) WriteBatch(ctx thrift.Context, req *WriteBatchRequest) error {
	ret := _m.ctrl.Call(_m, "WriteBatch", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTChanNodeRecorder) WriteBatch(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteBatch", arg0, arg1)
}
