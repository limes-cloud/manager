package logic

import (
	"fmt"
	v1 "manager/api/v1"
	"manager/config"
	"manager/consts"
	"manager/internal/model"
	"manager/pkg/md"
	"manager/pkg/util"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/kratos"
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
func (r *Role) Get(ctx kratos.Context, in *v1.GetRoleRequest) (*v1.GetRoleReply, error) {
	role := model.Role{}
	if err := role.OneByID(ctx, in.Id); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	reply := v1.GetRoleReply{}
	if err := util.Transform(role, &reply.Role); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	return &reply, nil
}

// Tree 查询当前用户的角色树
func (r *Role) Tree(ctx kratos.Context) (*v1.GetRoleTreeReply, error) {
	role := model.Role{}

	rid := md.RoleId(ctx)
	if rid == 0 {
		return nil, v1.ErrorMetadataFormat("role_id不存在")
	}

	tree, err := role.TreeByID(ctx, rid)
	if err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	reply := v1.GetRoleTreeReply{}
	if err = util.Transform(tree, &reply.Role); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}
	return &reply, nil
}

func (r *Role) Add(ctx kratos.Context, in *v1.AddRoleRequest) (*emptypb.Empty, error) {
	// 判断是否具有角色的管理信息
	user := model.User{}
	roleIds, err := user.RoleScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	if !util.InList(roleIds, in.ParentId) {
		return nil, v1.ErrorRolePermissionsFormat(fmt.Sprintf("role_id:%d", in.ParentId))
	}

	// 创建角色信息
	role := model.Role{}
	if err := util.Transform(in, &role); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	if err := role.Create(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return nil, nil
}

func (r *Role) Update(ctx kratos.Context, in *v1.UpdateRoleRequest) (*emptypb.Empty, error) {
	// 超级管理员不允许修改
	if in.Id == consts.SuperAdmin {
		return nil, v1.ErrorEditSystemData()
	}

	// 查询角色信息
	oldRole := model.Role{}
	if err := oldRole.OneByID(ctx, in.Id); err != nil {
		return nil, v1.ErrorNotFound()
	}

	// 判断是否具有角色的管理信息
	user := model.User{}
	roleIds, err := user.RoleScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	// 如果存在移动角色，判断是否有父级部门的权限
	if oldRole.ParentID != in.ParentId && !util.InList(roleIds, in.ParentId) {
		return nil, v1.ErrorDepartmentPermissions()
	}
	// 判断当前部门是否具有权限
	if !util.InList(roleIds, in.Id) {
		return nil, v1.ErrorDepartmentPermissions()
	}

	// 更新角色信息
	role := model.Role{}
	if err := util.Transform(in, &role); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	if err := role.Update(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return nil, nil
}

func (r *Role) Delete(ctx kratos.Context, in *v1.DeleteRoleRequest) (*emptypb.Empty, error) {
	if in.Id == consts.SuperAdmin {
		return nil, v1.ErrorDeleteSystemData()
	}

	if in.Id == md.RoleId(ctx) {
		return nil, v1.ErrorDeleteSelfRole()
	}

	// 判断是否具有角色的管理信息
	user := model.User{}
	roleIds, err := user.RoleScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	if !util.InList(roleIds, in.Id) {
		return nil, v1.ErrorRolePermissions()
	}

	// 删除角色
	role := model.Role{}
	if err := role.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return nil, nil
}
