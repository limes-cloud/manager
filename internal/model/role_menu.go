package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"
)

type RoleMenu struct {
	types.CreateModel
	RoleID uint32 `json:"role_id" gorm:"not null;comment:角色id"`
	MenuID uint32 `json:"menu_id" gorm:"not null;comment:菜单id"`
	Role   Role   `json:"role" gorm:"constraint:onDelete:cascade"`
	Menu   Menu   `json:"menu" gorm:"constraint:onDelete:cascade"`
}

// Update 批量更新角色所属菜单
func (rm *RoleMenu) Update(ctx kratosx.Context, roleId uint32, menuIds []uint32) error {
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
func (rm *RoleMenu) RoleMenus(ctx kratosx.Context, roleId uint32) ([]*RoleMenu, error) {
	var list []*RoleMenu
	return list, ctx.DB().Find(&list, "role_id=?", roleId).Error
}

// MenuRoles 通过菜单ID获取角色菜单列表
func (rm *RoleMenu) MenuRoles(ctx kratosx.Context, menuId uint32) ([]*RoleMenu, error) {
	var list []*RoleMenu
	return list, ctx.DB().Find(&list, "menu_id=?", menuId).Error
}

// DeleteByRoleID 通过角色id删除 角色所属菜单
func (rm *RoleMenu) DeleteByRoleID(ctx kratosx.Context, roleId uint32) error {
	return ctx.DB().Delete(rm, "role_id=?", roleId).Error
}
