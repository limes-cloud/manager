package role

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	"github.com/limes-cloud/manager/internal/biz/menu"
	"github.com/limes-cloud/manager/internal/biz/role"
)

func NewRepo() role.Repo {
	return &repo{}
}

type repo struct {
}

func (r repo) AllMenuApiByIds(ctx kratosx.Context, ids []uint32) ([]*role.MenuApi, error) {
	var list []*role.MenuApi
	return list, ctx.DB().Model(menu.Menu{}).Find(&list, "id in ? and type=?", ids, menu.MenuApi).Error
}

func (r repo) GetRole(ctx kratosx.Context, id uint32) (*role.Role, error) {
	var role role.Role
	return &role, ctx.DB().First(&role, id).Error
}

func (r repo) AllRole(ctx kratosx.Context, in *role.AllRoleRequest) ([]*role.Role, error) {
	var list []*role.Role

	db := ctx.DB().Model(role.Role{})
	if in != nil {
		if in.ParentId != nil {
			ids, _ := r.GetChildrenIds(ctx, *in.ParentId)
			db = db.Where("id in ?", ids)
		}
		if in.MenuId != nil {
			db = db.Where("id in (select role_id from role_menu where menu_id = ?)", in.MenuId)
		}
	}
	return list, db.Find(&list).Error
}

func (r repo) GetParentIds(ctx kratosx.Context, rid uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(role.RoleClosure{}).
		Select("parent").
		Where("children=?", rid).
		Scan(&ids).Error
}

func (r repo) GetChildrenIds(ctx kratosx.Context, rid uint32) ([]uint32, error) {
	var ids []uint32
	if err := ctx.DB().Model(role.RoleClosure{}).
		Select("children").
		Where("parent=?", rid).
		Scan(&ids).Error; err != nil {
		return ids, err
	}
	ids = append(ids, rid)
	return ids, nil
}

func (r repo) AddRole(ctx kratosx.Context, in *role.Role) (uint32, error) {
	return in.ID(), ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in).Error; err != nil {
			return err
		}

		dcs := r.roleClosure(ctx, in.Parent(), in.ID())
		if len(dcs) != 0 {
			return tx.Create(&dcs).Error
		}
		return nil
	})
}

func (r repo) roleClosure(ctx kratosx.Context, pid uint32, id uint32) []role.RoleClosure {
	rcs := []role.RoleClosure{{
		Parent:   pid,
		Children: id,
	}}
	list, _ := r.GetParentIds(ctx, pid)
	for _, item := range list {
		rcs = append(rcs, role.RoleClosure{
			Parent:   item,
			Children: id,
		})
	}
	return rcs
}

func (r repo) UpdateRole(ctx kratosx.Context, in *role.Role) error {
	old, err := r.GetRole(ctx, in.ID())
	if err != nil {
		return err
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		// 存在层级变更
		if old.ParentId != in.ParentId {
			if err := tx.Delete(role.RoleClosure{}, "children=?", in.ID()).Error; err != nil {
				return err
			}

			dcs := r.roleClosure(ctx, in.Parent(), in.ID())
			if len(dcs) != 0 {
				return tx.Create(&dcs).Error
			}
		}
		return tx.Updates(in).Error
	})
}

func (r repo) DeleteRole(ctx kratosx.Context, id uint32) error {
	ids, err := r.GetChildrenIds(ctx, id)
	if err != nil {
		return err
	}
	return ctx.DB().Where("id in ?", ids).Delete(role.Role{}).Error
}

// ParentStatus 获取指定角色的父角色状态
func (r repo) ParentStatus(ctx kratosx.Context, id uint32) bool {
	ids, _ := r.GetParentIds(ctx, id)
	total := int64(0)
	ctx.DB().Where("id in ? and status=false", ids).Count(&total)
	return total == 0
}

func (r repo) UpdateRoleMenus(ctx kratosx.Context, roleId uint32, menuIds []uint32) error {
	list := make([]role.RoleMenu, 0)
	for _, menuId := range menuIds {
		list = append(list, role.RoleMenu{
			RoleId: roleId,
			MenuId: menuId,
		})
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id=?", roleId).Delete(role.RoleMenu{}).Error; err != nil {
			return err
		}
		return tx.Create(&list).Error
	})
}

func (r repo) GetRoleMenuIds(ctx kratosx.Context, rid uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(role.RoleMenu{}).
		Select("menu_id").
		Where("role_id=?", rid).
		Scan(&ids).Error
}
