package service

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/limes-cloud/manager/api/auth"

	midauth "github.com/limes-cloud/manager/internal/middleware/auth"

	"github.com/forgoer/openssl"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/limes-cloud/kratosx/library/captcha"
	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/kratosx/pkg/ua"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
	bc "github.com/mojocn/base64Captcha"
)

const (
	captchaName  = "login"
	captchaScene = "userLogin"

	passwordCert = "login"
)

type Auth struct {
	auth    repository.Auth
	user    repository.User
	address repository.Address
	oauth   repository.OAuth
	menu    repository.Menu
	scope   repository.Scope
	tenant  repository.Tenant
}

func NewAuth(
	auth repository.Auth,
	user repository.User,
	address repository.Address,
	oauth repository.OAuth,
	menu repository.Menu,
	scope repository.Scope,
	tenant repository.Tenant,
) *Auth {
	return &Auth{
		auth:    auth,
		user:    user,
		address: address,
		oauth:   oauth,
		menu:    menu,
		scope:   scope,
		tenant:  tenant,
	}
}

// ApiAuth 外部接口rbac鉴权
func (u *Auth) ApiAuth(ctx core.Context, in *auth.ApiAuthRequest) (*auth.ApiAuthReply, error) {
	info := &auth.ApiAuthReply{
		TenantId: ctx.Auth().TenantId,
		UserId:   ctx.Auth().UserId,
		DeptId:   ctx.Auth().DeptId,
		Username: ctx.Auth().Username,
	}
	if ctx.JWT().IsWhitelist(in.Path, in.Method) {
		return info, nil
	}
	if ctx.IsSuperAdmin() {
		return info, nil
	}
	// 获取鉴权器，判断是否为白名单
	id, tp, has := u.menu.GetMenuTypeInfo(in.Path, in.Method)
	if !has {
		return nil, errors.ForbiddenError()
	}
	if tp == entity.MenuTypeBasic {
		return info, nil
	}

	// 判断是否拥有api权限
	if !u.scope.HasMenuScope(ctx, id) {
		return nil, errors.ForbiddenError()
	}

	_ = ctx.Pool().GoFunc(func() {
		ctx = ctx.Clone()
		log := entity.AuthLog{
			MenuId: id,
		}

		_, _ = u.auth.CreateAuthLog(ctx, &log)
	})

	return info, nil
}

// GetUserLoginCaptcha 获取用户登陆验证吗
func (u *Auth) GetUserLoginCaptcha(ctx core.Context) (*types.GetUserLoginCaptchaReply, error) {
	cpt, err := ctx.Captcha().Get(ctx, captchaName)
	if err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		return nil, errors.VerifyCaptchaError()
	}
	resp, err := cpt.GetCaptcha(&captcha.GetCaptchaRequest{
		Scene:    captchaScene,
		ClientIP: ctx.ClientIP(),
	})
	if err != nil {
		return nil, errors.GenCaptchaError(err.Error())
	}

	// 生成验证码图片
	dt := bc.NewDriverDigit(80, 200, 6, 0.7, 80)
	item, err := dt.DrawCaptcha(resp.VerifyCode)
	if err != nil {
		return nil, errors.GenCaptchaError(err.Error())
	}

	return &types.GetUserLoginCaptchaReply{
		Uuid:    resp.UUID,
		Expire:  uint32(resp.Expire.Seconds()),
		Captcha: item.EncodeB64string(),
	}, nil
}

// UserLogin 用户登陆
func (u *Auth) UserLogin(ctx core.Context, req *types.UserLoginRequest) (token string, err error) {
	var user *entity.User
	defer func() {
		if errors.IsUserDisableError(err) ||
			errors.IsVerifyCaptchaError(err) ||
			errors.IsPasswordError(err) ||
			errors.IsTenantNotFoundError(err) {
			return
		}

		if user == nil {
			return
		}
		u.addLoginLog(ctx, "username", user, err)
	}()

	// 判断验证码
	cpt, err := ctx.Captcha().Get(ctx, captchaName)
	if err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		err = errors.VerifyCaptchaError()
		return
	}
	if err = cpt.VerifyCaptcha(&captcha.VerifyCaptchaRequest{
		Scene:      captchaScene,
		ClientIP:   ctx.ClientIP(),
		UUID:       req.CaptchaId,
		VerifyCode: req.Captcha,
	}); err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		err = errors.VerifyCaptchaError()
		return
	}

	// 判断密码
	passByte, _ := base64.StdEncoding.DecodeString(req.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		ctx.Logger().Warnw("msg", "rsa decode error", "err", err.Error())
		err = errors.VerifyCaptchaError()
		return
	}
	pw := struct {
		Password string `json:"password"`
		Time     int64  `json:"time"`
	}{}
	if json.Unmarshal(decryptData, &pw) != nil {
		ctx.Logger().Errorw("msg", "password format error", "err", err.Error())
		err = errors.PasswordError()
		return
	}
	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		ctx.Logger().Errorw(
			"msg", "login pwd timeout",
			"current", time.Now().UnixMilli(),
			"submit", pw.Time,
			"dt", time.Now().UnixMilli()-pw.Time,
		)
		err = errors.PasswordExpireError()
		return
	}

	// 获取租户信息
	tenant, err := u.tenant.GetTenantByKeyword(ctx, req.Tenant)
	if err != nil {
		ctx.Logger().Errorw("msg", "get tenant info error", "err", err.Error())
		err = errors.TenantNotFoundError()
		return
	}

	// 获取用户信息
	user, err = u.user.GetUserByTU(ctx, tenant.Id, req.Username)
	if err != nil {
		ctx.Logger().Errorw("msg", "get user info error", "err", err.Error())
		err = errors.UsernameNotExistError()
		return
	}

	// 判断密码
	if !crypto.CompareHashPwd(user.Password, pw.Password) {
		err = errors.PasswordError()
		return
	}

	// 生产token
	return u.genToken(ctx, user)
}

// UserLogout 退出登陆
func (u *Auth) UserLogout(ctx core.Context) error {
	token := ctx.Token()
	if token != "" {
		ctx.JWT().AddBlacklist(token)
	}
	return nil
}

// UserRefreshToken 用户刷新token
func (u *Auth) UserRefreshToken(ctx core.Context) (string, error) {
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return "", errors.RefreshTokenError(err.Error())
	}

	uid := ctx.Auth().UserId

	// 更新用户表token信息
	if err := u.user.UpdateUser(ctx, &entity.User{
		BaseTenantDeptModel: model.BaseTenantDeptModel{Id: uid},
		Token:               proto.String(token),
	}); err != nil {
		ctx.Logger().Errorw("refresh token error", "err", err.Error())
		return "", errors.SystemError()
	}

	return token, nil
}

func (u *Auth) genToken(ctx core.Context, user *entity.User) (token string, err error) {
	if user.Status != nil && !*user.Status {
		err = errors.UserDisableError()
		return
	}

	// 判断用户token是否可用,可用直接返回当前token
	if user.Token != nil && *user.Token != "" && user.ExpireAt > (time.Now().Unix()+600) {
		return *user.Token, nil
	}

	// 构建token信息
	info := midauth.Info{
		TenantId: user.TenantId,
		UserId:   user.Id,
		Username: user.Username,
		DeptId:   user.DeptId,
	}
	token, err = ctx.JWT().NewToken(info.ToMap())
	if err != nil {
		ctx.Logger().Errorw("msg", "gen user token error", "err", err.Error())
		err = errors.GenTokenError()
		return
	}

	// 更新用户token信息
	expired := time.Now().Unix() + int64(ctx.JWT().Config().Expire.Seconds())
	data := &entity.User{
		BaseTenantDeptModel: model.BaseTenantDeptModel{Id: user.Id},
		Token:               &token,
		LoggedAt:            time.Now().Unix(),
		ExpireAt:            expired,
	}
	if err = u.user.UpdateUser(ctx, data); err != nil {
		ctx.Logger().Errorw("msg", "update user login info error", "err", err.Error())
		err = errors.SystemError()
		return
	}

	return token, nil
}

func (u *Auth) addLoginLog(ctx core.Context, tp string, user *entity.User, err error) {
	header, ok := transport.FromServerContext(ctx)
	if !ok {
		return
	}

	var (
		ip   = ctx.ClientIP()
		ug   = ua.Parse(header.RequestHeader().Get("User-Agent"))
		code = 200
		desc = "登陆成功"
	)

	if err != nil {
		var er *kerrors.Error
		if ok := kerrors.As(err, &er); ok {
			code = int(er.GetCode())
			desc = er.Error()
		} else {
			code = 400
			desc = err.Error()
		}
	}

	_, _ = u.auth.CreateLoginLog(ctx, &entity.LoginLog{
		CreateTenantUserModel: model.CreateTenantUserModel{
			UserId:   user.Id,
			DeptId:   user.DeptId,
			TenantId: user.TenantId,
		},
		Type:        tp,
		IP:          ctx.ClientIP(),
		Address:     u.address.GetAddressByIP(ip),
		Browser:     ug.Name + " " + ug.Version,
		Device:      ug.OS + " " + ug.OSVersion,
		Code:        code,
		Description: desc,
	})
}

// OAuthHandler 获取渠道授权方式
func (u *Auth) OAuthHandler(ctx core.Context, req *types.OAuthHandlerRequest) (*types.OAuthHandlerReply, error) {
	// 获取渠道信息
	channel, err := u.oauth.GetOAuthChannelByKeyword(ctx, req.Keyword)
	if err != nil {
		return nil, errors.SystemError()
	}

	// 获取授权方式
	author, err := u.oauth.GetOAuthor(channel)
	if err != nil {
		ctx.Logger().Errorw("get author error", "err", err.Error())
		return nil, errors.SystemError()
	}

	// 获取header
	header, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, errors.SystemError()
	}

	return author.OAuthHandler(ctx, &types.OAuthHandlerRequest{
		User:      req.User,
		UserAgent: header.RequestHeader().Get("User-Agent"),
		IP:        ctx.ClientIP(),
	})
}

// ReportOAuthCode 上报授权code
func (u *Auth) ReportOAuthCode(ctx core.Context, req *types.ReportOAuthCodeRequest) error {
	// 获取渠道信息
	channel, err := u.oauth.GetOAuthChannelByKeyword(ctx, req.Keyword)
	if err != nil {
		return errors.SystemError()
	}

	// 获取授权方式
	author, err := u.oauth.GetOAuthor(channel)
	if err != nil {
		ctx.Logger().Errorw("get author error", "err", err.Error())
		return errors.SystemError()
	}

	// 获取token
	token, err := author.GetOAuthToken(ctx, &types.GetOAuthTokenRequest{
		Code: req.Code,
	})
	if err != nil {
		return errors.GetOAuthTokenError()
	}

	// 存储当前token信息
	return u.oauth.SetTokenByUUID(ctx, req.UUID, token)
}

func (u *Auth) OAuthLogin(ctx core.Context, req *types.OAuthLoginRequest) (*types.OAuthLoginReply, error) {
	// 获取渠道信息
	channel, err := u.oauth.GetOAuthChannelByKeyword(ctx, req.Keyword)
	if err != nil {
		return nil, errors.SystemError()
	}

	author, err := u.oauth.GetOAuthor(channel)
	if err != nil {
		return nil, errors.SystemError()
	}

	// 获取token
	var ti *types.GetOAuthTokenReply
	if req.Code != "" {
		ti, err = author.GetOAuthToken(ctx, &types.GetOAuthTokenRequest{
			IP:   ctx.ClientIP(),
			UUID: req.UUID,
			Code: req.Code,
			User: req.User,
		})
	} else {
		ti, err = u.oauth.GetTokenByUUID(ctx, req.UUID)
	}

	// 获取token
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, errors.NotFoundReportOAuthError()
		}
		return nil, errors.OAuthLoginError()
	}

	// 获取用户信息
	ui, err := author.GetOAuthInfo(ctx, &types.GetOAuthInfoRequest{
		OID:   ti.OID,
		Token: ti.Token,
	})
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	// 判断是否绑定三方
	if !u.oauth.IsBindOAuth(ctx, channel.Id, ui.OID) {
		// 设置token，等待绑定
		if err := u.oauth.SetTokenByUUID(ctx, req.UUID, ti); err != nil {
			ctx.Logger().Errorw("set token error", "err", err.Error())
			return nil, errors.OAuthLoginError()
		}

		return &types.OAuthLoginReply{
			IsBind: false,
		}, nil
	}

	// 获取用户信息
	user, err := u.getUserByCA(ctx, channel.Id, ui.OID)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	reply := types.OAuthLoginReply{
		IsBind: true,
		Expire: proto.Uint32(uint32(time.Now().Unix() + int64(ctx.JWT().Config().Expire.Seconds()))),
	}
	if err := ctx.Transaction(func(ctx core.Context) error {
		// 生成token
		token, err := u.genToken(ctx, user)
		if err != nil {
			return err
		}

		// 更新渠道token信息
		if err := u.oauth.UpdateOAuth(ctx, &entity.OAuth{
			UserId:    user.Id,
			ChannelId: channel.Id,
			Token:     ti.Token,
			ExpiredAt: ti.Expire,
			LoggedAt:  time.Now().Unix(),
		}); err != nil {
			return errors.DatabaseError(err.Error())
		}

		// 更新用户token信息
		reply.Token = &token
		return nil
	}); err != nil {
		return nil, err
	}

	return &reply, nil
}

// OAuthBind 三方授权密码绑定
func (u *Auth) OAuthBind(ctx core.Context, req *types.OAuthBindRequest) (string, error) {
	// 获取渠道信息
	channel, err := u.oauth.GetOAuthChannelByKeyword(ctx, req.Keyword)
	if err != nil {
		return "", errors.SystemError()
	}

	// 获取渠道token信息
	ti, err := u.oauth.GetTokenByUUID(ctx, req.UUID)
	if err != nil {
		return "", errors.BindError("授权信息不存在或已失效")
	}

	// 获取用户信息
	author, err := u.oauth.GetOAuthor(channel)
	if err != nil {
		return "", errors.BindError()
	}
	ui, err := author.GetOAuthInfo(ctx, &types.GetOAuthInfoRequest{
		OID:   ti.OID,
		Token: ti.Token,
	})
	if err != nil {
		return "", errors.BindError("获取用户信息失败")
	}

	// 用户登陆
	token, err := u.UserLogin(ctx, req.UserLoginRequest)
	if err != nil {
		return "", err
	}

	// 获取token中的用户信息
	user := midauth.Info{}
	_ = ctx.JWT().ParseByToken(token, &user)
	if user.UserId == 0 {
		return "", errors.BindError()
	}

	// 绑定用户信息
	extra := entity.OAuthExtra{
		UnionID: ui.UnionId,
	}
	_, err = u.oauth.CreateOAuth(ctx, &entity.OAuth{
		UserId:    user.UserId,
		ChannelId: channel.Id,
		OID:       ui.OID,
		Token:     ti.Token,
		LoggedAt:  time.Now().Unix(),
		ExpiredAt: time.Now().Unix(),
		Extra:     extra.ToString(),
	})
	if err != nil {
		return "", errors.SystemError()
	}
	return token, nil
}

func (u *Auth) getUserByCA(ctx core.Context, cid uint32, aid string) (*entity.User, error) {
	oauth, err := u.oauth.GetOAuthByCO(ctx, cid, aid)
	if err != nil {
		ctx.Logger().Warnw("msg", "get oauth by ca error", "err", err.Error())
		return nil, errors.SystemError()
	}

	user, err := u.user.GetUser(ctx, oauth.UserId)
	if err != nil {
		ctx.Logger().Warnw("msg", "get user  error", "err", err.Error())
		return nil, errors.NotUserError()
	}

	if user.Status != nil && !*user.Status {
		return nil, errors.UserDisableError()
	}
	return user, nil
}

// ListLoginLog 获取用户登陆日志
func (u *Auth) ListLoginLog(ctx core.Context, req *types.ListLoginLogRequest) ([]*entity.LoginLog, uint32, error) {
	list, total, err := u.auth.ListLoginLog(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// ListAuthLog 获取用户登陆日志
func (u *Auth) ListAuthLog(ctx core.Context, req *types.ListAuthLogRequest) ([]*entity.AuthLog, uint32, error) {
	if req.Username != "" {
		user, err := u.user.GetUserByUsername(ctx, req.Username)
		if err != nil {
			ctx.Logger().Warnw("msg", "get user  error", "err", err.Error())
			return nil, 0, errors.NotUserError()
		}
		req.UserId = &user.Id
	}

	list, total, err := u.auth.ListAuthLog(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}
