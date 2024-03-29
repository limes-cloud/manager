// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: manager_dictionary_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Service_PageDictionary_FullMethodName        = "/manager_dictionary.Service/PageDictionary"
	Service_AddDictionary_FullMethodName         = "/manager_dictionary.Service/AddDictionary"
	Service_UpdateDictionary_FullMethodName      = "/manager_dictionary.Service/UpdateDictionary"
	Service_DeleteDictionary_FullMethodName      = "/manager_dictionary.Service/DeleteDictionary"
	Service_PageDictionaryValue_FullMethodName   = "/manager_dictionary.Service/PageDictionaryValue"
	Service_GetDictionaryValue_FullMethodName    = "/manager_dictionary.Service/GetDictionaryValue"
	Service_AddDictionaryValue_FullMethodName    = "/manager_dictionary.Service/AddDictionaryValue"
	Service_UpdateDictionaryValue_FullMethodName = "/manager_dictionary.Service/UpdateDictionaryValue"
	Service_DeleteDictionaryValue_FullMethodName = "/manager_dictionary.Service/DeleteDictionaryValue"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// PageDictionary 获取分页字典信息
	PageDictionary(ctx context.Context, in *PageDictionaryRequest, opts ...grpc.CallOption) (*PageDictionaryReply, error)
	// AddDictionary 添加字典信息
	AddDictionary(ctx context.Context, in *AddDictionaryRequest, opts ...grpc.CallOption) (*AddDictionaryReply, error)
	// UpdateDictionary 更新字典信息
	UpdateDictionary(ctx context.Context, in *UpdateDictionaryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteDictionary 删除字典信息
	DeleteDictionary(ctx context.Context, in *DeleteDictionaryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// PageDictionaryValue 获取分页字典值
	PageDictionaryValue(ctx context.Context, in *PageDictionaryValueRequest, opts ...grpc.CallOption) (*PageDictionaryValueReply, error)
	// GetDictionaryValue 获取指定字典值
	GetDictionaryValue(ctx context.Context, in *GetDictionaryValueRequest, opts ...grpc.CallOption) (*GetDictionaryValueReply, error)
	// AddDictionaryValue 删除字典值
	AddDictionaryValue(ctx context.Context, in *AddDictionaryValueRequest, opts ...grpc.CallOption) (*AddDictionaryValueReply, error)
	// UpdateDictionaryValue 更新字典值
	UpdateDictionaryValue(ctx context.Context, in *UpdateDictionaryValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteDictionaryValue 删除字典值
	DeleteDictionaryValue(ctx context.Context, in *DeleteDictionaryValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) PageDictionary(ctx context.Context, in *PageDictionaryRequest, opts ...grpc.CallOption) (*PageDictionaryReply, error) {
	out := new(PageDictionaryReply)
	err := c.cc.Invoke(ctx, Service_PageDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddDictionary(ctx context.Context, in *AddDictionaryRequest, opts ...grpc.CallOption) (*AddDictionaryReply, error) {
	out := new(AddDictionaryReply)
	err := c.cc.Invoke(ctx, Service_AddDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateDictionary(ctx context.Context, in *UpdateDictionaryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_UpdateDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteDictionary(ctx context.Context, in *DeleteDictionaryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_DeleteDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PageDictionaryValue(ctx context.Context, in *PageDictionaryValueRequest, opts ...grpc.CallOption) (*PageDictionaryValueReply, error) {
	out := new(PageDictionaryValueReply)
	err := c.cc.Invoke(ctx, Service_PageDictionaryValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetDictionaryValue(ctx context.Context, in *GetDictionaryValueRequest, opts ...grpc.CallOption) (*GetDictionaryValueReply, error) {
	out := new(GetDictionaryValueReply)
	err := c.cc.Invoke(ctx, Service_GetDictionaryValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddDictionaryValue(ctx context.Context, in *AddDictionaryValueRequest, opts ...grpc.CallOption) (*AddDictionaryValueReply, error) {
	out := new(AddDictionaryValueReply)
	err := c.cc.Invoke(ctx, Service_AddDictionaryValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateDictionaryValue(ctx context.Context, in *UpdateDictionaryValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_UpdateDictionaryValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteDictionaryValue(ctx context.Context, in *DeleteDictionaryValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_DeleteDictionaryValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// PageDictionary 获取分页字典信息
	PageDictionary(context.Context, *PageDictionaryRequest) (*PageDictionaryReply, error)
	// AddDictionary 添加字典信息
	AddDictionary(context.Context, *AddDictionaryRequest) (*AddDictionaryReply, error)
	// UpdateDictionary 更新字典信息
	UpdateDictionary(context.Context, *UpdateDictionaryRequest) (*emptypb.Empty, error)
	// DeleteDictionary 删除字典信息
	DeleteDictionary(context.Context, *DeleteDictionaryRequest) (*emptypb.Empty, error)
	// PageDictionaryValue 获取分页字典值
	PageDictionaryValue(context.Context, *PageDictionaryValueRequest) (*PageDictionaryValueReply, error)
	// GetDictionaryValue 获取指定字典值
	GetDictionaryValue(context.Context, *GetDictionaryValueRequest) (*GetDictionaryValueReply, error)
	// AddDictionaryValue 删除字典值
	AddDictionaryValue(context.Context, *AddDictionaryValueRequest) (*AddDictionaryValueReply, error)
	// UpdateDictionaryValue 更新字典值
	UpdateDictionaryValue(context.Context, *UpdateDictionaryValueRequest) (*emptypb.Empty, error)
	// DeleteDictionaryValue 删除字典值
	DeleteDictionaryValue(context.Context, *DeleteDictionaryValueRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) PageDictionary(context.Context, *PageDictionaryRequest) (*PageDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageDictionary not implemented")
}
func (UnimplementedServiceServer) AddDictionary(context.Context, *AddDictionaryRequest) (*AddDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDictionary not implemented")
}
func (UnimplementedServiceServer) UpdateDictionary(context.Context, *UpdateDictionaryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDictionary not implemented")
}
func (UnimplementedServiceServer) DeleteDictionary(context.Context, *DeleteDictionaryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDictionary not implemented")
}
func (UnimplementedServiceServer) PageDictionaryValue(context.Context, *PageDictionaryValueRequest) (*PageDictionaryValueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageDictionaryValue not implemented")
}
func (UnimplementedServiceServer) GetDictionaryValue(context.Context, *GetDictionaryValueRequest) (*GetDictionaryValueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDictionaryValue not implemented")
}
func (UnimplementedServiceServer) AddDictionaryValue(context.Context, *AddDictionaryValueRequest) (*AddDictionaryValueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDictionaryValue not implemented")
}
func (UnimplementedServiceServer) UpdateDictionaryValue(context.Context, *UpdateDictionaryValueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDictionaryValue not implemented")
}
func (UnimplementedServiceServer) DeleteDictionaryValue(context.Context, *DeleteDictionaryValueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDictionaryValue not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_PageDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PageDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_PageDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PageDictionary(ctx, req.(*PageDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AddDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddDictionary(ctx, req.(*AddDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateDictionary(ctx, req.(*UpdateDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteDictionary(ctx, req.(*DeleteDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PageDictionaryValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageDictionaryValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PageDictionaryValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_PageDictionaryValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PageDictionaryValue(ctx, req.(*PageDictionaryValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetDictionaryValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDictionaryValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetDictionaryValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetDictionaryValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetDictionaryValue(ctx, req.(*GetDictionaryValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddDictionaryValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDictionaryValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddDictionaryValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AddDictionaryValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddDictionaryValue(ctx, req.(*AddDictionaryValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateDictionaryValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDictionaryValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateDictionaryValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateDictionaryValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateDictionaryValue(ctx, req.(*UpdateDictionaryValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteDictionaryValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDictionaryValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteDictionaryValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteDictionaryValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteDictionaryValue(ctx, req.(*DeleteDictionaryValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager_dictionary.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PageDictionary",
			Handler:    _Service_PageDictionary_Handler,
		},
		{
			MethodName: "AddDictionary",
			Handler:    _Service_AddDictionary_Handler,
		},
		{
			MethodName: "UpdateDictionary",
			Handler:    _Service_UpdateDictionary_Handler,
		},
		{
			MethodName: "DeleteDictionary",
			Handler:    _Service_DeleteDictionary_Handler,
		},
		{
			MethodName: "PageDictionaryValue",
			Handler:    _Service_PageDictionaryValue_Handler,
		},
		{
			MethodName: "GetDictionaryValue",
			Handler:    _Service_GetDictionaryValue_Handler,
		},
		{
			MethodName: "AddDictionaryValue",
			Handler:    _Service_AddDictionaryValue_Handler,
		},
		{
			MethodName: "UpdateDictionaryValue",
			Handler:    _Service_UpdateDictionaryValue_Handler,
		},
		{
			MethodName: "DeleteDictionaryValue",
			Handler:    _Service_DeleteDictionaryValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager_dictionary_service.proto",
}
