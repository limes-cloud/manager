package logic

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/forgoer/openssl"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/manager/api/v1"
	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/internal/model"
	"github.com/limes-cloud/manager/pkg/md"
	"github.com/limes-cloud/manager/pkg/util"
)

type Auth struct {
	conf *config.Config
}

func NewAuth(conf *config.Config) *Auth {
	return &Auth{
		conf: conf,
	}
}

const (
	imageCaptchaTp = "login"
	passwordCert   = "login"
)

type Password struct {
	Password string `json:"password"`
	Time     int64  `json:"time"`
}

// Auth 退出登录
func (a *Auth) Auth(ctx kratosx.Context, in *v1.AuthRequest) (*emptypb.Empty, error) {
	author := ctx.Authentication().Enforce()
	role := md.RoleKeyword(ctx)

	var skipRoles []string
	_ = ctx.Config().Value("authentication.skipRole").Scan(&skipRoles)

	if util.InList(skipRoles, role) {
		return nil, nil
	}

	isAuth, _ := author.Enforce(role, in.Path, in.Method)
	if !isAuth {
		return nil, v1.ForbiddenError()
	}
	return nil, nil
}

// LoginCaptcha 登录验证码
func (a *Auth) LoginCaptcha(ctx kratosx.Context, in *emptypb.Empty) (*v1.LoginCaptchaReply, error) {
	res, err := ctx.Captcha().Image(imageCaptchaTp, ctx.ClientIP())
	if err != nil {
		return nil, v1.NewCaptchaErrorFormat(err.Error())
	}

	return &v1.LoginCaptchaReply{
		Uuid:    res.ID(),
		Captcha: res.Base64String(),
		Expire:  uint32(res.Expire().Seconds()),
	}, nil
}

// Login 用户登录
func (a *Auth) Login(ctx kratosx.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyImage(imageCaptchaTp, ctx.ClientIP(), in.CaptchaId, in.Captcha); err != nil {
		return nil, v1.VerifyCaptchaError()
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(in.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		return nil, v1.RsaDecodeErrorFormat(err.Error())
	}

	// 序列化密码
	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return nil, v1.PasswordFormatError()
	}

	// 判断当前时间戳是否过期,超过10s则拒绝
	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		return nil, v1.PasswordExpireError()
	}

	// 获取用户信息
	user := model.User{}
	if util.IsPhone(in.Username) {
		err = user.FindByPhone(ctx, in.Username)
	} else if util.IsEmail(in.Username) {
		err = user.FindByEmail(ctx, in.Username)
	} else {
		return nil, v1.UsernameFormatError()
	}

	// 查询不到用户信息
	if err != nil {
		return nil, v1.UsernameNotExistError()
	}

	// 用户被禁用则拒绝登陆
	if user.Status != nil && !*user.Status {
		return nil, v1.UserDisableErrorFormat(*user.Disabled)
	}

	// 角色被禁用则拒绝登录
	if user.Role.Status != nil {
		if !*user.Role.Status {
			return nil, v1.RoleDisableError()
		}
		// 上级被禁用下级也不允许登录
		if !user.Role.ParentStatus(ctx) {
			return nil, v1.RoleDisableError()
		}
	}

	// 对比用户密码，错误则拒绝登陆
	if !util.CompareHashPwd(user.Password, pw.Password) {
		return nil, v1.UserPasswordError()
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(md.New(&user))
	if err != nil {
		return nil, v1.NewTokenErrorFormat(err.Error())
	}

	// 更新用户的当前token
	nu := model.User{}
	nu.ID = user.ID
	nu.Token = token
	nu.LastLogin = time.Now().Unix()
	if err = nu.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return &v1.LoginReply{Token: token}, nil
}

// Logout 退出登录
func (a *Auth) Logout(ctx kratosx.Context) (*emptypb.Empty, error) {
	// 将此用户加入黑名单redis中
	token := ctx.Token()
	if token == "" {
		return nil, v1.EmptyTokenError()
	}
	ctx.JWT().AddBlacklist(token)
	return nil, nil
}

// ParseToken 对Token信息进行解析返回
func (a *Auth) ParseToken(ctx kratosx.Context) (*v1.ParseTokenReply, error) {
	return &v1.ParseTokenReply{
		UserId: md.UserId(ctx),
		RoleId: md.UserId(ctx),
	}, nil
}

// RefreshToken 刷新用户token
func (a *Auth) RefreshToken(ctx kratosx.Context) (*v1.RefreshTokenReply, error) {
	// 进行token续期
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return nil, v1.RefreshTokenErrorFormat(err.Error())
	}

	return &v1.RefreshTokenReply{
		Token: token,
	}, nil
}
