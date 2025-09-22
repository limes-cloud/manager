package types

import "github.com/limes-cloud/kratosx/model/page"

type ListJobRequest struct {
	page.Search
	Keyword *string `json:"keyword"`
	Name    *string `json:"name"`
	Status  *bool   `json:"status"`
}
