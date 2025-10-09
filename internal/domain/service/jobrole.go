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
}

func NewJobRole(
	repo repository.JobRole,
	scope repository.Scope,
) *JobRole {
	return &JobRole{
		repo:  repo,
		scope: scope,
	}
}

// ListJobRole 获取指定的部门角色列表
func (rm *JobRole) ListJobRole(ctx core.Context, req *types.ListJobRoleRequest) ([]*entity.Role, uint32, error) {
	if !ctx.IsSuperAdmin() {
		// 获取当前有权限的角色
		req.InRoleIds = rm.scope.RoleScopes(ctx)
	}

	// 获取当前角色有权限的菜单ID
	list, total, err := rm.repo.ListJobRole(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "get menu ids error", "err", err.Error())
		return nil, 0, errors.ListError()
	}

	return list, total, nil
}

// ListRoleJob 获取指定角色的所有部门列表
func (rm *JobRole) ListRoleJob(ctx core.Context, req *types.ListRoleJobRequest) ([]*entity.Job, uint32, error) {
	if !ctx.IsSuperAdmin() {
		// 获取当前的角色id
		if !rm.scope.HasRoleScope(ctx, req.RoleId) {
			return nil, 0, errors.RoleScopeError()
		}
	}

	// 获取菜单对应的角色列表
	list, total, err := rm.repo.ListRoleJob(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateJobRoles 批量创建指定部门的角色
func (rm *JobRole) CreateJobRoles(ctx core.Context, req *types.CreateJobRolesRequest) error {
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

	}

	// 组装为实体列表
	var rms []*entity.JobRole
	for _, v := range req.RoleIds {
		rms = append(rms, &entity.JobRole{
			RoleId: v,
			JobId:  req.JobId,
		})
	}
	if err := rm.repo.CreateJobRoles(ctx, rms); err != nil {
		return errors.CreateError()
	}
	return nil
}

// CreateRoleJobs 角色批量授权给菜单
func (rm *JobRole) CreateRoleJobs(ctx core.Context, req *types.CreateRoleJobsRequest) error {
	if !ctx.IsSuperAdmin() {
		// 判断是否拥有角色权限
		if !rm.scope.HasRoleScope(ctx, req.RoleId) {
			return errors.RoleScopeError()
		}
	}

	// 组装为实体列表
	var rms []*entity.JobRole
	for _, v := range req.JobIds {
		rms = append(rms, &entity.JobRole{
			JobId:  v,
			RoleId: req.RoleId,
		})
	}
	if err := rm.repo.CreateJobRoles(ctx, rms); err != nil {
		return errors.CreateError()
	}
	return nil
}

func (rm *JobRole) DeleteJobRoles(ctx core.Context, req *types.DeleteJobRolesRequest) error {
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

	}

	// 组装数据
	var rms []*entity.JobRole
	for _, v := range req.RoleIds {
		rms = append(rms, &entity.JobRole{
			RoleId: v,
			JobId:  req.JobId,
		})
	}
	if err := rm.repo.DeleteJobRoles(ctx, rms); err != nil {
		return errors.DeleteError()
	}
	return nil
}

func (rm *JobRole) DeleteRoleJobs(ctx core.Context, req *types.DeleteRoleJobsRequest) error {
	if !ctx.IsSuperAdmin() {
		// 判断是否拥有角色权限
		if !rm.scope.HasRoleScope(ctx, req.RoleId) {
			return errors.RoleScopeError()
		}

		// 不能取消当前角色的菜单
		if lo.Contains(rm.scope.RoleIds(ctx), req.RoleId) {
			return errors.DeleteError("不能取消当前角色的菜单")
		}

	}
	// 组装数据
	var rms []*entity.JobRole
	for _, v := range req.JobIds {
		rms = append(rms, &entity.JobRole{
			RoleId: req.RoleId,
			JobId:  v,
		})
	}
	if err := rm.repo.DeleteJobRoles(ctx, rms); err != nil {
		return errors.DeleteError()
	}
	return nil
}
