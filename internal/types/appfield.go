package types

import "github.com/limes-cloud/kratosx/model/page"

type ListAppFieldRequest struct {
	*page.Search
	App      string  `json:"app"`
	Keyword  *string `json:"keyword"`
	Name     *string `json:"name"`
	Required *bool   `json:"required"`
}
