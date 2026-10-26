package entity

import "github.com/limes-cloud/kratosx/model"

// UserForm 用户表单定义（替代原 FormSchema，以场景为中心）
type UserForm struct {
	Scene       string  `json:"scene" gorm:"column:scene"`             // 场景标识（关联 form_scene.keyword），可重复
	Name        string  `json:"name" gorm:"column:name"`               // 表单名称
	Description *string `json:"description" gorm:"column:description"` // 用途描述
	Schema      string  `json:"schema" gorm:"column:schema;type:json"` // 设计器导出的完整 JSON（nodes 树 + config）
	Version     uint32  `json:"version" gorm:"column:version"`         // 版本号，每次保存自增
	Status      *bool   `json:"status" gorm:"column:status"`           // false=草稿 true=已发布
	model.BaseModel
}
