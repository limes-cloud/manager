// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: basic_platform_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_basic_platform_service_proto protoreflect.FileDescriptor

var file_basic_platform_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa1, 0x20, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1e, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x7f, 0x0a,
	0x13, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x54, 0x72, 0x65, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x72, 0x65,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x62, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x12, 0x83,
	0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x64,
	0x73, 0x12, 0x25, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x62, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x6d, 0x65, 0x6e, 0x75,
	0x2f, 0x69, 0x64, 0x73, 0x12, 0x70, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x25, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01,
	0x2a, 0x1a, 0x13, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x12, 0x6a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x54, 0x72, 0x65, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x62, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x13, 0x2f,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x74, 0x72,
	0x65, 0x65, 0x12, 0x5c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1e, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x41,
	0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a,
	0x22, 0x0e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x62, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x21,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x3a, 0x01, 0x2a, 0x1a, 0x0e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75,
	0x54, 0x72, 0x65, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x62, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2f, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x74, 0x72, 0x65,
	0x65, 0x12, 0x5c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1e, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x41, 0x64,
	0x64, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22,
	0x0e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x12,
	0x62, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x21, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x3a, 0x01, 0x2a, 0x1a, 0x0e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x6e, 0x75, 0x12, 0x5f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e,
	0x75, 0x12, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x6e, 0x75, 0x12, 0x5b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x82, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x26, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72,
	0x65, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x62,
	0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x12, 0x6e, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x44, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a,
	0x22, 0x14, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x74, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x1a, 0x14, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x71, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x27, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x49, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x0b, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x62, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2f,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x63, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1f, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x5c, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x8c, 0x01, 0x0a, 0x19, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2e,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2f,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x7f, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x7f, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1e, 0x2f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x6c, 0x0a, 0x0b, 0x44, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22,
	0x16, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x69, 0x0a, 0x0a, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x6c, 0x0a, 0x0b, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x22, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x62, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x3a, 0x01, 0x2a, 0x1a, 0x0e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x5f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x7e, 0x0a, 0x10, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x62, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x98, 0x01, 0x0a, 0x15, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x2c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x53,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x12, 0x7f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0x23, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21,
	0x62, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x2f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x12, 0x56, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x12, 0x5d, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x6a, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x17, 0x2f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x63, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x12, 0x52, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x10, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x64, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x15, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x12, 0x6a,
	0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x22, 0x17, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_basic_platform_service_proto_goTypes = []interface{}{
	(*GetRoleRequest)(nil),                 // 0: basic_platform.GetRoleRequest
	(*emptypb.Empty)(nil),                  // 1: google.protobuf.Empty
	(*GetRoleMenuIdsRequest)(nil),          // 2: basic_platform.GetRoleMenuIdsRequest
	(*UpdateRoleMenuRequest)(nil),          // 3: basic_platform.UpdateRoleMenuRequest
	(*AddRoleRequest)(nil),                 // 4: basic_platform.AddRoleRequest
	(*UpdateRoleRequest)(nil),              // 5: basic_platform.UpdateRoleRequest
	(*DeleteRoleRequest)(nil),              // 6: basic_platform.DeleteRoleRequest
	(*AddMenuRequest)(nil),                 // 7: basic_platform.AddMenuRequest
	(*UpdateMenuRequest)(nil),              // 8: basic_platform.UpdateMenuRequest
	(*DeleteMenuRequest)(nil),              // 9: basic_platform.DeleteMenuRequest
	(*GetDepartmentRequest)(nil),           // 10: basic_platform.GetDepartmentRequest
	(*AddDepartmentRequest)(nil),           // 11: basic_platform.AddDepartmentRequest
	(*UpdateDepartmentRequest)(nil),        // 12: basic_platform.UpdateDepartmentRequest
	(*DeleteDepartmentRequest)(nil),        // 13: basic_platform.DeleteDepartmentRequest
	(*GetUserRequest)(nil),                 // 14: basic_platform.GetUserRequest
	(*PageUserRequest)(nil),                // 15: basic_platform.PageUserRequest
	(*AddUserRequest)(nil),                 // 16: basic_platform.AddUserRequest
	(*ResetUserPasswordRequest)(nil),       // 17: basic_platform.ResetUserPasswordRequest
	(*ChangeUserPasswordRequest)(nil),      // 18: basic_platform.ChangeUserPasswordRequest
	(*DisableUserRequest)(nil),             // 19: basic_platform.DisableUserRequest
	(*EnableUserRequest)(nil),              // 20: basic_platform.EnableUserRequest
	(*OfflineUserRequest)(nil),             // 21: basic_platform.OfflineUserRequest
	(*UpdateUserRequest)(nil),              // 22: basic_platform.UpdateUserRequest
	(*DeleteUserRequest)(nil),              // 23: basic_platform.DeleteUserRequest
	(*SwitchCurrentUserRoleRequest)(nil),   // 24: basic_platform.SwitchCurrentUserRoleRequest
	(*GetUserRolesRequest)(nil),            // 25: basic_platform.GetUserRolesRequest
	(*AuthRequest)(nil),                    // 26: basic_platform.AuthRequest
	(*LoginRequest)(nil),                   // 27: basic_platform.LoginRequest
	(*GetRoleReply)(nil),                   // 28: basic_platform.GetRoleReply
	(*CurrentRoleMenuTreeReply)(nil),       // 29: basic_platform.CurrentRoleMenuTreeReply
	(*GetRoleMenuIdsReply)(nil),            // 30: basic_platform.GetRoleMenuIdsReply
	(*GetRoleTreeReply)(nil),               // 31: basic_platform.GetRoleTreeReply
	(*GetMenuTreeReply)(nil),               // 32: basic_platform.GetMenuTreeReply
	(*GetDepartmentReply)(nil),             // 33: basic_platform.GetDepartmentReply
	(*GetDepartmentTreeReply)(nil),         // 34: basic_platform.GetDepartmentTreeReply
	(*GetUserReply)(nil),                   // 35: basic_platform.GetUserReply
	(*PageUserReply)(nil),                  // 36: basic_platform.PageUserReply
	(*ChangeUserPasswordCaptchaReply)(nil), // 37: basic_platform.ChangeUserPasswordCaptchaReply
	(*GetUserRolesReply)(nil),              // 38: basic_platform.GetUserRolesReply
	(*SwitchCurrentUserRoleReply)(nil),     // 39: basic_platform.SwitchCurrentUserRoleReply
	(*LoginReply)(nil),                     // 40: basic_platform.LoginReply
	(*LoginCaptchaReply)(nil),              // 41: basic_platform.LoginCaptchaReply
	(*ParseTokenReply)(nil),                // 42: basic_platform.ParseTokenReply
	(*RefreshTokenReply)(nil),              // 43: basic_platform.RefreshTokenReply
}
var file_basic_platform_service_proto_depIdxs = []int32{
	0,  // 0: basic_platform.Service.GetRole:input_type -> basic_platform.GetRoleRequest
	1,  // 1: basic_platform.Service.CurrentRoleMenuTree:input_type -> google.protobuf.Empty
	2,  // 2: basic_platform.Service.GetRoleMenuIds:input_type -> basic_platform.GetRoleMenuIdsRequest
	3,  // 3: basic_platform.Service.UpdateRoleMenus:input_type -> basic_platform.UpdateRoleMenuRequest
	1,  // 4: basic_platform.Service.GetRoleTree:input_type -> google.protobuf.Empty
	4,  // 5: basic_platform.Service.AddRole:input_type -> basic_platform.AddRoleRequest
	5,  // 6: basic_platform.Service.UpdateRole:input_type -> basic_platform.UpdateRoleRequest
	6,  // 7: basic_platform.Service.DeleteRole:input_type -> basic_platform.DeleteRoleRequest
	1,  // 8: basic_platform.Service.GetMenuTree:input_type -> google.protobuf.Empty
	7,  // 9: basic_platform.Service.AddMenu:input_type -> basic_platform.AddMenuRequest
	8,  // 10: basic_platform.Service.UpdateMenu:input_type -> basic_platform.UpdateMenuRequest
	9,  // 11: basic_platform.Service.DeleteMenu:input_type -> basic_platform.DeleteMenuRequest
	10, // 12: basic_platform.Service.GetDepartment:input_type -> basic_platform.GetDepartmentRequest
	1,  // 13: basic_platform.Service.GetDepartmentTree:input_type -> google.protobuf.Empty
	11, // 14: basic_platform.Service.AddDepartment:input_type -> basic_platform.AddDepartmentRequest
	12, // 15: basic_platform.Service.UpdateDepartment:input_type -> basic_platform.UpdateDepartmentRequest
	13, // 16: basic_platform.Service.DeleteDepartment:input_type -> basic_platform.DeleteDepartmentRequest
	14, // 17: basic_platform.Service.GetUser:input_type -> basic_platform.GetUserRequest
	1,  // 18: basic_platform.Service.CurrentUser:input_type -> google.protobuf.Empty
	15, // 19: basic_platform.Service.PageUser:input_type -> basic_platform.PageUserRequest
	16, // 20: basic_platform.Service.AddUser:input_type -> basic_platform.AddUserRequest
	1,  // 21: basic_platform.Service.ChangeUserPasswordCaptcha:input_type -> google.protobuf.Empty
	17, // 22: basic_platform.Service.ResetUserPassword:input_type -> basic_platform.ResetUserPasswordRequest
	18, // 23: basic_platform.Service.ChangeUserPassword:input_type -> basic_platform.ChangeUserPasswordRequest
	19, // 24: basic_platform.Service.DisableUser:input_type -> basic_platform.DisableUserRequest
	20, // 25: basic_platform.Service.EnableUser:input_type -> basic_platform.EnableUserRequest
	21, // 26: basic_platform.Service.OfflineUser:input_type -> basic_platform.OfflineUserRequest
	22, // 27: basic_platform.Service.UpdateUser:input_type -> basic_platform.UpdateUserRequest
	23, // 28: basic_platform.Service.DeleteUser:input_type -> basic_platform.DeleteUserRequest
	1,  // 29: basic_platform.Service.CurrentUserRoles:input_type -> google.protobuf.Empty
	24, // 30: basic_platform.Service.SwitchCurrentUserRole:input_type -> basic_platform.SwitchCurrentUserRoleRequest
	25, // 31: basic_platform.Service.GetUserRoles:input_type -> basic_platform.GetUserRolesRequest
	26, // 32: basic_platform.Service.Auth:input_type -> basic_platform.AuthRequest
	27, // 33: basic_platform.Service.Login:input_type -> basic_platform.LoginRequest
	1,  // 34: basic_platform.Service.LoginCaptcha:input_type -> google.protobuf.Empty
	1,  // 35: basic_platform.Service.Logout:input_type -> google.protobuf.Empty
	1,  // 36: basic_platform.Service.ParseToken:input_type -> google.protobuf.Empty
	1,  // 37: basic_platform.Service.RefreshToken:input_type -> google.protobuf.Empty
	28, // 38: basic_platform.Service.GetRole:output_type -> basic_platform.GetRoleReply
	29, // 39: basic_platform.Service.CurrentRoleMenuTree:output_type -> basic_platform.CurrentRoleMenuTreeReply
	30, // 40: basic_platform.Service.GetRoleMenuIds:output_type -> basic_platform.GetRoleMenuIdsReply
	1,  // 41: basic_platform.Service.UpdateRoleMenus:output_type -> google.protobuf.Empty
	31, // 42: basic_platform.Service.GetRoleTree:output_type -> basic_platform.GetRoleTreeReply
	1,  // 43: basic_platform.Service.AddRole:output_type -> google.protobuf.Empty
	1,  // 44: basic_platform.Service.UpdateRole:output_type -> google.protobuf.Empty
	1,  // 45: basic_platform.Service.DeleteRole:output_type -> google.protobuf.Empty
	32, // 46: basic_platform.Service.GetMenuTree:output_type -> basic_platform.GetMenuTreeReply
	1,  // 47: basic_platform.Service.AddMenu:output_type -> google.protobuf.Empty
	1,  // 48: basic_platform.Service.UpdateMenu:output_type -> google.protobuf.Empty
	1,  // 49: basic_platform.Service.DeleteMenu:output_type -> google.protobuf.Empty
	33, // 50: basic_platform.Service.GetDepartment:output_type -> basic_platform.GetDepartmentReply
	34, // 51: basic_platform.Service.GetDepartmentTree:output_type -> basic_platform.GetDepartmentTreeReply
	1,  // 52: basic_platform.Service.AddDepartment:output_type -> google.protobuf.Empty
	1,  // 53: basic_platform.Service.UpdateDepartment:output_type -> google.protobuf.Empty
	1,  // 54: basic_platform.Service.DeleteDepartment:output_type -> google.protobuf.Empty
	35, // 55: basic_platform.Service.GetUser:output_type -> basic_platform.GetUserReply
	35, // 56: basic_platform.Service.CurrentUser:output_type -> basic_platform.GetUserReply
	36, // 57: basic_platform.Service.PageUser:output_type -> basic_platform.PageUserReply
	1,  // 58: basic_platform.Service.AddUser:output_type -> google.protobuf.Empty
	37, // 59: basic_platform.Service.ChangeUserPasswordCaptcha:output_type -> basic_platform.ChangeUserPasswordCaptchaReply
	1,  // 60: basic_platform.Service.ResetUserPassword:output_type -> google.protobuf.Empty
	1,  // 61: basic_platform.Service.ChangeUserPassword:output_type -> google.protobuf.Empty
	1,  // 62: basic_platform.Service.DisableUser:output_type -> google.protobuf.Empty
	1,  // 63: basic_platform.Service.EnableUser:output_type -> google.protobuf.Empty
	1,  // 64: basic_platform.Service.OfflineUser:output_type -> google.protobuf.Empty
	1,  // 65: basic_platform.Service.UpdateUser:output_type -> google.protobuf.Empty
	1,  // 66: basic_platform.Service.DeleteUser:output_type -> google.protobuf.Empty
	38, // 67: basic_platform.Service.CurrentUserRoles:output_type -> basic_platform.GetUserRolesReply
	39, // 68: basic_platform.Service.SwitchCurrentUserRole:output_type -> basic_platform.SwitchCurrentUserRoleReply
	38, // 69: basic_platform.Service.GetUserRoles:output_type -> basic_platform.GetUserRolesReply
	1,  // 70: basic_platform.Service.Auth:output_type -> google.protobuf.Empty
	40, // 71: basic_platform.Service.Login:output_type -> basic_platform.LoginReply
	41, // 72: basic_platform.Service.LoginCaptcha:output_type -> basic_platform.LoginCaptchaReply
	1,  // 73: basic_platform.Service.Logout:output_type -> google.protobuf.Empty
	42, // 74: basic_platform.Service.ParseToken:output_type -> basic_platform.ParseTokenReply
	43, // 75: basic_platform.Service.RefreshToken:output_type -> basic_platform.RefreshTokenReply
	38, // [38:76] is the sub-list for method output_type
	0,  // [0:38] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_basic_platform_service_proto_init() }
func file_basic_platform_service_proto_init() {
	if File_basic_platform_service_proto != nil {
		return
	}
	file_basic_platform_role_proto_init()
	file_basic_platform_menu_proto_init()
	file_basic_platform_department_proto_init()
	file_basic_platform_user_proto_init()
	file_basic_platform_user_role_proto_init()
	file_basic_platform_role_menu_proto_init()
	file_basic_platform_auth_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basic_platform_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basic_platform_service_proto_goTypes,
		DependencyIndexes: file_basic_platform_service_proto_depIdxs,
	}.Build()
	File_basic_platform_service_proto = out.File
	file_basic_platform_service_proto_rawDesc = nil
	file_basic_platform_service_proto_goTypes = nil
	file_basic_platform_service_proto_depIdxs = nil
}
