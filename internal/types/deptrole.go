package types

import "github.com/limes-cloud/kratosx/model/page"

type ListDeptRoleRequest struct {
	page.Search
	InRoleIds []uint32 `json:"inRoleIds"`
	DeptId    uint32   `json:"deptId"`
	Name      *string  `json:"name"`
}

type CreateDeptRoleRequest struct {
	RoleId uint32 `json:"roleId"`
	DeptId uint32 `json:"deptId"`
}

type DeleteDeptRoleRequest struct {
	DeptId uint32 `json:"deptId"`
	RoleId uint32 `json:"roleId"`
}
