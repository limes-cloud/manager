package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type AppForm struct {
	repo     repository.AppForm
	formRepo repository.UserForm
}

func NewAppForm(repo repository.AppForm, formRepo repository.UserForm) *AppForm {
	return &AppForm{repo: repo, formRepo: formRepo}
}

// ListAppForm 查询某应用下的表单关联列表
func (s *AppForm) ListAppForm(ctx core.Context, req *types.ListAppFormRequest) ([]*entity.AppForm, error) {
	list, err := s.repo.ListAppForm(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list app form error", "err", err.Error())
		return nil, errors.ListError()
	}
	return list, nil
}

// CreateAppForm 分配表单到应用
func (s *AppForm) CreateAppForm(ctx core.Context, ent *entity.AppForm) (uint32, error) {
	// 从 user_form 获取 scene 冗余字段
	uf, err := s.formRepo.GetUserForm(ctx, &types.GetUserFormRequest{Id: &ent.FormId})
	if err != nil {
		ctx.Logger().Warnw("msg", "get user form for app form error", "err", err.Error())
		return 0, errors.GetError()
	}
	ent.Scene = uf.Scene

	id, err := s.repo.CreateAppForm(ctx, ent)
	if err != nil {
		ctx.Logger().Warnw("msg", "create app form error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// DeleteAppForm 取消分配
func (s *AppForm) DeleteAppForm(ctx core.Context, id uint32) error {
	if err := s.repo.DeleteAppForm(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete app form error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}

// GetUserFormByApp 通过 app_keyword + scene 获取表单
func (s *AppForm) GetUserFormByApp(ctx core.Context, appKeyword string, scene string) (*entity.UserForm, error) {
	// 先通过 app_form 找到 form_id
	af, err := s.repo.GetAppForm(ctx, appKeyword, scene)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app form error", "appKeyword", appKeyword, "scene", scene, "err", err.Error())
		return nil, errors.GetError()
	}

	// 再获取 user_form
	uf, err := s.formRepo.GetUserForm(ctx, &types.GetUserFormRequest{Id: &af.FormId})
	if err != nil {
		ctx.Logger().Warnw("msg", "get user form by app error", "formId", af.FormId, "err", err.Error())
		return nil, errors.GetError()
	}
	return uf, nil
}
