package entity

import "github.com/limes-cloud/kratosx/model"

type Tenant struct {
	model.DeleteModel
	Keyword     string `gorm:"column:keyword" json:"keyword"`         // 租户标识
	Logo        string `gorm:"column:logo"  json:"logo"`              // 租户头像
	Name        string `gorm:"column:name" json:"name"`               // 租户名称
	Status      *bool  `gorm:"column:status" json:"status"`           // 租户状态
	Description string `gorm:"column:description" json:"description"` // 租户简介
	Weight      *int32 `gorm:"column:weight" json:"weight"`           // 权重
}
