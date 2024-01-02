package handler

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/manager/api/v1"
)

func (s *Service) GetUserJob(ctx context.Context, in *v1.GetUserJobRequest) (*v1.GetUserJobReply, error) {
	return s.job.GetUserJob(kratosx.MustContext(ctx), in)
}

func (s *Service) GetJob(ctx context.Context, in *v1.GetJobRequest) (*v1.GetJobReply, error) {
	return s.job.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) PageJob(ctx context.Context, in *v1.PageJobRequest) (*v1.PageJobReply, error) {
	return s.job.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) AddJob(ctx context.Context, in *v1.AddJobRequest) (*emptypb.Empty, error) {
	return s.job.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateJob(ctx context.Context, in *v1.UpdateJobRequest) (*emptypb.Empty, error) {
	return s.job.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteJob(ctx context.Context, in *v1.DeleteJobRequest) (*emptypb.Empty, error) {
	return s.job.Delete(kratosx.MustContext(ctx), in)
}
