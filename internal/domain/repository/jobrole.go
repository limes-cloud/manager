package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type JobRole interface {
	// ListRoleJob 获取角色的菜单列表
	ListRoleJob(ctx core.Context, req *types.ListRoleJobRequest) ([]*entity.Job, uint32, error)
	// ListJobRole 获取指定菜单的角色列表
	ListJobRole(ctx core.Context, req *types.ListJobRoleRequest) ([]*entity.Role, uint32, error)
	// CreateJobRoles 创建角色菜单的关联信息
	CreateJobRoles(ctx core.Context, drs []*entity.JobRole) error
	// DeleteJobRoles 删除角色菜单关联信息
	DeleteJobRoles(ctx core.Context, drs []*entity.JobRole) error
}
