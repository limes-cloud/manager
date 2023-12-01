package main

import (
	v1 "manager/api/v1"
	srcConf "manager/config"
	"manager/internal/handler"
	"os"

	"github.com/limes-cloud/kratos/contrib/config/configure"

	"github.com/limes-cloud/kratos"
	"github.com/limes-cloud/kratos/config"
	"github.com/limes-cloud/kratos/log"
	"github.com/limes-cloud/kratos/middleware/tracing"
	"github.com/limes-cloud/kratos/transport/grpc"
	"github.com/limes-cloud/kratos/transport/http"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = os.Getenv("APP_NAME")

	id, _ = os.Hostname()
)

func main() {
	app := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Metadata(map[string]string{}),
		kratos.Config(configure.NewFromEnv()),
		kratos.RegistrarServer(RegisterServer),
		kratos.LoggerWith(kratos.LogField{
			"id":    id,
			"name":  Name,
			"trace": tracing.TraceID(),
			"span":  tracing.SpanID(),
		}),
	)

	if err := app.Run(); err != nil {
		log.Errorf("run service fail: %v", err)
	}
}

func RegisterServer(hs *http.Server, gs *grpc.Server, c config.Config) {
	conf := &srcConf.Config{}
	if err := c.ScanKey("business", conf); err != nil {
		panic("business config format error:" + err.Error())
	}

	srv := handler.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)
}
