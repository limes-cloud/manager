package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/limes-cloud/manager/api/manager/auth/v1"
	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/apis"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/infra/oauth"
	"github.com/limes-cloud/manager/internal/types"
)

type Auth struct {
	pb.UnimplementedAuthServer
	srv *service.Auth
}

func NewAuth() *Auth {
	return &Auth{
		srv: service.NewAuth(
			oauth.New(),
			dbs.NewChannel(),
			dbs.NewApp(),
			dbs.NewUser(),
			apis.NewAddress(),
			dbs.NewRole(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewAuth()
		pb.RegisterAuthHTTPServer(hs, srv)
		pb.RegisterAuthServer(gs, srv)
	})
}

// Auth 接口鉴权
func (auth *Auth) Auth(c context.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	res, err := auth.srv.Auth(kratosx.MustContext(c), &types.AuthRequest{
		Path:   req.Path,
		Method: req.Method,
	})
	if err != nil {
		return nil, err
	}
	if res == nil {
		return &pb.AuthReply{}, nil
	}

	return &pb.AuthReply{
		UserId:        res.UserId,
		UserName:      res.UserName,
		RoleIds:       res.RoleIds,
		DepartmentIds: res.DepartmentIds,
		JobIds:        res.JobIds,
	}, nil
}

// OAuthWay 获取渠道授权方式
func (auth *Auth) OAuthWay(c context.Context, req *pb.OAuthWayRequest) (*pb.OAuthWayReply, error) {
	resp, err := auth.srv.OAuthOAuthWay(kratosx.MustContext(c), &types.OAuthWayRequest{
		Keyword: req.Keyword,
		User:    req.User,
	})
	if err != nil {
		return nil, err
	}
	return &pb.OAuthWayReply{
		Uuid:      resp.UUID,
		Action:    resp.Action,
		Value:     resp.Value,
		Tip:       resp.Tip,
		CodeField: resp.CodeField,
		Keyword:   req.Keyword,
	}, nil
}

// ReportOAuthCode 上报登陆渠道信息
func (auth *Auth) ReportOAuthCode(c context.Context, req *pb.ReportOAuthCodeRequest) (*pb.ReportOAuthCodeReply, error) {
	err := auth.srv.ReportOAuthCode(kratosx.MustContext(c), &types.ReportOAuthCodeRequest{
		Code:    req.Code,
		Keyword: req.Keyword,
		UUID:    req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ReportOAuthCodeReply{}, nil
}

func (auth *Auth) OAuthLogin(c context.Context, req *pb.OAuthLoginRequest) (*pb.OAuthLoginReply, error) {
	resp, err := auth.srv.OAuthLogin(kratosx.MustContext(c), &types.OAuthLoginRequest{
		Code:    req.Code,
		User:    req.User,
		Keyword: req.Keyword,
		UUID:    req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	return &pb.OAuthLoginReply{
		IsBind: resp.IsBind,
		Token:  resp.Token,
	}, nil
}

func (auth *Auth) OAuthBind(c context.Context, req *pb.OAuthBindRequest) (*pb.OAuthBindReply, error) {
	token, err := auth.srv.OAuthBind(kratosx.MustContext(c), &types.OAuthBindRequest{
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
	return &pb.OAuthBindReply{
		Token: token,
	}, nil
}

func (auth *Auth) UserLogin(c context.Context, req *pb.UserLoginRequest) (*pb.UserLoginReply, error) {
	var (
		in  types.UserLoginRequest
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "request transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	token, err := auth.srv.UserLogin(ctx, &in)
	if err != nil {
		return nil, err
	}
	return &pb.UserLoginReply{Token: token}, nil
}

func (auth *Auth) UserLogout(c context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, auth.srv.UserLogout(kratosx.MustContext(c))
}

func (auth *Auth) UserRefreshToken(c context.Context, _ *emptypb.Empty) (*pb.UserRefreshTokenReply, error) {
	token, err := auth.srv.UserRefreshToken(kratosx.MustContext(c))
	if err != nil {
		return nil, err
	}
	return &pb.UserRefreshTokenReply{Token: token}, nil
}

func (auth *Auth) GetUserLoginCaptcha(c context.Context, _ *emptypb.Empty) (*pb.GetUserLoginCaptchaReply, error) {
	reply, err := auth.srv.GetUserLoginCaptcha(kratosx.MustContext(c))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserLoginCaptchaReply{
		Uuid:    reply.Uuid,
		Captcha: reply.Captcha,
		Expire:  reply.Expire,
	}, nil
}

func (auth *Auth) ListLoginLog(c context.Context, req *pb.ListLoginLogRequest) (*pb.ListLoginLogReply, error) {
	list, total, err := auth.srv.ListLoginLog(kratosx.MustContext(c), &types.ListLoginLogRequest{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Username:   req.Username,
		CreatedAts: req.CreatedAts,
	})
	if err != nil {
		return nil, err
	}
	reply := &pb.ListLoginLogReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListLoginLogReply_Log{
			Username:    item.Username,
			Type:        item.Type,
			Ip:          item.IP,
			Address:     item.Address,
			Browser:     item.Browser,
			Device:      item.Device,
			Code:        int32(item.Code),
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
		})
	}

	return reply, nil
}
