package dictionary

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/biz/dictionary"
)

func NewRepo() dictionary.Repo {
	return &repo{}
}

type repo struct {
}

func (r *repo) PageDictionary(ctx kratosx.Context, in *dictionary.PageDictionaryRequest) ([]*dictionary.Dictionary, uint32, error) {
	var list []*dictionary.Dictionary
	var total int64

	db := ctx.DB().Model(dictionary.Dictionary{})
	if in.Keyword != nil {
		db.Where("keyword=?", *in.Keyword)
	}
	if in.Name != nil {
		db.Where("name like ?", *in.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (r *repo) AddDictionary(ctx kratosx.Context, in *dictionary.Dictionary) (uint32, error) {
	return in.ID, ctx.DB().Create(in).Error
}

func (r *repo) UpdateDictionary(ctx kratosx.Context, in *dictionary.Dictionary) error {
	return ctx.DB().Updates(in).Error
}

func (r *repo) DeleteDictionary(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(dictionary.Dictionary{}, "id=?", id).Error
}

func (r *repo) GetDictionaryValue(ctx kratosx.Context, keyword string) ([]*dictionary.DictionaryValue, error) {
	var dict dictionary.Dictionary
	if err := ctx.DB().Where("keyword=?", keyword).First(&dict).Error; err != nil {
		return nil, err
	}

	var list []*dictionary.DictionaryValue
	if err := ctx.DB().Where("dictionary_id=?", dict.ID).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (r *repo) PageDictionaryValue(ctx kratosx.Context, in *dictionary.PageDictionaryValueRequest) ([]*dictionary.DictionaryValue, uint32, error) {
	var list []*dictionary.DictionaryValue
	var total int64

	db := ctx.DB().Model(dictionary.DictionaryValue{}).Where("dictionary_id=?", in.DictionaryId)

	if in.Label != nil {
		db = db.Where("label like ?", "%"+*in.Label+"%")
	}
	if in.Value != nil {
		db = db.Where("value=?", *in.Value)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (r *repo) AddDictionaryValue(ctx kratosx.Context, in *dictionary.DictionaryValue) (uint32, error) {
	return in.ID, ctx.DB().Create(in).Error
}

func (r *repo) UpdateDictionaryValue(ctx kratosx.Context, in *dictionary.DictionaryValue) error {
	return ctx.DB().Updates(in).Error
}

func (r *repo) DeleteDictionaryValue(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(dictionary.DictionaryValue{}, "id=?", id).Error
}
