package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/manager/internal/domain/entity"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"

	"github.com/limes-cloud/manager/api/role"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
)

type Role struct {
	role.UnimplementedRoleServer
	srv *service.Role
}

// NewRole 初始化角色应用
func NewRole() *Role {
	return &Role{
		srv: service.NewRole(
			dbs.NewRole(),
			dbs.NewRoleMenu(),
			dbs.NewScope(),
			dbs.NewTenantAdmin(),
		),
	}
}

// init 应用注册
func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewRole()
		role.RegisterRoleHTTPServer(hs, srv)
		role.RegisterRoleServer(gs, srv)
	})
}

// GetRole 获取指定角色信息
func (app *Role) GetRole(c context.Context, req *role.GetRoleRequest) (*role.GetRoleReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	ent, err := app.srv.GetRole(ctx, &types.GetRoleRequest{
		Id:      req.Id,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := role.GetRoleReply{}
	if err := value.Transform(ent, &reply); err != nil {
		ctx.Logger().Errorw("msg", "get role resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListCurrentRole 获取角色信息列表
func (app *Role) ListCurrentRole(c context.Context, req *role.ListRoleRequest) (*role.ListRoleReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListRoleRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list role req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, err := app.srv.ListCurrentRole(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := role.ListRoleReply{}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get role resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListRole 获取角色信息列表
func (app *Role) ListRole(c context.Context, req *role.ListRoleRequest) (*role.ListRoleReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListRoleRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list role req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, err := app.srv.ListRole(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := role.ListRoleReply{}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get role resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateRole 创建角色信息
func (app *Role) CreateRole(c context.Context, req *role.CreateRoleRequest) (*role.CreateRoleReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Role{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create role req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateRole(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &role.CreateRoleReply{Id: id}, nil
}

// UpdateRole 更新角色信息
func (app *Role) UpdateRole(c context.Context, req *role.UpdateRoleRequest) (*role.UpdateRoleReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.Role{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create role req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateRole(ctx, &in); err != nil {
		return nil, err
	}

	return &role.UpdateRoleReply{}, nil
}

// DeleteRole 删除角色信息
func (app *Role) DeleteRole(c context.Context, req *role.DeleteRoleRequest) (*role.DeleteRoleReply, error) {
	return &role.DeleteRoleReply{}, app.srv.DeleteRole(core.MustContext(c), req.Id)
}
