package types

import (
	"encoding/json"

	"github.com/limes-cloud/kratosx/model/page"
)

type OAuthor interface {
	Platform() string
	Name() string
}

const (
	GetOAuthWayBrowserWX = "wx"

	GetOAuthWayActionScan    = "scan"
	GetOAuthWayActionJump    = "jump"
	GetOAuthWayActionCaptcha = "captcha"
)

type ListOAuthChannelRequest struct {
	page.Search
	Keyword *string `json:"keyword"`
	Name    *string `json:"name"`
	Status  *bool   `json:"status"`
}

type ListOAuthRequest struct {
	page.Search
	UserId uint32 `json:"userId"`
}

type OAuthLoginRequest struct {
	User    string `json:"user"`
	Code    string `json:"code"`
	Keyword string `json:"keyword"`
	UUID    string `json:"uuid"`
}

type OAuthLoginReply struct {
	IsBind bool    `json:"isBind"`
	Token  *string `json:"token"`
	Expire *uint32 `json:"expire"`
}

type OAuthHandlerRequest struct {
	Keyword   string `json:"keyword"`
	UserAgent string `json:"user_agent"`
	User      string `json:"user"`
	IP        string `json:"ip"`
}

type OAuthHandlerReply struct {
	UUID      string `json:"uuid"`
	Action    string `json:"action"`
	Value     string `json:"value"`
	Tip       string `json:"tip"`
	CodeField string `json:"codeField"`
}

type ReportOAuthCodeRequest struct {
	Code    string `json:"code"`
	Keyword string `json:"keyword"`
	UUID    string `json:"uuid"`
}

type GetOAuthTokenRequest struct {
	IP   string `json:"ip"`
	UUID string `json:"uuid"`
	User string `json:"user"`
	Code string `json:"code"`
}

type GetOAuthTokenReply struct {
	Token  string `json:"token"`
	OID    string `json:"oid"`
	Expire int64  `json:"expire"`
}

type GetOAuthInfoRequest struct {
	OID   string `json:"oid"`
	Token string `json:"token"`
}

type GetOAuthInfoReply struct {
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

func (reply GetOAuthInfoReply) ToString() string {
	b, _ := json.Marshal(reply)
	return string(b)
}

type OAuthBindRequest struct {
	*UserLoginRequest
	UUID    string
	Keyword string
}
