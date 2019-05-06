// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chromiumos/common.proto

package chromiumos

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

// BuildTarget encapsulates a number of related arguments. At the moment, the
// usage of specific arguments is on a per endpoint basis, but will converge
// to support a standard set as time goes on.
// BuildTarget is a more appropriate name going forward for what we currently
// call "board".
type BuildTarget struct {
	// The name of the build target (a.k.a. board name).
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildTarget) Reset()         { *m = BuildTarget{} }
func (m *BuildTarget) String() string { return proto.CompactTextString(m) }
func (*BuildTarget) ProtoMessage()    {}
func (*BuildTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa0b57c3d7d4c63b, []int{0}
}

func (m *BuildTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildTarget.Unmarshal(m, b)
}
func (m *BuildTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildTarget.Marshal(b, m, deterministic)
}
func (m *BuildTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildTarget.Merge(m, src)
}
func (m *BuildTarget) XXX_Size() int {
	return xxx_messageInfo_BuildTarget.Size(m)
}
func (m *BuildTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildTarget.DiscardUnknown(m)
}

var xxx_messageInfo_BuildTarget proto.InternalMessageInfo

func (m *BuildTarget) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Chroot is used to define how to enter a chroot.
type Chroot struct {
	// The directory containing the chroot.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The cache directory override.
	CacheDir string `protobuf:"bytes,2,opt,name=cache_dir,json=cacheDir,proto3" json:"cache_dir,omitempty"`
	// Environment configuration.
	Env                  *Chroot_ChrootEnv `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Chroot) Reset()         { *m = Chroot{} }
func (m *Chroot) String() string { return proto.CompactTextString(m) }
func (*Chroot) ProtoMessage()    {}
func (*Chroot) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa0b57c3d7d4c63b, []int{1}
}

func (m *Chroot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chroot.Unmarshal(m, b)
}
func (m *Chroot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chroot.Marshal(b, m, deterministic)
}
func (m *Chroot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chroot.Merge(m, src)
}
func (m *Chroot) XXX_Size() int {
	return xxx_messageInfo_Chroot.Size(m)
}
func (m *Chroot) XXX_DiscardUnknown() {
	xxx_messageInfo_Chroot.DiscardUnknown(m)
}

var xxx_messageInfo_Chroot proto.InternalMessageInfo

func (m *Chroot) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Chroot) GetCacheDir() string {
	if m != nil {
		return m.CacheDir
	}
	return ""
}

func (m *Chroot) GetEnv() *Chroot_ChrootEnv {
	if m != nil {
		return m.Env
	}
	return nil
}

type Chroot_ChrootEnv struct {
	// USE flags to set.
	UseFlags []*UseFlag `protobuf:"bytes,1,rep,name=use_flags,json=useFlags,proto3" json:"use_flags,omitempty"`
	// FEATURES flags to set.
	Features             []*Feature `protobuf:"bytes,2,rep,name=features,proto3" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Chroot_ChrootEnv) Reset()         { *m = Chroot_ChrootEnv{} }
func (m *Chroot_ChrootEnv) String() string { return proto.CompactTextString(m) }
func (*Chroot_ChrootEnv) ProtoMessage()    {}
func (*Chroot_ChrootEnv) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa0b57c3d7d4c63b, []int{1, 0}
}

func (m *Chroot_ChrootEnv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chroot_ChrootEnv.Unmarshal(m, b)
}
func (m *Chroot_ChrootEnv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chroot_ChrootEnv.Marshal(b, m, deterministic)
}
func (m *Chroot_ChrootEnv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chroot_ChrootEnv.Merge(m, src)
}
func (m *Chroot_ChrootEnv) XXX_Size() int {
	return xxx_messageInfo_Chroot_ChrootEnv.Size(m)
}
func (m *Chroot_ChrootEnv) XXX_DiscardUnknown() {
	xxx_messageInfo_Chroot_ChrootEnv.DiscardUnknown(m)
}

var xxx_messageInfo_Chroot_ChrootEnv proto.InternalMessageInfo

func (m *Chroot_ChrootEnv) GetUseFlags() []*UseFlag {
	if m != nil {
		return m.UseFlags
	}
	return nil
}

func (m *Chroot_ChrootEnv) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

// FEATURES environment variable component message.
type Feature struct {
	Feature              string   `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa0b57c3d7d4c63b, []int{2}
}

func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (m *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(m, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetFeature() string {
	if m != nil {
		return m.Feature
	}
	return ""
}

// Message describing a package, The corresponding variable names are from
// https://devmanual.gentoo.org/ebuild-writing/variables/index.html
type PackageInfo struct {
	// The package name (PN variable).
	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	// The package category (CATEGORY variable).
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	// The package version and revision (PVR variable).
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PackageInfo) Reset()         { *m = PackageInfo{} }
func (m *PackageInfo) String() string { return proto.CompactTextString(m) }
func (*PackageInfo) ProtoMessage()    {}
func (*PackageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa0b57c3d7d4c63b, []int{3}
}

func (m *PackageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageInfo.Unmarshal(m, b)
}
func (m *PackageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageInfo.Marshal(b, m, deterministic)
}
func (m *PackageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageInfo.Merge(m, src)
}
func (m *PackageInfo) XXX_Size() int {
	return xxx_messageInfo_PackageInfo.Size(m)
}
func (m *PackageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PackageInfo proto.InternalMessageInfo

func (m *PackageInfo) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *PackageInfo) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *PackageInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Message for USE flags.
type UseFlag struct {
	// The use flag.
	Flag                 string   `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UseFlag) Reset()         { *m = UseFlag{} }
func (m *UseFlag) String() string { return proto.CompactTextString(m) }
func (*UseFlag) ProtoMessage()    {}
func (*UseFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa0b57c3d7d4c63b, []int{4}
}

func (m *UseFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseFlag.Unmarshal(m, b)
}
func (m *UseFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseFlag.Marshal(b, m, deterministic)
}
func (m *UseFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseFlag.Merge(m, src)
}
func (m *UseFlag) XXX_Size() int {
	return xxx_messageInfo_UseFlag.Size(m)
}
func (m *UseFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_UseFlag.DiscardUnknown(m)
}

var xxx_messageInfo_UseFlag proto.InternalMessageInfo

func (m *UseFlag) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildTarget)(nil), "chromiumos.BuildTarget")
	proto.RegisterType((*Chroot)(nil), "chromiumos.Chroot")
	proto.RegisterType((*Chroot_ChrootEnv)(nil), "chromiumos.Chroot.ChrootEnv")
	proto.RegisterType((*Feature)(nil), "chromiumos.Feature")
	proto.RegisterType((*PackageInfo)(nil), "chromiumos.PackageInfo")
	proto.RegisterType((*UseFlag)(nil), "chromiumos.UseFlag")
}

func init() { proto.RegisterFile("chromiumos/common.proto", fileDescriptor_fa0b57c3d7d4c63b) }

var fileDescriptor_fa0b57c3d7d4c63b = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xcf, 0x4e, 0xbb, 0x40,
	0x10, 0x0e, 0xed, 0x2f, 0x2d, 0x0c, 0xbf, 0xd3, 0x7a, 0x90, 0x54, 0x4d, 0x5a, 0xbc, 0xf4, 0x04,
	0xa6, 0x1a, 0x1f, 0xa0, 0x6a, 0x13, 0x2f, 0xc6, 0x10, 0xbd, 0x78, 0x69, 0x56, 0x3a, 0x2c, 0xc4,
	0xb2, 0xd3, 0xec, 0x42, 0x13, 0x1f, 0xd5, 0xb7, 0x31, 0x2c, 0x4b, 0x21, 0xf1, 0xc4, 0xf7, 0x67,
	0x99, 0xc9, 0xf7, 0x0d, 0x9c, 0xa7, 0xb9, 0xa2, 0xb2, 0xa8, 0x4b, 0xd2, 0x71, 0x4a, 0x65, 0x49,
	0x32, 0x3a, 0x28, 0xaa, 0x88, 0x41, 0x6f, 0x84, 0x0b, 0xf0, 0xd7, 0x75, 0xb1, 0xdf, 0xbd, 0x71,
	0x25, 0xb0, 0x62, 0x0c, 0xfe, 0x49, 0x5e, 0x62, 0xe0, 0xcc, 0x9d, 0xa5, 0x97, 0x18, 0x1c, 0xfe,
	0x38, 0x30, 0x79, 0xc8, 0x15, 0x91, 0xb1, 0x0f, 0xbc, 0xca, 0x3b, 0xbb, 0xc1, 0xec, 0x02, 0xbc,
	0x94, 0xa7, 0x39, 0x6e, 0x77, 0x85, 0x0a, 0x46, 0xc6, 0x70, 0x8d, 0xf0, 0x58, 0x28, 0x16, 0xc1,
	0x18, 0xe5, 0x31, 0x18, 0xcf, 0x9d, 0xa5, 0xbf, 0xba, 0x8c, 0xfa, 0xc5, 0x51, 0x3b, 0xd1, 0x7e,
	0x9e, 0xe4, 0x31, 0x69, 0x1e, 0xce, 0x24, 0x78, 0x27, 0x85, 0xdd, 0x80, 0x57, 0x6b, 0xdc, 0x66,
	0x7b, 0x2e, 0x74, 0xe0, 0xcc, 0xc7, 0x4b, 0x7f, 0x75, 0x36, 0x1c, 0xf1, 0xae, 0x71, 0xb3, 0xe7,
	0x22, 0x71, 0xeb, 0x16, 0x68, 0x16, 0x83, 0x9b, 0x21, 0xaf, 0x6a, 0x85, 0x3a, 0x18, 0xfd, 0xfd,
	0x61, 0xd3, 0x7a, 0xc9, 0xe9, 0x51, 0x78, 0x0d, 0x53, 0x2b, 0xb2, 0x00, 0xa6, 0x56, 0xb6, 0xf1,
	0x3a, 0x1a, 0x66, 0xe0, 0xbf, 0xf2, 0xf4, 0x8b, 0x0b, 0x7c, 0x96, 0x19, 0xb1, 0x05, 0xfc, 0x3f,
	0xb4, 0x74, 0x3b, 0xe8, 0xca, 0xb7, 0xda, 0x0b, 0x2f, 0x91, 0xcd, 0xc0, 0x4d, 0x79, 0x85, 0x82,
	0xd4, 0x77, 0x5f, 0x49, 0xcb, 0x9b, 0x3d, 0x47, 0x54, 0xba, 0x20, 0x69, 0x6a, 0xf1, 0x92, 0x8e,
	0x86, 0x57, 0x30, 0xb5, 0x91, 0x9a, 0xa2, 0x9b, 0xd8, 0x5d, 0xd1, 0x0d, 0x5e, 0xdf, 0x7f, 0xdc,
	0x09, 0x3a, 0xc5, 0x89, 0x48, 0x89, 0x78, 0x70, 0xe1, 0x42, 0x66, 0x8a, 0xc7, 0xe6, 0xc0, 0xb1,
	0xa0, 0x81, 0xf3, 0x39, 0x31, 0xe2, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0xb3, 0xff,
	0xda, 0x10, 0x02, 0x00, 0x00,
}
