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

	list, total, err := s.srv.ListAuthLog(ctx, in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := log.ListAuthLogReply{
		Total: total,
	}

	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list log resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}
