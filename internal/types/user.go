package types

import "github.com/limes-cloud/kratosx/model/page"

type GetCurrentUserRequest struct {
	App *string `json:"app"`
}

type UpdateCurrentUserRequest struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type GetUserRequest struct {
	Id       *uint32 `json:"id"`
	Username *string `json:"username"`
	App      *string `json:"app"`
}

type ListUserRequest struct {
	page.Search
	InIds     []uint32 `json:"idIds"`
	NotInIds  []uint32 `json:"notInIds"`
	InDeptIds []uint32 `json:"inDeptIds"`
	Username  *string  `json:"name"`
	Nickname  *string  `json:"nickname"`
	Status    *bool    `json:"status"`
}
