package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"

	pb "github.com/limes-cloud/manager/api/manager/auth/v1"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/types"
)

type Auth struct {
	pb.UnimplementedAuthServer
	srv *service.Auth
}

func NewAuth(conf *conf.Config) *Auth {
	return &Auth{
		srv: service.NewAuth(conf),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewAuth(c)
		pb.RegisterAuthHTTPServer(hs, srv)
		pb.RegisterAuthServer(gs, srv)
	})
}

// Auth 接口鉴权
func (s *Auth) Auth(c context.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	res, err := s.srv.Auth(kratosx.MustContext(c), &types.AuthRequest{
		Path:   req.Path,
		Method: req.Method,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AuthReply{
		UserId:            res.UserId,
		UserName:          res.UserName,
		RoleId:            res.RoleId,
		RoleKeyword:       res.RoleKeyword,
		DepartmentId:      res.DepartmentId,
		DepartmentKeyword: res.DepartmentKeyword,
	}, nil
}
