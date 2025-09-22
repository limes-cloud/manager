package entity

import "github.com/limes-cloud/kratosx/model"

type RoleEntity struct {
	RoleId   uint32  `json:"roleId" gorm:"column:role_id"`
	EntityId uint32  `json:"entityId" gorm:"column:entity_id"`
	Action   string  `json:"action" gorm:"column:action"`
	Scope    string  `json:"scope" gorm:"column:scope"`
	Fields   string  `json:"fields" gorm:"column:fields"`
	Rules    string  `json:"rules" gorm:"column:rules"`
	Entity   *Entity `json:"entity" gorm:"foreignKey:EntityId;references:Id"`
	model.BaseModel
}
