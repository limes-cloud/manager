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

	// ListMenu 获取菜单信息列表
	ListMenu(ctx core.Context, req *types.ListMenuRequest) ([]*entity.Menu, error)

	// GetMenuChildrenIds 获取指定目录下的id
	GetMenuChildrenIds(ctx core.Context, id uint32) ([]uint32, error)

	// CreateMenu 创建菜单信息
	CreateMenu(ctx core.Context, req *entity.Menu) (uint32, error)

	// UpdateMenu 更新菜单信息
	UpdateMenu(ctx core.Context, req *entity.Menu) error

	// DeleteMenu 删除菜单信息
	DeleteMenu(ctx core.Context, id uint32) error

	// SetHome 设置菜单首页
	SetHome(ctx core.Context, id uint32) error
}
