package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Auth struct{}

var (
	authIns  *Auth
	authOnce sync.Once
)

func NewAuth() *Auth {
	authOnce.Do(func() {
		authIns = &Auth{}
	})
	return authIns
}

// ListLoginLog 获取列表 fixed code
func (auth *Auth) ListLoginLog(ctx core.Context, req *types.ListLoginLogRequest) ([]*entity.LoginLog, uint32, error) {
	var (
		total int64
		fs    = []string{"*"}
		list  []*entity.LoginLog
	)

	db := ctx.DB().Model(entity.LoginLog{}).Select(fs)

	if req.Username != nil {
		db = db.Where("username = ?", *req.Username)
	}
	if len(req.CreatedAts) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", req.CreatedAts[0], req.CreatedAts[1])
	}

	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateLoginLog 创建数据
func (auth *Auth) CreateLoginLog(ctx core.Context, log *entity.LoginLog) (uint32, error) {
	return log.Id, ctx.DB().Create(log).Error
}

// ListAuthLog 获取列表 fixed code
func (auth *Auth) ListAuthLog(ctx core.Context, req *types.ListAuthLogRequest) ([]*entity.AuthLog, uint32, error) {
	var (
		total int64
		fs    = []string{"*"}
		list  []*entity.AuthLog
	)

	db := ctx.DB().Model(entity.LoginLog{}).Select(fs)

	if req.Username != nil {
		db = db.Where("username = ?", *req.Username)
	}
	if len(req.CreatedAts) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", req.CreatedAts[0], req.CreatedAts[1])
	}

	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateAuthLog 创建数据
func (auth *Auth) CreateAuthLog(ctx core.Context, log *entity.AuthLog) (uint32, error) {
	return log.Id, ctx.DB().Create(log).Error
}
