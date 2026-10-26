package types

import "github.com/limes-cloud/kratosx/model/page"

type ListFieldRequest struct {
	*page.Search
	Keywords []string `json:"keywords"`
	Keyword  *string  `json:"keyword"`
	Name     *string  `json:"name"`
	Type     *string  `json:"type"`
	Status   *bool    `json:"status"`
}

// FieldOptionItem 字段选项枚举项
type FieldOptionItem struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// FieldOptionsReply 字段枚举配置响应
type FieldOptionsReply struct {
	Types []FieldOptionItem `json:"types"` // 组件类型列表
	Rules []FieldOptionItem `json:"rules"` // 校验规则类型列表
}
