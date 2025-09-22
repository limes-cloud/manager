package app

import (
	"context"

	"github.com/limes-cloud/manager/internal/domain/entity"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/api/userdept"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type UserDept struct {
	userdept.UnimplementedUserDeptServer
	srv *service.UserDept
}

func NewUserDept() *UserDept {
	return &UserDept{
		srv: service.NewUserDept(
			dbs.NewUserDept(),
			dbs.NewScope(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewUserDept()
		userdept.RegisterUserDeptHTTPServer(hs, srv)
		userdept.RegisterUserDeptServer(gs, srv)
	})
}

// ListUserDept 获取角色菜单列表
func (s *UserDept) ListUserDept(c context.Context, req *userdept.ListUserDeptRequest) (*userdept.ListUserDeptReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	list, total, err := s.srv.ListUserDept(ctx, &types.ListUserDeptRequest{
		Search: page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		UserId: req.UserId,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := userdept.ListUserDeptReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list job role resp transform error", "err", err)
		return nil, errors.TransformError()
	}

	return &reply, nil
}

func (s *UserDept) ListDeptUser(c context.Context, req *userdept.ListDeptUserRequest) (*userdept.ListDeptUserReply, error) {
	ctx := core.MustContext(c)
	list, total, err := s.srv.ListDeptUser(ctx, &types.ListDeptUserRequest{
		Search: page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		DeptId: req.DeptId,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}
	reply := userdept.ListDeptUserReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "list role job \resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *UserDept) CreateUserDept(c context.Context, req *userdept.CreateUserDeptRequest) (*userdept.CreateUserDeptReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.CreateUserDept(ctx, &entity.UserDept{
		UserId: req.UserId,
		DeptId: req.DeptId,
		JobId:  req.JobId,
	})
	if err != nil {
		return nil, err
	}
	return &userdept.CreateUserDeptReply{}, nil
}

func (s *UserDept) DeleteUserDept(c context.Context, req *userdept.DeleteUserDeptRequest) (*userdept.DeleteUserDeptReply, error) {
	ctx := core.MustContext(c)
	err := s.srv.DeleteUserDept(ctx, &entity.UserDept{
		UserId: req.UserId,
		DeptId: req.DeptId,
	})
	if err != nil {
		return nil, err
	}
	return &userdept.DeleteUserDeptReply{}, nil
}
