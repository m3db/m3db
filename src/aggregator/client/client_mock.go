// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/aggregator/client (interfaces: Client,AdminClient)

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

// Package client is a generated GoMock package.
package client

import (
	"reflect"

	"github.com/m3db/m3/src/metrics/metadata"
	"github.com/m3db/m3/src/metrics/metric/aggregated"
	"github.com/m3db/m3/src/metrics/metric/unaggregated"
	"github.com/m3db/m3/src/metrics/policy"

	"github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// Flush mocks base method
func (m *MockClient) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockClientMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockClient)(nil).Flush))
}

// Init mocks base method
func (m *MockClient) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockClientMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockClient)(nil).Init))
}

// WritePassthrough mocks base method
func (m *MockClient) WritePassthrough(arg0 aggregated.Metric, arg1 policy.StoragePolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WritePassthrough", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WritePassthrough indicates an expected call of WritePassthrough
func (mr *MockClientMockRecorder) WritePassthrough(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePassthrough", reflect.TypeOf((*MockClient)(nil).WritePassthrough), arg0, arg1)
}

// WriteTimed mocks base method
func (m *MockClient) WriteTimed(arg0 aggregated.Metric, arg1 metadata.TimedMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTimed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTimed indicates an expected call of WriteTimed
func (mr *MockClientMockRecorder) WriteTimed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTimed", reflect.TypeOf((*MockClient)(nil).WriteTimed), arg0, arg1)
}

// WriteTimedWithStagedMetadatas mocks base method
func (m *MockClient) WriteTimedWithStagedMetadatas(arg0 aggregated.Metric, arg1 metadata.StagedMetadatas) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTimedWithStagedMetadatas", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTimedWithStagedMetadatas indicates an expected call of WriteTimedWithStagedMetadatas
func (mr *MockClientMockRecorder) WriteTimedWithStagedMetadatas(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTimedWithStagedMetadatas", reflect.TypeOf((*MockClient)(nil).WriteTimedWithStagedMetadatas), arg0, arg1)
}

// WriteUntimedBatchTimer mocks base method
func (m *MockClient) WriteUntimedBatchTimer(arg0 unaggregated.BatchTimer, arg1 metadata.StagedMetadatas) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteUntimedBatchTimer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteUntimedBatchTimer indicates an expected call of WriteUntimedBatchTimer
func (mr *MockClientMockRecorder) WriteUntimedBatchTimer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteUntimedBatchTimer", reflect.TypeOf((*MockClient)(nil).WriteUntimedBatchTimer), arg0, arg1)
}

// WriteUntimedCounter mocks base method
func (m *MockClient) WriteUntimedCounter(arg0 unaggregated.Counter, arg1 metadata.StagedMetadatas) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteUntimedCounter", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteUntimedCounter indicates an expected call of WriteUntimedCounter
func (mr *MockClientMockRecorder) WriteUntimedCounter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteUntimedCounter", reflect.TypeOf((*MockClient)(nil).WriteUntimedCounter), arg0, arg1)
}

// WriteUntimedGauge mocks base method
func (m *MockClient) WriteUntimedGauge(arg0 unaggregated.Gauge, arg1 metadata.StagedMetadatas) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteUntimedGauge", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteUntimedGauge indicates an expected call of WriteUntimedGauge
func (mr *MockClientMockRecorder) WriteUntimedGauge(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteUntimedGauge", reflect.TypeOf((*MockClient)(nil).WriteUntimedGauge), arg0, arg1)
}

// MockAdminClient is a mock of AdminClient interface
type MockAdminClient struct {
	ctrl     *gomock.Controller
	recorder *MockAdminClientMockRecorder
}

// MockAdminClientMockRecorder is the mock recorder for MockAdminClient
type MockAdminClientMockRecorder struct {
	mock *MockAdminClient
}

// NewMockAdminClient creates a new mock instance
func NewMockAdminClient(ctrl *gomock.Controller) *MockAdminClient {
	mock := &MockAdminClient{ctrl: ctrl}
	mock.recorder = &MockAdminClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdminClient) EXPECT() *MockAdminClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockAdminClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockAdminClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAdminClient)(nil).Close))
}

// Flush mocks base method
func (m *MockAdminClient) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockAdminClientMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockAdminClient)(nil).Flush))
}

// Init mocks base method
func (m *MockAdminClient) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockAdminClientMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAdminClient)(nil).Init))
}

// WriteForwarded mocks base method
func (m *MockAdminClient) WriteForwarded(arg0 aggregated.ForwardedMetric, arg1 metadata.ForwardMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteForwarded", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteForwarded indicates an expected call of WriteForwarded
func (mr *MockAdminClientMockRecorder) WriteForwarded(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteForwarded", reflect.TypeOf((*MockAdminClient)(nil).WriteForwarded), arg0, arg1)
}

// WritePassthrough mocks base method
func (m *MockAdminClient) WritePassthrough(arg0 aggregated.Metric, arg1 policy.StoragePolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WritePassthrough", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WritePassthrough indicates an expected call of WritePassthrough
func (mr *MockAdminClientMockRecorder) WritePassthrough(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePassthrough", reflect.TypeOf((*MockAdminClient)(nil).WritePassthrough), arg0, arg1)
}

// WriteTimed mocks base method
func (m *MockAdminClient) WriteTimed(arg0 aggregated.Metric, arg1 metadata.TimedMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTimed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTimed indicates an expected call of WriteTimed
func (mr *MockAdminClientMockRecorder) WriteTimed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTimed", reflect.TypeOf((*MockAdminClient)(nil).WriteTimed), arg0, arg1)
}

// WriteTimedWithStagedMetadatas mocks base method
func (m *MockAdminClient) WriteTimedWithStagedMetadatas(arg0 aggregated.Metric, arg1 metadata.StagedMetadatas) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTimedWithStagedMetadatas", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTimedWithStagedMetadatas indicates an expected call of WriteTimedWithStagedMetadatas
func (mr *MockAdminClientMockRecorder) WriteTimedWithStagedMetadatas(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTimedWithStagedMetadatas", reflect.TypeOf((*MockAdminClient)(nil).WriteTimedWithStagedMetadatas), arg0, arg1)
}

// WriteUntimedBatchTimer mocks base method
func (m *MockAdminClient) WriteUntimedBatchTimer(arg0 unaggregated.BatchTimer, arg1 metadata.StagedMetadatas) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteUntimedBatchTimer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteUntimedBatchTimer indicates an expected call of WriteUntimedBatchTimer
func (mr *MockAdminClientMockRecorder) WriteUntimedBatchTimer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteUntimedBatchTimer", reflect.TypeOf((*MockAdminClient)(nil).WriteUntimedBatchTimer), arg0, arg1)
}

// WriteUntimedCounter mocks base method
func (m *MockAdminClient) WriteUntimedCounter(arg0 unaggregated.Counter, arg1 metadata.StagedMetadatas) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteUntimedCounter", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteUntimedCounter indicates an expected call of WriteUntimedCounter
func (mr *MockAdminClientMockRecorder) WriteUntimedCounter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteUntimedCounter", reflect.TypeOf((*MockAdminClient)(nil).WriteUntimedCounter), arg0, arg1)
}

// WriteUntimedGauge mocks base method
func (m *MockAdminClient) WriteUntimedGauge(arg0 unaggregated.Gauge, arg1 metadata.StagedMetadatas) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteUntimedGauge", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteUntimedGauge indicates an expected call of WriteUntimedGauge
func (mr *MockAdminClientMockRecorder) WriteUntimedGauge(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteUntimedGauge", reflect.TypeOf((*MockAdminClient)(nil).WriteUntimedGauge), arg0, arg1)
}
