package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	pb "github.com/limes-cloud/manager/api/manager/channel/v1"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/infra/oauth"
	"github.com/limes-cloud/manager/internal/infra/rpc"
	"github.com/limes-cloud/manager/internal/types"
)

type Channel struct {
	pb.UnimplementedChannelServer
	srv *service.Channel
}

func NewChannel(conf *conf.Config) *Channel {
	return &Channel{
		srv: service.NewChannel(conf, dbs.NewChannel(), rpc.NewFile(), oauth.New()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewChannel(c)
		pb.RegisterChannelHTTPServer(hs, srv)
		pb.RegisterChannelServer(gs, srv)
	})
}

// ListChannelType 获取登陆渠道可用列表
func (ch *Channel) ListChannelType(_ context.Context, _ *pb.ListChannelTypeRequest) (*pb.ListChannelTypeReply, error) {
	tps := ch.srv.GetTypes()
	reply := pb.ListChannelTypeReply{}
	for _, item := range tps {
		reply.List = append(reply.List, &pb.ListChannelTypeReply_Type{
			Keyword: item.Platform(),
			Name:    item.Name(),
		})
	}
	return &reply, nil
}

// ListAdminChannel 获取后台登陆渠道列表
func (ch *Channel) ListAdminChannel(c context.Context, _ *pb.ListAdminChannelRequest) (*pb.ListAdminChannelReply, error) {
	list, total, err := ch.srv.ListAdminChannel(kratosx.MustContext(c))
	if err != nil {
		return nil, err
	}
	reply := pb.ListAdminChannelReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListAdminChannelReply_Channel{
			Logo:    item.Logo,
			LogoUrl: item.LogoUrl,
			Keyword: item.Keyword,
			Name:    item.Name,
			Type:    item.Type,
		})
	}
	return &reply, nil
}

// ListChannel 获取登陆渠道列表
func (ch *Channel) ListChannel(c context.Context, req *pb.ListChannelRequest) (*pb.ListChannelReply, error) {
	list, total, err := ch.srv.ListChannel(kratosx.MustContext(c), &types.ListChannelRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Keyword:  req.Keyword,
		Name:     req.Name,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListChannelReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListChannelReply_Channel{
			Id:          item.Id,
			Logo:        item.Logo,
			LogoUrl:     item.LogoUrl,
			Keyword:     item.Keyword,
			Type:        item.Type,
			Name:        item.Name,
			Status:      item.Status,
			Admin:       item.Admin,
			Ak:          item.Ak,
			Sk:          item.Sk,
			Extra:       item.Extra,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateChannel 创建登陆渠道
func (ch *Channel) CreateChannel(c context.Context, req *pb.CreateChannelRequest) (*pb.CreateChannelReply, error) {
	id, err := ch.srv.CreateChannel(kratosx.MustContext(c), &entity.Channel{
		Logo:        req.Logo,
		Keyword:     req.Keyword,
		Type:        req.Type,
		Name:        req.Name,
		Ak:          req.Ak,
		Status:      proto.Bool(false),
		Admin:       proto.Bool(req.Admin),
		Sk:          req.Sk,
		Extra:       req.Extra,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateChannelReply{Id: id}, nil
}

// UpdateChannel 更新登陆渠道
func (ch *Channel) UpdateChannel(c context.Context, req *pb.UpdateChannelRequest) (*pb.UpdateChannelReply, error) {
	if err := ch.srv.UpdateChannel(kratosx.MustContext(c), &entity.Channel{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Logo:        req.Logo,
		Keyword:     req.Keyword,
		Type:        req.Type,
		Name:        req.Name,
		Status:      req.Status,
		Admin:       req.Admin,
		Ak:          req.Ak,
		Sk:          req.Sk,
		Extra:       req.Extra,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateChannelReply{}, nil
}

// DeleteChannel 删除登陆渠道
func (ch *Channel) DeleteChannel(c context.Context, req *pb.DeleteChannelRequest) (*pb.DeleteChannelReply, error) {
	if err := ch.srv.DeleteChannel(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteChannelReply{}, nil
}
