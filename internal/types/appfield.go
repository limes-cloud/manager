package types

import "github.com/limes-cloud/kratosx/model/page"

type ListAppFieldRequest struct {
	*page.Search
	AppId    uint32  `json:"appId"`
	Keyword  *string `json:"keyword"`
	Name     *string `json:"name"`
	Required *bool   `json:"required"`
}
