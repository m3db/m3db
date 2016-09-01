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
// Source: github.com/m3db/m3db/encoding/types.go

package encoding

import (
	gomock "github.com/golang/mock/gomock"
	pool "github.com/m3db/m3db/pool"
	ts "github.com/m3db/m3db/ts"
	io0 "github.com/m3db/m3db/x/io"
	time0 "github.com/m3db/m3x/time"
	io "io"
	time "time"
)

// Mock of Encoder interface
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *_MockEncoderRecorder
}

// Recorder for MockEncoder (not exported)
type _MockEncoderRecorder struct {
	mock *MockEncoder
}

func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &_MockEncoderRecorder{mock}
	return mock
}

func (_m *MockEncoder) EXPECT() *_MockEncoderRecorder {
	return _m.recorder
}

func (_m *MockEncoder) Encode(dp ts.Datapoint, timeUnit time0.Unit, annotation ts.Annotation) error {
	ret := _m.ctrl.Call(_m, "Encode", dp, timeUnit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockEncoderRecorder) Encode(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Encode", arg0, arg1, arg2)
}

func (_m *MockEncoder) Stream() io0.SegmentReader {
	ret := _m.ctrl.Call(_m, "Stream")
	ret0, _ := ret[0].(io0.SegmentReader)
	return ret0
}

func (_mr *_MockEncoderRecorder) Stream() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stream")
}

func (_m *MockEncoder) Seal() {
	_m.ctrl.Call(_m, "Seal")
}

func (_mr *_MockEncoderRecorder) Seal() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Seal")
}

func (_m *MockEncoder) Unseal() error {
	ret := _m.ctrl.Call(_m, "Unseal")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockEncoderRecorder) Unseal() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unseal")
}

func (_m *MockEncoder) Reset(t time.Time, capacity int) {
	_m.ctrl.Call(_m, "Reset", t, capacity)
}

func (_mr *_MockEncoderRecorder) Reset(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0, arg1)
}

func (_m *MockEncoder) ResetSetData(t time.Time, data []byte, writable bool) {
	_m.ctrl.Call(_m, "ResetSetData", t, data, writable)
}

func (_mr *_MockEncoderRecorder) ResetSetData(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResetSetData", arg0, arg1, arg2)
}

func (_m *MockEncoder) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockEncoderRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
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

func (_m *MockOptions) DefaultTimeUnit(tu time0.Unit) Options {
	ret := _m.ctrl.Call(_m, "DefaultTimeUnit", tu)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) DefaultTimeUnit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DefaultTimeUnit", arg0)
}

func (_m *MockOptions) GetDefaultTimeUnit() time0.Unit {
	ret := _m.ctrl.Call(_m, "GetDefaultTimeUnit")
	ret0, _ := ret[0].(time0.Unit)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetDefaultTimeUnit() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDefaultTimeUnit")
}

func (_m *MockOptions) TimeEncodingSchemes(value TimeEncodingSchemes) Options {
	ret := _m.ctrl.Call(_m, "TimeEncodingSchemes", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) TimeEncodingSchemes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TimeEncodingSchemes", arg0)
}

func (_m *MockOptions) GetTimeEncodingSchemes() TimeEncodingSchemes {
	ret := _m.ctrl.Call(_m, "GetTimeEncodingSchemes")
	ret0, _ := ret[0].(TimeEncodingSchemes)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetTimeEncodingSchemes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTimeEncodingSchemes")
}

func (_m *MockOptions) MarkerEncodingScheme(value MarkerEncodingScheme) Options {
	ret := _m.ctrl.Call(_m, "MarkerEncodingScheme", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) MarkerEncodingScheme(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MarkerEncodingScheme", arg0)
}

func (_m *MockOptions) GetMarkerEncodingScheme() MarkerEncodingScheme {
	ret := _m.ctrl.Call(_m, "GetMarkerEncodingScheme")
	ret0, _ := ret[0].(MarkerEncodingScheme)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetMarkerEncodingScheme() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMarkerEncodingScheme")
}

func (_m *MockOptions) EncoderPool(value EncoderPool) Options {
	ret := _m.ctrl.Call(_m, "EncoderPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) EncoderPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EncoderPool", arg0)
}

func (_m *MockOptions) GetEncoderPool() EncoderPool {
	ret := _m.ctrl.Call(_m, "GetEncoderPool")
	ret0, _ := ret[0].(EncoderPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetEncoderPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEncoderPool")
}

func (_m *MockOptions) ReaderIteratorPool(value ReaderIteratorPool) Options {
	ret := _m.ctrl.Call(_m, "ReaderIteratorPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ReaderIteratorPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReaderIteratorPool", arg0)
}

func (_m *MockOptions) GetReaderIteratorPool() ReaderIteratorPool {
	ret := _m.ctrl.Call(_m, "GetReaderIteratorPool")
	ret0, _ := ret[0].(ReaderIteratorPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetReaderIteratorPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetReaderIteratorPool")
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

func (_m *MockOptions) SegmentReaderPool(value io0.SegmentReaderPool) Options {
	ret := _m.ctrl.Call(_m, "SegmentReaderPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SegmentReaderPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SegmentReaderPool", arg0)
}

func (_m *MockOptions) GetSegmentReaderPool() io0.SegmentReaderPool {
	ret := _m.ctrl.Call(_m, "GetSegmentReaderPool")
	ret0, _ := ret[0].(io0.SegmentReaderPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) GetSegmentReaderPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSegmentReaderPool")
}

// Mock of Iterator interface
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *_MockIteratorRecorder
}

// Recorder for MockIterator (not exported)
type _MockIteratorRecorder struct {
	mock *MockIterator
}

func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &_MockIteratorRecorder{mock}
	return mock
}

func (_m *MockIterator) EXPECT() *_MockIteratorRecorder {
	return _m.recorder
}

func (_m *MockIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockIteratorRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

func (_m *MockIterator) Current() (ts.Datapoint, time0.Unit, ts.Annotation) {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(ts.Datapoint)
	ret1, _ := ret[1].(time0.Unit)
	ret2, _ := ret[2].(ts.Annotation)
	return ret0, ret1, ret2
}

func (_mr *_MockIteratorRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

func (_m *MockIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIteratorRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockIterator) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockIteratorRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of ReaderIterator interface
type MockReaderIterator struct {
	ctrl     *gomock.Controller
	recorder *_MockReaderIteratorRecorder
}

// Recorder for MockReaderIterator (not exported)
type _MockReaderIteratorRecorder struct {
	mock *MockReaderIterator
}

func NewMockReaderIterator(ctrl *gomock.Controller) *MockReaderIterator {
	mock := &MockReaderIterator{ctrl: ctrl}
	mock.recorder = &_MockReaderIteratorRecorder{mock}
	return mock
}

func (_m *MockReaderIterator) EXPECT() *_MockReaderIteratorRecorder {
	return _m.recorder
}

func (_m *MockReaderIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockReaderIteratorRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

func (_m *MockReaderIterator) Current() (ts.Datapoint, time0.Unit, ts.Annotation) {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(ts.Datapoint)
	ret1, _ := ret[1].(time0.Unit)
	ret2, _ := ret[2].(ts.Annotation)
	return ret0, ret1, ret2
}

func (_mr *_MockReaderIteratorRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

func (_m *MockReaderIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockReaderIteratorRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockReaderIterator) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockReaderIteratorRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockReaderIterator) Reset(reader io.Reader) {
	_m.ctrl.Call(_m, "Reset", reader)
}

func (_mr *_MockReaderIteratorRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

// Mock of MultiReaderIterator interface
type MockMultiReaderIterator struct {
	ctrl     *gomock.Controller
	recorder *_MockMultiReaderIteratorRecorder
}

// Recorder for MockMultiReaderIterator (not exported)
type _MockMultiReaderIteratorRecorder struct {
	mock *MockMultiReaderIterator
}

func NewMockMultiReaderIterator(ctrl *gomock.Controller) *MockMultiReaderIterator {
	mock := &MockMultiReaderIterator{ctrl: ctrl}
	mock.recorder = &_MockMultiReaderIteratorRecorder{mock}
	return mock
}

func (_m *MockMultiReaderIterator) EXPECT() *_MockMultiReaderIteratorRecorder {
	return _m.recorder
}

func (_m *MockMultiReaderIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockMultiReaderIteratorRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

func (_m *MockMultiReaderIterator) Current() (ts.Datapoint, time0.Unit, ts.Annotation) {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(ts.Datapoint)
	ret1, _ := ret[1].(time0.Unit)
	ret2, _ := ret[2].(ts.Annotation)
	return ret0, ret1, ret2
}

func (_mr *_MockMultiReaderIteratorRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

func (_m *MockMultiReaderIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMultiReaderIteratorRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockMultiReaderIterator) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockMultiReaderIteratorRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockMultiReaderIterator) Reset(readers []io.Reader) {
	_m.ctrl.Call(_m, "Reset", readers)
}

func (_mr *_MockMultiReaderIteratorRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

func (_m *MockMultiReaderIterator) ResetSliceOfSlices(readers io0.ReaderSliceOfSlicesIterator) {
	_m.ctrl.Call(_m, "ResetSliceOfSlices", readers)
}

func (_mr *_MockMultiReaderIteratorRecorder) ResetSliceOfSlices(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResetSliceOfSlices", arg0)
}

// Mock of SeriesIterator interface
type MockSeriesIterator struct {
	ctrl     *gomock.Controller
	recorder *_MockSeriesIteratorRecorder
}

// Recorder for MockSeriesIterator (not exported)
type _MockSeriesIteratorRecorder struct {
	mock *MockSeriesIterator
}

func NewMockSeriesIterator(ctrl *gomock.Controller) *MockSeriesIterator {
	mock := &MockSeriesIterator{ctrl: ctrl}
	mock.recorder = &_MockSeriesIteratorRecorder{mock}
	return mock
}

func (_m *MockSeriesIterator) EXPECT() *_MockSeriesIteratorRecorder {
	return _m.recorder
}

func (_m *MockSeriesIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockSeriesIteratorRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

func (_m *MockSeriesIterator) Current() (ts.Datapoint, time0.Unit, ts.Annotation) {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(ts.Datapoint)
	ret1, _ := ret[1].(time0.Unit)
	ret2, _ := ret[2].(ts.Annotation)
	return ret0, ret1, ret2
}

func (_mr *_MockSeriesIteratorRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

func (_m *MockSeriesIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSeriesIteratorRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockSeriesIterator) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockSeriesIteratorRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockSeriesIterator) ID() string {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockSeriesIteratorRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockSeriesIterator) Start() time.Time {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockSeriesIteratorRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockSeriesIterator) End() time.Time {
	ret := _m.ctrl.Call(_m, "End")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockSeriesIteratorRecorder) End() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "End")
}

func (_m *MockSeriesIterator) Reset(id string, startInclusive time.Time, endExclusive time.Time, replicas []Iterator) {
	_m.ctrl.Call(_m, "Reset", id, startInclusive, endExclusive, replicas)
}

func (_mr *_MockSeriesIteratorRecorder) Reset(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0, arg1, arg2, arg3)
}

// Mock of SeriesIterators interface
type MockSeriesIterators struct {
	ctrl     *gomock.Controller
	recorder *_MockSeriesIteratorsRecorder
}

// Recorder for MockSeriesIterators (not exported)
type _MockSeriesIteratorsRecorder struct {
	mock *MockSeriesIterators
}

func NewMockSeriesIterators(ctrl *gomock.Controller) *MockSeriesIterators {
	mock := &MockSeriesIterators{ctrl: ctrl}
	mock.recorder = &_MockSeriesIteratorsRecorder{mock}
	return mock
}

func (_m *MockSeriesIterators) EXPECT() *_MockSeriesIteratorsRecorder {
	return _m.recorder
}

func (_m *MockSeriesIterators) Iters() []SeriesIterator {
	ret := _m.ctrl.Call(_m, "Iters")
	ret0, _ := ret[0].([]SeriesIterator)
	return ret0
}

func (_mr *_MockSeriesIteratorsRecorder) Iters() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Iters")
}

func (_m *MockSeriesIterators) Len() int {
	ret := _m.ctrl.Call(_m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockSeriesIteratorsRecorder) Len() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Len")
}

func (_m *MockSeriesIterators) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockSeriesIteratorsRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of MutableSeriesIterators interface
type MockMutableSeriesIterators struct {
	ctrl     *gomock.Controller
	recorder *_MockMutableSeriesIteratorsRecorder
}

// Recorder for MockMutableSeriesIterators (not exported)
type _MockMutableSeriesIteratorsRecorder struct {
	mock *MockMutableSeriesIterators
}

func NewMockMutableSeriesIterators(ctrl *gomock.Controller) *MockMutableSeriesIterators {
	mock := &MockMutableSeriesIterators{ctrl: ctrl}
	mock.recorder = &_MockMutableSeriesIteratorsRecorder{mock}
	return mock
}

func (_m *MockMutableSeriesIterators) EXPECT() *_MockMutableSeriesIteratorsRecorder {
	return _m.recorder
}

func (_m *MockMutableSeriesIterators) Iters() []SeriesIterator {
	ret := _m.ctrl.Call(_m, "Iters")
	ret0, _ := ret[0].([]SeriesIterator)
	return ret0
}

func (_mr *_MockMutableSeriesIteratorsRecorder) Iters() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Iters")
}

func (_m *MockMutableSeriesIterators) Len() int {
	ret := _m.ctrl.Call(_m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockMutableSeriesIteratorsRecorder) Len() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Len")
}

func (_m *MockMutableSeriesIterators) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockMutableSeriesIteratorsRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockMutableSeriesIterators) Reset(size int) {
	_m.ctrl.Call(_m, "Reset", size)
}

func (_mr *_MockMutableSeriesIteratorsRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

func (_m *MockMutableSeriesIterators) Cap() int {
	ret := _m.ctrl.Call(_m, "Cap")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockMutableSeriesIteratorsRecorder) Cap() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Cap")
}

func (_m *MockMutableSeriesIterators) SetAt(idx int, iter SeriesIterator) {
	_m.ctrl.Call(_m, "SetAt", idx, iter)
}

func (_mr *_MockMutableSeriesIteratorsRecorder) SetAt(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetAt", arg0, arg1)
}

// Mock of Decoder interface
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *_MockDecoderRecorder
}

// Recorder for MockDecoder (not exported)
type _MockDecoderRecorder struct {
	mock *MockDecoder
}

func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &_MockDecoderRecorder{mock}
	return mock
}

func (_m *MockDecoder) EXPECT() *_MockDecoderRecorder {
	return _m.recorder
}

func (_m *MockDecoder) Decode(reader io.Reader) ReaderIterator {
	ret := _m.ctrl.Call(_m, "Decode", reader)
	ret0, _ := ret[0].(ReaderIterator)
	return ret0
}

func (_mr *_MockDecoderRecorder) Decode(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Decode", arg0)
}

// Mock of IStream interface
type MockIStream struct {
	ctrl     *gomock.Controller
	recorder *_MockIStreamRecorder
}

// Recorder for MockIStream (not exported)
type _MockIStreamRecorder struct {
	mock *MockIStream
}

func NewMockIStream(ctrl *gomock.Controller) *MockIStream {
	mock := &MockIStream{ctrl: ctrl}
	mock.recorder = &_MockIStreamRecorder{mock}
	return mock
}

func (_m *MockIStream) EXPECT() *_MockIStreamRecorder {
	return _m.recorder
}

func (_m *MockIStream) ReadBit() (Bit, error) {
	ret := _m.ctrl.Call(_m, "ReadBit")
	ret0, _ := ret[0].(Bit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIStreamRecorder) ReadBit() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadBit")
}

func (_m *MockIStream) ReadByte() (byte, error) {
	ret := _m.ctrl.Call(_m, "ReadByte")
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIStreamRecorder) ReadByte() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadByte")
}

func (_m *MockIStream) ReadBits(numBits int) (uint64, error) {
	ret := _m.ctrl.Call(_m, "ReadBits", numBits)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIStreamRecorder) ReadBits(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadBits", arg0)
}

func (_m *MockIStream) PeekBits(numBits int) (uint64, error) {
	ret := _m.ctrl.Call(_m, "PeekBits", numBits)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIStreamRecorder) PeekBits(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PeekBits", arg0)
}

func (_m *MockIStream) Reset(r io.Reader) {
	_m.ctrl.Call(_m, "Reset", r)
}

func (_mr *_MockIStreamRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

// Mock of OStream interface
type MockOStream struct {
	ctrl     *gomock.Controller
	recorder *_MockOStreamRecorder
}

// Recorder for MockOStream (not exported)
type _MockOStreamRecorder struct {
	mock *MockOStream
}

func NewMockOStream(ctrl *gomock.Controller) *MockOStream {
	mock := &MockOStream{ctrl: ctrl}
	mock.recorder = &_MockOStreamRecorder{mock}
	return mock
}

func (_m *MockOStream) EXPECT() *_MockOStreamRecorder {
	return _m.recorder
}

func (_m *MockOStream) Clone() OStream {
	ret := _m.ctrl.Call(_m, "Clone")
	ret0, _ := ret[0].(OStream)
	return ret0
}

func (_mr *_MockOStreamRecorder) Clone() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Clone")
}

func (_m *MockOStream) Len() int {
	ret := _m.ctrl.Call(_m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOStreamRecorder) Len() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Len")
}

func (_m *MockOStream) Empty() bool {
	ret := _m.ctrl.Call(_m, "Empty")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockOStreamRecorder) Empty() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Empty")
}

func (_m *MockOStream) WriteBit(v Bit) {
	_m.ctrl.Call(_m, "WriteBit", v)
}

func (_mr *_MockOStreamRecorder) WriteBit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteBit", arg0)
}

func (_m *MockOStream) WriteBits(v uint64, numBits int) {
	_m.ctrl.Call(_m, "WriteBits", v, numBits)
}

func (_mr *_MockOStreamRecorder) WriteBits(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteBits", arg0, arg1)
}

func (_m *MockOStream) WriteByte(v byte) {
	_m.ctrl.Call(_m, "WriteByte", v)
}

func (_mr *_MockOStreamRecorder) WriteByte(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteByte", arg0)
}

func (_m *MockOStream) WriteBytes(bytes []byte) {
	_m.ctrl.Call(_m, "WriteBytes", bytes)
}

func (_mr *_MockOStreamRecorder) WriteBytes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteBytes", arg0)
}

func (_m *MockOStream) Reset(buffer []byte) {
	_m.ctrl.Call(_m, "Reset", buffer)
}

func (_mr *_MockOStreamRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

func (_m *MockOStream) Rawbytes() ([]byte, int) {
	ret := _m.ctrl.Call(_m, "Rawbytes")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

func (_mr *_MockOStreamRecorder) Rawbytes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rawbytes")
}

// Mock of EncoderPool interface
type MockEncoderPool struct {
	ctrl     *gomock.Controller
	recorder *_MockEncoderPoolRecorder
}

// Recorder for MockEncoderPool (not exported)
type _MockEncoderPoolRecorder struct {
	mock *MockEncoderPool
}

func NewMockEncoderPool(ctrl *gomock.Controller) *MockEncoderPool {
	mock := &MockEncoderPool{ctrl: ctrl}
	mock.recorder = &_MockEncoderPoolRecorder{mock}
	return mock
}

func (_m *MockEncoderPool) EXPECT() *_MockEncoderPoolRecorder {
	return _m.recorder
}

func (_m *MockEncoderPool) Init(alloc EncoderAllocate) {
	_m.ctrl.Call(_m, "Init", alloc)
}

func (_mr *_MockEncoderPoolRecorder) Init(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0)
}

func (_m *MockEncoderPool) Get() Encoder {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Encoder)
	return ret0
}

func (_mr *_MockEncoderPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockEncoderPool) Put(encoder Encoder) {
	_m.ctrl.Call(_m, "Put", encoder)
}

func (_mr *_MockEncoderPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of ReaderIteratorPool interface
type MockReaderIteratorPool struct {
	ctrl     *gomock.Controller
	recorder *_MockReaderIteratorPoolRecorder
}

// Recorder for MockReaderIteratorPool (not exported)
type _MockReaderIteratorPoolRecorder struct {
	mock *MockReaderIteratorPool
}

func NewMockReaderIteratorPool(ctrl *gomock.Controller) *MockReaderIteratorPool {
	mock := &MockReaderIteratorPool{ctrl: ctrl}
	mock.recorder = &_MockReaderIteratorPoolRecorder{mock}
	return mock
}

func (_m *MockReaderIteratorPool) EXPECT() *_MockReaderIteratorPoolRecorder {
	return _m.recorder
}

func (_m *MockReaderIteratorPool) Init(alloc ReaderIteratorAllocate) {
	_m.ctrl.Call(_m, "Init", alloc)
}

func (_mr *_MockReaderIteratorPoolRecorder) Init(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0)
}

func (_m *MockReaderIteratorPool) Get() ReaderIterator {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(ReaderIterator)
	return ret0
}

func (_mr *_MockReaderIteratorPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockReaderIteratorPool) Put(iter ReaderIterator) {
	_m.ctrl.Call(_m, "Put", iter)
}

func (_mr *_MockReaderIteratorPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of MultiReaderIteratorPool interface
type MockMultiReaderIteratorPool struct {
	ctrl     *gomock.Controller
	recorder *_MockMultiReaderIteratorPoolRecorder
}

// Recorder for MockMultiReaderIteratorPool (not exported)
type _MockMultiReaderIteratorPoolRecorder struct {
	mock *MockMultiReaderIteratorPool
}

func NewMockMultiReaderIteratorPool(ctrl *gomock.Controller) *MockMultiReaderIteratorPool {
	mock := &MockMultiReaderIteratorPool{ctrl: ctrl}
	mock.recorder = &_MockMultiReaderIteratorPoolRecorder{mock}
	return mock
}

func (_m *MockMultiReaderIteratorPool) EXPECT() *_MockMultiReaderIteratorPoolRecorder {
	return _m.recorder
}

func (_m *MockMultiReaderIteratorPool) Init(alloc ReaderIteratorAllocate) {
	_m.ctrl.Call(_m, "Init", alloc)
}

func (_mr *_MockMultiReaderIteratorPoolRecorder) Init(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0)
}

func (_m *MockMultiReaderIteratorPool) Get() MultiReaderIterator {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(MultiReaderIterator)
	return ret0
}

func (_mr *_MockMultiReaderIteratorPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockMultiReaderIteratorPool) Put(iter MultiReaderIterator) {
	_m.ctrl.Call(_m, "Put", iter)
}

func (_mr *_MockMultiReaderIteratorPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of SeriesIteratorPool interface
type MockSeriesIteratorPool struct {
	ctrl     *gomock.Controller
	recorder *_MockSeriesIteratorPoolRecorder
}

// Recorder for MockSeriesIteratorPool (not exported)
type _MockSeriesIteratorPoolRecorder struct {
	mock *MockSeriesIteratorPool
}

func NewMockSeriesIteratorPool(ctrl *gomock.Controller) *MockSeriesIteratorPool {
	mock := &MockSeriesIteratorPool{ctrl: ctrl}
	mock.recorder = &_MockSeriesIteratorPoolRecorder{mock}
	return mock
}

func (_m *MockSeriesIteratorPool) EXPECT() *_MockSeriesIteratorPoolRecorder {
	return _m.recorder
}

func (_m *MockSeriesIteratorPool) Init() {
	_m.ctrl.Call(_m, "Init")
}

func (_mr *_MockSeriesIteratorPoolRecorder) Init() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init")
}

func (_m *MockSeriesIteratorPool) Get() SeriesIterator {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(SeriesIterator)
	return ret0
}

func (_mr *_MockSeriesIteratorPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockSeriesIteratorPool) Put(iter SeriesIterator) {
	_m.ctrl.Call(_m, "Put", iter)
}

func (_mr *_MockSeriesIteratorPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of MutableSeriesIteratorsPool interface
type MockMutableSeriesIteratorsPool struct {
	ctrl     *gomock.Controller
	recorder *_MockMutableSeriesIteratorsPoolRecorder
}

// Recorder for MockMutableSeriesIteratorsPool (not exported)
type _MockMutableSeriesIteratorsPoolRecorder struct {
	mock *MockMutableSeriesIteratorsPool
}

func NewMockMutableSeriesIteratorsPool(ctrl *gomock.Controller) *MockMutableSeriesIteratorsPool {
	mock := &MockMutableSeriesIteratorsPool{ctrl: ctrl}
	mock.recorder = &_MockMutableSeriesIteratorsPoolRecorder{mock}
	return mock
}

func (_m *MockMutableSeriesIteratorsPool) EXPECT() *_MockMutableSeriesIteratorsPoolRecorder {
	return _m.recorder
}

func (_m *MockMutableSeriesIteratorsPool) Init() {
	_m.ctrl.Call(_m, "Init")
}

func (_mr *_MockMutableSeriesIteratorsPoolRecorder) Init() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init")
}

func (_m *MockMutableSeriesIteratorsPool) Get(size int) MutableSeriesIterators {
	ret := _m.ctrl.Call(_m, "Get", size)
	ret0, _ := ret[0].(MutableSeriesIterators)
	return ret0
}

func (_mr *_MockMutableSeriesIteratorsPoolRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockMutableSeriesIteratorsPool) Put(iters MutableSeriesIterators) {
	_m.ctrl.Call(_m, "Put", iters)
}

func (_mr *_MockMutableSeriesIteratorsPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of IteratorArrayPool interface
type MockIteratorArrayPool struct {
	ctrl     *gomock.Controller
	recorder *_MockIteratorArrayPoolRecorder
}

// Recorder for MockIteratorArrayPool (not exported)
type _MockIteratorArrayPoolRecorder struct {
	mock *MockIteratorArrayPool
}

func NewMockIteratorArrayPool(ctrl *gomock.Controller) *MockIteratorArrayPool {
	mock := &MockIteratorArrayPool{ctrl: ctrl}
	mock.recorder = &_MockIteratorArrayPoolRecorder{mock}
	return mock
}

func (_m *MockIteratorArrayPool) EXPECT() *_MockIteratorArrayPoolRecorder {
	return _m.recorder
}

func (_m *MockIteratorArrayPool) Init() {
	_m.ctrl.Call(_m, "Init")
}

func (_mr *_MockIteratorArrayPoolRecorder) Init() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init")
}

func (_m *MockIteratorArrayPool) Get(size int) []Iterator {
	ret := _m.ctrl.Call(_m, "Get", size)
	ret0, _ := ret[0].([]Iterator)
	return ret0
}

func (_mr *_MockIteratorArrayPoolRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockIteratorArrayPool) Put(iters []Iterator) {
	_m.ctrl.Call(_m, "Put", iters)
}

func (_mr *_MockIteratorArrayPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}
