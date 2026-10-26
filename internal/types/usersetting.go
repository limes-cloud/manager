package types

type GetUserSettingRequest struct {
	App    string `json:"app"`
	UserId uint32 `json:"userId"`
}

type UpdateCurrentUserSettingRequest struct {
	App     string `json:"app"`
	Setting string `json:"setting"`
}
