package model

import (
	"github.com/limes-cloud/kratosx/types"
)

type Department struct {
	ParentId    uint32  `json:"parentId" gorm:"column:parent_id"`
	Name        string  `json:"name" gorm:"column:name"`
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Description *string `json:"description" gorm:"column:description"`
	types.BaseModel
}

type DepartmentClosure struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Parent   uint32 `json:"parent" gorm:"column:parent"`
	Children uint32 `json:"children" gorm:"column:children"`
}
