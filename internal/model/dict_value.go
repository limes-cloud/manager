package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type DictValue struct {
	types.BaseModel
	DictID      uint32  `json:"dict_id" gorm:"not null;uniqueIndex:dv;comment:字典id"`
	Label       string  `json:"label" gorm:"not null;size:128;unique;comment:标签"`
	Value       string  `json:"value" gorm:"not null;uniqueIndex:dv;size:128;unique;comment:值"`
	Status      *bool   `json:"status" gorm:"default:true;comment:状态"`
	Weight      *uint32 `json:"weight" gorm:"default 0;comment:权重"`
	Type        string  `json:"type" gorm:"size:32;comment:字典类型"`
	Extra       string  `json:"extra" gorm:"size:256;comment:扩展信息"`
	Description string  `json:"description" gorm:"size:256;comment:描述"`
	Dict        *Dict   `json:"dict" gorm:"constraint:onDelete:cascade"`
}

// AllByDictId 通过dict_id查询所有字典值信息
func (u *DictValue) AllByDictId(ctx kratosx.Context, id uint32) ([]*DictValue, error) {
	var list []*DictValue
	return list, ctx.DB().Model(DictValue{}).Find(&list, "dict=?", id).Error
}

// Page 查询分页数据
func (u *DictValue) Page(ctx kratosx.Context, options *types.PageOptions) ([]*DictValue, uint32, error) {
	var list []*DictValue
	total := int64(0)

	db := ctx.DB().Model(u)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// Creates 批量创字典值位信息
func (u *DictValue) Creates(ctx kratosx.Context, jobs []*DictValue) error {
	return ctx.DB().Model(u).Create(&jobs).Error
}

// Create 创字典值位信息
func (u *DictValue) Create(ctx kratosx.Context) error {
	return ctx.DB().Model(u).Create(u).Error
}

// Update 更新字典值信息
func (u *DictValue) Update(ctx kratosx.Context) error {
	return ctx.DB().Updates(&u).Error
}

// DeleteByID 通过id删除字典值信息
func (u *DictValue) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}
