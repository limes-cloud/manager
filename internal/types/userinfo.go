package types

type ListUserinfoRequest struct {
	Fields []string `json:"fields"`
	AppId  uint32   `json:"appId"`
	UserId uint32   `json:"userId"`
}
