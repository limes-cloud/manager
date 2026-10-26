package dbs

import (
	"sync"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type AppForm struct{}

var (
	appFormIns  *AppForm
	appFormOnce sync.Once
)

func NewAppForm() *AppForm {
	appFormOnce.Do(func() {
		appFormIns = &AppForm{}
	})
	return appFormIns
}

// ListAppForm 查询某应用下的表单关联列表
func (r AppForm) ListAppForm(ctx core.Context, req *types.ListAppFormRequest) ([]*entity.AppForm, error) {
	var list []*entity.AppForm
	err := ctx.DB().Model(entity.AppForm{}).
		Where("app_keyword = ?", req.AppKeyword).
		Find(&list).Error
	return list, err
}

// GetAppForm 获取单条关联（by app_keyword + scene）
func (r AppForm) GetAppForm(ctx core.Context, appKeyword string, scene string) (*entity.AppForm, error) {
	var af entity.AppForm
	err := ctx.DB().Where("app_keyword = ? AND scene = ?", appKeyword, scene).First(&af).Error
	return &af, err
}

// CreateAppForm 分配表单到应用
func (r AppForm) CreateAppForm(ctx core.Context, ent *entity.AppForm) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// DeleteAppForm 取消分配（by id）
func (r AppForm) DeleteAppForm(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.AppForm{}).Error
}
