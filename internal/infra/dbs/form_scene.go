package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type FormScene struct{}

var (
	formSceneIns  *FormScene
	formSceneOnce sync.Once
)

func NewFormScene() *FormScene {
	formSceneOnce.Do(func() {
		formSceneIns = &FormScene{}
	})
	return formSceneIns
}

// GetFormScene 获取场景（by id）
func (r FormScene) GetFormScene(ctx core.Context, id uint32) (*entity.FormScene, error) {
	var fs entity.FormScene
	return &fs, ctx.DB().First(&fs, id).Error
}

// GetFormSceneByKeyword 获取场景（by keyword）
func (r FormScene) GetFormSceneByKeyword(ctx core.Context, keyword string) (*entity.FormScene, error) {
	var fs entity.FormScene
	return &fs, ctx.DB().Where("keyword = ?", keyword).First(&fs).Error
}

// ListFormScene 查询场景列表
func (r FormScene) ListFormScene(ctx core.Context, req *types.ListFormSceneRequest) ([]*entity.FormScene, uint32, error) {
	var (
		list  []*entity.FormScene
		total int64
	)

	db := ctx.DB().Model(entity.FormScene{})

	if req.Keyword != nil {
		db = db.Where("keyword LIKE ?", *req.Keyword+"%")
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	if req.Search != nil {
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}
		db = page.SearchScopes(db, req.Search)
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateFormScene 创建场景
func (r FormScene) CreateFormScene(ctx core.Context, ent *entity.FormScene) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateFormScene 更新场景
func (r FormScene) UpdateFormScene(ctx core.Context, ent *entity.FormScene) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteFormScene 删除场景
func (r FormScene) DeleteFormScene(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.FormScene{}).Error
}
