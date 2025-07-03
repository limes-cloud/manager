package entity

import (
	"strings"

	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/kratosx/types"
)

const (
	ScopeAssignAll   = "ASSIGN_ALL" // 指定层级下的所有层级
	ScopeCurrentDown = "CUR_DOWN"   // 当前层级以及下级
	ScopeDown        = "DOWN"       // 下级层级
	ScopeCustom      = "CUSTOM"     // 自定义权限
	ScopeCurrent     = "CUR"        // 当前层级
	ScopeSelf        = "SELF"       // 仅自己
)

var (
	ScopeLeaves = map[string]int{
		ScopeAssignAll:   -1,
		ScopeCustom:      -1,
		ScopeCurrentDown: 4,
		ScopeDown:        3,
		ScopeCurrent:     2,
		ScopeSelf:        1,
	}
)

type Role struct {
	ParentId      uint32  `json:"parentId,omitempty" gorm:"column:parent_id"`
	Name          string  `json:"name,omitempty" gorm:"column:name"`
	Keyword       string  `json:"keyword,omitempty" gorm:"column:keyword"`
	Status        *bool   `json:"status,omitempty" gorm:"column:status"`
	DataScope     string  `json:"dataScope,omitempty" gorm:"column:data_scope"`
	DepartmentIds *string `json:"departmentIds,omitempty" gorm:"column:department_ids"`
	JobScope      string  `json:"jobScope,omitempty" gorm:"column:job_scope"`
	JobIds        *string `json:"jobIds,omitempty" gorm:"column:job_ids"`
	Description   *string `json:"description,omitempty" gorm:"column:description"`
	Children      []*Role `json:"Children,omitempty" gorm:"-"`
	types.BaseModel
}

type RoleMenu struct {
	Id     uint32 `json:"id" gorm:"column:id"`
	MenuId uint32 `json:"menu_id" gorm:"column:menu_id"`
	RoleId uint32 `json:"role_id" gorm:"column:role_id"`
}

type RoleClosure struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Parent   uint32 `json:"parent" gorm:"column:parent"`
	Children uint32 `json:"children" gorm:"column:children"`
}

func (r *Role) GetCustomDataIds() []uint32 {
	ids := make([]uint32, 0)
	arr := strings.Split(*r.DepartmentIds, ",")
	for _, id := range arr {
		ids = append(ids, valx.ToUint32(id))
	}
	return ids
}

func (r *Role) GetCustomJobIds() []uint32 {
	ids := make([]uint32, 0)
	arr := strings.Split(*r.JobIds, ",")
	for _, id := range arr {
		ids = append(ids, valx.ToUint32(id))
	}
	return ids
}

// ID 获取ID
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
