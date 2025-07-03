package dbs

import (
	"errors"
	"sync"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Department struct {
}

var (
	departmentIns  *Department
	departmentOnce sync.Once
)

func NewDepartment() *Department {
	departmentOnce.Do(func() {
		departmentIns = &Department{}
	})
	return departmentIns
}

// ListDepartmentClassify 获取列表
func (infra *Department) ListDepartmentClassify(ctx kratosx.Context, req *types.ListDepartmentClassifyRequest) ([]*entity.DepartmentClassify, uint32, error) {
	var (
		list  []*entity.DepartmentClassify
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.DepartmentClassify{}).Select(fs)
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// CreateDepartmentClassify 创建数据
func (infra *Department) CreateDepartmentClassify(ctx kratosx.Context, tg *entity.DepartmentClassify) (uint32, error) {
	return tg.Id, ctx.DB().Create(tg).Error
}

// UpdateDepartmentClassify 更新数据
func (infra *Department) UpdateDepartmentClassify(ctx kratosx.Context, tg *entity.DepartmentClassify) error {
	return ctx.DB().Updates(tg).Error
}

// DeleteDepartmentClassify 删除数据
func (infra *Department) DeleteDepartmentClassify(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.DepartmentClassify{}).Error
}

// GetDepartment 获取指定的数据
func (infra *Department) GetDepartment(ctx kratosx.Context, id uint32) (*entity.Department, error) {
	var (
		ent = entity.Department{}
		fs  = []string{"*"}
	)
	return &ent, ctx.DB().Select(fs).First(&ent, id).Error
}

// GetDepartmentByKeyword 获取指定数据
func (infra *Department) GetDepartmentByKeyword(ctx kratosx.Context, keyword string) (*entity.Department, error) {
	var (
		m  = entity.Department{}
		fs = []string{"*"}
	)
	return &m, ctx.DB().Select(fs).Where("keyword = ?", keyword).First(&m).Error
}

// ListDepartment 获取列表 fixed code
func (infra *Department) ListDepartment(ctx kratosx.Context, req *types.ListDepartmentRequest) ([]*entity.Department, error) {
	var (
		ms []*entity.Department
		fs = []string{"*"}
	)

	db := ctx.DB().Model(entity.Department{}).
		Select(fs).
		Preload("Classify").
		Preload("Roles", "status=true")

	if req.Name != nil {
		db = db.Where("name LIKE ?", "%"+*req.Name+"%")
	}
	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Ids != nil {
		db = db.Where("id in ?", req.Ids)
	}
	if req.RootId != nil {
		ids, _ := infra.GetDepartmentChildrenIds(ctx, []uint32{*req.RootId})
		db = db.Where("id in ?", ids)
	}

	if req.ClassifyId != nil {
		db = db.Where("classify_id = ?", *req.ClassifyId)
	}

	return ms, db.Find(&ms).Error
}

// CreateDepartment 创建数据
func (infra *Department) CreateDepartment(ctx kratosx.Context, req *entity.Department) (uint32, error) {
	return req.Id, ctx.Transaction(func(ctx kratosx.Context) error {
		if err := ctx.DB().Create(req).Error; err != nil {
			return err
		}
		return infra.appendDepartmentChildren(ctx, req.ParentId, req.Id)
	})
}

// UpdateDepartment 更新数据
func (infra *Department) UpdateDepartment(ctx kratosx.Context, req *entity.Department) error {
	if req.Id == req.ParentId {
		return errors.New("父级不能为自己")
	}
	old, err := infra.GetDepartment(ctx, req.Id)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if len(req.DepartmentRoles) != 0 {
			if err := ctx.DB().Where("department_id=?", req.Id).Delete(entity.DepartmentRole{}).Error; err != nil {
				return err
			}
		}
		if old.ParentId != req.ParentId {
			if err := infra.removeDepartmentParent(ctx, req.Id); err != nil {
				return err
			}
			if err := infra.appendDepartmentChildren(ctx, req.ParentId, req.Id); err != nil {
				return err
			}
		}
		return ctx.DB().Updates(req).Error
	})
}

// DeleteDepartment 删除数据
func (infra *Department) DeleteDepartment(ctx kratosx.Context, id uint32) error {
	ids, err := infra.GetDepartmentChildrenIds(ctx, []uint32{id})
	if err != nil {
		return err
	}
	ids = append(ids, id)
	return ctx.DB().Where("id in ?", ids).Delete(&entity.Department{}).Error
}

// GetDepartmentChildrenIds 获取指定id的所有子id
func (infra *Department) GetDepartmentChildrenIds(ctx kratosx.Context, pids []uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.DepartmentClosure{}).
		Select("children").
		Where("parent in ?", pids).
		Scan(&ids).Error
}

// GetDepartmentParentIds 获取指定id的所有父id
func (infra *Department) GetDepartmentParentIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.DepartmentClosure{}).
		Select("parent").
		Where("children=?", id).
		Scan(&ids).Error
}

// appendDepartmentChildren 添加id到指定的父id下
func (infra *Department) appendDepartmentChildren(ctx kratosx.Context, pid uint32, id uint32) error {
	list := []*entity.DepartmentClosure{
		{
			Parent:   pid,
			Children: id,
		},
	}
	ids, _ := infra.GetDepartmentParentIds(ctx, pid)
	for _, item := range ids {
		list = append(list, &entity.DepartmentClosure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Create(&list).Error
}

// removeDepartmentParent 删除指定id的所有父层级
func (infra *Department) removeDepartmentParent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.DepartmentClosure{}, "children=?", id).Error
}

// mergeRoleScope 合并权限
func (infra *Department) mergeRoleScope(roles []*entity.Role) (string, []uint32) {
	var (
		scope = ""
		leave = 99
		ids   []uint32
	)

	for _, item := range roles {
		tl, ok := entity.ScopeLeaves[item.DataScope]
		if !ok {
			continue
		}
		if item.DataScope == entity.ScopeCustom {
			ids = append(ids, item.GetCustomDataIds()...)
		}
		if tl > leave {
			continue
		}
		scope = item.DataScope
	}
	return scope, ids
}

// GetDepartmentRoles 获取部门角色
func (infra *Department) GetDepartmentRoles(ctx kratosx.Context, uid uint32) ([]*entity.DepartmentRole, error) {
	// 获取当前用户所在的所有部门
	var depIds []uint32
	if err := ctx.DB().
		Model(entity.UserDepartment{}).
		Select([]string{"department_id"}).
		Where("user_id = ?", uid).
		Scan(&depIds).Error; err != nil {
		return nil, err
	}

	// 获取部门的所有角色
	var drs []*entity.DepartmentRole
	if err := ctx.DB().
		Model(entity.DepartmentRole{}).
		Where("department_id in ?", depIds).
		Find(&drs).Error; err != nil {
		return nil, err
	}

	return drs, nil
}
