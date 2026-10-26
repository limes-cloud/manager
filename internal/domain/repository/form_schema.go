package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type UserForm interface {
	// GetUserForm 获取表单（by id 或 scene）
	GetUserForm(ctx core.Context, req *types.GetUserFormRequest) (*entity.UserForm, error)

	// ListUserForm 分页查询表单列表
	ListUserForm(ctx core.Context, req *types.ListUserFormRequest) ([]*entity.UserForm, uint32, error)

	// CreateUserForm 创建表单
	CreateUserForm(ctx core.Context, ent *entity.UserForm) (uint32, error)

	// UpdateUserForm 更新表单
	UpdateUserForm(ctx core.Context, ent *entity.UserForm) error

	// DeleteUserForm 删除表单
	DeleteUserForm(ctx core.Context, id uint32) error
}
