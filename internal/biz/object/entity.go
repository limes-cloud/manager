package object

import "github.com/limes-cloud/kratosx/types"

type ObjectRule struct {
	Field   string `json:"field"`
	Operate string `json:"operate"`
	Object  uint32 `json:"object"`
}

type Object struct {
	types.BaseModel
	Keyword     string  `json:"keyword"`
	Name        string  `json:"name"`
	Api         string  `json:"api"`
	Method      string  `json:"method"`
	Params      *string `json:"params"`
	Label       string  `json:"label"`
	Value       string  `json:"value"`
	Description string  `json:"description"`
}
