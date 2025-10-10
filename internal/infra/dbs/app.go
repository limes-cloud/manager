package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"

	"gorm.io/gorm"
)

type App struct{}

var (
	appIns  *App
	appOnce sync.Once
)

func NewApp() *App {
	appOnce.Do(func() {
		appIns = &App{}
	})
	return appIns
}

// GetAppByKeyword 获取指定数据
func (r App) GetAppByKeyword(ctx core.Context, keyword string) (*entity.App, error) {
	var (
		app = entity.App{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &app, db.Where("keyword = ?", keyword).First(&app).Error
}

// GetApp 获取指定的数据
func (r App) GetApp(ctx core.Context, id uint32) (*entity.App, error) {
	var (
		app = entity.App{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &app, db.Where("id = ?", id).First(&app).Error
}

// ListApp 获取列表
func (r App) ListApp(ctx core.Context, req *types.ListAppRequest) ([]*entity.App, uint32, error) {
	var (
		list  []*entity.App
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.App{}).Select(fs)
	if req.InIds != nil {
		db = db.Where("id IN ?", req.InIds)
	}
	if req.Keyword != nil {
		db = db.Where("keyword LIKE ?", *req.Keyword+"%")
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if req.Search != nil {
		// 查询条件下数据总数
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}
		// 分页排序
		db = page.SearchScopes(db, req.Search)
		if err := db.Find(&list).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err := db.Find(&list).Error; err != nil {
			return nil, 0, err
		}
		total = int64(len(list))
	}

	return list, uint32(total), nil
}

// CreateApp 创建数据
func (r App) CreateApp(ctx core.Context, app *entity.App) (uint32, error) {
	return app.Id, ctx.DB().Create(app).Error
}

// UpdateApp 更新数据
func (r App) UpdateApp(ctx core.Context, app *entity.App) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		//if len(app.AppChannels) != 0 {
		//	if err := tx.Where("app_id=?", app.Id).Delete(entity.AppChannel{}).Error; err != nil {
		//		return err
		//	}
		//}
		//
		//if len(app.AppFields) != 0 {
		//	if err := tx.Where("app_id=?", app.Id).Delete(entity.AppField{}).Error; err != nil {
		//		return err
		//	}
		//}

		return tx.Updates(app).Error
	})
}

// DeleteApp 删除数据
func (r App) DeleteApp(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.App{}).Error
}

// ListAppOAuthChannel 创建数据
func (r App) ListAppOAuthChannel(ctx core.Context, req *types.ListAppOAuthChannelRequest) ([]*entity.AppOAuthChannel, uint32, error) {
	var (
		list  []*entity.AppOAuthChannel
		total int64
	)

	db := ctx.DB().Model(entity.AppOAuthChannel{}).Where("app_id = ?", req.AppId)

	join := ctx.DB().Where("Channel.status = true")
	if req.Keyword != nil {
		join = join.Where("keyword LIKE ?", *req.Keyword+"%")
	}
	if req.Name != nil {
		join = join.Where("name LIKE ?", *req.Name+"%")
	}
	db = db.Joins("Channel", join)
	db = db.Where("Channel.status is not null")

	// 查询条件下数据总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = page.SearchScopes(db, &req.Search)
	if err := db.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, uint32(total), nil
}

func (r App) ListTenantAppOAuthChannel(ctx core.Context, req *types.ListTenantAppOAuthChannelRequest) ([]*entity.AppOAuthChannel, error) {
	var list []*entity.AppOAuthChannel
	db := ctx.DB().Model(entity.AppOAuthChannel{}).
		Where("`app_oauth_channel`.`app_id` = ?", req.AppId).
		Where("`app_oauth_channel`.`tenant_id` = ?", req.TenantId)
	db = db.Joins("Channel", ctx.DB().Where("Channel.status = true"))
	db = db.Where("Channel.status is not null")
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

// CreateAppOAuthChannel 创建数据
func (r App) CreateAppOAuthChannel(ctx core.Context, aoc *entity.AppOAuthChannel) (uint32, error) {
	return aoc.Id, ctx.DB().Create(aoc).Error
}

// DeleteAppOAuthChannel 删除数据
func (r App) DeleteAppOAuthChannel(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.AppOAuthChannel{}).Error
}

// ListAppField 创建数据
func (r App) ListAppField(ctx core.Context, req *types.ListAppFieldRequest) ([]*entity.AppField, uint32, error) {
	var (
		list  []*entity.AppField
		total int64
	)

	db := ctx.DB().Model(entity.AppField{}).Where("app_id = ?", req.AppId)

	join := ctx.DB().Where("Field.status = true")
	if req.Keyword != nil {
		join = join.Where("keyword LIKE ?", *req.Keyword+"%")
	}
	if req.Name != nil {
		join = join.Where("name LIKE ?", *req.Name+"%")
	}
	db = db.Joins("Field", join)
	db = db.Where("Field.status is not null")

	// 查询条件下数据总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = page.SearchScopes(db, &req.Search)
	if err := db.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, uint32(total), nil
}

// CreateAppField 创建数据
func (r App) CreateAppField(ctx core.Context, aoc *entity.AppField) (uint32, error) {
	return aoc.Id, ctx.DB().Create(aoc).Error
}

// DeleteAppField 删除数据
func (r App) DeleteAppField(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.AppField{}).Error
}
