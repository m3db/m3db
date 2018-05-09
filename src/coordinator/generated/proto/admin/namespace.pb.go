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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: namespace.proto

/*
	Package admin is a generated protocol buffer package.

	It is generated from these files:
		namespace.proto
		placement.proto

	It has these top-level messages:
		NamespaceGetResponse
		NamespaceAddRequest
		PlacementInitRequest
		PlacementGetResponse
		PlacementAddRequest
*/
package admin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import namespace "github.com/m3db/m3db/generated/proto/namespace"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NamespaceGetResponse struct {
	Registry *namespace.Registry `protobuf:"bytes,1,opt,name=registry" json:"registry,omitempty"`
}

func (m *NamespaceGetResponse) Reset()                    { *m = NamespaceGetResponse{} }
func (m *NamespaceGetResponse) String() string            { return proto.CompactTextString(m) }
func (*NamespaceGetResponse) ProtoMessage()               {}
func (*NamespaceGetResponse) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{0} }

func (m *NamespaceGetResponse) GetRegistry() *namespace.Registry {
	if m != nil {
		return m.Registry
	}
	return nil
}

type NamespaceAddRequest struct {
	Name                  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RetentionPeriod       string `protobuf:"bytes,2,opt,name=retention_period,json=retentionPeriod,proto3" json:"retention_period,omitempty"`
	BlockSize             string `protobuf:"bytes,3,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	BufferFuture          string `protobuf:"bytes,4,opt,name=buffer_future,json=bufferFuture,proto3" json:"buffer_future,omitempty"`
	BufferPast            string `protobuf:"bytes,5,opt,name=buffer_past,json=bufferPast,proto3" json:"buffer_past,omitempty"`
	BlockDataExpiry       bool   `protobuf:"varint,6,opt,name=block_data_expiry,json=blockDataExpiry,proto3" json:"block_data_expiry,omitempty"`
	BlockDataExpiryPeriod string `protobuf:"bytes,7,opt,name=block_data_expiry_period,json=blockDataExpiryPeriod,proto3" json:"block_data_expiry_period,omitempty"`
	BootstrapEnabled      bool   `protobuf:"varint,8,opt,name=bootstrap_enabled,json=bootstrapEnabled,proto3" json:"bootstrap_enabled,omitempty"`
	CleanupEnabled        bool   `protobuf:"varint,9,opt,name=cleanup_enabled,json=cleanupEnabled,proto3" json:"cleanup_enabled,omitempty"`
	FlushEnabled          bool   `protobuf:"varint,10,opt,name=flush_enabled,json=flushEnabled,proto3" json:"flush_enabled,omitempty"`
	RepairEnabled         bool   `protobuf:"varint,11,opt,name=repair_enabled,json=repairEnabled,proto3" json:"repair_enabled,omitempty"`
	WritesToCommitlog     bool   `protobuf:"varint,12,opt,name=writes_to_commitlog,json=writesToCommitlog,proto3" json:"writes_to_commitlog,omitempty"`
}

func (m *NamespaceAddRequest) Reset()                    { *m = NamespaceAddRequest{} }
func (m *NamespaceAddRequest) String() string            { return proto.CompactTextString(m) }
func (*NamespaceAddRequest) ProtoMessage()               {}
func (*NamespaceAddRequest) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{1} }

func (m *NamespaceAddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamespaceAddRequest) GetRetentionPeriod() string {
	if m != nil {
		return m.RetentionPeriod
	}
	return ""
}

func (m *NamespaceAddRequest) GetBlockSize() string {
	if m != nil {
		return m.BlockSize
	}
	return ""
}

func (m *NamespaceAddRequest) GetBufferFuture() string {
	if m != nil {
		return m.BufferFuture
	}
	return ""
}

func (m *NamespaceAddRequest) GetBufferPast() string {
	if m != nil {
		return m.BufferPast
	}
	return ""
}

func (m *NamespaceAddRequest) GetBlockDataExpiry() bool {
	if m != nil {
		return m.BlockDataExpiry
	}
	return false
}

func (m *NamespaceAddRequest) GetBlockDataExpiryPeriod() string {
	if m != nil {
		return m.BlockDataExpiryPeriod
	}
	return ""
}

func (m *NamespaceAddRequest) GetBootstrapEnabled() bool {
	if m != nil {
		return m.BootstrapEnabled
	}
	return false
}

func (m *NamespaceAddRequest) GetCleanupEnabled() bool {
	if m != nil {
		return m.CleanupEnabled
	}
	return false
}

func (m *NamespaceAddRequest) GetFlushEnabled() bool {
	if m != nil {
		return m.FlushEnabled
	}
	return false
}

func (m *NamespaceAddRequest) GetRepairEnabled() bool {
	if m != nil {
		return m.RepairEnabled
	}
	return false
}

func (m *NamespaceAddRequest) GetWritesToCommitlog() bool {
	if m != nil {
		return m.WritesToCommitlog
	}
	return false
}

func init() {
	proto.RegisterType((*NamespaceGetResponse)(nil), "admin.NamespaceGetResponse")
	proto.RegisterType((*NamespaceAddRequest)(nil), "admin.NamespaceAddRequest")
}
func (m *NamespaceGetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceGetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Registry != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.Registry.Size()))
		n1, err := m.Registry.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *NamespaceAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceAddRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.RetentionPeriod) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.RetentionPeriod)))
		i += copy(dAtA[i:], m.RetentionPeriod)
	}
	if len(m.BlockSize) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.BlockSize)))
		i += copy(dAtA[i:], m.BlockSize)
	}
	if len(m.BufferFuture) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.BufferFuture)))
		i += copy(dAtA[i:], m.BufferFuture)
	}
	if len(m.BufferPast) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.BufferPast)))
		i += copy(dAtA[i:], m.BufferPast)
	}
	if m.BlockDataExpiry {
		dAtA[i] = 0x30
		i++
		if m.BlockDataExpiry {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.BlockDataExpiryPeriod) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.BlockDataExpiryPeriod)))
		i += copy(dAtA[i:], m.BlockDataExpiryPeriod)
	}
	if m.BootstrapEnabled {
		dAtA[i] = 0x40
		i++
		if m.BootstrapEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CleanupEnabled {
		dAtA[i] = 0x48
		i++
		if m.CleanupEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.FlushEnabled {
		dAtA[i] = 0x50
		i++
		if m.FlushEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RepairEnabled {
		dAtA[i] = 0x58
		i++
		if m.RepairEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.WritesToCommitlog {
		dAtA[i] = 0x60
		i++
		if m.WritesToCommitlog {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintNamespace(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NamespaceGetResponse) Size() (n int) {
	var l int
	_ = l
	if m.Registry != nil {
		l = m.Registry.Size()
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func (m *NamespaceAddRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	l = len(m.RetentionPeriod)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	l = len(m.BlockSize)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	l = len(m.BufferFuture)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	l = len(m.BufferPast)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	if m.BlockDataExpiry {
		n += 2
	}
	l = len(m.BlockDataExpiryPeriod)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	if m.BootstrapEnabled {
		n += 2
	}
	if m.CleanupEnabled {
		n += 2
	}
	if m.FlushEnabled {
		n += 2
	}
	if m.RepairEnabled {
		n += 2
	}
	if m.WritesToCommitlog {
		n += 2
	}
	return n
}

func sovNamespace(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNamespace(x uint64) (n int) {
	return sovNamespace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NamespaceGetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespaceGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Registry == nil {
				m.Registry = &namespace.Registry{}
			}
			if err := m.Registry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NamespaceAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespaceAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetentionPeriod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RetentionPeriod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockSize", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockSize = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferFuture", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BufferFuture = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferPast", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BufferPast = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockDataExpiry", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BlockDataExpiry = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockDataExpiryPeriod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockDataExpiryPeriod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BootstrapEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BootstrapEnabled = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CleanupEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CleanupEnabled = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlushEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FlushEnabled = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepairEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RepairEnabled = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WritesToCommitlog", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WritesToCommitlog = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNamespace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNamespace
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNamespace
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNamespace(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNamespace = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNamespace   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("namespace.proto", fileDescriptorNamespace) }

var fileDescriptorNamespace = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0x5d, 0xfb, 0xc7, 0xe4, 0x24, 0x69, 0x92, 0x89, 0xc2, 0x20, 0x18, 0x4b, 0x8b, 0x58,
	0x15, 0xb2, 0x60, 0x2f, 0xbc, 0x13, 0xfc, 0x53, 0x7b, 0x27, 0x65, 0xf5, 0x7e, 0x99, 0xdd, 0x3d,
	0x49, 0x07, 0x77, 0x67, 0xc6, 0x99, 0xb3, 0x68, 0xfb, 0x24, 0x3e, 0x92, 0x97, 0x3e, 0x82, 0xc4,
	0x17, 0xf0, 0x11, 0xc4, 0xb3, 0xd9, 0x2d, 0xa4, 0x37, 0xcb, 0xf2, 0xfb, 0x7e, 0xdf, 0x37, 0x0c,
	0x0c, 0x8c, 0x8d, 0xaa, 0x30, 0x38, 0x95, 0xe3, 0xc2, 0x79, 0x4b, 0x56, 0xec, 0xa9, 0xa2, 0xd2,
	0xe6, 0xe1, 0xeb, 0x95, 0xa6, 0xcb, 0x3a, 0x5b, 0xe4, 0xb6, 0x8a, 0xab, 0xd3, 0x22, 0x6b, 0x3e,
	0x2b, 0x34, 0xe8, 0x15, 0x61, 0x11, 0xb3, 0x1c, 0x77, 0xe5, 0x78, 0x6b, 0xe6, 0xe8, 0x1c, 0xee,
	0x7f, 0x6c, 0xd1, 0x39, 0x52, 0x82, 0xc1, 0x59, 0x13, 0x50, 0xc4, 0xd0, 0xf3, 0xb8, 0xd2, 0x81,
	0xfc, 0x95, 0x8c, 0x0e, 0xa3, 0x93, 0xc1, 0xcb, 0xd9, 0xe2, 0xa6, 0x9b, 0x6c, 0xa2, 0xa4, 0x93,
	0x8e, 0xfe, 0xee, 0xc0, 0xac, 0x5b, 0x7a, 0x53, 0x14, 0x09, 0x7e, 0xad, 0x31, 0x90, 0x10, 0xb0,
	0xfb, 0xbf, 0xc7, 0x23, 0xfd, 0x84, 0xff, 0xc5, 0x33, 0x98, 0x78, 0x24, 0x34, 0xa4, 0xad, 0x49,
	0x1d, 0x7a, 0x6d, 0x0b, 0x79, 0x97, 0xf3, 0x71, 0xc7, 0x2f, 0x18, 0x8b, 0x47, 0x00, 0x59, 0x69,
	0xf3, 0x2f, 0x69, 0xd0, 0xd7, 0x28, 0x77, 0x58, 0xea, 0x33, 0xf9, 0xa4, 0xaf, 0x51, 0x1c, 0xc3,
	0x28, 0xab, 0x97, 0x4b, 0xf4, 0xe9, 0xb2, 0xa6, 0xda, 0xa3, 0xdc, 0x65, 0x63, 0xd8, 0xc0, 0x0f,
	0xcc, 0xc4, 0x63, 0x18, 0x6c, 0x24, 0xa7, 0x02, 0xc9, 0x3d, 0x56, 0xa0, 0x41, 0x17, 0x2a, 0x90,
	0x78, 0x0e, 0xd3, 0xe6, 0x90, 0x42, 0x91, 0x4a, 0xf1, 0xbb, 0xd3, 0xfe, 0x4a, 0xee, 0x1f, 0x46,
	0x27, 0xbd, 0x64, 0xcc, 0xc1, 0x7b, 0x45, 0xea, 0x8c, 0xb1, 0x78, 0x05, 0xf2, 0x96, 0xdb, 0xde,
	0xe1, 0x1e, 0x2f, 0x3f, 0xd8, 0xaa, 0x6c, 0x6e, 0xf2, 0x02, 0xa6, 0x99, 0xb5, 0x14, 0xc8, 0x2b,
	0x97, 0xa2, 0x51, 0x59, 0x89, 0x85, 0xec, 0xf1, 0x21, 0x93, 0x2e, 0x38, 0x6b, 0xb8, 0x78, 0x0a,
	0xe3, 0xbc, 0x44, 0x65, 0xea, 0x1b, 0xb5, 0xcf, 0xea, 0xc1, 0x06, 0xb7, 0xe2, 0x31, 0x8c, 0x96,
	0x65, 0x1d, 0x2e, 0x3b, 0x0d, 0x58, 0x1b, 0x32, 0x6c, 0xa5, 0x27, 0x70, 0xe0, 0xd1, 0x29, 0xed,
	0x3b, 0x6b, 0xc0, 0xd6, 0xa8, 0xa1, 0xad, 0xb6, 0x80, 0xd9, 0x37, 0xaf, 0x09, 0x43, 0x4a, 0x36,
	0xcd, 0x6d, 0x55, 0x69, 0x2a, 0xed, 0x4a, 0x0e, 0xd9, 0x9d, 0x36, 0xd1, 0x67, 0xfb, 0xae, 0x0d,
	0xde, 0x4e, 0x7e, 0xae, 0xe7, 0xd1, 0xaf, 0xf5, 0x3c, 0xfa, 0xbd, 0x9e, 0x47, 0x3f, 0xfe, 0xcc,
	0xef, 0x64, 0xfb, 0xfc, 0xa8, 0x4e, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x25, 0x48, 0x59,
	0xae, 0x02, 0x00, 0x00,
}
