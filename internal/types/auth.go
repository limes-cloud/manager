package types

import "github.com/limes-cloud/kratosx/model/page"

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
	page.Search
	Username   *string `json:"username"`
	CreatedAts []int64 `json:"createdAts"`
}

type ListAuthLogRequest struct {
	page.Search
	Username   *string `json:"username"`
	CreatedAts []int64 `json:"createdAts"`
}
