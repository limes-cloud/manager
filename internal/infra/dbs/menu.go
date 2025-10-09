package dbs

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/samber/lo"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/kratosx/pkg/cache"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type cacheMenuType struct {
	Id   uint32 `json:"id"`
	Type string `json:"type"`
}

type Menu struct {
	cache *cache.Cache[string, *cacheMenuType]
}

var (
	menuIns      *Menu
	menuOnce     sync.Once
	menuCacheKey = "menu"
)

func NewMenu() *Menu {
	menuOnce.Do(func() {
		menuIns = &Menu{}

		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "menu"),
			kratosx.WithSkipDBHook(),
		)
		ctx.BeforeStart("load cache menu", func() {
			menuIns.cache = cache.NewCacheAndInit[string, *cacheMenuType](
				ctx,
				ctx.Redis(),
				menuCacheKey,
				cache.HookLoad(func() (map[string]*cacheMenuType, error) {
					return menuIns.load(ctx)
				}),
			)
		})
	})
	return menuIns
}

// ListMenu 获取列表
func (infra *Menu) ListMenu(ctx core.Context, req *types.ListMenuRequest) ([]*entity.Menu, error) {
	var (
		ms []*entity.Menu
		fs = []string{"*"}
	)
	db := ctx.DB().Model(entity.Menu{}).Select(fs)
	if req.AppId != nil {
		db = db.Where("app_id = ?", req.AppId)
	}
	if req.InIds != nil {
		db = db.Where("id in ?", req.InIds)
	}
	if req.InTypes != nil {
		db = db.Where("type in ?", req.InTypes)
	}
	if req.NotInTypes != nil {
		db = db.Where("type not in ?", req.NotInTypes)
	}

	db = db.Order("weight desc,id asc")
	return ms, db.Find(&ms).Error
}

// CreateMenu 创建数据
func (infra *Menu) CreateMenu(ctx core.Context, menu *entity.Menu) (uint32, error) {
	if menu.Type == entity.MenuTypeRoot {
		menu.ParentId = 0
	}

	return menu.Id, ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(menu).Error; err != nil {
			return err
		}

		if err := infra.appendMenuChildren(ctx, menu.ParentId, menu.Id); err != nil {
			return err
		}

		if lo.Contains([]string{entity.MenuTypeApi, entity.MenuTypeBasic}, menu.Type) {
			op := infra.cache.OP(ctx)
			return op.Put(infra.getCacheKey(menu.GetApi(), menu.GetMethod()), &cacheMenuType{
				Id:   menu.Id,
				Type: menu.Type,
			}).Do()
		}

		return nil
	})
}

// GetMenu 获取指定的数据
func (infra *Menu) GetMenu(ctx core.Context, id uint32) (*entity.Menu, error) {
	menu := entity.Menu{}
	return &menu, ctx.DB().First(&menu, id).Error
}

// UpdateMenu 更新数据
func (infra *Menu) UpdateMenu(ctx core.Context, menu *entity.Menu) error {
	if menu.Id == menu.ParentId {
		return errors.New("父级不能为自己")
	}

	// 获取之前的菜单信息
	old, err := infra.GetMenu(ctx, menu.Id)
	if err != nil {
		return err
	}

	// 出现了父级菜单变化
	if old.ParentId != menu.ParentId {
		if err := infra.removeMenuParent(ctx, menu.Id); err != nil {
			return err
		}
		if err := infra.appendMenuChildren(ctx, menu.ParentId, menu.Id); err != nil {
			return err
		}
	}

	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Updates(menu).Error; err != nil {
			return err
		}
		if old.GetApi() != menu.GetApi() || old.GetMethod() != menu.GetMethod() {
			op := infra.cache.OP(ctx)
			op = op.Delete(infra.getCacheKey(old.GetApi(), old.GetMethod()))
			op = op.Put(infra.getCacheKey(menu.GetApi(), menu.GetMethod()), &cacheMenuType{
				Id:   menu.Id,
				Type: menu.Type,
			})
			return op.Do()
		}
		return nil
	})
}

// DeleteMenu 删除数据
func (infra *Menu) DeleteMenu(ctx core.Context, id uint32) error {
	ids, err := infra.GetMenuChildrenIds(ctx, id)
	if err != nil {
		return err
	}
	ids = append(ids, id)

	// 查询当前条件下的菜单接口
	var apis []*entity.Menu
	if err := ctx.DB().
		Where("id in ?", ids).
		Where("type in ?", []string{
			entity.MenuTypeApi,
			entity.MenuTypeBasic,
		}).
		Find(&apis).Error; err != nil {
		return err
	}
	op := infra.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id in ?", ids).Delete(&entity.Menu{}).Error; err != nil {
			return err
		}
		for _, item := range apis {
			op = op.Delete(infra.getCacheKey(item.GetApi(), item.GetMethod()))
		}
		return op.Do()
	})
}

// GetMenuChildrenIds 获取指定id的所有子id
func (infra *Menu) GetMenuChildrenIds(ctx core.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.MenuClosure{}).
		Select("children").
		Where("parent=?", id).
		Scan(&ids).Error
}

func (infra *Menu) ListMenuChildrenApi(ctx core.Context, id uint32) ([]*entity.Menu, error) {
	ids, err := infra.GetMenuChildrenIds(ctx, id)
	if err != nil {
		return nil, err
	}
	ids = append(ids, id)
	var list []*entity.Menu
	return list, ctx.DB().Select([]string{"id", "type", "api", "method"}).Find(&list, "id in ?", ids).Error
}

// GetMenuIdsByAppId 获取指定app id的所有菜单id
func (infra *Menu) GetMenuIdsByAppId(ctx core.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.Menu{}).
		Select("id").
		Where("app_id=?", id).
		Scan(&ids).Error
}

// GetMenuParentIds 获取指定id的所有父id
func (infra *Menu) GetMenuParentIds(ctx core.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.MenuClosure{}).
		Select("parent").
		Where("children=?", id).
		Scan(&ids).Error
}

// appendMenuChildren 添加id到指定的父id下
func (infra *Menu) appendMenuChildren(ctx core.Context, pid uint32, id uint32) error {
	list := []*entity.MenuClosure{
		{
			Parent:   pid,
			Children: id,
		},
	}
	ids, _ := infra.GetMenuParentIds(ctx, pid)
	for _, item := range ids {
		list = append(list, &entity.MenuClosure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Create(&list).Error
}

// removeMenuParent 删除指定id的所有父层级
func (infra *Menu) removeMenuParent(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.MenuClosure{}, "children=?", id).Error
}

// SetHome 重置当前树的首页节点
func (infra *Menu) SetHome(ctx core.Context, id uint32) error {
	pids, err := infra.GetMenuParentIds(ctx, id)
	if err != nil {
		return err
	}

	var menu entity.Menu
	if err = ctx.DB().
		Where("id in ?", pids).
		Where("type=?", entity.MenuTypeRoot).
		Find(&menu).Error; err != nil {
		return err
	}

	cids, err := infra.GetMenuChildrenIds(ctx, menu.Id)
	if err != nil {
		return err
	}
	pids = append(pids, cids...)
	return ctx.DB().Model(entity.Menu{}).Where("id in ?", pids).Update("is_home", false).Error
}

// load 加载全量的数据
func (infra *Menu) load(ctx core.Context) (map[string]*cacheMenuType, error) {
	var list []*entity.Menu
	if err := ctx.DB().Where("type in ?", []string{
		entity.MenuTypeApi,
		entity.MenuTypeBasic,
	}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]*cacheMenuType{}
	for _, item := range list {
		bucket[infra.getCacheKey(item.GetApi(), item.GetMethod())] = &cacheMenuType{
			Id:   item.Id,
			Type: item.Type,
		}
	}

	return bucket, nil
}

// getCacheKey 获取缓存的key
func (infra *Menu) getCacheKey(api string, method string) string {
	return fmt.Sprintf("%s:%s", api, method)
}

// GetMenuTypeInfo 获取api类型
func (infra *Menu) GetMenuTypeInfo(api string, method string) (uint32, string, bool) {
	val, has := infra.cache.Get(infra.getCacheKey(api, method))
	if !has {
		return 0, "", false
	}
	return val.Id, val.Type, true
}
