package service

import (
	"context"
	"sync"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/tree"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/pkg/md"
	"github.com/limes-cloud/manager/internal/types"
)

type Menu struct {
	once sync.Once
	conf *conf.Config
	repo repository.Menu
	role repository.Role
	rbac repository.Rbac
}

func NewMenu(
	config *conf.Config,
	repo repository.Menu,
	role repository.Role,
	rbac repository.Rbac,
) *Menu {
	uc := &Menu{conf: config, repo: repo, role: role, rbac: rbac}
	uc.once.Do(func() {
		uc.repo.InitBasicMenu(kratosx.MustContext(context.Background()))
	})
	return uc
}

// ListMenuByCurRole 获取当前角色的菜单树
func (u *Menu) ListMenuByCurRole(ctx kratosx.Context) ([]*entity.Menu, error) {
	list, err := u.repo.ListMenuByRoleId(ctx, md.RoleId(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "list menu error", "err", err.Error())
		return nil, errors.ListError()
	}
	return tree.BuildArrayTree(list), nil
}

// ListMenu 获取菜单信息列表树
func (u *Menu) ListMenu(ctx kratosx.Context, req *types.ListMenuRequest) ([]*entity.Menu, error) {
	list, err := u.repo.ListMenu(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list menu error", "err", err.Error())
		return nil, errors.ListError()
	}
	return tree.BuildArrayTree(list), nil
}

// CreateMenu 创建菜单信息
func (u *Menu) CreateMenu(ctx kratosx.Context, menu *entity.Menu) (uint32, error) {
	var id uint32
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		// 创建菜单
		var err error
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

		// 添加到白名单
		if menu.Type == entity.MenuBasic {
			ctx.Authentication().AddWhitelist(*menu.Api, *menu.Method)
		}

		return nil
	}); err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateMenu 更新菜单信息
func (u *Menu) UpdateMenu(ctx kratosx.Context, menu *entity.Menu) error {
	old, err := u.repo.GetMenu(ctx, menu.Id)
	if err != nil {
		return err
	}

	err = ctx.Transaction(func(ctx kratosx.Context) error {
		// 之前是api，现在不是
		if old.Type == entity.MenuApi && menu.Type != entity.MenuApi {
			if err := u.rbac.DeleteRbacApi(ctx, *old.Api, *old.Method); err != nil {
				return err
			}
		}

		// 之前是api，现在也是，但是接口/方法出现变化
		if old.Type == entity.MenuApi && menu.Type == entity.MenuApi && (*old.Method != *menu.Method || *old.Api != *menu.Api) {
			if err := u.rbac.UpdateRbacApi(ctx,
				types.MenuApi{Api: *old.Api, Method: *old.Method},
				types.MenuApi{Api: *menu.Api, Method: *menu.Method},
			); err != nil {
				return err
			}
		}

		// 之前不是api，现在是，将api添加到所有已拥有的用户
		if old.Type != entity.MenuApi && menu.Type == entity.MenuApi {
			roles, err := u.role.AllRoleKeywordByMenuId(ctx, menu.Id)
			if err != nil {
				return err
			}
			if err := u.rbac.CreateRbacRolesApi(ctx, roles, types.MenuApi{Api: *old.Api, Method: *old.Method}); err != nil {
				return err
			}
		}

		if old.Type == entity.MenuBasic {
			ctx.Authentication().RemoveWhitelist(*old.Api, *old.Method)
		}

		if menu.Type == entity.MenuBasic {
			ctx.Authentication().AddWhitelist(*menu.Api, *menu.Method)
		}

		return u.repo.UpdateMenu(ctx, menu)
	})

	if err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteMenu 删除菜单信息
func (u *Menu) DeleteMenu(ctx kratosx.Context, id uint32) error {
	apis, err := u.repo.ListMenuChildrenApi(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "delete menu api error", "err", err.Error())
		return errors.DatabaseError()
	}

	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		for _, item := range apis {
			// 移除白名单
			if item.Type == entity.MenuBasic {
				ctx.Authentication().RemoveWhitelist(*item.Api, *item.Method)
			}

			// 移除api
			if item.Type == entity.MenuApi {
				if err := u.rbac.DeleteRbacApi(ctx, *item.Api, *item.Method); err != nil {
					return err
				}
			}
		}
		// 删除
		return u.repo.DeleteMenu(ctx, id)
	}); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
