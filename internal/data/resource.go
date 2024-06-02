package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	biz "github.com/limes-cloud/manager/internal/biz/resource"
	"github.com/limes-cloud/manager/internal/data/model"
	"gorm.io/gorm/clause"
)

type resourceRepo struct {
	departmentRepo
}

func NewResourceRepo() biz.Repo {
	return &resourceRepo{
		departmentRepo{},
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

func (r resourceRepo) UpdateResourceScopes(ctx kratosx.Context, userId uint32, req []*biz.Resource) error {
	var list []*model.Resource
	for _, item := range req {
		list = append(list, r.ToResourceModel(item))
	}

	all, scopes, err := r.GetDepartmentDataScope(ctx, userId)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if !all {
			if err := ctx.DB().Delete(model.Resource{}, "department_id in ?", scopes).Error; err != nil {
				return err
			}
		}
		return ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(list).Error
	})

}
