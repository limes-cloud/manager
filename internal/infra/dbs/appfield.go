package dbs

import (
	"sync"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type AppField struct{}

var (
	afIns  *AppField
	afOnce sync.Once
)

func NewAppField() *AppField {
	afOnce.Do(func() {
		afIns = &AppField{}
	})
	return afIns
}

// GetAppFields 获取列表
func (r AppField) GetAppFields(ctx core.Context, appId uint32) ([]string, error) {
	var list []string
	db := ctx.DB().
		Model(entity.AppField{}).
		Select("Field.keyword").
		Where("app_id = ?", appId)
	db = db.Joins("Field", ctx.DB().Where("Field.status = true"))
	db = db.Where("Field.status is not null")

	return list, db.Scan(&list).Error
}

// ListAppField 获取列表
func (r AppField) ListAppField(ctx core.Context, req *types.ListAppFieldRequest) ([]*entity.AppField, uint32, error) {
	var (
		list  []*entity.AppField
		total int64
	)

	db := ctx.DB().Model(entity.AppField{}).Where("app_id = ?", req.AppId)

	join := ctx.DB().Where("Field.status = true")
	if req.Keyword != nil {
		join = join.Where("keyword LIKE ?", *req.Keyword+"%")
	}
	if req.Name != nil {
		join = join.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Required != nil {
		join = join.Where("app_field.required = ?", *req.Required)
	}
	db = db.Joins("Field", join)
	db = db.Where("Field.status is not null")

	if req.Search != nil {
		// 查询条件下数据总数
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}

		// 分页排序
		db = page.SearchScopes(db, req.Search)
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateAppField 创建数据
func (r AppField) CreateAppField(ctx core.Context, afc *entity.AppField) (uint32, error) {
	return afc.Id, ctx.DB().Create(afc).Error
}

// UpdateAppField 更新数据
func (r AppField) UpdateAppField(ctx core.Context, afc *entity.AppField) error {
	return ctx.DB().Updates(afc).Error
}

// DeleteAppField 删除数据
func (r AppField) DeleteAppField(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.AppField{}).Error
}
