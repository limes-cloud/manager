package app

import (
	"context"

	"github.com/limes-cloud/kratosx/model"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/manager/api/oauther"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	execer "github.com/limes-cloud/manager/internal/infra/oauther"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/types"
	"google.golang.org/protobuf/proto"
)

type OAuther struct {
	oauther.UnimplementedOAutherServer
	srv *service.OAuther
}

func NewOAuther() *OAuther {
	return &OAuther{
		srv: service.NewOAuther(dbs.NewOAuther(), execer.New()),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewOAuther()
		oauther.RegisterOAutherHTTPServer(hs, srv)
		oauther.RegisterOAutherServer(gs, srv)
	})
}

// ListOAutherType 获取登陆渠道可用列表
func (ch *OAuther) ListOAutherType(_ context.Context, _ *oauther.ListOAutherTypeRequest) (*oauther.ListOAutherTypeReply, error) {
	tps := ch.srv.ListOAutherTypes()
	reply := oauther.ListOAutherTypeReply{}
	for _, item := range tps {
		reply.List = append(reply.List, &oauther.ListOAutherTypeReply_Type{
			Keyword: item.Keyword,
			Name:    item.Name,
		})
	}
	return &reply, nil
}

// ListOAuther 获取登陆渠道列表
func (ch *OAuther) ListOAuther(c context.Context, req *oauther.ListOAutherRequest) (*oauther.ListOAutherReply, error) {
	list, total, err := ch.srv.ListOAuther(core.MustContext(c), &types.ListOAutherRequest{
		Search: &page.Search{
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
	reply := oauther.ListOAutherReply{Total: total}
	for _, item := range list {
		setting := item.GetSetting()
		reply.List = append(reply.List, &oauther.ListOAutherReply_OAuther{
			Id:      item.Id,
			Logo:    item.Logo,
			Keyword: item.Keyword,
			Type:    item.Type,
			Name:    item.Name,
			Status:  item.Status,
			Ak:      item.Ak,
			Sk:      item.Sk,
			Setting: &oauther.OAutherSetting{
				Callback: setting.Callback,
				Email: &oauther.OAutherSetting_Email{
					Host:     setting.Email.Host,
					Port:     uint32(setting.Email.Port),
					Subject:  setting.Email.Subject,
					From:     setting.Email.From,
					Template: setting.Email.Template,
				},
			},
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateOAuther 创建登陆渠道
func (ch *OAuther) CreateOAuther(c context.Context, req *oauther.CreateOAutherRequest) (*oauther.CreateOAutherReply, error) {
	id, err := ch.srv.CreateOAuther(core.MustContext(c), &entity.OAuther{
		Logo:        req.Logo,
		Keyword:     req.Keyword,
		Type:        req.Type,
		Name:        req.Name,
		Ak:          req.Ak,
		Status:      proto.Bool(false),
		Sk:          req.Sk,
		Setting:     value.ObjToString(req.Setting),
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &oauther.CreateOAutherReply{Id: id}, nil
}

// UpdateOAuther 更新登陆渠道
func (ch *OAuther) UpdateOAuther(c context.Context, req *oauther.UpdateOAutherRequest) (*oauther.UpdateOAutherReply, error) {
	ctx := core.MustContext(c)
	if err := ch.srv.UpdateOAuther(ctx, &entity.OAuther{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		Logo:            value.Value(req.Logo),
		Type:            value.Value(req.Type),
		Name:            value.Value(req.Name),
		Ak:              value.Value(req.Ak),
		Status:          req.Status,
		Sk:              value.Value(req.Sk),
		Setting:         value.ObjToString(req.Setting),
		Description:     value.Value(req.Description),
	}); err != nil {
		return nil, err
	}

	return &oauther.UpdateOAutherReply{}, nil
}

// DeleteOAuther 删除登陆渠道
func (ch *OAuther) DeleteOAuther(c context.Context, req *oauther.DeleteOAutherRequest) (*oauther.DeleteOAutherReply, error) {
	if err := ch.srv.DeleteOAuther(core.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &oauther.DeleteOAutherReply{}, nil
}
