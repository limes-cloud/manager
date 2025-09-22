package types

import "github.com/limes-cloud/kratosx/model/page"

type ListRoleDeptRequest struct {
	page.Search
	InDeptIds []uint32 `json:"inDeptIds"`
	RoleId    uint32   `json:"roleId"`
	Name      *string  `json:"name"`
}

type ListDeptRoleRequest struct {
	page.Search
	InRoleIds []uint32 `json:"inRoleIds"`
	DeptId    uint32   `json:"deptId"`
	Name      *string  `json:"name"`
}

type CreateDeptRolesRequest struct {
	RoleIds []uint32 `json:"roleIds"`
	DeptId  uint32   `json:"deptId"`
}

type CreateRoleDeptsRequest struct {
	DeptIds []uint32 `json:"deptIds"`
	RoleId  uint32   `json:"roleId"`
}

type DeleteRoleDeptsRequest struct {
	DeptIds []uint32 `json:"deptIds"`
	RoleId  uint32   `json:"roleId"`
}

type DeleteDeptRolesRequest struct {
	DeptId  uint32   `json:"deptId"`
	RoleIds []uint32 `json:"roleIds"`
}
