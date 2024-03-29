// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device/partner.proto

package device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Partner_Role int32

const (
	Partner_ROLE_UNSPECIFIED Partner_Role = 0
	// Original equipment manufacturer
	Partner_OEM Partner_Role = 1
	// Original device manufacturer
	Partner_ODM Partner_Role = 2
	// System-on-a-chip manufacturer
	Partner_SOC_MANUFACTURER Partner_Role = 3
	// Component manufacturer
	Partner_COMPONENT_MANUFACTURER Partner_Role = 4
)

var Partner_Role_name = map[int32]string{
	0: "ROLE_UNSPECIFIED",
	1: "OEM",
	2: "ODM",
	3: "SOC_MANUFACTURER",
	4: "COMPONENT_MANUFACTURER",
}

var Partner_Role_value = map[string]int32{
	"ROLE_UNSPECIFIED":       0,
	"OEM":                    1,
	"ODM":                    2,
	"SOC_MANUFACTURER":       3,
	"COMPONENT_MANUFACTURER": 4,
}

func (x Partner_Role) String() string {
	return proto.EnumName(Partner_Role_name, int32(x))
}

func (Partner_Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_18c5c125885dd24f, []int{0, 0}
}

// Partner we interact with to develop devices.
type Partner struct {
	Id *PartnerId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required. Common name of the partner.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Role(s) the partner plays.
	Role                 []Partner_Role `protobuf:"varint,3,rep,packed,name=role,proto3,enum=device.Partner_Role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Partner) Reset()         { *m = Partner{} }
func (m *Partner) String() string { return proto.CompactTextString(m) }
func (*Partner) ProtoMessage()    {}
func (*Partner) Descriptor() ([]byte, []int) {
	return fileDescriptor_18c5c125885dd24f, []int{0}
}

func (m *Partner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Partner.Unmarshal(m, b)
}
func (m *Partner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Partner.Marshal(b, m, deterministic)
}
func (m *Partner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Partner.Merge(m, src)
}
func (m *Partner) XXX_Size() int {
	return xxx_messageInfo_Partner.Size(m)
}
func (m *Partner) XXX_DiscardUnknown() {
	xxx_messageInfo_Partner.DiscardUnknown(m)
}

var xxx_messageInfo_Partner proto.InternalMessageInfo

func (m *Partner) GetId() *PartnerId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Partner) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Partner) GetRole() []Partner_Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterEnum("device.Partner_Role", Partner_Role_name, Partner_Role_value)
	proto.RegisterType((*Partner)(nil), "device.Partner")
}

func init() { proto.RegisterFile("device/partner.proto", fileDescriptor_18c5c125885dd24f) }

var fileDescriptor_18c5c125885dd24f = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0xed, 0x0b, 0x1b, 0x46, 0x90, 0x18, 0x8a, 0x96, 0x9d, 0xea, 0x4e, 0x3d, 0x25, 0x52,
	0x3f, 0xc1, 0xec, 0x32, 0x28, 0xd8, 0xa6, 0x64, 0xeb, 0xc5, 0x4b, 0xad, 0x4b, 0xac, 0x81, 0xb5,
	0x19, 0x71, 0xfa, 0x29, 0xfd, 0x50, 0xb2, 0x44, 0x0f, 0xdb, 0xed, 0xff, 0xf2, 0xfb, 0xc3, 0xf3,
	0x80, 0x48, 0xc8, 0x6f, 0xb5, 0x95, 0x64, 0xdf, 0x99, 0xc3, 0x28, 0x0d, 0xde, 0x1b, 0x7d, 0xd0,
	0x68, 0xe2, 0xd2, 0xd9, 0xdd, 0x69, 0xdb, 0x2a, 0xe1, 0x80, 0xf9, 0x8f, 0x07, 0xa6, 0xb5, 0x0b,
	0xd1, 0x3d, 0xf0, 0x95, 0x88, 0xbd, 0xc4, 0x4b, 0xaf, 0xb2, 0x1b, 0xec, 0x16, 0xf8, 0xaf, 0x2c,
	0x04, 0xf7, 0x95, 0x40, 0x08, 0x84, 0x63, 0x37, 0xc8, 0xd8, 0x4f, 0xbc, 0xf4, 0x92, 0x5b, 0x8d,
	0x52, 0x10, 0x1a, 0xbd, 0x93, 0x71, 0x90, 0x04, 0xe9, 0x75, 0x16, 0x9d, 0x0d, 0x31, 0xd7, 0x3b,
	0xc9, 0x2d, 0x31, 0x7f, 0x05, 0xe1, 0xd1, 0xa1, 0x08, 0x40, 0xce, 0x9e, 0x69, 0xdb, 0x54, 0xeb,
	0x9a, 0xe6, 0xc5, 0xaa, 0xa0, 0x4b, 0x78, 0x81, 0xa6, 0x20, 0x60, 0xb4, 0x84, 0x9e, 0x15, 0xcb,
	0x12, 0xfa, 0x47, 0x6e, 0xcd, 0xf2, 0xb6, 0x5c, 0x54, 0xcd, 0x6a, 0x91, 0x6f, 0x1a, 0x4e, 0x39,
	0x0c, 0xd0, 0x0c, 0xdc, 0xe6, 0xac, 0xac, 0x59, 0x45, 0xab, 0xcd, 0x69, 0x17, 0x3e, 0x65, 0x2f,
	0x0f, 0xbd, 0xc6, 0xdb, 0x0f, 0xa3, 0x07, 0xf5, 0x35, 0x60, 0x6d, 0x7a, 0xf2, 0x6f, 0xf4, 0x27,
	0x51, 0xe3, 0xbb, 0xe9, 0x88, 0x7d, 0x9c, 0xf4, 0x9a, 0xb8, 0x43, 0xdf, 0x26, 0x36, 0x78, 0xfc,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x97, 0xf7, 0x67, 0x07, 0x42, 0x01, 0x00, 0x00,
}
