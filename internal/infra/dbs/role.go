package dbs

import (
	"errors"
	"sync"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Role struct{}

var (
	roleIns  *Role
	roleOnce sync.Once
)

func NewRole() *Role {
	roleOnce.Do(func() {
		roleIns = &Role{}
	})
	return roleIns
}

// GetRole 获取指定的数据
func (infra *Role) GetRole(ctx core.Context, id uint32) (*entity.Role, error) {
	var (
		ent = entity.Role{}
		fs  = []string{"*"}
	)
	return &ent, ctx.DB().Select(fs).First(&ent, id).Error
}

// GetRoleByKeyword 获取指定数据
func (infra *Role) GetRoleByKeyword(ctx core.Context, keyword string) (*entity.Role, error) {
	var (
		m  = entity.Role{}
		fs = []string{"*"}
	)
	return &m, ctx.DB().Select(fs).Where("keyword = ?", keyword).First(&m).Error
}

// ListRole 获取列表 fixed code
func (infra *Role) ListRole(ctx core.Context, req *types.ListRoleRequest) ([]*entity.Role, error) {
	var (
		ms []*entity.Role
		fs = []string{"*"}
	)

	db := ctx.DB().Model(entity.Role{}).Select(fs)

	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Keyword != nil {
		db = db.Where("keyword LIKE ?", *req.Keyword+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.InIds != nil {
		db = db.Where("id in ?", req.InIds)
	}

	return ms, db.Find(&ms).Error
}

// CreateRole 创建数据
func (infra *Role) CreateRole(ctx core.Context, req *entity.Role) (uint32, error) {
	return req.Id, ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(req).Error; err != nil {
			return err
		}
		if req.ParentId != 0 {
			return infra.appendRoleChildren(ctx, req.ParentId, req.Id)
		}
		return nil
	})
}

// UpdateRole 更新数据
func (infra *Role) UpdateRole(ctx core.Context, req *entity.Role) error {
	if req.Id == req.ParentId {
		return errors.New("父级不能为自己")
	}
	old, err := infra.GetRole(ctx, req.Id)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx core.Context) error {
		if old.ParentId != req.ParentId {
			if err := infra.removeRoleParent(ctx, req.Id); err != nil {
				return err
			}
			if err := infra.appendRoleChildren(ctx, req.ParentId, req.Id); err != nil {
				return err
			}
		}
		return ctx.DB().Updates(req).Error
	})
}

// DeleteRole 删除数据
func (infra *Role) DeleteRole(ctx core.Context, id uint32) error {
	ids, err := infra.GetRoleChildrenIds(ctx, []uint32{id})
	if err != nil {
		return err
	}
	ids = append(ids, id)
	return ctx.DB().Where("id in ?", ids).Delete(&entity.Role{}).Error
}

// GetRoleChildrenIds 获取指定id的所有子id
func (infra *Role) GetRoleChildrenIds(ctx core.Context, rids []uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.RoleClosure{}).
		Select("children").
		Where("parent in ?", rids).
		Scan(&ids).Error
}

// GetRoleParentIds 获取指定id的所有父id
func (infra *Role) GetRoleParentIds(ctx core.Context, rids []uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.RoleClosure{}).
		Select("parent").
		Where("children in ?", rids).
		Scan(&ids).Error
}

// GetRoleChildrenKeywords 获取指定id的所有子keyword
func (infra *Role) GetRoleChildrenKeywords(ctx core.Context, id uint32) ([]string, error) {
	ids, err := infra.GetRoleChildrenIds(ctx, []uint32{id})
	if err != nil {
		return nil, err
	}
	ids = append(ids, id)

	// 获取全部keyword
	var keywords []string
	return keywords, ctx.DB().Model(entity.Role{}).
		Select("keyword").
		Where("id in ?", ids).
		Scan(&keywords).Error
}

// appendRoleChildren 添加id到指定的父id下
func (infra *Role) appendRoleChildren(ctx core.Context, pid uint32, id uint32) error {
	list := []*entity.RoleClosure{
		{
			Parent:   pid,
			Children: id,
		},
	}
	ids, _ := infra.GetRoleParentIds(ctx, []uint32{pid})
	for _, item := range ids {
		list = append(list, &entity.RoleClosure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Create(&list).Error
}

// removeRoleParent 删除指定id的所有父层级
func (infra *Role) removeRoleParent(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.RoleClosure{}, "children=?", id).Error
}
