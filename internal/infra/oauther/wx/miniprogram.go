package wx

import (
	"github.com/silenceper/wechat/v2"

	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"

	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type MiniProgram struct {
	conf *entity.OAuther
}

func NewMiniProgram(req *entity.OAuther) (repository.OAutherFunc, error) {
	return &MiniProgram{conf: req}, nil
}

func (woa MiniProgram) config(ctx core.Context) *miniConfig.Config {
	opt := ctx.Redis().Options()

	// 使用redis
	memory := cache.NewRedis(&cache.RedisOpts{
		Host:     opt.Addr,
		Password: opt.Password,
		Database: opt.DB,
	})
	// 初始化配置
	cfg := &miniConfig.Config{
		AppID:     woa.conf.Ak,
		AppSecret: woa.conf.Sk,
		Cache:     memory,
	}
	return cfg
}

func (woa MiniProgram) Visible(_ core.Context, req *types.OAutherVisibleRequest) (*types.OAutherVisibleReply, error) {
	return &types.OAutherVisibleReply{
		Visible:       req.Platform == types.PlatformMPWeiXin,
		Recommend:     true,
		RecommendText: "一键授权登录",
	}, nil
}

func (woa MiniProgram) Handler(ctx core.Context, req *types.OAutherHandleRequest) (*types.OAutherHandleReply, error) {
	uid := crypto.MD5([]byte(uuid.NewString()))
	return &types.OAutherHandleReply{
		UUID:   uid,
		Action: types.OAutherWayActionUniLogin,
		Tip:    "授权登陆",
	}, nil
}

func (woa MiniProgram) GetToken(ctx core.Context, req *types.OAutherTokenRequest) (*types.OAutherTokenReply, error) {
	conf := woa.config(ctx)
	wc := wechat.NewWechat()
	mg := wc.GetMiniProgram(conf)

	cs, err := mg.GetAuth().Code2Session(req.Code)
	if err != nil {
		return nil, err
	}

	return &types.OAutherTokenReply{
		OID: cs.OpenID,
	}, nil
}

func (woa MiniProgram) GetInfo(ctx core.Context, req *types.OAutherInfoRequest) (*types.OAutherInfoReply, error) {
	return &types.OAutherInfoReply{
		OID: req.OID,
	}, nil
}
