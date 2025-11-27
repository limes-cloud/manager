package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type AppOAuther struct {
	repo    repository.AppOAuther
	app     repository.App
	tenant  repository.Tenant
	oauther repository.OAuther
}

func NewAppOAuther(
	repo repository.AppOAuther,
	app repository.App,
	tenant repository.Tenant,
	oauther repository.OAuther,
) *AppOAuther {
	return &AppOAuther{
		repo:    repo,
		app:     app,
		tenant:  tenant,
		oauther: oauther,
	}
}

// ListAppOAuther 获取应用渠道信息
func (u *AppOAuther) ListAppOAuther(ctx core.Context, req *types.ListAppOAutherRequest) ([]*entity.AppOAuther, uint32, error) {
	list, total, err := u.repo.ListAppOAuther(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list app oauth channel error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateAppOAuther 创建应用渠道信息
func (u *AppOAuther) CreateAppOAuther(ctx core.Context, app *entity.AppOAuther) (uint32, error) {
	// 查询渠道的类型
	oauther, err := u.oauther.GetOAuther(ctx, app.OAutherId)
	if err != nil {
		return 0, errors.DatabaseError("不存在授权渠道")
	}
	app.Type = oauther.Type

	// 查询是否存在类型相同的渠道
	_, err = u.repo.GetAppOAutherByAT(ctx, app.AppId, app.Type)
	if err == nil {
		return 0, errors.CreateError("已存在相同类型的授权渠道")
	}

	id, err := u.repo.CreateAppOAuther(ctx, app)
	if err != nil {
		ctx.Logger().Warnw("msg", "create app oauth channel error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// DeleteAppOAuther 删除应用渠道信息
func (u *AppOAuther) DeleteAppOAuther(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteAppOAuther(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete app oauth channel error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
