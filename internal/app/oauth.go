package app

import (
	"context"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/api/oauth"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	oauth2 "github.com/limes-cloud/manager/internal/infra/oauth"
	"github.com/limes-cloud/manager/internal/types"
)

type OAuth struct {
	oauth.UnimplementedOAuthServer
	srv *service.OAuth
}

func NewOAuth() *OAuth {
	return &OAuth{
		srv: service.NewOAuth(oauth2.New()),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewOAuth()
		oauth.RegisterOAuthHTTPServer(hs, srv)
		oauth.RegisterOAuthServer(gs, srv)
	})
}

// ListOAuthChannelType 获取登陆渠道可用列表
func (ch *OAuth) ListOAuthChannelType(_ context.Context, _ *oauth.ListOAuthChannelTypeRequest) (*oauth.ListOAuthChannelTypeReply, error) {
	tps := ch.srv.GetTypes()
	reply := oauth.ListOAuthChannelTypeReply{}
	for _, item := range tps {
		reply.List = append(reply.List, &oauth.ListOAuthChannelTypeReply_Type{
			Keyword: item.Platform(),
			Name:    item.Name(),
		})
	}
	return &reply, nil
}

// ListOAuthChannel 获取登陆渠道列表
func (ch *OAuth) ListOAuthChannel(c context.Context, req *oauth.ListOAuthChannelRequest) (*oauth.ListOAuthChannelReply, error) {
	list, total, err := ch.srv.ListOAuthChannel(core.MustContext(c), &types.ListOAuthChannelRequest{
		Search: page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
			Order:    req.Order,
			OrderBy:  req.OrderBy,
		},
		Keyword: req.Keyword,
		Name:    req.Name,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}
	reply := oauth.ListOAuthChannelReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &oauth.ListOAuthChannelReply_OAuthChannel{
			Id:          item.Id,
			Logo:        item.Logo,
			Keyword:     item.Keyword,
			Type:        item.Type,
			Name:        item.Name,
			Status:      item.Status,
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

// CreateOAuthChannel 创建登陆渠道
func (ch *OAuth) CreateOAuthChannel(c context.Context, req *oauth.CreateOAuthChannelRequest) (*oauth.CreateOAuthChannelReply, error) {
	id, err := ch.srv.CreateOAuthChannel(core.MustContext(c), &entity.OAuthChannel{
		Logo:        req.Logo,
		Keyword:     req.Keyword,
		Type:        req.Type,
		Name:        req.Name,
		Ak:          req.Ak,
		Status:      proto.Bool(false),
		Sk:          req.Sk,
		Extra:       req.Extra,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &oauth.CreateOAuthChannelReply{Id: id}, nil
}

// UpdateOAuthChannel 更新登陆渠道
func (ch *OAuth) UpdateOAuthChannel(c context.Context, req *oauth.UpdateOAuthChannelRequest) (*oauth.UpdateOAuthChannelReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.OAuthChannel{}
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "transform error", "err", err)
		return nil, errors.TransformError()
	}

	if err := ch.srv.UpdateOAuthChannel(ctx, &in); err != nil {
		return nil, err
	}

	return &oauth.UpdateOAuthChannelReply{}, nil
}

// DeleteOAuthChannel 删除登陆渠道
func (ch *OAuth) DeleteOAuthChannel(c context.Context, req *oauth.DeleteOAuthChannelRequest) (*oauth.DeleteOAuthChannelReply, error) {
	if err := ch.srv.DeleteOAuthChannel(core.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &oauth.DeleteOAuthChannelReply{}, nil
}
