package logic

import (
	"context"
	v1 "manager/api/v1"
	"manager/config"
	"manager/consts"
	"manager/internal/model"

	"gorm.io/gorm"

	"github.com/limes-cloud/kratos"
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
	ctx := kratos.MustContext(context.Background())
	if err := a.InitAuthentication(ctx); err != nil {
		return err
	}

	return nil
}

func (a *Init) InitAuthentication(ctx kratos.Context) error {
	menu := model.Menu{}
	list, err := menu.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("type=?", consts.MENU_BASIC)
	})

	if err != nil {
		return v1.ErrorDatabase()
	}

	for _, item := range list {
		if item.Method != nil && item.Api != nil {
			ctx.Authentication().AddWhitelist(*item.Api, *item.Method)
		}
	}

	return nil
}
