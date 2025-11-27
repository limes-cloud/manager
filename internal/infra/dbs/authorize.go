package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Authorize struct{}

var (
	azIns  *Authorize
	azOnce sync.Once
)

func NewAuthorize() *Authorize {
	azOnce.Do(func() {
		azIns = &Authorize{}
	})
	return azIns
}

func (infra *Authorize) GetAuthorize(ctx core.Context, req *types.GetAuthorizeRequest) (*entity.Authorize, error) {
	var res entity.Authorize
	return &res, ctx.DB().
		Where("tenant_id = ?", req.TenantId).
		Where("user_id = ?", req.UserId).
		Where("app_id = ?", req.AppId).
		First(&res).Error
}

// ListAuthorize 获取列表
func (infra *Authorize) ListAuthorize(ctx core.Context, req *types.ListAuthorizeRequest) ([]*entity.Authorize, uint32, error) {
	var (
		list  []*entity.Authorize
		total int64
	)

	db := ctx.DB().Model(entity.Authorize{})
	if req.UserId != nil {
		db = db.Where("user_id = ?", *req.UserId)
	}
	if req.AppIds != nil {
		db = db.Where("app_id in ?", req.AppIds)
	}
	if req.Search != nil {
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}
		db = page.SearchScopes(db, req.Search)
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateAuthorize 创建数据
func (infra *Authorize) CreateAuthorize(ctx core.Context, tg *entity.Authorize) (uint32, error) {
	return tg.Id, ctx.DB().Create(tg).Error
}

// UpdateAuthorize 更新数据
func (infra *Authorize) UpdateAuthorize(ctx core.Context, tg *entity.Authorize) error {
	return ctx.DB().Updates(tg).Error
}

// DeleteAuthorize 删除数据
func (infra *Authorize) DeleteAuthorize(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Authorize{}).Error
}
