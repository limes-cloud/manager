package yiban

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/forgoer/openssl"
	"github.com/google/uuid"
	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type YiBan struct {
	conf *entity.OAuther
}

func NewYiBan(req *entity.OAuther) (repository.OAutherFunc, error) {
	return &YiBan{conf: req}, nil
}

func (y YiBan) Visible(ctx core.Context, req *types.OAutherVisibleRequest) (*types.OAutherVisibleReply, error) {
	return &types.OAutherVisibleReply{
		Visible:   true,
		Recommend: req.Platform == types.PlatformWebYiBan,
	}, nil
}

// Start 获取鉴权方式
func (y YiBan) Start(_ core.Context, req *types.OAutherStartRequest) (*types.OAutherStartReply, error) {
	uid := crypto.MD5([]byte(uuid.NewString()))
	resp := types.OAutherStartReply{
		UUID:   uid,
		Action: types.OAutherWayActionJump,
		Tip:    "点击跳转授权",
	}

	// 不是 yiban app 打开
	if req.Platform == types.PlatformWebYiBan {
		resp.Action = types.OAutherWayActionScan
		resp.Tip = "打开易APP扫码登陆"
	}

	// 设置跳转url
	resp.Value = fmt.Sprintf(
		"https://oauth.yiban.cn/code/html?client_id=%s&redirect_uri=%s&state=%s.%s.%s.%s",
		y.conf.Ak,
		y.conf.GetSetting().Callback,
		y.conf.Keyword,
		resp.Action,
		uid,
		"verify_request",
	)
	return &resp, nil
}

func (y YiBan) Complete(ctx core.Context, req *types.OAutherCompleteRequest) (*types.OAutherIdentity, error) {
	token, expire, err := y.getToken(req.Code)
	if err != nil {
		return nil, err
	}
	return y.getInfo(ctx, token, expire)
}

func (y YiBan) getToken(code string) (string, int64, error) {
	res := struct {
		VisitOauth struct {
			AccessToken  string `json:"access_token"`
			TokenExpires string `json:"token_expires"`
		} `json:"visit_oauth"`
	}{}

	src := y.hexToByte(code)
	iv := []byte(y.conf.Ak)
	key := []byte(y.conf.Sk)

	body, _ := openssl.AesCBCDecrypt(src, key, iv, openssl.PKCS7_PADDING)
	if body == nil {
		return "", 0, errors.New("decrypt error")
	}

	// 解析数据
	if err := json.Unmarshal(body, &res); err != nil {
		return "", 0, err
	}

	return res.VisitOauth.AccessToken, y.toInt64(res.VisitOauth.TokenExpires), nil
}

func (y YiBan) getInfo(ctx core.Context, token string, expire int64) (*types.OAutherIdentity, error) {
	res := struct {
		Status string `json:"status"`
		Info   struct {
			MsgCN      string `json:"msgCN"`
			Userid     string `json:"yb_userid"`
			Username   string `json:"yb_username"`
			Usernick   string `json:"yb_usernick"`
			Sex        string `json:"yb_sex"`
			Birthday   string `json:"yb_birthday"`
			Money      string `json:"yb_money"`
			Exp        string `json:"yb_exp"`
			Userhead   string `json:"yb_userhead"`
			Regtime    string `json:"yb_regtime"`
			Schoolid   string `json:"yb_schoolid"`
			Schoolname string `json:"yb_schoolname"`
		} `json:"info"`
	}{}

	url := "https://openapi.yiban.cn/user/me?access_token=" + token
	resp, err := ctx.Request().Get(url)
	if err != nil {
		return nil, err
	}

	if err := resp.Result(&res); err != nil {
		return nil, err
	}

	if res.Status == "error" {
		return nil, errors.New(res.Info.MsgCN)
	}

	return &types.OAutherIdentity{
		OID:      res.Info.Userid,
		RealName: res.Info.Username,
		NickName: res.Info.Usernick,
		Gender:   res.Info.Sex,
		Birthday: res.Info.Birthday,
		Avatar:   res.Info.Userhead,
		Token:    token,
		Expire:   expire,
	}, nil
}

func (y YiBan) hexToByte(hex string) []byte {
	length := len(hex) / 2
	slice := make([]byte, length)
	rs := []rune(hex)
	for i := 0; i < length; i++ {
		s := string(rs[i*2 : i*2+2])
		value, _ := strconv.ParseInt(s, 16, 10)
		slice[i] = byte(value & 0xFF)
	}
	return slice
}

func (y YiBan) toInt64(in string) int64 {
	val, _ := strconv.ParseInt(in, 10, 64)
	return val
}
