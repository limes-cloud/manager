package service

import (
	"context"

	"github.com/limes-cloud/manager/internal/data"

	"github.com/limes-cloud/kratosx"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "github.com/limes-cloud/manager/api/manager/system/v1"
	"github.com/limes-cloud/manager/internal/biz/system"
	"github.com/limes-cloud/manager/internal/conf"
)

type SystemService struct {
	pb.UnimplementedSystemServer
	uc *system.UseCase
}

func NewSystemService(conf *conf.Config) *SystemService {
	return &SystemService{
		uc: system.NewUseCase(conf, data.NewSystemRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewSystemService(c)
		pb.RegisterSystemHTTPServer(hs, srv)
		pb.RegisterSystemServer(gs, srv)
	})
}

// GetSystemSetting 获取系统设置
func (s *SystemService) GetSystemSetting(c context.Context, req *pb.GetSystemSettingRequest) (*pb.GetSystemSettingReply, error) {
	setting := s.uc.GetSystemSetting(kratosx.MustContext(c), &system.GetSystemSettingRequest{
		DictionaryKeywords: req.DictionaryKeywords,
	})

	reply := pb.GetSystemSettingReply{
		Debug:              setting.Debug,
		Title:              setting.Title,
		Desc:               setting.Desc,
		Copyright:          setting.Copyright,
		Logo:               setting.Logo,
		Watermark:          setting.Watermark,
		ChangePasswordType: setting.ChangePasswordType,
	}
	if len(setting.Dictionaries) != 0 {
		dictArr := make(map[string]*pb.GetSystemSettingReply_DictionaryValueList)
		for _, item := range setting.Dictionaries {
			if dictArr[item.Keyword] == nil {
				dictArr[item.Keyword] = &pb.GetSystemSettingReply_DictionaryValueList{}
			}
			dv := &pb.DictionaryValue{
				Label: item.Label,
				Value: item.Value,
				Type:  item.Type,
				Extra: item.Extra,
			}
			dictArr[item.Keyword].List = append(dictArr[item.Keyword].List, dv)
		}
		reply.Dictionaries = dictArr
	}
	return &reply, nil
}
