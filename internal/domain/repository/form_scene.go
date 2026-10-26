package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type FormScene interface {
	// GetFormScene 获取场景（by id）
	GetFormScene(ctx core.Context, id uint32) (*entity.FormScene, error)

	// GetFormSceneByKeyword 获取场景（by keyword）
	GetFormSceneByKeyword(ctx core.Context, keyword string) (*entity.FormScene, error)

	// ListFormScene 查询场景列表
	ListFormScene(ctx core.Context, req *types.ListFormSceneRequest) ([]*entity.FormScene, uint32, error)

	// CreateFormScene 创建场景
	CreateFormScene(ctx core.Context, ent *entity.FormScene) (uint32, error)

	// UpdateFormScene 更新场景
	UpdateFormScene(ctx core.Context, ent *entity.FormScene) error

	// DeleteFormScene 删除场景
	DeleteFormScene(ctx core.Context, id uint32) error
}
