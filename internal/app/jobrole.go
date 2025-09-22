package app

import (
	"context"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
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

func (s *JobRole) ListRoleJob(c context.Context, req *jobrole.ListRoleJobRequest) (*jobrole.ListRoleJobReply, error) {
	ctx := core.MustContext(c)
	list, total, err := s.srv.ListRoleJob(ctx, &types.ListRoleJobRequest{
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
	reply := jobrole.ListRoleJobReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list role job \resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *JobRole) CreateRoleJobs(c context.Context, req *jobrole.CreateRoleJobsRequest) (*jobrole.CreateRoleJobsReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.CreateRoleJobs(ctx, &types.CreateRoleJobsRequest{
		RoleId: req.RoleId,
		JobIds: req.JobIds,
	})
	if err != nil {
		return nil, err
	}
	return &jobrole.CreateRoleJobsReply{}, nil
}

func (s *JobRole) CreateJobRoles(c context.Context, req *jobrole.CreateJobRolesRequest) (*jobrole.CreateJobRolesReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.CreateJobRoles(ctx, &types.CreateJobRolesRequest{
		JobId:   req.JobId,
		RoleIds: req.RoleIds,
	})
	if err != nil {
		return nil, err
	}
	return &jobrole.CreateJobRolesReply{}, nil
}

func (s *JobRole) DeleteJobRoles(c context.Context, req *jobrole.DeleteJobRolesRequest) (*jobrole.DeleteJobRolesReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.DeleteJobRoles(ctx, &types.DeleteJobRolesRequest{
		JobId:   req.JobId,
		RoleIds: req.RoleIds,
	})
	if err != nil {
		return nil, err
	}
	return &jobrole.DeleteJobRolesReply{}, nil
}

func (s *JobRole) DeleteRoleJobs(c context.Context, req *jobrole.DeleteRoleJobsRequest) (*jobrole.DeleteRoleJobsReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.DeleteRoleJobs(ctx, &types.DeleteRoleJobsRequest{
		JobIds: req.JobIds,
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	return &jobrole.DeleteRoleJobsReply{}, nil
}
