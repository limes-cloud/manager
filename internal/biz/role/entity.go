package role

import (
	"github.com/limes-cloud/kratosx/pkg/tree"
)

type Role struct {
	Id            uint32  `json:"id"`
	ParentId      uint32  `json:"parentId"`
	Name          string  `json:"name"`
	Keyword       string  `json:"keyword"`
	Status        *bool   `json:"status"`
	DataScope     string  `json:"dataScope"`
	DepartmentIds *string `json:"departmentIds"`
	Description   *string `json:"description"`
	CreatedAt     int64   `json:"createdAt"`
	UpdatedAt     int64   `json:"updatedAt"`
	Children      []*Role `json:"Children"`
}

// ID 获取菜单树ID
func (m *Role) ID() uint32 {
	return m.Id
}

// Parent 获取父ID
func (m *Role) Parent() uint32 {
	return m.ParentId
}

// AppendChildren 添加子节点
func (m *Role) AppendChildren(child any) {
	menu := child.(*Role)
	m.Children = append(m.Children, menu)
}

// ChildrenNode 获取子节点
func (m *Role) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range m.Children {
		list = append(list, item)
	}
	return list
}
