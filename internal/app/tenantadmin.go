package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/api/tenantadmin"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type TenantAdmin struct {
	tenantadmin.UnimplementedTenantAdminServer
	srv *service.TenantAdmin
}

// NewTenantAdmin 初始化租户应用
func NewTenantAdmin() *TenantAdmin {
	return &TenantAdmin{
		srv: service.NewTenantAdmin(
			dbs.NewTenantAdmin(),
		),
	}
}

// init 应用注册
func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewTenantAdmin()
		tenantadmin.RegisterTenantAdminHTTPServer(hs, srv)
		tenantadmin.RegisterTenantAdminServer(gs, srv)
	})
}

// ListTenantAdmin 获取租户信息列表
func (admin *TenantAdmin) ListTenantAdmin(c context.Context, req *tenantadmin.ListTenantAdminRequest) (*tenantadmin.ListTenantAdminReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, total, err := admin.srv.ListTenantAdmin(ctx, &types.ListTenantAdminRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		TenantId: req.TenantId,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := tenantadmin.ListTenantAdminReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &tenantadmin.ListTenantAdminReply_Data{
			Id:        item.User.Id,
			Avatar:    item.User.Avatar,
			Username:  item.User.Username,
			Nickname:  item.User.Nickname,
			CreatedAt: uint32(item.User.CreatedAt),
		})
	}
	return &reply, nil
}

// CreateTenantAdmin 创建租户信息
func (admin *TenantAdmin) CreateTenantAdmin(c context.Context, req *tenantadmin.CreateTenantAdminRequest) (*tenantadmin.CreateTenantAdminReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	id, err := admin.srv.CreateTenantAdmin(ctx, &types.CreateTenantAdminRequest{
		TenantId: req.TenantId,
		UserId:   req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &tenantadmin.CreateTenantAdminReply{Id: id}, nil
}

// DeleteTenantAdmin 删除租户信息
func (admin *TenantAdmin) DeleteTenantAdmin(c context.Context, req *tenantadmin.DeleteTenantAdminRequest) (*tenantadmin.DeleteTenantAdminReply, error) {
	return &tenantadmin.DeleteTenantAdminReply{}, admin.srv.DeleteTenantAdmin(core.MustContext(c), &types.DeleteTenantAdminRequest{
		TenantId: req.TenantId,
		UserId:   req.UserId,
	})
}
