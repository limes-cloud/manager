package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type OAuth struct {
	repo repository.OAuth
}

func NewOAuth(oauth repository.OAuth) *OAuth {
	return &OAuth{
		repo: oauth,
	}
}

// GetTypes 获取可以开通的登录渠道
func (u *OAuth) GetTypes() []types.OAuthor {
	return u.repo.ListOAuthor()
}

// ListOAuthChannel 获取登陆渠道列表
func (u *OAuth) ListOAuthChannel(ctx core.Context, req *types.ListOAuthChannelRequest) ([]*entity.OAuthChannel, uint32, error) {
	list, total, err := u.repo.ListOAuthChannel(ctx, req)
	if err != nil {
		ctx.Logger().Errorw("msg", "list channel error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateOAuthChannel 创建登陆渠道
func (u *OAuth) CreateOAuthChannel(ctx core.Context, channel *entity.OAuthChannel) (uint32, error) {
	channel.Status = proto.Bool(false)
	id, err := u.repo.CreateOAuthChannel(ctx, channel)
	if err != nil {
		ctx.Logger().Errorw("msg", "create channel error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateOAuthChannel 更新登陆渠道
func (u *OAuth) UpdateOAuthChannel(ctx core.Context, channel *entity.OAuthChannel) error {
	if err := u.repo.UpdateOAuthChannel(ctx, channel); err != nil {
		ctx.Logger().Errorw("msg", "update channel error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteOAuthChannel 删除登陆渠道
func (u *OAuth) DeleteOAuthChannel(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteOAuthChannel(ctx, id); err != nil {
		ctx.Logger().Errorw("msg", "delete channel error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
