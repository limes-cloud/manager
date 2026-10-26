package service

import (
	"github.com/samber/lo"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type RoleEntity struct {
	repo  repository.RoleEntity
	scope repository.Scope // 角色管辖范围（防越权授予）
}

func NewRoleEntity(
	repo repository.RoleEntity,
	scope repository.Scope,
) *RoleEntity {
	return &RoleEntity{
		repo:  repo,
		scope: scope,
	}
}

// ListRoleEntity 获取用户字段列表
func (srv *RoleEntity) ListRoleEntity(ctx core.Context, app string, req *types.ListRoleEntityRequest) ([]*entity.RoleEntity, uint32, error) {
	// 非超管只能查询自己管辖角色的数据权限，防止越权查看其他角色的授权配置
	if !ctx.Auth().IsAdmin && !lo.Contains(srv.scope.RoleScopes(ctx), req.RoleId) {
		return nil, 0, errors.RoleScopeError()
	}

	list, total, err := srv.repo.ListRoleEntity(ctx, app, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list role entity error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateRoleEntities 批量创建数据权限（同一实体下的多个操作各自独立配置）
func (srv *RoleEntity) CreateRoleEntities(ctx core.Context, fields []*entity.RoleEntity) ([]uint32, error) {
	if len(fields) == 0 {
		return nil, nil
	}

	// 校验所有角色都需在管辖范围内，防止越权批量授予
	for _, field := range fields {
		if err := srv.assertRoleScope(ctx, field.RoleId); err != nil {
			return nil, err
		}
	}

	ids := make([]uint32, 0, len(fields))
	err := ctx.Transaction(func(ctx core.Context) error {
		for _, field := range fields {
			id, err := srv.repo.CreateRoleEntity(ctx, field)
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		ctx.Logger().Warnw("msg", "create role entities error", "err", err.Error())
		return nil, errors.CreateError(err.Error())
	}
	return ids, nil
}

// CreateRoleEntity 创建用户字段
func (srv *RoleEntity) CreateRoleEntity(ctx core.Context, field *entity.RoleEntity) (uint32, error) {
	// 单条模式要求操作与范围必填（批量走 CreateRoleEntities，不经过此方法）
	if field.Action == "" || field.Scope == "" {
		return 0, errors.ParamsError()
	}
	if err := srv.assertRoleScope(ctx, field.RoleId); err != nil {
		return 0, err
	}

	id, err := srv.repo.CreateRoleEntity(ctx, field)
	if err != nil {
		ctx.Logger().Warnw("msg", "create role entity error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateRoleEntity 更新用户字段
func (srv *RoleEntity) UpdateRoleEntity(ctx core.Context, field *entity.RoleEntity) error {
	// 按 id 查询原记录，用原记录的 RoleId 做管辖校验（前端可能不传 RoleId）
	old, err := srv.repo.GetRoleEntity(ctx, field.App, field.Id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get role entity error", "err", err.Error())
		return errors.GetError()
	}
	if err := srv.assertRoleScope(ctx, old.RoleId); err != nil {
		return err
	}

	if err := srv.repo.UpdateRoleEntity(ctx, field); err != nil {
		ctx.Logger().Warnw("msg", "update role entity error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteRoleEntity 删除用户字段
func (srv *RoleEntity) DeleteRoleEntity(ctx core.Context, app string, id uint32) error {
	// 先查询实体所属角色，校验其是否在当前操作者的管辖范围内，防止越权删除
	re, err := srv.repo.GetRoleEntity(ctx, app, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get role entity error", "err", err.Error())
		return errors.GetError()
	}
	if err := srv.assertRoleScope(ctx, re.RoleId); err != nil {
		return err
	}

	if err := srv.repo.DeleteRoleEntity(ctx, app, id); err != nil {
		ctx.Logger().Warnw("msg", "delete role entity error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}

// assertRoleScope 校验被授权角色是否在当前操作者的角色管辖范围内
// 否则操作者可以给自己管不到的角色配置数据权限
func (srv *RoleEntity) assertRoleScope(ctx core.Context, roleId uint32) error {
	if ctx.Auth().IsAdmin {
		return nil
	}
	if !lo.Contains(srv.scope.RoleScopes(ctx), roleId) {
		return errors.RoleScopeError()
	}
	return nil
}
