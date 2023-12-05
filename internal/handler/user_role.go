package handler

import (
	"context"
	v1 "manager/api/v1"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CurrentUserRoles(ctx context.Context, in *emptypb.Empty) (*v1.GetUserRolesReply, error) {
	return s.userRole.CurrentUserRoles(kratosx.MustContext(ctx))
}

func (s *Service) GetUserRoles(ctx context.Context, in *v1.GetUserRolesRequest) (*v1.GetUserRolesReply, error) {
	return s.userRole.GetUserRoles(kratosx.MustContext(ctx), in)
}

func (s *Service) SwitchCurrentUserRole(ctx context.Context, in *v1.SwitchCurrentUserRoleRequest) (*v1.SwitchCurrentUserRoleReply, error) {
	return s.userRole.SwitchCurrentUserRole(kratosx.MustContext(ctx), in)
}
