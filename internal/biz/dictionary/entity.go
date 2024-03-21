package dictionary

import "github.com/limes-cloud/kratosx/types"

type Dictionary struct {
	types.BaseModel
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DictionaryValue struct {
	types.BaseModel
	DictionaryId uint32  `json:"dictionary_id"`
	Label        string  `json:"label"`
	Value        string  `json:"value"`
	Status       *bool   `json:"status"`
	Weight       *uint32 `json:"weight"`
	Type         string  `json:"type"`
	Extra        string  `json:"extra"`
	Description  string  `json:"description"`
}
