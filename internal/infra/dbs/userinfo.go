package dbs

import (
	"sync"

	"gorm.io/gorm/clause"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Userinfo struct{}

var (
	uiIns  *Userinfo
	uiOnce sync.Once
)

func NewUserinfo() *Userinfo {
	uiOnce.Do(func() {
		uiIns = &Userinfo{}
	})
	return uiIns
}

// ListUserinfo 获取指定的数据
func (u *Userinfo) ListUserinfo(ctx core.Context, req *types.ListUserinfoRequest) ([]*entity.Userinfo, error) {
	var (
		ms    []*entity.Userinfo
		model = entity.Userinfo{UserId: req.UserId}
	)
	db := ctx.DB().Table(model.TableName()).Where("user_id = ?", req.UserId)
	if len(req.Fields) != 0 {
		db.Where("field in ?", req.Fields)
	}
	return ms, db.Scan(&ms).Error
}

// UpsertUserinfo 更新数据
func (u *Userinfo) UpsertUserinfo(ctx core.Context, list []*entity.Userinfo) error {
	return ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(list).Error
}
