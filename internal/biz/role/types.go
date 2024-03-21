package role

type AllRoleRequest struct {
	ParentId *uint32 `json:"parent_id"`
	MenuId   *uint32 `json:"menu_id"`
}
