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
func (rm *UserDept) ListUserDept(ctx core.Context, req *types.ListUserDeptRequest) ([]*entity.UserDept, uint32, error) {
	// 获取当前角色有权限的菜单ID
	list, total, err := rm.repo.ListUserDept(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "get menu ids error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}

	return list, total, nil
}

// ListDeptUser 获取指定角色的所有部门列表
func (rm *UserDept) ListDeptUser(ctx core.Context, req *types.ListDeptUserRequest) ([]*entity.User, uint32, error) {
	list, total, err := rm.repo.ListDeptUser(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateUserDept 创建指定部门的角色
func (rm *UserDept) CreateUserDept(ctx core.Context, req *entity.UserDept) error {
	if err := rm.repo.CreateUserDept(ctx, req); err != nil {
		return errors.CreateError(err.Error())
	}
	return nil
}

// UpdateUserDept 修改指定部门的角色
func (rm *UserDept) UpdateUserDept(ctx core.Context, req *entity.UserDept) error {
	if err := rm.repo.UpdateUserDept(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

func (rm *UserDept) DeleteUserDept(ctx core.Context, id uint32) error {
	if err := rm.repo.DeleteUserDept(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
