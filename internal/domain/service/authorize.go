package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/limes-cloud/kratosx/library/db/gormtranserror"
	"github.com/limes-cloud/kratosx/pkg/ua"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/internal/pkg/jwt"
	"gorm.io/gorm"

	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/limes-cloud/kratosx/library/captcha"
	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
	bc "github.com/mojocn/base64Captcha"
)

const (
	captchaName          = "login"
	loginCaptchaScene    = "captcha:login"
	bindCaptchaScene     = "captcha:bind"
	registerCaptchaScene = "captcha:register"

	oautherCache = "oauther:cache"
)

type Authorize struct {
	repo     repository.Authorize
	scope    repository.Scope
	tenant   repository.Tenant
	app      repository.App
	user     repository.User
	ao       repository.AppOAuther
	oa       repository.OAuther
	oaexec   repository.OAuthExecer
	uo       repository.UserOAuther
	field    repository.Field
	appfield repository.AppField
	log      repository.Log
	address  repository.Address
	info     repository.Userinfo
	userdept repository.UserDept
	tad      repository.TenantAdmin
	menu     repository.Menu
}

func NewAuthorize(
	repo repository.Authorize,
	scope repository.Scope,
	tenant repository.Tenant,
	app repository.App,
	user repository.User,
	ao repository.AppOAuther,
	oa repository.OAuther,
	oaexec repository.OAuthExecer,
	uo repository.UserOAuther,
	field repository.Field,
	appfield repository.AppField,
	log repository.Log,
	address repository.Address,
	info repository.Userinfo,
	userdept repository.UserDept,
	tad repository.TenantAdmin,
	menu repository.Menu,
) *Authorize {
	return &Authorize{
		repo:     repo,
		scope:    scope,
		tenant:   tenant,
		app:      app,
		user:     user,
		ao:       ao,
		oa:       oa,
		oaexec:   oaexec,
		uo:       uo,
		field:    field,
		appfield: appfield,
		log:      log,
		address:  address,
		info:     info,
		userdept: userdept,
		tad:      tad,
		menu:     menu,
	}
}

type cacheData struct {
	TenantId  uint32 `json:"tenantId"`
	AppId     uint32 `json:"appId"`
	OAutherId uint32 `json:"oautherId"`
	UserId    uint32 `json:"userId"`
	OID       string `json:"oid"`
	Token     string `json:"token"`
	Expire    int64  `json:"expire"`
}

func (s *Authorize) ListOAuther(ctx core.Context, tn string, an string) ([]*entity.AppOAuther, error) {
	// 获取应用
	app, err := s.app.GetAppByKeyword(an)
	if err != nil {
		return nil, errors.GetError("获取应用失败")
	}

	// 获取租户
	tenant, err := s.tenant.GetTenantByKeyword(ctx, tn)
	if err != nil {
		return nil, errors.GetError("获取租户失败")
	}

	list, _, err := s.ao.ListAppOAuther(ctx, &types.ListAppOAutherRequest{
		AppId:    app.Id,
		TenantId: value.Pointer(tenant.Id),
	})
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return list, err
}

// OAutherHandle 授权处理
func (s *Authorize) OAutherHandle(ctx core.Context, req *types.OAutherHandleRequest) (*types.OAutherHandleReply, error) {
	// 获取租户信息
	tenant, err := s.tenant.GetTenantByKeyword(ctx, req.Tenant)
	if err != nil {
		ctx.Logger().Errorw("msg", "get tenant info error", "err", err.Error())
		return nil, errors.TenantNotFoundError()
	}
	if !value.Value(tenant.Status) {
		return nil, errors.SystemError("租户已禁用")
	}

	// 获取应用信息
	app, err := s.app.GetAppByKeyword(req.App)
	if err != nil {
		ctx.Logger().Errorw("msg", "get app info error", "err", err.Error())
		return nil, errors.AppNotFoundError()
	}
	if !value.Value(app.Status) {
		return nil, errors.SystemError(value.Value(app.Reason))
	}

	// 获取渠道信息
	oauther, err := s.oa.GetOAutherByKeyword(ctx, tenant.Id, req.Keyword)
	if err != nil {
		return nil, errors.SystemError(err.Error())
	}

	// 获取授权方式
	author, err := s.oaexec.GetOAuther(oauther)
	if err != nil {
		ctx.Logger().Errorw("get author error", "err", err.Error())
		return nil, errors.SystemError()
	}

	// 获取header
	header, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, errors.SystemError()
	}

	reply, err := author.Handler(ctx, &types.OAutherHandleRequest{
		Account:   req.Account,
		UserAgent: header.RequestHeader().Get("User-Agent"),
		IP:        ctx.ClientIP(),
	})
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *Authorize) OAutherLogin(ctx core.Context, req *types.OAutherLoginRequest) (*types.OAutherLoginReply, error) {
	// 获取租户信息
	tenant, err := s.tenant.GetTenantByKeyword(ctx, req.Tenant)
	if err != nil {
		ctx.Logger().Errorw("msg", "get tenant info error", "err", err.Error())
		return nil, errors.TenantNotFoundError()
	}
	if !value.Value(tenant.Status) {
		return nil, errors.SystemError("租户已禁用")
	}

	// 获取应用信息
	app, err := s.app.GetAppByKeyword(req.App)
	if err != nil {
		ctx.Logger().Errorw("msg", "get app info error", "err", err.Error())
		return nil, errors.AppNotFoundError()
	}
	if !value.Value(app.Status) {
		return nil, errors.SystemError(value.Value(app.Reason))
	}

	// 获取渠道信息
	oauter, err := s.oa.GetOAutherByKeyword(ctx, tenant.Id, req.Keyword)
	if err != nil {
		return nil, errors.SystemError()
	}

	// 获取授权器
	execer, err := s.oaexec.GetOAuther(oauter)
	if err != nil {
		return nil, errors.SystemError()
	}

	// 获取token
	ti, err := execer.GetToken(ctx, &types.OAutherTokenRequest{
		IP:      ctx.ClientIP(),
		UUID:    req.UUID,
		Code:    req.Code,
		Account: req.Account,
	})
	// 获取token
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, errors.NotFoundReportOAuthError()
		}
		return nil, errors.OAuthLoginError(err.Error())
	}

	// 获取用户信息
	ui, err := execer.GetInfo(ctx, &types.OAutherInfoRequest{
		OID:   ti.OID,
		Token: ti.Token,
	})
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	// 判断是否绑定三方
	if !s.uo.IsBindUserOAuther(ctx, oauter.Id, ui.OID) {
		// 设置token，等待绑定
		if err := s.setCacheData(ctx, req.UUID, &cacheData{
			TenantId:  tenant.Id,
			AppId:     app.Id,
			OAutherId: oauter.Id,
			Token:     ti.Token,
			OID:       ti.OID,
			Expire:    ti.Expire,
		}); err != nil {
			ctx.Logger().Errorw("set token error", "err", err.Error())
			return nil, errors.OAuthLoginError()
		}
		return &types.OAutherLoginReply{NeedBind: true}, nil
	}

	// 获取用户ID信息
	uoi, err := s.uo.GetUserOAutherByCO(ctx, oauter.Id, ui.OID)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	// 获取用户信息
	user, err := s.user.GetUser(ctx, uoi.UserId)
	if err != nil {
		ctx.Logger().Errorw("msg", "get user info error", "err", err.Error())
		return nil, errors.SystemError("用户不存在")
	}
	// 判断用户状态
	if !value.Value(user.Status) {
		return nil, errors.UserDisableError(value.Value(user.Reason))
	}

	// 获取授权信息
	az, err := s.repo.GetAuthorize(ctx, &types.GetAuthorizeRequest{
		TenantId: tenant.Id,
		AppId:    app.Id,
		UserId:   user.Id,
	})
	// 如果不是私有的应用，且用户第一次登陆，则自动创建授权信息
	if gormtranserror.Is(err, gorm.ErrRecordNotFound) && !value.Value(app.Private) {
		az.Id, err = s.repo.CreateAuthorize(ctx, &entity.Authorize{
			CreateTenantModel: model.CreateTenantModel{TenantId: tenant.Id},
			UserId:            user.Id,
			AppId:             app.Id,
		})
	}
	if err != nil {
		ctx.Logger().Errorw("msg", "get authorize info error", "err", err.Error())
		return nil, errors.SystemError(err.Error())
	}

	// 更新渠道信息
	// 更新渠道token信息
	if err := s.uo.UpdateUserOAuther(ctx, &entity.UserOAuther{
		UserId:    user.Id,
		OAutherId: oauter.Id,
		OID:       ti.OID,
		Token:     ti.Token,
		ExpiredAt: ti.Expire,
		LoggedAt:  time.Now().Unix(),
	}); err != nil {
		return nil, errors.DatabaseError(err.Error())
	}

	// 判断是否需要补充资料信息
	ni, err := s.needInfo(ctx, app.Id, user)
	if err != nil {
		return nil, err
	}
	if ni {
		err := s.setCacheData(ctx, req.UUID, &cacheData{
			TenantId: tenant.Id,
			AppId:    app.Id,
			UserId:   user.Id,
		})
		if err != nil {
			return nil, errors.SystemError(err.Error())
		}
		return &types.OAutherLoginReply{NeedInfo: true}, nil
	}

	// 生成token
	token, err := s.genToken(ctx, app, az, user)
	go s.AddLoginLog(ctx.Clone(), oauter.Type, app.Id, user, err)
	if err != nil {
		return nil, err
	}
	return &types.OAutherLoginReply{Token: &token}, nil
}

func (s *Authorize) OAutherBind(ctx core.Context, req *types.OAutherBindRequest) (*types.OAutherBindReply, error) {
	// 获取uuid对应的缓存信息
	cd, err := s.getCacheData(ctx, req.UUID)
	if err != nil {
		return nil, errors.BindError("绑定信息已过期，请重新退出绑定")
	}
	if cd.AppId == 0 || cd.TenantId == 0 {
		return nil, errors.SystemError()
	}

	// 判断验证码
	cpt, err := ctx.Captcha().Get(ctx, captchaName, captcha.WithVerifiedDelete(true))
	if err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		return nil, errors.VerifyCaptchaError()
	}
	if err = cpt.VerifyCaptcha(&captcha.VerifyCaptchaRequest{
		Scene:      bindCaptchaScene,
		ClientIP:   ctx.ClientIP(),
		UUID:       req.CaptchaId,
		VerifyCode: req.Captcha,
	}); err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		return nil, errors.VerifyCaptchaError()
	}

	// 获取租户信息
	tenant, err := s.tenant.GetTenant(ctx, cd.TenantId)
	if err != nil {
		ctx.Logger().Errorw("msg", "get tenant info error", "err", err.Error())
		return nil, errors.TenantNotFoundError()
	}
	if !value.Value(tenant.Status) {
		return nil, errors.SystemError("租户已禁用")
	}

	// 获取应用信息
	app, err := s.app.GetApp(ctx, cd.AppId)
	if err != nil {
		ctx.Logger().Errorw("msg", "get app info error", "err", err.Error())
		return nil, errors.AppNotFoundError()
	}
	if !value.Value(app.Status) {
		return nil, errors.SystemError(value.Value(app.Reason))
	}

	// 获取渠道信息
	oauter, err := s.oa.GetOAuther(ctx, cd.OAutherId)
	if err != nil {
		ctx.Logger().Errorw("msg", "get oauther info error", "err", err.Error())
		return nil, errors.OAuthLoginError()
	}

	// 获取用户信息
	user, err := s.user.GetUserByTU(ctx, tenant.Id, req.Username)
	if err != nil {
		ctx.Logger().Errorw("msg", "get user info error", "err", err.Error())
		// todo 用户不存在是否进行快速注册
		return nil, errors.UsernameNotExistError()
	}

	// 获取授权信息
	az, err := s.repo.GetAuthorize(ctx, &types.GetAuthorizeRequest{
		TenantId: tenant.Id,
		AppId:    app.Id,
		UserId:   user.Id,
	})
	// 如果不是私有的应用，且用户第一次登陆，则自动创建授权信息
	if gormtranserror.Is(err, gorm.ErrRecordNotFound) {
		if !value.Value(app.Private) {
			az.Id, err = s.repo.CreateAuthorize(ctx, &entity.Authorize{
				CreateTenantModel: model.CreateTenantModel{TenantId: tenant.Id},
				UserId:            user.Id,
				AppId:             app.Id,
			})
		} else {
			return nil, errors.AppScopeError()
		}
	}
	if err != nil {
		ctx.Logger().Errorw("msg", "get authorize info error", "err", err.Error())
		return nil, errors.SystemError(err.Error())
	}

	// 判断是否需要补充资料信息
	ni, err := s.needInfo(ctx, app.Id, user)
	if err != nil {
		return nil, err
	}
	if ni {
		err := s.setCacheData(ctx, req.CaptchaId, &cacheData{
			TenantId: tenant.Id,
			AppId:    app.Id,
			UserId:   user.Id,
		})
		if err != nil {
			return nil, errors.SystemError(err.Error())
		}
		return &types.OAutherBindReply{NeedInfo: true}, nil
	}

	reply := types.OAutherBindReply{}
	// 查询三方绑定信息
	if err := ctx.Transaction(func(ctx core.Context) error {
		// 写入三方绑定信息
		if _, err := s.uo.CreateUserOAuther(ctx, &entity.UserOAuther{
			UserId:    user.Id,
			OAutherId: cd.OAutherId,
			OID:       cd.OID,
			Token:     cd.Token,
			ExpiredAt: cd.Expire,
			LoggedAt:  time.Now().Unix(),
		}); err != nil {
			return errors.SystemError(err.Error())
		}

		// 生成token
		token, err := s.genToken(ctx, app, az, user)
		go s.AddLoginLog(ctx.Clone(), oauter.Type, app.Id, user, err)

		if err != nil {
			return err
		}

		// 更新用户token信息
		reply.Token = token
		return nil
	}); err != nil {
		return nil, err
	}

	return &reply, nil
}

func (s *Authorize) GetImageCaptcha(ctx core.Context, scene string) (*types.Captcha, error) {
	getScene := func(s string) (string, error) {
		switch s {
		case "login":
			return loginCaptchaScene, nil
		case "bind":
			return bindCaptchaScene, nil
		case "register":
			return registerCaptchaScene, nil
		default:
			return "", errors.SystemError("不支持的场景")
		}
	}

	scene, err := getScene(scene)
	if err != nil {
		return nil, err
	}

	cpt, err := ctx.Captcha().Get(ctx, captchaName)
	if err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		return nil, errors.VerifyCaptchaError()
	}
	resp, err := cpt.GetCaptcha(&captcha.GetCaptchaRequest{
		Scene:    scene,
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

	return &types.Captcha{
		Uuid:    resp.UUID,
		Expire:  uint32(resp.Expire.Seconds()),
		Captcha: item.EncodeB64string(),
	}, nil
}

// Login 用户密码登陆
func (s *Authorize) Login(ctx core.Context, req *types.LoginRequest) (reply *types.LoginReply, err error) {
	// 判断验证码
	cpt, err := ctx.Captcha().Get(ctx, captchaName, captcha.WithVerifiedDelete(true))
	if err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		err = errors.VerifyCaptchaError()
		return
	}
	if err = cpt.VerifyCaptcha(&captcha.VerifyCaptchaRequest{
		Scene:      loginCaptchaScene,
		ClientIP:   ctx.ClientIP(),
		UUID:       req.CaptchaId,
		VerifyCode: req.Captcha,
	}); err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		err = errors.VerifyCaptchaError()
		return
	}

	// 获取租户信息
	tenant, err := s.tenant.GetTenantByKeyword(ctx, req.Tenant)
	if err != nil {
		ctx.Logger().Errorw("msg", "get tenant info error", "err", err.Error())
		return nil, errors.TenantNotFoundError()
	}
	if !value.Value(tenant.Status) {
		return nil, errors.SystemError("租户已禁用")
	}

	// 获取应用信息
	app, err := s.app.GetAppByKeyword(req.App)
	if err != nil {
		ctx.Logger().Errorw("msg", "get app info error", "err", err.Error())
		return nil, errors.AppNotFoundError()
	}
	if !value.Value(app.Status) {
		return nil, errors.SystemError(value.Value(app.Reason))
	}

	var user *entity.User
	if value.IsPhone(req.Username) || value.IsEmail(req.Username) {
		var ao *entity.AppOAuther
		if value.IsPhone(req.Username) {
			ao, err = s.ao.GetAppOAutherByAT(ctx, app.Id, "email")
		} else {
			ao, err = s.ao.GetAppOAutherByAT(ctx, app.Id, "phone")
		}
		if err != nil {
			ctx.Logger().Errorw("msg", "get user info error", "err", err.Error())
			err = errors.UsernameNotExistError()
			return
		}

		// 获取渠道对应的用户信息
		info, err := s.uo.GetUserOAutherByCO(ctx, ao.OAutherId, req.Username)
		if err != nil {
			ctx.Logger().Errorw("msg", "get user info error", "err", err.Error())
			return nil, errors.UsernameNotExistError()
		}

		user, err = s.user.GetUser(ctx, info.UserId)
		if err != nil {
			ctx.Logger().Errorw("msg", "get user info error", "err", err.Error())
			return nil, errors.UsernameNotExistError()
		}
	} else {
		// 获取用户信息
		user, err = s.user.GetUserByTU(ctx, tenant.Id, req.Username)
		if err != nil {
			ctx.Logger().Errorw("msg", "get user info error", "err", err.Error())
			return nil, errors.UsernameNotExistError()
		}
	}

	// 判断用户状态
	if !value.Value(user.Status) {
		return nil, errors.UserDisableError(value.Value(user.Reason))
	}

	// 判断密码
	if !crypto.CompareHashPwd(user.Password, req.Password) {
		err = errors.PasswordError()
		return
	}

	// 获取授权信息
	az, err := s.repo.GetAuthorize(ctx, &types.GetAuthorizeRequest{
		TenantId: tenant.Id,
		AppId:    app.Id,
		UserId:   user.Id,
	})
	// 如果不是私有的应用，且用户第一次登陆，则自动创建授权信息
	if gormtranserror.Is(err, gorm.ErrRecordNotFound) {
		if !value.Value(app.Private) {
			az.Id, err = s.repo.CreateAuthorize(ctx, &entity.Authorize{
				CreateTenantModel: model.CreateTenantModel{TenantId: tenant.Id},
				UserId:            user.Id,
				AppId:             app.Id,
			})
		} else {
			return nil, errors.AppScopeError()
		}
	}
	if err != nil {
		ctx.Logger().Errorw("msg", "get authorize info error", "err", err.Error())
		err = errors.SystemError()
		return
	}

	// 判断是否需要补充资料信息
	ni, err := s.needInfo(ctx, app.Id, user)
	if err != nil {
		return nil, err
	}
	if ni {
		err := s.setCacheData(ctx, req.CaptchaId, &cacheData{
			TenantId: tenant.Id,
			AppId:    app.Id,
			UserId:   user.Id,
		})
		if err != nil {
			return nil, errors.SystemError(err.Error())
		}
		return &types.LoginReply{NeedInfo: true}, nil
	}

	// 生产token
	token, err := s.genToken(ctx, app, az, user)
	go s.AddLoginLog(ctx.Clone(), "username", app.Id, user, err)

	if err != nil {
		return nil, err
	}
	return &types.LoginReply{
		Token: token,
	}, nil
}

// Register 用户密码注册
func (s *Authorize) Register(ctx core.Context, req *types.RegisterRequest) (*types.RegisterReply, error) {
	// 判断验证码
	cpt, err := ctx.Captcha().Get(ctx, captchaName, captcha.WithVerifiedDelete(true))
	if err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		return nil, errors.VerifyCaptchaError()
	}
	if err = cpt.VerifyCaptcha(&captcha.VerifyCaptchaRequest{
		Scene:      registerCaptchaScene,
		ClientIP:   ctx.ClientIP(),
		UUID:       req.CaptchaId,
		VerifyCode: req.Captcha,
	}); err != nil {
		ctx.Logger().Warnw("msg", "captcha verify error", "err", err.Error())
		return nil, errors.VerifyCaptchaError()
	}

	// 获取租户信息
	tenant, err := s.tenant.GetTenantByKeyword(ctx, req.Tenant)
	if err != nil {
		ctx.Logger().Errorw("msg", "get tenant info error", "err", err.Error())
		return nil, errors.TenantNotFoundError()
	}
	if !value.Value(tenant.Status) {
		return nil, errors.SystemError("租户已禁用")
	}

	// 获取应用信息
	app, err := s.app.GetAppByKeyword(req.App)
	if err != nil {
		ctx.Logger().Errorw("msg", "get app info error", "err", err.Error())
		return nil, errors.AppNotFoundError()
	}
	if !value.Value(app.Status) {
		return nil, errors.SystemError(value.Value(app.Reason))
	}
	if value.Value(app.Private) {
		return nil, errors.AppScopeError("该应用未开放注册")
	}

	// 判断用户是否存在
	_, err = s.user.GetUserByTU(ctx, tenant.Id, req.Username)
	if err == nil {
		return nil, errors.RegisterError("用户已存在")
	}

	// 写入用户信息
	setting := tenant.GetSetting()
	user := &entity.User{
		BaseTenantModel: model.BaseTenantModel{TenantId: tenant.Id},
		Username:        req.Username,
		Password:        crypto.EncodePwd(req.Password),
		Status:          value.Pointer(true),
		Nickname:        setting.DefaultUserNickname,
		Avatar:          setting.DefaultUserAvatar,
		LoggedAt:        time.Now().Unix(),
	}
	user.Id, err = s.user.CreateUser(ctx, user)
	if err != nil {
		ctx.Logger().Errorw("msg", "create user error", "err", err.Error())
		return nil, errors.SystemError()
	}

	// 获取授权信息
	az, err := s.repo.GetAuthorize(ctx, &types.GetAuthorizeRequest{
		TenantId: tenant.Id,
		AppId:    app.Id,
		UserId:   user.Id,
	})
	// 如果不是私有的应用，且用户第一次登陆，则自动创建授权信息
	if gormtranserror.Is(err, gorm.ErrRecordNotFound) {
		if !value.Value(app.Private) {
			az.Id, err = s.repo.CreateAuthorize(ctx, &entity.Authorize{
				CreateTenantModel: model.CreateTenantModel{TenantId: tenant.Id},
				UserId:            user.Id,
				AppId:             app.Id,
			})
		} else {
			return nil, errors.AppScopeError()
		}
	}
	if err != nil {
		ctx.Logger().Errorw("msg", "get authorize info error", "err", err.Error())
		return nil, errors.SystemError()
	}

	// 判断是否需要补充资料信息
	ni, err := s.needInfo(ctx, app.Id, user)
	if err != nil {
		return nil, err
	}
	if ni {
		err := s.setCacheData(ctx, req.CaptchaId, &cacheData{
			TenantId: tenant.Id,
			AppId:    app.Id,
			UserId:   user.Id,
		})
		if err != nil {
			return nil, errors.SystemError(err.Error())
		}
		return &types.RegisterReply{NeedInfo: true}, nil
	}

	// 生产token
	token, err := s.genToken(ctx, app, az, user)
	go s.AddLoginLog(ctx.Clone(), "username", app.Id, user, err)

	if err != nil {
		return nil, err
	}
	return &types.RegisterReply{
		Token: token,
	}, nil
}

func (s *Authorize) needInfo(ctx core.Context, appid uint32, user *entity.User) (bool, error) {
	return false, nil
	var im map[string]struct{}
	for _, v := range user.Infos {
		im[v.Field] = struct{}{}
	}

	// 获取租户层级下的必填字段
	list, _, err := s.field.ListField(ctx, &types.ListFieldRequest{
		Required: value.Pointer(true),
	})
	if err != nil {
		return false, errors.SystemError(err.Error())
	}
	for _, v := range list {
		_, ok := im[v.Keyword]
		if !ok {
			return true, nil
		}
	}

	// 检查app必填写字段
	afs, _, err := s.appfield.ListAppField(ctx, &types.ListAppFieldRequest{
		AppId:    appid,
		Required: value.Pointer(true),
	})
	if err != nil {
		return false, errors.SystemError(err.Error())
	}
	for _, v := range afs {
		_, ok := im[v.Field.Keyword]
		if !ok {
			return true, nil
		}
	}

	return false, nil
}

func (s *Authorize) setCacheData(ctx core.Context, uuid string, data *cacheData) error {
	b, _ := json.Marshal(data)
	key := fmt.Sprintf("%s:%s", oautherCache, uuid)
	return ctx.Redis().Set(ctx, key, string(b), 10*time.Minute).Err()
}

func (s *Authorize) getCacheData(ctx core.Context, uuid string) (*cacheData, error) {
	cd := cacheData{}
	key := fmt.Sprintf("%s:%s", oautherCache, uuid)
	if err := json.Unmarshal([]byte(ctx.Redis().Get(ctx, key).Val()), &cd); err != nil {
		return nil, errors.SystemError(err.Error())
	}
	return &cd, nil
}

func (s *Authorize) genToken(ctx core.Context, app *entity.App, az *entity.Authorize, user *entity.User) (token string, err error) {
	if user.Status != nil && !*user.Status {
		err = errors.UserDisableError()
		return
	}

	//infos, err := s.GetUserinfo(ctx, app.Id, user.Id)
	//if err != nil {
	//	return "", errors.SystemError(err.Error())
	//}

	//userdepts, err := s.GetUserDepts(ctx, user.Id)
	//if err != nil {
	//	return "", errors.SystemError(err.Error())
	//}

	jc := app.GetSetting().JWT
	jwter := jwt.New(ctx.Context, &jwt.Config{
		Secret:  jc.Secret,
		Expire:  time.Duration(jc.Expire) * time.Second,
		Renewal: time.Duration(jc.Renewal) * time.Second,
		App:     app.Keyword,
	})

	//var mainDeptId uint32
	//for _, item := range userdepts {
	//	if value.Value(item.Main) {
	//		mainDeptId = item.DeptId
	//	}
	//}

	info := types.AuthorizeInfo{
		AppId:    app.Id,
		TenantId: user.TenantId,
		UserId:   user.Id,
		DeptId:   s.userdept.GetUserMainDeptId(user.Id),
	}
	token, err = jwter.NewToken(info.ToMap())
	if err != nil {
		ctx.Logger().Errorw("msg", "gen user token error", "err", err.Error())
		err = errors.GenTokenError()
		return
	}

	// 开启唯一设备登陆
	if err := jwter.SetUnique(ctx.ClientIP(), token); err != nil {
		return "", errors.SystemError(err.Error())
	}

	// 开启唯一平台登陆
	if err := jwter.SetUnique(ctx.UserAgent().String, token); err != nil {
		return "", errors.SystemError(err.Error())
	}

	// 更新应用授权token信息
	expired := time.Now().Unix() + jc.Expire

	tm := crypto.MD5([]byte(token))
	tokens := az.GetTokens()
	tokens[tm] = expired

	data := &entity.Authorize{
		CreateTenantModel: model.CreateTenantModel{Id: az.Id},
		Tokens:            value.ObjToString(tokens),
		LoggedAt:          time.Now().Unix(),
		ExpiredAt:         expired,
	}
	if err = s.repo.UpdateAuthorize(ctx, data); err != nil {
		ctx.Logger().Errorw("msg", "update user login info error", "err", err.Error())
		err = errors.SystemError()
		return
	}

	// 更新用户登陆时间
	if err = s.user.UpdateUser(ctx, &entity.User{
		BaseTenantModel: model.BaseTenantModel{Id: user.Id},
		LoggedAt:        time.Now().Unix(),
	}); err != nil {
		ctx.Logger().Errorw("msg", "update user login info error", "err", err.Error())
		err = errors.SystemError()
		return
	}

	return token, nil
}

func (s *Authorize) AddLoginLog(ctx core.Context, tp string, appId uint32, user *entity.User, err error) {
	if errors.IsUserDisableError(err) ||
		errors.IsVerifyCaptchaError(err) ||
		errors.IsPasswordError(err) ||
		errors.IsTenantNotFoundError(err) {
		return
	}

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

	_, _ = s.log.CreateLoginLog(ctx, &entity.LoginLog{
		AppId: appId,
		CreateTenantUserModel: model.CreateTenantUserModel{
			UserId:   user.Id,
			DeptId:   s.userdept.GetUserMainDeptId(user.Id),
			TenantId: user.TenantId,
		},
		Type:        tp,
		IP:          ctx.ClientIP(),
		Address:     s.address.GetAddressByIP(ip),
		Browser:     ug.Name + " " + ug.Version,
		Device:      ug.OS + " " + ug.OSVersion,
		Code:        code,
		Description: desc,
	})
}

func (s *Authorize) AddAuthLog(ctx core.Context, appId uint32, menuId uint32) {
	_, _ = s.log.CreateAuthLog(ctx, &entity.AuthLog{
		AppId: appId,
		CreateTenantUserModel: model.CreateTenantUserModel{
			UserId:   ctx.Auth().UserId,
			DeptId:   ctx.Auth().DeptId,
			TenantId: ctx.Auth().TenantId,
		},
		MenuId: menuId,
	})
}

func (u *Authorize) GetUserinfo(ctx core.Context, appId uint32, userId uint32) (map[string]any, error) {
	getFieldReq := &types.ListUserinfoRequest{
		UserId: userId,
	}

	// 获取用户的字段
	fields, _ := u.appfield.GetAppFields(ctx, appId)
	getFieldReq.Fields = fields

	// 获取用户应用信息
	infos, _ := u.info.ListUserinfo(ctx, &types.ListUserinfoRequest{
		UserId:   userId,
		TenantId: ctx.Auth().TenantId,
		Fields:   fields,
	})

	var res map[string]any
	for _, info := range infos {
		res[info.Field] = info.Value
	}
	return res, nil
}

func (u *Authorize) GetUserDepts(ctx core.Context, userId uint32) ([]*entity.UserDept, error) {
	// 获取用户部门信息
	depts, _, err := u.userdept.ListUserDept(ctx, &types.ListUserDeptRequest{
		UserId: userId,
	})
	return depts, err
}

// CheckAuth api鉴权
func (u *Authorize) CheckAuth(ctx core.Context, req *types.CheckAuthRequest) (*types.AuthorizeInfo, error) {
	validate := func(method string, path string) error {
		// 如果是管理员则默认获取全部权限
		if u.tad.IsAdmin(ctx.Auth().TenantId, ctx.Auth().UserId) {
			return nil
		}

		// 获取鉴权器，判断是否为白名单
		menu, has := u.menu.GetCacheMenu(req.Path, req.Method)
		if !has {
			return errors.ForbiddenError()
		}
		// 如果是基础api，则放行
		if menu.Type == entity.MenuTypeBasic {
			return nil
		}

		// 判断是否拥有api权限
		if !u.scope.HasMenuScope(ctx, menu.Id) {
			return errors.ForbiddenError(menu.Name)
		}

		go u.AddAuthLog(ctx, menu.AppId, menu.Id)

		return nil
	}

	key := fmt.Sprintf("%s:%s", req.Path, req.Method)
	if ctx.Config().JWT.Whitelist[key] {
		return nil, nil
	}

	// 验证接口
	if err := validate(req.Method, req.Path); err != nil {
		return nil, err
	}

	// 获取用户token信息
	return &types.AuthorizeInfo{
		UserId:   ctx.Auth().UserId,
		TenantId: ctx.Auth().TenantId,
		DeptId:   ctx.Auth().DeptId,
	}, nil
}
