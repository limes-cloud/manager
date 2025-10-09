package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type UserDept struct {
	user  repository.User
	repo  repository.UserDept
	scope repository.Scope
}

func NewUserDept(
	repo repository.UserDept,
	scope repository.Scope,
	user repository.User,
) *UserDept {
	return &UserDept{
		repo:  repo,
		scope: scope,
		user:  user,
	}
}

// ListUserDept 获取指定的部门角色列表
func (rm *UserDept) ListUserDept(ctx core.Context, req *types.ListUserDeptRequest) ([]*entity.Dept, uint32, error) {
	//if !ctx.IsSuperAdmin() {
	//	// 获取当前有权限的角色
	//	all, ids := rm.scope.DeptScopes(ctx, entity.UserEntityName)
	//	if !all {
	//		req.InDeptIds = ids
	//	}
	//}

	// 获取当前角色有权限的菜单ID
	list, total, err := rm.repo.ListUserDept(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "get menu ids error", "err", err.Error())
		return nil, 0, errors.ListError()
	}

	return list, total, nil
}

// ListDeptUser 获取指定角色的所有部门列表
func (rm *UserDept) ListDeptUser(ctx core.Context, req *types.ListDeptUserRequest) ([]*entity.User, uint32, error) {
	//if !ctx.IsSuperAdmin() {
	//	// 获取当前有权限的角色
	//	has := rm.scope.HasDeptScope(ctx, entity.UserEntityName, req.DeptId)
	//	if !has {
	//		return nil, 0, errors.DeptScopeError()
	//	}
	//}

	// 获取菜单对应的角色列表
	list, total, err := rm.repo.ListDeptUser(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateUserDept 创建指定部门的角色
func (rm *UserDept) CreateUserDept(ctx core.Context, req *entity.UserDept) error {
	//if !ctx.IsSuperAdmin() {
	//	// 获取当前有权限的角色
	//	has := rm.scope.HasDeptScope(ctx, entity.UserEntityName, req.DeptId)
	//	if !has {
	//		return errors.DeptScopeError()
	//	}
	//
	//	// 获取当前用户的部门
	//	user, err := rm.user.GetUser(ctx, req.UserId)
	//	if err != nil {
	//		return errors.GetError()
	//	}
	//	has = rm.scope.HasDeptScope(ctx, entity.UserEntityName, user.DeptId)
	//	if !has {
	//		return errors.DeptScopeError()
	//	}
	//}

	if err := rm.repo.CreateUserDept(ctx, req); err != nil {
		return errors.CreateError()
	}
	return nil
}

func (rm *UserDept) DeleteUserDept(ctx core.Context, req *entity.UserDept) error {
	//if !ctx.IsSuperAdmin() {
	//	// 获取当前有权限的角色
	//	has := rm.scope.HasDeptScope(ctx, entity.UserEntityName, req.DeptId)
	//	if !has {
	//		return errors.DeptScopeError()
	//	}
	//
	//	// 获取当前用户的部门
	//	user, err := rm.user.GetUser(ctx, req.UserId)
	//	if err != nil {
	//		return errors.GetError()
	//	}
	//	has = rm.scope.HasDeptScope(ctx, entity.UserEntityName, user.DeptId)
	//	if !has {
	//		return errors.DeptScopeError()
	//	}
	//}

	if err := rm.repo.DeleteUserDept(ctx, req); err != nil {
		return errors.DeleteError()
	}
	return nil
}
