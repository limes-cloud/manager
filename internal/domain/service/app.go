package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type App struct {
	repo repository.App
}

func NewApp(repo repository.App) *App {
	return &App{
		repo: repo,
	}
}

// GetAppByKeyword 获取指定的应用信息
func (u *App) GetAppByKeyword(ctx core.Context, keyword string) (*entity.App, error) {
	res, err := u.repo.GetAppByKeyword(keyword)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// GetApp 获取指定的应用信息
func (u *App) GetApp(ctx core.Context, id uint32) (*entity.App, error) {
	res, err := u.repo.GetApp(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// ListApp 获取应用信息列表
func (u *App) ListApp(ctx core.Context, req *types.ListAppRequest) ([]*entity.App, uint32, error) {
	list, total, err := u.repo.ListApp(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list app error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateApp 创建应用信息
func (u *App) CreateApp(ctx core.Context, app *entity.App) (uint32, error) {
	id, err := u.repo.CreateApp(ctx, app)
	if err != nil {
		ctx.Logger().Warnw("msg", "create app error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateApp 更新应用信息
func (u *App) UpdateApp(ctx core.Context, app *entity.App) error {
	if err := u.repo.UpdateApp(ctx, app); err != nil {
		ctx.Logger().Warnw("msg", "update app error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteApp 删除应用信息
func (u *App) DeleteApp(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteApp(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete app error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
