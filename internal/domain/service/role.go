package service

import (
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/samber/lo"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Role struct {
	repo  repository.Role
	rm    repository.RoleMenu
	scope repository.Scope
}

func NewRole(
	repo repository.Role,
	rm repository.RoleMenu,
	scope repository.Scope,
) *Role {
	return &Role{
		repo:  repo,
		rm:    rm,
		scope: scope,
	}
}

// GetRole 获取角色列表
func (u *Role) GetRole(ctx core.Context, req *types.GetRoleRequest) (*entity.Role, error) {
	// 按 App + Keyword 定位角色（keyword 在单个 app 内唯一，无需全局 id）
	res, err := u.repo.GetRoleByKeyword(ctx, req.App, req.Keyword)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError()
	}
	return res, nil
}

// ListCurrentRole 获取当前角色信息列表树
func (u *Role) ListCurrentRole(ctx core.Context, req *types.ListRoleRequest) ([]*entity.Role, error) {
	// 获取角色权限

	return u.ListRole(ctx, req)
}

// ListRole 获取角色列表
func (u *Role) ListRole(ctx core.Context, req *types.ListRoleRequest) ([]*entity.Role, error) {
	list, err := u.repo.ListRole(ctx, ctx.AppKey(), req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list role error", "err", err.Error())
		return nil, errors.ListError()
	}
	return tree.BuildArrayTree(list), nil
}

// CreateRole 创建角色
func (u *Role) CreateRole(ctx core.Context, req *entity.Role) (uint32, error) {
	// 非超级管理员只能在自己的管辖角色下创建子角色
	if !ctx.Auth().IsAdmin {
		if !u.scope.HasRoleScope(ctx, value.Value(req.ParentId)) {
			return 0, errors.RoleScopeError()
		}
	}

	// 创建角色
	id, err := u.repo.CreateRole(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create role error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateRole 更新角色
func (u *Role) UpdateRole(ctx core.Context, req *entity.Role) error {
	// 超级管理员可管理所有角色，跳过角色权限校验
	if !ctx.Auth().IsAdmin {
		// 不能修改自己的角色，防止主动提权
		if lo.Contains(u.scope.RoleIds(ctx), req.Id) {
			return errors.UpdateError("不能修改当前用户所属角色")
		}

		// 判断是否具有角色权限
		if !u.scope.HasRoleScope(ctx, req.Id) {
			return errors.RoleScopeError()
		}

		// 如果修改父级菜单,判断是否具有角色权限
		old, err := u.repo.GetRole(ctx, req.App, req.Id)
		if err != nil {
			return errors.UpdateError()
		}
		if value.Value(old.ParentId) != value.Value(req.ParentId) && !u.scope.HasRoleScope(ctx, value.Value(req.ParentId)) {
			return errors.RoleScopeError()
		}
	}

	// 更新角色
	if err := u.repo.UpdateRole(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update role error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// DeleteRole 删除角色
func (u *Role) DeleteRole(ctx core.Context, app string, id uint32) error {
	// 超级管理员可管理所有角色，跳过角色权限校验
	if !ctx.Auth().IsAdmin {
		// 不能删除自己的角色
		if lo.Contains(u.scope.RoleIds(ctx), id) {
			return errors.UpdateError("不能删除当前用户所属角色")
		}

		// 判断是否具有角色权限
		if !u.scope.HasRoleScope(ctx, id) {
			return errors.RoleScopeError()
		}
	}

	// 查询下级角色
	rids, err := u.repo.GetRoleChildrenIds(ctx, []uint32{id})
	if err != nil {
		return err
	}

	if err := ctx.Transaction(func(ctx core.Context) error {
		// 删除角色
		if err = u.repo.DeleteRole(ctx, app, id); err != nil {
			return err
		}

		// 删除权限
		return u.rm.DeleteRoles(ctx, app, rids)
	}); err != nil {
		ctx.Logger().Warnw("msg", "delete role error", "err", err.Error())
		return errors.DeleteError()
	}

	return nil
}
