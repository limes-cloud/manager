package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type AppOAuther interface {
	// ListAppOAuther 获取应用授权列表
	ListAppOAuther(ctx core.Context, req *types.ListAppOAutherRequest) ([]*entity.AppOAuther, uint32, error)

	// GetAppOAutherByAT 获取应用信息
	GetAppOAutherByAT(ctx core.Context, appId uint32, tp string) (*entity.AppOAuther, error)

	// CreateAppOAuther 创建应用信息
	CreateAppOAuther(ctx core.Context, req *entity.AppOAuther) (uint32, error)

	// DeleteAppOAuther 删除应用信息
	DeleteAppOAuther(ctx core.Context, id uint32) error
}
