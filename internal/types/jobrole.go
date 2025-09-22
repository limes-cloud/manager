package types

import "github.com/limes-cloud/kratosx/model/page"

type ListRoleJobRequest struct {
	page.Search
	InJobIds []uint32 `json:"inJobIds"`
	RoleId   uint32   `json:"roleId"`
	Name     *string  `json:"name"`
}

type ListJobRoleRequest struct {
	page.Search
	InRoleIds []uint32 `json:"inRoleIds"`
	JobId     uint32   `json:"jobId"`
	Name      *string  `json:"name"`
}

type CreateJobRolesRequest struct {
	RoleIds []uint32 `json:"roleIds"`
	JobId   uint32   `json:"jobId"`
}

type CreateRoleJobsRequest struct {
	JobIds []uint32 `json:"jobIds"`
	RoleId uint32   `json:"roleId"`
}

type DeleteRoleJobsRequest struct {
	JobIds []uint32 `json:"jobIds"`
	RoleId uint32   `json:"roleId"`
}

type DeleteJobRolesRequest struct {
	JobId   uint32   `json:"jobId"`
	RoleIds []uint32 `json:"roleIds"`
}
