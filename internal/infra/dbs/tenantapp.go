package dbs

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/limes-cloud/kratosx/pkg/cache"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type TenantApp struct {
	cache     *cache.Cache[string, struct{}]
	cacheMenu *cache.Cache[string, struct{}]
}

var (
	taIns          *TenantApp
	taOnce         sync.Once
	taCacheKey     = "tenantapp"
	taMenuCacheKey = "tenantappmenu"
)

func NewTenantApp() *TenantApp {
	taOnce.Do(func() {
		taIns = &TenantApp{}
		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "tenant app"),
			kratosx.WithSkipDBHook(),
		)

		ctx.BeforeStart("load cache tenant app", func() {
			taIns.cache = cache.NewCacheAndInit[string, struct{}](
				ctx,
				ctx.Redis(),
				taCacheKey,
				cache.HookLoad(func() (map[string]struct{}, error) {
					return taIns.load(ctx)
				}),
			)

			taIns.cacheMenu = cache.NewCacheAndInit[string, struct{}](
				ctx,
				ctx.Redis(),
				taMenuCacheKey,
				cache.HookLoad(func() (map[string]struct{}, error) {
					return taIns.loadMenu(ctx)
				}),
			)
		})
	})
	return taIns
}

// ListTenantApp 获取app列表
func (ta *TenantApp) ListTenantApp(ctx core.Context, req *types.ListTenantAppRequest) ([]*entity.TenantApp, uint32, error) {
	var (
		list  []*entity.TenantApp
		total int64
	)

	db := ctx.DB().Model(entity.TenantApp{}).Where("tenant_id=?", req.TenantId)

	if req.AppName != nil {
		joinWhere := ctx.DB().Where("App.name LIKE ?", *req.AppName+"%")
		db = db.Joins("App", joinWhere)
	} else {
		db = db.Joins("App")
	}

	// 查询条件下数据总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 搜索排序
	db = page.SearchScopes(db, &req.Search)

	return list, uint32(total), db.Find(&list).Error
}

// CreateTenantApp 创建租户的应用信息
func (ta *TenantApp) CreateTenantApp(ctx core.Context, ent *entity.TenantApp) (uint32, error) {
	op := ta.cache.OP(ctx)
	err := ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(ent).Error; err != nil {
			return err
		}
		return op.Put(ta.getCacheKey(ent.TenantId, ent.AppId), struct{}{}).Do()
	})
	if err != nil {
		return 0, err
	}
	return ent.Id, nil
}

// UpdateTenantApp 更新租户的应用信息
func (ta *TenantApp) UpdateTenantApp(ctx core.Context, ent *entity.TenantApp) error {
	return ctx.DB().
		Where("tenant_id=? AND app_id=?", ent.TenantId, ent.AppId).
		Updates(ent).Error
}

// DeleteTenantApp 删除租户的应用
func (ta *TenantApp) DeleteTenantApp(ctx core.Context, tenantId uint32, appId uint32) error {
	op := ta.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().
			Where("tenant_id=? AND app_id=?", tenantId, appId).
			Delete(&entity.TenantApp{}).Error; err != nil {
			return err
		}

		if err := ctx.DB().
			Where("tenant_id=? AND app_id=?", tenantId, appId).
			Delete(&entity.TenantAppMenu{}).Error; err != nil {
			return err
		}

		return op.Delete(ta.getCacheKey(tenantId, appId)).Do()
	})
}

func (ta *TenantApp) GetTenantAppMenuIds(ctx core.Context, tid uint32, aid uint32) ([]uint32, error) {
	var menuIds []uint32
	return menuIds, ctx.DB().Model(entity.TenantAppMenu{}).
		Where("tenant_id = ? and app_id = ?", tid, aid).
		Pluck("menu_id", &menuIds).Error
}

func (ta *TenantApp) GetTenantMenuIds(tid uint32) []uint32 {
	var (
		menuIds = make([]uint32, 0)
		keys    = ta.cacheMenu.Keys()
	)
	for _, key := range keys {
		ot, _, om := ta.splitMenuCacheKey(key)
		if ot == tid {
			menuIds = append(menuIds, om)
		}
	}

	return menuIds
}

func (ta *TenantApp) CreateTenantAppMenuIds(ctx core.Context, tid, aid uint32, mids []uint32) error {
	var (
		list []entity.TenantAppMenu
		op   = ta.cacheMenu.OP(ctx)
	)
	for _, mid := range mids {
		list = append(list, entity.TenantAppMenu{
			TenantId: tid,
			AppId:    aid,
			MenuId:   mid,
		})
		op = op.Put(ta.getMenuCacheKey(tid, aid, mid), struct{}{})
	}
	if err := ctx.DB().Create(&list).Error; err != nil {
		return err
	}
	return op.Do()
}

func (ta *TenantApp) DeleteTenantAppMenuIds(ctx core.Context, tid, aid uint32, mids []uint32) error {
	if err := ctx.DB().
		Where("tenant_id=? AND app_id=? AND menu_id in ?", tid, aid, mids).
		Delete(&entity.TenantAppMenu{}).Error; err != nil {
		return err
	}
	op := ta.cacheMenu.OP(ctx)
	for _, mid := range mids {
		op = op.Delete(ta.getMenuCacheKey(tid, aid, mid))
	}
	return op.Do()
}

// load 加载全量的数据
func (ta *TenantApp) load(ctx core.Context) (map[string]struct{}, error) {
	var list []*entity.TenantApp
	if err := ctx.DB().Model(&entity.TenantApp{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]struct{}{}
	for _, item := range list {
		bucket[ta.getCacheKey(item.TenantId, item.AppId)] = struct{}{}
	}

	return bucket, nil
}

// loadMenu 加载菜单的数据
func (ta *TenantApp) loadMenu(ctx core.Context) (map[string]struct{}, error) {
	var list []*entity.TenantAppMenu
	if err := ctx.DB().Model(&entity.TenantAppMenu{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]struct{}{}
	for _, item := range list {
		bucket[ta.getMenuCacheKey(item.TenantId, item.AppId, item.MenuId)] = struct{}{}
	}

	return bucket, nil
}

// getCacheKey 获取缓存的key
func (ta *TenantApp) getMenuCacheKey(tid uint32, aid uint32, mid uint32) string {
	return fmt.Sprintf("%d:%d:%d", tid, aid, mid)
}

// getCacheKey 获取缓存的key
func (ta *TenantApp) getCacheKey(tid uint32, aid uint32) string {
	return fmt.Sprintf("%d:%d", tid, aid)
}

// HasAppScope 判断是否具有指定的应用权限
func (ta *TenantApp) HasAppScope(tid uint32, aid uint32) bool {
	key := ta.getCacheKey(tid, aid)
	_, ok := ta.cache.Get(key)
	return ok
}

// splitCacheKey 通过key获取id
func (ta *TenantApp) splitMenuCacheKey(key string) (uint32, uint32, uint32) {
	arr := strings.Split(key, ":")
	if len(arr) != 3 {
		return 0, 0, 0
	}
	tid, _ := strconv.Atoi(arr[0])
	aid, _ := strconv.Atoi(arr[1])
	mid, _ := strconv.Atoi(arr[2])
	return uint32(tid), uint32(aid), uint32(mid)
}

// splitCacheKey 通过key获取id
func (ta *TenantApp) splitCacheKey(key string) (uint32, uint32) {
	arr := strings.Split(key, ":")
	if len(arr) != 2 {
		return 0, 0
	}
	rid, _ := strconv.Atoi(arr[0])
	mid, _ := strconv.Atoi(arr[1])
	return uint32(rid), uint32(mid)
}

func (ta *TenantApp) GetAppIds(tid uint32) []uint32 {
	var (
		ids  = make([]uint32, 0)
		keys = ta.cache.Keys()
	)

	for _, key := range keys {
		trid, aid := ta.splitCacheKey(key)
		if trid == tid {
			ids = append(ids, aid)
		}
	}
	return ids
}
