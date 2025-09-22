package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
	"github.com/samber/lo"
)

type DeptRole struct {
	repo  repository.DeptRole
	scope repository.Scope
}

func NewDeptRole(
	repo repository.DeptRole,
	scope repository.Scope,
) *DeptRole {
	return &DeptRole{
		repo:  repo,
		scope: scope,
	}
}

// ListDeptRole 获取指定的部门角色列表
func (rm *DeptRole) ListDeptRole(ctx core.Context, req *types.ListDeptRoleRequest) ([]*entity.Role, uint32, error) {
	if !rm.scope.IsSuperAdmin(ctx) {
		// 判断是否具有当前部门ID
		if !rm.scope.HasDeptScope(ctx, entity.DeptEntityName, req.DeptId) {
			return nil, 0, errors.DeptScopeError()
		}

		// 获取当前有权限的角色
		req.InRoleIds = rm.scope.RoleScopes(ctx)
	}

	// 获取当前角色有权限的菜单ID
	list, total, err := rm.repo.ListDeptRole(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "get menu ids error", "err", err.Error())
		return nil, 0, errors.ListError()
	}

	return list, total, nil
}

// ListRoleDept 获取指定角色的所有部门列表
func (rm *DeptRole) ListRoleDept(ctx core.Context, req *types.ListRoleDeptRequest) ([]*entity.Dept, uint32, error) {
	if !rm.scope.IsSuperAdmin(ctx) {
		// 获取当前的角色id
		if !rm.scope.HasRoleScope(ctx, req.RoleId) {
			return nil, 0, errors.RoleScopeError()
		}
		// 获取当前有权限的部门
		all, ids := rm.scope.DeptScopes(ctx, entity.DeptEntityName)
		if !all {
			req.InDeptIds = ids
		}
	}

	// 获取菜单对应的角色列表
	list, total, err := rm.repo.ListRoleDept(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateDeptRoles 批量创建指定部门的角色
func (rm *DeptRole) CreateDeptRoles(ctx core.Context, req *types.CreateDeptRolesRequest) error {
	if !rm.scope.IsSuperAdmin(ctx) {
		// 获取当前有权限的角色ID
		rids := rm.scope.RoleScopes(ctx)
		if len(rids) == 0 {
			return errors.DeptScopeError()
		}

		// 判断是否拥有角色权限
		if len(lo.Intersect(rids, req.RoleIds)) != len(req.RoleIds) {
			return errors.RoleScopeError()
		}

		// 验证当前角色是否具有指定菜单的权限
		if !rm.scope.HasDeptScope(ctx, entity.DeptEntityName, req.DeptId) {
			return errors.DeptScopeError()
		}
	}

	// 组装为实体列表
	var rms []*entity.DeptRole
	for _, v := range req.RoleIds {
		rms = append(rms, &entity.DeptRole{
			RoleId: v,
			DeptId: req.DeptId,
		})
	}
	if err := rm.repo.CreateDeptRoles(ctx, rms); err != nil {
		return errors.CreateError()
	}
	return nil
}

// CreateRoleDepts 角色批量授权给菜单
func (rm *DeptRole) CreateRoleDepts(ctx core.Context, req *types.CreateRoleDeptsRequest) error {
	if !rm.scope.IsSuperAdmin(ctx) {

		// 判断是否拥有部门权限
		all, dids := rm.scope.DeptScopes(ctx, entity.DeptEntityName)
		if !all {
			if len(lo.Intersect(dids, req.DeptIds)) != len(req.DeptIds) {
				return errors.DeptScopeError()
			}
		}

		// 判断是否拥有角色权限
		if !rm.scope.HasRoleScope(ctx, req.RoleId) {
			return errors.RoleScopeError()
		}
	}

	// 组装为实体列表
	var rms []*entity.DeptRole
	for _, v := range req.DeptIds {
		rms = append(rms, &entity.DeptRole{
			DeptId: v,
			RoleId: req.RoleId,
		})
	}
	if err := rm.repo.CreateDeptRoles(ctx, rms); err != nil {
		return errors.CreateError()
	}
	return nil
}

func (rm *DeptRole) DeleteDeptRoles(ctx core.Context, req *types.DeleteDeptRolesRequest) error {
	if !rm.scope.IsSuperAdmin(ctx) {

		// 获取当前有权限的角色ID
		rids := rm.scope.RoleScopes(ctx)
		if len(rids) == 0 {
			return errors.RoleScopeError()
		}

		// 判断是否拥有角色权限
		if len(lo.Intersect(rids, req.RoleIds)) != len(req.RoleIds) {
			return errors.RoleScopeError()
		}

		// 不能操作当前所在部门
		if ctx.Auth().DeptId == req.DeptId {
			return errors.DeptScopeError()
		}

		// 判断是否拥有部门权限
		all, dids := rm.scope.DeptScopes(ctx, entity.DeptEntityName)
		if !all {
			if lo.Contains(dids, req.DeptId) {
				return errors.DeptScopeError()
			}
		}
	}

	// 组装数据
	var rms []*entity.DeptRole
	for _, v := range req.RoleIds {
		rms = append(rms, &entity.DeptRole{
			RoleId: v,
			DeptId: req.DeptId,
		})
	}
	if err := rm.repo.DeleteDeptRoles(ctx, rms); err != nil {
		return errors.DeleteError()
	}
	return nil
}

func (rm *DeptRole) DeleteRoleDepts(ctx core.Context, req *types.DeleteRoleDeptsRequest) error {
	if !rm.scope.IsSuperAdmin(ctx) {
		// 判断是否拥有角色权限
		if !rm.scope.HasRoleScope(ctx, req.RoleId) {
			return errors.RoleScopeError()
		}

		// 不能取消当前角色的菜单
		if lo.Contains(rm.scope.RoleIds(ctx), req.RoleId) {
			return errors.DeleteError("不能取消当前角色的菜单")
		}

		// 判断部门是否在可操作的部门内
		all, dids := rm.scope.DeptScopes(ctx, entity.DeptEntityName)
		if !all {
			if len(lo.Intersect(dids, req.DeptIds)) != len(req.DeptIds) {
				return errors.DeptScopeError()
			}
		}
	}
	// 组装数据
	var rms []*entity.DeptRole
	for _, v := range req.DeptIds {
		rms = append(rms, &entity.DeptRole{
			RoleId: req.RoleId,
			DeptId: v,
		})
	}
	if err := rm.repo.DeleteDeptRoles(ctx, rms); err != nil {
		return errors.DeleteError()
	}
	return nil
}
