package logic

import (
	"github.com/limes-cloud/kratosx"

	v1 "github.com/limes-cloud/manager/api/v1"
	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/pkg/util"
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
	reply := v1.GetSettingReply{Id: 111}
	if util.Transform(a.conf.Setting, &reply) != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}
