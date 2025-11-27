package types

type ListUserinfoRequest struct {
	Fields   []string `json:"fields"`
	UserId   uint32   `json:"userId"`
	TenantId uint32   `json:"tenantId"`
}

type GetUserinfoByFieldValueRequest struct {
	TenantId uint32 `json:"tenantId"`

	Field string `json:"field"`
	Value string `json:"value"`
}
