package app

import (
	"context"
	"encoding/json"

	"github.com/limes-cloud/kratosx/model"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/app"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type App struct {
	app.UnimplementedAppServer
	srv *service.App
}

func NewApp() *App {
	return &App{
		srv: service.NewApp(
			dbs.NewApp(),
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

// GetAppByKeyword 获取指定的应用信息
func (s *App) GetAppByKeyword(c context.Context, keyword string) (*entity.App, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	data, err := s.srv.GetAppByKeyword(ctx, keyword)
	return data, err
}

// GetSampleApp 获取指定的应用信息
func (s *App) GetSampleApp(c context.Context, req *app.GetSampleAppRequest) (*app.GetSampleAppReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	data, err := s.srv.GetAppByKeyword(ctx, req.Keyword)
	if err != nil {
		return nil, err
	}
	reply := app.GetSampleAppReply{
		Id:          data.Id,
		Type:        data.Type,
		Logo:        data.Logo,
		Favicon:     data.Favicon,
		Keyword:     data.Keyword,
		Name:        data.Name,
		ShowName:    data.ShowName,
		Status:      value.Value(data.Status),
		Reason:      data.Reason,
		Private:     value.Value(data.Private),
		Description: data.Description,
		Comment:     data.Comment,
		CreatedAt:   uint32(data.CreatedAt),
		UpdatedAt:   uint32(data.UpdatedAt),
	}
	_ = json.Unmarshal([]byte(data.Setting), &reply.Setting)
	return &reply, nil
}

// GetApp 获取指定的应用信息
func (s *App) GetApp(c context.Context, req *app.GetAppRequest) (*app.GetAppReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	data, err := s.srv.GetApp(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	reply := app.GetAppReply{
		Id:          data.Id,
		Type:        data.Type,
		Logo:        data.Logo,
		Favicon:     data.Favicon,
		Keyword:     data.Keyword,
		Secret:      data.Secret,
		Name:        data.Name,
		ShowName:    data.ShowName,
		Status:      value.Value(data.Status),
		Reason:      data.Reason,
		Private:     value.Value(data.Private),
		Description: data.Description,
		Comment:     data.Comment,
		CreatedAt:   uint32(data.CreatedAt),
		UpdatedAt:   uint32(data.UpdatedAt),
	}
	_ = json.Unmarshal([]byte(data.Setting), &reply.Setting)
	return &reply, nil
}

// ListApp 获取应用信息列表
func (s *App) ListApp(c context.Context, req *app.ListAppRequest) (*app.ListAppReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, total, err := s.srv.ListApp(ctx, &types.ListAppRequest{
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

	// 处理返回数据
	reply := app.ListAppReply{Total: total}
	for _, ent := range list {
		reply.List = append(reply.List, &app.ListAppReply_Data{
			Id:          ent.Id,
			Logo:        ent.Logo,
			Favicon:     ent.Favicon,
			Keyword:     ent.Keyword,
			Name:        ent.Name,
			ShowName:    ent.ShowName,
			Status:      ent.Status,
			Reason:      ent.Reason,
			Private:     ent.Private,
			Description: ent.Description,
			Comment:     ent.Comment,
			CreatedAt:   uint32(ent.CreatedAt),
			UpdatedAt:   uint32(ent.UpdatedAt),
		})
	}

	return &reply, nil
}

// CreateApp 创建应用信息
func (s *App) CreateApp(c context.Context, req *app.CreateAppRequest) (*app.CreateAppReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	id, err := s.srv.CreateApp(ctx, &entity.App{
		Type:        req.Type,
		Logo:        req.Logo,
		Favicon:     req.Favicon,
		Keyword:     req.Keyword,
		Secret:      req.Secret,
		Name:        req.Name,
		ShowName:    req.ShowName,
		Status:      value.Pointer(false),
		Private:     value.Pointer(false),
		Reason:      value.Pointer("应用初始化"),
		Description: req.Description,
		Comment:     req.Comment,
		Setting:     value.ObjToString(req.Setting),
	})
	if err != nil {
		return nil, err
	}

	return &app.CreateAppReply{Id: id}, nil
}

// UpdateApp 更新应用信息
func (s *App) UpdateApp(c context.Context, req *app.UpdateAppRequest) (*app.UpdateAppReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	if err := s.srv.UpdateApp(ctx, &entity.App{
		BaseModel:   model.BaseModel{Id: req.Id},
		Type:        value.Value(req.Type),
		Logo:        value.Value(req.Logo),
		Favicon:     value.Value(req.Favicon),
		Secret:      value.Value(req.Secret),
		Name:        value.Value(req.Name),
		ShowName:    value.Value(req.ShowName),
		Status:      req.Status,
		Private:     req.Private,
		Reason:      req.Reason,
		Description: req.Description,
		Comment:     req.Comment,
		Setting:     value.ObjToString(req.Setting),
	}); err != nil {
		return nil, err
	}

	return &app.UpdateAppReply{}, nil
}

// DeleteApp 删除应用信息
func (s *App) DeleteApp(c context.Context, req *app.DeleteAppRequest) (*app.DeleteAppReply, error) {
	return &app.DeleteAppReply{}, s.srv.DeleteApp(core.MustContext(c), req.Id)
}
