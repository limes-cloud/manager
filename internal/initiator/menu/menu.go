package menu

import (
	"log"

	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/internal/model"
)

func initFunc(db *gorm.DB, menu *model.Menu) {
	if err := db.Model(menu).Create(menu).Error; err != nil {
		log.Printf("init menu error :%v\n", err.Error())
		db.Rollback()
		return
	}
	pid := menu.ID()
	for _, cm := range menu.Children {
		cm.ParentID = pid
		initFunc(db, cm)
	}
}

func Init(ctx kratosx.Context, config *config.Config) {
	db := ctx.DB().Begin()
	for _, menu := range menus {
		initFunc(db, menu)
	}

	db.Commit()
}
