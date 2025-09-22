package types

import "github.com/limes-cloud/kratosx/model/page"

type ListDeptRequest struct {
	InIds      []uint32 `json:"inIds"`
	Name       *string  `json:"name"`
	Keyword    *string  `json:"keyword"`
	ClassifyId *uint32  `json:"classifyId"`
	Status     *bool    `json:"status"`
}

type ListDeptClassifyRequest struct {
	page.Search
	Name   *string `json:"name"`
	Status *bool   `json:"status"`
}
