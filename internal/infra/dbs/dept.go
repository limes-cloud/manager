package dbs

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/kratosx/pkg/cache"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Dept struct {
	cache *cache.Cache[string, struct{}]
}

var (
	deptIns      *Dept
	deptOnce     sync.Once
	deptCacheKey = "dept"
)

func NewDept() *Dept {
	deptOnce.Do(func() {
		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "dept"),
			kratosx.WithSkipDBHook(),
		)
		deptIns = &Dept{}

		ctx.BeforeStart("load cache dept", func() {
			deptIns.cache = cache.NewCacheAndInit[string, struct{}](
				ctx,
				ctx.Redis(),
				deptCacheKey,
				cache.HookLoad(func() (map[string]struct{}, error) {
					return deptIns.load(ctx)
				}),
			)
		})
	})
	return deptIns
}

// ListDeptClassify 获取列表
func (dept *Dept) ListDeptClassify(ctx core.Context, req *types.ListDeptClassifyRequest) ([]*entity.DeptClassify, uint32, error) {
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
func (dept *Dept) CreateDeptClassify(ctx core.Context, tg *entity.DeptClassify) (uint32, error) {
	return tg.Id, ctx.DB().Create(tg).Error
}

// UpdateDeptClassify 更新数据
func (dept *Dept) UpdateDeptClassify(ctx core.Context, tg *entity.DeptClassify) error {
	return ctx.DB().Updates(tg).Error
}

// DeleteDeptClassify 删除数据
func (dept *Dept) DeleteDeptClassify(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.DeptClassify{}).Error
}

// GetDept 获取指定的数据
func (dept *Dept) GetDept(ctx core.Context, id uint32) (*entity.Dept, error) {
	var (
		ent = entity.Dept{}
		fs  = []string{"*"}
	)
	return &ent, ctx.DB().Select(fs).First(&ent, id).Error
}

// GetDeptByKeyword 获取指定数据
func (dept *Dept) GetDeptByKeyword(ctx core.Context, keyword string) (*entity.Dept, error) {
	var (
		m  = entity.Dept{}
		fs = []string{"*"}
	)
	return &m, ctx.DB().Select(fs).Where("keyword = ?", keyword).First(&m).Error
}

// ListDept 获取列表 fixed code
func (dept *Dept) ListDept(ctx core.Context, req *types.ListDeptRequest) ([]*entity.Dept, error) {
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
func (dept *Dept) CreateDept(ctx core.Context, req *entity.Dept) (uint32, error) {
	return req.Id, ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(req).Error; err != nil {
			return err
		}
		if req.ParentId != 0 {
			return dept.appendDeptChildren(ctx, req.ParentId, req.Id)
		}
		return nil
	})
}

// UpdateDept 更新数据
func (dept *Dept) UpdateDept(ctx core.Context, req *entity.Dept) error {
	if req.Id == req.ParentId {
		return errors.New("父级不能为自己")
	}
	old, err := dept.GetDept(ctx, req.Id)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx core.Context) error {
		if old.ParentId != req.ParentId {
			if err := dept.removeDeptParent(ctx, req.Id); err != nil {
				return err
			}
			if err := dept.appendDeptChildren(ctx, req.ParentId, req.Id); err != nil {
				return err
			}
		}
		return ctx.DB().Updates(req).Error
	})
}

// DeleteDept 删除数据
func (dept *Dept) DeleteDept(ctx core.Context, id uint32) error {
	ids := dept.GetDeptChildrenIds([]uint32{id})
	ids = append(ids, id)
	return ctx.DB().Where("id in ?", ids).Delete(&entity.Dept{}).Error
}

// GetDeptChildrenIds 获取指定id的所有子id
func (dept *Dept) GetDeptChildrenIds(pids []uint32) []uint32 {
	var (
		ids  = make([]uint32, 0)
		keys = dept.cache.Keys()
	)
	for _, key := range keys {
		pid, cid := dept.splitCacheKey(key)
		if value.InList(pids, pid) {
			ids = append(ids, cid)
		}
	}
	return ids
}

// GetDeptParentIds 获取指定id的所有父id
func (dept *Dept) GetDeptParentIds(id uint32) []uint32 {
	var (
		ids  = make([]uint32, 0)
		keys = dept.cache.Keys()
	)
	for _, key := range keys {
		pid, cid := dept.splitCacheKey(key)
		if cid == id {
			ids = append(ids, pid)
		}
	}
	return ids
}

// appendDeptChildren 添加id到指定的父id下
func (dept *Dept) appendDeptChildren(ctx core.Context, pid uint32, id uint32) error {
	op := dept.cache.OP(ctx)

	list := []*entity.DeptClosure{
		{Parent: pid, Children: id},
	}
	ids := dept.GetDeptParentIds(pid)
	for _, item := range ids {
		list = append(list, &entity.DeptClosure{
			Parent:   item,
			Children: id,
		})
	}

	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Model(entity.DeptClosure{}).Create(&list).Error; err != nil {
			return err
		}
		for _, item := range list {
			op.Put(dept.getCacheKey(item.Parent, item.Children), struct{}{})
		}
		return op.Do()
	})
}

// removeDeptParent 删除指定id的所有父层级
func (dept *Dept) removeDeptParent(ctx core.Context, id uint32) error {
	var (
		keys = dept.cache.Keys()
		op   = dept.cache.OP(ctx)
	)

	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().
			Where("children=?", id).
			Delete(&entity.DeptClosure{}).
			Error; err != nil {
			return err
		}

		for _, key := range keys {
			_, cid := dept.splitCacheKey(key)
			if cid == id {
				op.Deletes([]string{key})
			}
		}
		return op.Do()
	})
}

// splitCacheKey 通过key获取id
func (dept *Dept) splitCacheKey(key string) (uint32, uint32) {
	arr := strings.Split(key, ":")
	if len(arr) != 2 {
		return 0, 0
	}
	jobId, _ := strconv.Atoi(arr[0])
	roleId, _ := strconv.Atoi(arr[1])
	return uint32(jobId), uint32(roleId)
}

// getCacheKey 获取缓存的key
func (dept *Dept) getCacheKey(pid uint32, cid uint32) string {
	return fmt.Sprintf("%d:%d", pid, cid)
}

// load 加载全量的数据
func (dept *Dept) load(ctx core.Context) (map[string]struct{}, error) {
	var list []*entity.DeptClosure
	if err := ctx.DB().Model(&entity.DeptClosure{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]struct{}{}
	for _, item := range list {
		bucket[dept.getCacheKey(item.Parent, item.Children)] = struct{}{}
	}

	return bucket, nil
}
