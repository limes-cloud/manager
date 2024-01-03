package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type Job struct {
	types.BaseModel
	Keyword     string  `json:"keyword" gorm:"not null;size:32;unique;comment:关键字"`
	Name        string  `json:"name" gorm:"not null;size:32;unique;comment:名称"`
	Weight      *uint32 `json:"weight" gorm:"default 0;comment:权重"`
	Description string  `json:"description" gorm:"not null;size:256;comment:描述"`
}

// FindByID 通过id查询职位信息
func (u *Job) FindByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(u, id).Error
}

// FindByKeyword 通过keyword查询职位信息
func (u *Job) FindByKeyword(ctx kratosx.Context, email string) error {
	return ctx.DB().First(u, "keyword=?", email).Error
}

// Page 查询分页数据
func (u *Job) Page(ctx kratosx.Context, options *types.PageOptions) ([]*Job, uint32, error) {
	var list []*Job
	total := int64(0)

	db := ctx.DB().Model(u)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Order("weight desc").Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// Creates 批量创建职位信息
func (u *Job) Creates(ctx kratosx.Context, jobs []*Job) error {
	return ctx.DB().Model(u).Create(&jobs).Error
}

// Create 创建职位信息
func (u *Job) Create(ctx kratosx.Context) error {
	return ctx.DB().Model(u).Create(u).Error
}

// Update 更新职位信息
func (u *Job) Update(ctx kratosx.Context) error {
	return ctx.DB().Updates(&u).Error
}

// DeleteByID 通过id删除职位信息
func (u *Job) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}
