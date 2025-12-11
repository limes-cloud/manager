package types

import "github.com/limes-cloud/kratosx/model/page"

type ListFieldRequest struct {
	*page.Search
	TenantId uint32   `json:"tenantId"`
	Keywords []string `json:"keywords"`
	Keyword  *string  `json:"keyword"`
	Name     *string  `json:"name"`
	Status   *bool    `json:"status"`
	Required *bool    `json:"required"`
	Unique   *bool    `json:"unique"`
}

type FieldType struct {
	Type string
	Name string
}
