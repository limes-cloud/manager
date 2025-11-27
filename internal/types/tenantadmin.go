package types

import (
	"github.com/limes-cloud/kratosx/model/page"
)

type ListTenantAdminRequest struct {
	*page.Search
	TenantId uint32
}

type CreateTenantAdminRequest struct {
	TenantId uint32 `json:"tenantId"` // 租户ID
	UserId   uint32 `json:"userId"`   // 用户ID
}

type DeleteTenantAdminRequest struct {
	TenantId uint32 `json:"tenantId"` // 租户ID
	UserId   uint32 `json:"userId"`   // 用户ID
}
