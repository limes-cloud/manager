package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Channel interface {
	// GetChannelByKeyword 获取指定的登陆渠道
	GetChannelByKeyword(ctx kratosx.Context, keyword string) (*entity.Channel, error)

	// GetChannel 获取指定的登陆渠道
	GetChannel(ctx kratosx.Context, id uint32) (*entity.Channel, error)

	// ListChannel 获取登陆渠道列表
	ListChannel(ctx kratosx.Context, req *types.ListChannelRequest) ([]*entity.Channel, uint32, error)

	// CreateChannel 创建登陆渠道
	CreateChannel(ctx kratosx.Context, req *entity.Channel) (uint32, error)

	// UpdateChannel 更新登陆渠道
	UpdateChannel(ctx kratosx.Context, req *entity.Channel) error

	// DeleteChannel 删除登陆渠道
	DeleteChannel(ctx kratosx.Context, id uint32) error
}
