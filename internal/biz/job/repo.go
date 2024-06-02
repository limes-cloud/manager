package job

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// ListJob 获取职位信息列表
	ListJob(ctx kratosx.Context, req *ListJobRequest) ([]*Job, uint32, error)

	// CreateJob 创建职位信息
	CreateJob(ctx kratosx.Context, req *Job) (uint32, error)

	// UpdateJob 更新职位信息
	UpdateJob(ctx kratosx.Context, req *Job) error

	// DeleteJob 删除职位信息
	DeleteJob(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetJob 获取指定的职位信息
	GetJob(ctx kratosx.Context, id uint32) (*Job, error)

	// GetJobByKeyword 获取指定的职位信息
	GetJobByKeyword(ctx kratosx.Context, keyword string) (*Job, error)
}
