package entity

import (
	"github.com/limes-cloud/kratosx/model"
)

const (
	AppEntityName = "app"
)

type App struct {
	Logo        string  `json:"logo" gorm:"column:logo"`
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Name        string  `json:"name" gorm:"column:name"`
	Status      *bool   `json:"status" gorm:"column:status"`
	Reason      *string `json:"reason" gorm:"column:reason"`
	Private     *bool   `json:"private" gorm:"column:private"`
	Extra       *string `json:"extra" gorm:"column:extra"`
	Description *string `json:"description" gorm:"column:description"`
	model.BaseModel
}

// TableName 获取表名称
func (a App) TableName() string {
	return AppEntityName
}

type AppOAuthChannel struct {
	model.CreateTenantModel
	AppId     uint32        `json:"appId" gorm:"column:app_id"`
	ChannelId uint32        `json:"channelId" gorm:"column:channel_id"`
	Channel   *OAuthChannel `json:"channel" gorm:"foreignKey:channel_id;references:id"`
}

func (AppOAuthChannel) TableName() string {
	return "app_oauth_channel"
}

type AppField struct {
	model.CreateTenantModel
	AppId   uint32 `json:"appId" gorm:"column:app_id"`
	FieldId uint32 `json:"fieldId" gorm:"column:field_id"`
	Field   *Field `json:"field" gorm:"foreignKey:field_id;references:id"`
}
