package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Dictionary struct{}

var (
	dictionaryIns  *Dictionary
	dictionaryOnce sync.Once
)

func NewDictionary() *Dictionary {
	dictionaryOnce.Do(func() {
		dictionaryIns = &Dictionary{}
	})
	return dictionaryIns
}

// GetDictionary 获取指定的数据
func (infra *Dictionary) GetDictionary(ctx kratosx.Context, id uint32) (*entity.Dictionary, error) {
	var dict entity.Dictionary
	return &dict, ctx.DB().First(&dict, id).Error
}

// GetDictionaryByKeyword 获取指定数据
func (infra *Dictionary) GetDictionaryByKeyword(ctx kratosx.Context, keyword string) (*entity.Dictionary, error) {
	var dict entity.Dictionary
	return &dict, ctx.DB().First(&dict, "keyword=?", keyword).Error
}

// ListDictionary 获取列表
func (infra *Dictionary) ListDictionary(ctx kratosx.Context, req *types.ListDictionaryRequest) ([]*entity.Dictionary, uint32, error) {
	var (
		total int64
		fs    = []string{"*"}
		ms    []*entity.Dictionary
	)

	db := ctx.DB().Model(entity.Dictionary{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	return ms, uint32(total), nil
}

// CreateDictionary 创建数据
func (infra *Dictionary) CreateDictionary(ctx kratosx.Context, dict *entity.Dictionary) (uint32, error) {
	return dict.Id, ctx.DB().Create(&dict).Error
}

// UpdateDictionary 更新数据
func (infra *Dictionary) UpdateDictionary(ctx kratosx.Context, dict *entity.Dictionary) error {
	return ctx.DB().Updates(dict).Error
}

// DeleteDictionary 删除数据
func (infra *Dictionary) DeleteDictionary(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Dictionary{}).Error
}

// ListDictionaryValue 获取列表
func (infra *Dictionary) ListDictionaryValue(ctx kratosx.Context, req *types.ListDictionaryValueRequest) ([]*entity.DictionaryValue, uint32, error) {
	var (
		ms    []*entity.DictionaryValue
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.DictionaryValue{}).Select(fs).Where("dictionary_id = ?", req.DictionaryId)

	if req.Label != nil {
		db = db.Where("label LIKE ?", *req.Label+"%")
	}
	if req.Value != nil {
		db = db.Where("value = ?", *req.Value)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = db.Order("weight desc,id asc")
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return ms, uint32(total), db.Find(&ms).Error
}

// CreateDictionaryValue 创建数据
func (infra *Dictionary) CreateDictionaryValue(ctx kratosx.Context, dictValue *entity.DictionaryValue) (uint32, error) {
	return dictValue.Id, ctx.DB().Create(dictValue).Error
}

// GetDictionaryValue 获取指定的数据
func (infra *Dictionary) GetDictionaryValue(ctx kratosx.Context, id uint32) (*entity.DictionaryValue, error) {
	var dictValue entity.DictionaryValue
	return &dictValue, ctx.DB().First(&dictValue, "id=?", id).Error
}

// UpdateDictionaryValue 更新数据
func (infra *Dictionary) UpdateDictionaryValue(ctx kratosx.Context, dictValue *entity.DictionaryValue) error {
	return ctx.DB().Updates(dictValue).Error
}

// UpdateDictionaryValueStatus 更新数据状态
func (infra *Dictionary) UpdateDictionaryValueStatus(ctx kratosx.Context, id uint32, status bool) error {
	// 获取字典是否为树，变更全部子节点状态
	return ctx.DB().Model(entity.DictionaryValue{}).Where("id = ?", id).Update("status", status).Error
}

// DeleteDictionaryValue 删除数据
func (infra *Dictionary) DeleteDictionaryValue(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.DictionaryValue{}).Error
}

func (infra *Dictionary) AllDictionaryValueByKeyword(ctx kratosx.Context, keyword string) ([]*entity.DictionaryValue, error) {
	var (
		m  = entity.Dictionary{}
		ms []*entity.DictionaryValue
	)
	db := ctx.DB().Model(entity.Dictionary{}).Select("id")
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	if err := ctx.DB().Model(entity.DictionaryValue{}).Select("*").
		Where("status=true").
		Where("dictionary_id=?", m.Id).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (infra *Dictionary) AllDictionaryValue(ctx kratosx.Context, req *types.AllDictionaryValueRequest) ([]*entity.DictionaryValue, error) {
	var ms []*entity.DictionaryValue
	db := ctx.DB().Model(entity.DictionaryValue{}).Select("*").Where("dictionary_id=?", req.DictionaryId)
	if req.Status != nil {
		db = db.Where("status=?", *req.Status)
	}
	if req.Label != nil {
		db = db.Where("label LIKE ?", *req.Label+"%")
	}
	if req.Value != nil {
		db = db.Where("value = ?", *req.Value)
	}
	if err := db.Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (infra *Dictionary) ListDictionaryValues(ctx kratosx.Context, keywords []string) ([]*types.DictionaryValue, error) {
	var ids []uint32
	if err := ctx.DB().Model(entity.Dictionary{}).
		Select("id").
		Where("keyword in ?", keywords).
		Find(&ids).Error; err != nil {
		return nil, err
	}

	var list []*types.DictionaryValue
	if len(ids) == 0 {
		return list, nil
	}
	if err := ctx.DB().Model(entity.DictionaryValue{}).
		Select([]string{"keyword", "label", "value", "type", "extra"}).
		Where("dictionary_id in ?", ids).
		Where("status = true").
		Preload("Dictionary").
		Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
