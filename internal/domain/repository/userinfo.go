package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Userinfo interface {
	// ListUserinfo 获取用户信息列表
	ListUserinfo(ctx core.Context, req *types.ListUserinfoRequest) ([]*entity.Userinfo, error)

	// UpsertUserinfo 更新用户信息
	UpsertUserinfo(ctx core.Context, list []*entity.Userinfo) error
}
