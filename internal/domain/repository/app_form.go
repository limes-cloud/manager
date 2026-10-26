package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type AppForm interface {
	// ListAppForm 查询某应用下的表单关联列表
	ListAppForm(ctx core.Context, req *types.ListAppFormRequest) ([]*entity.AppForm, error)

	// GetAppForm 获取单条关联（by app_keyword + scene）
	GetAppForm(ctx core.Context, appKeyword string, scene string) (*entity.AppForm, error)

	// CreateAppForm 分配表单到应用
	CreateAppForm(ctx core.Context, ent *entity.AppForm) (uint32, error)

	// DeleteAppForm 取消分配（by id）
	DeleteAppForm(ctx core.Context, id uint32) error
}
