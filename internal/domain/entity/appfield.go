package entity

import "github.com/limes-cloud/kratosx/model"

type AppField struct {
	model.CreateTenantModel
	AppId    uint32 `json:"appId" gorm:"column:app_id"`
	FieldId  uint32 `json:"fieldId" gorm:"column:field_id"`
	Required *bool  `json:"required" gorm:"column:required" `
	Field    *Field `json:"field" gorm:"foreignKey:field_id;references:id"`
}
