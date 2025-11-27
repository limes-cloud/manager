package types

import (
	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/model/page"
)

type GetTenantAppRequest struct {
	AppId    uint32 `json:"appId"`
	TenantId uint32 `json:"tenantId"` // 租户ID
}

type ListTenantAppRequest struct {
	*page.Search
	AppName    *string `json:"appName"`
	AppKeyword *string `json:"appKeyword"`
	TenantId   uint32  `json:"tenantId"` // 租户ID
}

type CreateTenantAppRequest struct {
	model.BaseModel
	TenantId  uint32   `json:"tenantId"`  // 租户ID
	AppId     uint32   `json:"appId"`     // 应用ID
	ExpiredAt uint32   `json:"expiredAt"` // 过期时间
	MenuIds   []uint32 `json:"menuIds"`   // 菜单id
	Setting   string   `json:"setting"`   // 配置
}

type UpdateTenantAppRequest struct {
	model.BaseModel
	TenantId  uint32   `json:"tenantId"`  // 租户ID
	AppId     uint32   `json:"appId"`     // 应用ID
	ExpiredAt uint32   `json:"expiredAt"` // 过期时间
	MenuIds   []uint32 `json:"menuIds"`   // 菜单id
	Setting   string   `json:"setting"`   // 配置
}
