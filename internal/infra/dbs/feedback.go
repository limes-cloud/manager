package dbs

import (
	"fmt"
	"sync"

	"github.com/limes-cloud/manager/internal/core"

	"google.golang.org/protobuf/proto"

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

// ListFeedbackCategory 获取列表
func (r Feedback) ListFeedbackCategory(ctx core.Context, req *types.ListFeedbackCategoryRequest) ([]*entity.FeedbackCategory, uint32, error) {
	var (
		list  []*entity.FeedbackCategory
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.FeedbackCategory{}).Select(fs)

	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
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

	return list, uint32(total), db.Find(&list).Error
}

// CreateFeedbackCategory 创建数据
func (r Feedback) CreateFeedbackCategory(ctx core.Context, fc *entity.FeedbackCategory) (uint32, error) {
	return fc.Id, ctx.DB().Create(fc).Error
}

// GetFeedbackCategory 获取指定的数据
func (r Feedback) GetFeedbackCategory(ctx core.Context, id uint32) (*entity.FeedbackCategory, error) {
	var (
		fc = entity.FeedbackCategory{}
		fs = []string{"*"}
	)

	return &fc, ctx.DB().Select(fs).First(&fc, id).Error
}

// UpdateFeedbackCategory 更新数据
func (r Feedback) UpdateFeedbackCategory(ctx core.Context, fc *entity.FeedbackCategory) error {
	return ctx.DB().Updates(fc).Error
}

// DeleteFeedbackCategory 删除数据
func (r Feedback) DeleteFeedbackCategory(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.FeedbackCategory{}).Error
}

// GetFeedback 获取指定的数据
func (r Feedback) GetFeedback(ctx core.Context, id uint32) (*entity.Feedback, error) {
	var (
		feedback = entity.Feedback{}
		fs       = []string{"*"}
	)
	return &feedback, ctx.DB().Select(fs).First(&feedback, id).Error
}

// ListFeedback 获取列表
func (r Feedback) ListFeedback(ctx core.Context, req *types.ListFeedbackRequest) ([]*entity.Feedback, uint32, error) {
	var (
		list  []*entity.Feedback
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.Feedback{}).Select(fs)
	db = db.Preload("App").Preload("User").Preload("Category")

	if req.AppId != nil {
		db = db.Where("app_id = ?", *req.AppId)
	}
	if req.AppId == nil && req.AppIds != nil {
		db = db.Where("app_id in ?", req.AppIds)
	}
	if req.CategoryId != nil {
		db = db.Where("category_id = ?", *req.CategoryId)
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
