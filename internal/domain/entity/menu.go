package entity

import (
	"github.com/limes-cloud/kratosx/model"
	"gorm.io/gorm"
)

const (
	MenuTypeRoot   = "R"  // 应用根菜单
	MenuTypeApi    = "A"  // API接口
	MenuTypePage   = "M"  // 菜单页面
	MenuTypeGroup  = "G"  // 菜单分组
	MenuTypeBasic  = "BA" // 基础API
	MenuTypeButton = "B"  // 按钮权限
	MenuTypeLink   = "L"  // 外部链接
)

//type MenuApi struct {
//	Id     uint32 `json:"id"`
//	Api    string `json:"api"`
//	Method string `json:"method"`
//}

type Menu struct {
	AppId      uint32      `json:"appId" gorm:"column:app_id"`          // 应用ID
	ParentId   uint32      `json:"parentId" gorm:"column:parent_id"`    // 父节点ID
	Title      string      `json:"title" gorm:"column:title"`           // 菜单标题
	Type       string      `json:"type" gorm:"column:type"`             // 菜单类型
	Keyword    *string     `json:"keyword" gorm:"column:keyword"`       // 菜单标识
	Icon       *string     `json:"icon" gorm:"column:icon"`             // 菜单图标
	Api        *string     `json:"api" gorm:"column:api"`               // API接口信息
	Method     *string     `json:"method" gorm:"column:method"`         // API请求方法
	Path       *string     `json:"path" gorm:"column:path"`             // 路由路径
	Permission *string     `json:"permission" gorm:"column:permission"` // 权限标识
	Component  *string     `json:"component" gorm:"column:component"`   // 组件路径
	Redirect   *string     `json:"redirect" gorm:"column:redirect"`     // 重定向路径
	Weight     *int32      `json:"weight" gorm:"column:weight"`         // 菜单权重
	IsIframe   *bool       `json:"isIframe" gorm:"column:is_iframe"`    // 是否外链打开
	IsHidden   *bool       `json:"isHidden" gorm:"column:is_hidden"`    // 是否隐藏菜单
	IsCache    *bool       `json:"isCache" gorm:"column:is_cache"`      // 是否缓存页面
	IsHome     *bool       `json:"isHome" gorm:"column:is_home"`        // 是否首页
	IsAffix    *bool       `json:"isAffix" gorm:"column:is_affix"`      // 是否固定在标签页
	Children   []*Menu     `json:"children" gorm:"-"`                   // 子节点列表
	RoleMenus  []*RoleMenu `json:"roleMenus" gorm:"foreignKey:menu_id;references:id"`
	model.BaseModel
}

func (m *Menu) GetApi() string {
	if m.Api == nil {
		return ""
	}
	return *m.Api
}

func (m *Menu) GetMethod() string {
	if m.Method == nil {
		return ""
	}
	return *m.Method
}

type MenuClosure struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Parent   uint32 `json:"parent" gorm:"column:parent"`
	Children uint32 `json:"children" gorm:"column:children"`
}

// GormHook 获取ID
func (m *Menu) GormHook(db *gorm.DB) {
	// return m.Id
}

// ID 获取ID
func (m *Menu) ID() uint32 {
	return m.Id
}

// Parent 获取父ID
func (m *Menu) Parent() uint32 {
	return m.ParentId
}

// AppendChildren 添加子节点
func (m *Menu) AppendChildren(child *Menu) {
	m.Children = append(m.Children, child)
}

// ChildrenNode 获取子节点
func (m *Menu) ChildrenNode() []*Menu {
	var list []*Menu
	list = append(list, m.Children...)
	return list
}
