package model

import (
	"github.com/limes-cloud/kratosx/types"
)

type Menu struct {
	ParentId   uint32  `json:"parentId" gorm:"column:parent_id"`
	Title      string  `json:"title" gorm:"column:title"`
	Type       string  `json:"type" gorm:"column:type"`
	Keyword    *string `json:"keyword" gorm:"column:keyword"`
	Icon       *string `json:"icon" gorm:"column:icon"`
	Api        *string `json:"api" gorm:"column:api"`
	Method     *string `json:"method" gorm:"column:method"`
	Path       *string `json:"path" gorm:"column:path"`
	Permission *string `json:"permission" gorm:"column:permission"`
	Component  *string `json:"component" gorm:"column:component"`
	Redirect   *string `json:"redirect" gorm:"column:redirect"`
	Weight     *int32  `json:"weight" gorm:"column:weight"`
	IsHidden   *bool   `json:"isHidden" gorm:"column:is_hidden"`
	IsCache    *bool   `json:"isCache" gorm:"column:is_cache"`
	IsHome     *bool   `json:"isHome" gorm:"column:is_home"`
	IsAffix    *bool   `json:"isAffix" gorm:"column:is_affix"`
	types.BaseModel
}

type MenuClosure struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Parent   uint32 `json:"parent" gorm:"column:parent"`
	Children uint32 `json:"children" gorm:"column:children"`
}
