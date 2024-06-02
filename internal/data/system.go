package data

import (
	"github.com/limes-cloud/kratosx"
	biz "github.com/limes-cloud/manager/internal/biz/system"
	"github.com/limes-cloud/manager/internal/data/model"
)

type systemRepo struct {
}

func NewSystemRepo() biz.Repo {
	return &systemRepo{}
}

func (s systemRepo) GetDictionaryValues(ctx kratosx.Context, keywords []string) ([]*biz.DictionaryValue, error) {
	var list []*biz.DictionaryValue
	if err := ctx.DB().Model(model.DictionaryValue{}).
		Select([]string{"keyword", "label", "value", "type", "extra"}).
		Where("keyword in ?", keywords).
		Where("status = true").
		Joins("Dictionary").
		Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
