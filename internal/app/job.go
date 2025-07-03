package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	ktypes "github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/manager/api/manager/errors"
	pb "github.com/limes-cloud/manager/api/manager/job/v1"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type Job struct {
	pb.UnimplementedJobServer
	srv *service.Job
}

func NewJob(conf *conf.Config) *Job {
	return &Job{
		srv: service.NewJob(conf, dbs.NewJob()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewJob(c)
		pb.RegisterJobHTTPServer(hs, srv)
		pb.RegisterJobServer(gs, srv)
	})
}

// ListJob 获取职位信息列表
func (s *Job) ListJob(c context.Context, req *pb.ListJobRequest) (*pb.ListJobReply, error) {
	var ctx = kratosx.MustContext(c)
	result, err := s.srv.ListJob(ctx, &types.ListJobRequest{
		Keyword: req.Keyword,
		Name:    req.Name,
		RootId:  req.RootId,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListJobReply{}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// ListCurrentJob 获取职位信息列表
func (s *Job) ListCurrentJob(c context.Context, req *pb.ListJobRequest) (*pb.ListJobReply, error) {
	var ctx = kratosx.MustContext(c)
	result, err := s.srv.ListCurrentJob(ctx, &types.ListJobRequest{
		Keyword: req.Keyword,
		Name:    req.Name,
		RootId:  req.RootId,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListJobReply{}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateJob 创建职位信息
func (s *Job) CreateJob(c context.Context, req *pb.CreateJobRequest) (*pb.CreateJobReply, error) {
	id, err := s.srv.CreateJob(kratosx.MustContext(c), &entity.Job{
		ParentId:    req.ParentId,
		Keyword:     req.Keyword,
		Name:        req.Name,
		Weight:      req.Weight,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateJobReply{Id: id}, nil
}

// UpdateJob 更新职位信息
func (s *Job) UpdateJob(c context.Context, req *pb.UpdateJobRequest) (*pb.UpdateJobReply, error) {
	if err := s.srv.UpdateJob(kratosx.MustContext(c), &entity.Job{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		ParentId:    req.ParentId,
		Keyword:     req.Keyword,
		Name:        req.Name,
		Weight:      req.Weight,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateJobReply{}, nil
}

// DeleteJob 删除职位信息
func (s *Job) DeleteJob(c context.Context, req *pb.DeleteJobRequest) (*pb.DeleteJobReply, error) {
	return &pb.DeleteJobReply{}, s.srv.DeleteJob(kratosx.MustContext(c), req.Id)
}
