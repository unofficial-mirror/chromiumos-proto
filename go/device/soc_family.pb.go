// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device/soc_family.proto

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

// These are the major releases from chip manufacturers, loosely defined as a
// unique silicon die, which subsequently drive Reference Board designs and the
// overall product roadmap (e.g. Intel Kaby Lake, Rockchip RK3399).
type SocFamily struct {
	// Required. Unique ID of the SoC Family.
	Id *SocFamilyId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required. Common name used to refer to the family.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Required. ISA implemented by the SoC Family.
	Architecture         Architecture `protobuf:"varint,3,opt,name=architecture,proto3,enum=device.Architecture" json:"architecture,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SocFamily) Reset()         { *m = SocFamily{} }
func (m *SocFamily) String() string { return proto.CompactTextString(m) }
func (*SocFamily) ProtoMessage()    {}
func (*SocFamily) Descriptor() ([]byte, []int) {
	return fileDescriptor_6234c394af7e9cd3, []int{0}
}

func (m *SocFamily) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocFamily.Unmarshal(m, b)
}
func (m *SocFamily) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocFamily.Marshal(b, m, deterministic)
}
func (m *SocFamily) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocFamily.Merge(m, src)
}
func (m *SocFamily) XXX_Size() int {
	return xxx_messageInfo_SocFamily.Size(m)
}
func (m *SocFamily) XXX_DiscardUnknown() {
	xxx_messageInfo_SocFamily.DiscardUnknown(m)
}

var xxx_messageInfo_SocFamily proto.InternalMessageInfo

func (m *SocFamily) GetId() *SocFamilyId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SocFamily) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SocFamily) GetArchitecture() Architecture {
	if m != nil {
		return m.Architecture
	}
	return Architecture_ARCHITECTURE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*SocFamily)(nil), "device.SocFamily")
}

func init() { proto.RegisterFile("device/soc_family.proto", fileDescriptor_6234c394af7e9cd3) }

var fileDescriptor_6234c394af7e9cd3 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2f, 0xce, 0x4f, 0x8e, 0x4f, 0x4b, 0xcc, 0xcd, 0xcc, 0xa9, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x48, 0x48, 0x49, 0x42, 0x15, 0x24, 0x16, 0x25, 0x67, 0x64,
	0x96, 0xa4, 0x26, 0x97, 0x94, 0x16, 0xa5, 0x42, 0x94, 0x48, 0x49, 0x61, 0xe8, 0x8d, 0xcf, 0x4c,
	0x81, 0xc8, 0x29, 0xd5, 0x71, 0x71, 0x06, 0xe7, 0x27, 0xbb, 0x81, 0x45, 0x85, 0x94, 0xb9, 0x98,
	0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x84, 0xf5, 0x20, 0xba, 0xf4, 0xe0, 0xd2,
	0x9e, 0x29, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x05, 0x17, 0x0f, 0xb2, 0xbd, 0x12, 0xcc, 0x0a,
	0x8c, 0x1a, 0x7c, 0x46, 0x22, 0x30, 0x23, 0x1c, 0x91, 0xe4, 0x82, 0x50, 0x54, 0x3a, 0x19, 0x45,
	0x19, 0xa4, 0xe7, 0xeb, 0x25, 0x67, 0x14, 0xe5, 0xe7, 0x66, 0x96, 0xe6, 0xea, 0xe5, 0x17, 0xa5,
	0xeb, 0xc3, 0x38, 0xf9, 0xc5, 0xfa, 0x99, 0x79, 0x69, 0x45, 0x89, 0xfa, 0x60, 0x97, 0xea, 0xa7,
	0xe7, 0xeb, 0x43, 0x8c, 0x4b, 0x62, 0x03, 0x0b, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x05,
	0xa0, 0xce, 0xce, 0x14, 0x01, 0x00, 0x00,
}