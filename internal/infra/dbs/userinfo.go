package dbs

import (
	"errors"
	"sync"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/kratosx/pkg/crypto"

	"gorm.io/gorm/clause"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Userinfo struct {
	field *Field
}

var (
	uiIns  *Userinfo
	uiOnce sync.Once
)

func NewUserinfo() *Userinfo {
	uiOnce.Do(func() {
		uiIns = &Userinfo{
			field: NewField(),
		}
	})
	return uiIns
}

// ListUserinfo 获取指定的数据
func (u *Userinfo) ListUserinfo(ctx core.Context, req *types.ListUserinfoRequest) ([]*entity.Userinfo, error) {
	var (
		ms    []*entity.Userinfo
		model = entity.Userinfo{UserId: req.UserId}
	)
	db := ctx.DB().Table(model.GetTableName()).Where("user_id = ?", req.UserId)
	if len(req.Fields) != 0 {
		db.Where("field in ?", req.Fields)
	}
	return ms, db.Scan(&ms).Error
}

// UpsertUserinfo 更新数据
func (u *Userinfo) UpsertUserinfo(ctx core.Context, list []*entity.Userinfo) error {
	if len(list) == 0 {
		return nil
	}

	// 循环当前的数组
	var fields []string
	for _, item := range list {
		item.ValueMd5 = crypto.MD5([]byte(item.Value))
		fields = append(fields, item.Field)
	}

	// 查询当前字段中需要校验唯一性的值
	respList, _, err := u.field.ListField(ctx, &types.ListFieldRequest{
		Keywords: fields,
		Unique:   value.Pointer(true),
	})
	if err != nil {
		return err
	}

	b := make(map[string]*entity.Field)
	for _, item := range respList {
		b[item.Keyword] = item
	}

	// 校验唯一性
	for _, item := range list {
		field, ok := b[item.Field]
		if !ok {
			continue
		}

		// 查询是否已经存在值
		var id int
		ctx.DB().Table(item.GetTableName()).
			Select("id").
			Where("tenant_id = ?", item.TenantId).
			Where("user_id = ?", item.UserId).
			Where("field = ?", item.Field).
			Where("value_md5 = ?", item.ValueMd5).Scan(&id)

		// 已经存在值
		if id != 0 {
			return errors.New(field.Name + "已经存在")
		}

	}

	return ctx.DB().Table(list[0].GetTableName()).Clauses(clause.OnConflict{UpdateAll: true}).Create(list).Error
}
