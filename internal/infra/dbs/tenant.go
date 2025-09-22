package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Tenant struct{}

var (
	tenantIns  *Tenant
	tenantOnce sync.Once
)

func NewTenant() *Tenant {
	tenantOnce.Do(func() {
		tenantIns = &Tenant{}
	})
	return tenantIns
}

// ListTenant 获取列表
func (r *Tenant) ListTenant(ctx core.Context, req *types.ListTenantRequest) ([]*entity.Tenant, uint32, error) {
	var (
		list  []*entity.Tenant
		total int64
	)

	db := ctx.DB().Model(entity.Tenant{})

	// 查询名称
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	if req.Keyword != nil {
		db = db.Where("keyword LIKE ?", *req.Keyword+"%")
	}

	// 查询状态
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	// 查询条件下数据总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 搜索排序
	db = page.SearchScopes(db, &req.Search)

	return list, uint32(total), db.Find(&list).Error
}

// CreateTenant 创建数据
func (r *Tenant) CreateTenant(ctx core.Context, et *entity.Tenant) (uint32, error) {
	return et.Id, ctx.DB().Create(et).Error
}

// GetTenant 获取指定的数据
func (r *Tenant) GetTenant(ctx core.Context, id uint32) (*entity.Tenant, error) {
	et := entity.Tenant{}
	return &et, ctx.DB().First(&et, id).Error
}

func (r *Tenant) GetTenantByKeyword(ctx core.Context, keyword string) (*entity.Tenant, error) {
	et := entity.Tenant{}
	return &et, ctx.DB().First(&et, "keyword=?", keyword).Error
}

// UpdateTenant 更新数据
func (r *Tenant) UpdateTenant(ctx core.Context, et *entity.Tenant) error {
	return ctx.DB().Updates(et).Error
}

// DeleteTenant 删除数据
func (r *Tenant) DeleteTenant(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Tenant{}, id).Error
}

// ExistTenant 是否存在
func (r *Tenant) ExistTenant(ctx core.Context, keyword string) (bool, error) {
	et := entity.Tenant{}
	return et.Id > 0, ctx.DB().Select("id").Where("keyword = ?", keyword).First(&et).Error
}

// ListTenantApp 获取app列表
func (r *Tenant) ListTenantApp(ctx core.Context, req *types.ListTenantAppRequest) ([]*entity.TenantApp, uint32, error) {
	var (
		list  []*entity.TenantApp
		total int64
	)

	db := ctx.DB().Model(entity.TenantApp{})

	if req.AppName != nil {
		joinWhere := ctx.DB().Where("App.name LIKE ?", *req.AppName+"%")
		db = db.Joins("App", joinWhere)
	} else {
		db = db.Joins("App")
	}

	// 查询条件下数据总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 搜索排序
	db = page.SearchScopes(db, &req.Search)

	return list, uint32(total), db.Find(&list).Error
}

func (r *Tenant) CreateTenantApp(ctx core.Context, ent *entity.TenantApp) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

func (r *Tenant) UpdateTenantApp(ctx core.Context, ent *entity.TenantApp) error {
	return ctx.DB().Updates(ent).Error
}

func (r *Tenant) DeleteTenantApp(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.TenantApp{}, id).Error
}
