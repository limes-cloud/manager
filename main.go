package main

import (
	"context"

	"github.com/limes-cloud/manager/internal/middleware"

	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"

	"github.com/limes-cloud/kratosx/library"
	"github.com/limes-cloud/kratosx/library/db"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/app"
	"github.com/limes-cloud/manager/internal/core"
)

func main() {
	srv := core.InitApp(
		kratosx.WithRegistrarServer(app.Register),
		kratosx.WithValidateErrHook(func(ctx context.Context, err error) error {
			c := core.MustContext(ctx)
			c.Logger().Warnw("msg", "params validate error", "err", err)
			return errors.ParamsError()
		}),
		kratosx.WithLibraryOptions(
			library.WithDBOptions(db.WithHookScope(service.NewScope(dbs.NewScope(), dbs.NewUser()).Hook)),
		),
		kratosx.WithMiddleware(middleware.Middleware()...),
	)

	if err := srv.App().Run(); err != nil {
		log.Fatal(err)
	}
}
