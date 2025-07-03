package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Role interface {
	// ListCacheRole 获取缓存的role
	ListCacheRole(ctx kratosx.Context) ([]*entity.Role, error)

	// GetRole 获取指定的角色信息
	GetRole(ctx kratosx.Context, id uint32) (*entity.Role, error)

	// GetRoleByKeyword 获取指定的角色信息
	GetRoleByKeyword(ctx kratosx.Context, keyword string) (*entity.Role, error)

	// GetRoleChildrenKeywords 获取指定角色的下级所有keyword
	GetRoleChildrenKeywords(ctx kratosx.Context, id uint32) ([]string, error)

	// ListRole 获取角色信息列表
	ListRole(ctx kratosx.Context, req *types.ListRoleRequest) ([]*entity.Role, error)

	// CreateRole 创建角色信息
	CreateRole(ctx kratosx.Context, req *entity.Role) (uint32, error)

	// UpdateRole 更新角色信息
	UpdateRole(ctx kratosx.Context, req *entity.Role) error

	// UpdateRoleStatus 更新角色信息状态
	UpdateRoleStatus(ctx kratosx.Context, id uint32, status bool) error

	// DeleteRole 删除角色信息
	DeleteRole(ctx kratosx.Context, id uint32) error

	// GetRoleMenuIds 获取指定角色的菜单id列表
	GetRoleMenuIds(ctx kratosx.Context, rids []uint32) ([]uint32, error)

	// UpdateRoleMenu 更新角色的菜单
	UpdateRoleMenu(ctx kratosx.Context, id uint32, menuIds []uint32) error

	// GetRoleScopeByUID 获取指定角色组的角色权限
	GetRoleScopeByUID(ctx kratosx.Context, uid uint32) (bool, []uint32, error)

	// GetRoleScope 获取指定角色组的角色权限
	GetRoleScope(ctx kratosx.Context, rids []uint32) (bool, []uint32, error)

	// HasRolePurview 用户是否具有指定的角色权限
	HasRolePurview(ctx kratosx.Context, uid uint32, rids []uint32) (bool, error)

	// AllRoleKeywordByMenuId 获取指定菜单id的所有角色keyword
	AllRoleKeywordByMenuId(ctx kratosx.Context, id uint32) ([]string, error)

	// GetDataScope 获取数据权限列表
	GetDataScope(ctx kratosx.Context, uid uint32) (bool, []uint32, error)

	// HasDataPurview 用户是否具有指定的数据权限
	HasDataPurview(ctx kratosx.Context, uid uint32, rids []uint32) (bool, error)

	GetEnableRoleIdsByUID(ctx kratosx.Context, uid uint32) ([]uint32, error)
}
