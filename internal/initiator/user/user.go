package user

import (
	"log"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/internal/model"
)

func Init(ctx kratosx.Context, config *config.Config) {
	db := ctx.DB().Begin()

	if err := db.Create(user).Error; err != nil {
		log.Printf("init user error :%s\n", err.Error())
		db.Rollback()
		return
	}

	if err := db.Create(&model.UserRole{
		RoleID: consts.SuperAdmin,
		UserID: consts.SuperAdmin,
	}).Error; err != nil {
		log.Printf("init UserRole error :%s\n", err.Error())
		db.Rollback()
		return
	}

	if err := db.Create(&model.UserJob{
		JobID:  consts.SuperAdmin,
		UserID: consts.SuperAdmin,
	}).Error; err != nil {
		log.Printf("init UserJob error :%s\n", err.Error())
		db.Rollback()
		return
	}

	db.Commit()
}
