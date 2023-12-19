package handler

import (
	"context"
	v1 "github.com/limes-cloud/manager/api/v1"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetRole(ctx context.Context, in *v1.GetRoleRequest) (*v1.GetRoleReply, error) {
	return s.role.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) GetRoleTree(ctx context.Context, in *emptypb.Empty) (*v1.GetRoleTreeReply, error) {
	return s.role.Tree(kratosx.MustContext(ctx))
}

func (s *Service) AddRole(ctx context.Context, in *v1.AddRoleRequest) (*emptypb.Empty, error) {
	return s.role.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateRole(ctx context.Context, in *v1.UpdateRoleRequest) (*emptypb.Empty, error) {
	return s.role.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteRole(ctx context.Context, in *v1.DeleteRoleRequest) (*emptypb.Empty, error) {
	return s.role.Delete(kratosx.MustContext(ctx), in)
}
