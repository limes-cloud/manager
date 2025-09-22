package types

import "github.com/limes-cloud/kratosx/model/page"

type ListFeedbackCategoryRequest struct {
	page.Search
	Name *string `json:"name"`
}

type ListFeedbackRequest struct {
	page.Search
	AppId      *uint32  `json:"appId"`
	CategoryId *uint32  `json:"categoryId"`
	Status     *string  `json:"status"`
	Platform   *string  `json:"platform"`
	AppIds     []uint32 `json:"appIds"`
}
