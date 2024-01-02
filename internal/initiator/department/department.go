package department

import (
	"log"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/config"
)

func Init(ctx kratosx.Context, config *config.Config) {
	if err := department.Create(ctx); err != nil {
		log.Printf("init department error :%s\n", err.Error())
	}
}
