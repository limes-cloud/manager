package service

import (
	"time"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/kratosx/pkg/valx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/pkg/md"
	"github.com/limes-cloud/manager/internal/types"
)

const (
	ChangePwCaptchaType  = "captcha"
	ChangePwPasswordType = "password"

	pwdCaptchaKey   = "changePassword"
	loginCaptchaKey = "login"

	passwordCert = "login"
)

type User struct {
	conf    *conf.Config
	repo    repository.User
	dept    repository.Department
	role    repository.Role
	job     repository.Job
	file    repository.File
	address repository.Address
}

func NewUser(
	config *conf.Config,
	repo repository.User,
	dept repository.Department,
	role repository.Role,
	job repository.Job,
	file repository.File,
	address repository.Address,
) *User {
	return &User{
		conf:    config,
		repo:    repo,
		dept:    dept,
		role:    role,
		job:     job,
		file:    file,
		address: address,
	}
}

// GetUser 获取指定的用户信息
func (u *User) GetUser(ctx kratosx.Context, req *types.GetUserRequest) (*entity.User, error) {
	var (
		user *entity.User
		err  error
	)
	isPurview := func(ctx kratosx.Context, user *entity.User) error {
		// 获取当前用户部门全下
		all, scope, err := u.role.GetDataScope(ctx, md.UserId(ctx))
		if err != nil {
			return errors.DatabaseError()
		}

		if all {
			return nil
		}

		uni := valx.New(scope)
		for _, id := range user.GetDepartmentIds() {
			if uni.Has(id) {
				return nil
			}
		}
		return errors.DepartmentPurviewError()
	}

	if req.Id != nil {
		user, err = u.repo.GetUser(ctx, *req.Id)
		if err != nil {
			return nil, errors.DatabaseError()
		}
		if err := isPurview(ctx, user); err != nil {
			return nil, err
		}
	} else if req.Phone != nil {
		user, err = u.repo.GetUserByPhone(ctx, *req.Phone)
		if err != nil {
			return nil, errors.DatabaseError()
		}
		if err := isPurview(ctx, user); err != nil {
			return nil, err
		}
	} else if req.Email != nil {
		user, err = u.repo.GetUserByEmail(ctx, *req.Email)
		if err != nil {
			return nil, errors.DatabaseError()
		}
		if err := isPurview(ctx, user); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.ParamsError()
	}
	return user, err
}

// ListUser 获取用户信息列表
func (u *User) ListUser(ctx kratosx.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	all, scopes, err := u.role.GetDataScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, 0, errors.DatabaseError(err.Error())
	}
	if !all {
		// 没有权限仅查看当前用户
		if scopes == nil {
			req.ID = proto.Uint32(md.UserId(ctx))
		} else {
			req.DepartmentIds = scopes
		}
	}

	list, total, err := u.repo.ListUser(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}

	for ind, item := range list {
		if item.Avatar != nil {
			url := u.file.GetFileURL(ctx, *item.Avatar)
			list[ind].Avatar = &url
		}
	}
	return list, total, nil
}

// CreateUser 创建用户信息
func (u *User) CreateUser(ctx kratosx.Context, req *entity.User) (uint32, error) {
	// 判断是否具有部门权限
	hasDeptPurview, err := u.role.HasDataPurview(ctx, md.UserId(ctx), req.GetDepartmentIds())
	if err != nil {
		ctx.Logger().Warnw("msg", "get dept purview error", "err", err.Error())
		return 0, errors.DatabaseError()
	}
	if !hasDeptPurview {
		return 0, errors.DepartmentPurviewError()
	}

	// 判断是否具有角色权限
	all, scopes, err := u.role.GetRoleScope(ctx, md.RoleIds(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "get role scopes error", "err", err.Error())
		return 0, errors.DatabaseError()
	}
	if !all {
		inr := valx.New(scopes)
		for _, ur := range req.UserRoles {
			if !inr.Has(ur.RoleId) {
				return 0, errors.RolePurviewError()
			}
		}
	}

	// 创建用户信息
	req.Nickname = req.Name
	req.Avatar = &u.conf.DefaultUserAvatar
	req.Password = crypto.EncodePwd(u.conf.DefaultUserPassword)
	req.Status = proto.Bool(true)

	id, err := u.repo.CreateUser(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateUser 更新用户信息
func (u *User) UpdateUser(ctx kratosx.Context, user *entity.User) error {
	var curUserId = md.UserId(ctx)

	// 系统数据不允许修改
	if user.Id == 1 {
		return errors.EditSystemDataError()
	}

	if curUserId == user.Id {
		return errors.UpdateError("不能修改当前用户")
	}

	// 获取用户基础信息
	oldUser, err := u.repo.GetUser(ctx, user.Id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get base user error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 获取当前用的部门权限
	all, scopes, err := u.role.GetDataScope(ctx, curUserId)
	if err != nil {
		ctx.Logger().Warnw("msg", "get dept purview error", "err", err.Error())
		return errors.DatabaseError()
	}
	if !all {
		var (
			inr      = valx.New(scopes)
			hasScope = false
		)

		// 判断是否具有用户修改权限
		for _, id := range oldUser.GetDepartmentIds() {
			if inr.Has(id) {
				hasScope = true
				break
			}
		}
		if !hasScope {
			return errors.UserPurviewError()
		}

		// 判断是否具有修改后的部门权限
		for _, id := range user.GetDepartmentIds() {
			if !inr.Has(id) {
				return errors.DepartmentPurviewError()
			}
		}

		// 获取修改的用户部门在当前用户不可见的id
		// 将不可见的ID追加到新建里面
		oldDeptIds := oldUser.GetDepartmentIds()
		for _, id := range oldDeptIds {
			if !inr.Has(id) {
				user.UserDepartments = append(user.UserDepartments, &entity.UserDepartment{
					UserId:       curUserId,
					DepartmentId: id,
				})
			}
		}
	}

	// 获取当前的角色权限
	all, scopes, err = u.role.GetRoleScope(ctx, md.RoleIds(ctx))
	if !all {
		var inr = valx.New(scopes)
		// 判断是否具有修改后的角色权限
		for _, id := range user.GetRoleIds() {
			if !inr.Has(id) {
				return errors.RolePurviewError()
			}
		}

		// 获取修改的用户角色在当前用户不可见的id
		// 将不可见的ID追加到新建里面
		oldRoleIds := oldUser.GetRoleIds()
		for _, id := range oldRoleIds {
			if !inr.Has(id) {
				user.UserRoles = append(user.UserRoles, &entity.UserRole{
					UserId: curUserId,
					RoleId: id,
				})
			}
		}
	}

	// 获取当前的职位权限
	// all, scopes, err = u.job.GetJobDataScope(ctx, curUserId)
	// if !all {
	//	var inr = valx.New(scopes)
	//	// 判断是否具有修改后的职位权限
	//	for _, id := range user.GetJobIds() {
	//		if !inr.Has(id) {
	//			return errors.JobPurviewError()
	//		}
	//	}
	//
	//	// 获取修改的用户部门在当前用户不可见的id
	//	// 将不可见的ID追加到新建里面
	//	oldJobIds := oldUser.GetJobIds()
	//	for _, id := range oldJobIds {
	//		if !inr.Has(id) {
	//			user.UserJobs = append(user.UserJobs, &entity.UserJob{
	//				UserId: curUserId,
	//				JobId:  id,
	//			})
	//		}
	//	}
	// }

	// 更新用户
	user.Token = proto.String("")
	if err := u.repo.ForceUpdateUser(ctx, user); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateUserStatus 更新用户信息状态 fixed code
func (u *User) UpdateUserStatus(ctx kratosx.Context, id uint32, status bool) error {
	// 系统数据不允许修改
	if id == 1 {
		return errors.EditSystemDataError()
	}
	if md.UserId(ctx) == id {
		return errors.UpdateError("不能修改当前用户")
	}

	// 获取用户基础信息
	oldUser, err := u.repo.GetBaseUser(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get base user error", "err", err.Error())
		return errors.DatabaseError()
	}

	if err := u.hasUserScope(ctx, id); err != nil {
		return err
	}

	// 更新角色状态
	if err := u.repo.UpdateUserStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}

	// 如果是禁用用户
	expire := ctx.Config().App().JWT.Expire.Seconds()
	if !status && oldUser.Token != nil && oldUser.LoggedAt > time.Now().Unix()-int64(expire) {
		ctx.JWT().AddBlacklist(*oldUser.Token)
	}

	return nil
}

// DeleteUser 删除用户信息
func (u *User) DeleteUser(ctx kratosx.Context, id uint32) error {
	// 系统数据不允许修改
	if id == 1 {
		return errors.EditSystemDataError()
	}
	if md.UserId(ctx) == id {
		return errors.DeleteError("不能删除当前用户")
	}

	if err := u.hasUserScope(ctx, id); err != nil {
		return err
	}

	if err := u.repo.DeleteUser(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ResetUserPassword 重置用户密码
func (u *User) ResetUserPassword(ctx kratosx.Context, id uint32) error {
	// 系统数据不允许修改
	if id == 1 {
		return errors.EditSystemDataError()
	}

	if err := u.hasUserScope(ctx, id); err != nil {
		return err
	}

	if err := u.repo.UpdateUser(ctx, &entity.User{
		BaseModel: ktypes.BaseModel{Id: id},
		Password:  crypto.EncodePwd(u.conf.DefaultUserPassword),
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}

	return nil
}

// GetCurrentUser 获取当前的用户信息
func (u *User) GetCurrentUser(ctx kratosx.Context) (*entity.User, error) {
	res, err := u.repo.GetUser(ctx, md.UserId(ctx))
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	if res.Avatar != nil {
		url := u.file.GetFileURL(ctx, *res.Avatar)
		res.Avatar = &url
	}
	return res, nil
}

// UpdateCurrentUser 更新当前的基础信息
func (u *User) UpdateCurrentUser(ctx kratosx.Context, req *types.UpdateCurrentUserRequest) error {
	if err := u.repo.UpdateUser(ctx, &entity.User{
		BaseModel: ktypes.BaseModel{Id: md.UserId(ctx)},
		Avatar:    req.Avatar,
		Nickname:  req.Nickname,
		Gender:    req.Gender,
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}

	return nil
}

// UpdateCurrentUserSetting 保存当前用户设置
func (u *User) UpdateCurrentUserSetting(ctx kratosx.Context, setting string) error {
	if err := u.repo.UpdateUser(ctx, &entity.User{
		BaseModel: ktypes.BaseModel{Id: md.UserId(ctx)},
		Setting:   &setting,
	}); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// SendCurrentUserCaptcha 发送当前用户验证吗
func (u *User) SendCurrentUserCaptcha(ctx kratosx.Context, tp string) (*types.SendCurrentUserCaptchaReply, error) {
	tps := []string{pwdCaptchaKey, loginCaptchaKey}
	if !valx.InList(tps, tp) {
		return nil, errors.ParamsError()
	}

	user, err := u.repo.GetUser(ctx, md.UserId(ctx))
	if err != nil {
		return nil, errors.GetError(err.Error())
	}

	resp, err := ctx.Captcha().Email(tp, ctx.ClientIP(), user.Email)
	if err != nil {
		return nil, errors.SendCaptchaError(err.Error())
	}

	return &types.SendCurrentUserCaptchaReply{
		Uuid:   resp.ID(),
		Expire: uint32(resp.Expire().Seconds()),
	}, nil
}

// UpdateCurrentUserPassword 修改当前用户密码
func (u *User) UpdateCurrentUserPassword(ctx kratosx.Context, req *types.UpdateCurrentUserPasswordRequest) error {
	user, err := u.repo.GetBaseUser(ctx, md.UserId(ctx))
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	switch u.conf.ChangePasswordType {
	case ChangePwCaptchaType:
		if req.CaptchaId == nil || req.Captcha == nil {
			return errors.ParamsError()
		}
		if err := ctx.Captcha().VerifyEmail(pwdCaptchaKey, ctx.ClientIP(), *req.CaptchaId, *req.Captcha, user.Email); err != nil {
			return errors.VerifyCaptchaError()
		}
	case ChangePwPasswordType:
		if req.OldPassword == nil {
			return errors.ParamsError()
		}
		if !crypto.CompareHashPwd(user.Password, *req.OldPassword) {
			return errors.PasswordError()
		}
	default:
		return errors.SystemError("验证方式配置错误")
	}

	nu := entity.User{
		BaseModel: ktypes.BaseModel{Id: md.UserId(ctx)},
		Password:  crypto.EncodePwd(req.Password),
	}
	if err := u.repo.UpdateUser(ctx, &nu); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

func (u *User) hasUserScope(ctx kratosx.Context, uid uint32) error {
	// 获取用户基础信息
	oldUser, err := u.repo.GetUser(ctx, uid)
	if err != nil {
		ctx.Logger().Warnw("msg", "get base user error", "err", err.Error())
		return errors.DatabaseError()
	}

	// 获取当前用的部门权限
	all, scopes, err := u.role.GetDataScope(ctx, md.UserId(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "get dept purview error", "err", err.Error())
		return errors.DatabaseError()
	}

	if !all {
		var (
			inr      = valx.New(scopes)
			hasScope = false
		)

		// 判断是否具有用户修改权限
		for _, id := range oldUser.GetDepartmentIds() {
			if inr.Has(id) {
				hasScope = true
				break
			}
		}
		if !hasScope {
			return errors.UserPurviewError()
		}
	}
	return nil
}
