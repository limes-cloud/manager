package entity

import (
	"github.com/limes-cloud/kratosx/model"
)

type Dictionary struct {
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Name        string  `json:"name" gorm:"column:name"`
	Type        string  `json:"type" gorm:"column:type"`
	Description *string `json:"description" gorm:"column:description"`
	model.BaseTenantModel
}

type DictionaryValue struct {
	DictionaryId uint32             `json:"dictionaryId" gorm:"column:dictionary_id"`
	ParentId     uint32             `json:"parentId" gorm:"column:parent_id"`
	Label        string             `json:"label" gorm:"column:label"`
	Value        string             `json:"value" gorm:"column:value"`
	Status       *bool              `json:"status" gorm:"column:status"`
	Weight       *int32             `json:"weight" gorm:"column:weight"`
	Type         *string            `json:"type" gorm:"column:type"`
	Extra        *string            `json:"extra" gorm:"column:extra"`
	Description  *string            `json:"description" gorm:"column:description"`
	Dictionary   *Dictionary        `json:"dictionary"`
	Children     []*DictionaryValue `json:"children" gorm:"-"`
	model.BaseModel
}

// ID 获取ID
func (m *DictionaryValue) ID() uint32 {
	return m.Id
}

// Parent 获取父ID
func (m *DictionaryValue) Parent() uint32 {
	return m.ParentId
}

// AppendChildren 添加子节点
func (m *DictionaryValue) AppendChildren(child *DictionaryValue) {
	m.Children = append(m.Children, child)
}

// ChildrenNode 获取子节点
func (m *DictionaryValue) ChildrenNode() []*DictionaryValue {
	var list []*DictionaryValue
	list = append(list, m.Children...)
	return list
}
