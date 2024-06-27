package data

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"

	biz "github.com/limes-cloud/manager/internal/biz/dictionary"
	"github.com/limes-cloud/manager/internal/data/model"
)

type dictionaryRepo struct {
}

func NewDictionaryRepo() biz.Repo {
	return &dictionaryRepo{}
}

// ToDictionaryEntity model转entity
func (r dictionaryRepo) ToDictionaryEntity(m *model.Dictionary) *biz.Dictionary {
	e := &biz.Dictionary{}
	_ = valx.Transform(m, e)
	return e
}

// ToDictionaryModel entity转model
func (r dictionaryRepo) ToDictionaryModel(e *biz.Dictionary) *model.Dictionary {
	m := &model.Dictionary{}
	_ = valx.Transform(e, m)
	return m
}

// ListDictionary 获取列表
func (r dictionaryRepo) ListDictionary(ctx kratosx.Context, req *biz.ListDictionaryRequest) ([]*biz.Dictionary, uint32, error) {
	var (
		bs    []*biz.Dictionary
		ms    []*model.Dictionary
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Dictionary{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	if err := db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: *req.OrderBy},
		Desc:   *req.Order == "desc",
	}).Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToDictionaryEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateDictionary 创建数据
func (r dictionaryRepo) CreateDictionary(ctx kratosx.Context, req *biz.Dictionary) (uint32, error) {
	m := r.ToDictionaryModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// GetDictionary 获取指定的数据
func (r dictionaryRepo) GetDictionary(ctx kratosx.Context, id uint32) (*biz.Dictionary, error) {
	var (
		m  = model.Dictionary{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToDictionaryEntity(&m), nil
}

// UpdateDictionary 更新数据
func (r dictionaryRepo) UpdateDictionary(ctx kratosx.Context, req *biz.Dictionary) error {
	return ctx.DB().Updates(r.ToDictionaryModel(req)).Error
}

// DeleteDictionary 删除数据
func (r dictionaryRepo) DeleteDictionary(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.Dictionary{})
	return uint32(db.RowsAffected), db.Error
}

// ToDictionaryValueEntity model转entity
func (r dictionaryRepo) ToDictionaryValueEntity(m *model.DictionaryValue) *biz.DictionaryValue {
	e := &biz.DictionaryValue{}
	_ = valx.Transform(m, e)
	return e
}

// ToDictionaryValueModel entity转model
func (r dictionaryRepo) ToDictionaryValueModel(e *biz.DictionaryValue) *model.DictionaryValue {
	m := &model.DictionaryValue{}
	_ = valx.Transform(e, m)
	return m
}

// ListDictionaryValue 获取列表
func (r dictionaryRepo) ListDictionaryValue(ctx kratosx.Context, req *biz.ListDictionaryValueRequest) ([]*biz.DictionaryValue, uint32, error) {
	var (
		bs    []*biz.DictionaryValue
		ms    []*model.DictionaryValue
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.DictionaryValue{}).Select(fs)

	if req.DictionaryId != nil {
		db = db.Where("dictionary_id = ?", *req.DictionaryId)
	}
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

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToDictionaryValueEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateDictionaryValue 创建数据
func (r dictionaryRepo) CreateDictionaryValue(ctx kratosx.Context, req *biz.DictionaryValue) (uint32, error) {
	m := r.ToDictionaryValueModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// GetDictionaryValue 获取指定的数据
func (r dictionaryRepo) GetDictionaryValue(ctx kratosx.Context, id uint32) (*biz.DictionaryValue, error) {
	var (
		m  = model.DictionaryValue{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToDictionaryValueEntity(&m), nil
}

// UpdateDictionaryValue 更新数据
func (r dictionaryRepo) UpdateDictionaryValue(ctx kratosx.Context, req *biz.DictionaryValue) error {
	return ctx.DB().Updates(r.ToDictionaryValueModel(req)).Error
}

// UpdateDictionaryValueStatus 更新数据状态
func (r dictionaryRepo) UpdateDictionaryValueStatus(ctx kratosx.Context, id uint32, status bool) error {
	return ctx.DB().Model(model.DictionaryValue{}).Where("id=?", id).Update("status", status).Error
}

// DeleteDictionaryValue 删除数据
func (r dictionaryRepo) DeleteDictionaryValue(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.DictionaryValue{})
	return uint32(db.RowsAffected), db.Error
}

// GetDictionaryByKeyword 获取指定数据
func (r dictionaryRepo) GetDictionaryByKeyword(ctx kratosx.Context, keyword string) (*biz.Dictionary, error) {
	var (
		m  = model.Dictionary{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToDictionaryEntity(&m), nil
}

func (r dictionaryRepo) AllDictionaryValue(ctx kratosx.Context, keyword string) ([]*biz.DictionaryValue, error) {
	var (
		m  = model.Dictionary{}
		ms []*model.DictionaryValue
		bs []*biz.DictionaryValue
	)
	db := ctx.DB().Select("id")
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	if err := ctx.DB().Select("*").
		Where("status=true").
		Where("dictionary_id=?", m.Id).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToDictionaryValueEntity(m))
	}
	return bs, nil
}
