// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chromite/api/image.proto

package chromite_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "go.chromium.org/chromiumos/infra/proto/go/chromite/api"
	chromiumos "go.chromium.org/chromiumos/infra/proto/go/chromiumos"
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

// The types of images that can be built.
type Image_Type int32

const (
	Image_IMAGE_TYPE_UNDEFINED Image_Type = 0
	Image_BASE                 Image_Type = 1
	Image_DEV                  Image_Type = 2
	Image_TEST                 Image_Type = 3
)

var Image_Type_name = map[int32]string{
	0: "IMAGE_TYPE_UNDEFINED",
	1: "BASE",
	2: "DEV",
	3: "TEST",
}

var Image_Type_value = map[string]int32{
	"IMAGE_TYPE_UNDEFINED": 0,
	"BASE":                 1,
	"DEV":                  2,
	"TEST":                 3,
}

func (x Image_Type) String() string {
	return proto.EnumName(Image_Type_name, int32(x))
}

func (Image_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_abef4ffabc69f3ed, []int{0, 0}
}

// Image argument - encapsulate data about an image.
type Image struct {
	// Path to the image file.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The image type.
	Type                 Image_Type `protobuf:"varint,2,opt,name=type,proto3,enum=chromite.api.Image_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_abef4ffabc69f3ed, []int{0}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Image) GetType() Image_Type {
	if m != nil {
		return m.Type
	}
	return Image_IMAGE_TYPE_UNDEFINED
}

// The image test arguments.
type CreateImageRequest struct {
	// The build target whose image is being built.
	BuildTarget *chromiumos.BuildTarget `protobuf:"bytes,1,opt,name=build_target,json=buildTarget,proto3" json:"build_target,omitempty"`
	// The types of images to build, defaults to building base image.
	ImageTypes []Image_Type `protobuf:"varint,2,rep,packed,name=image_types,json=imageTypes,proto3,enum=chromite.api.Image_Type" json:"image_types,omitempty"`
	// Whether rootfs verification should be disabled (enabled by default).
	DisableRootfsVerification bool `protobuf:"varint,3,opt,name=disable_rootfs_verification,json=disableRootfsVerification,proto3" json:"disable_rootfs_verification,omitempty"`
	// The image version.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Disk layout option. See README.disk_layout and legacy_disk_layout.json in
	// src/scripts/build_library.
	DiskLayout string `protobuf:"bytes,5,opt,name=disk_layout,json=diskLayout,proto3" json:"disk_layout,omitempty"`
	// Used to set the LSB builder path key in /etc/lsb-release. See
	// chromite/scripts/cros_set_lsb_release.py.
	BuilderPath          string   `protobuf:"bytes,6,opt,name=builder_path,json=builderPath,proto3" json:"builder_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImageRequest) Reset()         { *m = CreateImageRequest{} }
func (m *CreateImageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateImageRequest) ProtoMessage()    {}
func (*CreateImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abef4ffabc69f3ed, []int{1}
}

func (m *CreateImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageRequest.Unmarshal(m, b)
}
func (m *CreateImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageRequest.Marshal(b, m, deterministic)
}
func (m *CreateImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageRequest.Merge(m, src)
}
func (m *CreateImageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateImageRequest.Size(m)
}
func (m *CreateImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageRequest proto.InternalMessageInfo

func (m *CreateImageRequest) GetBuildTarget() *chromiumos.BuildTarget {
	if m != nil {
		return m.BuildTarget
	}
	return nil
}

func (m *CreateImageRequest) GetImageTypes() []Image_Type {
	if m != nil {
		return m.ImageTypes
	}
	return nil
}

func (m *CreateImageRequest) GetDisableRootfsVerification() bool {
	if m != nil {
		return m.DisableRootfsVerification
	}
	return false
}

func (m *CreateImageRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CreateImageRequest) GetDiskLayout() string {
	if m != nil {
		return m.DiskLayout
	}
	return ""
}

func (m *CreateImageRequest) GetBuilderPath() string {
	if m != nil {
		return m.BuilderPath
	}
	return ""
}

type CreateImageResult struct {
	// Whether it completed successfully.
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Images that were built. Will contain no more than one per image type.
	Images               []*Image `protobuf:"bytes,2,rep,name=images,proto3" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImageResult) Reset()         { *m = CreateImageResult{} }
func (m *CreateImageResult) String() string { return proto.CompactTextString(m) }
func (*CreateImageResult) ProtoMessage()    {}
func (*CreateImageResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_abef4ffabc69f3ed, []int{2}
}

func (m *CreateImageResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageResult.Unmarshal(m, b)
}
func (m *CreateImageResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageResult.Marshal(b, m, deterministic)
}
func (m *CreateImageResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageResult.Merge(m, src)
}
func (m *CreateImageResult) XXX_Size() int {
	return xxx_messageInfo_CreateImageResult.Size(m)
}
func (m *CreateImageResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageResult.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageResult proto.InternalMessageInfo

func (m *CreateImageResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreateImageResult) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

// The image test arguments.
type TestImageRequest struct {
	// The image to be tested.
	Image *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// The build target whose image is being tested.
	BuildTarget          *chromiumos.BuildTarget  `protobuf:"bytes,2,opt,name=build_target,json=buildTarget,proto3" json:"build_target,omitempty"`
	Result               *TestImageRequest_Result `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TestImageRequest) Reset()         { *m = TestImageRequest{} }
func (m *TestImageRequest) String() string { return proto.CompactTextString(m) }
func (*TestImageRequest) ProtoMessage()    {}
func (*TestImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abef4ffabc69f3ed, []int{3}
}

func (m *TestImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestImageRequest.Unmarshal(m, b)
}
func (m *TestImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestImageRequest.Marshal(b, m, deterministic)
}
func (m *TestImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestImageRequest.Merge(m, src)
}
func (m *TestImageRequest) XXX_Size() int {
	return xxx_messageInfo_TestImageRequest.Size(m)
}
func (m *TestImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestImageRequest proto.InternalMessageInfo

func (m *TestImageRequest) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *TestImageRequest) GetBuildTarget() *chromiumos.BuildTarget {
	if m != nil {
		return m.BuildTarget
	}
	return nil
}

func (m *TestImageRequest) GetResult() *TestImageRequest_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

// Test results options specifications.
type TestImageRequest_Result struct {
	// Location where the test results should be written.
	Directory            string   `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestImageRequest_Result) Reset()         { *m = TestImageRequest_Result{} }
func (m *TestImageRequest_Result) String() string { return proto.CompactTextString(m) }
func (*TestImageRequest_Result) ProtoMessage()    {}
func (*TestImageRequest_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_abef4ffabc69f3ed, []int{3, 0}
}

func (m *TestImageRequest_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestImageRequest_Result.Unmarshal(m, b)
}
func (m *TestImageRequest_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestImageRequest_Result.Marshal(b, m, deterministic)
}
func (m *TestImageRequest_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestImageRequest_Result.Merge(m, src)
}
func (m *TestImageRequest_Result) XXX_Size() int {
	return xxx_messageInfo_TestImageRequest_Result.Size(m)
}
func (m *TestImageRequest_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_TestImageRequest_Result.DiscardUnknown(m)
}

var xxx_messageInfo_TestImageRequest_Result proto.InternalMessageInfo

func (m *TestImageRequest_Result) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

type TestImageResult struct {
	// Whether all tests passed.
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestImageResult) Reset()         { *m = TestImageResult{} }
func (m *TestImageResult) String() string { return proto.CompactTextString(m) }
func (*TestImageResult) ProtoMessage()    {}
func (*TestImageResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_abef4ffabc69f3ed, []int{4}
}

func (m *TestImageResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestImageResult.Unmarshal(m, b)
}
func (m *TestImageResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestImageResult.Marshal(b, m, deterministic)
}
func (m *TestImageResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestImageResult.Merge(m, src)
}
func (m *TestImageResult) XXX_Size() int {
	return xxx_messageInfo_TestImageResult.Size(m)
}
func (m *TestImageResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TestImageResult.DiscardUnknown(m)
}

var xxx_messageInfo_TestImageResult proto.InternalMessageInfo

func (m *TestImageResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterEnum("chromite.api.Image_Type", Image_Type_name, Image_Type_value)
	proto.RegisterType((*Image)(nil), "chromite.api.Image")
	proto.RegisterType((*CreateImageRequest)(nil), "chromite.api.CreateImageRequest")
	proto.RegisterType((*CreateImageResult)(nil), "chromite.api.CreateImageResult")
	proto.RegisterType((*TestImageRequest)(nil), "chromite.api.TestImageRequest")
	proto.RegisterType((*TestImageRequest_Result)(nil), "chromite.api.TestImageRequest.Result")
	proto.RegisterType((*TestImageResult)(nil), "chromite.api.TestImageResult")
}

func init() { proto.RegisterFile("chromite/api/image.proto", fileDescriptor_abef4ffabc69f3ed) }

var fileDescriptor_abef4ffabc69f3ed = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xe1, 0x6e, 0xd3, 0x30,
	0x10, 0x26, 0x69, 0x96, 0x76, 0x97, 0x0a, 0x82, 0x41, 0x5a, 0x08, 0x83, 0x95, 0x48, 0xa0, 0xa2,
	0xa1, 0x54, 0x2a, 0xbf, 0x40, 0x1a, 0xd2, 0x46, 0x03, 0xaa, 0xc4, 0xa6, 0x29, 0x0d, 0x93, 0xe0,
	0x4f, 0xe4, 0xa6, 0xde, 0x6a, 0xd1, 0xce, 0xc1, 0x76, 0x2a, 0xf5, 0x0d, 0x78, 0x15, 0x9e, 0x02,
	0x69, 0x4f, 0xc2, 0x0b, 0xec, 0x1d, 0x90, 0x9d, 0x54, 0x6d, 0xd9, 0xa8, 0xc4, 0x3f, 0xdf, 0x7d,
	0x9f, 0xef, 0xee, 0xfb, 0xce, 0x06, 0x2f, 0x1b, 0x73, 0x36, 0xa5, 0x92, 0x74, 0x70, 0x4e, 0x3b,
	0x74, 0x8a, 0x2f, 0x48, 0x98, 0x73, 0x26, 0x19, 0x6a, 0x2e, 0x90, 0x10, 0xe7, 0xd4, 0xdf, 0x5d,
	0xe3, 0x0d, 0x0b, 0x3a, 0x19, 0xa5, 0x38, 0xa7, 0x25, 0xd7, 0xdf, 0x29, 0xd1, 0x62, 0xca, 0x44,
	0x27, 0x63, 0xd3, 0x29, 0xbb, 0x2c, 0x81, 0xe0, 0x87, 0x01, 0x5b, 0x7d, 0x55, 0x14, 0x21, 0xb0,
	0x72, 0x2c, 0xc7, 0x9e, 0xd1, 0x32, 0xda, 0xdb, 0xb1, 0x3e, 0xa3, 0x57, 0x60, 0xc9, 0x79, 0x4e,
	0x3c, 0xb3, 0x65, 0xb4, 0xef, 0x76, 0xbd, 0x70, 0xb5, 0x63, 0xa8, 0xaf, 0x85, 0xc9, 0x3c, 0x27,
	0xb1, 0x66, 0x05, 0x07, 0x60, 0xa9, 0x08, 0x79, 0xf0, 0xb0, 0x7f, 0x7c, 0xf8, 0x31, 0x4a, 0x93,
	0x2f, 0xa7, 0x51, 0xfa, 0xf9, 0xa4, 0x17, 0x7d, 0xe8, 0x9f, 0x44, 0x3d, 0xf7, 0x0e, 0x6a, 0x80,
	0x75, 0x74, 0x38, 0x88, 0x5c, 0x03, 0xd5, 0xa1, 0xd6, 0x8b, 0xce, 0x5c, 0x53, 0xa5, 0x92, 0x68,
	0x90, 0xb8, 0xb5, 0xe0, 0xa7, 0x09, 0xe8, 0x3d, 0x27, 0x58, 0x12, 0x5d, 0x39, 0x26, 0xdf, 0x0b,
	0x22, 0x24, 0x7a, 0x0b, 0xcd, 0x52, 0x8d, 0xc4, 0xfc, 0x82, 0x48, 0x3d, 0x9f, 0xd3, 0xdd, 0x09,
	0x97, 0x8a, 0xc2, 0x23, 0x85, 0x27, 0x1a, 0x8e, 0x9d, 0xe1, 0x32, 0x40, 0x6f, 0xc0, 0xd1, 0x8e,
	0xa5, 0x6a, 0x3e, 0xe1, 0x99, 0xad, 0xda, 0x46, 0x19, 0xa0, 0xc9, 0xea, 0x28, 0xd0, 0x3b, 0x78,
	0x3c, 0xa2, 0x02, 0x0f, 0x27, 0x24, 0xe5, 0x8c, 0xc9, 0x73, 0x91, 0xce, 0x08, 0xa7, 0xe7, 0x34,
	0xc3, 0x92, 0xb2, 0x4b, 0xaf, 0xd6, 0x32, 0xda, 0x8d, 0xf8, 0x51, 0x45, 0x89, 0x35, 0xe3, 0x6c,
	0x85, 0x80, 0x3c, 0xa8, 0xcf, 0x08, 0x17, 0x8a, 0x6b, 0x69, 0x47, 0x17, 0x21, 0xda, 0x03, 0x67,
	0x44, 0xc5, 0xb7, 0x74, 0x82, 0xe7, 0xac, 0x90, 0xde, 0x96, 0x46, 0x41, 0xa5, 0x3e, 0xe9, 0x0c,
	0x7a, 0x56, 0x29, 0x26, 0x3c, 0xd5, 0x1b, 0xb1, 0x35, 0xc3, 0xa9, 0x72, 0xa7, 0x58, 0x8e, 0x83,
	0xaf, 0x70, 0x7f, 0xcd, 0x2a, 0x51, 0x4c, 0xa4, 0x6a, 0x29, 0x8a, 0x2c, 0x23, 0x42, 0x68, 0x93,
	0x1a, 0xf1, 0x22, 0x44, 0xfb, 0x60, 0x6b, 0x69, 0xa5, 0x05, 0x4e, 0xf7, 0xc1, 0x2d, 0x16, 0xc4,
	0x15, 0x25, 0xf8, 0x6d, 0x80, 0x9b, 0x10, 0x21, 0xd7, 0xb6, 0xf0, 0x12, 0xb6, 0x34, 0x5c, 0xd9,
	0x7f, 0x6b, 0x81, 0x92, 0x71, 0x63, 0x61, 0xe6, 0x7f, 0x2c, 0xec, 0x00, 0x6c, 0xae, 0xc5, 0x68,
	0x83, 0x9d, 0xee, 0xf3, 0xf5, 0x3e, 0x7f, 0x8f, 0x15, 0x96, 0xca, 0xe3, 0xea, 0x92, 0xff, 0x02,
	0xec, 0xca, 0x8b, 0x5d, 0xd8, 0x1e, 0x51, 0x4e, 0x32, 0xc9, 0xf8, 0xbc, 0x7a, 0xd2, 0xcb, 0x44,
	0xb0, 0x0f, 0xf7, 0x56, 0x4a, 0x6d, 0x36, 0xaf, 0xfb, 0xcb, 0x80, 0xa6, 0x66, 0x0e, 0x08, 0x9f,
	0xd1, 0x8c, 0xa0, 0x01, 0xd8, 0xa5, 0xf9, 0xa8, 0xb5, 0x3e, 0xde, 0xcd, 0xd7, 0xeb, 0xef, 0x6d,
	0x60, 0xa8, 0xbe, 0x81, 0x7d, 0x75, 0xed, 0x9b, 0xae, 0x81, 0x8e, 0xc1, 0x52, 0x23, 0xa1, 0xa7,
	0x9b, 0x15, 0xfb, 0x4f, 0xfe, 0x89, 0xaf, 0x96, 0xf3, 0x9d, 0xab, 0x6b, 0xbf, 0x5e, 0xed, 0x6c,
	0x68, 0xeb, 0xbf, 0xfe, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x41, 0x86, 0xf7, 0x4c,
	0x04, 0x00, 0x00,
}
