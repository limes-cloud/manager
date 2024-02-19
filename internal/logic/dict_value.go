package logic

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"
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

func (r *DictValue) GetValue(ctx kratosx.Context, in *v1.GetDictValueRequest) (*v1.GetDictValueReply, error) {
	dict := model.Dict{}
	if err := dict.FindByKeyword(ctx, in.Keyword); err != nil {
		return nil, v1.NotFoundError()
	}

	dv := model.DictValue{}
	list, err := dv.AllByDictId(ctx, dict.ID)
	if err != nil {
		return nil, v1.DatabaseError()
	}

	reply := &v1.GetDictValueReply{Dict: make(map[string]string)}
	for _, value := range list {
		reply.Dict[value.Value] = value.Label
	}

	if err = util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return reply, nil
}

func (r *DictValue) Page(ctx kratosx.Context, in *v1.PageDictValueRequest) (*v1.PageDictValueReply, error) {
	dv := model.DictValue{}
	list, total, err := dv.Page(ctx, &types.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			db = db.Where("dict_id=?", in.DictId)
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

// Add 添加字典值信息
func (r *DictValue) Add(ctx kratosx.Context, in *v1.AddDictValueRequest) (*emptypb.Empty, error) {
	dv := model.DictValue{}

	// 进行数据转换
	if err := util.Transform(in, &dv); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 创建字典值
	if err := dv.Create(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Import 导入字典值信息
func (r *DictValue) Import(ctx kratosx.Context, in *v1.ImportDictValueRequest) (*v1.ImportDictValueReply, error) {
	count := 0
	for _, item := range in.List {
		dv := model.DictValue{
			DictID:      in.DictId,
			Label:       item.Label,
			Value:       item.Value,
			Weight:      proto.Uint32(0),
			Description: item.Description,
		}
		if dv.Create(ctx) == nil {
			count++
		}
	}
	return &v1.ImportDictValueReply{Count: uint32(count)}, nil
}

// Update 更新字典值信息
func (r *DictValue) Update(ctx kratosx.Context, in *v1.UpdateDictValueRequest) (*emptypb.Empty, error) {
	dv := model.DictValue{}

	// 转换数据格式
	if err := util.Transform(in, &dv); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 更新字典值信息
	if err := dv.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Delete 删除字典值
func (r *DictValue) Delete(ctx kratosx.Context, in *v1.DeleteDictValueRequest) (*emptypb.Empty, error) {
	dv := model.DictValue{}

	if err := dv.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}
