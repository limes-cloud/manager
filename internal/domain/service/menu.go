package service

import (
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/samber/lo"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Menu struct {
	repo  repository.Menu
	app   repository.App
	rm    repository.RoleMenu
	scope repository.Scope
	ta    repository.TenantApp
}

func NewMenu(
	repo repository.Menu,
	app repository.App,
	rm repository.RoleMenu,
	scope repository.Scope,
	ta repository.TenantApp,
) *Menu {
	uc := &Menu{
		repo:  repo,
		app:   app,
		rm:    rm,
		scope: scope,
		ta:    ta,
	}
	return uc
}

// ListCurrentMenu 获取当前的菜单信息列表树
func (u *Menu) ListCurrentMenu(ctx core.Context, req *types.ListMenuRequest) ([]*entity.Menu, error) {
	// 获取当前的角色列表
	if ctx.IsSuperAdmin() && req.FilterTenant {
		req.InIds = u.ta.GetTenantMenuIds(ctx.Auth().TenantId)
	}
	if !ctx.IsSuperAdmin() {
		rids := u.scope.RoleScopes(ctx)
		mids := u.rm.GetMenuIdsByRoleIds(rids)
		tmids := u.ta.GetTenantMenuIds(ctx.Auth().TenantId)

		if req.AppId != nil {
			ids, err := u.ta.GetTenantAppMenuIds(ctx, ctx.Auth().TenantId, *req.AppId)
			if err != nil {
				return nil, errors.SystemError()
			}
			tmids = ids
		}
		req.InIds = lo.Intersect(mids, tmids)

	}

	// 获取当前角色有权限的菜单ID
	list, err := u.ListMenu(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "get menu ids error", "err", err.Error())
		return nil, errors.ListError()
	}

	// 构建树
	if req.FilterRoot {
		return list, nil
	}

	// 获取应用，设置为顶级菜单
	var (
		appIds []uint32
		mb     = map[uint32][]*entity.Menu{}
		nm     []*entity.Menu
	)
	for _, menu := range list {
		appIds = append(appIds, menu.AppId)
		mb[menu.AppId] = append(mb[menu.AppId], menu)
	}
	apps, _, err := u.app.ListApp(ctx, &types.ListAppRequest{
		InIds: lo.Uniq(appIds),
	})
	if err != nil {
		ctx.Logger().Warnw("msg", "list app error", "err", err.Error())
		return nil, errors.ListError()
	}

	for _, app := range apps {
		menu := entity.Menu{
			Title:     app.Name,
			Icon:      &app.Logo,
			Keyword:   &app.Keyword,
			Type:      entity.MenuTypeRoot,
			AppId:     app.Id,
			Children:  mb[app.Id],
			Component: value.Pointer("Layout"),
			Path:      value.Pointer("/" + app.Keyword),
		}
		nm = append(nm, &menu)
	}

	return nm, nil
}

// ListMenu 获取菜单信息列表树
func (u *Menu) ListMenu(ctx core.Context, req *types.ListMenuRequest) ([]*entity.Menu, error) {
	list, err := u.repo.ListMenu(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list menu error", "err", err.Error())
		return nil, errors.ListError()
	}

	var (
		gms = map[uint32]int{}
		mb  = map[uint32][]uint32{}
	)

	for index, item := range list {
		if item.Type == entity.MenuTypeGroup {
			gms[item.Id] = index
		}
		mb[item.ParentId] = append(mb[item.ParentId], item.Id)
	}

	// 如果是G，但是没有子节点，则去除G
	for id, ind := range gms {
		if len(mb[id]) == 0 {
			list = append(list[:ind], list[ind+1:]...)
		}
	}

	return tree.BuildArrayTree(list), nil
}

// CreateMenu 创建菜单信息
func (u *Menu) CreateMenu(ctx core.Context, menu *entity.Menu) (uint32, error) {
	var (
		id  uint32
		err error
	)

	// 判断是否具有应用权限
	if !u.scope.HasAppScope(ctx, menu.AppId) {
		return 0, errors.AppScopeError()
	}

	// 开启事务
	if err = ctx.Transaction(func(ctx core.Context) error {
		// 创建菜单
		id, err = u.repo.CreateMenu(ctx, menu)
		if err != nil {
			return err
		}

		// 设置为首页
		if menu.IsHome != nil && *menu.IsHome {
			if err := u.repo.SetHome(ctx, id); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateMenu 更新菜单信息
func (u *Menu) UpdateMenu(ctx core.Context, menu *entity.Menu) error {
	// 获取之前的菜单信息
	old, err := u.repo.GetMenu(ctx, menu.Id)
	if err != nil {
		return err
	}

	// 判断是否具有应用权限
	if !u.scope.HasAppScope(ctx, old.AppId) {
		return errors.AppScopeError()
	}

	// 置空应用ID,不允许修改为其他的应用ID
	menu.AppId = 0

	// 变更菜单
	if err := u.repo.UpdateMenu(ctx, menu); err != nil {
		return errors.UpdateError(err.Error())
	}

	return nil
}

// DeleteMenu 删除菜单信息
func (u *Menu) DeleteMenu(ctx core.Context, id uint32) error {
	// 获取之前的菜单信息
	old, err := u.repo.GetMenu(ctx, id)
	if err != nil {
		return err
	}

	// 判断是否具有应用权限
	if !u.scope.HasAppScope(ctx, old.AppId) {
		return errors.AppScopeError()
	}

	// 查询下级菜单
	mids, err := u.repo.GetMenuChildrenIds(ctx, id)
	if err != nil {
		return err
	}

	mids = append(mids, id)

	return ctx.Transaction(func(ctx core.Context) error {
		// 删除角色
		if err = u.repo.DeleteMenu(ctx, id); err != nil {
			return err
		}

		// 删除权限
		return u.rm.DeleteMenus(ctx, mids)
	})
}
