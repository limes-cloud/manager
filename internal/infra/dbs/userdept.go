package dbs

import (
	"context"
	"sync"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/cache"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
	"gorm.io/gorm/clause"
)

type UserDept struct {
	cache *cache.Cache[uint32, []*entity.UserDept]
}

var (
	udIns      *UserDept
	udOnce     sync.Once
	udCacheKey = "userdept"
)

func NewUserDept() *UserDept {
	udOnce.Do(func() {
		// 初始化属性
		udIns = &UserDept{}

		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "user dept"),
			kratosx.WithSkipDBHook(),
		)
		ctx.BeforeStart("load cache user dept", func() {
			udIns.cache = cache.NewCacheAndInit[uint32, []*entity.UserDept](
				ctx,
				ctx.Redis(),
				udCacheKey,
				cache.HookLoad(func() (map[uint32][]*entity.UserDept, error) {
					return udIns.load(ctx)
				}),
			)
		})
	})
	return udIns
}

// ListDeptUser 获取指定角色绑定的所有部门
func (ud *UserDept) ListDeptUser(ctx core.Context, req *types.ListDeptUserRequest) ([]*entity.User, uint32, error) {
	var (
		total int64
		list  []*entity.User
	)
	db := ctx.DB().Model(&entity.User{}).
		Joins("left join user_dept on user.id = user_dept.user_id").
		Joins("Classify", ctx.DB().Where("Classify.status=1")).
		Where("user_dept.dept_id is not null").
		Where("Classify.id is not null").
		Where("user.status = 1")

	if req.Name != nil {
		db = db.Where("user.name like ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// ListUserDept 获取指定部门的角色列表
func (ud *UserDept) ListUserDept(ctx core.Context, req *types.ListUserDeptRequest) ([]*entity.UserDept, uint32, error) {
	var (
		total int64
		list  []*entity.UserDept
	)
	db := ctx.DB().
		Model(&entity.UserDept{}).Where("user_id = ?", req.UserId).
		Joins("Dept", ctx.DB().Where("Dept.status=1")).
		Joins("Job", ctx.DB().Where("Job.status=1")).
		Where("Job.status is not null").
		Where("Dept.status is not null")

	if req.Dept != nil {
		db = db.Where("dept.name like ?", *req.Dept+"%")
	}
	if req.Job != nil {
		db = db.Where("job.name like ?", *req.Job+"%")
	}
	if req.InDeptIds != nil {
		db = db.Where("dept.id in ?", req.InDeptIds)
	}

	if req.Search != nil {
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}
		db = page.SearchScopes(db, req.Search)
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateUserDept 批量绑定部门部门角色关系
func (ud *UserDept) CreateUserDept(ctx core.Context, ent *entity.UserDept) error {
	op := ud.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Model(&entity.UserDept{}).
			Clauses(clause.OnConflict{UpdateAll: true}).
			Create(&ent).Error; err != nil {
			return err
		}

		if value.Value(ent.Main) {
			if err := ctx.DB().Model(ent).
				Where("user_id = ? and id != ?", ent.UserId, ent.Id).
				Update("main", false).Error; err != nil {
				return err
			}
		}

		// 获取当前用户的全量部门更新。
		list, _, err := ud.ListUserDept(ctx, &types.ListUserDeptRequest{
			UserId: ent.UserId,
		})
		if err != nil {
			return err
		}

		return op.Put(ent.UserId, list).Do()
	})
}

// UpdateUserDept 更新绑定部门部门角色关系
func (ud *UserDept) UpdateUserDept(ctx core.Context, ent *entity.UserDept) error {
	// 查询之前的数据
	var oud entity.UserDept
	if err := ctx.DB().Where("id = ?", ent.Id).First(&oud).Error; err != nil {
		return err
	}

	op := ud.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id = ?", ent.Id).Updates(ent).Error; err != nil {
			return err
		}

		if value.Value(ent.Main) {
			if err := ctx.DB().Model(&entity.UserDept{}).
				Where("user_id = ? and id != ?", ent.UserId, ent.Id).
				Update("main", false).Error; err != nil {
				return err
			}
		}

		// 获取当前用户的全量部门更新。
		list, _, err := ud.ListUserDept(ctx, &types.ListUserDeptRequest{
			UserId: ent.UserId,
		})
		if err != nil {
			return err
		}

		return op.Put(ent.UserId, list).Do()
	})
}

// DeleteUserDept 获取指定菜单的角色列表
func (ud *UserDept) DeleteUserDept(ctx core.Context, id uint32) error {
	ent := &entity.UserDept{}
	if err := ctx.DB().First(&ent, id).Error; err != nil {
		return err
	}

	op := ud.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("dept_id = ? AND user_id = ?", ent.DeptId, ent.UserId).Delete(&entity.UserDept{}).Error; err != nil {
			return err
		}

		// 获取当前用户的全量部门更新。
		list, _, err := ud.ListUserDept(ctx, &types.ListUserDeptRequest{
			UserId: ent.UserId,
		})
		if err != nil {
			return err
		}

		return op.Put(ent.UserId, list).Do()
	})
}

func (ud *UserDept) GetUserMainDeptId(uid uint32) uint32 {
	list, _ := ud.cache.Get(uid)
	for _, item := range list {
		if value.Value(item.Main) {
			return item.DeptId
		}
	}
	return 0
}

// GetDeptIdsByUserId 获取指定部门绑定的所有角色ID
func (ud *UserDept) GetDeptIdsByUserId(userId uint32) []uint32 {
	var ids []uint32
	list, _ := ud.cache.Get(userId)
	for _, item := range list {
		ids = append(ids, item.DeptId)
	}
	return ids
}

// ListUserDeptByUserId 获取指定部门绑定的所有角色ID
func (ud *UserDept) ListUserDeptByUserId(id uint32) []*entity.UserDept {
	list, _ := ud.cache.Get(id)
	return list
}

// load 加载全量的数据
func (ud *UserDept) load(ctx core.Context) (map[uint32][]*entity.UserDept, error) {
	var list []*entity.UserDept
	if err := ctx.DB().Model(&entity.UserDept{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[uint32][]*entity.UserDept{}
	for _, item := range list {
		bucket[item.UserId] = append(bucket[item.UserId], item)
	}

	return bucket, nil
}
