package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type TenantApp interface {
	// GetAppIds 获取租户应用ID列表
	GetAppIds(tid uint32) []uint32

	// GetTenantApp 获取租户应用
	GetTenantApp(ctx core.Context, req *types.GetTenantAppRequest) (*entity.TenantApp, error)

	// ListTenantApp 获取租户应用列表
	ListTenantApp(ctx core.Context, req *types.ListTenantAppRequest) ([]*entity.TenantApp, uint32, error)

	// CreateTenantApp 创建租户应用
	CreateTenantApp(ctx core.Context, req *entity.TenantApp) (uint32, error)

	// UpdateTenantApp 更新租户应用
	UpdateTenantApp(ctx core.Context, req *entity.TenantApp) error

	// DeleteTenantApp 删除租户应用
	DeleteTenantApp(ctx core.Context, tid uint32, aid uint32) error

	// GetTenantMenuIds 获取租户的菜单ids
	GetTenantMenuIds(tid uint32) []uint32

	// GetTenantAppMenuIds 获取租户应用的菜单ids
	GetTenantAppMenuIds(ctx core.Context, tid uint32, aid uint32) ([]uint32, error)

	// CreateTenantAppMenuIds 添加租户应用的菜单ids
	CreateTenantAppMenuIds(ctx core.Context, tid, aid uint32, mids []uint32) error

	// DeleteTenantAppMenuIds 获取租户应用的菜单ids
	DeleteTenantAppMenuIds(ctx core.Context, tid, aid uint32, mids []uint32) error
}
