package dbs

import (
	"sync"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Field struct{}

var (
	fieldIns  *Field
	fieldOnce sync.Once
)

func NewField() *Field {
	fieldOnce.Do(func() {
		fieldIns = &Field{}
	})
	return fieldIns
}

// ListField 获取列表
func (r Field) ListField(ctx core.Context, req *types.ListFieldRequest) ([]*entity.Field, uint32, error) {
	var (
		list  []*entity.Field
		total int64
	)

	db := ctx.DB().Model(entity.Field{})

	if req.TenantId != 0 {
		db = db.Where("tenant_id = ?", req.TenantId)
	}
	if req.Keywords != nil {
		db = db.Where("keyword in ?", req.Keywords)
	}
	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Required != nil {
		db = db.Where("required = ?", *req.Required)
	}
	if req.Unique != nil {
		db = db.Where("`unique` = ?", *req.Unique)
	}

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

// CreateField 创建数据
func (r Field) CreateField(ctx core.Context, field *entity.Field) (uint32, error) {
	return field.Id, ctx.DB().Create(field).Error
}

// GetField 获取指定的数据
func (r Field) GetField(ctx core.Context, id uint32) (*entity.Field, error) {
	var (
		field = entity.Field{}
		fs    = []string{"*"}
	)
	return &field, ctx.DB().Select(fs).First(&field, id).Error
}

// UpdateField 更新数据
func (r Field) UpdateField(ctx core.Context, field *entity.Field) error {
	return ctx.DB().Updates(field).Error
}

// DeleteField 删除数据
func (r Field) DeleteField(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Field{}).Error
}
