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
	list, total, err := fd.srv.ListRoleEntity(ctx, &types.ListRoleEntityRequest{
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
		_ = json.Unmarshal([]byte(v.Fields), &item.Fields)
		_ = json.Unmarshal([]byte(v.Rules), &item.Rules)
		_ = json.Unmarshal([]byte(v.Depts), &item.Depts)
		reply.List = append(reply.List, item)
	}

	return &reply, nil
}

// CreateRoleEntity 创建用户字段
func (fd *RoleEntity) CreateRoleEntity(c context.Context, req *roleentity.CreateRoleEntityRequest) (*roleentity.CreateRoleEntityReply, error) {
	ctx := core.MustContext(c)

	id, err := fd.srv.CreateRoleEntity(ctx, &entity.RoleEntity{
		RoleId:   req.RoleId,
		EntityId: req.EntityId,
		Action:   req.Action,
		Scope:    req.Scope,
		Fields:   value.ObjToString(req.Fields),
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
		RoleId:    value.Value(req.RoleId),
		EntityId:  value.Value(req.EntityId),
		Action:    value.Value(req.Action),
		Scope:     value.Value(req.Scope),
		Fields:    value.ObjToString(req.Fields),
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
	if err := fd.srv.DeleteRoleEntity(core.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &roleentity.DeleteRoleEntityReply{}, nil
}
