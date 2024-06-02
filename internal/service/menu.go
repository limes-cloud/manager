package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/api/manager/errors"
	pb "github.com/limes-cloud/manager/api/manager/menu/v1"
	"github.com/limes-cloud/manager/internal/biz/menu"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/data"
)

type MenuService struct {
	pb.UnimplementedMenuServer
	uc *menu.UseCase
}

func NewMenuService(conf *conf.Config) *MenuService {
	return &MenuService{
		uc: menu.NewUseCase(conf, data.NewMenuRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewMenuService(c)
		pb.RegisterMenuHTTPServer(hs, srv)
		pb.RegisterMenuServer(gs, srv)
	})
}

// ListMenu 获取菜单信息列表
func (s *MenuService) ListMenu(c context.Context, req *pb.ListMenuRequest) (*pb.ListMenuReply, error) {
	var (
		in  = menu.ListMenuRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListMenu(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListMenuReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// ListMenuByCurRole 获取当前角色的菜单信息列表
func (s *MenuService) ListMenuByCurRole(c context.Context, _ *pb.ListMenuByCurRoleRequest) (*pb.ListMenuByCurRoleReply, error) {
	var (
		ctx = kratosx.MustContext(c)
	)

	result, total, err := s.uc.ListMenuByCurRole(ctx)
	if err != nil {
		return nil, err
	}

	reply := pb.ListMenuByCurRoleReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateMenu 创建菜单信息
func (s *MenuService) CreateMenu(c context.Context, req *pb.CreateMenuRequest) (*pb.CreateMenuReply, error) {
	var (
		in  = menu.Menu{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateMenu(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateMenuReply{Id: id}, nil
}

// UpdateMenu 更新菜单信息
func (s *MenuService) UpdateMenu(c context.Context, req *pb.UpdateMenuRequest) (*pb.UpdateMenuReply, error) {
	var (
		in  = menu.Menu{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateMenu(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateMenuReply{}, nil
}

// DeleteMenu 删除菜单信息
func (s *MenuService) DeleteMenu(c context.Context, req *pb.DeleteMenuRequest) (*pb.DeleteMenuReply, error) {
	total, err := s.uc.DeleteMenu(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteMenuReply{Total: total}, nil
}
