// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: manager_job_service.proto

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

var File_manager_job_service_proto protoreflect.FileDescriptor

var file_manager_job_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xff, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5b, 0x0a, 0x07, 0x50, 0x61, 0x67, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x1b, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x12,
	0x5a, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4a, 0x6f, 0x62, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x6a, 0x6f, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x12, 0x5e, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x12, 0x5b, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_manager_job_service_proto_goTypes = []interface{}{
	(*PageJobRequest)(nil),   // 0: manager_job.PageJobRequest
	(*AddJobRequest)(nil),    // 1: manager_job.AddJobRequest
	(*UpdateJobRequest)(nil), // 2: manager_job.UpdateJobRequest
	(*DeleteJobRequest)(nil), // 3: manager_job.DeleteJobRequest
	(*PageJobReply)(nil),     // 4: manager_job.PageJobReply
	(*AddJobReply)(nil),      // 5: manager_job.AddJobReply
	(*emptypb.Empty)(nil),    // 6: google.protobuf.Empty
}
var file_manager_job_service_proto_depIdxs = []int32{
	0, // 0: manager_job.Service.PageJob:input_type -> manager_job.PageJobRequest
	1, // 1: manager_job.Service.AddJob:input_type -> manager_job.AddJobRequest
	2, // 2: manager_job.Service.UpdateJob:input_type -> manager_job.UpdateJobRequest
	3, // 3: manager_job.Service.DeleteJob:input_type -> manager_job.DeleteJobRequest
	4, // 4: manager_job.Service.PageJob:output_type -> manager_job.PageJobReply
	5, // 5: manager_job.Service.AddJob:output_type -> manager_job.AddJobReply
	6, // 6: manager_job.Service.UpdateJob:output_type -> google.protobuf.Empty
	6, // 7: manager_job.Service.DeleteJob:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_manager_job_service_proto_init() }
func file_manager_job_service_proto_init() {
	if File_manager_job_service_proto != nil {
		return
	}
	file_manager_job_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manager_job_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manager_job_service_proto_goTypes,
		DependencyIndexes: file_manager_job_service_proto_depIdxs,
	}.Build()
	File_manager_job_service_proto = out.File
	file_manager_job_service_proto_rawDesc = nil
	file_manager_job_service_proto_goTypes = nil
	file_manager_job_service_proto_depIdxs = nil
}
