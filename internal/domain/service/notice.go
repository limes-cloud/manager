package service

import (
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Notice struct {
	repo  repository.Notice
	scope repository.Scope
}

func NewNotice(
	repo repository.Notice,
) *Notice {
	return &Notice{
		repo: repo,
	}
}

// ListNoticeClassify 获取通知列表
func (u *Notice) ListNoticeClassify(ctx core.Context) ([]*entity.NoticeClassify, error) {
	list, err := u.repo.ListNoticeClassify(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "list dept classify error", "err", err.Error())
		return nil, errors.ListError(err.Error())
	}
	return list, nil
}

// CreateNoticeClassify 创建通知
func (u *Notice) CreateNoticeClassify(ctx core.Context, req *entity.NoticeClassify) (uint32, error) {
	id, err := u.repo.CreateNoticeClassify(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create dept classify error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateNoticeClassify 更新通知
func (u *Notice) UpdateNoticeClassify(ctx core.Context, req *entity.NoticeClassify) error {
	if err := u.repo.UpdateNoticeClassify(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update dept classify error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteNoticeClassify 删除通知
func (u *Notice) DeleteNoticeClassify(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteNoticeClassify(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete dept classify error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// GetNotice 获取通知列表
func (u *Notice) GetNotice(ctx core.Context, id uint32) (*entity.Notice, error) {
	res, err := u.repo.GetNotice(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// GetVisibleNotice 获取通知列表
func (u *Notice) GetVisibleNotice(ctx core.Context, id uint32) (*entity.Notice, error) {
	res, err := u.repo.GetNotice(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	if !value.Value(res.Status) {
		return nil, errors.GetError("通知不可见")
	}

	// 标记当前内容为已读
	if err := u.repo.AddUserNotice(ctx, ctx.Auth().UserId, id); err != nil {
		ctx.Logger().Warnw("msg", "update dept error", "err", err.Error())
		return nil, errors.UpdateError(err.Error())
	}

	return res, nil
}

// ListNotice 获取通知列表
func (u *Notice) ListNotice(ctx core.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error) {
	list, total, err := u.repo.ListNotice(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list dept error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}

	return list, total, nil
}

// CreateNotice 创建通知
func (u *Notice) CreateNotice(ctx core.Context, req *entity.Notice) (uint32, error) {
	req.Status = value.Pointer(false)
	id, err := u.repo.CreateNotice(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create dept error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateNotice 更新通知
func (u *Notice) UpdateNotice(ctx core.Context, req *entity.Notice) error {
	if err := u.repo.UpdateNotice(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update dept error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteNotice 删除通知
func (u *Notice) DeleteNotice(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteNotice(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete dept error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
