package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type OAuther struct {
	repo repository.OAuther
	exec repository.OAuthExecer
}

func NewOAuther(ch repository.OAuther, exec repository.OAuthExecer) *OAuther {
	return &OAuther{
		repo: ch,
		exec: exec,
	}
}

// ListOAutherTypes 获取可以开通的渠道列表
func (u *OAuther) ListOAutherTypes() []*types.OAutherType {
	return u.exec.ListOAutherType()
}

// ListOAuther 获取授权渠道列表
func (u *OAuther) ListOAuther(ctx core.Context, req *types.ListOAutherRequest) ([]*entity.OAuther, uint32, error) {
	list, total, err := u.repo.ListOAuther(ctx, req)
	if err != nil {
		ctx.Logger().Errorw("msg", "list oauther error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateOAuther 创建授权渠道
func (u *OAuther) CreateOAuther(ctx core.Context, oauther *entity.OAuther) (uint32, error) {
	oauther.Status = proto.Bool(false)
	id, err := u.repo.CreateOAuther(ctx, oauther)
	if err != nil {
		ctx.Logger().Errorw("msg", "create oauther error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateOAuther 更新授权渠道
func (u *OAuther) UpdateOAuther(ctx core.Context, oauther *entity.OAuther) error {
	if err := u.repo.UpdateOAuther(ctx, oauther); err != nil {
		ctx.Logger().Errorw("msg", "update oauther error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteOAuther 删除授权渠道
func (u *OAuther) DeleteOAuther(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteOAuther(ctx, id); err != nil {
		ctx.Logger().Errorw("msg", "delete oauther error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
