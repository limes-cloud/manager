package job

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// ListJob 获取职位信息列表
func (u *UseCase) ListJob(ctx kratosx.Context, req *ListJobRequest) ([]*Job, uint32, error) {
	list, total, err := u.repo.ListJob(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateJob 创建职位信息
func (u *UseCase) CreateJob(ctx kratosx.Context, req *Job) (uint32, error) {
	id, err := u.repo.CreateJob(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateJob 更新职位信息
func (u *UseCase) UpdateJob(ctx kratosx.Context, req *Job) error {
	if err := u.repo.UpdateJob(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteJob 删除职位信息
func (u *UseCase) DeleteJob(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteJob(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// GetJob 获取指定的职位信息
func (u *UseCase) GetJob(ctx kratosx.Context, req *GetJobRequest) (*Job, error) {
	var (
		res *Job
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetJob(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetJobByKeyword(ctx, *req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}
