package dbs

import (
	"fmt"
	"sync"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Field struct {
}

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
func (r Field) ListField(ctx kratosx.Context, req *types.ListFieldRequest) ([]*entity.Field, uint32, error) {
	var (
		list  []*entity.Field
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Field{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateField 创建数据
func (r Field) CreateField(ctx kratosx.Context, field *entity.Field) (uint32, error) {
	return field.Id, ctx.DB().Create(field).Error
}

// GetField 获取指定的数据
func (r Field) GetField(ctx kratosx.Context, id uint32) (*entity.Field, error) {
	var (
		field = entity.Field{}
		fs    = []string{"*"}
	)
	return &field, ctx.DB().Select(fs).First(&field, id).Error
}

// UpdateField 更新数据
func (r Field) UpdateField(ctx kratosx.Context, field *entity.Field) error {
	return ctx.DB().Updates(field).Error
}

// DeleteField 删除数据
func (r Field) DeleteField(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Field{}).Error
}
