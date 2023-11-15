package model

import (
	"github.com/limes-cloud/kratos"
	"gorm.io/gorm"
)

type RoleMenu struct {
	CreateModel
	RoleID uint32 `json:"role_id"`
	MenuID uint32 `json:"menu_id"`
	Role   Role   `json:"role" gorm:"->;constraint:OnDelete:cascade"`
	Menu   Menu   `json:"menu" gorm:"->;constraint:OnDelete:cascade"`
}

// Update 批量更新角色所属菜单
func (rm *RoleMenu) Update(ctx kratos.Context, roleId uint32, menuIds []uint32) error {

	// 组装新的菜单数据
	list := make([]RoleMenu, 0)
	for _, menuId := range menuIds {
		list = append(list, RoleMenu{
			RoleID: roleId,
			MenuID: menuId,
		})
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id=?", roleId).Delete(rm).Error; err != nil {
			return err
		}
		return tx.Create(&list).Error
	})
}

// RoleMenus 通过角色ID获取角色菜单
func (rm *RoleMenu) RoleMenus(ctx kratos.Context, roleId uint32) ([]*RoleMenu, error) {
	var list []*RoleMenu
	return list, ctx.DB().Find(&list, "role_id=?", roleId).Error
}

// MenuRoles 通过菜单ID获取角色菜单列表
func (rm *RoleMenu) MenuRoles(ctx kratos.Context, menuId uint32) ([]*RoleMenu, error) {
	var list []*RoleMenu
	return list, ctx.DB().Find(&list, "menu_id=?", menuId).Error
}

// DeleteByRoleID 通过角色id删除 角色所属菜单
func (rm *RoleMenu) DeleteByRoleID(ctx kratos.Context, roleId uint32) error {
	return ctx.DB().Delete(rm, "role_id=?", roleId).Error
}
