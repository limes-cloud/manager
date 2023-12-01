// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: manager_setting.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type GetSettingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Debug     bool   `protobuf:"varint,2,opt,name=debug,proto3" json:"debug,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Desc      string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Copyright string `protobuf:"bytes,5,opt,name=copyright,proto3" json:"copyright,omitempty"`
	Logo      string `protobuf:"bytes,6,opt,name=logo,proto3" json:"logo,omitempty"`
}

func (x *GetSettingReply) Reset() {
	*x = GetSettingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingReply) ProtoMessage() {}

func (x *GetSettingReply) ProtoReflect() protoreflect.Message {
	mi := &file_manager_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingReply.ProtoReflect.Descriptor instead.
func (*GetSettingReply) Descriptor() ([]byte, []int) {
	return file_manager_setting_proto_rawDescGZIP(), []int{0}
}

func (x *GetSettingReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSettingReply) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *GetSettingReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetSettingReply) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *GetSettingReply) GetCopyright() string {
	if x != nil {
		return x.Copyright
	}
	return ""
}

func (x *GetSettingReply) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

var File_manager_setting_proto protoreflect.FileDescriptor

var file_manager_setting_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x6f, 0x67, 0x6f, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manager_setting_proto_rawDescOnce sync.Once
	file_manager_setting_proto_rawDescData = file_manager_setting_proto_rawDesc
)

func file_manager_setting_proto_rawDescGZIP() []byte {
	file_manager_setting_proto_rawDescOnce.Do(func() {
		file_manager_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_setting_proto_rawDescData)
	})
	return file_manager_setting_proto_rawDescData
}

var file_manager_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_manager_setting_proto_goTypes = []interface{}{
	(*GetSettingReply)(nil), // 0: manager.GetSettingReply
}
var file_manager_setting_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_manager_setting_proto_init() }
func file_manager_setting_proto_init() {
	if File_manager_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingReply); i {
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
			RawDescriptor: file_manager_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manager_setting_proto_goTypes,
		DependencyIndexes: file_manager_setting_proto_depIdxs,
		MessageInfos:      file_manager_setting_proto_msgTypes,
	}.Build()
	File_manager_setting_proto = out.File
	file_manager_setting_proto_rawDesc = nil
	file_manager_setting_proto_goTypes = nil
	file_manager_setting_proto_depIdxs = nil
}
