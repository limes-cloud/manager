package role

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/config"
	"github.com/limes-cloud/manager/internal/pkg/md"
)

type UseCase struct {
	repo Repo
	conf *config.Config
}

func NewUseCase(conf *config.Config, repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
		conf: conf,
	}
}

// RoleTree 查询当前用户的角色树
func (u *UseCase) RoleTree(ctx kratosx.Context) (tree.Tree, error) {
	// 查询子角色id列表信息
	roleId := md.RoleId(ctx)
	list, err := u.repo.AllRole(ctx, &AllRoleRequest{ParentId: proto.Uint32(roleId)})
	if err != nil {
		return nil, err
	}

	// 构造角色树
	var tl []tree.Tree
	for _, item := range list {
		if item.ID() == roleId {
			item.ParentId = 0
		}
		tl = append(tl, item)
	}
	return tree.BuildTree(tl), nil
}

func (u *UseCase) AddRole(ctx kratosx.Context, in *Role) (uint32, error) {
	roleId := md.RoleId(ctx)

	ids, err := u.repo.GetChildrenIds(ctx, roleId)
	if err != nil {
		return 0, errors.Database()
	}
	if !util.InList(ids, in.ParentId) {
		return 0, errors.RolePermissions()
	}

	// 创建角色信息
	id, err := u.repo.AddRole(ctx, in)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}

	return id, nil
}

func (u *UseCase) UpdateRole(ctx kratosx.Context, in *Role) error {
	// 超级管理员不允许修改
	if in.ID() == 1 {
		return errors.EditSystemData()
	}

	//  不能修改当前角色
	roleId := md.RoleId(ctx)
	if in.ID() == md.RoleId(ctx) {
		return errors.RolePermissions()
	}

	// 判断当前部门是否具有权限
	ids, err := u.repo.GetChildrenIds(ctx, roleId)
	if err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	if !util.InList(ids, in.ID()) {
		return errors.RolePermissions()
	}

	// 查询历史角色信息
	oldRole, err := u.repo.GetRole(ctx, roleId)
	if err != nil {
		return errors.NotFound()
	}

	// 如果存在移动角色，判断是否有父级部门的权限
	if oldRole.ParentId != in.ParentId && !util.InList(ids, in.ParentId) {
		return errors.RolePermissions()
	}

	// 更新角色信息
	if err := u.repo.UpdateRole(ctx, in); err != nil {
		return errors.DatabaseFormat(err.Error())
	}

	return nil
}

func (u *UseCase) DeleteRole(ctx kratosx.Context, id uint32) error {
	// 超级管理员不允许删除
	if id == 1 {
		return errors.DeleteSystemData()
	}

	// 不能删除当前角色
	roleId := md.RoleId(ctx)
	if id == roleId {
		return errors.DeleteSelfRole()
	}

	// 获取角色信息
	role, err := u.repo.GetRole(ctx, id)
	if err != nil {
		return errors.NotFound()
	}

	// 是否具有角色管理权限
	ids, err := u.repo.GetChildrenIds(ctx, roleId)
	if err != nil {
		return errors.Database()
	}
	if !util.InList(ids, id) {
		return errors.RolePermissions()
	}

	// 删除角色
	if err := u.repo.DeleteRole(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}

	// 删除rbac
	_, _ = ctx.Authentication().Enforce().RemoveFilteredPolicy(0, role.Keyword)
	return nil
}

// GetRoleMenuIds 获取指定角色的菜单ID
func (u *UseCase) GetRoleMenuIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	// 判断是否具有此角色的权限
	roleId := md.RoleId(ctx)
	rids, err := u.repo.GetChildrenIds(ctx, roleId)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	if !util.InList(rids, id) {
		return nil, errors.RolePermissions()
	}

	ids, err := u.repo.GetRoleMenuIds(ctx, id)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}

	return ids, nil
}

func (u *UseCase) UpdateRoleMenus(ctx kratosx.Context, roleId uint32, menuIds []uint32) error {
	// 超级管理员不存在菜单权限，自动获取全部菜单，所以禁止修改
	if roleId == 1 {
		return errors.EditSystemData()
	}

	// 判断是否具有此角色的权限
	curRoleId := md.RoleId(ctx)
	rids, _ := u.repo.GetChildrenIds(ctx, curRoleId)
	if !util.InList(rids, roleId) {
		return errors.RolePermissions()
	}

	// 不是超级管理员，则需要判断菜单是否都合法
	if curRoleId != 1 {
		// 获取当前用户的菜单id
		ids, err := u.repo.GetRoleMenuIds(ctx, curRoleId)
		if err != nil {
			return errors.DatabaseFormat(err.Error())
		}
		for _, id := range menuIds {
			if !util.InList(ids, id) {
				return errors.MenuPermissionsFormat(fmt.Sprintf("menu_id=%d", id))
			}
		}
	}

	// 获取当前role的数据
	role, err := u.repo.GetRole(ctx, roleId)
	if err != nil {
		return errors.DatabaseFormat(err.Error())
	}

	// 获取新的菜单信息
	var policies [][]string
	apiList, _ := u.repo.AllMenuApiByIds(ctx, menuIds)
	for _, item := range apiList {
		policies = append(policies, []string{role.Keyword, item.Api, item.Method})
	}

	// 进行菜单变更
	if err = u.repo.UpdateRoleMenus(ctx, roleId, menuIds); err != nil {
		return errors.DatabaseFormat(err.Error())
	}

	// 删除当前用户的全部rbac权限
	enforce := ctx.Authentication().Enforce()
	_, _ = enforce.RemoveFilteredPolicy(0, role.Keyword)

	// 添加新的rbac权限信息
	_, _ = enforce.AddPolicies(policies)

	return nil
}
