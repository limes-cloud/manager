package service

import (
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

const (
	StatusUntreated  = "untreated"
	StatusProcessing = "processing"
	StatusProcessed  = "processed"
)

type Feedback struct {
	repo repository.Feedback
}

func NewFeedback(
	repo repository.Feedback,
) *Feedback {
	return &Feedback{
		repo: repo,
	}
}

// ListFeedbackCategory 获取反馈建议分类列表
func (u *Feedback) ListFeedbackCategory(ctx core.Context, req *types.ListFeedbackCategoryRequest) ([]*entity.FeedbackCategory, uint32, error) {
	list, total, err := u.repo.ListFeedbackCategory(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list feedback category error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateFeedbackCategory 创建反馈建议分类
func (u *Feedback) CreateFeedbackCategory(ctx core.Context, req *entity.FeedbackCategory) (uint32, error) {
	id, err := u.repo.CreateFeedbackCategory(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create feedback category error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateFeedbackCategory 更新反馈建议分类
func (u *Feedback) UpdateFeedbackCategory(ctx core.Context, req *entity.FeedbackCategory) error {
	if err := u.repo.UpdateFeedbackCategory(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update feedback category error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteFeedbackCategory 删除反馈建议分类
func (u *Feedback) DeleteFeedbackCategory(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteFeedbackCategory(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete feedback category error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListFeedback 获取反馈建议列表
func (u *Feedback) ListFeedback(ctx core.Context, req *types.ListFeedbackRequest) ([]*entity.Feedback, uint32, error) {
	list, total, err := u.repo.ListFeedback(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list feedback error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}

	return list, total, nil
}

// CreateFeedback 创建反馈建议
func (u *Feedback) CreateFeedback(ctx core.Context, feedback *entity.Feedback) (uint32, error) {
	content := feedback.Title + feedback.Content + feedback.Device
	if feedback.Images != nil {
		content += *feedback.Images
	}
	if feedback.Contact != nil {
		content += *feedback.Contact
	}

	feedback.Status = StatusUntreated
	feedback.Md5 = crypto.MD5([]byte(content))

	if u.repo.IsExistFeedbackByMd5(ctx, feedback.Md5) {
		return 0, errors.ExistFeedbackError()
	}

	id, err := u.repo.CreateFeedback(ctx, feedback)
	if err != nil {
		ctx.Logger().Warnw("msg", "create feedback error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// DeleteFeedback 删除反馈建议
func (u *Feedback) DeleteFeedback(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteFeedback(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete feedback error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// UpdateFeedback 更新反馈建议
func (u *Feedback) UpdateFeedback(ctx core.Context, feedback *entity.Feedback) error {
	if feedback.Status != "" && value.InList([]string{
		StatusUntreated,
		StatusProcessing,
		StatusProcessed,
	}, feedback.Status) {
		return errors.ParamsError()
	}
	feedback.ProcessedBy = &ctx.Auth().UserId
	if err := u.repo.UpdateFeedback(ctx, feedback); err != nil {
		ctx.Logger().Warnw("msg", "update feedback error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}
