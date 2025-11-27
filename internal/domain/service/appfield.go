package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type AppField struct {
	repo   repository.AppField
	app    repository.App
	tenant repository.Tenant
}

func NewAppField(repo repository.AppField, app repository.App, tenant repository.Tenant) *AppField {
	return &AppField{
		repo:   repo,
		app:    app,
		tenant: tenant,
	}
}

// ListAppField 获取应用字段信息
func (u *AppField) ListAppField(ctx core.Context, req *types.ListAppFieldRequest) ([]*entity.AppField, uint32, error) {
	list, total, err := u.repo.ListAppField(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list app field error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateAppField 创建应用字段信息
func (u *AppField) CreateAppField(ctx core.Context, app *entity.AppField) (uint32, error) {
	id, err := u.repo.CreateAppField(ctx, app)
	if err != nil {
		ctx.Logger().Warnw("msg", "create app field error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateAppField 更新应用字段信息
func (u *AppField) UpdateAppField(ctx core.Context, app *entity.AppField) error {
	err := u.repo.UpdateAppField(ctx, app)
	if err != nil {
		ctx.Logger().Warnw("msg", "create app field error", "err", err.Error())
		return errors.CreateError(err.Error())
	}
	return nil
}

// DeleteAppField 删除应用字段信息
func (u *AppField) DeleteAppField(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteAppField(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete app field error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
