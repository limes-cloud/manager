package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/manager/api/errors"
	v1 "github.com/limes-cloud/manager/api/role/v1"
	biz "github.com/limes-cloud/manager/internal/biz/role"
	"github.com/limes-cloud/manager/internal/config"
	data "github.com/limes-cloud/manager/internal/data/role"
)

type RoleService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase

	conf *config.Config
}

func NewRoleService(conf *config.Config) *RoleService {
	return &RoleService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (r RoleService) GetRoleTree(ctx context.Context, _ *emptypb.Empty) (*v1.Role, error) {
	tree, err := r.uc.RoleTree(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := v1.Role{}
	if err := util.Transform(tree, &reply); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (r RoleService) AddRole(ctx context.Context, request *v1.AddRoleRequest) (*v1.AddRoleReply, error) {
	var role biz.Role
	if err := util.Transform(request, &role); err != nil {
		return nil, errors.Transform()
	}

	id, err := r.uc.AddRole(kratosx.MustContext(ctx), &role)
	if err != nil {
		return nil, err
	}

	return &v1.AddRoleReply{Id: id}, nil
}

func (r RoleService) DeleteRole(ctx context.Context, request *v1.DeleteRoleRequest) (*emptypb.Empty, error) {
	return nil, r.uc.DeleteRole(kratosx.MustContext(ctx), request.Id)
}

func (r RoleService) UpdateRole(ctx context.Context, request *v1.UpdateRoleRequest) (*emptypb.Empty, error) {
	var req biz.Role
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}

	return nil, r.uc.UpdateRole(kratosx.MustContext(ctx), &req)
}

func (r RoleService) GetRoleMenuIds(ctx context.Context, request *v1.GetRoleMenuIdsRequest) (*v1.GetRoleMenuIdsReply, error) {
	ids, err := r.uc.GetRoleMenuIds(kratosx.MustContext(ctx), request.RoleId)
	if err != nil {
		return nil, err
	}
	return &v1.GetRoleMenuIdsReply{List: ids}, nil
}

func (r RoleService) UpdateRoleMenus(ctx context.Context, request *v1.UpdateRoleMenuRequest) (*emptypb.Empty, error) {
	return nil, r.uc.UpdateRoleMenus(kratosx.MustContext(ctx), request.RoleId, request.MenuIds)
}
