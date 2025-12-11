package types

import "github.com/limes-cloud/kratosx/model/page"

type ListFeedbackClassifyRequest struct {
	*page.Search
	Name *string `json:"name"`
}

type ListFeedbackRequest struct {
	*page.Search
	AppId      *uint32  `json:"appId"`
	ClassifyId *uint32  `json:"classifyId"`
	Status     *string  `json:"status"`
	Platform   *string  `json:"platform"`
	AppIds     []uint32 `json:"appIds"`
}
