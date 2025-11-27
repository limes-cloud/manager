package service

import (
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Tenant struct {
	repo repository.Tenant
	app  repository.App
}

func NewTenant(repo repository.Tenant, app repository.App) *Tenant {
	return &Tenant{repo: repo, app: app}
}

// GetTenant 获取租户列表
func (u *Tenant) GetTenant(ctx core.Context, req *types.GetTenantRequest) (*entity.Tenant, error) {
	var (
		res *entity.Tenant
		err error
	)

	if req.Id != 0 {
		res, err = u.repo.GetTenant(ctx, req.Id)
	} else if req.Keyword != "" {
		res, err = u.repo.GetTenantByKeyword(ctx, req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// ListTenant 获取租户列表
func (u *Tenant) ListTenant(ctx core.Context, req *types.ListTenantRequest) ([]*entity.Tenant, uint32, error) {
	if req.App != nil {
		// 查询当前app信息
		app, err := u.app.GetAppByKeyword(*req.App)
		if err != nil {
			ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
			return nil, 0, errors.GetError(err.Error())
		}
		if !value.Value(app.Status) {
			return nil, 0, errors.GetError(app.Reason)
		}
		req.AppId = &app.Id
	}
	list, total, err := u.repo.ListTenant(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list tenant error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateTenant 创建租户
func (u *Tenant) CreateTenant(ctx core.Context, req *entity.Tenant) (uint32, error) {
	id, err := u.repo.CreateTenant(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create tenant error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateTenant 更新租户
func (u *Tenant) UpdateTenant(ctx core.Context, req *entity.Tenant) error {
	if err := u.repo.UpdateTenant(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update tenant error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteTenant 删除租户
func (u *Tenant) DeleteTenant(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteTenant(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete tenant error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
