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

	// GetMenuIdsByRoleIds 通过角色id获取菜单id列表
	GetMenuIdsByRoleIds(ctx core.Context, app string, rids []uint32) []uint32

	// GetRoleIdsByMenuIds 通过菜单id获取角色id列表
	GetRoleIdsByMenuIds(ctx core.Context, app string, mids []uint32) []uint32

	// DeleteRoles 批量删除角色关联的菜单
	DeleteRoles(ctx core.Context, app string, rids []uint32) error

	// DeleteMenus 批量删除菜单关联的角色
	DeleteMenus(ctx core.Context, app string, mids []uint32) error
}
