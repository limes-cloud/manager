// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: manager_dict.proto

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

type Dict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Keyword     string  `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type        string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Extra       string  `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
	Description string  `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Weight      *uint32 `protobuf:"varint,7,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
	CreatedAt   *uint32 `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt   *uint32 `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
}

func (x *Dict) Reset() {
	*x = Dict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dict) ProtoMessage() {}

func (x *Dict) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dict.ProtoReflect.Descriptor instead.
func (*Dict) Descriptor() ([]byte, []int) {
	return file_manager_dict_proto_rawDescGZIP(), []int{0}
}

func (x *Dict) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dict) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *Dict) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dict) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Dict) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *Dict) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Dict) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

func (x *Dict) GetCreatedAt() uint32 {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return 0
}

func (x *Dict) GetUpdatedAt() uint32 {
	if x != nil && x.UpdatedAt != nil {
		return *x.UpdatedAt
	}
	return 0
}

type GetDictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDictRequest) Reset() {
	*x = GetDictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDictRequest) ProtoMessage() {}

func (x *GetDictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDictRequest.ProtoReflect.Descriptor instead.
func (*GetDictRequest) Descriptor() ([]byte, []int) {
	return file_manager_dict_proto_rawDescGZIP(), []int{1}
}

func (x *GetDictRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PageDictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keyword  *string `protobuf:"bytes,3,opt,name=keyword,proto3,oneof" json:"keyword,omitempty"`
	Name     *string `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *PageDictRequest) Reset() {
	*x = PageDictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageDictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageDictRequest) ProtoMessage() {}

func (x *PageDictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageDictRequest.ProtoReflect.Descriptor instead.
func (*PageDictRequest) Descriptor() ([]byte, []int) {
	return file_manager_dict_proto_rawDescGZIP(), []int{2}
}

func (x *PageDictRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageDictRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageDictRequest) GetKeyword() string {
	if x != nil && x.Keyword != nil {
		return *x.Keyword
	}
	return ""
}

func (x *PageDictRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type PageDictReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*Dict `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PageDictReply) Reset() {
	*x = PageDictReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageDictReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageDictReply) ProtoMessage() {}

func (x *PageDictReply) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageDictReply.ProtoReflect.Descriptor instead.
func (*PageDictReply) Descriptor() ([]byte, []int) {
	return file_manager_dict_proto_rawDescGZIP(), []int{3}
}

func (x *PageDictReply) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageDictReply) GetList() []*Dict {
	if x != nil {
		return x.List
	}
	return nil
}

type AddDictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword     string  `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type        string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Extra       *string `protobuf:"bytes,4,opt,name=extra,proto3,oneof" json:"extra,omitempty"`
	Description string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Weight      *uint32 `protobuf:"varint,6,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
}

func (x *AddDictRequest) Reset() {
	*x = AddDictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDictRequest) ProtoMessage() {}

func (x *AddDictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDictRequest.ProtoReflect.Descriptor instead.
func (*AddDictRequest) Descriptor() ([]byte, []int) {
	return file_manager_dict_proto_rawDescGZIP(), []int{4}
}

func (x *AddDictRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *AddDictRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddDictRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AddDictRequest) GetExtra() string {
	if x != nil && x.Extra != nil {
		return *x.Extra
	}
	return ""
}

func (x *AddDictRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddDictRequest) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

type UpdateDictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type        string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Extra       *string `protobuf:"bytes,4,opt,name=extra,proto3,oneof" json:"extra,omitempty"`
	Description string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Weight      *uint32 `protobuf:"varint,6,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
}

func (x *UpdateDictRequest) Reset() {
	*x = UpdateDictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDictRequest) ProtoMessage() {}

func (x *UpdateDictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDictRequest.ProtoReflect.Descriptor instead.
func (*UpdateDictRequest) Descriptor() ([]byte, []int) {
	return file_manager_dict_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDictRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDictRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDictRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateDictRequest) GetExtra() string {
	if x != nil && x.Extra != nil {
		return *x.Extra
	}
	return ""
}

func (x *UpdateDictRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateDictRequest) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

type DeleteDictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDictRequest) Reset() {
	*x = DeleteDictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_dict_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDictRequest) ProtoMessage() {}

func (x *DeleteDictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dict_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDictRequest.ProtoReflect.Descriptor instead.
func (*DeleteDictRequest) Descriptor() ([]byte, []int) {
	return file_manager_dict_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDictRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_manager_dict_proto protoreflect.FileDescriptor

var file_manager_dict_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x04, 0x44, 0x69, 0x63, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x69,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x65, 0x44, 0x69, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x32, 0x20,
	0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x65,
	0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x21, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0xee, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0xf0, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x48, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manager_dict_proto_rawDescOnce sync.Once
	file_manager_dict_proto_rawDescData = file_manager_dict_proto_rawDesc
)

func file_manager_dict_proto_rawDescGZIP() []byte {
	file_manager_dict_proto_rawDescOnce.Do(func() {
		file_manager_dict_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_dict_proto_rawDescData)
	})
	return file_manager_dict_proto_rawDescData
}

var file_manager_dict_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_manager_dict_proto_goTypes = []interface{}{
	(*Dict)(nil),              // 0: manager.Dict
	(*GetDictRequest)(nil),    // 1: manager.GetDictRequest
	(*PageDictRequest)(nil),   // 2: manager.PageDictRequest
	(*PageDictReply)(nil),     // 3: manager.PageDictReply
	(*AddDictRequest)(nil),    // 4: manager.AddDictRequest
	(*UpdateDictRequest)(nil), // 5: manager.UpdateDictRequest
	(*DeleteDictRequest)(nil), // 6: manager.DeleteDictRequest
}
var file_manager_dict_proto_depIdxs = []int32{
	0, // 0: manager.PageDictReply.list:type_name -> manager.Dict
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_manager_dict_proto_init() }
func file_manager_dict_proto_init() {
	if File_manager_dict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_dict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dict); i {
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
		file_manager_dict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDictRequest); i {
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
		file_manager_dict_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageDictRequest); i {
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
		file_manager_dict_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageDictReply); i {
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
		file_manager_dict_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDictRequest); i {
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
		file_manager_dict_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDictRequest); i {
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
		file_manager_dict_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDictRequest); i {
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
	file_manager_dict_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_manager_dict_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_manager_dict_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_manager_dict_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manager_dict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manager_dict_proto_goTypes,
		DependencyIndexes: file_manager_dict_proto_depIdxs,
		MessageInfos:      file_manager_dict_proto_msgTypes,
	}.Build()
	File_manager_dict_proto = out.File
	file_manager_dict_proto_rawDesc = nil
	file_manager_dict_proto_goTypes = nil
	file_manager_dict_proto_depIdxs = nil
}
