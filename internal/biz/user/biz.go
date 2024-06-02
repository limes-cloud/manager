package user

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/forgoer/openssl"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/pkg/md"
	"google.golang.org/protobuf/proto"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// GetUser 获取指定的用户信息 fixed code
func (u *UseCase) GetUser(ctx kratosx.Context, req *GetUserRequest) (*User, error) {
	var (
		res *User
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetUser(ctx, *req.Id)
	} else if req.Phone != nil {
		res, err = u.repo.GetUserByPhone(ctx, *req.Phone)
	} else if req.Email != nil {
		res, err = u.repo.GetUserByEmail(ctx, *req.Email)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}

	for _, role := range res.Roles {
		if role.Id == res.RoleId {
			res.Role = role
		}
	}

	return res, nil
}

// ListUser 获取用户信息列表 fixed code
func (u *UseCase) ListUser(ctx kratosx.Context, req *ListUserRequest) ([]*User, uint32, error) {
	all, scopes, err := u.repo.GetDepartmentDataScope(ctx, md.UserId(ctx))
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
	return list, total, nil
}

// CreateUser 创建用户信息  fixed code
func (u *UseCase) CreateUser(ctx kratosx.Context, req *User) (uint32, error) {
	all, scopes, err := u.repo.GetDepartmentDataScope(ctx, md.UserId(ctx))
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	if !all && !valx.InList(scopes, req.DepartmentId) {
		return 0, errors.DepartmentPermissionsError()
	}

	all, scopes, err = u.repo.GetRoleDataScope(ctx, md.RoleId(ctx))
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	if !all {
		for _, ur := range req.UserRoles {
			if !valx.InList(scopes, ur.RoleId) {
				return 0, errors.RolePermissionsError()
			}
		}
	}

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

// UpdateUser 更新用户信息 fixed code
func (u *UseCase) UpdateUser(ctx kratosx.Context, req *User) error {
	if req.Id == 1 {
		return errors.EditSystemDataError()
	}

	curUserId := md.UserId(ctx)
	is, err := u.repo.HasUserDataScope(ctx, curUserId, req.Id)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	if !is {
		return errors.UserPermissionsError()
	}

	all, scopes, err := u.repo.GetDepartmentDataScope(ctx, curUserId)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	if !all && valx.InList(scopes, req.DepartmentId) {
		return errors.DepartmentPermissionsError()
	}

	if len(req.UserRoles) != 0 {
		all, scopes, err = u.repo.GetRoleDataScope(ctx, curUserId)
		if err != nil {
			return errors.DatabaseError(err.Error())
		}
		if !all {
			for _, ur := range req.UserRoles {
				if !all && valx.InList(scopes, ur.RoleId) {
					return errors.DepartmentPermissionsError()
				}
			}
		}
	}

	if err := u.repo.UpdateUser(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateUserStatus 更新用户信息状态 fixed code
func (u *UseCase) UpdateUserStatus(ctx kratosx.Context, id uint32, status bool) error {
	if id == 1 {
		return errors.DeleteSystemDataError()
	}
	is, err := u.repo.HasUserDataScope(ctx, md.UserId(ctx), id)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	if !is {
		return errors.UserPermissionsError()
	}

	if err := u.repo.UpdateUserStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}

	if status == false {
		token, logged, err := u.repo.GetUserToken(ctx, id)
		if err != nil {
			return errors.DatabaseError(err.Error())
		}
		if token == nil {
			return nil
		}
		expire := ctx.Config().App().JWT.Expire
		if logged > time.Now().Unix()-int64(expire.Seconds()) {
			return nil
		}
		ctx.JWT().AddBlacklist(*token)
	}

	return nil
}

// DeleteUser 删除用户信息 fixed code
func (u *UseCase) DeleteUser(ctx kratosx.Context, ids []uint32) (uint32, error) {
	for _, id := range ids {
		if id == 1 {
			return 0, errors.DeleteSystemDataError()
		}

		is, err := u.repo.HasUserDataScope(ctx, md.UserId(ctx), id)
		if err != nil {
			return 0, errors.DatabaseError(err.Error())
		}
		if !is {
			return 0, errors.UserPermissionsError()
		}
	}

	total, err := u.repo.DeleteUser(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// ResetUserPassword 重置用户密码
func (u *UseCase) ResetUserPassword(ctx kratosx.Context, id uint32) error {
	if id == 1 {
		return errors.DeleteSystemDataError()
	}

	is, err := u.repo.HasUserDataScope(ctx, md.UserId(ctx), id)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	if !is {
		return errors.UserPermissionsError()
	}

	if err = u.repo.UpdateUser(ctx, &User{
		Id:       id,
		Password: crypto.EncodePwd(u.conf.DefaultUserPassword),
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}

	return nil
}

// GetCurrentUser 获取当前的用户信息
func (u *UseCase) GetCurrentUser(ctx kratosx.Context) (*User, error) {
	res, err := u.repo.GetUser(ctx, md.UserId(ctx))
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	for _, role := range res.Roles {
		if role.Id == res.RoleId {
			res.Role = role
		}
	}
	return res, nil
}

// UpdateCurrentUser 更新当前的基础信息
func (u *UseCase) UpdateCurrentUser(ctx kratosx.Context, req *UpdateCurrentUserRequest) error {
	if err := u.repo.UpdateUser(ctx, &User{
		Id:       md.UserId(ctx),
		Avatar:   req.Avatar,
		Nickname: req.Nickname,
		Gender:   req.Gender,
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}

	return nil
}

// UpdateCurrentUserRole 切换当前用户角色
func (u *UseCase) UpdateCurrentUserRole(ctx kratosx.Context, rid uint32) error {
	curUserId := md.UserId(ctx)
	all, scopes, err := u.repo.GetRoleDataScope(ctx, curUserId)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	if !all && valx.InList(scopes, rid) {
		return errors.DepartmentPermissionsError()
	}

	if err = u.repo.UpdateUser(ctx, &User{
		Id:     curUserId,
		RoleId: rid,
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// UpdateCurrentUserSetting 保存当前用户设置
func (u *UseCase) UpdateCurrentUserSetting(ctx kratosx.Context, setting string) error {
	if err := u.repo.UpdateUser(ctx, &User{
		Id:      md.UserId(ctx),
		Setting: &setting,
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// SendCurrentUserCaptcha 发送当前用户验证吗
func (u *UseCase) SendCurrentUserCaptcha(ctx kratosx.Context, tp string) (*SendCurrentUserCaptchaReply, error) {
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

	return &SendCurrentUserCaptchaReply{
		Uuid:   resp.ID(),
		Expire: uint32(resp.Expire().Seconds()),
	}, nil
}

// UpdateCurrentUserPassword 修改当前用户密码
func (u *UseCase) UpdateCurrentUserPassword(ctx kratosx.Context, req *UpdateCurrentUserPasswordRequest) error {
	switch u.conf.ChangePasswordType {
	case ChangePwCaptchaType:
		if req.CaptchaId == nil || req.Captcha == nil {
			return errors.ParamsError()
		}
		if err := ctx.Captcha().VerifyEmail(pwdCaptchaKey, ctx.ClientIP(), *req.CaptchaId, *req.Captcha); err != nil {
			return errors.VerifyCaptchaError()
		}
	case ChangePwPasswordType:
		if req.OldPassword == nil {
			return errors.ParamsError()
		}
		password, err := u.repo.GetUserPassword(ctx, md.UserId(ctx))
		if err != nil {
			return errors.DatabaseError(err.Error())
		}
		if !crypto.CompareHashPwd(password, *req.OldPassword) {
			return errors.PasswordError()
		}
	default:
		return errors.SystemError("验证方式配置错误")
	}

	user := User{
		Id:       md.UserId(ctx),
		Password: crypto.EncodePwd(req.Password),
	}
	if err := u.repo.UpdateUser(ctx, &user); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// GetUserLoginCaptcha 获取用户登陆验证吗
func (u *UseCase) GetUserLoginCaptcha(ctx kratosx.Context) (*GetUserLoginCaptchaReply, error) {
	resp, err := ctx.Captcha().Image(loginCaptchaKey, ctx.ClientIP())
	if err != nil {
		return nil, errors.GenCaptchaError(err.Error())
	}

	return &GetUserLoginCaptchaReply{
		Uuid:    resp.ID(),
		Expire:  uint32(resp.Expire().Seconds()),
		Captcha: resp.Base64String(),
	}, nil
}

func (u *UseCase) UserLogin(ctx kratosx.Context, in *UserLoginRequest) (string, error) {
	if err := ctx.Captcha().VerifyImage(loginCaptchaKey, ctx.ClientIP(), in.CaptchaId, in.Captcha); err != nil {
		return "", errors.VerifyCaptchaError()
	}

	passByte, _ := base64.StdEncoding.DecodeString(in.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		ctx.Logger().Error("rsa解密失败:", err.Error())
		return "", errors.SystemError()
	}

	pw := struct {
		Password string `json:"password"`
		Time     int64  `json:"time"`
	}{}
	if json.Unmarshal(decryptData, &pw) != nil {
		return "", errors.PasswordError()
	}

	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		ctx.Logger().Errorw(
			"msg", "login pwd timeout",
			"current", time.Now().UnixMilli(),
			"submit", pw.Time,
			"dt", time.Now().UnixMilli()-pw.Time,
		)
		return "", errors.PasswordExpireError()
	}

	// 获取用户信息
	var user *User
	if valx.IsPhone(in.Username) {
		user, err = u.repo.GetUserByPhone(ctx, in.Username)
	} else if valx.IsEmail(in.Username) {
		user, err = u.repo.GetUserByEmail(ctx, in.Username)
	} else {
		return "", errors.UsernameFormatError()
	}

	if err != nil {
		return "", errors.UsernameNotExistError()
	}

	if user.Status != nil && !*user.Status {
		return "", errors.UserDisableError()
	}

	var (
		notSwitch   bool
		enableRoles []*Role
	)
	for _, role := range user.Roles {
		if role.Status != nil && *role.Status == true {
			enableRoles = append(enableRoles, role)
			if role.Id == user.RoleId {
				notSwitch = true
				user.Role = role
			}
		}
	}
	if len(enableRoles) == 0 {
		return "", errors.RoleDisableError()
	}

	if !notSwitch {
		user.RoleId = enableRoles[0].Id
		user.Role = enableRoles[0]
	}

	password, err := u.repo.GetUserPassword(ctx, user.Id)
	if err != nil {
		return "", errors.DatabaseError(err.Error())
	}

	if !crypto.CompareHashPwd(password, pw.Password) {
		return "", errors.PasswordError()
	}

	token, err := ctx.JWT().NewToken(md.New(&md.Auth{
		UserId:            user.Id,
		RoleId:            user.RoleId,
		RoleKeyword:       user.Role.Keyword,
		DepartmentId:      user.DepartmentId,
		DepartmentKeyword: user.Department.Keyword,
	}))
	if err != nil {
		return "", errors.GenTokenError()
	}

	data := &User{
		Id:       user.Id,
		RoleId:   user.RoleId,
		Token:    &token,
		LoggedAt: time.Now().Unix(),
	}

	if err := u.repo.UpdateUser(ctx, data); err != nil {
		return "", errors.DatabaseError(err.Error())
	}
	return token, nil
}

// UserLogout 退出登陆
func (u *UseCase) UserLogout(ctx kratosx.Context) error {
	token := ctx.Token()
	if token != "" {
		ctx.JWT().AddBlacklist(token)
	}
	return nil
}

// UserRefreshToken 用户刷新token
func (u *UseCase) UserRefreshToken(ctx kratosx.Context) (string, error) {
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return "", errors.RefreshTokenError(err.Error())
	}
	return token, nil
}
