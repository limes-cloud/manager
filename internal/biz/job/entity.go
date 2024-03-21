package job

import "github.com/limes-cloud/kratosx/types"

type Job struct {
	types.BaseModel
	Keyword     string  `json:"keyword"`
	Name        string  `json:"name"`
	Weight      *uint32 `json:"weight"`
	Description string  `json:"description"`
}
