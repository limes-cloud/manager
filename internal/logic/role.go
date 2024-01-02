package logic

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	v1 "github.com/limes-cloud/manager/api/v1"
	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/internal/model"
	"github.com/limes-cloud/manager/pkg/md"
	"github.com/limes-cloud/manager/pkg/tree"
	"github.com/limes-cloud/manager/pkg/util"
)

type Role struct {
	conf *config.Config
}

func NewRole(conf *config.Config) *Role {
	return &Role{
		conf: conf,
	}
}

// Get 查询指定的角色信息
func (r *Role) Get(ctx kratosx.Context, in *v1.GetRoleRequest) (*v1.GetRoleReply, error) {
	role := model.Role{}
	if err := role.FindByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.GetRoleReply{}
	if err := util.Transform(role, &reply.Role); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// Tree 查询当前用户的角色树
func (r *Role) Tree(ctx kratosx.Context) (*v1.GetRoleTreeReply, error) {
	role := model.Role{
		BaseModel: types.BaseModel{
			ID: md.RoleId(ctx),
		},
	}
	if role.ID() == 0 {
		return nil, v1.MetadataErrorFormat("role_id不存在")
	}

	// 获取当前角色的子角色id
	ids, err := role.FindManagerIds(ctx)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 查询子角色id列表信息
	list, err := role.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in ?", ids)
	})
	if err != nil {
		return nil, err
	}

	// 构造角色树
	var tl []tree.Tree
	for _, item := range list {
		tl = append(tl, item)
	}
	trs := tree.BuildTree(tl)

	reply := v1.GetRoleTreeReply{}
	if err = util.Transform(trs, &reply.Role); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	return &reply, nil
}

func (r *Role) Add(ctx kratosx.Context, in *v1.AddRoleRequest) (*emptypb.Empty, error) {
	// 获取当前角色的子角色
	role := model.Role{
		BaseModel: types.BaseModel{
			ID: md.RoleId(ctx),
		},
	}
	ids, err := role.FindManagerIds(ctx)
	if err != nil {
		return nil, v1.DatabaseError()
	}
	if !util.InList(ids, in.ParentId) {
		return nil, v1.RolePermissionsErrorFormat(fmt.Sprintf("role_id:%d", in.ParentId))
	}

	// 创建角色信息
	nr := model.Role{}
	if err := util.Transform(in, &nr); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := nr.Create(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

func (r *Role) Update(ctx kratosx.Context, in *v1.UpdateRoleRequest) (*emptypb.Empty, error) {
	// 超级管理员不允许修改
	if in.Id == consts.SuperAdmin {
		return nil, v1.EditSystemDataError()
	}

	//  不能修改当前角色
	if in.Id == md.RoleId(ctx) {
		return nil, v1.RolePermissionsError()
	}

	// 获取当前角色的子角色
	role := model.Role{
		BaseModel: types.BaseModel{
			ID: md.RoleId(ctx),
		},
	}
	ids, err := role.FindManagerIds(ctx)
	if err != nil {
		return nil, v1.DatabaseError()
	}

	// 判断当前部门是否具有权限
	if !util.InList(ids, in.Id) {
		return nil, v1.DepartmentPermissionsError()
	}

	// 查询历史角色信息
	oldRole := model.Role{}
	if err := oldRole.FindByID(ctx, in.Id); err != nil {
		return nil, v1.NotFoundError()
	}

	// 如果存在移动角色，判断是否有父级部门的权限
	if oldRole.ParentID != in.ParentId && !util.InList(ids, in.ParentId) {
		return nil, v1.DepartmentPermissionsError()
	}

	// 更新角色信息
	nr := model.Role{}
	if err := util.Transform(in, &nr); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	if err := nr.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

func (r *Role) Delete(ctx kratosx.Context, in *v1.DeleteRoleRequest) (*emptypb.Empty, error) {
	// 超级管理员不允许删除
	if in.Id == consts.SuperAdmin {
		return nil, v1.DeleteSystemDataError()
	}

	// 不能删除当前角色
	if in.Id == md.RoleId(ctx) {
		return nil, v1.DeleteSelfRoleError()
	}

	// 判断是否具有角色的管理信息
	role := model.Role{
		BaseModel: types.BaseModel{
			ID: md.RoleId(ctx),
		},
	}

	// 判断是否未当前角色的下级角色
	ids, err := role.FindManagerIds(ctx)
	if err != nil {
		return nil, v1.DatabaseError()
	}
	if !util.InList(ids, in.Id) {
		return nil, v1.RolePermissionsError()
	}

	// 删除角色
	if err := role.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 删除rbac
	_, _ = ctx.Authentication().Enforce().RemoveFilteredPolicy(0, role.Keyword)
	return nil, nil
}
