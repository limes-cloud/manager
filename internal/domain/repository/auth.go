package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Auth interface {
	ListLoginLog(ctx core.Context, req *types.ListLoginLogRequest) ([]*entity.LoginLog, uint32, error)
	CreateLoginLog(ctx core.Context, log *entity.LoginLog) (uint32, error)
	ListAuthLog(ctx core.Context, req *types.ListAuthLogRequest) ([]*entity.AuthLog, uint32, error)
	CreateAuthLog(ctx core.Context, log *entity.AuthLog) (uint32, error)
}
