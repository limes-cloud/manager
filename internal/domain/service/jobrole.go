package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
	"github.com/samber/lo"
)

type JobRole struct {
	repo  repository.JobRole
	scope repository.Scope
	tad   repository.TenantAdmin
}

func NewJobRole(
	repo repository.JobRole,
	scope repository.Scope,
	tad repository.TenantAdmin,
) *JobRole {
	return &JobRole{
		repo:  repo,
		scope: scope,
		tad:   tad,
	}
}

// ListJobRole 获取指定的部门角色列表
func (rm *JobRole) ListJobRole(ctx core.Context, req *types.ListJobRoleRequest) ([]*entity.Role, uint32, error) {
	if !rm.tad.IsAdmin(ctx.Auth().TenantId, ctx.Auth().UserId) {
		// 获取当前有权限的角色
		req.InRoleIds = rm.scope.RoleScopes(ctx)
	}

	// 获取当前角色有权限的菜单ID
	list, total, err := rm.repo.ListJobRole(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "get menu ids error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}

	return list, total, nil
}

// CreateJobRole 批量创建指定部门的角色
func (rm *JobRole) CreateJobRole(ctx core.Context, req *types.CreateJobRoleRequest) error {
	if !rm.tad.IsAdmin(ctx.Auth().TenantId, ctx.Auth().UserId) {
		// 获取当前有权限的角色ID
		rids := rm.scope.RoleScopes(ctx)
		if len(rids) == 0 {
			return errors.RoleScopeError()
		}

		// 判断是否拥有角色权限
		if !lo.Contains(rids, req.RoleId) {
			return errors.RoleScopeError()
		}

	}

	// 组装为实体列表
	if err := rm.repo.CreateJobRole(ctx, req.JobId, req.RoleId); err != nil {
		return errors.CreateError(err.Error())
	}
	return nil
}

func (rm *JobRole) DeleteJobRole(ctx core.Context, req *types.DeleteJobRoleRequest) error {
	if !rm.tad.IsAdmin(ctx.Auth().TenantId, ctx.Auth().UserId) {

		// 获取当前有权限的角色ID
		rids := rm.scope.RoleScopes(ctx)
		if len(rids) == 0 {
			return errors.RoleScopeError()
		}

		// 判断是否拥有角色权限
		if !lo.Contains(rids, req.RoleId) {
			return errors.RoleScopeError()
		}
	}

	if err := rm.repo.DeleteJobRole(ctx, req.JobId, req.RoleId); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
