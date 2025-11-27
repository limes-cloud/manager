package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type DeptRole interface {
	// ListDeptRole 获取部门角色列表
	ListDeptRole(ctx core.Context, req *types.ListDeptRoleRequest) ([]*entity.Role, uint32, error)
	// CreateDeptRole 创建部门角色信息
	CreateDeptRole(ctx core.Context, deptId, roleId uint32) error
	// DeleteDeptRole 删除部门角色信息
	DeleteDeptRole(ctx core.Context, deptId, roleId uint32) error
}
