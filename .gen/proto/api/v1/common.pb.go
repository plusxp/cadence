// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: uber/cadence/api/v1/common.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EncodingType int32

const (
	EncodingType_ENCODING_TYPE_INVALID  EncodingType = 0
	EncodingType_ENCODING_TYPE_THRIFTRW EncodingType = 1
	EncodingType_ENCODING_TYPE_JSON     EncodingType = 2
	EncodingType_ENCODING_TYPE_PROTO3   EncodingType = 3
)

// Enum value maps for EncodingType.
var (
	EncodingType_name = map[int32]string{
		0: "ENCODING_TYPE_INVALID",
		1: "ENCODING_TYPE_THRIFTRW",
		2: "ENCODING_TYPE_JSON",
		3: "ENCODING_TYPE_PROTO3",
	}
	EncodingType_value = map[string]int32{
		"ENCODING_TYPE_INVALID":  0,
		"ENCODING_TYPE_THRIFTRW": 1,
		"ENCODING_TYPE_JSON":     2,
		"ENCODING_TYPE_PROTO3":   3,
	}
)

func (x EncodingType) Enum() *EncodingType {
	p := new(EncodingType)
	*p = x
	return p
}

func (x EncodingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncodingType) Descriptor() protoreflect.EnumDescriptor {
	return file_uber_cadence_api_v1_common_proto_enumTypes[0].Descriptor()
}

func (EncodingType) Type() protoreflect.EnumType {
	return &file_uber_cadence_api_v1_common_proto_enumTypes[0]
}

func (x EncodingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncodingType.Descriptor instead.
func (EncodingType) EnumDescriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{0}
}

type WorkflowExecution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId      string `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *WorkflowExecution) Reset() {
	*x = WorkflowExecution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowExecution) ProtoMessage() {}

func (x *WorkflowExecution) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowExecution.ProtoReflect.Descriptor instead.
func (*WorkflowExecution) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowExecution) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *WorkflowExecution) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

type WorkflowType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *WorkflowType) Reset() {
	*x = WorkflowType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowType) ProtoMessage() {}

func (x *WorkflowType) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowType.ProtoReflect.Descriptor instead.
func (*WorkflowType) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ActivityType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ActivityType) Reset() {
	*x = ActivityType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityType) ProtoMessage() {}

func (x *ActivityType) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityType.ProtoReflect.Descriptor instead.
func (*ActivityType) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *ActivityType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *Payload) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Failure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason  string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	Details []byte `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Failure) Reset() {
	*x = Failure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failure) ProtoMessage() {}

func (x *Failure) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Failure.ProtoReflect.Descriptor instead.
func (*Failure) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *Failure) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Failure) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

type Memo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields map[string]*Payload `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Memo) Reset() {
	*x = Memo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memo) ProtoMessage() {}

func (x *Memo) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memo.ProtoReflect.Descriptor instead.
func (*Memo) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{5}
}

func (x *Memo) GetFields() map[string]*Payload {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields map[string]*Payload `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{6}
}

func (x *Header) GetFields() map[string]*Payload {
	if x != nil {
		return x.Fields
	}
	return nil
}

type SearchAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IndexedFields map[string]*Payload `protobuf:"bytes,1,rep,name=indexed_fields,json=indexedFields,proto3" json:"indexed_fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SearchAttributes) Reset() {
	*x = SearchAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAttributes) ProtoMessage() {}

func (x *SearchAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAttributes.ProtoReflect.Descriptor instead.
func (*SearchAttributes) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{7}
}

func (x *SearchAttributes) GetIndexedFields() map[string]*Payload {
	if x != nil {
		return x.IndexedFields
	}
	return nil
}

type DataBlob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncodingType EncodingType `protobuf:"varint,1,opt,name=encoding_type,json=encodingType,proto3,enum=uber.cadence.api.v1.EncodingType" json:"encoding_type,omitempty"`
	Data         []byte       `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataBlob) Reset() {
	*x = DataBlob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBlob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBlob) ProtoMessage() {}

func (x *DataBlob) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBlob.ProtoReflect.Descriptor instead.
func (*DataBlob) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{8}
}

func (x *DataBlob) GetEncodingType() EncodingType {
	if x != nil {
		return x.EncodingType
	}
	return EncodingType_ENCODING_TYPE_INVALID
}

func (x *DataBlob) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type WorkerVersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Impl           string `protobuf:"bytes,1,opt,name=impl,proto3" json:"impl,omitempty"`
	FeatureVersion string `protobuf:"bytes,2,opt,name=feature_version,json=featureVersion,proto3" json:"feature_version,omitempty"`
}

func (x *WorkerVersionInfo) Reset() {
	*x = WorkerVersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerVersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerVersionInfo) ProtoMessage() {}

func (x *WorkerVersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerVersionInfo.ProtoReflect.Descriptor instead.
func (*WorkerVersionInfo) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{9}
}

func (x *WorkerVersionInfo) GetImpl() string {
	if x != nil {
		return x.Impl
	}
	return ""
}

func (x *WorkerVersionInfo) GetFeatureVersion() string {
	if x != nil {
		return x.FeatureVersion
	}
	return ""
}

type SupportedClientVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates the highest Go SDK version server will accept requests from.
	GoSdk string `protobuf:"bytes,1,opt,name=go_sdk,json=goSdk,proto3" json:"go_sdk,omitempty"`
	// Indicates the highest Java SDK version server will accept requests from.
	JavaSdk string `protobuf:"bytes,2,opt,name=java_sdk,json=javaSdk,proto3" json:"java_sdk,omitempty"`
}

func (x *SupportedClientVersions) Reset() {
	*x = SupportedClientVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportedClientVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportedClientVersions) ProtoMessage() {}

func (x *SupportedClientVersions) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportedClientVersions.ProtoReflect.Descriptor instead.
func (*SupportedClientVersions) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{10}
}

func (x *SupportedClientVersions) GetGoSdk() string {
	if x != nil {
		return x.GoSdk
	}
	return ""
}

func (x *SupportedClientVersions) GetJavaSdk() string {
	if x != nil {
		return x.JavaSdk
	}
	return ""
}

type RetryPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interval of the first retry. If backoffCoefficient is 1.0 then it is used for all retries.
	InitialInterval *durationpb.Duration `protobuf:"bytes,1,opt,name=initial_interval,json=initialInterval,proto3" json:"initial_interval,omitempty"`
	// Coefficient used to calculate the next retry interval.
	// The next retry interval is previous interval multiplied by the coefficient.
	// Must be 1 or larger.
	BackoffCoefficient float64 `protobuf:"fixed64,2,opt,name=backoff_coefficient,json=backoffCoefficient,proto3" json:"backoff_coefficient,omitempty"`
	// Maximum interval between retries. Exponential backoff leads to interval increase.
	// This value is the cap of the increase. Default is 100x of initial interval.
	MaximumInterval *durationpb.Duration `protobuf:"bytes,3,opt,name=maximum_interval,json=maximumInterval,proto3" json:"maximum_interval,omitempty"`
	// Maximum number of attempts. When exceeded the retries stop even if not expired yet.
	// Must be 1 or bigger. Default is unlimited.
	MaximumAttempts int32 `protobuf:"varint,4,opt,name=maximum_attempts,json=maximumAttempts,proto3" json:"maximum_attempts,omitempty"`
	// Non-Retryable errors. Will stop retrying if error type matches this list.
	NonRetryableErrorReasons []string `protobuf:"bytes,5,rep,name=non_retryable_error_reasons,json=nonRetryableErrorReasons,proto3" json:"non_retryable_error_reasons,omitempty"`
	// Expiration time for the whole retry process.
	ExpirationInterval *durationpb.Duration `protobuf:"bytes,6,opt,name=expiration_interval,json=expirationInterval,proto3" json:"expiration_interval,omitempty"`
}

func (x *RetryPolicy) Reset() {
	*x = RetryPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_common_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryPolicy) ProtoMessage() {}

func (x *RetryPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_common_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryPolicy.ProtoReflect.Descriptor instead.
func (*RetryPolicy) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_common_proto_rawDescGZIP(), []int{11}
}

func (x *RetryPolicy) GetInitialInterval() *durationpb.Duration {
	if x != nil {
		return x.InitialInterval
	}
	return nil
}

func (x *RetryPolicy) GetBackoffCoefficient() float64 {
	if x != nil {
		return x.BackoffCoefficient
	}
	return 0
}

func (x *RetryPolicy) GetMaximumInterval() *durationpb.Duration {
	if x != nil {
		return x.MaximumInterval
	}
	return nil
}

func (x *RetryPolicy) GetMaximumAttempts() int32 {
	if x != nil {
		return x.MaximumAttempts
	}
	return 0
}

func (x *RetryPolicy) GetNonRetryableErrorReasons() []string {
	if x != nil {
		return x.NonRetryableErrorReasons
	}
	return nil
}

func (x *RetryPolicy) GetExpirationInterval() *durationpb.Duration {
	if x != nil {
		return x.ExpirationInterval
	}
	return nil
}

var File_uber_cadence_api_v1_common_proto protoreflect.FileDescriptor

var file_uber_cadence_api_v1_common_proto_rawDesc = []byte{
	0x0a, 0x20, 0x75, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x75, 0x6e, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x07,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3b, 0x0a, 0x07, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x6d,
	0x6f, 0x12, 0x3d, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x1a, 0x57, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa2, 0x01, 0x0a, 0x06, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x57, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd3,
	0x01, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x5f, 0x0a, 0x0e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x75, 0x62,
	0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x1a, 0x5e, 0x0a, 0x12, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x75, 0x62,
	0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x66, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6c, 0x6f, 0x62,
	0x12, 0x46, 0x0a, 0x0d, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63,
	0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x50, 0x0a, 0x11,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x70, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x6d, 0x70, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4b,
	0x0a, 0x17, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x67, 0x6f, 0x5f,
	0x73, 0x64, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6f, 0x53, 0x64, 0x6b,
	0x12, 0x19, 0x0a, 0x08, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x73, 0x64, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6a, 0x61, 0x76, 0x61, 0x53, 0x64, 0x6b, 0x22, 0x80, 0x03, 0x0a, 0x0b,
	0x52, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x44, 0x0a, 0x10, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x2f, 0x0a, 0x13, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x5f, 0x63, 0x6f, 0x65,
	0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12,
	0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x43, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x41, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x6e, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x6e, 0x6f, 0x6e, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x73, 0x12, 0x4a, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2a, 0x77,
	0x0a, 0x0c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x15, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x43,
	0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x48, 0x52, 0x49, 0x46,
	0x54, 0x52, 0x57, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x18, 0x0a,
	0x14, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x33, 0x10, 0x03, 0x42, 0x56, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x75,
	0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x62, 0x65, 0x72, 0x2f,
	0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x2e, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uber_cadence_api_v1_common_proto_rawDescOnce sync.Once
	file_uber_cadence_api_v1_common_proto_rawDescData = file_uber_cadence_api_v1_common_proto_rawDesc
)

func file_uber_cadence_api_v1_common_proto_rawDescGZIP() []byte {
	file_uber_cadence_api_v1_common_proto_rawDescOnce.Do(func() {
		file_uber_cadence_api_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_uber_cadence_api_v1_common_proto_rawDescData)
	})
	return file_uber_cadence_api_v1_common_proto_rawDescData
}

var file_uber_cadence_api_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_uber_cadence_api_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_uber_cadence_api_v1_common_proto_goTypes = []interface{}{
	(EncodingType)(0),               // 0: uber.cadence.api.v1.EncodingType
	(*WorkflowExecution)(nil),       // 1: uber.cadence.api.v1.WorkflowExecution
	(*WorkflowType)(nil),            // 2: uber.cadence.api.v1.WorkflowType
	(*ActivityType)(nil),            // 3: uber.cadence.api.v1.ActivityType
	(*Payload)(nil),                 // 4: uber.cadence.api.v1.Payload
	(*Failure)(nil),                 // 5: uber.cadence.api.v1.Failure
	(*Memo)(nil),                    // 6: uber.cadence.api.v1.Memo
	(*Header)(nil),                  // 7: uber.cadence.api.v1.Header
	(*SearchAttributes)(nil),        // 8: uber.cadence.api.v1.SearchAttributes
	(*DataBlob)(nil),                // 9: uber.cadence.api.v1.DataBlob
	(*WorkerVersionInfo)(nil),       // 10: uber.cadence.api.v1.WorkerVersionInfo
	(*SupportedClientVersions)(nil), // 11: uber.cadence.api.v1.SupportedClientVersions
	(*RetryPolicy)(nil),             // 12: uber.cadence.api.v1.RetryPolicy
	nil,                             // 13: uber.cadence.api.v1.Memo.FieldsEntry
	nil,                             // 14: uber.cadence.api.v1.Header.FieldsEntry
	nil,                             // 15: uber.cadence.api.v1.SearchAttributes.IndexedFieldsEntry
	(*durationpb.Duration)(nil),     // 16: google.protobuf.Duration
}
var file_uber_cadence_api_v1_common_proto_depIdxs = []int32{
	13, // 0: uber.cadence.api.v1.Memo.fields:type_name -> uber.cadence.api.v1.Memo.FieldsEntry
	14, // 1: uber.cadence.api.v1.Header.fields:type_name -> uber.cadence.api.v1.Header.FieldsEntry
	15, // 2: uber.cadence.api.v1.SearchAttributes.indexed_fields:type_name -> uber.cadence.api.v1.SearchAttributes.IndexedFieldsEntry
	0,  // 3: uber.cadence.api.v1.DataBlob.encoding_type:type_name -> uber.cadence.api.v1.EncodingType
	16, // 4: uber.cadence.api.v1.RetryPolicy.initial_interval:type_name -> google.protobuf.Duration
	16, // 5: uber.cadence.api.v1.RetryPolicy.maximum_interval:type_name -> google.protobuf.Duration
	16, // 6: uber.cadence.api.v1.RetryPolicy.expiration_interval:type_name -> google.protobuf.Duration
	4,  // 7: uber.cadence.api.v1.Memo.FieldsEntry.value:type_name -> uber.cadence.api.v1.Payload
	4,  // 8: uber.cadence.api.v1.Header.FieldsEntry.value:type_name -> uber.cadence.api.v1.Payload
	4,  // 9: uber.cadence.api.v1.SearchAttributes.IndexedFieldsEntry.value:type_name -> uber.cadence.api.v1.Payload
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_uber_cadence_api_v1_common_proto_init() }
func file_uber_cadence_api_v1_common_proto_init() {
	if File_uber_cadence_api_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uber_cadence_api_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowExecution); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Failure); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchAttributes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBlob); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerVersionInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportedClientVersions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_uber_cadence_api_v1_common_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_uber_cadence_api_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uber_cadence_api_v1_common_proto_goTypes,
		DependencyIndexes: file_uber_cadence_api_v1_common_proto_depIdxs,
		EnumInfos:         file_uber_cadence_api_v1_common_proto_enumTypes,
		MessageInfos:      file_uber_cadence_api_v1_common_proto_msgTypes,
	}.Build()
	File_uber_cadence_api_v1_common_proto = out.File
	file_uber_cadence_api_v1_common_proto_rawDesc = nil
	file_uber_cadence_api_v1_common_proto_goTypes = nil
	file_uber_cadence_api_v1_common_proto_depIdxs = nil
}
