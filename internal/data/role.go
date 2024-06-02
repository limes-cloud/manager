package data

import (
	"errors"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/internal/biz/menu"
	biz "github.com/limes-cloud/manager/internal/biz/role"
	"github.com/limes-cloud/manager/internal/data/model"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"
)

type roleRepo struct {
}

func NewRoleRepo() biz.Repo {
	return &roleRepo{}
}

// ToRoleEntity model转entity
func (r roleRepo) ToRoleEntity(m *model.Role) *biz.Role {
	e := &biz.Role{}
	_ = valx.Transform(m, e)
	return e
}

// ToRoleModel entity转model
func (r roleRepo) ToRoleModel(e *biz.Role) *model.Role {
	m := &model.Role{}
	_ = valx.Transform(e, m)
	return m
}

// UpdateRoleMenu 更新所有角色的id
func (r roleRepo) UpdateRoleMenu(ctx kratosx.Context, roleId uint32, menuIds []uint32) error {
	defer func() {
		_ = ctx.Authentication().Enforce().LoadPolicy()
	}()

	role, err := r.GetRole(ctx, roleId)
	if err != nil {
		return err
	}

	var apis []model.MenuApi
	if err = ctx.DB().Model(model.Menu{}).
		Select([]string{"api", "method"}).
		Where("type=? and id in ?", menu.MenuApi, menuIds).
		Scan(&apis).Error; err != nil {
		return err
	}

	var list []*model.RoleMenu
	for _, mid := range menuIds {
		list = append(list, &model.RoleMenu{
			RoleId: roleId,
			MenuId: mid,
		})
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if err := ctx.DB().Delete(model.RoleMenu{}, "role_id=?", role.Id).Error; err != nil {
			return err
		}
		if err := ctx.DB().Create(&list).Error; err != nil {
			return err
		}
		if err := r.deleteRbac(ctx, role.Keyword); err != nil {
			return err
		}
		if len(apis) != 0 {
			return r.addRbac(ctx, role.Keyword, apis)
		}
		return nil
	})

}

// GetRoleMenuIds 获取指定角色的所有id
func (r roleRepo) GetRoleMenuIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(model.RoleMenu{}).
		Select("menu_id").
		Where("role_id=?", id).
		Scan(&ids).Error
}

// ListRole 获取列表 fixed code
func (r roleRepo) ListRole(ctx kratosx.Context, req *biz.ListRoleRequest) ([]*biz.Role, uint32, error) {
	var (
		bs    []*biz.Role
		ms    []*model.Role
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Role{}).Select(fs)
	if req.Ids != nil {
		db = db.Where("id in ?", req.Ids)
	}

	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	if err := db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: *req.OrderBy},
		Desc:   *req.Order == "desc",
	}).Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToRoleEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateRole 创建数据
func (r roleRepo) CreateRole(ctx kratosx.Context, req *biz.Role) (uint32, error) {
	m := r.ToRoleModel(req)
	return m.Id, ctx.Transaction(func(ctx kratosx.Context) error {
		if err := ctx.DB().Create(m).Error; err != nil {
			return err
		}
		return r.appendRoleChildren(ctx, req.ParentId, m.Id)
	})
}

// GetRole 获取指定的数据
func (r roleRepo) GetRole(ctx kratosx.Context, id uint32) (*biz.Role, error) {
	var (
		m  = model.Role{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToRoleEntity(&m), nil
}

// UpdateRole 更新数据
func (r roleRepo) UpdateRole(ctx kratosx.Context, req *biz.Role) error {
	if req.Id == req.ParentId {
		return errors.New("父级不能为自己")
	}
	old, err := r.GetRole(ctx, req.Id)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if old.ParentId != req.ParentId {
			if err := r.removeRoleParent(ctx, req.Id); err != nil {
				return err
			}
			if err := r.appendRoleChildren(ctx, req.ParentId, req.Id); err != nil {
				return err
			}
		}
		return ctx.DB().Updates(r.ToRoleModel(req)).Error
	})
}

// UpdateRoleStatus 更新数据状态
func (r roleRepo) UpdateRoleStatus(ctx kratosx.Context, id uint32, status bool) error {
	return ctx.DB().Model(model.Role{}).Where("id=?", id).Update("status", status).Error
}

// DeleteRole 删除数据
func (r roleRepo) DeleteRole(ctx kratosx.Context, ids []uint32) (uint32, error) {
	var del []uint32
	for _, id := range ids {
		del = append(del, id)
		childrenIds, err := r.GetRoleChildrenIds(ctx, id)
		if err != nil {
			return 0, err
		}
		del = append(del, childrenIds...)
	}
	db := ctx.DB().Where("id in ?", del).Delete(&model.Role{})
	return uint32(db.RowsAffected), db.Error
}

// GetRoleChildrenIds 获取指定id的所有子id
func (r roleRepo) GetRoleChildrenIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(model.RoleClosure{}).
		Select("children").
		Where("parent=?", id).
		Scan(&ids).Error
}

// GetRoleParentIds 获取指定id的所有父id
func (r roleRepo) GetRoleParentIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(model.RoleClosure{}).
		Select("parent").
		Where("children=?", id).
		Scan(&ids).Error
}

// appendRoleChildren 添加id到指定的父id下
func (r roleRepo) appendRoleChildren(ctx kratosx.Context, pid uint32, id uint32) error {
	list := []*model.RoleClosure{
		{
			Parent:   pid,
			Children: id,
		},
	}
	ids, _ := r.GetRoleParentIds(ctx, pid)
	for _, item := range ids {
		list = append(list, &model.RoleClosure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Create(&list).Error
}

// removeRoleParent 删除指定id的所有父层级
func (r roleRepo) removeRoleParent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&model.RoleClosure{}, "children=?", id).Error
}

// addRbac 添加权限
func (r roleRepo) addRbac(ctx kratosx.Context, role string, menus []model.MenuApi) error {
	var list []*model.CasbinRule
	for _, item := range menus {
		list = append(list, &model.CasbinRule{
			Ptype: "p",
			V0:    role,
			V1:    item.Api,
			V2:    item.Method,
		})
	}
	return ctx.DB().Create(&list).Error
}

// deleteRbac 删除权限
func (r roleRepo) deleteRbac(ctx kratosx.Context, role string) error {
	return ctx.DB().Where("v0=?", role).Delete(&model.CasbinRule{}).Error
}

// GetRoleByKeyword 获取指定数据
func (r roleRepo) GetRoleByKeyword(ctx kratosx.Context, keyword string) (*biz.Role, error) {
	var (
		m  = model.Role{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToRoleEntity(&m), nil
}

func (r roleRepo) GetRoleDataScope(ctx kratosx.Context, rid uint32) (bool, []uint32, error) {
	if rid == 1 {
		return true, []uint32{}, nil
	}
	ids, err := r.GetRoleChildrenIds(ctx, rid)
	ids = append(ids, rid)
	return false, ids, err
}
