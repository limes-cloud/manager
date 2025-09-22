package types

import "github.com/limes-cloud/kratosx/model/page"

type ListFieldRequest struct {
	page.Search
	Keyword *string `json:"keyword"`
	Name    *string `json:"name"`
	Status  *bool   `json:"status"`
}

type FieldType struct {
	Type string
	Name string
}
