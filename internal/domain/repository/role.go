package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Role interface {
	// GetRoleChildrenIds 获取指定觉得的下级所有角色id
	GetRoleChildrenIds(ctx core.Context, ids []uint32) ([]uint32, error)

	// GetRoleParentIds 获取指定觉得的下级所有角色id
	GetRoleParentIds(ctx core.Context, ids []uint32) ([]uint32, error)

	// GetRoleByKeyword 获取指定的部门信息
	GetRoleByKeyword(ctx core.Context, keyword string) (*entity.Role, error)

	// GetRole 获取指定的部门信息
	GetRole(ctx core.Context, id uint32) (*entity.Role, error)

	// ListRole 获取部门信息列表
	ListRole(ctx core.Context, req *types.ListRoleRequest) ([]*entity.Role, error)

	// CreateRole 创建部门信息
	CreateRole(ctx core.Context, req *entity.Role) (uint32, error)

	// UpdateRole 更新部门信息
	UpdateRole(ctx core.Context, req *entity.Role) error

	// DeleteRole 删除部门信息
	DeleteRole(ctx core.Context, id uint32) error
}
