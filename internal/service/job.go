package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/manager/api/errors"
	v1 "github.com/limes-cloud/manager/api/job/v1"
	biz "github.com/limes-cloud/manager/internal/biz/job"
	"github.com/limes-cloud/manager/internal/config"
	data "github.com/limes-cloud/manager/internal/data/job"
)

type JobService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase

	conf *config.Config
}

func NewJobService(conf *config.Config) *JobService {
	return &JobService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (o JobService) GetJob(ctx context.Context, request *v1.GetJobRequest) (*v1.Job, error) {
	obj, err := o.uc.GetJobById(kratosx.MustContext(ctx), request.Id)
	if err != nil {
		return nil, err
	}

	reply := v1.Job{}
	if err := util.Transform(obj, &reply); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (o JobService) PageJob(ctx context.Context, request *v1.PageJobRequest) (*v1.PageJobReply, error) {
	var in biz.PageJobRequest
	if err := util.Transform(request, &in); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := o.uc.PageJob(kratosx.MustContext(ctx), &in)
	if err != nil {
		return nil, err
	}

	reply := v1.PageJobReply{Total: total}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (o JobService) AddJob(ctx context.Context, request *v1.AddJobRequest) (*v1.AddJobReply, error) {
	var job biz.Job
	if err := util.Transform(request, &job); err != nil {
		return nil, errors.Transform()
	}

	id, err := o.uc.AddJob(kratosx.MustContext(ctx), &job)
	if err != nil {
		return nil, err
	}

	return &v1.AddJobReply{Id: id}, nil
}

func (o JobService) UpdateJob(ctx context.Context, request *v1.UpdateJobRequest) (*emptypb.Empty, error) {
	var job biz.Job
	if err := util.Transform(request, &job); err != nil {
		return nil, errors.Transform()
	}

	return nil, o.uc.UpdateJob(kratosx.MustContext(ctx), &job)
}

func (o JobService) DeleteJob(ctx context.Context, request *v1.DeleteJobRequest) (*emptypb.Empty, error) {
	return nil, o.uc.DeleteJob(kratosx.MustContext(ctx), request.Id)
}
