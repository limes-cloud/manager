package dict

import (
	"log"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/config"
)

func Init(ctx kratosx.Context, config *config.Config) {
	db := ctx.DB().Begin()

	if err := db.Create(dict).Error; err != nil {
		db.Rollback()
		log.Printf("init dict error :%s\n", err.Error())
		return
	}

	if err := db.Create(dictValue).Error; err != nil {
		db.Rollback()
		log.Printf("init dictValue error :%s\n", err.Error())
		return
	}

	db.Commit()
}
