package wx

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type ServiceAccount struct {
	conf *entity.OAuther
}

func NewServiceAccount(req *entity.OAuther) (repository.OAutherFunc, error) {
	return &ServiceAccount{conf: req}, nil
}

func (woa ServiceAccount) Visible(_ core.Context, req *types.OAutherVisibleRequest) (*types.OAutherVisibleReply, error) {
	return &types.OAutherVisibleReply{
		Visible:   true,
		Recommend: req.Platform == types.PlatformWebWeiXin,
	}, nil
}

func (woa ServiceAccount) Start(_ core.Context, req *types.OAutherStartRequest) (*types.OAutherStartReply, error) {
	uid := crypto.MD5([]byte(uuid.NewString()))
	resp := types.OAutherStartReply{
		UUID:   uid,
		Action: types.OAutherWayActionJump,
		Tip:    "点击跳转授权",
	}

	// 不是微信打开
	if req.Platform != types.PlatformWebWeiXin {
		resp.Action = types.OAutherWayActionScan
		resp.Tip = "打开微信扫码登陆"
	}

	// 设置跳转url
	callback := woa.conf.GetSetting().Callback
	resp.Value = fmt.Sprintf(
		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=%s.%s.%s.%s#wechat_redirect",
		woa.conf.Ak,
		callback,
		woa.conf.Keyword,
		resp.Action,
		uid,
		"code",
	)
	return &resp, nil
}

func (woa ServiceAccount) Complete(ctx core.Context, req *types.OAutherCompleteRequest) (*types.OAutherIdentity, error) {
	token, oid, expire, err := woa.getToken(ctx, req.Code)
	if err != nil {
		return nil, err
	}
	return woa.getInfo(ctx, oid, token, expire)
}

func (woa ServiceAccount) getToken(ctx core.Context, code string) (string, string, int64, error) {
	res := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		OpenId      string `json:"openid"`
	}{}

	response, err := ctx.Request().Option(func(request *resty.Request) {
		request.SetQueryParam("appid", woa.conf.Ak).
			SetQueryParam("secret", woa.conf.Sk).
			SetQueryParam("grant_type", "authorization_code").
			SetQueryParam("code", code)
	}).Post("https://api.weixin.qq.com/sns/oauth2/access_token", nil)
	if err != nil {
		return "", "", 0, err
	}

	if err = response.Result(&res); err != nil {
		return "", "", 0, err
	}

	return res.AccessToken, res.OpenId, time.Now().Unix() + int64(res.ExpiresIn), nil
}

func (woa ServiceAccount) getInfo(ctx core.Context, oid string, token string, expire int64) (*types.OAutherIdentity, error) {
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

	response, err := ctx.Request().Option(func(request *resty.Request) {
		request.SetQueryParam("access_token", token).
			SetQueryParam("openid", oid).
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

	return &types.OAutherIdentity{
		OID:      res.Openid,
		NickName: res.Nickname,
		Gender:   transSex(res.Sex),
		Province: res.Province,
		City:     res.City,
		Country:  res.Country,
		Avatar:   res.Headimgurl,
		UnionId:  res.Unionid,
		Token:    token,
		Expire:   expire,
	}, nil
}
