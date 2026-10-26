package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type UserDept interface {
	// GetDeptIdsByUserId 获取用户部门id
	GetDeptIdsByUserId(ctx core.Context, app string, userId uint32) []uint32

	// GetUserMainDeptId 获取指定用户的主部门
	GetUserMainDeptId(ctx core.Context, app string, uid uint32) uint32

	// GetUserMainJobId 获取指定用户的主岗位
	GetUserMainJobId(ctx core.Context, app string, uid uint32) uint32

	// ListDeptUser 获取角色的菜单列表
	ListDeptUser(ctx core.Context, app string, req *types.ListDeptUserRequest) ([]*entity.User, uint32, error)

	// ListUserDept 获取指定菜单的角色列表
	ListUserDept(ctx core.Context, app string, req *types.ListUserDeptRequest) ([]*entity.UserDept, uint32, error)

	// CreateUserDept 创建角色菜单的关联信息
	CreateUserDept(ctx core.Context, ud *entity.UserDept) error

	// UpdateUserDept 更新角色菜单的关联信息
	UpdateUserDept(ctx core.Context, ud *entity.UserDept) error

	// DeleteUserDept 删除角色菜单关联信息
	DeleteUserDept(ctx core.Context, app string, id uint32) error
}
