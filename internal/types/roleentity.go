package types

import "github.com/limes-cloud/kratosx/model/page"

type ListRoleEntityRequest struct {
	page.Search
	RoleId   uint32  `json:"roleId"`
	EntityId *uint32 `json:"entityId"`
}
