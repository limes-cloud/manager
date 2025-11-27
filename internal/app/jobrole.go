package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/api/jobrole"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type JobRole struct {
	jobrole.UnimplementedJobRoleServer
	srv *service.JobRole
}

func NewJobRole() *JobRole {
	return &JobRole{
		srv: service.NewJobRole(
			dbs.NewJobRole(),
			dbs.NewScope(),
			dbs.NewTenantAdmin(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewJobRole()
		jobrole.RegisterJobRoleHTTPServer(hs, srv)
		jobrole.RegisterJobRoleServer(gs, srv)
	})
}

// ListJobRole 获取角色菜单列表
func (s *JobRole) ListJobRole(c context.Context, req *jobrole.ListJobRoleRequest) (*jobrole.ListJobRoleReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, total, err := s.srv.ListJobRole(ctx, &types.ListJobRoleRequest{
		Search: page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		JobId: req.JobId,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := jobrole.ListJobRoleReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list job role resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

func (s *JobRole) CreateJobRole(c context.Context, req *jobrole.CreateJobRoleRequest) (*jobrole.CreateJobRoleReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.CreateJobRole(ctx, &types.CreateJobRoleRequest{
		JobId:  req.JobId,
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	return &jobrole.CreateJobRoleReply{}, nil
}

func (s *JobRole) DeleteJobRole(c context.Context, req *jobrole.DeleteJobRoleRequest) (*jobrole.DeleteJobRoleReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.DeleteJobRole(ctx, &types.DeleteJobRoleRequest{
		JobId:  req.JobId,
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	return &jobrole.DeleteJobRoleReply{}, nil
}
