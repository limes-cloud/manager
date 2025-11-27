package types

import "github.com/limes-cloud/kratosx/model/page"

const (
	InfoKey = "x-md-global-info"
)

const (
	OAutherWayBrowserWX = "wx"

	OAutherWayActionScan    = "scan"
	OAutherWayActionJump    = "jump"
	OAutherWayActionCaptcha = "captcha"
)

type CheckAuthRequest struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type AuthorizeBase struct {
	Iat int64 `json:"iat,omitempty"`
	Exp int64 `json:"exp,omitempty"`
}

type AuthorizeInfo struct {
	*AuthorizeBase
	AppId    uint32 `json:"appId"`
	TenantId uint32 `json:"tenantId"`
	UserId   uint32 `json:"userId"`
	DeptId   uint32 `json:"deptId"`
	JobId    uint32 `json:"jobId"`
}

func (az AuthorizeInfo) ToMap() map[string]any {
	return map[string]any{
		"appId":    az.AppId,
		"tenantId": az.TenantId,
		"userId":   az.UserId,
		"deptId":   az.DeptId,
	}
}

type Captcha struct {
	Uuid    string `json:"uuid"`
	Captcha string `json:"captcha"`
	Expire  uint32 `json:"expire"`
}

type LoginRequest struct {
	Tenant    string `json:"tenant"`
	App       string `json:"app"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaId string `json:"captchaId"`
	Captcha   string `json:"captcha"`
}

type LoginReply struct {
	NeedInfo bool   `json:"needInfo"`
	Token    string `json:"token"`
}

type RegisterRequest struct {
	Tenant    string `json:"tenant"`
	App       string `json:"app"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaId string `json:"captchaId"`
	Captcha   string `json:"captcha"`
}

type RegisterReply struct {
	NeedInfo bool   `json:"needInfo"`
	Token    string `json:"token"`
}

type ListAuthorizeRequest struct {
	*page.Search
	UserId *uint32  `json:"userId"`
	AppIds []uint32 `json:"appIds"`
}

type GetAuthorizeRequest struct {
	TenantId uint32 `json:"tenantId"`
	AppId    uint32 `json:"appId"`
	UserId   uint32 `json:"userId"`
}

type OAutherHandleRequest struct {
	Tenant    string `json:"tenant"`
	App       string `json:"app"`
	Keyword   string `json:"keyword"`
	UserAgent string `json:"userAgent"`
	Account   string `json:"account"`
	IP        string `json:"ip"`
}

type OAutherHandleReply struct {
	UUID      string `json:"uuid"`
	Action    string `json:"action"`
	Value     string `json:"value"`
	Tip       string `json:"tip"`
	CodeField string `json:"codeField"`
}

type OAutherLoginRequest struct {
	Account string `json:"account"`
	Code    string `json:"code"`
	Keyword string `json:"keyword"`
	UUID    string `json:"uuid"`
	Tenant  string `json:"tenant"`
	App     string `json:"app"`
}

type OAutherLoginReply struct {
	NeedBind bool    `json:"needBind"`
	NeedInfo bool    `json:"needInfo"`
	Token    *string `json:"token"`
	Expire   *uint32 `json:"expire"`
}

type OAutherBindRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaId string `json:"captchaId"`
	Captcha   string `json:"captcha"`
	UUID      string `json:"uuid"`
}

type OAutherBindReply struct {
	NeedRegister bool
	NeedInfo     bool
	Token        string
}
