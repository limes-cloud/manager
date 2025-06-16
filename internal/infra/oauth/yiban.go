package oauth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/forgoer/openssl"
	"github.com/google/uuid"
	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

func init() {
	register("yiban", "易班", NewYiBan)
}

type YiBan struct {
	conf *entity.Channel
}

func NewYiBan(req *entity.Channel) repository.OAuthor {
	return &YiBan{conf: req}
}

// GetOAuthWay 获取鉴权方式
func (y YiBan) GetOAuthWay(_ kratosx.Context, req *types.GetOAuthWayRequest) (*types.GetOAuthWayReply, error) {
	uid := crypto.MD5([]byte(uuid.NewString()))
	var resp = types.GetOAuthWayReply{
		UUID: uid,
		Type: types.GetOAuthWayTypeJump,
		Tip:  "点击跳转授权",
	}

	// 不是 yiban app 打开
	if !strings.Contains(req.UserAgent, "yiban") {
		resp.Type = types.GetOAuthWayTypeScan
		resp.Tip = "打开易APP扫码登陆"
	}

	// 设置跳转url
	resp.Value = fmt.Sprintf(
		"https://oauth.yiban.cn/code/html?client_id=%s&redirect_uri=%s&state=%s.%s.%s",
		y.conf.Ak,
		y.conf.GetExtra().CallBack,
		y.conf.Keyword,
		resp.Type,
		uid,
	)
	return &resp, nil
}

func (y YiBan) GetOAuthToken(_ kratosx.Context, req *types.GetOAuthTokenRequest) (*types.GetOAuthTokenReply, error) {
	res := struct {
		VisitOauth struct {
			AccessToken  string `json:"access_token"`
			TokenExpires string `json:"token_expires"`
		} `json:"visit_oauth"`
	}{}

	src := y.hexToByte(req.Code)
	iv := []byte(y.conf.Ak)
	key := []byte(y.conf.Sk)

	body, _ := openssl.AesCBCDecrypt(src, key, iv, openssl.PKCS7_PADDING)
	if body == nil {
		return nil, errors.New("decrypt error")
	}

	// 解析数据
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &types.GetOAuthTokenReply{
		Token:  res.VisitOauth.AccessToken,
		Expire: y.toInt64(res.VisitOauth.TokenExpires),
	}, nil
}

func (y YiBan) GetOAuthInfo(ctx kratosx.Context, req *types.GetOAuthInfoRequest) (*types.GetOAuthInfoReply, error) {
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

	url := "https://openapi.yiban.cn/user/me?access_token=" + req.Token
	resp, err := ctx.Http().Get(url)
	if err != nil {
		return nil, err
	}

	if err := resp.Result(&res); err != nil {
		return nil, err
	}

	if res.Status == "error" {
		return nil, errors.New(res.Info.MsgCN)
	}

	return &types.GetOAuthInfoReply{
		OID:      res.Info.Userid,
		RealName: res.Info.Username,
		NickName: res.Info.Usernick,
		Gender:   res.Info.Sex,
		Birthday: res.Info.Birthday,
		Avatar:   res.Info.Userhead,
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
