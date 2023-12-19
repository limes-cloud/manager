package model

import (
	"github.com/limes-cloud/manager/pkg/tree"

	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"
)

type Menu struct {
	BaseModel
	ParentID   uint32  `json:"parent_id"`
	App        string  `json:"app"`
	Title      string  `json:"title"`
	Type       string  `json:"type"`
	Icon       *string `json:"icon"`
	Path       *string `json:"path"`
	Keyword    *string `json:"keyword"`
	Permission *string `json:"permission"`
	Component  *string `json:"component"`
	Api        *string `json:"api"`
	Method     *string `json:"method"`
	Redirect   *string `json:"redirect"`
	Weight     *int    `json:"weight"`
	IsHidden   *bool   `json:"is_hidden"`
	IsCache    *bool   `json:"is_cache"`
	IsHome     *bool   `json:"is_home"`
	IsAffix    *bool   `json:"is_affix"`
	Children   []*Menu `json:"children"  gorm:"-"`
}

func (m *Menu) ID() uint32 {
	return m.BaseModel.ID
}

func (m *Menu) Parent() uint32 {
	return m.ParentID
}

func (m *Menu) AppendChildren(child any) {
	menu := child.(*Menu)
	m.Children = append(m.Children, menu)
}

func (m *Menu) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range m.Children {
		list = append(list, item)
	}
	return list
}

// Create 创建菜单
func (m *Menu) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(m).Error
}

// OneByID 通过id查询指定菜单
func (m *Menu) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(m, id).Error
}

// OneByKeyword 通过keyword条件查询指定菜单
func (m *Menu) OneByKeyword(ctx kratosx.Context, keyword string) error {
	return ctx.DB().First(m, "keyword=?", keyword).Error
}

// All 获取全部的菜单列表
func (m *Menu) All(ctx kratosx.Context, scopes Scopes) ([]*Menu, error) {
	var list []*Menu
	db := ctx.DB()
	if scopes != nil {
		db = db.Scopes(scopes)
	}
	return list, db.Order("weight desc").Find(&list).Error
}

// Tree 获取菜单树
func (m *Menu) Tree(ctx kratosx.Context, scopes Scopes) (tree.Tree, error) {
	list, err := m.All(ctx, scopes)

	if err != nil {
		return nil, err
	}
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	return tree.BuildTree(treeList), nil
}

// Update 更新菜单
func (m *Menu) Update(ctx kratosx.Context) error {
	return ctx.DB().Updates(m).Error
}

// UseHome 讲此菜单用于首页菜单
func (m *Menu) UseHome(ctx kratosx.Context, srvKey string, menuId uint32) error {
	return ctx.DB().Model(Menu{}).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("server_keyword=?", srvKey).
			Where("id!=?", menuId).
			Update("is_home", false).Error; err != nil {
			return err
		}

		return tx.Where("server_keyword=?", srvKey).
			Where("id=?", menuId).
			Update("is_home", true).Error
	})
}

// DeleteByIds 通过条件删除菜单
func (m *Menu) DeleteByIds(ctx kratosx.Context, ids []uint32) error {
	return ctx.DB().Where("id in ?", ids).Delete(Menu{}).Error
}
