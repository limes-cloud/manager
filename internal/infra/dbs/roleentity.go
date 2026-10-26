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
func (r *RoleEntity) ListRoleEntity(ctx core.Context, app string, req *types.ListRoleEntityRequest) ([]*entity.RoleEntity, uint32, error) {
	var (
		list  []*entity.RoleEntity
		total int64
	)

	db := ctx.DB().Model(entity.RoleEntity{}).
		Preload("Entity").
		Where("role_id=?", req.RoleId).
		Where("role_entity.app = ?", app)

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
	if req.App == "" {
		return 0, fmt.Errorf("app is required")
	}
	if err := r.validate(ctx, req.App, req.RoleId, req.EntityId); err != nil {
		return 0, err
	}
	op := r.cache.OP(ctx)
	err := ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(req).Error; err != nil {
			return err
		}
		return op.Put(r.getCacheKey(req), req).Do()
	})
	if err != nil {
		return 0, err
	}
	return req.Id, nil
}

// GetRoleEntity 获取指定的数据
func (r *RoleEntity) GetRoleEntity(ctx core.Context, app string, id uint32) (*entity.RoleEntity, error) {
	var (
		re = entity.RoleEntity{}
		fs = []string{"*"}
	)
	return &re, ctx.DB().Select(fs).Where("app = ?", app).First(&re, id).Error
}

// UpdateRoleEntity 更新数据
func (r *RoleEntity) UpdateRoleEntity(ctx core.Context, req *entity.RoleEntity) error {
	if req.App == "" {
		return fmt.Errorf("app is required")
	}
	// 查询当前id的详细数据
	var re entity.RoleEntity
	app := req.App
	if err := ctx.DB().Where("app = ?", app).First(&re, req.Id).Error; err != nil {
		return err
	}
	newReq := *req
	_ = mergo.Merge(&newReq, re)
	newReq.App = app
	req.App = app
	if err := r.validate(ctx, app, newReq.RoleId, newReq.EntityId); err != nil {
		return err
	}

	op := r.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("app = ?", app).Updates(req).Error; err != nil {
			return err
		}
		if newReq.RoleId != re.RoleId || newReq.EntityId != re.EntityId || newReq.Action != re.Action {
			op = op.Delete(r.getCacheKey(&re))
		}
		return op.Put(r.getCacheKey(&newReq), &newReq).Do()
	})
}

// DeleteRoleEntity 删除数据
func (r *RoleEntity) DeleteRoleEntity(ctx core.Context, app string, id uint32) error {
	// 查询当前id的详细数据
	var re entity.RoleEntity
	if err := ctx.DB().Where("app = ?", app).First(&re, id).Error; err != nil {
		return err
	}

	op := r.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("app = ?", app).Where("id = ?", id).Delete(&entity.RoleEntity{}).Error; err != nil {
			return err
		}
		return op.Delete(r.getCacheKey(&re)).Do()
	})
}

func (r *RoleEntity) load(ctx core.Context) (map[string]*entity.RoleEntity, error) {
	var list []*entity.RoleEntity
	if err := ctx.DB().Find(&list).Error; err != nil {
		return nil, err
	}
	bucket := map[string]*entity.RoleEntity{}
	for _, item := range list {
		bucket[r.getCacheKey(item)] = item
	}
	return bucket, nil
}

func (r *RoleEntity) getCacheKey(re *entity.RoleEntity) string {
	return fmt.Sprintf("%s:%d:%d:%s", re.App, re.RoleId, re.EntityId, re.Action)
}

func (r *RoleEntity) GetRoleEntityByRDA(ctx core.Context, re *entity.RoleEntity) (*entity.RoleEntity, bool) {
	return r.cache.Get(r.getCacheKey(re))
}

// validate 校验角色与数据实体的归属
// role_entity.app 语义为「实体所属模块」，与角色所属应用可以不同（跨模块授权）：
// 角色始终属于当前登录应用 ctx.AppKey()，实体属于被订阅的基础设施模块 module。
func (r *RoleEntity) validate(ctx core.Context, module string, roleId, entityId uint32) error {
	cur, err := NewApp().GetAppByKeyword(ctx.AppKey())
	if err != nil {
		return err
	}
	if !cur.HasModule(module) {
		return fmt.Errorf("module %s is not subscribed by app %s", module, cur.Keyword)
	}

	// 角色属于当前登录应用
	var role entity.Role
	if err := ctx.DB().Where("app = ? AND id = ?", cur.Keyword, roleId).First(&role).Error; err != nil {
		return err
	}

	// 实体属于目标模块
	var ent entity.Entity
	return ctx.DB().Where("app = ? AND id = ?", module, entityId).First(&ent).Error
}
