package types

import (
	"encoding/json"

	"github.com/limes-cloud/kratosx/model/page"
)

type ListOAutherRequest struct {
	*page.Search
	Keyword *string `json:"keyword"`
	Name    *string `json:"name"`
	Status  *bool   `json:"status"`
}

type OAutherType struct {
	Keyword string
	Name    string
}

type OAutherTokenRequest struct {
	IP      string `json:"ip"`
	UUID    string `json:"uuid"`
	Account string `json:"account"`
	Code    string `json:"code"`
}

type OAutherTokenReply struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
	OID    string `json:"oid"`
}

type OAutherInfoRequest struct {
	OID   string `json:"oid"`
	Token string `json:"token"`
}

type OAutherInfoReply struct {
	OID      string `json:"oid,omitempty"`      // 授权唯一ID
	Birthday string `json:"birthday,omitempty"` // 生日
	RealName string `json:"realName,omitempty"` // 姓名
	NickName string `json:"nickName,omitempty"` // 昵称
	Gender   string `json:"gender,omitempty"`   // 性别
	Province string `json:"province,omitempty"` // 省
	City     string `json:"city,omitempty"`     // 市
	Country  string `json:"country,omitempty"`  // 区
	Avatar   string `json:"avatar,omitempty"`   // 头像
	UnionId  string `json:"unionId,omitempty"`  // 联合ID
}

func (reply OAutherInfoReply) ToString() string {
	b, _ := json.Marshal(reply)
	return string(b)
}
