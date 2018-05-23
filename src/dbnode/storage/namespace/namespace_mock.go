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
// Source: github.com/m3db/m3db/src/dbnode/storage/namespace/types.go

// Package namespace is a generated GoMock package.
package namespace

import (
	"reflect"
	"time"

	"github.com/m3db/m3cluster/client"
	"github.com/m3db/m3db/src/dbnode/retention"
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

// Equal mocks base method
func (m *MockOptions) Equal(value Options) bool {
	ret := m.ctrl.Call(m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockOptionsMockRecorder) Equal(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockOptions)(nil).Equal), value)
}

// SetBootstrapEnabled mocks base method
func (m *MockOptions) SetBootstrapEnabled(value bool) Options {
	ret := m.ctrl.Call(m, "SetBootstrapEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBootstrapEnabled indicates an expected call of SetBootstrapEnabled
func (mr *MockOptionsMockRecorder) SetBootstrapEnabled(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBootstrapEnabled", reflect.TypeOf((*MockOptions)(nil).SetBootstrapEnabled), value)
}

// BootstrapEnabled mocks base method
func (m *MockOptions) BootstrapEnabled() bool {
	ret := m.ctrl.Call(m, "BootstrapEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// BootstrapEnabled indicates an expected call of BootstrapEnabled
func (mr *MockOptionsMockRecorder) BootstrapEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapEnabled", reflect.TypeOf((*MockOptions)(nil).BootstrapEnabled))
}

// SetFlushEnabled mocks base method
func (m *MockOptions) SetFlushEnabled(value bool) Options {
	ret := m.ctrl.Call(m, "SetFlushEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetFlushEnabled indicates an expected call of SetFlushEnabled
func (mr *MockOptionsMockRecorder) SetFlushEnabled(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFlushEnabled", reflect.TypeOf((*MockOptions)(nil).SetFlushEnabled), value)
}

// FlushEnabled mocks base method
func (m *MockOptions) FlushEnabled() bool {
	ret := m.ctrl.Call(m, "FlushEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// FlushEnabled indicates an expected call of FlushEnabled
func (mr *MockOptionsMockRecorder) FlushEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushEnabled", reflect.TypeOf((*MockOptions)(nil).FlushEnabled))
}

// SetSnapshotEnabled mocks base method
func (m *MockOptions) SetSnapshotEnabled(value bool) Options {
	ret := m.ctrl.Call(m, "SetSnapshotEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetSnapshotEnabled indicates an expected call of SetSnapshotEnabled
func (mr *MockOptionsMockRecorder) SetSnapshotEnabled(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSnapshotEnabled", reflect.TypeOf((*MockOptions)(nil).SetSnapshotEnabled), value)
}

// SnapshotEnabled mocks base method
func (m *MockOptions) SnapshotEnabled() bool {
	ret := m.ctrl.Call(m, "SnapshotEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SnapshotEnabled indicates an expected call of SnapshotEnabled
func (mr *MockOptionsMockRecorder) SnapshotEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotEnabled", reflect.TypeOf((*MockOptions)(nil).SnapshotEnabled))
}

// SetWritesToCommitLog mocks base method
func (m *MockOptions) SetWritesToCommitLog(value bool) Options {
	ret := m.ctrl.Call(m, "SetWritesToCommitLog", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetWritesToCommitLog indicates an expected call of SetWritesToCommitLog
func (mr *MockOptionsMockRecorder) SetWritesToCommitLog(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWritesToCommitLog", reflect.TypeOf((*MockOptions)(nil).SetWritesToCommitLog), value)
}

// WritesToCommitLog mocks base method
func (m *MockOptions) WritesToCommitLog() bool {
	ret := m.ctrl.Call(m, "WritesToCommitLog")
	ret0, _ := ret[0].(bool)
	return ret0
}

// WritesToCommitLog indicates an expected call of WritesToCommitLog
func (mr *MockOptionsMockRecorder) WritesToCommitLog() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritesToCommitLog", reflect.TypeOf((*MockOptions)(nil).WritesToCommitLog))
}

// SetCleanupEnabled mocks base method
func (m *MockOptions) SetCleanupEnabled(value bool) Options {
	ret := m.ctrl.Call(m, "SetCleanupEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetCleanupEnabled indicates an expected call of SetCleanupEnabled
func (mr *MockOptionsMockRecorder) SetCleanupEnabled(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCleanupEnabled", reflect.TypeOf((*MockOptions)(nil).SetCleanupEnabled), value)
}

// CleanupEnabled mocks base method
func (m *MockOptions) CleanupEnabled() bool {
	ret := m.ctrl.Call(m, "CleanupEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CleanupEnabled indicates an expected call of CleanupEnabled
func (mr *MockOptionsMockRecorder) CleanupEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupEnabled", reflect.TypeOf((*MockOptions)(nil).CleanupEnabled))
}

// SetRepairEnabled mocks base method
func (m *MockOptions) SetRepairEnabled(value bool) Options {
	ret := m.ctrl.Call(m, "SetRepairEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetRepairEnabled indicates an expected call of SetRepairEnabled
func (mr *MockOptionsMockRecorder) SetRepairEnabled(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRepairEnabled", reflect.TypeOf((*MockOptions)(nil).SetRepairEnabled), value)
}

// RepairEnabled mocks base method
func (m *MockOptions) RepairEnabled() bool {
	ret := m.ctrl.Call(m, "RepairEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// RepairEnabled indicates an expected call of RepairEnabled
func (mr *MockOptionsMockRecorder) RepairEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepairEnabled", reflect.TypeOf((*MockOptions)(nil).RepairEnabled))
}

// SetRetentionOptions mocks base method
func (m *MockOptions) SetRetentionOptions(value retention.Options) Options {
	ret := m.ctrl.Call(m, "SetRetentionOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetRetentionOptions indicates an expected call of SetRetentionOptions
func (mr *MockOptionsMockRecorder) SetRetentionOptions(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRetentionOptions", reflect.TypeOf((*MockOptions)(nil).SetRetentionOptions), value)
}

// RetentionOptions mocks base method
func (m *MockOptions) RetentionOptions() retention.Options {
	ret := m.ctrl.Call(m, "RetentionOptions")
	ret0, _ := ret[0].(retention.Options)
	return ret0
}

// RetentionOptions indicates an expected call of RetentionOptions
func (mr *MockOptionsMockRecorder) RetentionOptions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetentionOptions", reflect.TypeOf((*MockOptions)(nil).RetentionOptions))
}

// SetIndexOptions mocks base method
func (m *MockOptions) SetIndexOptions(value IndexOptions) Options {
	ret := m.ctrl.Call(m, "SetIndexOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetIndexOptions indicates an expected call of SetIndexOptions
func (mr *MockOptionsMockRecorder) SetIndexOptions(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIndexOptions", reflect.TypeOf((*MockOptions)(nil).SetIndexOptions), value)
}

// IndexOptions mocks base method
func (m *MockOptions) IndexOptions() IndexOptions {
	ret := m.ctrl.Call(m, "IndexOptions")
	ret0, _ := ret[0].(IndexOptions)
	return ret0
}

// IndexOptions indicates an expected call of IndexOptions
func (mr *MockOptionsMockRecorder) IndexOptions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexOptions", reflect.TypeOf((*MockOptions)(nil).IndexOptions))
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
func (m *MockIndexOptions) EXPECT() *MockIndexOptionsMockRecorder {
	return m.recorder
}

// Equal mocks base method
func (m *MockIndexOptions) Equal(value IndexOptions) bool {
	ret := m.ctrl.Call(m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockIndexOptionsMockRecorder) Equal(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockIndexOptions)(nil).Equal), value)
}

// SetEnabled mocks base method
func (m *MockIndexOptions) SetEnabled(value bool) IndexOptions {
	ret := m.ctrl.Call(m, "SetEnabled", value)
	ret0, _ := ret[0].(IndexOptions)
	return ret0
}

// SetEnabled indicates an expected call of SetEnabled
func (mr *MockIndexOptionsMockRecorder) SetEnabled(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEnabled", reflect.TypeOf((*MockIndexOptions)(nil).SetEnabled), value)
}

// Enabled mocks base method
func (m *MockIndexOptions) Enabled() bool {
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled
func (mr *MockIndexOptionsMockRecorder) Enabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockIndexOptions)(nil).Enabled))
}

// SetBlockSize mocks base method
func (m *MockIndexOptions) SetBlockSize(value time.Duration) IndexOptions {
	ret := m.ctrl.Call(m, "SetBlockSize", value)
	ret0, _ := ret[0].(IndexOptions)
	return ret0
}

// SetBlockSize indicates an expected call of SetBlockSize
func (mr *MockIndexOptionsMockRecorder) SetBlockSize(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBlockSize", reflect.TypeOf((*MockIndexOptions)(nil).SetBlockSize), value)
}

// BlockSize mocks base method
func (m *MockIndexOptions) BlockSize() time.Duration {
	ret := m.ctrl.Call(m, "BlockSize")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BlockSize indicates an expected call of BlockSize
func (mr *MockIndexOptionsMockRecorder) BlockSize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockSize", reflect.TypeOf((*MockIndexOptions)(nil).BlockSize))
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
func (m *MockMetadata) EXPECT() *MockMetadataMockRecorder {
	return m.recorder
}

// Equal mocks base method
func (m *MockMetadata) Equal(value Metadata) bool {
	ret := m.ctrl.Call(m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockMetadataMockRecorder) Equal(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockMetadata)(nil).Equal), value)
}

// ID mocks base method
func (m *MockMetadata) ID() ident.ID {
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(ident.ID)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockMetadataMockRecorder) ID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockMetadata)(nil).ID))
}

// Options mocks base method
func (m *MockMetadata) Options() Options {
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(Options)
	return ret0
}

// Options indicates an expected call of Options
func (mr *MockMetadataMockRecorder) Options() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockMetadata)(nil).Options))
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
func (m *MockMap) EXPECT() *MockMapMockRecorder {
	return m.recorder
}

// Equal mocks base method
func (m *MockMap) Equal(value Map) bool {
	ret := m.ctrl.Call(m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockMapMockRecorder) Equal(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockMap)(nil).Equal), value)
}

// Get mocks base method
func (m *MockMap) Get(arg0 ident.ID) (Metadata, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMapMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMap)(nil).Get), arg0)
}

// IDs mocks base method
func (m *MockMap) IDs() []ident.ID {
	ret := m.ctrl.Call(m, "IDs")
	ret0, _ := ret[0].([]ident.ID)
	return ret0
}

// IDs indicates an expected call of IDs
func (mr *MockMapMockRecorder) IDs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDs", reflect.TypeOf((*MockMap)(nil).IDs))
}

// Metadatas mocks base method
func (m *MockMap) Metadatas() []Metadata {
	ret := m.ctrl.Call(m, "Metadatas")
	ret0, _ := ret[0].([]Metadata)
	return ret0
}

// Metadatas indicates an expected call of Metadatas
func (mr *MockMapMockRecorder) Metadatas() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadatas", reflect.TypeOf((*MockMap)(nil).Metadatas))
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
func (m *MockWatch) EXPECT() *MockWatchMockRecorder {
	return m.recorder
}

// C mocks base method
func (m *MockWatch) C() <-chan struct{} {
	ret := m.ctrl.Call(m, "C")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// C indicates an expected call of C
func (mr *MockWatchMockRecorder) C() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockWatch)(nil).C))
}

// Get mocks base method
func (m *MockWatch) Get() Map {
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(Map)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockWatchMockRecorder) Get() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockWatch)(nil).Get))
}

// Close mocks base method
func (m *MockWatch) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockWatchMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWatch)(nil).Close))
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
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockRegistry) Watch() (Watch, error) {
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(Watch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockRegistryMockRecorder) Watch() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockRegistry)(nil).Watch))
}

// Close mocks base method
func (m *MockRegistry) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockRegistryMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRegistry)(nil).Close))
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
func (m *MockInitializer) EXPECT() *MockInitializerMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockInitializer) Init() (Registry, error) {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(Registry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init
func (mr *MockInitializerMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockInitializer)(nil).Init))
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
func (m *MockDynamicOptions) EXPECT() *MockDynamicOptionsMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockDynamicOptions) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockDynamicOptionsMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockDynamicOptions)(nil).Validate))
}

// SetInstrumentOptions mocks base method
func (m *MockDynamicOptions) SetInstrumentOptions(value instrument.Options) DynamicOptions {
	ret := m.ctrl.Call(m, "SetInstrumentOptions", value)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetInstrumentOptions indicates an expected call of SetInstrumentOptions
func (mr *MockDynamicOptionsMockRecorder) SetInstrumentOptions(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInstrumentOptions", reflect.TypeOf((*MockDynamicOptions)(nil).SetInstrumentOptions), value)
}

// InstrumentOptions mocks base method
func (m *MockDynamicOptions) InstrumentOptions() instrument.Options {
	ret := m.ctrl.Call(m, "InstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

// InstrumentOptions indicates an expected call of InstrumentOptions
func (mr *MockDynamicOptionsMockRecorder) InstrumentOptions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstrumentOptions", reflect.TypeOf((*MockDynamicOptions)(nil).InstrumentOptions))
}

// SetConfigServiceClient mocks base method
func (m *MockDynamicOptions) SetConfigServiceClient(c client.Client) DynamicOptions {
	ret := m.ctrl.Call(m, "SetConfigServiceClient", c)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetConfigServiceClient indicates an expected call of SetConfigServiceClient
func (mr *MockDynamicOptionsMockRecorder) SetConfigServiceClient(c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfigServiceClient", reflect.TypeOf((*MockDynamicOptions)(nil).SetConfigServiceClient), c)
}

// ConfigServiceClient mocks base method
func (m *MockDynamicOptions) ConfigServiceClient() client.Client {
	ret := m.ctrl.Call(m, "ConfigServiceClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// ConfigServiceClient indicates an expected call of ConfigServiceClient
func (mr *MockDynamicOptionsMockRecorder) ConfigServiceClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigServiceClient", reflect.TypeOf((*MockDynamicOptions)(nil).ConfigServiceClient))
}

// SetNamespaceRegistryKey mocks base method
func (m *MockDynamicOptions) SetNamespaceRegistryKey(k string) DynamicOptions {
	ret := m.ctrl.Call(m, "SetNamespaceRegistryKey", k)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetNamespaceRegistryKey indicates an expected call of SetNamespaceRegistryKey
func (mr *MockDynamicOptionsMockRecorder) SetNamespaceRegistryKey(k interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespaceRegistryKey", reflect.TypeOf((*MockDynamicOptions)(nil).SetNamespaceRegistryKey), k)
}

// NamespaceRegistryKey mocks base method
func (m *MockDynamicOptions) NamespaceRegistryKey() string {
	ret := m.ctrl.Call(m, "NamespaceRegistryKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// NamespaceRegistryKey indicates an expected call of NamespaceRegistryKey
func (mr *MockDynamicOptionsMockRecorder) NamespaceRegistryKey() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamespaceRegistryKey", reflect.TypeOf((*MockDynamicOptions)(nil).NamespaceRegistryKey))
}

// SetInitTimeout mocks base method
func (m *MockDynamicOptions) SetInitTimeout(value time.Duration) DynamicOptions {
	ret := m.ctrl.Call(m, "SetInitTimeout", value)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetInitTimeout indicates an expected call of SetInitTimeout
func (mr *MockDynamicOptionsMockRecorder) SetInitTimeout(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInitTimeout", reflect.TypeOf((*MockDynamicOptions)(nil).SetInitTimeout), value)
}

// InitTimeout mocks base method
func (m *MockDynamicOptions) InitTimeout() time.Duration {
	ret := m.ctrl.Call(m, "InitTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// InitTimeout indicates an expected call of InitTimeout
func (mr *MockDynamicOptionsMockRecorder) InitTimeout() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitTimeout", reflect.TypeOf((*MockDynamicOptions)(nil).InitTimeout))
}
