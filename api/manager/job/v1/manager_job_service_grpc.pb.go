// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/manager/job/manager_job_service.proto

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
	Job_ListJob_FullMethodName   = "/manager.api.manager.job.v1.Job/ListJob"
	Job_CreateJob_FullMethodName = "/manager.api.manager.job.v1.Job/CreateJob"
	Job_UpdateJob_FullMethodName = "/manager.api.manager.job.v1.Job/UpdateJob"
	Job_DeleteJob_FullMethodName = "/manager.api.manager.job.v1.Job/DeleteJob"
	Job_GetJob_FullMethodName    = "/manager.api.manager.job.v1.Job/GetJob"
)

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobClient interface {
	// ListJob 获取职位信息列表
	ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*ListJobReply, error)
	// CreateJob 创建职位信息
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobReply, error)
	// UpdateJob 更新职位信息
	UpdateJob(ctx context.Context, in *UpdateJobRequest, opts ...grpc.CallOption) (*UpdateJobReply, error)
	// DeleteJob 删除职位信息
	DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*DeleteJobReply, error)
	// GetJob 获取指定的职位信息
	GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobReply, error)
}

type jobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobClient(cc grpc.ClientConnInterface) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*ListJobReply, error) {
	out := new(ListJobReply)
	err := c.cc.Invoke(ctx, Job_ListJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobReply, error) {
	out := new(CreateJobReply)
	err := c.cc.Invoke(ctx, Job_CreateJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) UpdateJob(ctx context.Context, in *UpdateJobRequest, opts ...grpc.CallOption) (*UpdateJobReply, error) {
	out := new(UpdateJobReply)
	err := c.cc.Invoke(ctx, Job_UpdateJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*DeleteJobReply, error) {
	out := new(DeleteJobReply)
	err := c.cc.Invoke(ctx, Job_DeleteJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobReply, error) {
	out := new(GetJobReply)
	err := c.cc.Invoke(ctx, Job_GetJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
// All implementations must embed UnimplementedJobServer
// for forward compatibility
type JobServer interface {
	// ListJob 获取职位信息列表
	ListJob(context.Context, *ListJobRequest) (*ListJobReply, error)
	// CreateJob 创建职位信息
	CreateJob(context.Context, *CreateJobRequest) (*CreateJobReply, error)
	// UpdateJob 更新职位信息
	UpdateJob(context.Context, *UpdateJobRequest) (*UpdateJobReply, error)
	// DeleteJob 删除职位信息
	DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobReply, error)
	// GetJob 获取指定的职位信息
	GetJob(context.Context, *GetJobRequest) (*GetJobReply, error)
	mustEmbedUnimplementedJobServer()
}

// UnimplementedJobServer must be embedded to have forward compatible implementations.
type UnimplementedJobServer struct {
}

func (UnimplementedJobServer) ListJob(context.Context, *ListJobRequest) (*ListJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJob not implemented")
}
func (UnimplementedJobServer) CreateJob(context.Context, *CreateJobRequest) (*CreateJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedJobServer) UpdateJob(context.Context, *UpdateJobRequest) (*UpdateJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJob not implemented")
}
func (UnimplementedJobServer) DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJob not implemented")
}
func (UnimplementedJobServer) GetJob(context.Context, *GetJobRequest) (*GetJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedJobServer) mustEmbedUnimplementedJobServer() {}

// UnsafeJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServer will
// result in compilation errors.
type UnsafeJobServer interface {
	mustEmbedUnimplementedJobServer()
}

func RegisterJobServer(s grpc.ServiceRegistrar, srv JobServer) {
	s.RegisterService(&Job_ServiceDesc, srv)
}

func _Job_ListJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).ListJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_ListJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).ListJob(ctx, req.(*ListJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_CreateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_UpdateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).UpdateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_UpdateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).UpdateJob(ctx, req.(*UpdateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_DeleteJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).DeleteJob(ctx, req.(*DeleteJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_GetJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetJob(ctx, req.(*GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Job_ServiceDesc is the grpc.ServiceDesc for Job service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Job_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.api.manager.job.v1.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListJob",
			Handler:    _Job_ListJob_Handler,
		},
		{
			MethodName: "CreateJob",
			Handler:    _Job_CreateJob_Handler,
		},
		{
			MethodName: "UpdateJob",
			Handler:    _Job_UpdateJob_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _Job_DeleteJob_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _Job_GetJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/manager/job/manager_job_service.proto",
}
