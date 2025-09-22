package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
)

type RoleMenu interface {
	// CreateRoleMenus 创建角色菜单的关联信息
	CreateRoleMenus(ctx core.Context, rms []*entity.RoleMenu) error

	// DeleteRoleMenus 删除角色菜单关联信息
	DeleteRoleMenus(ctx core.Context, rms []*entity.RoleMenu) error

	// GetMenuIdsByRoleIds 获取角色id
	GetMenuIdsByRoleIds(mids []uint32) []uint32

	// GetRoleIdsByMenuIds 获取菜单id
	GetRoleIdsByMenuIds(rids []uint32) []uint32

	// DeleteRoles 批量删除角色
	DeleteRoles(ctx core.Context, rids []uint32) error

	// DeleteMenus 批量删除菜单
	DeleteMenus(ctx core.Context, rids []uint32) error
}
