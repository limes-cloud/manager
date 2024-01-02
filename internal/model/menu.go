package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"

	"github.com/limes-cloud/manager/pkg/tree"
)

type Menu struct {
	types.BaseModel
	ParentID   uint32  `json:"parent_id" gorm:"not null;comment:父id"`
	App        string  `json:"app" gorm:"not null;type:char(32) binary;comment:服务"`
	Title      string  `json:"title" gorm:"not null;size:128;comment:标题"`
	Type       string  `json:"type" gorm:"not null;type:char(32);comment:类型"`
	Keyword    *string `json:"keyword" gorm:"default null;unique;size:64;comment:关键词"`
	Icon       *string `json:"icon" gorm:"default null;type:char(32);comment:图标"`
	Api        *string `json:"api" gorm:"default null;size:128;comment:接口"`
	Method     *string `json:"method" gorm:"default null;size:12;comment:接口方法"`
	Path       *string `json:"path" gorm:"default null;unique;size:128;comment:路径"`
	Permission *string `json:"permission" gorm:"default null;size:128;comment:指令"`
	Component  *string `json:"component" gorm:"default null;size:128;comment:组件"`
	Redirect   *string `json:"redirect" gorm:"default null;size:128;comment:重定向地址"`
	Weight     *uint32 `json:"weight" gorm:"default 0;comment:权重"`
	IsHidden   *bool   `json:"is_hidden" gorm:"default false;comment:是否隐藏"`
	IsCache    *bool   `json:"is_cache" gorm:"default false;comment:是否缓存"`
	IsHome     *bool   `json:"is_home" gorm:"default false;comment:是否为首页"`
	IsAffix    *bool   `json:"is_affix" gorm:"default false;comment:是否为标签"`
	Children   []*Menu `json:"children"  gorm:"-"`
}

// ID 获取菜单树ID
func (m *Menu) ID() uint32 {
	return m.BaseModel.ID
}

// Parent 获取父ID
func (m *Menu) Parent() uint32 {
	return m.ParentID
}

// AppendChildren 添加子节点
func (m *Menu) AppendChildren(child any) {
	menu := child.(*Menu)
	m.Children = append(m.Children, menu)
}

// ChildrenNode 获取子节点
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

// FindByID 通过id查询指定菜单
func (m *Menu) FindByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(m, id).Error
}

// All 获取全部的菜单列表
func (m *Menu) All(ctx kratosx.Context, scopes types.Scopes) ([]*Menu, error) {
	var list []*Menu
	db := ctx.DB()
	if scopes != nil {
		db = db.Scopes(scopes)
	}
	return list, db.Order("weight desc").Find(&list).Error
}

// Tree 获取菜单树
func (m *Menu) Tree(ctx kratosx.Context, scopes types.Scopes) (tree.Tree, error) {
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

// UseHome 将此菜单用于首页菜单
func (m *Menu) UseHome(ctx kratosx.Context, srvKey string, menuId uint32) error {
	return ctx.DB().Model(Menu{}).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("app=?", srvKey).Where("id!=?", menuId).Update("is_home", false).Error; err != nil {
			return err
		}
		return tx.Where("app=?", srvKey).Where("id=?", menuId).Update("is_home", true).Error
	})
}

// DeleteByIds 通过条件删除菜单
func (m *Menu) DeleteByIds(ctx kratosx.Context, ids []uint32) error {
	return ctx.DB().Where("id in ?", ids).Delete(Menu{}).Error
}
