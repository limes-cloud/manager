package main

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/app"
	"github.com/limes-cloud/manager/internal/core"
	initentity "github.com/limes-cloud/manager/internal/initiator/entity"
	initmenu "github.com/limes-cloud/manager/internal/initiator/menu"
	"github.com/limes-cloud/manager/internal/middleware"
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
		// library.WithDBOptions(db.WithHookScope(service.NewScope(
		// 	dbs.NewScope(),
		// 	dbs.NewUserDept(),
		// ).Hook)),
		),
		kratosx.WithMiddleware(middleware.Middleware()...),
	)

	initentity.Register(core.MustContext(context.Background()))
	initmenu.Register(core.MustContext(context.Background()))

	if err := srv.App().Run(); err != nil {
		log.Fatal(err)
	}
}
