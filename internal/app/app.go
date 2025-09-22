package app

import (
	"context"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/manager/api/app"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type App struct {
	app.UnimplementedAppServer
	srv *service.App
}

func NewApp() *App {
	return &App{
		srv: service.NewApp(
			dbs.NewApp(),
			dbs.NewTenantApp(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewApp()
		app.RegisterAppHTTPServer(hs, srv)
		app.RegisterAppServer(gs, srv)
	})
}

// GetApp 获取指定的应用信息
func (s *App) GetApp(c context.Context, req *app.GetAppRequest) (*app.GetAppReply, error) {
	ctx := core.MustContext(c)
	data, err := s.srv.GetApp(ctx, &types.GetAppRequest{
		Id:      req.Id,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}
	reply := app.GetAppReply{}
	if err := value.Transform(data, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListApp 获取应用信息列表
func (s *App) ListCurrentApp(c context.Context, req *app.ListAppRequest) (*app.ListAppReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListAppRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list app req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := s.srv.ListCurrentApp(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := app.ListAppReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list app resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// ListApp 获取应用信息列表
func (s *App) ListApp(c context.Context, req *app.ListAppRequest) (*app.ListAppReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListAppRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list app req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := s.srv.ListApp(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := app.ListAppReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list app resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateApp 创建应用信息
func (s *App) CreateApp(c context.Context, req *app.CreateAppRequest) (*app.CreateAppReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.App{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create app req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := s.srv.CreateApp(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &app.CreateAppReply{Id: id}, nil
}

// UpdateApp 更新应用信息
func (s *App) UpdateApp(c context.Context, req *app.UpdateAppRequest) (*app.UpdateAppReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.App{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "update app req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := s.srv.UpdateApp(ctx, &in); err != nil {
		return nil, err
	}

	return &app.UpdateAppReply{}, nil
}

// DeleteApp 删除应用信息
func (s *App) DeleteApp(c context.Context, req *app.DeleteAppRequest) (*app.DeleteAppReply, error) {
	return &app.DeleteAppReply{}, s.srv.DeleteApp(core.MustContext(c), req.Id)
}

// ListAppOAuthChannel 获取应用渠道信息
func (s *App) ListAppOAuthChannel(c context.Context, req *app.ListAppOAuthChannelRequest) (*app.ListAppOAuthChannelReply, error) {
	// 调用服务
	var (
		ctx = core.MustContext(c)
		in  = types.ListAppOAuthChannelRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list app req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := s.srv.ListAppOAuthChannel(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := app.ListAppOAuthChannelReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list app resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateAppOAuthChannel 创建应用渠道信息
func (s *App) CreateAppOAuthChannel(c context.Context, req *app.CreateAppOAuthChannelRequest) (*app.CreateAppOAuthChannelReply, error) {
	// 调用服务
	id, err := s.srv.CreateAppOAuthChannel(core.MustContext(c), &entity.AppOAuthChannel{
		AppId:     req.AppId,
		ChannelId: req.ChannelId,
	})
	if err != nil {
		return nil, err
	}

	return &app.CreateAppOAuthChannelReply{Id: id}, nil
}

// DeleteAppOAuthChannel 删除应用信息
func (s *App) DeleteAppOAuthChannel(c context.Context, req *app.DeleteAppOAuthChannelRequest) (*app.DeleteAppOAuthChannelReply, error) {
	return &app.DeleteAppOAuthChannelReply{}, s.srv.DeleteAppOAuthChannel(core.MustContext(c), req.Id)
}

// ListAppField 获取应用字段信息
func (s *App) ListAppField(c context.Context, req *app.ListAppFieldRequest) (*app.ListAppFieldReply, error) {
	// 调用服务
	var (
		ctx = core.MustContext(c)
		in  = types.ListAppFieldRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list app req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := s.srv.ListAppField(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := app.ListAppFieldReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list app resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateAppField 创建应用字段信息
func (s *App) CreateAppField(c context.Context, req *app.CreateAppFieldRequest) (*app.CreateAppFieldReply, error) {
	// 调用服务
	id, err := s.srv.CreateAppField(core.MustContext(c), &entity.AppField{
		AppId:   req.AppId,
		FieldId: req.FieldId,
	})
	if err != nil {
		return nil, err
	}

	return &app.CreateAppFieldReply{Id: id}, nil
}

// DeleteAppField 删除应用信息
func (s *App) DeleteAppField(c context.Context, req *app.DeleteAppFieldRequest) (*app.DeleteAppFieldReply, error) {
	return &app.DeleteAppFieldReply{}, s.srv.DeleteAppField(core.MustContext(c), req.Id)
}
