package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Resource struct {
}

var (
	resourceIns  *Resource
	resourceOnce sync.Once
)

func NewResource() *Resource {
	resourceOnce.Do(func() {
		resourceIns = &Resource{}
	})
	return resourceIns
}

func (infra *Resource) GetResourceScopes(ctx kratosx.Context, req *types.GetResourceScopesRequest) ([]uint32, error) {
	var ids []uint32
	db := ctx.DB().Model(entity.Resource{}).
		Select("resource_id").
		Where("keyword=?", req.Keyword)

	if req.DepartmentIds != nil {
		db = db.Where("department_id in ?", req.DepartmentIds)
	}

	return ids, db.Scan(&ids).Error
}

func (infra *Resource) DeleteResource(ctx kratosx.Context, req *types.DeleteResourceRequest) error {
	db := ctx.DB().Where("keyword=?", req.Keyword).Where("resource_id=?", req.ResourceId)
	if req.DepartmentIds != nil {
		db = db.Where("department_id in ?", req.DepartmentIds)
	}
	return db.Delete(entity.Resource{}).Error
}

func (infra *Resource) CreateResources(ctx kratosx.Context, list []*entity.Resource) error {
	return ctx.DB().Create(list).Error
}

func (infra *Resource) GetResource(ctx kratosx.Context, req *types.GetResourceRequest) ([]uint32, error) {
	var ids []uint32
	db := ctx.DB().Select("department_id").
		Model(entity.Resource{}).
		Where("keyword=? and resource_id=?", req.Keyword, req.ResourceId)
	if req.DepartmentIds != nil {
		db.Where("department_id in ?", req.DepartmentIds)
	}

	if err := db.Scan(&ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}
