package dbs

import (
	"fmt"
	"sync"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Channel struct {
}

var (
	channelIns  *Channel
	channelOnce sync.Once
)

func NewChannel() *Channel {
	channelOnce.Do(func() {
		channelIns = &Channel{}
	})
	return channelIns
}

// GetChannelByKeyword 获取指定数据
func (r Channel) GetChannelByKeyword(ctx kratosx.Context, keyword string) (*entity.Channel, error) {
	var (
		channel = entity.Channel{}
		fs      = []string{"*"}
	)
	return &channel, ctx.DB().Select(fs).Where("keyword = ?", keyword).First(&channel).Error
}

// GetChannel 获取指定的数据
func (r Channel) GetChannel(ctx kratosx.Context, id uint32) (*entity.Channel, error) {
	var (
		channel = entity.Channel{}
		fs      = []string{"*"}
	)
	return &channel, ctx.DB().Select(fs).First(&channel, id).Error
}

// ListChannel 获取列表
func (r Channel) ListChannel(ctx kratosx.Context, req *types.ListChannelRequest) ([]*entity.Channel, uint32, error) {
	var (
		list  []*entity.Channel
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Channel{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
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

	return list, uint32(total), db.Find(&list).Error
}

// CreateChannel 创建数据
func (r Channel) CreateChannel(ctx kratosx.Context, channel *entity.Channel) (uint32, error) {
	return channel.Id, ctx.DB().Create(channel).Error
}

// UpdateChannel 更新数据
func (r Channel) UpdateChannel(ctx kratosx.Context, channel *entity.Channel) error {
	return ctx.DB().Updates(channel).Error
}

// DeleteChannel 删除数据
func (r Channel) DeleteChannel(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Channel{}).Error
}
