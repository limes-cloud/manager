package menu

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// ListMenu 获取菜单信息列表
	ListMenu(ctx kratosx.Context, req *ListMenuRequest) ([]*Menu, uint32, error)

	// ListMenuByRoleId 获取指定角色的菜单列表
	ListMenuByRoleId(ctx kratosx.Context, id uint32) ([]*Menu, uint32, error)

	// CreateMenu 创建菜单信息
	CreateMenu(ctx kratosx.Context, req *Menu) (uint32, error)

	// UpdateMenu 更新菜单信息
	UpdateMenu(ctx kratosx.Context, req *Menu) error

	// DeleteMenu 删除菜单信息
	DeleteMenu(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetMenuParentIds 获取父菜单信息ID列表
	GetMenuParentIds(ctx kratosx.Context, id uint32) ([]uint32, error)

	// GetMenuChildrenIds 获取子菜单信息ID列表
	GetMenuChildrenIds(ctx kratosx.Context, id uint32) ([]uint32, error)

	// InitBasicMenu 初始化基础菜单api
	InitBasicMenu(ctx kratosx.Context)
}
