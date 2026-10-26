package dbs

import (
	"context"
	"fmt"
	"sync"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/samber/lo"

	"dario.cat/mergo"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/kratosx/pkg/cache"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Entity struct {
	entCache   *cache.Cache[string, uint32]
	fieldCache *cache.Cache[uint32, string]
	ruleCache  *cache.Cache[uint32, string]
}

var (
	entityIns           *Entity
	entityOnce          sync.Once
	entityCacheKey      = "entity"
	entityFieldCacheKey = "entityfield"
	entityRuleCacheKey  = "entityrule"
)

func NewEntity() *Entity {
	entityOnce.Do(func() {
		entityIns = &Entity{}
		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "entity"),
		)
		ctx.BeforeStart("load entity", func() {
			entityIns.entCache = cache.NewCacheAndInit[string, uint32](
				ctx,
				ctx.Redis(),
				entityCacheKey,
				cache.HookLoad(func() (map[string]uint32, error) {
					return entityIns.loadEntity(ctx)
				}),
			)
			entityIns.fieldCache = cache.NewCacheAndInit[uint32, string](
				ctx,
				ctx.Redis(),
				entityFieldCacheKey,
				cache.HookLoad(func() (map[uint32]string, error) {
					return entityIns.loadEntityField(ctx)
				}),
			)
			entityIns.ruleCache = cache.NewCacheAndInit[uint32, string](
				ctx,
				ctx.Redis(),
				entityRuleCacheKey,
				cache.HookLoad(func() (map[uint32]string, error) {
					return entityIns.loadEntityRule(ctx)
				}),
			)
		})
	})
	return entityIns
}

func (infra *Entity) GetEntityRuleById(id uint32) (string, bool) {
	return infra.ruleCache.Get(id)
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
	)

	db := ctx.DB().Model(entity.EntityRule{})
	if req.EntityId != nil {
		db = db.Where("entity_id = ?", req.EntityId)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	if req.Search != nil {
		// 查询条件下数据总数
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}

		// 分页排序
		db = page.SearchScopes(db, req.Search)
	}

	return list, uint32(total), db.Order("id asc").Find(&list).Error
}

// CreateEntityRule 创建数据
func (infra *Entity) CreateEntityRule(ctx core.Context, tg *entity.EntityRule) (uint32, error) {
	op := infra.ruleCache.OP(ctx)
	err := ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(tg).Error; err != nil {
			return err
		}
		return op.Put(tg.Id, tg.Expression).Do()
	})
	if err != nil {
		return 0, err
	}
	return tg.Id, nil
}

// UpdateEntityRule 更新数据
func (infra *Entity) UpdateEntityRule(ctx core.Context, tg *entity.EntityRule) error {
	op := infra.ruleCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Updates(tg).Error; err != nil {
			return err
		}
		if tg.Expression != "" {
			return op.Put(tg.Id, tg.Expression).Do()
		}
		return nil
	})
}

// DeleteEntityRule 删除数据
func (infra *Entity) DeleteEntityRule(ctx core.Context, id uint32) error {
	op := infra.ruleCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id = ?", id).Delete(&entity.EntityRule{}).Error; err != nil {
			return err
		}
		return op.Delete(id).Do()
	})
}

// DeleteEntityRules 删除多条实体规则
func (infra *Entity) DeleteEntityRules(ctx core.Context, ids []uint32) error {
	ids = lo.Uniq(ids)
	if len(ids) == 0 {
		return nil
	}
	op := infra.ruleCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id in ?", ids).Delete(&entity.EntityRule{}).Error; err != nil {
			return err
		}
		for _, id := range ids {
			op = op.Delete(id)
		}
		return op.Do()
	})
}

// UpdateRegisteredEntityRule 更新注册实体规则
func (infra *Entity) UpdateRegisteredEntityRule(ctx core.Context, tg *entity.EntityRule) error {
	op := infra.ruleCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Model(&entity.EntityRule{}).
			Where("id = ?", tg.Id).
			Omit("name").
			Updates(tg).Error; err != nil {
			return err
		}
		if tg.Expression != "" {
			return op.Put(tg.Id, tg.Expression).Do()
		}
		return nil
	})
}

func (infra *Entity) loadEntityRule(ctx core.Context) (map[uint32]string, error) {
	var list []*entity.EntityRule
	if err := ctx.DB().Find(&list).Error; err != nil {
		return nil, err
	}
	bucket := map[uint32]string{}
	for _, item := range list {
		bucket[item.Id] = item.Expression
	}
	return bucket, nil
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
	op := infra.fieldCache.OP(ctx)
	err := ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(tg).Error; err != nil {
			return err
		}
		return op.Put(tg.Id, tg.Name).Do()
	})
	if err != nil {
		return 0, err
	}
	return tg.Id, nil
}

// UpdateEntityField 更新数据
func (infra *Entity) UpdateEntityField(ctx core.Context, tg *entity.EntityField) error {
	op := infra.fieldCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Updates(tg).Error; err != nil {
			return err
		}
		if tg.Name != "" {
			return op.Put(tg.Id, tg.Name).Do()
		}
		return nil
	})
}

// DeleteEntityField 删除数据
func (infra *Entity) DeleteEntityField(ctx core.Context, id uint32) error {
	op := infra.fieldCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id = ?", id).Delete(&entity.EntityField{}).Error; err != nil {
			return err
		}
		return op.Delete(id).Do()
	})
}

// DeleteEntityFields 删除多条实体字段
func (infra *Entity) DeleteEntityFields(ctx core.Context, ids []uint32) error {
	ids = lo.Uniq(ids)
	if len(ids) == 0 {
		return nil
	}
	op := infra.fieldCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id in ?", ids).Delete(&entity.EntityField{}).Error; err != nil {
			return err
		}
		for _, id := range ids {
			op = op.Delete(id)
		}
		return op.Do()
	})
}

// UpdateRegisteredEntityField 更新注册实体字段
func (infra *Entity) UpdateRegisteredEntityField(ctx core.Context, tg *entity.EntityField) error {
	return ctx.DB().Model(&entity.EntityField{}).
		Where("id = ?", tg.Id).
		Select("comment", "index").
		Updates(tg).Error
}

func (infra *Entity) loadEntityField(ctx core.Context) (map[uint32]string, error) {
	var list []*entity.EntityField
	if err := ctx.DB().Find(&list).Error; err != nil {
		return nil, err
	}
	bucket := map[uint32]string{}
	for _, item := range list {
		bucket[item.Id] = item.Name
	}
	return bucket, nil
}

// GetEntity 获取指定的数据
func (infra *Entity) GetEntity(ctx core.Context, id uint32) (*entity.Entity, error) {
	var (
		ent = entity.Entity{}
		fs  = []string{"*"}
	)
	return &ent, ctx.DB().Select(fs).First(&ent, id).Error
}

// GetEntityByName 获取指定应用下的实体
func (infra *Entity) GetEntityByName(ctx core.Context, app string, db string, name string) (*entity.Entity, error) {
	ent := entity.Entity{}
	return &ent, ctx.DB().
		Where("app = ? AND `database` = ? AND name = ?", app, db, name).
		First(&ent).Error
}

// ListEntity 获取列表 fixed code
func (infra *Entity) ListEntity(ctx core.Context, req *types.ListEntityRequest) ([]*entity.Entity, error) {
	var ms []*entity.Entity

	db := ctx.DB().Model(&entity.Entity{}).Where("app = ?", req.App)

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
	op := infra.entCache.OP(ctx)
	err := ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(req).Error; err != nil {
			return err
		}
		return op.Put(infra.getEntityCacheKey(req.App, req.Database, req.Name), req.Id).Do()
	})
	if err != nil {
		return 0, err
	}
	return req.Id, nil
}

// UpdateEntity 更新数据
func (infra *Entity) UpdateEntity(ctx core.Context, req *entity.Entity) error {
	// 查询当前id的详细数据
	var ent entity.Entity
	if err := ctx.DB().First(&ent, req.Id).Error; err != nil {
		return err
	}

	newReq := *req
	_ = mergo.Merge(&newReq, ent)

	op := infra.entCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Updates(req).Error; err != nil {
			return err
		}
		if newReq.Database != ent.Database || newReq.Name != ent.Name {
			return op.Delete(infra.getEntityCacheKey(ent.App, ent.Database, ent.Name)).
				Put(infra.getEntityCacheKey(newReq.App, newReq.Database, newReq.Name), newReq.Id).Do()
		}
		return nil
	})
}

// UpdateRegisteredEntity 更新注册实体
func (infra *Entity) UpdateRegisteredEntity(ctx core.Context, req *entity.Entity) error {
	return ctx.DB().Model(&entity.Entity{}).
		Where("id = ?", req.Id).
		Omit("app", "database", "name").
		Updates(req).Error
}

// DeleteEntity 删除数据
func (infra *Entity) DeleteEntity(ctx core.Context, id uint32) error {
	// 查询当前id的详细数据
	var ent entity.Entity
	if err := ctx.DB().First(&ent, id).Error; err != nil {
		return err
	}

	// 删除
	op := infra.entCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id = ?", id).Delete(&entity.Entity{}).Error; err != nil {
			return err
		}
		return op.Delete(infra.getEntityCacheKey(ent.App, ent.Database, ent.Name)).Do()
	})
}

// DeleteEntities 删除多条实体数据
func (infra *Entity) DeleteEntities(ctx core.Context, ids []uint32) error {
	ids = lo.Uniq(ids)
	if len(ids) == 0 {
		return nil
	}

	var ents []*entity.Entity
	if err := ctx.DB().Where("id in ?", ids).Find(&ents).Error; err != nil {
		return err
	}
	var fields []*entity.EntityField
	if err := ctx.DB().Where("entity_id in ?", ids).Find(&fields).Error; err != nil {
		return err
	}
	var rules []*entity.EntityRule
	if err := ctx.DB().Where("entity_id in ?", ids).Find(&rules).Error; err != nil {
		return err
	}

	op := infra.entCache.OP(ctx)
	fieldOp := infra.fieldCache.OP(ctx)
	ruleOp := infra.ruleCache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id in ?", ids).Delete(&entity.Entity{}).Error; err != nil {
			return err
		}
		for _, item := range ents {
			op = op.Delete(infra.getEntityCacheKey(item.App, item.Database, item.Name))
		}
		for _, item := range fields {
			fieldOp = fieldOp.Delete(item.Id)
		}
		for _, item := range rules {
			ruleOp = ruleOp.Delete(item.Id)
		}
		if err := op.Do(); err != nil {
			return err
		}
		if err := fieldOp.Do(); err != nil {
			return err
		}
		return ruleOp.Do()
	})
}

func (infra *Entity) GetEntityIdByName(ctx core.Context, app string, db string, name string) (uint32, bool) {
	return infra.entCache.Get(infra.getEntityCacheKey(app, db, name))
}

// load 加载全量的数据
func (infra *Entity) loadEntity(ctx core.Context) (map[string]uint32, error) {
	var list []*entity.Entity
	if err := ctx.DB().Find(&list).Error; err != nil {
		return nil, err
	}
	bucket := map[string]uint32{}
	for _, item := range list {
		bucket[infra.getEntityCacheKey(item.App, item.Database, item.Name)] = item.Id
	}
	return bucket, nil
}

// getCacheKey 获取缓存的key
func (infra *Entity) getEntityCacheKey(app string, db string, mid string) string {
	return fmt.Sprintf("%s:%s:%s", app, db, mid)
}
