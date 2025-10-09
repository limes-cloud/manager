package types

type GetUserSettingRequest struct {
	AppId  uint32 `json:"appId"`
	UserId uint32 `json:"userId"`
}

type UpdateCurrentUserSettingRequest struct {
	App     string `json:"app"`
	Setting string `json:"setting"`
}
