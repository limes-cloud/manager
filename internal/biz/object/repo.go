package object

import "github.com/limes-cloud/kratosx"

type Repo interface {
	GetObjectById(ctx kratosx.Context, id uint32) (*Object, error)
	GetObjectByKeyword(ctx kratosx.Context, keyword string) (*Object, error)
	PageObject(ctx kratosx.Context, req *PageObjectRequest) ([]*Object, uint32, error)
	AddObject(ctx kratosx.Context, c *Object) (uint32, error)
	UpdateObject(ctx kratosx.Context, c *Object) error
	DeleteObject(ctx kratosx.Context, uint322 uint32) error
}
