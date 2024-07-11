package md

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
)

type Auth struct {
	UserId            uint32 `json:"userId"`
	RoleId            uint32 `json:"roleId"`
	RoleKeyword       string `json:"roleKeyword"`
	DepartmentId      uint32 `json:"departmentId"`
	DepartmentKeyword string `json:"departmentKeyword"`
}

func New(info *Auth) map[string]any {
	var res map[string]any
	_ = valx.Transform(info, &res)
	return res
}

func Get(ctx kratosx.Context) (*Auth, error) {
	var auth Auth
	return &auth, ctx.JWT().Parse(ctx, &auth)
}

func UserId(ctx kratosx.Context) uint32 {
	m, err := Get(ctx)
	if err != nil {
		return 0
	}
	return m.UserId
}

func RoleId(ctx kratosx.Context) uint32 {
	m, err := Get(ctx)
	if err != nil {
		return 0
	}
	return m.RoleId
}

func RoleKeyword(ctx kratosx.Context) string {
	m, err := Get(ctx)
	if err != nil {
		return ""
	}
	return m.RoleKeyword
}

func DepartmentId(ctx kratosx.Context) uint32 {
	m, err := Get(ctx)
	if err != nil {
		return 0
	}
	return m.DepartmentId
}

func DepartmentKeyword(ctx kratosx.Context) string {
	m, err := Get(ctx)
	if err != nil {
		return ""
	}
	return m.DepartmentKeyword
}
