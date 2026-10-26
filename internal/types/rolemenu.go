package types

type GetRoleMenuIdsRequest struct {
	RoleId uint32  `json:"roleId"`
	App    *string `json:"app"`
}
type GetMenuRoleIdsRequest struct {
	MenuId uint32 `json:"menuId"`
}

type CreateMenuRolesRequest struct {
	RoleIds []uint32 `json:"roleIds"`
	MenuId  uint32   `json:"menuId"`
}

type CreateRoleMenusRequest struct {
	MenuIds []uint32 `json:"menuIds"`
	RoleId  uint32   `json:"roleId"`
	App     string   `json:"app"`
}

type DeleteRoleMenusRequest struct {
	MenuIds []uint32 `json:"menuIds"`
	RoleId  uint32   `json:"roleId"`
}

type DeleteMenuRolesRequest struct {
	MenuId  uint32   `json:"menuId"`
	RoleIds []uint32 `json:"roleIds"`
}
