package entity

import (
	"encoding/json"

	"github.com/limes-cloud/kratosx/model"
)

type OAuthChannel struct {
	Logo        string `json:"logo" gorm:"column:logo"`               // 渠道logo
	Keyword     string `json:"keyword" gorm:"column:keyword"`         // 渠道标识
	Type        string `json:"type"  gorm:"column:type"`              // 渠道类型
	Name        string `json:"name" gorm:"column:name"`               // 渠道名称
	Status      *bool  `json:"status" gorm:"column:status"`           // 启用状态
	Ak          string `json:"ak" gorm:"column:ak"`                   // 渠道AK
	Sk          string `json:"sk" gorm:"column:sk"`                   // 渠道SK
	Extra       string `json:"extra" gorm:"column:extra"`             // 扩展信息
	Description string `json:"description" gorm:"column:description"` // 描述信息
	model.BaseTenantModel
}

func (c OAuthChannel) TableName() string {
	return "oauth_channel"
}

type ChannelExtra struct {
	CallBack string `json:"callback"`
	Captcha  string `json:"captcha"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Subject  string `json:"subject"`
}

func (c OAuthChannel) GetExtra() *ChannelExtra {
	var extra ChannelExtra
	_ = json.Unmarshal([]byte(c.Extra), &extra)
	return &extra
}

func (extra ChannelExtra) ToString() string {
	b, _ := json.Marshal(extra)
	return string(b)
}

type OAuth struct {
	UserId    uint32        `json:"userId" gorm:"column:user_id"`
	ChannelId uint32        `json:"channelId" gorm:"column:channel_id"`
	OID       string        `json:"oid" gorm:"column:oid"`
	Token     string        `json:"token" gorm:"column:token"`
	LoggedAt  int64         `json:"loggedAt" gorm:"column:logged_at"`
	ExpiredAt int64         `json:"expiredAt" gorm:"column:expired_at"`
	Extra     string        `json:"extra" gorm:"column:extra"`
	Channel   *OAuthChannel `json:"channel"`
	model.BaseTenantModel
}

func (o OAuth) GetExtra() *OAuthExtra {
	var extra OAuthExtra
	_ = json.Unmarshal([]byte(o.Extra), &extra)
	return &extra
}

type OAuthExtra struct {
	UnionID string `json:"unionId"`
}

func (extra OAuthExtra) ToString() string {
	b, _ := json.Marshal(extra)
	return string(b)
}
