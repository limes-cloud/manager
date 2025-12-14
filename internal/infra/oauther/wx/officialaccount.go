package wx

//
//import (
//	"fmt"
//	"strings"
//	"time"
//
//	"github.com/silenceper/wechat/v2/officialaccount"
//
//	"github.com/silenceper/wechat/v2"
//	"github.com/silenceper/wechat/v2/cache"
//	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
//
//	"github.com/go-resty/resty/v2"
//	"github.com/google/uuid"
//	"github.com/limes-cloud/kratosx/pkg/crypto"
//	"github.com/limes-cloud/manager/internal/core"
//	"github.com/limes-cloud/manager/internal/domain/entity"
//	"github.com/limes-cloud/manager/internal/domain/repository"
//	"github.com/limes-cloud/manager/internal/types"
//)
//
//type OfficialAccount struct {
//	conf *entity.OAuther
//}
//
//func NewOfficialAccount(req *entity.OAuther) (repository.OAutherFunc, error) {
//	return &OfficialAccount{conf: req}, nil
//}
//
//func (woa OfficialAccount) config(ctx core.Context) *offConfig.Config {
//	opt := ctx.Redis().Options()
//
//	// 使用redis
//	memory := cache.NewRedis(&cache.RedisOpts{
//		Host:     opt.Addr,
//		Password: opt.Password,
//		Database: opt.DB,
//	})
//	setting := woa.conf.GetSetting()
//
//	// 初始化配置
//	cfg := &offConfig.Config{
//		AppID:          woa.conf.Ak,
//		AppSecret:      woa.conf.Sk,
//		Token:          woa.conf.Setting,
//		EncodingAESKey: setting.AesKey,
//		Cache:          memory,
//	}
//	return cfg
//}
//
//func (woa OfficialAccount) Handler(ctx core.Context, req *types.OAutherHandleRequest) (*types.OAutherHandleReply, error) {
//	client := woa.client(ctx)
//
//	wc := wechat.NewWechat()
//	a := wc.GetOpenPlatform(woa.config(ctx))
//	a.
//
//	ticket, err := qr.CreateTemporaryQRCode(expireSeconds, sceneID)
//
//	uid := crypto.MD5([]byte(uuid.NewString()))
//	resp := types.OAutherHandleReply{
//		UUID:   uid,
//		Action: types.OAutherWayActionJump,
//		Tip:    "点击跳转授权",
//	}
//
//	// 不是微信打开
//	if !strings.Contains(req.UserAgent, "MicroMessenger") {
//		resp.Action = types.OAutherWayActionScan
//		resp.Tip = "打开微信扫码登陆"
//	}
//
//	// 设置跳转url
//	callback := woa.conf.GetSetting().Callback
//	resp.Value = fmt.Sprintf(
//		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=%s.%s.%s.%s#wechat_redirect",
//		woa.conf.Ak,
//		callback,
//		woa.conf.Keyword,
//		resp.Action,
//		uid,
//		"code",
//	)
//	return &resp, nil
//}
//
//func (woa OfficialAccount) GetToken(ctx core.Context, req *types.OAutherTokenRequest) (*types.OAutherTokenReply, error) {
//	res := struct {
//		AccessToken string `json:"access_token"`
//		ExpiresIn   int    `json:"expires_in"`
//		OpenId      string `json:"openid"`
//	}{}
//
//	response, err := ctx.Request().Option(func(request *resty.Request) {
//		request.SetQueryParam("appid", woa.conf.Ak).
//			SetQueryParam("secret", woa.conf.Sk).
//			SetQueryParam("grant_type", "authorization_code").
//			SetQueryParam("code", req.Code)
//	}).Post("https://api.weixin.qq.com/sns/oauth2/access_token", nil)
//	if err != nil {
//		return nil, err
//	}
//
//	if err = response.Result(&res); err != nil {
//		return nil, err
//	}
//
//	return &types.OAutherTokenReply{
//		OID:    res.OpenId,
//		Token:  res.AccessToken,
//		Expire: time.Now().Unix() + int64(res.ExpiresIn),
//	}, nil
//}
//
//func (woa OfficialAccount) GetInfo(ctx core.Context, req *types.OAutherInfoRequest) (*types.OAutherInfoReply, error) {
//	res := struct {
//		Openid     string `json:"openid"`
//		Nickname   string `json:"nickname"`
//		Sex        int    `json:"sex"`
//		Province   string `json:"province"`
//		City       string `json:"city"`
//		Country    string `json:"country"`
//		Headimgurl string `json:"headimgurl"`
//		Unionid    string `json:"unionid"`
//	}{}
//
//	response, err := ctx.Request().Option(func(request *resty.Request) {
//		request.SetQueryParam("access_token", req.Token).
//			SetQueryParam("openid", req.OID).
//			SetQueryParam("lang", "zh_CN")
//	}).Post("https://api.weixin.qq.com/sns/userinfo", nil)
//	if err != nil {
//		return nil, err
//	}
//	if err = response.Result(&res); err != nil {
//		return nil, err
//	}
//
//	transSex := func(sex int) string {
//		switch sex {
//		case 1:
//			return "M"
//		case 2:
//			return "F"
//		default:
//			return "U"
//		}
//	}
//
//	return &types.OAutherInfoReply{
//		OID:      res.Openid,
//		NickName: res.Nickname,
//		Gender:   transSex(res.Sex),
//		Province: res.Province,
//		City:     res.City,
//		Country:  res.Country,
//		Avatar:   res.Headimgurl,
//		UnionId:  res.Unionid,
//	}, nil
//}
