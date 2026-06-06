package types

type ListUserinfoRequest struct {
	Fields []string `json:"fields"`
	UserId uint32   `json:"userId"`
}

type GetUserinfoByFieldValueRequest struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
