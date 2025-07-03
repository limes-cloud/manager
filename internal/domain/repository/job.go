package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Job interface {
	// GetJob 获取指定的部门信息
	GetJob(ctx kratosx.Context, id uint32) (*entity.Job, error)

	// GetJobByKeyword 获取指定的部门信息
	GetJobByKeyword(ctx kratosx.Context, keyword string) (*entity.Job, error)

	// ListJob 获取部门信息列表
	ListJob(ctx kratosx.Context, req *types.ListJobRequest) ([]*entity.Job, error)

	// CreateJob 创建部门信息
	CreateJob(ctx kratosx.Context, req *entity.Job) (uint32, error)

	// UpdateJob 更新部门信息
	UpdateJob(ctx kratosx.Context, req *entity.Job) error

	// DeleteJob 删除部门信息
	DeleteJob(ctx kratosx.Context, id uint32) error

	// GetJobDataScope 获取指定用户的部门权限
	GetJobDataScope(ctx kratosx.Context, uid uint32) (bool, []uint32, error)

	// HasJobPurview 是否具有指定的部门权限
	HasJobPurview(ctx kratosx.Context, uid uint32, jids []uint32) (bool, error)
}
