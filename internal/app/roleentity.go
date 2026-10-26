package app

import (
	"context"
	"encoding/json"

	"github.com/limes-cloud/kratosx/model"

	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/roleentity"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type RoleEntity struct {
	roleentity.UnimplementedRoleEntityServer
	srv *service.RoleEntity
}

func NewRoleEntity() *RoleEntity {
	return &RoleEntity{
		srv: service.NewRoleEntity(
			dbs.NewRoleEntity(),
			dbs.NewScope(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewRoleEntity()
		roleentity.RegisterRoleEntityHTTPServer(hs, srv)
		roleentity.RegisterRoleEntityServer(gs, srv)
	})
}

// ListRoleEntity 获取用户字段列表
func (fd *RoleEntity) ListRoleEntity(c context.Context, req *roleentity.ListRoleEntityRequest) (*roleentity.ListRoleEntityReply, error) {
	ctx := core.MustContext(c)
	list, total, err := fd.srv.ListRoleEntity(ctx, ctx.AppKey(req.GetApp()), &types.ListRoleEntityRequest{
		Search: page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		RoleId:   req.RoleId,
		EntityId: req.EntityId,
	})
	if err != nil {
		return nil, err
	}

	// 处理请求参数
	reply := roleentity.ListRoleEntityReply{Total: total}
	for _, v := range list {
		item := &roleentity.ListRoleEntityReply_Data{
			Id:        v.Id,
			RoleId:    v.RoleId,
			EntityId:  v.EntityId,
			Action:    v.Action,
			Scope:     v.Scope,
			CreatedAt: uint32(v.CreatedAt),
			UpdatedAt: uint32(v.UpdatedAt),
			Entity: &roleentity.ListRoleEntityReply_Entity{
				Id:      v.Entity.Id,
				Name:    v.Entity.Name,
				Comment: v.Entity.Comment,
			},
		}
		_ = json.Unmarshal([]byte(v.Rules), &item.Rules)
		_ = json.Unmarshal([]byte(v.Depts), &item.Depts)
		reply.List = append(reply.List, item)
	}

	return &reply, nil
}

// CreateRoleEntity 创建用户字段
func (fd *RoleEntity) CreateRoleEntity(c context.Context, req *roleentity.CreateRoleEntityRequest) (*roleentity.CreateRoleEntityReply, error) {
	ctx := core.MustContext(c)
	app := ctx.AppKey(req.GetApp())

	// 批量模式：items 非空时按各自的 action/scope 逐条创建（事务内一次提交）
	if len(req.GetItems()) > 0 {
		var fields []*entity.RoleEntity
		for _, item := range req.GetItems() {
			fields = append(fields, &entity.RoleEntity{
				App:      app,
				RoleId:   req.RoleId,
				EntityId: req.EntityId,
				Action:   item.Action,
				Scope:    item.Scope,
				Depts:    value.ObjToString(item.Depts),
				Rules:    value.ObjToString(item.Rules),
			})
		}
		ids, err := fd.srv.CreateRoleEntities(ctx, fields)
		if err != nil {
			return nil, err
		}
		if len(ids) > 0 {
			return &roleentity.CreateRoleEntityReply{Id: ids[0]}, nil
		}
		return &roleentity.CreateRoleEntityReply{}, nil
	}

	id, err := fd.srv.CreateRoleEntity(ctx, &entity.RoleEntity{
		App:      app,
		RoleId:   req.RoleId,
		EntityId: req.EntityId,
		Action:   value.Value(req.Action),
		Scope:    value.Value(req.Scope),
		Rules:    value.ObjToString(req.Rules),
		Depts:    value.ObjToString(req.Depts),
	})
	if err != nil {
		return nil, err
	}
	return &roleentity.CreateRoleEntityReply{Id: id}, nil
}

// UpdateRoleEntity 更新用户字段
func (fd *RoleEntity) UpdateRoleEntity(c context.Context, req *roleentity.UpdateRoleEntityRequest) (*roleentity.UpdateRoleEntityReply, error) {
	ctx := core.MustContext(c)

	err := fd.srv.UpdateRoleEntity(ctx, &entity.RoleEntity{
		BaseModel: model.BaseModel{Id: req.Id},
		App:       ctx.AppKey(req.GetApp()),
		RoleId:    value.Value(req.RoleId),
		EntityId:  value.Value(req.EntityId),
		Scope:     value.Value(req.Scope),
		Rules:     value.ObjToString(req.Rules),
		Depts:     value.ObjToString(req.Depts),
	})
	if err != nil {
		return nil, err
	}
	return &roleentity.UpdateRoleEntityReply{}, nil
}

// DeleteRoleEntity 删除用户字段
func (fd *RoleEntity) DeleteRoleEntity(c context.Context, req *roleentity.DeleteRoleEntityRequest) (*roleentity.DeleteRoleEntityReply, error) {
	ctx := core.MustContext(c)
	if err := fd.srv.DeleteRoleEntity(ctx, ctx.AppKey(), req.Id); err != nil {
		return nil, err
	}
	return &roleentity.DeleteRoleEntityReply{}, nil
}
