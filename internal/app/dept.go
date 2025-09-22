package app

import (
	"context"

	"github.com/limes-cloud/manager/internal/domain/entity"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/manager/api/dept"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
)

type Dept struct {
	dept.UnimplementedDeptServer
	srv *service.Dept
}

// NewDept 初始化租户应用
func NewDept() *Dept {
	return &Dept{
		srv: service.NewDept(
			dbs.NewDept(),
			dbs.NewScope(),
		),
	}
}

// init 应用注册
func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewDept()
		dept.RegisterDeptHTTPServer(hs, srv)
		dept.RegisterDeptServer(gs, srv)
	})
}

// GetDept 获取指定租户信息
func (app *Dept) GetDept(c context.Context, req *dept.GetDeptRequest) (*dept.GetDeptReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	ent, err := app.srv.GetDept(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := dept.GetDeptReply{}
	if err := value.Transform(ent, &reply); err != nil {
		ctx.Logger().Errorw("msg", "get dept resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListCurrentDept 获取租户信息列表
func (app *Dept) ListCurrentDept(c context.Context, req *dept.ListDeptRequest) (*dept.ListDeptReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListDeptRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list dept req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, err := app.srv.ListCurrentDept(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := dept.ListDeptReply{}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get dept resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListDept 获取租户信息列表
func (app *Dept) ListDept(c context.Context, req *dept.ListDeptRequest) (*dept.ListDeptReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListDeptRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list dept req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, err := app.srv.ListDept(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := dept.ListDeptReply{}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get dept resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateDept 创建租户信息
func (app *Dept) CreateDept(c context.Context, req *dept.CreateDeptRequest) (*dept.CreateDeptReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Dept{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create dept req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateDept(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &dept.CreateDeptReply{Id: id}, nil
}

// UpdateDept 更新租户信息
func (app *Dept) UpdateDept(c context.Context, req *dept.UpdateDeptRequest) (*dept.UpdateDeptReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Dept{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create dept req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateDept(ctx, &in); err != nil {
		return nil, err
	}

	return &dept.UpdateDeptReply{}, nil
}

// DeleteDept 删除租户信息
func (app *Dept) DeleteDept(c context.Context, req *dept.DeleteDeptRequest) (*dept.DeleteDeptReply, error) {
	return &dept.DeleteDeptReply{}, app.srv.DeleteDept(core.MustContext(c), req.Id)
}

// ListDeptClassify 获取租户信息列表
func (app *Dept) ListDeptClassify(c context.Context, req *dept.ListDeptClassifyRequest) (*dept.ListDeptClassifyReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListDeptClassifyRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list dept classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := app.srv.ListDeptClassify(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := dept.ListDeptClassifyReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get dept classify resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateDeptClassify 创建租户信息
func (app *Dept) CreateDeptClassify(c context.Context, req *dept.CreateDeptClassifyRequest) (*dept.CreateDeptClassifyReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.DeptClassify{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create dept classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateDeptClassify(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &dept.CreateDeptClassifyReply{Id: id}, nil
}

// UpdateDeptClassify 更新租户信息
func (app *Dept) UpdateDeptClassify(c context.Context, req *dept.UpdateDeptClassifyRequest) (*dept.UpdateDeptClassifyReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.DeptClassify{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create dept classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateDeptClassify(ctx, &in); err != nil {
		return nil, err
	}

	return &dept.UpdateDeptClassifyReply{}, nil
}

// DeleteDeptClassify 删除租户信息
func (app *Dept) DeleteDeptClassify(c context.Context, req *dept.DeleteDeptClassifyRequest) (*dept.DeleteDeptClassifyReply, error) {
	return &dept.DeleteDeptClassifyReply{}, app.srv.DeleteDeptClassify(core.MustContext(c), req.Id)
}
