package email

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

type Email struct {
	ak       string
	sk       string
	subject  string
	from     string
	template string
	host     string
	port     int
}

func NewEmail(req *entity.OAuther) (repository.OAutherFunc, error) {
	info := req.GetSetting()
	return &Email{
		ak:       req.Ak,
		sk:       req.Sk,
		from:     info.Email.From,
		host:     info.Email.Host,
		port:     info.Email.Port,
		subject:  info.Email.Subject,
		template: info.Email.Template,
	}, nil
}

// Handler 前置处理
func (e *Email) Handler(ctx core.Context, req *types.OAutherHandleRequest) (*types.OAutherHandleReply, error) {
	// 判断是否为邮箱账号
	if !value.IsEmail(req.Account) {
		return nil, errors.New("邮箱输入错误")
	}

	if e.host == "" || e.port == 0 {
		return nil, errors.New("邮件服务配置错误")
	}

	// 获取验证器
	cpt, err := ctx.Captcha().Get(ctx, email)
	if err != nil {
		return nil, err
	}

	// 获取验证码
	cptReply, err := cpt.GetCaptcha(&captcha.GetCaptchaRequest{
		Scene:    oauthEmailScene,
		User:     req.Account,
		ClientIP: req.IP,
	})
	if err != nil {
		return nil, err
	}

	// 发送验证码
	if err := e.Send(req.Account, cptReply); err != nil {
		_ = cpt.CancelCaptcha(cptReply.UUID)
		return nil, err
	}

	// 返回数据
	return &types.OAutherHandleReply{
		UUID:   cptReply.UUID,
		Action: types.OAutherWayActionCaptcha,
		Tip:    "验证码登陆",
		Value:  fmt.Sprint(int64(cpt.GetCaptchaDuration().Seconds())),
	}, nil
}

func (e *Email) GetToken(ctx core.Context, req *types.OAutherTokenRequest) (*types.OAutherTokenReply, error) {
	// 获取验证器
	cpt, err := ctx.Captcha().Get(ctx, email)
	if err != nil {
		return nil, err
	}

	// 验证
	err = cpt.VerifyCaptcha(&captcha.VerifyCaptchaRequest{
		Scene:      oauthEmailScene,
		User:       req.Account,
		ClientIP:   req.IP,
		VerifyCode: req.Code,
		UUID:       req.UUID,
	})
	if err != nil {
		return nil, err
	}

	return &types.OAutherTokenReply{
		OID: req.Account,
	}, nil
}

func (e *Email) GetInfo(_ core.Context, req *types.OAutherInfoRequest) (*types.OAutherInfoReply, error) {
	return &types.OAutherInfoReply{OID: req.OID}, nil
}

func (e *Email) Send(email string, req *captcha.GetCaptchaResponse) error {
	n := template.New("email")
	parser, err := n.Parse(e.template)
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
	m.SetHeader("From", m.FormatAddress(e.ak, e.from))
	m.SetHeader("To", m.FormatAddress(email, ""))
	m.SetHeader("Subject", e.subject)
	m.SetBody(fmt.Sprintf("text/html; charset=UTF-8"), html.String())
	d := gomail.NewDialer(
		e.host,
		e.port,
		e.ak,
		e.sk,
	)
	return d.DialAndSend(m)
}
