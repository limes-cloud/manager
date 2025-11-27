package dbs

import (
	"sync"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
)

const (
	tokenKey = "oauth:token:"
)

type UserOAuther struct{}

var (
	uoIns  *UserOAuther
	uoOnce sync.Once
)

func NewUserOAuther() *UserOAuther {
	uoOnce.Do(func() {
		uoIns = &UserOAuther{}
	})
	return uoIns
}

func (oa UserOAuther) CreateUserOAuther(ctx core.Context, ua *entity.UserOAuther) (uint32, error) {
	return ua.Id, ctx.DB().Create(ua).Error
}

func (oa UserOAuther) UpdateUserOAuther(ctx core.Context, ua *entity.UserOAuther) error {
	return ctx.DB().
		Where("user_id=?", ua.UserId).
		Where("oauther_id=?", ua.OAutherId).
		Updates(ua).Error
}

func (oa UserOAuther) DeleteUserOAuther(ctx core.Context, userId uint32, oautherId uint32) error {
	return ctx.DB().Where("user_id = ? and oauther_id = ?", userId, oautherId).Delete(&entity.UserOAuther{}).Error
}

func (oa UserOAuther) IsBindUserOAuther(ctx core.Context, cid uint32, oid string) bool {
	var count int64
	ctx.DB().Model(entity.UserOAuther{}).
		Where("user_id is not null").
		Where("oauther_id=?", cid).
		Where("oid=?", oid).
		Count(&count)
	return count != 0
}

// GetUserOAutherByCO 通过oauther_id 和 oid获取 授权信息
func (oa UserOAuther) GetUserOAutherByCO(ctx core.Context, cid uint32, oid string) (*entity.UserOAuther, error) {
	var ua entity.UserOAuther
	return &ua, ctx.DB().Where("oauther_id=?", cid).Where("oid=?", oid).First(&ua).Error
}
