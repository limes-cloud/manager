package app

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/internal/core"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/manager/api/auth"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/apis"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type Auth struct {
	auth.UnimplementedAuthServer
	srv *service.Auth
}

func NewAuth() *Auth {
	return &Auth{
		srv: service.NewAuth(
			dbs.NewAuth(),
			dbs.NewUser(),
			apis.NewAddress(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewAuth()
		auth.RegisterAuthHTTPServer(hs, srv)
		auth.RegisterAuthServer(gs, srv)
	})
}

// Auth 接口鉴权
func (a *Auth) Auth(c context.Context, req *auth.AuthRequest) (*auth.AuthReply, error) {
	//res, err := a.srv.Auth(core.MustContext(c), &types.AuthRequest{
	//	Path:   req.Path,
	//	Method: req.Method,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//if res == nil {
	//	return &auth.AuthReply{}, nil
	//}

	return &auth.AuthReply{
		// UserId:        res.UserId,
		// Username:      res.UserName,
		// RoleIds:       res.RoleIds,
		// DepartmentIds: res.DepartmentIds,
		// JobIds:        res.JobIds,
	}, nil
}

func (a *Auth) GetUserLoginCaptcha(c context.Context, _ *empty.Empty) (*auth.GetUserLoginCaptchaReply, error) {
	reply, err := a.srv.GetUserLoginCaptcha(core.MustContext(c))
	if err != nil {
		return nil, err
	}
	return &auth.GetUserLoginCaptchaReply{
		Uuid:    reply.Uuid,
		Captcha: reply.Captcha,
		Expire:  reply.Expire,
	}, nil
}

func (a *Auth) UserLogin(c context.Context, req *auth.UserLoginRequest) (*auth.UserLoginReply, error) {
	var (
		in  types.UserLoginRequest
		ctx = core.MustContext(c)
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "request transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	token, err := a.srv.UserLogin(ctx, &in)
	if err != nil {
		return nil, err
	}
	return &auth.UserLoginReply{Token: token}, nil
}

func (a *Auth) UserLogout(c context.Context, _ *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, a.srv.UserLogout(core.MustContext(c))
}

func (a *Auth) UserRefreshToken(c context.Context, _ *empty.Empty) (*auth.UserRefreshTokenReply, error) {
	token, err := a.srv.UserRefreshToken(core.MustContext(c))
	if err != nil {
		return nil, err
	}
	return &auth.UserRefreshTokenReply{Token: token}, nil
}

func (a *Auth) ListLoginLog(c context.Context, req *auth.ListLoginLogRequest) (*auth.ListLoginLogReply, error) {
	var (
		in  types.ListLoginLogRequest
		ctx = core.MustContext(c)
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "request transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	list, total, err := a.srv.ListLoginLog(ctx, &in)
	if err != nil {
		return nil, err
	}
	reply := &auth.ListLoginLogReply{Total: total}

	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "resp transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return reply, nil
}

// OAuthHandler 处理授权
func (a *Auth) OAuthHandler(c context.Context, req *auth.OAuthHandlerRequest) (*auth.OAuthHandlerReply, error) {
	resp, err := a.srv.OAuthHandler(core.MustContext(c), &types.OAuthHandlerRequest{
		Keyword: req.Keyword,
		User:    req.User,
	})
	if err != nil {
		return nil, err
	}
	return &auth.OAuthHandlerReply{
		Uuid:      resp.UUID,
		Action:    resp.Action,
		Value:     resp.Value,
		Tip:       resp.Tip,
		CodeField: resp.CodeField,
		Keyword:   req.Keyword,
	}, nil
}

// ReportOAuthCode 上报登陆渠道信息
func (a *Auth) ReportOAuthCode(c context.Context, req *auth.ReportOAuthCodeRequest) (*auth.ReportOAuthCodeReply, error) {
	err := a.srv.ReportOAuthCode(core.MustContext(c), &types.ReportOAuthCodeRequest{
		Code:    req.Code,
		Keyword: req.Keyword,
		UUID:    req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	return &auth.ReportOAuthCodeReply{}, nil
}

func (a *Auth) OAuthLogin(c context.Context, req *auth.OAuthLoginRequest) (*auth.OAuthLoginReply, error) {
	resp, err := a.srv.OAuthLogin(core.MustContext(c), &types.OAuthLoginRequest{
		Code:    req.Code,
		User:    req.User,
		Keyword: req.Keyword,
		UUID:    req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	return &auth.OAuthLoginReply{
		IsBind: resp.IsBind,
		Token:  resp.Token,
	}, nil
}

func (a *Auth) OAuthBind(c context.Context, req *auth.OAuthBindRequest) (*auth.OAuthBindReply, error) {
	token, err := a.srv.OAuthBind(core.MustContext(c), &types.OAuthBindRequest{
		UserLoginRequest: &types.UserLoginRequest{
			Username:  req.Username,
			Password:  req.Password,
			CaptchaId: req.CaptchaId,
			Captcha:   req.Captcha,
		},
		Keyword: req.Keyword,
		UUID:    req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	return &auth.OAuthBindReply{
		Token: token,
	}, nil
}

func (a *Auth) ListAuthLog(c context.Context, req *auth.ListAuthLogRequest) (*auth.ListAuthLogReply, error) {
	var (
		in  types.ListAuthLogRequest
		ctx = core.MustContext(c)
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "request transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	list, total, err := a.srv.ListAuthLog(ctx, &in)
	if err != nil {
		return nil, err
	}
	reply := &auth.ListAuthLogReply{Total: total}

	if err := value.Transform(list, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "resp transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return reply, nil
}
