// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chromite/api/sdk.proto

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

// Contains some common cros_sdk path arguments.
type ChrootPaths struct {
	// Override for the cache directory.
	Cache string `protobuf:"bytes,1,opt,name=cache,proto3" json:"cache,omitempty"`
	// Set the chrome root.
	Chrome string `protobuf:"bytes,2,opt,name=chrome,proto3" json:"chrome,omitempty"`
	// The absolute chroot path.
	Chroot               string   `protobuf:"bytes,3,opt,name=chroot,proto3" json:"chroot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChrootPaths) Reset()         { *m = ChrootPaths{} }
func (m *ChrootPaths) String() string { return proto.CompactTextString(m) }
func (*ChrootPaths) ProtoMessage()    {}
func (*ChrootPaths) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d6bdc7c0e50b1e, []int{0}
}

func (m *ChrootPaths) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChrootPaths.Unmarshal(m, b)
}
func (m *ChrootPaths) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChrootPaths.Marshal(b, m, deterministic)
}
func (m *ChrootPaths) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChrootPaths.Merge(m, src)
}
func (m *ChrootPaths) XXX_Size() int {
	return xxx_messageInfo_ChrootPaths.Size(m)
}
func (m *ChrootPaths) XXX_DiscardUnknown() {
	xxx_messageInfo_ChrootPaths.DiscardUnknown(m)
}

var xxx_messageInfo_ChrootPaths proto.InternalMessageInfo

func (m *ChrootPaths) GetCache() string {
	if m != nil {
		return m.Cache
	}
	return ""
}

func (m *ChrootPaths) GetChrome() string {
	if m != nil {
		return m.Chrome
	}
	return ""
}

func (m *ChrootPaths) GetChroot() string {
	if m != nil {
		return m.Chroot
	}
	return ""
}

// The chroot version information.
type ChrootVersion struct {
	// The version number.
	Version              uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChrootVersion) Reset()         { *m = ChrootVersion{} }
func (m *ChrootVersion) String() string { return proto.CompactTextString(m) }
func (*ChrootVersion) ProtoMessage()    {}
func (*ChrootVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d6bdc7c0e50b1e, []int{1}
}

func (m *ChrootVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChrootVersion.Unmarshal(m, b)
}
func (m *ChrootVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChrootVersion.Marshal(b, m, deterministic)
}
func (m *ChrootVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChrootVersion.Merge(m, src)
}
func (m *ChrootVersion) XXX_Size() int {
	return xxx_messageInfo_ChrootVersion.Size(m)
}
func (m *ChrootVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ChrootVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ChrootVersion proto.InternalMessageInfo

func (m *ChrootVersion) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// Create request message.
type CreateRequest struct {
	// The chroot-create flag arguments.
	Flags *CreateRequest_Flags `protobuf:"bytes,1,opt,name=flags,proto3" json:"flags,omitempty"`
	// The chroot path arguments.
	Paths                *ChrootPaths `protobuf:"bytes,2,opt,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d6bdc7c0e50b1e, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetFlags() *CreateRequest_Flags {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *CreateRequest) GetPaths() *ChrootPaths {
	if m != nil {
		return m.Paths
	}
	return nil
}

// Options that affect how the chroot is created.
type CreateRequest_Flags struct {
	// Whether or not to replace the chroot if it already exists.
	NoReplace bool `protobuf:"varint,1,opt,name=no_replace,json=noReplace,proto3" json:"no_replace,omitempty"`
	// Whether to do a full build of the SDK or use prebuilts.
	Bootstrap bool `protobuf:"varint,2,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	// Whether the chroot should be mounted on a loopback image or created
	// directly inside a directory. Set to true to create in a directory.
	NoUseImage           bool     `protobuf:"varint,3,opt,name=no_use_image,json=noUseImage,proto3" json:"no_use_image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest_Flags) Reset()         { *m = CreateRequest_Flags{} }
func (m *CreateRequest_Flags) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_Flags) ProtoMessage()    {}
func (*CreateRequest_Flags) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d6bdc7c0e50b1e, []int{2, 0}
}

func (m *CreateRequest_Flags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_Flags.Unmarshal(m, b)
}
func (m *CreateRequest_Flags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_Flags.Marshal(b, m, deterministic)
}
func (m *CreateRequest_Flags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_Flags.Merge(m, src)
}
func (m *CreateRequest_Flags) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_Flags.Size(m)
}
func (m *CreateRequest_Flags) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_Flags.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_Flags proto.InternalMessageInfo

func (m *CreateRequest_Flags) GetNoReplace() bool {
	if m != nil {
		return m.NoReplace
	}
	return false
}

func (m *CreateRequest_Flags) GetBootstrap() bool {
	if m != nil {
		return m.Bootstrap
	}
	return false
}

func (m *CreateRequest_Flags) GetNoUseImage() bool {
	if m != nil {
		return m.NoUseImage
	}
	return false
}

// Create response message.
type CreateResponse struct {
	// The resulting chroot version.
	Version              *ChrootVersion `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d6bdc7c0e50b1e, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetVersion() *ChrootVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

// Chroot update request message.
// Example json:
// {"toolchain_targets": [{"name": "eve"}]}
type UpdateRequest struct {
	// The flags.
	Flags *UpdateRequest_Flags `protobuf:"bytes,1,opt,name=flags,proto3" json:"flags,omitempty"`
	// The targets whose toolchains should be updated in the chroot.
	ToolchainTargets     []*chromiumos.BuildTarget `protobuf:"bytes,2,rep,name=toolchain_targets,json=toolchainTargets,proto3" json:"toolchain_targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d6bdc7c0e50b1e, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetFlags() *UpdateRequest_Flags {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *UpdateRequest) GetToolchainTargets() []*chromiumos.BuildTarget {
	if m != nil {
		return m.ToolchainTargets
	}
	return nil
}

// Update flag arguments.
type UpdateRequest_Flags struct {
	// Whether to build from source or use prebuilt packages.
	BuildSource          bool     `protobuf:"varint,1,opt,name=build_source,json=buildSource,proto3" json:"build_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest_Flags) Reset()         { *m = UpdateRequest_Flags{} }
func (m *UpdateRequest_Flags) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest_Flags) ProtoMessage()    {}
func (*UpdateRequest_Flags) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d6bdc7c0e50b1e, []int{4, 0}
}

func (m *UpdateRequest_Flags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest_Flags.Unmarshal(m, b)
}
func (m *UpdateRequest_Flags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest_Flags.Marshal(b, m, deterministic)
}
func (m *UpdateRequest_Flags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest_Flags.Merge(m, src)
}
func (m *UpdateRequest_Flags) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest_Flags.Size(m)
}
func (m *UpdateRequest_Flags) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest_Flags.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest_Flags proto.InternalMessageInfo

func (m *UpdateRequest_Flags) GetBuildSource() bool {
	if m != nil {
		return m.BuildSource
	}
	return false
}

// Chroot update response message.
type UpdateResponse struct {
	// The chroot version after update is complete.
	Version              *ChrootVersion `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d6bdc7c0e50b1e, []int{5}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetVersion() *ChrootVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func init() {
	proto.RegisterType((*ChrootPaths)(nil), "chromite.api.ChrootPaths")
	proto.RegisterType((*ChrootVersion)(nil), "chromite.api.ChrootVersion")
	proto.RegisterType((*CreateRequest)(nil), "chromite.api.CreateRequest")
	proto.RegisterType((*CreateRequest_Flags)(nil), "chromite.api.CreateRequest.Flags")
	proto.RegisterType((*CreateResponse)(nil), "chromite.api.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "chromite.api.UpdateRequest")
	proto.RegisterType((*UpdateRequest_Flags)(nil), "chromite.api.UpdateRequest.Flags")
	proto.RegisterType((*UpdateResponse)(nil), "chromite.api.UpdateResponse")
}

func init() { proto.RegisterFile("chromite/api/sdk.proto", fileDescriptor_b3d6bdc7c0e50b1e) }

var fileDescriptor_b3d6bdc7c0e50b1e = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x44, 0x49, 0x93, 0x71, 0x82, 0xc2, 0x0a, 0xb5, 0xc1, 0x0d, 0x52, 0xeb, 0x53,
	0xe1, 0xe0, 0x48, 0x46, 0x88, 0x3b, 0x41, 0x20, 0xc4, 0x05, 0x6d, 0x28, 0x57, 0x6b, 0x63, 0x6f,
	0xe3, 0x55, 0x62, 0xcf, 0xb2, 0xbb, 0xee, 0x4b, 0xf1, 0x06, 0x1c, 0xfb, 0x26, 0x3c, 0x40, 0xdf,
	0x01, 0x79, 0x37, 0xae, 0x71, 0x54, 0x21, 0x24, 0x6e, 0xde, 0x7f, 0xfe, 0x19, 0xfd, 0xf3, 0x69,
	0x0c, 0xa7, 0x69, 0xae, 0xb0, 0x10, 0x86, 0x2f, 0x99, 0x14, 0x4b, 0x9d, 0xed, 0x22, 0xa9, 0xd0,
	0x20, 0x99, 0x34, 0x7a, 0xc4, 0xa4, 0x08, 0x16, 0x1d, 0xd7, 0xa6, 0x12, 0xfb, 0x2c, 0x61, 0x52,
	0x38, 0x6f, 0x70, 0xe6, 0xaa, 0x55, 0x81, 0x7a, 0x99, 0x62, 0x51, 0x60, 0xe9, 0x0a, 0xe1, 0x1a,
	0xfc, 0x55, 0xae, 0x10, 0xcd, 0x17, 0x66, 0x72, 0x4d, 0x9e, 0xc1, 0x20, 0x65, 0x69, 0xce, 0xe7,
	0xde, 0x85, 0x77, 0x35, 0xa6, 0xee, 0x41, 0x4e, 0x61, 0x68, 0xfb, 0xf9, 0xbc, 0x67, 0xe5, 0xc3,
	0xab, 0xd1, 0xd1, 0xcc, 0xfb, 0xad, 0x8e, 0x26, 0x7c, 0x09, 0x53, 0x37, 0xf4, 0x1b, 0x57, 0x5a,
	0x60, 0x49, 0xe6, 0x70, 0x72, 0xeb, 0x3e, 0xed, 0xe0, 0x29, 0x6d, 0x9e, 0xe1, 0x2f, 0x0f, 0xa6,
	0x2b, 0xc5, 0x99, 0xe1, 0x94, 0x7f, 0xaf, 0xb8, 0x36, 0xe4, 0x2d, 0x0c, 0x6e, 0xf6, 0x6c, 0xab,
	0xad, 0xd3, 0x8f, 0x2f, 0xa3, 0x3f, 0xd7, 0x8c, 0x3a, 0xde, 0xe8, 0x43, 0x6d, 0xa4, 0xce, 0x4f,
	0x96, 0x30, 0x90, 0xf5, 0x12, 0x36, 0xa4, 0x1f, 0x3f, 0x3f, 0x6a, 0x6c, 0xb7, 0xa4, 0xce, 0x17,
	0xdc, 0xc0, 0xc0, 0x0e, 0x20, 0x2f, 0x00, 0x4a, 0x4c, 0x14, 0x97, 0x7b, 0x96, 0xba, 0xd5, 0x47,
	0x74, 0x5c, 0x22, 0x75, 0x02, 0x59, 0xc0, 0x78, 0x83, 0x68, 0xb4, 0x51, 0x4c, 0xda, 0xe1, 0x23,
	0xda, 0x0a, 0xe4, 0x02, 0x26, 0x25, 0x26, 0x95, 0xe6, 0x89, 0x28, 0xd8, 0x96, 0x5b, 0x14, 0x23,
	0x0a, 0x25, 0x5e, 0x6b, 0xfe, 0xa9, 0x56, 0xc2, 0x8f, 0xf0, 0xa4, 0x89, 0xad, 0x25, 0x96, 0x9a,
	0x93, 0x37, 0x5d, 0x1e, 0x7e, 0x7c, 0xfe, 0x58, 0xd8, 0x03, 0xbd, 0x16, 0xd6, 0x4f, 0x0f, 0xa6,
	0xd7, 0x32, 0xfb, 0x67, 0x58, 0x1d, 0x6f, 0x17, 0xd6, 0x7b, 0x78, 0x6a, 0x10, 0xf7, 0x69, 0xce,
	0x44, 0x99, 0x18, 0xa6, 0xb6, 0xdc, 0xd4, 0xe0, 0xfa, 0x57, 0x7e, 0x7c, 0x16, 0xb5, 0xc7, 0x12,
	0xbd, 0xab, 0x0f, 0xe9, 0xab, 0xad, 0xd3, 0xd9, 0x43, 0x87, 0x13, 0x74, 0xf0, 0xaa, 0x21, 0x78,
	0x09, 0x13, 0x77, 0x72, 0x1a, 0x2b, 0xf5, 0xc0, 0xd0, 0xb7, 0xda, 0xda, 0x4a, 0x35, 0x85, 0x26,
	0xcf, 0x7f, 0x51, 0x88, 0x7f, 0x78, 0x00, 0xeb, 0x6c, 0xb7, 0xe6, 0xea, 0x56, 0xa4, 0x9c, 0xac,
	0x60, 0xe8, 0xe8, 0x92, 0xf3, 0xbf, 0x9c, 0x4a, 0xb0, 0x78, 0xbc, 0x78, 0x88, 0xf2, 0x19, 0x86,
	0x2e, 0xdc, 0xf1, 0x90, 0x0e, 0xc2, 0xe3, 0x21, 0xdd, 0x7d, 0xc2, 0xe1, 0xdd, 0x7d, 0xd0, 0x9b,
	0x79, 0x81, 0x7f, 0x77, 0x1f, 0x9c, 0x40, 0x5f, 0x67, 0xbb, 0x59, 0x6f, 0x33, 0xb4, 0xff, 0xd9,
	0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x7e, 0x17, 0xc0, 0xc6, 0x03, 0x00, 0x00,
}
