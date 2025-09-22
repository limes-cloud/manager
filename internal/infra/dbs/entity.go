package dbs

import (
	"fmt"
	"sync"

	"github.com/limes-cloud/kratosx/pkg/cache"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Entity struct {
	entCache *cache.Cache[string, *entity.Entity]
}

var (
	entityIns          *Entity
	entityOnce         sync.Once
	entityCacheKey     = "entity"
	entityRuleCacheKey = "entityrule"
)

func NewEntity() *Entity {
	entityOnce.Do(func() {
		entityIns = &Entity{}
	})
	return entityIns
}

func (infra *Entity) GetEntityRule(ctx core.Context, id uint32, name string) (*entity.EntityRule, error) {
	var (
		ent = entity.EntityRule{}
		fs  = []string{"*"}
	)
	return &ent, ctx.DB().
		Select(fs).
		Where("entity_id=? and name=?", id, name).
		First(&ent).Error
}

// ListEntityRule 获取列表
func (infra *Entity) ListEntityRule(ctx core.Context, req *types.ListEntityRuleRequest) ([]*entity.EntityRule, uint32, error) {
	var (
		list  []*entity.EntityRule
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.EntityRule{}).Where("entity_id = ?", req.EntityId).Select(fs)
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	return list, uint32(total), db.Order("id asc").Find(&list).Error
}

// CreateEntityRule 创建数据
func (infra *Entity) CreateEntityRule(ctx core.Context, tg *entity.EntityRule) (uint32, error) {
	return tg.Id, ctx.DB().Create(tg).Error
}

// UpdateEntityRule 更新数据
func (infra *Entity) UpdateEntityRule(ctx core.Context, tg *entity.EntityRule) error {
	return ctx.DB().Updates(tg).Error
}

// DeleteEntityRule 删除数据
func (infra *Entity) DeleteEntityRule(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.EntityRule{}).Error
}

func (infra *Entity) GetEntityField(ctx core.Context, id uint32, name string) (*entity.EntityField, error) {
	var (
		ent = entity.EntityField{}
		fs  = []string{"*"}
	)
	return &ent, ctx.DB().
		Select(fs).
		Where("entity_id=? and name=?", id, name).
		First(&ent).Error
}

// ListEntityField 获取列表
func (infra *Entity) ListEntityField(ctx core.Context, req *types.ListEntityFieldRequest) ([]*entity.EntityField, uint32, error) {
	var (
		list  []*entity.EntityField
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.EntityField{}).Where("entity_id = ?", req.EntityId).Select(fs)
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	return list, uint32(total), db.Order("`index` asc,`id` asc").Find(&list).Error
}

// CreateEntityField 创建数据
func (infra *Entity) CreateEntityField(ctx core.Context, tg *entity.EntityField) (uint32, error) {
	return tg.Id, ctx.DB().Create(tg).Error
}

// UpdateEntityField 更新数据
func (infra *Entity) UpdateEntityField(ctx core.Context, tg *entity.EntityField) error {
	return ctx.DB().Updates(tg).Error
}

// DeleteEntityField 删除数据
func (infra *Entity) DeleteEntityField(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.EntityField{}).Error
}

func (infra *Entity) GetEntityByName(ctx core.Context, db string, name string) (*entity.Entity, error) {
	var (
		ent = entity.Entity{}
		fs  = []string{"*"}
	)
	return &ent, ctx.DB().
		Select(fs).
		Where("`database` = ? and `name` = ?", db, name).
		First(&ent).Error
}

// GetEntity 获取指定的数据
func (infra *Entity) GetEntity(ctx core.Context, id uint32) (*entity.Entity, error) {
	var (
		ent = entity.Entity{}
		fs  = []string{"*"}
	)
	return &ent, ctx.DB().Select(fs).First(&ent, id).Error
}

// GetEntityByKeyword 获取指定数据
func (infra *Entity) GetEntityByKeyword(ctx core.Context, keyword string) (*entity.Entity, error) {
	var (
		m  = entity.Entity{}
		fs = []string{"*"}
	)
	return &m, ctx.DB().Select(fs).Where("keyword = ?", keyword).First(&m).Error
}

// ListEntity 获取列表 fixed code
func (infra *Entity) ListEntity(ctx core.Context, req *types.ListEntityRequest) ([]*entity.Entity, error) {
	var ms []*entity.Entity

	db := ctx.DB().Model(&entity.Entity{}).Where("app_id = ?", req.AppId)

	if req.Name != nil {
		db = db.Where("name LIKE ?", "%"+*req.Name+"%")
	}
	if req.Database != nil {
		db = db.Where("database = ?", *req.Database)
	}
	return ms, db.Find(&ms).Error
}

// CreateEntity 创建数据
func (infra *Entity) CreateEntity(ctx core.Context, req *entity.Entity) (uint32, error) {
	return req.Id, ctx.DB().Create(req).Error
}

// UpdateEntity 更新数据
func (infra *Entity) UpdateEntity(ctx core.Context, req *entity.Entity) error {
	// 查询当前id的详细数据
	var ent entity.Entity
	if err := ctx.DB().Model(&entity.Entity{}).First(&ent, req.Id).Error; err != nil {
		return err
	}

	// 修改
	op := infra.entCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Model(&entity.Entity{}).Updates(req).Error; err != nil {
			return err
		}
		if ent.Database != req.Database || ent.Name != req.Name {
			if err := op.Delete(infra.getEntityCacheKey(ent.Database, ent.Name)).Do(); err != nil {
				return err
			}
		}

		ent.Database = req.Database
		ent.Name = req.Name

		return op.Puts(map[string]*entity.Entity{
			infra.getEntityCacheKey(ent.Database, ent.Name): &ent,
		}).Do()
	})
}

// DeleteEntity 删除数据
func (infra *Entity) DeleteEntity(ctx core.Context, id uint32) error {
	// 查询当前id的详细数据
	var ent entity.Entity
	if err := ctx.DB().Model(&entity.Entity{}).First(&ent, id).Error; err != nil {
		return err
	}

	// 删除
	op := infra.entCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id = ?", id).Delete(&entity.Entity{}).Error; err != nil {
			return err
		}
		return op.Delete(infra.getEntityCacheKey(ent.Database, ent.Name)).Do()
	})
}

// load 加载全量的数据
func (infra *Entity) loadEntity(ctx core.Context) (map[string]*entity.Entity, error) {
	var list []*entity.Entity
	if err := ctx.DB().Model(&entity.Entity{}).Find(&list).Error; err != nil {
		return nil, err
	}
	bucket := map[string]*entity.Entity{}
	for _, item := range list {
		bucket[infra.getEntityCacheKey(item.Database, item.Name)] = item
	}
	return bucket, nil
}

// getCacheKey 获取缓存的key
func (infra *Entity) getEntityCacheKey(db string, mid string) string {
	return fmt.Sprintf("%s:%s", db, mid)
}
