package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/api/tenantapp"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type TenantApp struct {
	tenantapp.UnimplementedTenantServer
	srv *service.TenantApp
}

// NewTenantApp 初始化租户应用
func NewTenantApp() *TenantApp {
	return &TenantApp{
		srv: service.NewTenantApp(
			dbs.NewTenantApp(),
		),
	}
}

// init 应用注册
func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewTenantApp()
		tenantapp.RegisterTenantHTTPServer(hs, srv)
		tenantapp.RegisterTenantServer(gs, srv)
	})
}

// ListTenantApp 获取租户信息列表
func (app *TenantApp) ListTenantApp(c context.Context, req *tenantapp.ListTenantAppRequest) (*tenantapp.ListTenantAppReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListTenantAppRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list tenantapp app req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := app.srv.ListTenantApp(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := tenantapp.ListTenantAppReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get tenantapp app resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateTenantApp 创建租户信息
func (app *TenantApp) CreateTenantApp(c context.Context, req *tenantapp.CreateTenantAppRequest) (*tenantapp.CreateTenantAppReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.CreateTenantAppRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create tenantapp app req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateTenantApp(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &tenantapp.CreateTenantAppReply{Id: id}, nil
}

// UpdateTenantApp 更新租户信息
func (app *TenantApp) UpdateTenantApp(c context.Context, req *tenantapp.UpdateTenantAppRequest) (*tenantapp.UpdateTenantAppReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.UpdateTenantAppRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create tenantapp app req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateTenantApp(ctx, &in); err != nil {
		return nil, err
	}

	return &tenantapp.UpdateTenantAppReply{}, nil
}

// DeleteTenantApp 删除租户信息
func (app *TenantApp) DeleteTenantApp(c context.Context, req *tenantapp.DeleteTenantAppRequest) (*tenantapp.DeleteTenantAppReply, error) {
	return &tenantapp.DeleteTenantAppReply{}, app.srv.DeleteTenantApp(core.MustContext(c), req.TenantId, req.AppId)
}

// GetTenantAppMenuIds 获取租户菜单
func (app *TenantApp) GetTenantAppMenuIds(c context.Context, req *tenantapp.GetTenantAppMenuIdsRequest) (*tenantapp.GetTenantAppMenuIdsReply, error) {
	ids, err := app.srv.GetTenantAppMenuIds(core.MustContext(c), req.TenantId, req.AppId)
	if err != nil {
		return nil, err
	}
	return &tenantapp.GetTenantAppMenuIdsReply{MenuIds: ids}, nil
}
