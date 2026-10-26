package menu

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

func Register(ctx core.Context) {
	ctx.BeforeStart("register manager menu", func() {
		srv := service.NewMenu(
			dbs.NewMenu(),
			dbs.NewApp(),
			dbs.NewRoleMenu(),
			dbs.NewScope(),
		)
		reply, err := srv.RegisterMenu(ctx, &types.RegisterMenuRequest{
			AppKeyword: managerAppKeyword,
			Mode:       service.RegisterMenuModeSync,
			Menus:      managerMenus,
		})
		if err != nil {
			ctx.Logger().Errorw("msg", "register manager menu error", "err", err.Error())
			return
		}
		ctx.Logger().Infow(
			"msg", "register manager menu success",
			"created", reply.Created,
			"updated", reply.Updated,
			"deleted", reply.Deleted,
		)
	})
}
