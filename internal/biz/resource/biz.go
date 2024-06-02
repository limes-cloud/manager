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

func NewUseCase(config *conf.Config) *UseCase {
	return &UseCase{conf: config}
}

// GetResourceScopes 获取指定用户的资源权限
func (u *UseCase) GetResourceScopes(ctx kratosx.Context, keyword string) (bool, []uint32, error) {
	all, scopes, err := u.repo.GetResourceScopes(ctx, md.UserId(ctx), keyword)
	if err != nil {
		return false, nil, errors.DatabaseError(err.Error())
	}
	return all, scopes, nil
}

// UpdateResourceScopes 更新指定用户的资源权限
func (u *UseCase) UpdateResourceScopes(ctx kratosx.Context, req *UpdateResourceScopesRequest) error {
	var list []*Resource
	for _, scope := range req.Scopes {
		list = append(list, &Resource{
			Keyword:      req.Keyword,
			ResourceId:   req.ResourceId,
			DepartmentId: scope,
		})
	}
	if err := u.repo.UpdateResourceScopes(ctx, md.UserId(ctx), list); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}
