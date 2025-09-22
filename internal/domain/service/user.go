package service

import (
	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type User struct {
	repo  repository.User
	scope repository.Scope
}

func NewUser(repo repository.User, scope repository.Scope) *User {
	return &User{repo: repo, scope: scope}
}

// GetCurrentUser 获取当前的用户信息
func (u *User) GetCurrentUser(ctx core.Context) (*entity.User, error) {
	res, err := u.repo.GetUser(ctx, ctx.Auth().UserId)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
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

	// 判断是否具有用户权限
	if !u.scope.HasDeptScope(ctx, entity.UserEntityName, res.Dept.Id) {
		return nil, errors.DeptScopeError()
	}

	return res, nil
}

// ListUser 获取租户列表
func (u *User) ListUser(ctx core.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	list, total, err := u.repo.ListUser(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list user error", "err", err.Error())
		return nil, 0, errors.ListError()
	}

	return list, total, nil
}

// CreateUser 创建租户
func (u *User) CreateUser(ctx core.Context, req *entity.User) (uint32, error) {
	req.Password = crypto.EncodePwd(ctx.Config().DefaultUserPassword)

	id, err := u.repo.CreateUser(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create user error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// ResetPassword 更新用户
func (u *User) ResetPassword(ctx core.Context, uid uint32) error {
	req := &entity.User{
		BaseTenantDeptModel: model.BaseTenantDeptModel{
			Id: uid,
		},
		Password: crypto.EncodePwd(ctx.Config().DefaultUserPassword),
	}
	if err := u.repo.UpdateUser(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update user error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// UpdateUser 更新租户
func (u *User) UpdateUser(ctx core.Context, req *entity.User) error {
	if err := u.repo.UpdateUser(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update user error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// DeleteUser 删除租户
func (u *User) DeleteUser(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteUser(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete user error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}
