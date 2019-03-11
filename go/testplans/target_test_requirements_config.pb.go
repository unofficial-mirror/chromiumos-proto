// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testplans/target_test_requirements_config.proto

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

type GceTestCfg struct {
	GceTest              []*GceTestCfg_GceTest `protobuf:"bytes,1,rep,name=gce_test,json=gceTest,proto3" json:"gce_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GceTestCfg) Reset()         { *m = GceTestCfg{} }
func (m *GceTestCfg) String() string { return proto.CompactTextString(m) }
func (*GceTestCfg) ProtoMessage()    {}
func (*GceTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{0}
}

func (m *GceTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GceTestCfg.Unmarshal(m, b)
}
func (m *GceTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GceTestCfg.Marshal(b, m, deterministic)
}
func (m *GceTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GceTestCfg.Merge(m, src)
}
func (m *GceTestCfg) XXX_Size() int {
	return xxx_messageInfo_GceTestCfg.Size(m)
}
func (m *GceTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_GceTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_GceTestCfg proto.InternalMessageInfo

func (m *GceTestCfg) GetGceTest() []*GceTestCfg_GceTest {
	if m != nil {
		return m.GceTest
	}
	return nil
}

type GceTestCfg_GceTest struct {
	// Test type to be run.
	TestType string `protobuf:"bytes,1,opt,name=test_type,json=testType,proto3" json:"test_type,omitempty"`
	// Test suite to be run in GCETest.
	TestSuite string `protobuf:"bytes,2,opt,name=test_suite,json=testSuite,proto3" json:"test_suite,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec int32 `protobuf:"varint,3,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	// Use the old ctest code path rather than the new chromite one.
	UseCtest             bool     `protobuf:"varint,4,opt,name=use_ctest,json=useCtest,proto3" json:"use_ctest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GceTestCfg_GceTest) Reset()         { *m = GceTestCfg_GceTest{} }
func (m *GceTestCfg_GceTest) String() string { return proto.CompactTextString(m) }
func (*GceTestCfg_GceTest) ProtoMessage()    {}
func (*GceTestCfg_GceTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{0, 0}
}

func (m *GceTestCfg_GceTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GceTestCfg_GceTest.Unmarshal(m, b)
}
func (m *GceTestCfg_GceTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GceTestCfg_GceTest.Marshal(b, m, deterministic)
}
func (m *GceTestCfg_GceTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GceTestCfg_GceTest.Merge(m, src)
}
func (m *GceTestCfg_GceTest) XXX_Size() int {
	return xxx_messageInfo_GceTestCfg_GceTest.Size(m)
}
func (m *GceTestCfg_GceTest) XXX_DiscardUnknown() {
	xxx_messageInfo_GceTestCfg_GceTest.DiscardUnknown(m)
}

var xxx_messageInfo_GceTestCfg_GceTest proto.InternalMessageInfo

func (m *GceTestCfg_GceTest) GetTestType() string {
	if m != nil {
		return m.TestType
	}
	return ""
}

func (m *GceTestCfg_GceTest) GetTestSuite() string {
	if m != nil {
		return m.TestSuite
	}
	return ""
}

func (m *GceTestCfg_GceTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

func (m *GceTestCfg_GceTest) GetUseCtest() bool {
	if m != nil {
		return m.UseCtest
	}
	return false
}

type HwTestCfg struct {
	HwTest               []*HwTestCfg_HwTest `protobuf:"bytes,1,rep,name=hw_test,json=hwTest,proto3" json:"hw_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HwTestCfg) Reset()         { *m = HwTestCfg{} }
func (m *HwTestCfg) String() string { return proto.CompactTextString(m) }
func (*HwTestCfg) ProtoMessage()    {}
func (*HwTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{1}
}

func (m *HwTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HwTestCfg.Unmarshal(m, b)
}
func (m *HwTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HwTestCfg.Marshal(b, m, deterministic)
}
func (m *HwTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HwTestCfg.Merge(m, src)
}
func (m *HwTestCfg) XXX_Size() int {
	return xxx_messageInfo_HwTestCfg.Size(m)
}
func (m *HwTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_HwTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_HwTestCfg proto.InternalMessageInfo

func (m *HwTestCfg) GetHwTest() []*HwTestCfg_HwTest {
	if m != nil {
		return m.HwTest
	}
	return nil
}

// Configuration for a hardware test suite.
type HwTestCfg_HwTest struct {
	// Name of the test suite to run.
	Suite string `protobuf:"bytes,1,opt,name=suite,proto3" json:"suite,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec int32 `protobuf:"varint,2,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	// Failure on HW tests warns only (does not generate error).
	WarnOnly bool `protobuf:"varint,3,opt,name=warn_only,json=warnOnly,proto3" json:"warn_only,omitempty"`
	// Usually we consider structural failures here as OK.
	Critical bool `protobuf:"varint,4,opt,name=critical,proto3" json:"critical,omitempty"`
	// Should we file bugs if a test fails in a suite run.
	FileBugs bool `protobuf:"varint,5,opt,name=file_bugs,json=fileBugs,proto3" json:"file_bugs,omitempty"`
	// Minimum number of DUTs required for testing in the hw lab.
	MinimumDuts int32 `protobuf:"varint,6,opt,name=minimum_duts,json=minimumDuts,proto3" json:"minimum_duts,omitempty"`
	// Whether we should retry tests that fail in a suite run.
	Retry bool `protobuf:"varint,7,opt,name=retry,proto3" json:"retry,omitempty"`
	// Maximum job retries allowed at suite level. 0 for no max.
	MaxRetries int32 `protobuf:"varint,8,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Preferred minimum duts. Lab will prioritize on getting such
	// number of duts even if the suite is competing with
	// other suites that have higher priority.
	SuiteMinDuts int32 `protobuf:"varint,9,opt,name=suite_min_duts,json=suiteMinDuts,proto3" json:"suite_min_duts,omitempty"`
	// Only offload failed tests to Google Storage.
	OffloadFailuresOnly  bool     `protobuf:"varint,10,opt,name=offload_failures_only,json=offloadFailuresOnly,proto3" json:"offload_failures_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HwTestCfg_HwTest) Reset()         { *m = HwTestCfg_HwTest{} }
func (m *HwTestCfg_HwTest) String() string { return proto.CompactTextString(m) }
func (*HwTestCfg_HwTest) ProtoMessage()    {}
func (*HwTestCfg_HwTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{1, 0}
}

func (m *HwTestCfg_HwTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HwTestCfg_HwTest.Unmarshal(m, b)
}
func (m *HwTestCfg_HwTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HwTestCfg_HwTest.Marshal(b, m, deterministic)
}
func (m *HwTestCfg_HwTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HwTestCfg_HwTest.Merge(m, src)
}
func (m *HwTestCfg_HwTest) XXX_Size() int {
	return xxx_messageInfo_HwTestCfg_HwTest.Size(m)
}
func (m *HwTestCfg_HwTest) XXX_DiscardUnknown() {
	xxx_messageInfo_HwTestCfg_HwTest.DiscardUnknown(m)
}

var xxx_messageInfo_HwTestCfg_HwTest proto.InternalMessageInfo

func (m *HwTestCfg_HwTest) GetSuite() string {
	if m != nil {
		return m.Suite
	}
	return ""
}

func (m *HwTestCfg_HwTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

func (m *HwTestCfg_HwTest) GetWarnOnly() bool {
	if m != nil {
		return m.WarnOnly
	}
	return false
}

func (m *HwTestCfg_HwTest) GetCritical() bool {
	if m != nil {
		return m.Critical
	}
	return false
}

func (m *HwTestCfg_HwTest) GetFileBugs() bool {
	if m != nil {
		return m.FileBugs
	}
	return false
}

func (m *HwTestCfg_HwTest) GetMinimumDuts() int32 {
	if m != nil {
		return m.MinimumDuts
	}
	return 0
}

func (m *HwTestCfg_HwTest) GetRetry() bool {
	if m != nil {
		return m.Retry
	}
	return false
}

func (m *HwTestCfg_HwTest) GetMaxRetries() int32 {
	if m != nil {
		return m.MaxRetries
	}
	return 0
}

func (m *HwTestCfg_HwTest) GetSuiteMinDuts() int32 {
	if m != nil {
		return m.SuiteMinDuts
	}
	return 0
}

func (m *HwTestCfg_HwTest) GetOffloadFailuresOnly() bool {
	if m != nil {
		return m.OffloadFailuresOnly
	}
	return false
}

type MoblabVmTestCfg struct {
	MoblabTest           []*MoblabVmTestCfg_MoblabTest `protobuf:"bytes,1,rep,name=moblab_test,json=moblabTest,proto3" json:"moblab_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MoblabVmTestCfg) Reset()         { *m = MoblabVmTestCfg{} }
func (m *MoblabVmTestCfg) String() string { return proto.CompactTextString(m) }
func (*MoblabVmTestCfg) ProtoMessage()    {}
func (*MoblabVmTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{2}
}

func (m *MoblabVmTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoblabVmTestCfg.Unmarshal(m, b)
}
func (m *MoblabVmTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoblabVmTestCfg.Marshal(b, m, deterministic)
}
func (m *MoblabVmTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoblabVmTestCfg.Merge(m, src)
}
func (m *MoblabVmTestCfg) XXX_Size() int {
	return xxx_messageInfo_MoblabVmTestCfg.Size(m)
}
func (m *MoblabVmTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_MoblabVmTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_MoblabVmTestCfg proto.InternalMessageInfo

func (m *MoblabVmTestCfg) GetMoblabTest() []*MoblabVmTestCfg_MoblabTest {
	if m != nil {
		return m.MoblabTest
	}
	return nil
}

type MoblabVmTestCfg_MoblabTest struct {
	// Test type to be run.
	TestType string `protobuf:"bytes,1,opt,name=test_type,json=testType,proto3" json:"test_type,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec           int32    `protobuf:"varint,2,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoblabVmTestCfg_MoblabTest) Reset()         { *m = MoblabVmTestCfg_MoblabTest{} }
func (m *MoblabVmTestCfg_MoblabTest) String() string { return proto.CompactTextString(m) }
func (*MoblabVmTestCfg_MoblabTest) ProtoMessage()    {}
func (*MoblabVmTestCfg_MoblabTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{2, 0}
}

func (m *MoblabVmTestCfg_MoblabTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoblabVmTestCfg_MoblabTest.Unmarshal(m, b)
}
func (m *MoblabVmTestCfg_MoblabTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoblabVmTestCfg_MoblabTest.Marshal(b, m, deterministic)
}
func (m *MoblabVmTestCfg_MoblabTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoblabVmTestCfg_MoblabTest.Merge(m, src)
}
func (m *MoblabVmTestCfg_MoblabTest) XXX_Size() int {
	return xxx_messageInfo_MoblabVmTestCfg_MoblabTest.Size(m)
}
func (m *MoblabVmTestCfg_MoblabTest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoblabVmTestCfg_MoblabTest.DiscardUnknown(m)
}

var xxx_messageInfo_MoblabVmTestCfg_MoblabTest proto.InternalMessageInfo

func (m *MoblabVmTestCfg_MoblabTest) GetTestType() string {
	if m != nil {
		return m.TestType
	}
	return ""
}

func (m *MoblabVmTestCfg_MoblabTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

type TastVmTestCfg struct {
	TastVmTest           []*TastVmTestCfg_TastVmTest `protobuf:"bytes,1,rep,name=tast_vm_test,json=tastVmTest,proto3" json:"tast_vm_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TastVmTestCfg) Reset()         { *m = TastVmTestCfg{} }
func (m *TastVmTestCfg) String() string { return proto.CompactTextString(m) }
func (*TastVmTestCfg) ProtoMessage()    {}
func (*TastVmTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{3}
}

func (m *TastVmTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TastVmTestCfg.Unmarshal(m, b)
}
func (m *TastVmTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TastVmTestCfg.Marshal(b, m, deterministic)
}
func (m *TastVmTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TastVmTestCfg.Merge(m, src)
}
func (m *TastVmTestCfg) XXX_Size() int {
	return xxx_messageInfo_TastVmTestCfg.Size(m)
}
func (m *TastVmTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_TastVmTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_TastVmTestCfg proto.InternalMessageInfo

func (m *TastVmTestCfg) GetTastVmTest() []*TastVmTestCfg_TastVmTest {
	if m != nil {
		return m.TastVmTest
	}
	return nil
}

type TastVmTestCfg_TastTestExpr struct {
	// A single tast test expression. See https://goo.gl/UPNEgT
	TestExpr             string   `protobuf:"bytes,1,opt,name=test_expr,json=testExpr,proto3" json:"test_expr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TastVmTestCfg_TastTestExpr) Reset()         { *m = TastVmTestCfg_TastTestExpr{} }
func (m *TastVmTestCfg_TastTestExpr) String() string { return proto.CompactTextString(m) }
func (*TastVmTestCfg_TastTestExpr) ProtoMessage()    {}
func (*TastVmTestCfg_TastTestExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{3, 0}
}

func (m *TastVmTestCfg_TastTestExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TastVmTestCfg_TastTestExpr.Unmarshal(m, b)
}
func (m *TastVmTestCfg_TastTestExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TastVmTestCfg_TastTestExpr.Marshal(b, m, deterministic)
}
func (m *TastVmTestCfg_TastTestExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TastVmTestCfg_TastTestExpr.Merge(m, src)
}
func (m *TastVmTestCfg_TastTestExpr) XXX_Size() int {
	return xxx_messageInfo_TastVmTestCfg_TastTestExpr.Size(m)
}
func (m *TastVmTestCfg_TastTestExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_TastVmTestCfg_TastTestExpr.DiscardUnknown(m)
}

var xxx_messageInfo_TastVmTestCfg_TastTestExpr proto.InternalMessageInfo

func (m *TastVmTestCfg_TastTestExpr) GetTestExpr() string {
	if m != nil {
		return m.TestExpr
	}
	return ""
}

type TastVmTestCfg_TastVmTest struct {
	// String containing short human-readable name describing test suite.
	SuiteName string `protobuf:"bytes,1,opt,name=suite_name,json=suiteName,proto3" json:"suite_name,omitempty"`
	// List of string expressions describing which tests to run; this
	// is passed directly to the 'tast run' command. See
	// https://goo.gl/UPNEgT for info about test expressions.
	TastTestExpr []*TastVmTestCfg_TastTestExpr `protobuf:"bytes,2,rep,name=tast_test_expr,json=tastTestExpr,proto3" json:"tast_test_expr,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec           int32    `protobuf:"varint,3,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TastVmTestCfg_TastVmTest) Reset()         { *m = TastVmTestCfg_TastVmTest{} }
func (m *TastVmTestCfg_TastVmTest) String() string { return proto.CompactTextString(m) }
func (*TastVmTestCfg_TastVmTest) ProtoMessage()    {}
func (*TastVmTestCfg_TastVmTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{3, 1}
}

func (m *TastVmTestCfg_TastVmTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TastVmTestCfg_TastVmTest.Unmarshal(m, b)
}
func (m *TastVmTestCfg_TastVmTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TastVmTestCfg_TastVmTest.Marshal(b, m, deterministic)
}
func (m *TastVmTestCfg_TastVmTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TastVmTestCfg_TastVmTest.Merge(m, src)
}
func (m *TastVmTestCfg_TastVmTest) XXX_Size() int {
	return xxx_messageInfo_TastVmTestCfg_TastVmTest.Size(m)
}
func (m *TastVmTestCfg_TastVmTest) XXX_DiscardUnknown() {
	xxx_messageInfo_TastVmTestCfg_TastVmTest.DiscardUnknown(m)
}

var xxx_messageInfo_TastVmTestCfg_TastVmTest proto.InternalMessageInfo

func (m *TastVmTestCfg_TastVmTest) GetSuiteName() string {
	if m != nil {
		return m.SuiteName
	}
	return ""
}

func (m *TastVmTestCfg_TastVmTest) GetTastTestExpr() []*TastVmTestCfg_TastTestExpr {
	if m != nil {
		return m.TastTestExpr
	}
	return nil
}

func (m *TastVmTestCfg_TastVmTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

type VmTestCfg struct {
	VmTest               []*VmTestCfg_VmTest `protobuf:"bytes,1,rep,name=vm_test,json=vmTest,proto3" json:"vm_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VmTestCfg) Reset()         { *m = VmTestCfg{} }
func (m *VmTestCfg) String() string { return proto.CompactTextString(m) }
func (*VmTestCfg) ProtoMessage()    {}
func (*VmTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{4}
}

func (m *VmTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmTestCfg.Unmarshal(m, b)
}
func (m *VmTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmTestCfg.Marshal(b, m, deterministic)
}
func (m *VmTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmTestCfg.Merge(m, src)
}
func (m *VmTestCfg) XXX_Size() int {
	return xxx_messageInfo_VmTestCfg.Size(m)
}
func (m *VmTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_VmTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_VmTestCfg proto.InternalMessageInfo

func (m *VmTestCfg) GetVmTest() []*VmTestCfg_VmTest {
	if m != nil {
		return m.VmTest
	}
	return nil
}

type VmTestCfg_VmTest struct {
	// Test type to be run.
	TestType string `protobuf:"bytes,1,opt,name=test_type,json=testType,proto3" json:"test_type,omitempty"`
	// Test suite to be run in VMTest.
	TestSuite string `protobuf:"bytes,2,opt,name=test_suite,json=testSuite,proto3" json:"test_suite,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec int32 `protobuf:"varint,3,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	// Whether we should retry tests that fail in a suite run.
	Retry bool `protobuf:"varint,4,opt,name=retry,proto3" json:"retry,omitempty"`
	// Maximum job retries allowed at suite level. 0 for no max.
	MaxRetries int32 `protobuf:"varint,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Failure on VM tests warns only.
	WarnOnly bool `protobuf:"varint,6,opt,name=warn_only,json=warnOnly,proto3" json:"warn_only,omitempty"`
	// Use the old ctest code path rather than the new chromite one.
	UseCtest             bool     `protobuf:"varint,7,opt,name=use_ctest,json=useCtest,proto3" json:"use_ctest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmTestCfg_VmTest) Reset()         { *m = VmTestCfg_VmTest{} }
func (m *VmTestCfg_VmTest) String() string { return proto.CompactTextString(m) }
func (*VmTestCfg_VmTest) ProtoMessage()    {}
func (*VmTestCfg_VmTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{4, 0}
}

func (m *VmTestCfg_VmTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmTestCfg_VmTest.Unmarshal(m, b)
}
func (m *VmTestCfg_VmTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmTestCfg_VmTest.Marshal(b, m, deterministic)
}
func (m *VmTestCfg_VmTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmTestCfg_VmTest.Merge(m, src)
}
func (m *VmTestCfg_VmTest) XXX_Size() int {
	return xxx_messageInfo_VmTestCfg_VmTest.Size(m)
}
func (m *VmTestCfg_VmTest) XXX_DiscardUnknown() {
	xxx_messageInfo_VmTestCfg_VmTest.DiscardUnknown(m)
}

var xxx_messageInfo_VmTestCfg_VmTest proto.InternalMessageInfo

func (m *VmTestCfg_VmTest) GetTestType() string {
	if m != nil {
		return m.TestType
	}
	return ""
}

func (m *VmTestCfg_VmTest) GetTestSuite() string {
	if m != nil {
		return m.TestSuite
	}
	return ""
}

func (m *VmTestCfg_VmTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

func (m *VmTestCfg_VmTest) GetRetry() bool {
	if m != nil {
		return m.Retry
	}
	return false
}

func (m *VmTestCfg_VmTest) GetMaxRetries() int32 {
	if m != nil {
		return m.MaxRetries
	}
	return 0
}

func (m *VmTestCfg_VmTest) GetWarnOnly() bool {
	if m != nil {
		return m.WarnOnly
	}
	return false
}

func (m *VmTestCfg_VmTest) GetUseCtest() bool {
	if m != nil {
		return m.UseCtest
	}
	return false
}

// Specifies a CrOS target, either by reference design or by a specific build
// target.
type TargetCriteria struct {
	// Types that are valid to be assigned to TargetType:
	//	*TargetCriteria_ReferenceDesign
	//	*TargetCriteria_BuildTarget
	TargetType           isTargetCriteria_TargetType `protobuf_oneof:"target_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TargetCriteria) Reset()         { *m = TargetCriteria{} }
func (m *TargetCriteria) String() string { return proto.CompactTextString(m) }
func (*TargetCriteria) ProtoMessage()    {}
func (*TargetCriteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{5}
}

func (m *TargetCriteria) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCriteria.Unmarshal(m, b)
}
func (m *TargetCriteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCriteria.Marshal(b, m, deterministic)
}
func (m *TargetCriteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCriteria.Merge(m, src)
}
func (m *TargetCriteria) XXX_Size() int {
	return xxx_messageInfo_TargetCriteria.Size(m)
}
func (m *TargetCriteria) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCriteria.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCriteria proto.InternalMessageInfo

type isTargetCriteria_TargetType interface {
	isTargetCriteria_TargetType()
}

type TargetCriteria_ReferenceDesign struct {
	ReferenceDesign string `protobuf:"bytes,1,opt,name=reference_design,json=referenceDesign,proto3,oneof"`
}

type TargetCriteria_BuildTarget struct {
	BuildTarget string `protobuf:"bytes,2,opt,name=build_target,json=buildTarget,proto3,oneof"`
}

func (*TargetCriteria_ReferenceDesign) isTargetCriteria_TargetType() {}

func (*TargetCriteria_BuildTarget) isTargetCriteria_TargetType() {}

func (m *TargetCriteria) GetTargetType() isTargetCriteria_TargetType {
	if m != nil {
		return m.TargetType
	}
	return nil
}

func (m *TargetCriteria) GetReferenceDesign() string {
	if x, ok := m.GetTargetType().(*TargetCriteria_ReferenceDesign); ok {
		return x.ReferenceDesign
	}
	return ""
}

func (m *TargetCriteria) GetBuildTarget() string {
	if x, ok := m.GetTargetType().(*TargetCriteria_BuildTarget); ok {
		return x.BuildTarget
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TargetCriteria) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TargetCriteria_ReferenceDesign)(nil),
		(*TargetCriteria_BuildTarget)(nil),
	}
}

// Details which testing is needed for a single CrOS build target.
type PerTargetTestRequirements struct {
	// Specifies the reference design or build target to which these testing
	// requirements should be applied.
	TargetCriteria *TargetCriteria `protobuf:"bytes,1,opt,name=target_criteria,json=targetCriteria,proto3" json:"target_criteria,omitempty"`
	// These configure what testing is needed for these BuildCriteria.
	GceTestCfg           *GceTestCfg      `protobuf:"bytes,2,opt,name=gce_test_cfg,json=gceTestCfg,proto3" json:"gce_test_cfg,omitempty"`
	HwTestCfg            *HwTestCfg       `protobuf:"bytes,3,opt,name=hw_test_cfg,json=hwTestCfg,proto3" json:"hw_test_cfg,omitempty"`
	MoblabVmTestCfg      *MoblabVmTestCfg `protobuf:"bytes,4,opt,name=moblab_vm_test_cfg,json=moblabVmTestCfg,proto3" json:"moblab_vm_test_cfg,omitempty"`
	TastVmTestCfg        *TastVmTestCfg   `protobuf:"bytes,6,opt,name=tast_vm_test_cfg,json=tastVmTestCfg,proto3" json:"tast_vm_test_cfg,omitempty"`
	VmTestCfg            *VmTestCfg       `protobuf:"bytes,5,opt,name=vm_test_cfg,json=vmTestCfg,proto3" json:"vm_test_cfg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PerTargetTestRequirements) Reset()         { *m = PerTargetTestRequirements{} }
func (m *PerTargetTestRequirements) String() string { return proto.CompactTextString(m) }
func (*PerTargetTestRequirements) ProtoMessage()    {}
func (*PerTargetTestRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{6}
}

func (m *PerTargetTestRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerTargetTestRequirements.Unmarshal(m, b)
}
func (m *PerTargetTestRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerTargetTestRequirements.Marshal(b, m, deterministic)
}
func (m *PerTargetTestRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerTargetTestRequirements.Merge(m, src)
}
func (m *PerTargetTestRequirements) XXX_Size() int {
	return xxx_messageInfo_PerTargetTestRequirements.Size(m)
}
func (m *PerTargetTestRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_PerTargetTestRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_PerTargetTestRequirements proto.InternalMessageInfo

func (m *PerTargetTestRequirements) GetTargetCriteria() *TargetCriteria {
	if m != nil {
		return m.TargetCriteria
	}
	return nil
}

func (m *PerTargetTestRequirements) GetGceTestCfg() *GceTestCfg {
	if m != nil {
		return m.GceTestCfg
	}
	return nil
}

func (m *PerTargetTestRequirements) GetHwTestCfg() *HwTestCfg {
	if m != nil {
		return m.HwTestCfg
	}
	return nil
}

func (m *PerTargetTestRequirements) GetMoblabVmTestCfg() *MoblabVmTestCfg {
	if m != nil {
		return m.MoblabVmTestCfg
	}
	return nil
}

func (m *PerTargetTestRequirements) GetTastVmTestCfg() *TastVmTestCfg {
	if m != nil {
		return m.TastVmTestCfg
	}
	return nil
}

func (m *PerTargetTestRequirements) GetVmTestCfg() *VmTestCfg {
	if m != nil {
		return m.VmTestCfg
	}
	return nil
}

// A listing of all testing that should be done for all CrOS builds.
type TargetTestRequirementsCfg struct {
	// The testing that should be performed for a single CrOS build target or
	// reference design.
	PerTargetTestRequirements []*PerTargetTestRequirements `protobuf:"bytes,1,rep,name=per_target_test_requirements,json=perTargetTestRequirements,proto3" json:"per_target_test_requirements,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                     `json:"-"`
	XXX_unrecognized          []byte                       `json:"-"`
	XXX_sizecache             int32                        `json:"-"`
}

func (m *TargetTestRequirementsCfg) Reset()         { *m = TargetTestRequirementsCfg{} }
func (m *TargetTestRequirementsCfg) String() string { return proto.CompactTextString(m) }
func (*TargetTestRequirementsCfg) ProtoMessage()    {}
func (*TargetTestRequirementsCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af54ecd79c6f0f9, []int{7}
}

func (m *TargetTestRequirementsCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetTestRequirementsCfg.Unmarshal(m, b)
}
func (m *TargetTestRequirementsCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetTestRequirementsCfg.Marshal(b, m, deterministic)
}
func (m *TargetTestRequirementsCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetTestRequirementsCfg.Merge(m, src)
}
func (m *TargetTestRequirementsCfg) XXX_Size() int {
	return xxx_messageInfo_TargetTestRequirementsCfg.Size(m)
}
func (m *TargetTestRequirementsCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetTestRequirementsCfg.DiscardUnknown(m)
}

var xxx_messageInfo_TargetTestRequirementsCfg proto.InternalMessageInfo

func (m *TargetTestRequirementsCfg) GetPerTargetTestRequirements() []*PerTargetTestRequirements {
	if m != nil {
		return m.PerTargetTestRequirements
	}
	return nil
}

func init() {
	proto.RegisterType((*GceTestCfg)(nil), "testplans.GceTestCfg")
	proto.RegisterType((*GceTestCfg_GceTest)(nil), "testplans.GceTestCfg.GceTest")
	proto.RegisterType((*HwTestCfg)(nil), "testplans.HwTestCfg")
	proto.RegisterType((*HwTestCfg_HwTest)(nil), "testplans.HwTestCfg.HwTest")
	proto.RegisterType((*MoblabVmTestCfg)(nil), "testplans.MoblabVmTestCfg")
	proto.RegisterType((*MoblabVmTestCfg_MoblabTest)(nil), "testplans.MoblabVmTestCfg.MoblabTest")
	proto.RegisterType((*TastVmTestCfg)(nil), "testplans.TastVmTestCfg")
	proto.RegisterType((*TastVmTestCfg_TastTestExpr)(nil), "testplans.TastVmTestCfg.TastTestExpr")
	proto.RegisterType((*TastVmTestCfg_TastVmTest)(nil), "testplans.TastVmTestCfg.TastVmTest")
	proto.RegisterType((*VmTestCfg)(nil), "testplans.VmTestCfg")
	proto.RegisterType((*VmTestCfg_VmTest)(nil), "testplans.VmTestCfg.VmTest")
	proto.RegisterType((*TargetCriteria)(nil), "testplans.TargetCriteria")
	proto.RegisterType((*PerTargetTestRequirements)(nil), "testplans.PerTargetTestRequirements")
	proto.RegisterType((*TargetTestRequirementsCfg)(nil), "testplans.TargetTestRequirementsCfg")
}

func init() {
	proto.RegisterFile("testplans/target_test_requirements_config.proto", fileDescriptor_4af54ecd79c6f0f9)
}

var fileDescriptor_4af54ecd79c6f0f9 = []byte{
	// 865 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4b, 0x6f, 0xdb, 0x46,
	0x10, 0x0e, 0x25, 0xeb, 0x35, 0x94, 0xed, 0x80, 0x4d, 0x00, 0x5a, 0x69, 0x50, 0x57, 0x49, 0x01,
	0x03, 0x01, 0x24, 0xc0, 0x71, 0xd1, 0x5e, 0x6b, 0xe7, 0x61, 0xb4, 0x48, 0x5b, 0x6c, 0x8c, 0x1e,
	0x7a, 0x59, 0x50, 0xf4, 0x92, 0x5e, 0x80, 0x4b, 0x32, 0xbb, 0x4b, 0xdb, 0xba, 0xb6, 0xe7, 0x02,
	0xbd, 0xb6, 0x97, 0x9e, 0xfa, 0x63, 0x0a, 0xf4, 0xdc, 0xdf, 0x53, 0xec, 0x2c, 0x5f, 0x92, 0xa5,
	0xe4, 0x94, 0x93, 0x76, 0xbe, 0xd9, 0x99, 0xfd, 0xbe, 0x1d, 0xce, 0xac, 0x60, 0xae, 0x99, 0xd2,
	0x79, 0x12, 0xa4, 0x6a, 0xae, 0x03, 0x19, 0x33, 0x4d, 0x0d, 0x40, 0x25, 0x7b, 0x57, 0x70, 0xc9,
	0x04, 0x4b, 0xb5, 0xa2, 0x61, 0x96, 0x46, 0x3c, 0x9e, 0xe5, 0x32, 0xd3, 0x99, 0x37, 0xaa, 0x03,
	0xa6, 0xff, 0x3a, 0x00, 0xaf, 0x43, 0x76, 0xc1, 0x94, 0x3e, 0x8b, 0x62, 0xef, 0x6b, 0x18, 0xc6,
	0x21, 0xc3, 0x78, 0xdf, 0x39, 0xec, 0x1e, 0xb9, 0xc7, 0x8f, 0x67, 0xf5, 0xe6, 0x59, 0xb3, 0xb1,
	0x5a, 0x92, 0x41, 0x6c, 0x17, 0x93, 0x5f, 0x1d, 0x18, 0x94, 0xa0, 0xf7, 0x08, 0xf0, 0x04, 0xaa,
	0x97, 0x39, 0xf3, 0x9d, 0x43, 0xe7, 0x68, 0x44, 0x86, 0x06, 0xb8, 0x58, 0xe6, 0xcc, 0x7b, 0x0c,
	0x80, 0x4e, 0x55, 0x70, 0xcd, 0xfc, 0x0e, 0x7a, 0x71, 0xfb, 0x5b, 0x03, 0x78, 0x9f, 0x81, 0xab,
	0xb9, 0x60, 0x59, 0xa1, 0xa9, 0x62, 0xa1, 0xdf, 0x3d, 0x74, 0x8e, 0x7a, 0x04, 0x4a, 0xe8, 0x2d,
	0x0b, 0x4d, 0xf2, 0x42, 0x31, 0x1a, 0x22, 0xc7, 0x9d, 0x43, 0xe7, 0x68, 0x48, 0x86, 0x85, 0x62,
	0x67, 0xc6, 0x9e, 0xfe, 0xd6, 0x85, 0xd1, 0xf9, 0x4d, 0xa5, 0xe6, 0x04, 0x06, 0x57, 0x37, 0x6d,
	0x31, 0x8f, 0x5a, 0x62, 0xea, 0x6d, 0xe5, 0x8a, 0xf4, 0xaf, 0xf0, 0x77, 0xf2, 0x4f, 0x07, 0xfa,
	0x16, 0xf2, 0x1e, 0x40, 0xcf, 0xd2, 0xb4, 0x22, 0xac, 0xb1, 0x4e, 0xb1, 0xb3, 0x89, 0xe2, 0x4d,
	0x20, 0x53, 0x9a, 0xa5, 0xc9, 0x12, 0x15, 0x0c, 0xc9, 0xd0, 0x00, 0x3f, 0xa4, 0xc9, 0xd2, 0x9b,
	0xc0, 0x30, 0x94, 0x5c, 0xf3, 0x30, 0x48, 0x2a, 0xfa, 0x95, 0x6d, 0x02, 0x23, 0x9e, 0x30, 0xba,
	0x28, 0x62, 0xe5, 0xf7, 0xac, 0xd3, 0x00, 0xa7, 0x45, 0xac, 0xbc, 0xcf, 0x61, 0x2c, 0x78, 0xca,
	0x45, 0x21, 0xe8, 0x65, 0xa1, 0x95, 0xdf, 0xc7, 0x73, 0xdd, 0x12, 0x7b, 0x51, 0x68, 0x65, 0xf8,
	0x4a, 0xa6, 0xe5, 0xd2, 0x1f, 0x60, 0xac, 0x35, 0x0c, 0x5f, 0x11, 0xdc, 0x52, 0x63, 0x70, 0xa6,
	0xfc, 0xa1, 0xe5, 0x2b, 0x82, 0x5b, 0x62, 0x11, 0xef, 0x29, 0xec, 0xa1, 0x32, 0x2a, 0x78, 0x6a,
	0x73, 0x8f, 0x70, 0xcf, 0x18, 0xd1, 0x37, 0x3c, 0xc5, 0xe4, 0xc7, 0xf0, 0x30, 0x8b, 0xa2, 0x24,
	0x0b, 0x2e, 0x69, 0x14, 0xf0, 0xa4, 0x90, 0x4c, 0x59, 0x85, 0x80, 0x87, 0x7d, 0x52, 0x3a, 0x5f,
	0x95, 0x3e, 0x23, 0x76, 0xfa, 0xb7, 0x03, 0xfb, 0x6f, 0xb2, 0x45, 0x12, 0x2c, 0x7e, 0x12, 0x55,
	0x55, 0x5e, 0x81, 0x2b, 0x10, 0x6a, 0x57, 0xe6, 0x8b, 0x56, 0x65, 0xd6, 0x02, 0x4a, 0x1b, 0x6b,
	0x04, 0xa2, 0x5e, 0x4f, 0xbe, 0x05, 0x68, 0x3c, 0xef, 0xff, 0xe6, 0x3e, 0x54, 0xb1, 0xe9, 0x5f,
	0x1d, 0xd8, 0xbd, 0x08, 0x94, 0x6e, 0x58, 0xbe, 0x84, 0xb1, 0x0e, 0x94, 0xa6, 0xd7, 0xa2, 0x4d,
	0xf3, 0x49, 0x8b, 0xe6, 0xca, 0xfe, 0x96, 0x45, 0x40, 0xd7, 0xeb, 0xc9, 0x33, 0x18, 0x1b, 0x8f,
	0x59, 0xbf, 0xbc, 0xcd, 0x65, 0x4d, 0x93, 0xdd, 0xe6, 0xb2, 0x4d, 0xd3, 0x38, 0x27, 0x7f, 0x38,
	0x00, 0x4d, 0x1e, 0xd3, 0x29, 0xb6, 0x2c, 0x69, 0x20, 0x2a, 0x4d, 0x23, 0x44, 0xbe, 0x0f, 0x04,
	0xf3, 0xbe, 0x83, 0x3d, 0x64, 0xd8, 0xe4, 0xeb, 0xdc, 0xb9, 0xca, 0xbb, 0x1c, 0x2b, 0x26, 0x04,
	0xe5, 0xd5, 0xbc, 0x3e, 0xd4, 0x76, 0xd3, 0x3f, 0x3b, 0x30, 0x6a, 0x6e, 0xe7, 0x04, 0x06, 0xab,
	0x17, 0xd3, 0xee, 0xac, 0xe6, 0xc0, 0xf2, 0x42, 0xfa, 0xd7, 0xf6, 0x32, 0xfe, 0x73, 0xa0, 0x5f,
	0x6a, 0xfb, 0xa8, 0x23, 0xa2, 0x6e, 0x83, 0x9d, 0xf7, 0xb4, 0x41, 0xef, 0x4e, 0x1b, 0xac, 0xb4,
	0x6d, 0x7f, 0xad, 0x6d, 0x57, 0xc6, 0xce, 0x60, 0x6d, 0xec, 0xbc, 0x83, 0xbd, 0x0b, 0x9c, 0xbc,
	0x67, 0x92, 0x6b, 0x26, 0x79, 0xe0, 0x3d, 0x83, 0xfb, 0x92, 0x45, 0x4c, 0xb2, 0x34, 0x64, 0xf4,
	0x92, 0x29, 0x1e, 0xa7, 0x56, 0xe6, 0xf9, 0x3d, 0xb2, 0x5f, 0x7b, 0x5e, 0xa0, 0xc3, 0x7b, 0x02,
	0xe3, 0x45, 0xc1, 0x93, 0x4b, 0x6a, 0xc7, 0xb7, 0x55, 0x7c, 0x7e, 0x8f, 0xb8, 0x88, 0xda, 0xcc,
	0xa7, 0xbb, 0xe0, 0x56, 0xd3, 0x7d, 0x99, 0xb3, 0xe9, 0xef, 0x5d, 0x38, 0xf8, 0x91, 0x49, 0xeb,
	0xc4, 0x5b, 0x6e, 0x8d, 0x7b, 0xef, 0x14, 0xf6, 0xcb, 0xcd, 0x61, 0xc9, 0x08, 0x4f, 0x77, 0x8f,
	0x0f, 0x56, 0x3e, 0x8e, 0x36, 0x65, 0xb2, 0xa7, 0x57, 0x25, 0x7c, 0x05, 0xe3, 0xea, 0x2d, 0xa0,
	0x61, 0x14, 0x23, 0x2b, 0xf7, 0xf8, 0xe1, 0xc6, 0xf7, 0x80, 0x40, 0xdc, 0x3c, 0x22, 0x27, 0xe0,
	0x96, 0x63, 0x17, 0xe3, 0xba, 0x18, 0xf7, 0x60, 0xd3, 0xe8, 0x25, 0xa3, 0xab, 0x7a, 0x58, 0xbf,
	0x06, 0xaf, 0x1c, 0x0b, 0xe5, 0x97, 0x85, 0xc1, 0x3b, 0x18, 0x3c, 0xd9, 0x3e, 0x1d, 0xc8, 0xbe,
	0x58, 0x9b, 0x2f, 0xdf, 0xc0, 0xfd, 0x76, 0xe7, 0x62, 0x9a, 0x3e, 0xa6, 0xf1, 0xb7, 0x75, 0x06,
	0xd9, 0xd5, 0x2b, 0xcd, 0x7f, 0x02, 0x6e, 0x3b, 0xba, 0x77, 0x47, 0x41, 0x13, 0x39, 0xba, 0xae,
	0x96, 0xd3, 0x5f, 0x1c, 0x38, 0xd8, 0x5c, 0x0f, 0x93, 0x93, 0xc1, 0xa7, 0x39, 0x93, 0x74, 0xdb,
	0x0b, 0x5d, 0xf6, 0xd1, 0xd3, 0xd6, 0x21, 0x5b, 0xcb, 0x4b, 0x0e, 0xf2, 0x6d, 0xae, 0xd3, 0x2f,
	0x7f, 0x7e, 0x1e, 0x67, 0xb3, 0xf0, 0x4a, 0x66, 0x82, 0x17, 0x62, 0x96, 0xc9, 0x78, 0x5e, 0x19,
	0x99, 0x9a, 0xf3, 0x34, 0x92, 0xc1, 0x1c, 0xff, 0x05, 0xcc, 0xe3, 0xac, 0xf9, 0xe3, 0xb0, 0xe8,
	0x23, 0xf6, 0xfc, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x41, 0x2c, 0x83, 0x4c, 0x08, 0x00,
	0x00,
}
