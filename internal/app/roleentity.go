package app

import (
	"context"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"

	"github.com/limes-cloud/manager/api/roleentity"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
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
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list role req transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateRoleEntity 创建用户字段
func (fd *RoleEntity) CreateRoleEntity(c context.Context, req *roleentity.CreateRoleEntityRequest) (*roleentity.CreateRoleEntityReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.RoleEntity{}
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create role req transform error", "err", err)
		return nil, errors.TransformError()
	}

	id, err := fd.srv.CreateRoleEntity(ctx, &in)
	if err != nil {
		return nil, err
	}
	return &roleentity.CreateRoleEntityReply{Id: id}, nil
}

// UpdateRoleEntity 更新用户字段
func (fd *RoleEntity) UpdateRoleEntity(c context.Context, req *roleentity.UpdateRoleEntityRequest) (*roleentity.UpdateRoleEntityReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.RoleEntity{}
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "update role req transform error", "err", err)
		return nil, errors.TransformError()
	}

	err := fd.srv.UpdateRoleEntity(ctx, &in)
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
