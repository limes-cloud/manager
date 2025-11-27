package types

import "github.com/limes-cloud/kratosx/model/page"

type GetTenantRequest struct {
	Id      uint32 `json:"id"`
	Keyword string `json:"keyword"`
}

type ListTenantRequest struct {
	*page.Search
	Keyword *string `json:"keyword"`
	Name    *string `json:"name"`
	Status  *bool   `json:"status"`
	AppId   *uint32 `json:"appId"`

	// 中间转换字段
	App *string `json:"app"`
}
