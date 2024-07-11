package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	biz "github.com/limes-cloud/manager/internal/biz/resource"
	"github.com/limes-cloud/manager/internal/data/model"
)

type resourceRepo struct {
	departmentRepo
}

func NewResourceRepo() biz.Repo {
	return &resourceRepo{
		departmentRepo: departmentRepo{},
	}
}

// ToResourceEntity model转entity
func (r resourceRepo) ToResourceEntity(m *model.Resource) *biz.Resource {
	e := &biz.Resource{}
	_ = valx.Transform(m, e)
	return e
}

// ToResourceModel entity转model
func (r resourceRepo) ToResourceModel(e *biz.Resource) *model.Resource {
	m := &model.Resource{}
	_ = valx.Transform(e, m)
	return m
}

func (r resourceRepo) GetResourceScopes(ctx kratosx.Context, userId uint32, keyword string) (bool, []uint32, error) {
	all, scopes, err := r.GetDepartmentDataScope(ctx, userId)
	if err != nil {
		return false, nil, err
	}
	if all {
		return true, nil, nil
	}
	var ids []uint32
	if err := ctx.DB().Select("resource_id").
		Model(model.Resource{}).
		Where("keyword=? and department_id in ?", keyword, scopes).
		Scan(&ids).Error; err != nil {
		return false, nil, err
	}
	return false, ids, nil
}

func (r resourceRepo) UpdateResource(ctx kratosx.Context, req *biz.UpdateResourceRequest) error {
	var list []*model.Resource
	all, scopes, err := r.GetDepartmentDataScope(ctx, req.UserId)
	if err != nil {
		return err
	}
	for _, id := range req.DepartmentIds {
		if all || valx.InList(scopes, id) {
			list = append(list, &model.Resource{
				Keyword:      req.Keyword,
				ResourceId:   req.ResourceId,
				DepartmentId: id,
			})
		}
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		db := ctx.DB().Where("keyword=?", req.Keyword).Where("resource_id=?", req.ResourceId)
		if !all {
			db = db.Where("department_id in ?", scopes)
		}
		if err := db.Delete(model.Resource{}).Error; err != nil {
			return err
		}
		return ctx.DB().Create(list).Error
	})
}

func (r resourceRepo) GetResource(ctx kratosx.Context, req *biz.GetResourceRequest) ([]uint32, error) {
	all, scopes, err := r.GetDepartmentDataScope(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var ids []uint32
	db := ctx.DB().Select("department_id").
		Model(model.Resource{}).
		Where("keyword=? and resource_id=?", req.Keyword, req.ResourceId)
	if !all {
		db.Where("department_id in ?", scopes)
	}

	if err = db.Scan(&ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}
