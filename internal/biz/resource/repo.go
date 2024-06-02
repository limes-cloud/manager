package resource

import "github.com/limes-cloud/kratosx"

type Repo interface {
	// GetResourceScopes 获取资源权限
	GetResourceScopes(ctx kratosx.Context, userId uint32, keyword string) (bool, []uint32, error)
	// UpdateResourceScopes 更新资源权限
	UpdateResourceScopes(ctx kratosx.Context, userId uint32, req []*Resource) error
}
