package oauth

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"time"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"gopkg.in/gomail.v2"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

const (
	email         = "email"
	emailTemplate = `
<div style="margin:auto;max-width: 700px;">
    <div style="padding:10px 0; margin:0 10px; display: flex;align-items: center; ">
        <image src="http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image"
               width='35px' height='35px' />
        <span style="margin: 0 10px;font-size: 20px; font-weight: 700;">统一应用管理平台</span>
    </div>
    <div
            style="box-sizing: border-box;min-height: 100px;padding:15px;margin:0 10px 25px 10px; line-height: 24px; text-align: justify; border-top: 2px solid #1e90ff; background-color: #fefefe;box-shadow: 0 1px 8px rgba(0,0,0,0.1);">
        <div style="font-size:14px;white-space: normal;margin: 0px;padding: 0px;box-sizing: border-box;">
            您的邮箱验证码为： {{.captcha}}，该验证码在{{.minute}}分钟内有效，为了账号安全，请勿向其他人泄漏验证码信息。
        </div>
    </div>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
        <tbody >
        <tr>
            <td align="center" style="color:#76808E;font-size:13px;">开放协作，拥抱未来</td>
        </tr>
        <tr>
            <td align="center" style="color:#76808E;font-size:13px;">
                Copyright © 2025 qlime.cn. All rights reserved.
            </td>
        </tr>
        </tbody>
    </table>
</div>
`
)

func init() {
	register(email, "邮箱", NewEmail)
}

type Email struct {
	conf         *entity.Channel
	captcha      *captcha
	fromUser     string
	fromName     string
	subject      string
	length       int
	duration     int
	limit        int
	refreshTime  int
	host         string
	port         int
	uniqueDevice bool
}

func NewEmail(req *entity.Channel) repository.OAuthor {
	info := req.GetExtra()
	if info.Length < 4 {
		info.Length = 4
	}
	if info.Duration < 60 {
		info.Duration = 180
	}
	if info.RefreshTime == 0 {
		info.RefreshTime = info.Duration
	}
	if info.Limit == 0 {
		info.Limit = 5
	}
	if info.Subject == "" {
		info.Subject = "验证码通知"
	}

	return &Email{
		conf:         req,
		fromUser:     req.Ak,
		length:       info.Length,
		duration:     info.Duration,
		limit:        info.Limit,
		uniqueDevice: info.UniqueDevice,
		host:         info.Host,
		port:         info.Port,
		refreshTime:  info.RefreshTime,
		subject:      info.Subject,
		captcha:      &captcha{},
	}
}

// GetOAuthWay 获取鉴权方式
func (e *Email) GetOAuthWay(ctx kratosx.Context, req *types.GetOAuthWayRequest) (*types.GetOAuthWayReply, error) {
	// 判断是否为邮箱账号
	if !valx.IsEmail(req.User) {
		return nil, errors.New("邮箱输入错误")
	}

	if e.host == "" || e.port == 0 {
		return nil, errors.New("邮件服务配置错误")
	}

	// 生成随机验证码
	verifyCode := e.captcha.randomCode(e.length)

	// 获取当前场景下客户端的唯一id
	sid := e.captcha.sid(e.conf.Keyword, req.IP)

	// 判断是否超过最大的获取次数
	if e.limit != 0 {
		var count int
		_ = ctx.Redis().Get(ctx, sid).Scan(&count)
		if count > e.limit {
			return nil, errors.New("当前获取验证码次数已超限")
		}
	}

	// 获取当前场景下客户端用户的唯一id
	uid := e.captcha.uid(sid, req.User)

	// 判断是否允许刷新验证码
	if ttl := ctx.Redis().TTL(ctx, uid).Val(); ttl.Seconds() > 0 {
		if e.refreshTime > 0 && (float64(e.duration)-ttl.Seconds()) < float64(e.refreshTime) {
			return nil, errors.New("请勿重复请求验证码")
		}
		// 清除旧的验证码
		ctx.Redis().Del(ctx, uid)
	}

	// 生成验证码的唯一id
	aid := e.captcha.aid(e.conf.Keyword, req.User, verifyCode)

	// 发送验证码
	if err := e.Send(req.User, "", map[string]any{
		"captcha": verifyCode,
		"minute":  e.duration / 60,
	}); err != nil {
		return nil, err
	}

	// 设置验证码
	if err := ctx.Redis().Set(ctx, uid, aid, time.Duration(e.duration)*time.Second).Err(); err != nil {
		return nil, err
	}

	// 当天请求次数累计
	now := time.Now()
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
	if !ctx.Redis().SetNX(ctx, sid, 1, endTime.Sub(now)).Val() {
		_ = ctx.Redis().Incr(ctx, sid)
	}

	var resp = types.GetOAuthWayReply{
		UUID:   e.captcha.getUUIDByUID(uid),
		Action: types.GetOAuthWayActionCaptcha,
		Tip:    "验证码登陆",
		Value:  fmt.Sprint(e.duration),
	}

	return &resp, nil
}

func (e *Email) GetOAuthToken(ctx kratosx.Context, req *types.GetOAuthTokenRequest) (*types.GetOAuthTokenReply, error) {
	// 获取当前场景下客户端的唯一id
	sid := e.captcha.sid(e.conf.Keyword, req.IP)

	// 通过uuid生成uid
	uid := e.captcha.getUIDByUUID(req.UUID)

	// 唯一设备校验
	if e.uniqueDevice && uid != e.captcha.uid(sid, req.User) {
		return nil, errors.New("验证码错误")
	}

	// 获取验证码的唯一aid
	aid := e.captcha.aid(e.conf.Keyword, req.User, req.Code)

	// 判断验证码是否正确,或过期
	oriAid, _ := ctx.Redis().Get(ctx, uid).Result()
	if oriAid != aid {
		return nil, errors.New("验证码错误")
	}

	// 清除验证码
	_ = ctx.Redis().Del(ctx, uid).Err()

	return &types.GetOAuthTokenReply{
		OID: req.User,
	}, nil
}

func (e *Email) GetOAuthInfo(_ kratosx.Context, req *types.GetOAuthInfoRequest) (*types.GetOAuthInfoReply, error) {
	return &types.GetOAuthInfoReply{
		OID: req.OID,
	}, nil
}

func (e *Email) Send(email string, name string, variable map[string]any) error {
	n := template.New("email")
	parser, err := n.Parse(emailTemplate)
	if err != nil {
		return err
	}
	html := bytes.NewBuffer([]byte(""))
	if err = parser.Execute(html, variable); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(e.fromUser, e.fromName))
	m.SetHeader("To", m.FormatAddress(email, name))
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
