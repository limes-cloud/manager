package job

import (
	"log"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/config"
)

func Init(ctx kratosx.Context, config *config.Config) {
	if err := job.Create(ctx); err != nil {
		log.Printf("init job error :%s\n", err.Error())
	}
}
