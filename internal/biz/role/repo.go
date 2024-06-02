package role

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// ListRole 获取角色信息列表
	ListRole(ctx kratosx.Context, req *ListRoleRequest) ([]*Role, uint32, error)

	// CreateRole 创建角色信息
	CreateRole(ctx kratosx.Context, req *Role) (uint32, error)

	// UpdateRole 更新角色信息
	UpdateRole(ctx kratosx.Context, req *Role) error

	// UpdateRoleStatus 更新角色信息状态
	UpdateRoleStatus(ctx kratosx.Context, id uint32, status bool) error

	// DeleteRole 删除角色信息
	DeleteRole(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetRoleParentIds 获取父角色信息ID列表
	GetRoleParentIds(ctx kratosx.Context, id uint32) ([]uint32, error)

	// GetRoleChildrenIds 获取子角色信息ID列表
	GetRoleChildrenIds(ctx kratosx.Context, id uint32) ([]uint32, error)

	// GetRoleMenuIds 获取指定角色的菜单id列表
	GetRoleMenuIds(ctx kratosx.Context, id uint32) ([]uint32, error)

	// UpdateRoleMenu 更新角色的菜单
	UpdateRoleMenu(ctx kratosx.Context, roleId uint32, menuIds []uint32) error

	// GetRole 获取指定的角色信息
	GetRole(ctx kratosx.Context, id uint32) (*Role, error)

	// GetRoleByKeyword 获取指定的角色信息
	GetRoleByKeyword(ctx kratosx.Context, keyword string) (*Role, error)

	// GetRoleDataScope 获取指定角色的权限
	GetRoleDataScope(ctx kratosx.Context, rid uint32) (bool, []uint32, error)
}
