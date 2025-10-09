package dbs

import (
	"context"
	"fmt"
	"sync"

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
			kratosx.WithSkipDBHook(),
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

func (infra *Entity) GetEntityFieldName(id uint32) (string, bool) {
	return infra.fieldCache.Get(id)
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
	op := infra.entCache.OP(ctx)
	err := ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(req).Error; err != nil {
			return err
		}
		return op.Put(infra.getEntityCacheKey(req.Database, req.Name), req.Id).Do()
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
			return op.Delete(infra.getEntityCacheKey(ent.Database, ent.Name)).
				Put(infra.getEntityCacheKey(newReq.Database, newReq.Name), newReq.Id).Do()
		}
		return nil
	})
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
		return op.Delete(infra.getEntityCacheKey(ent.Database, ent.Name)).Do()
	})
}

func (infra *Entity) GetEntityIdByName(db string, name string) (uint32, bool) {
	return infra.entCache.Get(infra.getEntityCacheKey(db, name))
}

// load 加载全量的数据
func (infra *Entity) loadEntity(ctx core.Context) (map[string]uint32, error) {
	var list []*entity.Entity
	if err := ctx.DB().Find(&list).Error; err != nil {
		return nil, err
	}
	bucket := map[string]uint32{}
	for _, item := range list {
		bucket[infra.getEntityCacheKey(item.Database, item.Name)] = item.Id
	}
	return bucket, nil
}

// getCacheKey 获取缓存的key
func (infra *Entity) getEntityCacheKey(db string, mid string) string {
	return fmt.Sprintf("%s:%s", db, mid)
}
