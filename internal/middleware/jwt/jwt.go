package jwt

import (
	"context"
	"encoding/base64"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/manager/api/authorize"

	"github.com/limes-cloud/manager/internal/app"

	"github.com/limes-cloud/kratosx/pkg/crypto"

	"github.com/limes-cloud/manager/internal/infra/dbs"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/limes-cloud/manager/api/errors"
	pkgtjwt "github.com/limes-cloud/manager/internal/pkg/jwt"
)

const (
	TokenKey = "x-md-global-token"
)

type jwterKey struct{}

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
			token := md.Get(TokenKey)
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
			md.Set(TokenKey, token)

			// 判断token格式
			if strings.Count(token, ".") != 3 {
				return nil, kratosJwt.ErrTokenInvalid
			}

			// 获取app key
			appEncode := token[strings.LastIndex(token, ".")+1:]
			token = token[:strings.LastIndex(token, ".")]

			// 切分最后一个.
			appByte, _ := base64.StdEncoding.DecodeString(appEncode)

			// 获取app信息
			data, err := appIns.GetAppByKeyword(string(appByte))
			if err != nil {
				return nil, errors.SystemError()
			}
			jc := data.GetSetting().JWT

			// 初始化app jtw
			jwter := pkgtjwt.New(kratosx.MustContext(ctx), &pkgtjwt.Config{
				Secret:         jc.Secret,
				Expire:         time.Duration(jc.Expire) * time.Second,
				Renewal:        time.Duration(jc.Renewal) * time.Second,
				App:            data.Keyword,
				UniqueDevice:   jc.UniqueDevice,
				UniquePlatform: jc.UniquePlatform,
			})

			// 设置jwter到ctx
			ctx = context.WithValue(ctx, jwterKey{}, jwter)

			return handler(ctx, req)
		}
	}
}

// TokenValidate jwt校验
func TokenValidate() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			// 设置token到metadata
			md, ok := metadata.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			token := md.Get(TokenKey)
			if token == "" {
				return handler(ctx, req)
			}

			jwter := ctx.Value(jwterKey{}).(pkgtjwt.JWT)

			// 设置token信息到ctx
			info := &types.AuthorizeInfo{}
			if err := jwter.ParseByToken(token, info); err != nil {
				return nil, err
			}

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

			// 是否开启了唯一设备校验
			if jwter.Config().UniqueDevice {
				if !jwter.CompareUniqueToken(kctx.ClientIP(), token) {
					return nil, kratosJwt.ErrTokenInvalid
				}
			}

			// 是否开启了唯一平台校验
			if jwter.Config().UniquePlatform {
				if !jwter.CompareUniqueToken(kctx.UserAgent().String, token) {
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
