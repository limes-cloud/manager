package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Entity interface {
	// GetEntityIdByName 查询部门信息
	GetEntityIdByName(db string, name string) (uint32, bool)

	// GetEntity 查询部门信息
	GetEntity(ctx core.Context, id uint32) (*entity.Entity, error)

	// ListEntity 获取部门信息列表
	ListEntity(ctx core.Context, req *types.ListEntityRequest) ([]*entity.Entity, error)

	// CreateEntity 创建部门信息
	CreateEntity(ctx core.Context, req *entity.Entity) (uint32, error)

	// UpdateEntity 更新部门信息
	UpdateEntity(ctx core.Context, req *entity.Entity) error

	// DeleteEntity 删除部门信息
	DeleteEntity(ctx core.Context, id uint32) error

	// ListEntityField 获取部门分类列表
	ListEntityField(ctx core.Context, req *types.ListEntityFieldRequest) ([]*entity.EntityField, uint32, error)

	// CreateEntityField 创建部门分类
	CreateEntityField(ctx core.Context, req *entity.EntityField) (uint32, error)

	// UpdateEntityField 更新部门分类
	UpdateEntityField(ctx core.Context, req *entity.EntityField) error

	// DeleteEntityField 删除部门分类
	DeleteEntityField(ctx core.Context, id uint32) error

	GetEntityField(ctx core.Context, id uint32, name string) (*entity.EntityField, error)

	// ListEntityRule 获取实体规则列表
	ListEntityRule(ctx core.Context, req *types.ListEntityRuleRequest) ([]*entity.EntityRule, uint32, error)

	// CreateEntityRule 创建实体规则
	CreateEntityRule(ctx core.Context, req *entity.EntityRule) (uint32, error)

	// UpdateEntityRule 更新实体规则
	UpdateEntityRule(ctx core.Context, req *entity.EntityRule) error

	// DeleteEntityRule 删除实体规则
	DeleteEntityRule(ctx core.Context, id uint32) error

	GetEntityRule(ctx core.Context, id uint32, name string) (*entity.EntityRule, error)
}
