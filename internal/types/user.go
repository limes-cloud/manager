package types

import "github.com/limes-cloud/kratosx/model/page"

type OfflineUserRequest struct {
	UserId uint32   `json:"userId"`
	AppIds []uint32 `json:"appIds"`
}

type GetCurrentUserRequest struct {
	App *string `json:"app"`
}

type UpdateCurrentUserRequest struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type UpdateCurrentUserPasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
}

type GetUserRequest struct {
	Id       *uint32 `json:"id"`
	Username *string `json:"username"`
	App      *string `json:"app"`
}

type ListUserRequest struct {
	page.Search
	InIds      []uint32 `json:"idIds"`
	NotInIds   []uint32 `json:"notInIds"`
	InDeptIds  []uint32 `json:"inDeptIds"`
	InJobIds   []uint32 `json:"inJobIds"`
	Username   *string  `json:"username"`
	Status     *bool    `json:"status"`
	LoggedAts  []uint32 `json:"loggedAts"`
	CreatedAts []uint32 `json:"createdAts"`
	AppId      *uint32  `json:"appId"`

	// 中转参数
	DeptId *uint32 `json:"deptId"`
	JobId  *uint32 `json:"jobId"`
	App    *string `json:"app"`
}
