package entity

import (
	"github.com/limes-cloud/kratosx/model"
)

type Entity struct {
	AppId    uint32         `json:"appId" gorm:"column:app_id"`
	Database string         `json:"database"  gorm:"column:database"`
	Name     string         `json:"name" gorm:"column:name"`
	Comment  string         `json:"comment" gorm:"column:comment"`
	Fields   []*EntityField `json:"fields" gorm:"foreignKey:entity_id;references:Id"`
	App      *App           `json:"app" gorm:"foreignKey:AppId;references:Id"`
	model.BaseModel
}

type EntityField struct {
	EntityId uint32 `json:"entityId" gorm:"column:entity_id"`
	Name     string `json:"name" gorm:"column:name"`
	Comment  string `json:"comment" gorm:"column:comment"`
	Index    int    `json:"index" gorm:"column:index"`
	Entity   Entity `gorm:"foreignKey:EntityId;references:Id"`
	model.BaseModel
}

type EntityRule struct {
	EntityId    uint32 `json:"entityId" gorm:"column:entity_id"`
	Name        string `json:"name" gorm:"column:name"`
	Expression  string `json:"expression" gorm:"column:expression"`
	Description string `json:"description" gorm:"column:description"`
	Entity      Entity `gorm:"foreignKey:EntityId;references:Id"`
	model.BaseTenantModel
}
