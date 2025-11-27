package dbs

import (
	"sync"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type AppOAuther struct{}

var (
	aoIns  *AppOAuther
	aoOnce sync.Once
)

func NewAppOAuther() *AppOAuther {
	aoOnce.Do(func() {
		aoIns = &AppOAuther{}
	})
	return aoIns
}

// ListAppOAuther 获取列表
func (r AppOAuther) ListAppOAuther(ctx core.Context, req *types.ListAppOAutherRequest) ([]*entity.AppOAuther, uint32, error) {
	var (
		list  []*entity.AppOAuther
		total int64
	)

	db := ctx.DB().Model(entity.AppOAuther{}).Where("app_id = ?", req.AppId)

	join := ctx.DB().Where("OAuther.status = true")
	if req.Keyword != nil {
		join = join.Where("keyword LIKE ?", *req.Keyword+"%")
	}
	if req.Name != nil {
		join = join.Where("name LIKE ?", *req.Name+"%")
	}
	db = db.Joins("OAuther", join)
	db = db.Where("OAuther.status is not null")

	if req.TenantId != nil {
		db = db.Where("app_oauther.tenant_id = ?", *req.TenantId)
	}

	if req.Search != nil {
		// 查询条件下数据总数
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}

		// 分页排序
		db = page.SearchScopes(db, req.Search)
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateAppOAuther 创建数据
func (r AppOAuther) CreateAppOAuther(ctx core.Context, aoc *entity.AppOAuther) (uint32, error) {
	return aoc.Id, ctx.DB().Create(aoc).Error
}

// DeleteAppOAuther 删除数据
func (r AppOAuther) DeleteAppOAuther(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.AppOAuther{}).Error
}

func (r AppOAuther) GetAppOAutherByAT(ctx core.Context, appId uint32, tp string) (*entity.AppOAuther, error) {
	var res entity.AppOAuther
	return &res, ctx.DB().
		Where("app_id = ?", appId).
		Where("type = ?", tp).
		First(&res).Error
}
