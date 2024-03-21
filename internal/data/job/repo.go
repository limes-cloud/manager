package job

import (
	"github.com/limes-cloud/kratosx"

	biz "github.com/limes-cloud/manager/internal/biz/job"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

func (u *repo) GetJobById(ctx kratosx.Context, id uint32) (*biz.Job, error) {
	var job biz.Job
	return &job, ctx.DB().First(&job, "id=?", id).Error
}

func (u *repo) GetJobByKeyword(ctx kratosx.Context, keyword string) (*biz.Job, error) {
	var job biz.Job
	return &job, ctx.DB().First(&job, "keyword=?", keyword).Error
}

func (u *repo) PageJob(ctx kratosx.Context, req *biz.PageJobRequest) ([]*biz.Job, uint32, error) {
	var list []*biz.Job
	var total int64
	db := ctx.DB().Model(biz.Job{})
	if req.Keyword != nil {
		db.Where("keyword=?", *req.Keyword)
	}
	if req.Name != nil {
		db.Where("name like ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *repo) AddJob(ctx kratosx.Context, job *biz.Job) (uint32, error) {
	return job.ID, ctx.DB().Create(job).Error
}

func (u *repo) UpdateJob(ctx kratosx.Context, job *biz.Job) error {
	return ctx.DB().Updates(job).Error
}

func (u *repo) DeleteJob(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Job{}, "id=?", id).Error
}
