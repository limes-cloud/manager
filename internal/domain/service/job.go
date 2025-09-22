package service

import (
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Job struct {
	repo repository.Job
}

func NewJob(repo repository.Job) *Job {
	return &Job{repo: repo}
}

// ListJob 获取租户列表
func (u *Job) ListJob(ctx core.Context, req *types.ListJobRequest) ([]*entity.Job, uint32, error) {
	list, total, err := u.repo.ListJob(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list job error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateJob 创建租户
func (u *Job) CreateJob(ctx core.Context, req *entity.Job) (uint32, error) {
	req.Status = value.Bool(false)
	id, err := u.repo.CreateJob(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create job error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateJob 更新租户
func (u *Job) UpdateJob(ctx core.Context, req *entity.Job) error {
	if err := u.repo.UpdateJob(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update job error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// DeleteJob 删除租户
func (u *Job) DeleteJob(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteJob(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete job error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}
