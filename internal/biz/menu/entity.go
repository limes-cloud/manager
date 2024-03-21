package menu

import (
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/kratosx/types"
)

type Menu struct {
	types.BaseModel
	ParentId        uint32  `json:"parent_id"`
	App             string  `json:"app"`
	Title           string  `json:"title"`
	Type            string  `json:"type"`
	Keyword         *string `json:"keyword"`
	Icon            *string `json:"icon"`
	Api             *string `json:"api"`
	Method          *string `json:"method"`
	Path            *string `json:"path"`
	Permission      *string `json:"permission"`
	CheckObject     *bool   `json:"check_object"`
	CheckObjectRule *string `json:"check_object_rule"`
	Component       *string `json:"component"`
	Redirect        *string `json:"redirect"`
	Weight          *uint32 `json:"weight"`
	IsHidden        *bool   `json:"is_hidden"`
	IsCache         *bool   `json:"is_cache"`
	IsHome          *bool   `json:"is_home"`
	IsAffix         *bool   `json:"is_affix"`
	Children        []*Menu `json:"children"  gorm:"-"`
}

// ID 获取菜单树ID
func (m *Menu) ID() uint32 {
	return m.BaseModel.ID
}

// Parent 获取父ID
func (m *Menu) Parent() uint32 {
	return m.ParentId
}

// AppendChildren 添加子节点
func (m *Menu) AppendChildren(child any) {
	menu := child.(*Menu)
	m.Children = append(m.Children, menu)
}

// ChildrenNode 获取子节点
func (m *Menu) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range m.Children {
		list = append(list, item)
	}
	return list
}
