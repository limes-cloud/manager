package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "github.com/limes-cloud/manager/api/manager/auth/v1"
	"github.com/limes-cloud/manager/internal/app"
)

func Authentication() middleware.Middleware {
	author := app.NewAuth()
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			// 跳过jwt白名单的也不检测
			_, is := kratosJwt.FromContext(ctx)
			if !is {
				return handler(ctx, req)
			}

			path, method := "", ""
			if tr, ok := transport.FromServerContext(ctx); ok {
				path = tr.Operation()
			}

			h, is := http.RequestFromServerContext(ctx)
			if is {
				path = h.URL.Path
				method = h.Method
			} else {
				method = "GRPC"
			}

			_, err := author.Auth(ctx, &v1.AuthRequest{
				Path:   path,
				Method: method,
			})

			if err != nil {
				return nil, err
			}

			return handler(ctx, req)
		}
	}
}
