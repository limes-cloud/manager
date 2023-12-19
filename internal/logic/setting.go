package logic

import (
	v1 "github.com/limes-cloud/manager/api/v1"
	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/pkg/util"

	"github.com/limes-cloud/kratosx"
)

type Setting struct {
	conf *config.Config
}

func NewSetting(conf *config.Config) *Setting {
	return &Setting{
		conf: conf,
	}
}

// Get 获取配置
func (a *Setting) Get(ctx kratosx.Context) (*v1.GetSettingReply, error) {
	reply := v1.GetSettingReply{}
	if util.Transform(a.conf.Setting, &reply) != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}
