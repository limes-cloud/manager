package menu

import (
	"context"
	"sync"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/tree"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/config"
	"github.com/limes-cloud/manager/internal/pkg/md"
)

type UseCase struct {
	repo Repo
	conf *config.Config
	once sync.Once
}

func NewUseCase(conf *config.Config, repo Repo) *UseCase {
	u := &UseCase{
		repo: repo,
		conf: conf,
	}
	u.once.Do(func() {
		repo.InitBasicMenu(kratosx.MustContext(context.Background()))
	})
	return u
}

// MenuTree 查询系统菜单树
func (u *UseCase) MenuTree(ctx kratosx.Context, in *AllMenuRequest) ([]tree.Tree, error) {
	list, err := u.repo.AllMenu(ctx, in)
	if err != nil {
		return nil, errors.Database()
	}

	var trees []tree.Tree

	// 构建树枝，并选取根节点
	for _, item := range list {
		trees = append(trees, item)
	}

	return tree.BuildArrayTree(trees), nil
}

// // MenuTree 查询系统菜单树
//
//	func (u *UseCase) MenuTree(ctx kratosx.Context, in *AllMenuRequest) ([]tree.Tree, error) {
//		list, err := u.repo.AllMenu(ctx, in)
//		if err != nil {
//			return nil, errors.Database()
//		}
//
//		var (
//			group = make(map[string][]tree.Tree)
//			roots []tree.Tree
//			apps  []*Menu
//		)
//
//		// 构建树枝，并选取根节点
//		for _, item := range list {
//			group[item.App] = append(group[item.App], item)
//			if item.ParentId == 0 {
//				apps = append(apps, item)
//			}
//		}
//
//		// 通过根节点构造树
//		for _, app := range apps {
//			root := tree.BuildTreeByID(group[app.App], app.ID())
//			roots = append(roots, root)
//		}
//
//		return roots, nil
//	}
func (u *UseCase) MenuTreeFromRole(ctx kratosx.Context) ([]tree.Tree, error) {
	req := &AllMenuRequest{NotType: proto.String(MenuBasic)}
	if md.RoleId(ctx) != 1 {
		req.RoleId = proto.Uint32(md.RoleId(ctx))
	}
	return u.MenuTree(ctx, req)
}

func (u *UseCase) AddMenu(ctx kratosx.Context, in *Menu) (uint32, error) {
	if in.Type == MenuRoot {
		in.ParentId = 0
	}

	id, err := u.repo.AddMenu(ctx, in)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}

	// 更新菜单首页
	if in.IsHome != nil && *in.IsHome {
		_ = u.repo.UseHome(ctx, *in.Keyword, id)
	}

	// 如果是基础api，则添加到白名单中
	if in.Type == MenuBasic {
		ctx.Authentication().AddWhitelist(*in.Api, *in.Method)
	}

	return id, nil
}

func (u *UseCase) UpdateMenu(ctx kratosx.Context, in *Menu) error {
	old, err := u.repo.GetMenu(ctx, in.ID())
	if err != nil {
		return errors.NotFound()
	}

	// 不能修改菜单副菜单为的当前id
	if old.ID() == in.Parent() {
		return errors.ParentMenu()
	}

	enforce := ctx.Authentication().Enforce()
	// 之前是接口，变更后不是
	if old.Type == MenuApi && in.Type != MenuApi {
		_, _ = enforce.RemoveFilteredPolicy(1, *old.Api, *old.Method)
	}

	// 存在方法或者路径变更时则更新rbac数据
	if old.Type == MenuApi && in.Type == MenuApi && (*old.Method != *in.Method || *old.Api != *in.Api) {
		oldPolices := enforce.GetFilteredPolicy(1, *old.Api, *old.Method)
		if len(oldPolices) != 0 {
			var newPolices [][]string
			for _, val := range oldPolices {
				newPolices = append(newPolices, []string{val[0], *in.Api, *in.Method})
			}
			_, _ = enforce.UpdatePolicies(oldPolices, newPolices)
		}
	}

	// 之前不是接口，现在是接口的情况下，则对拥有此菜单的角色全部新增
	if old.Type != MenuApi && in.Type == MenuApi {
		keys, _ := u.repo.AllRoleKeyword(ctx, in.ID())
		if len(keys) != 0 {
			var newPolices [][]string
			for _, val := range keys {
				newPolices = append(newPolices, []string{val, *in.Api, *in.Method})
			}
			_, _ = enforce.AddPolicies(newPolices)
		}
	}

	// 如果之前是基础api，则删除白名单
	if old.Type == MenuBasic {
		ctx.Authentication().RemoveWhitelist(*old.Api, *old.Method)
	}

	// 如果现在是基础api，则添加白名单
	if in.Type == MenuBasic {
		ctx.Authentication().AddWhitelist(*in.Api, *in.Method)
	}

	// 提交修改
	if err := u.repo.UpdateMenu(ctx, in); err != nil {
		return errors.DatabaseFormat(err.Error())
	}

	// 更新为菜单首页
	if in.IsHome != nil && *in.IsHome {
		_ = u.repo.UseHome(ctx, *old.Keyword, in.ID())
	}

	return nil
}

func (u *UseCase) DeleteMenu(ctx kratosx.Context, id uint32) error {
	if id == 1 {
		return errors.DeleteSystemData()
	}

	list, err := u.repo.AllMenu(ctx, &AllMenuRequest{})
	if err != nil {
		return errors.DatabaseFormat(err.Error())
	}

	// 构造菜单树
	var (
		tl     []tree.Tree
		apis   []*Menu
		apiSet = map[uint32]*Menu{}
	)
	for _, item := range list {
		tl = append(tl, item)
		if item.Type == MenuApi || item.Type == MenuBasic {
			apiSet[item.ID()] = item
		}
	}

	// 获取删除菜单的下级菜单id
	ids := tree.GetTreeID(tree.BuildTreeByID(tl, id))

	// 筛选出其中是Api的路由
	for _, id := range ids {
		if apiSet[id] != nil {
			apis = append(apis, apiSet[id])
		}
	}

	// 删除菜单
	if err := u.repo.DeleteMenu(ctx, ids); err != nil {
		return errors.DatabaseFormat(err.Error())
	}

	// 删除接口的rbac权限表
	auth := ctx.Authentication()
	for _, item := range apis {
		if item.Type == MenuBasic {
			auth.RemoveWhitelist(*item.Api, *item.Method)
		}
		if item.Type == MenuApi {
			_, _ = auth.Enforce().RemoveFilteredPolicy(1, *item.Api, *item.Method)
		}
	}
	return nil
}
