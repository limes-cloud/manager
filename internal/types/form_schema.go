package types

import "github.com/limes-cloud/kratosx/model/page"

// GetUserFormRequest 获取单个表单请求
type GetUserFormRequest struct {
	Id    *uint32 `json:"id"`
	Scene *string `json:"scene"`
}

// ListUserFormRequest 分页查询表单列表请求
type ListUserFormRequest struct {
	*page.Search
	Scene  *string `json:"scene"`
	Name   *string `json:"name"`
	Status *bool   `json:"status"`
}
