package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Authorize interface {
	// GetAuthorize 获取授权列表
	GetAuthorize(ctx core.Context, req *types.GetAuthorizeRequest) (*entity.Authorize, error)

	// ListAuthorize 获取授权列表
	ListAuthorize(ctx core.Context, req *types.ListAuthorizeRequest) ([]*entity.Authorize, uint32, error)

	// CreateAuthorize 创建授权
	CreateAuthorize(ctx core.Context, req *entity.Authorize) (uint32, error)

	// UpdateAuthorize 更新授权
	UpdateAuthorize(ctx core.Context, req *entity.Authorize) error

	// DeleteAuthorize 删除授权
	DeleteAuthorize(ctx core.Context, id uint32) error
}
