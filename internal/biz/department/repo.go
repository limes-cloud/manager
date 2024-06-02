package department

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// ListDepartment 获取部门信息列表
	ListDepartment(ctx kratosx.Context, req *ListDepartmentRequest) ([]*Department, uint32, error)

	// CreateDepartment 创建部门信息
	CreateDepartment(ctx kratosx.Context, req *Department) (uint32, error)

	// UpdateDepartment 更新部门信息
	UpdateDepartment(ctx kratosx.Context, req *Department) error

	// DeleteDepartment 删除部门信息
	DeleteDepartment(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetDepartmentParentIds 获取父部门信息ID列表
	GetDepartmentParentIds(ctx kratosx.Context, id uint32) ([]uint32, error)

	// GetDepartmentChildrenIds 获取子部门信息ID列表
	GetDepartmentChildrenIds(ctx kratosx.Context, id uint32) ([]uint32, error)

	// GetDepartment 获取指定的部门信息
	GetDepartment(ctx kratosx.Context, id uint32) (*Department, error)

	// GetDepartmentByKeyword 获取指定的部门信息
	GetDepartmentByKeyword(ctx kratosx.Context, keyword string) (*Department, error)

	// GetDepartmentDataScope 获取指定用户的部门权限
	GetDepartmentDataScope(ctx kratosx.Context, uid uint32) (bool, []uint32, error)
}
