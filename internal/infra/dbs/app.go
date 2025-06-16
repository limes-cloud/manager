package dbs

import (
	"fmt"
	"sync"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type App struct {
}

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
func (r App) GetAppByKeyword(ctx kratosx.Context, keyword string) (*entity.App, error) {
	var (
		app = entity.App{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs).
		Preload("Channels", "status=true").
		Preload("Fields", "status=true")
	return &app, db.Where("keyword = ?", keyword).First(&app).Error
}

// GetApp 获取指定的数据
func (r App) GetApp(ctx kratosx.Context, id uint32) (*entity.App, error) {
	var (
		app = entity.App{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs).
		Preload("Channels", "status=true").
		Preload("Fields", "status=true")
	return &app, db.Where("id = ?", id).First(&app).Error
}

// ListApp 获取列表
func (r App) ListApp(ctx kratosx.Context, req *types.ListAppRequest) ([]*entity.App, uint32, error) {
	var (
		list  []*entity.App
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.App{}).Select(fs)
	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Ids != nil {
		db = db.Where("id in ?", req.Ids)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	if err := db.Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, uint32(total), nil
}

// CreateApp 创建数据
func (r App) CreateApp(ctx kratosx.Context, app *entity.App) (uint32, error) {
	return app.Id, ctx.DB().Create(app).Error
}

// UpdateApp 更新数据
func (r App) UpdateApp(ctx kratosx.Context, app *entity.App) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if len(app.AppChannels) != 0 {
			if err := tx.Where("app_id=?", app.Id).Delete(entity.AppChannel{}).Error; err != nil {
				return err
			}
		}

		if len(app.AppFields) != 0 {
			if err := tx.Where("app_id=?", app.Id).Delete(entity.AppField{}).Error; err != nil {
				return err
			}
		}

		return tx.Updates(app).Error
	})
}

// UpdateAppStatus 更新数据状态
func (r App) UpdateAppStatus(ctx kratosx.Context, req *types.UpdateAppStatusRequest) error {
	return ctx.DB().Model(entity.App{}).
		Where("id=?", req.Id).
		Updates(map[string]any{
			"status":       req.Status,
			"disable_desc": req.DisableDesc,
		}).Error
}

// DeleteApp 删除数据
func (r App) DeleteApp(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.App{}).Error
}
