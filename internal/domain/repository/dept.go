package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Dept interface {
	// ListDeptClassify 获取部门分类列表
	ListDeptClassify(ctx core.Context, req *types.ListDeptClassifyRequest) ([]*entity.DeptClassify, uint32, error)

	// CreateDeptClassify 创建部门分类
	CreateDeptClassify(ctx core.Context, req *entity.DeptClassify) (uint32, error)

	// UpdateDeptClassify 更新部门分类
	UpdateDeptClassify(ctx core.Context, req *entity.DeptClassify) error

	// DeleteDeptClassify 删除部门分类
	DeleteDeptClassify(ctx core.Context, id uint32) error

	// GetDept 获取指定的部门信息
	GetDept(ctx core.Context, id uint32) (*entity.Dept, error)

	// ListDept 获取部门信息列表
	ListDept(ctx core.Context, req *types.ListDeptRequest) ([]*entity.Dept, error)

	// CreateDept 创建部门信息
	CreateDept(ctx core.Context, req *entity.Dept) (uint32, error)

	// UpdateDept 更新部门信息
	UpdateDept(ctx core.Context, req *entity.Dept) error

	// DeleteDept 删除部门信息
	DeleteDept(ctx core.Context, id uint32) error
}
