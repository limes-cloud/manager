package service

import (
	"time"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/kratosx/pkg/valx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/pkg/md"
	"github.com/limes-cloud/manager/internal/types"
)

const (
	ChangePwCaptchaType  = "captcha"
	ChangePwPasswordType = "password"

	pwdCaptchaKey   = "changePassword"
	loginCaptchaKey = "login"

	passwordCert = "login"
)

type User struct {
	conf    *conf.Config
	repo    repository.User
	dept    repository.Department
	role    repository.Role
	file    repository.File
	address repository.Address
}

func NewUser(
	config *conf.Config,
	repo repository.User,
	dept repository.Department,
	role repository.Role,
	file repository.File,
	address repository.Address,
) *User {
	return &User{
		conf:    config,
		repo:    repo,
		dept:    dept,
		role:    role,
		file:    file,
		address: address,
	}
}

// GetUser 获取指定的用户信息
func (u *User) GetUser(ctx kratosx.Context, req *types.GetUserRequest) (*entity.User, error) {
	var (
		user *entity.User
		err  error
	)

	isPurview := func(userId uint32) error {
		has, err := u.dept.HasDepartmentPurview(ctx, md.UserId(ctx), userId)
		if err != nil {
			ctx.Logger().Errorw("msg", "get dept purview error", "err", err.Error())
			return errors.DatabaseError()
		}
		if !has {
			return errors.DepartmentPurviewError()
		}
		return nil
	}

	if req.Id != nil {
		if err := isPurview(*req.Id); err != nil {
			return nil, err
		}
		user, err = u.repo.GetUser(ctx, *req.Id)
	} else if req.Phone != nil {
		user, err = u.repo.GetUserByPhone(ctx, *req.Phone)
	} else if req.Email != nil {
		user, err = u.repo.GetUserByEmail(ctx, *req.Email)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}

	for _, role := range user.Roles {
		if role.Id == user.RoleId {
			user.Role = role
		}
	}

	if err := isPurview(user.Id); err != nil {
		return nil, err
	}

	return user, nil
}

// ListUser 获取用户信息列表
func (u *User) ListUser(ctx kratosx.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	all, scopes, err := u.dept.GetDepartmentDataScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, 0, errors.DatabaseError(err.Error())
	}
	if !all {
		req.DepartmentIds = scopes
	}

	list, total, err := u.repo.ListUser(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}

	for ind, item := range list {
		if item.Avatar != nil {
			url := u.file.GetFileURL(ctx, *item.Avatar)
			list[ind].Avatar = &url
		}
	}
	return list, total, nil
}

// CreateUser 创建用户信息
func (u *User) CreateUser(ctx kratosx.Context, req *entity.User) (uint32, error) {
	// 判断是否具有部门权限
	hasDeptPurview, err := u.dept.HasDepartmentPurview(ctx, md.UserId(ctx), req.DepartmentId)
	if err != nil {
		ctx.Logger().Warnw("msg", "get dept purview error", "err", err.Error())
		return 0, errors.DatabaseError()
	}
	if !hasDeptPurview {
		return 0, errors.DepartmentPurviewError()
	}

	// 判断是否具有角色权限
	all, scopes, err := u.role.GetRoleDataScope(ctx, md.RoleId(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "get role scopes error", "err", err.Error())
		return 0, errors.DatabaseError()
	}
	if !all {
		inr := valx.New(scopes)
		for _, ur := range req.UserRoles {
			if !inr.Has(ur.RoleId) {
				return 0, errors.RolePurviewError()
			}
		}
	}

	// 创建用户信息
	req.Nickname = req.Name
	req.Avatar = &u.conf.DefaultUserAvatar
	req.Password = crypto.EncodePwd(u.conf.DefaultUserPassword)
	req.RoleId = req.UserRoles[0].RoleId
	req.Status = proto.Bool(true)

	id, err := u.repo.CreateUser(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateUser 更新用户信息
func (u *User) UpdateUser(ctx kratosx.Context, user *entity.User) error {
	var curUserId = md.UserId(ctx)

	// 系统数据不允许修改
	if user.Id == 1 {
		return errors.EditSystemDataError()
	}

	// 获取用户基础信息
	oldUser, err := u.repo.GetBaseUser(ctx, user.Id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get base user error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 获取当前用的部门权限
	all, scopes, err := u.dept.GetDepartmentDataScope(ctx, curUserId)
	if err != nil {
		ctx.Logger().Warnw("msg", "get dept purview error", "err", err.Error())
		return errors.DatabaseError()
	}

	inr := valx.New(scopes)

	// 判断是否具体操作用户权限
	if !all && !inr.Has(oldUser.DepartmentId) {
		return errors.DepartmentPurviewError()
	}

	// 判断是否具有变更后的部门权限
	if !all && !inr.Has(user.DepartmentId) {
		return errors.DepartmentPurviewError()
	}

	// 判断是否具有变更后的角色权限
	if len(user.UserRoles) != 0 {
		all, scopes, err = u.role.GetRoleDataScope(ctx, curUserId)
		if err != nil {
			ctx.Logger().Warnw("msg", "get role scopes error", "err", err.Error())
			return errors.DatabaseError()
		}

		rin := valx.New(scopes)
		for _, ur := range user.UserRoles {
			if !all && rin.Has(ur.RoleId) {
				return errors.RolePurviewError()
			}
		}
	}

	// 更新用户
	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateUserStatus 更新用户信息状态 fixed code
func (u *User) UpdateUserStatus(ctx kratosx.Context, id uint32, status bool) error {
	// 系统数据不允许修改
	if id == 1 {
		return errors.EditSystemDataError()
	}

	// 获取用户基础信息
	oldUser, err := u.repo.GetBaseUser(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get base user error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 获取当前用的部门权限
	hasPurview, err := u.dept.HasDepartmentPurview(ctx, md.UserId(ctx), oldUser.DepartmentId)
	if err != nil {
		ctx.Logger().Warnw("msg", "get dept purview error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 判断是否具体操作用户权限
	if !hasPurview {
		return errors.DepartmentPurviewError()
	}

	// 更新角色状态
	if err := u.repo.UpdateUserStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}

	// 如果是禁用用户
	expire := ctx.Config().App().JWT.Expire.Seconds()
	if !status && oldUser.Token != nil && oldUser.LoggedAt > time.Now().Unix()-int64(expire) {
		ctx.JWT().AddBlacklist(*oldUser.Token)
	}

	return nil
}

// DeleteUser 删除用户信息
func (u *User) DeleteUser(ctx kratosx.Context, id uint32) error {
	// 系统数据不允许修改
	if id == 1 {
		return errors.EditSystemDataError()
	}

	// 获取用户基础信息
	oldUser, err := u.repo.GetBaseUser(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get base user error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 获取当前用的部门权限
	hasPurview, err := u.dept.HasDepartmentPurview(ctx, md.UserId(ctx), oldUser.DepartmentId)
	if err != nil {
		ctx.Logger().Warnw("msg", "get dept purview error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 判断是否具体操作用户权限
	if !hasPurview {
		return errors.DepartmentPurviewError()
	}

	if err := u.repo.DeleteUser(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ResetUserPassword 重置用户密码
func (u *User) ResetUserPassword(ctx kratosx.Context, id uint32) error {
	// 系统数据不允许修改
	if id == 1 {
		return errors.EditSystemDataError()
	}

	// 获取用户基础信息
	oldUser, err := u.repo.GetBaseUser(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get base user error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 获取当前用的部门权限
	hasPurview, err := u.dept.HasDepartmentPurview(ctx, md.UserId(ctx), oldUser.DepartmentId)
	if err != nil {
		ctx.Logger().Warnw("msg", "get dept purview error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 判断是否具体操作用户权限
	if !hasPurview {
		return errors.DepartmentPurviewError()
	}

	if err = u.repo.UpdateUser(ctx, &entity.User{
		BaseModel: ktypes.BaseModel{Id: id},
		Password:  crypto.EncodePwd(u.conf.DefaultUserPassword),
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}

	return nil
}

// GetCurrentUser 获取当前的用户信息
func (u *User) GetCurrentUser(ctx kratosx.Context) (*entity.User, error) {
	res, err := u.repo.GetUser(ctx, md.UserId(ctx))
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	for _, role := range res.Roles {
		if role.Id == res.RoleId {
			res.Role = role
		}
	}
	if res.Avatar != nil {
		url := u.file.GetFileURL(ctx, *res.Avatar)
		res.Avatar = &url
	}
	return res, nil
}

// UpdateCurrentUser 更新当前的基础信息
func (u *User) UpdateCurrentUser(ctx kratosx.Context, req *types.UpdateCurrentUserRequest) error {
	if err := u.repo.UpdateUser(ctx, &entity.User{
		BaseModel: ktypes.BaseModel{Id: md.UserId(ctx)},
		Avatar:    req.Avatar,
		Nickname:  req.Nickname,
		Gender:    req.Gender,
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}

	return nil
}

// UpdateCurrentUserRole 切换当前用户角色
func (u *User) UpdateCurrentUserRole(ctx kratosx.Context, rid uint32) error {
	rids, err := u.repo.GetUserRoleIds(ctx, md.UserId(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "get user roles error", "err", err.Error())
		return errors.DatabaseError()
	}
	if !valx.InList(rids, rid) {
		return errors.RolePurviewError()
	}

	if err = u.repo.UpdateUser(ctx, &entity.User{
		BaseModel: ktypes.BaseModel{Id: md.UserId(ctx)},
		RoleId:    rid,
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// UpdateCurrentUserSetting 保存当前用户设置
func (u *User) UpdateCurrentUserSetting(ctx kratosx.Context, setting string) error {
	if err := u.repo.UpdateUser(ctx, &entity.User{
		BaseModel: ktypes.BaseModel{Id: md.UserId(ctx)},
		Setting:   &setting,
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// SendCurrentUserCaptcha 发送当前用户验证吗
func (u *User) SendCurrentUserCaptcha(ctx kratosx.Context, tp string) (*types.SendCurrentUserCaptchaReply, error) {
	tps := []string{pwdCaptchaKey, loginCaptchaKey}
	if !valx.InList(tps, tp) {
		return nil, errors.ParamsError()
	}

	user, err := u.repo.GetUser(ctx, md.UserId(ctx))
	if err != nil {
		return nil, errors.GetError(err.Error())
	}

	resp, err := ctx.Captcha().Email(tp, ctx.ClientIP(), user.Email)
	if err != nil {
		return nil, errors.SendCaptchaError(err.Error())
	}

	return &types.SendCurrentUserCaptchaReply{
		Uuid:   resp.ID(),
		Expire: uint32(resp.Expire().Seconds()),
	}, nil
}

// UpdateCurrentUserPassword 修改当前用户密码
func (u *User) UpdateCurrentUserPassword(ctx kratosx.Context, req *types.UpdateCurrentUserPasswordRequest) error {
	user, err := u.repo.GetBaseUser(ctx, md.UserId(ctx))
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	switch u.conf.ChangePasswordType {
	case ChangePwCaptchaType:
		if req.CaptchaId == nil || req.Captcha == nil {
			return errors.ParamsError()
		}
		if err := ctx.Captcha().VerifyEmail(pwdCaptchaKey, ctx.ClientIP(), *req.CaptchaId, *req.Captcha, user.Email); err != nil {
			return errors.VerifyCaptchaError()
		}
	case ChangePwPasswordType:
		if req.OldPassword == nil {
			return errors.ParamsError()
		}
		if !crypto.CompareHashPwd(user.Password, *req.OldPassword) {
			return errors.PasswordError()
		}
	default:
		return errors.SystemError("验证方式配置错误")
	}

	nu := entity.User{
		BaseModel: ktypes.BaseModel{Id: md.UserId(ctx)},
		Password:  crypto.EncodePwd(req.Password),
	}
	if err := u.repo.UpdateUser(ctx, &nu); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}
