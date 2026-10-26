package service

import (
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/samber/lo"
	lom "github.com/samber/lo/mutable"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type RoleMenu struct {
	repo  repository.RoleMenu
	app   repository.App
	scope repository.Scope
	menu  repository.Menu
	role  repository.Role
}

func NewRoleMenu(
	repo repository.RoleMenu,
	app repository.App,
	scope repository.Scope,
	menu repository.Menu,
	role repository.Role,
) *RoleMenu {
	return &RoleMenu{
		app:   app,
		repo:  repo,
		scope: scope,
		menu:  menu,
		role:  role,
	}
}

// GetRoleMenuIds 获取角色的菜单id列表
func (rm *RoleMenu) GetRoleMenuIds(ctx core.Context, req *types.GetRoleMenuIdsRequest) ([]uint32, error) {
	app := ctx.AppKey(value.Value(req.App))
	mrs := rm.repo.GetMenuIdsByRoleIds(ctx, app, []uint32{req.RoleId})
	appMenuIds, err := rm.menu.GetMenuIdsByApp(ctx, app)
	if err != nil {
		return nil, errors.SystemError()
	}
	return lo.Intersect(mrs, appMenuIds), nil
}

// GetMenuRoleIds 获取指定菜单的所有角色ID
func (rm *RoleMenu) GetMenuRoleIds(ctx core.Context, req *types.GetMenuRoleIdsRequest) ([]uint32, error) {
	app := ctx.AppKey()
	rids := rm.scope.RoleScopes(ctx)
	mrs := rm.repo.GetRoleIdsByMenuIds(ctx, app, []uint32{req.MenuId})
	// 验证当前角色是否具有指定菜单的权限
	mrs = lom.Filter(mrs, func(item uint32) bool {
		return lo.Contains(rids, item)
	})
	return mrs, nil
}

// CreateMenuRoles 菜单批量授权给角色
func (rm *RoleMenu) CreateMenuRoles(ctx core.Context, req *types.CreateMenuRolesRequest) error {
	app := ctx.AppKey()
	rids := rm.scope.RoleIds(ctx)
	if len(lo.Intersect(rids, req.RoleIds)) != len(req.RoleIds) {
		return errors.RoleScopeError()
	}
	mids := rm.repo.GetMenuIdsByRoleIds(ctx, app, rids)
	if !lo.Contains(mids, req.MenuId) {
		return errors.MenuScopeError()
	}

	// 组装为实体列表
	var rms []*entity.RoleMenu
	for _, v := range req.RoleIds {
		rms = append(rms, &entity.RoleMenu{
			App:    app,
			RoleId: v,
			MenuId: req.MenuId,
		})
	}
	if err := rm.repo.CreateRoleMenus(ctx, rms); err != nil {
		return errors.CreateError()
	}
	return nil
}

// CreateRoleMenus 角色批量授权给菜单
func (rm *RoleMenu) CreateRoleMenus(ctx core.Context, req *types.CreateRoleMenusRequest) error {
	// req.App 是「授权目标模块」这个业务参数，取值范围由当前应用的模块订阅约束（dbs 层 validate 校验）
	app := ctx.AppKey(req.App)

	// 被授权角色必须在当前操作者的角色管辖范围内，防止越权授予非管辖角色
	if !ctx.Auth().IsAdmin && !lo.Contains(rm.scope.RoleScopes(ctx), req.RoleId) {
		return errors.RoleScopeError()
	}

	// 获取目标模块下的所有菜单
	appMenuIds, err := rm.menu.GetMenuIdsByApp(ctx, app)
	if err != nil {
		return errors.SystemError()
	}

	// 计算可授予的菜单上限：超管为模块全量，普通用户为自身（含子角色）已有的菜单
	// 否则操作者可以把自己没有的权限授予他人，形成提权路径
	roleMenuIds := appMenuIds
	if !ctx.Auth().IsAdmin {
		myMenuIds := rm.repo.GetMenuIdsByRoleIds(ctx, app, rm.scope.RoleScopes(ctx))
		roleMenuIds = lo.Intersect(appMenuIds, myMenuIds)
	}

	// 获取修改的角色的菜单权限
	editRoleMenuIds := rm.repo.GetMenuIdsByRoleIds(ctx, app, []uint32{req.RoleId})

	// 剔除当前角色无权限的菜单
	editRoleMenuIds = lo.Intersect(editRoleMenuIds, roleMenuIds)
	req.MenuIds = lo.Intersect(req.MenuIds, roleMenuIds)

	// 获取增加、删除的菜单
	rmIds, addIds := lo.Difference(editRoleMenuIds, req.MenuIds)

	// 获取当前角色到被修改角色的节点，同步添加被增加的菜单
	// 获取当前的角色下的所有子角色id
	childRoles, err := rm.role.GetRoleChildrenIds(ctx, rm.scope.RoleScopes(ctx))
	if err != nil {
		return errors.SystemError()
	}
	// 获取被修改角色所有父角色id
	parentRoles, err := rm.role.GetRoleParentIds(ctx, []uint32{req.RoleId})
	if err != nil {
		return errors.SystemError()
	}
	// 获取当前角色到被修改角色的节点，同步添加被增加的菜单
	addRoles := lo.Intersect(childRoles, parentRoles)
	addRoles = append(addRoles, req.RoleId)

	// 删除当前角色菜单
	err = ctx.Transaction(func(ctx core.Context) error {
		if len(rmIds) != 0 {
			var rms []*entity.RoleMenu
			for _, v := range rmIds {
				rms = append(rms, &entity.RoleMenu{
					App:    app,
					RoleId: req.RoleId,
					MenuId: v,
				})
			}
			if err := rm.repo.DeleteRoleMenus(ctx, rms); err != nil {
				return err
			}
		}
		if len(addIds) != 0 {
			var rms []*entity.RoleMenu
			for _, v := range addIds {
				for _, roleId := range addRoles {
					rms = append(rms, &entity.RoleMenu{
						App:    app,
						RoleId: roleId,
						MenuId: v,
					})
				}
			}
			if err := rm.repo.CreateRoleMenus(ctx, rms); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return errors.CreateError()
	}
	return nil
}

func (rm *RoleMenu) DeleteRoleMenus(ctx core.Context, req *types.DeleteRoleMenusRequest) error {
	app := ctx.AppKey()
	rids := rm.scope.RoleIds(ctx)
	if !lo.Contains(rids, req.RoleId) {
		return errors.RoleScopeError()
	}
	mids := rm.repo.GetMenuIdsByRoleIds(ctx, app, rids)
	if len(lo.Intersect(mids, req.MenuIds)) != len(req.MenuIds) {
		return errors.MenuScopeError()
	}

	// 组装数据
	var rms []*entity.RoleMenu
	for _, v := range req.MenuIds {
		rms = append(rms, &entity.RoleMenu{
			App:    app,
			RoleId: req.RoleId,
			MenuId: v,
		})
	}
	if err := rm.repo.DeleteRoleMenus(ctx, rms); err != nil {
		return errors.DeleteError()
	}
	return nil
}

func (rm *RoleMenu) DeleteMenuRoles(ctx core.Context, req *types.DeleteMenuRolesRequest) error {
	app := ctx.AppKey()
	rids := rm.scope.RoleIds(ctx)
	if len(lo.Intersect(rids, req.RoleIds)) != len(req.RoleIds) {
		return errors.RoleScopeError()
	}
	mids := rm.repo.GetMenuIdsByRoleIds(ctx, app, rids)
	if !lo.Contains(mids, req.MenuId) {
		return errors.MenuScopeError()
	}

	// 组装数据
	var rms []*entity.RoleMenu
	for _, v := range req.RoleIds {
		rms = append(rms, &entity.RoleMenu{
			App:    app,
			RoleId: v,
			MenuId: req.MenuId,
		})
	}
	if err := rm.repo.DeleteRoleMenus(ctx, rms); err != nil {
		return errors.DeleteError()
	}
	return nil
}
