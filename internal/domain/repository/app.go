package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type App interface {
	// GetApp 获取指定的应用信息
	GetApp(ctx kratosx.Context, id uint32) (*entity.App, error)

	// ListApp 获取应用信息列表
	ListApp(ctx kratosx.Context, req *types.ListAppRequest) ([]*entity.App, uint32, error)

	// CreateApp 创建应用信息
	CreateApp(ctx kratosx.Context, req *entity.App) (uint32, error)

	// UpdateApp 更新应用信息
	UpdateApp(ctx kratosx.Context, req *entity.App) error

	// UpdateAppStatus 更新应用信息状态
	UpdateAppStatus(ctx kratosx.Context, req *types.UpdateAppStatusRequest) error

	// DeleteApp 删除应用信息
	DeleteApp(ctx kratosx.Context, id uint32) error

	// GetAppByKeyword 获取指定的应用信息
	GetAppByKeyword(ctx kratosx.Context, keyword string) (*entity.App, error)
}
