package logic

import (
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	v1 "github.com/limes-cloud/manager/api/v1"
	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/internal/model"
	"github.com/limes-cloud/manager/pkg/tree"
	"github.com/limes-cloud/manager/pkg/util"
)

type Menu struct {
	conf *config.Config
}

func NewMenu(conf *config.Config) *Menu {
	return &Menu{
		conf: conf,
	}
}

// Tree 查询系统菜单树
func (r *Menu) Tree(ctx kratosx.Context) (*v1.GetMenuTreeReply, error) {
	menu := model.Menu{}
	list, err := menu.All(ctx, nil)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	var (
		group = make(map[string][]tree.Tree)
		roots []tree.Tree
		apps  []*model.Menu
	)

	// 构建树枝，并选取根节点
	for _, item := range list {
		group[item.App] = append(group[item.App], item)
		if item.ParentID == 0 {
			apps = append(apps, item)
		}
	}

	// 通过根及诶单构造树
	for _, app := range apps {
		root := tree.BuildTreeByID(group[app.App], app.ID())
		roots = append(roots, root)
	}

	reply := v1.GetMenuTreeReply{}
	if err = util.Transform(roots, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	return &reply, nil
}

func (r *Menu) Add(ctx kratosx.Context, in *v1.AddMenuRequest) (*emptypb.Empty, error) {
	menu := model.Menu{}
	if err := util.Transform(in, &menu); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := menu.Create(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 更新菜单首页
	if in.IsHome != nil && *in.IsHome {
		_ = menu.UseHome(ctx, in.App, menu.ID())
	}

	// 如果是基础api，则添加到白名单中
	if in.Type == consts.MenuBasic {
		ctx.Authentication().AddWhitelist(*in.Api, *in.Method)
	}

	return nil, nil
}

func (r *Menu) Update(ctx kratosx.Context, in *v1.UpdateMenuRequest) (*emptypb.Empty, error) {
	oldMenu := model.Menu{}
	if err := oldMenu.FindByID(ctx, in.Id); err != nil {
		return nil, v1.NotFoundErrorFormat(err.Error())
	}

	if in.Id == in.ParentId {
		return nil, v1.ParentMenuError()
	}

	enforce := ctx.Authentication().Enforce()

	// 菜单之前是接口，现在不是，则删除rbac
	if oldMenu.Type == "A" && in.Type != "A" {
		_, _ = enforce.RemoveFilteredPolicy(1, *oldMenu.Api, *oldMenu.Method)
	}

	// 之前和现在都为接口，且存在方法或者路径变更时则更新rbac数据
	if oldMenu.Type == "A" && in.Type == "A" && (*oldMenu.Method != *in.Method || *oldMenu.Api != *in.Api) {
		oldPolices := enforce.GetFilteredPolicy(1, *oldMenu.Api, *oldMenu.Method)
		if len(oldPolices) != 0 {
			var newPolices [][]string
			for _, val := range oldPolices {
				newPolices = append(newPolices, []string{val[0], *in.Api, *in.Method})
			}
			_, _ = enforce.UpdatePolicies(oldPolices, newPolices)
		}
	}

	// 当之前不是接口，现在是接口的情况下，则进行新增
	if oldMenu.Type != "A" && in.Type == "A" {
		// 获取选中当前菜单的角色
		rm := model.RoleMenu{}
		rms, _ := rm.MenuRoles(ctx, in.Id)
		if len(rms) != 0 {
			var roleIds []uint32
			for _, item := range rms {
				roleIds = append(roleIds, item.RoleID)
			}

			// 获取当前菜单的全部角色信息
			role := model.Role{}
			roles, _ := role.All(ctx, func(db *gorm.DB) *gorm.DB {
				return db.Where("id in ?", roleIds)
			})

			// 添加菜单到rbac权限表
			var newPolices [][]string
			for _, val := range roles {
				newPolices = append(newPolices, []string{val.Keyword, *in.Api, *in.Method})
			}
			_, _ = enforce.AddPolicies(newPolices)
		}
	}

	// 如果之前是基础api，则删除白名单
	if oldMenu.Type == consts.MenuBasic {
		ctx.Authentication().RemoveWhitelist(*oldMenu.Api, *oldMenu.Method)
	}

	// 如果现在是基础api，则添加白名单
	if in.Type == consts.MenuBasic {
		ctx.Authentication().AddWhitelist(*in.Api, *in.Method)
	}

	// 修改数据库
	menu := model.Menu{}
	if err := util.Transform(in, &menu); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := menu.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 更新为菜单首页
	if in.IsHome != nil && *in.IsHome {
		_ = menu.UseHome(ctx, oldMenu.App, menu.ID())
	}

	return nil, nil
}

func (r *Menu) Delete(ctx kratosx.Context, in *v1.DeleteMenuRequest) (*emptypb.Empty, error) {
	if in.Id == consts.SuperAdmin {
		return nil, v1.DeleteSystemDataError()
	}
	menu := model.Menu{}
	// 查询当前菜单
	if err := menu.FindByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 获取当前服务的所有菜单
	list, _ := menu.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("app=?", menu.App)
	})

	// 构造菜单树
	var (
		tl     []tree.Tree
		apis   []*model.Menu
		apiSet = map[uint32]*model.Menu{}
	)
	for _, item := range list {
		tl = append(tl, item)
		if item.Type == "A" {
			apiSet[item.ID()] = item
		}
	}
	// 获取删除菜单的下级菜单id
	t := tree.BuildTreeByID(tl, in.Id)
	ids := tree.GetTreeID(t)

	// 筛选出其中是Api的路由
	for _, id := range ids {
		if apiSet[id] != nil {
			apis = append(apis, apiSet[id])
		}
	}

	// 删除数据库
	if err := menu.DeleteByIds(ctx, ids); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 删除api的rbac权限表
	auth := ctx.Authentication()
	for _, item := range apis {
		if item.Type == consts.MenuBasic {
			auth.RemoveWhitelist(*item.Api, *item.Method)
		}
		_, _ = auth.Enforce().RemoveFilteredPolicy(1, *item.Api, *item.Method)
	}

	return nil, nil
}
