package model

import (
	"github.com/limes-cloud/kratosx/types"
)

type Role struct {
	ParentId      uint32  `json:"parentId" gorm:"column:parent_id"`
	Name          string  `json:"name" gorm:"column:name"`
	Keyword       string  `json:"keyword" gorm:"column:keyword"`
	Status        *bool   `json:"status" gorm:"column:status"`
	DataScope     string  `json:"dataScope" gorm:"column:data_scope"`
	DepartmentIds *string `json:"departmentIds" gorm:"column:department_ids"`
	Description   *string `json:"description" gorm:"column:description"`
	types.BaseModel
}

type RoleClosure struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Parent   uint32 `json:"parent" gorm:"column:parent"`
	Children uint32 `json:"children" gorm:"column:children"`
}
