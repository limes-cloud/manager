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
		// 监听变更时间
		ctx := core.MustContext(context.Background(), kratosx.WithTrace("watch role_menu policy", ""))

		// 初始化属性
		rmIns = &RoleMenu{}

		// 缓存相关操作
		{
			lc := cache.NewCache[string, struct{}](
				ctx.Redis(),
				rmCacheKey,
				cache.HookLoad(func() (map[string]struct{}, error) {
					return rmIns.load(ctx)
				}),
			)

			// 加载缓存,加载失败则直接报错，避免线上隐式错误。
			if err := lc.Init(ctx); err != nil {
				panic(err)
			}
			// 监听缓存变更
			go lc.Subscribe(ctx)
			// 定期修复缓存
			go lc.Repair(ctx)

			rmIns.cache = lc
		}
	})
	return rmIns
}

// CreateRoleMenus 获取指定菜单的角色列表
func (rm *RoleMenu) CreateRoleMenus(ctx core.Context, rms []*entity.RoleMenu) error {
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
	op := rm.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		db := ctx.DB()
		for _, item := range rms {
			db = db.Or(ctx.DB().Where("role_id = ? AND menu_id = ?", item.RoleId, item.MenuId))
		}
		if err := db.Delete(&entity.RoleMenu{}).Error; err != nil {
			return err
		}

		// 从缓存中删除
		return op.Deletes(rm.transkeys(rms)).Do()
	})
}

func (rm *RoleMenu) GetMenuIdsByRoleIds(rids []uint32) []uint32 {
	var (
		ids  []uint32
		keys = rm.cache.Keys()
	)

	include := value.NewInclude(rids)
	for _, key := range keys {
		rid, mid := rm.splitCacheKey(key)
		if mid != 0 && include.Has(rid) {
			ids = append(ids, mid)
		}
	}
	return ids
}

func (rm *RoleMenu) GetRoleIdsByMenuIds(mids []uint32) []uint32 {
	var (
		ids  []uint32
		keys = rm.cache.Keys()
	)

	include := value.NewInclude(mids)
	for _, key := range keys {
		rid, mid := rm.splitCacheKey(key)
		if rid != 0 && include.Has(mid) {
			ids = append(ids, rid)
		}
	}
	return ids
}

func (rm *RoleMenu) DeleteRoles(ctx core.Context, rids []uint32) error {
	var (
		keys    = rm.cache.Keys()
		include = value.NewInclude(rids)
		op      = rm.cache.OP(ctx)
	)

	for _, key := range keys {
		rid, _ := rm.splitCacheKey(key)
		if include.Has(rid) {
			op.Delete(key)
		}
	}

	return op.Do()
}

func (rm *RoleMenu) DeleteMenus(ctx core.Context, mids []uint32) error {
	var (
		keys    = rm.cache.Keys()
		include = value.NewInclude(mids)
		op      = rm.cache.OP(ctx)
	)

	for _, key := range keys {
		_, mid := rm.splitCacheKey(key)
		if include.Has(mid) {
			op.Delete(key)
		}
	}

	return op.Do()
}

// load 加载全量的数据
func (rm *RoleMenu) load(ctx core.Context) (map[string]struct{}, error) {
	var list []*entity.RoleMenu
	if err := ctx.DB().Model(&entity.RoleMenu{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]struct{}{}
	for _, item := range list {
		bucket[rm.getCacheKey(item.RoleId, item.MenuId)] = struct{}{}
	}

	return bucket, nil
}

func (rm *RoleMenu) transvals(drs []*entity.RoleMenu) map[string]struct{} {
	kvs := map[string]struct{}{}
	for _, item := range drs {
		kvs[rm.getCacheKey(item.RoleId, item.MenuId)] = struct{}{}
	}
	return kvs
}

func (rm *RoleMenu) transkeys(drs []*entity.RoleMenu) []string {
	var keys []string
	for _, item := range drs {
		keys = append(keys, rm.getCacheKey(item.RoleId, item.MenuId))
	}
	return keys
}

// getCacheKey 获取缓存的key
func (rm *RoleMenu) getCacheKey(rid uint32, mid uint32) string {
	return fmt.Sprintf("%d:%d", rid, mid)
}

// splitCacheKey 通过key获取id
func (rm *RoleMenu) splitCacheKey(key string) (uint32, uint32) {
	arr := strings.Split(key, ":")
	if len(arr) != 2 {
		return 0, 0
	}
	rid, _ := strconv.Atoi(arr[0])
	mid, _ := strconv.Atoi(arr[1])
	return uint32(rid), uint32(mid)
}
