package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type TenantAdmin interface {
	// GetAdminIds 获取租户管理员ID列表
	GetAdminIds(tid uint32) []uint32

	// IsAdmin 判断用户是否为租户管理员
	IsAdmin(tid uint32, uid uint32) bool

	// ListTenantAdmin 获取租户管理员列表
	ListTenantAdmin(ctx core.Context, req *types.ListTenantAdminRequest) ([]*entity.TenantAdmin, uint32, error)

	// CreateTenantAdmin 创建租户管理员
	CreateTenantAdmin(ctx core.Context, req *entity.TenantAdmin) (uint32, error)

	// DeleteTenantAdmin 删除租户管理员
	DeleteTenantAdmin(ctx core.Context, req *types.DeleteTenantAdminRequest) error
}
