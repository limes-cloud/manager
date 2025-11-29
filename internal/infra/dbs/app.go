package dbs

import (
	"context"
	"sync"

	"gorm.io/gorm"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/cache"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type App struct {
	cache *cache.Cache[string, *entity.App]
}

var (
	appIns      *App
	appOnce     sync.Once
	appCacheKey = "app"
)

func NewApp() *App {
	appOnce.Do(func() {
		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "app"),
			kratosx.WithSkipDBHook(),
		)
		appIns = &App{}

		ctx.BeforeStart("load cache app", func() {
			appIns.cache = cache.NewCacheAndInit[string, *entity.App](
				ctx,
				ctx.Redis(),
				appCacheKey,
				cache.HookLoad(func() (map[string]*entity.App, error) {
					return appIns.load(ctx)
				}),
			)
		})
	})
	return appIns
}

// GetApp 获取指定的数据
func (r App) GetApp(ctx core.Context, id uint32) (*entity.App, error) {
	app := entity.App{}
	return &app, ctx.DB().Where("id = ?", id).First(&app).Error
}

// GetAppByKeyword 获取指定数据
func (r App) GetAppByKeyword(keyword string) (*entity.App, error) {
	app, is := r.cache.Get(keyword)
	if !is {
		return nil, gorm.ErrRecordNotFound
	}
	return app, nil
}

// ListApp 获取列表
func (r App) ListApp(ctx core.Context, req *types.ListAppRequest) ([]*entity.App, uint32, error) {
	var (
		list  []*entity.App
		total int64
	)

	db := ctx.DB().Model(entity.App{})
	if req.InIds != nil {
		db = db.Where("id IN ?", req.InIds)
	}
	if req.Keyword != nil {
		db = db.Where("keyword LIKE ?", *req.Keyword+"%")
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if req.Search != nil {
		// 查询条件下数据总数
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}

		// 分页排序
		db = page.SearchScopes(db, req.Search)
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateApp 创建数据
func (r App) CreateApp(ctx core.Context, app *entity.App) (uint32, error) {
	return app.Id, ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(app).Error; err != nil {
			return err
		}
		return r.cache.OP(ctx).Put(app.Keyword, app).Do()
	})
}

// UpdateApp 更新数据
func (r App) UpdateApp(ctx core.Context, app *entity.App) error {
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id = ?", app.Id).Updates(app).Error; err != nil {
			return err
		}

		// 查询应用最新数据
		var na entity.App
		if err := ctx.DB().Where("id = ?", app.Id).First(&na).Error; err != nil {
			return err
		}
		return r.cache.OP(ctx).Put(na.Keyword, &na).Do()
	})
}

// DeleteApp 删除数据
func (r App) DeleteApp(ctx core.Context, id uint32) error {
	// 查询应用最新数据
	var na entity.App
	if err := ctx.DB().Where("id = ?", id).First(&na).Error; err != nil {
		return err
	}

	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id = ?", id).Delete(&entity.App{}).Error; err != nil {
			return err
		}
		return r.cache.OP(ctx).Delete(na.Keyword).Do()
	})
}

// load 加载全量的数据
func (r App) load(ctx core.Context) (map[string]*entity.App, error) {
	var list []*entity.App
	if err := ctx.DB().Model(&entity.App{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]*entity.App{}
	for _, item := range list {
		bucket[item.Keyword] = item
	}
	return bucket, nil
}
