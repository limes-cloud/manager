package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/api/manager/errors"
	pb "github.com/limes-cloud/manager/api/manager/user/v1"
	"github.com/limes-cloud/manager/internal/biz/user"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/data"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *user.UseCase
}

func NewUserService(conf *conf.Config) *UserService {
	return &UserService{
		uc: user.NewUseCase(conf, data.NewUserRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewUserService(c)
		pb.RegisterUserHTTPServer(hs, srv)
		pb.RegisterUserServer(gs, srv)
	})
}

// ListUser 获取用户信息列表
func (s *UserService) ListUser(c context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	var (
		in  = user.ListUserRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListUser(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListUserReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateUser 创建用户信息 fixed code
func (s *UserService) CreateUser(c context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	var (
		in  = user.User{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	for _, id := range req.RoleIds {
		in.UserRoles = append(in.UserRoles, &user.UserRole{
			RoleId: id,
		})
	}

	for _, id := range req.JobIds {
		in.UserJobs = append(in.UserJobs, &user.UserJob{
			JobId: id,
		})
	}

	id, err := s.uc.CreateUser(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserReply{Id: id}, nil
}

// UpdateUser 更新用户信息 fixed code
func (s *UserService) UpdateUser(c context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	var (
		in  = user.User{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	for _, id := range req.RoleIds {
		in.UserRoles = append(in.UserRoles, &user.UserRole{
			RoleId: id,
		})
	}

	for _, id := range req.JobIds {
		in.UserJobs = append(in.UserJobs, &user.UserJob{
			JobId: id,
		})
	}

	if err := s.uc.UpdateUser(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateUserReply{}, nil
}

// DeleteUser 删除用户信息
func (s *UserService) DeleteUser(c context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	total, err := s.uc.DeleteUser(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserReply{Total: total}, nil
}

// UpdateUserStatus 更新用户信息状态
func (s *UserService) UpdateUserStatus(c context.Context, req *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusReply, error) {
	return &pb.UpdateUserStatusReply{}, s.uc.UpdateUserStatus(kratosx.MustContext(c), req.Id, req.Status)
}

// GetUser 获取指定的用户信息
func (s *UserService) GetUser(c context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	var (
		in  = user.GetUserRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "request transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetUser(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetUserReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *UserService) GetCurrentUser(c context.Context, _ *emptypb.Empty) (*pb.GetUserReply, error) {
	var (
		ctx = kratosx.MustContext(c)
	)

	result, err := s.uc.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	reply := pb.GetUserReply{}
	if err = valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *UserService) ResetUserPassword(c context.Context, req *pb.ResetUserPasswordRequest) (*pb.ResetUserPasswordReply, error) {
	return &pb.ResetUserPasswordReply{}, s.uc.ResetUserPassword(kratosx.MustContext(c), req.Id)
}

func (s *UserService) UpdateCurrentUser(c context.Context, req *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserReply, error) {
	var (
		in  user.UpdateCurrentUserRequest
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "request transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &pb.UpdateCurrentUserReply{}, s.uc.UpdateCurrentUser(kratosx.MustContext(c), &in)
}

func (s *UserService) UpdateCurrentUserRole(c context.Context, req *pb.UpdateCurrentUserRoleRequest) (*pb.UpdateCurrentUserRoleReply, error) {
	return &pb.UpdateCurrentUserRoleReply{}, s.uc.UpdateCurrentUserRole(kratosx.MustContext(c), req.RoleId)
}

func (s *UserService) UpdateCurrentUserPassword(c context.Context, req *pb.UpdateCurrentUserPasswordRequest) (*pb.UpdateCurrentUserPasswordReply, error) {
	var (
		in  user.UpdateCurrentUserPasswordRequest
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "request transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &pb.UpdateCurrentUserPasswordReply{}, s.uc.UpdateCurrentUserPassword(ctx, &in)
}

func (s *UserService) UpdateCurrentUserSetting(c context.Context, req *pb.UpdateCurrentUserSettingRequest) (*pb.UpdateCurrentUserSettingReply, error) {
	return &pb.UpdateCurrentUserSettingReply{}, s.uc.UpdateCurrentUserSetting(kratosx.MustContext(c), req.Setting)
}

func (s *UserService) UserLogin(c context.Context, req *pb.UserLoginRequest) (*pb.UserLoginReply, error) {
	var (
		in  user.UserLoginRequest
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "request transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	token, err := s.uc.UserLogin(ctx, &in)
	if err != nil {
		return nil, err
	}
	return &pb.UserLoginReply{Token: token}, nil
}

func (s *UserService) UserLogout(c context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.uc.UserLogout(kratosx.MustContext(c))
}

func (s *UserService) UserRefreshToken(c context.Context, _ *emptypb.Empty) (*pb.UserRefreshTokenReply, error) {
	token, err := s.uc.UserRefreshToken(kratosx.MustContext(c))
	if err != nil {
		return nil, err
	}
	return &pb.UserRefreshTokenReply{Token: token}, nil
}

func (s *UserService) SendCurrentUserCaptcha(c context.Context, req *pb.SendCurrentUserCaptchaRequest) (*pb.SendCurrentUserCaptchaReply, error) {
	reply, err := s.uc.SendCurrentUserCaptcha(kratosx.MustContext(c), req.Type)
	if err != nil {
		return nil, err
	}
	return &pb.SendCurrentUserCaptchaReply{
		Uuid:    reply.Uuid,
		Captcha: reply.Captcha,
		Expire:  reply.Expire,
	}, nil
}

func (s *UserService) GetUserLoginCaptcha(c context.Context, _ *emptypb.Empty) (*pb.GetUserLoginCaptchaReply, error) {
	reply, err := s.uc.GetUserLoginCaptcha(kratosx.MustContext(c))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserLoginCaptchaReply{
		Uuid:    reply.Uuid,
		Captcha: reply.Captcha,
		Expire:  reply.Expire,
	}, nil
}
