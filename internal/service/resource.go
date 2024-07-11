package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"

	pb "github.com/limes-cloud/manager/api/manager/resource/v1"
	"github.com/limes-cloud/manager/internal/biz/resource"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/data"
)

type ResourceService struct {
	pb.UnimplementedResourceServer
	uc *resource.UseCase
}

func NewResourceService(conf *conf.Config) *ResourceService {
	return &ResourceService{
		uc: resource.NewUseCase(conf, data.NewResourceRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewResourceService(c)
		pb.RegisterResourceHTTPServer(hs, srv)
		pb.RegisterResourceServer(gs, srv)
	})
}

// GetResourceScopes 获取当前用户的资源权限
func (r ResourceService) GetResourceScopes(c context.Context, req *pb.GetResourceScopesRequest) (*pb.GetResourceScopesReply, error) {
	all, ids, err := r.uc.GetResourceScopes(kratosx.MustContext(c), req.UserId, req.Keyword)
	if err != nil {
		return nil, err
	}
	return &pb.GetResourceScopesReply{All: all, Scopes: ids}, nil
}

// GetResource 获取指定用户的资源权限
func (r ResourceService) GetResource(c context.Context, req *pb.GetResourceRequest) (*pb.GetResourceReply, error) {
	ids, err := r.uc.GetResource(kratosx.MustContext(c), &resource.GetResourceRequest{
		Keyword:    req.Keyword,
		ResourceId: req.ResourceId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetResourceReply{DepartmentIds: ids}, nil
}

// UpdateResource 更新用户的资源权限
func (r ResourceService) UpdateResource(c context.Context, req *pb.UpdateResourceRequest) (*pb.UpdateResourceReply, error) {
	if err := r.uc.UpdateResource(kratosx.MustContext(c), &resource.UpdateResourceRequest{
		Keyword:       req.Keyword,
		ResourceId:    req.ResourceId,
		DepartmentIds: req.DepartmentIds,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateResourceReply{}, nil
}
