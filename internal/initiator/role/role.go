package role

import (
	"log"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/config"
)

func Init(ctx kratosx.Context, config *config.Config) {
	if err := role.Create(ctx); err != nil {
		log.Printf("init role error :%s\n", err.Error())
	}
}
