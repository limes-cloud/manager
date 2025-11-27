package entity

import "github.com/limes-cloud/kratosx/model"

const (
	ActionCreate = "create"
	ActionRead   = "read"
	ActionUpdate = "update"
	ActionDelete = "delete"
)

type RoleEntity struct {
	RoleId   uint32  `json:"roleId" gorm:"column:role_id"`
	EntityId uint32  `json:"entityId" gorm:"column:entity_id"`
	Action   string  `json:"action" gorm:"column:action"`
	Scope    string  `json:"scope" gorm:"column:scope"`
	Depts    string  `json:"depts" gorm:"column:depts"`
	Fields   string  `json:"fields" gorm:"column:fields"`
	Rules    string  `json:"rules" gorm:"column:rules"`
	Entity   *Entity `json:"entity" gorm:"foreignKey:EntityId;references:Id"`
	model.BaseModel
}
