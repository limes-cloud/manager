package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type OAutherFunc interface {
	// Handler 授权处理
	Handler(core.Context, *types.OAutherHandleRequest) (*types.OAutherHandleReply, error)

	// GetToken 获取授权token
	GetToken(core.Context, *types.OAutherTokenRequest) (*types.OAutherTokenReply, error)

	// GetInfo 获取授权信息
	GetInfo(core.Context, *types.OAutherInfoRequest) (*types.OAutherInfoReply, error)
}

type OAuthExecer interface {
	// ListOAutherType 获取授权列表类型
	ListOAutherType() []*types.OAutherType

	// GetOAuther 获取授权器
	GetOAuther(req *entity.OAuther) (OAutherFunc, error)
}

type OAuther interface {
	// GetOAutherByKeyword 获取指定的登陆渠道
	GetOAutherByKeyword(ctx core.Context, tid uint32, keyword string) (*entity.OAuther, error)

	// GetOAuther 获取指定的登陆渠道
	GetOAuther(ctx core.Context, id uint32) (*entity.OAuther, error)

	// ListOAuther 获取登陆渠道列表
	ListOAuther(ctx core.Context, req *types.ListOAutherRequest) ([]*entity.OAuther, uint32, error)

	// CreateOAuther 创建登陆渠道
	CreateOAuther(ctx core.Context, req *entity.OAuther) (uint32, error)

	// UpdateOAuther 更新登陆渠道
	UpdateOAuther(ctx core.Context, req *entity.OAuther) error

	// DeleteOAuther 删除登陆渠道
	DeleteOAuther(ctx core.Context, id uint32) error
}
