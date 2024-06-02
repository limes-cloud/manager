package service

import (
	"context"

	"github.com/limes-cloud/kratosx"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "github.com/limes-cloud/manager/api/manager/resource/v1"
	"github.com/limes-cloud/manager/internal/biz/resource"
	"github.com/limes-cloud/manager/internal/conf"
)

type ResourceService struct {
	pb.UnimplementedResourceServer
	uc *resource.UseCase
}

func NewResourceService(conf *conf.Config) *ResourceService {
	return &ResourceService{
		uc: resource.NewUseCase(conf),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewResourceService(c)
		pb.RegisterResourceHTTPServer(hs, srv)
		pb.RegisterResourceServer(gs, srv)
	})
}

// GetResourceScopes 获取指定用户的资源权限
func (r ResourceService) GetResourceScopes(c context.Context, req *pb.GetResourceScopesRequest) (*pb.GetResourceScopesReply, error) {
	all, scopes, err := r.uc.GetResourceScopes(kratosx.MustContext(c), req.Keyword)
	if err != nil {
		return nil, err
	}
	return &pb.GetResourceScopesReply{All: all, Scopes: scopes}, nil
}

// UpdateResourceScopes 更新用户的资源权限
func (r ResourceService) UpdateResourceScopes(c context.Context, req *pb.UpdateResourceScopesRequest) (*pb.UpdateResourceScopesReply, error) {
	if err := r.uc.UpdateResourceScopes(kratosx.MustContext(c), &resource.UpdateResourceScopesRequest{
		Keyword:    req.Keyword,
		Scopes:     req.Scopes,
		ResourceId: req.ResourceId,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateResourceScopesReply{}, nil
}
