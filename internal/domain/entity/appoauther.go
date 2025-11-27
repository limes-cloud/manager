package entity

import "github.com/limes-cloud/kratosx/model"

type AppOAuther struct {
	model.CreateTenantModel
	AppId     uint32   `json:"appId" gorm:"column:app_id"`
	OAutherId uint32   `json:"oautherId" gorm:"column:oauther_id"`
	Type      string   `json:"type"  gorm:"column:type"`
	OAuther   *OAuther `json:"oauther" gorm:"foreignKey:oauther_id;references:id"`
}

func (AppOAuther) TableName() string {
	return "app_oauther"
}
