package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type App interface {
	// GetApp 获取指定的应用信息
	GetApp(ctx core.Context, id uint32) (*entity.App, error)

	// GetAppByKeyword 获取指定的应用信息
	GetAppByKeyword(keyword string) (*entity.App, error)

	// ListApp 获取应用信息列表
	ListApp(ctx core.Context, req *types.ListAppRequest) ([]*entity.App, uint32, error)

	// CreateApp 创建应用信息
	CreateApp(ctx core.Context, req *entity.App) (uint32, error)

	// UpdateApp 更新应用信息
	UpdateApp(ctx core.Context, req *entity.App) error

	// DeleteApp 删除应用信息
	DeleteApp(ctx core.Context, id uint32) error
}
