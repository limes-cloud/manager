package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Log interface {
	// ListLoginLog 获取登陆日志
	ListLoginLog(ctx core.Context, req *types.ListLoginLogRequest) ([]*entity.LoginLog, uint32, error)

	// CreateLoginLog 创建登陆日志
	CreateLoginLog(ctx core.Context, log *entity.LoginLog) (uint32, error)

	// ListAuthLog 获取鉴权日志
	ListAuthLog(ctx core.Context, req *types.ListAuthLogRequest) ([]*entity.AuthLog, uint32, error)

	// CreateAuthLog 创建鉴权日志
	CreateAuthLog(ctx core.Context, log *entity.AuthLog) (uint32, error)
}
