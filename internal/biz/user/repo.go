package user

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// GetUser 获取指定的用户信息
	GetUser(ctx kratosx.Context, id uint32) (*User, error)

	// ListUser 获取用户信息列表
	ListUser(ctx kratosx.Context, req *ListUserRequest) ([]*User, uint32, error)

	// CreateUser 创建用户信息
	CreateUser(ctx kratosx.Context, req *User) (uint32, error)

	// UpdateUser 更新用户信息
	UpdateUser(ctx kratosx.Context, req *User) error

	// UpdateUserStatus 更新用户信息状态
	UpdateUserStatus(ctx kratosx.Context, id uint32, status bool) error

	// DeleteUser 删除用户信息
	DeleteUser(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetUserByPhone 获取指定的用户信息
	GetUserByPhone(ctx kratosx.Context, phone string) (*User, error)

	// GetUserByEmail 获取指定的用户信息
	GetUserByEmail(ctx kratosx.Context, email string) (*User, error)

	// GetDepartmentDataScope 获取指定用户的部门权限
	GetDepartmentDataScope(ctx kratosx.Context, uid uint32) (bool, []uint32, error)

	// GetRoleDataScope 获取指定角色的角色权限
	GetRoleDataScope(ctx kratosx.Context, uid uint32) (bool, []uint32, error)

	// HasUserDataScope 获取某个用户是否具有另一个用户的权限
	HasUserDataScope(ctx kratosx.Context, pid, uid uint32) (bool, error)

	// GetUserToken 获取用户的token信息
	GetUserToken(ctx kratosx.Context, id uint32) (*string, int64, error)

	// GetUserPassword 获取用户的password
	GetUserPassword(ctx kratosx.Context, id uint32) (string, error)
}
