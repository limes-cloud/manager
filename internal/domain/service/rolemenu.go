package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
	"github.com/samber/lo"
	lom "github.com/samber/lo/mutable"
)

type RoleMenu struct {
	repo  repository.RoleMenu
	app   repository.App
	scope repository.Scope
	menu  repository.Menu
	role  repository.Role
}

func NewRoleMenu(
	repo repository.RoleMenu,
	app repository.App,
	scope repository.Scope,
	menu repository.Menu,
	role repository.Role,
) *RoleMenu {
	return &RoleMenu{
		app:   app,
		repo:  repo,
		scope: scope,
		menu:  menu,
		role:  role,
	}
}

// GetRoleMenuIds 获取角色的菜单id列表
func (rm *RoleMenu) GetRoleMenuIds(ctx core.Context, req *types.GetRoleMenuIdsRequest) ([]uint32, error) {
	rids := rm.scope.RoleScopes(ctx)
	if !ctx.IsSuperAdmin() && !lo.Contains(rids, req.RoleId) {
		return nil, errors.RoleScopeError()
	}

	mrs := rm.repo.GetMenuIdsByRoleIds([]uint32{req.RoleId})
	if req.AppId != nil {
		appMenuIds, err := rm.menu.GetMenuIdsByAppId(ctx, *req.AppId)
		if err != nil {
			return nil, errors.SystemError()
		}
		mrs = lo.Intersect(mrs, appMenuIds)
	}
	return mrs, nil
}

// GetMenuRoleIds 获取指定菜单的所有角色ID
func (rm *RoleMenu) GetMenuRoleIds(ctx core.Context, req *types.GetMenuRoleIdsRequest) ([]uint32, error) {
	rids := rm.scope.RoleScopes(ctx)
	mrs := rm.repo.GetRoleIdsByMenuIds([]uint32{req.MenuId})
	if !ctx.IsSuperAdmin() {
		// 验证当前角色是否具有指定菜单的权限
		lom.Filter(mrs, func(item uint32) bool {
			return !lo.Contains(rids, item)
		})
	}
	return mrs, nil
}

// CreateMenuRoles 菜单批量授权给角色
func (rm *RoleMenu) CreateMenuRoles(ctx core.Context, req *types.CreateMenuRolesRequest) error {
	// 超级管理员不做权限校验
	if !ctx.IsSuperAdmin() {
		// 获取当前有权限的角色ID
		rids := rm.scope.RoleScopes(ctx)
		if len(rids) == 0 {
			return errors.RoleScopeError()
		}

		// 判断是否拥有角色权限
		if len(lo.Intersect(rids, req.RoleIds)) != len(req.RoleIds) {
			return errors.RoleScopeError()
		}

		// 获取当前角色有权限的菜单ID,验证当前角色是否具有指定菜单的权限
		mids := rm.repo.GetMenuIdsByRoleIds(rids)
		if !lo.Contains(mids, req.MenuId) {
			return errors.MenuScopeError()
		}
	}

	// 组装为实体列表
	var rms []*entity.RoleMenu
	for _, v := range req.RoleIds {
		rms = append(rms, &entity.RoleMenu{
			RoleId: v,
			MenuId: req.MenuId,
		})
	}
	if err := rm.repo.CreateRoleMenus(ctx, rms); err != nil {
		return errors.CreateError()
	}
	return nil
}

// CreateRoleMenus 角色批量授权给菜单
func (rm *RoleMenu) CreateRoleMenus(ctx core.Context, req *types.CreateRoleMenusRequest) error {
	// 获取appid下的所有菜单
	appMenuIds, err := rm.menu.GetMenuIdsByAppId(ctx, req.AppId)
	if err != nil {
		return errors.SystemError()
	}

	// 获取当前角色的菜单
	roleMenuIds := appMenuIds
	if !ctx.IsSuperAdmin() {
		roleMenuIds = rm.repo.GetMenuIdsByRoleIds(rm.scope.RoleScopes(ctx))
	}

	// 获取应用下当前角色的最大菜单权限
	roleMenuIds = lo.Intersect(appMenuIds, roleMenuIds)

	// 获取修改的角色的菜单权限
	editRoleMenuIds := rm.repo.GetMenuIdsByRoleIds([]uint32{req.RoleId})

	// 剔除当前角色无权限的菜单
	editRoleMenuIds = lo.Intersect(editRoleMenuIds, roleMenuIds)
	req.MenuIds = lo.Intersect(req.MenuIds, roleMenuIds)

	// 获取增加、删除的菜单
	rmIds, addIds := lo.Difference(editRoleMenuIds, req.MenuIds)

	// 获取当前角色到被修改角色的节点，同步添加被增加的菜单
	// 获取当前的角色下的所有子角色id
	childRoles, err := rm.role.GetRoleChildrenIds(ctx, rm.scope.RoleScopes(ctx))
	if err != nil {
		return errors.SystemError()
	}
	// 获取被修改角色所有父角色id
	parentRoles, err := rm.role.GetRoleParentIds(ctx, []uint32{req.RoleId})
	if err != nil {
		return errors.SystemError()
	}
	// 获取当前角色到被修改角色的节点，同步添加被增加的菜单
	addRoles := lo.Intersect(childRoles, parentRoles)
	addRoles = append(addRoles, req.RoleId)

	// 删除当前角色菜单
	err = ctx.Transaction(func(ctx core.Context) error {
		if len(rmIds) != 0 {
			var rms []*entity.RoleMenu
			for _, v := range rmIds {
				rms = append(rms, &entity.RoleMenu{
					RoleId: req.RoleId,
					MenuId: v,
				})
			}
			if err := rm.repo.DeleteRoleMenus(ctx, rms); err != nil {
				return err
			}
		}
		if len(addIds) != 0 {
			var rms []*entity.RoleMenu
			for _, v := range addIds {
				for _, roleId := range addRoles {
					rms = append(rms, &entity.RoleMenu{
						RoleId: roleId,
						MenuId: v,
					})
				}
			}
			if err := rm.repo.CreateRoleMenus(ctx, rms); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return errors.CreateError()
	}
	return nil
}

func (rm *RoleMenu) DeleteRoleMenus(ctx core.Context, req *types.DeleteRoleMenusRequest) error {
	if !ctx.IsSuperAdmin() {
		// 获取当前有权限的角色ID
		rids := rm.scope.RoleScopes(ctx)
		if len(rids) == 0 {
			return errors.RoleScopeError()
		}

		// 判断是否拥有角色权限
		if !lo.Contains(rids, req.RoleId) {
			return errors.RoleScopeError()
		}

		// 不能取消当前角色的菜单
		if !lo.Contains(rm.scope.RoleIds(ctx), req.RoleId) {
			return errors.DeleteError("不能取消当前角色的菜单")
		}

		// 判断菜单是否在可操作的菜单内
		mids := rm.repo.GetMenuIdsByRoleIds(rids)
		if len(lo.Intersect(mids, req.MenuIds)) != len(req.MenuIds) {
			return errors.MenuScopeError()
		}
	}

	// 组装数据
	var rms []*entity.RoleMenu
	for _, v := range req.MenuIds {
		rms = append(rms, &entity.RoleMenu{
			RoleId: req.RoleId,
			MenuId: v,
		})
	}
	if err := rm.repo.DeleteRoleMenus(ctx, rms); err != nil {
		return errors.DeleteError()
	}
	return nil
}

func (rm *RoleMenu) DeleteMenuRoles(ctx core.Context, req *types.DeleteMenuRolesRequest) error {
	if !ctx.IsSuperAdmin() {
		// 获取当前有权限的角色ID
		rids := rm.scope.RoleScopes(ctx)
		if len(rids) == 0 {
			return errors.RoleScopeError()
		}

		// 判断是否拥有角色权限
		if len(lo.Intersect(rids, req.RoleIds)) != len(req.RoleIds) {
			return errors.RoleScopeError()
		}

		// 不能取消当前角色的菜单
		if len(lo.Intersect(rm.scope.RoleIds(ctx), req.RoleIds)) > 1 {
			return errors.DeleteError("不能取消当前角色的菜单")
		}

		// 判断菜单是否在可操作的菜单内
		mids := rm.repo.GetMenuIdsByRoleIds(rids)
		if lo.Contains(mids, req.MenuId) {
			return errors.MenuScopeError()
		}
	}

	// 组装数据
	var rms []*entity.RoleMenu
	for _, v := range req.RoleIds {
		rms = append(rms, &entity.RoleMenu{
			RoleId: v,
			MenuId: req.MenuId,
		})
	}
	if err := rm.repo.DeleteRoleMenus(ctx, rms); err != nil {
		return errors.DeleteError()
	}
	return nil
}
