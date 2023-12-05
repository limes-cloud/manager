package md

import (
	"manager/internal/model"

	"github.com/limes-cloud/kratosx"
)

const (
	uid  = "user_id"
	rid  = "role_id"
	rkey = "role_keyword"
	did  = "department_id"
	dkey = "department_keyword"
)

func New(info *model.User) map[string]any {
	return map[string]any{
		uid:  info.ID,
		rid:  info.RoleID,
		rkey: info.Role.Keyword,
		did:  info.DepartmentID,
		dkey: info.Department.Keyword,
	}
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
