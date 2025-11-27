package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/limes-cloud/manager/internal/middleware/jwt"
)

func Middleware() []middleware.Middleware {
	return []middleware.Middleware{
		// 获取token
		jwt.Token(),

		// jwt校验
		jwt.TokenValidate(),

		// api鉴权
		jwt.ApiValidate(),
	}
}
