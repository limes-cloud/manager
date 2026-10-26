package authorize

import (
	"context"
	"encoding/json"
	"strings"

	md "github.com/go-kratos/kratos/v2/metadata"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"

	"github.com/limes-cloud/manager/api/errors"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/limes-cloud/kratosx"
)

const (
	TokenKey = "x-md-global-token"
	AppKey   = kratosx.AppKeyMetadataKey

	InfoKey = "x-md-global-info"

	// ClientIPKey、UserAgentKey 用于下游服务把原始请求身份透传给 manager，
	// 避免 manager 在 ParseToken 时误用内部 gRPC 调用的 IP/UA 做唯一性校验。
	ClientIPKey  = "x-md-global-client-ip"
	UserAgentKey = "x-md-global-user-agent"
)

type authInfo struct{}

// AuthMiddleware 鉴权
func AuthMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(c context.Context, req any) (any, error) {
			// 从header获取info,存在则认为已经通过鉴权体系
			header, ok := transport.FromServerContext(c)
			if !ok {
				return nil, errors.SystemError()
			}

			info := header.RequestHeader().Get(InfoKey)
			if info != "" {
				reply := CheckAuthReply{}
				if err := json.Unmarshal([]byte(info), &reply); err != nil {
					return nil, errors.SystemError()
				}
				cctx := context.WithValue(c, authInfo{}, &reply)
				return handler(cctx, req)
			}

			token := header.RequestHeader().Get(TokenKey)
			if token == "" {
				token = header.RequestHeader().Get("Authorization")
				token = strings.TrimPrefix(token, "Bearer ")
				if token != "" {
					c = md.AppendToClientContext(c, TokenKey, token)
				}
			}
			if token == "" {
				return handler(c, req)
			}

			ctx := kratosx.MustContext(c)
			appKey := strings.TrimSpace(ctx.AppKey())
			if appKey == "" {
				return nil, kratosJwt.ErrTokenInvalid
			}
			ctx.SetAppKey(appKey)
			ctx.SetMetadata(ClientIPKey, ctx.ClientIP())
			ctx.SetMetadata(UserAgentKey, ctx.UserAgent().String)
			c = ctx.Ctx()

			// 解析token
			conn, err := ctx.GrpcConn("Manager")
			if err != nil {
				return nil, err
			}

			client := NewAuthorizeClient(conn)
			reply, err := client.ParseToken(ctx, &ParseTokenRequest{})
			if err != nil {
				return nil, err
			}
			// ParseTokenReply 与 CheckAuthReply 是不同 proto 类型，Get() 按 CheckAuthReply 断言，
			// 必须显式转换，否则下游 Auth() 永远取不到用户信息（表现为 NotLoginError 403）。
			cctx := context.WithValue(ctx.Ctx(), authInfo{}, &CheckAuthReply{
				UserId: reply.UserId,
				DeptId: reply.DeptId,
			})
			return handler(cctx, req)
		}
	}
}

func Get(ctx context.Context) *CheckAuthReply {
	v, _ := ctx.Value(authInfo{}).(*CheckAuthReply)
	return v
}
