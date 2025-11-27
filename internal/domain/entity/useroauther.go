package entity

import (
	"encoding/json"

	"github.com/limes-cloud/kratosx/model"
)

type UserOAuther struct {
	UserId    uint32   `json:"userId" gorm:"column:user_id"`
	OAutherId uint32   `json:"oautherId" gorm:"column:oauther_id"`
	OID       string   `json:"oid" gorm:"column:oid"`
	Token     string   `json:"token" gorm:"column:token"`
	LoggedAt  int64    `json:"loggedAt" gorm:"column:logged_at"`
	ExpiredAt int64    `json:"expiredAt" gorm:"column:expired_at"`
	Extra     string   `json:"extra" gorm:"column:extra"`
	OAuther   *OAuther `json:"oauther"`
	model.BaseTenantModel
}

func (c UserOAuther) TableName() string {
	return "user_oauther"
}

func (o UserOAuther) GetExtra() *UserOAutherExtra {
	var extra UserOAutherExtra
	_ = json.Unmarshal([]byte(o.Extra), &extra)
	return &extra
}

type UserOAutherExtra struct {
	UnionID string `json:"unionId"`
}
