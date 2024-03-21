package role

import (
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/kratosx/types"
)

type Role struct {
	types.BaseModel
	ParentId      uint32  `json:"parent_id"`
	Name          string  `json:"name"`
	Keyword       string  `json:"keyword"`
	Status        *bool   `json:"status"`
	Description   *string `json:"description"`
	DepartmentIds *string `json:"department_ids"`
	DataScope     string  `json:"data_scope"`
	Children      []*Role `json:"children" gorm:"-"`
}

// ID 获取ID
func (r *Role) ID() uint32 {
	return r.BaseModel.ID
}

// Parent 获取父ID
func (r *Role) Parent() uint32 {
	return r.ParentId
}

// AppendChildren 添加子树
func (r *Role) AppendChildren(child any) {
	menu := child.(*Role)
	r.Children = append(r.Children, menu)
}

// ChildrenNode 获取子树列表
func (r *Role) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range r.Children {
		list = append(list, item)
	}
	return list
}

type RoleClosure struct {
	ID       uint32 `json:"id"`
	Parent   uint32 `json:"parent"`
	Children uint32 `json:"children"`
}

type RoleMenu struct {
	types.CreateModel
	MenuId uint32 `json:"menu_id"`
	RoleId uint32 `json:"role_id"`
}

type MenuApi struct {
	Api    string `json:"api"`
	Method string `json:"method"`
}
