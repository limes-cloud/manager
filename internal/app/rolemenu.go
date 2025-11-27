package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/manager/api/rolemenu"

	"github.com/limes-cloud/manager/internal/infra/dbs"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/types"
)

type RoleMenu struct {
	rolemenu.UnimplementedRoleMenuServer
	srv *service.RoleMenu
}

func NewRoleMenu() *RoleMenu {
	return &RoleMenu{
		srv: service.NewRoleMenu(
			dbs.NewRoleMenu(),
			dbs.NewApp(),
			dbs.NewScope(),
			dbs.NewMenu(),
			dbs.NewRole(),
			dbs.NewTenantAdmin(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewRoleMenu()
		rolemenu.RegisterRoleMenuHTTPServer(hs, srv)
		rolemenu.RegisterRoleMenuServer(gs, srv)
	})
}

// GetRoleMenuIds 获取角色菜单列表
func (s *RoleMenu) GetRoleMenuIds(c context.Context, req *rolemenu.GetRoleMenuIdsRequest) (*rolemenu.GetRoleMenuIdsReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, err := s.srv.GetRoleMenuIds(ctx, &types.GetRoleMenuIdsRequest{
		AppId:  req.AppId,
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	return &rolemenu.GetRoleMenuIdsReply{MenuIds: list}, nil
}

// GetMenuRoleIds 获取角色菜单列表
func (s *RoleMenu) GetMenuRoleIds(c context.Context, req *rolemenu.GetMenuRoleIdsRequest) (*rolemenu.GetMenuRoleIdsReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, err := s.srv.GetMenuRoleIds(ctx, &types.GetMenuRoleIdsRequest{MenuId: req.MenuId})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	return &rolemenu.GetMenuRoleIdsReply{RoleIds: list}, nil
}

func (s *RoleMenu) CreateMenuRoles(c context.Context, req *rolemenu.CreateMenuRolesRequest) (*rolemenu.CreateMenuRolesReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.CreateMenuRoles(ctx, &types.CreateMenuRolesRequest{
		MenuId:  req.MenuId,
		RoleIds: req.RoleIds,
	})
	if err != nil {
		return nil, err
	}
	return &rolemenu.CreateMenuRolesReply{}, nil
}

func (s *RoleMenu) CreateRoleMenus(c context.Context, req *rolemenu.CreateRoleMenusRequest) (*rolemenu.CreateRoleMenusReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.CreateRoleMenus(ctx, &types.CreateRoleMenusRequest{
		MenuIds: req.MenuIds,
		RoleId:  req.RoleId,
		AppId:   req.AppId,
	})
	if err != nil {
		return nil, err
	}
	return &rolemenu.CreateRoleMenusReply{}, nil
}

func (s *RoleMenu) DeleteMenuRoles(c context.Context, req *rolemenu.DeleteMenuRolesRequest) (*rolemenu.DeleteMenuRolesReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.DeleteMenuRoles(ctx, &types.DeleteMenuRolesRequest{
		MenuId:  req.MenuId,
		RoleIds: req.RoleIds,
	})
	if err != nil {
		return nil, err
	}
	return &rolemenu.DeleteMenuRolesReply{}, nil
}

func (s *RoleMenu) DeleteRoleMenus(c context.Context, req *rolemenu.DeleteRoleMenusRequest) (*rolemenu.DeleteRoleMenusReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.DeleteRoleMenus(ctx, &types.DeleteRoleMenusRequest{
		MenuIds: req.MenuIds,
		RoleId:  req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	return &rolemenu.DeleteRoleMenusReply{}, nil
}
