package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/manager/internal/domain/entity"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	apientity "github.com/limes-cloud/manager/api/entity"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
)

type Entity struct {
	apientity.UnimplementedEntityServer
	srv *service.Entity
}

// NewEntity 初始化租户应用
func NewEntity() *Entity {
	return &Entity{
		srv: service.NewEntity(
			dbs.NewEntity(),
		),
	}
}

// init 应用注册
func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewEntity()
		apientity.RegisterEntityHTTPServer(hs, srv)
		apientity.RegisterEntityServer(gs, srv)
	})
}

// ImportEntity 载入全部实体信息
func (app *Entity) ImportEntity(c context.Context, req *apientity.ImportEntityRequest) (*emptypb.Empty, error) {
	var (
		ctx  = core.MustContext(c)
		list []*entity.Entity
	)

	// 处理请求参数
	if err := value.Transform(req.List, &list); err != nil {
		ctx.Logger().Errorw("msg", "list entity req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	for _, item := range list {
		item.AppId = req.AppId
	}

	if err := app.srv.ImportEntity(ctx, list); err != nil {
		return nil, err
	}

	// 处理返回数据
	return &emptypb.Empty{}, nil
}

// LoadEntity 载入全部实体信息
func (app *Entity) LoadEntity(c context.Context, _ *emptypb.Empty) (*apientity.LoadEntityReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, err := app.srv.LoadEntity(ctx)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := apientity.LoadEntityReply{}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get entity resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListEntity 获取租户信息列表
func (app *Entity) ListEntity(c context.Context, req *apientity.ListEntityRequest) (*apientity.ListEntityReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListEntityRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list entity req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, err := app.srv.ListEntity(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := apientity.ListEntityReply{Total: uint32(len(list))}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get entity resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateEntity 创建租户信息
func (app *Entity) CreateEntity(c context.Context, req *apientity.CreateEntityRequest) (*apientity.CreateEntityReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Entity{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create entity req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateEntity(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &apientity.CreateEntityReply{Id: id}, nil
}

// UpdateEntity 更新租户信息
func (app *Entity) UpdateEntity(c context.Context, req *apientity.UpdateEntityRequest) (*apientity.UpdateEntityReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Entity{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create entity req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateEntity(ctx, &in); err != nil {
		return nil, err
	}

	return &apientity.UpdateEntityReply{}, nil
}

// DeleteEntity 删除租户信息
func (app *Entity) DeleteEntity(c context.Context, req *apientity.DeleteEntityRequest) (*apientity.DeleteEntityReply, error) {
	return &apientity.DeleteEntityReply{}, app.srv.DeleteEntity(core.MustContext(c), req.Id)
}

// ListEntityField 获取租户信息列表
func (app *Entity) ListEntityField(c context.Context, req *apientity.ListEntityFieldRequest) (*apientity.ListEntityFieldReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListEntityFieldRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list entity classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := app.srv.ListEntityField(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := apientity.ListEntityFieldReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get entity classify resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateEntityField 创建租户信息
func (app *Entity) CreateEntityField(c context.Context, req *apientity.CreateEntityFieldRequest) (*apientity.CreateEntityFieldReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.EntityField{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create entity classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateEntityField(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &apientity.CreateEntityFieldReply{Id: id}, nil
}

// UpdateEntityField 更新租户信息
func (app *Entity) UpdateEntityField(c context.Context, req *apientity.UpdateEntityFieldRequest) (*apientity.UpdateEntityFieldReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.EntityField{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create entity classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateEntityField(ctx, &in); err != nil {
		return nil, err
	}

	return &apientity.UpdateEntityFieldReply{}, nil
}

// DeleteEntityField 删除租户信息
func (app *Entity) DeleteEntityField(c context.Context, req *apientity.DeleteEntityFieldRequest) (*apientity.DeleteEntityFieldReply, error) {
	return &apientity.DeleteEntityFieldReply{}, app.srv.DeleteEntityField(core.MustContext(c), req.Id)
}

// ListEntityRule 获取实体规则列表
func (app *Entity) ListEntityRule(c context.Context, req *apientity.ListEntityRuleRequest) (*apientity.ListEntityRuleReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListEntityRuleRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list entity classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := app.srv.ListEntityRule(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := apientity.ListEntityRuleReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get entity classify resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateEntityRule 创建实体规则
func (app *Entity) CreateEntityRule(c context.Context, req *apientity.CreateEntityRuleRequest) (*apientity.CreateEntityRuleReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.EntityRule{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create entity classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateEntityRule(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &apientity.CreateEntityRuleReply{Id: id}, nil
}

// UpdateEntityRule 更新实体规则
func (app *Entity) UpdateEntityRule(c context.Context, req *apientity.UpdateEntityRuleRequest) (*apientity.UpdateEntityRuleReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.EntityRule{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create entity classify req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateEntityRule(ctx, &in); err != nil {
		return nil, err
	}

	return &apientity.UpdateEntityRuleReply{}, nil
}

// DeleteEntityRule 删除实体规则
func (app *Entity) DeleteEntityRule(c context.Context, req *apientity.DeleteEntityRuleRequest) (*apientity.DeleteEntityRuleReply, error) {
	return &apientity.DeleteEntityRuleReply{}, app.srv.DeleteEntityRule(core.MustContext(c), req.Id)
}
