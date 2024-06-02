package data

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	biz "github.com/limes-cloud/manager/internal/biz/job"
	"github.com/limes-cloud/manager/internal/data/model"
	"google.golang.org/protobuf/proto"
)

type jobRepo struct {
}

func NewJobRepo() biz.Repo {
	return &jobRepo{}
}

// ToJobEntity model转entity
func (r jobRepo) ToJobEntity(m *model.Job) *biz.Job {
	e := &biz.Job{}
	_ = valx.Transform(m, e)
	return e
}

// ToJobModel entity转model
func (r jobRepo) ToJobModel(e *biz.Job) *model.Job {
	m := &model.Job{}
	_ = valx.Transform(e, m)
	return m
}

// ListJob 获取列表
func (r jobRepo) ListJob(ctx kratosx.Context, req *biz.ListJobRequest) ([]*biz.Job, uint32, error) {
	var (
		bs    []*biz.Job
		ms    []*model.Job
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Job{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToJobEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateJob 创建数据
func (r jobRepo) CreateJob(ctx kratosx.Context, req *biz.Job) (uint32, error) {
	m := r.ToJobModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// GetJob 获取指定的数据
func (r jobRepo) GetJob(ctx kratosx.Context, id uint32) (*biz.Job, error) {
	var (
		m  = model.Job{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToJobEntity(&m), nil
}

// UpdateJob 更新数据
func (r jobRepo) UpdateJob(ctx kratosx.Context, req *biz.Job) error {
	return ctx.DB().Updates(r.ToJobModel(req)).Error
}

// DeleteJob 删除数据
func (r jobRepo) DeleteJob(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.Job{})
	return uint32(db.RowsAffected), db.Error
}

// GetJobByKeyword 获取指定数据
func (r jobRepo) GetJobByKeyword(ctx kratosx.Context, keyword string) (*biz.Job, error) {
	var (
		m  = model.Job{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToJobEntity(&m), nil
}
