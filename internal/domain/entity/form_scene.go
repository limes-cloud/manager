package entity

import "github.com/limes-cloud/kratosx/model"

// FormScene 表单场景枚举
type FormScene struct {
	Keyword     string  `json:"keyword" gorm:"column:keyword"`         // 场景唯一标识（如 register / profile / activity）
	Name        string  `json:"name" gorm:"column:name"`               // 场景名称（如 "用户注册"、"个人资料"）
	Description *string `json:"description" gorm:"column:description"` // 场景描述
	model.BaseModel
}
