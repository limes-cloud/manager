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

// Get 获取字典信息
func (r *Dict) Get(ctx kratosx.Context, in *v1.GetDictRequest) (*v1.Dict, error) {
	dict := model.Dict{}

	if err := dict.FindByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 进行数据转换
	reply := v1.Dict{}
	if err := util.Transform(dict, &reply); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	return &reply, nil
}

func (r *Dict) Page(ctx kratosx.Context, in *v1.PageDictRequest) (*v1.PageDictReply, error) {
	dict := model.Dict{}
	list, total, err := dict.Page(ctx, &types.PageOptions{
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

// Add 添加字典信息
func (r *Dict) Add(ctx kratosx.Context, in *v1.AddDictRequest) (*emptypb.Empty, error) {
	dict := model.Dict{}

	// 进行数据转换
	if err := util.Transform(in, &dict); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 创建字典
	if err := dict.Create(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Update 更新字典信息
func (r *Dict) Update(ctx kratosx.Context, in *v1.UpdateDictRequest) (*emptypb.Empty, error) {
	dict := model.Dict{}

	// 转换数据格式
	if err := util.Transform(in, &dict); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 更新字典信息
	if err := dict.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Delete 删除字典
func (r *Dict) Delete(ctx kratosx.Context, in *v1.DeleteDictRequest) (*emptypb.Empty, error) {
	dict := model.Dict{}

	if err := dict.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}
