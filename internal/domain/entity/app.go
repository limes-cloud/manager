package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type App struct {
	Logo          string        `json:"logo" gorm:"column:logo"`
	LogoUrl       string        `json:"logoUrl" gorm:"-"`
	Keyword       string        `json:"keyword" gorm:"column:keyword"`
	Name          string        `json:"name" gorm:"column:name"`
	Status        *bool         `json:"status" gorm:"column:status"`
	DisableDesc   *string       `json:"disableDesc" gorm:"column:disable_desc"`
	AllowRegistry *bool         `json:"allowRegistry" gorm:"column:allow_registry"`
	Version       string        `json:"version" gorm:"column:version"`
	Copyright     string        `json:"copyright" gorm:"column:copyright"`
	Extra         *string       `json:"extra" gorm:"column:extra"`
	Description   *string       `json:"description" gorm:"column:description"`
	AppChannels   []*AppChannel `json:"appChannels"`
	AppFields     []*AppField   `json:"appFields"`
	Channels      []*Channel    `json:"channels" gorm:"many2many:app_channel"`
	Fields        []*Field      `json:"fields" gorm:"many2many:app_field"`
	types.BaseModel
}

type AppChannel struct {
	AppId     uint32 `json:"appId" gorm:"column:app_id"`
	ChannelId uint32 `json:"channelId" gorm:"column:channel_id"`
}

type AppField struct {
	AppId   uint32 `json:"appId" gorm:"column:app_id"`
	FieldId uint32 `json:"fieldId" gorm:"column:field_id"`
}
