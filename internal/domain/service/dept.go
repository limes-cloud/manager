package service

import (
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Dept struct {
	repo  repository.Dept
	scope repository.Scope
}

func NewDept(
	repo repository.Dept,
	scope repository.Scope,
) *Dept {
	return &Dept{
		repo:  repo,
		scope: scope,
	}
}

// ListDeptClassify 获取租户列表
func (u *Dept) ListDeptClassify(ctx core.Context, req *types.ListDeptClassifyRequest) ([]*entity.DeptClassify, uint32, error) {
	list, total, err := u.repo.ListDeptClassify(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list dept classify error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateDeptClassify 创建租户
func (u *Dept) CreateDeptClassify(ctx core.Context, req *entity.DeptClassify) (uint32, error) {
	id, err := u.repo.CreateDeptClassify(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create dept classify error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateDeptClassify 更新租户
func (u *Dept) UpdateDeptClassify(ctx core.Context, req *entity.DeptClassify) error {
	if err := u.repo.UpdateDeptClassify(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update dept classify error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteDeptClassify 删除租户
func (u *Dept) DeleteDeptClassify(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteDeptClassify(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete dept classify error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// GetDept 获取租户列表
func (u *Dept) GetDept(ctx core.Context, id uint32) (*entity.Dept, error) {
	res, err := u.repo.GetDept(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// ListDept 获取租户列表
func (u *Dept) ListDept(ctx core.Context, req *types.ListDeptRequest) ([]*entity.Dept, error) {
	list, err := u.repo.ListDept(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list dept error", "err", err.Error())
		return nil, errors.ListError(err.Error())
	}

	return tree.BuildArrayTree(list), nil
}

// ListCurrentDept 获取当前用户的部门信息列表树
func (u *Dept) ListCurrentDept(ctx core.Context, req *types.ListDeptRequest) ([]*entity.Dept, error) {
	return u.ListDept(ctx, req)
}

// CreateDept 创建租户
func (u *Dept) CreateDept(ctx core.Context, req *entity.Dept) (uint32, error) {
	id, err := u.repo.CreateDept(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create dept error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateDept 更新租户
func (u *Dept) UpdateDept(ctx core.Context, req *entity.Dept) error {
	if err := u.repo.UpdateDept(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update dept error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteDept 删除租户
func (u *Dept) DeleteDept(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteDept(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete dept error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
