package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Menu interface {
	GetMenuIdsByAppId(ctx core.Context, id uint32) ([]uint32, error)

	// GetMenu 获取菜单列表
	GetMenu(ctx core.Context, id uint32) (*entity.Menu, error)

	// GetMenuByKeyword 获取指定应用下指定标识的菜单
	GetMenuByKeyword(ctx core.Context, appId uint32, keyword string) (*entity.Menu, error)

	// GetMenuByApi 获取指定应用下指定接口的菜单
	GetMenuByApi(ctx core.Context, appId uint32, api string, method string) (*entity.Menu, error)

	// ListMenu 获取菜单信息列表
	ListMenu(ctx core.Context, req *types.ListMenuRequest) ([]*entity.Menu, error)

	// GetMenuChildrenIds 获取指定目录下的id
	GetMenuChildrenIds(ctx core.Context, id uint32) ([]uint32, error)

	// CreateMenu 创建菜单信息
	CreateMenu(ctx core.Context, req *entity.Menu) (uint32, error)

	// UpdateMenu 更新菜单信息
	UpdateMenu(ctx core.Context, req *entity.Menu) error

	// UpdateRegisteredMenu 更新注册菜单信息
	UpdateRegisteredMenu(ctx core.Context, req *entity.Menu) error

	// DeleteMenu 删除菜单信息
	DeleteMenu(ctx core.Context, id uint32) error

	// DeleteMenus 删除多个菜单信息
	DeleteMenus(ctx core.Context, ids []uint32) error

	// SetHome 设置菜单首页
	SetHome(ctx core.Context, id uint32) error

	GetCacheMenu(api string, method string) (*types.CacheMenu, bool)
}
