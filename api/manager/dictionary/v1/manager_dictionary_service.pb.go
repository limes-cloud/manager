// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: api/manager/dictionary/manager_dictionary_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_manager_dictionary_manager_dictionary_service_proto protoreflect.FileDescriptor

var file_api_manager_dictionary_manager_dictionary_service_proto_rawDesc = []byte{
	0x0a, 0x37, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x72, 0x79, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xe1, 0x0e, 0x0a, 0x0a, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x12, 0xa8, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x72, 0x79, 0x12, 0x38, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0xaf, 0x01, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x12, 0x3a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a,
	0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x12, 0xaf,
	0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x72, 0x79, 0x12, 0x3a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x38, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x3a, 0x01, 0x2a, 0x1a, 0x1a, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79,
	0x12, 0xac, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x3a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x38, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x2a, 0x1a, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x12,
	0xbc, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0xc4,
	0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0xc4, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x3f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x1a, 0x20, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0xdd, 0x01, 0x0a,
	0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x45, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x3a, 0x01, 0x2a, 0x1a, 0x27, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0xc1, 0x01, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x2a, 0x20,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0xa3, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x72, 0x79, 0x12, 0x37, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x42, 0x3c, 0x0a, 0x21, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x44, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_manager_dictionary_manager_dictionary_service_proto_goTypes = []interface{}{
	(*ListDictionaryRequest)(nil),              // 0: manager.api.manager.dictionary.v1.ListDictionaryRequest
	(*CreateDictionaryRequest)(nil),            // 1: manager.api.manager.dictionary.v1.CreateDictionaryRequest
	(*UpdateDictionaryRequest)(nil),            // 2: manager.api.manager.dictionary.v1.UpdateDictionaryRequest
	(*DeleteDictionaryRequest)(nil),            // 3: manager.api.manager.dictionary.v1.DeleteDictionaryRequest
	(*ListDictionaryValueRequest)(nil),         // 4: manager.api.manager.dictionary.v1.ListDictionaryValueRequest
	(*CreateDictionaryValueRequest)(nil),       // 5: manager.api.manager.dictionary.v1.CreateDictionaryValueRequest
	(*UpdateDictionaryValueRequest)(nil),       // 6: manager.api.manager.dictionary.v1.UpdateDictionaryValueRequest
	(*UpdateDictionaryValueStatusRequest)(nil), // 7: manager.api.manager.dictionary.v1.UpdateDictionaryValueStatusRequest
	(*DeleteDictionaryValueRequest)(nil),       // 8: manager.api.manager.dictionary.v1.DeleteDictionaryValueRequest
	(*GetDictionaryRequest)(nil),               // 9: manager.api.manager.dictionary.v1.GetDictionaryRequest
	(*ListDictionaryReply)(nil),                // 10: manager.api.manager.dictionary.v1.ListDictionaryReply
	(*CreateDictionaryReply)(nil),              // 11: manager.api.manager.dictionary.v1.CreateDictionaryReply
	(*UpdateDictionaryReply)(nil),              // 12: manager.api.manager.dictionary.v1.UpdateDictionaryReply
	(*DeleteDictionaryReply)(nil),              // 13: manager.api.manager.dictionary.v1.DeleteDictionaryReply
	(*ListDictionaryValueReply)(nil),           // 14: manager.api.manager.dictionary.v1.ListDictionaryValueReply
	(*CreateDictionaryValueReply)(nil),         // 15: manager.api.manager.dictionary.v1.CreateDictionaryValueReply
	(*UpdateDictionaryValueReply)(nil),         // 16: manager.api.manager.dictionary.v1.UpdateDictionaryValueReply
	(*UpdateDictionaryValueStatusReply)(nil),   // 17: manager.api.manager.dictionary.v1.UpdateDictionaryValueStatusReply
	(*DeleteDictionaryValueReply)(nil),         // 18: manager.api.manager.dictionary.v1.DeleteDictionaryValueReply
	(*GetDictionaryReply)(nil),                 // 19: manager.api.manager.dictionary.v1.GetDictionaryReply
}
var file_api_manager_dictionary_manager_dictionary_service_proto_depIdxs = []int32{
	0,  // 0: manager.api.manager.dictionary.v1.Dictionary.ListDictionary:input_type -> manager.api.manager.dictionary.v1.ListDictionaryRequest
	1,  // 1: manager.api.manager.dictionary.v1.Dictionary.CreateDictionary:input_type -> manager.api.manager.dictionary.v1.CreateDictionaryRequest
	2,  // 2: manager.api.manager.dictionary.v1.Dictionary.UpdateDictionary:input_type -> manager.api.manager.dictionary.v1.UpdateDictionaryRequest
	3,  // 3: manager.api.manager.dictionary.v1.Dictionary.DeleteDictionary:input_type -> manager.api.manager.dictionary.v1.DeleteDictionaryRequest
	4,  // 4: manager.api.manager.dictionary.v1.Dictionary.ListDictionaryValue:input_type -> manager.api.manager.dictionary.v1.ListDictionaryValueRequest
	5,  // 5: manager.api.manager.dictionary.v1.Dictionary.CreateDictionaryValue:input_type -> manager.api.manager.dictionary.v1.CreateDictionaryValueRequest
	6,  // 6: manager.api.manager.dictionary.v1.Dictionary.UpdateDictionaryValue:input_type -> manager.api.manager.dictionary.v1.UpdateDictionaryValueRequest
	7,  // 7: manager.api.manager.dictionary.v1.Dictionary.UpdateDictionaryValueStatus:input_type -> manager.api.manager.dictionary.v1.UpdateDictionaryValueStatusRequest
	8,  // 8: manager.api.manager.dictionary.v1.Dictionary.DeleteDictionaryValue:input_type -> manager.api.manager.dictionary.v1.DeleteDictionaryValueRequest
	9,  // 9: manager.api.manager.dictionary.v1.Dictionary.GetDictionary:input_type -> manager.api.manager.dictionary.v1.GetDictionaryRequest
	10, // 10: manager.api.manager.dictionary.v1.Dictionary.ListDictionary:output_type -> manager.api.manager.dictionary.v1.ListDictionaryReply
	11, // 11: manager.api.manager.dictionary.v1.Dictionary.CreateDictionary:output_type -> manager.api.manager.dictionary.v1.CreateDictionaryReply
	12, // 12: manager.api.manager.dictionary.v1.Dictionary.UpdateDictionary:output_type -> manager.api.manager.dictionary.v1.UpdateDictionaryReply
	13, // 13: manager.api.manager.dictionary.v1.Dictionary.DeleteDictionary:output_type -> manager.api.manager.dictionary.v1.DeleteDictionaryReply
	14, // 14: manager.api.manager.dictionary.v1.Dictionary.ListDictionaryValue:output_type -> manager.api.manager.dictionary.v1.ListDictionaryValueReply
	15, // 15: manager.api.manager.dictionary.v1.Dictionary.CreateDictionaryValue:output_type -> manager.api.manager.dictionary.v1.CreateDictionaryValueReply
	16, // 16: manager.api.manager.dictionary.v1.Dictionary.UpdateDictionaryValue:output_type -> manager.api.manager.dictionary.v1.UpdateDictionaryValueReply
	17, // 17: manager.api.manager.dictionary.v1.Dictionary.UpdateDictionaryValueStatus:output_type -> manager.api.manager.dictionary.v1.UpdateDictionaryValueStatusReply
	18, // 18: manager.api.manager.dictionary.v1.Dictionary.DeleteDictionaryValue:output_type -> manager.api.manager.dictionary.v1.DeleteDictionaryValueReply
	19, // 19: manager.api.manager.dictionary.v1.Dictionary.GetDictionary:output_type -> manager.api.manager.dictionary.v1.GetDictionaryReply
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_manager_dictionary_manager_dictionary_service_proto_init() }
func file_api_manager_dictionary_manager_dictionary_service_proto_init() {
	if File_api_manager_dictionary_manager_dictionary_service_proto != nil {
		return
	}
	file_api_manager_dictionary_manager_dictionary_proto_init()
	file_api_manager_dictionary_manager_dictionary_value_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_manager_dictionary_manager_dictionary_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_manager_dictionary_manager_dictionary_service_proto_goTypes,
		DependencyIndexes: file_api_manager_dictionary_manager_dictionary_service_proto_depIdxs,
	}.Build()
	File_api_manager_dictionary_manager_dictionary_service_proto = out.File
	file_api_manager_dictionary_manager_dictionary_service_proto_rawDesc = nil
	file_api_manager_dictionary_manager_dictionary_service_proto_goTypes = nil
	file_api_manager_dictionary_manager_dictionary_service_proto_depIdxs = nil
}
