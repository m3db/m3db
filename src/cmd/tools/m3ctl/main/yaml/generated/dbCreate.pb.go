// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dbCreate.proto

package yaml

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	admin "github.com/m3db/m3/src/query/generated/proto/admin"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// enum OperationType {
//CREATE = 0;
//INIT = 1;
//}
type DatabaseCreateRequestYaml struct {
	Operation            string                       `protobuf:"bytes,1,opt,name=Operation,proto3" json:"Operation,omitempty"`
	Request              *admin.DatabaseCreateRequest `protobuf:"bytes,2,opt,name=Request,proto3" json:"Request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DatabaseCreateRequestYaml) Reset()         { *m = DatabaseCreateRequestYaml{} }
func (m *DatabaseCreateRequestYaml) String() string { return proto.CompactTextString(m) }
func (*DatabaseCreateRequestYaml) ProtoMessage()    {}
func (*DatabaseCreateRequestYaml) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac8e05fdfec588ce, []int{0}
}

func (m *DatabaseCreateRequestYaml) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseCreateRequestYaml.Unmarshal(m, b)
}
func (m *DatabaseCreateRequestYaml) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseCreateRequestYaml.Marshal(b, m, deterministic)
}
func (m *DatabaseCreateRequestYaml) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseCreateRequestYaml.Merge(m, src)
}
func (m *DatabaseCreateRequestYaml) XXX_Size() int {
	return xxx_messageInfo_DatabaseCreateRequestYaml.Size(m)
}
func (m *DatabaseCreateRequestYaml) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseCreateRequestYaml.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseCreateRequestYaml proto.InternalMessageInfo

func (m *DatabaseCreateRequestYaml) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *DatabaseCreateRequestYaml) GetRequest() *admin.DatabaseCreateRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*DatabaseCreateRequestYaml)(nil), "yaml.DatabaseCreateRequestYaml")
}

func init() { proto.RegisterFile("dbCreate.proto", fileDescriptor_ac8e05fdfec588ce) }

var fileDescriptor_ac8e05fdfec588ce = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x49, 0x72, 0x2e,
	0x4a, 0x4d, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xa9, 0x4c, 0xcc, 0xcd,
	0x91, 0x72, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x35, 0x4e,
	0x49, 0xd2, 0xcf, 0x35, 0xd6, 0x2f, 0x2e, 0x4a, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x4f,
	0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0x49, 0x4d, 0xd1, 0x07, 0xeb, 0xd1, 0x4f, 0x4c, 0xc9, 0xcd,
	0xcc, 0xd3, 0x4f, 0x49, 0x2c, 0x49, 0x4c, 0x4a, 0x2c, 0x86, 0x1a, 0xa4, 0x54, 0xc8, 0x25, 0xe9,
	0x02, 0x15, 0x81, 0x58, 0x10, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x12, 0x99, 0x98, 0x9b, 0x23,
	0x24, 0xc3, 0xc5, 0xe9, 0x5f, 0x00, 0x32, 0x23, 0x33, 0x3f, 0x4f, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0x21, 0x20, 0x64, 0xc6, 0xc5, 0x0e, 0x55, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d,
	0x24, 0xa3, 0x07, 0xb6, 0x42, 0x0f, 0xab, 0x81, 0x41, 0x30, 0xc5, 0x49, 0x6c, 0x60, 0x9b, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xb1, 0x77, 0xc1, 0xd4, 0x00, 0x00, 0x00,
}
