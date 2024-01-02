package auth

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/internal/model"
)

func Init(ctx kratosx.Context, config *config.Config) {
	menu := model.Menu{}
	list, _ := menu.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("type=?", consts.MenuBasic)
	})

	for _, item := range list {
		if item.Method != nil && item.Api != nil {
			ctx.Authentication().AddWhitelist(*item.Api, *item.Method)
		}
	}
}
