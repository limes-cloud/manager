package entity

import "github.com/limes-cloud/kratosx/model"

// AppForm 应用-表单关联（中间表）
type AppForm struct {
	AppKeyword        string `json:"appKeyword" gorm:"column:app_keyword"` // 应用标识（关联 app.keyword）
	FormId            uint32 `json:"formId" gorm:"column:form_id"`         // 表单 ID（关联 user_form.id）
	Scene             string `json:"scene" gorm:"column:scene"`            // 场景标识（冗余自 user_form.scene，用于唯一约束）
	model.CreateModel        // Id(uint32 自增主键) + CreatedAt(int64 Unix 秒)
}
