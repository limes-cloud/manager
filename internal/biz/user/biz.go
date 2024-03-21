package user

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/forgoer/openssl"
	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/biz/department"
	"github.com/limes-cloud/manager/internal/biz/role"
	"github.com/limes-cloud/manager/internal/config"
	"github.com/limes-cloud/manager/internal/pkg/md"
)

type UseCase struct {
	repo     Repo
	depRepo  department.Repo
	roleRepo role.Repo
	conf     *config.Config
}

func NewUseCase(conf *config.Config, repo Repo, depRepo department.Repo, roleRepo role.Repo) *UseCase {
	return &UseCase{
		repo:     repo,
		depRepo:  depRepo,
		roleRepo: roleRepo,
		conf:     conf,
	}
}

const (
	captcha        = "changePassword"
	imageCaptchaTp = "login"

	passwordCert = "login"
)

// CurrentUser 查询当前用户信息
func (u *UseCase) CurrentUser(ctx kratosx.Context) (*User, error) {
	user, err := u.repo.GetUser(ctx, md.UserId(ctx))
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	return user, nil
}

func (u *UseCase) PageUser(ctx kratosx.Context, in *PageUserRequest) ([]*User, uint32, error) {
	ids, err := u.depRepo.AllManagerDepartmentIds(ctx, md.UserId(ctx))
	if err != nil {
		return nil, 0, errors.DatabaseFormat(err.Error())
	}

	in.DepartmentIds = ids
	list, total, err := u.repo.PageUser(ctx, in)
	if err != nil {
		return nil, 0, errors.DatabaseFormat(err.Error())
	}
	return list, total, nil
}

func (u *UseCase) AddUser(ctx kratosx.Context, user *User) (uint32, error) {
	// 判读是否具有部门权限
	depIds, err := u.depRepo.AllManagerDepartmentIds(ctx, md.UserId(ctx))
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	if !util.InList(depIds, user.DepartmentId) {
		return 0, errors.DepartmentPermissions()
	}

	// 判断当前用户是否具有角色的权限
	rids, err := u.roleRepo.GetChildrenIds(ctx, md.RoleId(ctx))
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	for _, item := range user.UserRoles {
		if !util.InList(rids, item.RoleID) {
			return 0, errors.RolePermissionsFormat(fmt.Sprintf("role_id:%d", item.RoleID))
		}
	}

	// 设置默认值
	user.Nickname = user.Name
	user.Avatar = u.conf.DefaultUserAvatar
	user.Password = util.ParsePwd(u.conf.DefaultUserPassword)
	user.RoleId = user.UserRoles[0].RoleID
	user.Status = proto.Bool(true)

	id, err := u.repo.AddUser(ctx, user)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

func (u *UseCase) UpdateUser(ctx kratosx.Context, user *User) error {
	// 超级管理员不允许修改
	if user.ID == 1 {
		return errors.EditSystemData()
	}

	// 查询用户信息
	old, err := u.repo.GetUser(ctx, user.ID)
	if err != nil {
		return errors.NotFound()
	}

	// 判读是否具有用户修改权限
	depIds, err := u.depRepo.AllManagerDepartmentIds(ctx, md.UserId(ctx))
	if err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	if !util.InList(depIds, old.DepartmentId) {
		return errors.UserPermissions()
	}

	// 存在修改部门
	if user.DepartmentId != 0 && !util.InList(depIds, user.DepartmentId) {
		return errors.DepartmentPermissions()
	}

	// 存在修改角色
	if len(user.UserRoles) != 0 {
		rids, err := u.roleRepo.GetChildrenIds(ctx, md.RoleId(ctx))
		if err != nil {
			return errors.DatabaseFormat(err.Error())
		}
		for _, item := range user.UserRoles {
			if !util.InList(rids, item.RoleID) {
				return errors.RolePermissionsFormat(fmt.Sprintf("role_id:%d", item.RoleID))
			}
		}
		user.RoleId = user.UserRoles[0].RoleID
	}

	if user.Password != "" {
		user.Password = util.ParsePwd(user.Password)
	}

	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

func (u *UseCase) UpdateCurrentUser(ctx kratosx.Context, in *UpdateCurrentUserRequest) error {
	return u.repo.UpdateUser(ctx, &User{
		BaseModel: types.BaseModel{ID: md.UserId(ctx)},
		Nickname:  in.Nickname,
		Gender:    in.Gender,
	})
}

func (u *UseCase) DeleteUser(ctx kratosx.Context, id uint32) error {
	if id == 1 {
		return errors.DeleteSystemData()
	}
	if id == md.UserId(ctx) {
		return errors.DeleteSelfUser()
	}

	// 查询用户信息
	old, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return errors.NotFound()
	}

	// 判读是否具有用户修改权限
	depIds, err := u.depRepo.AllManagerDepartmentIds(ctx, md.UserId(ctx))
	if err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	if !util.InList(depIds, old.DepartmentId) {
		return errors.UserPermissions()
	}

	if err := u.repo.DeleteUser(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

func (u *UseCase) OfflineUser(ctx kratosx.Context, id uint32) error {
	user, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return errors.NotFound()
	}

	ctx.JWT().AddBlacklist(user.Token)
	return nil
}

func (u *UseCase) ChangePasswordCaptcha(ctx kratosx.Context) (*CaptchaReply, error) {
	user, err := u.repo.GetUser(ctx, md.UserId(ctx))
	if err != nil {
		return nil, errors.NotFound()
	}

	resp, err := ctx.Captcha().Email(captcha, ctx.ClientIP(), user.Email)
	if err != nil {
		return nil, errors.SendEmailCaptchaFormat(err.Error())
	}

	return &CaptchaReply{
		Uuid:   resp.ID(),
		Expire: uint32(resp.Expire().Seconds()),
	}, nil
}

func (u *UseCase) ChangeUserPassword(ctx kratosx.Context, in *ChangeUserPasswordRequest) error {
	if err := ctx.Captcha().VerifyEmail(captcha, ctx.ClientIP(), in.CaptchaId, in.Captcha); err != nil {
		return errors.VerifyCaptcha()
	}

	old, err := u.repo.GetUser(ctx, md.UserId(ctx))
	if err != nil {
		return errors.NotFound()
	}

	user := User{
		BaseModel: types.BaseModel{ID: old.ID},
		Password:  in.Password,
	}
	if err := u.repo.UpdateUser(ctx, &user); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

func (u *UseCase) EnableUser(ctx kratosx.Context, id uint32) error {
	return u.UpdateUser(ctx, &User{
		BaseModel: types.BaseModel{ID: id},
		Status:    proto.Bool(true),
		Disabled:  proto.String(""),
	})
}

func (u *UseCase) DisableUser(ctx kratosx.Context, id uint32, desc string) error {
	return u.UpdateUser(ctx, &User{
		BaseModel: types.BaseModel{ID: id},
		Status:    proto.Bool(false),
		Disabled:  proto.String(desc),
	})
}

func (u *UseCase) ResetUserPassword(ctx kratosx.Context, id uint32) error {
	return u.UpdateUser(ctx, &User{
		BaseModel: types.BaseModel{ID: id},
		Password:  u.conf.DefaultUserPassword,
	})
}

func (u *UseCase) SwitchCurrentUserRole(ctx kratosx.Context, id uint32) error {
	if !u.repo.HasRole(ctx, md.UserId(ctx), id) {
		return errors.RolePermissions()
	}

	return u.UpdateUser(ctx, &User{
		BaseModel: types.BaseModel{ID: id},
		RoleId:    id,
	})
}

func (u *UseCase) UserLoginCaptcha(ctx kratosx.Context) (*CaptchaReply, error) {
	res, err := ctx.Captcha().Image(imageCaptchaTp, ctx.ClientIP())
	if err != nil {
		return nil, errors.NewCaptchaFormat(err.Error())
	}

	return &CaptchaReply{
		Uuid:    res.ID(),
		Captcha: res.Base64String(),
		Expire:  uint32(res.Expire().Seconds()),
	}, nil
}

func (u *UseCase) UserLogin(ctx kratosx.Context, in *UserLoginRequest) (string, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyImage(imageCaptchaTp, ctx.ClientIP(), in.CaptchaId, in.Captcha); err != nil {
		return "", errors.VerifyCaptcha()
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(in.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		return "", errors.RsaDecodeFormat(err.Error())
	}

	// 序列化密码
	pw := struct {
		Password string `json:"password"`
		Time     int64  `json:"time"`
	}{}

	if json.Unmarshal(decryptData, &pw) != nil {
		return "", errors.PasswordFormat()
	}

	// 判断当前时间戳是否过期,超过10s则拒绝
	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		ctx.Logger().Errorw(
			"login pwd timeout", time.Now().UnixMilli(),
			"submit", pw.Time,
			"dt", time.Now().UnixMilli()-pw.Time,
		)
		return "", errors.PasswordExpire()
	}

	// 获取用户信息
	var user *User
	if util.IsPhone(in.Username) {
		user, err = u.repo.GetUserByPhone(ctx, in.Username)
	} else if util.IsEmail(in.Username) {
		user, err = u.repo.GetUserByEmail(ctx, in.Username)
	} else {
		return "", errors.UsernameFormat()
	}

	// 查询不到用户信息
	if err != nil {
		return "", errors.UsernameNotExist()
	}

	// 用户被禁用则拒绝登陆
	if user.Status != nil && !*user.Status {
		return "", errors.UserDisableFormat(*user.Disabled)
	}

	// 角色被禁用则拒绝登录
	if user.Role.Status != nil {
		if !*user.Role.Status {
			return "", errors.RoleDisable()
		}
		// 上级被禁用下级也不允许登录
		if !u.roleRepo.ParentStatus(ctx, user.RoleId) {
			return "", errors.RoleDisable()
		}
	}

	// 对比用户密码，错误则拒绝登陆
	if !util.CompareHashPwd(user.Password, pw.Password) {
		return "", errors.UserPassword()
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(md.New(&md.Auth{
		UserId:            user.ID,
		RoleId:            user.RoleId,
		RoleKeyword:       user.Role.Keyword,
		DepartmentId:      user.DepartmentId,
		DepartmentKeyword: user.Department.Keyword,
	}))
	if err != nil {
		return "", errors.NewTokenFormat(err.Error())
	}

	// 更新用户的当前token
	if err := u.repo.UpdateUser(ctx, &User{
		BaseModel: types.BaseModel{ID: user.ID},
		Token:     token,
		LastLogin: time.Now().Unix(),
	}); err != nil {
		return "", errors.DatabaseFormat(err.Error())
	}

	return token, nil
}

func (u *UseCase) UserLogout(ctx kratosx.Context) error {
	token := ctx.Token()
	if token != "" {
		ctx.JWT().AddBlacklist(token)
	}
	return nil
}

func (u *UseCase) UserRefreshToken(ctx kratosx.Context) (string, error) {
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return "", errors.RefreshTokenFormat(err.Error())
	}
	return token, nil
}
