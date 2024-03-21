// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.24.4
// source: manager_setting_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationServiceGetSetting = "/manager_setting.Service/GetSetting"

type ServiceHTTPServer interface {
	// GetSetting GetSetting 获取当前系统的配置
	GetSetting(context.Context, *emptypb.Empty) (*GetSettingReply, error)
}

func RegisterServiceHTTPServer(s *http.Server, srv ServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/manager/v1/setting", _Service_GetSetting0_HTTP_Handler(srv))
}

func _Service_GetSetting0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceGetSetting)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetSetting(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSettingReply)
		return ctx.Result(200, reply)
	}
}

type ServiceHTTPClient interface {
	GetSetting(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetSettingReply, err error)
}

type ServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewServiceHTTPClient(client *http.Client) ServiceHTTPClient {
	return &ServiceHTTPClientImpl{client}
}

func (c *ServiceHTTPClientImpl) GetSetting(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetSettingReply, error) {
	var out GetSettingReply
	pattern := "/manager/v1/setting"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceGetSetting))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}