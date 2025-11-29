package entity

import (
	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/pkg/value"
)

type DeptClassify struct {
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
	Weight      uint32  `json:"weight" gorm:"column:weight"`
	model.BaseTenantModel
}

type Dept struct {
	Id          uint32        `json:"id" gorm:"column:id;hook:dept[rud:w]"`
	ParentId    *uint32       `json:"parentId" gorm:"column:parent_id;hook:dept[cu:c]"`
	ClassifyId  uint32        `json:"classifyId"  gorm:"column:classify_id"`
	Logo        string        `json:"logo" gorm:"column:logo"`
	Keyword     string        `json:"keyword" gorm:"column:keyword"`
	Name        string        `json:"name" gorm:"column:name"`
	Status      *bool         `json:"status" gorm:"column:status"`
	Description *string       `json:"description" gorm:"column:description"`
	TenantId    uint32        `json:"tenantId" gorm:"column:tenant_id;hook:tenant"`
	CreatedAt   int64         `json:"createdAt,omitempty" gorm:"column:created_at;"`
	UpdatedAt   int64         `json:"updatedAt,omitempty" gorm:"column:updated_at;"`
	Classify    *DeptClassify `json:"classify"`
	Children    []*Dept       `json:"children" gorm:"-"`
}

type DeptClosure struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Parent   uint32 `json:"parent" gorm:"column:parent"`
	Children uint32 `json:"children" gorm:"column:children"`
}

// ID 获取树ID
func (m *Dept) ID() uint32 {
	return m.Id
}

// Parent 获取父ID
func (m *Dept) Parent() uint32 {
	return value.Value(m.ParentId)
}

// AppendChildren 添加子节点
func (m *Dept) AppendChildren(child *Dept) {
	m.Children = append(m.Children, child)
}

// ChildrenNode 获取子节点
func (m *Dept) ChildrenNode() []*Dept {
	var list []*Dept
	list = append(list, m.Children...)
	return list
}
