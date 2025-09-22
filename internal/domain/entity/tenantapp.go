package entity

import "github.com/limes-cloud/kratosx/model"

type TenantApp struct {
	model.BaseModel
	TenantId  uint32   `gorm:"column:tenant_id" json:"tenantId"`   // 租户ID
	AppId     uint32   `gorm:"column:app_id" json:"appId"`         // 应用ID
	ExpiredAt uint32   `gorm:"column:expired_at" json:"expiredAt"` // 过期时间
	App       *App     `json:"app"`
	MenuIds   []uint32 `json:"menuIds" gorm:"-"`
}

type TenantAppMenu struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	TenantId uint32 `json:"tenantId" gorm:"column:tenant_id"` // 租户ID
	AppId    uint32 `json:"appId" gorm:"column:app_id"`       // 应用ID
	MenuId   uint32 `json:"menuId"  gorm:"column:menu_id" `   // 菜单ID
}
