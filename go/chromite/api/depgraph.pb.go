// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chromite/api/depgraph.proto

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

// Message contains data about the portage dependency graph and the packages'
// dependency source paths.
type DepGraph struct {
	// The board whose dependency graph is being created for.
	BuildTarget *chromiumos.BuildTarget `protobuf:"bytes,1,opt,name=build_target,json=buildTarget,proto3" json:"build_target,omitempty"`
	// List of packages in the dependency graph and their infos, which include
	// dependency packages and the source paths.
	PackageDeps []*PackageDepInfo `protobuf:"bytes,2,rep,name=package_deps,json=packageDeps,proto3" json:"package_deps,omitempty"`
	// The chroot to use to execute the endpoint.
	Chroot               *chromiumos.Chroot `protobuf:"bytes,3,opt,name=chroot,proto3" json:"chroot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DepGraph) Reset()         { *m = DepGraph{} }
func (m *DepGraph) String() string { return proto.CompactTextString(m) }
func (*DepGraph) ProtoMessage()    {}
func (*DepGraph) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a5308b238ddba97, []int{0}
}

func (m *DepGraph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepGraph.Unmarshal(m, b)
}
func (m *DepGraph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepGraph.Marshal(b, m, deterministic)
}
func (m *DepGraph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepGraph.Merge(m, src)
}
func (m *DepGraph) XXX_Size() int {
	return xxx_messageInfo_DepGraph.Size(m)
}
func (m *DepGraph) XXX_DiscardUnknown() {
	xxx_messageInfo_DepGraph.DiscardUnknown(m)
}

var xxx_messageInfo_DepGraph proto.InternalMessageInfo

func (m *DepGraph) GetBuildTarget() *chromiumos.BuildTarget {
	if m != nil {
		return m.BuildTarget
	}
	return nil
}

func (m *DepGraph) GetPackageDeps() []*PackageDepInfo {
	if m != nil {
		return m.PackageDeps
	}
	return nil
}

func (m *DepGraph) GetChroot() *chromiumos.Chroot {
	if m != nil {
		return m.Chroot
	}
	return nil
}

// Message describing a package and its dependencies.
type PackageDepInfo struct {
	// The package itself.
	PackageInfo *chromiumos.PackageInfo `protobuf:"bytes,1,opt,name=package_info,json=packageInfo,proto3" json:"package_info,omitempty"`
	// List of packages this package depends on.
	DependencyPackages []*chromiumos.PackageInfo `protobuf:"bytes,2,rep,name=dependency_packages,json=dependencyPackages,proto3" json:"dependency_packages,omitempty"`
	// List of source paths the package depends on.
	DependencySourcePaths []*SourcePath `protobuf:"bytes,3,rep,name=dependency_source_paths,json=dependencySourcePaths,proto3" json:"dependency_source_paths,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}      `json:"-"`
	XXX_unrecognized      []byte        `json:"-"`
	XXX_sizecache         int32         `json:"-"`
}

func (m *PackageDepInfo) Reset()         { *m = PackageDepInfo{} }
func (m *PackageDepInfo) String() string { return proto.CompactTextString(m) }
func (*PackageDepInfo) ProtoMessage()    {}
func (*PackageDepInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a5308b238ddba97, []int{1}
}

func (m *PackageDepInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageDepInfo.Unmarshal(m, b)
}
func (m *PackageDepInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageDepInfo.Marshal(b, m, deterministic)
}
func (m *PackageDepInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageDepInfo.Merge(m, src)
}
func (m *PackageDepInfo) XXX_Size() int {
	return xxx_messageInfo_PackageDepInfo.Size(m)
}
func (m *PackageDepInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageDepInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PackageDepInfo proto.InternalMessageInfo

func (m *PackageDepInfo) GetPackageInfo() *chromiumos.PackageInfo {
	if m != nil {
		return m.PackageInfo
	}
	return nil
}

func (m *PackageDepInfo) GetDependencyPackages() []*chromiumos.PackageInfo {
	if m != nil {
		return m.DependencyPackages
	}
	return nil
}

func (m *PackageDepInfo) GetDependencySourcePaths() []*SourcePath {
	if m != nil {
		return m.DependencySourcePaths
	}
	return nil
}

// Message describes a source path.
type SourcePath struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourcePath) Reset()         { *m = SourcePath{} }
func (m *SourcePath) String() string { return proto.CompactTextString(m) }
func (*SourcePath) ProtoMessage()    {}
func (*SourcePath) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a5308b238ddba97, []int{2}
}

func (m *SourcePath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourcePath.Unmarshal(m, b)
}
func (m *SourcePath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourcePath.Marshal(b, m, deterministic)
}
func (m *SourcePath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourcePath.Merge(m, src)
}
func (m *SourcePath) XXX_Size() int {
	return xxx_messageInfo_SourcePath.Size(m)
}
func (m *SourcePath) XXX_DiscardUnknown() {
	xxx_messageInfo_SourcePath.DiscardUnknown(m)
}

var xxx_messageInfo_SourcePath proto.InternalMessageInfo

func (m *SourcePath) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// Message for creating dependency graph json map.
// Example Json:
// {"build_target":{"name":"board"},"output_file":"/tmp/depgraph.json"}
type GetBuildDependencyGraphRequest struct {
	// Required.
	// The board whose dependency graph is being created.
	BuildTarget *chromiumos.BuildTarget `protobuf:"bytes,1,opt,name=build_target,json=buildTarget,proto3" json:"build_target,omitempty"`
	// Required.
	// The path where the dep graph json file is being saved to.
	OutputPath           string   `protobuf:"bytes,2,opt,name=output_path,json=outputPath,proto3" json:"output_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBuildDependencyGraphRequest) Reset()         { *m = GetBuildDependencyGraphRequest{} }
func (m *GetBuildDependencyGraphRequest) String() string { return proto.CompactTextString(m) }
func (*GetBuildDependencyGraphRequest) ProtoMessage()    {}
func (*GetBuildDependencyGraphRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a5308b238ddba97, []int{3}
}

func (m *GetBuildDependencyGraphRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBuildDependencyGraphRequest.Unmarshal(m, b)
}
func (m *GetBuildDependencyGraphRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBuildDependencyGraphRequest.Marshal(b, m, deterministic)
}
func (m *GetBuildDependencyGraphRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBuildDependencyGraphRequest.Merge(m, src)
}
func (m *GetBuildDependencyGraphRequest) XXX_Size() int {
	return xxx_messageInfo_GetBuildDependencyGraphRequest.Size(m)
}
func (m *GetBuildDependencyGraphRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBuildDependencyGraphRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBuildDependencyGraphRequest proto.InternalMessageInfo

func (m *GetBuildDependencyGraphRequest) GetBuildTarget() *chromiumos.BuildTarget {
	if m != nil {
		return m.BuildTarget
	}
	return nil
}

func (m *GetBuildDependencyGraphRequest) GetOutputPath() string {
	if m != nil {
		return m.OutputPath
	}
	return ""
}

// Response from creating dependency graph json map.
type GetBuildDependencyGraphResponse struct {
	// The created dependency graph file
	BuildDependencyGraphFile string   `protobuf:"bytes,1,opt,name=build_dependency_graph_file,json=buildDependencyGraphFile,proto3" json:"build_dependency_graph_file,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *GetBuildDependencyGraphResponse) Reset()         { *m = GetBuildDependencyGraphResponse{} }
func (m *GetBuildDependencyGraphResponse) String() string { return proto.CompactTextString(m) }
func (*GetBuildDependencyGraphResponse) ProtoMessage()    {}
func (*GetBuildDependencyGraphResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a5308b238ddba97, []int{4}
}

func (m *GetBuildDependencyGraphResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBuildDependencyGraphResponse.Unmarshal(m, b)
}
func (m *GetBuildDependencyGraphResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBuildDependencyGraphResponse.Marshal(b, m, deterministic)
}
func (m *GetBuildDependencyGraphResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBuildDependencyGraphResponse.Merge(m, src)
}
func (m *GetBuildDependencyGraphResponse) XXX_Size() int {
	return xxx_messageInfo_GetBuildDependencyGraphResponse.Size(m)
}
func (m *GetBuildDependencyGraphResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBuildDependencyGraphResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBuildDependencyGraphResponse proto.InternalMessageInfo

func (m *GetBuildDependencyGraphResponse) GetBuildDependencyGraphFile() string {
	if m != nil {
		return m.BuildDependencyGraphFile
	}
	return ""
}

func init() {
	proto.RegisterType((*DepGraph)(nil), "chromite.api.DepGraph")
	proto.RegisterType((*PackageDepInfo)(nil), "chromite.api.PackageDepInfo")
	proto.RegisterType((*SourcePath)(nil), "chromite.api.SourcePath")
	proto.RegisterType((*GetBuildDependencyGraphRequest)(nil), "chromite.api.GetBuildDependencyGraphRequest")
	proto.RegisterType((*GetBuildDependencyGraphResponse)(nil), "chromite.api.GetBuildDependencyGraphResponse")
}

func init() { proto.RegisterFile("chromite/api/depgraph.proto", fileDescriptor_7a5308b238ddba97) }

var fileDescriptor_7a5308b238ddba97 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0x25, 0x3b, 0xb2, 0x68, 0x65, 0x10, 0x6d, 0xd1, 0x09, 0xd9, 0x65, 0x77, 0xc8, 0x69, 0x11,
	0x4d, 0x60, 0x04, 0x11, 0x41, 0x84, 0x75, 0x70, 0xf5, 0x36, 0x64, 0x3d, 0x79, 0x89, 0x9d, 0xa4,
	0x26, 0xd3, 0x38, 0x93, 0x6a, 0xbb, 0x3b, 0x0b, 0x1e, 0xfc, 0x22, 0x8f, 0xfe, 0x81, 0x5f, 0xe3,
	0xc5, 0x7f, 0x90, 0x74, 0x32, 0x93, 0x44, 0x76, 0x44, 0xd8, 0x5b, 0xba, 0xde, 0xab, 0xf7, 0xea,
	0x55, 0xa7, 0xe1, 0x28, 0x5b, 0x29, 0xda, 0x08, 0x83, 0x11, 0x97, 0x22, 0xca, 0x51, 0x16, 0x8a,
	0xcb, 0x55, 0x28, 0x15, 0x19, 0x62, 0xe3, 0x2d, 0x18, 0x72, 0x29, 0xfc, 0xe3, 0x01, 0x35, 0xad,
	0xc4, 0x3a, 0x4f, 0xb8, 0x14, 0x0d, 0xd7, 0x9f, 0x34, 0x68, 0xb5, 0x21, 0x1d, 0x65, 0xb4, 0xd9,
	0x50, 0xd9, 0x00, 0xc1, 0x0f, 0x07, 0x6e, 0xcf, 0x51, 0x5e, 0xd4, 0xba, 0xec, 0x25, 0x8c, 0x9b,
	0x46, 0xc3, 0x55, 0x81, 0xc6, 0x73, 0xa6, 0xce, 0x99, 0x3b, 0x9b, 0x84, 0x5d, 0x73, 0x78, 0x5e,
	0xe3, 0x1f, 0x2c, 0x1c, 0xbb, 0x69, 0x77, 0x60, 0xaf, 0x61, 0x2c, 0x79, 0xf6, 0x99, 0x17, 0x98,
	0xe4, 0x28, 0xb5, 0x77, 0x30, 0x1d, 0x9d, 0xb9, 0xb3, 0xe3, 0xb0, 0x3f, 0x64, 0xb8, 0x68, 0x18,
	0x73, 0x94, 0xef, 0xcb, 0x25, 0xc5, 0xae, 0xdc, 0x9d, 0x35, 0x7b, 0x0c, 0x87, 0x35, 0x97, 0x8c,
	0x37, 0xb2, 0xb6, 0xac, 0x6f, 0xfb, 0xc6, 0x22, 0x71, 0xcb, 0x08, 0x7e, 0x39, 0x70, 0x77, 0xa8,
	0x55, 0xcf, 0xbe, 0xf5, 0x17, 0xe5, 0x92, 0xae, 0x9b, 0xbd, 0xed, 0x18, 0x58, 0xdb, 0xde, 0x77,
	0xf0, 0x20, 0x47, 0x89, 0x65, 0x8e, 0x65, 0xf6, 0x35, 0x69, 0x91, 0x6d, 0x84, 0xbd, 0x12, 0xac,
	0xeb, 0x69, 0xcb, 0x9a, 0x2d, 0x60, 0xd2, 0x53, 0xd2, 0x54, 0xa9, 0x0c, 0x13, 0xc9, 0xcd, 0x4a,
	0x7b, 0x23, 0xab, 0xe6, 0x0d, 0x17, 0x72, 0x69, 0x19, 0x0b, 0x6e, 0x56, 0xf1, 0xc3, 0xae, 0xb1,
	0xab, 0xea, 0x60, 0x0a, 0xd0, 0x1d, 0x19, 0x83, 0x5b, 0xb5, 0x9a, 0x4d, 0x77, 0x27, 0xb6, 0xdf,
	0xc1, 0x37, 0x38, 0xb9, 0x40, 0x63, 0x2f, 0x66, 0xbe, 0x93, 0xb0, 0x17, 0x1a, 0xe3, 0x97, 0x0a,
	0xb5, 0xb9, 0xd1, 0xbd, 0x9e, 0x82, 0x4b, 0x95, 0x91, 0x95, 0xb1, 0x31, 0xbc, 0x03, 0x6b, 0x0c,
	0x4d, 0xa9, 0x1e, 0x29, 0xf8, 0x04, 0xa7, 0x7b, 0xed, 0xb5, 0xa4, 0x52, 0x23, 0x7b, 0x05, 0x47,
	0x8d, 0x7f, 0x6f, 0x37, 0xf6, 0x4f, 0x4e, 0x96, 0x62, 0x8d, 0x6d, 0x18, 0x2f, 0xbd, 0x46, 0xe2,
	0xad, 0x58, 0xe3, 0xec, 0xbb, 0x03, 0xf7, 0xbb, 0xfa, 0x25, 0xaa, 0x2b, 0x91, 0x21, 0xbb, 0x82,
	0xc9, 0x1e, 0x5f, 0xf6, 0x64, 0xb8, 0xe4, 0x7f, 0x6f, 0xc7, 0x7f, 0xfa, 0x9f, 0xec, 0x26, 0x8c,
	0x7f, 0xf2, 0xf3, 0xb7, 0xef, 0xc3, 0xa3, 0x94, 0xb8, 0xca, 0x93, 0xbf, 0x63, 0xdd, 0x73, 0xce,
	0x5f, 0x7c, 0x7c, 0x5e, 0xd0, 0x6e, 0xb5, 0x21, 0xa9, 0x22, 0xea, 0x3d, 0x3e, 0x51, 0x2e, 0x15,
	0x8f, 0xec, 0xdb, 0x8b, 0x0a, 0x8a, 0xfa, 0x8f, 0x36, 0x3d, 0xb4, 0xe5, 0x67, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x54, 0x1e, 0xa7, 0xa7, 0xf6, 0x03, 0x00, 0x00,
}
