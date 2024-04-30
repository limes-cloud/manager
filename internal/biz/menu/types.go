package menu

type AllMenuRequest struct {
	NotType *string `json:"not_type"`
	RoleId  *uint32 `json:"role_id"`
}
