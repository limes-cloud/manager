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

// ListUserinfo 获取指定用户的数据
func (u *Userinfo) ListUserinfo(ctx core.Context, req *types.ListUserinfoRequest) ([]*entity.Userinfo, error) {
	var (
		ms    []*entity.Userinfo
		model = entity.Userinfo{UserId: req.UserId}
	)
	db := ctx.DB().Table(model.GetTableName()).Where("user_id = ?", req.UserId)
	if len(req.Fields) != 0 {
		db = db.Where("field in ?", req.Fields)
	}
	return ms, db.Scan(&ms).Error
}

// UpsertUserinfo 批量 upsert 用户数据（按 user_id + field 唯一约束，同 key 覆盖）
func (u *Userinfo) UpsertUserinfo(ctx core.Context, list []*entity.Userinfo) error {
	if len(list) == 0 {
		return nil
	}

	return ctx.DB().Table(list[0].GetTableName()).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "field"}},
			DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
		}).
		Create(list).Error
}
