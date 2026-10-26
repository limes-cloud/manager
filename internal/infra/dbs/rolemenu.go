package dbs

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/kratosx/pkg/cache"

	"github.com/limes-cloud/kratosx"

	"gorm.io/gorm/clause"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
)

type RoleMenu struct {
	cache *cache.Cache[string, struct{}]
}

var (
	rmIns      *RoleMenu
	rmOnce     sync.Once
	rmCacheKey = "rolemenu"
)

func NewRoleMenu() *RoleMenu {
	rmOnce.Do(func() {
		// 初始化属性
		rmIns = &RoleMenu{}

		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("watch role_menu policy", ""),
		)
		ctx.BeforeStart("load cache role menu", func() {
			rmIns.cache = cache.NewCacheAndInit[string, struct{}](
				ctx,
				ctx.Redis(),
				rmCacheKey,
				cache.HookLoad(func() (map[string]struct{}, error) {
					return rmIns.load(ctx)
				}),
			)
		})
	})
	return rmIns
}

// CreateRoleMenus 获取指定菜单的角色列表
func (rm *RoleMenu) CreateRoleMenus(ctx core.Context, rms []*entity.RoleMenu) error {
	if len(rms) == 0 {
		return nil
	}
	app := rms[0].App
	if app == "" {
		return fmt.Errorf("app is required")
	}
	for _, item := range rms {
		item.App = app
		if err := rm.validate(ctx, app, item.RoleId, item.MenuId); err != nil {
			return err
		}
	}
	op := rm.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Model(&entity.RoleMenu{}).
			Clauses(clause.OnConflict{UpdateAll: true}).
			Create(rms).Error; err != nil {
			return err
		}

		return op.Puts(rm.transvals(rms)).Do()
	})
}

// DeleteRoleMenus 获取指定菜单的角色列表
func (rm *RoleMenu) DeleteRoleMenus(ctx core.Context, rms []*entity.RoleMenu) error {
	if len(rms) == 0 {
		return nil
	}
	app := rms[0].App
	if app == "" {
		return fmt.Errorf("app is required")
	}
	for _, item := range rms {
		item.App = app
	}
	op := rm.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		db := ctx.DB()
		for _, item := range rms {
			db = db.Or(ctx.DB().Where("app = ? AND role_id = ? AND menu_id = ?", app, item.RoleId, item.MenuId))
		}
		if err := db.Delete(&entity.RoleMenu{}).Error; err != nil {
			return err
		}

		// 从缓存中删除
		return op.Deletes(rm.transkeys(rms)).Do()
	})
}

func (rm *RoleMenu) GetMenuIdsByRoleIds(ctx core.Context, app string, rids []uint32) []uint32 {
	var (
		ids  = make([]uint32, 0)
		keys = rm.cache.Keys()
	)

	include := value.NewInclude(rids)
	for _, key := range keys {
		itemApp, rid, mid := rm.splitCacheKey(key)
		if itemApp == app && mid != 0 && include.Has(rid) {
			ids = append(ids, mid)
		}
	}
	return ids
}

func (rm *RoleMenu) GetRoleIdsByMenuIds(ctx core.Context, app string, mids []uint32) []uint32 {
	var (
		ids  = make([]uint32, 0)
		keys = rm.cache.Keys()
	)

	include := value.NewInclude(mids)
	for _, key := range keys {
		itemApp, rid, mid := rm.splitCacheKey(key)
		if itemApp == app && rid != 0 && include.Has(mid) {
			ids = append(ids, rid)
		}
	}
	return ids
}

func (rm *RoleMenu) hasRoleMenu(app string, rid uint32, mid uint32) bool {
	_, ok := rm.cache.Get(rm.getCacheKey(app, rid, mid))
	return ok
}

func (rm *RoleMenu) DeleteRoles(ctx core.Context, app string, rids []uint32) error {
	var (
		keys    = rm.cache.Keys()
		include = value.NewInclude(rids)
		op      = rm.cache.OP(ctx)
	)

	for _, key := range keys {
		itemApp, rid, _ := rm.splitCacheKey(key)
		if itemApp == app && include.Has(rid) {
			op.Delete(key)
		}
	}

	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("app = ?", app).Where("role_id in ?", rids).Delete(&entity.RoleMenu{}).Error; err != nil {
			return err
		}
		return op.Do()
	})
}

func (rm *RoleMenu) DeleteMenus(ctx core.Context, app string, mids []uint32) error {
	var (
		keys    = rm.cache.Keys()
		include = value.NewInclude(mids)
		op      = rm.cache.OP(ctx)
	)

	for _, key := range keys {
		itemApp, _, mid := rm.splitCacheKey(key)
		if itemApp == app && include.Has(mid) {
			op.Delete(key)
		}
	}

	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("app = ?", app).Where("menu_id in ?", mids).Delete(&entity.RoleMenu{}).Error; err != nil {
			return err
		}
		return op.Do()
	})
}

// load 加载全量的数据
func (rm *RoleMenu) load(ctx core.Context) (map[string]struct{}, error) {
	var list []*entity.RoleMenu
	if err := ctx.DB().Model(&entity.RoleMenu{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]struct{}{}
	for _, item := range list {
		bucket[rm.getCacheKey(item.App, item.RoleId, item.MenuId)] = struct{}{}
	}

	return bucket, nil
}

func (rm *RoleMenu) transvals(drs []*entity.RoleMenu) map[string]struct{} {
	kvs := map[string]struct{}{}
	for _, item := range drs {
		kvs[rm.getCacheKey(item.App, item.RoleId, item.MenuId)] = struct{}{}
	}
	return kvs
}

func (rm *RoleMenu) transkeys(drs []*entity.RoleMenu) []string {
	var keys []string
	for _, item := range drs {
		keys = append(keys, rm.getCacheKey(item.App, item.RoleId, item.MenuId))
	}
	return keys
}

// getCacheKey 获取缓存的key
func (rm *RoleMenu) getCacheKey(app string, rid uint32, mid uint32) string {
	return fmt.Sprintf("%s:%d:%d", app, rid, mid)
}

// splitCacheKey 通过key获取id
func (rm *RoleMenu) splitCacheKey(key string) (string, uint32, uint32) {
	arr := strings.Split(key, ":")
	if len(arr) != 3 {
		return "", 0, 0
	}
	rid, _ := strconv.Atoi(arr[1])
	mid, _ := strconv.Atoi(arr[2])
	return arr[0], uint32(rid), uint32(mid)
}

// validate 校验角色与菜单的归属
// role_menu.app 语义为「菜单所属模块」，与角色所属应用可以不同（跨模块授权）：
// 角色始终属于当前登录应用 ctx.AppKey()，菜单属于被订阅的基础设施模块 module。
// 同时校验 module 在当前应用的可触达范围内，防止授予未订阅模块的菜单。
func (rm *RoleMenu) validate(ctx core.Context, module string, roleId, menuId uint32) error {
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

	// 菜单属于目标模块
	var menu entity.Menu
	return ctx.DB().Where("app = ? AND id = ?", module, menuId).First(&menu).Error
}
