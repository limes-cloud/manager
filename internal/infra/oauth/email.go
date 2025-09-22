package oauth

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"

	"github.com/limes-cloud/kratosx/library/captcha"
	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/pkg/value"
	"gopkg.in/gomail.v2"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

const (
	email           = "email"
	oauthEmailScene = "oauth:email"
)

func init() {
	register(email, "邮箱", NewEmail)
}

type Email struct {
	conf     *entity.OAuthChannel
	subject  string
	captcha  string
	fromUser string
	fromName string
	tpl      string
	host     string
	port     int
}

func NewEmail(req *entity.OAuthChannel) repository.OAuthor {
	info := req.GetExtra()
	return &Email{
		conf:     req,
		host:     info.Host,
		port:     info.Port,
		captcha:  info.Captcha,
		fromUser: req.Ak,
		fromName: info.Name,
		subject:  info.Subject,
	}
}

// OAuthHandler 前置处理
func (e *Email) OAuthHandler(ctx core.Context, req *types.OAuthHandlerRequest) (*types.OAuthHandlerReply, error) {
	// 判断是否为邮箱账号
	if !value.IsEmail(req.User) {
		return nil, errors.New("邮箱输入错误")
	}

	if e.host == "" || e.port == 0 {
		return nil, errors.New("邮件服务配置错误")
	}

	// 获取验证器
	cpt, err := ctx.Captcha().Get(ctx, e.captcha)
	if err != nil {
		return nil, err
	}

	// 获取验证码
	cptReply, err := cpt.GetCaptcha(&captcha.GetCaptchaRequest{
		Scene:    oauthEmailScene,
		User:     req.User,
		ClientIP: req.IP,
	})
	if err != nil {
		return nil, err
	}

	// 发送验证码
	if err := e.Send(ctx, req.User, cptReply); err != nil {
		_ = cpt.CancelCaptcha(cptReply.UUID)
		return nil, err
	}

	// 返回数据
	return &types.OAuthHandlerReply{
		UUID:   cptReply.UUID,
		Action: types.GetOAuthWayActionCaptcha,
		Tip:    "验证码登陆",
		Value:  fmt.Sprint(int64(cptReply.Expire.Seconds())),
	}, nil
}

func (e *Email) GetOAuthToken(ctx core.Context, req *types.GetOAuthTokenRequest) (*types.GetOAuthTokenReply, error) {
	// 获取验证器
	cpt, err := ctx.Captcha().Get(ctx, e.captcha)
	if err != nil {
		return nil, err
	}

	// 验证
	err = cpt.VerifyCaptcha(&captcha.VerifyCaptchaRequest{
		Scene:      oauthEmailScene,
		User:       req.User,
		ClientIP:   req.IP,
		VerifyCode: req.Code,
		UUID:       req.UUID,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetOAuthTokenReply{
		OID: req.User,
	}, nil
}

func (e *Email) GetOAuthInfo(_ core.Context, req *types.GetOAuthInfoRequest) (*types.GetOAuthInfoReply, error) {
	return &types.GetOAuthInfoReply{OID: req.OID}, nil
}

func (e *Email) Send(ctx core.Context, email string, req *captcha.GetCaptchaResponse) error {
	tpl := ctx.Loader(e.tpl)
	if tpl == nil || len(tpl) == 0 {
		return errors.New("not fount tpl")
	}

	n := template.New("email")
	parser, err := n.Parse(string(tpl))
	if err != nil {
		return err
	}
	html := bytes.NewBuffer([]byte(""))
	if err = parser.Execute(html, map[string]any{
		"captcha": req.VerifyCode,
		"minute":  int64(req.Expire.Minutes()),
	}); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(e.fromUser, e.fromName))
	m.SetHeader("To", m.FormatAddress(email, ""))
	m.SetHeader("Subject", e.subject)
	m.SetBody(fmt.Sprintf("text/html; charset=UTF-8"), html.String())
	d := gomail.NewDialer(
		e.host,
		e.port,
		e.conf.Ak,
		e.conf.Sk,
	)
	return d.DialAndSend(m)
}
