package model

import (
	"github.com/limes-cloud/manager/pkg/tree"

	"github.com/limes-cloud/kratosx"
)

type Department struct {
	BaseModel
	ParentID    uint32        `json:"parent_id"`
	Keyword     string        `json:"keyword"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Children    []*Department `json:"children"  gorm:"-"`
}

func (t *Department) ID() uint32 {
	return t.BaseModel.ID
}

func (t *Department) Parent() uint32 {
	return t.ParentID
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

// Create 创建部门
func (t *Department) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(t).Error
}

// Tree 获取部门树
func (t *Department) Tree(ctx kratosx.Context) (tree.Tree, error) {
	// 获取部门列表
	list, err := t.All(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 根据列表构建部门树
	var trees []tree.Tree
	for _, item := range list {
		trees = append(trees, item)
	}
	return tree.BuildTree(trees), nil
}

func (t *Department) TreeByID(ctx kratosx.Context, id uint32) (tree.Tree, error) {
	// 获取部门列表
	list, err := t.All(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 根据列表构建部门树
	var trees []tree.Tree
	for _, item := range list {
		trees = append(trees, item)
	}
	return tree.BuildTreeByID(trees, id), nil
}

// All 获取全部部门
func (t *Department) All(ctx kratosx.Context, scopes Scopes) ([]*Department, error) {
	list := make([]*Department, 0)
	db := ctx.DB().Model(Department{})
	if scopes != nil {
		db = scopes(db)
	}
	return list, db.Find(&list).Error
}

// OneByID 通过部门id查询
func (t *Department) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(t, "id=?", id).Error
}

// OneByKeyword 通过部门keyword查询
func (t *Department) OneByKeyword(ctx kratosx.Context, keyword string) error {
	return ctx.DB().First(t, "keyword=?", keyword).Error
}

// Update 更新部门信息
func (t *Department) Update(ctx kratosx.Context) error {
	return ctx.DB().Updates(t).Error
}

// DeleteByID 通过id删除指定部门 以及部门下的部门
func (t *Department) DeleteByID(ctx kratosx.Context, id uint32) error {
	list, err := t.All(ctx, nil)
	if err != nil {
		return err
	}

	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	team := tree.BuildTreeByID(treeList, id)
	ids := tree.GetTreeID(team)

	// 进行数据删除
	return ctx.DB().Where("id in ?", ids).Delete(t).Error
}
