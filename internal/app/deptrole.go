package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/kratosx/model/page"

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
			dbs.NewTenantAdmin(),
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

func (s *DeptRole) CreateDeptRole(c context.Context, req *deptrole.CreateDeptRoleRequest) (*deptrole.CreateDeptRoleReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.CreateDeptRole(ctx, &types.CreateDeptRoleRequest{
		DeptId: req.DeptId,
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	return &deptrole.CreateDeptRoleReply{}, nil
}

func (s *DeptRole) DeleteDeptRole(c context.Context, req *deptrole.DeleteDeptRoleRequest) (*deptrole.DeleteDeptRoleReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.DeleteDeptRole(ctx, &types.DeleteDeptRoleRequest{
		DeptId: req.DeptId,
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	return &deptrole.DeleteDeptRoleReply{}, nil
}
