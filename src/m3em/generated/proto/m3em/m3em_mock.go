// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/m3em/generated/proto/m3em (interfaces: OperatorClient,Operator_PushFileClient,Operator_PullFileClient,Operator_PullFileServer)

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

// Package m3em is a generated GoMock package.
package m3em

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// MockOperatorClient is a mock of OperatorClient interface.
type MockOperatorClient struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorClientMockRecorder
}

// MockOperatorClientMockRecorder is the mock recorder for MockOperatorClient.
type MockOperatorClientMockRecorder struct {
	mock *MockOperatorClient
}

// NewMockOperatorClient creates a new mock instance.
func NewMockOperatorClient(ctrl *gomock.Controller) *MockOperatorClient {
	mock := &MockOperatorClient{ctrl: ctrl}
	mock.recorder = &MockOperatorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorClient) EXPECT() *MockOperatorClientMockRecorder {
	return m.recorder
}

// PullFile mocks base method.
func (m *MockOperatorClient) PullFile(arg0 context.Context, arg1 *PullFileRequest, arg2 ...grpc.CallOption) (Operator_PullFileClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PullFile", varargs...)
	ret0, _ := ret[0].(Operator_PullFileClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullFile indicates an expected call of PullFile.
func (mr *MockOperatorClientMockRecorder) PullFile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullFile", reflect.TypeOf((*MockOperatorClient)(nil).PullFile), varargs...)
}

// PushFile mocks base method.
func (m *MockOperatorClient) PushFile(arg0 context.Context, arg1 ...grpc.CallOption) (Operator_PushFileClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PushFile", varargs...)
	ret0, _ := ret[0].(Operator_PushFileClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PushFile indicates an expected call of PushFile.
func (mr *MockOperatorClientMockRecorder) PushFile(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushFile", reflect.TypeOf((*MockOperatorClient)(nil).PushFile), varargs...)
}

// Setup mocks base method.
func (m *MockOperatorClient) Setup(arg0 context.Context, arg1 *SetupRequest, arg2 ...grpc.CallOption) (*SetupResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Setup", varargs...)
	ret0, _ := ret[0].(*SetupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Setup indicates an expected call of Setup.
func (mr *MockOperatorClientMockRecorder) Setup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockOperatorClient)(nil).Setup), varargs...)
}

// Start mocks base method.
func (m *MockOperatorClient) Start(arg0 context.Context, arg1 *StartRequest, arg2 ...grpc.CallOption) (*StartResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Start", varargs...)
	ret0, _ := ret[0].(*StartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *MockOperatorClientMockRecorder) Start(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockOperatorClient)(nil).Start), varargs...)
}

// Stop mocks base method.
func (m *MockOperatorClient) Stop(arg0 context.Context, arg1 *StopRequest, arg2 ...grpc.CallOption) (*StopResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stop", varargs...)
	ret0, _ := ret[0].(*StopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *MockOperatorClientMockRecorder) Stop(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockOperatorClient)(nil).Stop), varargs...)
}

// Teardown mocks base method.
func (m *MockOperatorClient) Teardown(arg0 context.Context, arg1 *TeardownRequest, arg2 ...grpc.CallOption) (*TeardownResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Teardown", varargs...)
	ret0, _ := ret[0].(*TeardownResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Teardown indicates an expected call of Teardown.
func (mr *MockOperatorClientMockRecorder) Teardown(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teardown", reflect.TypeOf((*MockOperatorClient)(nil).Teardown), varargs...)
}

// MockOperator_PushFileClient is a mock of Operator_PushFileClient interface.
type MockOperator_PushFileClient struct {
	ctrl     *gomock.Controller
	recorder *MockOperator_PushFileClientMockRecorder
}

// MockOperator_PushFileClientMockRecorder is the mock recorder for MockOperator_PushFileClient.
type MockOperator_PushFileClientMockRecorder struct {
	mock *MockOperator_PushFileClient
}

// NewMockOperator_PushFileClient creates a new mock instance.
func NewMockOperator_PushFileClient(ctrl *gomock.Controller) *MockOperator_PushFileClient {
	mock := &MockOperator_PushFileClient{ctrl: ctrl}
	mock.recorder = &MockOperator_PushFileClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperator_PushFileClient) EXPECT() *MockOperator_PushFileClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method.
func (m *MockOperator_PushFileClient) CloseAndRecv() (*PushFileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*PushFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv.
func (mr *MockOperator_PushFileClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockOperator_PushFileClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method.
func (m *MockOperator_PushFileClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockOperator_PushFileClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockOperator_PushFileClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockOperator_PushFileClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockOperator_PushFileClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockOperator_PushFileClient)(nil).Context))
}

// Header mocks base method.
func (m *MockOperator_PushFileClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockOperator_PushFileClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockOperator_PushFileClient)(nil).Header))
}

// RecvMsg mocks base method.
func (m *MockOperator_PushFileClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockOperator_PushFileClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockOperator_PushFileClient)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockOperator_PushFileClient) Send(arg0 *PushFileRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockOperator_PushFileClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockOperator_PushFileClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m *MockOperator_PushFileClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockOperator_PushFileClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockOperator_PushFileClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockOperator_PushFileClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockOperator_PushFileClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockOperator_PushFileClient)(nil).Trailer))
}

// MockOperator_PullFileClient is a mock of Operator_PullFileClient interface.
type MockOperator_PullFileClient struct {
	ctrl     *gomock.Controller
	recorder *MockOperator_PullFileClientMockRecorder
}

// MockOperator_PullFileClientMockRecorder is the mock recorder for MockOperator_PullFileClient.
type MockOperator_PullFileClientMockRecorder struct {
	mock *MockOperator_PullFileClient
}

// NewMockOperator_PullFileClient creates a new mock instance.
func NewMockOperator_PullFileClient(ctrl *gomock.Controller) *MockOperator_PullFileClient {
	mock := &MockOperator_PullFileClient{ctrl: ctrl}
	mock.recorder = &MockOperator_PullFileClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperator_PullFileClient) EXPECT() *MockOperator_PullFileClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockOperator_PullFileClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockOperator_PullFileClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockOperator_PullFileClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockOperator_PullFileClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockOperator_PullFileClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockOperator_PullFileClient)(nil).Context))
}

// Header mocks base method.
func (m *MockOperator_PullFileClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockOperator_PullFileClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockOperator_PullFileClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockOperator_PullFileClient) Recv() (*PullFileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*PullFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockOperator_PullFileClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockOperator_PullFileClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockOperator_PullFileClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockOperator_PullFileClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockOperator_PullFileClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockOperator_PullFileClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockOperator_PullFileClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockOperator_PullFileClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockOperator_PullFileClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockOperator_PullFileClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockOperator_PullFileClient)(nil).Trailer))
}

// MockOperator_PullFileServer is a mock of Operator_PullFileServer interface.
type MockOperator_PullFileServer struct {
	ctrl     *gomock.Controller
	recorder *MockOperator_PullFileServerMockRecorder
}

// MockOperator_PullFileServerMockRecorder is the mock recorder for MockOperator_PullFileServer.
type MockOperator_PullFileServerMockRecorder struct {
	mock *MockOperator_PullFileServer
}

// NewMockOperator_PullFileServer creates a new mock instance.
func NewMockOperator_PullFileServer(ctrl *gomock.Controller) *MockOperator_PullFileServer {
	mock := &MockOperator_PullFileServer{ctrl: ctrl}
	mock.recorder = &MockOperator_PullFileServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperator_PullFileServer) EXPECT() *MockOperator_PullFileServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockOperator_PullFileServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockOperator_PullFileServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockOperator_PullFileServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockOperator_PullFileServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockOperator_PullFileServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockOperator_PullFileServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockOperator_PullFileServer) Send(arg0 *PullFileResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockOperator_PullFileServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockOperator_PullFileServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockOperator_PullFileServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockOperator_PullFileServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockOperator_PullFileServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockOperator_PullFileServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockOperator_PullFileServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockOperator_PullFileServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockOperator_PullFileServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockOperator_PullFileServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockOperator_PullFileServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockOperator_PullFileServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockOperator_PullFileServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockOperator_PullFileServer)(nil).SetTrailer), arg0)
}
