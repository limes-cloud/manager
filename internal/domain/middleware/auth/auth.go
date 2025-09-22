package auth

import (
	"context"

	"github.com/limes-cloud/manager/api/errors"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/limes-cloud/kratosx"
)

const (
	metaKey = "x-md-global-auth"
)

type Info struct {
	TenantId uint32 `json:"tenantId"`
	UserId   uint32 `json:"userId"`
	DeptId   uint32 `json:"deptId"`
	JobId    uint32 `json:"jobId"`
	Username string `json:"username"`
}

func (a Info) ToMap() map[string]any {
	return map[string]any{
		"tenantId": a.TenantId,
		"userId":   a.UserId,
		"deptId":   a.DeptId,
		"username": a.Username,
	}
}

type key struct{}

func Parse() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			kctx := kratosx.MustContext(ctx)
			token := kctx.JWT().GetToken(ctx)
			if token == "" {
				return handler(ctx, req)
			}

			// 解析token的信息
			info := &Info{}
			if err := kctx.JWT().ParseByToken(token, info); err != nil {
				return nil, errors.SystemError(err)
			}
			ctx = context.WithValue(ctx, key{}, info)
			return handler(ctx, req)
		}
	}
}

func GetAuth(ctx context.Context) *Info {
	val, ok := ctx.Value(key{}).(*Info)
	if !ok {
		return nil
	}
	return val
}

func HasLogin(ctx context.Context) bool {
	_, ok := ctx.Value(key{}).(*Info)
	return ok
}
