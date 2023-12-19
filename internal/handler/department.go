package handler

import (
	"context"
	v1 "github.com/limes-cloud/manager/api/v1"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetDepartment(ctx context.Context, in *v1.GetDepartmentRequest) (*v1.GetDepartmentReply, error) {
	return s.department.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) GetUserDepartmentTree(ctx context.Context, in *emptypb.Empty) (*v1.GetUserDepartmentTreeReply, error) {
	return s.department.UserTree(kratosx.MustContext(ctx))
}

func (s *Service) GetDepartmentTree(ctx context.Context, in *emptypb.Empty) (*v1.GetDepartmentTreeReply, error) {
	return s.department.Tree(kratosx.MustContext(ctx))
}

func (s *Service) AddDepartment(ctx context.Context, in *v1.AddDepartmentRequest) (*emptypb.Empty, error) {
	return s.department.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateDepartment(ctx context.Context, in *v1.UpdateDepartmentRequest) (*emptypb.Empty, error) {
	return s.department.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteDepartment(ctx context.Context, in *v1.DeleteDepartmentRequest) (*emptypb.Empty, error) {
	return s.department.Delete(kratosx.MustContext(ctx), in)
}
