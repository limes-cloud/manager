package entity

import (
	"github.com/limes-cloud/kratosx/model"
)

const (
	FieldEmail = "email"
	FieldPhone = "phone"
)

type Field struct {
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Type        string  `json:"type" gorm:"column:type"`
	Name        string  `json:"name" gorm:"column:name"`
	Status      *bool   `json:"status" gorm:"column:status"`
	Required    *bool   `json:"required" gorm:"column:required"`
	Unique      *bool   `json:"unique" gorm:"column:unique"`
	Description *string `json:"description" gorm:"column:description"`
	model.BaseTenantModel
}
