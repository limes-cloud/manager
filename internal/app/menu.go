package app

import (
	"context"

	"github.com/limes-cloud/manager/internal/infra/dbs"

	"github.com/limes-cloud/manager/api/menu"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/types"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Menu struct {
	menu.UnimplementedMenuServer
	srv *service.Menu
}

func NewMenu() *Menu {
	return &Menu{
		srv: service.NewMenu(
			dbs.NewMenu(),
			dbs.NewApp(),
			dbs.NewRoleMenu(),
			dbs.NewScope(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewMenu()
		menu.RegisterMenuHTTPServer(hs, srv)
		menu.RegisterMenuServer(gs, srv)
	})
}

// ListCurrentMenu 获取应用信息列表
func (s *Menu) ListCurrentMenu(c context.Context, req *menu.ListCurrentMenuRequest) (*menu.ListMenuReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	in := &types.ListMenuRequest{
		AppId:      req.AppId,
		FilterRoot: req.FilterRoot,
	}
	if req.GetFilterBaseApi() {
		in.NotInTypes = []string{entity.MenuTypeBasic}
	}
	list, err := s.srv.ListCurrentMenu(ctx, in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := menu.ListMenuReply{}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list menu resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// ListMenu 获取应用信息列表
func (s *Menu) ListMenu(c context.Context, req *menu.ListMenuRequest) (*menu.ListMenuReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	in := &types.ListMenuRequest{
		AppId: &req.AppId,
	}
	if req.GetFilterBaseApi() {
		in.NotInTypes = []string{entity.MenuTypeBasic}
	}
	list, err := s.srv.ListMenu(ctx, in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := menu.ListMenuReply{}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list menu resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateMenu 创建应用信息
func (s *Menu) CreateMenu(c context.Context, req *menu.CreateMenuRequest) (*menu.CreateMenuReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Menu{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create menu req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := s.srv.CreateMenu(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &menu.CreateMenuReply{Id: id}, nil
}

// UpdateMenu 更新应用信息
func (s *Menu) UpdateMenu(c context.Context, req *menu.UpdateMenuRequest) (*menu.UpdateMenuReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Menu{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "update menu req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := s.srv.UpdateMenu(ctx, &in); err != nil {
		return nil, err
	}

	return &menu.UpdateMenuReply{}, nil
}

// DeleteMenu 删除应用信息
func (s *Menu) DeleteMenu(c context.Context, req *menu.DeleteMenuRequest) (*menu.DeleteMenuReply, error) {
	return &menu.DeleteMenuReply{}, s.srv.DeleteMenu(core.MustContext(c), req.Id)
}
