package app

import (
	"context"

	"github.com/limes-cloud/manager/internal/domain/entity"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/manager/api/job"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
)

type Job struct {
	job.UnimplementedJobServer
	srv *service.Job
}

// NewJob 初始化租户应用
func NewJob() *Job {
	return &Job{
		srv: service.NewJob(
			dbs.NewJob(),
		),
	}
}

// init 应用注册
func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewJob()
		job.RegisterJobHTTPServer(hs, srv)
		job.RegisterJobServer(gs, srv)
	})
}

// ListJob 获取租户信息列表
func (app *Job) ListJob(c context.Context, req *job.ListJobRequest) (*job.ListJobReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListJobRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list job req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := app.srv.ListJob(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := job.ListJobReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get job resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateJob 创建租户信息
func (app *Job) CreateJob(c context.Context, req *job.CreateJobRequest) (*job.CreateJobReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Job{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create job req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateJob(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &job.CreateJobReply{Id: id}, nil
}

// UpdateJob 更新租户信息
func (app *Job) UpdateJob(c context.Context, req *job.UpdateJobRequest) (*job.UpdateJobReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Job{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create job req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateJob(ctx, &in); err != nil {
		return nil, err
	}

	return &job.UpdateJobReply{}, nil
}

// DeleteJob 删除租户信息
func (app *Job) DeleteJob(c context.Context, req *job.DeleteJobRequest) (*job.DeleteJobReply, error) {
	return &job.DeleteJobReply{}, app.srv.DeleteJob(core.MustContext(c), req.Id)
}
