package entity

import (
	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx/types"
)

type OAuth struct {
	UserId    uint32   `json:"userId" gorm:"column:user_id"`
	ChannelId uint32   `json:"channelId" gorm:"column:channel_id"`
	OID       string   `json:"oid" gorm:"column:oid"`
	Token     string   `json:"token" gorm:"column:token"`
	LoggedAt  int64    `json:"loggedAt" gorm:"column:logged_at"`
	ExpiredAt int64    `json:"expiredAt" gorm:"column:expired_at"`
	Extra     string   `json:"extra" gorm:"column:extra"`
	Channel   *Channel `json:"channel"`
	types.CreateModel
}

func (OAuth) TableName() string {
	return "oauth"
}

func (o OAuth) GetExtra() *OAuthExtra {
	var extra OAuthExtra
	_ = json.Unmarshal([]byte(o.Extra), &extra)
	return &extra
}

type OAuthExtra struct {
	UnionID string `json:"union_id"`
}

func (extra OAuthExtra) ToString() string {
	b, _ := json.Marshal(extra)
	return string(b)
}
