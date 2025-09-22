package types

type GetRoleRequest struct {
	Id      *uint32 `json:"id"`
	Keyword *string `json:"keyword"`
}

type ListRoleRequest struct {
	InIds   []uint32 `json:"inIds"`
	Name    *string  `json:"name"`
	Status  *bool    `json:"status"`
	Keyword *string  `json:"keyword"`
}
