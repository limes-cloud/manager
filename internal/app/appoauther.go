package app

import (
	"context"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/api/errors"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/appoauther"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type AppOAuther struct {
	appoauther.UnimplementedAppOAutherServer
	srv *service.AppOAuther
}

func NewAppOAuther() *AppOAuther {
	return &AppOAuther{
		srv: service.NewAppOAuther(
			dbs.NewAppOAuther(),
			dbs.NewApp(),
			dbs.NewTenant(),
			dbs.NewOAuther(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewAppOAuther()
		appoauther.RegisterAppOAutherHTTPServer(hs, srv)
		appoauther.RegisterAppOAutherServer(gs, srv)
	})
}

// ListAppOAuther 获取应用渠道信息
func (s *AppOAuther) ListAppOAuther(c context.Context, req *appoauther.ListAppOAutherRequest) (*appoauther.ListAppOAutherReply, error) {
	// 调用服务
	ctx := core.MustContext(c)

	// 调用服务
	list, total, err := s.srv.ListAppOAuther(ctx, &types.ListAppOAutherRequest{
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
	reply := appoauther.ListAppOAutherReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list appoauther resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateAppOAuther 创建应用渠道信息
func (s *AppOAuther) CreateAppOAuther(c context.Context, req *appoauther.CreateAppOAutherRequest) (*appoauther.CreateAppOAutherReply, error) {
	// 调用服务
	id, err := s.srv.CreateAppOAuther(core.MustContext(c), &entity.AppOAuther{
		AppId:     req.AppId,
		OAutherId: req.OautherId,
	})
	if err != nil {
		return nil, err
	}

	return &appoauther.CreateAppOAutherReply{Id: id}, nil
}

// DeleteAppOAuther 删除应用信息
func (s *AppOAuther) DeleteAppOAuther(c context.Context, req *appoauther.DeleteAppOAutherRequest) (*appoauther.DeleteAppOAutherReply, error) {
	return &appoauther.DeleteAppOAutherReply{}, s.srv.DeleteAppOAuther(core.MustContext(c), req.Id)
}
