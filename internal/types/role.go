package types

type GetRoleRequest struct {
	App     string `json:"app"`
	Keyword string `json:"keyword"`
}

type ListRoleRequest struct {
	InIds   []uint32 `json:"inIds"`
	Name    *string  `json:"name"`
	Status  *bool    `json:"status"`
	Keyword *string  `json:"keyword"`
}
