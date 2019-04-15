// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testplans/generate_test_plan.proto

package testplans

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

type ProtoBytes struct {
	SerializedProto      []byte   `protobuf:"bytes,1,opt,name=serialized_proto,json=serializedProto,proto3" json:"serialized_proto,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoBytes) Reset()         { *m = ProtoBytes{} }
func (m *ProtoBytes) String() string { return proto.CompactTextString(m) }
func (*ProtoBytes) ProtoMessage()    {}
func (*ProtoBytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8083f85c67203f6, []int{0}
}

func (m *ProtoBytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoBytes.Unmarshal(m, b)
}
func (m *ProtoBytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoBytes.Marshal(b, m, deterministic)
}
func (m *ProtoBytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoBytes.Merge(m, src)
}
func (m *ProtoBytes) XXX_Size() int {
	return xxx_messageInfo_ProtoBytes.Size(m)
}
func (m *ProtoBytes) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoBytes.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoBytes proto.InternalMessageInfo

func (m *ProtoBytes) GetSerializedProto() []byte {
	if m != nil {
		return m.SerializedProto
	}
	return nil
}

type GenerateTestPlanRequest struct {
	// Serialized buildbucket Build protos that are part of this orchestrator run.
	// See https://chromium.googlesource.com/infra/luci/luci-go/+/master/buildbucket/proto/build.proto
	BuildbucketProtos []*ProtoBytes `protobuf:"bytes,5,rep,name=buildbucket_protos,json=buildbucketProtos,proto3" json:"buildbucket_protos,omitempty"`
	// Path on disk to the chromiumos repo checkout.
	ChromiumosCheckoutRoot string `protobuf:"bytes,4,opt,name=chromiumos_checkout_root,json=chromiumosCheckoutRoot,proto3" json:"chromiumos_checkout_root,omitempty"`
	// Path on disk to the "repo" CLI tool.
	RepoToolPath         string   `protobuf:"bytes,6,opt,name=repo_tool_path,json=repoToolPath,proto3" json:"repo_tool_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateTestPlanRequest) Reset()         { *m = GenerateTestPlanRequest{} }
func (m *GenerateTestPlanRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateTestPlanRequest) ProtoMessage()    {}
func (*GenerateTestPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8083f85c67203f6, []int{1}
}

func (m *GenerateTestPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateTestPlanRequest.Unmarshal(m, b)
}
func (m *GenerateTestPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateTestPlanRequest.Marshal(b, m, deterministic)
}
func (m *GenerateTestPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateTestPlanRequest.Merge(m, src)
}
func (m *GenerateTestPlanRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateTestPlanRequest.Size(m)
}
func (m *GenerateTestPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateTestPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateTestPlanRequest proto.InternalMessageInfo

func (m *GenerateTestPlanRequest) GetBuildbucketProtos() []*ProtoBytes {
	if m != nil {
		return m.BuildbucketProtos
	}
	return nil
}

func (m *GenerateTestPlanRequest) GetChromiumosCheckoutRoot() string {
	if m != nil {
		return m.ChromiumosCheckoutRoot
	}
	return ""
}

func (m *GenerateTestPlanRequest) GetRepoToolPath() string {
	if m != nil {
		return m.RepoToolPath
	}
	return ""
}

// The final test plan.
type GenerateTestPlanResponse struct {
	TestUnit             []*TestUnit `protobuf:"bytes,1,rep,name=test_unit,json=testUnit,proto3" json:"test_unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GenerateTestPlanResponse) Reset()         { *m = GenerateTestPlanResponse{} }
func (m *GenerateTestPlanResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateTestPlanResponse) ProtoMessage()    {}
func (*GenerateTestPlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8083f85c67203f6, []int{2}
}

func (m *GenerateTestPlanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateTestPlanResponse.Unmarshal(m, b)
}
func (m *GenerateTestPlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateTestPlanResponse.Marshal(b, m, deterministic)
}
func (m *GenerateTestPlanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateTestPlanResponse.Merge(m, src)
}
func (m *GenerateTestPlanResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateTestPlanResponse.Size(m)
}
func (m *GenerateTestPlanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateTestPlanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateTestPlanResponse proto.InternalMessageInfo

func (m *GenerateTestPlanResponse) GetTestUnit() []*TestUnit {
	if m != nil {
		return m.TestUnit
	}
	return nil
}

// The files that should be tested in a test plan.
type BuildPayload struct {
	// The GS bucket in which artifacts for the build are stored, e.g.
	// gs://chromeos-image-archive
	ArtifactsGsBucket string `protobuf:"bytes,1,opt,name=artifacts_gs_bucket,json=artifactsGsBucket,proto3" json:"artifacts_gs_bucket,omitempty"`
	// The path in the bucket in which artifacts for the build are stored, e.g.
	// eve-paladin/R73-11588.0.0-rc4
	ArtifactsGsPath      string   `protobuf:"bytes,2,opt,name=artifacts_gs_path,json=artifactsGsPath,proto3" json:"artifacts_gs_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildPayload) Reset()         { *m = BuildPayload{} }
func (m *BuildPayload) String() string { return proto.CompactTextString(m) }
func (*BuildPayload) ProtoMessage()    {}
func (*BuildPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8083f85c67203f6, []int{3}
}

func (m *BuildPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildPayload.Unmarshal(m, b)
}
func (m *BuildPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildPayload.Marshal(b, m, deterministic)
}
func (m *BuildPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildPayload.Merge(m, src)
}
func (m *BuildPayload) XXX_Size() int {
	return xxx_messageInfo_BuildPayload.Size(m)
}
func (m *BuildPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildPayload.DiscardUnknown(m)
}

var xxx_messageInfo_BuildPayload proto.InternalMessageInfo

func (m *BuildPayload) GetArtifactsGsBucket() string {
	if m != nil {
		return m.ArtifactsGsBucket
	}
	return ""
}

func (m *BuildPayload) GetArtifactsGsPath() string {
	if m != nil {
		return m.ArtifactsGsPath
	}
	return ""
}

// Hardware or VM requirements for running a test plan.
type SchedulingRequirements struct {
	// Types that are valid to be assigned to TargetType:
	//	*SchedulingRequirements_BuildTarget
	TargetType           isSchedulingRequirements_TargetType `protobuf_oneof:"target_type"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *SchedulingRequirements) Reset()         { *m = SchedulingRequirements{} }
func (m *SchedulingRequirements) String() string { return proto.CompactTextString(m) }
func (*SchedulingRequirements) ProtoMessage()    {}
func (*SchedulingRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8083f85c67203f6, []int{4}
}

func (m *SchedulingRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulingRequirements.Unmarshal(m, b)
}
func (m *SchedulingRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulingRequirements.Marshal(b, m, deterministic)
}
func (m *SchedulingRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulingRequirements.Merge(m, src)
}
func (m *SchedulingRequirements) XXX_Size() int {
	return xxx_messageInfo_SchedulingRequirements.Size(m)
}
func (m *SchedulingRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulingRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulingRequirements proto.InternalMessageInfo

type isSchedulingRequirements_TargetType interface {
	isSchedulingRequirements_TargetType()
}

type SchedulingRequirements_BuildTarget struct {
	BuildTarget string `protobuf:"bytes,2,opt,name=build_target,json=buildTarget,proto3,oneof"`
}

func (*SchedulingRequirements_BuildTarget) isSchedulingRequirements_TargetType() {}

func (m *SchedulingRequirements) GetTargetType() isSchedulingRequirements_TargetType {
	if m != nil {
		return m.TargetType
	}
	return nil
}

func (m *SchedulingRequirements) GetBuildTarget() string {
	if x, ok := m.GetTargetType().(*SchedulingRequirements_BuildTarget); ok {
		return x.BuildTarget
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SchedulingRequirements) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SchedulingRequirements_BuildTarget)(nil),
	}
}

// Metadata for executing one test configuration for one target type.
type TestUnit struct {
	// Types that are valid to be assigned to TestCfg:
	//	*TestUnit_GceTestCfg
	//	*TestUnit_HwTestCfg
	//	*TestUnit_MoblabVmTestCfg
	//	*TestUnit_TastVmTestCfg
	//	*TestUnit_VmTestCfg
	TestCfg isTestUnit_TestCfg `protobuf_oneof:"TestCfg"`
	// The scheduling requirements for this test plan.
	SchedulingRequirements *SchedulingRequirements `protobuf:"bytes,6,opt,name=scheduling_requirements,json=schedulingRequirements,proto3" json:"scheduling_requirements,omitempty"`
	// The build files provided to run this test plan.
	BuildPayload         *BuildPayload `protobuf:"bytes,7,opt,name=build_payload,json=buildPayload,proto3" json:"build_payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TestUnit) Reset()         { *m = TestUnit{} }
func (m *TestUnit) String() string { return proto.CompactTextString(m) }
func (*TestUnit) ProtoMessage()    {}
func (*TestUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8083f85c67203f6, []int{5}
}

func (m *TestUnit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestUnit.Unmarshal(m, b)
}
func (m *TestUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestUnit.Marshal(b, m, deterministic)
}
func (m *TestUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestUnit.Merge(m, src)
}
func (m *TestUnit) XXX_Size() int {
	return xxx_messageInfo_TestUnit.Size(m)
}
func (m *TestUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_TestUnit.DiscardUnknown(m)
}

var xxx_messageInfo_TestUnit proto.InternalMessageInfo

type isTestUnit_TestCfg interface {
	isTestUnit_TestCfg()
}

type TestUnit_GceTestCfg struct {
	GceTestCfg *GceTestCfg `protobuf:"bytes,1,opt,name=gce_test_cfg,json=gceTestCfg,proto3,oneof"`
}

type TestUnit_HwTestCfg struct {
	HwTestCfg *HwTestCfg `protobuf:"bytes,2,opt,name=hw_test_cfg,json=hwTestCfg,proto3,oneof"`
}

type TestUnit_MoblabVmTestCfg struct {
	MoblabVmTestCfg *MoblabVmTestCfg `protobuf:"bytes,3,opt,name=moblab_vm_test_cfg,json=moblabVmTestCfg,proto3,oneof"`
}

type TestUnit_TastVmTestCfg struct {
	TastVmTestCfg *TastVmTestCfg `protobuf:"bytes,4,opt,name=tast_vm_test_cfg,json=tastVmTestCfg,proto3,oneof"`
}

type TestUnit_VmTestCfg struct {
	VmTestCfg *VmTestCfg `protobuf:"bytes,5,opt,name=vm_test_cfg,json=vmTestCfg,proto3,oneof"`
}

func (*TestUnit_GceTestCfg) isTestUnit_TestCfg() {}

func (*TestUnit_HwTestCfg) isTestUnit_TestCfg() {}

func (*TestUnit_MoblabVmTestCfg) isTestUnit_TestCfg() {}

func (*TestUnit_TastVmTestCfg) isTestUnit_TestCfg() {}

func (*TestUnit_VmTestCfg) isTestUnit_TestCfg() {}

func (m *TestUnit) GetTestCfg() isTestUnit_TestCfg {
	if m != nil {
		return m.TestCfg
	}
	return nil
}

func (m *TestUnit) GetGceTestCfg() *GceTestCfg {
	if x, ok := m.GetTestCfg().(*TestUnit_GceTestCfg); ok {
		return x.GceTestCfg
	}
	return nil
}

func (m *TestUnit) GetHwTestCfg() *HwTestCfg {
	if x, ok := m.GetTestCfg().(*TestUnit_HwTestCfg); ok {
		return x.HwTestCfg
	}
	return nil
}

func (m *TestUnit) GetMoblabVmTestCfg() *MoblabVmTestCfg {
	if x, ok := m.GetTestCfg().(*TestUnit_MoblabVmTestCfg); ok {
		return x.MoblabVmTestCfg
	}
	return nil
}

func (m *TestUnit) GetTastVmTestCfg() *TastVmTestCfg {
	if x, ok := m.GetTestCfg().(*TestUnit_TastVmTestCfg); ok {
		return x.TastVmTestCfg
	}
	return nil
}

func (m *TestUnit) GetVmTestCfg() *VmTestCfg {
	if x, ok := m.GetTestCfg().(*TestUnit_VmTestCfg); ok {
		return x.VmTestCfg
	}
	return nil
}

func (m *TestUnit) GetSchedulingRequirements() *SchedulingRequirements {
	if m != nil {
		return m.SchedulingRequirements
	}
	return nil
}

func (m *TestUnit) GetBuildPayload() *BuildPayload {
	if m != nil {
		return m.BuildPayload
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestUnit) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestUnit_GceTestCfg)(nil),
		(*TestUnit_HwTestCfg)(nil),
		(*TestUnit_MoblabVmTestCfg)(nil),
		(*TestUnit_TastVmTestCfg)(nil),
		(*TestUnit_VmTestCfg)(nil),
	}
}

func init() {
	proto.RegisterType((*ProtoBytes)(nil), "testplans.ProtoBytes")
	proto.RegisterType((*GenerateTestPlanRequest)(nil), "testplans.GenerateTestPlanRequest")
	proto.RegisterType((*GenerateTestPlanResponse)(nil), "testplans.GenerateTestPlanResponse")
	proto.RegisterType((*BuildPayload)(nil), "testplans.BuildPayload")
	proto.RegisterType((*SchedulingRequirements)(nil), "testplans.SchedulingRequirements")
	proto.RegisterType((*TestUnit)(nil), "testplans.TestUnit")
}

func init() { proto.RegisterFile("testplans/generate_test_plan.proto", fileDescriptor_e8083f85c67203f6) }

var fileDescriptor_e8083f85c67203f6 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdd, 0x4e, 0xdb, 0x30,
	0x18, 0xa5, 0x3f, 0x40, 0xfb, 0xb5, 0x8c, 0x62, 0x36, 0x88, 0xb8, 0x62, 0xd9, 0x2e, 0xd8, 0x2e,
	0xda, 0x09, 0xb4, 0x3f, 0x69, 0x57, 0x65, 0x12, 0x0c, 0x31, 0xa9, 0xf2, 0xd8, 0x2e, 0xb8, 0xb1,
	0xdc, 0xd4, 0x75, 0x32, 0x92, 0x38, 0xd8, 0x5f, 0x40, 0xec, 0x15, 0xf7, 0x26, 0x7b, 0x8a, 0x29,
	0x4e, 0x9b, 0xb8, 0x82, 0xbb, 0xf8, 0xf8, 0x9c, 0x53, 0x9f, 0xef, 0xa7, 0xe0, 0xa3, 0x30, 0x98,
	0xc5, 0x3c, 0x35, 0x23, 0x29, 0x52, 0xa1, 0x39, 0x0a, 0x56, 0x40, 0xac, 0xc0, 0x86, 0x99, 0x56,
	0xa8, 0x48, 0xb7, 0xe2, 0x1c, 0x8c, 0x6a, 0x3a, 0x72, 0x2d, 0x05, 0x96, 0x64, 0x2d, 0x6e, 0xf3,
	0x48, 0x8b, 0x44, 0xa4, 0x68, 0x58, 0xa0, 0xd2, 0x79, 0x24, 0x4b, 0xad, 0xff, 0x11, 0x60, 0x52,
	0x7c, 0x8c, 0x1f, 0x50, 0x18, 0xf2, 0x06, 0x06, 0x46, 0xe8, 0x88, 0xc7, 0xd1, 0x1f, 0x31, 0x63,
	0x96, 0xe1, 0x35, 0x0e, 0x1b, 0x47, 0x7d, 0xba, 0x5d, 0xe3, 0x96, 0xef, 0xff, 0x6d, 0xc0, 0xfe,
	0xd9, 0xe2, 0x45, 0x57, 0xc2, 0xe0, 0x24, 0xe6, 0x29, 0x15, 0xb7, 0xb9, 0x30, 0x48, 0xbe, 0x02,
	0x99, 0xe6, 0x51, 0x3c, 0x9b, 0xe6, 0xc1, 0x8d, 0xc0, 0xd2, 0xc7, 0x78, 0xeb, 0x87, 0xad, 0xa3,
	0xde, 0xf1, 0x8b, 0x61, 0xf5, 0xc4, 0x61, 0xfd, 0xcb, 0x74, 0xc7, 0x11, 0x58, 0xd8, 0x90, 0x4f,
	0xe0, 0x05, 0xa1, 0x56, 0x49, 0x94, 0x27, 0xca, 0xb0, 0x20, 0x14, 0xc1, 0x8d, 0xca, 0x91, 0x69,
	0xa5, 0xd0, 0x6b, 0x1f, 0x36, 0x8e, 0xba, 0x74, 0xaf, 0xbe, 0x3f, 0x5d, 0x5c, 0x53, 0xa5, 0x90,
	0xbc, 0x86, 0x67, 0x5a, 0x64, 0x8a, 0xa1, 0x52, 0x31, 0xcb, 0x38, 0x86, 0xde, 0x86, 0xe5, 0xf7,
	0x0b, 0xf4, 0x4a, 0xa9, 0x78, 0xc2, 0x31, 0xbc, 0x68, 0x77, 0x1a, 0x83, 0xe6, 0x45, 0xbb, 0xd3,
	0x1c, 0xb4, 0xfc, 0x4b, 0xf0, 0x1e, 0x87, 0x31, 0x99, 0x4a, 0x8d, 0x20, 0xef, 0xc0, 0x16, 0x98,
	0xe5, 0x69, 0x84, 0x5e, 0xc3, 0x86, 0xd8, 0x75, 0x42, 0x14, 0xfc, 0x9f, 0x69, 0x84, 0xb4, 0x83,
	0x8b, 0x2f, 0xff, 0x37, 0xf4, 0xc7, 0x45, 0x9c, 0x09, 0x7f, 0x88, 0x15, 0x9f, 0x91, 0x21, 0xec,
	0x72, 0x8d, 0xd1, 0x9c, 0x07, 0x68, 0x98, 0x34, 0xac, 0x8c, 0x69, 0x2b, 0xdb, 0xa5, 0x3b, 0xd5,
	0xd5, 0x99, 0x19, 0xdb, 0x0b, 0xf2, 0x16, 0x76, 0x56, 0xf8, 0x36, 0x42, 0xd3, 0xb2, 0xb7, 0x1d,
	0x76, 0x91, 0xc2, 0xbf, 0x84, 0xbd, 0x1f, 0x41, 0x28, 0x66, 0x79, 0x1c, 0xa5, 0x92, 0x3a, 0x7d,
	0x26, 0xaf, 0xa0, 0x6f, 0x8b, 0xca, 0xca, 0x49, 0x28, 0x0d, 0xce, 0xd7, 0x68, 0xcf, 0xa2, 0x57,
	0x16, 0x1c, 0x6f, 0x41, 0x6f, 0x39, 0x28, 0x0f, 0x99, 0xf0, 0xff, 0xb5, 0xa0, 0xb3, 0x0c, 0x44,
	0x3e, 0x43, 0x5f, 0x06, 0x8b, 0x71, 0x0b, 0xe6, 0xd2, 0xbe, 0x77, 0xb5, 0x81, 0x67, 0x81, 0x2d,
	0xd7, 0xe9, 0x5c, 0x9e, 0xaf, 0x51, 0x90, 0xd5, 0x89, 0x7c, 0x80, 0x5e, 0x78, 0x5f, 0x2b, 0x9b,
	0x56, 0xf9, 0xdc, 0x51, 0x9e, 0xdf, 0xd7, 0xc2, 0x6e, 0xb8, 0x3c, 0x90, 0x6f, 0x40, 0x12, 0x35,
	0x8d, 0xf9, 0x94, 0xdd, 0x25, 0xb5, 0xbc, 0x65, 0xe5, 0x07, 0x8e, 0xfc, 0xbb, 0x25, 0xfd, 0x4a,
	0x6a, 0x93, 0xed, 0x64, 0x15, 0x22, 0xa7, 0x30, 0x40, 0x6e, 0x70, 0xc5, 0xa8, 0x6d, 0x8d, 0x3c,
	0xb7, 0x7b, 0xdc, 0xa0, 0x6b, 0xb3, 0x85, 0x2e, 0x50, 0xe4, 0x70, 0xf5, 0xeb, 0x8f, 0x72, 0xb8,
	0xda, 0xee, 0x5d, 0xa5, 0xbb, 0x86, 0x7d, 0x53, 0x75, 0x65, 0x65, 0xfd, 0xec, 0x28, 0xf6, 0x8e,
	0x5f, 0x3a, 0x1e, 0x4f, 0xf7, 0x8f, 0xee, 0x99, 0xa7, 0xfb, 0xfa, 0x05, 0xb6, 0xca, 0xbe, 0x66,
	0xe5, 0x78, 0x79, 0x9b, 0xd6, 0x71, 0xdf, 0x71, 0x74, 0xa7, 0x8f, 0x96, 0x53, 0xb0, 0x38, 0x8d,
	0xbb, 0xb0, 0xb9, 0x78, 0xe4, 0xf8, 0xfd, 0xf5, 0x89, 0x54, 0xc3, 0xe5, 0x0e, 0x0d, 0x95, 0x96,
	0xa3, 0x7a, 0xa1, 0x46, 0x51, 0x3a, 0xd7, 0x7c, 0x64, 0x77, 0x77, 0x24, 0x55, 0xfd, 0xc7, 0x32,
	0xdd, 0xb0, 0xd8, 0xc9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xc7, 0xb0, 0xdf, 0x9b, 0x04,
	0x00, 0x00,
}
