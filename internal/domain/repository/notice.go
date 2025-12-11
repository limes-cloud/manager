package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Notice interface {
	// ListNoticeClassify 获取资料分组列表
	ListNoticeClassify(ctx core.Context) ([]*entity.NoticeClassify, error)

	// CreateNoticeClassify 创建资料分组
	CreateNoticeClassify(ctx core.Context, req *entity.NoticeClassify) (uint32, error)

	// UpdateNoticeClassify 更新资料分组
	UpdateNoticeClassify(ctx core.Context, req *entity.NoticeClassify) error

	// DeleteNoticeClassify 删除资料分组
	DeleteNoticeClassify(ctx core.Context, id uint32) error

	// GetNotice 获取指定的资料信息
	GetNotice(ctx core.Context, id uint32) (*entity.Notice, error)

	// ListNotice 获取资料信息列表
	ListNotice(ctx core.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error)

	// CreateNotice 创建资料信息
	CreateNotice(ctx core.Context, req *entity.Notice) (uint32, error)

	// UpdateNotice 更新资料信息
	UpdateNotice(ctx core.Context, req *entity.Notice) error

	// DeleteNotice 删除资料信息
	DeleteNotice(ctx core.Context, id uint32) error

	// AddUserNotice 标记文章为已读
	AddUserNotice(ctx core.Context, uid, nid uint32) error
}
