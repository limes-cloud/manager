package dbs

import (
	"sync"

	"gorm.io/gorm/clause"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type UserSetting struct{}

var (
	usIns  *UserSetting
	usOnce sync.Once
)

func NewUserSetting() *UserSetting {
	usOnce.Do(func() {
		usIns = &UserSetting{}
	})
	return usIns
}

// GetUserSetting 获取设置信息
func (u *UserSetting) GetUserSetting(ctx core.Context, req *types.GetUserSettingRequest) (*entity.UserSetting, error) {
	ms := entity.UserSetting{}
	db := ctx.DB().Model(&entity.UserSetting{UserId: req.UserId}).
		Where("app_id = ?", req.AppId).
		Where("user_id = ?", req.UserId)
	return &ms, db.First(&ms).Error
}

// UpsertUserSetting 更新数据
func (u *UserSetting) UpsertUserSetting(ctx core.Context, req *entity.UserSetting) error {
	return ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(req).Error
}
