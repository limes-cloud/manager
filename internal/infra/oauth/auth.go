package oauth

import (
	"errors"
	"sort"
	"time"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/model/page"

	json "github.com/json-iterator/go"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

const (
	tokenKey = "oauth:token:"
)

var ins = make(map[string]*oauthor)

type OAuthorInitFunc func(req *entity.OAuthChannel) repository.OAuthor

type oauthor struct {
	platform string
	name     string
	OAuthorInitFunc
}

func (o oauthor) Platform() string {
	return o.platform
}

func (o oauthor) Name() string {
	return o.name
}

func register(platform string, name string, init OAuthorInitFunc) {
	ins[platform] = &oauthor{
		platform:        platform,
		name:            name,
		OAuthorInitFunc: init,
	}
}

type OAuth struct{}

func New() *OAuth {
	return &OAuth{}
}

func (oa OAuth) CreateOAuth(ctx core.Context, oauth *entity.OAuth) (uint32, error) {
	return oauth.Id, ctx.DB().Create(oauth).Error
}

func (oa OAuth) ListOAuth(ctx core.Context, req *types.ListOAuthRequest) ([]*entity.OAuth, uint32, error) {
	var (
		list  []*entity.OAuth
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.OAuth{}).Select(fs).
		Preload("Channel", "status=true").
		Where("user_id = ?", req.UserId)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = page.SearchScopes(db, &req.Search)
	if err := db.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, uint32(total), nil
}

func (oa OAuth) UpdateOAuth(ctx core.Context, oauth *entity.OAuth) error {
	return ctx.DB().
		Where("user_id=?", oauth.UserId).
		Where("channel_id=?", oauth.ChannelId).
		Updates(oauth).Error
}

func (oa OAuth) DeleteOAuth(ctx core.Context, userId uint32, channelId uint32) error {
	return ctx.DB().Where("user_id = ? and channel_id = ?", userId, channelId).Delete(&entity.OAuth{}).Error
}

func (oa OAuth) IsBindOAuth(ctx core.Context, cid uint32, oid string) bool {
	var count int64
	ctx.DB().Model(entity.OAuth{}).
		Where("user_id is not null").
		Where("channel_id=?", cid).
		Where("oid=?", oid).
		Count(&count)
	return count != 0
}

// GetOAuthByCO 通过channel_id 和 oid获取 授权信息
func (oa OAuth) GetOAuthByCO(ctx core.Context, cid uint32, oid string) (*entity.OAuth, error) {
	var oauth entity.OAuth
	return &oauth, ctx.DB().Where("channel_id=?", cid).Where("oid=?", oid).First(&oauth).Error
}

// GetTokenByUUID 通过uuid获取token
func (oa OAuth) GetTokenByUUID(ctx core.Context, s string) (*types.GetOAuthTokenReply, error) {
	res, err := ctx.Redis().Get(ctx, tokenKey+s).Result()
	if err != nil {
		return nil, err
	}
	var reply types.GetOAuthTokenReply
	if err := json.Unmarshal([]byte(res), &reply); err != nil {
		return nil, err
	}
	return &reply, nil
}

// SetTokenByUUID 通过uuid设置token
func (oa OAuth) SetTokenByUUID(ctx core.Context, s string, reply *types.GetOAuthTokenReply) error {
	b, _ := json.Marshal(reply)
	return ctx.Redis().SetNX(ctx, tokenKey+s, string(b), 5*time.Minute).Err()
}

// ListOAuthor 获取实现的授权器名称
func (oa OAuth) ListOAuthor() []types.OAuthor {
	var list []types.OAuthor
	for _, item := range ins {
		list = append(list, item)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Platform() > list[j].Platform()
	})
	return list
}

// GetOAuthor 获取指定的授权器
func (oa OAuth) GetOAuthor(req *entity.OAuthChannel) (repository.OAuthor, error) {
	if req == nil || ins[req.Type] == nil {
		return nil, errors.New("not fount oauthor " + req.Type)
	}
	return ins[req.Type].OAuthorInitFunc(req), nil
}

// GetOAuthChannelByKeyword 获取指定数据
func (r OAuth) GetOAuthChannelByKeyword(ctx core.Context, keyword string) (*entity.OAuthChannel, error) {
	var (
		channel = entity.OAuthChannel{}
		fs      = []string{"*"}
	)
	return &channel, ctx.DB().Select(fs).Where("keyword = ?", keyword).First(&channel).Error
}

// GetOAuthChannel 获取指定的数据
func (r OAuth) GetOAuthChannel(ctx core.Context, id uint32) (*entity.OAuthChannel, error) {
	var (
		channel = entity.OAuthChannel{}
		fs      = []string{"*"}
	)
	return &channel, ctx.DB().Select(fs).First(&channel, id).Error
}

// ListOAuthChannel 获取列表
func (r OAuth) ListOAuthChannel(ctx core.Context, req *types.ListOAuthChannelRequest) ([]*entity.OAuthChannel, uint32, error) {
	var (
		list  []*entity.OAuthChannel
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.OAuthChannel{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword LIKE ?", *req.Keyword+"%")
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateOAuthChannel 创建数据
func (r OAuth) CreateOAuthChannel(ctx core.Context, channel *entity.OAuthChannel) (uint32, error) {
	return channel.Id, ctx.DB().Create(channel).Error
}

// UpdateOAuthChannel 更新数据
func (r OAuth) UpdateOAuthChannel(ctx core.Context, channel *entity.OAuthChannel) error {
	return ctx.DB().Updates(channel).Error
}

// DeleteOAuthChannel 删除数据
func (r OAuth) DeleteOAuthChannel(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.OAuthChannel{}).Error
}
