package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type UserForm struct{}

var (
	userFormIns  *UserForm
	userFormOnce sync.Once
)

func NewUserForm() *UserForm {
	userFormOnce.Do(func() {
		userFormIns = &UserForm{}
	})
	return userFormIns
}

// GetUserForm 获取表单（by id 或 scene）
func (r UserForm) GetUserForm(ctx core.Context, req *types.GetUserFormRequest) (*entity.UserForm, error) {
	var uf entity.UserForm
	db := ctx.DB().Model(entity.UserForm{})
	if req.Id != nil {
		db = db.Where("id = ?", *req.Id)
	}
	if req.Scene != nil {
		db = db.Where("scene = ?", *req.Scene)
	}
	return &uf, db.First(&uf).Error
}

// ListUserForm 分页查询表单列表
func (r UserForm) ListUserForm(ctx core.Context, req *types.ListUserFormRequest) ([]*entity.UserForm, uint32, error) {
	var (
		list  []*entity.UserForm
		total int64
	)

	db := ctx.DB().Model(entity.UserForm{})

	if req.Scene != nil {
		db = db.Where("scene = ?", *req.Scene)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if req.Search != nil {
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}
		db = page.SearchScopes(db, req.Search)
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateUserForm 创建表单
func (r UserForm) CreateUserForm(ctx core.Context, ent *entity.UserForm) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateUserForm 更新表单
func (r UserForm) UpdateUserForm(ctx core.Context, ent *entity.UserForm) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteUserForm 删除表单
func (r UserForm) DeleteUserForm(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.UserForm{}).Error
}
