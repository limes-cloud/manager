package model

import (
	"github.com/limes-cloud/manager/pkg/tree"

	"github.com/limes-cloud/kratosx"
)

type Role struct {
	BaseModel
	ParentID      uint32  `json:"parent_id"`
	Name          string  `json:"name"`
	Keyword       string  `json:"keyword"`
	Status        *bool   `json:"status"`
	Description   *string `json:"description"`
	DepartmentIds *string `json:"department_ids"`
	DataScope     string  `json:"data_scope"`
	Children      []*Role `json:"children" gorm:"-"`
}

func (r *Role) ID() uint32 {
	return r.BaseModel.ID
}

func (r *Role) Parent() uint32 {
	return r.ParentID
}

func (r *Role) AppendChildren(child any) {
	menu := child.(*Role)
	r.Children = append(r.Children, menu)
}

func (r *Role) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range r.Children {
		list = append(list, item)
	}
	return list
}

// Create 创建角色信息
func (r *Role) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(r).Error
}

// OneByID 通过ID查询角色信息
func (r *Role) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(r, "id=?", id).Error
}

// All 查询全部角色信息
func (r *Role) All(ctx kratosx.Context, scopes Scopes) ([]*Role, error) {
	var list []*Role
	db := ctx.DB()
	if scopes != nil {
		db.Scopes(scopes)
	}
	return list, db.Find(&list).Error
}

// TreeByID 查询指定角色为根节点的角色树
func (r *Role) TreeByID(ctx kratosx.Context, id uint32) (tree.Tree, error) {
	list, err := r.All(ctx, nil)
	if err != nil {
		return nil, err
	}
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	return tree.BuildTreeByID(treeList, id), nil
}

// Tree 查询角色为根节点的角色树
func (r *Role) Tree(ctx kratosx.Context) (tree.Tree, error) {
	list, err := r.All(ctx, nil)
	if err != nil {
		return nil, err
	}
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	return tree.BuildTree(treeList), nil
}

// Update 更新角色信息
func (r *Role) Update(ctx kratosx.Context) error {
	return ctx.DB().Updates(r).Error
}

// DeleteByID 通过ID删除角色信息
func (r *Role) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(r, "id=?", id).Error
}

// RoleStatus 获取角色状态
func (r *Role) RoleStatus(ctx kratosx.Context, id uint32) bool {
	roleTree, err := r.Tree(ctx)
	if err != nil {
		return false
	}
	res := false

	rt := roleTree.(*Role)
	rt.dfsRoleStatus(id, true, &res)
	return res
}

// dfsRoleStatus 深度遍历获取角色状态
func (r *Role) dfsRoleStatus(id uint32, status bool, res *bool) bool {
	if id == r.BaseModel.ID {
		is := *r.Status && status
		*res = is
	}

	for _, item := range r.Children {
		item.dfsRoleStatus(id, status && *item.Status, res)
	}
	return status
}
