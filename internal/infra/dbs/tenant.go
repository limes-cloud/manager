package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/pkg/value"

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

	// 查询指定应用的租户
	if req.AppId != nil {
		db = db.Joins("LEFT JOIN tenant_app ON tenant_app.tenant_id = tenant.id and tenant_app.app_id = ?", req.AppId)
		db = db.Where("tenant_app.app_id is not null")
	}

	// 搜索排序
	if req.Search != nil {
		// 查询条件下数据总数
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}
		// 默认排序
		if req.Order == nil {
			req.Order = value.Pointer("desc")
		}
		// 排序字段
		if req.OrderBy == nil {
			req.Order = value.Pointer("weight")
		}
		// 分页查询
		db = page.SearchScopes(db, req.Search)
	}

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
