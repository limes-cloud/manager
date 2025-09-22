package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/pkg/field"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type App struct {
	repo repository.App
	ta   repository.TenantApp
}

func NewApp(repo repository.App, scope repository.TenantApp) *App {
	return &App{
		repo: repo,
		ta:   scope,
	}
}

// GetApp 获取指定的应用信息
func (u *App) GetApp(ctx core.Context, req *types.GetAppRequest) (*entity.App, error) {
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
		return nil, errors.GetError()
	}

	return res, nil
}

// ListCurrentApp 获取应用信息列表
func (u *App) ListCurrentApp(ctx core.Context, req *types.ListAppRequest) ([]*entity.App, uint32, error) {
	req.InIds = u.ta.GetAppIds(ctx.Auth().TenantId)
	list, total, err := u.repo.ListApp(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list app error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// ListApp 获取应用信息列表
func (u *App) ListApp(ctx core.Context, req *types.ListAppRequest) ([]*entity.App, uint32, error) {
	list, total, err := u.repo.ListApp(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list app error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateApp 创建应用信息
func (u *App) CreateApp(ctx core.Context, app *entity.App) (uint32, error) {
	app.Status = proto.Bool(false)
	app.Reason = proto.String("应用未发布")
	id, err := u.repo.CreateApp(ctx, app)
	if err != nil {
		ctx.Logger().Warnw("msg", "create app error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateApp 更新应用信息
func (u *App) UpdateApp(ctx core.Context, app *entity.App) error {
	if err := u.repo.UpdateApp(ctx, app); err != nil {
		ctx.Logger().Warnw("msg", "update app error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// DeleteApp 删除应用信息
func (u *App) DeleteApp(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteApp(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete app error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}

// ListAppOAuthChannel 获取应用渠道信息
func (u *App) ListAppOAuthChannel(ctx core.Context, req *types.ListAppOAuthChannelRequest) ([]*entity.AppOAuthChannel, uint32, error) {
	list, total, err := u.repo.ListAppOAuthChannel(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create app oauth channel error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateAppOAuthChannel 创建应用渠道信息
func (u *App) CreateAppOAuthChannel(ctx core.Context, app *entity.AppOAuthChannel) (uint32, error) {
	id, err := u.repo.CreateAppOAuthChannel(ctx, app)
	if err != nil {
		ctx.Logger().Warnw("msg", "create app oauth channel error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// DeleteAppOAuthChannel 删除应用渠道信息
func (u *App) DeleteAppOAuthChannel(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteAppOAuthChannel(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete app oauth channel error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}

// ListAppField 获取应用字段信息
func (u *App) ListAppField(ctx core.Context, req *types.ListAppFieldRequest) ([]*entity.AppField, uint32, error) {
	fd := field.New()
	list, total, err := u.repo.ListAppField(ctx, req)
	for _, item := range list {
		if item.Field != nil {
			item.Field.Type = fd.GetType(item.Field.Type).Name()
		}
	}
	if err != nil {
		ctx.Logger().Warnw("msg", "create app oauth channel error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateAppField 创建应用字段信息
func (u *App) CreateAppField(ctx core.Context, app *entity.AppField) (uint32, error) {
	id, err := u.repo.CreateAppField(ctx, app)
	if err != nil {
		ctx.Logger().Warnw("msg", "create app oauth channel error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// DeleteAppField 删除应用字段信息
func (u *App) DeleteAppField(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteAppField(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete app oauth channel error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}
