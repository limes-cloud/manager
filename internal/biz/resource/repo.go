package resource

import "github.com/limes-cloud/kratosx"

type Repo interface {
	// GetResourceScopes 获取指定用户的资源权限
	GetResourceScopes(ctx kratosx.Context, uid uint32, keyword string) (bool, []uint32, error)

	// GetResource 获取资源权限
	GetResource(ctx kratosx.Context, req *GetResourceRequest) ([]uint32, error)

	// UpdateResource 更新资源权限
	UpdateResource(ctx kratosx.Context, req *UpdateResourceRequest) error
}
