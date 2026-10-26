package entity

import "github.com/limes-cloud/kratosx/model"

type Field struct {
	Keyword     string       `json:"keyword" gorm:"column:keyword"`               // 字段唯一标识
	Type        string       `json:"type" gorm:"column:type"`                     // 组件类型：input/textarea/number/password/radio/checkbox/select/time/date/datetime/upload
	Name        string       `json:"name" gorm:"column:name"`                     // 字段标签
	Config      *FieldConfig `json:"config" gorm:"column:config;serializer:json"` // 组件配置
	Status      *bool        `json:"status" gorm:"column:status"`                 // 是否启用：0 禁用 1 启用
	Description *string      `json:"description" gorm:"column:description"`       // 字段说明
	model.BaseModel
}
