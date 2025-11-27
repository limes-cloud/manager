package dbs

import (
	"sync"

	"gorm.io/gorm"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type User struct{}

var (
	userIns  *User
	userOnce sync.Once
)

func NewUser() *User {
	userOnce.Do(func() {
		userIns = &User{}
	})
	return userIns
}

// GetUser 获取指定的数据
func (u *User) GetUser(ctx core.Context, id uint32) (*entity.User, error) {
	ent := entity.User{}
	return &ent, ctx.DB().First(&ent, id).Error
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

// ListUser 获取列表
func (u *User) ListUser(ctx core.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	var (
		total int64
		ms    []*entity.User
	)

	db := ctx.DB().Model(&entity.User{})

	if req.Username != nil {
		db = db.Where("username = ?", req.Username)
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
	if len(req.InDeptIds) > 0 || len(req.InJobIds) > 0 {
		cdb := ctx.DB().ToSQL(func(tx *gorm.DB) *gorm.DB {
			tx = tx.Model(&entity.UserDept{}).Select("user_id")
			if len(req.InDeptIds) > 0 {
				tx = tx.Where("dept_id in ?", req.InDeptIds)
			}
			if len(req.InJobIds) > 0 {
				tx = tx.Where("job_id in ?", req.InJobIds)
			}
			return tx.Scan(&[]uint32{})
		})

		db = db.Where("(id in (" + cdb + ")")
	}

	if len(req.LoggedAts) == 2 {
		db = db.Where("loggedAt between ?", req.LoggedAts)
	}
	if len(req.CreatedAts) == 2 {
		db = db.Where("createdAt between ?", req.CreatedAts)
	}
	if req.AppId != nil {
		db = db.Joins("Authorize", ctx.DB().Where("Authorize.app_id = ?", *req.AppId)).
			Where("Authorize.app_id is not null")
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
	return req.Id, ctx.DB().Create(req).Error
}

// UpdateUser 更新数据
func (u *User) UpdateUser(ctx core.Context, req *entity.User) error {
	return ctx.DB().Updates(req).Error
}

// DeleteUser 删除数据
func (u *User) DeleteUser(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.User{}).Error
}
