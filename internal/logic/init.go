package logic

import (
	"context"
	v1 "github.com/limes-cloud/manager/api/v1"
	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/internal/model"

	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"
)

type Init struct {
	conf *config.Config
}

func NewInit(conf *config.Config) *Init {
	return &Init{
		conf: conf,
	}
}

// Init 执行系统初始化
func (a *Init) Init() error {
	ctx := kratosx.MustContext(context.Background())
	if err := a.InitAuthentication(ctx); err != nil {
		return err
	}

	return nil
}

func (a *Init) InitAuthentication(ctx kratosx.Context) error {
	menu := model.Menu{}
	list, err := menu.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("type=?", consts.MENU_BASIC)
	})

	if err != nil {
		return v1.DatabaseError()
	}

	for _, item := range list {
		if item.Method != nil && item.Api != nil {
			ctx.Authentication().AddWhitelist(*item.Api, *item.Method)
		}
	}

	return nil
}
