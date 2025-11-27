package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
)

type UserOAuther interface {
	// CreateUserOAuther 创建用户授权信息
	CreateUserOAuther(ctx core.Context, req *entity.UserOAuther) (uint32, error)

	// DeleteUserOAuther 删除用户授权信息
	DeleteUserOAuther(ctx core.Context, userId uint32, appId uint32) error

	// IsBindUserOAuther 是否绑定授权信息
	IsBindUserOAuther(ctx core.Context, channel uint32, oid string) bool

	// GetUserOAutherByCO 获取指定的授权信息
	GetUserOAutherByCO(ctx core.Context, cid uint32, oid string) (*entity.UserOAuther, error)

	// UpdateUserOAuther 更新用户授权信息
	UpdateUserOAuther(ctx core.Context, req *entity.UserOAuther) error
}
