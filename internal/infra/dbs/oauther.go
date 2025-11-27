package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type OAuther struct{}

var (
	chIns  *OAuther
	chOnce sync.Once
)

func NewOAuther() *OAuther {
	chOnce.Do(func() {
		chIns = &OAuther{}
	})
	return chIns
}

// GetOAutherByKeyword 获取指定数据
func (r *OAuther) GetOAutherByKeyword(ctx core.Context, tid uint32, keyword string) (*entity.OAuther, error) {
	channel := entity.OAuther{}
	return &channel, ctx.DB().
		Where("tenant_id = ?", tid).
		Where("keyword = ?", keyword).
		First(&channel).Error
}

// GetOAuther 获取指定的数据
func (r *OAuther) GetOAuther(ctx core.Context, id uint32) (*entity.OAuther, error) {
	var (
		channel = entity.OAuther{}
		fs      = []string{"*"}
	)
	return &channel, ctx.DB().Select(fs).First(&channel, id).Error
}

// ListOAuther 获取列表
func (r *OAuther) ListOAuther(ctx core.Context, req *types.ListOAutherRequest) ([]*entity.OAuther, uint32, error) {
	var (
		list  []*entity.OAuther
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.OAuther{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword LIKE ?", *req.Keyword+"%")
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if req.Search != nil {
		// 查询条件下数据总数
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}

		// 分页排序
		db = page.SearchScopes(db, req.Search)
	}
	return list, uint32(total), db.Find(&list).Error
}

// CreateOAuther 创建数据
func (r *OAuther) CreateOAuther(ctx core.Context, channel *entity.OAuther) (uint32, error) {
	return channel.Id, ctx.DB().Create(channel).Error
}

// UpdateOAuther 更新数据
func (r *OAuther) UpdateOAuther(ctx core.Context, channel *entity.OAuther) error {
	return ctx.DB().Updates(channel).Error
}

// DeleteOAuther 删除数据
func (r *OAuther) DeleteOAuther(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.OAuther{}).Error
}
