package oauth

import (
	"errors"
	"fmt"
	"sort"
	"time"

	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

const (
	tokenKey = "oauth:token:"
)

var ins = make(map[string]*creator)

type Creator func(req *entity.Channel) repository.OAuthor

type creator struct {
	platform string
	name     string
	Creator
}

func (o creator) Platform() string {
	return o.platform
}

func (o creator) Name() string {
	return o.name
}

func register(platform string, name string, at Creator) {
	ins[platform] = &creator{
		platform: platform,
		name:     name,
		Creator:  at,
	}
}

type OAuth struct {
}

func New() *OAuth {
	return &OAuth{}
}

func (oa OAuth) CreateOAuth(ctx kratosx.Context, oauth *entity.OAuth) (uint32, error) {
	return oauth.Id, ctx.DB().Create(oauth).Error
}

func (oa OAuth) ListOAuth(ctx kratosx.Context, req *types.ListOAuthRequest) ([]*entity.OAuth, uint32, error) {
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
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}
	if err := db.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, uint32(total), nil
}

func (oa OAuth) UpdateOAuth(ctx kratosx.Context, oauth *entity.OAuth) error {
	return ctx.DB().
		Where("user_id=?", oauth.UserId).
		Where("channel_id=?", oauth.ChannelId).
		Updates(oauth).Error
}

func (oa OAuth) DeleteOAuth(ctx kratosx.Context, userId uint32, channelId uint32) error {
	return ctx.DB().Where("user_id = ? and channel_id = ?", userId, channelId).Delete(&entity.OAuth{}).Error
}

func (oa OAuth) IsBindOAuth(ctx kratosx.Context, cid uint32, oid string) bool {
	var count int64
	ctx.DB().Model(entity.OAuth{}).
		Where("user_id is not null").
		Where("channel_id=?", cid).
		Where("oid=?", oid).
		Count(&count)
	return count != 0
}

func (oa OAuth) GetOAuthByCO(ctx kratosx.Context, cid uint32, oid string) (*entity.OAuth, error) {
	var oauth entity.OAuth
	return &oauth, ctx.DB().Where("channel_id=?", cid).Where("oid=?", oid).First(&oauth).Error
}

func (oa OAuth) GetTokenByUUID(ctx kratosx.Context, s string) (*types.GetOAuthTokenReply, error) {
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

func (oa OAuth) SetTokenByUUID(ctx kratosx.Context, s string, reply *types.GetOAuthTokenReply) error {
	b, _ := json.Marshal(reply)
	return ctx.Redis().SetNX(ctx, tokenKey+s, string(b), 5*time.Minute).Err()
}

// ListOAuthorName 获取实现的授权器名称
func (oa OAuth) ListOAuthorName() []types.OAuthorName {
	var list []types.OAuthorName
	for _, item := range ins {
		list = append(list, item)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Platform() > list[j].Platform()
	})
	return list
}

// GetOAuthor 获取指定的授权器
func (oa OAuth) GetOAuthor(req *entity.Channel) (repository.OAuthor, error) {
	if req == nil || ins[req.Type] == nil {
		return nil, errors.New("ot fount platform oauthor")
	}

	return ins[req.Type].Creator(req), nil
}
