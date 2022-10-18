// Copyright 2022 Tigris Data, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: server/v1/health.proto

package api

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HealthCheckInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckInput) Reset() {
	*x = HealthCheckInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckInput) ProtoMessage() {}

func (x *HealthCheckInput) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckInput.ProtoReflect.Descriptor instead.
func (*HealthCheckInput) Descriptor() ([]byte, []int) {
	return file_server_v1_health_proto_rawDescGZIP(), []int{0}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_server_v1_health_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_server_v1_health_proto protoreflect.FileDescriptor

var file_server_v1_health_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x31, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x72, 0x0a, 0x09, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x41, 0x50, 0x49, 0x12, 0x65, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x11, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x14, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0xba, 0x47,
	0x1d, 0x0a, 0x0d, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x41,
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x67, 0x72, 0x69, 0x73, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x67, 0x72,
	0x69, 0x73, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x69, 0x67, 0x72, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_v1_health_proto_rawDescOnce sync.Once
	file_server_v1_health_proto_rawDescData = file_server_v1_health_proto_rawDesc
)

func file_server_v1_health_proto_rawDescGZIP() []byte {
	file_server_v1_health_proto_rawDescOnce.Do(func() {
		file_server_v1_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_v1_health_proto_rawDescData)
	})
	return file_server_v1_health_proto_rawDescData
}

var file_server_v1_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_v1_health_proto_goTypes = []interface{}{
	(*HealthCheckInput)(nil),    // 0: HealthCheckInput
	(*HealthCheckResponse)(nil), // 1: HealthCheckResponse
}
var file_server_v1_health_proto_depIdxs = []int32{
	0, // 0: HealthAPI.Health:input_type -> HealthCheckInput
	1, // 1: HealthAPI.Health:output_type -> HealthCheckResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_v1_health_proto_init() }
func file_server_v1_health_proto_init() {
	if File_server_v1_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_v1_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckInput); i {
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
		file_server_v1_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
			RawDescriptor: file_server_v1_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_v1_health_proto_goTypes,
		DependencyIndexes: file_server_v1_health_proto_depIdxs,
		MessageInfos:      file_server_v1_health_proto_msgTypes,
	}.Build()
	File_server_v1_health_proto = out.File
	file_server_v1_health_proto_rawDesc = nil
	file_server_v1_health_proto_goTypes = nil
	file_server_v1_health_proto_depIdxs = nil
}
