package dbs

import (
	"errors"
	"sync"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Dept struct{}

var (
	deptIns  *Dept
	deptOnce sync.Once
)

func NewDept() *Dept {
	deptOnce.Do(func() {
		deptIns = &Dept{}
	})
	return deptIns
}

// ListDeptClassify 获取列表
func (infra *Dept) ListDeptClassify(ctx core.Context, req *types.ListDeptClassifyRequest) ([]*entity.DeptClassify, uint32, error) {
	var (
		list  []*entity.DeptClassify
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.DeptClassify{}).Select(fs)
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateDeptClassify 创建数据
func (infra *Dept) CreateDeptClassify(ctx core.Context, tg *entity.DeptClassify) (uint32, error) {
	return tg.Id, ctx.DB().Create(tg).Error
}

// UpdateDeptClassify 更新数据
func (infra *Dept) UpdateDeptClassify(ctx core.Context, tg *entity.DeptClassify) error {
	return ctx.DB().Updates(tg).Error
}

// DeleteDeptClassify 删除数据
func (infra *Dept) DeleteDeptClassify(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.DeptClassify{}).Error
}

// GetDept 获取指定的数据
func (infra *Dept) GetDept(ctx core.Context, id uint32) (*entity.Dept, error) {
	var (
		ent = entity.Dept{}
		fs  = []string{"*"}
	)
	return &ent, ctx.DB().Select(fs).First(&ent, id).Error
}

// GetDeptByKeyword 获取指定数据
func (infra *Dept) GetDeptByKeyword(ctx core.Context, keyword string) (*entity.Dept, error) {
	var (
		m  = entity.Dept{}
		fs = []string{"*"}
	)
	return &m, ctx.DB().Select(fs).Where("keyword = ?", keyword).First(&m).Error
}

// ListDept 获取列表 fixed code
func (infra *Dept) ListDept(ctx core.Context, req *types.ListDeptRequest) ([]*entity.Dept, error) {
	var (
		ms []*entity.Dept
		fs = []string{"*"}
	)

	db := ctx.DB().Model(entity.Dept{}).
		Select(fs).
		Preload("Classify")

	if req.Name != nil {
		db = db.Where("name LIKE ?", "%"+*req.Name+"%")
	}
	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.ClassifyId != nil {
		db = db.Where("classify_id = ?", *req.ClassifyId)
	}

	return ms, db.Find(&ms).Error
}

// CreateDept 创建数据
func (infra *Dept) CreateDept(ctx core.Context, req *entity.Dept) (uint32, error) {
	return req.Id, ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(req).Error; err != nil {
			return err
		}
		if req.ParentId != 0 {
			return infra.appendDeptChildren(ctx, req.ParentId, req.Id)
		}
		return nil
	})
}

// UpdateDept 更新数据
func (infra *Dept) UpdateDept(ctx core.Context, req *entity.Dept) error {
	if req.Id == req.ParentId {
		return errors.New("父级不能为自己")
	}
	old, err := infra.GetDept(ctx, req.Id)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx core.Context) error {
		if old.ParentId != req.ParentId {
			if err := infra.removeDeptParent(ctx, req.Id); err != nil {
				return err
			}
			if err := infra.appendDeptChildren(ctx, req.ParentId, req.Id); err != nil {
				return err
			}
		}
		return ctx.DB().Updates(req).Error
	})
}

// DeleteDept 删除数据
func (infra *Dept) DeleteDept(ctx core.Context, id uint32) error {
	ids, err := infra.GetDeptChildrenIds(ctx, []uint32{id})
	if err != nil {
		return err
	}
	ids = append(ids, id)
	return ctx.DB().Where("id in ?", ids).Delete(&entity.Dept{}).Error
}

// GetDeptChildrenIds 获取指定id的所有子id
func (infra *Dept) GetDeptChildrenIds(ctx core.Context, pids []uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.DeptClosure{}).
		Select("children").
		Where("parent in ?", pids).
		Scan(&ids).Error
}

// GetDeptParentIds 获取指定id的所有父id
func (infra *Dept) GetDeptParentIds(ctx core.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.DeptClosure{}).
		Select("parent").
		Where("children=?", id).
		Scan(&ids).Error
}

// appendDeptChildren 添加id到指定的父id下
func (infra *Dept) appendDeptChildren(ctx core.Context, pid uint32, id uint32) error {
	list := []*entity.DeptClosure{
		{
			Parent:   pid,
			Children: id,
		},
	}
	ids, _ := infra.GetDeptParentIds(ctx, pid)
	for _, item := range ids {
		list = append(list, &entity.DeptClosure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Create(&list).Error
}

// removeDeptParent 删除指定id的所有父层级
func (infra *Dept) removeDeptParent(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.DeptClosure{}, "children=?", id).Error
}
