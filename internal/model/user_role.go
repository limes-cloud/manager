package model

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"
)

type UserRole struct {
	CreateModel
	RoleID uint32 `json:"role_id" gour:"not null;size:32;comment:角色id"`
	UserID uint32 `json:"user_id" gour:"not null;size:32;comment:菜单id"`
	Role   *Role  `json:"role" gour:"->;constraint:OnDelete:cascade"`
}

func (ur *UserRole) OneByUserAndRole(ctx kratosx.Context, userId, roleId uint32) error {
	return ctx.DB().First(ur, "user_id=? and role_id=?", userId, roleId).Error
}

func (ur *UserRole) Add(ctx kratosx.Context) error {
	return ctx.DB().Create(&ur).Error
}

// Update 批量更新角色所属菜单
func (ur *UserRole) Update(ctx kratosx.Context, userId uint32, roleIds []uint32) error {
	// 组装新的菜单数据
	list := make([]UserRole, 0)
	for _, menuId := range roleIds {
		list = append(list, UserRole{
			UserID: userId,
			RoleID: menuId,
		})
	}

	// 删除之后再重新创建
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id=?", userId).Delete(ur).Error; err != nil {
			return err
		}
		return tx.Create(&list).Error
	})
}

// UserRoles 通过角色ID获取角色菜单
func (ur *UserRole) UserRoles(ctx kratosx.Context, userId uint32) ([]*UserRole, error) {
	var list []*UserRole
	return list, ctx.DB().Preload("Role").Find(&list, "user_id=?", userId).Error
}

// DeleteByID 通过id删除
func (ur *UserRole) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(ur, id).Error
}
