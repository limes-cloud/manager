package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type TenantAdmin struct {
	repo repository.TenantAdmin
}

func NewTenantAdmin(repo repository.TenantAdmin) *TenantAdmin {
	return &TenantAdmin{repo: repo}
}

// ListTenantAdmin 获取租户列表
func (u *TenantAdmin) ListTenantAdmin(ctx core.Context, req *types.ListTenantAdminRequest) ([]*entity.TenantAdmin, uint32, error) {
	list, total, err := u.repo.ListTenantAdmin(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list tenant admin error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateTenantAdmin 创建租户
func (u *TenantAdmin) CreateTenantAdmin(ctx core.Context, req *types.CreateTenantAdminRequest) (uint32, error) {
	var (
		id  uint32
		err error
	)
	err = ctx.Transaction(func(ctx core.Context) error {
		// 创建租户应用
		id, err = u.repo.CreateTenantAdmin(ctx, &entity.TenantAdmin{
			TenantId: req.TenantId,
			UserId:   req.UserId,
		})
		return err
	})
	if err != nil {
		ctx.Logger().Warnw("msg", "create tenant admin error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// DeleteTenantAdmin 删除租户
func (u *TenantAdmin) DeleteTenantAdmin(ctx core.Context, req *types.DeleteTenantAdminRequest) error {
	if err := u.repo.DeleteTenantAdmin(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "delete tenant admin error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}
