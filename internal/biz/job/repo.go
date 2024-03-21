package job

import "github.com/limes-cloud/kratosx"

type Repo interface {
	GetJobById(ctx kratosx.Context, id uint32) (*Job, error)
	GetJobByKeyword(ctx kratosx.Context, keyword string) (*Job, error)
	PageJob(ctx kratosx.Context, in *PageJobRequest) ([]*Job, uint32, error)
	AddJob(ctx kratosx.Context, in *Job) (uint32, error)
	UpdateJob(ctx kratosx.Context, in *Job) error
	DeleteJob(ctx kratosx.Context, id uint32) error
}
