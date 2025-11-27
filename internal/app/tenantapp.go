package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/kratosx/pkg/value"
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

// GetTenantApp 获取租户菜单
func (app *TenantApp) GetTenantApp(c context.Context, req *tenantapp.GetTenantAppRequest) (*tenantapp.GetTenantAppReply, error) {
	reply, err := app.srv.GetTenantApp(core.MustContext(c), &types.GetTenantAppRequest{
		AppId:    req.AppId,
		TenantId: req.TenantId,
	})
	if err != nil {
		return nil, err
	}

	setting := reply.GetSetting()
	return &tenantapp.GetTenantAppReply{
		Id:        reply.Id,
		AppId:     reply.AppId,
		TenantId:  reply.TenantId,
		ExpiredAt: reply.ExpiredAt,
		CreatedAt: uint32(reply.CreatedAt),
		UpdatedAt: uint32(reply.UpdatedAt),
		App: &tenantapp.App{
			Id:      reply.App.Id,
			Name:    reply.App.Name,
			Keyword: reply.App.Keyword,
			Logo:    reply.App.Logo,
		},
		MenuIds: reply.MenuIds,
		Setting: &tenantapp.TenantAppSetting{
			EnableNotice: setting.EnableNotice,
			NoticeEmail:  setting.NoticeEmail,
		},
	}, nil
}

// ListTenantApp 获取租户信息列表
func (app *TenantApp) ListTenantApp(c context.Context, req *tenantapp.ListTenantAppRequest) (*tenantapp.ListTenantAppReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, total, err := app.srv.ListTenantApp(ctx, &types.ListTenantAppRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
			Order:    req.Order,
			OrderBy:  req.OrderBy,
		},
		AppKeyword: req.AppKeyword,
		AppName:    req.AppName,
		TenantId:   req.TenantId,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := tenantapp.ListTenantAppReply{Total: total}
	for _, item := range list {
		setting := item.GetSetting()
		reply.List = append(reply.List, &tenantapp.ListTenantAppReply_Data{
			Id:        item.Id,
			AppId:     item.AppId,
			TenantId:  item.TenantId,
			ExpiredAt: item.ExpiredAt,
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
			App: &tenantapp.App{
				Id:      item.App.Id,
				Name:    item.App.Name,
				Keyword: item.App.Keyword,
				Logo:    item.App.Logo,
			},
			Setting: &tenantapp.TenantAppSetting{
				EnableNotice: setting.EnableNotice,
				NoticeEmail:  setting.NoticeEmail,
			},
		})
	}
	return &reply, nil
}

// CreateTenantApp 创建租户信息
func (app *TenantApp) CreateTenantApp(c context.Context, req *tenantapp.CreateTenantAppRequest) (*tenantapp.CreateTenantAppReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	id, err := app.srv.CreateTenantApp(ctx, &types.CreateTenantAppRequest{
		TenantId:  req.TenantId,
		AppId:     req.AppId,
		ExpiredAt: req.ExpiredAt,
		MenuIds:   req.MenuIds,
		Setting:   value.ObjToString(req.Setting),
	})
	if err != nil {
		return nil, err
	}

	return &tenantapp.CreateTenantAppReply{Id: id}, nil
}

// UpdateTenantApp 更新租户信息
func (app *TenantApp) UpdateTenantApp(c context.Context, req *tenantapp.UpdateTenantAppRequest) (*tenantapp.UpdateTenantAppReply, error) {
	ctx := core.MustContext(c)

	if err := app.srv.UpdateTenantApp(ctx, &types.UpdateTenantAppRequest{
		TenantId:  req.TenantId,
		AppId:     req.AppId,
		ExpiredAt: req.ExpiredAt,
		MenuIds:   req.MenuIds,
		Setting:   value.ObjToString(req.Setting),
	}); err != nil {
		return nil, err
	}

	return &tenantapp.UpdateTenantAppReply{}, nil
}

// DeleteTenantApp 删除租户信息
func (app *TenantApp) DeleteTenantApp(c context.Context, req *tenantapp.DeleteTenantAppRequest) (*tenantapp.DeleteTenantAppReply, error) {
	return &tenantapp.DeleteTenantAppReply{}, app.srv.DeleteTenantApp(core.MustContext(c), req.TenantId, req.AppId)
}
