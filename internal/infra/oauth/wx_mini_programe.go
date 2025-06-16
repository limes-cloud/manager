package oauth

func init() {
	// register("wx_official_account", "微信公众号", NewWXOfficialAccount)
}

//
// type WXOfficialAccount struct {
//	conf *types.GetOAuthorRequest
// }
//
// func NewWXOfficialAccount(req *types.GetOAuthorRequest) OAuthor {
//	return &WXOfficialAccount{conf: req}
// }
//
// func (woa WXOfficialAccount) GetOAuthToken(ctx kratosx.Context, req *types.GetOAuthTokenRequest) (*types.GetOAuthTokenReply, error) {
//	data := struct {
//		AccessToken string `json:"access_token"`
//		ExpiresIn   int    `json:"expires_in"`
//	}{}
//
//	response, err := ctx.Http().Option(func(request *resty.Request) {
//		request.SetQueryParam("appid", woa.conf.Ak).
//			SetQueryParam("secret", woa.conf.Sk).
//			SetQueryParam("grant_type", "client_credential")
//	}).Post("https://api.weixin.qq.com/cgi-bin/token", nil)
//	if err != nil {
//		return nil, err
//	}
//
//	if err = response.Result(&data); err != nil {
//		return nil, err
//	}
//
//	return &types.GetOAuthTokenReply{
//		Token:  data.AccessToken,
//		Expire: time.Now().Unix() + int64(data.ExpiresIn),
//	}, nil
// }
//
// func (woa WXOfficialAccount) GetOAuthInfo(ctx kratosx.Context, req *types.GetOAuthInfoRequest) (*types.GetOAuthInfoReply, error) {
//	data := struct {
//		Openid  string `json:"openid"`
//		Unionid string `json:"unionid"`
//	}{}
//	response, err := ctx.Http().Option(func(request *resty.Request) {
//		request.SetQueryParam("appid", woa.conf.Ak).
//			SetQueryParam("secret", woa.conf.Sk).
//			SetQueryParam("js_code", req.Code).
//			SetQueryParam("grant_type", "authorization_code")
//	}).Post("https://api.weixin.qq.com/sns/jscode2session", nil)
//	if err != nil {
//		return nil, err
//	}
//	if err = response.Result(&data); err != nil {
//		return nil, err
//	}
//	return &GetAuthInfoReply{
//		AuthId:  data.Openid,
//		UnionId: &data.Unionid,
//	}, nil
// }
