package department

// ListDepartmentRequest fixed code
type ListDepartmentRequest struct {
	Order   *string  `json:"order"`
	OrderBy *string  `json:"orderBy"`
	Name    *string  `json:"name"`
	Keyword *string  `json:"keyword"`
	Ids     []uint32 `json:"ids"`
}

type GetDepartmentRequest struct {
	Id      *uint32 `json:"id"`
	Keyword *string `json:"keyword"`
}

type DataScope struct {
	All    bool
	Scopes []uint32
}
