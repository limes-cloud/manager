package entity

import "github.com/limes-cloud/kratosx/model"

const (
	DeptEntityName = "dept"
)

type DeptClassify struct {
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *bool   `json:"status" gorm:"column:status"`
	Weight      uint32  `json:"weight" gorm:"column:weight"`
	model.BaseTenantModel
}

type Dept struct {
	Id          uint32        `json:"id" gorm:"column:id;hook:dept[rud:w]"`
	ParentId    uint32        `json:"parentId" gorm:"column:parent_id;hook:dept[cu:c]"`
	ClassifyId  uint32        `json:"classifyId"  gorm:"column:classify_id"`
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

// TableName 获取表名称
func (m *Dept) TableName() string {
	return DeptEntityName
}

// ID 获取树ID
func (m *Dept) ID() uint32 {
	return m.Id
}

// Parent 获取父ID
func (m *Dept) Parent() uint32 {
	return m.ParentId
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
