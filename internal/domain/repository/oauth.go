package repository

import (
	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type OAuth interface {
	// GetOAuthChannelByKeyword 获取指定的登陆渠道
	GetOAuthChannelByKeyword(ctx core.Context, keyword string) (*entity.OAuthChannel, error)

	// GetOAuthChannel 获取指定的登陆渠道
	GetOAuthChannel(ctx core.Context, id uint32) (*entity.OAuthChannel, error)

	// ListOAuthChannel 获取登陆渠道列表
	ListOAuthChannel(ctx core.Context, req *types.ListOAuthChannelRequest) ([]*entity.OAuthChannel, uint32, error)

	// CreateOAuthChannel 创建登陆渠道
	CreateOAuthChannel(ctx core.Context, req *entity.OAuthChannel) (uint32, error)

	// UpdateOAuthChannel 更新登陆渠道
	UpdateOAuthChannel(ctx core.Context, req *entity.OAuthChannel) error

	// DeleteOAuthChannel 删除登陆渠道
	DeleteOAuthChannel(ctx core.Context, id uint32) error

	// CreateOAuth 创建应用授权信息
	CreateOAuth(ctx core.Context, req *entity.OAuth) (uint32, error)

	// ListOAuth 获取应用授权信息列表
	ListOAuth(ctx core.Context, req *types.ListOAuthRequest) ([]*entity.OAuth, uint32, error)

	// DeleteOAuth 删除应用授权信息
	DeleteOAuth(ctx core.Context, userId uint32, appId uint32) error

	// IsBindOAuth 是否绑定授权信息
	IsBindOAuth(ctx core.Context, channel uint32, oid string) bool

	// GetOAuthByCO 获取指定的授权信息
	GetOAuthByCO(ctx core.Context, cid uint32, oid string) (*entity.OAuth, error)

	// UpdateOAuth 更新应用授权信息
	UpdateOAuth(ctx core.Context, req *entity.OAuth) error

	// GetTokenByUUID 获取token
	GetTokenByUUID(core.Context, string) (*types.GetOAuthTokenReply, error)

	// SetTokenByUUID 设置缓存的token
	SetTokenByUUID(core.Context, string, *types.GetOAuthTokenReply) error

	// ListOAuthor 获取授权列表
	ListOAuthor() []types.OAuthor

	// GetOAuthor 获取授权器
	GetOAuthor(req *entity.OAuthChannel) (OAuthor, error)
}

type OAuthor interface {
	OAuthHandler(core.Context, *types.OAuthHandlerRequest) (*types.OAuthHandlerReply, error)
	GetOAuthToken(core.Context, *types.GetOAuthTokenRequest) (*types.GetOAuthTokenReply, error)
	GetOAuthInfo(core.Context, *types.GetOAuthInfoRequest) (*types.GetOAuthInfoReply, error)
}
