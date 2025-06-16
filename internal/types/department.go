package types

// ListDepartmentRequest fixed code
type ListDepartmentRequest struct {
	Name    *string  `json:"name"`
	Keyword *string  `json:"keyword"`
	Ids     []uint32 `json:"ids"`
}

type ListDepartmentClassifyRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Name     *string `json:"name"`
}
