package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/manager/api/errors"
	v1 "github.com/limes-cloud/manager/api/user/v1"
	biz "github.com/limes-cloud/manager/internal/biz/user"
	"github.com/limes-cloud/manager/internal/config"
	depdata "github.com/limes-cloud/manager/internal/data/department"
	roledata "github.com/limes-cloud/manager/internal/data/role"
	data "github.com/limes-cloud/manager/internal/data/user"
)

type UserService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase

	conf *config.Config
}

func NewUserService(conf *config.Config) *UserService {
	return &UserService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo(), depdata.NewRepo(), roledata.NewRepo()),
	}
}

func (u UserService) PageUser(ctx context.Context, request *v1.PageUserRequest) (*v1.PageUserReply, error) {
	var req biz.PageUserRequest
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := u.uc.PageUser(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageUserReply{Total: total}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (u UserService) AddUser(ctx context.Context, request *v1.AddUserRequest) (*v1.AddUserReply, error) {
	var user biz.User
	if err := util.Transform(request, &user); err != nil {
		return nil, errors.Transform()
	}

	for _, id := range request.RoleIds {
		user.UserRoles = append(user.UserRoles, &biz.UserRole{
			RoleID: id,
		})
	}

	for _, id := range request.JobIds {
		user.UserJobs = append(user.UserJobs, &biz.UserJob{
			JobID: id,
		})
	}

	id, err := u.uc.AddUser(kratosx.MustContext(ctx), &user)
	if err != nil {
		return nil, err
	}

	return &v1.AddUserReply{Id: id}, nil
}

func (u UserService) ChangeUserPassword(ctx context.Context, request *v1.ChangeUserPasswordRequest) (*emptypb.Empty, error) {
	var req biz.ChangeUserPasswordRequest
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}
	return nil, u.uc.ChangeUserPassword(kratosx.MustContext(ctx), &req)
}

func (u UserService) ChangeUserPasswordCaptcha(ctx context.Context, _ *emptypb.Empty) (*v1.CaptchaReply, error) {
	res, err := u.uc.ChangePasswordCaptcha(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := v1.CaptchaReply{}
	if err := util.Transform(res, &reply); err != nil {
		return nil, errors.Transform()
	}
	return &reply, nil
}

func (u UserService) CurrentUser(ctx context.Context, _ *emptypb.Empty) (*v1.User, error) {
	res, err := u.uc.CurrentUser(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := v1.User{}
	if err := util.Transform(res, &reply); err != nil {
		return nil, errors.Transform()
	}
	return &reply, nil
}

func (u UserService) DeleteUser(ctx context.Context, request *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, u.uc.DeleteUser(kratosx.MustContext(ctx), request.Id)
}

func (u UserService) DisableUser(ctx context.Context, request *v1.DisableUserRequest) (*emptypb.Empty, error) {
	return nil, u.uc.DisableUser(kratosx.MustContext(ctx), request.Id, request.Desc)
}

func (u UserService) EnableUser(ctx context.Context, request *v1.EnableUserRequest) (*emptypb.Empty, error) {
	return nil, u.uc.EnableUser(kratosx.MustContext(ctx), request.Id)
}

func (u UserService) OfflineUser(ctx context.Context, request *v1.OfflineUserRequest) (*emptypb.Empty, error) {
	return nil, u.uc.OfflineUser(kratosx.MustContext(ctx), request.Id)
}

func (u UserService) ResetUserPassword(ctx context.Context, request *v1.ResetUserPasswordRequest) (*emptypb.Empty, error) {
	return nil, u.uc.ResetUserPassword(kratosx.MustContext(ctx), request.Id)
}

func (u UserService) SwitchCurrentUserRole(ctx context.Context, request *v1.SwitchCurrentUserRoleRequest) (*v1.SwitchCurrentUserRoleReply, error) {
	return nil, u.uc.SwitchCurrentUserRole(kratosx.MustContext(ctx), request.RoleId)
}

func (u UserService) UpdateUser(ctx context.Context, request *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	var user biz.User
	if err := util.Transform(request, &user); err != nil {
		return nil, errors.Transform()
	}

	for _, id := range request.RoleIds {
		user.UserRoles = append(user.UserRoles, &biz.UserRole{
			RoleID: id,
		})
	}

	for _, id := range request.JobIds {
		user.UserJobs = append(user.UserJobs, &biz.UserJob{
			JobID: id,
		})
	}
	return nil, u.uc.UpdateUser(kratosx.MustContext(ctx), &user)
}

func (u UserService) UpdateCurrentUser(ctx context.Context, request *v1.UpdateCurrentUserRequest) (*emptypb.Empty, error) {
	var req biz.UpdateCurrentUserRequest
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}

	return nil, u.uc.UpdateCurrentUser(kratosx.MustContext(ctx), &req)
}

func (u UserService) UserLogin(ctx context.Context, request *v1.UserLoginRequest) (*v1.UserLoginReply, error) {
	var req biz.UserLoginRequest
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}

	token, err := u.uc.UserLogin(kratosx.MustContext(ctx), &req)
	return &v1.UserLoginReply{
		Token: token,
	}, err
}

func (u UserService) UserLoginCaptcha(ctx context.Context, _ *emptypb.Empty) (*v1.CaptchaReply, error) {
	res, err := u.uc.UserLoginCaptcha(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := v1.CaptchaReply{}
	if err := util.Transform(res, &reply); err != nil {
		return nil, errors.Transform()
	}
	return &reply, nil
}

func (u UserService) UserLogout(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, u.uc.UserLogout(kratosx.MustContext(ctx))
}

func (u UserService) UserRefreshToken(ctx context.Context, _ *emptypb.Empty) (*v1.UserRefreshTokenReply, error) {
	token, err := u.uc.UserRefreshToken(kratosx.MustContext(ctx))
	return &v1.UserRefreshTokenReply{Token: token}, err
}
