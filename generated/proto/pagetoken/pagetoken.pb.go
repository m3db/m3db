// Copyright (c) 2017 Uber Technologies, Inc.
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

// Code generated by protoc-gen-go.
// source: pagetoken.proto
// DO NOT EDIT!

/*
Package pagetoken is a generated protocol buffer package.

It is generated from these files:
	pagetoken.proto

It has these top-level messages:
	PageToken
*/
package pagetoken

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PageToken struct {
	// Types that are valid to be assigned to Phase:
	//	*PageToken_ActiveSeriesPhase_
	//	*PageToken_FlushedSeriesPhase_
	Phase isPageToken_Phase `protobuf_oneof:"phase"`
}

func (m *PageToken) Reset()                    { *m = PageToken{} }
func (m *PageToken) String() string            { return proto.CompactTextString(m) }
func (*PageToken) ProtoMessage()               {}
func (*PageToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isPageToken_Phase interface {
	isPageToken_Phase()
}

type PageToken_ActiveSeriesPhase_ struct {
	ActiveSeriesPhase *PageToken_ActiveSeriesPhase `protobuf:"bytes,1,opt,name=active_series_phase,json=activeSeriesPhase,oneof"`
}
type PageToken_FlushedSeriesPhase_ struct {
	FlushedSeriesPhase *PageToken_FlushedSeriesPhase `protobuf:"bytes,2,opt,name=flushed_series_phase,json=flushedSeriesPhase,oneof"`
}

func (*PageToken_ActiveSeriesPhase_) isPageToken_Phase()  {}
func (*PageToken_FlushedSeriesPhase_) isPageToken_Phase() {}

func (m *PageToken) GetPhase() isPageToken_Phase {
	if m != nil {
		return m.Phase
	}
	return nil
}

func (m *PageToken) GetActiveSeriesPhase() *PageToken_ActiveSeriesPhase {
	if x, ok := m.GetPhase().(*PageToken_ActiveSeriesPhase_); ok {
		return x.ActiveSeriesPhase
	}
	return nil
}

func (m *PageToken) GetFlushedSeriesPhase() *PageToken_FlushedSeriesPhase {
	if x, ok := m.GetPhase().(*PageToken_FlushedSeriesPhase_); ok {
		return x.FlushedSeriesPhase
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PageToken) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PageToken_OneofMarshaler, _PageToken_OneofUnmarshaler, _PageToken_OneofSizer, []interface{}{
		(*PageToken_ActiveSeriesPhase_)(nil),
		(*PageToken_FlushedSeriesPhase_)(nil),
	}
}

func _PageToken_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PageToken)
	// phase
	switch x := m.Phase.(type) {
	case *PageToken_ActiveSeriesPhase_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ActiveSeriesPhase); err != nil {
			return err
		}
	case *PageToken_FlushedSeriesPhase_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FlushedSeriesPhase); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PageToken.Phase has unexpected type %T", x)
	}
	return nil
}

func _PageToken_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PageToken)
	switch tag {
	case 1: // phase.active_series_phase
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PageToken_ActiveSeriesPhase)
		err := b.DecodeMessage(msg)
		m.Phase = &PageToken_ActiveSeriesPhase_{msg}
		return true, err
	case 2: // phase.flushed_series_phase
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PageToken_FlushedSeriesPhase)
		err := b.DecodeMessage(msg)
		m.Phase = &PageToken_FlushedSeriesPhase_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PageToken_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PageToken)
	// phase
	switch x := m.Phase.(type) {
	case *PageToken_ActiveSeriesPhase_:
		s := proto.Size(x.ActiveSeriesPhase)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PageToken_FlushedSeriesPhase_:
		s := proto.Size(x.FlushedSeriesPhase)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PageToken_ActiveSeriesPhase struct {
	IndexCursor int64 `protobuf:"varint,1,opt,name=indexCursor" json:"indexCursor,omitempty"`
}

func (m *PageToken_ActiveSeriesPhase) Reset()                    { *m = PageToken_ActiveSeriesPhase{} }
func (m *PageToken_ActiveSeriesPhase) String() string            { return proto.CompactTextString(m) }
func (*PageToken_ActiveSeriesPhase) ProtoMessage()               {}
func (*PageToken_ActiveSeriesPhase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type PageToken_FlushedSeriesPhase struct {
	CurrBlockStartUnixNanos int64 `protobuf:"varint,1,opt,name=currBlockStartUnixNanos" json:"currBlockStartUnixNanos,omitempty"`
	CurrBlockIndexRead      int64 `protobuf:"varint,2,opt,name=currBlockIndexRead" json:"currBlockIndexRead,omitempty"`
}

func (m *PageToken_FlushedSeriesPhase) Reset()                    { *m = PageToken_FlushedSeriesPhase{} }
func (m *PageToken_FlushedSeriesPhase) String() string            { return proto.CompactTextString(m) }
func (*PageToken_FlushedSeriesPhase) ProtoMessage()               {}
func (*PageToken_FlushedSeriesPhase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func init() {
	proto.RegisterType((*PageToken)(nil), "pagetoken.PageToken")
	proto.RegisterType((*PageToken_ActiveSeriesPhase)(nil), "pagetoken.PageToken.ActiveSeriesPhase")
	proto.RegisterType((*PageToken_FlushedSeriesPhase)(nil), "pagetoken.PageToken.FlushedSeriesPhase")
}

func init() { proto.RegisterFile("pagetoken.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x48, 0x4c, 0x4f,
	0x2d, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x35, 0x31, 0x73, 0x71, 0x06, 0x24, 0xa6, 0xa7, 0x86, 0x80, 0x78, 0x42, 0x11, 0x5c, 0xc2, 0x89,
	0xc9, 0x25, 0x99, 0x65, 0xa9, 0xf1, 0xc5, 0xa9, 0x45, 0x99, 0xa9, 0xc5, 0xf1, 0x05, 0x19, 0x89,
	0xc5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x6a, 0x7a, 0x08, 0x73, 0xe0, 0x5a, 0xf4,
	0x1c, 0xc1, 0xea, 0x83, 0xc1, 0xca, 0x03, 0x40, 0xaa, 0x3d, 0x18, 0x82, 0x04, 0x13, 0xd1, 0x05,
	0x85, 0xa2, 0xb9, 0x44, 0xd2, 0x72, 0x4a, 0x8b, 0x33, 0x52, 0x53, 0x50, 0x8d, 0x66, 0x02, 0x1b,
	0xad, 0x8e, 0xd5, 0x68, 0x37, 0x88, 0x06, 0x54, 0xb3, 0x85, 0xd2, 0x30, 0x44, 0xa5, 0x4c, 0xb9,
	0x04, 0x31, 0x9c, 0x21, 0xa4, 0xc0, 0xc5, 0x9d, 0x99, 0x97, 0x92, 0x5a, 0xe1, 0x5c, 0x5a, 0x54,
	0x9c, 0x5f, 0x04, 0xf6, 0x03, 0x73, 0x10, 0xb2, 0x90, 0x54, 0x1d, 0x97, 0x10, 0xa6, 0x15, 0x42,
	0x16, 0x5c, 0xe2, 0xc9, 0xa5, 0x45, 0x45, 0x4e, 0x39, 0xf9, 0xc9, 0xd9, 0xc1, 0x25, 0x89, 0x45,
	0x25, 0xa1, 0x79, 0x99, 0x15, 0x7e, 0x89, 0x79, 0xf9, 0xc5, 0x50, 0x33, 0x70, 0x49, 0x0b, 0xe9,
	0x71, 0x09, 0xc1, 0xa5, 0x3c, 0x41, 0xf6, 0x04, 0xa5, 0x26, 0xa6, 0x80, 0x7d, 0xc8, 0x1c, 0x84,
	0x45, 0xc6, 0x89, 0x9d, 0x8b, 0x15, 0x1c, 0x08, 0x49, 0x6c, 0xe0, 0x68, 0x31, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xf7, 0x99, 0xc8, 0x22, 0xa9, 0x01, 0x00, 0x00,
}
