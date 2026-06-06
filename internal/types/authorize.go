package types

import (
	"github.com/limes-cloud/kratosx/model/page"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	InfoKey = "x-md-global-info"
)

const (
	OAutherWayBrowserWX = "wx"

	OAutherWayActionUniLogin = "unilogin"
	OAutherWayActionScan     = "scan"
	OAutherWayActionJump     = "jump"
	OAutherWayActionCaptcha  = "captcha"

	PlatformApp        = "app"              // App
	PlatformWeb        = "web"              // Web
	PlatformWebYiBan   = "web-yiban"        // 易班 app-web
	PlatformWebWeiXin  = "web-weixin"       // 微信 app-web
	PlatformWebAliPay  = "web-alipay"       // 支付宝 app-web
	PlatformWebDouYin  = "web-douyin"       // 抖音 app-web
	PlatformMPWeiXin   = "mp-weixin"        // 微信小程序
	PlatformMPAlipay   = "mp-alipay"        // 支付宝小程序
	PlatformMPBaidu    = "mp-baidu"         // 百度小程序
	PlatformMPTouTiao  = "mp-toutiao"       // 抖音小程序
	PlatformMPLark     = "mp-lark"          // 飞书小程序
	PlatformMPQQ       = "mp-qq"            // QQ小程序
	PlatformMPKuaiShou = "mp-kuaishou"      // 快手小程序
	PlatformMPJD       = "mp-jd"            // 京东小程序
	PlatformMP360      = "mp-360"           // 360小程序
	PlatformMPHarmony  = "mp-harmony"       // 鸿蒙元服务
	PlatformMPQuickApp = "quickapp-webview" // 快应用通用(包含联盟、华为)
)

type ListVisibleOAutherRequest struct {
	App      string `json:"app"`
	Platform string `json:"platform"`
}

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
	UserId   uint32 `json:"userId"`
	DeptId   uint32 `json:"deptId"`
	JobId    uint32 `json:"jobId"`
}

func (az AuthorizeInfo) ToMap() map[string]any {
	return map[string]any{
		"appId":  az.AppId,
		"userId": az.UserId,
		"deptId": az.DeptId,
		"jobId":  az.JobId,
	}
}

type Captcha struct {
	Uuid    string `json:"uuid"`
	Captcha string `json:"captcha"`
	Expire  uint32 `json:"expire"`
}

type LoginRequest struct {
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

type OAutherVisibleRequest struct {
	UserAgent string `json:"userAgent"`
	Platform  string `json:"platform"`
}

type OAutherVisibleReply struct {
	Visible       bool   `json:"visible"`
	Recommend     bool   `json:"recommend"`
	RecommendText string `json:"recommendText"`
}

type GetAuthorizeRequest struct {
	AppId    uint32 `json:"appId"`
	UserId   uint32 `json:"userId"`
}

type OAutherHandleRequest struct {
	App       string `json:"app"`
	Keyword   string `json:"keyword"`
	UserAgent string `json:"userAgent"`
	Account   string `json:"account"`
	IP        string `json:"ip"`
	Platform  string `json:"platform"`
}

type OAutherHandleReply struct {
	UUID   string `json:"uuid"`
	Action string `json:"action"`
	Value  string `json:"value"`
	Tip    string `json:"tip"`
}

type OAutherLoginRequest struct {
	Account string `json:"account"`
	Code    string `json:"code"`
	UUID    string `json:"uuid"`
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
	Register  bool   `json:"register"`
}

type OAutherBindReply struct {
	NeedRegister bool
	NeedInfo     bool
	Token        string
}

type OAutherReportRequest struct {
	Code string `json:"code"`
	UUID string `json:"uuid"`
}

type FillInfo struct {
	Type    string          `json:"type"`
	Keyword string          `json:"keyword"`
	Name    string          `json:"name"`
	Value   *structpb.Value `json:"value"`
}
