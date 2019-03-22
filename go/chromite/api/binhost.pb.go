// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chromite/api/binhost.proto

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

// Portage environment variable that points to a remote binhost.
type BinhostKey int32

const (
	BinhostKey_POSTSUBMIT_BINHOST            BinhostKey = 0
	BinhostKey_LATEST_RELEASE_CHROME_BINHOST BinhostKey = 1
	BinhostKey_PREFLIGHT_BINHOST             BinhostKey = 2
)

var BinhostKey_name = map[int32]string{
	0: "POSTSUBMIT_BINHOST",
	1: "LATEST_RELEASE_CHROME_BINHOST",
	2: "PREFLIGHT_BINHOST",
}

var BinhostKey_value = map[string]int32{
	"POSTSUBMIT_BINHOST":            0,
	"LATEST_RELEASE_CHROME_BINHOST": 1,
	"PREFLIGHT_BINHOST":             2,
}

func (x BinhostKey) String() string {
	return proto.EnumName(BinhostKey_name, int32(x))
}

func (BinhostKey) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d632782a0b1177ef, []int{0}
}

// Describes where prebuilts will be uploaded so package index and other
// Portage metadata can be updated appropriately.
type PrepareBinhostUploadsRequest struct {
	// Build target to prepare prebuilts for.
	BuildTarget *chromiumos.BuildTarget `protobuf:"bytes,1,opt,name=build_target,json=buildTarget,proto3" json:"build_target,omitempty"`
	// Full URI where we wish to store prebuilts. Note that this service
	// call does NOT upload them, it only updates metadata.
	// Example: gs://chromeos-prebuilt/board/amd64-generic/packages/
	Uri                  string   `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareBinhostUploadsRequest) Reset()         { *m = PrepareBinhostUploadsRequest{} }
func (m *PrepareBinhostUploadsRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareBinhostUploadsRequest) ProtoMessage()    {}
func (*PrepareBinhostUploadsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d632782a0b1177ef, []int{0}
}

func (m *PrepareBinhostUploadsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareBinhostUploadsRequest.Unmarshal(m, b)
}
func (m *PrepareBinhostUploadsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareBinhostUploadsRequest.Marshal(b, m, deterministic)
}
func (m *PrepareBinhostUploadsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareBinhostUploadsRequest.Merge(m, src)
}
func (m *PrepareBinhostUploadsRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareBinhostUploadsRequest.Size(m)
}
func (m *PrepareBinhostUploadsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareBinhostUploadsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareBinhostUploadsRequest proto.InternalMessageInfo

func (m *PrepareBinhostUploadsRequest) GetBuildTarget() *chromiumos.BuildTarget {
	if m != nil {
		return m.BuildTarget
	}
	return nil
}

func (m *PrepareBinhostUploadsRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

// Return all files to upload.
type PrepareBinhostUploadsResponse struct {
	// Absolute chroot path to the local directory containing files to upload.
	UploadsDir string `protobuf:"bytes,1,opt,name=uploads_dir,json=uploadsDir,proto3" json:"uploads_dir,omitempty"`
	// All targets to be uploaded to the binhost.
	UploadTargets        []*PrepareBinhostUploadsResponse_UploadTarget `protobuf:"bytes,2,rep,name=upload_targets,json=uploadTargets,proto3" json:"upload_targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *PrepareBinhostUploadsResponse) Reset()         { *m = PrepareBinhostUploadsResponse{} }
func (m *PrepareBinhostUploadsResponse) String() string { return proto.CompactTextString(m) }
func (*PrepareBinhostUploadsResponse) ProtoMessage()    {}
func (*PrepareBinhostUploadsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d632782a0b1177ef, []int{1}
}

func (m *PrepareBinhostUploadsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareBinhostUploadsResponse.Unmarshal(m, b)
}
func (m *PrepareBinhostUploadsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareBinhostUploadsResponse.Marshal(b, m, deterministic)
}
func (m *PrepareBinhostUploadsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareBinhostUploadsResponse.Merge(m, src)
}
func (m *PrepareBinhostUploadsResponse) XXX_Size() int {
	return xxx_messageInfo_PrepareBinhostUploadsResponse.Size(m)
}
func (m *PrepareBinhostUploadsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareBinhostUploadsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareBinhostUploadsResponse proto.InternalMessageInfo

func (m *PrepareBinhostUploadsResponse) GetUploadsDir() string {
	if m != nil {
		return m.UploadsDir
	}
	return ""
}

func (m *PrepareBinhostUploadsResponse) GetUploadTargets() []*PrepareBinhostUploadsResponse_UploadTarget {
	if m != nil {
		return m.UploadTargets
	}
	return nil
}

// An upload target is a file in the uploads_dir.
type PrepareBinhostUploadsResponse_UploadTarget struct {
	// Paths relative to uploads_dir.
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareBinhostUploadsResponse_UploadTarget) Reset() {
	*m = PrepareBinhostUploadsResponse_UploadTarget{}
}
func (m *PrepareBinhostUploadsResponse_UploadTarget) String() string {
	return proto.CompactTextString(m)
}
func (*PrepareBinhostUploadsResponse_UploadTarget) ProtoMessage() {}
func (*PrepareBinhostUploadsResponse_UploadTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_d632782a0b1177ef, []int{1, 0}
}

func (m *PrepareBinhostUploadsResponse_UploadTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareBinhostUploadsResponse_UploadTarget.Unmarshal(m, b)
}
func (m *PrepareBinhostUploadsResponse_UploadTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareBinhostUploadsResponse_UploadTarget.Marshal(b, m, deterministic)
}
func (m *PrepareBinhostUploadsResponse_UploadTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareBinhostUploadsResponse_UploadTarget.Merge(m, src)
}
func (m *PrepareBinhostUploadsResponse_UploadTarget) XXX_Size() int {
	return xxx_messageInfo_PrepareBinhostUploadsResponse_UploadTarget.Size(m)
}
func (m *PrepareBinhostUploadsResponse_UploadTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareBinhostUploadsResponse_UploadTarget.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareBinhostUploadsResponse_UploadTarget proto.InternalMessageInfo

func (m *PrepareBinhostUploadsResponse_UploadTarget) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// Set a binhost URL for the given build target.
type SetBinhostRequest struct {
	// Build target to update binhost for.
	BuildTarget *chromiumos.BuildTarget `protobuf:"bytes,1,opt,name=build_target,json=buildTarget,proto3" json:"build_target,omitempty"`
	// Whether or not the binhost is private.
	Private bool `protobuf:"varint,2,opt,name=private,proto3" json:"private,omitempty"`
	// Binhost key to set or update.
	Key BinhostKey `protobuf:"varint,3,opt,name=key,proto3,enum=chromite.api.BinhostKey" json:"key,omitempty"`
	// URI storing all the binaries.
	Uri                  string   `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetBinhostRequest) Reset()         { *m = SetBinhostRequest{} }
func (m *SetBinhostRequest) String() string { return proto.CompactTextString(m) }
func (*SetBinhostRequest) ProtoMessage()    {}
func (*SetBinhostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d632782a0b1177ef, []int{2}
}

func (m *SetBinhostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetBinhostRequest.Unmarshal(m, b)
}
func (m *SetBinhostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetBinhostRequest.Marshal(b, m, deterministic)
}
func (m *SetBinhostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBinhostRequest.Merge(m, src)
}
func (m *SetBinhostRequest) XXX_Size() int {
	return xxx_messageInfo_SetBinhostRequest.Size(m)
}
func (m *SetBinhostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBinhostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetBinhostRequest proto.InternalMessageInfo

func (m *SetBinhostRequest) GetBuildTarget() *chromiumos.BuildTarget {
	if m != nil {
		return m.BuildTarget
	}
	return nil
}

func (m *SetBinhostRequest) GetPrivate() bool {
	if m != nil {
		return m.Private
	}
	return false
}

func (m *SetBinhostRequest) GetKey() BinhostKey {
	if m != nil {
		return m.Key
	}
	return BinhostKey_POSTSUBMIT_BINHOST
}

func (m *SetBinhostRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

// Response for SetBinhost service call.
type SetBinhostResponse struct {
	// Absolute path to the updated binhost conf file.
	OutputFile           string   `protobuf:"bytes,1,opt,name=output_file,json=outputFile,proto3" json:"output_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetBinhostResponse) Reset()         { *m = SetBinhostResponse{} }
func (m *SetBinhostResponse) String() string { return proto.CompactTextString(m) }
func (*SetBinhostResponse) ProtoMessage()    {}
func (*SetBinhostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d632782a0b1177ef, []int{3}
}

func (m *SetBinhostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetBinhostResponse.Unmarshal(m, b)
}
func (m *SetBinhostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetBinhostResponse.Marshal(b, m, deterministic)
}
func (m *SetBinhostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBinhostResponse.Merge(m, src)
}
func (m *SetBinhostResponse) XXX_Size() int {
	return xxx_messageInfo_SetBinhostResponse.Size(m)
}
func (m *SetBinhostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBinhostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetBinhostResponse proto.InternalMessageInfo

func (m *SetBinhostResponse) GetOutputFile() string {
	if m != nil {
		return m.OutputFile
	}
	return ""
}

func init() {
	proto.RegisterEnum("chromite.api.BinhostKey", BinhostKey_name, BinhostKey_value)
	proto.RegisterType((*PrepareBinhostUploadsRequest)(nil), "chromite.api.PrepareBinhostUploadsRequest")
	proto.RegisterType((*PrepareBinhostUploadsResponse)(nil), "chromite.api.PrepareBinhostUploadsResponse")
	proto.RegisterType((*PrepareBinhostUploadsResponse_UploadTarget)(nil), "chromite.api.PrepareBinhostUploadsResponse.UploadTarget")
	proto.RegisterType((*SetBinhostRequest)(nil), "chromite.api.SetBinhostRequest")
	proto.RegisterType((*SetBinhostResponse)(nil), "chromite.api.SetBinhostResponse")
}

func init() { proto.RegisterFile("chromite/api/binhost.proto", fileDescriptor_d632782a0b1177ef) }

var fileDescriptor_d632782a0b1177ef = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0xc5, 0xed, 0xc4, 0xd8, 0x4d, 0x29, 0xdd, 0x95, 0xc6, 0xa2, 0x68, 0xd3, 0x42, 0x9f, 0xaa,
	0x22, 0xa5, 0x52, 0x11, 0x12, 0xe2, 0x6d, 0x85, 0x8c, 0x56, 0x74, 0xb4, 0x72, 0xb2, 0x57, 0xa2,
	0xb4, 0x35, 0xd4, 0x22, 0xad, 0x8d, 0xe3, 0x4c, 0xda, 0x17, 0xf1, 0x1f, 0x48, 0xfc, 0x09, 0xaf,
	0xfc, 0x03, 0x6a, 0x9c, 0x92, 0x54, 0x30, 0x04, 0xd2, 0xde, 0x7c, 0xef, 0x39, 0x3e, 0xf7, 0xe4,
	0xdc, 0x18, 0x9c, 0xf9, 0x52, 0x89, 0x15, 0xd7, 0xac, 0x17, 0x4b, 0xde, 0x9b, 0xf1, 0xf5, 0x52,
	0xa4, 0xda, 0x93, 0x4a, 0x68, 0x81, 0x8d, 0x2d, 0xe6, 0xc5, 0x92, 0x3b, 0x27, 0xbb, 0xcc, 0x8c,
	0x27, 0x8b, 0x28, 0x96, 0xdc, 0x70, 0x9d, 0x63, 0x83, 0x66, 0x2b, 0x91, 0xf6, 0xe6, 0x62, 0xb5,
	0x12, 0x6b, 0x03, 0xb4, 0x13, 0x38, 0x99, 0x2a, 0x26, 0x63, 0xc5, 0x06, 0x46, 0xfc, 0x4a, 0x26,
	0x22, 0x5e, 0xa4, 0x94, 0x7d, 0xce, 0x58, 0xaa, 0xf1, 0x25, 0x34, 0x8c, 0x96, 0x8e, 0xd5, 0x47,
	0xa6, 0x6d, 0xe2, 0x92, 0x8e, 0xd5, 0x3f, 0xf6, 0x4a, 0x3d, 0x6f, 0xb0, 0xc1, 0xc3, 0x1c, 0xa6,
	0xd6, 0xac, 0x2c, 0xb0, 0x05, 0xf5, 0x4c, 0x71, 0xbb, 0xe6, 0x92, 0xce, 0x01, 0xdd, 0x1c, 0xdb,
	0xdf, 0x08, 0x9c, 0xde, 0x32, 0x2e, 0x95, 0x62, 0x9d, 0x32, 0x3c, 0x03, 0x2b, 0x33, 0xad, 0x68,
	0xc1, 0x55, 0x3e, 0xee, 0x80, 0x42, 0xd1, 0x7a, 0xcd, 0x15, 0x46, 0xd0, 0x34, 0x55, 0xe1, 0x28,
	0xb5, 0x6b, 0x6e, 0xbd, 0x63, 0xf5, 0x5f, 0x78, 0xd5, 0x38, 0xbc, 0xbf, 0x4e, 0xf1, 0x4c, 0x5d,
	0x78, 0x7e, 0x98, 0x55, 0xaa, 0xd4, 0x69, 0x43, 0xa3, 0x0a, 0x23, 0xc2, 0x9e, 0x8c, 0xf5, 0xb2,
	0xb0, 0x92, 0x9f, 0xdb, 0x5f, 0x08, 0x1c, 0x06, 0x4c, 0x17, 0xea, 0x77, 0x91, 0x95, 0x0d, 0xfb,
	0x52, 0xf1, 0xeb, 0x58, 0xb3, 0x3c, 0xaf, 0x07, 0x74, 0x5b, 0x62, 0x17, 0xea, 0x9f, 0xd8, 0x8d,
	0x5d, 0x77, 0x49, 0xa7, 0xd9, 0xb7, 0x77, 0xbf, 0xb2, 0x30, 0xf0, 0x96, 0xdd, 0xd0, 0x0d, 0x69,
	0x9b, 0xf8, 0x5e, 0x99, 0xf8, 0x73, 0xc0, 0xaa, 0xd1, 0x32, 0x65, 0x91, 0x69, 0x99, 0xe9, 0xe8,
	0x03, 0x4f, 0xd8, 0x36, 0x65, 0xd3, 0xba, 0xe0, 0x09, 0xeb, 0xbe, 0x07, 0x28, 0xb5, 0xf1, 0x31,
	0xe0, 0x74, 0x12, 0x84, 0xc1, 0xd5, 0xe0, 0x72, 0x14, 0x46, 0x83, 0xd1, 0xbb, 0xe1, 0x24, 0x08,
	0x5b, 0xf7, 0xf0, 0x09, 0x9c, 0x8e, 0xcf, 0x43, 0x3f, 0x08, 0x23, 0xea, 0x8f, 0xfd, 0xf3, 0xc0,
	0x8f, 0x5e, 0x0d, 0xe9, 0xe4, 0xd2, 0xff, 0x45, 0x21, 0x78, 0x04, 0x87, 0x53, 0xea, 0x5f, 0x8c,
	0x47, 0x6f, 0x86, 0xe5, 0xcd, 0x5a, 0xff, 0x3b, 0x81, 0x66, 0x31, 0x20, 0x60, 0xea, 0x9a, 0xcf,
	0x19, 0x4a, 0x38, 0xfa, 0xe3, 0xd2, 0xb0, 0xfb, 0x4f, 0x9b, 0xcd, 0x57, 0xe0, 0x3c, 0xfd, 0x8f,
	0xbf, 0x00, 0x27, 0x00, 0x65, 0x36, 0x78, 0xb6, 0x7b, 0xf5, 0xb7, 0xf5, 0x3a, 0xee, 0xed, 0x04,
	0x23, 0xe8, 0x3c, 0xfa, 0xfa, 0xc3, 0xb1, 0x60, 0xbf, 0x78, 0xa6, 0x2d, 0x32, 0xbb, 0x9f, 0x3f,
	0xb2, 0x67, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x8b, 0xde, 0x1d, 0xc7, 0x03, 0x00, 0x00,
}
