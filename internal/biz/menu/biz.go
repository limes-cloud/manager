package menu

import (
	"context"
	"sync"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/pkg/md"
)

type UseCase struct {
	once sync.Once
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	uc := &UseCase{conf: config, repo: repo}
	uc.once.Do(func() {
		uc.repo.InitBasicMenu(kratosx.MustContext(context.Background()))
	})
	return uc
}

// ListMenuByCurRole 获取当前角色的菜单树
func (u *UseCase) ListMenuByCurRole(ctx kratosx.Context) ([]tree.Tree, uint32, error) {
	list, total, err := u.repo.ListMenuByRoleId(ctx, md.RoleId(ctx))
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	var ts []tree.Tree
	for _, item := range list {
		ts = append(ts, item)
	}
	return tree.BuildArrayTree(ts), total, nil
}

// ListMenu 获取菜单信息列表树
func (u *UseCase) ListMenu(ctx kratosx.Context, req *ListMenuRequest) ([]tree.Tree, uint32, error) {
	list, total, err := u.repo.ListMenu(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	var ts []tree.Tree
	for _, item := range list {
		ts = append(ts, item)
	}
	return tree.BuildArrayTree(ts), total, nil
}

// CreateMenu 创建菜单信息
func (u *UseCase) CreateMenu(ctx kratosx.Context, req *Menu) (uint32, error) {
	id, err := u.repo.CreateMenu(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateMenu 更新菜单信息
func (u *UseCase) UpdateMenu(ctx kratosx.Context, req *Menu) error {
	if err := u.repo.UpdateMenu(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteMenu 删除菜单信息
func (u *UseCase) DeleteMenu(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteMenu(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}
