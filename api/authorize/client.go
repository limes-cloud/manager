package authorize

import (
	"context"
	"encoding/json"
	"strings"

	md "github.com/go-kratos/kratos/v2/metadata"

	"github.com/limes-cloud/manager/api/errors"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/limes-cloud/kratosx"
)

const (
	TokenKey = "x-md-global-token"

	InfoKey = "x-md-global-info"
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

			// 解析token
			ctx := kratosx.MustContext(c)
			conn, err := ctx.GrpcConn("Manager")
			if err != nil {
				return nil, err
			}

			client := NewAuthorizeClient(conn)
			reply, err := client.ParseToken(ctx, &ParseTokenRequest{})
			if err != nil {
				return nil, err
			}
			cctx := context.WithValue(ctx.Ctx(), authInfo{}, reply)
			return handler(cctx, req)
		}
	}
}

func Get(ctx context.Context) *CheckAuthReply {
	v, _ := ctx.Value(authInfo{}).(*CheckAuthReply)
	return v
}
