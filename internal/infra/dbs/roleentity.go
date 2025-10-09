package dbs

import (
	"context"
	"fmt"
	"sync"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/cache"

	"github.com/limes-cloud/manager/internal/core"

	"dario.cat/mergo"
	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type RoleEntity struct {
	cache *cache.Cache[string, *entity.RoleEntity]
}

var (
	reIns      *RoleEntity
	reOnce     sync.Once
	reCacheKey = "roleentity"
)

func NewRoleEntity() *RoleEntity {
	reOnce.Do(func() {
		reIns = &RoleEntity{}

		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "role entity"),
			kratosx.WithSkipDBHook(),
		)
		ctx.BeforeStart("load cache role entity", func() {
			reIns.cache = cache.NewCacheAndInit[string, *entity.RoleEntity](
				ctx,
				ctx.Redis(),
				reCacheKey,
				cache.HookLoad(func() (map[string]*entity.RoleEntity, error) {
					return reIns.load(ctx)
				}),
			)
		})
	})
	return reIns
}

// ListRoleEntity 获取列表
func (r *RoleEntity) ListRoleEntity(ctx core.Context, req *types.ListRoleEntityRequest) ([]*entity.RoleEntity, uint32, error) {
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
func (r *RoleEntity) CreateRoleEntity(ctx core.Context, req *entity.RoleEntity) (uint32, error) {
	op := r.cache.OP(ctx)
	err := ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(req).Error; err != nil {
			return err
		}
		return op.Put(r.getCacheKey(req.RoleId, req.EntityId, req.Action), req).Do()
	})
	if err != nil {
		return 0, err
	}
	return req.Id, nil
}

// GetRoleEntity 获取指定的数据
func (r *RoleEntity) GetRoleEntity(ctx core.Context, id uint32) (*entity.RoleEntity, error) {
	var (
		re = entity.RoleEntity{}
		fs = []string{"*"}
	)
	return &re, ctx.DB().Select(fs).First(&re, id).Error
}

// UpdateRoleEntity 更新数据
func (r *RoleEntity) UpdateRoleEntity(ctx core.Context, req *entity.RoleEntity) error {
	// 查询当前id的详细数据
	var re entity.RoleEntity
	if err := ctx.DB().First(&re, req.Id).Error; err != nil {
		return err
	}
	newReq := *req
	_ = mergo.Merge(&newReq, re)

	op := r.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Updates(req).Error; err != nil {
			return err
		}
		if newReq.RoleId != re.RoleId || newReq.EntityId != re.EntityId || newReq.Action != re.Action {
			op = op.Delete(r.getCacheKey(re.RoleId, re.EntityId, re.Action))
		}
		return op.Put(r.getCacheKey(newReq.RoleId, newReq.EntityId, newReq.Action), &newReq).Do()
	})
}

// DeleteRoleEntity 删除数据
func (r *RoleEntity) DeleteRoleEntity(ctx core.Context, id uint32) error {
	// 查询当前id的详细数据
	var re entity.RoleEntity
	if err := ctx.DB().First(&re, id).Error; err != nil {
		return err
	}

	op := r.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id = ?", id).Delete(&entity.RoleEntity{}).Error; err != nil {
			return err
		}
		return op.Delete(r.getCacheKey(re.RoleId, re.EntityId, re.Action)).Do()
	})
}

func (r *RoleEntity) load(ctx core.Context) (map[string]*entity.RoleEntity, error) {
	var list []*entity.RoleEntity
	if err := ctx.DB().Find(&list).Error; err != nil {
		return nil, err
	}
	bucket := map[string]*entity.RoleEntity{}
	for _, item := range list {
		bucket[r.getCacheKey(item.RoleId, item.EntityId, item.Action)] = item
	}
	return bucket, nil
}

func (r *RoleEntity) getCacheKey(rid uint32, eid uint32, act string) string {
	return fmt.Sprintf("%d:%d:%s", rid, eid, act)
}

func (r *RoleEntity) GetRoleEntityByRDA(rid uint32, eid uint32, act string) (*entity.RoleEntity, bool) {
	return r.cache.Get(r.getCacheKey(rid, eid, act))
}
