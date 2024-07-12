package data

import (
	"errors"
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"

	biz "github.com/limes-cloud/manager/internal/biz/menu"
	"github.com/limes-cloud/manager/internal/data/model"
)

type menuRepo struct {
}

func NewMenuRepo() biz.Repo {
	return &menuRepo{}
}

// ToMenuEntity model转entity
func (r menuRepo) ToMenuEntity(m *model.Menu) *biz.Menu {
	e := &biz.Menu{}
	_ = valx.Transform(m, e)
	return e
}

// ToMenuModel entity转model
func (r menuRepo) ToMenuModel(e *biz.Menu) *model.Menu {
	m := &model.Menu{}
	_ = valx.Transform(e, m)
	return m
}

func (r menuRepo) InitBasicMenu(ctx kratosx.Context) {
	var list []*model.Menu
	if err := ctx.DB().Model(model.Menu{}).Find(&list, "type=?", biz.MenuBasic).Error; err != nil {
		ctx.Logger().Errorf("init basic api error %s", err.Error())
		return
	}
	for _, item := range list {
		if item.Method != nil && item.Api != nil {
			ctx.Authentication().AddWhitelist(*item.Api, *item.Method)
		}
	}
}

// ListMenu 获取列表
func (r menuRepo) ListMenu(ctx kratosx.Context, req *biz.ListMenuRequest) ([]*biz.Menu, uint32, error) {
	var (
		bs    []*biz.Menu
		ms    []*model.Menu
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Menu{}).Select(fs)

	if req.Title != nil {
		db = db.Where("title LIKE ?", *req.Title+"%")
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
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToMenuEntity(m))
	}
	return bs, uint32(total), nil
}

// ListMenuByRoleId 获取指定的角色的菜单
func (r menuRepo) ListMenuByRoleId(ctx kratosx.Context, id uint32) ([]*biz.Menu, uint32, error) {
	var (
		bs    []*biz.Menu
		ms    []*model.Menu
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Menu{}).Select(fs).Where("type!=?", biz.MenuBasic)
	if id != 1 {
		var ids []uint32
		if err := ctx.DB().Model(model.RoleMenu{}).
			Select("menu_id").
			Where("role_id=?", id).
			Scan(&ids).Error; err != nil {
			return nil, 0, err
		}
		db = db.Where("id in ?", ids)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Order("weight desc").Order("id asc").Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToMenuEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateMenu 创建数据 fixed code
func (r menuRepo) CreateMenu(ctx kratosx.Context, req *biz.Menu) (uint32, error) {
	m := r.ToMenuModel(req)
	if m.Type == biz.MenuRoot {
		m.ParentId = 0
	}
	return m.Id, ctx.Transaction(func(ctx kratosx.Context) error {
		if err := ctx.DB().Create(m).Error; err != nil {
			return err
		}
		if m.IsHome != nil && *m.IsHome {
			if err := r.resetHome(ctx, m.Id); err != nil {
				return err
			}
		}
		if req.Type == biz.MenuBasic && req.Api != nil && req.Method != nil {
			ctx.Authentication().AddWhitelist(*req.Api, *req.Method)
		}
		return r.appendMenuChildren(ctx, req.ParentId, m.Id)
	})
}

// GetMenu 获取指定的数据
func (r menuRepo) GetMenu(ctx kratosx.Context, id uint32) (*biz.Menu, error) {
	var (
		m  = model.Menu{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToMenuEntity(&m), nil
}

// UpdateMenu 更新数据 fixed code
func (r menuRepo) UpdateMenu(ctx kratosx.Context, req *biz.Menu) error {
	defer func() {
		_ = ctx.Authentication().Enforce().LoadPolicy()
	}()

	if req.Id == req.ParentId {
		return errors.New("父级不能为自己")
	}
	old, err := r.GetMenu(ctx, req.Id)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if old.Type == biz.MenuApi && req.Type != biz.MenuApi {
			if err := r.deleteRbac(ctx, model.MenuApi{Api: *old.Api, Method: *old.Method}); err != nil {
				return err
			}
		}

		if old.Type == biz.MenuApi && req.Type == biz.MenuApi && (*old.Method != *req.Method || *old.Api != *req.Api) {
			if err := r.updateRbac(ctx,
				model.MenuApi{Api: *old.Api, Method: *old.Method},
				model.MenuApi{Api: *req.Api, Method: *req.Method},
			); err != nil {
				return err
			}
		}

		if old.Type != biz.MenuApi && req.Type == biz.MenuApi {
			roles, err := r.allRoleKeyword(ctx, req.Id)
			if err != nil {
				return err
			}
			if err := r.addRbac(ctx, roles, model.MenuApi{Api: *old.Api, Method: *old.Method}); err != nil {
				return err
			}
		}

		if old.ParentId != req.ParentId {
			if err := r.removeMenuParent(ctx, req.Id); err != nil {
				return err
			}
			if err := r.appendMenuChildren(ctx, req.ParentId, req.Id); err != nil {
				return err
			}
		}

		if req.IsHome != nil && *req.IsHome {
			if err := r.resetHome(ctx, req.Id); err != nil {
				return err
			}
		}

		if old.Type == biz.MenuBasic {
			ctx.Authentication().RemoveWhitelist(*old.Api, *old.Method)
		}

		if req.Type == biz.MenuBasic {
			ctx.Authentication().AddWhitelist(*req.Api, *req.Method)
		}

		return ctx.DB().Updates(r.ToMenuModel(req)).Error
	})
}

// DeleteMenu 删除数据 fixed code
func (r menuRepo) DeleteMenu(ctx kratosx.Context, ids []uint32) (uint32, error) {
	defer func() {
		_ = ctx.Authentication().Enforce().LoadPolicy()
	}()

	var del []uint32
	for _, id := range ids {
		del = append(del, id)
		childrenIds, err := r.GetMenuChildrenIds(ctx, id)
		if err != nil {
			return 0, err
		}
		del = append(del, childrenIds...)
	}
	var list []*model.Menu
	if err := ctx.DB().
		Select([]string{"id", "type", "api", "method"}).
		Find(&list, "id in ?", del).Error; err != nil {
		return 0, err
	}

	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		del = []uint32{}
		for _, item := range list {
			del = append(del, item.Id)
			if item.Type == biz.MenuBasic {
				ctx.Authentication().RemoveWhitelist(*item.Api, *item.Method)
			}
			if item.Type == biz.MenuApi {
				if err := r.deleteRbac(ctx, model.MenuApi{Api: *item.Api, Method: *item.Method}); err != nil {
					return err
				}
			}
		}
		return ctx.DB().Where("id in ?", del).Delete(&model.Menu{}).Error
	}); err != nil {
		return 0, err
	}
	return uint32(len(del)), nil
}

// GetMenuChildrenIds 获取指定id的所有子id
func (r menuRepo) GetMenuChildrenIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(model.MenuClosure{}).
		Select("children").
		Where("parent=?", id).
		Scan(&ids).Error
}

// GetMenuParentIds 获取指定id的所有父id
func (r menuRepo) GetMenuParentIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(model.MenuClosure{}).
		Select("parent").
		Where("children=?", id).
		Scan(&ids).Error
}

// appendMenuChildren 添加id到指定的父id下
func (r menuRepo) appendMenuChildren(ctx kratosx.Context, pid uint32, id uint32) error {
	list := []*model.MenuClosure{
		{
			Parent:   pid,
			Children: id,
		},
	}
	ids, _ := r.GetMenuParentIds(ctx, pid)
	for _, item := range ids {
		list = append(list, &model.MenuClosure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Create(&list).Error
}

// removeMenuParent 删除指定id的所有父层级
func (r menuRepo) removeMenuParent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&model.MenuClosure{}, "children=?", id).Error
}

// resetHome 重置当前树的首页节点
func (r menuRepo) resetHome(ctx kratosx.Context, id uint32) error {
	pids, err := r.GetMenuParentIds(ctx, id)
	if err != nil {
		return err
	}

	var menu model.Menu
	if err = ctx.DB().
		Where("id in ?", pids).
		Where("type=?", biz.MenuRoot).
		Find(&menu).Error; err != nil {
		return err
	}

	cids, err := r.GetMenuChildrenIds(ctx, menu.Id)
	if err != nil {
		return err
	}
	pids = append(pids, cids...)
	return ctx.DB().Model(model.Menu{}).Where("id in ?", pids).Update("is_home", false).Error
}

func (r menuRepo) addRbac(ctx kratosx.Context, roles []string, req model.MenuApi) error {
	var list []*model.CasbinRule
	for _, role := range roles {
		list = append(list, &model.CasbinRule{
			Ptype: "p",
			V0:    role,
			V1:    req.Api,
			V2:    req.Method,
		})
	}
	return ctx.DB().Create(&list).Error
}

func (r menuRepo) deleteRbac(ctx kratosx.Context, req model.MenuApi) error {
	return ctx.DB().Where("v1=? and v2=?", req.Api, req.Method).Delete(model.CasbinRule{}).Error
}

func (r menuRepo) updateRbac(ctx kratosx.Context, old model.MenuApi, now model.MenuApi) error {
	return ctx.DB().
		Model(model.CasbinRule{}).
		Where("v1=? and v2=?", old.Api, old.Method).
		UpdateColumn("v1", now.Api).
		UpdateColumn("v2", now.Method).
		Error
}

func (r menuRepo) allRoleKeyword(ctx kratosx.Context, id uint32) ([]string, error) {
	var (
		keys []string
		ids  []uint32
	)

	if err := ctx.DB().Model(model.RoleMenu{}).
		Scan("menu_id").
		Where("role_id=?", id).
		Scan(&ids).Error; err != nil {
		return nil, err
	}

	return keys, ctx.DB().Model(model.Role{}).Select("keyword").
		Where("id in ?", ids).
		Scan(&keys).
		Error
}
