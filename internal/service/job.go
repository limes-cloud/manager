package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/api/manager/errors"
	pb "github.com/limes-cloud/manager/api/manager/job/v1"
	"github.com/limes-cloud/manager/internal/biz/job"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/data"
)

type JobService struct {
	pb.UnimplementedJobServer
	uc *job.UseCase
}

func NewJobService(conf *conf.Config) *JobService {
	return &JobService{
		uc: job.NewUseCase(conf, data.NewJobRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewJobService(c)
		pb.RegisterJobHTTPServer(hs, srv)
		pb.RegisterJobServer(gs, srv)
	})
}

// GetJob 获取指定的职位信息
func (s *JobService) GetJob(c context.Context, req *pb.GetJobRequest) (*pb.GetJobReply, error) {
	var (
		in  = job.GetJobRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetJob(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetJobReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListJob 获取职位信息列表
func (s *JobService) ListJob(c context.Context, req *pb.ListJobRequest) (*pb.ListJobReply, error) {
	var (
		in  = job.ListJobRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListJob(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListJobReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateJob 创建职位信息
func (s *JobService) CreateJob(c context.Context, req *pb.CreateJobRequest) (*pb.CreateJobReply, error) {
	var (
		in  = job.Job{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateJob(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateJobReply{Id: id}, nil
}

// UpdateJob 更新职位信息
func (s *JobService) UpdateJob(c context.Context, req *pb.UpdateJobRequest) (*pb.UpdateJobReply, error) {
	var (
		in  = job.Job{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateJob(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateJobReply{}, nil
}

// DeleteJob 删除职位信息
func (s *JobService) DeleteJob(c context.Context, req *pb.DeleteJobRequest) (*pb.DeleteJobReply, error) {
	total, err := s.uc.DeleteJob(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteJobReply{Total: total}, nil
}
