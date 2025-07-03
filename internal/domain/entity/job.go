package entity

import "github.com/limes-cloud/kratosx/types"

type Job struct {
	ParentId    uint32  `json:"parentId" gorm:"column:parent_id"`
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Name        string  `json:"name" gorm:"column:name"`
	Weight      *int32  `json:"weight" gorm:"column:weight"`
	Description *string `json:"description" gorm:"column:description"`
	Children    []*Job  `json:"children" gorm:"-"`
	types.BaseModel
}

type JobClosure struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Parent   uint32 `json:"parent" gorm:"column:parent"`
	Children uint32 `json:"children" gorm:"column:children"`
}

// ID 获取树ID
func (m *Job) ID() uint32 {
	return m.Id
}

// Parent 获取父ID
func (m *Job) Parent() uint32 {
	return m.ParentId
}

// AppendChildren 添加子节点
func (m *Job) AppendChildren(child *Job) {
	m.Children = append(m.Children, child)
}

// ChildrenNode 获取子节点
func (m *Job) ChildrenNode() []*Job {
	var list []*Job
	list = append(list, m.Children...)
	return list
}
