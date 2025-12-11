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

// ListFeedbackClassify 获取反馈建议分类列表
func (u *Feedback) ListFeedbackClassify(ctx core.Context, req *types.ListFeedbackClassifyRequest) ([]*entity.FeedbackClassify, uint32, error) {
	list, total, err := u.repo.ListFeedbackClassify(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list feedback classify error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateFeedbackClassify 创建反馈建议分类
func (u *Feedback) CreateFeedbackClassify(ctx core.Context, req *entity.FeedbackClassify) (uint32, error) {
	id, err := u.repo.CreateFeedbackClassify(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create feedback classify error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateFeedbackClassify 更新反馈建议分类
func (u *Feedback) UpdateFeedbackClassify(ctx core.Context, req *entity.FeedbackClassify) error {
	if err := u.repo.UpdateFeedbackClassify(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update feedback classify error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteFeedbackClassify 删除反馈建议分类
func (u *Feedback) DeleteFeedbackClassify(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteFeedbackClassify(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete feedback classify error", "err", err.Error())
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
	if feedback.Status != "" && !value.InList([]string{
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
