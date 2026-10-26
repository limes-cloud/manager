package types

import "github.com/limes-cloud/kratosx/model/page"

type ListDeptUserRequest struct {
	page.Search
	DeptId uint32  `json:"deptId"`
	Name   *string `json:"name"`
	App    string  `json:"app"`
}

type ListUserDeptRequest struct {
	*page.Search
	InDeptIds []uint32 `json:"inDeptIds"`
	UserId    uint32   `json:"userId"`
	Dept      *string  `json:"dept"`
	Job       *string  `json:"job"`
	App       string   `json:"app"`
}

type DeleteDeptUsersRequest struct {
	UserId uint32 `json:"userId"`
	DeptId uint32 `json:"deptId"`
}

type DeleteUserDeptsRequest struct {
	UserId uint32 `json:"userId"`
	DeptId uint32 `json:"deptId"`
}
