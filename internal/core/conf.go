package core

import (
	"os"

	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/limes-cloud/configure/api/configure/client"
	"github.com/limes-cloud/kratosx/config"
)

var conf = &Conf{}

type Conf struct {
	Setting struct {
		Name      string
		Debug     bool
		Title     string
		Desc      string
		Copyright string
		Logo      string
		Watermark string
	}
	ChangePasswordType string

	// jwt 白名单
	JWT struct {
		Whitelist map[string]bool
	}
}

func configSource() kconfig.Source {
	host := os.Getenv("CONF_HOST")
	token := os.Getenv("CONF_TOKEN")
	name := os.Getenv("APP_NAME")
	if host != "" && token != "" && name != "" {
		return client.New(host, name, token)
	}
	return file.NewSource("conf/")
}

// configScanWatch 初始化
func configScanWatch(watch config.Watcher) {
	watch("business", func(value config.Value) {
		if err := value.Scan(&conf); err != nil {
			panic(err)
		}
	})
}
