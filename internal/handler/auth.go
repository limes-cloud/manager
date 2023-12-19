package handler

import (
	"context"
	v1 "github.com/limes-cloud/manager/api/v1"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) Auth(ctx context.Context, in *v1.AuthRequest) (*emptypb.Empty, error) {
	return s.auth.Auth(kratosx.MustContext(ctx), in)
}

func (s *Service) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return s.auth.Login(kratosx.MustContext(ctx), in)
}

func (s *Service) Logout(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	return s.auth.Logout(kratosx.MustContext(ctx))
}

func (s *Service) LoginCaptcha(ctx context.Context, in *emptypb.Empty) (*v1.LoginCaptchaReply, error) {
	return s.auth.LoginCaptcha(kratosx.MustContext(ctx), in)
}

func (s *Service) ParseToken(ctx context.Context, in *emptypb.Empty) (*v1.ParseTokenReply, error) {
	return s.auth.ParseToken(kratosx.MustContext(ctx))
}

func (s *Service) RefreshToken(ctx context.Context, in *emptypb.Empty) (*v1.RefreshTokenReply, error) {
	return s.auth.RefreshToken(kratosx.MustContext(ctx))
}
