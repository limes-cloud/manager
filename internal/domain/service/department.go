package service

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/pkg/md"
	"github.com/limes-cloud/manager/internal/types"
)

type Department struct {
	conf *conf.Config
	repo repository.Department
	role repository.Role
}

func NewDepartment(config *conf.Config, repo repository.Department, role repository.Role) *Department {
	return &Department{conf: config, repo: repo, role: role}
}

// ListDepartmentClassify 获取部门分类列表
func (t *Department) ListDepartmentClassify(ctx kratosx.Context, req *types.ListDepartmentClassifyRequest) ([]*entity.DepartmentClassify, uint32, error) {
	list, total, err := t.repo.ListDepartmentClassify(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateDepartmentClassify 创建部门分类
func (t *Department) CreateDepartmentClassify(ctx kratosx.Context, tg *entity.DepartmentClassify) (uint32, error) {
	id, err := t.repo.CreateDepartmentClassify(ctx, tg)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateDepartmentClassify 更新部门分类
func (t *Department) UpdateDepartmentClassify(ctx kratosx.Context, tg *entity.DepartmentClassify) error {
	if err := t.repo.UpdateDepartmentClassify(ctx, tg); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteDepartmentClassify 删除部门分类
func (t *Department) DeleteDepartmentClassify(ctx kratosx.Context, id uint32) error {
	err := t.repo.DeleteDepartmentClassify(ctx, id)
	if err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListDepartment 获取部门信息列表树
func (u *Department) ListDepartment(ctx kratosx.Context, req *types.ListDepartmentRequest) ([]*entity.Department, error) {
	// 获取部门列表
	list, err := u.repo.ListDepartment(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list department error", "err", err.Error())
		return nil, errors.ListError()
	}

	return tree.BuildArrayTree(list), nil
}

// ListCurrentDepartment 获取当前用户的部门信息列表树
func (u *Department) ListCurrentDepartment(ctx kratosx.Context, req *types.ListDepartmentRequest) ([]*entity.Department, error) {
	// 获取当前用户的数据权限列表
	all, scopes, err := u.role.GetDataScope(ctx, md.UserId(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "list department error", "err", err.Error())
		return nil, errors.DatabaseError()
	}

	// 通过指定权限列表的部门
	if !all {
		// 没有权限时，默认可查看当前所属部门
		if len(scopes) == 0 {
			scopes = md.DepartmentIds(ctx)
		}
		req.Ids = scopes
	}
	return u.ListDepartment(ctx, req)
}

// CreateDepartment 创建部门信息
func (u *Department) CreateDepartment(ctx kratosx.Context, req *entity.Department) (uint32, error) {
	// 是否具有父级部门权限
	hasPurview, err := u.role.HasDataPurview(ctx, md.UserId(ctx), []uint32{req.ParentId})
	if err != nil {
		ctx.Logger().Warnw("msg", "get department repo error", "err", err.Error())
		return 0, errors.DatabaseError()
	}

	// 权限判断
	if !hasPurview {
		return 0, errors.DepartmentPurviewError()
	}

	// 是否具有角色权限
	has, err := u.role.HasRolePurview(ctx, md.UserId(ctx), req.GetRoleIds())
	if err != nil {
		return 0, errors.DatabaseError()
	}
	if !has {
		return 0, errors.RolePurviewError()
	}

	// 新增数据
	id, err := u.repo.CreateDepartment(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create department error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateDepartment 更新部门信息
func (u *Department) UpdateDepartment(ctx kratosx.Context, req *entity.Department) error {
	if req.Id == 1 {
		return errors.EditSystemDataError()
	}

	// 获取历史信息
	oldDept, err := u.repo.GetDepartment(ctx, req.Id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get base dept error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 是否具有部门数据权限
	has, err := u.role.HasDataPurview(ctx, md.UserId(ctx), []uint32{req.Id, req.ParentId})
	if err != nil {
		ctx.Logger().Warnw("msg", "get department repo error", "err", err.Error())
		return errors.DatabaseError()
	}
	// 权限判断
	if !has {
		return errors.DepartmentPurviewError()
	}

	// 获取当前的角色权限
	all, scopes, err := u.role.GetRoleScope(ctx, md.RoleIds(ctx))
	if !all {
		var inr = valx.New(scopes)
		// 判断是否具有修改后的角色权限
		for _, id := range req.GetRoleIds() {
			if !inr.Has(id) {
				return errors.RolePurviewError()
			}
		}

		// 获取修改的用户角色在当前用户不可见的id
		// 将不可见的ID追加到新建里面
		oldRoleIds := oldDept.GetRoleIds()
		for _, id := range oldRoleIds {
			if !inr.Has(id) {
				req.DepartmentRoles = append(req.DepartmentRoles, &entity.DepartmentRole{
					DepartmentId: req.Id,
					RoleId:       id,
				})
			}
		}
	}

	// 更新数据
	if err := u.repo.UpdateDepartment(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteDepartment 删除部门信息
func (u *Department) DeleteDepartment(ctx kratosx.Context, id uint32) error {
	if id == 1 {
		return errors.DeleteSystemDataError()
	}

	// 不能删除当前部门
	if valx.InList(md.DepartmentIds(ctx), id) {
		return errors.DepartmentPurviewError()
	}

	// 获取是否具有部门权限
	hasPurview, err := u.role.HasDataPurview(ctx, md.UserId(ctx), []uint32{id})
	if err != nil {
		ctx.Logger().Warnw("msg", "get department repo error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 判定权限
	if !hasPurview {
		return errors.DepartmentPurviewError()
	}

	// 更新数据
	if err := u.repo.DeleteDepartment(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// GetDepartment 获取指定的部门信息
func (u *Department) GetDepartment(ctx kratosx.Context, id uint32) (*entity.Department, error) {
	res, err := u.repo.GetDepartment(ctx, id)

	if err != nil {
		ctx.Logger().Warnw("msg", "get department error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// GetDepartmentByKeyword 获取指定的部门信息
func (u *Department) GetDepartmentByKeyword(ctx kratosx.Context, keyword string) (*entity.Department, error) {
	res, err := u.repo.GetDepartmentByKeyword(ctx, keyword)
	if err != nil {
		ctx.Logger().Warnw("msg", "get department error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}
