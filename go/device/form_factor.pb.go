// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device/form_factor.proto

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

// For more details, see:
// https://www.google.com/chromeos/partner/fe/docs/latest-requirements/formfactors.html
type FormFactor int32

const (
	FormFactor_FORM_FACTOR_UNSPECIFIED FormFactor = 0
	FormFactor_CLAMSHELL               FormFactor = 1
	FormFactor_CONVERTIBLE             FormFactor = 2
	FormFactor_DETACHABLE              FormFactor = 3
	FormFactor_CHROMEBASE              FormFactor = 4
	FormFactor_CHROMEBOX               FormFactor = 5
	FormFactor_CHROMEBIT               FormFactor = 6
)

var FormFactor_name = map[int32]string{
	0: "FORM_FACTOR_UNSPECIFIED",
	1: "CLAMSHELL",
	2: "CONVERTIBLE",
	3: "DETACHABLE",
	4: "CHROMEBASE",
	5: "CHROMEBOX",
	6: "CHROMEBIT",
}

var FormFactor_value = map[string]int32{
	"FORM_FACTOR_UNSPECIFIED": 0,
	"CLAMSHELL":               1,
	"CONVERTIBLE":             2,
	"DETACHABLE":              3,
	"CHROMEBASE":              4,
	"CHROMEBOX":               5,
	"CHROMEBIT":               6,
}

func (x FormFactor) String() string {
	return proto.EnumName(FormFactor_name, int32(x))
}

func (FormFactor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_70c20bca674d535a, []int{0}
}

func init() {
	proto.RegisterEnum("device.FormFactor", FormFactor_name, FormFactor_value)
}

func init() { proto.RegisterFile("device/form_factor.proto", fileDescriptor_70c20bca674d535a) }

var fileDescriptor_70c20bca674d535a = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0x71, 0xd7, 0x3f, 0x05, 0x47, 0xd4, 0x90, 0x8b, 0x82, 0x6f, 0xe0, 0xa1, 0x11, 0x7d,
	0x82, 0x34, 0x3b, 0xa1, 0x81, 0x74, 0x23, 0x69, 0x14, 0xf1, 0x52, 0xd6, 0xba, 0xad, 0x3d, 0xc4,
	0x91, 0xb8, 0xfa, 0x0a, 0xbe, 0xb6, 0x6c, 0x8a, 0xe0, 0xf1, 0x33, 0x3f, 0x06, 0xbe, 0x70, 0xf9,
	0xba, 0xf9, 0x9e, 0xfa, 0x8d, 0x18, 0x28, 0xc5, 0x6e, 0x58, 0xf7, 0x5b, 0x4a, 0xe5, 0x47, 0xa2,
	0x2d, 0xf1, 0x62, 0x5e, 0xae, 0x7f, 0x16, 0x00, 0x9a, 0x52, 0xd4, 0x79, 0xe4, 0x57, 0x70, 0xa1,
	0x9d, 0x6f, 0x3a, 0x2d, 0x55, 0x70, 0xbe, 0x7b, 0x58, 0xb5, 0xf7, 0xa8, 0x8c, 0x36, 0xb8, 0x64,
	0x7b, 0xfc, 0x14, 0x8e, 0x95, 0x95, 0x4d, 0x5b, 0xa3, 0xb5, 0x6c, 0xc1, 0xcf, 0xe1, 0x44, 0xb9,
	0xd5, 0x23, 0xfa, 0x60, 0x2a, 0x8b, 0x6c, 0x9f, 0x9f, 0x01, 0x2c, 0x31, 0x48, 0x55, 0xcb, 0x9d,
	0x0f, 0x76, 0x56, 0xb5, 0x77, 0x0d, 0x56, 0xb2, 0x45, 0x76, 0x98, 0xff, 0x67, 0xbb, 0x27, 0x76,
	0xf4, 0x8f, 0x26, 0xb0, 0xa2, 0xba, 0x7d, 0xbe, 0x19, 0xa9, 0xec, 0xdf, 0x12, 0xc5, 0xe9, 0x2b,
	0x96, 0x94, 0x46, 0xf1, 0x07, 0xfa, 0x14, 0xd3, 0xfb, 0x90, 0xd6, 0x22, 0xb7, 0x8b, 0x91, 0xc4,
	0x5c, 0xff, 0x52, 0xe4, 0xc3, 0xdd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x3e, 0x8c, 0x58,
	0xe8, 0x00, 0x00, 0x00,
}