package app

import (
	"context"
	"encoding/json"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/spf13/cast"

	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/manager/internal/domain/entity"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/pkg/value"
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
			dbs.NewApp(),
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
		Id:      req.GetId(),
		Keyword: req.GetKeyword(),
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := tenant.GetTenantReply{
		Id:          ent.Id,
		Keyword:     ent.Keyword,
		Logo:        ent.Logo,
		Name:        ent.Name,
		Description: ent.Description,
		Status:      value.Value(ent.Status),
		Weight:      value.Value(ent.Weight),
		CreatedAt:   cast.ToUint32(ent.CreatedAt),
		UpdatedAt:   cast.ToUint32(ent.UpdatedAt),
	}
	_ = json.Unmarshal([]byte(ent.Setting), &reply.Setting)
	return &reply, nil
}

// ListAppTenant 获取可见的租户信息列表
func (app *Tenant) ListAppTenant(c context.Context, req *tenant.ListAppTenantRequest) (*tenant.ListAppTenantReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, _, err := app.srv.ListTenant(ctx, &types.ListTenantRequest{
		Status: value.Pointer(true),
		App:    value.Pointer(req.App),
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := tenant.ListAppTenantReply{}
	for _, ent := range list {
		reply.List = append(reply.List, &tenant.ListAppTenantReply_Data{
			Id:      ent.Id,
			Keyword: ent.Keyword,
			Logo:    ent.Logo,
			Name:    ent.Name,
		})
	}

	return &reply, nil
}

// ListTenant 获取租户信息列表
func (app *Tenant) ListTenant(c context.Context, req *tenant.ListTenantRequest) (*tenant.ListTenantReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, total, err := app.srv.ListTenant(ctx, &types.ListTenantRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
			Order:    req.Order,
			OrderBy:  req.OrderBy,
		},
		Keyword: req.Keyword,
		Name:    req.Name,
		Status:  req.Status,
		App:     req.App,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := tenant.ListTenantReply{Total: total}
	for _, ent := range list {
		reply.List = append(reply.List, &tenant.ListTenantReply_Data{
			Id:          ent.Id,
			Keyword:     ent.Keyword,
			Logo:        ent.Logo,
			Name:        ent.Name,
			Description: ent.Description,
			Status:      value.Value(ent.Status),
			Weight:      value.Value(ent.Weight),
			CreatedAt:   cast.ToUint32(ent.CreatedAt),
			UpdatedAt:   cast.ToUint32(ent.UpdatedAt),
		})
	}

	return &reply, nil
}

// CreateTenant 创建租户信息
func (app *Tenant) CreateTenant(c context.Context, req *tenant.CreateTenantRequest) (*tenant.CreateTenantReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	id, err := app.srv.CreateTenant(ctx, &entity.Tenant{
		Keyword:     req.Keyword,
		Logo:        req.Logo,
		Name:        req.Name,
		Description: req.Description,
		Status:      value.Pointer(false),
		Weight:      value.Pointer(req.Weight),
		Setting:     value.ObjToString(req.Setting),
	})
	if err != nil {
		return nil, err
	}

	return &tenant.CreateTenantReply{Id: id}, nil
}

// UpdateTenant 更新租户信息
func (app *Tenant) UpdateTenant(c context.Context, req *tenant.UpdateTenantRequest) (*tenant.UpdateTenantReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	if err := app.srv.UpdateTenant(ctx, &entity.Tenant{
		DeleteModel: model.DeleteModel{Id: req.Id},
		Logo:        value.Value(req.Logo),
		Name:        value.Value(req.Name),
		Description: value.Value(req.Description),
		Status:      req.Status,
		Weight:      req.Weight,
		Setting:     value.ObjToString(req.Setting),
	}); err != nil {
		return nil, err
	}

	return &tenant.UpdateTenantReply{}, nil
}

// DeleteTenant 删除租户信息
func (app *Tenant) DeleteTenant(c context.Context, req *tenant.DeleteTenantRequest) (*tenant.DeleteTenantReply, error) {
	return &tenant.DeleteTenantReply{}, app.srv.DeleteTenant(core.MustContext(c), req.Id)
}
