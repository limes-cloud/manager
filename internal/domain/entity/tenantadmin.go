package entity

import (
	"github.com/limes-cloud/kratosx/model"
)

type TenantAdmin struct {
	model.CreateModel
	TenantId uint32 `gorm:"column:tenant_id" json:"tenantId"` // 租户ID
	UserId   uint32 `gorm:"column:user_id" json:"userId"`     // 应用ID
	User     *User  `json:"user"`
}
