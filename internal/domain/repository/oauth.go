package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type OAuth interface {
	// CreateOAuth 创建应用授权信息
	CreateOAuth(ctx kratosx.Context, req *entity.OAuth) (uint32, error)

	// ListOAuth 获取应用授权信息列表
	ListOAuth(ctx kratosx.Context, req *types.ListOAuthRequest) ([]*entity.OAuth, uint32, error)

	// DeleteOAuth 删除应用授权信息
	DeleteOAuth(ctx kratosx.Context, userId uint32, appId uint32) error

	// IsBindOAuth 是否绑定授权信息
	IsBindOAuth(ctx kratosx.Context, channel uint32, oid string) bool

	// GetOAuthByCO 获取指定的授权信息
	GetOAuthByCO(ctx kratosx.Context, cid uint32, oid string) (*entity.OAuth, error)

	// UpdateOAuth 更新应用授权信息
	UpdateOAuth(ctx kratosx.Context, req *entity.OAuth) error

	// GetTokenByUUID 获取token
	GetTokenByUUID(kratosx.Context, string) (*types.GetOAuthTokenReply, error)

	// SetTokenByUUID 设置缓存的token
	SetTokenByUUID(kratosx.Context, string, *types.GetOAuthTokenReply) error

	// ListOAuthorName 获取授权列表
	ListOAuthorName() []types.OAuthorName

	// GetOAuthor 获取授权器
	GetOAuthor(req *entity.Channel) (OAuthor, error)
}

type OAuthor interface {
	GetOAuthWay(kratosx.Context, *types.GetOAuthWayRequest) (*types.GetOAuthWayReply, error)
	GetOAuthToken(kratosx.Context, *types.GetOAuthTokenRequest) (*types.GetOAuthTokenReply, error)
	GetOAuthInfo(kratosx.Context, *types.GetOAuthInfoRequest) (*types.GetOAuthInfoReply, error)
}
