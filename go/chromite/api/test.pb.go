// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chromite/api/test.proto

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

// The program that manages and runs the tests in the VM.
type VmTestRequest_TestHarness int32

const (
	// No test harness specified.
	VmTestRequest_UNSPECIFIED VmTestRequest_TestHarness = 0
	// Run tests with Tast.
	VmTestRequest_TAST VmTestRequest_TestHarness = 1
	// Run tests with Autotest.
	VmTestRequest_AUTOTEST VmTestRequest_TestHarness = 2
)

var VmTestRequest_TestHarness_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "TAST",
	2: "AUTOTEST",
}

var VmTestRequest_TestHarness_value = map[string]int32{
	"UNSPECIFIED": 0,
	"TAST":        1,
	"AUTOTEST":    2,
}

func (x VmTestRequest_TestHarness) String() string {
	return proto.EnumName(VmTestRequest_TestHarness_name, int32(x))
}

func (VmTestRequest_TestHarness) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{6, 0}
}

type BuildTargetUnitTestRequest struct {
	// The build target being tested.
	BuildTarget *chromiumos.BuildTarget `protobuf:"bytes,1,opt,name=build_target,json=buildTarget,proto3" json:"build_target,omitempty"`
	// The path where the result tarball should be saved.
	ResultPath string `protobuf:"bytes,2,opt,name=result_path,json=resultPath,proto3" json:"result_path,omitempty"`
	// The chroot containing the sysroot.
	Chroot               *chromiumos.Chroot `protobuf:"bytes,3,opt,name=chroot,proto3" json:"chroot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BuildTargetUnitTestRequest) Reset()         { *m = BuildTargetUnitTestRequest{} }
func (m *BuildTargetUnitTestRequest) String() string { return proto.CompactTextString(m) }
func (*BuildTargetUnitTestRequest) ProtoMessage()    {}
func (*BuildTargetUnitTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{0}
}

func (m *BuildTargetUnitTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildTargetUnitTestRequest.Unmarshal(m, b)
}
func (m *BuildTargetUnitTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildTargetUnitTestRequest.Marshal(b, m, deterministic)
}
func (m *BuildTargetUnitTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildTargetUnitTestRequest.Merge(m, src)
}
func (m *BuildTargetUnitTestRequest) XXX_Size() int {
	return xxx_messageInfo_BuildTargetUnitTestRequest.Size(m)
}
func (m *BuildTargetUnitTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildTargetUnitTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildTargetUnitTestRequest proto.InternalMessageInfo

func (m *BuildTargetUnitTestRequest) GetBuildTarget() *chromiumos.BuildTarget {
	if m != nil {
		return m.BuildTarget
	}
	return nil
}

func (m *BuildTargetUnitTestRequest) GetResultPath() string {
	if m != nil {
		return m.ResultPath
	}
	return ""
}

func (m *BuildTargetUnitTestRequest) GetChroot() *chromiumos.Chroot {
	if m != nil {
		return m.Chroot
	}
	return nil
}

type BuildTargetUnitTestResponse struct {
	// The unittest tarball that was created.
	TarballPath string `protobuf:"bytes,1,opt,name=tarball_path,json=tarballPath,proto3" json:"tarball_path,omitempty"`
	// The list of packages that failed.
	FailedPackages       []*chromiumos.PackageInfo `protobuf:"bytes,2,rep,name=failed_packages,json=failedPackages,proto3" json:"failed_packages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BuildTargetUnitTestResponse) Reset()         { *m = BuildTargetUnitTestResponse{} }
func (m *BuildTargetUnitTestResponse) String() string { return proto.CompactTextString(m) }
func (*BuildTargetUnitTestResponse) ProtoMessage()    {}
func (*BuildTargetUnitTestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{1}
}

func (m *BuildTargetUnitTestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildTargetUnitTestResponse.Unmarshal(m, b)
}
func (m *BuildTargetUnitTestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildTargetUnitTestResponse.Marshal(b, m, deterministic)
}
func (m *BuildTargetUnitTestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildTargetUnitTestResponse.Merge(m, src)
}
func (m *BuildTargetUnitTestResponse) XXX_Size() int {
	return xxx_messageInfo_BuildTargetUnitTestResponse.Size(m)
}
func (m *BuildTargetUnitTestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildTargetUnitTestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuildTargetUnitTestResponse proto.InternalMessageInfo

func (m *BuildTargetUnitTestResponse) GetTarballPath() string {
	if m != nil {
		return m.TarballPath
	}
	return ""
}

func (m *BuildTargetUnitTestResponse) GetFailedPackages() []*chromiumos.PackageInfo {
	if m != nil {
		return m.FailedPackages
	}
	return nil
}

// Chromite UnitTest request.
type ChromiteUnitTestRequest struct {
	// The chroot to use to execute the endpoint.
	Chroot               *chromiumos.Chroot `protobuf:"bytes,1,opt,name=chroot,proto3" json:"chroot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ChromiteUnitTestRequest) Reset()         { *m = ChromiteUnitTestRequest{} }
func (m *ChromiteUnitTestRequest) String() string { return proto.CompactTextString(m) }
func (*ChromiteUnitTestRequest) ProtoMessage()    {}
func (*ChromiteUnitTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{2}
}

func (m *ChromiteUnitTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChromiteUnitTestRequest.Unmarshal(m, b)
}
func (m *ChromiteUnitTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChromiteUnitTestRequest.Marshal(b, m, deterministic)
}
func (m *ChromiteUnitTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChromiteUnitTestRequest.Merge(m, src)
}
func (m *ChromiteUnitTestRequest) XXX_Size() int {
	return xxx_messageInfo_ChromiteUnitTestRequest.Size(m)
}
func (m *ChromiteUnitTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChromiteUnitTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChromiteUnitTestRequest proto.InternalMessageInfo

func (m *ChromiteUnitTestRequest) GetChroot() *chromiumos.Chroot {
	if m != nil {
		return m.Chroot
	}
	return nil
}

type ChromiteUnitTestResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChromiteUnitTestResponse) Reset()         { *m = ChromiteUnitTestResponse{} }
func (m *ChromiteUnitTestResponse) String() string { return proto.CompactTextString(m) }
func (*ChromiteUnitTestResponse) ProtoMessage()    {}
func (*ChromiteUnitTestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{3}
}

func (m *ChromiteUnitTestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChromiteUnitTestResponse.Unmarshal(m, b)
}
func (m *ChromiteUnitTestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChromiteUnitTestResponse.Marshal(b, m, deterministic)
}
func (m *ChromiteUnitTestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChromiteUnitTestResponse.Merge(m, src)
}
func (m *ChromiteUnitTestResponse) XXX_Size() int {
	return xxx_messageInfo_ChromiteUnitTestResponse.Size(m)
}
func (m *ChromiteUnitTestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChromiteUnitTestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChromiteUnitTestResponse proto.InternalMessageInfo

// Run the debug_info_test script.
type DebugInfoTestRequest struct {
	// The sysroot to tests.
	Sysroot *Sysroot `protobuf:"bytes,1,opt,name=sysroot,proto3" json:"sysroot,omitempty"`
	// The chroot to use to execute the endpoint.
	Chroot               *chromiumos.Chroot `protobuf:"bytes,2,opt,name=chroot,proto3" json:"chroot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DebugInfoTestRequest) Reset()         { *m = DebugInfoTestRequest{} }
func (m *DebugInfoTestRequest) String() string { return proto.CompactTextString(m) }
func (*DebugInfoTestRequest) ProtoMessage()    {}
func (*DebugInfoTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{4}
}

func (m *DebugInfoTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugInfoTestRequest.Unmarshal(m, b)
}
func (m *DebugInfoTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugInfoTestRequest.Marshal(b, m, deterministic)
}
func (m *DebugInfoTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugInfoTestRequest.Merge(m, src)
}
func (m *DebugInfoTestRequest) XXX_Size() int {
	return xxx_messageInfo_DebugInfoTestRequest.Size(m)
}
func (m *DebugInfoTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugInfoTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DebugInfoTestRequest proto.InternalMessageInfo

func (m *DebugInfoTestRequest) GetSysroot() *Sysroot {
	if m != nil {
		return m.Sysroot
	}
	return nil
}

func (m *DebugInfoTestRequest) GetChroot() *chromiumos.Chroot {
	if m != nil {
		return m.Chroot
	}
	return nil
}

type DebugInfoTestResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugInfoTestResponse) Reset()         { *m = DebugInfoTestResponse{} }
func (m *DebugInfoTestResponse) String() string { return proto.CompactTextString(m) }
func (*DebugInfoTestResponse) ProtoMessage()    {}
func (*DebugInfoTestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{5}
}

func (m *DebugInfoTestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugInfoTestResponse.Unmarshal(m, b)
}
func (m *DebugInfoTestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugInfoTestResponse.Marshal(b, m, deterministic)
}
func (m *DebugInfoTestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugInfoTestResponse.Merge(m, src)
}
func (m *DebugInfoTestResponse) XXX_Size() int {
	return xxx_messageInfo_DebugInfoTestResponse.Size(m)
}
func (m *DebugInfoTestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugInfoTestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DebugInfoTestResponse proto.InternalMessageInfo

type VmTestRequest struct {
	// The build target under test. Required.
	BuildTarget *chromiumos.BuildTarget `protobuf:"bytes,1,opt,name=build_target,json=buildTarget,proto3" json:"build_target,omitempty"`
	// The chroot in which to run commands. Required.
	Chroot *chromiumos.Chroot `protobuf:"bytes,2,opt,name=chroot,proto3" json:"chroot,omitempty"`
	// The VM image to test. Required.
	VmImage *VmImage `protobuf:"bytes,3,opt,name=vm_image,json=vmImage,proto3" json:"vm_image,omitempty"`
	// Options for SSHing into the VM.
	SshOptions  *VmTestRequest_SshOptions `protobuf:"bytes,4,opt,name=ssh_options,json=sshOptions,proto3" json:"ssh_options,omitempty"`
	TestHarness VmTestRequest_TestHarness `protobuf:"varint,5,opt,name=test_harness,json=testHarness,proto3,enum=chromite.api.VmTestRequest_TestHarness" json:"test_harness,omitempty"`
	// All VM tests that should be run. At least one must be specified.
	VmTests              []*VmTestRequest_VmTest `protobuf:"bytes,6,rep,name=vm_tests,json=vmTests,proto3" json:"vm_tests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *VmTestRequest) Reset()         { *m = VmTestRequest{} }
func (m *VmTestRequest) String() string { return proto.CompactTextString(m) }
func (*VmTestRequest) ProtoMessage()    {}
func (*VmTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{6}
}

func (m *VmTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmTestRequest.Unmarshal(m, b)
}
func (m *VmTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmTestRequest.Marshal(b, m, deterministic)
}
func (m *VmTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmTestRequest.Merge(m, src)
}
func (m *VmTestRequest) XXX_Size() int {
	return xxx_messageInfo_VmTestRequest.Size(m)
}
func (m *VmTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VmTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VmTestRequest proto.InternalMessageInfo

func (m *VmTestRequest) GetBuildTarget() *chromiumos.BuildTarget {
	if m != nil {
		return m.BuildTarget
	}
	return nil
}

func (m *VmTestRequest) GetChroot() *chromiumos.Chroot {
	if m != nil {
		return m.Chroot
	}
	return nil
}

func (m *VmTestRequest) GetVmImage() *VmImage {
	if m != nil {
		return m.VmImage
	}
	return nil
}

func (m *VmTestRequest) GetSshOptions() *VmTestRequest_SshOptions {
	if m != nil {
		return m.SshOptions
	}
	return nil
}

func (m *VmTestRequest) GetTestHarness() VmTestRequest_TestHarness {
	if m != nil {
		return m.TestHarness
	}
	return VmTestRequest_UNSPECIFIED
}

func (m *VmTestRequest) GetVmTests() []*VmTestRequest_VmTest {
	if m != nil {
		return m.VmTests
	}
	return nil
}

// Options for SSHing into the VM.
type VmTestRequest_SshOptions struct {
	// Path to the private key to the VM.
	PrivateKeyPath string `protobuf:"bytes,1,opt,name=private_key_path,json=privateKeyPath,proto3" json:"private_key_path,omitempty"`
	// SSH port to communicate with VM.
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmTestRequest_SshOptions) Reset()         { *m = VmTestRequest_SshOptions{} }
func (m *VmTestRequest_SshOptions) String() string { return proto.CompactTextString(m) }
func (*VmTestRequest_SshOptions) ProtoMessage()    {}
func (*VmTestRequest_SshOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{6, 0}
}

func (m *VmTestRequest_SshOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmTestRequest_SshOptions.Unmarshal(m, b)
}
func (m *VmTestRequest_SshOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmTestRequest_SshOptions.Marshal(b, m, deterministic)
}
func (m *VmTestRequest_SshOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmTestRequest_SshOptions.Merge(m, src)
}
func (m *VmTestRequest_SshOptions) XXX_Size() int {
	return xxx_messageInfo_VmTestRequest_SshOptions.Size(m)
}
func (m *VmTestRequest_SshOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_VmTestRequest_SshOptions.DiscardUnknown(m)
}

var xxx_messageInfo_VmTestRequest_SshOptions proto.InternalMessageInfo

func (m *VmTestRequest_SshOptions) GetPrivateKeyPath() string {
	if m != nil {
		return m.PrivateKeyPath
	}
	return ""
}

func (m *VmTestRequest_SshOptions) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// A specific VM test to run.
type VmTestRequest_VmTest struct {
	// The name pattern for the VM test.
	// For Autotest, this pattern matches against the test suite name.
	// For Tast, this pattern matches against the Tast test name.
	Pattern              string   `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmTestRequest_VmTest) Reset()         { *m = VmTestRequest_VmTest{} }
func (m *VmTestRequest_VmTest) String() string { return proto.CompactTextString(m) }
func (*VmTestRequest_VmTest) ProtoMessage()    {}
func (*VmTestRequest_VmTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{6, 1}
}

func (m *VmTestRequest_VmTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmTestRequest_VmTest.Unmarshal(m, b)
}
func (m *VmTestRequest_VmTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmTestRequest_VmTest.Marshal(b, m, deterministic)
}
func (m *VmTestRequest_VmTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmTestRequest_VmTest.Merge(m, src)
}
func (m *VmTestRequest_VmTest) XXX_Size() int {
	return xxx_messageInfo_VmTestRequest_VmTest.Size(m)
}
func (m *VmTestRequest_VmTest) XXX_DiscardUnknown() {
	xxx_messageInfo_VmTestRequest_VmTest.DiscardUnknown(m)
}

var xxx_messageInfo_VmTestRequest_VmTest proto.InternalMessageInfo

func (m *VmTestRequest_VmTest) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

type VmTestResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmTestResponse) Reset()         { *m = VmTestResponse{} }
func (m *VmTestResponse) String() string { return proto.CompactTextString(m) }
func (*VmTestResponse) ProtoMessage()    {}
func (*VmTestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5c91b4ac6233d9e, []int{7}
}

func (m *VmTestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmTestResponse.Unmarshal(m, b)
}
func (m *VmTestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmTestResponse.Marshal(b, m, deterministic)
}
func (m *VmTestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmTestResponse.Merge(m, src)
}
func (m *VmTestResponse) XXX_Size() int {
	return xxx_messageInfo_VmTestResponse.Size(m)
}
func (m *VmTestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VmTestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VmTestResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("chromite.api.VmTestRequest_TestHarness", VmTestRequest_TestHarness_name, VmTestRequest_TestHarness_value)
	proto.RegisterType((*BuildTargetUnitTestRequest)(nil), "chromite.api.BuildTargetUnitTestRequest")
	proto.RegisterType((*BuildTargetUnitTestResponse)(nil), "chromite.api.BuildTargetUnitTestResponse")
	proto.RegisterType((*ChromiteUnitTestRequest)(nil), "chromite.api.ChromiteUnitTestRequest")
	proto.RegisterType((*ChromiteUnitTestResponse)(nil), "chromite.api.ChromiteUnitTestResponse")
	proto.RegisterType((*DebugInfoTestRequest)(nil), "chromite.api.DebugInfoTestRequest")
	proto.RegisterType((*DebugInfoTestResponse)(nil), "chromite.api.DebugInfoTestResponse")
	proto.RegisterType((*VmTestRequest)(nil), "chromite.api.VmTestRequest")
	proto.RegisterType((*VmTestRequest_SshOptions)(nil), "chromite.api.VmTestRequest.SshOptions")
	proto.RegisterType((*VmTestRequest_VmTest)(nil), "chromite.api.VmTestRequest.VmTest")
	proto.RegisterType((*VmTestResponse)(nil), "chromite.api.VmTestResponse")
}

func init() { proto.RegisterFile("chromite/api/test.proto", fileDescriptor_f5c91b4ac6233d9e) }

var fileDescriptor_f5c91b4ac6233d9e = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x25, 0x5d, 0xd7, 0x96, 0x9b, 0xae, 0x8b, 0x0c, 0x53, 0x23, 0x6f, 0x12, 0x25, 0x88, 0x51,
	0x78, 0x48, 0xd0, 0x90, 0x26, 0x84, 0x84, 0xc4, 0xd6, 0x15, 0xe8, 0x90, 0xd8, 0x94, 0x76, 0x13,
	0xe2, 0xa5, 0x72, 0x3b, 0xaf, 0x8d, 0xd6, 0xc4, 0x21, 0x76, 0x2b, 0xed, 0x95, 0x3f, 0xe0, 0x23,
	0xf6, 0x23, 0xfb, 0x9f, 0xfd, 0x03, 0x4a, 0xec, 0xae, 0x69, 0xe8, 0x3a, 0x90, 0x78, 0xaa, 0x7d,
	0xcf, 0xf1, 0xb9, 0xf7, 0x3a, 0xe7, 0xd6, 0x50, 0xed, 0x0f, 0x23, 0xe6, 0x7b, 0x82, 0x3a, 0x24,
	0xf4, 0x1c, 0x41, 0xb9, 0xb0, 0xc3, 0x88, 0x09, 0x86, 0xca, 0x53, 0xc0, 0x26, 0xa1, 0x87, 0xb7,
	0xe6, 0x68, 0xbd, 0xb1, 0x37, 0x3a, 0xeb, 0x92, 0xd0, 0x93, 0x5c, 0x6c, 0xce, 0xa1, 0x9e, 0x4f,
	0x06, 0x54, 0x21, 0x78, 0x0e, 0xe1, 0x97, 0x3c, 0x62, 0x4c, 0x65, 0xc0, 0x2a, 0xf5, 0xd8, 0x67,
	0xdc, 0xe9, 0x33, 0xdf, 0x67, 0x81, 0x04, 0xac, 0x2b, 0x0d, 0xf0, 0x7e, 0x9c, 0xa2, 0x43, 0xa2,
	0x01, 0x15, 0x27, 0x81, 0x27, 0x3a, 0x94, 0x0b, 0x97, 0xfe, 0x18, 0x53, 0x2e, 0xd0, 0x3b, 0x28,
	0xcb, 0x02, 0x44, 0x02, 0x9b, 0x5a, 0x4d, 0xab, 0xeb, 0x3b, 0x55, 0x7b, 0x26, 0x67, 0xa7, 0x4e,
	0xbb, 0x7a, 0x6f, 0xb6, 0x41, 0x4f, 0x40, 0x8f, 0x28, 0x1f, 0x8f, 0x44, 0x37, 0x24, 0x62, 0x68,
	0xe6, 0x6a, 0x5a, 0xfd, 0xa1, 0x0b, 0x32, 0x74, 0x4c, 0xc4, 0x10, 0xbd, 0x82, 0x42, 0xac, 0xc3,
	0x84, 0xb9, 0x92, 0xc8, 0xa2, 0xb4, 0x6c, 0x23, 0x41, 0x5c, 0xc5, 0xb0, 0x7e, 0x6a, 0xb0, 0xb9,
	0xb0, 0x4e, 0x1e, 0xb2, 0x80, 0x53, 0xf4, 0x14, 0xca, 0x82, 0x44, 0x3d, 0x32, 0x1a, 0xc9, 0x6c,
	0x5a, 0x92, 0x4d, 0x57, 0xb1, 0x24, 0xdd, 0x07, 0x58, 0x3f, 0x27, 0xde, 0x88, 0x9e, 0x75, 0x43,
	0xd2, 0xbf, 0x20, 0x03, 0xca, 0xcd, 0x5c, 0x6d, 0x25, 0xdb, 0xce, 0xb1, 0xc4, 0x5a, 0xc1, 0x39,
	0x73, 0x2b, 0x92, 0xaf, 0x42, 0xdc, 0x6a, 0x42, 0xb5, 0xa1, 0xee, 0x38, 0x7b, 0x51, 0xb3, 0x5e,
	0xb4, 0x7b, 0x7b, 0xc1, 0x60, 0xfe, 0x29, 0x23, 0xfb, 0xb0, 0x38, 0x3c, 0x3e, 0xa0, 0xbd, 0xf1,
	0x20, 0xce, 0x9f, 0xd6, 0x77, 0xa0, 0xa8, 0xbe, 0xa8, 0x4a, 0xb0, 0x61, 0xa7, 0x4d, 0x63, 0xb7,
	0x25, 0xe8, 0x4e, 0x59, 0xa9, 0x82, 0x72, 0xf7, 0x16, 0x54, 0x85, 0x8d, 0x4c, 0x52, 0x55, 0xcd,
	0x55, 0x1e, 0xd6, 0x4e, 0xfd, 0xff, 0x65, 0x88, 0x7f, 0x28, 0x09, 0xbd, 0x86, 0xd2, 0xc4, 0xef,
	0x26, 0xf6, 0x56, 0xee, 0xc8, 0x34, 0x7c, 0xea, 0xb7, 0x62, 0xd0, 0x2d, 0x4e, 0xe4, 0x02, 0x7d,
	0x02, 0x9d, 0xf3, 0x61, 0x97, 0x85, 0xc2, 0x63, 0x01, 0x37, 0xf3, 0xc9, 0xa1, 0xed, 0xec, 0xa1,
	0x54, 0x2f, 0x76, 0x9b, 0x0f, 0x8f, 0x24, 0xdb, 0x05, 0x7e, 0xbb, 0x46, 0x87, 0x50, 0x8e, 0x67,
	0xb3, 0x3b, 0x24, 0x51, 0x40, 0x39, 0x37, 0x57, 0x6b, 0x5a, 0xbd, 0xb2, 0xf3, 0x62, 0x99, 0x52,
	0xbc, 0xfe, 0x2c, 0xe9, 0xae, 0x2e, 0x66, 0x1b, 0xf4, 0x3e, 0x69, 0x23, 0x8e, 0x70, 0xb3, 0x90,
	0x98, 0xcd, 0x5a, 0xa6, 0xa3, 0x76, 0xc5, 0x49, 0xf2, 0xcb, 0xf1, 0x21, 0xc0, 0xac, 0x48, 0x54,
	0x07, 0x23, 0x8c, 0xbc, 0x09, 0x11, 0xb4, 0x7b, 0x41, 0x2f, 0xd3, 0x3e, 0xaf, 0xa8, 0xf8, 0x17,
	0x7a, 0x99, 0x58, 0x1d, 0x41, 0x3e, 0x64, 0x91, 0xbc, 0xe7, 0x55, 0x37, 0x59, 0x63, 0x0b, 0x0a,
	0x52, 0x1e, 0x99, 0x50, 0x0c, 0x89, 0x10, 0x34, 0x0a, 0xd4, 0xf1, 0xe9, 0xd6, 0xda, 0x05, 0x3d,
	0xd5, 0x0a, 0x5a, 0x07, 0xfd, 0xe4, 0x6b, 0xfb, 0xb8, 0xd9, 0x68, 0x7d, 0x6c, 0x35, 0x0f, 0x8c,
	0x07, 0xa8, 0x04, 0xf9, 0xce, 0x5e, 0xbb, 0x63, 0x68, 0xa8, 0x0c, 0xa5, 0xbd, 0x93, 0xce, 0x51,
	0xa7, 0xd9, 0xee, 0x18, 0x39, 0xcb, 0x80, 0xca, 0xb4, 0x11, 0xe9, 0x9c, 0x9d, 0x5f, 0x2b, 0x52,
	0xaa, 0x4d, 0xa3, 0x89, 0xd7, 0xa7, 0x28, 0x82, 0x47, 0x0b, 0xc6, 0x17, 0xd5, 0xe7, 0x6f, 0xe3,
	0xee, 0x7f, 0x22, 0xfc, 0xf2, 0x2f, 0x98, 0xca, 0xb5, 0x85, 0xeb, 0x1b, 0x9c, 0x33, 0x72, 0x88,
	0x80, 0x91, 0x9d, 0x33, 0xf4, 0x7c, 0x5e, 0xe6, 0x8e, 0x71, 0xc6, 0xdb, 0xf7, 0xd1, 0xd4, 0xdf,
	0xce, 0x37, 0x58, 0x9b, 0x9b, 0x1c, 0x94, 0xf9, 0xbc, 0x8b, 0x66, 0x19, 0x3f, 0x5b, 0xca, 0x51,
	0xca, 0x8d, 0xdb, 0xcf, 0xb5, 0xb9, 0xc4, 0x31, 0x78, 0x6b, 0x31, 0x28, 0x45, 0x70, 0xf9, 0xfa,
	0x06, 0x97, 0x20, 0x1f, 0xfb, 0xcf, 0xd0, 0xf6, 0xdf, 0x7e, 0xdf, 0x1d, 0xb0, 0xdb, 0x99, 0xb3,
	0x59, 0x34, 0x70, 0x52, 0xcf, 0x82, 0x17, 0x9c, 0x47, 0xc4, 0x49, 0x5e, 0x05, 0x67, 0xc0, 0x9c,
	0xf4, 0x63, 0xd2, 0x2b, 0x24, 0xe1, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x63, 0x7d, 0xe0,
	0x53, 0xc2, 0x06, 0x00, 0x00,
}
