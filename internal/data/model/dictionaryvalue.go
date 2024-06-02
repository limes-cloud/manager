package model

import (
	"github.com/limes-cloud/kratosx/types"
)

type DictionaryValue struct {
	DictionaryId uint32      `json:"dictionaryId" gorm:"column:dictionary_id"`
	Label        string      `json:"label" gorm:"column:label"`
	Value        string      `json:"value" gorm:"column:value"`
	Status       *bool       `json:"status" gorm:"column:status"`
	Weight       *int32      `json:"weight" gorm:"column:weight"`
	Type         *string     `json:"type" gorm:"column:type"`
	Extra        *string     `json:"extra" gorm:"column:extra"`
	Description  *string     `json:"description" gorm:"column:description"`
	Dictionary   *Dictionary //fixed code
	types.BaseModel
}
