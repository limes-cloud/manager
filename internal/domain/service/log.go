package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Log struct {
	repo repository.Log
}

func NewLog(
	repo repository.Log,
) *Log {
	uc := &Log{
		repo: repo,
	}
	return uc
}

// ListLoginLog 获取登陆日志
func (u *Log) ListLoginLog(ctx core.Context, req *types.ListLoginLogRequest) ([]*entity.LoginLog, uint32, error) {
	list, total, err := u.repo.ListLoginLog(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list log error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// ListAuthLog 获取鉴权日志
func (u *Log) ListAuthLog(ctx core.Context, req *types.ListAuthLogRequest) ([]*entity.AuthLog, uint32, error) {
	list, total, err := u.repo.ListAuthLog(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list log error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}
