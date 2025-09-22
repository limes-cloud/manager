package core

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/manager/internal/domain/middleware"
)

// InitApp 初始化系统
func InitApp(opts ...kratosx.Option) *kratosx.App {
	defOpts := []kratosx.Option{
		// kratosx.WithConfigSource(configSource()),
		kratosx.WithConfigWatch(configScanWatch),
		kratosx.WithMiddleware(middleware.Middleware()...),
	}
	return kratosx.New(append(defOpts, opts...)...)
}
