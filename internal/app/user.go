package app

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/manager/api/user"

	"github.com/limes-cloud/manager/internal/domain/entity"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
)

type User struct {
	user.UnimplementedUserServer
	srv *service.User
}

// NewUser 初始化角色应用
func NewUser() *User {
	return &User{
		srv: service.NewUser(
			dbs.NewUser(),
			dbs.NewScope(),
			dbs.NewApp(),
			dbs.NewUserinfo(),
			dbs.NewUserSetting(),
		),
	}
}

// init 应用注册
func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewUser()
		user.RegisterUserHTTPServer(hs, srv)
		user.RegisterUserServer(gs, srv)
	})
}

func (app *User) GetCurrentUser(c context.Context, req *user.GetCurrentUserRequest) (*user.UserObject, error) {
	ctx := core.MustContext(c)

	result, err := app.srv.GetCurrentUser(ctx, &types.GetCurrentUserRequest{
		App: req.App,
	})
	if err != nil {
		return nil, err
	}

	reply := user.UserObject{}
	if err = value.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (app *User) UpdateCurrentUserSetting(c context.Context, req *user.UpdateCurrentUserSettingRequest) (*emptypb.Empty, error) {
	ctx := core.MustContext(c)
	err := app.srv.UpdateCurrentUserSetting(ctx, &types.UpdateCurrentUserSettingRequest{
		App:     req.App,
		Setting: req.Setting,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (app *User) UpdateCurrentUser(c context.Context, req *user.UpdateCurrentUserRequest) (*emptypb.Empty, error) {
	ctx := core.MustContext(c)
	err := app.srv.UpdateCurrentUser(ctx, &types.UpdateCurrentUserRequest{
		Avatar:   req.GetAvatar(),
		Nickname: req.GetNickname(),
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// GetUser 获取指定角色信息
func (app *User) GetUser(c context.Context, req *user.GetUserRequest) (*user.UserObject, error) {
	ctx := core.MustContext(c)

	// 调用服务
	ent, err := app.srv.GetUser(ctx, &types.GetUserRequest{
		Id:       req.Id,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := user.UserObject{}
	if err := value.Transform(ent, &reply); err != nil {
		ctx.Logger().Errorw("msg", "get user resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListUser 获取角色信息列表
func (app *User) ListUser(c context.Context, req *user.ListUserRequest) (*user.ListUserReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = types.ListUserRequest{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "list user req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	list, total, err := app.srv.ListUser(ctx, &in)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	reply := user.ListUserReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Errorw("msg", "get user resp transform error", "err", err)
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateUser 创建角色信息
func (app *User) CreateUser(c context.Context, req *user.CreateUserRequest) (*user.CreateUserReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.User{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create user req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	id, err := app.srv.CreateUser(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &user.CreateUserReply{Id: id}, nil
}

// ResetPassword 重置密码
func (app *User) ResetPassword(c context.Context, req *user.ResetPasswordRequest) (*emptypb.Empty, error) {
	ctx := core.MustContext(c)

	// 调用服务
	if err := app.srv.ResetPassword(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// UpdateUser 更新角色信息
func (app *User) UpdateUser(c context.Context, req *user.UpdateUserRequest) (*user.UpdateUserReply, error) {
	var (
		ctx = core.MustContext(c)
		in  = entity.User{}
	)

	// 处理请求参数
	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Errorw("msg", "create user req transform error", "err", err)
		return nil, errors.TransformError()
	}

	// 调用服务
	if err := app.srv.UpdateUser(ctx, &in); err != nil {
		return nil, err
	}

	return &user.UpdateUserReply{}, nil
}

// DeleteUser 删除角色信息
func (app *User) DeleteUser(c context.Context, req *user.DeleteUserRequest) (*user.DeleteUserReply, error) {
	return &user.DeleteUserReply{}, app.srv.DeleteUser(core.MustContext(c), req.Id)
}
