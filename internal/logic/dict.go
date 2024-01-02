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

type Dict struct {
	conf *config.Config
}

func NewDict(conf *config.Config) *Dict {
	return &Dict{
		conf: conf,
	}
}

func (r *Dict) Page(ctx kratosx.Context, in *v1.PageDictRequest) (*v1.PageDictReply, error) {
	job := model.Dict{}
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

	reply := v1.PageDictReply{Total: total}
	// 进行数据转换
	if err = util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// Add 添加职位信息
func (r *Dict) Add(ctx kratosx.Context, in *v1.AddDictRequest) (*emptypb.Empty, error) {
	job := model.Dict{}

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
func (r *Dict) Update(ctx kratosx.Context, in *v1.UpdateDictRequest) (*emptypb.Empty, error) {
	job := model.Dict{}

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
func (r *Dict) Delete(ctx kratosx.Context, in *v1.DeleteDictRequest) (*emptypb.Empty, error) {
	job := model.Dict{}

	if err := job.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}
