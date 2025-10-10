package types

import "github.com/limes-cloud/kratosx/model/page"

type GetAppRequest struct {
	Id      *uint32 `json:"id"`
	Keyword *string `json:"keyword"`
}

type ListAppRequest struct {
	*page.Search
	InIds   []uint32 `json:"inIds"`
	Keyword *string  `json:"keyword"`
	Name    *string  `json:"name"`
	Status  *bool    `json:"status"`
}

type ListAppOAuthChannelRequest struct {
	page.Search
	AppId   uint32  `json:"appId"`
	Keyword *string `json:"keyword"`
	Name    *string `json:"name"`
}

type ListTenantAppOAuthChannelRequest struct {
	TenantId uint32 `json:"tenantId"`
	AppId    uint32 `json:"appId"`
}

type ListAppFieldRequest struct {
	page.Search
	AppId   uint32  `json:"appId"`
	Keyword *string `json:"keyword"`
	Name    *string `json:"name"`
}
