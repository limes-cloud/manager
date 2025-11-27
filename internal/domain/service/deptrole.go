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
	tad   repository.TenantAdmin
}

func NewDeptRole(
	repo repository.DeptRole,
	scope repository.Scope,
	tad repository.TenantAdmin,
) *DeptRole {
	return &DeptRole{
		repo:  repo,
		scope: scope,
		tad:   tad,
	}
}

// ListDeptRole 获取指定的部门角色列表
func (rm *DeptRole) ListDeptRole(ctx core.Context, req *types.ListDeptRoleRequest) ([]*entity.Role, uint32, error) {
	if !rm.tad.IsAdmin(ctx.Auth().TenantId, ctx.Auth().UserId) {
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

// CreateDeptRole 批量创建指定部门的角色
func (rm *DeptRole) CreateDeptRole(ctx core.Context, req *types.CreateDeptRoleRequest) error {
	if !rm.tad.IsAdmin(ctx.Auth().TenantId, ctx.Auth().UserId) {
		// 获取当前有权限的角色ID
		rids := rm.scope.RoleScopes(ctx)
		if len(rids) == 0 {
			return errors.DeptScopeError()
		}

		// 判断是否拥有角色权限
		if !lo.Contains(rids, req.RoleId) {
			return errors.RoleScopeError()
		}

	}

	if err := rm.repo.CreateDeptRole(ctx, req.DeptId, req.RoleId); err != nil {
		return errors.CreateError()
	}
	return nil
}

func (rm *DeptRole) DeleteDeptRole(ctx core.Context, req *types.DeleteDeptRoleRequest) error {
	if !rm.tad.IsAdmin(ctx.Auth().TenantId, ctx.Auth().UserId) {

		// 获取当前有权限的角色ID
		rids := rm.scope.RoleScopes(ctx)
		if len(rids) == 0 {
			return errors.RoleScopeError()
		}

		// 判断是否拥有角色权限
		if lo.Contains(rids, req.RoleId) {
			return errors.RoleScopeError()
		}

		// 不能操作当前所在部门
		if ctx.Auth().DeptId == req.DeptId {
			return errors.DeptScopeError()
		}

	}

	if err := rm.repo.DeleteDeptRole(ctx, req.DeptId, req.RoleId); err != nil {
		return errors.DeleteError()
	}
	return nil
}
