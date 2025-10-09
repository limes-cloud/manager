package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/types"
)

type Scope interface {
	GetScope(ctx core.Context, database string, entity string, action string) *types.GetScopeResponse

	HasMenuScope(ctx core.Context, menuId uint32) bool

	HasAppScope(ctx core.Context, roleId uint32) bool

	// HasRoleScope 判断是否具有角色权限
	HasRoleScope(ctx core.Context, roleId uint32) bool

	// RoleScopes 获取应用的权限列表
	RoleScopes(ctx core.Context) []uint32

	// DeptScopes 获取部门的权限列表
	DeptScopes(ctx core.Context, database, entity, action string) (bool, []uint32)

	// RoleIds 用户当前用户的角色列表
	RoleIds(ctx core.Context) []uint32

	// DeptIds 获取当前用户的部门列表
	DeptIds(ctx core.Context) []uint32
}
