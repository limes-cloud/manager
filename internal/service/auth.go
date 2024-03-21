package service

import (
	"context"

	"github.com/limes-cloud/kratosx"

	v1 "github.com/limes-cloud/manager/api/auth/v1"
	biz "github.com/limes-cloud/manager/internal/biz/auth"
	"github.com/limes-cloud/manager/internal/config"
	"github.com/limes-cloud/manager/internal/data/department"
	"github.com/limes-cloud/manager/internal/data/menu"
)

type AuthService struct {
	v1.UnimplementedServiceServer
	uc   *biz.UseCase
	conf *config.Config
}

func NewAuthService(conf *config.Config) *AuthService {
	return &AuthService{
		conf: conf,
		uc:   biz.NewUseCase(conf, department.NewRepo(), menu.NewRepo()),
	}
}

func (s AuthService) Auth(ctx context.Context, in *v1.AuthRequest) (*v1.AuthReply, error) {
	return s.uc.Auth(kratosx.MustContext(ctx), in)
}
