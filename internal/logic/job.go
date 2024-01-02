package logic

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	v1 "github.com/limes-cloud/manager/api/v1"
	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/internal/model"
	"github.com/limes-cloud/manager/pkg/util"
)

type Job struct {
	conf *config.Config
}

func NewJob(conf *config.Config) *Job {
	return &Job{
		conf: conf,
	}
}

// GetUserJob 查询指定用户的职位信息，提供给外部rpc调用
func (r *Job) GetUserJob(ctx kratosx.Context, in *v1.GetUserJobRequest) (*v1.GetUserJobReply, error) {
	var jobs []*model.Job

	uj := model.UserJob{}
	ujs, err := uj.UserJobs(ctx, in.Id)
	if err != nil {
		return nil, v1.DatabaseError()
	}

	for _, item := range ujs {
		jobs = append(jobs, item.Job)
	}

	reply := v1.GetUserJobReply{}
	if err := util.Transform(jobs, &reply.Jobs); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// Get 查询指定的职位信息，提供给外部rpc调用
func (r *Job) Get(ctx kratosx.Context, in *v1.GetJobRequest) (*v1.GetJobReply, error) {
	job := model.Job{}
	if err := job.FindByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.GetJobReply{}
	if err := util.Transform(job, &reply.Job); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

func (r *Job) Page(ctx kratosx.Context, in *v1.PageJobRequest) (*v1.PageJobReply, error) {
	job := model.Job{}
	list, total, err := job.Page(ctx, &types.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.Keyword != nil {
				db = db.Where("keyword=?", *in.Keyword)
			}
			return db
		},
	})
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.PageJobReply{Total: total}
	// 进行数据转换
	if err = util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// Add 添加职位信息
func (r *Job) Add(ctx kratosx.Context, in *v1.AddJobRequest) (*emptypb.Empty, error) {
	job := model.Job{}

	// 进行数据转换
	if err := util.Transform(in, &job); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 创建职位
	if err := job.Create(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Update 更新职位信息
func (r *Job) Update(ctx kratosx.Context, in *v1.UpdateJobRequest) (*emptypb.Empty, error) {
	job := model.Job{}

	// 转换数据格式
	if err := util.Transform(in, &job); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 更新职位信息
	if err := job.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Delete 删除职位
func (r *Job) Delete(ctx kratosx.Context, in *v1.DeleteJobRequest) (*emptypb.Empty, error) {
	job := model.Job{}

	if err := job.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}
