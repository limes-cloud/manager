package types

import (
	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/model/page"
)

type ListTenantAppRequest struct {
	AppName  *string `json:"appName"`
	TenantId uint32  `json:"tenantId"` // 租户ID
	page.Search
}

type CreateTenantAppRequest struct {
	model.BaseModel
	TenantId  uint32   `json:"tenantId"`  // 租户ID
	AppId     uint32   `json:"appId"`     // 应用ID
	ExpiredAt uint32   `json:"expiredAt"` // 过期时间
	MenuIds   []uint32 `json:"menuIds"`   // 菜单id
}

type UpdateTenantAppRequest struct {
	model.BaseModel
	TenantId  uint32   `json:"tenantId"`  // 租户ID
	AppId     uint32   `json:"appId"`     // 应用ID
	ExpiredAt uint32   `json:"expiredAt"` // 过期时间
	MenuIds   []uint32 `json:"menuIds"`   // 菜单id
}
