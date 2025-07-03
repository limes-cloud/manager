package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "github.com/limes-cloud/manager/api/manager/app/v1"
	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/infra/rpc"
	"github.com/limes-cloud/manager/internal/types"
)

type App struct {
	pb.UnimplementedAppServer
	srv *service.App
}

func NewApp(conf *conf.Config) *App {
	return &App{
		srv: service.NewApp(
			conf,
			dbs.NewApp(),
			service.NewResource(conf, dbs.NewResource(), dbs.NewDepartment(), dbs.NewRole()),
			rpc.NewFile(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewApp(c)
		pb.RegisterAppHTTPServer(hs, srv)
		pb.RegisterAppServer(gs, srv)
	})
}

// GetApp 获取指定的应用信息
func (s *App) GetApp(c context.Context, req *pb.GetAppRequest) (*pb.GetAppReply, error) {
	result, err := s.srv.GetApp(kratosx.MustContext(c), &types.GetAppRequest{
		Id:      req.Id,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.GetAppReply{}
	if err := valx.Transform(result, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListApp 获取应用信息列表
func (s *App) ListApp(c context.Context, req *pb.ListAppRequest) (*pb.ListAppReply, error) {
	list, total, err := s.srv.ListApp(kratosx.MustContext(c), &types.ListAppRequest{
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

	reply := pb.ListAppReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListAppReply_App{
			Id:            item.Id,
			Logo:          item.Logo,
			LogoUrl:       item.LogoUrl,
			Keyword:       item.Keyword,
			Name:          item.Name,
			Status:        item.Status,
			DisableDesc:   item.DisableDesc,
			AllowRegistry: item.AllowRegistry,
			Version:       item.Version,
			Copyright:     item.Copyright,
			Extra:         item.Extra,
			Description:   item.Description,
			CreatedAt:     uint32(item.CreatedAt),
			UpdatedAt:     uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateApp 创建应用信息
func (s *App) CreateApp(c context.Context, req *pb.CreateAppRequest) (*pb.CreateAppReply, error) {
	app := entity.App{
		Logo:          req.Logo,
		Keyword:       req.Keyword,
		Name:          req.Name,
		AllowRegistry: req.AllowRegistry,
		Version:       req.Version,
		Copyright:     req.Copyright,
		Extra:         req.Extra,
		Description:   req.Description,
	}

	for _, item := range req.ChannelIds {
		app.AppChannels = append(app.AppChannels, &entity.AppChannel{
			ChannelId: item,
		})
	}

	for _, item := range req.FieldIds {
		app.AppFields = append(app.AppFields, &entity.AppField{
			FieldId: item,
		})
	}

	id, err := s.srv.CreateApp(kratosx.MustContext(c), &app)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAppReply{Id: id}, nil
}

// UpdateApp 更新应用信息
func (s *App) UpdateApp(c context.Context, req *pb.UpdateAppRequest) (*pb.UpdateAppReply, error) {
	app := entity.App{
		BaseModel:     ktypes.BaseModel{Id: req.Id},
		Logo:          req.Logo,
		Keyword:       req.Keyword,
		Name:          req.Name,
		AllowRegistry: req.AllowRegistry,
		Version:       req.Version,
		Copyright:     req.Copyright,
		Extra:         req.Extra,
		Description:   req.Description,
	}

	for _, item := range req.ChannelIds {
		app.AppChannels = append(app.AppChannels, &entity.AppChannel{
			ChannelId: item,
		})
	}

	for _, item := range req.FieldIds {
		app.AppFields = append(app.AppFields, &entity.AppField{
			FieldId: item,
		})
	}

	if err := s.srv.UpdateApp(kratosx.MustContext(c), &app); err != nil {
		return nil, err
	}

	return &pb.UpdateAppReply{}, nil
}

// UpdateAppStatus 更新应用信息状态
func (s *App) UpdateAppStatus(c context.Context, req *pb.UpdateAppStatusRequest) (*pb.UpdateAppStatusReply, error) {
	return &pb.UpdateAppStatusReply{}, s.srv.UpdateAppStatus(kratosx.MustContext(c), &types.UpdateAppStatusRequest{
		Id:          req.Id,
		Status:      req.Status,
		DisableDesc: req.DisableDesc,
	})
}

// DeleteApp 删除应用信息
func (s *App) DeleteApp(c context.Context, req *pb.DeleteAppRequest) (*pb.DeleteAppReply, error) {
	if err := s.srv.DeleteApp(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteAppReply{}, nil
}
