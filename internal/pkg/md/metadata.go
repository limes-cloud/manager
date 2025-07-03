package md

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/manager/api/manager/errors"
)

type Auth struct {
	DepartmentIds []uint32 `json:"departmentIds"`
	JobIds        []uint32 `json:"jobIds"`
	RoleIds       []uint32 `json:"roleIds"`
	UserId        uint32   `json:"userId"`
	UserName      string   `json:"userName"`
}

func New(info *Auth) map[string]any {
	var res map[string]any
	_ = valx.Transform(info, &res)
	return res
}

func Get(ctx kratosx.Context) *Auth {
	var (
		data Auth
		err  error
	)
	if ctx.Token() != "" {
		err = ctx.JWT().Parse(ctx, &data)
	} else {
		// 三方服务调用的时候通过auth信息获取
		err = ctx.Authentication().ParseAuth(ctx, &data)
	}
	if err != nil {
		panic(errors.ForbiddenError())
	}

	if data.UserId == 0 {
		panic(errors.ForbiddenError())
	}
	return &data
}

func UserId(ctx kratosx.Context) uint32 {
	return Get(ctx).UserId
}

func RoleIds(ctx kratosx.Context) []uint32 {
	return Get(ctx).RoleIds
}

func DepartmentIds(ctx kratosx.Context) []uint32 {
	return Get(ctx).DepartmentIds
}

func JobIds(ctx kratosx.Context) []uint32 {
	return Get(ctx).JobIds
}
