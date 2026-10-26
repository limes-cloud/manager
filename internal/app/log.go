package app

import (
	"context"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/api/log"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
)

type Log struct {
	log.UnimplementedLogServer
	srv *service.Log
}

func NewLog() *Log {
	return &Log{
		srv: service.NewLog(
			dbs.NewLog(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewLog()
		log.RegisterLogHTTPServer(hs, srv)
		log.RegisterLogServer(gs, srv)
	})
}

// ListLoginLog 获取登陆日志
func (s *Log) ListLoginLog(c context.Context, req *log.ListLoginLogRequest) (*log.ListLoginLogReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = &types.ListLoginLogRequest{}
	)
	if err := value.Transform(req, in); err != nil {
		ctx.Logger().Errorw("msg", "list log resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 解析应用范围：显式 app 优先，否则回退到当前请求 app
	in.App = value.Pointer(ctx.AppKey(req.GetApp()))

	list, total, err := s.srv.ListLoginLog(ctx, in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := log.ListLoginLogReply{
		Total: total,
	}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list log resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// ListAuthLog 获取鉴权日志
func (s *Log) ListAuthLog(c context.Context, req *log.ListAuthLogRequest) (*log.ListAuthLogReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = &types.ListAuthLogRequest{}
	)
	if err := value.Transform(req, in); err != nil {
		ctx.Logger().Errorw("msg", "list log resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 解析应用范围：显式 app 优先，否则回退到当前请求 app
	in.App = value.Pointer(ctx.AppKey(req.GetApp()))

	list, total, err := s.srv.ListAuthLog(ctx, in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := log.ListAuthLogReply{
		Total: total,
	}
	for _, v := range list {
		item := &log.ListAuthLogReply_Log{
			CreatedAt: uint32(v.CreatedAt),
		}
		// 实体 App 为字符串（应用标识），AppInfo 才是应用对象，逐字段映射避免 JSON 键名冲突
		if v.AppInfo != nil {
			item.App = &log.ListAuthLogReply_Log_App{
				Id:      v.AppInfo.Id,
				Name:    v.AppInfo.Name,
				Keyword: v.AppInfo.Keyword,
				Logo:    v.AppInfo.Logo,
			}
		}
		if v.User != nil {
			item.User = &log.ListAuthLogReply_Log_User{
				Id:       v.User.Id,
				Username: v.User.Username,
				Nickname: value.Value(v.User.Nickname),
				Avatar:   value.Value(v.User.Avatar),
			}
		}
		if v.Menu != nil {
			item.Menu = &log.ListAuthLogReply_Log_Menu{
				Api:    value.Value(v.Menu.Api),
				Method: value.Value(v.Menu.Method),
				Title:  v.Menu.Title,
			}
		}
		reply.List = append(reply.List, item)
	}

	return &reply, nil
}
