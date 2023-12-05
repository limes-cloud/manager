package handler

import (
	"context"
	v1 "manager/api/v1"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetMenuTree(ctx context.Context, in *emptypb.Empty) (*v1.GetMenuTreeReply, error) {
	return s.menu.Tree(kratosx.MustContext(ctx))
}

func (s *Service) AddMenu(ctx context.Context, in *v1.AddMenuRequest) (*emptypb.Empty, error) {
	return s.menu.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateMenu(ctx context.Context, in *v1.UpdateMenuRequest) (*emptypb.Empty, error) {
	return s.menu.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteMenu(ctx context.Context, in *v1.DeleteMenuRequest) (*emptypb.Empty, error) {
	return s.menu.Delete(kratosx.MustContext(ctx), in)
}
