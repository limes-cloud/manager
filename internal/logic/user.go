package logic

import (
	"fmt"
	v1 "manager/api/v1"
	"manager/config"
	"manager/consts"
	"manager/internal/model"
	"manager/pkg/md"
	"manager/pkg/util"

	"google.golang.org/protobuf/proto"

	"gorm.io/gorm"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/kratos"
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
func (r *User) Get(ctx kratos.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user := model.User{}
	if err := user.OneByID(ctx, in.Id); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	// 获取用户的全部角色
	roles := []*model.Role{}
	rm := model.UserRole{}
	rms, _ := rm.UserRoles(ctx, in.Id)
	for _, item := range rms {
		roles = append(roles, item.Role)
	}

	reply := v1.GetUserReply{}
	if err := util.Transform(user, &reply.User); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	if err := util.Transform(roles, &reply.User.Roles); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	return &reply, nil
}

// ChangePasswordCaptcha 重置用户密码验证码
func (r *User) ChangePasswordCaptcha(ctx kratos.Context, in *emptypb.Empty) (*v1.ChangeUserPasswordCaptchaReply, error) {
	user := model.User{}
	if err := user.OneByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.ErrorNotFound()
	}

	resp, err := ctx.Captcha().Email(captcha, ctx.ClientIP(), user.Email)
	if err != nil {
		return nil, v1.ErrorSendEmailCaptchaFormat(err.Error())
	}

	return &v1.ChangeUserPasswordCaptchaReply{
		Uuid:   resp.ID(),
		Expire: uint32(resp.Expire().Seconds()),
	}, nil
}

// ChangePassword 重置用户密码
func (r *User) ChangePassword(ctx kratos.Context, in *v1.ChangeUserPasswordRequest) (*emptypb.Empty, error) {
	// 查询用户
	user := model.User{}
	if err := user.OneByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.ErrorNotFound()
	}

	// 验证验证码
	if err := ctx.Captcha().VerifyEmail(captcha, ctx.ClientIP(), in.CaptchaId, in.Captcha); err != nil {
		return nil, v1.ErrorVerifyCaptcha()
	}

	// 修改密码
	nu := model.User{}
	nu.ID = user.ID
	nu.Password = in.Password
	if err := nu.Update(ctx); err != nil {
		return nil, v1.ErrorDatabase()
	}

	return nil, nil
}

// Enable 启用用户
func (r *User) Enable(ctx kratos.Context, in *v1.EnableUserRequest) (*emptypb.Empty, error) {
	// 查询用户
	user := model.User{}
	if err := user.OneByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.ErrorNotFound()
	}

	// 判断当前用户是否具有部门的权限
	depIds, _ := user.DataScope(ctx, md.UserId(ctx))
	if !util.InList(depIds, user.DepartmentID) {
		return nil, v1.ErrorDepartmentPermissions()
	}

	// 判断是否具有当前用户的管理权限
	if !user.HasUserScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.ErrorUserPermissions()
	}

	// 将用户启用
	nu := model.User{}
	nu.ID = in.Id
	nu.Status = proto.Bool(true)
	nu.Disabled = proto.String("")
	if err := nu.Update(ctx); err != nil {
		return nil, v1.ErrorDatabase()
	}

	return nil, nil
}

// Disable 禁用用户
func (r *User) Disable(ctx kratos.Context, in *v1.DisableUserRequest) (*emptypb.Empty, error) {
	// 超级管理员不允许修改
	if in.Id == consts.SuperAdmin {
		return nil, v1.ErrorEditSystemData()
	}

	// 超级管理员不允许修改
	if in.Id == md.UserId(ctx) {
		return nil, v1.ErrorDisableSelfUser()
	}

	// 查询用户
	user := model.User{}
	if err := user.OneByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.ErrorNotFound()
	}

	// 判断是否具有当前用户的管理权限
	if !user.HasUserScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.ErrorUserPermissions()
	}

	// 将用户禁用
	nu := model.User{}
	nu.ID = in.Id
	nu.Status = proto.Bool(false)
	nu.Disabled = proto.String(in.Desc)
	if err := nu.Update(ctx); err != nil {
		return nil, v1.ErrorDatabase()
	}

	return nil, nil
}

// Offline 对当前登陆用户进行下线
func (r *User) Offline(ctx kratos.Context, in *v1.OfflineUserRequest) (*emptypb.Empty, error) {
	// 查询用户
	user := model.User{}
	if err := user.OneByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.ErrorNotFound()
	}

	// 判断是否具有当前用户的管理权限
	if !user.HasUserScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.ErrorUserPermissions()
	}

	// 将用户下线
	ctx.JWT().AddBlacklist(user.Token)
	return nil, nil
}

// ResetPassword 重置用户密码
func (r *User) ResetPassword(ctx kratos.Context, in *v1.ResetUserPasswordRequest) (*emptypb.Empty, error) {
	// 超级管理员不允许修改
	if in.Id == consts.SuperAdmin {
		return nil, v1.ErrorEditSystemData()
	}

	// 查询用户信息
	user := model.User{}
	if err := user.OneByID(ctx, in.Id); err != nil {
		return nil, v1.ErrorNotFound()
	}

	// 判断是否具有当前用户的管理权限
	if !user.HasUserScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.ErrorUserPermissions()
	}

	// 重置密码
	nu := model.User{}
	nu.ID = in.Id
	nu.Password = r.conf.DefaultUserPassword
	if err := nu.Update(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return nil, nil
}

// Current 查询当前用户信息
func (r *User) Current(ctx kratos.Context) (*v1.GetUserReply, error) {
	return r.Get(ctx, &v1.GetUserRequest{
		Id: md.UserId(ctx),
	})
}

func (r *User) Page(ctx kratos.Context, in *v1.PageUserRequest) (*v1.PageUserReply, error) {
	user := model.User{}
	// 获取用户所管理的部门
	ids, err := user.DataScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	list, total, err := user.Page(ctx, &model.PageOptions{
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
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	reply := v1.PageUserReply{Total: total}
	// 进行数据转换
	if err = util.Transform(list, &reply.List); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	return &reply, nil
}

// Add 添加用户信息
func (r *User) Add(ctx kratos.Context, in *v1.AddUserRequest) (*emptypb.Empty, error) {
	user := model.User{}

	// 判断当前用户是否具有部门的权限
	depIds, _ := user.DataScope(ctx, md.UserId(ctx))
	if !util.InList(depIds, in.DepartmentId) {
		return nil, v1.ErrorDepartmentPermissions()
	}

	// 判断当前用户是否具有角色的权限
	roleIds, _ := user.RoleScope(ctx, md.UserId(ctx))
	for _, id := range in.RoleIds {
		if !util.InList(roleIds, id) {
			return nil, v1.ErrorDepartmentPermissionsFormat(fmt.Sprintf("role_id:%d", id))
		}
	}

	// 进行数据转换
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	// 设置默认值
	user.Nickname = user.Name
	user.Avatar = r.conf.DefaultUserAvatar
	user.Password = util.ParsePwd(r.conf.DefaultUserPassword)
	user.RoleID = in.RoleIds[0]

	// 创建用户
	if err := user.Create(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	// 更新用户角色信息
	userRole := model.UserRole{}
	if err := userRole.Update(ctx, user.ID, in.RoleIds); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}
	return nil, nil
}

// Update 更新用户信息
func (r *User) Update(ctx kratos.Context, in *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	// 超级管理员不允许修改
	if in.Id == consts.SuperAdmin {
		return nil, v1.ErrorEditSystemData()
	}

	user := model.User{}
	// 判断当前用户是否具有部门的权限
	depIds, _ := user.DataScope(ctx, md.UserId(ctx))
	if !util.InList(depIds, in.DepartmentId) {
		return nil, v1.ErrorDepartmentPermissions()
	}

	// 判断当前用户是否具有角色的权限
	roleIds, _ := user.RoleScope(ctx, md.UserId(ctx))
	for _, id := range in.RoleIds {
		if !util.InList(roleIds, id) {
			return nil, v1.ErrorDepartmentPermissionsFormat(fmt.Sprintf("role_id:%d", id))
		}
	}

	// 判断是否具有当前用户的管理权限
	if !user.HasUserScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.ErrorUserPermissions()
	}

	// 转换数据格式
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	// 将角色设置为更新后的第一个角色
	user.RoleID = in.RoleIds[0]

	// 更新用户信息
	if err := user.Update(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	// 更新用户角色信息
	userRole := model.UserRole{}
	if err := userRole.Update(ctx, user.ID, in.RoleIds); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return nil, nil
}

// UpdateBasic 更新用户基础信息
func (r *User) UpdateBasic(ctx kratos.Context, in *v1.UpdateUserBasicRequest) (*emptypb.Empty, error) {
	user := model.User{
		BaseModel: model.BaseModel{ID: md.UserId(ctx)},
		Nickname:  in.Nickname,
		Gender:    in.Gender,
	}

	// 更新用户信息
	if err := user.Update(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return nil, nil
}

// Delete 删除用户
func (r *User) Delete(ctx kratos.Context, in *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	if in.Id == consts.SuperAdmin {
		return nil, v1.ErrorDeleteSystemData()
	}

	user := model.User{}
	// 判断是否具有当前用户的管理权限
	if !user.HasUserScope(ctx, md.UserId(ctx), in.Id) {
		return nil, v1.ErrorUserPermissions()
	}

	if err := user.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return nil, nil
}
