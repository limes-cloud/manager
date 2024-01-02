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

type DictValue struct {
	conf *config.Config
}

func NewDictValue(conf *config.Config) *DictValue {
	return &DictValue{
		conf: conf,
	}
}

func (r *DictValue) Page(ctx kratosx.Context, in *v1.PageDictValueRequest) (*v1.PageDictValueReply, error) {
	job := model.DictValue{}
	list, total, err := job.Page(ctx, &types.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.Label != nil {
				db = db.Where("label like ?", "%"+*in.Label+"%")
			}
			if in.Value != nil {
				db = db.Where("value=?", *in.Value)
			}
			return db
		},
	})
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.PageDictValueReply{Total: total}
	// 进行数据转换
	if err = util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// Add 添加职位信息
func (r *DictValue) Add(ctx kratosx.Context, in *v1.AddDictValueRequest) (*emptypb.Empty, error) {
	job := model.DictValue{}

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
func (r *DictValue) Update(ctx kratosx.Context, in *v1.UpdateDictValueRequest) (*emptypb.Empty, error) {
	job := model.DictValue{}

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
func (r *DictValue) Delete(ctx kratosx.Context, in *v1.DeleteDictValueRequest) (*emptypb.Empty, error) {
	job := model.DictValue{}

	if err := job.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}
