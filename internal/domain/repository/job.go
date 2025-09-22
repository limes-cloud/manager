package repository

import (
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Job interface {
	// ListJob 获取职位列表
	ListJob(ctx core.Context, req *types.ListJobRequest) ([]*entity.Job, uint32, error)

	// CreateJob 创建职位
	CreateJob(ctx core.Context, req *entity.Job) (uint32, error)

	// UpdateJob 更新职位
	UpdateJob(ctx core.Context, req *entity.Job) error

	// DeleteJob 删除职位
	DeleteJob(ctx core.Context, id uint32) error
}
