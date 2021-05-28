// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/query/block (interfaces: Block,StepIter,Builder,Step,SeriesIter)

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

// Package block is a generated GoMock package.
package block

import (
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
)

// MockBlock is a mock of Block interface.
type MockBlock struct {
	ctrl     *gomock.Controller
	recorder *MockBlockMockRecorder
}

// MockBlockMockRecorder is the mock recorder for MockBlock.
type MockBlockMockRecorder struct {
	mock *MockBlock
}

// NewMockBlock creates a new mock instance.
func NewMockBlock(ctrl *gomock.Controller) *MockBlock {
	mock := &MockBlock{ctrl: ctrl}
	mock.recorder = &MockBlockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlock) EXPECT() *MockBlockMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockBlock) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockBlockMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBlock)(nil).Close))
}

// Info mocks base method.
func (m *MockBlock) Info() BlockInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(BlockInfo)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockBlockMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockBlock)(nil).Info))
}

// Meta mocks base method.
func (m *MockBlock) Meta() Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(Metadata)
	return ret0
}

// Meta indicates an expected call of Meta.
func (mr *MockBlockMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockBlock)(nil).Meta))
}

// MultiSeriesIter mocks base method.
func (m *MockBlock) MultiSeriesIter(arg0 int) ([]SeriesIterBatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiSeriesIter", arg0)
	ret0, _ := ret[0].([]SeriesIterBatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiSeriesIter indicates an expected call of MultiSeriesIter.
func (mr *MockBlockMockRecorder) MultiSeriesIter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiSeriesIter", reflect.TypeOf((*MockBlock)(nil).MultiSeriesIter), arg0)
}

// SeriesIter mocks base method.
func (m *MockBlock) SeriesIter() (SeriesIter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesIter")
	ret0, _ := ret[0].(SeriesIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeriesIter indicates an expected call of SeriesIter.
func (mr *MockBlockMockRecorder) SeriesIter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesIter", reflect.TypeOf((*MockBlock)(nil).SeriesIter))
}

// StepIter mocks base method.
func (m *MockBlock) StepIter() (StepIter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StepIter")
	ret0, _ := ret[0].(StepIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StepIter indicates an expected call of StepIter.
func (mr *MockBlockMockRecorder) StepIter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StepIter", reflect.TypeOf((*MockBlock)(nil).StepIter))
}

// MockStepIter is a mock of StepIter interface.
type MockStepIter struct {
	ctrl     *gomock.Controller
	recorder *MockStepIterMockRecorder
}

// MockStepIterMockRecorder is the mock recorder for MockStepIter.
type MockStepIterMockRecorder struct {
	mock *MockStepIter
}

// NewMockStepIter creates a new mock instance.
func NewMockStepIter(ctrl *gomock.Controller) *MockStepIter {
	mock := &MockStepIter{ctrl: ctrl}
	mock.recorder = &MockStepIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStepIter) EXPECT() *MockStepIterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStepIter) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockStepIterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStepIter)(nil).Close))
}

// Current mocks base method.
func (m *MockStepIter) Current() Step {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(Step)
	return ret0
}

// Current indicates an expected call of Current.
func (mr *MockStepIterMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockStepIter)(nil).Current))
}

// Err mocks base method.
func (m *MockStepIter) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockStepIterMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockStepIter)(nil).Err))
}

// Next mocks base method.
func (m *MockStepIter) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockStepIterMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockStepIter)(nil).Next))
}

// SeriesMeta mocks base method.
func (m *MockStepIter) SeriesMeta() []SeriesMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesMeta")
	ret0, _ := ret[0].([]SeriesMeta)
	return ret0
}

// SeriesMeta indicates an expected call of SeriesMeta.
func (mr *MockStepIterMockRecorder) SeriesMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesMeta", reflect.TypeOf((*MockStepIter)(nil).SeriesMeta))
}

// StepCount mocks base method.
func (m *MockStepIter) StepCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StepCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// StepCount indicates an expected call of StepCount.
func (mr *MockStepIterMockRecorder) StepCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StepCount", reflect.TypeOf((*MockStepIter)(nil).StepCount))
}

// MockBuilder is a mock of Builder interface.
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder.
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance.
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// AddCols mocks base method.
func (m *MockBuilder) AddCols(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCols", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCols indicates an expected call of AddCols.
func (mr *MockBuilderMockRecorder) AddCols(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCols", reflect.TypeOf((*MockBuilder)(nil).AddCols), arg0)
}

// AppendValue mocks base method.
func (m *MockBuilder) AppendValue(arg0 int, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendValue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendValue indicates an expected call of AppendValue.
func (mr *MockBuilderMockRecorder) AppendValue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendValue", reflect.TypeOf((*MockBuilder)(nil).AppendValue), arg0, arg1)
}

// AppendValues mocks base method.
func (m *MockBuilder) AppendValues(arg0 int, arg1 []float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendValues", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendValues indicates an expected call of AppendValues.
func (mr *MockBuilderMockRecorder) AppendValues(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendValues", reflect.TypeOf((*MockBuilder)(nil).AppendValues), arg0, arg1)
}

// Build mocks base method.
func (m *MockBuilder) Build() Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(Block)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockBuilderMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBuilder)(nil).Build))
}

// BuildAsType mocks base method.
func (m *MockBuilder) BuildAsType(arg0 BlockType) Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildAsType", arg0)
	ret0, _ := ret[0].(Block)
	return ret0
}

// BuildAsType indicates an expected call of BuildAsType.
func (mr *MockBuilderMockRecorder) BuildAsType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildAsType", reflect.TypeOf((*MockBuilder)(nil).BuildAsType), arg0)
}

// PopulateColumns mocks base method.
func (m *MockBuilder) PopulateColumns(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PopulateColumns", arg0)
}

// PopulateColumns indicates an expected call of PopulateColumns.
func (mr *MockBuilderMockRecorder) PopulateColumns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopulateColumns", reflect.TypeOf((*MockBuilder)(nil).PopulateColumns), arg0)
}

// SetRow mocks base method.
func (m *MockBuilder) SetRow(arg0 int, arg1 []float64, arg2 SeriesMeta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRow", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRow indicates an expected call of SetRow.
func (mr *MockBuilderMockRecorder) SetRow(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRow", reflect.TypeOf((*MockBuilder)(nil).SetRow), arg0, arg1, arg2)
}

// MockStep is a mock of Step interface.
type MockStep struct {
	ctrl     *gomock.Controller
	recorder *MockStepMockRecorder
}

// MockStepMockRecorder is the mock recorder for MockStep.
type MockStepMockRecorder struct {
	mock *MockStep
}

// NewMockStep creates a new mock instance.
func NewMockStep(ctrl *gomock.Controller) *MockStep {
	mock := &MockStep{ctrl: ctrl}
	mock.recorder = &MockStepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStep) EXPECT() *MockStepMockRecorder {
	return m.recorder
}

// Time mocks base method.
func (m *MockStep) Time() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Time indicates an expected call of Time.
func (mr *MockStepMockRecorder) Time() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockStep)(nil).Time))
}

// Values mocks base method.
func (m *MockStep) Values() []float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].([]float64)
	return ret0
}

// Values indicates an expected call of Values.
func (mr *MockStepMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockStep)(nil).Values))
}

// MockSeriesIter is a mock of SeriesIter interface.
type MockSeriesIter struct {
	ctrl     *gomock.Controller
	recorder *MockSeriesIterMockRecorder
}

// MockSeriesIterMockRecorder is the mock recorder for MockSeriesIter.
type MockSeriesIterMockRecorder struct {
	mock *MockSeriesIter
}

// NewMockSeriesIter creates a new mock instance.
func NewMockSeriesIter(ctrl *gomock.Controller) *MockSeriesIter {
	mock := &MockSeriesIter{ctrl: ctrl}
	mock.recorder = &MockSeriesIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeriesIter) EXPECT() *MockSeriesIterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSeriesIter) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSeriesIterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSeriesIter)(nil).Close))
}

// Current mocks base method.
func (m *MockSeriesIter) Current() UnconsolidatedSeries {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(UnconsolidatedSeries)
	return ret0
}

// Current indicates an expected call of Current.
func (mr *MockSeriesIterMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockSeriesIter)(nil).Current))
}

// Err mocks base method.
func (m *MockSeriesIter) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockSeriesIterMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSeriesIter)(nil).Err))
}

// Next mocks base method.
func (m *MockSeriesIter) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockSeriesIterMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockSeriesIter)(nil).Next))
}

// SeriesCount mocks base method.
func (m *MockSeriesIter) SeriesCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// SeriesCount indicates an expected call of SeriesCount.
func (mr *MockSeriesIterMockRecorder) SeriesCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesCount", reflect.TypeOf((*MockSeriesIter)(nil).SeriesCount))
}

// SeriesMeta mocks base method.
func (m *MockSeriesIter) SeriesMeta() []SeriesMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesMeta")
	ret0, _ := ret[0].([]SeriesMeta)
	return ret0
}

// SeriesMeta indicates an expected call of SeriesMeta.
func (mr *MockSeriesIterMockRecorder) SeriesMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesMeta", reflect.TypeOf((*MockSeriesIter)(nil).SeriesMeta))
}
