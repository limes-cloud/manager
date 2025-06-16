package service

import (
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Channel struct {
	conf  *conf.Config
	repo  repository.Channel
	file  repository.File
	oauth repository.OAuth
}

func NewChannel(
	conf *conf.Config,
	repo repository.Channel,
	file repository.File,
	oauth repository.OAuth,
) *Channel {
	return &Channel{
		conf:  conf,
		repo:  repo,
		file:  file,
		oauth: oauth,
	}
}

// GetTypes 获取可以开通的登录渠道
func (u *Channel) GetTypes() []types.OAuthorName {
	return u.oauth.ListOAuthorName()
}

// ListAdminChannel 获取后台登陆渠道列表
func (u *Channel) ListAdminChannel(ctx kratosx.Context) ([]*entity.Channel, uint32, error) {
	list, total, err := u.repo.ListChannel(ctx, &types.ListChannelRequest{
		Page:     1,
		PageSize: 10,
		Admin:    proto.Bool(true),
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "list channel error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	for _, item := range list {
		item.LogoUrl = u.file.GetFileURL(ctx, item.Logo)
	}
	return list, total, nil
}

// ListChannel 获取登陆渠道列表
func (u *Channel) ListChannel(ctx kratosx.Context, req *types.ListChannelRequest) ([]*entity.Channel, uint32, error) {
	list, total, err := u.repo.ListChannel(ctx, req)
	if err != nil {
		ctx.Logger().Errorw("msg", "list channel error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}

	for _, item := range list {
		item.LogoUrl = u.file.GetFileURL(ctx, item.Logo)
	}
	return list, total, nil
}

// CreateChannel 创建登陆渠道
func (u *Channel) CreateChannel(ctx kratosx.Context, channel *entity.Channel) (uint32, error) {
	channel.Status = proto.Bool(false)
	id, err := u.repo.CreateChannel(ctx, channel)
	if err != nil {
		ctx.Logger().Errorw("msg", "create channel error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateChannel 更新登陆渠道
func (u *Channel) UpdateChannel(ctx kratosx.Context, channel *entity.Channel) error {
	if err := u.repo.UpdateChannel(ctx, channel); err != nil {
		ctx.Logger().Errorw("msg", "update channel error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteChannel 删除登陆渠道
func (u *Channel) DeleteChannel(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteChannel(ctx, id); err != nil {
		ctx.Logger().Errorw("msg", "delete channel error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
