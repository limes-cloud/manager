// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: manager_user_service.proto

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

var File_manager_user_service_proto protoreflect.FileDescriptor

var file_manager_user_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x82, 0x0f, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x61, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x12, 0x11, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x60, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x7a, 0x0a, 0x19, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23,
	0x22, 0x21, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2f, 0x63, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x12, 0x7f, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24,
	0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2f, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x12, 0x82, 0x01, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x6c, 0x0a, 0x0b, 0x44, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x69, 0x0a, 0x0a, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x6c, 0x0a, 0x0b, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x20, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x62, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a,
	0x01, 0x2a, 0x1a, 0x10, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x78, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x3a, 0x01, 0x2a, 0x1a, 0x18, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x5f,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x96, 0x01, 0x0a, 0x15, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x6c, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x6e, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x5d, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x17, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x77, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x23, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1e,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_manager_user_service_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),                // 0: google.protobuf.Empty
	(*PageUserRequest)(nil),              // 1: manager_user.PageUserRequest
	(*AddUserRequest)(nil),               // 2: manager_user.AddUserRequest
	(*ResetUserPasswordRequest)(nil),     // 3: manager_user.ResetUserPasswordRequest
	(*ChangeUserPasswordRequest)(nil),    // 4: manager_user.ChangeUserPasswordRequest
	(*DisableUserRequest)(nil),           // 5: manager_user.DisableUserRequest
	(*EnableUserRequest)(nil),            // 6: manager_user.EnableUserRequest
	(*OfflineUserRequest)(nil),           // 7: manager_user.OfflineUserRequest
	(*UpdateUserRequest)(nil),            // 8: manager_user.UpdateUserRequest
	(*UpdateCurrentUserRequest)(nil),     // 9: manager_user.UpdateCurrentUserRequest
	(*DeleteUserRequest)(nil),            // 10: manager_user.DeleteUserRequest
	(*SwitchCurrentUserRoleRequest)(nil), // 11: manager_user.SwitchCurrentUserRoleRequest
	(*UserLoginRequest)(nil),             // 12: manager_user.UserLoginRequest
	(*User)(nil),                         // 13: manager_user.User
	(*PageUserReply)(nil),                // 14: manager_user.PageUserReply
	(*AddUserReply)(nil),                 // 15: manager_user.AddUserReply
	(*CaptchaReply)(nil),                 // 16: manager_user.CaptchaReply
	(*SwitchCurrentUserRoleReply)(nil),   // 17: manager_user.SwitchCurrentUserRoleReply
	(*UserLoginReply)(nil),               // 18: manager_user.UserLoginReply
	(*UserRefreshTokenReply)(nil),        // 19: manager_user.UserRefreshTokenReply
}
var file_manager_user_service_proto_depIdxs = []int32{
	0,  // 0: manager_user.Service.CurrentUser:input_type -> google.protobuf.Empty
	1,  // 1: manager_user.Service.PageUser:input_type -> manager_user.PageUserRequest
	2,  // 2: manager_user.Service.AddUser:input_type -> manager_user.AddUserRequest
	0,  // 3: manager_user.Service.ChangeUserPasswordCaptcha:input_type -> google.protobuf.Empty
	3,  // 4: manager_user.Service.ResetUserPassword:input_type -> manager_user.ResetUserPasswordRequest
	4,  // 5: manager_user.Service.ChangeUserPassword:input_type -> manager_user.ChangeUserPasswordRequest
	5,  // 6: manager_user.Service.DisableUser:input_type -> manager_user.DisableUserRequest
	6,  // 7: manager_user.Service.EnableUser:input_type -> manager_user.EnableUserRequest
	7,  // 8: manager_user.Service.OfflineUser:input_type -> manager_user.OfflineUserRequest
	8,  // 9: manager_user.Service.UpdateUser:input_type -> manager_user.UpdateUserRequest
	9,  // 10: manager_user.Service.UpdateCurrentUser:input_type -> manager_user.UpdateCurrentUserRequest
	10, // 11: manager_user.Service.DeleteUser:input_type -> manager_user.DeleteUserRequest
	11, // 12: manager_user.Service.SwitchCurrentUserRole:input_type -> manager_user.SwitchCurrentUserRoleRequest
	12, // 13: manager_user.Service.UserLogin:input_type -> manager_user.UserLoginRequest
	0,  // 14: manager_user.Service.UserLoginCaptcha:input_type -> google.protobuf.Empty
	0,  // 15: manager_user.Service.UserLogout:input_type -> google.protobuf.Empty
	0,  // 16: manager_user.Service.UserRefreshToken:input_type -> google.protobuf.Empty
	13, // 17: manager_user.Service.CurrentUser:output_type -> manager_user.User
	14, // 18: manager_user.Service.PageUser:output_type -> manager_user.PageUserReply
	15, // 19: manager_user.Service.AddUser:output_type -> manager_user.AddUserReply
	16, // 20: manager_user.Service.ChangeUserPasswordCaptcha:output_type -> manager_user.CaptchaReply
	0,  // 21: manager_user.Service.ResetUserPassword:output_type -> google.protobuf.Empty
	0,  // 22: manager_user.Service.ChangeUserPassword:output_type -> google.protobuf.Empty
	0,  // 23: manager_user.Service.DisableUser:output_type -> google.protobuf.Empty
	0,  // 24: manager_user.Service.EnableUser:output_type -> google.protobuf.Empty
	0,  // 25: manager_user.Service.OfflineUser:output_type -> google.protobuf.Empty
	0,  // 26: manager_user.Service.UpdateUser:output_type -> google.protobuf.Empty
	0,  // 27: manager_user.Service.UpdateCurrentUser:output_type -> google.protobuf.Empty
	0,  // 28: manager_user.Service.DeleteUser:output_type -> google.protobuf.Empty
	17, // 29: manager_user.Service.SwitchCurrentUserRole:output_type -> manager_user.SwitchCurrentUserRoleReply
	18, // 30: manager_user.Service.UserLogin:output_type -> manager_user.UserLoginReply
	16, // 31: manager_user.Service.UserLoginCaptcha:output_type -> manager_user.CaptchaReply
	0,  // 32: manager_user.Service.UserLogout:output_type -> google.protobuf.Empty
	19, // 33: manager_user.Service.UserRefreshToken:output_type -> manager_user.UserRefreshTokenReply
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_manager_user_service_proto_init() }
func file_manager_user_service_proto_init() {
	if File_manager_user_service_proto != nil {
		return
	}
	file_manager_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manager_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manager_user_service_proto_goTypes,
		DependencyIndexes: file_manager_user_service_proto_depIdxs,
	}.Build()
	File_manager_user_service_proto = out.File
	file_manager_user_service_proto_rawDesc = nil
	file_manager_user_service_proto_goTypes = nil
	file_manager_user_service_proto_depIdxs = nil
}