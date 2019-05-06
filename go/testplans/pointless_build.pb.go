// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testplans/pointless_build.proto

package testplans

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	api "go.chromium.org/chromiumos/infra/proto/go/chromite/api"
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

type PointlessBuildCheckResponse_PointlessBuildReason int32

const (
	PointlessBuildCheckResponse_POINTLESS_BUILD_REASON_UNSPECIFIED          PointlessBuildCheckResponse_PointlessBuildReason = 0
	PointlessBuildCheckResponse_IRRELEVANT_TO_DEPS_GRAPH                    PointlessBuildCheckResponse_PointlessBuildReason = 1
	PointlessBuildCheckResponse_IRRELEVANT_TO_KNOWN_NON_PORTAGE_DIRECTORIES PointlessBuildCheckResponse_PointlessBuildReason = 2
)

var PointlessBuildCheckResponse_PointlessBuildReason_name = map[int32]string{
	0: "POINTLESS_BUILD_REASON_UNSPECIFIED",
	1: "IRRELEVANT_TO_DEPS_GRAPH",
	2: "IRRELEVANT_TO_KNOWN_NON_PORTAGE_DIRECTORIES",
}

var PointlessBuildCheckResponse_PointlessBuildReason_value = map[string]int32{
	"POINTLESS_BUILD_REASON_UNSPECIFIED":          0,
	"IRRELEVANT_TO_DEPS_GRAPH":                    1,
	"IRRELEVANT_TO_KNOWN_NON_PORTAGE_DIRECTORIES": 2,
}

func (x PointlessBuildCheckResponse_PointlessBuildReason) String() string {
	return proto.EnumName(PointlessBuildCheckResponse_PointlessBuildReason_name, int32(x))
}

func (PointlessBuildCheckResponse_PointlessBuildReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65492a4cea7c5e67, []int{1, 0}
}

// The input proto for invocations of the Pointless Build Checker program.
type PointlessBuildCheckRequest struct {
	// Serialized buildbucket Build proto for this child builder run.
	// See https://chromium.googlesource.com/infra/luci/luci-go/+/master/buildbucket/proto/build.proto
	BuildbucketProto *ProtoBytes `protobuf:"bytes,1,opt,name=buildbucket_proto,json=buildbucketProto,proto3" json:"buildbucket_proto,omitempty"`
	// Path on disk to a JSON-encoded chromite.api.DepGraph for this child
	// builder.
	// DEPRECATED: Use the DepGraph message below.
	DepGraphPath string `protobuf:"bytes,2,opt,name=dep_graph_path,json=depGraphPath,proto3" json:"dep_graph_path,omitempty"`
	// The DepGraph for this child builder.
	DepGraph *api.DepGraph `protobuf:"bytes,5,opt,name=dep_graph,json=depGraph,proto3" json:"dep_graph,omitempty"`
	// Path on disk to the chromiumos workspace repo checkout.
	ChromiumosWorkspaceCheckoutRoot string `protobuf:"bytes,3,opt,name=chromiumos_workspace_checkout_root,json=chromiumosWorkspaceCheckoutRoot,proto3" json:"chromiumos_workspace_checkout_root,omitempty"`
	// Path on disk to the "repo" CLI tool.
	RepoToolPath         string   `protobuf:"bytes,4,opt,name=repo_tool_path,json=repoToolPath,proto3" json:"repo_tool_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PointlessBuildCheckRequest) Reset()         { *m = PointlessBuildCheckRequest{} }
func (m *PointlessBuildCheckRequest) String() string { return proto.CompactTextString(m) }
func (*PointlessBuildCheckRequest) ProtoMessage()    {}
func (*PointlessBuildCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65492a4cea7c5e67, []int{0}
}

func (m *PointlessBuildCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointlessBuildCheckRequest.Unmarshal(m, b)
}
func (m *PointlessBuildCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointlessBuildCheckRequest.Marshal(b, m, deterministic)
}
func (m *PointlessBuildCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointlessBuildCheckRequest.Merge(m, src)
}
func (m *PointlessBuildCheckRequest) XXX_Size() int {
	return xxx_messageInfo_PointlessBuildCheckRequest.Size(m)
}
func (m *PointlessBuildCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PointlessBuildCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PointlessBuildCheckRequest proto.InternalMessageInfo

func (m *PointlessBuildCheckRequest) GetBuildbucketProto() *ProtoBytes {
	if m != nil {
		return m.BuildbucketProto
	}
	return nil
}

func (m *PointlessBuildCheckRequest) GetDepGraphPath() string {
	if m != nil {
		return m.DepGraphPath
	}
	return ""
}

func (m *PointlessBuildCheckRequest) GetDepGraph() *api.DepGraph {
	if m != nil {
		return m.DepGraph
	}
	return nil
}

func (m *PointlessBuildCheckRequest) GetChromiumosWorkspaceCheckoutRoot() string {
	if m != nil {
		return m.ChromiumosWorkspaceCheckoutRoot
	}
	return ""
}

func (m *PointlessBuildCheckRequest) GetRepoToolPath() string {
	if m != nil {
		return m.RepoToolPath
	}
	return ""
}

// The output proto for invocations of the Pointless Build Checker program.
type PointlessBuildCheckResponse struct {
	// Whether the build is pointless and can be terminated without proceeding to
	// building packages and testing.
	BuildIsPointless *wrappers.BoolValue `protobuf:"bytes,1,opt,name=build_is_pointless,json=buildIsPointless,proto3" json:"build_is_pointless,omitempty"`
	// If build_is_pointless, this is the reason that the Pointless Build Checker
	// came to that conclusion. Otherwise, this is unspecified.
	PointlessBuildReason PointlessBuildCheckResponse_PointlessBuildReason `protobuf:"varint,2,opt,name=pointless_build_reason,json=pointlessBuildReason,proto3,enum=testplans.PointlessBuildCheckResponse_PointlessBuildReason" json:"pointless_build_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *PointlessBuildCheckResponse) Reset()         { *m = PointlessBuildCheckResponse{} }
func (m *PointlessBuildCheckResponse) String() string { return proto.CompactTextString(m) }
func (*PointlessBuildCheckResponse) ProtoMessage()    {}
func (*PointlessBuildCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65492a4cea7c5e67, []int{1}
}

func (m *PointlessBuildCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointlessBuildCheckResponse.Unmarshal(m, b)
}
func (m *PointlessBuildCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointlessBuildCheckResponse.Marshal(b, m, deterministic)
}
func (m *PointlessBuildCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointlessBuildCheckResponse.Merge(m, src)
}
func (m *PointlessBuildCheckResponse) XXX_Size() int {
	return xxx_messageInfo_PointlessBuildCheckResponse.Size(m)
}
func (m *PointlessBuildCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PointlessBuildCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PointlessBuildCheckResponse proto.InternalMessageInfo

func (m *PointlessBuildCheckResponse) GetBuildIsPointless() *wrappers.BoolValue {
	if m != nil {
		return m.BuildIsPointless
	}
	return nil
}

func (m *PointlessBuildCheckResponse) GetPointlessBuildReason() PointlessBuildCheckResponse_PointlessBuildReason {
	if m != nil {
		return m.PointlessBuildReason
	}
	return PointlessBuildCheckResponse_POINTLESS_BUILD_REASON_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("testplans.PointlessBuildCheckResponse_PointlessBuildReason", PointlessBuildCheckResponse_PointlessBuildReason_name, PointlessBuildCheckResponse_PointlessBuildReason_value)
	proto.RegisterType((*PointlessBuildCheckRequest)(nil), "testplans.PointlessBuildCheckRequest")
	proto.RegisterType((*PointlessBuildCheckResponse)(nil), "testplans.PointlessBuildCheckResponse")
}

func init() { proto.RegisterFile("testplans/pointless_build.proto", fileDescriptor_65492a4cea7c5e67) }

var fileDescriptor_65492a4cea7c5e67 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xdd, 0x6a, 0xdb, 0x30,
	0x18, 0x5d, 0xb2, 0x1f, 0x56, 0x6d, 0x94, 0xcc, 0x74, 0x25, 0xa4, 0x63, 0x2d, 0x61, 0x8c, 0xc2,
	0x40, 0x86, 0x86, 0x5d, 0xed, 0x2a, 0x4e, 0xbc, 0xd4, 0x34, 0xd8, 0x46, 0x76, 0x5b, 0xd8, 0x8d,
	0x50, 0x1c, 0xd5, 0x31, 0x71, 0xfc, 0xa9, 0x92, 0x4c, 0xd9, 0x43, 0xec, 0x55, 0xf6, 0x60, 0x7b,
	0x8a, 0x61, 0x39, 0x76, 0xba, 0x92, 0x5d, 0xea, 0x7c, 0x47, 0xdf, 0x39, 0x3a, 0x47, 0xe8, 0x54,
	0x73, 0xa5, 0x45, 0xce, 0x0a, 0x65, 0x0b, 0xc8, 0x0a, 0x9d, 0x73, 0xa5, 0xe8, 0xa2, 0xcc, 0xf2,
	0x25, 0x16, 0x12, 0x34, 0x58, 0x07, 0x2d, 0x61, 0x70, 0x92, 0xac, 0x24, 0x6c, 0x32, 0xcd, 0x6d,
	0x26, 0x32, 0x7b, 0xc9, 0x45, 0x2a, 0x99, 0x58, 0xd5, 0xbc, 0xc1, 0xf1, 0x6e, 0x51, 0x02, 0x9b,
	0x0d, 0x14, 0x5b, 0xfc, 0x63, 0x0a, 0x90, 0xe6, 0xdc, 0x36, 0xa7, 0x45, 0x79, 0x67, 0x3f, 0x48,
	0x26, 0x04, 0x97, 0xaa, 0x9e, 0x0f, 0x7f, 0x77, 0xd1, 0x20, 0x6c, 0x94, 0x9d, 0x4a, 0x78, 0xb2,
	0xe2, 0xc9, 0x9a, 0xf0, 0xfb, 0x92, 0x2b, 0x6d, 0x39, 0xe8, 0x9d, 0x71, 0xb3, 0x28, 0x93, 0x35,
	0xd7, 0xd4, 0xdc, 0xe9, 0x77, 0xce, 0x3a, 0xe7, 0x6f, 0x2e, 0xde, 0xe3, 0x56, 0x12, 0x87, 0x15,
	0xee, 0xfc, 0xd4, 0x5c, 0x91, 0xde, 0x23, 0xbe, 0x81, 0xad, 0x4f, 0xe8, 0x70, 0xc9, 0x05, 0x35,
	0x6e, 0xa9, 0x60, 0x7a, 0xd5, 0xef, 0x9e, 0x75, 0xce, 0x0f, 0xc8, 0xdb, 0x25, 0x17, 0xb3, 0x0a,
	0x0c, 0x99, 0x5e, 0x59, 0x23, 0x74, 0xd0, 0xb2, 0xfa, 0x2f, 0x8d, 0xc2, 0x31, 0x6e, 0x5e, 0x8c,
	0x99, 0xc8, 0xf0, 0x74, 0x4b, 0x27, 0xaf, 0x9b, 0x8b, 0xd6, 0x15, 0x1a, 0xd6, 0x94, 0x72, 0x03,
	0x8a, 0x3e, 0x80, 0x5c, 0x2b, 0xc1, 0x12, 0x4e, 0x93, 0xea, 0x0d, 0x50, 0x6a, 0x2a, 0x01, 0x74,
	0xff, 0xb9, 0x91, 0x3b, 0xdd, 0x31, 0x6f, 0x1b, 0xe2, 0x64, 0xcb, 0x23, 0x00, 0xba, 0xf2, 0x29,
	0xb9, 0x00, 0xaa, 0x01, 0xf2, 0xda, 0xe7, 0x8b, 0xda, 0x67, 0x85, 0xc6, 0x00, 0x79, 0xe5, 0x73,
	0xf8, 0xa7, 0x8b, 0x4e, 0xf6, 0x06, 0xa6, 0x04, 0x14, 0x8a, 0x5b, 0x97, 0xc8, 0x32, 0x09, 0xd0,
	0x4c, 0xd1, 0xb6, 0xd2, 0x6d, 0x64, 0x03, 0x5c, 0xb7, 0x81, 0x9b, 0x36, 0xb0, 0x03, 0x90, 0xdf,
	0xb0, 0xbc, 0xe4, 0xdb, 0xdc, 0x3c, 0xd5, 0xee, 0xb6, 0xee, 0xd1, 0xf1, 0x93, 0x3f, 0x41, 0x25,
	0x67, 0x0a, 0x0a, 0x93, 0xdf, 0xe1, 0xc5, 0xb7, 0xc7, 0x05, 0xfc, 0xdf, 0xd1, 0x93, 0x19, 0x31,
	0x2b, 0xc8, 0x91, 0xd8, 0x83, 0x0e, 0x7f, 0x75, 0xd0, 0xd1, 0x3e, 0xba, 0xf5, 0x19, 0x0d, 0xc3,
	0xc0, 0xf3, 0xe3, 0xb9, 0x1b, 0x45, 0xd4, 0xb9, 0xf6, 0xe6, 0x53, 0x4a, 0xdc, 0x71, 0x14, 0xf8,
	0xf4, 0xda, 0x8f, 0x42, 0x77, 0xe2, 0x7d, 0xf7, 0xdc, 0x69, 0xef, 0x99, 0xf5, 0x01, 0xf5, 0x3d,
	0x42, 0xdc, 0xb9, 0x7b, 0x33, 0xf6, 0x63, 0x1a, 0x07, 0x74, 0xea, 0x86, 0x11, 0x9d, 0x91, 0x71,
	0x78, 0xd9, 0xeb, 0x58, 0x36, 0xfa, 0xf2, 0xef, 0xf4, 0xca, 0x0f, 0x6e, 0x7d, 0xea, 0x07, 0x3e,
	0x0d, 0x03, 0x12, 0x8f, 0x67, 0x2e, 0x9d, 0x7a, 0xc4, 0x9d, 0xc4, 0x01, 0xf1, 0xdc, 0xa8, 0xd7,
	0x75, 0xbe, 0xfe, 0x18, 0xa5, 0x80, 0x9b, 0xe2, 0x30, 0xc8, 0xd4, 0xde, 0xb5, 0x68, 0x67, 0xc5,
	0x9d, 0x64, 0xf5, 0xcf, 0xb6, 0x53, 0xb0, 0xdb, 0x34, 0x16, 0xaf, 0x0c, 0x36, 0xfa, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xe5, 0x5e, 0x2d, 0x40, 0x5e, 0x03, 0x00, 0x00,
}
