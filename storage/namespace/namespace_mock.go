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
// Source: github.com/m3db/m3db/storage/namespace/types.go

package namespace

import (
	"reflect"
	"time"

	"github.com/m3db/m3cluster/client"
	"github.com/m3db/m3db/retention"
	"github.com/m3db/m3x/ident"
	"github.com/m3db/m3x/instrument"

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
func (_m *MockOptions) EXPECT() *MockOptionsMockRecorder {
	return _m.recorder
}

// Validate mocks base method
func (_m *MockOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (_mr *MockOptionsMockRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Validate", reflect.TypeOf((*MockOptions)(nil).Validate))
}

// Equal mocks base method
func (_m *MockOptions) Equal(value Options) bool {
	ret := _m.ctrl.Call(_m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (_mr *MockOptionsMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Equal", reflect.TypeOf((*MockOptions)(nil).Equal), arg0)
}

// SetBootstrapEnabled mocks base method
func (_m *MockOptions) SetBootstrapEnabled(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetBootstrapEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBootstrapEnabled indicates an expected call of SetBootstrapEnabled
func (_mr *MockOptionsMockRecorder) SetBootstrapEnabled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBootstrapEnabled", reflect.TypeOf((*MockOptions)(nil).SetBootstrapEnabled), arg0)
}

// BootstrapEnabled mocks base method
func (_m *MockOptions) BootstrapEnabled() bool {
	ret := _m.ctrl.Call(_m, "BootstrapEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// BootstrapEnabled indicates an expected call of BootstrapEnabled
func (_mr *MockOptionsMockRecorder) BootstrapEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BootstrapEnabled", reflect.TypeOf((*MockOptions)(nil).BootstrapEnabled))
}

// SetFlushEnabled mocks base method
func (_m *MockOptions) SetFlushEnabled(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetFlushEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetFlushEnabled indicates an expected call of SetFlushEnabled
func (_mr *MockOptionsMockRecorder) SetFlushEnabled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetFlushEnabled", reflect.TypeOf((*MockOptions)(nil).SetFlushEnabled), arg0)
}

// FlushEnabled mocks base method
func (_m *MockOptions) FlushEnabled() bool {
	ret := _m.ctrl.Call(_m, "FlushEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// FlushEnabled indicates an expected call of FlushEnabled
func (_mr *MockOptionsMockRecorder) FlushEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FlushEnabled", reflect.TypeOf((*MockOptions)(nil).FlushEnabled))
}

// SetSnapshotEnabled mocks base method
func (_m *MockOptions) SetSnapshotEnabled(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetSnapshotEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetSnapshotEnabled indicates an expected call of SetSnapshotEnabled
func (_mr *MockOptionsMockRecorder) SetSnapshotEnabled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetSnapshotEnabled", reflect.TypeOf((*MockOptions)(nil).SetSnapshotEnabled), arg0)
}

// SnapshotEnabled mocks base method
func (_m *MockOptions) SnapshotEnabled() bool {
	ret := _m.ctrl.Call(_m, "SnapshotEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SnapshotEnabled indicates an expected call of SnapshotEnabled
func (_mr *MockOptionsMockRecorder) SnapshotEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SnapshotEnabled", reflect.TypeOf((*MockOptions)(nil).SnapshotEnabled))
}

// SetWritesToCommitLog mocks base method
func (_m *MockOptions) SetWritesToCommitLog(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetWritesToCommitLog", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetWritesToCommitLog indicates an expected call of SetWritesToCommitLog
func (_mr *MockOptionsMockRecorder) SetWritesToCommitLog(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetWritesToCommitLog", reflect.TypeOf((*MockOptions)(nil).SetWritesToCommitLog), arg0)
}

// WritesToCommitLog mocks base method
func (_m *MockOptions) WritesToCommitLog() bool {
	ret := _m.ctrl.Call(_m, "WritesToCommitLog")
	ret0, _ := ret[0].(bool)
	return ret0
}

// WritesToCommitLog indicates an expected call of WritesToCommitLog
func (_mr *MockOptionsMockRecorder) WritesToCommitLog() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WritesToCommitLog", reflect.TypeOf((*MockOptions)(nil).WritesToCommitLog))
}

// SetCleanupEnabled mocks base method
func (_m *MockOptions) SetCleanupEnabled(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetCleanupEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetCleanupEnabled indicates an expected call of SetCleanupEnabled
func (_mr *MockOptionsMockRecorder) SetCleanupEnabled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetCleanupEnabled", reflect.TypeOf((*MockOptions)(nil).SetCleanupEnabled), arg0)
}

// CleanupEnabled mocks base method
func (_m *MockOptions) CleanupEnabled() bool {
	ret := _m.ctrl.Call(_m, "CleanupEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CleanupEnabled indicates an expected call of CleanupEnabled
func (_mr *MockOptionsMockRecorder) CleanupEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CleanupEnabled", reflect.TypeOf((*MockOptions)(nil).CleanupEnabled))
}

// SetRepairEnabled mocks base method
func (_m *MockOptions) SetRepairEnabled(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetRepairEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetRepairEnabled indicates an expected call of SetRepairEnabled
func (_mr *MockOptionsMockRecorder) SetRepairEnabled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRepairEnabled", reflect.TypeOf((*MockOptions)(nil).SetRepairEnabled), arg0)
}

// RepairEnabled mocks base method
func (_m *MockOptions) RepairEnabled() bool {
	ret := _m.ctrl.Call(_m, "RepairEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// RepairEnabled indicates an expected call of RepairEnabled
func (_mr *MockOptionsMockRecorder) RepairEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RepairEnabled", reflect.TypeOf((*MockOptions)(nil).RepairEnabled))
}

// SetRetentionOptions mocks base method
func (_m *MockOptions) SetRetentionOptions(value retention.Options) Options {
	ret := _m.ctrl.Call(_m, "SetRetentionOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetRetentionOptions indicates an expected call of SetRetentionOptions
func (_mr *MockOptionsMockRecorder) SetRetentionOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRetentionOptions", reflect.TypeOf((*MockOptions)(nil).SetRetentionOptions), arg0)
}

// RetentionOptions mocks base method
func (_m *MockOptions) RetentionOptions() retention.Options {
	ret := _m.ctrl.Call(_m, "RetentionOptions")
	ret0, _ := ret[0].(retention.Options)
	return ret0
}

// RetentionOptions indicates an expected call of RetentionOptions
func (_mr *MockOptionsMockRecorder) RetentionOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RetentionOptions", reflect.TypeOf((*MockOptions)(nil).RetentionOptions))
}

// SetIndexOptions mocks base method
func (_m *MockOptions) SetIndexOptions(value IndexOptions) Options {
	ret := _m.ctrl.Call(_m, "SetIndexOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetIndexOptions indicates an expected call of SetIndexOptions
func (_mr *MockOptionsMockRecorder) SetIndexOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetIndexOptions", reflect.TypeOf((*MockOptions)(nil).SetIndexOptions), arg0)
}

// IndexOptions mocks base method
func (_m *MockOptions) IndexOptions() IndexOptions {
	ret := _m.ctrl.Call(_m, "IndexOptions")
	ret0, _ := ret[0].(IndexOptions)
	return ret0
}

// IndexOptions indicates an expected call of IndexOptions
func (_mr *MockOptionsMockRecorder) IndexOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "IndexOptions", reflect.TypeOf((*MockOptions)(nil).IndexOptions))
}

// MockIndexOptions is a mock of IndexOptions interface
type MockIndexOptions struct {
	ctrl     *gomock.Controller
	recorder *MockIndexOptionsMockRecorder
}

// MockIndexOptionsMockRecorder is the mock recorder for MockIndexOptions
type MockIndexOptionsMockRecorder struct {
	mock *MockIndexOptions
}

// NewMockIndexOptions creates a new mock instance
func NewMockIndexOptions(ctrl *gomock.Controller) *MockIndexOptions {
	mock := &MockIndexOptions{ctrl: ctrl}
	mock.recorder = &MockIndexOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIndexOptions) EXPECT() *MockIndexOptionsMockRecorder {
	return _m.recorder
}

// Equal mocks base method
func (_m *MockIndexOptions) Equal(value IndexOptions) bool {
	ret := _m.ctrl.Call(_m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (_mr *MockIndexOptionsMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Equal", reflect.TypeOf((*MockIndexOptions)(nil).Equal), arg0)
}

// SetEnabled mocks base method
func (_m *MockIndexOptions) SetEnabled(value bool) IndexOptions {
	ret := _m.ctrl.Call(_m, "SetEnabled", value)
	ret0, _ := ret[0].(IndexOptions)
	return ret0
}

// SetEnabled indicates an expected call of SetEnabled
func (_mr *MockIndexOptionsMockRecorder) SetEnabled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetEnabled", reflect.TypeOf((*MockIndexOptions)(nil).SetEnabled), arg0)
}

// Enabled mocks base method
func (_m *MockIndexOptions) Enabled() bool {
	ret := _m.ctrl.Call(_m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled
func (_mr *MockIndexOptionsMockRecorder) Enabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Enabled", reflect.TypeOf((*MockIndexOptions)(nil).Enabled))
}

// SetBlockSize mocks base method
func (_m *MockIndexOptions) SetBlockSize(value time.Duration) IndexOptions {
	ret := _m.ctrl.Call(_m, "SetBlockSize", value)
	ret0, _ := ret[0].(IndexOptions)
	return ret0
}

// SetBlockSize indicates an expected call of SetBlockSize
func (_mr *MockIndexOptionsMockRecorder) SetBlockSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBlockSize", reflect.TypeOf((*MockIndexOptions)(nil).SetBlockSize), arg0)
}

// BlockSize mocks base method
func (_m *MockIndexOptions) BlockSize() time.Duration {
	ret := _m.ctrl.Call(_m, "BlockSize")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BlockSize indicates an expected call of BlockSize
func (_mr *MockIndexOptionsMockRecorder) BlockSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BlockSize", reflect.TypeOf((*MockIndexOptions)(nil).BlockSize))
}

// MockMetadata is a mock of Metadata interface
type MockMetadata struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataMockRecorder
}

// MockMetadataMockRecorder is the mock recorder for MockMetadata
type MockMetadataMockRecorder struct {
	mock *MockMetadata
}

// NewMockMetadata creates a new mock instance
func NewMockMetadata(ctrl *gomock.Controller) *MockMetadata {
	mock := &MockMetadata{ctrl: ctrl}
	mock.recorder = &MockMetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockMetadata) EXPECT() *MockMetadataMockRecorder {
	return _m.recorder
}

// Equal mocks base method
func (_m *MockMetadata) Equal(value Metadata) bool {
	ret := _m.ctrl.Call(_m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (_mr *MockMetadataMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Equal", reflect.TypeOf((*MockMetadata)(nil).Equal), arg0)
}

// ID mocks base method
func (_m *MockMetadata) ID() ident.ID {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(ident.ID)
	return ret0
}

// ID indicates an expected call of ID
func (_mr *MockMetadataMockRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ID", reflect.TypeOf((*MockMetadata)(nil).ID))
}

// Options mocks base method
func (_m *MockMetadata) Options() Options {
	ret := _m.ctrl.Call(_m, "Options")
	ret0, _ := ret[0].(Options)
	return ret0
}

// Options indicates an expected call of Options
func (_mr *MockMetadataMockRecorder) Options() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Options", reflect.TypeOf((*MockMetadata)(nil).Options))
}

// MockMap is a mock of Map interface
type MockMap struct {
	ctrl     *gomock.Controller
	recorder *MockMapMockRecorder
}

// MockMapMockRecorder is the mock recorder for MockMap
type MockMapMockRecorder struct {
	mock *MockMap
}

// NewMockMap creates a new mock instance
func NewMockMap(ctrl *gomock.Controller) *MockMap {
	mock := &MockMap{ctrl: ctrl}
	mock.recorder = &MockMapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockMap) EXPECT() *MockMapMockRecorder {
	return _m.recorder
}

// Equal mocks base method
func (_m *MockMap) Equal(value Map) bool {
	ret := _m.ctrl.Call(_m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (_mr *MockMapMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Equal", reflect.TypeOf((*MockMap)(nil).Equal), arg0)
}

// Get mocks base method
func (_m *MockMap) Get(_param0 ident.ID) (Metadata, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (_mr *MockMapMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockMap)(nil).Get), arg0)
}

// IDs mocks base method
func (_m *MockMap) IDs() []ident.ID {
	ret := _m.ctrl.Call(_m, "IDs")
	ret0, _ := ret[0].([]ident.ID)
	return ret0
}

// IDs indicates an expected call of IDs
func (_mr *MockMapMockRecorder) IDs() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "IDs", reflect.TypeOf((*MockMap)(nil).IDs))
}

// Metadatas mocks base method
func (_m *MockMap) Metadatas() []Metadata {
	ret := _m.ctrl.Call(_m, "Metadatas")
	ret0, _ := ret[0].([]Metadata)
	return ret0
}

// Metadatas indicates an expected call of Metadatas
func (_mr *MockMapMockRecorder) Metadatas() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Metadatas", reflect.TypeOf((*MockMap)(nil).Metadatas))
}

// MockWatch is a mock of Watch interface
type MockWatch struct {
	ctrl     *gomock.Controller
	recorder *MockWatchMockRecorder
}

// MockWatchMockRecorder is the mock recorder for MockWatch
type MockWatchMockRecorder struct {
	mock *MockWatch
}

// NewMockWatch creates a new mock instance
func NewMockWatch(ctrl *gomock.Controller) *MockWatch {
	mock := &MockWatch{ctrl: ctrl}
	mock.recorder = &MockWatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockWatch) EXPECT() *MockWatchMockRecorder {
	return _m.recorder
}

// C mocks base method
func (_m *MockWatch) C() <-chan struct{} {
	ret := _m.ctrl.Call(_m, "C")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// C indicates an expected call of C
func (_mr *MockWatchMockRecorder) C() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "C", reflect.TypeOf((*MockWatch)(nil).C))
}

// Get mocks base method
func (_m *MockWatch) Get() Map {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Map)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockWatchMockRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockWatch)(nil).Get))
}

// Close mocks base method
func (_m *MockWatch) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockWatchMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockWatch)(nil).Close))
}

// MockRegistry is a mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return _m.recorder
}

// Watch mocks base method
func (_m *MockRegistry) Watch() (Watch, error) {
	ret := _m.ctrl.Call(_m, "Watch")
	ret0, _ := ret[0].(Watch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (_mr *MockRegistryMockRecorder) Watch() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Watch", reflect.TypeOf((*MockRegistry)(nil).Watch))
}

// Close mocks base method
func (_m *MockRegistry) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockRegistryMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockRegistry)(nil).Close))
}

// MockInitializer is a mock of Initializer interface
type MockInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockInitializerMockRecorder
}

// MockInitializerMockRecorder is the mock recorder for MockInitializer
type MockInitializerMockRecorder struct {
	mock *MockInitializer
}

// NewMockInitializer creates a new mock instance
func NewMockInitializer(ctrl *gomock.Controller) *MockInitializer {
	mock := &MockInitializer{ctrl: ctrl}
	mock.recorder = &MockInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockInitializer) EXPECT() *MockInitializerMockRecorder {
	return _m.recorder
}

// Init mocks base method
func (_m *MockInitializer) Init() (Registry, error) {
	ret := _m.ctrl.Call(_m, "Init")
	ret0, _ := ret[0].(Registry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init
func (_mr *MockInitializerMockRecorder) Init() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Init", reflect.TypeOf((*MockInitializer)(nil).Init))
}

// MockDynamicOptions is a mock of DynamicOptions interface
type MockDynamicOptions struct {
	ctrl     *gomock.Controller
	recorder *MockDynamicOptionsMockRecorder
}

// MockDynamicOptionsMockRecorder is the mock recorder for MockDynamicOptions
type MockDynamicOptionsMockRecorder struct {
	mock *MockDynamicOptions
}

// NewMockDynamicOptions creates a new mock instance
func NewMockDynamicOptions(ctrl *gomock.Controller) *MockDynamicOptions {
	mock := &MockDynamicOptions{ctrl: ctrl}
	mock.recorder = &MockDynamicOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDynamicOptions) EXPECT() *MockDynamicOptionsMockRecorder {
	return _m.recorder
}

// Validate mocks base method
func (_m *MockDynamicOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (_mr *MockDynamicOptionsMockRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Validate", reflect.TypeOf((*MockDynamicOptions)(nil).Validate))
}

// SetInstrumentOptions mocks base method
func (_m *MockDynamicOptions) SetInstrumentOptions(value instrument.Options) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetInstrumentOptions", value)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetInstrumentOptions indicates an expected call of SetInstrumentOptions
func (_mr *MockDynamicOptionsMockRecorder) SetInstrumentOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetInstrumentOptions", reflect.TypeOf((*MockDynamicOptions)(nil).SetInstrumentOptions), arg0)
}

// InstrumentOptions mocks base method
func (_m *MockDynamicOptions) InstrumentOptions() instrument.Options {
	ret := _m.ctrl.Call(_m, "InstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

// InstrumentOptions indicates an expected call of InstrumentOptions
func (_mr *MockDynamicOptionsMockRecorder) InstrumentOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InstrumentOptions", reflect.TypeOf((*MockDynamicOptions)(nil).InstrumentOptions))
}

// SetConfigServiceClient mocks base method
func (_m *MockDynamicOptions) SetConfigServiceClient(c client.Client) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetConfigServiceClient", c)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetConfigServiceClient indicates an expected call of SetConfigServiceClient
func (_mr *MockDynamicOptionsMockRecorder) SetConfigServiceClient(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetConfigServiceClient", reflect.TypeOf((*MockDynamicOptions)(nil).SetConfigServiceClient), arg0)
}

// ConfigServiceClient mocks base method
func (_m *MockDynamicOptions) ConfigServiceClient() client.Client {
	ret := _m.ctrl.Call(_m, "ConfigServiceClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// ConfigServiceClient indicates an expected call of ConfigServiceClient
func (_mr *MockDynamicOptionsMockRecorder) ConfigServiceClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ConfigServiceClient", reflect.TypeOf((*MockDynamicOptions)(nil).ConfigServiceClient))
}

// SetNamespaceRegistryKey mocks base method
func (_m *MockDynamicOptions) SetNamespaceRegistryKey(k string) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetNamespaceRegistryKey", k)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetNamespaceRegistryKey indicates an expected call of SetNamespaceRegistryKey
func (_mr *MockDynamicOptionsMockRecorder) SetNamespaceRegistryKey(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetNamespaceRegistryKey", reflect.TypeOf((*MockDynamicOptions)(nil).SetNamespaceRegistryKey), arg0)
}

// NamespaceRegistryKey mocks base method
func (_m *MockDynamicOptions) NamespaceRegistryKey() string {
	ret := _m.ctrl.Call(_m, "NamespaceRegistryKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// NamespaceRegistryKey indicates an expected call of NamespaceRegistryKey
func (_mr *MockDynamicOptionsMockRecorder) NamespaceRegistryKey() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "NamespaceRegistryKey", reflect.TypeOf((*MockDynamicOptions)(nil).NamespaceRegistryKey))
}

// SetInitTimeout mocks base method
func (_m *MockDynamicOptions) SetInitTimeout(value time.Duration) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetInitTimeout", value)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetInitTimeout indicates an expected call of SetInitTimeout
func (_mr *MockDynamicOptionsMockRecorder) SetInitTimeout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetInitTimeout", reflect.TypeOf((*MockDynamicOptions)(nil).SetInitTimeout), arg0)
}

// InitTimeout mocks base method
func (_m *MockDynamicOptions) InitTimeout() time.Duration {
	ret := _m.ctrl.Call(_m, "InitTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// InitTimeout indicates an expected call of InitTimeout
func (_mr *MockDynamicOptionsMockRecorder) InitTimeout() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InitTimeout", reflect.TypeOf((*MockDynamicOptions)(nil).InitTimeout))
}
