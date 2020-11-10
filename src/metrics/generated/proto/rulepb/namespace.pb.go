// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/metrics/generated/proto/rulepb/namespace.proto

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

/*
	Package rulepb is a generated protocol buffer package.

	It is generated from these files:
		github.com/m3db/m3/src/metrics/generated/proto/rulepb/namespace.proto
		github.com/m3db/m3/src/metrics/generated/proto/rulepb/rule.proto

	It has these top-level messages:
		NamespaceSnapshot
		Namespace
		Namespaces
		MappingRuleSnapshot
		MappingRule
		RollupTarget
		RollupTargetV2
		RollupRuleSnapshot
		RollupRule
		RuleSet
*/
package rulepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type NamespaceSnapshot struct {
	ForRulesetVersion  int32  `protobuf:"varint,1,opt,name=for_ruleset_version,json=forRulesetVersion,proto3" json:"for_ruleset_version,omitempty"`
	Tombstoned         bool   `protobuf:"varint,2,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	LastUpdatedAtNanos int64  `protobuf:"varint,3,opt,name=last_updated_at_nanos,json=lastUpdatedAtNanos,proto3" json:"last_updated_at_nanos,omitempty"`
	LastUpdatedBy      string `protobuf:"bytes,4,opt,name=last_updated_by,json=lastUpdatedBy,proto3" json:"last_updated_by,omitempty"`
}

func (m *NamespaceSnapshot) Reset()                    { *m = NamespaceSnapshot{} }
func (m *NamespaceSnapshot) String() string            { return proto.CompactTextString(m) }
func (*NamespaceSnapshot) ProtoMessage()               {}
func (*NamespaceSnapshot) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{0} }

func (m *NamespaceSnapshot) GetForRulesetVersion() int32 {
	if m != nil {
		return m.ForRulesetVersion
	}
	return 0
}

func (m *NamespaceSnapshot) GetTombstoned() bool {
	if m != nil {
		return m.Tombstoned
	}
	return false
}

func (m *NamespaceSnapshot) GetLastUpdatedAtNanos() int64 {
	if m != nil {
		return m.LastUpdatedAtNanos
	}
	return 0
}

func (m *NamespaceSnapshot) GetLastUpdatedBy() string {
	if m != nil {
		return m.LastUpdatedBy
	}
	return ""
}

type Namespace struct {
	Name      string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Snapshots []*NamespaceSnapshot `protobuf:"bytes,2,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *Namespace) Reset()                    { *m = Namespace{} }
func (m *Namespace) String() string            { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()               {}
func (*Namespace) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{1} }

func (m *Namespace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Namespace) GetSnapshots() []*NamespaceSnapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type Namespaces struct {
	Namespaces []*Namespace `protobuf:"bytes,1,rep,name=namespaces" json:"namespaces,omitempty"`
}

func (m *Namespaces) Reset()                    { *m = Namespaces{} }
func (m *Namespaces) String() string            { return proto.CompactTextString(m) }
func (*Namespaces) ProtoMessage()               {}
func (*Namespaces) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{2} }

func (m *Namespaces) GetNamespaces() []*Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterType((*NamespaceSnapshot)(nil), "rulepb.NamespaceSnapshot")
	proto.RegisterType((*Namespace)(nil), "rulepb.Namespace")
	proto.RegisterType((*Namespaces)(nil), "rulepb.Namespaces")
}
func (m *NamespaceSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceSnapshot) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ForRulesetVersion != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.ForRulesetVersion))
	}
	if m.Tombstoned {
		dAtA[i] = 0x10
		i++
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.LastUpdatedAtNanos != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.LastUpdatedAtNanos))
	}
	if len(m.LastUpdatedBy) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.LastUpdatedBy)))
		i += copy(dAtA[i:], m.LastUpdatedBy)
	}
	return i, nil
}

func (m *Namespace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Namespace) MarshalTo(dAtA []byte) (int, error) {
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
	if len(m.Snapshots) > 0 {
		for _, msg := range m.Snapshots {
			dAtA[i] = 0x12
			i++
			i = encodeVarintNamespace(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Namespaces) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Namespaces) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespaces) > 0 {
		for _, msg := range m.Namespaces {
			dAtA[i] = 0xa
			i++
			i = encodeVarintNamespace(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
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
func (m *NamespaceSnapshot) Size() (n int) {
	var l int
	_ = l
	if m.ForRulesetVersion != 0 {
		n += 1 + sovNamespace(uint64(m.ForRulesetVersion))
	}
	if m.Tombstoned {
		n += 2
	}
	if m.LastUpdatedAtNanos != 0 {
		n += 1 + sovNamespace(uint64(m.LastUpdatedAtNanos))
	}
	l = len(m.LastUpdatedBy)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func (m *Namespace) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	if len(m.Snapshots) > 0 {
		for _, e := range m.Snapshots {
			l = e.Size()
			n += 1 + l + sovNamespace(uint64(l))
		}
	}
	return n
}

func (m *Namespaces) Size() (n int) {
	var l int
	_ = l
	if len(m.Namespaces) > 0 {
		for _, e := range m.Namespaces {
			l = e.Size()
			n += 1 + l + sovNamespace(uint64(l))
		}
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
func (m *NamespaceSnapshot) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NamespaceSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForRulesetVersion", wireType)
			}
			m.ForRulesetVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ForRulesetVersion |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
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
			m.Tombstoned = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdatedAtNanos", wireType)
			}
			m.LastUpdatedAtNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastUpdatedAtNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdatedBy", wireType)
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
			m.LastUpdatedBy = string(dAtA[iNdEx:postIndex])
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
func (m *Namespace) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Namespace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Namespace: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Snapshots", wireType)
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
			m.Snapshots = append(m.Snapshots, &NamespaceSnapshot{})
			if err := m.Snapshots[len(m.Snapshots)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Namespaces) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Namespaces: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Namespaces: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespaces", wireType)
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
			m.Namespaces = append(m.Namespaces, &Namespace{})
			if err := m.Namespaces[len(m.Namespaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/metrics/generated/proto/rulepb/namespace.proto", fileDescriptorNamespace)
}

var fileDescriptorNamespace = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0xbf, 0xb4, 0xfd, 0x8a, 0x73, 0x45, 0xa4, 0x11, 0x61, 0xdc, 0x0c, 0x43, 0x17, 0x32,
	0xab, 0x09, 0xb5, 0x88, 0x4b, 0xb1, 0x20, 0xee, 0xba, 0x88, 0x28, 0xe2, 0x66, 0x48, 0x66, 0xd2,
	0x3f, 0xd0, 0x24, 0x43, 0x6e, 0x2a, 0xf4, 0x2d, 0x7c, 0x22, 0xd7, 0x2e, 0x7d, 0x04, 0xa9, 0x2f,
	0x22, 0x33, 0xa3, 0xb5, 0xd2, 0x9d, 0xbb, 0xe4, 0xfc, 0x4e, 0x2e, 0xe7, 0xe4, 0xc2, 0xf5, 0x74,
	0xee, 0x67, 0x4b, 0x99, 0xe6, 0x56, 0x33, 0x3d, 0x2c, 0x24, 0xd3, 0x43, 0x86, 0x2e, 0x67, 0x5a,
	0x79, 0x37, 0xcf, 0x91, 0x4d, 0x95, 0x51, 0x4e, 0x78, 0x55, 0xb0, 0xd2, 0x59, 0x6f, 0x99, 0x5b,
	0x2e, 0x54, 0x29, 0x99, 0x11, 0x5a, 0x61, 0x29, 0x72, 0x95, 0xd6, 0x32, 0xed, 0x36, 0x7a, 0xff,
	0x85, 0x40, 0x6f, 0xfc, 0xcd, 0x6e, 0x8d, 0x28, 0x71, 0x66, 0x3d, 0x4d, 0xe1, 0x68, 0x62, 0x5d,
	0x56, 0x79, 0x50, 0xf9, 0xec, 0x49, 0x39, 0x9c, 0x5b, 0x13, 0x92, 0x98, 0x24, 0xff, 0x79, 0x6f,
	0x62, 0x1d, 0x6f, 0xc8, 0x7d, 0x03, 0x68, 0x04, 0xe0, 0xad, 0x96, 0xe8, 0xad, 0x51, 0x45, 0xd8,
	0x8a, 0x49, 0xb2, 0xc7, 0xb7, 0x14, 0x3a, 0x80, 0xe3, 0x85, 0x40, 0x9f, 0x2d, 0xcb, 0xa2, 0x8a,
	0x96, 0x09, 0x9f, 0x19, 0x61, 0x2c, 0x86, 0xed, 0x98, 0x24, 0x6d, 0x4e, 0x2b, 0x78, 0xd7, 0xb0,
	0x2b, 0x3f, 0xae, 0x08, 0x3d, 0x85, 0xc3, 0x5f, 0x4f, 0xe4, 0x2a, 0xec, 0xc4, 0x24, 0x09, 0xf8,
	0xc1, 0x96, 0x79, 0xb4, 0xea, 0x3f, 0x40, 0xb0, 0xc9, 0x4f, 0x29, 0x74, 0xaa, 0xa2, 0x75, 0xd0,
	0x80, 0xd7, 0x67, 0x7a, 0x01, 0x01, 0x7e, 0xf5, 0xc2, 0xb0, 0x15, 0xb7, 0x93, 0xfd, 0xb3, 0x93,
	0xb4, 0x69, 0x9f, 0xee, 0x34, 0xe7, 0x3f, 0xde, 0xfe, 0x25, 0xc0, 0x86, 0x23, 0x1d, 0x00, 0x6c,
	0xfe, 0x10, 0x43, 0x52, 0xcf, 0xe9, 0xed, 0xcc, 0xe1, 0x5b, 0xa6, 0xd1, 0xcd, 0xeb, 0x3a, 0x22,
	0x6f, 0xeb, 0x88, 0xbc, 0xaf, 0x23, 0xf2, 0xfc, 0x11, 0xfd, 0x7b, 0x3c, 0xff, 0xd3, 0xf2, 0x64,
	0xb7, 0xbe, 0x0d, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x2c, 0xd7, 0x1c, 0xfc, 0x01, 0x00,
	0x00,
}
