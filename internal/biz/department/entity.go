package department

import (
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/kratosx/types"
)

type Department struct {
	types.BaseModel
	ParentId    uint32        `json:"parent_id"`
	Keyword     string        `json:"keyword"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Children    []*Department `json:"children" gorm:"-"`
}

func (t *Department) ID() uint32 {
	return t.BaseModel.ID
}

func (t *Department) Parent() uint32 {
	return t.ParentId
}

func (t *Department) AppendChildren(child any) {
	team := child.(*Department)
	t.Children = append(t.Children, team)
}

func (t *Department) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range t.Children {
		list = append(list, item)
	}
	return list
}

type DepartmentClosure struct {
	ID       uint32 `json:"id"`
	Parent   uint32 `json:"parent"`
	Children uint32 `json:"children"`
}

type DepartmentObject struct {
	ID           uint32 `json:"id"`
	DepartmentId uint32 `json:"department_id"`
	ObjectId     uint32 `json:"object_id"`
	Value        string `json:"value"`
}
