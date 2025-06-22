package types

import "encoding/json"

type OAuthorName interface {
	Platform() string
	Name() string
}

const (
	GetOAuthWayBrowserWX = "wx"

	GetOAuthWayTypeScan    = "scan"
	GetOAuthWayTypeJump    = "jump"
	GetOAuthWayTypeCaptcha = "captcha"
)

type ListOAuthRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	UserId   uint32  `json:"userId"`
}

type OAuthLoginRequest struct {
	User    string
	Code    string
	Keyword string
	UUID    string
}

type OAuthLoginReply struct {
	IsBind bool    `json:"isBind"`
	Token  *string `json:"token"`
	Expire *uint32 `json:"expire"`
}

type OAuthWayRequest struct {
	Keyword string
	User    string
}

type GetOAuthWayRequest struct {
	UserAgent string
	User      string
	IP        string
}

type ReportOAuthCodeRequest struct {
	Code    string
	Keyword string
	UUID    string
}

type GetOAuthWayReply struct {
	UUID  string
	Type  string
	Value string
	Tip   string
}

type GetOAuthTokenRequest struct {
	IP   string
	UUID string
	User string
	Code string
}

type GetOAuthTokenReply struct {
	Token  string
	OID    string
	Expire int64
}

type GetOAuthInfoRequest struct {
	OID   string
	Token string
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

type AuthRequest struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type GetUserLoginCaptchaReply struct {
	Uuid    string `json:"uuid"`
	Captcha string `json:"captcha"`
	Expire  uint32 `json:"expire"`
}

type UserLoginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaId string `json:"captchaId"`
	Captcha   string `json:"captcha"`
}

type ListLoginLogRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"pageSize"`
	Username   *string `json:"username"`
	CreatedAts []int64 `json:"createdAts"`
}

type OAuthBindRequest struct {
	*UserLoginRequest
	UUID    string
	Keyword string
}
