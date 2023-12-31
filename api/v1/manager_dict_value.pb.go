// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: manager_dict_value.proto

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

type DictValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DictId      uint32  `protobuf:"varint,2,opt,name=dict_id,json=dictId,proto3" json:"dict_id,omitempty"`
	Label       string  `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Value       string  `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Status      bool    `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	Weight      *uint32 `protobuf:"varint,6,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
	Type        string  `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Extra       string  `protobuf:"bytes,8,opt,name=extra,proto3" json:"extra,omitempty"`
	Description string  `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *uint32 `protobuf:"varint,10,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt   *uint32 `protobuf:"varint,11,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
}

func (x *DictValue) Reset() {
	*x = DictValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictValue) ProtoMessage() {}

func (x *DictValue) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictValue.ProtoReflect.Descriptor instead.
func (*DictValue) Descriptor() ([]byte, []int) {
	return file_manager_dict_value_proto_rawDescGZIP(), []int{0}
}

func (x *DictValue) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DictValue) GetDictId() uint32 {
	if x != nil {
		return x.DictId
	}
	return 0
}

func (x *DictValue) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *DictValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *DictValue) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *DictValue) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

func (x *DictValue) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DictValue) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *DictValue) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DictValue) GetCreatedAt() uint32 {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return 0
}

func (x *DictValue) GetUpdatedAt() uint32 {
	if x != nil && x.UpdatedAt != nil {
		return *x.UpdatedAt
	}
	return 0
}

type PageDictValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Label    *string `protobuf:"bytes,3,opt,name=label,proto3,oneof" json:"label,omitempty"`
	Value    *string `protobuf:"bytes,4,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *PageDictValueRequest) Reset() {
	*x = PageDictValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageDictValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageDictValueRequest) ProtoMessage() {}

func (x *PageDictValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageDictValueRequest.ProtoReflect.Descriptor instead.
func (*PageDictValueRequest) Descriptor() ([]byte, []int) {
	return file_manager_dict_value_proto_rawDescGZIP(), []int{1}
}

func (x *PageDictValueRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageDictValueRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageDictValueRequest) GetLabel() string {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return ""
}

func (x *PageDictValueRequest) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type PageDictValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*DictValue `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PageDictValueReply) Reset() {
	*x = PageDictValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_value_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageDictValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageDictValueReply) ProtoMessage() {}

func (x *PageDictValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_value_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageDictValueReply.ProtoReflect.Descriptor instead.
func (*PageDictValueReply) Descriptor() ([]byte, []int) {
	return file_manager_dict_value_proto_rawDescGZIP(), []int{2}
}

func (x *PageDictValueReply) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageDictValueReply) GetList() []*DictValue {
	if x != nil {
		return x.List
	}
	return nil
}

type AddDictValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DictId      uint32  `protobuf:"varint,2,opt,name=dict_id,json=dictId,proto3" json:"dict_id,omitempty"`
	Label       string  `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Value       string  `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Type        *string `protobuf:"bytes,5,opt,name=type,proto3,oneof" json:"type,omitempty"`
	Extra       *string `protobuf:"bytes,6,opt,name=extra,proto3,oneof" json:"extra,omitempty"`
	Description *string `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Weight      *uint32 `protobuf:"varint,8,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
	Status      *bool   `protobuf:"varint,9,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *AddDictValueRequest) Reset() {
	*x = AddDictValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_value_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDictValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDictValueRequest) ProtoMessage() {}

func (x *AddDictValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_value_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDictValueRequest.ProtoReflect.Descriptor instead.
func (*AddDictValueRequest) Descriptor() ([]byte, []int) {
	return file_manager_dict_value_proto_rawDescGZIP(), []int{3}
}

func (x *AddDictValueRequest) GetDictId() uint32 {
	if x != nil {
		return x.DictId
	}
	return 0
}

func (x *AddDictValueRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *AddDictValueRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AddDictValueRequest) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *AddDictValueRequest) GetExtra() string {
	if x != nil && x.Extra != nil {
		return *x.Extra
	}
	return ""
}

func (x *AddDictValueRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *AddDictValueRequest) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

func (x *AddDictValueRequest) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

type UpdateDictValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DictId      uint32  `protobuf:"varint,2,opt,name=dict_id,json=dictId,proto3" json:"dict_id,omitempty"`
	Label       string  `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Value       string  `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Type        *string `protobuf:"bytes,5,opt,name=type,proto3,oneof" json:"type,omitempty"`
	Extra       *string `protobuf:"bytes,6,opt,name=extra,proto3,oneof" json:"extra,omitempty"`
	Description *string `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Weight      *uint32 `protobuf:"varint,8,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
	Status      *bool   `protobuf:"varint,9,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *UpdateDictValueRequest) Reset() {
	*x = UpdateDictValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_value_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDictValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDictValueRequest) ProtoMessage() {}

func (x *UpdateDictValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_value_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDictValueRequest.ProtoReflect.Descriptor instead.
func (*UpdateDictValueRequest) Descriptor() ([]byte, []int) {
	return file_manager_dict_value_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDictValueRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDictValueRequest) GetDictId() uint32 {
	if x != nil {
		return x.DictId
	}
	return 0
}

func (x *UpdateDictValueRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *UpdateDictValueRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *UpdateDictValueRequest) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *UpdateDictValueRequest) GetExtra() string {
	if x != nil && x.Extra != nil {
		return *x.Extra
	}
	return ""
}

func (x *UpdateDictValueRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateDictValueRequest) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

func (x *UpdateDictValueRequest) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

type DeleteDictValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDictValueRequest) Reset() {
	*x = DeleteDictValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_value_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDictValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDictValueRequest) ProtoMessage() {}

func (x *DeleteDictValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_value_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDictValueRequest.ProtoReflect.Descriptor instead.
func (*DeleteDictValueRequest) Descriptor() ([]byte, []int) {
	return file_manager_dict_value_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDictValueRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_manager_dict_value_proto protoreflect.FileDescriptor

var file_manager_dict_value_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x02, 0x0a,
	0x09, 0x44, 0x69, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x69,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x64, 0x69, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x22, 0xa5, 0x01, 0x0a, 0x14, 0x50, 0x61, 0x67, 0x65, 0x44, 0x69, 0x63, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20,
	0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a,
	0x04, 0x18, 0x32, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x19, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x52, 0x0a, 0x12, 0x50, 0x61, 0x67,
	0x65, 0x44, 0x69, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x69,
	0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xc3, 0x02,
	0x0a, 0x13, 0x41, 0x64, 0x64, 0x44, 0x69, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52,
	0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0xdf, 0x02, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x64, 0x69, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20,
	0x00, 0x52, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x31, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x69, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manager_dict_value_proto_rawDescOnce sync.Once
	file_manager_dict_value_proto_rawDescData = file_manager_dict_value_proto_rawDesc
)

func file_manager_dict_value_proto_rawDescGZIP() []byte {
	file_manager_dict_value_proto_rawDescOnce.Do(func() {
		file_manager_dict_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_dict_value_proto_rawDescData)
	})
	return file_manager_dict_value_proto_rawDescData
}

var file_manager_dict_value_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_manager_dict_value_proto_goTypes = []interface{}{
	(*DictValue)(nil),              // 0: manager.DictValue
	(*PageDictValueRequest)(nil),   // 1: manager.PageDictValueRequest
	(*PageDictValueReply)(nil),     // 2: manager.PageDictValueReply
	(*AddDictValueRequest)(nil),    // 3: manager.AddDictValueRequest
	(*UpdateDictValueRequest)(nil), // 4: manager.UpdateDictValueRequest
	(*DeleteDictValueRequest)(nil), // 5: manager.DeleteDictValueRequest
}
var file_manager_dict_value_proto_depIdxs = []int32{
	0, // 0: manager.PageDictValueReply.list:type_name -> manager.DictValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_manager_dict_value_proto_init() }
func file_manager_dict_value_proto_init() {
	if File_manager_dict_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_dict_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictValue); i {
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
		file_manager_dict_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageDictValueRequest); i {
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
		file_manager_dict_value_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageDictValueReply); i {
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
		file_manager_dict_value_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDictValueRequest); i {
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
		file_manager_dict_value_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDictValueRequest); i {
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
		file_manager_dict_value_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDictValueRequest); i {
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
	file_manager_dict_value_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_manager_dict_value_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_manager_dict_value_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_manager_dict_value_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manager_dict_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manager_dict_value_proto_goTypes,
		DependencyIndexes: file_manager_dict_value_proto_depIdxs,
		MessageInfos:      file_manager_dict_value_proto_msgTypes,
	}.Build()
	File_manager_dict_value_proto = out.File
	file_manager_dict_value_proto_rawDesc = nil
	file_manager_dict_value_proto_goTypes = nil
	file_manager_dict_value_proto_depIdxs = nil
}
