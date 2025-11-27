package types

import (
	"github.com/limes-cloud/kratosx/model/page"
)

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

type ReportOAuthCodeRequest struct {
	Code    string `json:"code"`
	Keyword string `json:"keyword"`
	UUID    string `json:"uuid"`
}
