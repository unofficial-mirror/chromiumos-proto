// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chromite/api/android.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MarkStableStatusType int32

const (
	// Unspecified
	MarkStableStatusType_MARK_STABLE_STATUS_UNSPECIFIED MarkStableStatusType = 0
	// Success
	MarkStableStatusType_MARK_STABLE_STATUS_SUCCESS MarkStableStatusType = 1
	// Pinned (at android_atom)
	MarkStableStatusType_MARK_STABLE_STATUS_PINNED MarkStableStatusType = 2
	// Early exit
	MarkStableStatusType_MARK_STABLE_STATUS_EARLY_EXIT MarkStableStatusType = 3
)

var MarkStableStatusType_name = map[int32]string{
	0: "MARK_STABLE_STATUS_UNSPECIFIED",
	1: "MARK_STABLE_STATUS_SUCCESS",
	2: "MARK_STABLE_STATUS_PINNED",
	3: "MARK_STABLE_STATUS_EARLY_EXIT",
}

var MarkStableStatusType_value = map[string]int32{
	"MARK_STABLE_STATUS_UNSPECIFIED": 0,
	"MARK_STABLE_STATUS_SUCCESS":     1,
	"MARK_STABLE_STATUS_PINNED":      2,
	"MARK_STABLE_STATUS_EARLY_EXIT":  3,
}

func (x MarkStableStatusType) String() string {
	return proto.EnumName(MarkStableStatusType_name, int32(x))
}

func (MarkStableStatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f74c095a284c1bd1, []int{0}
}

type MarkStableRequest struct {
	// Required.
	// The manifest branch being used.
	TrackingBranch string `protobuf:"bytes,1,opt,name=tracking_branch,json=trackingBranch,proto3" json:"tracking_branch,omitempty"`
	// Required.
	// Portage package name for Android container.
	PackageName string `protobuf:"bytes,2,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	// Required.
	// Android branch to import from.
	AndroidBuildBranch string `protobuf:"bytes,3,opt,name=android_build_branch,json=androidBuildBranch,proto3" json:"android_build_branch,omitempty"`
	// Force set the android build id that will be used.
	AndroidVersion string `protobuf:"bytes,4,opt,name=android_version,json=androidVersion,proto3" json:"android_version,omitempty"`
	// Android GTS branch to copy artifacts from.
	AndroidGtsBuildBranch string `protobuf:"bytes,5,opt,name=android_gts_build_branch,json=androidGtsBuildBranch,proto3" json:"android_gts_build_branch,omitempty"`
	// The set of relevant build targets. Used to clean old version and for a
	// emerge-able sanity check for the new version.
	// Recommended.
	Boards []*chromiumos.BuildTarget `protobuf:"bytes,6,rep,name=boards,proto3" json:"boards,omitempty"`
	// The buildroot to use to execute the endpoint.  The chroot is the directory
	// "chroot" in the buildroot directory.
	Buildroot            string   `protobuf:"bytes,7,opt,name=buildroot,proto3" json:"buildroot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkStableRequest) Reset()         { *m = MarkStableRequest{} }
func (m *MarkStableRequest) String() string { return proto.CompactTextString(m) }
func (*MarkStableRequest) ProtoMessage()    {}
func (*MarkStableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f74c095a284c1bd1, []int{0}
}

func (m *MarkStableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkStableRequest.Unmarshal(m, b)
}
func (m *MarkStableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkStableRequest.Marshal(b, m, deterministic)
}
func (m *MarkStableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkStableRequest.Merge(m, src)
}
func (m *MarkStableRequest) XXX_Size() int {
	return xxx_messageInfo_MarkStableRequest.Size(m)
}
func (m *MarkStableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkStableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MarkStableRequest proto.InternalMessageInfo

func (m *MarkStableRequest) GetTrackingBranch() string {
	if m != nil {
		return m.TrackingBranch
	}
	return ""
}

func (m *MarkStableRequest) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *MarkStableRequest) GetAndroidBuildBranch() string {
	if m != nil {
		return m.AndroidBuildBranch
	}
	return ""
}

func (m *MarkStableRequest) GetAndroidVersion() string {
	if m != nil {
		return m.AndroidVersion
	}
	return ""
}

func (m *MarkStableRequest) GetAndroidGtsBuildBranch() string {
	if m != nil {
		return m.AndroidGtsBuildBranch
	}
	return ""
}

func (m *MarkStableRequest) GetBoards() []*chromiumos.BuildTarget {
	if m != nil {
		return m.Boards
	}
	return nil
}

func (m *MarkStableRequest) GetBuildroot() string {
	if m != nil {
		return m.Buildroot
	}
	return ""
}

type MarkStableResponse struct {
	// Possible errors.
	Status MarkStableStatusType `protobuf:"varint,1,opt,name=status,proto3,enum=chromite.api.MarkStableStatusType" json:"status,omitempty"`
	// The new package atom.
	AndroidAtom          *chromiumos.PackageInfo `protobuf:"bytes,2,opt,name=android_atom,json=androidAtom,proto3" json:"android_atom,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MarkStableResponse) Reset()         { *m = MarkStableResponse{} }
func (m *MarkStableResponse) String() string { return proto.CompactTextString(m) }
func (*MarkStableResponse) ProtoMessage()    {}
func (*MarkStableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f74c095a284c1bd1, []int{1}
}

func (m *MarkStableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkStableResponse.Unmarshal(m, b)
}
func (m *MarkStableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkStableResponse.Marshal(b, m, deterministic)
}
func (m *MarkStableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkStableResponse.Merge(m, src)
}
func (m *MarkStableResponse) XXX_Size() int {
	return xxx_messageInfo_MarkStableResponse.Size(m)
}
func (m *MarkStableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkStableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MarkStableResponse proto.InternalMessageInfo

func (m *MarkStableResponse) GetStatus() MarkStableStatusType {
	if m != nil {
		return m.Status
	}
	return MarkStableStatusType_MARK_STABLE_STATUS_UNSPECIFIED
}

func (m *MarkStableResponse) GetAndroidAtom() *chromiumos.PackageInfo {
	if m != nil {
		return m.AndroidAtom
	}
	return nil
}

type UnpinVersionRequest struct {
	// The chroot to use to execute the endpoint.
	Chroot               *chromiumos.Chroot `protobuf:"bytes,1,opt,name=chroot,proto3" json:"chroot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UnpinVersionRequest) Reset()         { *m = UnpinVersionRequest{} }
func (m *UnpinVersionRequest) String() string { return proto.CompactTextString(m) }
func (*UnpinVersionRequest) ProtoMessage()    {}
func (*UnpinVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f74c095a284c1bd1, []int{2}
}

func (m *UnpinVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnpinVersionRequest.Unmarshal(m, b)
}
func (m *UnpinVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnpinVersionRequest.Marshal(b, m, deterministic)
}
func (m *UnpinVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnpinVersionRequest.Merge(m, src)
}
func (m *UnpinVersionRequest) XXX_Size() int {
	return xxx_messageInfo_UnpinVersionRequest.Size(m)
}
func (m *UnpinVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnpinVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnpinVersionRequest proto.InternalMessageInfo

func (m *UnpinVersionRequest) GetChroot() *chromiumos.Chroot {
	if m != nil {
		return m.Chroot
	}
	return nil
}

type UnpinVersionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnpinVersionResponse) Reset()         { *m = UnpinVersionResponse{} }
func (m *UnpinVersionResponse) String() string { return proto.CompactTextString(m) }
func (*UnpinVersionResponse) ProtoMessage()    {}
func (*UnpinVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f74c095a284c1bd1, []int{3}
}

func (m *UnpinVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnpinVersionResponse.Unmarshal(m, b)
}
func (m *UnpinVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnpinVersionResponse.Marshal(b, m, deterministic)
}
func (m *UnpinVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnpinVersionResponse.Merge(m, src)
}
func (m *UnpinVersionResponse) XXX_Size() int {
	return xxx_messageInfo_UnpinVersionResponse.Size(m)
}
func (m *UnpinVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnpinVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnpinVersionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("chromite.api.MarkStableStatusType", MarkStableStatusType_name, MarkStableStatusType_value)
	proto.RegisterType((*MarkStableRequest)(nil), "chromite.api.MarkStableRequest")
	proto.RegisterType((*MarkStableResponse)(nil), "chromite.api.MarkStableResponse")
	proto.RegisterType((*UnpinVersionRequest)(nil), "chromite.api.UnpinVersionRequest")
	proto.RegisterType((*UnpinVersionResponse)(nil), "chromite.api.UnpinVersionResponse")
}

func init() { proto.RegisterFile("chromite/api/android.proto", fileDescriptor_f74c095a284c1bd1) }

var fileDescriptor_f74c095a284c1bd1 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0x07, 0x99, 0x76, 0x5a, 0x6d, 0xc1, 0x14, 0x16, 0xa2, 0x6d, 0xb4, 0xb9, 0xd9,
	0xb4, 0x8b, 0x04, 0x15, 0x09, 0xd0, 0xee, 0xd2, 0x2e, 0xa0, 0x8a, 0xad, 0xaa, 0x92, 0x96, 0x7f,
	0x12, 0x8a, 0xdc, 0xd4, 0xcb, 0xa2, 0x2e, 0x71, 0x70, 0xdc, 0x49, 0xbc, 0x03, 0x8f, 0xc0, 0xd3,
	0xf4, 0x2d, 0x78, 0x88, 0xbd, 0x03, 0xaa, 0xe3, 0xd0, 0x14, 0x75, 0x5c, 0x45, 0x3a, 0xdf, 0xef,
	0x7c, 0x3e, 0xfe, 0x8e, 0x1c, 0x30, 0xc2, 0x6b, 0x46, 0x93, 0x98, 0x13, 0x1b, 0x67, 0xb1, 0x8d,
	0xd3, 0x29, 0xa3, 0xf1, 0xd4, 0xca, 0x18, 0xe5, 0x14, 0x35, 0x4a, 0xcd, 0xc2, 0x59, 0x6c, 0x1c,
	0xac, 0x91, 0x93, 0x79, 0x7c, 0x33, 0x0d, 0x70, 0x16, 0x17, 0xac, 0xb1, 0x5f, 0xa8, 0xf3, 0x84,
	0xe6, 0x76, 0x48, 0x93, 0x84, 0xa6, 0x85, 0x60, 0x2e, 0x6a, 0xf0, 0xf8, 0x12, 0xb3, 0x99, 0xcf,
	0xf1, 0xe4, 0x86, 0x78, 0xe4, 0xfb, 0x9c, 0xe4, 0x1c, 0x1d, 0xc3, 0x1e, 0x67, 0x38, 0x9c, 0xc5,
	0x69, 0x14, 0x4c, 0x18, 0x4e, 0xc3, 0x6b, 0x5d, 0x69, 0x29, 0x27, 0x3b, 0xde, 0x6e, 0x59, 0xee,
	0x8a, 0x2a, 0x6a, 0x43, 0x23, 0xc3, 0xe1, 0x0c, 0x47, 0x24, 0x48, 0x71, 0x42, 0xf4, 0x9a, 0xa0,
	0xea, 0xb2, 0x36, 0xc0, 0x09, 0x41, 0x2f, 0xa1, 0x29, 0xe7, 0x0e, 0x8a, 0xa9, 0xa4, 0xe1, 0x96,
	0x40, 0x91, 0xd4, 0xba, 0x4b, 0x49, 0x9a, 0x1e, 0xc3, 0x5e, 0xd9, 0x71, 0x4b, 0x58, 0x1e, 0xd3,
	0x54, 0x7f, 0x58, 0x9c, 0x2e, 0xcb, 0x1f, 0x8b, 0x2a, 0x7a, 0x03, 0x7a, 0x09, 0x46, 0x3c, 0x5f,
	0xb7, 0x7f, 0x24, 0x3a, 0x9e, 0x4a, 0xfd, 0x3d, 0xcf, 0xab, 0x27, 0xd8, 0xa0, 0x4e, 0x28, 0x66,
	0xd3, 0x5c, 0x57, 0x5b, 0x5b, 0x27, 0xf5, 0xce, 0xbe, 0xb5, 0xca, 0xc7, 0x12, 0xe0, 0x08, 0xb3,
	0x88, 0x70, 0x4f, 0x62, 0xe8, 0x00, 0x76, 0x84, 0x3b, 0xa3, 0x94, 0xeb, 0xdb, 0xc2, 0x7a, 0x55,
	0x30, 0x7f, 0x2a, 0x80, 0xaa, 0x21, 0xe6, 0x19, 0x4d, 0x73, 0x82, 0xce, 0x40, 0xcd, 0x39, 0xe6,
	0xf3, 0x5c, 0x84, 0xb7, 0xdb, 0x31, 0xad, 0xea, 0xc6, 0xac, 0x55, 0x87, 0x2f, 0xa8, 0xd1, 0x8f,
	0x8c, 0x78, 0xb2, 0x03, 0x9d, 0x41, 0xa3, 0xbc, 0x1a, 0xe6, 0x34, 0x11, 0xc1, 0xfe, 0x33, 0xe7,
	0xb0, 0x08, 0xb9, 0x9f, 0x5e, 0x51, 0xaf, 0x2e, 0x61, 0x87, 0xd3, 0xc4, 0x74, 0xe0, 0xc9, 0x38,
	0xcd, 0xe2, 0x54, 0xc6, 0x54, 0x2e, 0xf5, 0x14, 0xd4, 0x65, 0x37, 0xe5, 0x62, 0x9c, 0x7a, 0x07,
	0x55, 0xcd, 0x7a, 0x42, 0xf1, 0x24, 0x61, 0x3e, 0x83, 0xe6, 0xba, 0x45, 0x71, 0xa5, 0xd3, 0x5f,
	0x0a, 0x34, 0x37, 0xcd, 0x8d, 0x4c, 0x38, 0xba, 0x74, 0xbc, 0x0f, 0x81, 0x3f, 0x72, 0xba, 0x17,
	0xee, 0xf2, 0x33, 0x1a, 0xfb, 0xc1, 0x78, 0xe0, 0x0f, 0xdd, 0x5e, 0xff, 0x5d, 0xdf, 0x3d, 0xd7,
	0x1e, 0xa0, 0x23, 0x30, 0x36, 0x30, 0xfe, 0xb8, 0xd7, 0x73, 0x7d, 0x5f, 0x53, 0xd0, 0x21, 0x3c,
	0xdf, 0xa0, 0x0f, 0xfb, 0x83, 0x81, 0x7b, 0xae, 0xd5, 0x50, 0x1b, 0x0e, 0x37, 0xc8, 0xae, 0xe3,
	0x5d, 0x7c, 0x09, 0xdc, 0xcf, 0xfd, 0x91, 0xb6, 0xd5, 0xf9, 0xad, 0xc0, 0xae, 0x53, 0x24, 0xe1,
	0x13, 0x76, 0x1b, 0x87, 0x04, 0x7d, 0x02, 0x58, 0x0d, 0x8c, 0x5e, 0xdc, 0xb7, 0x02, 0x19, 0x92,
	0xd1, 0xba, 0x1f, 0x28, 0x22, 0x30, 0xd5, 0xc5, 0x9d, 0x51, 0xd3, 0x6a, 0xe8, 0x1b, 0x34, 0xaa,
	0x11, 0xa1, 0xf6, 0x7a, 0xe7, 0x86, 0x0d, 0x18, 0xe6, 0xff, 0x90, 0x35, 0x7b, 0xc5, 0xd8, 0x5b,
	0xdc, 0x19, 0x75, 0xd8, 0x96, 0x7b, 0xd5, 0x94, 0xee, 0xdb, 0xaf, 0xaf, 0x23, 0xfa, 0x77, 0x65,
	0x16, 0x65, 0x91, 0x5d, 0x79, 0xd4, 0x71, 0x7a, 0xc5, 0xb0, 0x2d, 0xde, 0xb4, 0x1d, 0x51, 0xbb,
	0xfa, 0x33, 0x98, 0xa8, 0xa2, 0xfc, 0xea, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xab, 0x41,
	0x98, 0x4d, 0x04, 0x00, 0x00,
}
