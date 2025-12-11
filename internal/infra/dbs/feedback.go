package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Feedback struct{}

var (
	feedbackIns  *Feedback
	feedbackOnce sync.Once
)

func NewFeedback() *Feedback {
	feedbackOnce.Do(func() {
		feedbackIns = &Feedback{}
	})
	return feedbackIns
}

// ListFeedbackClassify 获取列表
func (r Feedback) ListFeedbackClassify(ctx core.Context, req *types.ListFeedbackClassifyRequest) ([]*entity.FeedbackClassify, uint32, error) {
	var (
		list  []*entity.FeedbackClassify
		total int64
	)

	db := ctx.DB().Model(entity.FeedbackClassify{})

	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	db = page.SearchScopes(db, req.Search)

	return list, uint32(total), db.Find(&list).Error
}

// CreateFeedbackClassify 创建数据
func (r Feedback) CreateFeedbackClassify(ctx core.Context, fc *entity.FeedbackClassify) (uint32, error) {
	return fc.Id, ctx.DB().Create(fc).Error
}

// GetFeedbackClassify 获取指定的数据
func (r Feedback) GetFeedbackClassify(ctx core.Context, id uint32) (*entity.FeedbackClassify, error) {
	fc := entity.FeedbackClassify{}
	return &fc, ctx.DB().First(&fc, id).Error
}

// UpdateFeedbackClassify 更新数据
func (r Feedback) UpdateFeedbackClassify(ctx core.Context, fc *entity.FeedbackClassify) error {
	return ctx.DB().Updates(fc).Error
}

// DeleteFeedbackClassify 删除数据
func (r Feedback) DeleteFeedbackClassify(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.FeedbackClassify{}).Error
}

// GetFeedback 获取指定的数据
func (r Feedback) GetFeedback(ctx core.Context, id uint32) (*entity.Feedback, error) {
	feedback := entity.Feedback{}
	return &feedback, ctx.DB().First(&feedback, id).Error
}

// ListFeedback 获取列表
func (r Feedback) ListFeedback(ctx core.Context, req *types.ListFeedbackRequest) ([]*entity.Feedback, uint32, error) {
	var (
		list  []*entity.Feedback
		total int64
	)

	db := ctx.DB().Model(entity.Feedback{})
	db = db.Preload("App").Preload("User").Preload("Classify")

	if req.AppId != nil {
		db = db.Where("app_id = ?", *req.AppId)
	}
	if req.AppId == nil && req.AppIds != nil {
		db = db.Where("app_id in ?", req.AppIds)
	}
	if req.ClassifyId != nil {
		db = db.Where("category_id = ?", *req.ClassifyId)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Platform != nil {
		db = db.Where("platform = ?", *req.Platform)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = page.SearchScopes(db, req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateFeedback 创建数据
func (r Feedback) CreateFeedback(ctx core.Context, feedback *entity.Feedback) (uint32, error) {
	return feedback.Id, ctx.DB().Create(feedback).Error
}

// DeleteFeedback 删除数据
func (r Feedback) DeleteFeedback(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Feedback{}).Error
}

// UpdateFeedback 更新数据
func (r Feedback) UpdateFeedback(ctx core.Context, feedback *entity.Feedback) error {
	return ctx.DB().Updates(feedback).Error
}

func (r Feedback) IsExistFeedbackByMd5(ctx core.Context, md5 string) bool {
	var count int64
	ctx.DB().Model(entity.Feedback{}).Where("md5=?", md5).Count(&count)
	return count != 0
}
