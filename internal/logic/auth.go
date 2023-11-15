package logic

import (
	"encoding/base64"
	"encoding/json"
	v1 "manager/api/v1"
	"manager/config"
	"manager/internal/model"
	"manager/internal/types"
	"manager/pkg/md"
	"manager/pkg/util"
	"time"

	"github.com/forgoer/openssl"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
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

// Auth 退出登录
func (a *Auth) Auth(ctx kratos.Context, in *v1.AuthRequest) (*emptypb.Empty, error) {
	author := ctx.Authentication().Enforce()
	role := md.RoleKeyword(ctx)

	skipRoles := []string{}
	_ = ctx.Config().Value("authentication.skipRole").Scan(&skipRoles)

	if util.InList(skipRoles, role) {
		return nil, nil
	}

	isAuth, _ := author.Enforce(role, in.Path, in.Method)
	if !isAuth {
		return nil, v1.ErrorForbidden()
	}
	return nil, nil
}

// LoginCaptcha 登录验证码
func (a *Auth) LoginCaptcha(ctx kratos.Context, in *emptypb.Empty) (*v1.LoginCaptchaReply, error) {
	ctx.Logger().Errorf("captcha ip:", ctx.ClientIP())
	res, err := ctx.Captcha().Image(imageCaptchaTp, ctx.ClientIP())
	if err != nil {
		return nil, v1.ErrorNewCaptchaFormat(err.Error())
	}

	return &v1.LoginCaptchaReply{
		Uuid:    res.ID(),
		Captcha: res.Base64String(),
		Expire:  uint32(res.Expire().Seconds()),
	}, nil
}

// Login 用户登录
func (a *Auth) Login(ctx kratos.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	ctx.Logger().Errorf("login ip:", ctx.ClientIP())
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyImage(imageCaptchaTp, ctx.ClientIP(), in.CaptchaId, in.Captcha); err != nil {
		return nil, v1.ErrorVerifyCaptcha()
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(in.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		return nil, v1.ErrorRsaDecodeFormat(err.Error())
	}

	// 序列化密码
	var pw types.Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return nil, v1.ErrorPasswordFormat()
	}

	// 判断当前时间戳是否过期,超过10s则拒绝
	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		return nil, v1.ErrorPasswordExpire()
	}

	// 获取用户信息
	user := model.User{}
	if util.IsPhone(in.Username) {
		err = user.OneByPhone(ctx, in.Username)
	} else if util.IsEmail(in.Username) {
		err = user.OneByEmail(ctx, in.Username)
	} else {
		return nil, v1.ErrorUsernameFormat()
	}

	// 查询不到用户信息
	if err != nil {
		return nil, v1.ErrorUsernameNotExist()
	}

	// 用户被禁用则拒绝登陆
	if user.Status != nil && !*user.Status {
		return nil, v1.ErrorUserDisableFormat(*user.Disabled)
	}

	// 角色被禁用则拒绝登录
	role := model.Role{}
	if !role.RoleStatus(ctx, user.RoleID) {
		return nil, v1.ErrorRoleDisable()
	}

	// 对比用户密码，错误则拒绝登陆
	if !util.CompareHashPwd(user.Password, pw.Password) {
		return nil, v1.ErrorUserPassword()
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(md.New(&user))
	if err != nil {
		return nil, v1.ErrorNewTokenFormat(err.Error())
	}

	// 更新用户的当前token
	nu := model.User{}
	nu.ID = user.ID
	nu.Token = token
	nu.LastLogin = uint32(time.Now().Unix())
	if err = nu.Update(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return &v1.LoginReply{Token: token}, nil
}

// Logout 退出登录
func (a *Auth) Logout(ctx kratos.Context) (*emptypb.Empty, error) {
	// 将此用户加入黑名单redis中
	token := ctx.Token()
	if token == "" {
		return nil, v1.ErrorEmptyToken()
	}
	ctx.JWT().AddBlacklist(token)
	return nil, nil
}

// ParseToken 对Token信息进行解析返回
func (a *Auth) ParseToken(ctx kratos.Context) (*v1.ParseTokenReply, error) {
	return &v1.ParseTokenReply{
		UserId: md.UserId(ctx),
		RoleId: md.UserId(ctx),
	}, nil
}

// RefreshToken 刷新用户token
func (a *Auth) RefreshToken(ctx kratos.Context) (*v1.RefreshTokenReply, error) {
	// 进行token续期
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return nil, v1.ErrorRefreshTokenFormat(err.Error())
	}

	return &v1.RefreshTokenReply{
		Token: token,
	}, nil
}
