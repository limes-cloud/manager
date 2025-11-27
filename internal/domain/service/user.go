package service

import (
	"time"

	"github.com/samber/lo"

	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/pkg/jwt"
	"github.com/limes-cloud/manager/internal/types"
)

type User struct {
	repo     repository.User
	app      repository.App
	af       repository.AppField
	info     repository.Userinfo
	setting  repository.UserSetting
	tenant   repository.Tenant
	userdept repository.UserDept
	az       repository.Authorize
	scope    repository.Scope
	tad      repository.TenantAdmin
}

func NewUser(
	repo repository.User,
	app repository.App,
	af repository.AppField,
	info repository.Userinfo,
	setting repository.UserSetting,
	tenant repository.Tenant,
	userdept repository.UserDept,
	az repository.Authorize,
	scope repository.Scope,
	tad repository.TenantAdmin,
) *User {
	return &User{
		repo:     repo,
		app:      app,
		af:       af,
		info:     info,
		setting:  setting,
		tenant:   tenant,
		userdept: userdept,
		az:       az,
		scope:    scope,
		tad:      tad,
	}
}

// GetCurrentUser 获取当前的用户信息
func (u *User) GetCurrentUser(ctx core.Context, req *types.GetCurrentUserRequest) (*entity.User, error) {
	res, err := u.GetUser(ctx, &types.GetUserRequest{
		Id:  &ctx.Auth().UserId,
		App: req.App,
	})
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// UpdateCurrentUser 更新租户
func (u *User) UpdateCurrentUser(ctx core.Context, req *types.UpdateCurrentUserRequest) error {
	if err := u.repo.UpdateUser(ctx, &entity.User{
		BaseTenantModel: model.BaseTenantModel{
			Id: ctx.Auth().UserId,
		},
		Avatar:   req.Avatar,
		Nickname: req.Nickname,
	}); err != nil {
		ctx.Logger().Warnw("msg", "update user error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// UpdateCurrentUserPassword 更新用户密码
func (u *User) UpdateCurrentUserPassword(ctx core.Context, req *types.UpdateCurrentUserPasswordRequest) error {
	// 获取当前用户信息
	user, err := u.repo.GetUser(ctx, ctx.Auth().UserId)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}

	// 对比密码
	if !crypto.CompareHashPwd(user.Password, req.OldPassword) {
		return errors.PasswordError()
	}

	// 更新密码
	if err := u.repo.UpdateUser(ctx, &entity.User{
		BaseTenantModel: model.BaseTenantModel{
			Id: ctx.Auth().UserId,
		},
		Password: crypto.EncodePwd(req.Password),
	}); err != nil {
		ctx.Logger().Warnw("msg", "update user error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// UpdateCurrentUserSetting 更新租户
func (u *User) UpdateCurrentUserSetting(ctx core.Context, req *types.UpdateCurrentUserSettingRequest) error {
	// 获取应用信息
	app, err := u.app.GetAppByKeyword(req.App)
	if err != nil {
		return errors.GetError("获取应用信息失败")
	}

	if err := u.setting.UpsertUserSetting(ctx, &entity.UserSetting{
		UserId:  ctx.Auth().UserId,
		AppId:   app.Id,
		Setting: req.Setting,
	}); err != nil {
		ctx.Logger().Warnw("msg", "update user error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// GetUser 获取租户列表
func (u *User) GetUser(ctx core.Context, req *types.GetUserRequest) (*entity.User, error) {
	var (
		res *entity.User
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetUser(ctx, *req.Id)
	} else if req.Username != nil {
		res, err = u.repo.GetUserByUsername(ctx, *req.Username)
	} else {
		return nil, errors.ParamsError()
	}
	if err != nil {
		ctx.Logger().Warnw("msg", "get user error", "err", err.Error())
		return nil, errors.GetError()
	}

	getFieldReq := &types.ListUserinfoRequest{
		UserId: res.Id,
	}

	// 不需要获取应用相关信息，则直接返回
	if req.App != nil {
		app, err := u.app.GetAppByKeyword(*req.App)
		if err != nil {
			return nil, errors.GetError("获取应用信息失败")
		}

		// 获取用户的字段
		fields, err := u.af.GetAppFields(ctx, app.Id)
		getFieldReq.Fields = fields

		// 获取用户应用设置
		us, _ := u.setting.GetUserSetting(ctx, &types.GetUserSettingRequest{
			AppId:  app.Id,
			UserId: res.Id,
		})
		if us != nil {
			res.Setting = us.Setting
		}
	}

	// 获取用户应用信息
	infos, _ := u.info.ListUserinfo(ctx, &types.ListUserinfoRequest{
		UserId:   res.Id,
		TenantId: ctx.Auth().TenantId,
	})
	if infos != nil {
		res.Infos = infos
	}

	// 获取用户部门信息
	depts, _, _ := u.userdept.ListUserDept(ctx, &types.ListUserDeptRequest{
		UserId: res.Id,
	})
	if depts != nil {
		res.UserDepts = depts
	}

	return res, nil
}

// ListUser 获取租户列表
func (u *User) ListUser(ctx core.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	if !u.tad.IsAdmin(ctx.Auth().TenantId, ctx.Auth().UserId) {
		all, ids := u.scope.SystemDeptScopes(ctx, "user", entity.ActionRead)
		if !all {
			if len(ids) == 0 {
				req.InIds = []uint32{ctx.Auth().UserId}
			} else {
				req.InDeptIds = lo.Intersect(req.InDeptIds, ids)
			}
		}
	}

	if req.DeptId != nil {
		req.InDeptIds = append(req.InDeptIds, *req.DeptId)
	}
	if req.JobId != nil {
		req.InJobIds = append(req.InJobIds, *req.JobId)
	}
	if req.App != nil {
		app, err := u.app.GetAppByKeyword(*req.App)
		if err != nil {
			return nil, 0, errors.GetError(err.Error())
		}
		req.AppId = &app.Id
	}

	list, total, err := u.repo.ListUser(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list user error", "err", err.Error())
		return nil, 0, errors.ListError()
	}

	return list, total, nil
}

// CreateUser 创建租户
func (u *User) CreateUser(ctx core.Context, req *entity.User) (uint32, error) {
	// 获取租户信息
	tenant, err := u.tenant.GetTenant(ctx, ctx.Auth().TenantId)
	if err != nil {
		ctx.Logger().Warnw("msg", "get tenant error", "err", err.Error())
		return 0, errors.CreateError()
	}

	setting := tenant.GetSetting()
	req.Password = crypto.EncodePwd(setting.DefaultUserPassword)
	if req.Avatar == "" {
		req.Avatar = setting.DefaultUserAvatar
	}
	if req.Nickname == "" {
		req.Nickname = setting.DefaultUserNickname
	}

	// 创建用户信息
	var id uint32
	_ = ctx.Transaction(func(ctx core.Context) error {
		id, err = u.repo.CreateUser(ctx, req)
		if err != nil {
			ctx.Logger().Warnw("msg", "create user error", "err", err.Error())
			return err
		}
		for _, info := range req.Infos {
			info.UserId = id
			info.TenantId = ctx.Auth().TenantId
		}
		return u.info.UpsertUserinfo(ctx, req.Infos)
	})
	if err != nil {
		ctx.Logger().Warnw("msg", "create user error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateUser 更新租户
func (u *User) UpdateUser(ctx core.Context, req *entity.User) error {
	// 获取用户所属的部门
	deptIds := u.userdept.GetDeptIdsByUserId(req.Id)
	if !u.tad.IsAdmin(ctx.Auth().TenantId, ctx.Auth().UserId) {
		all, ids := u.scope.SystemDeptScopes(ctx, "user", entity.ActionUpdate)
		if !all && len(lo.Intersect(ids, deptIds)) == 0 {
			return errors.UserScopeError()
		}
	}

	return ctx.Transaction(func(ctx core.Context) error {
		if err := u.repo.UpdateUser(ctx, req); err != nil {
			ctx.Logger().Warnw("msg", "update user error", "err", err.Error())
			return errors.UpdateError(err.Error())
		}
		if err := u.info.UpsertUserinfo(ctx, req.Infos); err != nil {
			ctx.Logger().Warnw("msg", "update user info error", "err", err.Error())
			return errors.UpdateError(err.Error())
		}
		return nil
	})
}

// DeleteUser 删除租户
func (u *User) DeleteUser(ctx core.Context, id uint32) error {
	// 获取用户所属的部门
	deptIds := u.userdept.GetDeptIdsByUserId(id)
	if !u.tad.IsAdmin(ctx.Auth().TenantId, ctx.Auth().UserId) {
		all, ids := u.scope.SystemDeptScopes(ctx, "user", entity.ActionDelete)
		if !all && len(lo.Intersect(ids, deptIds)) == 0 {
			return errors.UserScopeError()
		}
	}

	if err := u.repo.DeleteUser(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete user error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}

// UpdateUserinfo 更新用户信息
func (u *User) UpdateUserinfo(ctx core.Context, list []*entity.Userinfo) error {
	if err := u.info.UpsertUserinfo(ctx, list); err != nil {
		ctx.Logger().Warnw("msg", "update user info error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// ResetPassword 更新用户
func (u *User) ResetPassword(ctx core.Context, uid uint32) error {
	// 获取租户信息
	tenant, err := u.tenant.GetTenant(ctx, ctx.Auth().TenantId)
	if err != nil {
		ctx.Logger().Warnw("msg", "get tenant error", "err", err.Error())
		return errors.CreateError()
	}

	setting := tenant.GetSetting()

	req := &entity.User{
		BaseTenantModel: model.BaseTenantModel{Id: uid},
		Password:        crypto.EncodePwd(setting.DefaultUserPassword),
	}
	if err := u.repo.UpdateUser(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update user error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// OfflineUser 下线用户
func (u *User) OfflineUser(ctx core.Context, req *types.OfflineUserRequest) error {
	// 获取当前用户的信息
	list, _, err := u.az.ListAuthorize(ctx, &types.ListAuthorizeRequest{
		UserId: value.Pointer(req.UserId),
		AppIds: req.AppIds,
	})
	if err != nil {
		ctx.Logger().Warnw("msg", "list authorize error", "err", err.Error())
		return errors.ListError()
	}

	for _, item := range list {
		// 获取应用信息
		app, err := u.app.GetApp(ctx, item.AppId)
		if err != nil {
			ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
			continue
		}

		jwtconf := app.GetSetting().JWT
		jwter := jwt.New(ctx.Context, &jwt.Config{
			App:     app.Keyword,
			Secret:  jwtconf.Secret,
			Expire:  time.Duration(jwtconf.Expire) * time.Second,
			Renewal: time.Duration(jwtconf.Renewal) * time.Second,
		})

		tokens := item.GetTokens()
		for token := range tokens {
			jwter.AddBlacklist(token)
		}
	}

	return nil
}
