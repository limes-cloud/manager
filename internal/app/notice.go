package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/api/notice"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type Notice struct {
	notice.UnimplementedNoticeServer
	srv *service.Notice
}

// NewNotice 初始化通知应用
func NewNotice() *Notice {
	return &Notice{
		srv: service.NewNotice(
			dbs.NewNotice(),
		),
	}
}

// init 应用注册
func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewNotice()
		notice.RegisterNoticeHTTPServer(hs, srv)
		notice.RegisterNoticeServer(gs, srv)
	})
}

// GetNotice 获取指定通知信息
func (app *Notice) GetNotice(c context.Context, req *notice.GetNoticeRequest) (*notice.GetNoticeReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	ent, err := app.srv.GetNotice(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := notice.GetNoticeReply{}
	if err := value.Transform(ent, &reply); err != nil {
		ctx.Logger().Errorw("msg", "get notice resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// GetVisibleNotice 获取指定通知信息
func (app *Notice) GetVisibleNotice(c context.Context, req *notice.GetNoticeRequest) (*notice.GetNoticeReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	ent, err := app.srv.GetVisibleNotice(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := notice.GetNoticeReply{}
	if err := value.Transform(ent, &reply); err != nil {
		ctx.Logger().Errorw("msg", "get notice resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListNotice 获取通知信息列表
func (app *Notice) ListNotice(c context.Context, req *notice.ListNoticeRequest) (*notice.ListNoticeReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListNoticeRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list notice req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := app.srv.ListNotice(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := notice.ListNoticeReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get notice resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListVisibleNotice 获取可见的通知信息列表
func (app *Notice) ListVisibleNotice(c context.Context, req *notice.ListVisibleNoticeRequest) (*notice.ListNoticeReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListNoticeRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list notice req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 设置默认参数
	in.Status = value.Bool(true)
	in.Order = value.String("desc")
	in.OrderBy = value.String("is_top")
	in.UserId = value.Pointer(ctx.Auth().UserId)

	// 调用服务
	list, total, err := app.srv.ListNotice(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := notice.ListNoticeReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get notice resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateNotice 创建通知信息
func (app *Notice) CreateNotice(c context.Context, req *notice.CreateNoticeRequest) (*notice.CreateNoticeReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Notice{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create notice req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateNotice(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &notice.CreateNoticeReply{Id: id}, nil
}

// UpdateNotice 更新通知信息
func (app *Notice) UpdateNotice(c context.Context, req *notice.UpdateNoticeRequest) (*notice.UpdateNoticeReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Notice{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create notice req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateNotice(ctx, &in); err != nil {
		return nil, err
	}

	return &notice.UpdateNoticeReply{}, nil
}

// DeleteNotice 删除通知信息
func (app *Notice) DeleteNotice(c context.Context, req *notice.DeleteNoticeRequest) (*notice.DeleteNoticeReply, error) {
	return &notice.DeleteNoticeReply{}, app.srv.DeleteNotice(core.MustContext(c), req.Id)
}

// ListNoticeClassify 获取通知信息列表
func (app *Notice) ListNoticeClassify(c context.Context, req *notice.ListNoticeClassifyRequest) (*notice.ListNoticeClassifyReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, err := app.srv.ListNoticeClassify(ctx)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := notice.ListNoticeClassifyReply{}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get notice classify resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateNoticeClassify 创建通知信息
func (app *Notice) CreateNoticeClassify(c context.Context, req *notice.CreateNoticeClassifyRequest) (*notice.CreateNoticeClassifyReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.NoticeClassify{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create notice classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateNoticeClassify(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &notice.CreateNoticeClassifyReply{Id: id}, nil
}

// UpdateNoticeClassify 更新通知信息
func (app *Notice) UpdateNoticeClassify(c context.Context, req *notice.UpdateNoticeClassifyRequest) (*notice.UpdateNoticeClassifyReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.NoticeClassify{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create notice classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateNoticeClassify(ctx, &in); err != nil {
		return nil, err
	}

	return &notice.UpdateNoticeClassifyReply{}, nil
}

// DeleteNoticeClassify 删除通知信息
func (app *Notice) DeleteNoticeClassify(c context.Context, req *notice.DeleteNoticeClassifyRequest) (*notice.DeleteNoticeClassifyReply, error) {
	return &notice.DeleteNoticeClassifyReply{}, app.srv.DeleteNoticeClassify(core.MustContext(c), req.Id)
}
