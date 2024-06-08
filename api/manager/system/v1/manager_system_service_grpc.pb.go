// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/manager/system/manager_system_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	System_GetSystemSetting_FullMethodName = "/manager.api.manager.system.v1.System/GetSystemSetting"
)

// SystemClient is the client API for System service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemClient interface {
	// GetSystemSetting 获取系统设置
	GetSystemSetting(ctx context.Context, in *GetSystemSettingRequest, opts ...grpc.CallOption) (*GetSystemSettingReply, error)
}

type systemClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemClient(cc grpc.ClientConnInterface) SystemClient {
	return &systemClient{cc}
}

func (c *systemClient) GetSystemSetting(ctx context.Context, in *GetSystemSettingRequest, opts ...grpc.CallOption) (*GetSystemSettingReply, error) {
	out := new(GetSystemSettingReply)
	err := c.cc.Invoke(ctx, System_GetSystemSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServer is the server API for System service.
// All implementations must embed UnimplementedSystemServer
// for forward compatibility
type SystemServer interface {
	// GetSystemSetting 获取系统设置
	GetSystemSetting(context.Context, *GetSystemSettingRequest) (*GetSystemSettingReply, error)
	mustEmbedUnimplementedSystemServer()
}

// UnimplementedSystemServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServer struct {
}

func (UnimplementedSystemServer) GetSystemSetting(context.Context, *GetSystemSettingRequest) (*GetSystemSettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemSetting not implemented")
}
func (UnimplementedSystemServer) mustEmbedUnimplementedSystemServer() {}

// UnsafeSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemServer will
// result in compilation errors.
type UnsafeSystemServer interface {
	mustEmbedUnimplementedSystemServer()
}

func RegisterSystemServer(s grpc.ServiceRegistrar, srv SystemServer) {
	s.RegisterService(&System_ServiceDesc, srv)
}

func _System_GetSystemSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).GetSystemSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_GetSystemSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).GetSystemSetting(ctx, req.(*GetSystemSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// System_ServiceDesc is the grpc.ServiceDesc for System service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var System_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.api.manager.system.v1.System",
	HandlerType: (*SystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSystemSetting",
			Handler:    _System_GetSystemSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/manager/system/manager_system_service.proto",
}