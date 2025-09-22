package app

import (
	"context"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/deptrole"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type DeptRole struct {
	deptrole.UnimplementedDeptRoleServer
	srv *service.DeptRole
}

func NewDeptRole() *DeptRole {
	return &DeptRole{
		srv: service.NewDeptRole(
			dbs.NewDeptRole(),
			dbs.NewScope(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewDeptRole()
		deptrole.RegisterDeptRoleHTTPServer(hs, srv)
		deptrole.RegisterDeptRoleServer(gs, srv)
	})
}

// ListDeptRole 获取角色菜单列表
func (s *DeptRole) ListDeptRole(c context.Context, req *deptrole.ListDeptRoleRequest) (*deptrole.ListDeptRoleReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, total, err := s.srv.ListDeptRole(ctx, &types.ListDeptRoleRequest{
		Search: page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		DeptId: req.DeptId,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := deptrole.ListDeptRoleReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list dept role resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

func (s *DeptRole) ListRoleDept(c context.Context, req *deptrole.ListRoleDeptRequest) (*deptrole.ListRoleDeptReply, error) {
	ctx := core.MustContext(c)
	list, total, err := s.srv.ListRoleDept(ctx, &types.ListRoleDeptRequest{
		Search: page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		RoleId: req.RoleId,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}
	reply := deptrole.ListRoleDeptReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list role dept \resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *DeptRole) CreateRoleDepts(c context.Context, req *deptrole.CreateRoleDeptsRequest) (*deptrole.CreateRoleDeptsReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.CreateRoleDepts(ctx, &types.CreateRoleDeptsRequest{
		RoleId:  req.RoleId,
		DeptIds: req.DeptIds,
	})
	if err != nil {
		return nil, err
	}
	return &deptrole.CreateRoleDeptsReply{}, nil
}

func (s *DeptRole) CreateDeptRoles(c context.Context, req *deptrole.CreateDeptRolesRequest) (*deptrole.CreateDeptRolesReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.CreateDeptRoles(ctx, &types.CreateDeptRolesRequest{
		DeptId:  req.DeptId,
		RoleIds: req.RoleIds,
	})
	if err != nil {
		return nil, err
	}
	return &deptrole.CreateDeptRolesReply{}, nil
}

func (s *DeptRole) DeleteDeptRoles(c context.Context, req *deptrole.DeleteDeptRolesRequest) (*deptrole.DeleteDeptRolesReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.DeleteDeptRoles(ctx, &types.DeleteDeptRolesRequest{
		DeptId:  req.DeptId,
		RoleIds: req.RoleIds,
	})
	if err != nil {
		return nil, err
	}
	return &deptrole.DeleteDeptRolesReply{}, nil
}

func (s *DeptRole) DeleteRoleDepts(c context.Context, req *deptrole.DeleteRoleDeptsRequest) (*deptrole.DeleteRoleDeptsReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.DeleteRoleDepts(ctx, &types.DeleteRoleDeptsRequest{
		DeptIds: req.DeptIds,
		RoleId:  req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	return &deptrole.DeleteRoleDeptsReply{}, nil
}
