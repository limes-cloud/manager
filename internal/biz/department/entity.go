package department

import (
	"github.com/limes-cloud/kratosx/pkg/tree"
)

type Department struct {
	Id          uint32        `json:"id"`
	ParentId    uint32        `json:"parentId"`
	Name        string        `json:"name"`
	Keyword     string        `json:"keyword"`
	Description *string       `json:"description"`
	CreatedAt   int64         `json:"createdAt"`
	UpdatedAt   int64         `json:"updatedAt"`
	Children    []*Department `json:"Children"`
}

// ID 获取菜单树ID
func (m *Department) ID() uint32 {
	return m.Id
}

// Parent 获取父ID
func (m *Department) Parent() uint32 {
	return m.ParentId
}

// AppendChildren 添加子节点
func (m *Department) AppendChildren(child any) {
	menu := child.(*Department)
	m.Children = append(m.Children, menu)
}

// ChildrenNode 获取子节点
func (m *Department) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range m.Children {
		list = append(list, item)
	}
	return list
}
