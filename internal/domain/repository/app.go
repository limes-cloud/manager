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
	GetAppByKeyword(ctx core.Context, keyword string) (*entity.App, error)

	// ListApp 获取应用信息列表
	ListApp(ctx core.Context, req *types.ListAppRequest) ([]*entity.App, uint32, error)

	// CreateApp 创建应用信息
	CreateApp(ctx core.Context, req *entity.App) (uint32, error)

	// UpdateApp 更新应用信息
	UpdateApp(ctx core.Context, req *entity.App) error

	// DeleteApp 删除应用信息
	DeleteApp(ctx core.Context, id uint32) error

	// ListAppOAuthChannel 获取应用授权列表
	ListAppOAuthChannel(ctx core.Context, req *types.ListAppOAuthChannelRequest) ([]*entity.AppOAuthChannel, uint32, error)

	// ListTenantAppOAuthChannel 获取应用授权列表
	ListTenantAppOAuthChannel(ctx core.Context, req *types.ListTenantAppOAuthChannelRequest) ([]*entity.AppOAuthChannel, error)

	// CreateAppOAuthChannel 创建应用信息
	CreateAppOAuthChannel(ctx core.Context, req *entity.AppOAuthChannel) (uint32, error)

	// DeleteAppOAuthChannel 删除应用信息
	DeleteAppOAuthChannel(ctx core.Context, id uint32) error

	// ListAppField 获取应用字段列表
	ListAppField(ctx core.Context, req *types.ListAppFieldRequest) ([]*entity.AppField, uint32, error)

	// CreateAppField 创建应用字段信息
	CreateAppField(ctx core.Context, req *entity.AppField) (uint32, error)

	// DeleteAppField 删除应用字段信息
	DeleteAppField(ctx core.Context, id uint32) error
}
