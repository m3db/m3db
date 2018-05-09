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
// Source: github.com/m3db/m3db/runtime/types.go

// Package runtime is a generated GoMock package.
package runtime

import (
	"reflect"
	"time"

	"github.com/m3db/m3db/ratelimit"
	"github.com/m3db/m3db/topology"
	"github.com/m3db/m3x/close"

	"github.com/golang/mock/gomock"
)

// MockOptions is a mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsMockRecorder
}

// MockOptionsMockRecorder is the mock recorder for MockOptions
type MockOptionsMockRecorder struct {
	mock *MockOptions
}

// NewMockOptions creates a new mock instance
func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &MockOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOptions) EXPECT() *MockOptionsMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockOptions) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockOptionsMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockOptions)(nil).Validate))
}

// SetPersistRateLimitOptions mocks base method
func (m *MockOptions) SetPersistRateLimitOptions(value ratelimit.Options) Options {
	ret := m.ctrl.Call(m, "SetPersistRateLimitOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetPersistRateLimitOptions indicates an expected call of SetPersistRateLimitOptions
func (mr *MockOptionsMockRecorder) SetPersistRateLimitOptions(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPersistRateLimitOptions", reflect.TypeOf((*MockOptions)(nil).SetPersistRateLimitOptions), value)
}

// PersistRateLimitOptions mocks base method
func (m *MockOptions) PersistRateLimitOptions() ratelimit.Options {
	ret := m.ctrl.Call(m, "PersistRateLimitOptions")
	ret0, _ := ret[0].(ratelimit.Options)
	return ret0
}

// PersistRateLimitOptions indicates an expected call of PersistRateLimitOptions
func (mr *MockOptionsMockRecorder) PersistRateLimitOptions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistRateLimitOptions", reflect.TypeOf((*MockOptions)(nil).PersistRateLimitOptions))
}

// SetWriteNewSeriesAsync mocks base method
func (m *MockOptions) SetWriteNewSeriesAsync(value bool) Options {
	ret := m.ctrl.Call(m, "SetWriteNewSeriesAsync", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetWriteNewSeriesAsync indicates an expected call of SetWriteNewSeriesAsync
func (mr *MockOptionsMockRecorder) SetWriteNewSeriesAsync(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteNewSeriesAsync", reflect.TypeOf((*MockOptions)(nil).SetWriteNewSeriesAsync), value)
}

// WriteNewSeriesAsync mocks base method
func (m *MockOptions) WriteNewSeriesAsync() bool {
	ret := m.ctrl.Call(m, "WriteNewSeriesAsync")
	ret0, _ := ret[0].(bool)
	return ret0
}

// WriteNewSeriesAsync indicates an expected call of WriteNewSeriesAsync
func (mr *MockOptionsMockRecorder) WriteNewSeriesAsync() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteNewSeriesAsync", reflect.TypeOf((*MockOptions)(nil).WriteNewSeriesAsync))
}

// SetWriteNewSeriesBackoffDuration mocks base method
func (m *MockOptions) SetWriteNewSeriesBackoffDuration(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetWriteNewSeriesBackoffDuration", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetWriteNewSeriesBackoffDuration indicates an expected call of SetWriteNewSeriesBackoffDuration
func (mr *MockOptionsMockRecorder) SetWriteNewSeriesBackoffDuration(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteNewSeriesBackoffDuration", reflect.TypeOf((*MockOptions)(nil).SetWriteNewSeriesBackoffDuration), value)
}

// WriteNewSeriesBackoffDuration mocks base method
func (m *MockOptions) WriteNewSeriesBackoffDuration() time.Duration {
	ret := m.ctrl.Call(m, "WriteNewSeriesBackoffDuration")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// WriteNewSeriesBackoffDuration indicates an expected call of WriteNewSeriesBackoffDuration
func (mr *MockOptionsMockRecorder) WriteNewSeriesBackoffDuration() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteNewSeriesBackoffDuration", reflect.TypeOf((*MockOptions)(nil).WriteNewSeriesBackoffDuration))
}

// SetWriteNewSeriesLimitPerShardPerSecond mocks base method
func (m *MockOptions) SetWriteNewSeriesLimitPerShardPerSecond(value int) Options {
	ret := m.ctrl.Call(m, "SetWriteNewSeriesLimitPerShardPerSecond", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetWriteNewSeriesLimitPerShardPerSecond indicates an expected call of SetWriteNewSeriesLimitPerShardPerSecond
func (mr *MockOptionsMockRecorder) SetWriteNewSeriesLimitPerShardPerSecond(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteNewSeriesLimitPerShardPerSecond", reflect.TypeOf((*MockOptions)(nil).SetWriteNewSeriesLimitPerShardPerSecond), value)
}

// WriteNewSeriesLimitPerShardPerSecond mocks base method
func (m *MockOptions) WriteNewSeriesLimitPerShardPerSecond() int {
	ret := m.ctrl.Call(m, "WriteNewSeriesLimitPerShardPerSecond")
	ret0, _ := ret[0].(int)
	return ret0
}

// WriteNewSeriesLimitPerShardPerSecond indicates an expected call of WriteNewSeriesLimitPerShardPerSecond
func (mr *MockOptionsMockRecorder) WriteNewSeriesLimitPerShardPerSecond() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteNewSeriesLimitPerShardPerSecond", reflect.TypeOf((*MockOptions)(nil).WriteNewSeriesLimitPerShardPerSecond))
}

// SetIndexNewSeriesBackoffDuration mocks base method
func (m *MockOptions) SetIndexNewSeriesBackoffDuration(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetIndexNewSeriesBackoffDuration", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetIndexNewSeriesBackoffDuration indicates an expected call of SetIndexNewSeriesBackoffDuration
func (mr *MockOptionsMockRecorder) SetIndexNewSeriesBackoffDuration(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIndexNewSeriesBackoffDuration", reflect.TypeOf((*MockOptions)(nil).SetIndexNewSeriesBackoffDuration), value)
}

// IndexNewSeriesBackoffDuration mocks base method
func (m *MockOptions) IndexNewSeriesBackoffDuration() time.Duration {
	ret := m.ctrl.Call(m, "IndexNewSeriesBackoffDuration")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// IndexNewSeriesBackoffDuration indicates an expected call of IndexNewSeriesBackoffDuration
func (mr *MockOptionsMockRecorder) IndexNewSeriesBackoffDuration() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexNewSeriesBackoffDuration", reflect.TypeOf((*MockOptions)(nil).IndexNewSeriesBackoffDuration))
}

// SetIndexNewSeriesLimitPerSecond mocks base method
func (m *MockOptions) SetIndexNewSeriesLimitPerSecond(value int) Options {
	ret := m.ctrl.Call(m, "SetIndexNewSeriesLimitPerSecond", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetIndexNewSeriesLimitPerSecond indicates an expected call of SetIndexNewSeriesLimitPerSecond
func (mr *MockOptionsMockRecorder) SetIndexNewSeriesLimitPerSecond(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIndexNewSeriesLimitPerSecond", reflect.TypeOf((*MockOptions)(nil).SetIndexNewSeriesLimitPerSecond), value)
}

// IndexNewSeriesLimitPerSecond mocks base method
func (m *MockOptions) IndexNewSeriesLimitPerSecond() int {
	ret := m.ctrl.Call(m, "IndexNewSeriesLimitPerSecond")
	ret0, _ := ret[0].(int)
	return ret0
}

// IndexNewSeriesLimitPerSecond indicates an expected call of IndexNewSeriesLimitPerSecond
func (mr *MockOptionsMockRecorder) IndexNewSeriesLimitPerSecond() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexNewSeriesLimitPerSecond", reflect.TypeOf((*MockOptions)(nil).IndexNewSeriesLimitPerSecond))
}

// SetTickSeriesBatchSize mocks base method
func (m *MockOptions) SetTickSeriesBatchSize(value int) Options {
	ret := m.ctrl.Call(m, "SetTickSeriesBatchSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetTickSeriesBatchSize indicates an expected call of SetTickSeriesBatchSize
func (mr *MockOptionsMockRecorder) SetTickSeriesBatchSize(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTickSeriesBatchSize", reflect.TypeOf((*MockOptions)(nil).SetTickSeriesBatchSize), value)
}

// TickSeriesBatchSize mocks base method
func (m *MockOptions) TickSeriesBatchSize() int {
	ret := m.ctrl.Call(m, "TickSeriesBatchSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// TickSeriesBatchSize indicates an expected call of TickSeriesBatchSize
func (mr *MockOptionsMockRecorder) TickSeriesBatchSize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TickSeriesBatchSize", reflect.TypeOf((*MockOptions)(nil).TickSeriesBatchSize))
}

// SetTickPerSeriesSleepDuration mocks base method
func (m *MockOptions) SetTickPerSeriesSleepDuration(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetTickPerSeriesSleepDuration", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetTickPerSeriesSleepDuration indicates an expected call of SetTickPerSeriesSleepDuration
func (mr *MockOptionsMockRecorder) SetTickPerSeriesSleepDuration(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTickPerSeriesSleepDuration", reflect.TypeOf((*MockOptions)(nil).SetTickPerSeriesSleepDuration), value)
}

// TickPerSeriesSleepDuration mocks base method
func (m *MockOptions) TickPerSeriesSleepDuration() time.Duration {
	ret := m.ctrl.Call(m, "TickPerSeriesSleepDuration")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TickPerSeriesSleepDuration indicates an expected call of TickPerSeriesSleepDuration
func (mr *MockOptionsMockRecorder) TickPerSeriesSleepDuration() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TickPerSeriesSleepDuration", reflect.TypeOf((*MockOptions)(nil).TickPerSeriesSleepDuration))
}

// SetTickMinimumInterval mocks base method
func (m *MockOptions) SetTickMinimumInterval(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetTickMinimumInterval", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetTickMinimumInterval indicates an expected call of SetTickMinimumInterval
func (mr *MockOptionsMockRecorder) SetTickMinimumInterval(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTickMinimumInterval", reflect.TypeOf((*MockOptions)(nil).SetTickMinimumInterval), value)
}

// TickMinimumInterval mocks base method
func (m *MockOptions) TickMinimumInterval() time.Duration {
	ret := m.ctrl.Call(m, "TickMinimumInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TickMinimumInterval indicates an expected call of TickMinimumInterval
func (mr *MockOptionsMockRecorder) TickMinimumInterval() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TickMinimumInterval", reflect.TypeOf((*MockOptions)(nil).TickMinimumInterval))
}

// SetMaxWiredBlocks mocks base method
func (m *MockOptions) SetMaxWiredBlocks(value uint) Options {
	ret := m.ctrl.Call(m, "SetMaxWiredBlocks", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetMaxWiredBlocks indicates an expected call of SetMaxWiredBlocks
func (mr *MockOptionsMockRecorder) SetMaxWiredBlocks(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxWiredBlocks", reflect.TypeOf((*MockOptions)(nil).SetMaxWiredBlocks), value)
}

// MaxWiredBlocks mocks base method
func (m *MockOptions) MaxWiredBlocks() uint {
	ret := m.ctrl.Call(m, "MaxWiredBlocks")
	ret0, _ := ret[0].(uint)
	return ret0
}

// MaxWiredBlocks indicates an expected call of MaxWiredBlocks
func (mr *MockOptionsMockRecorder) MaxWiredBlocks() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxWiredBlocks", reflect.TypeOf((*MockOptions)(nil).MaxWiredBlocks))
}

// SetClientBootstrapConsistencyLevel mocks base method
func (m *MockOptions) SetClientBootstrapConsistencyLevel(value topology.ReadConsistencyLevel) Options {
	ret := m.ctrl.Call(m, "SetClientBootstrapConsistencyLevel", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetClientBootstrapConsistencyLevel indicates an expected call of SetClientBootstrapConsistencyLevel
func (mr *MockOptionsMockRecorder) SetClientBootstrapConsistencyLevel(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClientBootstrapConsistencyLevel", reflect.TypeOf((*MockOptions)(nil).SetClientBootstrapConsistencyLevel), value)
}

// ClientBootstrapConsistencyLevel mocks base method
func (m *MockOptions) ClientBootstrapConsistencyLevel() topology.ReadConsistencyLevel {
	ret := m.ctrl.Call(m, "ClientBootstrapConsistencyLevel")
	ret0, _ := ret[0].(topology.ReadConsistencyLevel)
	return ret0
}

// ClientBootstrapConsistencyLevel indicates an expected call of ClientBootstrapConsistencyLevel
func (mr *MockOptionsMockRecorder) ClientBootstrapConsistencyLevel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientBootstrapConsistencyLevel", reflect.TypeOf((*MockOptions)(nil).ClientBootstrapConsistencyLevel))
}

// SetClientReadConsistencyLevel mocks base method
func (m *MockOptions) SetClientReadConsistencyLevel(value topology.ReadConsistencyLevel) Options {
	ret := m.ctrl.Call(m, "SetClientReadConsistencyLevel", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetClientReadConsistencyLevel indicates an expected call of SetClientReadConsistencyLevel
func (mr *MockOptionsMockRecorder) SetClientReadConsistencyLevel(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClientReadConsistencyLevel", reflect.TypeOf((*MockOptions)(nil).SetClientReadConsistencyLevel), value)
}

// ClientReadConsistencyLevel mocks base method
func (m *MockOptions) ClientReadConsistencyLevel() topology.ReadConsistencyLevel {
	ret := m.ctrl.Call(m, "ClientReadConsistencyLevel")
	ret0, _ := ret[0].(topology.ReadConsistencyLevel)
	return ret0
}

// ClientReadConsistencyLevel indicates an expected call of ClientReadConsistencyLevel
func (mr *MockOptionsMockRecorder) ClientReadConsistencyLevel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientReadConsistencyLevel", reflect.TypeOf((*MockOptions)(nil).ClientReadConsistencyLevel))
}

// SetClientWriteConsistencyLevel mocks base method
func (m *MockOptions) SetClientWriteConsistencyLevel(value topology.ConsistencyLevel) Options {
	ret := m.ctrl.Call(m, "SetClientWriteConsistencyLevel", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetClientWriteConsistencyLevel indicates an expected call of SetClientWriteConsistencyLevel
func (mr *MockOptionsMockRecorder) SetClientWriteConsistencyLevel(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClientWriteConsistencyLevel", reflect.TypeOf((*MockOptions)(nil).SetClientWriteConsistencyLevel), value)
}

// ClientWriteConsistencyLevel mocks base method
func (m *MockOptions) ClientWriteConsistencyLevel() topology.ConsistencyLevel {
	ret := m.ctrl.Call(m, "ClientWriteConsistencyLevel")
	ret0, _ := ret[0].(topology.ConsistencyLevel)
	return ret0
}

// ClientWriteConsistencyLevel indicates an expected call of ClientWriteConsistencyLevel
func (mr *MockOptionsMockRecorder) ClientWriteConsistencyLevel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientWriteConsistencyLevel", reflect.TypeOf((*MockOptions)(nil).ClientWriteConsistencyLevel))
}

// MockOptionsManager is a mock of OptionsManager interface
type MockOptionsManager struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsManagerMockRecorder
}

// MockOptionsManagerMockRecorder is the mock recorder for MockOptionsManager
type MockOptionsManagerMockRecorder struct {
	mock *MockOptionsManager
}

// NewMockOptionsManager creates a new mock instance
func NewMockOptionsManager(ctrl *gomock.Controller) *MockOptionsManager {
	mock := &MockOptionsManager{ctrl: ctrl}
	mock.recorder = &MockOptionsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOptionsManager) EXPECT() *MockOptionsManagerMockRecorder {
	return m.recorder
}

// Update mocks base method
func (m *MockOptionsManager) Update(value Options) error {
	ret := m.ctrl.Call(m, "Update", value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockOptionsManagerMockRecorder) Update(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOptionsManager)(nil).Update), value)
}

// Get mocks base method
func (m *MockOptionsManager) Get() Options {
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(Options)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockOptionsManagerMockRecorder) Get() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOptionsManager)(nil).Get))
}

// RegisterListener mocks base method
func (m *MockOptionsManager) RegisterListener(l OptionsListener) close.SimpleCloser {
	ret := m.ctrl.Call(m, "RegisterListener", l)
	ret0, _ := ret[0].(close.SimpleCloser)
	return ret0
}

// RegisterListener indicates an expected call of RegisterListener
func (mr *MockOptionsManagerMockRecorder) RegisterListener(l interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterListener", reflect.TypeOf((*MockOptionsManager)(nil).RegisterListener), l)
}

// Close mocks base method
func (m *MockOptionsManager) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockOptionsManagerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockOptionsManager)(nil).Close))
}

// MockOptionsListener is a mock of OptionsListener interface
type MockOptionsListener struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsListenerMockRecorder
}

// MockOptionsListenerMockRecorder is the mock recorder for MockOptionsListener
type MockOptionsListenerMockRecorder struct {
	mock *MockOptionsListener
}

// NewMockOptionsListener creates a new mock instance
func NewMockOptionsListener(ctrl *gomock.Controller) *MockOptionsListener {
	mock := &MockOptionsListener{ctrl: ctrl}
	mock.recorder = &MockOptionsListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOptionsListener) EXPECT() *MockOptionsListenerMockRecorder {
	return m.recorder
}

// SetRuntimeOptions mocks base method
func (m *MockOptionsListener) SetRuntimeOptions(value Options) {
	m.ctrl.Call(m, "SetRuntimeOptions", value)
}

// SetRuntimeOptions indicates an expected call of SetRuntimeOptions
func (mr *MockOptionsListenerMockRecorder) SetRuntimeOptions(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRuntimeOptions", reflect.TypeOf((*MockOptionsListener)(nil).SetRuntimeOptions), value)
}
