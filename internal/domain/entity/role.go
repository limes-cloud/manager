package entity

import "github.com/limes-cloud/kratosx/model"

type Role struct {
	ParentId    uint32      `json:"parentId" gorm:"column:parent_id"`
	Keyword     string      `json:"keyword"  gorm:"column:keyword"`
	Name        string      `json:"name" gorm:"column:name"`
	Status      *bool       `json:"status" gorm:"column:status"`
	Description *string     `json:"description" gorm:"column:description"`
	Children    []*Role     `json:"children" gorm:"-"`
	RoleMenus   []*RoleMenu `json:"roleMenus" gorm:"foreignKey:role_id;references:id"`
	DeptRoles   []*DeptRole
	model.BaseTenantModel
}

type RoleClosure struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Parent   uint32 `json:"parent" gorm:"column:parent"`
	Children uint32 `json:"children" gorm:"column:children"`
}

// ID 获取树ID
func (m *Role) ID() uint32 {
	return m.Id
}

// Parent 获取父ID
func (m *Role) Parent() uint32 {
	return m.ParentId
}

// AppendChildren 添加子节点
func (m *Role) AppendChildren(child *Role) {
	m.Children = append(m.Children, child)
}

// ChildrenNode 获取子节点
func (m *Role) ChildrenNode() []*Role {
	var list []*Role
	list = append(list, m.Children...)
	return list
}
