package md

import (
	"manager/internal/model"

	"github.com/limes-cloud/kratos"
)

const (
	uid  = "user_id"
	rid  = "role_id"
	rkey = "role_keyword"
	did  = "department_id"
	dkey = "department_keyword"
)

//type Metadata struct {
//	UserID uint32 `json:"user_id"`
//	RoleID uint32 `json:"role_id"`
//}

func New(info *model.User) map[string]any {
	return map[string]any{
		uid:  info.ID,
		rid:  info.RoleID,
		rkey: info.Role.Keyword,
		did:  info.DepartmentID,
		dkey: info.Department.Keyword,
	}
}

func UserId(ctx kratos.Context) uint32 {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return 0
	}
	return uint32(m[uid].(float64))
}

func RoleId(ctx kratos.Context) uint32 {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return 0
	}
	return uint32(m[rid].(float64))
}

func RoleKeyword(ctx kratos.Context) string {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return ""
	}
	return m[rkey].(string)
}

func DepartmentId(ctx kratos.Context) uint32 {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return 0
	}
	return uint32(m[did].(float64))
}

func DepartmentKeyword(ctx kratos.Context) string {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return ""
	}
	return m[dkey].(string)
}
