package menu

import (
	"github.com/limes-cloud/kratosx/pkg/tree"
)

type Menu struct {
	Id         uint32  `json:"id"`
	ParentId   uint32  `json:"parentId"`
	Title      string  `json:"title"`
	Type       string  `json:"type"`
	Keyword    *string `json:"keyword"`
	Icon       *string `json:"icon"`
	Api        *string `json:"api"`
	Method     *string `json:"method"`
	Path       *string `json:"path"`
	Permission *string `json:"permission"`
	Component  *string `json:"component"`
	Redirect   *string `json:"redirect"`
	Weight     *int32  `json:"weight"`
	IsHidden   *bool   `json:"isHidden"`
	IsCache    *bool   `json:"isCache"`
	IsHome     *bool   `json:"isHome"`
	IsAffix    *bool   `json:"isAffix"`
	CreatedAt  int64   `json:"createdAt"`
	UpdatedAt  int64   `json:"updatedAt"`
	Children   []*Menu `json:"Children"`
}

// ID 获取菜单树ID
func (m *Menu) ID() uint32 {
	return m.Id
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
