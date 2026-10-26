package types

import "github.com/limes-cloud/kratosx/model/page"

type ListAppOAutherRequest struct {
	*page.Search
	App     string  `json:"app"`
	Keyword *string `json:"keyword"`
	Name    *string `json:"name"`
}
