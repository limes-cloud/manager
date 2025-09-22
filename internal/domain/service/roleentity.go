package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type RoleEntity struct {
	repo repository.RoleEntity
}

func NewRoleEntity(
	repo repository.RoleEntity,
) *RoleEntity {
	return &RoleEntity{
		repo: repo,
	}
}

// ListRoleEntity 获取用户字段列表
func (srv *RoleEntity) ListRoleEntity(ctx core.Context, req *types.ListRoleEntityRequest) ([]*entity.RoleEntity, uint32, error) {
	list, total, err := srv.repo.ListRoleEntity(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list role entity error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateRoleEntity 创建用户字段
func (srv *RoleEntity) CreateRoleEntity(ctx core.Context, field *entity.RoleEntity) (uint32, error) {
	id, err := srv.repo.CreateRoleEntity(ctx, field)
	if err != nil {
		ctx.Logger().Warnw("msg", "create role entity error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateRoleEntity 更新用户字段
func (srv *RoleEntity) UpdateRoleEntity(ctx core.Context, field *entity.RoleEntity) error {
	if err := srv.repo.UpdateRoleEntity(ctx, field); err != nil {
		ctx.Logger().Warnw("msg", "update role entity error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteRoleEntity 删除用户字段
func (srv *RoleEntity) DeleteRoleEntity(ctx core.Context, id uint32) error {
	if err := srv.repo.DeleteRoleEntity(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete role entity error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}
