package types

type ListUserinfoRequest struct {
	Fields []string `json:"fields"`
	UserId uint32   `json:"userId"`
}
