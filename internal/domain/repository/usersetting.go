package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type UserSetting interface {
	// GetUserSetting 获取用户信息列表
	GetUserSetting(ctx core.Context, req *types.GetUserSettingRequest) (*entity.UserSetting, error)

	// UpsertUserSetting 更新用户信息
	UpsertUserSetting(ctx core.Context, req *entity.UserSetting) error
}
