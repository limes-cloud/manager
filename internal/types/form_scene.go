package types

import "github.com/limes-cloud/kratosx/model/page"

// ListFormSceneRequest 查询场景列表请求
type ListFormSceneRequest struct {
	*page.Search
	Keyword *string `json:"keyword"`
	Name    *string `json:"name"`
}
