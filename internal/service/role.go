package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/api/manager/errors"
	pb "github.com/limes-cloud/manager/api/manager/role/v1"
	"github.com/limes-cloud/manager/internal/biz/role"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/data"
)

type RoleService struct {
	pb.UnimplementedRoleServer
	uc *role.UseCase
}

func NewRoleService(conf *conf.Config) *RoleService {
	return &RoleService{
		uc: role.NewUseCase(conf, data.NewRoleRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewRoleService(c)
		pb.RegisterRoleHTTPServer(hs, srv)
		pb.RegisterRoleServer(gs, srv)
	})
}

// GetRoleMenuIds 获取指定角色的菜单id列表
func (s *RoleService) GetRoleMenuIds(c context.Context, req *pb.GetRoleMenuIdsRequest) (*pb.GetRoleMenuIdsReply, error) {
	list, err := s.uc.GetRoleMenuIds(kratosx.MustContext(c), req.RoleId)
	return &pb.GetRoleMenuIdsReply{List: list}, err
}

// GetRole 获取指定的角色信息
func (s *RoleService) GetRole(c context.Context, req *pb.GetRoleRequest) (*pb.GetRoleReply, error) {
	var (
		in  = role.GetRoleRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "request transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetRole(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetRoleReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListRole 获取角色信息列表
func (s *RoleService) ListRole(c context.Context, req *pb.ListRoleRequest) (*pb.ListRoleReply, error) {
	var (
		in  = role.ListRoleRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListRole(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListRoleReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateRole 创建角色信息
func (s *RoleService) CreateRole(c context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleReply, error) {
	var (
		in  = role.Role{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateRole(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoleReply{Id: id}, nil
}

// UpdateRole 更新角色信息
func (s *RoleService) UpdateRole(c context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleReply, error) {
	var (
		in  = role.Role{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateRole(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateRoleReply{}, nil
}

// UpdateRoleStatus 更新角色信息状态
func (s *RoleService) UpdateRoleStatus(c context.Context, req *pb.UpdateRoleStatusRequest) (*pb.UpdateRoleStatusReply, error) {
	return &pb.UpdateRoleStatusReply{}, s.uc.UpdateRoleStatus(kratosx.MustContext(c), req.Id, req.Status)
}

// UpdateRoleMenu 更新角色菜单
func (s *RoleService) UpdateRoleMenu(c context.Context, req *pb.UpdateRoleMenuRequest) (*pb.UpdateRoleMenuReply, error) {
	return &pb.UpdateRoleMenuReply{}, s.uc.UpdateRoleMenu(kratosx.MustContext(c), req.RoleId, req.MenuIds)
}

// DeleteRole 删除角色信息
func (s *RoleService) DeleteRole(c context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleReply, error) {
	total, err := s.uc.DeleteRole(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteRoleReply{Total: total}, nil
}
