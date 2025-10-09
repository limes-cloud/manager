package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/limes-cloud/manager/internal/middleware/api"
	"github.com/limes-cloud/manager/internal/middleware/auth"
)

func Middleware() []middleware.Middleware {
	return []middleware.Middleware{
		auth.Parse(),
		api.Check(),
	}
}
