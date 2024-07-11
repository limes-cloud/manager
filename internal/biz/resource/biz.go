package resource

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/pkg/md"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// GetResourceScopes 获取指定的资源权限
func (u *UseCase) GetResourceScopes(ctx kratosx.Context, userId uint32, keyword string) (bool, []uint32, error) {
	all, scopes, err := u.repo.GetResourceScopes(ctx, userId, keyword)
	if err != nil {
		return false, nil, errors.DatabaseError(err.Error())
	}
	return all, scopes, nil
}

// GetResource 获取指定的资源权限
func (u *UseCase) GetResource(ctx kratosx.Context, req *GetResourceRequest) ([]uint32, error) {
	req.UserId = md.UserId(ctx)
	ids, err := u.repo.GetResource(ctx, req)
	if err != nil {
		return nil, errors.DatabaseError(err.Error())
	}
	return ids, nil
}

// UpdateResource 更新资源权限
func (u *UseCase) UpdateResource(ctx kratosx.Context, req *UpdateResourceRequest) error {
	req.UserId = md.UserId(ctx)
	if err := u.repo.UpdateResource(ctx, req); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}
