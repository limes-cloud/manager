package entity

import "github.com/limes-cloud/kratosx/model"

type AppField struct {
	App      string `json:"app" gorm:"column:app"`           // 应用标识
	FieldId  uint32 `json:"fieldId" gorm:"column:field_id"`  // 关联字段 ID
	Required *bool  `json:"required" gorm:"column:required"` // 是否必填
	Unique   *bool  `json:"unique" gorm:"column:unique"`     // 是否唯一
	Sort     uint32 `json:"sort" gorm:"column:sort"`         // 排序（越小越前）
	Field    *Field `json:"field" gorm:"foreignKey:field_id;references:id"`
	model.BaseModel
}
