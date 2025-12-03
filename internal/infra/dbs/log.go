package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Log struct{}

var (
	logIns  *Log
	logOnce sync.Once
)

func NewLog() *Log {
	logOnce.Do(func() {
		logIns = &Log{}
	})
	return logIns
}

// ListLoginLog 获取列表 fixed code
func (log *Log) ListLoginLog(ctx core.Context, req *types.ListLoginLogRequest) ([]*entity.LoginLog, uint32, error) {
	var (
		total int64
		list  []*entity.LoginLog
	)

	db := ctx.DB().Model(entity.LoginLog{}).Preload("User").Preload("App")

	if req.Username != nil {
		db = db.Where("username = ?", *req.Username)
	}
	if len(req.CreatedAts) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", req.CreatedAts[0], req.CreatedAts[1])
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	req.Order = value.Pointer("desc")
	req.OrderBy = value.Pointer("id")
	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateLoginLog 创建数据
func (log *Log) CreateLoginLog(ctx core.Context, ll *entity.LoginLog) (uint32, error) {
	return ll.Id, ctx.DB().Create(ll).Error
}

// ListAuthLog 获取列表 fixed code
func (log *Log) ListAuthLog(ctx core.Context, req *types.ListAuthLogRequest) ([]*entity.AuthLog, uint32, error) {
	var (
		total int64
		list  []*entity.AuthLog
	)

	db := ctx.DB().Model(entity.AuthLog{}).
		Preload("User").
		Preload("Menu").
		Preload("App")

	if req.UserId != nil {
		db = db.Where("user_id = ?", *req.UserId)
	}
	if len(req.CreatedAts) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", req.CreatedAts[0], req.CreatedAts[1])
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	req.Order = value.Pointer("desc")
	req.OrderBy = value.Pointer("id")
	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateAuthLog 创建数据
func (log *Log) CreateAuthLog(ctx core.Context, al *entity.AuthLog) (uint32, error) {
	return al.Id, ctx.DB().Create(al).Error
}
