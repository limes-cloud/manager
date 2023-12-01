package handler

import (
	"context"
	v1 "manager/api/v1"

	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetSetting(ctx context.Context, in *emptypb.Empty) (*v1.GetSettingReply, error) {
	return s.setting.Get(kratos.MustContext(ctx))
}