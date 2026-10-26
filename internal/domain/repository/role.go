package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Role interface {
	// GetRoleChildrenIds 获取指定角色的所有下级角色id（role_closure 表为全局闭包，id 全局唯一，无需 app）
	GetRoleChildrenIds(ctx core.Context, ids []uint32) ([]uint32, error)

	// GetRoleParentIds 获取指定角色的所有上级角色id（role_closure 表为全局闭包，id 全局唯一，无需 app）
	GetRoleParentIds(ctx core.Context, ids []uint32) ([]uint32, error)

	// GetRoleByKeyword 按 App + Keyword 获取指定的角色信息
	GetRoleByKeyword(ctx core.Context, app string, keyword string) (*entity.Role, error)

	// GetRole 按 App + ID 获取指定的角色信息
	GetRole(ctx core.Context, app string, id uint32) (*entity.Role, error)

	// ListRole 获取指定应用的角色信息列表
	ListRole(ctx core.Context, app string, req *types.ListRoleRequest) ([]*entity.Role, error)

	// CreateRole 创建角色信息（app 由 req.App 指定）
	CreateRole(ctx core.Context, req *entity.Role) (uint32, error)

	// UpdateRole 更新角色信息（app 由 req.App 指定）
	UpdateRole(ctx core.Context, req *entity.Role) error

	// DeleteRole 删除指定应用的角色信息
	DeleteRole(ctx core.Context, app string, id uint32) error
}
