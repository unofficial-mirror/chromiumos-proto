// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device/device_variant.proto

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

type DeviceVariant_Feature int32

const (
	DeviceVariant_FEATURE_UNSPECIFIED DeviceVariant_Feature = 0
	DeviceVariant_RUGGED              DeviceVariant_Feature = 1
	DeviceVariant_STYLUS              DeviceVariant_Feature = 2
	DeviceVariant_TOUCHPAD            DeviceVariant_Feature = 3
	DeviceVariant_TOUCHSCREEN         DeviceVariant_Feature = 4
	DeviceVariant_REAR_CAMERA         DeviceVariant_Feature = 5
	DeviceVariant_KEYBOARD_BACKLIGHT  DeviceVariant_Feature = 6
	DeviceVariant_BASE_ACCELEROMETER  DeviceVariant_Feature = 7
	DeviceVariant_BASE_GYROSCOPE      DeviceVariant_Feature = 8
	DeviceVariant_BASE_MAGNETOMETER   DeviceVariant_Feature = 9
	DeviceVariant_LID_ACCELEROMETER   DeviceVariant_Feature = 10
)

var DeviceVariant_Feature_name = map[int32]string{
	0:  "FEATURE_UNSPECIFIED",
	1:  "RUGGED",
	2:  "STYLUS",
	3:  "TOUCHPAD",
	4:  "TOUCHSCREEN",
	5:  "REAR_CAMERA",
	6:  "KEYBOARD_BACKLIGHT",
	7:  "BASE_ACCELEROMETER",
	8:  "BASE_GYROSCOPE",
	9:  "BASE_MAGNETOMETER",
	10: "LID_ACCELEROMETER",
}

var DeviceVariant_Feature_value = map[string]int32{
	"FEATURE_UNSPECIFIED": 0,
	"RUGGED":              1,
	"STYLUS":              2,
	"TOUCHPAD":            3,
	"TOUCHSCREEN":         4,
	"REAR_CAMERA":         5,
	"KEYBOARD_BACKLIGHT":  6,
	"BASE_ACCELEROMETER":  7,
	"BASE_GYROSCOPE":      8,
	"BASE_MAGNETOMETER":   9,
	"LID_ACCELEROMETER":   10,
}

func (x DeviceVariant_Feature) String() string {
	return proto.EnumName(DeviceVariant_Feature_name, int32(x))
}

func (DeviceVariant_Feature) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d254428042c205dd, []int{0, 0}
}

// These are the distinct feature combinations that will be manufactured
// for a given Device Model.
type DeviceVariant struct {
	// Required. Unique ID of the Variant.
	Id *DeviceVariantId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required. Parent Device Model for the given Variant.
	DeviceId *DeviceModelId `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Required. Features of the given Variant.
	Feature              []DeviceVariant_Feature `protobuf:"varint,3,rep,packed,name=feature,proto3,enum=device.DeviceVariant_Feature" json:"feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DeviceVariant) Reset()         { *m = DeviceVariant{} }
func (m *DeviceVariant) String() string { return proto.CompactTextString(m) }
func (*DeviceVariant) ProtoMessage()    {}
func (*DeviceVariant) Descriptor() ([]byte, []int) {
	return fileDescriptor_d254428042c205dd, []int{0}
}

func (m *DeviceVariant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceVariant.Unmarshal(m, b)
}
func (m *DeviceVariant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceVariant.Marshal(b, m, deterministic)
}
func (m *DeviceVariant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceVariant.Merge(m, src)
}
func (m *DeviceVariant) XXX_Size() int {
	return xxx_messageInfo_DeviceVariant.Size(m)
}
func (m *DeviceVariant) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceVariant.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceVariant proto.InternalMessageInfo

func (m *DeviceVariant) GetId() *DeviceVariantId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DeviceVariant) GetDeviceId() *DeviceModelId {
	if m != nil {
		return m.DeviceId
	}
	return nil
}

func (m *DeviceVariant) GetFeature() []DeviceVariant_Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

func init() {
	proto.RegisterEnum("device.DeviceVariant_Feature", DeviceVariant_Feature_name, DeviceVariant_Feature_value)
	proto.RegisterType((*DeviceVariant)(nil), "device.DeviceVariant")
}

func init() { proto.RegisterFile("device/device_variant.proto", fileDescriptor_d254428042c205dd) }

var fileDescriptor_d254428042c205dd = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0xae, 0xd2, 0x40,
	0x18, 0x86, 0x6d, 0xab, 0x85, 0x33, 0x47, 0x8f, 0xe3, 0x98, 0xe3, 0x21, 0xf8, 0x13, 0xc2, 0x46,
	0x56, 0xad, 0xa9, 0x0b, 0xd7, 0xd3, 0xe9, 0x47, 0x69, 0x68, 0x29, 0x99, 0xb6, 0x26, 0xb8, 0x69,
	0x2a, 0x53, 0x70, 0x12, 0x61, 0x4c, 0x05, 0xae, 0xc0, 0xcb, 0xf4, 0x62, 0x4c, 0x7f, 0x30, 0xe1,
	0x84, 0x55, 0xfb, 0x3d, 0xef, 0x33, 0xef, 0x4c, 0xf2, 0xa1, 0xb7, 0xa2, 0x3c, 0xc9, 0x75, 0x69,
	0xb7, 0x9f, 0xfc, 0x54, 0x54, 0xb2, 0xd8, 0x1f, 0xac, 0x5f, 0x95, 0x3a, 0x28, 0x62, 0xb6, 0x74,
	0xf8, 0xee, 0x52, 0xda, 0x29, 0x51, 0xfe, 0xcc, 0xa5, 0x68, 0xad, 0xe1, 0x87, 0xab, 0x15, 0xff,
	0xf3, 0xf1, 0x1f, 0x03, 0xbd, 0xf0, 0x9a, 0xec, 0x6b, 0x1b, 0x91, 0x8f, 0x48, 0x97, 0x62, 0xa0,
	0x8d, 0xb4, 0xc9, 0xad, 0xf3, 0x60, 0xb5, 0xe7, 0xac, 0x0b, 0x25, 0x10, 0x5c, 0x97, 0x82, 0x38,
	0xe8, 0xa6, 0x6b, 0x95, 0x62, 0xa0, 0x37, 0xfe, 0xfd, 0xa5, 0x1f, 0xd5, 0x6f, 0x09, 0x04, 0xef,
	0xb7, 0x34, 0x10, 0xe4, 0x0b, 0xea, 0x6d, 0xca, 0xe2, 0x70, 0xac, 0xca, 0x81, 0x31, 0x32, 0x26,
	0x77, 0xce, 0xfb, 0xab, 0x37, 0x58, 0xd3, 0x56, 0xe2, 0x67, 0x7b, 0xfc, 0x57, 0x43, 0xbd, 0x0e,
	0x92, 0x07, 0xf4, 0x7a, 0x0a, 0x34, 0xcd, 0x38, 0xe4, 0xd9, 0x22, 0x59, 0x02, 0x0b, 0xa6, 0x01,
	0x78, 0xf8, 0x09, 0x41, 0xc8, 0xe4, 0x99, 0xef, 0x83, 0x87, 0xb5, 0xfa, 0x3f, 0x49, 0x57, 0x61,
	0x96, 0x60, 0x9d, 0x3c, 0x47, 0xfd, 0x34, 0xce, 0xd8, 0x6c, 0x49, 0x3d, 0x6c, 0x90, 0x97, 0xe8,
	0xb6, 0x99, 0x12, 0xc6, 0x01, 0x16, 0xf8, 0x69, 0x0d, 0x38, 0x50, 0x9e, 0x33, 0x1a, 0x01, 0xa7,
	0xf8, 0x19, 0x79, 0x83, 0xc8, 0x1c, 0x56, 0x6e, 0x4c, 0xb9, 0x97, 0xbb, 0x94, 0xcd, 0xc3, 0xc0,
	0x9f, 0xa5, 0xd8, 0xac, 0xb9, 0x4b, 0x13, 0xc8, 0x29, 0x63, 0x10, 0x02, 0x8f, 0x23, 0x48, 0x81,
	0xe3, 0x1e, 0x21, 0xe8, 0xae, 0xe1, 0xfe, 0x8a, 0xc7, 0x09, 0x8b, 0x97, 0x80, 0xfb, 0xe4, 0x1e,
	0xbd, 0x6a, 0x58, 0x44, 0xfd, 0x05, 0xa4, 0x9d, 0x7a, 0x53, 0xe3, 0x30, 0xf0, 0x1e, 0x35, 0x20,
	0xd7, 0xf9, 0xf6, 0x69, 0xab, 0xac, 0xf5, 0x8f, 0x4a, 0xed, 0xe4, 0x71, 0x67, 0xa9, 0x6a, 0x6b,
	0x9f, 0x07, 0xf5, 0xdb, 0x96, 0xfb, 0x4d, 0x55, 0xd8, 0xcd, 0xc2, 0xec, 0xad, 0xea, 0x76, 0xf9,
	0xdd, 0x6c, 0xc0, 0xe7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0xa0, 0xc6, 0x9b, 0x26, 0x02,
	0x00, 0x00,
}