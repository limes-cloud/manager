package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Notice struct{}

var (
	noticeIns  *Notice
	noticeOnce sync.Once
)

func NewNotice() *Notice {
	noticeOnce.Do(func() {
		noticeIns = &Notice{}
	})
	return noticeIns
}

// ListNoticeClassify 获取列表
func (r Notice) ListNoticeClassify(ctx core.Context) ([]*entity.NoticeClassify, error) {
	var list []*entity.NoticeClassify
	return list, ctx.DB().Model(entity.NoticeClassify{}).Find(&list).Error
}

// CreateNoticeClassify 创建数据
func (r Notice) CreateNoticeClassify(ctx core.Context, fc *entity.NoticeClassify) (uint32, error) {
	return fc.Id, ctx.DB().Create(fc).Error
}

// GetNoticeClassify 获取指定的数据
func (r Notice) GetNoticeClassify(ctx core.Context, id uint32) (*entity.NoticeClassify, error) {
	fc := entity.NoticeClassify{}
	return &fc, ctx.DB().First(&fc, id).Error
}

// UpdateNoticeClassify 更新数据
func (r Notice) UpdateNoticeClassify(ctx core.Context, fc *entity.NoticeClassify) error {
	return ctx.DB().Updates(fc).Error
}

// DeleteNoticeClassify 删除数据
func (r Notice) DeleteNoticeClassify(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.NoticeClassify{}).Error
}

// GetNotice 获取指定的数据
func (r Notice) GetNotice(ctx core.Context, id uint32) (*entity.Notice, error) {
	notice := entity.Notice{}
	return &notice, ctx.DB().First(&notice, id).Error
}

// ListNotice 获取列表
func (r Notice) ListNotice(ctx core.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error) {
	var (
		list  []*entity.Notice
		total int64
	)

	db := ctx.DB().Model(entity.Notice{})
	db = db.Preload("App").Preload("Classify")

	if req.AppId != nil {
		db = db.Where("app_id = ?", *req.AppId)
	}
	if req.Title != nil {
		db = db.Where("title like ?", "%"+*req.Title+"%")
	}
	if req.ClassifyId != nil {
		db = db.Where("category_id = ?", *req.ClassifyId)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.NotRead != nil && req.UserId != nil && *req.NotRead {
		db = db.Joins("left join notice_user on notice_user.notice_id = notice.id and notice_user.user_id = ?", *req.UserId)
		db = db.Where("notice_user.id is null")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = page.SearchScopes(db, req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateNotice 创建数据
func (r Notice) CreateNotice(ctx core.Context, notice *entity.Notice) (uint32, error) {
	return notice.Id, ctx.DB().Create(notice).Error
}

// DeleteNotice 删除数据
func (r Notice) DeleteNotice(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Notice{}).Error
}

// UpdateNotice 更新数据
func (r Notice) UpdateNotice(ctx core.Context, notice *entity.Notice) error {
	return ctx.DB().Updates(notice).Error
}

func (r Notice) IsExistNoticeByMd5(ctx core.Context, md5 string) bool {
	var count int64
	ctx.DB().Model(entity.Notice{}).Where("md5=?", md5).Count(&count)
	return count != 0
}

func (r Notice) AddUserNotice(ctx core.Context, uid, nid uint32) error {
	return ctx.DB().Create(&entity.NoticeUser{NoticeId: nid, UserId: uid}).Error
}
