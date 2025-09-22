package app

import (
	"context"

	"github.com/limes-cloud/manager/internal/domain/entity"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/manager/api/tenant"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
)

type Tenant struct {
	tenant.UnimplementedTenantServer
	srv *service.Tenant
}

// NewTenant 初始化租户应用
func NewTenant() *Tenant {
	return &Tenant{
		srv: service.NewTenant(
			dbs.NewTenant(),
		),
	}
}

// init 应用注册
func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewTenant()
		tenant.RegisterTenantHTTPServer(hs, srv)
		tenant.RegisterTenantServer(gs, srv)
	})
}

// GetTenant 获取指定租户信息
func (app *Tenant) GetTenant(c context.Context, req *tenant.GetTenantRequest) (*tenant.GetTenantReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	ent, err := app.srv.GetTenant(ctx, &types.GetTenantRequest{
		Id:      req.Id,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := tenant.GetTenantReply{}
	if err := value.Transform(ent, &reply); err != nil {
		ctx.Logger().Errorw("msg", "get tenant resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListTenant 获取租户信息列表
func (app *Tenant) ListTenant(c context.Context, req *tenant.ListTenantRequest) (*tenant.ListTenantReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListTenantRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list tenant req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := app.srv.ListTenant(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := tenant.ListTenantReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get tenant resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateTenant 创建租户信息
func (app *Tenant) CreateTenant(c context.Context, req *tenant.CreateTenantRequest) (*tenant.CreateTenantReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Tenant{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create tenant req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateTenant(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &tenant.CreateTenantReply{Id: id}, nil
}

// UpdateTenant 更新租户信息
func (app *Tenant) UpdateTenant(c context.Context, req *tenant.UpdateTenantRequest) (*tenant.UpdateTenantReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Tenant{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create tenant req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateTenant(ctx, &in); err != nil {
		return nil, err
	}

	return &tenant.UpdateTenantReply{}, nil
}

// DeleteTenant 删除租户信息
func (app *Tenant) DeleteTenant(c context.Context, req *tenant.DeleteTenantRequest) (*tenant.DeleteTenantReply, error) {
	return &tenant.DeleteTenantReply{}, app.srv.DeleteTenant(core.MustContext(c), req.Id)
}
