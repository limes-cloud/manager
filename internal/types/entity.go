package types

import "github.com/limes-cloud/kratosx/model/page"

type ListEntityRequest struct {
	page.Search
	AppId    uint32  `json:"appId"`
	Name     *string `json:"name"`
	Database *string `json:"database"`
}

type ListEntityFieldRequest struct {
	EntityId uint32  `json:"entityId"`
	Name     *string `json:"name"`
}

type ListEntityRuleRequest struct {
	*page.Search
	EntityId *uint32 `json:"entityId"`
	Name     *string `json:"name"`
	Status   *bool   `json:"status"`
}
