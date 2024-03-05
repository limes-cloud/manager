package logic

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	v1 "github.com/limes-cloud/manager/api/v1"
	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/internal/model"
	"github.com/limes-cloud/manager/pkg/md"
	"github.com/limes-cloud/manager/pkg/util"
)

type User struct {
	conf *config.Config
}

func NewUser(conf *config.Config) *User {
	return &User{
		conf: conf,
	}
}

const (
	captcha = "changePassword"
)

// Get 查询指定的用户信息，提供给外部rpc调用
func (r *User) Get(ctx kratosx.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user := model.User{}
	if err := user.FindByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 获取用户的全部角色
	var roles []*model.Role
	rm := model.UserRole{}
	rms, _ := rm.UserRoles(ctx, in.Id)
	for _, item := range rms {
		roles = append(roles, item.Role)
	}

	reply := v1.GetUserReply{}
	if err := util.Transform(user, &reply.User); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := util.Transform(roles, &reply.User.Roles); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// ChangePasswordCaptcha 重置用户密码验证码
func (r *User) ChangePasswordCaptcha(ctx kratosx.Context, in *emptypb.Empty) (*v1.ChangeUserPasswordCaptchaReply, error) {
	user := model.User{}
	if err := user.FindByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.NotFoundError()
	}

	resp, err := ctx.Captcha().Email(captcha, ctx.ClientIP(), user.Email)
	if err != nil {
		return nil, v1.SendEmailCaptchaErrorFormat(err.Error())
	}

	return &v1.ChangeUserPasswordCaptchaReply{
		Uuid:   resp.ID(),
		Expire: uint32(resp.Expire().Seconds()),
	}, nil
}

// ChangePassword 重置用户密码
func (r *User) ChangePassword(ctx kratosx.Context, in *v1.ChangeUserPasswordRequest) (*emptypb.Empty, error) {
	// 查询用户
	user := model.User{}
	if err := user.FindByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.NotFoundError()
	}

	// 验证验证码
	if err := ctx.Captcha().VerifyEmail(captcha, ctx.ClientIP(), in.CaptchaId, in.Captcha); err != nil {
		return nil, v1.VerifyCaptchaError()
	}

	// 修改密码
	nu := model.User{}
	nu.ID = user.ID
	nu.Password = in.Password
	if err := nu.Update(ctx); err != nil {
		return nil, v1.DatabaseError()
	}

	return nil, nil
}

// Enable 启用用户
func (r *User) Enable(ctx kratosx.Context, in *v1.EnableUserRequest) (*emptypb.Empty, error) {
	// 查询用户
	user := model.User{}
	if err := user.FindByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.NotFoundError()
	}

	// 判断当前用户是否具有部门的权限
	depIds, _ := user.ManagerDepartmentIds(ctx, md.UserId(ctx))
	if !util.InList(depIds, user.DepartmentID) {
		return nil, v1.DepartmentPermissionsError()
	}

	// 判断是否具有当前用户的管理权限
	if !user.HasUserManagerScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.UserPermissionsError()
	}

	// 将用户启用
	nu := model.User{}
	nu.ID = in.Id
	nu.Status = proto.Bool(true)
	nu.Disabled = proto.String("")
	if err := nu.Update(ctx); err != nil {
		return nil, v1.DatabaseError()
	}

	return nil, nil
}

// Disable 禁用用户
func (r *User) Disable(ctx kratosx.Context, in *v1.DisableUserRequest) (*emptypb.Empty, error) {
	// 超级管理员不允许修改
	if in.Id == consts.SuperAdmin {
		return nil, v1.EditSystemDataError()
	}

	// 超级管理员不允许修改
	if in.Id == md.UserId(ctx) {
		return nil, v1.DisableSelfUserError()
	}

	// 查询用户
	user := model.User{}
	if err := user.FindByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.NotFoundError()
	}

	// 判断是否具有当前用户的管理权限
	if !user.HasUserManagerScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.UserPermissionsError()
	}

	// 将用户禁用
	nu := model.User{}
	nu.ID = in.Id
	nu.Status = proto.Bool(false)
	nu.Disabled = proto.String(in.Desc)
	if err := nu.Update(ctx); err != nil {
		return nil, v1.DatabaseError()
	}

	return nil, nil
}

// Offline 对当前登陆用户进行下线
func (r *User) Offline(ctx kratosx.Context, in *v1.OfflineUserRequest) (*emptypb.Empty, error) {
	// 查询用户
	user := model.User{}
	if err := user.FindByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.NotFoundError()
	}

	// 判断是否具有当前用户的管理权限
	if !user.HasUserManagerScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.UserPermissionsError()
	}

	// 将用户下线
	ctx.JWT().AddBlacklist(user.Token)
	return nil, nil
}

// ResetPassword 重置用户密码
func (r *User) ResetPassword(ctx kratosx.Context, in *v1.ResetUserPasswordRequest) (*emptypb.Empty, error) {
	// 超级管理员不允许修改
	if in.Id == consts.SuperAdmin {
		return nil, v1.EditSystemDataError()
	}

	if in.Id == md.UserId(ctx) {
		return nil, v1.ResetSelfUserPasswordError()
	}

	// 查询用户信息
	user := model.User{}
	if err := user.FindByID(ctx, in.Id); err != nil {
		return nil, v1.NotFoundError()
	}

	// 判断是否具有当前用户的管理权限
	if !user.HasUserManagerScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.UserPermissionsError()
	}

	// 重置密码
	nu := model.User{
		BaseModel: types.BaseModel{
			ID: in.Id,
		},
		Password: r.conf.DefaultUserPassword,
	}
	if err := nu.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Current 查询当前用户信息
func (r *User) Current(ctx kratosx.Context) (*v1.GetUserReply, error) {
	return r.Get(ctx, &v1.GetUserRequest{
		Id: md.UserId(ctx),
	})
}

func (r *User) Page(ctx kratosx.Context, in *v1.PageUserRequest) (*v1.PageUserReply, error) {
	user := model.User{}
	// 获取用户所管理的部门
	ids, err := user.ManagerDepartmentIds(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	list, total, err := user.Page(ctx, &types.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.Username != nil {
				db = db.Where("email=? or phone=?", *in.Username, *in.Username)
			}
			if in.Status != nil {
				db = db.Where("status=?", *in.Status)
			}
			if in.RoleId != nil {
				db = db.Where("role_id=?", *in.RoleId)
			}
			if in.DepartmentId != nil {
				db = db.Where("department_id=?", *in.DepartmentId)
			}
			if in.Name != nil {
				db = db.Where("name like ?", "%"+*in.Name+"%")
			}
			if in.StartTime != nil {
				db = db.Where("created_at >= ?", *in.StartTime)
			}
			if in.EndTime != nil {
				db = db.Where("created_at <= ?", *in.EndTime)
			}
			return db.Where("department_id in ?", ids)
		},
	})
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.PageUserReply{Total: total}
	// 进行数据转换
	if err = util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// Add 添加用户信息
func (r *User) Add(ctx kratosx.Context, in *v1.AddUserRequest) (*emptypb.Empty, error) {
	user := model.User{}

	// 判断当前用户是否具有部门的权限
	depIds, _ := user.ManagerDepartmentIds(ctx, md.UserId(ctx))
	if !util.InList(depIds, in.DepartmentId) {
		return nil, v1.DepartmentPermissionsError()
	}

	// 判断当前用户是否具有角色的权限
	role := model.Role{BaseModel: types.BaseModel{ID: md.RoleId(ctx)}}
	rids, _ := role.FindManagerIds(ctx)
	for _, id := range in.RoleIds {
		if !util.InList(rids, id) {
			return nil, v1.DepartmentPermissionsErrorFormat(fmt.Sprintf("role_id:%d", id))
		}
	}

	// 进行数据转换
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 设置默认值
	user.Nickname = user.Name
	user.Avatar = r.conf.DefaultUserAvatar
	user.Password = r.conf.DefaultUserPassword
	user.RoleID = in.RoleIds[0]
	user.Status = proto.Bool(true)

	// 创建用户
	if err := user.Create(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 更新用户角色信息
	userRole := model.UserRole{}
	if err := userRole.Update(ctx, user.ID, in.RoleIds); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 更新用户职位信息
	userJob := model.UserJob{}
	if err := userJob.Update(ctx, user.ID, in.JobIds); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Update 更新用户信息
func (r *User) Update(ctx kratosx.Context, in *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	// 超级管理员不允许修改
	if in.Id == consts.SuperAdmin {
		return nil, v1.EditSystemDataError()
	}

	user := model.User{}
	// 判断当前用户是否具有部门的权限
	depIds, _ := user.ManagerDepartmentIds(ctx, md.UserId(ctx))
	if !util.InList(depIds, in.DepartmentId) {
		return nil, v1.DepartmentPermissionsError()
	}

	// 判断当前用户是否具有角色的权限
	role := model.Role{BaseModel: types.BaseModel{ID: md.RoleId(ctx)}}
	rids, _ := role.FindManagerIds(ctx)
	for _, id := range in.RoleIds {
		if !util.InList(rids, id) {
			return nil, v1.DepartmentPermissionsErrorFormat(fmt.Sprintf("role_id:%d", id))
		}
	}

	// 判断是否具有当前用户的管理权限
	if !user.HasUserManagerScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.UserPermissionsError()
	}

	// 转换数据格式
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 将角色设置为更新后的第一个角色
	user.RoleID = in.RoleIds[0]

	// 更新用户信息
	if err := user.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 更新用户角色信息
	userRole := model.UserRole{}
	if err := userRole.Update(ctx, user.ID, in.RoleIds); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 更新用户职位信息
	userJob := model.UserJob{}
	if err := userJob.Update(ctx, user.ID, in.JobIds); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// UpdateBasic 更新用户基础信息
func (r *User) UpdateBasic(ctx kratosx.Context, in *v1.UpdateUserBasicRequest) (*emptypb.Empty, error) {
	user := model.User{
		BaseModel: types.BaseModel{ID: md.UserId(ctx)},
		Nickname:  in.Nickname,
		Gender:    in.Gender,
	}

	// 更新用户信息
	if err := user.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Delete 删除用户
func (r *User) Delete(ctx kratosx.Context, in *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	if in.Id == consts.SuperAdmin {
		return nil, v1.DeleteSystemDataError()
	}

	if in.Id == md.UserId(ctx) {
		return nil, v1.DeleteSelfUserError()
	}

	user := model.User{}
	// 判断是否具有当前用户的管理权限
	if !user.HasUserManagerScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.UserPermissionsError()
	}

	if err := user.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}
