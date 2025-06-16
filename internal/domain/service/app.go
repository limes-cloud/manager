package service

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

const (
	permissionApp = "manager_app"
)

type App struct {
	conf       *conf.Config
	repo       repository.App
	permission repository.Permission
	file       repository.File
}

func NewApp(
	conf *conf.Config,
	repo repository.App,
	permission repository.Permission,
	file repository.File,
) *App {
	return &App{
		conf:       conf,
		repo:       repo,
		permission: permission,
		file:       file,
	}
}

// GetApp 获取指定的应用信息
func (u *App) GetApp(ctx kratosx.Context, req *types.GetAppRequest) (*entity.App, error) {
	var (
		res *entity.App
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetApp(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetAppByKeyword(ctx, *req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	res.LogoUrl = u.file.GetFileURL(ctx, res.Logo)
	for _, item := range res.Channels {
		item.Logo = u.file.GetFileURL(ctx, item.Logo)
	}
	return res, nil
}

// ListApp 获取应用信息列表
func (u *App) ListApp(ctx kratosx.Context, req *types.ListAppRequest) ([]*entity.App, uint32, error) {
	all, scopes, err := u.permission.GetResourceScopes(ctx, permissionApp)
	if err != nil {
		return nil, 0, err
	}
	if !all {
		req.Ids = scopes
	}
	list, total, err := u.repo.ListApp(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list app error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	for _, item := range list {
		item.LogoUrl = u.file.GetFileURL(ctx, item.Logo)
	}
	return list, total, nil
}

// CreateApp 创建应用信息
func (u *App) CreateApp(ctx kratosx.Context, app *entity.App) (uint32, error) {
	app.Status = proto.Bool(false)
	app.DisableDesc = proto.String("应用未发布")
	id, err := u.repo.CreateApp(ctx, app)
	if err != nil {
		ctx.Logger().Warnw("msg", "create app error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateApp 更新应用信息
func (u *App) UpdateApp(ctx kratosx.Context, app *entity.App) error {
	if !u.HasApp(ctx, app.Id) {
		return errors.NotPermissionError()
	}
	if err := u.repo.UpdateApp(ctx, app); err != nil {
		ctx.Logger().Warnw("msg", "update app error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateAppStatus 更新应用信息状态
func (u *App) UpdateAppStatus(ctx kratosx.Context, req *types.UpdateAppStatusRequest) error {
	if !u.HasApp(ctx, req.Id) {
		return errors.NotPermissionError()
	}

	if err := u.repo.UpdateAppStatus(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update app status error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteApp 删除应用信息
func (u *App) DeleteApp(ctx kratosx.Context, id uint32) error {
	if !u.HasApp(ctx, id) {
		return errors.NotPermissionError()
	}
	if err := u.repo.DeleteApp(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete app error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// HasApp 获取当前用户是否具有指定env的权限
func (u *App) HasApp(ctx kratosx.Context, id uint32) bool {
	all, ids, err := u.permission.GetResourceScopes(ctx, permissionApp)
	if err != nil {
		return false
	}
	if all {
		return true
	}
	return valx.InList(ids, id)
}
