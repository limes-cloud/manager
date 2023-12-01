package logic

import (
	v1 "manager/api/v1"
	"manager/config"
	"manager/pkg/util"

	"github.com/limes-cloud/kratos"
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
func (a *Setting) Get(ctx kratos.Context) (*v1.GetSettingReply, error) {
	reply := v1.GetSettingReply{}
	if util.Transform(a.conf.Setting, &reply) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}
