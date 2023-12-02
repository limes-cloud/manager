package handler

import (
	"context"
	v1 "manager/api/v1"

	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	return s.user.Get(kratos.MustContext(ctx), in)
}

func (s *Service) ChangeUserPasswordCaptcha(ctx context.Context, in *emptypb.Empty) (*v1.ChangeUserPasswordCaptchaReply, error) {
	return s.user.ChangePasswordCaptcha(kratos.MustContext(ctx), in)
}

func (s *Service) ChangeUserPassword(ctx context.Context, in *v1.ChangeUserPasswordRequest) (*emptypb.Empty, error) {
	return s.user.ChangePassword(kratos.MustContext(ctx), in)
}

func (s *Service) ResetUserPassword(ctx context.Context, in *v1.ResetUserPasswordRequest) (*emptypb.Empty, error) {
	return s.user.ResetPassword(kratos.MustContext(ctx), in)
}

func (s *Service) CurrentUser(ctx context.Context, in *emptypb.Empty) (*v1.GetUserReply, error) {
	return s.user.Current(kratos.MustContext(ctx))
}

func (s *Service) PageUser(ctx context.Context, in *v1.PageUserRequest) (*v1.PageUserReply, error) {
	return s.user.Page(kratos.MustContext(ctx), in)
}

func (s *Service) AddUser(ctx context.Context, in *v1.AddUserRequest) (*emptypb.Empty, error) {
	return s.user.Add(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	return s.user.Update(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateUserBasic(ctx context.Context, in *v1.UpdateUserBasicRequest) (*emptypb.Empty, error) {
	return s.user.UpdateBasic(kratos.MustContext(ctx), in)
}

func (s *Service) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	return s.user.Delete(kratos.MustContext(ctx), in)
}

func (s *Service) DisableUser(ctx context.Context, in *v1.DisableUserRequest) (*emptypb.Empty, error) {
	return s.user.Disable(kratos.MustContext(ctx), in)
}

func (s *Service) EnableUser(ctx context.Context, in *v1.EnableUserRequest) (*emptypb.Empty, error) {
	return s.user.Enable(kratos.MustContext(ctx), in)
}

func (s *Service) OfflineUser(ctx context.Context, in *v1.OfflineUserRequest) (*emptypb.Empty, error) {
	return s.user.Offline(kratos.MustContext(ctx), in)
}
