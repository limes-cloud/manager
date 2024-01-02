package dict

import (
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/internal/model"
)

var dict = &model.Dict{
	BaseModel: types.BaseModel{
		ID: 1,
	},
	Name:        "性别",
	Keyword:     "gender",
	Description: "性别选项",
}

var dictValue = &[]*model.DictValue{
	{
		DictID:      1,
		Label:       "男",
		Value:       "M",
		Status:      proto.Bool(true),
		Weight:      proto.Uint32(0),
		Extra:       "",
		Description: "男性",
	},
	{
		DictID:      1,
		Label:       "女",
		Value:       "F",
		Status:      proto.Bool(true),
		Weight:      proto.Uint32(0),
		Extra:       "",
		Description: "女性",
	},
	{
		DictID:      1,
		Label:       "未知",
		Value:       "U",
		Status:      proto.Bool(true),
		Weight:      proto.Uint32(0),
		Extra:       "",
		Description: "未知性别",
	},
}
