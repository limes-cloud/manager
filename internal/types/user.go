package types

import "github.com/limes-cloud/kratosx/model/page"

type GetUserRequest struct {
	Id       *uint32 `json:"id"`
	Username *string `json:"username"`
}

type ListUserRequest struct {
	page.Search
	InIds     []uint32 `json:"id"`
	InDeptIds []uint32 `json:"inDeptIds"`
	Username  *string  `json:"name"`
	Status    *bool    `json:"status"`
}
