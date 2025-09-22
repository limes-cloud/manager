package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Job struct{}

var (
	jobIns  *Job
	jobOnce sync.Once
)

func NewJob() *Job {
	jobOnce.Do(func() {
		jobIns = &Job{}
	})
	return jobIns
}

// ListJob 获取列表
func (infra *Job) ListJob(ctx core.Context, req *types.ListJobRequest) ([]*entity.Job, uint32, error) {
	var (
		list  []*entity.Job
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.Job{}).Select(fs)
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Keyword != nil {
		db = db.Where("keyword LIKE ?", *req.Keyword+"%")
	}

	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateJob 创建数据
func (infra *Job) CreateJob(ctx core.Context, tg *entity.Job) (uint32, error) {
	return tg.Id, ctx.DB().Create(tg).Error
}

// UpdateJob 更新数据
func (infra *Job) UpdateJob(ctx core.Context, tg *entity.Job) error {
	return ctx.DB().Updates(tg).Error
}

// DeleteJob 删除数据
func (infra *Job) DeleteJob(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Job{}).Error
}
