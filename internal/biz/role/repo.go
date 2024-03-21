package role

import "github.com/limes-cloud/kratosx"

type Repo interface {
	GetRole(ctx kratosx.Context, id uint32) (*Role, error)
	AllRole(ctx kratosx.Context, in *AllRoleRequest) ([]*Role, error)
	AddRole(ctx kratosx.Context, in *Role) (uint32, error)
	UpdateRole(ctx kratosx.Context, in *Role) error
	DeleteRole(ctx kratosx.Context, id uint32) error
	ParentStatus(ctx kratosx.Context, id uint32) bool
	GetChildrenIds(ctx kratosx.Context, rid uint32) ([]uint32, error)
	GetParentIds(ctx kratosx.Context, rid uint32) ([]uint32, error)
	UpdateRoleMenus(ctx kratosx.Context, rid uint32, menuIds []uint32) error
	GetRoleMenuIds(ctx kratosx.Context, rid uint32) ([]uint32, error)
	AllMenuApiByIds(ctx kratosx.Context, ids []uint32) ([]*MenuApi, error)
}
