package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type UserDept interface {
	// ListDeptUser 获取角色的菜单列表
	ListDeptUser(ctx core.Context, req *types.ListDeptUserRequest) ([]*entity.User, uint32, error)
	// ListUserDept 获取指定菜单的角色列表
	ListUserDept(ctx core.Context, req *types.ListUserDeptRequest) ([]*entity.UserDept, uint32, error)
	// CreateUserDept 创建角色菜单的关联信息
	CreateUserDept(ctx core.Context, ud *entity.UserDept) error
	// DeleteUserDept 删除角色菜单关联信息
	DeleteUserDept(ctx core.Context, ud *entity.UserDept) error
}
