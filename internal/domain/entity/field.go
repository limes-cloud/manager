package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type Field struct {
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Type        string  `json:"type" gorm:"column:type"`
	Name        string  `json:"name" gorm:"column:name"`
	Status      *bool   `json:"status" gorm:"column:status"`
	Description *string `json:"description" gorm:"column:description"`
	types.BaseModel
}
