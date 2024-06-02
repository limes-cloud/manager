package role

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

// ListRole 获取角色信息列表树 fixed code
func (u *UseCase) ListRole(ctx kratosx.Context, req *ListRoleRequest) ([]tree.Tree, uint32, error) {
	all, scopes, err := u.repo.GetRoleDataScope(ctx, md.RoleId(ctx))
	if err != nil {
		return nil, 0, err
	}
	if !all {
		req.Ids = scopes
	}

	list, total, err := u.repo.ListRole(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	var ts []tree.Tree
	for _, item := range list {
		ts = append(ts, item)
	}
	return tree.BuildArrayTree(ts), total, nil
}

// CreateRole 创建角色信息 fixed code
func (u *UseCase) CreateRole(ctx kratosx.Context, req *Role) (uint32, error) {
	all, scopes, err := u.repo.GetRoleDataScope(ctx, md.RoleId(ctx))
	if err != nil {
		return 0, err
	}
	if !all && !valx.InList(scopes, req.ParentId) {
		return 0, errors.RolePermissionsError()
	}

	id, err := u.repo.CreateRole(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateRole 更新角色信息 fixed code
func (u *UseCase) UpdateRole(ctx kratosx.Context, req *Role) error {
	if req.Id == 1 {
		return errors.EditSystemDataError()
	}

	all, scopes, err := u.repo.GetRoleDataScope(ctx, md.RoleId(ctx))
	if err != nil {
		return err
	}
	if !all && (!valx.InList(scopes, req.Id) || !valx.InList(scopes, req.ParentId)) {
		return errors.RolePermissionsError()
	}

	if err := u.repo.UpdateRole(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateRoleStatus 更新角色信息状态 fixed code
func (u *UseCase) UpdateRoleStatus(ctx kratosx.Context, id uint32, status bool) error {
	if id == 1 {
		return errors.EditSystemDataError()
	}

	all, scopes, err := u.repo.GetRoleDataScope(ctx, md.RoleId(ctx))
	if err != nil {
		return err
	}
	if !all && (!valx.InList(scopes, id)) {
		return errors.RolePermissionsError()
	}

	if err := u.repo.UpdateRoleStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteRole 删除角色信息 fixed code
func (u *UseCase) DeleteRole(ctx kratosx.Context, ids []uint32) (uint32, error) {
	curRoleId := md.RoleId(ctx)
	all, scopes, err := u.repo.GetRoleDataScope(ctx, curRoleId)
	if err != nil {
		return 0, err
	}

	for _, id := range scopes {
		if id == 1 {
			return 0, errors.DeleteSystemDataError()
		}
		if curRoleId == id {
			return 0, errors.RolePermissionsError()
		}
		if !all && (!valx.InList(scopes, id)) {
			return 0, errors.RolePermissionsError()
		}
	}

	total, err := u.repo.DeleteRole(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// GetRoleMenuIds 获取指定角色的菜单id列表
func (u *UseCase) GetRoleMenuIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	list, err := u.repo.GetRoleMenuIds(ctx, id)
	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	return list, err
}

// UpdateRoleMenu 更新角色的菜单列表
func (u *UseCase) UpdateRoleMenu(ctx kratosx.Context, roleId uint32, menuIds []uint32) error {
	if roleId == 1 {
		return errors.EditSystemDataError()
	}
	curRoleId := md.RoleId(ctx)
	if roleId == curRoleId {
		return errors.RolePermissionsError()
	}

	all, scopes, err := u.repo.GetRoleDataScope(ctx, curRoleId)
	if err != nil {
		return err
	}
	if !all && !valx.InList(scopes, roleId) {
		return errors.RolePermissionsError()
	}

	rids, err := u.repo.GetRoleChildrenIds(ctx, curRoleId)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	if !valx.InList(rids, roleId) {
		return errors.RolePermissionsError()
	}

	if curRoleId != 1 {
		curMenuIds, err := u.repo.GetRoleMenuIds(ctx, curRoleId)
		if err != nil {
			return errors.DatabaseError(err.Error())
		}
		for _, id := range menuIds {
			if !valx.InList(curMenuIds, id) {
				return errors.MenuPermissionsError()
			}
		}
	}

	if err := u.repo.UpdateRoleMenu(ctx, roleId, menuIds); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// GetRole 获取指定的角色信息
func (u *UseCase) GetRole(ctx kratosx.Context, req *GetRoleRequest) (*Role, error) {
	var (
		res *Role
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetRole(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetRoleByKeyword(ctx, *req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}
