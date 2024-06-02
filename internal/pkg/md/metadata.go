package md

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
)

const (
	uid  = "userId"
	rid  = "roleId"
	rkey = "roleKeyword"
	did  = "departmentId"
	dkey = "departmentKeyword"
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
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return 0
	}
	return uint32(m[uid].(float64))
}

func RoleId(ctx kratosx.Context) uint32 {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return 0
	}
	return uint32(m[rid].(float64))
}

func RoleKeyword(ctx kratosx.Context) string {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return ""
	}
	return m[rkey].(string)
}

func DepartmentId(ctx kratosx.Context) uint32 {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return 0
	}
	return uint32(m[did].(float64))
}

func DepartmentKeyword(ctx kratosx.Context) string {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return ""
	}
	return m[dkey].(string)
}
