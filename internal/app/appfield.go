package app

import (
	"context"

	"github.com/limes-cloud/kratosx/model"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/api/errors"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/appfield"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type AppField struct {
	appfield.UnimplementedAppFieldServer
	srv *service.AppField
}

func NewAppField() *AppField {
	return &AppField{
		srv: service.NewAppField(
			dbs.NewAppField(),
			dbs.NewApp(),
			dbs.NewTenant(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewAppField()
		appfield.RegisterAppFieldHTTPServer(hs, srv)
		appfield.RegisterAppFieldServer(gs, srv)
	})
}

// ListAppField 获取应用渠道信息
func (s *AppField) ListAppField(c context.Context, req *appfield.ListAppFieldRequest) (*appfield.ListAppFieldReply, error) {
	// 调用服务
	ctx := core.MustContext(c)

	// 调用服务
	list, total, err := s.srv.ListAppField(ctx, &types.ListAppFieldRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
			Order:    req.Order,
			OrderBy:  req.OrderBy,
		},
		Keyword: req.Keyword,
		Name:    req.Name,
		AppId:   req.AppId,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := appfield.ListAppFieldReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list appfield resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateAppField 创建应用渠道信息
func (s *AppField) CreateAppField(c context.Context, req *appfield.CreateAppFieldRequest) (*appfield.CreateAppFieldReply, error) {
	// 调用服务
	id, err := s.srv.CreateAppField(core.MustContext(c), &entity.AppField{
		AppId:    req.AppId,
		FieldId:  req.FieldId,
		Required: &req.Required,
	})
	if err != nil {
		return nil, err
	}

	return &appfield.CreateAppFieldReply{Id: id}, nil
}

// UpdateAppField 更新应用渠道信息
func (s *AppField) UpdateAppField(c context.Context, req *appfield.UpdateAppFieldRequest) (*appfield.UpdateAppFieldReply, error) {
	// 调用服务
	err := s.srv.UpdateAppField(core.MustContext(c), &entity.AppField{
		CreateTenantModel: model.CreateTenantModel{Id: req.Id},
		Required:          &req.Required,
	})
	if err != nil {
		return nil, err
	}

	return &appfield.UpdateAppFieldReply{}, nil
}

// DeleteAppField 删除应用信息
func (s *AppField) DeleteAppField(c context.Context, req *appfield.DeleteAppFieldRequest) (*appfield.DeleteAppFieldReply, error) {
	return &appfield.DeleteAppFieldReply{}, s.srv.DeleteAppField(core.MustContext(c), req.Id)
}
