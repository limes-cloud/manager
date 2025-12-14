package app

import (
	"context"

	"github.com/limes-cloud/manager/internal/infra/apis"

	"github.com/limes-cloud/manager/internal/infra/oauther"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/manager/internal/infra/dbs"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/manager/api/authorize"
	"github.com/limes-cloud/manager/internal/domain/service"
)

type Authorize struct {
	authorize.UnimplementedAuthorizeServer
	srv *service.Authorize
}

func NewAuthorize() *Authorize {
	return &Authorize{
		srv: service.NewAuthorize(
			dbs.NewAuthorize(),
			dbs.NewScope(),
			dbs.NewTenant(),
			dbs.NewApp(),
			dbs.NewUser(),
			dbs.NewAppOAuther(),
			dbs.NewOAuther(),
			oauther.New(),
			dbs.NewUserOAuther(),
			dbs.NewField(),
			dbs.NewAppField(),
			dbs.NewLog(),
			apis.NewAddress(),
			dbs.NewUserinfo(),
			dbs.NewUserDept(),
			dbs.NewTenantAdmin(),
			dbs.NewMenu(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewAuthorize()
		authorize.RegisterAuthorizeHTTPServer(hs, srv)
		authorize.RegisterAuthorizeServer(gs, srv)
	})
}

// ListOAuther 获取登陆渠道信息
func (az *Authorize) ListOAuther(c context.Context, req *authorize.ListOAutherRequest) (*authorize.ListOAutherReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())

	// 调用服务
	list, err := az.srv.ListOAuther(ctx, req.Tenant, req.App)
	if err != nil {
		return nil, err
	}
	reply := authorize.ListOAutherReply{}
	for _, item := range list {
		reply.List = append(reply.List, &authorize.ListOAutherReply_OAuther{
			Id:      item.OAuther.Id,
			Logo:    item.OAuther.Logo,
			Name:    item.OAuther.Name,
			Keyword: item.OAuther.Keyword,
			Type:    item.OAuther.Type,
		})
	}

	return &reply, nil
}

func (az *Authorize) OAutherHandle(c context.Context, req *authorize.OAutherHandleRequest) (*authorize.OAutherHandleReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	resp, err := az.srv.OAutherHandle(ctx, &types.OAutherHandleRequest{
		Tenant:  req.Tenant,
		App:     req.App,
		Keyword: req.Keyword,
		Account: req.Account,
	})
	if err != nil {
		return nil, err
	}
	return &authorize.OAutherHandleReply{
		Uuid:    resp.UUID,
		Action:  resp.Action,
		Value:   resp.Value,
		Tip:     resp.Tip,
		Keyword: req.Keyword,
	}, nil
}

func (az *Authorize) OAutherLogin(c context.Context, req *authorize.OAutherLoginRequest) (*authorize.OAutherLoginReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	reply, err := az.srv.OAutherLogin(ctx, &types.OAutherLoginRequest{
		Account: req.Account,
		Code:    req.Code,
		UUID:    req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	return &authorize.OAutherLoginReply{
		NeedBind: reply.NeedBind,
		NeedInfo: reply.NeedInfo,
		Token:    reply.Token,
	}, nil
}

func (az *Authorize) OAutherReport(c context.Context, req *authorize.OAutherReportRequest) (*authorize.OAutherReportReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	err := az.srv.OAutherReport(ctx, &types.OAutherReportRequest{
		Code: req.Code,
		UUID: req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	return &authorize.OAutherReportReply{}, nil
}

func (az *Authorize) OAutherBind(c context.Context, req *authorize.OAutherBindRequest) (*authorize.OAutherBindReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	reply, err := az.srv.OAutherBind(ctx, &types.OAutherBindRequest{
		Username:  req.Username,
		Password:  req.Password,
		CaptchaId: req.CaptchaId,
		Captcha:   req.Captcha,
		UUID:      req.Uuid,
		Register:  req.Register,
	})
	if err != nil {
		return nil, err
	}
	return &authorize.OAutherBindReply{
		NeedInfo: reply.NeedInfo,
		Token:    reply.Token,
	}, nil
}

func (az *Authorize) GetImageCaptcha(c context.Context, req *authorize.GetImageCaptchaRequest) (*authorize.GetImageCaptchaReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	reply, err := az.srv.GetImageCaptcha(ctx, req.Scene)
	if err != nil {
		return nil, err
	}
	return &authorize.GetImageCaptchaReply{
		Uuid:    reply.Uuid,
		Captcha: reply.Captcha,
		Expire:  reply.Expire,
	}, nil
}

func (az *Authorize) Login(c context.Context, req *authorize.LoginRequest) (*authorize.LoginReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	reply, err := az.srv.Login(ctx, &types.LoginRequest{
		Tenant:    req.Tenant,
		App:       req.App,
		CaptchaId: req.CaptchaId,
		Captcha:   req.Captcha,
		Username:  req.Username,
		Password:  req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &authorize.LoginReply{
		NeedInfo: reply.NeedInfo,
		Token:    reply.Token,
	}, nil
}

func (az *Authorize) Register(c context.Context, req *authorize.RegisterRequest) (*authorize.RegisterReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	reply, err := az.srv.Register(ctx, &types.RegisterRequest{
		Tenant:    req.Tenant,
		App:       req.App,
		CaptchaId: req.CaptchaId,
		Captcha:   req.Captcha,
		Username:  req.Username,
		Password:  req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &authorize.RegisterReply{
		NeedInfo: reply.NeedInfo,
		Token:    reply.Token,
	}, nil
}

func (az *Authorize) CheckAuth(c context.Context, req *authorize.CheckAuthRequest) (*authorize.CheckAuthReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	reply, err := az.srv.CheckAuth(ctx, &types.CheckAuthRequest{
		Path:   req.Path,
		Method: req.Method,
	})
	if err != nil {
		return nil, err
	}

	if reply == nil {
		return nil, nil
	}

	return &authorize.CheckAuthReply{
		TenantId: reply.TenantId,
		UserId:   reply.UserId,
		DeptId:   reply.DeptId,
	}, nil
}

func (az *Authorize) ParseToken(c context.Context, req *authorize.ParseTokenRequest) (*authorize.ParseTokenReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())

	return &authorize.ParseTokenReply{
		UserId:   ctx.Auth().UserId,
		TenantId: ctx.Auth().TenantId,
		DeptId:   ctx.Auth().DeptId,
	}, nil
}

func (az *Authorize) GetFillInfo(c context.Context, req *authorize.GetFillInfoRequest) (*authorize.GetFillInfoReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	list, err := az.srv.GetFillInfo(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	reply := authorize.GetFillInfoReply{}
	for _, item := range list {
		reply.List = append(reply.List, &authorize.GetFillInfoReply_Field{
			Type:    item.Type,
			Name:    item.Name,
			Keyword: item.Keyword,
			Value:   item.Value,
		})
	}

	return &reply, nil
}

func (az *Authorize) FillInfo(c context.Context, req *authorize.FillInfoRequest) (*authorize.FillInfoReply, error) {
	ctx := core.MustContext(c, kratosx.WithSkipDBHook())
	token, err := az.srv.FillInfo(ctx, req.Uuid, req.Infos)
	if err != nil {
		return nil, err
	}

	return &authorize.FillInfoReply{Token: token}, nil
}
