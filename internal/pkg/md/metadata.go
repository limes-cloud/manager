package md

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
)

const (
	uid  = "user_id"
	rid  = "role_id"
	rkey = "role_keyword"
	did  = "department_id"
	dkey = "department_keyword"
)

type Auth struct {
	UserId            uint32 `json:"user_id"`
	RoleId            uint32 `json:"role_id"`
	RoleKeyword       string `json:"role_keyword"`
	DepartmentId      uint32 `json:"department_id"`
	DepartmentKeyword string `json:"department_keyword"`
}

func New(info *Auth) map[string]any {
	var res map[string]any
	_ = util.Transform(info, &res)
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
