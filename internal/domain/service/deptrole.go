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
	rids := rm.scope.RoleScopes(ctx)
	if !lo.Contains(rids, req.RoleId) {
		return errors.RoleScopeError()
	}
	if err := rm.repo.CreateDeptRole(ctx, req.DeptId, req.RoleId); err != nil {
		return errors.CreateError()
	}
	return nil
}

func (rm *DeptRole) DeleteDeptRole(ctx core.Context, req *types.DeleteDeptRoleRequest) error {
	rids := rm.scope.RoleScopes(ctx)
	if !lo.Contains(rids, req.RoleId) {
		return errors.RoleScopeError()
	}
	if ctx.Auth().DeptId == req.DeptId {
		return errors.DeptScopeError()
	}
	if err := rm.repo.DeleteDeptRole(ctx, req.DeptId, req.RoleId); err != nil {
		return errors.DeleteError()
	}
	return nil
}
