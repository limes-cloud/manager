package department

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/pkg/md"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// ListDepartment 获取部门信息列表树 fixed code
func (u *UseCase) ListDepartment(ctx kratosx.Context, req *ListDepartmentRequest) ([]tree.Tree, uint32, error) {
	all, scopes, err := u.repo.GetDepartmentDataScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, 0, errors.DatabaseError(err.Error())
	}
	if !all {
		req.Ids = scopes
	}
	list, total, err := u.repo.ListDepartment(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	var ts []tree.Tree
	for _, item := range list {
		ts = append(ts, item)
	}
	return tree.BuildArrayTree(ts), total, nil
}

// CreateDepartment 创建部门信息 fixed code
func (u *UseCase) CreateDepartment(ctx kratosx.Context, req *Department) (uint32, error) {
	all, scopes, err := u.repo.GetDepartmentDataScope(ctx, md.UserId(ctx))
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	if !all && !valx.InList(scopes, req.ParentId) {
		return 0, errors.DepartmentPermissionsError()
	}

	id, err := u.repo.CreateDepartment(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateDepartment 更新部门信息 fixed code
func (u *UseCase) UpdateDepartment(ctx kratosx.Context, req *Department) error {
	if req.Id == md.DepartmentId(ctx) {
		return errors.DepartmentPermissionsError()
	}

	all, scopes, err := u.repo.GetDepartmentDataScope(ctx, md.UserId(ctx))
	if err != nil {
		return errors.DatabaseError(err.Error())
	}

	if !all && (!valx.InList(scopes, req.ParentId) || !valx.InList(scopes, req.Id)) {
		return errors.DepartmentPermissionsError()
	}

	if err := u.repo.UpdateDepartment(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteDepartment 删除部门信息 fixed code
func (u *UseCase) DeleteDepartment(ctx kratosx.Context, ids []uint32) (uint32, error) {
	all, scopes, err := u.repo.GetDepartmentDataScope(ctx, md.UserId(ctx))
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}

	for _, id := range ids {
		if id == md.DepartmentId(ctx) {
			return 0, errors.DepartmentPermissionsError()
		}
		if !all && !valx.InList(scopes, id) {
			return 0, errors.DepartmentPermissionsError()
		}
	}

	total, err := u.repo.DeleteDepartment(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// GetDepartment 获取指定的部门信息
func (u *UseCase) GetDepartment(ctx kratosx.Context, req *GetDepartmentRequest) (*Department, error) {
	var (
		res *Department
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetDepartment(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetDepartmentByKeyword(ctx, *req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}
