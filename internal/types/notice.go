package types

import "github.com/limes-cloud/kratosx/model/page"

type ListNoticeRequest struct {
	*page.Search
	ClassifyId *uint32 `json:"classifyId"`
	AppId      *uint32 `json:"appId"`
	Title      *string `json:"title"`
	IsTop      *bool   `json:"isTop"`
	Status     *bool   `json:"status"`
	NotRead    *bool   `json:"notRead"`
	UserId     *uint32 `json:"userId"`
}
