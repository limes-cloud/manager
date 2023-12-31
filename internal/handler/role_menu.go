package handler

import (
	"context"
	v1 "github.com/limes-cloud/manager/api/v1"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CurrentRoleMenuTree(ctx context.Context, in *emptypb.Empty) (*v1.CurrentRoleMenuTreeReply, error) {
	return s.roleMenu.CurrenMenuTree(kratosx.MustContext(ctx))
}

func (s *Service) UpdateRoleMenus(ctx context.Context, in *v1.UpdateRoleMenuRequest) (*emptypb.Empty, error) {
	return s.roleMenu.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) GetRoleMenuIds(ctx context.Context, in *v1.GetRoleMenuIdsRequest) (*v1.GetRoleMenuIdsReply, error) {
	return s.roleMenu.GetMenuIds(kratosx.MustContext(ctx), in)
}
