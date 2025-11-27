package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type AppField interface {
	// GetAppFields 获取应用字段列表
	GetAppFields(ctx core.Context, appId uint32) ([]string, error)

	// ListAppField 获取应用字段列表
	ListAppField(ctx core.Context, req *types.ListAppFieldRequest) ([]*entity.AppField, uint32, error)

	// CreateAppField 创建应用字段信息
	CreateAppField(ctx core.Context, req *entity.AppField) (uint32, error)

	// UpdateAppField 更新应用字段信息
	UpdateAppField(ctx core.Context, req *entity.AppField) error

	// DeleteAppField 删除应用字段信息
	DeleteAppField(ctx core.Context, id uint32) error
}
