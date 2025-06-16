package repository

import "github.com/limes-cloud/kratosx"

type Permission interface {
	GetResourceScopes(ctx kratosx.Context, keyword string) (bool, []uint32, error)
}
