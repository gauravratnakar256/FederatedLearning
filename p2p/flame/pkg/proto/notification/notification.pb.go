// Copyright 2022 Cisco Systems, Inc. and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: notification.proto

package notification

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobEventType int32

const (
	JobEventType_UNKNOWN_EVENT_TYPE JobEventType = 0 // default
	JobEventType_START_JOB          JobEventType = 1
	JobEventType_STOP_JOB           JobEventType = 2
	JobEventType_UPDATE_JOB         JobEventType = 3
)

// Enum value maps for JobEventType.
var (
	JobEventType_name = map[int32]string{
		0: "UNKNOWN_EVENT_TYPE",
		1: "START_JOB",
		2: "STOP_JOB",
		3: "UPDATE_JOB",
	}
	JobEventType_value = map[string]int32{
		"UNKNOWN_EVENT_TYPE": 0,
		"START_JOB":          1,
		"STOP_JOB":           2,
		"UPDATE_JOB":         3,
	}
)

func (x JobEventType) Enum() *JobEventType {
	p := new(JobEventType)
	*p = x
	return p
}

func (x JobEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_proto_enumTypes[0].Descriptor()
}

func (JobEventType) Type() protoreflect.EnumType {
	return &file_notification_proto_enumTypes[0]
}

func (x JobEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobEventType.Descriptor instead.
func (JobEventType) EnumDescriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{0}
}

type DeployEventType int32

const (
	DeployEventType_UNKNOWN_DEPLOYMENT_TYPE DeployEventType = 0 // default
	DeployEventType_ADD_RESOURCE            DeployEventType = 1
	DeployEventType_REVOKE_RESOURCE         DeployEventType = 2
)

// Enum value maps for DeployEventType.
var (
	DeployEventType_name = map[int32]string{
		0: "UNKNOWN_DEPLOYMENT_TYPE",
		1: "ADD_RESOURCE",
		2: "REVOKE_RESOURCE",
	}
	DeployEventType_value = map[string]int32{
		"UNKNOWN_DEPLOYMENT_TYPE": 0,
		"ADD_RESOURCE":            1,
		"REVOKE_RESOURCE":         2,
	}
)

func (x DeployEventType) Enum() *DeployEventType {
	p := new(DeployEventType)
	*p = x
	return p
}

func (x DeployEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeployEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_proto_enumTypes[1].Descriptor()
}

func (DeployEventType) Type() protoreflect.EnumType {
	return &file_notification_proto_enumTypes[1]
}

func (x DeployEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeployEventType.Descriptor instead.
func (DeployEventType) EnumDescriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{1}
}

type JobResponse_Status int32

const (
	JobResponse_ERROR           JobResponse_Status = 0 // default
	JobResponse_PARTIAL_SUCCESS JobResponse_Status = 1 // e.g., not all agents received event
	JobResponse_SUCCESS         JobResponse_Status = 2
)

// Enum value maps for JobResponse_Status.
var (
	JobResponse_Status_name = map[int32]string{
		0: "ERROR",
		1: "PARTIAL_SUCCESS",
		2: "SUCCESS",
	}
	JobResponse_Status_value = map[string]int32{
		"ERROR":           0,
		"PARTIAL_SUCCESS": 1,
		"SUCCESS":         2,
	}
)

func (x JobResponse_Status) Enum() *JobResponse_Status {
	p := new(JobResponse_Status)
	*p = x
	return p
}

func (x JobResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_proto_enumTypes[2].Descriptor()
}

func (JobResponse_Status) Type() protoreflect.EnumType {
	return &file_notification_proto_enumTypes[2]
}

func (x JobResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobResponse_Status.Descriptor instead.
func (JobResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{5, 0}
}

type DeployResponse_Status int32

const (
	DeployResponse_ERROR           DeployResponse_Status = 0 // default
	DeployResponse_PARTIAL_SUCCESS DeployResponse_Status = 1 // e.g., not all deployers received event
	DeployResponse_SUCCESS         DeployResponse_Status = 2
)

// Enum value maps for DeployResponse_Status.
var (
	DeployResponse_Status_name = map[int32]string{
		0: "ERROR",
		1: "PARTIAL_SUCCESS",
		2: "SUCCESS",
	}
	DeployResponse_Status_value = map[string]int32{
		"ERROR":           0,
		"PARTIAL_SUCCESS": 1,
		"SUCCESS":         2,
	}
)

func (x DeployResponse_Status) Enum() *DeployResponse_Status {
	p := new(DeployResponse_Status)
	*p = x
	return p
}

func (x DeployResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeployResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_proto_enumTypes[3].Descriptor()
}

func (DeployResponse_Status) Type() protoreflect.EnumType {
	return &file_notification_proto_enumTypes[3]
}

func (x DeployResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeployResponse_Status.Descriptor instead.
func (DeployResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{7, 0}
}

type JobTaskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *JobTaskInfo) Reset() {
	*x = JobTaskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobTaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTaskInfo) ProtoMessage() {}

func (x *JobTaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTaskInfo.ProtoReflect.Descriptor instead.
func (*JobTaskInfo) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{0}
}

func (x *JobTaskInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobTaskInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type JobEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  JobEventType `protobuf:"varint,1,opt,name=type,proto3,enum=grpcNotification.JobEventType" json:"type,omitempty"`
	JobId string       `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *JobEvent) Reset() {
	*x = JobEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobEvent) ProtoMessage() {}

func (x *JobEvent) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobEvent.ProtoReflect.Descriptor instead.
func (*JobEvent) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{1}
}

func (x *JobEvent) GetType() JobEventType {
	if x != nil {
		return x.Type
	}
	return JobEventType_UNKNOWN_EVENT_TYPE
}

func (x *JobEvent) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type DeployInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComputeId string `protobuf:"bytes,1,opt,name=compute_id,json=computeId,proto3" json:"compute_id,omitempty"`
	ApiKey    string `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *DeployInfo) Reset() {
	*x = DeployInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployInfo) ProtoMessage() {}

func (x *DeployInfo) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployInfo.ProtoReflect.Descriptor instead.
func (*DeployInfo) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{2}
}

func (x *DeployInfo) GetComputeId() string {
	if x != nil {
		return x.ComputeId
	}
	return ""
}

func (x *DeployInfo) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type DeployEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  DeployEventType `protobuf:"varint,1,opt,name=type,proto3,enum=grpcNotification.DeployEventType" json:"type,omitempty"`
	JobId string          `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *DeployEvent) Reset() {
	*x = DeployEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployEvent) ProtoMessage() {}

func (x *DeployEvent) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployEvent.ProtoReflect.Descriptor instead.
func (*DeployEvent) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{3}
}

func (x *DeployEvent) GetType() DeployEventType {
	if x != nil {
		return x.Type
	}
	return DeployEventType_UNKNOWN_DEPLOYMENT_TYPE
}

func (x *DeployEvent) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type JobEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    JobEventType `protobuf:"varint,1,opt,name=type,proto3,enum=grpcNotification.JobEventType" json:"type,omitempty"`
	JobId   string       `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	TaskIds []string     `protobuf:"bytes,3,rep,name=task_ids,json=taskIds,proto3" json:"task_ids,omitempty"`
}

func (x *JobEventRequest) Reset() {
	*x = JobEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobEventRequest) ProtoMessage() {}

func (x *JobEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobEventRequest.ProtoReflect.Descriptor instead.
func (*JobEventRequest) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{4}
}

func (x *JobEventRequest) GetType() JobEventType {
	if x != nil {
		return x.Type
	}
	return JobEventType_UNKNOWN_EVENT_TYPE
}

func (x *JobEventRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobEventRequest) GetTaskIds() []string {
	if x != nil {
		return x.TaskIds
	}
	return nil
}

type JobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      JobResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=grpcNotification.JobResponse_Status" json:"status,omitempty"`
	Message     string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FailedTasks []string           `protobuf:"bytes,3,rep,name=failed_tasks,json=failedTasks,proto3" json:"failed_tasks,omitempty"`
}

func (x *JobResponse) Reset() {
	*x = JobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResponse) ProtoMessage() {}

func (x *JobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResponse.ProtoReflect.Descriptor instead.
func (*JobResponse) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{5}
}

func (x *JobResponse) GetStatus() JobResponse_Status {
	if x != nil {
		return x.Status
	}
	return JobResponse_ERROR
}

func (x *JobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *JobResponse) GetFailedTasks() []string {
	if x != nil {
		return x.FailedTasks
	}
	return nil
}

type DeployEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       DeployEventType `protobuf:"varint,1,opt,name=type,proto3,enum=grpcNotification.DeployEventType" json:"type,omitempty"`
	ComputeIds []string        `protobuf:"bytes,2,rep,name=compute_ids,json=computeIds,proto3" json:"compute_ids,omitempty"`
	JobId      string          `protobuf:"bytes,3,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *DeployEventRequest) Reset() {
	*x = DeployEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployEventRequest) ProtoMessage() {}

func (x *DeployEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployEventRequest.ProtoReflect.Descriptor instead.
func (*DeployEventRequest) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{6}
}

func (x *DeployEventRequest) GetType() DeployEventType {
	if x != nil {
		return x.Type
	}
	return DeployEventType_UNKNOWN_DEPLOYMENT_TYPE
}

func (x *DeployEventRequest) GetComputeIds() []string {
	if x != nil {
		return x.ComputeIds
	}
	return nil
}

func (x *DeployEventRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type DeployResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          DeployResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=grpcNotification.DeployResponse_Status" json:"status,omitempty"`
	Message         string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FailedDeployers []string              `protobuf:"bytes,3,rep,name=failed_deployers,json=failedDeployers,proto3" json:"failed_deployers,omitempty"`
}

func (x *DeployResponse) Reset() {
	*x = DeployResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployResponse) ProtoMessage() {}

func (x *DeployResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployResponse.ProtoReflect.Descriptor instead.
func (*DeployResponse) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{7}
}

func (x *DeployResponse) GetStatus() DeployResponse_Status {
	if x != nil {
		return x.Status
	}
	return DeployResponse_ERROR
}

func (x *DeployResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeployResponse) GetFailedDeployers() []string {
	if x != nil {
		return x.FailedDeployers
	}
	return nil
}

var File_notification_proto protoreflect.FileDescriptor

var file_notification_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x72, 0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x55, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a,
	0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x22, 0x5b,
	0x0a, 0x0b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x77, 0x0a, 0x0f, 0x4a,
	0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4a, 0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x22,
	0x35, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x49, 0x64, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0xcd, 0x01, 0x0a,
	0x0e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x72, 0x73, 0x22, 0x35, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41,
	0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x2a, 0x53, 0x0a, 0x0c,
	0x4a, 0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x4a, 0x4f,
	0x42, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x4a, 0x4f, 0x42, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4a, 0x4f, 0x42, 0x10,
	0x03, 0x2a, 0x55, 0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x44, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x02, 0x32, 0x5d, 0x0a, 0x0d, 0x4a, 0x6f, 0x62, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x4a, 0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x54,
	0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x32, 0x65, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x32, 0x62,
	0x0a, 0x0f, 0x4a, 0x6f, 0x62, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x4f, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4a, 0x6f, 0x62, 0x12, 0x21,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0x6e, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x58, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x66, 0x6c, 0x61, 0x6d,
	0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_proto_rawDescOnce sync.Once
	file_notification_proto_rawDescData = file_notification_proto_rawDesc
)

func file_notification_proto_rawDescGZIP() []byte {
	file_notification_proto_rawDescOnce.Do(func() {
		file_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_proto_rawDescData)
	})
	return file_notification_proto_rawDescData
}

var file_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_notification_proto_goTypes = []interface{}{
	(JobEventType)(0),          // 0: grpcNotification.JobEventType
	(DeployEventType)(0),       // 1: grpcNotification.DeployEventType
	(JobResponse_Status)(0),    // 2: grpcNotification.JobResponse.Status
	(DeployResponse_Status)(0), // 3: grpcNotification.DeployResponse.Status
	(*JobTaskInfo)(nil),        // 4: grpcNotification.JobTaskInfo
	(*JobEvent)(nil),           // 5: grpcNotification.JobEvent
	(*DeployInfo)(nil),         // 6: grpcNotification.DeployInfo
	(*DeployEvent)(nil),        // 7: grpcNotification.DeployEvent
	(*JobEventRequest)(nil),    // 8: grpcNotification.JobEventRequest
	(*JobResponse)(nil),        // 9: grpcNotification.JobResponse
	(*DeployEventRequest)(nil), // 10: grpcNotification.DeployEventRequest
	(*DeployResponse)(nil),     // 11: grpcNotification.DeployResponse
}
var file_notification_proto_depIdxs = []int32{
	0,  // 0: grpcNotification.JobEvent.type:type_name -> grpcNotification.JobEventType
	1,  // 1: grpcNotification.DeployEvent.type:type_name -> grpcNotification.DeployEventType
	0,  // 2: grpcNotification.JobEventRequest.type:type_name -> grpcNotification.JobEventType
	2,  // 3: grpcNotification.JobResponse.status:type_name -> grpcNotification.JobResponse.Status
	1,  // 4: grpcNotification.DeployEventRequest.type:type_name -> grpcNotification.DeployEventType
	3,  // 5: grpcNotification.DeployResponse.status:type_name -> grpcNotification.DeployResponse.Status
	4,  // 6: grpcNotification.JobEventRoute.GetJobEvent:input_type -> grpcNotification.JobTaskInfo
	6,  // 7: grpcNotification.DeployEventRoute.GetDeployEvent:input_type -> grpcNotification.DeployInfo
	8,  // 8: grpcNotification.JobTriggerRoute.NotifyJob:input_type -> grpcNotification.JobEventRequest
	10, // 9: grpcNotification.DeployTriggerRoute.NotifyDeploy:input_type -> grpcNotification.DeployEventRequest
	5,  // 10: grpcNotification.JobEventRoute.GetJobEvent:output_type -> grpcNotification.JobEvent
	7,  // 11: grpcNotification.DeployEventRoute.GetDeployEvent:output_type -> grpcNotification.DeployEvent
	9,  // 12: grpcNotification.JobTriggerRoute.NotifyJob:output_type -> grpcNotification.JobResponse
	11, // 13: grpcNotification.DeployTriggerRoute.NotifyDeploy:output_type -> grpcNotification.DeployResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_notification_proto_init() }
func file_notification_proto_init() {
	if File_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobTaskInfo); i {
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
		file_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobEvent); i {
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
		file_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployInfo); i {
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
		file_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployEvent); i {
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
		file_notification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobEventRequest); i {
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
		file_notification_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobResponse); i {
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
		file_notification_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployEventRequest); i {
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
		file_notification_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployResponse); i {
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
			RawDescriptor: file_notification_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_notification_proto_goTypes,
		DependencyIndexes: file_notification_proto_depIdxs,
		EnumInfos:         file_notification_proto_enumTypes,
		MessageInfos:      file_notification_proto_msgTypes,
	}.Build()
	File_notification_proto = out.File
	file_notification_proto_rawDesc = nil
	file_notification_proto_goTypes = nil
	file_notification_proto_depIdxs = nil
}