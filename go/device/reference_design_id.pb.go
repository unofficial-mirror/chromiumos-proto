// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device/reference_design_id.proto

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

// Globally unique identifier.
type ReferenceDesignId struct {
	// Required.
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReferenceDesignId) Reset()         { *m = ReferenceDesignId{} }
func (m *ReferenceDesignId) String() string { return proto.CompactTextString(m) }
func (*ReferenceDesignId) ProtoMessage()    {}
func (*ReferenceDesignId) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f8abec352e0f4a, []int{0}
}

func (m *ReferenceDesignId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReferenceDesignId.Unmarshal(m, b)
}
func (m *ReferenceDesignId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReferenceDesignId.Marshal(b, m, deterministic)
}
func (m *ReferenceDesignId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReferenceDesignId.Merge(m, src)
}
func (m *ReferenceDesignId) XXX_Size() int {
	return xxx_messageInfo_ReferenceDesignId.Size(m)
}
func (m *ReferenceDesignId) XXX_DiscardUnknown() {
	xxx_messageInfo_ReferenceDesignId.DiscardUnknown(m)
}

var xxx_messageInfo_ReferenceDesignId proto.InternalMessageInfo

func (m *ReferenceDesignId) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ReferenceDesignId)(nil), "device.ReferenceDesignId")
}

func init() { proto.RegisterFile("device/reference_design_id.proto", fileDescriptor_96f8abec352e0f4a) }

var fileDescriptor_96f8abec352e0f4a = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2f, 0x4a, 0x4d, 0x4b, 0x2d, 0x4a, 0xcd, 0x4b, 0x4e, 0x8d, 0x4f, 0x49, 0x2d,
	0xce, 0x4c, 0xcf, 0x8b, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xa8,
	0x50, 0xd2, 0xe4, 0x12, 0x0c, 0x82, 0x29, 0x72, 0x01, 0xab, 0xf1, 0x4c, 0x11, 0x12, 0xe1, 0x62,
	0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x9c, 0x8c,
	0xa2, 0x0c, 0xd2, 0xf3, 0xf5, 0x92, 0x33, 0x8a, 0xf2, 0x73, 0x33, 0x4b, 0x73, 0xf5, 0xf2, 0x8b,
	0xd2, 0xf5, 0x61, 0x9c, 0xfc, 0x62, 0xfd, 0xcc, 0xbc, 0xb4, 0xa2, 0x44, 0x7d, 0xb0, 0xe1, 0xfa,
	0xe9, 0xf9, 0xfa, 0x10, 0xe3, 0x93, 0xd8, 0xc0, 0x02, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x89, 0x2a, 0x2e, 0x85, 0x91, 0x00, 0x00, 0x00,
}
