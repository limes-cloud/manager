package service

import (
	"fmt"
	"strings"

	"github.com/limes-cloud/kratosx/library/db/gormtranserror"
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/samber/lo"
	"gorm.io/gorm"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Menu struct {
	repo  repository.Menu
	app   repository.App
	rm    repository.RoleMenu
	scope repository.Scope
}

const (
	RegisterMenuModeAppend = "append"
	RegisterMenuModeSync   = "sync"
)

type registerMenuState struct {
	App     string
	SeenKey map[string]struct{}
	SeenApi map[string]struct{}
	Reply   *types.RegisterMenuReply
}

func NewMenu(
	repo repository.Menu,
	app repository.App,
	rm repository.RoleMenu,
	scope repository.Scope,
) *Menu {
	uc := &Menu{
		repo:  repo,
		app:   app,
		rm:    rm,
		scope: scope,
	}
	return uc
}

// ListCurrentMenu 获取当前的菜单信息列表树
func (u *Menu) ListCurrentMenu(ctx core.Context, req *types.ListMenuRequest) ([]*entity.Menu, error) {
	info := ctx.Auth()

	// 超级管理员直接查全量菜单，不经过角色过滤
	if !info.IsAdmin {
		rids := u.scope.RoleIds(ctx)
		mids := u.rm.GetMenuIdsByRoleIds(ctx, ctx.AppKey(), rids)
		req.InIds = lo.Intersect(mids, req.InIds)
	}

	// 获取当前角色有权限的菜单ID
	list, err := u.ListMenu(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "get menu ids error", "err", err.Error())
		return nil, errors.ListError()
	}

	// 构建树
	if req.OnlyMenu {
		return list, nil
	}

	// 获取应用，设置为顶级菜单
	var (
		mb     = map[string][]*entity.Menu{}
		nm     []*entity.Menu
	)
	for _, menu := range list {
		mb[menu.App] = append(mb[menu.App], menu)
	}
	listApps, _, err := u.app.ListApp(ctx, &types.ListAppRequest{})
	if err != nil {
		ctx.Logger().Warnw("msg", "list app error", "err", err.Error())
		return nil, errors.ListError()
	}

	for _, app := range listApps {
		menu := entity.Menu{
			Title:     app.Name,
			Icon:      &app.Logo,
			Keyword:   &app.Keyword,
			Type:      entity.MenuTypeRoot,
			App:       app.Keyword,
			Children:  mb[app.Keyword],
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
		pid := value.Value(item.ParentId)
		mb[pid] = append(mb[pid], item.Id)
	}

	// 如果是G，但是没有子节点，则去除G
	// 注：必须按 ID 收集后统一重建，或倒序删除；直接用 range gms 的索引边遍历边删
	// 会在存在≥2个孤立G时 slice 越界（[ind:len] 超过收缩后的长度）
	var orphanG []uint32
	for id, ind := range gms {
		if len(mb[id]) == 0 {
			id := list[ind].Id
			orphanG = append(orphanG, id)
		}
	}
	if len(orphanG) > 0 {
		filtered := make([]*entity.Menu, 0, len(list)-len(orphanG))
		for _, item := range list {
			if !value.InList(orphanG, item.Id) {
				filtered = append(filtered, item)
			}
		}
		list = filtered
	}

	return tree.BuildArrayTree(list), nil
}

// CreateMenu 创建菜单信息
func (u *Menu) CreateMenu(ctx core.Context, menu *entity.Menu) (uint32, error) {
	var (
		id  uint32
		err error
	)

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
		ctx.Logger().Warnw("msg", "create menu error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateMenu 更新菜单信息
func (u *Menu) UpdateMenu(ctx core.Context, menu *entity.Menu) error {
	// 变更菜单
	if err := u.repo.UpdateMenu(ctx, menu); err != nil {
		ctx.Logger().Warnw("msg", "update menu error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteMenu 删除菜单信息
func (u *Menu) DeleteMenu(ctx core.Context, id uint32) error {
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
		return u.rm.DeleteMenus(ctx, ctx.AppKey(), mids)
	})
}

// RegisterMenu 注册应用菜单树
func (u *Menu) RegisterMenu(ctx core.Context, req *types.RegisterMenuRequest) (*types.RegisterMenuReply, error) {
	app, err := u.app.GetAppByKeyword(req.AppKeyword)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError("应用不存在")
	}

	mode := strings.ToLower(strings.TrimSpace(req.Mode))
	if mode == "" {
		mode = RegisterMenuModeAppend
	}
	if mode != RegisterMenuModeAppend && mode != RegisterMenuModeSync {
		return nil, errors.ParamsError("mode仅支持append或sync")
	}

	state := &registerMenuState{
		App:     app.Keyword,
		SeenKey: map[string]struct{}{},
		SeenApi: map[string]struct{}{},
		Reply:   &types.RegisterMenuReply{},
	}
	if err = ctx.Transaction(func(ctx core.Context) error {
		for _, item := range req.Menus {
			if err := u.registerMenuNode(ctx, state, 0, item); err != nil {
				return err
			}
		}

		if mode != RegisterMenuModeSync {
			return nil
		}
		return u.deleteUnregisteredMenus(ctx, state)
	}); err != nil {
		ctx.Logger().Warnw("msg", "register menu error", "err", err.Error())
		return nil, errors.UpdateError(err.Error())
	}

	return state.Reply, nil
}

func (u *Menu) registerMenuNode(
	ctx core.Context,
	state *registerMenuState,
	parentId uint32,
	node *types.RegisterMenuNode,
) error {
	if node == nil {
		return nil
	}

	menu, err := u.toRegisterMenu(state.App, parentId, node)
	if err != nil {
		return err
	}

	old, err := u.matchRegisterMenu(ctx, state, menu)
	if err != nil {
		return err
	}
	if err := u.saveRegisterMenu(ctx, menu, old, state.Reply); err != nil {
		return err
	}

	for _, child := range node.Children {
		if err := u.registerMenuNode(ctx, state, menu.Id, child); err != nil {
			return err
		}
	}
	return nil
}

func (u *Menu) matchRegisterMenu(ctx core.Context, state *registerMenuState, menu *entity.Menu) (*entity.Menu, error) {
	var (
		old *entity.Menu
		err error
	)

	if key := value.Value(menu.Keyword); key != "" {
		if _, ok := state.SeenKey[key]; ok {
			return nil, fmt.Errorf("菜单key重复: %s", key)
		}
		state.SeenKey[key] = struct{}{}
		old, err = u.repo.GetMenuByKeyword(ctx, state.App, key)
	} else if u.isApiMenu(menu.Type) {
		apiKey := u.getRegisterApiKey(menu)
		if _, ok := state.SeenApi[apiKey]; ok {
			return nil, fmt.Errorf("接口菜单重复: %s", apiKey)
		}
		state.SeenApi[apiKey] = struct{}{}
		old, err = u.repo.GetMenuByApi(ctx, state.App, menu.GetApi(), menu.GetMethod())
	} else {
		return nil, nil
	}

	if gormtranserror.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return old, err
}

func (u *Menu) saveRegisterMenu(
	ctx core.Context,
	menu *entity.Menu,
	old *entity.Menu,
	reply *types.RegisterMenuReply,
) error {
	if old == nil {
		id, err := u.repo.CreateMenu(ctx, menu)
		if err != nil {
			return err
		}
		menu.Id = id
		reply.Created++
		return nil
	}

	menu.Id = old.Id
	if err := u.repo.UpdateRegisteredMenu(ctx, menu); err != nil {
		return err
	}
	reply.Updated++
	return nil
}

func (u *Menu) toRegisterMenu(app string, parentId uint32, node *types.RegisterMenuNode) (*entity.Menu, error) {
	var key *string
	if node.Key != nil {
		trimKey := strings.TrimSpace(*node.Key)
		if trimKey != "" {
			key = value.Pointer(trimKey)
		}
	}

	menu := &entity.Menu{
		App:        app,
		ParentId:   value.Pointer(parentId),
		Title:      node.Title,
		Type:       node.Type,
		Keyword:    key,
		Icon:       node.Icon,
		Api:        node.Api,
		Method:     node.Method,
		Path:       node.Path,
		Permission: node.Permission,
		Component:  node.Component,
		Redirect:   node.Redirect,
		Url:        node.Url,
		Weight:     node.Weight,
		IsIframe:   node.IsIframe,
		IsHidden:   node.IsHidden,
		IsCache:    node.IsCache,
		IsHome:     node.IsHome,
		IsAffix:    node.IsAffix,
	}
	if u.isApiMenu(menu.Type) {
		method := strings.ToUpper(strings.TrimSpace(value.Value(menu.Method)))
		api := strings.TrimSpace(value.Value(menu.Api))
		if api == "" || method == "" {
			return nil, fmt.Errorf("接口菜单%s缺少api或method", node.Title)
		}
		menu.Api = value.Pointer(api)
		menu.Method = value.Pointer(method)
	}
	return menu, nil
}

func (u *Menu) deleteUnregisteredMenus(
	ctx core.Context,
	state *registerMenuState,
) error {
	list, err := u.repo.ListMenu(ctx, &types.ListMenuRequest{App: value.Pointer(state.App)})
	if err != nil {
		return err
	}

	var ids []uint32
	for _, item := range list {
		if item.Keyword != nil && strings.TrimSpace(*item.Keyword) != "" {
			if _, ok := state.SeenKey[*item.Keyword]; ok {
				continue
			}
			ids = append(ids, item.Id)
			continue
		}

		if u.isApiMenu(item.Type) {
			if _, ok := state.SeenApi[u.getRegisterApiKey(item)]; ok {
				continue
			}
			ids = append(ids, item.Id)
		}
	}
	if len(ids) == 0 {
		return nil
	}
	if err := u.repo.DeleteMenus(ctx, ids); err != nil {
		return err
	}
	if err := u.rm.DeleteMenus(ctx, state.App, ids); err != nil {
		return err
	}
	state.Reply.Deleted = uint32(len(ids))
	return nil
}

func (u *Menu) getRegisterApiKey(menu *entity.Menu) string {
	return strings.ToUpper(menu.GetMethod()) + ":" + menu.GetApi()
}

func (u *Menu) isApiMenu(tp string) bool {
	return tp == entity.MenuTypeApi || tp == entity.MenuTypeBasic
}
