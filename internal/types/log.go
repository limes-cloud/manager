package types

import "github.com/limes-cloud/kratosx/model/page"

type ListLoginLogRequest struct {
	page.Search
	Username   *string `json:"username"`
	CreatedAts []int64 `json:"createdAts"`
}

type ListAuthLogRequest struct {
	page.Search
	Username   string  `json:"username"`
	UserId     *uint32 `json:"userId"`
	CreatedAts []int64 `json:"createdAts"`
}
