package app

import (
	"context"
	"encoding/json"

	"github.com/limes-cloud/kratosx/model"

	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/manager/internal/domain/entity"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"

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
			dbs.NewApp(),
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

// ListEntity 获取实体列表
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
	in.App = ctx.AppKey(in.App)

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

// CreateEntity 创建实体
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
	in.App = ctx.AppKey(in.App)

	// 调用服务
	id, err := app.srv.CreateEntity(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &apientity.CreateEntityReply{Id: id}, nil
}

// UpdateEntity 更新实体
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

// DeleteEntity 删除实体
func (app *Entity) DeleteEntity(c context.Context, req *apientity.DeleteEntityRequest) (*apientity.DeleteEntityReply, error) {
	return &apientity.DeleteEntityReply{}, app.srv.DeleteEntity(core.MustContext(c), req.Id)
}

// ListEntityField 获取实体列表
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

// CreateEntityField 创建实体
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

// UpdateEntityField 更新实体
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

// DeleteEntityField 删除实体
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
	for _, v := range list {
		item := &apientity.EntityRule{
			Id:          v.Id,
			EntityId:    v.EntityId,
			Name:        v.Name,
			Description: v.Description,
			Status:      value.Value(v.Status),
			CreatedAt:   uint32(v.CreatedAt),
			UpdatedAt:   uint32(v.UpdatedAt),
		}
		if err := json.Unmarshal([]byte(v.Expression), &item.Expression); err != nil {
			ctx.Logger().Warnw("msg", "parse entity rule expression error", "id", v.Id, "err", err.Error())
		}
		reply.List = append(reply.List, item)
	}

	return &reply, nil
}

// CreateEntityRule 创建实体规则
func (app *Entity) CreateEntityRule(c context.Context, req *apientity.CreateEntityRuleRequest) (*apientity.CreateEntityRuleReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	id, err := app.srv.CreateEntityRule(ctx, &entity.EntityRule{
		EntityId:    req.EntityId,
		Name:        req.Name,
		Expression:  value.ObjToString(req.Expression),
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &apientity.CreateEntityRuleReply{Id: id}, nil
}

// UpdateEntityRule 更新实体规则
func (app *Entity) UpdateEntityRule(c context.Context, req *apientity.UpdateEntityRuleRequest) (*apientity.UpdateEntityRuleReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	if err := app.srv.UpdateEntityRule(ctx, &entity.EntityRule{
		BaseModel:   model.BaseModel{Id: req.Id},
		EntityId:    value.Value(req.EntityId),
		Name:        value.Value(req.Name),
		Status:      req.Status,
		Expression:  value.ObjToString(req.Expression),
		Description: value.Value(req.Description),
	}); err != nil {
		return nil, err
	}

	return &apientity.UpdateEntityRuleReply{}, nil
}

// DeleteEntityRule 删除实体规则
func (app *Entity) DeleteEntityRule(c context.Context, req *apientity.DeleteEntityRuleRequest) (*apientity.DeleteEntityRuleReply, error) {
	return &apientity.DeleteEntityRuleReply{}, app.srv.DeleteEntityRule(core.MustContext(c), req.Id)
}

// RegisterEntity 应用启动自注册实体与字段（mode: append|sync）
func (app *Entity) RegisterEntity(c context.Context, req *apientity.RegisterEntityRequest) (*apientity.RegisterEntityReply, error) {
	ctx := core.MustContext(c)

	// api 消息 → types（嵌套节点逐层转换）
	in := &types.RegisterEntityRequest{
		AppKeyword: req.GetAppKeyword(),
		Mode:       req.GetMode(),
	}
	for _, item := range req.GetEntities() {
		if item == nil {
			continue
		}
		node := &types.RegisterEntityNode{
			Database: item.GetDatabase(),
			Name:     item.GetName(),
			Comment:  item.GetComment(),
		}
		for ind, f := range item.GetFields() {
			if f == nil {
				continue
			}
			field := &types.RegisterEntityFieldNode{Name: f.GetName(), Comment: f.GetComment()}
			if f.GetIndex() > 0 {
				field.Index = int(f.GetIndex())
			} else {
				field.Index = ind
			}
			node.Fields = append(node.Fields, field)
		}
		for _, r := range item.GetRules() {
			if r == nil {
				continue
			}
			node.Rules = append(node.Rules, &types.RegisterEntityRuleNode{
				Name:        r.GetName(),
				Status:      r.Status,
				Expression:  r.GetExpression(),
				Description: r.GetDescription(),
			})
		}
		in.Entities = append(in.Entities, node)
	}

	reply, err := app.srv.RegisterEntity(ctx, in)
	if err != nil {
		return nil, err
	}
	return &apientity.RegisterEntityReply{
		Created: reply.Created,
		Updated: reply.Updated,
		Deleted: reply.Deleted,
	}, nil
}
