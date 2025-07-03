package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type DepartmentClassify struct {
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
	types.BaseModel
}

type Department struct {
	ParentId        uint32              `json:"parentId" gorm:"column:parent_id"`
	ClassifyId      uint32              `json:"classifyId"  gorm:"column:classify_id"`
	Name            string              `json:"name" gorm:"column:name"`
	Keyword         string              `json:"keyword" gorm:"column:keyword"`
	Description     *string             `json:"description" gorm:"column:description"`
	Classify        *DepartmentClassify `json:"classify"`
	DepartmentRoles []*DepartmentRole   `json:"departmentRoles"`
	Roles           []*Role             `json:"roles" gorm:"many2many:department_role"`
	Children        []*Department       `json:"children" gorm:"-"`
	types.BaseModel
}

type DepartmentClosure struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Parent   uint32 `json:"parent" gorm:"column:parent"`
	Children uint32 `json:"children" gorm:"column:children"`
}

type DepartmentRole struct {
	DepartmentId uint32 `json:"departmentId" gorm:"column:department_id"`
	RoleId       uint32 `json:"roleId"  gorm:"column:role_id"`
}

func (u *Department) GetRoleIds() []uint32 {
	var ids []uint32
	if len(u.DepartmentRoles) != 0 {
		for _, item := range u.DepartmentRoles {
			ids = append(ids, item.RoleId)
		}
		return ids
	}
	for _, item := range u.Roles {
		ids = append(ids, item.Id)
	}
	return ids
}

// ID 获取树ID
func (m *Department) ID() uint32 {
	return m.Id
}

// Parent 获取父ID
func (m *Department) Parent() uint32 {
	return m.ParentId
}

// AppendChildren 添加子节点
func (m *Department) AppendChildren(child *Department) {
	m.Children = append(m.Children, child)
}

// ChildrenNode 获取子节点
func (m *Department) ChildrenNode() []*Department {
	var list []*Department
	list = append(list, m.Children...)
	return list
}
