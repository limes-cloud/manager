package oauth

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

func init() {
	register("wx_official_account", "微信公众号", NewWXOfficialAccount)
}

type WXOfficialAccount struct {
	conf *entity.Channel
}

func NewWXOfficialAccount(req *entity.Channel) repository.OAuthor {
	return &WXOfficialAccount{conf: req}
}

func (woa WXOfficialAccount) GetOAuthWay(_ kratosx.Context, req *types.GetOAuthWayRequest) (*types.GetOAuthWayReply, error) {
	uid := crypto.MD5([]byte(uuid.NewString()))
	var resp = types.GetOAuthWayReply{
		UUID:      uid,
		Action:    types.GetOAuthWayActionJump,
		Tip:       "点击跳转授权",
		CodeField: "code",
	}

	// 不是 yiban app 打开
	if !strings.Contains(req.UserAgent, "MicroMessenger") {
		resp.Action = types.GetOAuthWayActionScan
		resp.Tip = "打开微信扫码登陆"
	}

	// 设置跳转url
	resp.Value = fmt.Sprintf(
		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=%s.%s.%s#wechat_redirect",
		woa.conf.Ak,
		woa.conf.GetExtra().CallBack,
		woa.conf.Keyword,
		resp.Action,
		uid,
	)
	return &resp, nil
}

func (woa WXOfficialAccount) GetOAuthToken(ctx kratosx.Context, req *types.GetOAuthTokenRequest) (*types.GetOAuthTokenReply, error) {
	res := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		OpenId      string `json:"openid"`
	}{}

	response, err := ctx.Http().Option(func(request *resty.Request) {
		request.SetQueryParam("appid", woa.conf.Ak).
			SetQueryParam("secret", woa.conf.Sk).
			SetQueryParam("grant_type", "authorization_code").
			SetQueryParam("code", req.Code)
	}).Post("https://api.weixin.qq.com/sns/oauth2/access_token", nil)
	if err != nil {
		return nil, err
	}

	if err = response.Result(&res); err != nil {
		return nil, err
	}

	return &types.GetOAuthTokenReply{
		OID:    res.OpenId,
		Token:  res.AccessToken,
		Expire: time.Now().Unix() + int64(res.ExpiresIn),
	}, nil
}

func (woa WXOfficialAccount) GetOAuthInfo(ctx kratosx.Context, req *types.GetOAuthInfoRequest) (*types.GetOAuthInfoReply, error) {
	res := struct {
		Openid     string `json:"openid"`
		Nickname   string `json:"nickname"`
		Sex        int    `json:"sex"`
		Province   string `json:"province"`
		City       string `json:"city"`
		Country    string `json:"country"`
		Headimgurl string `json:"headimgurl"`
		Unionid    string `json:"unionid"`
	}{}

	response, err := ctx.Http().Option(func(request *resty.Request) {
		request.SetQueryParam("access_token", req.Token).
			SetQueryParam("openid", req.OID).
			SetQueryParam("lang", "zh_CN")
	}).Post("https://api.weixin.qq.com/sns/userinfo", nil)
	if err != nil {
		return nil, err
	}
	if err = response.Result(&res); err != nil {
		return nil, err
	}

	transSex := func(sex int) string {
		switch sex {
		case 1:
			return "M"
		case 2:
			return "F"
		default:
			return "U"
		}
	}

	return &types.GetOAuthInfoReply{
		OID:      res.Openid,
		NickName: res.Nickname,
		Gender:   transSex(res.Sex),
		Province: res.Province,
		City:     res.City,
		Country:  res.Country,
		Avatar:   res.Headimgurl,
		UnionId:  res.Unionid,
	}, nil
}
