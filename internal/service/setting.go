package service

import (
	"context"

	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/manager/api/setting/v1"
	"github.com/limes-cloud/manager/internal/config"
)

type SettingService struct {
	v1.UnimplementedServiceServer

	conf *config.Config
}

func NewSettingService(conf *config.Config) *SettingService {
	return &SettingService{
		conf: conf,
	}
}

func (s SettingService) GetSetting(ctx context.Context, _ *emptypb.Empty) (*v1.GetSettingReply, error) {
	reply := v1.GetSettingReply{}
	_ = util.Transform(s.conf.Setting, &reply)
	return &reply, nil
}
