package api

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/manager/api/auth"
	"github.com/limes-cloud/manager/internal/app"
	midauth "github.com/limes-cloud/manager/internal/middleware/auth"
)

// Check api鉴权
func Check() middleware.Middleware {
	author := app.NewAuth()
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			// 不存在token,则跳过鉴权
			has := midauth.HasAuth(ctx)
			if !has {
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

			_, err := author.ApiAuth(ctx, &auth.ApiAuthRequest{
				Path:   path,
				Method: method,
			})
			if err != nil {
				return nil, err
			}

			_, err = author.ApiAuth(ctx, &auth.ApiAuthRequest{
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
