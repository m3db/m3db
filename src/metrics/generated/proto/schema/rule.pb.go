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
// source: rule.proto
// DO NOT EDIT!

package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MappingRuleSnapshot struct {
	Name               string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tombstoned         bool              `protobuf:"varint,2,opt,name=tombstoned" json:"tombstoned,omitempty"`
	CutoverNanos       int64             `protobuf:"varint,3,opt,name=cutover_nanos,json=cutoverNanos" json:"cutover_nanos,omitempty"`
	TagFilters         map[string]string `protobuf:"bytes,4,rep,name=tag_filters,json=tagFilters" json:"tag_filters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Policies           []*Policy         `protobuf:"bytes,5,rep,name=policies" json:"policies,omitempty"`
	LastUpdatedAtNanos int64             `protobuf:"varint,6,opt,name=last_updated_at_nanos,json=lastUpdatedAtNanos" json:"last_updated_at_nanos,omitempty"`
	LastUpdatedBy      string            `protobuf:"bytes,7,opt,name=last_updated_by,json=lastUpdatedBy" json:"last_updated_by,omitempty"`
}

func (m *MappingRuleSnapshot) Reset()                    { *m = MappingRuleSnapshot{} }
func (m *MappingRuleSnapshot) String() string            { return proto.CompactTextString(m) }
func (*MappingRuleSnapshot) ProtoMessage()               {}
func (*MappingRuleSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *MappingRuleSnapshot) GetTagFilters() map[string]string {
	if m != nil {
		return m.TagFilters
	}
	return nil
}

func (m *MappingRuleSnapshot) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type MappingRule struct {
	Uuid      string                 `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Snapshots []*MappingRuleSnapshot `protobuf:"bytes,2,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *MappingRule) Reset()                    { *m = MappingRule{} }
func (m *MappingRule) String() string            { return proto.CompactTextString(m) }
func (*MappingRule) ProtoMessage()               {}
func (*MappingRule) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *MappingRule) GetSnapshots() []*MappingRuleSnapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type RollupTarget struct {
	Name     string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags     []string  `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
	Policies []*Policy `protobuf:"bytes,3,rep,name=policies" json:"policies,omitempty"`
}

func (m *RollupTarget) Reset()                    { *m = RollupTarget{} }
func (m *RollupTarget) String() string            { return proto.CompactTextString(m) }
func (*RollupTarget) ProtoMessage()               {}
func (*RollupTarget) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *RollupTarget) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type RollupRuleSnapshot struct {
	Name               string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tombstoned         bool              `protobuf:"varint,2,opt,name=tombstoned" json:"tombstoned,omitempty"`
	CutoverNanos       int64             `protobuf:"varint,3,opt,name=cutover_nanos,json=cutoverNanos" json:"cutover_nanos,omitempty"`
	TagFilters         map[string]string `protobuf:"bytes,4,rep,name=tag_filters,json=tagFilters" json:"tag_filters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Targets            []*RollupTarget   `protobuf:"bytes,5,rep,name=targets" json:"targets,omitempty"`
	LastUpdatedAtNanos int64             `protobuf:"varint,6,opt,name=last_updated_at_nanos,json=lastUpdatedAtNanos" json:"last_updated_at_nanos,omitempty"`
	LastUpdatedBy      string            `protobuf:"bytes,7,opt,name=last_updated_by,json=lastUpdatedBy" json:"last_updated_by,omitempty"`
}

func (m *RollupRuleSnapshot) Reset()                    { *m = RollupRuleSnapshot{} }
func (m *RollupRuleSnapshot) String() string            { return proto.CompactTextString(m) }
func (*RollupRuleSnapshot) ProtoMessage()               {}
func (*RollupRuleSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *RollupRuleSnapshot) GetTagFilters() map[string]string {
	if m != nil {
		return m.TagFilters
	}
	return nil
}

func (m *RollupRuleSnapshot) GetTargets() []*RollupTarget {
	if m != nil {
		return m.Targets
	}
	return nil
}

type RollupRule struct {
	Uuid      string                `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Snapshots []*RollupRuleSnapshot `protobuf:"bytes,2,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *RollupRule) Reset()                    { *m = RollupRule{} }
func (m *RollupRule) String() string            { return proto.CompactTextString(m) }
func (*RollupRule) ProtoMessage()               {}
func (*RollupRule) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *RollupRule) GetSnapshots() []*RollupRuleSnapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type RuleSet struct {
	Uuid               string         `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Namespace          string         `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	CreatedAtNanos     int64          `protobuf:"varint,3,opt,name=created_at_nanos,json=createdAtNanos" json:"created_at_nanos,omitempty"`
	LastUpdatedAtNanos int64          `protobuf:"varint,4,opt,name=last_updated_at_nanos,json=lastUpdatedAtNanos" json:"last_updated_at_nanos,omitempty"`
	Tombstoned         bool           `protobuf:"varint,5,opt,name=tombstoned" json:"tombstoned,omitempty"`
	CutoverNanos       int64          `protobuf:"varint,6,opt,name=cutover_nanos,json=cutoverNanos" json:"cutover_nanos,omitempty"`
	MappingRules       []*MappingRule `protobuf:"bytes,7,rep,name=mapping_rules,json=mappingRules" json:"mapping_rules,omitempty"`
	RollupRules        []*RollupRule  `protobuf:"bytes,8,rep,name=rollup_rules,json=rollupRules" json:"rollup_rules,omitempty"`
	LastUpdatedBy      string         `protobuf:"bytes,9,opt,name=last_updated_by,json=lastUpdatedBy" json:"last_updated_by,omitempty"`
}

func (m *RuleSet) Reset()                    { *m = RuleSet{} }
func (m *RuleSet) String() string            { return proto.CompactTextString(m) }
func (*RuleSet) ProtoMessage()               {}
func (*RuleSet) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *RuleSet) GetMappingRules() []*MappingRule {
	if m != nil {
		return m.MappingRules
	}
	return nil
}

func (m *RuleSet) GetRollupRules() []*RollupRule {
	if m != nil {
		return m.RollupRules
	}
	return nil
}

func init() {
	proto.RegisterType((*MappingRuleSnapshot)(nil), "schema.MappingRuleSnapshot")
	proto.RegisterType((*MappingRule)(nil), "schema.MappingRule")
	proto.RegisterType((*RollupTarget)(nil), "schema.RollupTarget")
	proto.RegisterType((*RollupRuleSnapshot)(nil), "schema.RollupRuleSnapshot")
	proto.RegisterType((*RollupRule)(nil), "schema.RollupRule")
	proto.RegisterType((*RuleSet)(nil), "schema.RuleSet")
}

func init() { proto.RegisterFile("rule.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x49, 0xd3, 0x5f, 0x99, 0xa6, 0x77, 0xc7, 0xdc, 0x09, 0xa1, 0x8a, 0x94, 0x0a, 0x12,
	0x4e, 0x08, 0xa8, 0x08, 0x55, 0xf0, 0x41, 0x41, 0x5f, 0xfc, 0x81, 0xac, 0xe7, 0x8b, 0x08, 0x61,
	0x9b, 0xae, 0xb9, 0x60, 0x7e, 0xb1, 0xbb, 0x39, 0xc8, 0xab, 0xf8, 0x37, 0xfb, 0x2c, 0xd9, 0x4d,
	0xda, 0xc6, 0xa6, 0xf5, 0x4d, 0xee, 0x6d, 0x32, 0x33, 0x3b, 0xf9, 0xce, 0x67, 0xbf, 0x09, 0x00,
	0x2f, 0x62, 0xe6, 0xe5, 0x3c, 0x93, 0x19, 0x0e, 0x45, 0x70, 0xcd, 0x12, 0x3a, 0xb3, 0xf3, 0x2c,
	0x8e, 0x82, 0x52, 0x67, 0x17, 0x3f, 0x4d, 0x38, 0xff, 0x40, 0xf3, 0x3c, 0x4a, 0x43, 0x52, 0xc4,
	0xec, 0x73, 0x4a, 0x73, 0x71, 0x9d, 0x49, 0x44, 0xe8, 0xa7, 0x34, 0x61, 0x8e, 0x31, 0x37, 0x5c,
	0x8b, 0xa8, 0x18, 0xef, 0x03, 0xc8, 0x2c, 0x59, 0x09, 0x99, 0xa5, 0x6c, 0xed, 0xf4, 0xe6, 0x86,
	0x3b, 0x26, 0x3b, 0x19, 0x7c, 0x00, 0xd3, 0xa0, 0x90, 0xd9, 0x0d, 0xe3, 0x7e, 0x4a, 0xd3, 0x4c,
	0x38, 0xe6, 0xdc, 0x70, 0x4d, 0x62, 0xd7, 0xc9, 0x8f, 0x55, 0x0e, 0xdf, 0xc3, 0x44, 0xd2, 0xd0,
	0xff, 0x1e, 0xc5, 0x92, 0x71, 0xe1, 0xf4, 0xe7, 0xa6, 0x3b, 0x79, 0xf2, 0xc8, 0xd3, 0xe2, 0xbc,
	0x0e, 0x29, 0xde, 0x15, 0x0d, 0xdf, 0xea, 0xee, 0x37, 0xa9, 0xe4, 0x25, 0x01, 0xb9, 0x49, 0xe0,
	0x25, 0x8c, 0xd5, 0x3a, 0x11, 0x13, 0xce, 0x40, 0x8d, 0x3a, 0x69, 0x46, 0x7d, 0x52, 0x6b, 0x92,
	0x4d, 0x1d, 0x1f, 0xc3, 0x9d, 0x98, 0x0a, 0xe9, 0x17, 0xf9, 0x9a, 0x4a, 0xb6, 0xf6, 0xa9, 0xac,
	0x65, 0x0e, 0x95, 0x4c, 0xac, 0x8a, 0x5f, 0x74, 0xed, 0x95, 0xd4, 0x62, 0x1f, 0xc2, 0x69, 0xeb,
	0xc8, 0xaa, 0x74, 0x46, 0x0a, 0xc8, 0x74, 0xa7, 0xf9, 0x75, 0x39, 0x7b, 0x09, 0xa7, 0x7f, 0xa9,
	0xc4, 0x33, 0x30, 0x7f, 0xb0, 0xb2, 0xe6, 0x57, 0x85, 0x78, 0x01, 0x83, 0x1b, 0x1a, 0x17, 0x4c,
	0x91, 0xb3, 0x88, 0x7e, 0x78, 0xd1, 0x5b, 0x1a, 0x8b, 0x6f, 0x30, 0xd9, 0x59, 0xbc, 0x62, 0x5f,
	0x14, 0xd1, 0xba, 0x61, 0x5f, 0xc5, 0xf8, 0x1c, 0x2c, 0x51, 0x03, 0x11, 0x4e, 0x4f, 0x6d, 0x7a,
	0xf7, 0x08, 0x34, 0xb2, 0xed, 0x5e, 0xac, 0xc0, 0x26, 0x59, 0x1c, 0x17, 0xf9, 0x15, 0xe5, 0x21,
	0xeb, 0xbe, 0x5a, 0x84, 0xbe, 0xa4, 0xa1, 0x9e, 0x6c, 0x11, 0x15, 0xb7, 0xd8, 0x9a, 0xc7, 0xd9,
	0x2e, 0x7e, 0x99, 0x80, 0xfa, 0x25, 0xff, 0xc7, 0x45, 0xef, 0xba, 0x5c, 0x74, 0xd9, 0xc8, 0xdb,
	0x57, 0x72, 0xd4, 0x44, 0x1e, 0x8c, 0xa4, 0x42, 0xd3, 0x78, 0xe8, 0xa2, 0x3d, 0x48, 0x73, 0x23,
	0x4d, 0xd3, 0x2d, 0x36, 0xd2, 0x57, 0x80, 0xed, 0xee, 0x9d, 0x3e, 0x5a, 0xee, 0xfb, 0x68, 0x76,
	0x18, 0xdb, 0xae, 0x8d, 0x7e, 0xf7, 0x60, 0xa4, 0x6a, 0xda, 0x42, 0x7b, 0x93, 0xef, 0x81, 0x55,
	0xdd, 0xaf, 0xc8, 0x69, 0xd0, 0x28, 0xdb, 0x26, 0xd0, 0x85, 0xb3, 0x80, 0xb3, 0x36, 0x2e, 0x7d,
	0xb1, 0x27, 0x75, 0xbe, 0x41, 0x75, 0x90, 0x6e, 0xff, 0x20, 0xdd, 0xb6, 0xa5, 0x06, 0xff, 0xb6,
	0xd4, 0xb0, 0xc3, 0x52, 0x4b, 0x98, 0x26, 0xfa, 0x43, 0xf2, 0xab, 0xbf, 0xa6, 0x70, 0x46, 0x8a,
	0xce, 0x79, 0xc7, 0x57, 0x46, 0xec, 0x64, 0xfb, 0x20, 0xf0, 0x19, 0xd8, 0x5c, 0xa1, 0xab, 0x0f,
	0x8e, 0xd5, 0x41, 0xdc, 0xc7, 0x4a, 0x26, 0x7c, 0x13, 0x77, 0x7a, 0xc2, 0xea, 0xf0, 0xc4, 0x6a,
	0xa8, 0xfe, 0xd4, 0x4f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x03, 0x43, 0x52, 0xcd, 0x05,
	0x00, 0x00,
}
