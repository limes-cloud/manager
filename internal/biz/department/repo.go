package department

import "github.com/limes-cloud/kratosx"

type Repo interface {
	GetDepartment(ctx kratosx.Context, id uint32) (*Department, error)
	AllManagerDepartment(ctx kratosx.Context, uid uint32) ([]*Department, error)
	AllManagerDepartmentIds(ctx kratosx.Context, uid uint32) ([]uint32, error)
	AllDepartment(ctx kratosx.Context, in *AllDepartmentRequest) ([]*Department, error)
	AddDepartment(ctx kratosx.Context, in *Department) (uint32, error)
	UpdateDepartment(ctx kratosx.Context, in *Department) error
	DeleteDepartment(ctx kratosx.Context, id uint32) error
	GetScope(ctx kratosx.Context, uid, oid uint32) ([]string, error)
	AllDepartmentObjectValue(ctx kratosx.Context, did, oid uint32) ([]string, error)
	ImportDepartmentObject(ctx kratosx.Context, in []*DepartmentObject) error
	AddDepartmentObject(ctx kratosx.Context, in *DepartmentObject) (uint32, error)
	DeleteDepartmentObjectValue(ctx kratosx.Context, oid uint32, value string) error
}
