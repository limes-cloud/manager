package role

// ListRoleRequest fixed code
type ListRoleRequest struct {
	Order   *string  `json:"order"`
	OrderBy *string  `json:"orderBy"`
	Name    *string  `json:"name"`
	Keyword *string  `json:"keyword"`
	Ids     []uint32 `json:"ids"`
}

type GetRoleRequest struct {
	Id      *uint32 `json:"id"`
	Keyword *string `json:"keyword"`
}
