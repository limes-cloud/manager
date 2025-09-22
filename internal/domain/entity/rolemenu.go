package entity

type RoleMenu struct {
	ID     uint32 `json:"id" gorm:"column:id"`
	RoleId uint32 `json:"roleId" gorm:"column:role_id"`
	MenuId uint32 `json:"menuId" gorm:"column:menu_id"`
}
