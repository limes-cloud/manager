package repository

import "github.com/limes-cloud/manager/internal/core"

type Scope interface {
	IsSuperAdmin(ctx core.Context) bool

	// HasAppScope 判断是否具有应用权限
	HasAppScope(ctx core.Context, appId uint32) bool

	// HasDeptScope 判断是否具有部门权限
	HasDeptScope(ctx core.Context, entity string, deptId uint32) bool

	// HasRoleScope 判断是否具有角色权限
	HasRoleScope(ctx core.Context, roleId uint32) bool

	// AppScope 判断是否具有应用权限
	AppScope(ctx core.Context, tid uint32) []uint32

	// RoleScopes 获取应用的权限列表
	RoleScopes(ctx core.Context) []uint32

	// DeptScopes 获取部门的权限列表
	DeptScopes(ctx core.Context, entity string) (bool, []uint32)

	// RoleIds 用户当前用户的角色列表
	RoleIds(ctx core.Context) []uint32

	// DeptIds 获取当前用户的部门列表
	DeptIds(ctx core.Context) []uint32
}
