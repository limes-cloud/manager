package menu

import "github.com/limes-cloud/kratosx"

type Repo interface {
	InitBasicMenu(ctx kratosx.Context)
	AllRoleKeyword(ctx kratosx.Context, id uint32) ([]string, error)
	AllMenu(ctx kratosx.Context, in *AllMenuRequest) ([]*Menu, error)
	GetMenu(ctx kratosx.Context, id uint32) (*Menu, error)
	GetMenuByApi(ctx kratosx.Context, api, method string) (*Menu, error)
	AddMenu(ctx kratosx.Context, in *Menu) (uint32, error)
	UseHome(ctx kratosx.Context, app string, id uint32) error
	UpdateMenu(ctx kratosx.Context, in *Menu) error
	DeleteMenu(ctx kratosx.Context, ids []uint32) error
}
