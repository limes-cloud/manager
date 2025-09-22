package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Tenant struct {
	repo repository.Tenant
}

func NewTenant(repo repository.Tenant) *Tenant {
	return &Tenant{repo: repo}
}

// GetTenant 获取租户列表
func (u *Tenant) GetTenant(ctx core.Context, req *types.GetTenantRequest) (*entity.Tenant, error) {
	var (
		res *entity.Tenant
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetTenant(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetTenantByKeyword(ctx, *req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError()
	}
	return res, nil
}

// ListTenant 获取租户列表
func (u *Tenant) ListTenant(ctx core.Context, req *types.ListTenantRequest) ([]*entity.Tenant, uint32, error) {
	list, total, err := u.repo.ListTenant(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list tenant error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateTenant 创建租户
func (u *Tenant) CreateTenant(ctx core.Context, req *entity.Tenant) (uint32, error) {
	id, err := u.repo.CreateTenant(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create tenant error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateTenant 更新租户
func (u *Tenant) UpdateTenant(ctx core.Context, req *entity.Tenant) error {
	if err := u.repo.UpdateTenant(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update tenant error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// DeleteTenant 删除租户
func (u *Tenant) DeleteTenant(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteTenant(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete tenant error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}
