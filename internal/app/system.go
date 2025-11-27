package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/api/system"

	"github.com/limes-cloud/manager/internal/domain/service"
)

type System struct {
	system.UnimplementedSystemServer
	srv *service.System
}

func NewSystem() *System {
	return &System{
		srv: service.NewSystem(),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewSystem()
		system.RegisterSystemHTTPServer(hs, srv)
		system.RegisterSystemServer(gs, srv)
	})
}

// GetSystemSetting 获取系统设置
func (s *System) GetSystemSetting(c context.Context, _ *system.GetSystemSettingRequest) (*system.GetSystemSettingReply, error) {
	setting := s.srv.GetSystemSetting(core.MustContext(c))

	reply := system.GetSystemSettingReply{
		Title:       setting.Title,
		Description: setting.Description,
		Copyright:   setting.Copyright,
		Logo:        setting.Logo,
		Watermark:   setting.Watermark,
	}
	//if len(setting.Dictionaries) != 0 {
	//	dictArr := make(map[string]*system.GetSystemSettingReply_DictionaryValueList)
	//	for _, item := range setting.Dictionaries {
	//		if dictArr[item.Keyword] == nil {
	//			dictArr[item.Keyword] = &system.GetSystemSettingReply_DictionaryValueList{}
	//		}
	//		dv := &system.DictionaryValue{
	//			Label: item.Label,
	//			Value: item.Value,
	//			Type:  item.Type,
	//			Extra: item.Extra,
	//		}
	//		dictArr[item.Keyword].List = append(dictArr[item.Keyword].List, dv)
	//	}
	//	reply.Dictionaries = dictArr
	//}
	return &reply, nil
}
