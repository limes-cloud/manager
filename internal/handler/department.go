package handler

import (
	"context"
	v1 "manager/api/v1"

	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetDepartment(ctx context.Context, in *v1.GetDepartmentRequest) (*v1.GetDepartmentReply, error) {
	return s.department.Get(kratos.MustContext(ctx), in)
}

func (s *Service) GetDepartmentTree(ctx context.Context, in *emptypb.Empty) (*v1.GetDepartmentTreeReply, error) {
	return s.department.Tree(kratos.MustContext(ctx))
}

func (s *Service) AddDepartment(ctx context.Context, in *v1.AddDepartmentRequest) (*emptypb.Empty, error) {
	return s.department.Add(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateDepartment(ctx context.Context, in *v1.UpdateDepartmentRequest) (*emptypb.Empty, error) {
	return s.department.Update(kratos.MustContext(ctx), in)
}

func (s *Service) DeleteDepartment(ctx context.Context, in *v1.DeleteDepartmentRequest) (*emptypb.Empty, error) {
	return s.department.Delete(kratos.MustContext(ctx), in)
}
