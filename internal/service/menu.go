package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/manager/api/errors"
	v1 "github.com/limes-cloud/manager/api/menu/v1"
	biz "github.com/limes-cloud/manager/internal/biz/menu"
	"github.com/limes-cloud/manager/internal/config"
	data "github.com/limes-cloud/manager/internal/data/menu"
)

type MenuService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase

	conf *config.Config
}

func NewMenuService(conf *config.Config) *MenuService {
	return &MenuService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (m MenuService) AddMenu(ctx context.Context, request *v1.AddMenuRequest) (*v1.AddMenuReply, error) {
	var menu biz.Menu
	if err := util.Transform(request, &menu); err != nil {
		return nil, errors.Transform()
	}

	id, err := m.uc.AddMenu(kratosx.MustContext(ctx), &menu)
	if err != nil {
		return nil, err
	}

	return &v1.AddMenuReply{Id: id}, nil
}

func (m MenuService) GetMenuTree(ctx context.Context, _ *emptypb.Empty) (*v1.GetMenuTreeReply, error) {
	tree, err := m.uc.MenuTree(kratosx.MustContext(ctx), nil)
	if err != nil {
		return nil, err
	}

	reply := v1.GetMenuTreeReply{}
	if err := util.Transform(tree, &reply.List); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (m MenuService) GetMenuTreeFromRole(ctx context.Context, _ *emptypb.Empty) (*v1.GetMenuTreeReply, error) {
	tree, err := m.uc.MenuTreeFromRole(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := v1.GetMenuTreeReply{}
	if err := util.Transform(tree, &reply.List); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (m MenuService) DeleteMenu(ctx context.Context, request *v1.DeleteMenuRequest) (*emptypb.Empty, error) {
	return nil, m.uc.DeleteMenu(kratosx.MustContext(ctx), request.Id)
}

func (m MenuService) UpdateMenu(ctx context.Context, request *v1.UpdateMenuRequest) (*emptypb.Empty, error) {
	var menu biz.Menu
	if err := util.Transform(request, &menu); err != nil {
		return nil, errors.Transform()
	}

	return nil, m.uc.UpdateMenu(kratosx.MustContext(ctx), &menu)
}
