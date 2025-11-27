package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Tenant interface {
	// GetTenant 获取指定的租户
	GetTenant(ctx core.Context, id uint32) (*entity.Tenant, error)

	// GetTenantByKeyword 获取指定的租户
	GetTenantByKeyword(ctx core.Context, keyword string) (*entity.Tenant, error)

	// ListTenant 获取租户列表
	ListTenant(ctx core.Context, req *types.ListTenantRequest) ([]*entity.Tenant, uint32, error)

	// CreateTenant 创建租户
	CreateTenant(ctx core.Context, req *entity.Tenant) (uint32, error)

	// UpdateTenant 更新租户
	UpdateTenant(ctx core.Context, req *entity.Tenant) error

	// DeleteTenant 删除租户
	DeleteTenant(ctx core.Context, id uint32) error
}
