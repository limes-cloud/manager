package dbs

import (
	"sync"

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
	return &ent, ctx.DB().Preload("Job").Preload("Dept").First(&ent, id).Error
}

// GetUserByUsername 获取指定数据
func (u *User) GetUserByUsername(ctx core.Context, keyword string) (*entity.User, error) {
	var (
		m  = entity.User{}
		fs = []string{"*"}
	)
	return &m, ctx.DB().Select(fs).Where("username = ?", keyword).First(&m).Error
}

// ListUser 获取列表 fixed code
func (u *User) ListUser(ctx core.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	var (
		total int64
		ms    []*entity.User
	)

	db := ctx.DB().Model(entity.User{}).Preload("Job").Preload("Dept")

	if req.Username != nil {
		db = db.Where("username = ?", req.Username)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
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
