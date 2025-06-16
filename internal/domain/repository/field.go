package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Field interface {
	// ListField 获取用户字段列表
	ListField(ctx kratosx.Context, req *types.ListFieldRequest) ([]*entity.Field, uint32, error)

	// CreateField 创建用户字段
	CreateField(ctx kratosx.Context, req *entity.Field) (uint32, error)

	// UpdateField 更新用户字段
	UpdateField(ctx kratosx.Context, req *entity.Field) error

	// DeleteField 删除用户字段
	DeleteField(ctx kratosx.Context, id uint32) error
}
