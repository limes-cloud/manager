package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	ktypes "github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/manager/api/manager/errors"
	pb "github.com/limes-cloud/manager/api/manager/role/v1"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type Role struct {
	pb.UnimplementedRoleServer
	srv *service.Role
}

func NewRole(conf *conf.Config) *Role {
	return &Role{
		srv: service.NewRole(conf, dbs.NewRole(), dbs.NewMenu(), dbs.NewRbac()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewRole(c)
		pb.RegisterRoleHTTPServer(hs, srv)
		pb.RegisterRoleServer(gs, srv)
	})
}

// GetRoleMenuIds 获取指定角色的菜单id列表
func (s *Role) GetRoleMenuIds(c context.Context, req *pb.GetRoleMenuIdsRequest) (*pb.GetRoleMenuIdsReply, error) {
	list, err := s.srv.GetRoleMenuIds(kratosx.MustContext(c), req.RoleId)
	return &pb.GetRoleMenuIdsReply{List: list}, err
}

// GetRole 获取指定的角色信息
func (s *Role) GetRole(c context.Context, req *pb.GetRoleRequest) (*pb.GetRoleReply, error) {
	var (
		ent *entity.Role
		err error
	)
	switch req.Params.(type) {
	case *pb.GetRoleRequest_Id:
		ent, err = s.srv.GetRole(kratosx.MustContext(c), req.GetId())
	case *pb.GetRoleRequest_Keyword:
		ent, err = s.srv.GetRoleByKeyword(kratosx.MustContext(c), req.GetKeyword())
	default:
		return nil, errors.ParamsError()
	}
	if err != nil {
		return nil, err
	}

	return &pb.GetRoleReply{
		Id:            ent.Id,
		ParentId:      ent.ParentId,
		Name:          ent.Name,
		Keyword:       ent.Keyword,
		Status:        ent.Status,
		DataScope:     ent.DataScope,
		DepartmentIds: ent.DepartmentIds,
		Description:   ent.Description,
		CreatedAt:     uint32(ent.CreatedAt),
		UpdatedAt:     uint32(ent.UpdatedAt),
	}, nil
}

// ListRole 获取角色信息列表
func (s *Role) ListRole(c context.Context, req *pb.ListRoleRequest) (*pb.ListRoleReply, error) {
	var ctx = kratosx.MustContext(c)
	result, err := s.srv.ListRole(ctx, &types.ListRoleRequest{
		Name:    req.Name,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListRoleReply{}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// ListCurrentRole 获取角色信息列表
func (s *Role) ListCurrentRole(c context.Context, req *pb.ListRoleRequest) (*pb.ListRoleReply, error) {
	var ctx = kratosx.MustContext(c)
	result, err := s.srv.ListCurrentRole(ctx, &types.ListRoleRequest{
		Name:    req.Name,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListRoleReply{}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateRole 创建角色信息
func (s *Role) CreateRole(c context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleReply, error) {
	var ctx = kratosx.MustContext(c)

	id, err := s.srv.CreateRole(ctx, &entity.Role{
		ParentId:      req.ParentId,
		Name:          req.Name,
		Keyword:       req.Keyword,
		Status:        req.Status,
		DataScope:     req.DataScope,
		DepartmentIds: req.DepartmentIds,
		// JobScope:      req.JobScope,
		// JobIds:        req.JobIds,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoleReply{Id: id}, nil
}

// UpdateRole 更新角色信息
func (s *Role) UpdateRole(c context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleReply, error) {
	if err := s.srv.UpdateRole(kratosx.MustContext(c), &entity.Role{
		BaseModel:     ktypes.BaseModel{Id: req.Id},
		ParentId:      req.ParentId,
		Name:          req.Name,
		DataScope:     req.DataScope,
		DepartmentIds: req.DepartmentIds,
		// JobScope:      req.JobScope,
		// JobIds:        req.JobIds,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateRoleReply{}, nil
}

// UpdateRoleStatus 更新角色信息状态
func (s *Role) UpdateRoleStatus(c context.Context, req *pb.UpdateRoleStatusRequest) (*pb.UpdateRoleStatusReply, error) {
	return &pb.UpdateRoleStatusReply{}, s.srv.UpdateRoleStatus(kratosx.MustContext(c), req.Id, req.Status)
}

// UpdateRoleMenu 更新角色菜单
func (s *Role) UpdateRoleMenu(c context.Context, req *pb.UpdateRoleMenuRequest) (*pb.UpdateRoleMenuReply, error) {
	return &pb.UpdateRoleMenuReply{}, s.srv.UpdateRoleMenu(kratosx.MustContext(c), req.RoleId, req.MenuIds)
}

// DeleteRole 删除角色信息
func (s *Role) DeleteRole(c context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleReply, error) {
	return &pb.DeleteRoleReply{}, s.srv.DeleteRole(kratosx.MustContext(c), req.Id)
}
