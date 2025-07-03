package service

import (
	"encoding/base64"
	"time"

	"github.com/forgoer/openssl"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/kratosx/pkg/valx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/pkg/md"
	"github.com/limes-cloud/manager/internal/pkg/ua"
	"github.com/limes-cloud/manager/internal/types"
)

type Auth struct {
	oauth   repository.OAuth
	channel repository.Channel
	app     repository.App
	user    repository.User
	address repository.Address
	role    repository.Role
}

func NewAuth(
	oauth repository.OAuth,
	channel repository.Channel,
	app repository.App,
	user repository.User,
	address repository.Address,
	role repository.Role,
) *Auth {
	return &Auth{
		oauth:   oauth,
		channel: channel,
		app:     app,
		user:    user,
		address: address,
		role:    role,
	}
}

// Auth 外部接口rbac鉴权
func (u *Auth) Auth(ctx kratosx.Context, in *types.AuthRequest) (*md.Auth, error) {
	// 获取鉴权器，判断是否为白名单
	author := ctx.Authentication()
	if author.IsWhitelist(in.Path, in.Method) {
		return nil, nil
	}

	info := md.Get(ctx)

	// 获取当前用户的所有角色key
	var roleKeys []string
	list, _ := u.role.ListCacheRole(ctx)
	for _, item := range list {
		roleKeys = append(roleKeys, item.Keyword)
	}

	// 获取需要跳过的角色
	skips := ctx.Config().App().Authentication.SkipRole
	uniquer := valx.New(roleKeys)
	for _, key := range skips {
		if uniquer.Has(key) {
			return info, nil
		}
	}

	// 进行rbac鉴权
	enforce := ctx.Authentication().Enforce()

	isPass := false
	for _, role := range roleKeys {
		isAuth, _ := enforce.Enforce(role, in.Path, in.Method)
		if isAuth {
			isPass = true
			break
		}
	}
	if !isPass {
		return nil, errors.ForbiddenError()
	}

	return info, nil
}

// OAuthOAuthWay 获取渠道授权方式
func (u *Auth) OAuthOAuthWay(ctx kratosx.Context, req *types.OAuthWayRequest) (*types.GetOAuthWayReply, error) {
	// 获取渠道信息
	channel, err := u.channel.GetChannelByKeyword(ctx, req.Keyword)
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

	return author.GetOAuthWay(ctx, &types.GetOAuthWayRequest{
		User:      req.User,
		UserAgent: header.RequestHeader().Get("User-Agent"),
		IP:        ctx.ClientIP(),
	})
}

// ReportOAuthCode 上报授权code
func (u *Auth) ReportOAuthCode(ctx kratosx.Context, req *types.ReportOAuthCodeRequest) error {
	// 获取渠道信息
	channel, err := u.channel.GetChannelByKeyword(ctx, req.Keyword)
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

func (u *Auth) OAuthLogin(ctx kratosx.Context, req *types.OAuthLoginRequest) (*types.OAuthLoginReply, error) {
	// 获取渠道信息
	channel, err := u.channel.GetChannelByKeyword(ctx, req.Keyword)
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
		Expire: proto.Uint32(uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))),
	}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
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
func (u *Auth) OAuthBind(ctx kratosx.Context, req *types.OAuthBindRequest) (string, error) {
	// 获取渠道信息
	channel, err := u.channel.GetChannelByKeyword(ctx, req.Keyword)
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

	// 用户瞪目
	token, err := u.UserLogin(ctx, req.UserLoginRequest)
	if err != nil {
		return "", err
	}

	// 获取token中的用户信息
	user := md.Auth{}
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

// GetUserLoginCaptcha 获取用户登陆验证吗
func (u *Auth) GetUserLoginCaptcha(ctx kratosx.Context) (*types.GetUserLoginCaptchaReply, error) {
	resp, err := ctx.Captcha().Image(loginCaptchaKey, ctx.ClientIP())
	if err != nil {
		return nil, errors.GenCaptchaError(err.Error())
	}

	return &types.GetUserLoginCaptchaReply{
		Uuid:    resp.ID(),
		Expire:  uint32(resp.Expire().Seconds()),
		Captcha: resp.Base64String(),
	}, nil
}

func (u *Auth) UserLogin(ctx kratosx.Context, in *types.UserLoginRequest) (token string, rerr error) {
	var (
		user  *entity.User
		utype string
	)

	defer func() {
		if errors.IsUserDisableError(rerr) || errors.IsVerifyCaptchaError(rerr) {
			return
		}
		u.addLoginLog(ctx, utype, in, rerr)
	}()

	if err := ctx.Captcha().VerifyImage(loginCaptchaKey, ctx.ClientIP(), in.CaptchaId, in.Captcha); err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		rerr = errors.VerifyCaptchaError()
		return
	}

	passByte, _ := base64.StdEncoding.DecodeString(in.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		ctx.Logger().Errorw("msg", "rsa decode error", "err", err.Error())
		rerr = errors.VerifyCaptchaError()
		return
	}

	pw := struct {
		Password string `json:"password"`
		Time     int64  `json:"time"`
	}{}
	if json.Unmarshal(decryptData, &pw) != nil {
		ctx.Logger().Errorw("msg", "password format error", "err", err.Error())
		rerr = errors.PasswordError()
		return
	}

	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		ctx.Logger().Errorw(
			"msg", "login pwd timeout",
			"current", time.Now().UnixMilli(),
			"submit", pw.Time,
			"dt", time.Now().UnixMilli()-pw.Time,
		)
		rerr = errors.PasswordExpireError()
		return
	}

	// 获取用户信息
	if valx.IsPhone(in.Username) {
		utype = "phone"
		user, err = u.user.GetUserByPhone(ctx, in.Username)
	} else if valx.IsEmail(in.Username) {
		utype = "email"
		user, err = u.user.GetUserByEmail(ctx, in.Username)
	} else {
		rerr = errors.UsernameFormatError()
		return
	}

	if err != nil {
		ctx.Logger().Errorw("msg", "get user info error", "err", err.Error())
		rerr = errors.UsernameNotExistError()
		return
	}

	// 判断密码
	if !crypto.CompareHashPwd(user.Password, pw.Password) {
		rerr = errors.PasswordError()
		return
	}

	// 生产token
	token, rerr = u.genToken(ctx, user)
	if rerr != nil {
		return
	}

	return token, nil
}

// UserLogout 退出登陆
func (u *Auth) UserLogout(ctx kratosx.Context) error {
	token := ctx.Token()
	if token != "" {
		ctx.JWT().AddBlacklist(token)
	}
	return nil
}

// UserRefreshToken 用户刷新token
func (u *Auth) UserRefreshToken(ctx kratosx.Context) (string, error) {
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return "", errors.RefreshTokenError(err.Error())
	}

	uid := md.UserId(ctx)

	// 更新用户表token信息
	if err := u.user.UpdateUser(ctx, &entity.User{
		BaseModel: ktypes.BaseModel{Id: uid},
		Token:     proto.String(token),
	}); err != nil {
		ctx.Logger().Errorw("refresh token error", "err", err.Error())
		return "", errors.SystemError()
	}

	return token, nil
}

// ListLoginLog 获取用户登陆日志
func (u *Auth) ListLoginLog(ctx kratosx.Context, req *types.ListLoginLogRequest) ([]*entity.LoginLog, uint32, error) {
	list, total, err := u.user.ListLoginLog(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

func (u *Auth) genToken(ctx kratosx.Context, user *entity.User) (token string, err error) {
	if user.Status != nil && !*user.Status {
		err = errors.UserDisableError()
		return
	}

	// 判断用户token是否可用,可用直接返回当前token
	if user.Token != nil && *user.Token != "" && user.ExpireAt > (time.Now().Unix()+600) {
		return *user.Token, nil
	}

	// 获取角色信息
	roleIds, err := u.role.GetEnableRoleIdsByUID(ctx, user.Id)
	if err != nil {
		return "", errors.SystemError()
	}
	if len(roleIds) == 0 {
		return "", errors.SystemError("当前用户未分配部门或角色")
	}

	// 获取职位信息
	// var (
	//	jobIds []uint32
	// )
	// for _, job := range user.Jobs {
	//	jobIds = append(jobIds, job.Id)
	// }

	// 获取部门信息
	var depIds []uint32
	for _, job := range user.Departments {
		depIds = append(depIds, job.Id)
	}

	token, err = ctx.JWT().NewToken(md.New(&md.Auth{
		UserId:   user.Id,
		UserName: user.Name,
		RoleIds:  roleIds,
		// JobIds:        jobIds,
		DepartmentIds: depIds,
	}))
	if err != nil {
		ctx.Logger().Errorw("msg", "gen user token error", "err", err.Error())
		err = errors.GenTokenError()
		return
	}

	// 保存登陆信息
	expired := time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds())
	data := &entity.User{
		BaseModel: ktypes.BaseModel{Id: user.Id},
		Token:     &token,
		LoggedAt:  time.Now().Unix(),
		ExpireAt:  expired,
	}

	if err = u.user.UpdateUser(ctx, data); err != nil {
		ctx.Logger().Errorw("msg", "update user login info error", "err", err.Error())
		err = errors.SystemError()
		return
	}

	return token, nil
}

func (u *Auth) getUserByCA(ctx kratosx.Context, cid uint32, aid string) (*entity.User, error) {
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

func (u *Auth) addLoginLog(ctx kratosx.Context, tp string, user *types.UserLoginRequest, rerr error) {
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

	if rerr != nil {
		var er *kerrors.Error
		if ok := kerrors.As(rerr, &er); ok {
			code = int(er.GetCode())
			desc = er.Error()
		} else {
			code = 400
			desc = rerr.Error()
		}
	}

	_, _ = u.user.CreateLoginLog(ctx, &entity.LoginLog{
		Username:    user.Username,
		Type:        tp,
		IP:          ctx.ClientIP(),
		Address:     u.address.GetIPAddress(ip),
		Browser:     ug.Name + " " + ug.Version,
		Device:      ug.OS + " " + ug.OSVersion,
		Code:        code,
		Description: desc,
	})
}
