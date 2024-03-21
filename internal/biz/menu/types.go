package menu

type AllMenuRequest struct {
	NotType *string `json:"not_type"`
	App     *string `json:"app"`
	RoleId  *uint32 `json:"role_id"`
}
