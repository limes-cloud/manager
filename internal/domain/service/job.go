package service

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/tree"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/pkg/md"
	"github.com/limes-cloud/manager/internal/types"
)

type Job struct {
	conf *conf.Config
	repo repository.Job
	role repository.Role
}

func NewJob(config *conf.Config, repo repository.Job) *Job {
	return &Job{conf: config, repo: repo}
}

// ListJob 获取部门信息列表树
func (u *Job) ListJob(ctx kratosx.Context, req *types.ListJobRequest) ([]*entity.Job, error) {
	// 获取部门列表
	list, err := u.repo.ListJob(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list department error", "err", err.Error())
		return nil, errors.ListError()
	}
	return tree.BuildArrayTree(list), nil
}

// ListCurrentJob 获取当前用户的部门信息列表树
func (u *Job) ListCurrentJob(ctx kratosx.Context, req *types.ListJobRequest) ([]*entity.Job, error) {
	all, scopes, err := u.repo.GetJobDataScope(ctx, md.UserId(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "list department error", "err", err.Error())
		return nil, errors.DatabaseError()
	}

	// 通过指定权限列表的部门
	if !all {
		req.Ids = scopes
	}
	return u.ListJob(ctx, req)
}

// CreateJob 创建职位信息
func (u *Job) CreateJob(ctx kratosx.Context, req *entity.Job) (uint32, error) {
	id, err := u.repo.CreateJob(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create job error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateJob 更新职位信息
func (u *Job) UpdateJob(ctx kratosx.Context, req *entity.Job) error {
	if err := u.repo.UpdateJob(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update job error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteJob 删除职位信息
func (u *Job) DeleteJob(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteJob(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete job error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
