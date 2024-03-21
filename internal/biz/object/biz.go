package object

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/config"
)

type UseCase struct {
	config *config.Config
	repo   Repo
}

func NewUseCase(config *config.Config, repo Repo) *UseCase {
	return &UseCase{config: config, repo: repo}
}

// GetObjectById 获取指定资源对象
func (u *UseCase) GetObjectById(ctx kratosx.Context, id uint32) (*Object, error) {
	object, err := u.repo.GetObjectById(ctx, id)
	if err != nil {
		return nil, errors.NotFound()
	}
	return object, nil
}

// GetObjectByKeyword 获取指定资源对象
func (u *UseCase) GetObjectByKeyword(ctx kratosx.Context, keyword string) (*Object, error) {
	object, err := u.repo.GetObjectByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.NotFound()
	}
	return object, nil
}

// PageObject 获取全部登录资源对象
func (u *UseCase) PageObject(ctx kratosx.Context, req *PageObjectRequest) ([]*Object, uint32, error) {
	object, total, err := u.repo.PageObject(ctx, req)
	if err != nil {
		return nil, 0, errors.NotFound()
	}
	return object, total, nil
}

// AddObject 添加登录资源对象信息
func (u *UseCase) AddObject(ctx kratosx.Context, object *Object) (uint32, error) {
	id, err := u.repo.AddObject(ctx, object)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

// UpdateObject 更新登录资源对象信息
func (u *UseCase) UpdateObject(ctx kratosx.Context, object *Object) error {
	if err := u.repo.UpdateObject(ctx, object); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// DeleteObject 删除登录资源对象信息
func (u *UseCase) DeleteObject(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteObject(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}
