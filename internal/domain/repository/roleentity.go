package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type RoleEntity interface {
	// ListRoleEntity 获取部门信息列表
	ListRoleEntity(ctx core.Context, app string, req *types.ListRoleEntityRequest) ([]*entity.RoleEntity, uint32, error)

	// GetRoleEntity 获取指定的数据权限
	GetRoleEntity(ctx core.Context, app string, id uint32) (*entity.RoleEntity, error)

	// CreateRoleEntity 创建部门信息
	CreateRoleEntity(ctx core.Context, req *entity.RoleEntity) (uint32, error)

	// UpdateRoleEntity 更新部门信息
	UpdateRoleEntity(ctx core.Context, req *entity.RoleEntity) error

	// DeleteRoleEntity 删除部门信息
	DeleteRoleEntity(ctx core.Context, app string, id uint32) error
}
