package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type User interface {
	GetUserDeptId(id uint32) uint32

	// GetUser 获取指定的用户信息
	GetUser(ctx core.Context, id uint32) (*entity.User, error)

	// GetUserByUsername 获取指定的用户信息
	GetUserByUsername(ctx core.Context, un string) (*entity.User, error)

	// GetUserByTU 获取指定的用户信息
	GetUserByTU(ctx core.Context, tid uint32, un string) (*entity.User, error)

	// ListUser 获取用户信息列表
	ListUser(ctx core.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error)

	// CreateUser 创建用户信息
	CreateUser(ctx core.Context, req *entity.User) (uint32, error)

	// UpdateUser 更新用户信息
	UpdateUser(ctx core.Context, req *entity.User) error

	// DeleteUser 删除用户信息
	DeleteUser(ctx core.Context, id uint32) error
}
