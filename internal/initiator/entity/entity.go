package entity

import (
	"github.com/limes-cloud/manager/internal/core"
	domainentity "github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

const managerAppKeyword = "manager"

func Register(ctx core.Context) {
	ctx.BeforeStart("register manager entity", func() {
		srv := service.NewEntity(dbs.NewEntity(), dbs.NewApp())
		list, err := srv.LoadEntity(ctx)
		if err != nil {
			ctx.Logger().Errorw("msg", "load manager entity error", "err", err.Error())
			return
		}

		reply, err := srv.RegisterEntity(ctx, &types.RegisterEntityRequest{
			AppKeyword: managerAppKeyword,
			Mode:       service.RegisterEntityModeAppend,
			Entities:   transEntities(list),
		})
		if err != nil {
			ctx.Logger().Errorw("msg", "register manager entity error", "err", err.Error())
			return
		}

		ctx.Logger().Infow(
			"msg", "register manager entity success",
			"created", reply.Created,
			"updated", reply.Updated,
			"deleted", reply.Deleted,
		)
	})
}

func transEntities(list []*domainentity.Entity) []*types.RegisterEntityNode {
	var res []*types.RegisterEntityNode
	for _, item := range list {
		if item == nil {
			continue
		}
		res = append(res, &types.RegisterEntityNode{
			Database: item.Database,
			Name:     item.Name,
			Comment:  item.Comment,
			Fields:   transFields(item.Fields),
		})
	}
	return res
}

func transFields(list []*domainentity.EntityField) []*types.RegisterEntityFieldNode {
	var res []*types.RegisterEntityFieldNode
	for _, item := range list {
		if item == nil {
			continue
		}
		res = append(res, &types.RegisterEntityFieldNode{
			Name:    item.Name,
			Comment: item.Comment,
		})
	}
	return res
}
