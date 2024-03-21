package object

import (
	"github.com/limes-cloud/kratosx"

	biz "github.com/limes-cloud/manager/internal/biz/object"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

func (u *repo) GetObjectById(ctx kratosx.Context, id uint32) (*biz.Object, error) {
	var object biz.Object
	return &object, ctx.DB().First(&object, "id=?", id).Error
}

func (u *repo) GetObjectByKeyword(ctx kratosx.Context, keyword string) (*biz.Object, error) {
	var object biz.Object
	return &object, ctx.DB().First(&object, "keyword=?", keyword).Error
}

func (u *repo) PageObject(ctx kratosx.Context, req *biz.PageObjectRequest) ([]*biz.Object, uint32, error) {
	var list []*biz.Object
	var total int64
	db := ctx.DB().Model(biz.Object{})
	if req.Keyword != nil {
		db.Where("keyword=?", *req.Keyword)
	}
	if req.Name != nil {
		db.Where("name like ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *repo) AddObject(ctx kratosx.Context, object *biz.Object) (uint32, error) {
	return object.ID, ctx.DB().Create(object).Error
}

func (u *repo) UpdateObject(ctx kratosx.Context, object *biz.Object) error {
	return ctx.DB().Updates(object).Error
}

func (u *repo) DeleteObject(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Object{}, "id=?", id).Error
}
