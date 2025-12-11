package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Feedback interface {
	// ListFeedbackClassify 获取反馈建议分类列表
	ListFeedbackClassify(ctx core.Context, req *types.ListFeedbackClassifyRequest) ([]*entity.FeedbackClassify, uint32, error)

	// CreateFeedbackClassify 创建反馈建议分类
	CreateFeedbackClassify(ctx core.Context, req *entity.FeedbackClassify) (uint32, error)

	// UpdateFeedbackClassify 更新反馈建议分类
	UpdateFeedbackClassify(ctx core.Context, req *entity.FeedbackClassify) error

	// DeleteFeedbackClassify 删除反馈建议分类
	DeleteFeedbackClassify(ctx core.Context, id uint32) error

	// IsExistFeedbackByMd5 是否存在反馈
	IsExistFeedbackByMd5(ctx core.Context, md5 string) bool

	// GetFeedback 获取反馈建议
	GetFeedback(ctx core.Context, id uint32) (*entity.Feedback, error)

	// ListFeedback 获取反馈建议列表
	ListFeedback(ctx core.Context, req *types.ListFeedbackRequest) ([]*entity.Feedback, uint32, error)

	// CreateFeedback 创建反馈建议
	CreateFeedback(ctx core.Context, req *entity.Feedback) (uint32, error)

	// DeleteFeedback 删除反馈建议
	DeleteFeedback(ctx core.Context, id uint32) error

	// UpdateFeedback 更新反馈建议
	UpdateFeedback(ctx core.Context, req *entity.Feedback) error
}
