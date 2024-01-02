package main

import (
	"log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/configure/client"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	_ "go.uber.org/automaxprocs"

	v1 "github.com/limes-cloud/manager/api/v1"
	srcConf "github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/internal/handler"
	"github.com/limes-cloud/manager/internal/initiator"
)

func main() {
	app := kratosx.New(
		kratosx.Config(client.NewFromEnv()),
		kratosx.RegistrarServer(RegisterServer),
	)

	if err := app.Run(); err != nil {
		log.Println("run service fail", err.Error())
	}
}

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	conf := &srcConf.Config{}
	// 配置初始化
	if err := c.Value("business").Scan(conf); err != nil {
		panic("business配置初始化失败：" + err.Error())
	}

	// 监听服务
	c.Watch("business", func(value config.Value) {
		if err := value.Scan(conf); err != nil {
			log.Printf("business配置变更失败：%s", err.Error())
		}
	})

	// 初始化逻辑
	ior := initiator.New(conf)
	if err := ior.Run(); err != nil {
		panic("initiator error:" + err.Error())
	}

	srv := handler.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)
}
