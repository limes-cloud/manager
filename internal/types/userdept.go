package types

import "github.com/limes-cloud/kratosx/model/page"

type ListDeptUserRequest struct {
	page.Search
	DeptId uint32  `json:"deptId"`
	Name   *string `json:"name"`
}

type ListUserDeptRequest struct {
	page.Search
	InDeptIds []uint32 `json:"inDeptIds"`
	UserId    uint32   `json:"userId"`
	Name      *string  `json:"name"`
}

type DeleteDeptUsersRequest struct {
	UserId uint32 `json:"userId"`
	DeptId uint32 `json:"deptId"`
}

type DeleteUserDeptsRequest struct {
	UserId uint32 `json:"userId"`
	DeptId uint32 `json:"deptId"`
}
