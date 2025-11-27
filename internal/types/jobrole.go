package types

import "github.com/limes-cloud/kratosx/model/page"

type ListJobRoleRequest struct {
	page.Search
	InRoleIds []uint32 `json:"inRoleIds"`

	JobId uint32  `json:"jobId"`
	Name  *string `json:"name"`
}

type CreateJobRoleRequest struct {
	RoleId uint32 `json:"roleId"`
	JobId  uint32 `json:"jobId"`
}

type DeleteJobRoleRequest struct {
	JobId  uint32 `json:"jobId"`
	RoleId uint32 `json:"roleId"`
}
