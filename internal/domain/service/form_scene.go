package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type FormScene struct {
	repo repository.FormScene
}

func NewFormScene(repo repository.FormScene) *FormScene {
	return &FormScene{repo: repo}
}

// ListFormScene 查询场景列表
func (s *FormScene) ListFormScene(ctx core.Context, req *types.ListFormSceneRequest) ([]*entity.FormScene, uint32, error) {
	list, total, err := s.repo.ListFormScene(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list form scene error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateFormScene 创建场景
func (s *FormScene) CreateFormScene(ctx core.Context, ent *entity.FormScene) (uint32, error) {
	id, err := s.repo.CreateFormScene(ctx, ent)
	if err != nil {
		ctx.Logger().Warnw("msg", "create form scene error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateFormScene 更新场景
func (s *FormScene) UpdateFormScene(ctx core.Context, ent *entity.FormScene) error {
	if err := s.repo.UpdateFormScene(ctx, ent); err != nil {
		ctx.Logger().Warnw("msg", "update form scene error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// DeleteFormScene 删除场景
func (s *FormScene) DeleteFormScene(ctx core.Context, id uint32) error {
	if err := s.repo.DeleteFormScene(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete form scene error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}
