package logic

import (
	"github.com/limes-cloud/kratosx"

	v1 "github.com/limes-cloud/manager/api/v1"
	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/internal/model"
	"github.com/limes-cloud/manager/pkg/md"
	"github.com/limes-cloud/manager/pkg/util"
)

type UserRole struct {
	conf *config.Config
}

func NewUserRole(conf *config.Config) *UserRole {
	return &UserRole{
		conf: conf,
	}
}

// GetUserRoles 查询指定用户的所有角色
func (r *UserRole) GetUserRoles(ctx kratosx.Context, in *v1.GetUserRolesRequest) (*v1.GetUserRolesReply, error) {
	ur := model.UserRole{}
	urs, err := ur.UserRoles(ctx, in.UserId)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 进行数据转换
	reply := v1.GetUserRolesReply{}
	if err = util.Transform(urs, &reply.UserRole); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// CurrentUserRoles 查询当前用户的所有角色
func (r *UserRole) CurrentUserRoles(ctx kratosx.Context) (*v1.GetUserRolesReply, error) {
	return r.GetUserRoles(ctx, &v1.GetUserRolesRequest{
		UserId: md.UserId(ctx),
	})
}

// SwitchCurrentUserRole 切换当前用户的角色
func (r *UserRole) SwitchCurrentUserRole(ctx kratosx.Context, in *v1.SwitchCurrentUserRoleRequest) (*v1.SwitchCurrentUserRoleReply, error) {
	// 判断是具有被切换的角色
	ur := model.UserRole{}
	if err := ur.FindByUserAndRole(ctx, md.UserId(ctx), in.RoleId); err != nil {
		return nil, v1.NotFoundError()
	}

	// 获取当前角色
	role := model.Role{}
	if err := role.FindByID(ctx, in.RoleId); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 获取当前用户信息
	user := model.User{}
	if err := user.FindByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 设置生成的用户信息为切换后角色信息
	user.RoleID = role.ID()
	user.Role.Keyword = role.Keyword
	user.Role.BaseModel.ID = role.ID()

	// 重新生成token
	token, err := ctx.JWT().NewToken(md.New(&user))
	if err != nil {
		return nil, v1.NewTokenErrorFormat(err.Error())
	}

	// 更新用户的当前token,以及角色
	nu := model.User{}
	nu.ID = user.ID
	nu.RoleID = in.RoleId
	nu.Token = token
	if err = nu.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return &v1.SwitchCurrentUserRoleReply{Token: token}, nil
}
