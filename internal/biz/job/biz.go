package job

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/config"
)

type UseCase struct {
	repo Repo
	conf *config.Config
}

func NewUseCase(conf *config.Config, repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
		conf: conf,
	}
}

// GetJobById 获取指定资源对象
func (u *UseCase) GetJobById(ctx kratosx.Context, id uint32) (*Job, error) {
	object, err := u.repo.GetJobById(ctx, id)
	if err != nil {
		return nil, errors.NotFound()
	}
	return object, nil
}

// GetJobByKeyword 获取指定资源对象
func (u *UseCase) GetJobByKeyword(ctx kratosx.Context, keyword string) (*Job, error) {
	object, err := u.repo.GetJobByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.NotFound()
	}
	return object, nil
}

// PageJob 获取全部登录资源对象
func (u *UseCase) PageJob(ctx kratosx.Context, req *PageJobRequest) ([]*Job, uint32, error) {
	object, total, err := u.repo.PageJob(ctx, req)
	if err != nil {
		return nil, 0, errors.NotFound()
	}
	return object, total, nil
}

// AddJob 添加登录资源对象信息
func (u *UseCase) AddJob(ctx kratosx.Context, object *Job) (uint32, error) {
	id, err := u.repo.AddJob(ctx, object)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

// UpdateJob 更新登录资源对象信息
func (u *UseCase) UpdateJob(ctx kratosx.Context, object *Job) error {
	if err := u.repo.UpdateJob(ctx, object); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// DeleteJob 删除登录资源对象信息
func (u *UseCase) DeleteJob(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteJob(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}
