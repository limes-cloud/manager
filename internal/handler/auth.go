package handler

import (
	"context"
	v1 "manager/api/v1"

	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) Auth(ctx context.Context, in *v1.AuthRequest) (*emptypb.Empty, error) {
	return s.auth.Auth(kratos.MustContext(ctx), in)
}

func (s *Service) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return s.auth.Login(kratos.MustContext(ctx), in)
}

func (s *Service) Logout(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	return s.auth.Logout(kratos.MustContext(ctx))
}

func (s *Service) LoginCaptcha(ctx context.Context, in *emptypb.Empty) (*v1.LoginCaptchaReply, error) {
	return s.auth.LoginCaptcha(kratos.MustContext(ctx), in)
}

func (s *Service) ParseToken(ctx context.Context, in *emptypb.Empty) (*v1.ParseTokenReply, error) {
	return s.auth.ParseToken(kratos.MustContext(ctx))
}

func (s *Service) RefreshToken(ctx context.Context, in *emptypb.Empty) (*v1.RefreshTokenReply, error) {
	return s.auth.RefreshToken(kratos.MustContext(ctx))
}
