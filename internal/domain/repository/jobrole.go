package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type JobRole interface {
	// ListJobRole 获取指定菜单的角色列表
	ListJobRole(ctx core.Context, req *types.ListJobRoleRequest) ([]*entity.Role, uint32, error)
	// CreateJobRole 创建角色菜单的关联信息
	CreateJobRole(ctx core.Context, jobId, roleId uint32) error
	// DeleteJobRole 删除角色菜单关联信息
	DeleteJobRole(ctx core.Context, jobId, roleId uint32) error
}
