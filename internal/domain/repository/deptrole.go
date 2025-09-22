package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type DeptRole interface {
	// ListRoleDept 获取角色的菜单列表
	ListRoleDept(ctx core.Context, req *types.ListRoleDeptRequest) ([]*entity.Dept, uint32, error)
	// ListDeptRole 获取指定菜单的角色列表
	ListDeptRole(ctx core.Context, req *types.ListDeptRoleRequest) ([]*entity.Role, uint32, error)
	// CreateDeptRoles 创建角色菜单的关联信息
	CreateDeptRoles(ctx core.Context, drs []*entity.DeptRole) error
	// DeleteDeptRoles 删除角色菜单关联信息
	DeleteDeptRoles(ctx core.Context, drs []*entity.DeptRole) error
}
