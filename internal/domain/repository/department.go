package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Department interface {
	// ListDepartmentClassify 获取部门分类列表
	ListDepartmentClassify(ctx kratosx.Context, req *types.ListDepartmentClassifyRequest) ([]*entity.DepartmentClassify, uint32, error)

	// CreateDepartmentClassify 创建部门分类
	CreateDepartmentClassify(ctx kratosx.Context, req *entity.DepartmentClassify) (uint32, error)

	// UpdateDepartmentClassify 更新部门分类
	UpdateDepartmentClassify(ctx kratosx.Context, req *entity.DepartmentClassify) error

	// DeleteDepartmentClassify 删除部门分类
	DeleteDepartmentClassify(ctx kratosx.Context, id uint32) error

	// GetDepartment 获取指定的部门信息
	GetDepartment(ctx kratosx.Context, id uint32) (*entity.Department, error)

	// GetDepartmentByKeyword 获取指定的部门信息
	GetDepartmentByKeyword(ctx kratosx.Context, keyword string) (*entity.Department, error)

	// ListDepartment 获取部门信息列表
	ListDepartment(ctx kratosx.Context, req *types.ListDepartmentRequest) ([]*entity.Department, error)

	// CreateDepartment 创建部门信息
	CreateDepartment(ctx kratosx.Context, req *entity.Department) (uint32, error)

	// UpdateDepartment 更新部门信息
	UpdateDepartment(ctx kratosx.Context, req *entity.Department) error

	// DeleteDepartment 删除部门信息
	DeleteDepartment(ctx kratosx.Context, id uint32) error

	// GetDepartmentDataScope 获取指定用户的部门权限
	GetDepartmentDataScope(ctx kratosx.Context, uid uint32) (bool, []uint32, error)

	// HasDepartmentPurview 是否具有指定的部门权限
	HasDepartmentPurview(ctx kratosx.Context, uid uint32, did uint32) (bool, error)
}
