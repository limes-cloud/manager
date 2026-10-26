package jwt

import (
	"context"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"

	"github.com/limes-cloud/manager/api/authorize"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/app"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	pkgtjwt "github.com/limes-cloud/manager/internal/pkg/jwt"
	"github.com/limes-cloud/manager/internal/types"
)

const appHeaderKey = kratosx.AppKeyHeaderKey

// Token 获取并设置token
func Token() middleware.Middleware {
	appIns := dbs.NewApp()
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			// 设置token到metadata
			md, ok := metadata.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}
			token := md.Get(core.TokenKey)
			// md中没有，则从header获取
			if token == "" {
				// 从header获取token
				header, ok := transport.FromServerContext(ctx)
				if !ok {
					return handler(ctx, req)
				}
				token = header.RequestHeader().Get("Authorization")
			}

			token = strings.TrimPrefix(token, "Bearer ")
			if token == "" {
				return handler(ctx, req)
			}
			md.Set(core.TokenKey, token)

			appKey := md.Get(core.AppKey)
			if appKey == "" {
				header, ok := transport.FromServerContext(ctx)
				if !ok {
					return nil, kratosJwt.ErrTokenInvalid
				}
				appKey = strings.TrimSpace(header.RequestHeader().Get(appHeaderKey))
				if appKey == "" {
					return nil, kratosJwt.ErrTokenInvalid
				}
				md.Set(core.AppKey, appKey)
			}
			kratosx.MustContext(ctx).SetAppKey(appKey)

			// 获取app信息
			data, err := appIns.GetAppByKeyword(appKey)
			if err != nil {
				return nil, errors.SystemError()
			}

			// 配置解析失败会导致 JWT 密钥为空，必须记录而非静默降级
			setting, err := data.ParseSetting()
			if err != nil {
				kratosx.MustContext(ctx).Logger().Errorw("msg", "parse app setting error", "err", err.Error(), "app", appKey)
				return nil, errors.SystemError()
			}
			jc := setting.JWT

			// 初始化app jwt
			jwter := pkgtjwt.New(kratosx.MustContext(ctx), &pkgtjwt.Config{
				Secret:         jc.Secret,
				Expire:         time.Duration(jc.Expire) * time.Second,
				Renewal:        time.Duration(jc.Renewal) * time.Second,
				UniqueDevice:   jc.UniqueDevice,
				UniquePlatform: jc.UniquePlatform,
			})

			// 设置jwter到ctx
			ctx = context.WithValue(ctx, core.JwterKey, jwter)

			return handler(ctx, req)
		}
	}
}

func getPathInfo(ctx context.Context) (string, string) {
	// 设置路径信息
	path, method := "", ""
	if tr, ok := transport.FromServerContext(ctx); ok {
		path = tr.Operation()
	}
	h, is := http.RequestFromServerContext(ctx)
	if is {
		path = h.URL.Path
		method = h.Method
	} else {
		method = "GRPC"
	}
	return path, method
}

// TokenValidate jwt校验
func TokenValidate() middleware.Middleware {
	userDeptIns := dbs.NewUserDept()
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			// 设置token到metadata
			md, ok := metadata.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			token := md.Get(core.TokenKey)

			// 判断是否在白名单内
			path, method := getPathInfo(ctx)
			pm := path + ":" + method
			cctx := core.MustContext(ctx)
			if cctx.Config().JWT.Whitelist[pm] {
				return handler(ctx, req)
			}

			if token == "" {
				return nil, errors.NotLoginError()
			}
			jwter := ctx.Value(core.JwterKey).(pkgtjwt.JWT)

			// 设置token信息到ctx
			info := &types.AuthorizeInfo{}
			if err := jwter.ParseByToken(token, info); err != nil {
				return nil, err
			}
			if info.AuthorizeBase == nil {
				info.AuthorizeBase = &types.AuthorizeBase{}
			}
			info.AppKey = cctx.AppKey()
			info.DeptId = userDeptIns.GetUserMainDeptId(cctx, info.AppKey, info.UserId)
			info.JobId = userDeptIns.GetUserMainJobId(cctx, info.AppKey, info.UserId)

			// 判断jwt是否过期
			if info.Exp < time.Now().Unix() {
				return nil, kratosJwt.ErrTokenExpired
			}

			// 写入用户信息
			ctx = context.WithValue(ctx, types.InfoKey, info)

			// 判断token是否在黑名单内
			tm := crypto.MD5([]byte(token))
			if jwter.IsBlacklist(tm) {
				return nil, kratosJwt.ErrTokenInvalid
			}

			kctx := kratosx.MustContext(ctx)
			clientIP := pkgtjwt.NormalizeClientIP(kctx.ClientIP())
			if forwarded := strings.TrimSpace(kctx.GetMetadata(authorize.ClientIPKey)); forwarded != "" {
				clientIP = pkgtjwt.NormalizeClientIP(forwarded)
			}
			userAgent := kctx.UserAgent().String
			if forwarded := kctx.GetMetadata(authorize.UserAgentKey); forwarded != "" {
				userAgent = forwarded
			}

			// 是否开启了唯一设备校验
			if jwter.Config().UniqueDevice {
				if !jwter.CompareUniqueToken(clientIP, token) {
					return nil, kratosJwt.ErrTokenInvalid
				}
			}

			// 是否开启了唯一平台校验
			if jwter.Config().UniquePlatform {
				if !jwter.CompareUniqueToken(userAgent, token) {
					return nil, kratosJwt.ErrTokenInvalid
				}
			}

			return handler(ctx, req)
		}
	}
}

// ApiValidate api校验
func ApiValidate() middleware.Middleware {
	author := app.NewAuthorize()
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			path, method := getPathInfo(ctx)
			if method == "GRPC" {
				return handler(ctx, req)
			}
			_, err := author.CheckAuth(ctx, &authorize.CheckAuthRequest{
				Path:   path,
				Method: method,
			})
			if err != nil {
				return nil, err
			}

			return handler(ctx, req)
		}
	}
}
