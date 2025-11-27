package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
	"github.com/samber/lo"
)

type TenantApp struct {
	repo repository.TenantApp
}

func NewTenantApp(repo repository.TenantApp) *TenantApp {
	return &TenantApp{repo: repo}
}

// GetTenantApp 获取租户列表
func (u *TenantApp) GetTenantApp(ctx core.Context, req *types.GetTenantAppRequest) (*entity.TenantApp, error) {
	res, err := u.repo.GetTenantApp(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list TenantApp error", "err", err.Error())
		return nil, errors.ListError(err.Error())
	}

	// 获取菜单ID
	menuIds, err := u.repo.GetTenantAppMenuIds(ctx, req.TenantId, req.AppId)
	if err != nil {
		ctx.Logger().Warnw("msg", "list TenantApp error", "err", err.Error())
		return nil, errors.ListError(err.Error())
	}
	res.MenuIds = menuIds
	return res, nil
}

// ListTenantApp 获取租户列表
func (u *TenantApp) ListTenantApp(ctx core.Context, req *types.ListTenantAppRequest) ([]*entity.TenantApp, uint32, error) {
	list, total, err := u.repo.ListTenantApp(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list TenantApp error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateTenantApp 创建租户
func (u *TenantApp) CreateTenantApp(ctx core.Context, req *types.CreateTenantAppRequest) (uint32, error) {
	var (
		id  uint32
		err error
	)
	err = ctx.Transaction(func(ctx core.Context) error {
		// 创建租户应用
		id, err = u.repo.CreateTenantApp(ctx, &entity.TenantApp{
			TenantId:  req.TenantId,
			AppId:     req.AppId,
			ExpiredAt: req.ExpiredAt,
			Setting:   req.Setting,
		})
		if err != nil {
			return err
		}
		// 创建租户菜单
		return u.repo.CreateTenantAppMenuIds(ctx, req.TenantId, req.AppId, req.MenuIds)
	})
	if err != nil {
		ctx.Logger().Warnw("msg", "create TenantApp error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateTenantApp 更新租户
func (u *TenantApp) UpdateTenantApp(ctx core.Context, req *types.UpdateTenantAppRequest) error {
	// 获取当前的全部菜单ID
	menuIds, err := u.repo.GetTenantAppMenuIds(ctx, req.TenantId, req.AppId)
	if err != nil {
		return errors.SystemError()
	}

	// 判断需要增加的菜单id
	rmIds, addIds := lo.Difference(menuIds, req.MenuIds)

	err = ctx.Transaction(func(ctx core.Context) error {
		// 创建租户应用
		err = u.repo.UpdateTenantApp(ctx, &entity.TenantApp{
			TenantId:  req.TenantId,
			AppId:     req.AppId,
			ExpiredAt: req.ExpiredAt,
			Setting:   req.Setting,
		})
		if err != nil {
			return err
		}

		if len(rmIds) != 0 {
			err = u.repo.DeleteTenantAppMenuIds(ctx, req.TenantId, req.AppId, rmIds)
			if err != nil {
				return err
			}
		}

		if len(addIds) != 0 {
			err = u.repo.CreateTenantAppMenuIds(ctx, req.TenantId, req.AppId, addIds)
			if err != nil {
				return err
			}
		}

		// 创建租户菜单
		return nil
	})
	if err != nil {
		ctx.Logger().Warnw("msg", "create TenantApp error", "err", err.Error())
		return errors.CreateError()
	}
	return nil
}

// DeleteTenantApp 删除租户
func (u *TenantApp) DeleteTenantApp(ctx core.Context, tid uint32, aid uint32) error {
	if err := u.repo.DeleteTenantApp(ctx, tid, aid); err != nil {
		ctx.Logger().Warnw("msg", "delete TenantApp error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}

// GetTenantAppMenuIds 获取租户菜单列表
func (u *TenantApp) GetTenantAppMenuIds(ctx core.Context, tid uint32, aid uint32) ([]uint32, error) {
	ids, err := u.repo.GetTenantAppMenuIds(ctx, tid, aid)
	if err != nil {
		ctx.Logger().Warnw("msg", "delete TenantApp error", "err", err.Error())
		return nil, errors.DeleteError()
	}
	return ids, nil
}
