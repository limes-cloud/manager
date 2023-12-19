package handler

import (
	"context"
	v1 "github.com/limes-cloud/manager/api/v1"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	return s.user.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) ChangeUserPasswordCaptcha(ctx context.Context, in *emptypb.Empty) (*v1.ChangeUserPasswordCaptchaReply, error) {
	return s.user.ChangePasswordCaptcha(kratosx.MustContext(ctx), in)
}

func (s *Service) ChangeUserPassword(ctx context.Context, in *v1.ChangeUserPasswordRequest) (*emptypb.Empty, error) {
	return s.user.ChangePassword(kratosx.MustContext(ctx), in)
}

func (s *Service) ResetUserPassword(ctx context.Context, in *v1.ResetUserPasswordRequest) (*emptypb.Empty, error) {
	return s.user.ResetPassword(kratosx.MustContext(ctx), in)
}

func (s *Service) CurrentUser(ctx context.Context, in *emptypb.Empty) (*v1.GetUserReply, error) {
	return s.user.Current(kratosx.MustContext(ctx))
}

func (s *Service) PageUser(ctx context.Context, in *v1.PageUserRequest) (*v1.PageUserReply, error) {
	return s.user.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) AddUser(ctx context.Context, in *v1.AddUserRequest) (*emptypb.Empty, error) {
	return s.user.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	return s.user.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateUserBasic(ctx context.Context, in *v1.UpdateUserBasicRequest) (*emptypb.Empty, error) {
	return s.user.UpdateBasic(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	return s.user.Delete(kratosx.MustContext(ctx), in)
}

func (s *Service) DisableUser(ctx context.Context, in *v1.DisableUserRequest) (*emptypb.Empty, error) {
	return s.user.Disable(kratosx.MustContext(ctx), in)
}

func (s *Service) EnableUser(ctx context.Context, in *v1.EnableUserRequest) (*emptypb.Empty, error) {
	return s.user.Enable(kratosx.MustContext(ctx), in)
}

func (s *Service) OfflineUser(ctx context.Context, in *v1.OfflineUserRequest) (*emptypb.Empty, error) {
	return s.user.Offline(kratosx.MustContext(ctx), in)
}
