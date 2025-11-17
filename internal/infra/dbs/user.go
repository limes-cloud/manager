package dbs

import (
	"context"
	"errors"
	"sync"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/kratosx/pkg/cache"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type User struct {
	cache *cache.Cache[uint32, uint32]
}

var (
	userIns      *User
	userOnce     sync.Once
	userCacheKey = "user"
)

func NewUser() *User {
	userOnce.Do(func() {
		userIns = &User{}
		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "user"),
			kratosx.WithSkipDBHook(),
		)
		ctx.BeforeStart("load cache user", func() {
			userIns.cache = cache.NewCacheAndInit[uint32, uint32](
				ctx,
				ctx.Redis(),
				userCacheKey,
				cache.HookLoad(func() (map[uint32]uint32, error) {
					return userIns.load(ctx)
				}),
			)
		})
	})
	return userIns
}

// GetUser 获取指定的数据
func (u *User) GetUser(ctx core.Context, id uint32) (*entity.User, error) {
	ent := entity.User{}
	return &ent, ctx.DB().
		Preload("Job").
		Preload("Dept").
		First(&ent, id).Error
}

// GetUserByUsername 获取指定数据
func (u *User) GetUserByUsername(ctx core.Context, keyword string) (*entity.User, error) {
	m := entity.User{}
	return &m, ctx.DB().
		Where("username = ?", keyword).
		First(&m).Error
}

func (u *User) GetUserByTU(ctx core.Context, tid uint32, un string) (*entity.User, error) {
	m := entity.User{}
	ctx = core.MustContext(ctx, kratosx.WithSkipDBHook())
	return &m, ctx.DB().
		Where("tenant_id = ?", tid).
		Where("username = ?", un).
		First(&m).Error
}

// ListUser 获取列表 fixed code
func (u *User) ListUser(ctx core.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	var (
		total int64
		ms    []*entity.User
	)

	db := ctx.DB().Model(entity.User{}).Preload("Job").Preload("Dept")

	if req.Username != nil {
		db = db.Where("username = ?", *req.Username)
	}
	if req.Nickname != nil {
		db = db.Where("nickname like ?", *req.Nickname+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if len(req.InIds) > 0 {
		db = db.Where("id in ?", req.InIds)
	}
	if len(req.NotInIds) > 0 {
		db = db.Where("id not in ?", req.NotInIds)
	}
	if req.InDeptIds != nil {
		db = db.Where("dept_id in ?", req.InDeptIds)
	}

	// 查询条件下数据总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 搜索排序
	db = page.SearchScopes(db, &req.Search)

	return ms, uint32(total), db.Find(&ms).Error
}

// CreateUser 创建数据
func (u *User) CreateUser(ctx core.Context, req *entity.User) (uint32, error) {
	if req.Username == entity.AdminUsername {
		return 0, errors.New("username is error")
	}
	op := u.cache.OP(ctx)
	err := ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(req).Error; err != nil {
			return err
		}
		return op.Put(req.Id, req.DeptId).Do()
	})
	if err != nil {
		return 0, err
	}
	return req.Id, nil
}

// UpdateUser 更新数据
func (u *User) UpdateUser(ctx core.Context, req *entity.User) error {
	op := u.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Updates(req).Error; err != nil {
			return err
		}
		if req.DeptId != 0 {
			return op.Put(req.Id, req.DeptId).Do()
		}
		return nil
	})
}

// DeleteUser 删除数据
func (u *User) DeleteUser(ctx core.Context, id uint32) error {
	op := u.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			return err
		}
		return op.Delete(id).Do()
	})
}

// ClearTokenExpire 清空tokens有效时间
func (u *User) ClearTokenExpire(ctx core.Context, id uint32) error {
	return ctx.DB().Model(entity.User{}).
		Where("id = ?", id).
		UpdateColumn("expire_at", 0).Error
}

func (u *User) GetUserDeptId(id uint32) uint32 {
	id, has := u.cache.Get(id)
	if !has {
		return 0
	}
	return id
}

// load 加载全量的数据
func (u *User) load(ctx core.Context) (map[uint32]uint32, error) {
	var list []*entity.User
	if err := ctx.DB().Select("id", "dept_id").Find(&list).Error; err != nil {
		return nil, err
	}
	bucket := map[uint32]uint32{}
	for _, item := range list {
		bucket[item.Id] = item.DeptId
	}
	return bucket, nil
}
