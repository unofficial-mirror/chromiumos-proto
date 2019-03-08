// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chromite/api/test_archive.proto

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

// Message for creating archives.
// Example Json:
// {"build_target":{"name":"board"},"output_directory":"/tmp/archives"}
type CreateArchiveRequest struct {
	// The board whose tarballs are being created.
	BuildTarget *chromiumos.BuildTarget `protobuf:"bytes,1,opt,name=build_target,json=buildTarget,proto3" json:"build_target,omitempty"`
	// The directory where the completed archive(s) should be saved.
	OutputDirectory      string   `protobuf:"bytes,2,opt,name=output_directory,json=outputDirectory,proto3" json:"output_directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateArchiveRequest) Reset()         { *m = CreateArchiveRequest{} }
func (m *CreateArchiveRequest) String() string { return proto.CompactTextString(m) }
func (*CreateArchiveRequest) ProtoMessage()    {}
func (*CreateArchiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e2bd7d1e1e8889, []int{0}
}

func (m *CreateArchiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArchiveRequest.Unmarshal(m, b)
}
func (m *CreateArchiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArchiveRequest.Marshal(b, m, deterministic)
}
func (m *CreateArchiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArchiveRequest.Merge(m, src)
}
func (m *CreateArchiveRequest) XXX_Size() int {
	return xxx_messageInfo_CreateArchiveRequest.Size(m)
}
func (m *CreateArchiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArchiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArchiveRequest proto.InternalMessageInfo

func (m *CreateArchiveRequest) GetBuildTarget() *chromiumos.BuildTarget {
	if m != nil {
		return m.BuildTarget
	}
	return nil
}

func (m *CreateArchiveRequest) GetOutputDirectory() string {
	if m != nil {
		return m.OutputDirectory
	}
	return ""
}

// Information about an archive file.
type ArchiveFile struct {
	// File path.
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArchiveFile) Reset()         { *m = ArchiveFile{} }
func (m *ArchiveFile) String() string { return proto.CompactTextString(m) }
func (*ArchiveFile) ProtoMessage()    {}
func (*ArchiveFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e2bd7d1e1e8889, []int{1}
}

func (m *ArchiveFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveFile.Unmarshal(m, b)
}
func (m *ArchiveFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveFile.Marshal(b, m, deterministic)
}
func (m *ArchiveFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveFile.Merge(m, src)
}
func (m *ArchiveFile) XXX_Size() int {
	return xxx_messageInfo_ArchiveFile.Size(m)
}
func (m *ArchiveFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveFile.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveFile proto.InternalMessageInfo

func (m *ArchiveFile) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// Response from creating archives.
type CreateArchiveResponse struct {
	// The created archive files.
	Files                []*ArchiveFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateArchiveResponse) Reset()         { *m = CreateArchiveResponse{} }
func (m *CreateArchiveResponse) String() string { return proto.CompactTextString(m) }
func (*CreateArchiveResponse) ProtoMessage()    {}
func (*CreateArchiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e2bd7d1e1e8889, []int{2}
}

func (m *CreateArchiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArchiveResponse.Unmarshal(m, b)
}
func (m *CreateArchiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArchiveResponse.Marshal(b, m, deterministic)
}
func (m *CreateArchiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArchiveResponse.Merge(m, src)
}
func (m *CreateArchiveResponse) XXX_Size() int {
	return xxx_messageInfo_CreateArchiveResponse.Size(m)
}
func (m *CreateArchiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArchiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArchiveResponse proto.InternalMessageInfo

func (m *CreateArchiveResponse) GetFiles() []*ArchiveFile {
	if m != nil {
		return m.Files
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateArchiveRequest)(nil), "chromite.api.CreateArchiveRequest")
	proto.RegisterType((*ArchiveFile)(nil), "chromite.api.ArchiveFile")
	proto.RegisterType((*CreateArchiveResponse)(nil), "chromite.api.CreateArchiveResponse")
}

func init() { proto.RegisterFile("chromite/api/test_archive.proto", fileDescriptor_b9e2bd7d1e1e8889) }

var fileDescriptor_b9e2bd7d1e1e8889 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0xa9, 0x7f, 0x30, 0xe9, 0xa0, 0x43, 0x50, 0x67, 0x2c, 0x82, 0x63, 0xdd, 0xd4, 0x4d,
	0x0a, 0x75, 0xe7, 0xce, 0x1f, 0x64, 0xd6, 0x75, 0xf6, 0x25, 0xed, 0x5c, 0x6d, 0xa0, 0x9d, 0xc4,
	0xe4, 0x66, 0xc4, 0x85, 0x0f, 0xe0, 0x6b, 0xf9, 0x3e, 0xbe, 0x83, 0xd8, 0xb4, 0x4c, 0x47, 0xc4,
	0x5d, 0x72, 0xce, 0xe1, 0xbb, 0x37, 0x27, 0xe4, 0xac, 0x28, 0xb5, 0xac, 0x05, 0x42, 0xcc, 0x95,
	0x88, 0x11, 0x0c, 0x66, 0x5c, 0x17, 0xa5, 0x58, 0x01, 0x53, 0x5a, 0xa2, 0xa4, 0xc3, 0x2e, 0xc0,
	0xb8, 0x12, 0xc1, 0xe9, 0x46, 0x3c, 0xb7, 0xa2, 0x5a, 0x64, 0x5c, 0x09, 0x97, 0x0d, 0xc6, 0xce,
	0xb5, 0xb5, 0x34, 0x71, 0x21, 0xeb, 0x5a, 0x2e, 0x9d, 0x11, 0xbe, 0x93, 0xc3, 0x3b, 0x0d, 0x1c,
	0xe1, 0xc6, 0xb1, 0x53, 0x78, 0xb1, 0x60, 0x90, 0x5e, 0x93, 0xa1, 0x63, 0x20, 0xd7, 0xcf, 0x80,
	0x13, 0x6f, 0xea, 0x45, 0x7e, 0x32, 0x66, 0x6b, 0x0e, 0xbb, 0xfd, 0xf1, 0xe7, 0x8d, 0x9d, 0xfa,
	0xf9, 0xfa, 0x42, 0x2f, 0xc9, 0x48, 0x5a, 0x54, 0x16, 0xb3, 0x85, 0xd0, 0x50, 0xa0, 0xd4, 0x6f,
	0x93, 0xad, 0xa9, 0x17, 0x0d, 0xd2, 0x03, 0xa7, 0xdf, 0x77, 0x72, 0x78, 0x4e, 0xfc, 0x76, 0xf0,
	0x83, 0xa8, 0x80, 0x52, 0xb2, 0xa3, 0x38, 0x96, 0xcd, 0xb4, 0x41, 0xda, 0x9c, 0xc3, 0x19, 0x39,
	0xfa, 0xb5, 0xa1, 0x51, 0x72, 0x69, 0x80, 0xc6, 0x64, 0xf7, 0x49, 0x54, 0x60, 0x26, 0xde, 0x74,
	0x3b, 0xf2, 0x93, 0x13, 0xd6, 0xef, 0x83, 0xf5, 0xb0, 0xa9, 0xcb, 0x25, 0x1f, 0x1e, 0xd9, 0x6f,
	0xe5, 0x47, 0xd0, 0x2b, 0x51, 0x00, 0xcd, 0xba, 0xe7, 0xcf, 0x5e, 0xe7, 0x60, 0xb0, 0x75, 0x0d,
	0x0d, 0x37, 0x61, 0x7f, 0x55, 0x14, 0x5c, 0xfc, 0x9b, 0x71, 0x4b, 0x06, 0xc7, 0x9f, 0x5f, 0x01,
	0x25, 0x23, 0x6e, 0x51, 0xf6, 0xbf, 0x30, 0xdf, 0x6b, 0xea, 0xbf, 0xfa, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xf2, 0x88, 0x0f, 0xbf, 0xe6, 0x01, 0x00, 0x00,
}
