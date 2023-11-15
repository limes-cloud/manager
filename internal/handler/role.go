package handler

import (
	"context"
	v1 "manager/api/v1"

	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetRole(ctx context.Context, in *v1.GetRoleRequest) (*v1.GetRoleReply, error) {
	return s.role.Get(kratos.MustContext(ctx), in)
}

func (s *Service) GetRoleTree(ctx context.Context, in *emptypb.Empty) (*v1.GetRoleTreeReply, error) {
	return s.role.Tree(kratos.MustContext(ctx))
}

func (s *Service) AddRole(ctx context.Context, in *v1.AddRoleRequest) (*emptypb.Empty, error) {
	return s.role.Add(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateRole(ctx context.Context, in *v1.UpdateRoleRequest) (*emptypb.Empty, error) {
	return s.role.Update(kratos.MustContext(ctx), in)
}

func (s *Service) DeleteRole(ctx context.Context, in *v1.DeleteRoleRequest) (*emptypb.Empty, error) {
	return s.role.Delete(kratos.MustContext(ctx), in)
}
