package service

import (
	"context"

	"github.com/limes-cloud/kratosx"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "github.com/limes-cloud/manager/api/manager/auth/v1"
	"github.com/limes-cloud/manager/internal/biz/auth"
	"github.com/limes-cloud/manager/internal/conf"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	uc *auth.UseCase
}

func NewAuthService(conf *conf.Config) *AuthService {
	return &AuthService{
		uc: auth.NewUseCase(conf),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewAuthService(c)
		pb.RegisterAuthHTTPServer(hs, srv)
		pb.RegisterAuthServer(gs, srv)
	})
}

// Auth 接口鉴权
func (s *AuthService) Auth(c context.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	res, err := s.uc.Auth(kratosx.MustContext(c), &auth.AuthRequest{
		Path:   req.Path,
		Method: req.Method,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AuthReply{
		UserId:            res.UserId,
		RoleId:            res.RoleId,
		RoleKeyword:       res.RoleKeyword,
		DepartmentId:      res.DepartmentId,
		DepartmentKeyword: res.DepartmentKeyword,
	}, nil
}
