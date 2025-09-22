package dbs

import (
	"sync"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type RoleEntity struct{}

var (
	reIns  *RoleEntity
	reOnce sync.Once
)

func NewRoleEntity() *RoleEntity {
	reOnce.Do(func() {
		reIns = &RoleEntity{}
	})
	return reIns
}

// ListRoleEntity 获取列表
func (r RoleEntity) ListRoleEntity(ctx core.Context, req *types.ListRoleEntityRequest) ([]*entity.RoleEntity, uint32, error) {
	var (
		list  []*entity.RoleEntity
		total int64
	)

	db := ctx.DB().Model(entity.RoleEntity{}).
		Preload("Entity").
		Preload("Entity.App").
		Where("role_id=?", req.RoleId)

	if req.EntityId != nil {
		db = db.Where("entity_id = ?", *req.EntityId)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = page.SearchScopes(db, &req.Search)

	return list, uint32(total), db.Find(&list).Error
}

// CreateRoleEntity 创建数据
func (r RoleEntity) CreateRoleEntity(ctx core.Context, re *entity.RoleEntity) (uint32, error) {
	return re.Id, ctx.DB().Create(re).Error
}

// GetRoleEntity 获取指定的数据
func (r RoleEntity) GetRoleEntity(ctx core.Context, id uint32) (*entity.RoleEntity, error) {
	var (
		re = entity.RoleEntity{}
		fs = []string{"*"}
	)
	return &re, ctx.DB().Select(fs).First(&re, id).Error
}

// UpdateRoleEntity 更新数据
func (r RoleEntity) UpdateRoleEntity(ctx core.Context, re *entity.RoleEntity) error {
	return ctx.DB().Updates(re).Error
}

// DeleteRoleEntity 删除数据
func (r RoleEntity) DeleteRoleEntity(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.RoleEntity{}).Error
}
