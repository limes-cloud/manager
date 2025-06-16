package service

import (
	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/manager/api/manager/auth"
	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/pkg/md"
	"github.com/limes-cloud/manager/internal/types"
)

const (
	StatusUntreated  = "untreated"
	StatusProcessing = "processing"
	StatusProcessed  = "processed"
)

type Feedback struct {
	conf       *conf.Config
	repo       repository.Feedback
	permission repository.Permission
	file       repository.File
}

func NewFeedback(
	conf *conf.Config,
	repo repository.Feedback,
	permission repository.Permission,
	file repository.File,
) *Feedback {
	return &Feedback{
		conf:       conf,
		repo:       repo,
		permission: permission,
		file:       file,
	}
}

// ListFeedbackCategory 获取反馈建议分类列表
func (u *Feedback) ListFeedbackCategory(ctx kratosx.Context, req *types.ListFeedbackCategoryRequest) ([]*entity.FeedbackCategory, uint32, error) {
	list, total, err := u.repo.ListFeedbackCategory(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list feedback category error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateFeedbackCategory 创建反馈建议分类
func (u *Feedback) CreateFeedbackCategory(ctx kratosx.Context, req *entity.FeedbackCategory) (uint32, error) {
	id, err := u.repo.CreateFeedbackCategory(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create feedback category error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateFeedbackCategory 更新反馈建议分类
func (u *Feedback) UpdateFeedbackCategory(ctx kratosx.Context, req *entity.FeedbackCategory) error {
	if err := u.repo.UpdateFeedbackCategory(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update feedback category error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteFeedbackCategory 删除反馈建议分类
func (u *Feedback) DeleteFeedbackCategory(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteFeedbackCategory(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete feedback category error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListFeedback 获取反馈建议列表
func (u *Feedback) ListFeedback(ctx kratosx.Context, req *types.ListFeedbackRequest) ([]*entity.Feedback, uint32, error) {
	all, scopes, err := u.permission.GetResourceScopes(ctx, permissionApp)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app permission error", "err", err.Error())
		return nil, 0, errors.SystemError()
	}
	if !all {
		req.AppIds = scopes
	}

	list, total, err := u.repo.ListFeedback(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list feedback error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	for _, item := range list {
		if item.Images == nil {
			continue
		}
		var arr []string
		if json.Unmarshal([]byte(*item.Images), &arr) != nil {
			continue
		}
		for _, image := range arr {
			url := u.file.GetFileURL(ctx, image)
			if url != "" {
				item.ImageUrls = append(item.ImageUrls, url)
			}
		}
	}

	return list, total, nil
}

// CreateFeedback 创建反馈建议
func (u *Feedback) CreateFeedback(ctx kratosx.Context, feedback *entity.Feedback) (uint32, error) {
	content := feedback.Title + feedback.Content + feedback.Device
	if feedback.Images != nil {
		content += *feedback.Images
	}
	if feedback.Contact != nil {
		content += *feedback.Contact
	}
	feedback.UserId = md.UserId(ctx)
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
func (u *Feedback) DeleteFeedback(ctx kratosx.Context, id uint32) error {
	feedback, err := u.repo.GetFeedback(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get feedback error", "err", err.Error())
		return errors.DatabaseError()
	}
	if !u.HasApp(ctx, feedback.AppId) {
		return errors.NotAppScopeError()
	}

	if err := u.repo.DeleteFeedback(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete feedback error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// UpdateFeedback 更新反馈建议
func (u *Feedback) UpdateFeedback(ctx kratosx.Context, feedback *entity.Feedback) error {
	if feedback.Status != "" && valx.InList([]string{
		StatusUntreated,
		StatusProcessing,
		StatusProcessed,
	}, feedback.Status) {
		return errors.ParamsError()
	}

	if !u.HasApp(ctx, feedback.AppId) {
		return errors.NotAppScopeError()
	}

	adminInfo, err := auth.GetAuthInfo(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "get auth info error", "err", err.Error())
		return errors.SystemError()
	}
	feedback.ProcessedBy = &adminInfo.UserId
	if err := u.repo.UpdateFeedback(ctx, feedback); err != nil {
		ctx.Logger().Warnw("msg", "update feedback error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// HasApp 获取当前用户是否具有指定env的权限
func (u *Feedback) HasApp(ctx kratosx.Context, id uint32) bool {
	all, ids, err := u.permission.GetResourceScopes(ctx, permissionApp)
	if err != nil {
		return false
	}
	if all {
		return true
	}
	return valx.InList(ids, id)
}
