package logic

import (
	"fmt"
	v1 "manager/api/v1"
	"manager/config"
	"manager/consts"
	"manager/internal/model"
	"manager/pkg/md"
	"manager/pkg/tree"
	"manager/pkg/util"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	"gorm.io/gorm"
)

type RoleMenu struct {
	conf *config.Config
}

func NewRoleMenu(conf *config.Config) *RoleMenu {
	return &RoleMenu{
		conf: conf,
	}
}

// CurrenMenuTree 查询当前角色的菜单树
func (r *RoleMenu) CurrenMenuTree(ctx kratosx.Context) (*v1.CurrentRoleMenuTreeReply, error) {
	menu := model.Menu{}
	list, err := menu.All(ctx, func(db *gorm.DB) *gorm.DB {
		db = db.Where("type!=? ", consts.MENU_BASIC)
		if md.RoleId(ctx) == consts.SuperAdmin {
			return db
		}
		return db.Where("id in (select menu_id from role_menu where role_id = ?)", md.RoleId(ctx))
	})
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

	reply := v1.CurrentRoleMenuTreeReply{}
	if err = util.Transform(roots, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	return &reply, nil
}

// GetMenuIds 获取指定角色的菜单ID
func (r *RoleMenu) GetMenuIds(ctx kratosx.Context, in *v1.GetRoleMenuIdsRequest) (*v1.GetRoleMenuIdsReply, error) {
	// 判断是否具有此角色的权限
	user := model.User{}
	roleIds, _ := user.RoleScope(ctx, md.UserId(ctx))
	if !util.InList(roleIds, in.RoleId) {
		return nil, v1.DepartmentPermissionsError()
	}

	rm := model.RoleMenu{}
	rms, err := rm.RoleMenus(ctx, in.RoleId)
	if err != nil {
		return nil, err
	}

	// 组装所有的菜单id
	var ids = make([]uint32, 0)
	for _, item := range rms {
		ids = append(ids, item.MenuID)
	}

	return &v1.GetRoleMenuIdsReply{
		List: ids,
	}, nil
}

func (r *RoleMenu) Update(ctx kratosx.Context, in *v1.UpdateRoleMenuRequest) (*emptypb.Empty, error) {
	// 超级管理员不存在菜单权限，自动获取全部菜单，所以禁止修改
	if in.RoleId == consts.SuperAdmin {
		return nil, v1.EditSystemDataError()
	}

	// 判断是否具有此角色的权限
	user := model.User{}
	roleIds, _ := user.RoleScope(ctx, md.UserId(ctx))
	if !util.InList(roleIds, in.RoleId) {
		return nil, v1.DepartmentPermissionsError()
	}

	// 获取当前用户的菜单id
	reply, err := r.GetMenuIds(ctx, &v1.GetRoleMenuIdsRequest{RoleId: md.RoleId(ctx)})
	if err != nil {
		return nil, err
	}

	//不晒超级管理员添加，则需要判断菜单是否都合法
	if md.RoleId(ctx) != consts.SuperAdmin {
		for _, id := range in.MenuIds {
			if !util.InList(reply.List, id) {
				return nil, v1.MenuPermissionsErrorFormat(fmt.Sprintf("menu_id=%d", id))
			}
		}
	}

	// 获取当前role的数据
	role := model.Role{}
	if err := role.OneByID(ctx, in.RoleId); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 获取新的菜单信息
	menu := model.Menu{}
	var policies [][]string
	apiList, _ := menu.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in ? and type = 'A'", in.MenuIds)
	})
	for _, item := range apiList {
		policies = append(policies, []string{role.Keyword, *item.Api, *item.Method})
	}

	// 进行菜单变更
	rm := model.RoleMenu{}
	if err = rm.Update(ctx, in.RoleId, in.MenuIds); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 删除当前用户的全部rbac权限
	enforce := ctx.Authentication().Enforce()
	_, _ = enforce.RemoveFilteredPolicy(0, role.Keyword)

	// 添加新的rbac权限信息
	_, _ = enforce.AddPolicies(policies)

	return nil, nil
}
