package types

import "github.com/limes-cloud/kratosx/model/page"

type ListAppRequest struct {
	*page.Search
	InIds   []uint32 `json:"inIds"`
	Keyword *string  `json:"keyword"`
	Name    *string  `json:"name"`
	Status  *bool    `json:"status"`
}
